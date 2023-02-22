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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/juncachain/juncachain/accounts"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/consensus"
	"github.com/juncachain/juncachain/consensus/misc"
	"github.com/juncachain/juncachain/core/state"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/crypto"
	"github.com/juncachain/juncachain/ethclient"
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
	//nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	//nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

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

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

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

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errUnauthorizedVerifier is returned if a header is verified by a non-authorized entity.
	errUnauthorizedVerifier = errors.New("unauthorized verifier")

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

	endpoint string
	once     sync.Once
	client   *ethclient.Client

	HookCandidates          func(number *big.Int) (MasterNodes, error)
	HookRandomizeSigners    func(number *big.Int, masternodes []common.Address) ([]common.Address, error)
	HookPenalty             func(chain consensus.ChainHeaderReader, number uint64) ([]common.Address, error)
	HookReward              func(chain consensus.ChainHeaderReader, state *state.StateDB, header *types.Header) (map[string]interface{}, error)
	HookGetBlockSigners     func(chain consensus.ChainHeaderReader, stateBlock *state.StateDB, header *types.Header) (map[common.Address]int, error)
	HookBlockSign           func(signer common.Address, toSign, current *types.Header) error
	HookSetRandomizeSecret  func(signer common.Address, header *types.Header) error
	HookSetRandomizeOpening func(signer common.Address, header *types.Header) error
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
	return c.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (c *PoSV) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
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

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}
	// Checkpoint blocks need to enforce zero beneficiary
	checkpoint := (number % c.config.Epoch) == 0
	//if checkpoint && header.Coinbase != (common.Address{}) {
	//	return errInvalidCheckpointBeneficiary
	//}
	// Nonces must be 0x00..0 or 0xff..f, zeroes enforced on checkpoints
	//if !bytes.Equal(header.Nonce[:], nonceAuthVote) && !bytes.Equal(header.Nonce[:], nonceDropVote) {
	//	return errInvalidVote
	//}
	//if checkpoint && !bytes.Equal(header.Nonce[:], nonceDropVote) {
	//	return errInvalidCheckpointVote
	//}
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	if len(header.Verification) < extraSeal {
		return consensus.ErrMissVerification
	}
	// Ensure that the extra-data contains an epoch on checkpoint, but none otherwise
	epochBytes := len(header.Extra) - extraVanity - extraSeal
	if !checkpoint && epochBytes != 0 {
		return errExtraSigners
	}
	if checkpoint && epochBytes == 0 {
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
	return c.verifyCascadingFields(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *PoSV) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
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

	// If the block is a checkpoint block, verify the signer list
	if number%c.config.Epoch == 0 {
		var extra Extra
		if err := extra.FromBytes(header.Extra); err != nil {
			return err
		}
		if len(extra.Epoch.M1s) == 0 || len(extra.Epoch.M2s) == 0 {
			return errMismatchingCheckpointSigners
		}
		newEpoch, err := c.makeEpoch(chain, number)
		if err != nil {
			return err
		}
		if !bytes.Equal(newEpoch.ToBytes(), extra.Epoch.ToBytes()) {
			return fmt.Errorf("invalid epoch before fork: have %s, want %s", extra.Epoch.String(), newEpoch.String())
		}
	}
	// All basic checks passed, verify the seal and return
	return c.verifySeal(chain, header, parents)
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (c *PoSV) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *PoSV) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
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
	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), header.Verification)
	if err != nil {
		return err
	}
	var verifier common.Address
	copy(verifier[:], crypto.Keccak256(pubkey[1:])[12:])
	if verifier == common.ZeroAddress {
		return errors.New("nil verifier")
	}

	// when sync blocks,cannot get epoch header sometimes
	epoch, _ := c.getEpoch(chain, c.LastEpoch(number))
	if epoch != nil {
		if !epoch.IsM1(signer) {
			return errUnauthorizedSigner
		}
		if !epoch.IsM1(verifier) {
			return errUnauthorizedVerifier
		}
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *PoSV) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Coinbase = common.Address{}
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()

	epoch, err := c.getEpoch(chain, c.LastEpoch(number))
	if err != nil {
		return err
	}
	c.lock.RLock()

	// Copy signer protected by mutex to avoid race condition
	signer := c.signer
	c.lock.RUnlock()

	//if number%c.config.Epoch != 0 {
	header.Coinbase = epoch.M1(c.config.Epoch, number)

	// Set the correct difficulty
	if signer == header.Coinbase {
		header.Difficulty = new(big.Int).Set(diffInTurn)
	} else {
		header.Difficulty = new(big.Int).Set(diffNoTurn)
	}
	//}

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	if number%c.config.Epoch == 0 {
		// On Junca restart,coinbase is ZeroAddress,node rpc not start yet
		if header.Coinbase == common.ZeroAddress {
			return nil
		}
		newEpoch, err := c.makeEpoch(chain, number)
		if err != nil {
			return err
		}
		header.Extra = append(header.Extra, newEpoch.ToBytes()...)
	}
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

	if c.HookSetRandomizeSecret != nil && epoch.IsM1(signer) {
		if err := c.HookSetRandomizeSecret(signer, header); err != nil {
			log.Error("[PoSV]HookSetRandomizeSecret", "err", err.Error())
		}
	}

	if c.HookSetRandomizeOpening != nil && epoch.IsM1(signer) {
		if err := c.HookSetRandomizeOpening(signer, header); err != nil {
			log.Error("[PoSV]HookSetRandomizeOpening", "err", err.Error())
		}
	}

	if c.HookBlockSign != nil && epoch.IsM2(signer) {
		if number > common.EpochBlockSignWiggle {
			toSignBlockNumber := number - common.EpochBlockSignWiggle
			toSignBlock := chain.GetHeaderByNumber(toSignBlockNumber)
			if toSignBlock != nil {
				if m2s := epoch.M2(c.config.Epoch, toSignBlockNumber); len(m2s) > 0 {
					for _, m2 := range m2s {
						if m2 == signer {
							log.Info("[PoSV]It's turn to validator block", "number", toSignBlockNumber, "validator", signer)
							if err := c.HookBlockSign(signer, toSignBlock, header); err != nil {
								log.Error("[PoSV]HookBlockSign", "err", err.Error())
							}
						}
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
	if c.HookReward != nil && header.Number.Uint64()%c.config.Epoch == 0 {
		log.Info("[PoSV]HookReward", "number", header.Number)
		rewards, err := c.HookReward(chain, state, header)
		if err != nil {
			log.Crit("[PoSV]HookReward", "err", err)
		}

		data, err := json.MarshalIndent(rewards, "", "  ")
		if err == nil {
			err = ioutil.WriteFile(filepath.Join(c.config.InstanceDir, header.Number.String()+"_epoch.json"), data, 0644)
		}
		if err != nil {
			log.Error("Error when save reward info ", "number", header.Number, "hash", header.Hash().Hex(), "err", err)
		}
	}
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
func (c *PoSV) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
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
	epoch, err := c.getEpoch(chain, c.LastEpoch(number))
	if err != nil {
		return err
	}
	if !epoch.IsM1(signer) {
		return errors.Errorf("Etherbase %s is not M1 at %d", signer.Hex(), number)
	}

	// If we're amongst the recent signers, wait for the next block
	// only check recent signers if there are more than one signer.
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

// LastEpoch return the epoch checkpoint number
func (c *PoSV) LastEpoch(number uint64) uint64 {
	if number <= c.config.Epoch {
		return 0
	} else if number%c.config.Epoch == 0 {
		return number - c.config.Epoch
	}
	return number - number%c.config.Epoch
}

// GetEpoch return the epoch data at checkpoint header,the param epoch must be checkpoint number
func (c *PoSV) GetEpoch(chain consensus.ChainHeaderReader, epoch uint64) (*Epoch, error) {
	return c.getEpoch(chain, epoch)
}

func (c *PoSV) getEpoch(chain consensus.ChainHeaderReader, epoch uint64) (*Epoch, error) {
	v, ok := c.epochs.Get(epoch)
	if ok {
		return v.(*Epoch), nil
	}

	header := chain.GetHeaderByNumber(epoch)
	if header != nil {
		extra := new(Extra)
		if err := extra.FromBytes(header.Extra); err != nil {
			return nil, err
		}
		c.epochs.Add(epoch, &extra.Epoch)
		return &extra.Epoch, nil
	}
	return nil, errors.Errorf("Not found epoch extra data:%d", epoch)
}

// makeEpoch generate epoch data and return it, number must be checkpoint
func (c *PoSV) makeEpoch(chain consensus.ChainHeaderReader, number uint64) (*Epoch, error) {
	if number%c.config.Epoch != 0 {
		return nil, errors.New("Not checkpoint")
	}
	ms, err := c.HookCandidates(new(big.Int).SetUint64(number - 1))
	if err != nil {
		return nil, err
	}

	var masterNodes []common.Address

	epoch := new(Epoch)
	epoch.Checkpoint = number
	epoch.Reward = c.CalcReward(number)
	for i, v := range ms {
		if v.Stake.Cmp(c.config.MinStaked) < 0 {
			continue
		}
		if i < common.MaxMasternodes {
			masterNodes = append(masterNodes, v.Address)
		}
		if i < common.MaxSigners {
			epoch.M1s = append(epoch.M1s, MasterNode{Address: v.Address, Stake: v.Stake})
		}
	}

	m2s, err := c.HookRandomizeSigners(new(big.Int).SetInt64(int64(number-1)), masterNodes)
	if err != nil {
		return nil, err
	}
	epoch.M2s = m2s

	newEpoch, err := c.penalty(chain, number, epoch)
	if err != nil {
		return nil, err
	}

	c.epochs.Add(newEpoch, newEpoch)
	return newEpoch, nil
}

func (c *PoSV) Config() *params.PoSVConfig {
	return c.config
}

func (c *PoSV) SetIPCEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *PoSV) initClient() {
	c.once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				log.Crit("PoSV init ethclient timeout")
			default:
				cli, err := ethclient.Dial(c.endpoint)
				if err != nil {
					log.Warn("PoSV ethclient.Dial", "err", err)
					continue
				}
				c.client = cli
				return
			}
		}
	})
}

func (c *PoSV) lastSignedBlock(chain consensus.ChainHeaderReader, number uint64, signer common.Address) *types.Header {
	for i := 0; i < common.MaxSigners; i++ {
		header := chain.GetHeaderByNumber(number)
		if header == nil {
			break
		}
		address, err := c.Author(header)
		if err != nil {
			log.Crit("Author", "err", err, "number", header.Number.Uint64())
		}
		if address == signer {
			c.recents.Add(signer, header.Number.Uint64())
			return header
		}
		number = number - 1
		if number == 0 {
			break
		}
	}
	return nil
}

// penalty penaltyHook , update epoch data and return it
func (c *PoSV) penalty(chain consensus.ChainHeaderReader, number uint64, epoch *Epoch) (*Epoch, error) {
	if c.HookPenalty != nil {
		log.Info("[PoSV]HookPenalty", "number", number)
		penaltys, err := c.HookPenalty(chain, number)
		if err != nil {
			return epoch, err
		}
		if len(penaltys) > 0 {
			epoch.Penalties = make([]common.Address, len(penaltys))
			for i, v := range penaltys {
				epoch.Penalties[i] = v
				log.Info("[PoSV]HookPenalty", "penalized", v.Hex())
			}
		}
		penalized := func(address common.Address) bool {
			for _, v := range penaltys {
				if v == address {
					return true
				}
			}
			return false
		}

		if len(penaltys) > 0 {
			var m1s MasterNodes
			var m2s []common.Address
			for _, v := range epoch.M1s {
				if !penalized(v.Address) {
					m1s = append(m1s, v)
				}
			}
			for _, v := range epoch.M2s {
				if !penalized(v) {
					m2s = append(m2s, v)
				}
			}
			if len(m1s) > 0 && len(m2s) > 0 {
				epoch.M1s = m1s
				epoch.M2s = m2s
			}
		}
	}
	return epoch, nil
}

// CalcReward return reward of current epoch
func (c *PoSV) CalcReward(number uint64) *big.Int {
	epochPerYear := 3600 * 24 * 365 / c.config.Period / c.config.Epoch
	currentEpoch := (number - 1) / c.config.Epoch
	var minted = new(big.Int)
	var mintThisEpoch = c.config.Reward
	for i := uint64(0); i < currentEpoch/epochPerYear; i++ {
		mintedThisYear := new(big.Int).Mul(mintThisEpoch, big.NewInt(int64(epochPerYear)))
		minted = minted.Add(minted, mintedThisYear)

		mintThisEpoch = new(big.Int).Mul(mintThisEpoch, big.NewInt(3))
		mintThisEpoch = new(big.Int).Div(mintThisEpoch, big.NewInt(4))
	}
	mod := currentEpoch % epochPerYear
	mintedThisYear := new(big.Int).Mul(mintThisEpoch, big.NewInt(int64(mod-1)))
	minted = minted.Add(minted, mintedThisYear)
	if minted.Cmp(c.config.TotalReward) >= 0 {
		return new(big.Int)
	}
	if new(big.Int).Sub(c.config.TotalReward, minted).Cmp(mintThisEpoch) <= 0 {
		return new(big.Int).Sub(c.config.TotalReward, minted)
	}
	return mintThisEpoch
}
