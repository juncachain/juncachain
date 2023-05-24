// Package posv Copyright (c) 2023 Juncachain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package posv implements the Proof-of-Stake Voting consensus engine.
package posv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/juncachain/juncachain/core"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/juncachain/juncachain/accounts"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/consensus"
	"github.com/juncachain/juncachain/consensus/misc"
	"github.com/juncachain/juncachain/core/state"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/crypto"
	"github.com/juncachain/juncachain/ethdb"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/params"
	"github.com/juncachain/juncachain/rlp"
	"github.com/juncachain/juncachain/rpc"
	"github.com/juncachain/juncachain/trie"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

const (
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory
	inmemoryEpochs     = 128  // Number of recent Epoch to keep in memory

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

const (
	epochLength = uint64(900)            // Default number of blocks after which to checkpoint and reset the pending votes
	extraVanity = 32                     // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal
)

// PoSV Proof-of-Stake Voting protocol constants.
var (
	nonceCheckpoint = hexutil.MustDecode("0xffffffffffffffff")
	nonceDefault    = hexutil.MustDecode("0x0000000000000000")

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidNonce is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidNonce = errors.New("nonce not 0x00..0 or 0xff..f")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errMismatchingCheckpointSigners is returned if a checkpoint block contains a
	// list of signers different than the one the local node calculated.
	errMismatchingCheckpointSigners = errors.New("mismatching signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block neither 1 or 2.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// errWrongDifficulty is returned if the difficulty of a block doesn't match the
	// turn of the signer.
	errWrongDifficulty = errors.New("wrong difficulty")

	// errInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	errInvalidTimestamp = errors.New("invalid timestamp")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")
)

// SignerFn hashes and signs the data to be signed by a backing account.
type SignerFn func(signer accounts.Account, mimeType string, message []byte) ([]byte, error)

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// PoSV is the proof-of-authority consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type PoSV struct {
	config *params.PoSVConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining
	epochs     *lru.ARCCache // Epochs of recent blocks to speed up mining

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields

	chainID *big.Int
	txPool  *core.TxPool
	acm     *accounts.Manager
}

// New creates a PoSV proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.PoSVConfig, db ethdb.Database) *PoSV {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)
	epochs, _ := lru.NewARC(inmemoryEpochs)

	return &PoSV{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		epochs:     epochs,
	}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (c *PoSV) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, c.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (c *PoSV) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	log.Info("----------VerifyHeader", "number", header.Number.Uint64())
	return c.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (c *PoSV) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	log.Info("----------VerifyHeaders", "begin number", headers[0].Number.Uint64(), "end number", headers[len(headers)-1].Number.Uint64())
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *PoSV) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()
	isCheckpoint := (number % c.config.Epoch) == 0

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}
	// Coinbase must be M1
	// getEpoch may be return nil and error when synchronizing blocks
	// The next time you sync to number, you can getEpoch correctly
	epoch, err := c.getEpoch(chain, c.LastCheckpoint(number))
	if epoch == nil || err != nil {
		return errors.Errorf("Got a nil epoch at %d, err %v", number, err)
	}
	if epoch != nil && !epoch.IsM1(header.Coinbase) {
		return errInvalidCheckpointBeneficiary
	}
	// Nonces must be 0x00..0 , zeroes enforced on checkpoints
	if !isCheckpoint && !bytes.Equal(header.Nonce[:], nonceDefault) {
		return errInvalidNonce
	}
	if isCheckpoint && !bytes.Equal(header.Nonce[:], nonceCheckpoint) {
		return errInvalidNonce
	}
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains an epoch on checkpoint, but none otherwise
	epochBytes := len(header.Extra) - extraVanity - extraSeal
	if !isCheckpoint && epochBytes != 0 {
		return errExtraSigners
	}
	if isCheckpoint && epochBytes == 0 {
		return errInvalidCheckpointSigners
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}
	// Verify that the gas limit is <= 2^63-1
	if header.GasLimit > params.MaxGasLimit {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, params.MaxGasLimit)
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields(chain, header, parents, epoch)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *PoSV) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, epoch *Epoch) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to its parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+c.config.Period > header.Time {
		return errInvalidTimestamp
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}
	if !chain.Config().IsLondon(header.Number) {
		// Verify BaseFee not present before EIP-1559 fork.
		if header.BaseFee != nil {
			return fmt.Errorf("invalid baseFee before fork: have %d, want <nil>", header.BaseFee)
		}
		if err := misc.VerifyGaslimit(parent.GasLimit, header.GasLimit); err != nil {
			return err
		}
	} else if err := misc.VerifyEip1559Header(chain.Config(), parent, header); err != nil {
		// Verify the header's EIP-1559 attributes.
		return err
	}

	// If the block is a checkpoint block, verify the new epoch data
	if number%c.config.Epoch == 0 {
		var extra Extra
		if err := extra.FromBytes(header.Extra); err != nil {
			return err
		}
		if len(extra.Epoch.M1s) == 0 || len(extra.Epoch.M2s) == 0 {
			return errMismatchingCheckpointSigners
		}
		// makeEpoch may be return nil and error when synchronizing blocks
		//if newEpoch, err := c.makeEpoch(chain, number); newEpoch != nil &&
		//	!bytes.Equal(newEpoch.ToBytes(), extra.Epoch.ToBytes()) {
		//	return fmt.Errorf("invalid epoch before fork: have %s, want %s", extra.Epoch.String(), newEpoch.String())
		//} else if err != nil {
		//	// If return an error,the synchronizer first writes the block before number into the blockchain
		//	// The next time you sync to number, you can makeEpoch correctly
		//	return errors.Errorf("Make a nil epoch at %d err %+v", number, err)
		//}
	}
	// All basic checks passed, verify the seal and return
	return c.verifySeal(chain, header, parents, epoch)
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (c *PoSV) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	log.Info("----------VerifyUncles", "number", block.Number().Uint64())
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *PoSV) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, epoch *Epoch) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, c.signatures)
	if err != nil {
		return err
	}
	if signer == common.ZeroAddress {
		return errors.New("nil signer")
	}

	// Signer must be M1
	if epoch != nil && !epoch.IsM1(signer) {
		return errUnauthorizedSigner
	}

	// It's turn to signer seal
	if epoch != nil && epoch.M1(c.config.Epoch, number) == signer {
		return nil
	}

	// Signer cannot sign two consecutive blocks unless
	// it's your turn or the block is the first in an epoch
	if number%c.config.Epoch != 1 {
		if parent := chain.GetHeader(header.ParentHash, number-1); parent != nil {
			if lastSigner, err := ecrecover(parent, c.signatures); err != nil {
				return err
			} else if lastSigner == signer {
				return errRecentlySigned
			}
		}
	}

	// Signer Prevent Signer from preempting blocks
	if parent := chain.GetHeader(header.ParentHash, number-1); parent != nil && epoch != nil {
		nextSealBlock := epoch.M1NextTurn(c.config.Epoch, number, signer)
		nextSealTime := parent.Time + (nextSealBlock-number)*c.config.Period
		if header.Time < nextSealTime {
			return errors.New("signed timeout not reached")
		}
	}

	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *PoSV) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	log.Info("----------Prepare", "number", header.Number.Uint64())
	// If the block isn't a isCheckpoint, cast a random vote (good enough for now)
	header.Coinbase = common.Address{}
	header.Nonce = types.BlockNonce{}
	number := header.Number.Uint64()
	isCheckpoint := number%c.config.Epoch == 0
	if isCheckpoint {
		copy(header.Nonce[:], nonceCheckpoint)
	}

	epoch, err := c.getEpoch(chain, c.LastCheckpoint(number))
	if err != nil {
		return err
	}

	c.lock.RLock()

	// Copy signer protected by mutex to avoid race condition
	signer := c.signer
	c.lock.RUnlock()

	// Set the coinbase in order
	header.Coinbase = epoch.M1(c.config.Epoch, number)

	// Set the correct difficulty
	if signer == header.Coinbase {
		header.Difficulty = new(big.Int).Set(diffInTurn)
	} else {
		header.Difficulty = new(big.Int).Set(diffNoTurn)
	}

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	//if isCheckpoint {
	//	// On Junca restart,coinbase is ZeroAddress,node rpc not start yet
	//	if header.Coinbase == common.ZeroAddress {
	//		return nil
	//	}
	//	nextEpoch, err := c.makeEpoch(chain, number)
	//	if err != nil {
	//		return err
	//	}
	//	header.Extra = append(header.Extra, nextEpoch.ToBytes()...)
	//}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = parent.Time + c.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// Set secret
	modNumber := header.Number.Uint64() % c.Config().Epoch
	if c.Config().Epoch-modNumber == common.EpochBlockSecret &&
		c.HookSetRandomizeSecret != nil &&
		epoch.IsM2(signer) {
		log.Info("[PoSV]Is number to set secret", "number", header.Number.Uint64(), "signer", signer)
		if err := c.HookSetRandomizeSecret(signer, header); err != nil {
			log.Error("[PoSV]HookSetRandomizeSecret", "err", err.Error())
		}
		return nil
	}

	// Set opening
	if c.Config().Epoch-modNumber == common.EpochBlockOpening &&
		c.HookSetRandomizeOpening != nil &&
		epoch.IsM2(signer) {
		log.Info("[PoSV]Is number to set opening", "number", header.Number.Uint64(), "signer", signer)
		if err := c.HookSetRandomizeOpening(signer, header); err != nil {
			log.Error("[PoSV]HookSetRandomizeOpening", "err", err.Error())
		}
	}

	// Sign parent block
	if c.HookBlockSign != nil && epoch.IsM2(signer) {
		if m2s := epoch.M2(c.config.Epoch, parent.Number.Uint64()); len(m2s) > 0 {
			for _, m2 := range m2s {
				if m2 == signer {
					log.Info("[PoSV]Is turn to validator block", "number", parent.Number.Uint64(), "validator", signer)
					if err := c.HookBlockSign(signer, parent, header); err != nil {
						log.Error("[PoSV]HookBlockSign", "err", err.Error())
					}
				}
			}
		}
	}

	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given.
func (c *PoSV) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	log.Info("----------Finalize", "number", header.Number.Uint64(), "txs", len(txs))

	if c.HookReward != nil && header.Number.Uint64()%c.config.Epoch == 0 {
		number := header.Number.Uint64()
		{
			// On Junca restart,coinbase is ZeroAddress,node rpc not start yet
			if header.Coinbase == common.ZeroAddress {
				return
			}
			nextEpoch, err := c.makeEpoch(chain, state, number)
			if err != nil {
				return
			}
			epochBzLen := len(nextEpoch.ToBytes())
			header.Extra = append(header.Extra, make([]byte, epochBzLen)...)
			copy(header.Extra[extraVanity:], nextEpoch.ToBytes())
			log.Info("[PoSV]HookReward", "number", header.Number)
			rewards, err := c.HookReward(chain, state, header, nextEpoch)
			if err != nil {
				log.Crit("[PoSV]HookReward", "err", err)
			}

			data, err := json.MarshalIndent(rewards, "", "  ")
			if err == nil {
				err = os.WriteFile(filepath.Join(c.config.InstanceDir, header.Number.String()+"_epoch.json"), data, 0644)
			}
			if err != nil {
				log.Error("Error when save reward info ", "number", header.Number, "hash", header.Hash().Hex(), "err", err)
			}
		}

	}
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
func (c *PoSV) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	log.Info("----------FinalizeAndAssemble", "number", header.Number.Uint64(), "txs", len(txs))
	// Finalize block
	c.Finalize(chain, header, state, txs, uncles)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (c *PoSV) Authorize(signer common.Address, signFn SignerFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *PoSV) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Info("----------Seal", "number", block.Number().Uint64(), "txs", len(block.Transactions()))
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if c.config.Period == 0 && len(block.Transactions()) == 0 {
		return errors.New("sealing paused while waiting for transactions")
	}
	// Don't hold the signer fields for the entire sealing procedure
	c.lock.RLock()
	signer, signFn := c.signer, c.signFn
	c.lock.RUnlock()

	// Get last epoch-header extra data
	epoch, err := c.getEpoch(chain, c.LastCheckpoint(number))
	if err != nil {
		return err
	}
	if !epoch.IsM1(signer) {
		return errors.Errorf("Etherbase %s is not M1 at %d", signer.Hex(), number)
	}

	// If we're amongst the recent signers, wait for the next block
	// only check recent signers if there are more than one signer.
	//if number%c.config.Epoch != 1 && epoch.M1(c.config.Epoch, number) != signer {
	if number%c.config.Epoch != 1 && epoch.M1(c.config.Epoch, number) != signer {
		parent := chain.GetHeaderByNumber(number - 1)
		if parent != nil && parent.Number.Uint64() > 0 && epoch.M1Length() > 1 {
			if pSigner, err := c.Author(parent); err != nil {
				return err
			} else if pSigner == signer {
				log.Info("[PoSV]Signed recently, must wait for others", "m1 length", epoch.M1Length())
				return nil
			}
		}
	}

	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	nextNumber := epoch.M1NextTurn(c.config.Epoch, number, signer)
	if number == 1 && nextNumber != number {
		log.Info("[PoSV]First block must be sealed in order")
		return nil
	}
	if nextNumber > number {
		delay = delay + time.Duration((nextNumber-number)*c.config.Period)*time.Second + wiggleTime
	}

	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypePoSV, CliqueRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)

	// Wait until sealing is terminated or delay timeout.
	log.Info("[PoSV]Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have:
// * DIFF_NOTURN(2) if BLOCK_NUMBER % SIGNER_COUNT != SIGNER_INDEX
// * DIFF_INTURN(1) if BLOCK_NUMBER % SIGNER_COUNT == SIGNER_INDEX
func (c *PoSV) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	log.Info("----------CalcDifficulty", "number", parent.Number.Uint64())
	c.lock.RLock()
	signer := c.signer
	c.lock.RUnlock()

	if signer == chain.CurrentHeader().Coinbase {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

// SealHash returns the hash of a block prior to it being sealed.
func (c *PoSV) SealHash(header *types.Header) common.Hash {
	log.Info("----------CalcDifficulty", "number", header.Number.Uint64())
	return SealHash(header)
}

// Close implements consensus.Engine. It's a noop for posv as there are no background threads.
func (c *PoSV) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (c *PoSV) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{{
		Namespace: "posv",
		Service:   &API{chain: chain, posv: c},
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.(crypto.KeccakState).Read(hash[:])
	return hash
}

// CliqueRLP returns the rlp bytes which needs to be signed for the proof-of-authority
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func CliqueRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	enc := []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	}
	if header.BaseFee != nil {
		enc = append(enc, header.BaseFee)
	}
	if err := rlp.Encode(w, enc); err != nil {
		panic("can't encode: " + err.Error())
	}
}

// LastCheckpoint return the epoch checkpoint number
func (c *PoSV) LastCheckpoint(currentNumber uint64) uint64 {
	if currentNumber <= c.config.Epoch {
		return 0
	} else if currentNumber%c.config.Epoch == 0 {
		return currentNumber - c.config.Epoch
	}
	return currentNumber - currentNumber%c.config.Epoch
}

// GetEpoch return the epoch data at checkpoint header,the param checkpoint must be checkpoint number
func (c *PoSV) GetEpoch(chain consensus.ChainHeaderReader, checkpoint uint64) (*Epoch, error) {
	return c.getEpoch(chain, checkpoint)
}

func (c *PoSV) getEpoch(chain consensus.ChainHeaderReader, checkpoint uint64) (*Epoch, error) {
	v, ok := c.epochs.Get(checkpoint)
	if ok {
		return v.(*Epoch), nil
	}

	header := chain.GetHeaderByNumber(checkpoint)
	if header != nil {
		extra := new(Extra)
		if err := extra.FromBytes(header.Extra); err != nil {
			return nil, err
		}
		c.epochs.Add(checkpoint, &extra.Epoch)
		return &extra.Epoch, nil
	}
	return nil, errors.Errorf("Not found checkpoint extra data:%d", checkpoint)
}

// makeEpoch generate epoch data and return it, number must be checkpoint
func (c *PoSV) makeEpoch(chain consensus.ChainHeaderReader, state *state.StateDB, number uint64) (*Epoch, error) {
	// Generate the next epoch
	epoch := new(Epoch)
	epoch.Checkpoint = number
	epoch.Reward = CalcReward(c.config, number)

	// Must shift block. number may be not sealed
	candidates := c.HookGetCandidates(state)

	// Get penalty list,penalties is sorted. number may be not sealed
	penalties, err := c.HookPenalty(chain, number-1)
	if err != nil {
		log.Error("[PoSV]HookPenalty", "err", err)
		return nil, err
	}
	epoch.Penalties = make([]common.Address, len(penalties))
	for i, v := range penalties {
		epoch.Penalties[i] = v
		log.Info("[PoSV]HookPenalty", "penalized", v.Hex())
	}

	penalized := func(address common.Address) bool {
		for _, v := range penalties {
			if v == address {
				return true
			}
		}
		return false
	}
	// If all candidates are penalied,penaly nothing
	if len(candidates)-len(penalties) == 0 {
		penalized = func(address common.Address) bool {
			return false
		}
	}

	// Max 150 masterNodes
	var masterNodes []common.Address
	for i, candidate := range candidates {
		if penalized(candidate.Address) {
			continue
		}
		if i < common.MaxMasternodes {
			masterNodes = append(masterNodes, candidate.Address)
		}
		if i < common.MaxSigners {
			epoch.M1s = append(epoch.M1s, MasterNode{Address: candidate.Address, Stake: candidate.Stake})
		}
	}

	// Randomize masterNodes. number may be not sealed
	m2s, err := c.HookRandomizeSigners(state, masterNodes)
	if err != nil {
		log.Error("[PoSV]HookRandomizeSigners", "err", err)
		return nil, err
	}
	epoch.M2s = m2s

	c.epochs.Add(epoch, epoch)
	return epoch, nil
}

func (c *PoSV) Config() *params.PoSVConfig {
	return c.config
}

func (c *PoSV) SetChainID(chainID *big.Int) {
	c.chainID = chainID
}

func (c *PoSV) SetTxPool(txPool *core.TxPool) {
	c.txPool = txPool
}

func (c *PoSV) SetAccountManager(acm *accounts.Manager) {
	c.acm = acm
}

// CalcReward return reward of current epoch
func CalcReward(cfg *params.PoSVConfig, checkpoint uint64) *big.Int {
	epochPerYear := 3600 * 24 * 365 / cfg.Period / cfg.Epoch
	currentEpoch := checkpoint / cfg.Epoch
	var minted = new(big.Int)
	var mintThisEpoch = cfg.Reward
	for i := uint64(0); i < currentEpoch/epochPerYear; i++ {
		mintedThisYear := new(big.Int).Mul(mintThisEpoch, big.NewInt(int64(epochPerYear)))
		minted = minted.Add(minted, mintedThisYear)

		// reduction 10% per year
		mintThisEpoch = new(big.Int).Mul(mintThisEpoch, big.NewInt(90))
		mintThisEpoch = new(big.Int).Div(mintThisEpoch, big.NewInt(100))
	}
	mod := currentEpoch % epochPerYear
	mintedThisYear := new(big.Int).Mul(mintThisEpoch, big.NewInt(int64(mod-1)))
	minted = minted.Add(minted, mintedThisYear)
	if minted.Cmp(cfg.TotalReward) >= 0 {
		return new(big.Int)
	}
	if new(big.Int).Sub(cfg.TotalReward, minted).Cmp(mintThisEpoch) <= 0 {
		return new(big.Int).Sub(cfg.TotalReward, minted)
	}
	return mintThisEpoch
}
