// Copyright (c) 2018 Juncachain
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

package contracts

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/juncachain/juncachain/accounts"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/consensus/posv"
	"github.com/juncachain/juncachain/contracts/blocksigner/contract"
	randomizeContract "github.com/juncachain/juncachain/contracts/randomize/contract"
	contractValidator "github.com/juncachain/juncachain/contracts/validator/contract"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/core/state"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/ethclient"
	"github.com/juncachain/juncachain/ethdb"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/params"
)

var (
	randomizeKeyName = []byte("randomizeKey")
	txMutex          sync.RWMutex
)

type rewardLog struct {
	Sign   uint64   `json:"sign"`
	Reward *big.Int `json:"reward"`
}

func etherbaseWallet(manager *accounts.Manager, etherbase common.Address) accounts.Wallet {
	etherbaseAccount := accounts.Account{
		Address: etherbase,
		URL:     accounts.URL{},
	}
	if w, err := manager.Find(etherbaseAccount); err == nil && w != nil {
		return w
	}
	return nil
}

// CreateTransactionSign Send tx sign for block number to smart contract blockSigner.
func CreateTransactionSign(chainConfig *params.ChainConfig, pool *core.TxPool, manager *accounts.Manager, header *types.Header, chainDb ethdb.Database, eb common.Address) error {
	txMutex.Lock()
	defer txMutex.Unlock()

	wallet := etherbaseWallet(manager, eb)
	if wallet == nil {
		return fmt.Errorf("Cannot find etherbase wallet:%s", eb.Hex())
	}

	nonce := pool.Nonce(eb)
	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = new(big.Int).SetUint64(params.InitialBaseFee)
	}
	gasPrice := new(big.Int).Add(baseFee, big.NewInt(1))
	tx := BuildTxSignBlockSigner(header.Number, header.Hash(), nonce, gasPrice, common.HexToAddress(common.BlockSignerSMC))
	txSigned, err := wallet.SignTx(accounts.Account{Address: eb}, tx, chainConfig.ChainID)
	if err != nil {
		log.Error("Fail to create tx sign", "error", err)
		return err
	}
	// Add tx signed to local tx pool.
	err = pool.AddLocal(txSigned)
	if err != nil {
		log.Error("Fail to add tx sign to local pool.", "error", err, "number", header.Number, "hash", txSigned.Hash().Hex(), "from", eb, "nonce", nonce)
		return err
	}
	return nil
}

// CreateTransactionSetSecret Send tx set secret to smart contract randomize.
func CreateTransactionSetSecret(chainConfig *params.ChainConfig, pool *core.TxPool, manager *accounts.Manager, header *types.Header, chainDb ethdb.Database, eb common.Address) error {
	txMutex.Lock()
	defer txMutex.Unlock()

	wallet := etherbaseWallet(manager, eb)
	if wallet == nil {
		return fmt.Errorf("Cannot find etherbase wallet:%s", eb.Hex())
	}

	// Generate random private key and save into chaindb.
	exist, _ := chainDb.Has(randomizeKeyName)
	if exist {
		return nil
	}

	// Only process when private key empty in state db.
	// Save randomize key into state db.
	randomizeKeyValue := RandStringByte(32)

	nonce := pool.Nonce(eb)
	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = new(big.Int).SetUint64(params.InitialBaseFee)
	}
	gasPrice := new(big.Int).Add(baseFee, big.NewInt(1))
	tx, err := BuildTxSecretRandomize(nonce, gasPrice, common.HexToAddress(common.RandomizeSMC), chainConfig.Posv.Epoch, randomizeKeyValue)
	if err != nil {
		log.Error("Fail to get tx secret for randomize", "error", err)
		return err
	}
	txSigned, err := wallet.SignTx(accounts.Account{Address: eb}, tx, chainConfig.ChainID)
	if err != nil {
		log.Error("Fail to create tx secret", "error", err)
		return err
	}
	// Add tx signed to local tx pool.
	err = pool.AddLocal(txSigned)
	if err != nil {
		log.Error("Fail to add tx secret to local pool.", "error", err, "number", header.Number, "hash", txSigned.Hash().Hex(), "from", eb, "nonce", nonce)
		return err
	}

	// Put randomize key into chainDb.
	chainDb.Put(randomizeKeyName, randomizeKeyValue)
	return nil
}

// CreateTransactionSetOpening Send tx set opening to smart contract randomize.
func CreateTransactionSetOpening(chainConfig *params.ChainConfig, pool *core.TxPool, manager *accounts.Manager, header *types.Header, chainDb ethdb.Database, eb common.Address) error {
	txMutex.Lock()
	defer txMutex.Unlock()

	wallet := etherbaseWallet(manager, eb)
	if wallet == nil {
		return fmt.Errorf("Cannot find etherbase wallet:%s", eb.Hex())
	}

	// Get random private key from chaindb.
	randomizeKeyValue, err := chainDb.Get(randomizeKeyName)
	if err != nil {
		log.Error("Fail to get randomize key from state db.", "error", err)
		return err
	}

	nonce := pool.Nonce(eb)
	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = new(big.Int).SetUint64(params.InitialBaseFee)
	}
	gasPrice := new(big.Int).Add(baseFee, big.NewInt(1))
	tx, err := BuildTxOpeningRandomize(nonce, gasPrice, common.HexToAddress(common.RandomizeSMC), randomizeKeyValue)
	if err != nil {
		log.Error("Fail to get tx opening for randomize", "error", err)
		return err
	}
	txSigned, err := wallet.SignTx(accounts.Account{Address: eb}, tx, chainConfig.ChainID)
	if err != nil {
		log.Error("Fail to create tx opening", "error", err)
		return err
	}
	// Add tx signed to local tx pool.
	err = pool.AddLocal(txSigned)
	if err != nil {
		log.Error("Fail to add tx opening to local pool.", "error", err, "number", header.Number, "hash", txSigned.Hash().Hex(), "from", eb, "nonce", nonce)
		return err
	}

	// Clear randomize key in state db.
	chainDb.Delete(randomizeKeyName)
	return nil
}

// BuildTxSignBlockSigner Build tx to sign into BlockSigner.
func BuildTxSignBlockSigner(blockNumber *big.Int, blockHash common.Hash, nonce uint64, gasPrice *big.Int, blockSigner common.Address) *types.Transaction {
	data := common.Hex2Bytes(common.HexSignMethod)
	inputData := append(data, common.LeftPadBytes(blockNumber.Bytes(), 32)...)
	inputData = append(inputData, common.LeftPadBytes(blockHash.Bytes(), 32)...)

	tx := types.NewTransaction(nonce, blockSigner, big.NewInt(0), 200000, gasPrice, inputData)
	return tx
}

// BuildTxSecretRandomize Build tx to set secret into Randomize.
func BuildTxSecretRandomize(nonce uint64, gasPrice *big.Int, randomizeAddr common.Address, epocNumber uint64, randomizeKey []byte) (*types.Transaction, error) {
	data := common.Hex2Bytes(common.HexSetSecret)
	rand.Seed(time.Now().UnixNano())
	secretNumb := rand.Intn(int(epocNumber))

	// Append randomize suffix in -1, 0, 1.
	secrets := []int64{int64(secretNumb)}
	sizeOfArray := int64(32)

	// Build extra data for tx with first position is size of array byte and second position are length of array byte.
	arrSizeOfSecrets := common.LeftPadBytes(new(big.Int).SetInt64(sizeOfArray).Bytes(), 32)
	arrLengthOfSecrets := common.LeftPadBytes(new(big.Int).SetInt64(int64(len(secrets))).Bytes(), 32)
	inputData := append(data, arrSizeOfSecrets...)
	inputData = append(inputData, arrLengthOfSecrets...)
	for _, secret := range secrets {
		encryptSecret := Encrypt(randomizeKey, new(big.Int).SetInt64(secret).String())
		inputData = append(inputData, common.LeftPadBytes([]byte(encryptSecret), int(sizeOfArray))...)
	}
	tx := types.NewTransaction(nonce, randomizeAddr, big.NewInt(0), 200000, gasPrice, inputData)

	return tx, nil
}

// BuildTxOpeningRandomize Build tx to set opening into Randomize.
func BuildTxOpeningRandomize(nonce uint64, gasPrice *big.Int, randomizeAddr common.Address, randomizeKey []byte) (*types.Transaction, error) {
	data := common.Hex2Bytes(common.HexSetOpening)
	inputData := append(data, randomizeKey...)
	tx := types.NewTransaction(nonce, randomizeAddr, big.NewInt(0), 200000, gasPrice, inputData)

	return tx, nil
}

// GetCandidates get candidates and caps from contract
func GetCandidates(client *ethclient.Client, number *big.Int) (map[common.Address]*big.Int, error) {
	addr := common.HexToAddress(common.MasternodeVotingSMC)
	validator, err := contractValidator.NewJuncaValidator(addr, client)
	if err != nil {
		return nil, err
	}
	opts := new(bind.CallOpts)
	candidates, err := validator.GetCandidates(opts)
	if err != nil {
		return nil, err
	}
	ret := make(map[common.Address]*big.Int)
	var zeroAddress = common.Address{}
	for _, v := range candidates {
		if bytes.Equal(v.Bytes(), zeroAddress.Bytes()) {
			continue
		}
		cap, err := validator.GetCandidateCap(opts, v)
		if err != nil {
			return nil, err
		}
		ret[v] = cap
	}
	return ret, nil
}

// GetSignersFromContract Get signers signed for blockNumber from blockSigner contract.
func GetSignersFromContract(stateDB *state.StateDB, block *types.Block) ([]common.Address, error) {
	return state.GetSigners(stateDB, block), nil
}

// GetSignersByExecutingEVM Get signers signed for blockNumber from blockSigner contract.
func GetSignersByExecutingEVM(addrBlockSigner common.Address, client bind.ContractBackend, blockHash common.Hash) ([]common.Address, error) {
	blockSigner, err := contract.NewBlockSigner(addrBlockSigner, client)
	if err != nil {
		log.Error("Fail get instance of blockSigner", "error", err)
		return nil, err
	}
	opts := new(bind.CallOpts)
	addrs, err := blockSigner.GetSigners(opts, blockHash)
	if err != nil {
		log.Error("Fail get block signers", "error", err)
		return nil, err
	}
	return addrs, nil
}

// GetRandomizeFromContract  random from randomize contract.
func GetRandomizeFromContract(client bind.ContractBackend, addrMasternode common.Address) (int64, error) {
	randomize, err := randomizeContract.NewJuncaRandomize(common.HexToAddress(common.RandomizeSMC), client)
	if err != nil {
		log.Error("Fail to get instance of randomize", "error", err)
	}
	opts := new(bind.CallOpts)
	secrets, err := randomize.GetSecret(opts, addrMasternode)
	if err != nil {
		log.Error("Fail get secrets from randomize", "error", err)
	}
	opening, err := randomize.GetOpening(opts, addrMasternode)
	if err != nil {
		log.Error("Fail get opening from randomize", "error", err)
	}

	return DecryptRandomizeFromSecretsAndOpening(secrets, opening)
}

func GetValidators(client *ethclient.Client, masters []common.Address) ([]common.Address, error) {
	// Check m2 exists on chaindb.
	// Get secrets and opening at epoc block checkpoint.
	var candidates []int64
	lenSigners := int64(len(masters))
	if lenSigners > 0 {
		for _, addr := range masters {
			random, err := GetRandomizeFromContract(client, addr)
			if err != nil {
				return nil, err
			}
			candidates = append(candidates, random)
		}
		// Get randomize m2 list.
		return GenM2FromRandomize(candidates, masters)
	}
	return nil, errors.New("Not found M1")
}

// GenM2FromRandomize Generate m2 listing from randomize array.
func GenM2FromRandomize(randomizes []int64, masternodes []common.Address) ([]common.Address, error) {
	total := int64(0)
	for _, j := range randomizes {
		total += j
	}

	r := rand.New(rand.NewSource(total))
	indexs := r.Perm(len(masternodes))

	var m2s = make([]common.Address, len(masternodes))
	for i, v := range masternodes {
		m2s[indexs[i]] = v
	}
	return m2s, nil
}

// DecodeValidatorsHexData Decode validator hex string.
func DecodeValidatorsHexData(validatorsStr string) ([]int64, error) {
	validatorsByte, err := hexutil.Decode(validatorsStr)
	if err != nil {
		return nil, err
	}

	return posv.ExtractValidatorsFromBytes(validatorsByte), nil
}

// DecryptRandomizeFromSecretsAndOpening Decrypt randomize from secrets and opening.
func DecryptRandomizeFromSecretsAndOpening(secrets [][32]byte, opening [32]byte) (int64, error) {
	var random int64
	if len(secrets) > 0 {
		for _, secret := range secrets {
			trimSecret := bytes.TrimLeft(secret[:], "\x00")
			decryptSecret := Decrypt(opening[:], string(trimSecret))
			if isInt(decryptSecret) {
				intNumber, err := strconv.Atoi(decryptSecret)
				if err != nil {
					log.Error("Can not convert string to integer", "error", err)
					return -1, err
				}
				random = int64(intNumber)
			}
		}
	}

	return random, nil
}

// Calculate reward for reward checkpoint.
//func GetRewardForCheckpoint(c *posv.PoSV, chain consensus.ChainReader, header *types.Header, rCheckpoint uint64, totalSigner *uint64) (map[common.Address]*rewardLog, error) {
//	// Not reward for singer of genesis block and only calculate reward at checkpoint block.
//	number := header.Number.Uint64()
//	prevCheckpoint := number - (rCheckpoint * 2)
//	startBlockNumber := prevCheckpoint + 1
//	endBlockNumber := startBlockNumber + rCheckpoint - 1
//	signers := make(map[common.Address]*rewardLog)
//	mapBlkHash := map[uint64]common.Hash{}
//
//	data := make(map[common.Hash][]common.Address)
//	for i := prevCheckpoint + (rCheckpoint * 2) - 1; i >= startBlockNumber; i-- {
//		header = chain.GetHeader(header.ParentHash, i)
//		mapBlkHash[i] = header.Hash()
//		signData, ok := c.BlockSigners.Get(header.Hash())
//		if !ok {
//			log.Debug("Failed get from cached", "hash", header.Hash().String(), "number", i)
//			block := chain.GetBlock(header.Hash(), i)
//			txs := block.Transactions()
//			if !chain.Config().IsTIPSigning(header.Number) {
//				receipts := core.GetBlockReceipts(c.GetDb(), header.Hash(), i)
//				signData = c.CacheData(header, txs, receipts)
//			} else {
//				signData = c.CacheSigner(header.Hash(), txs)
//			}
//		}
//		txs := signData.([]*types.Transaction)
//		for _, tx := range txs {
//			blkHash := common.BytesToHash(tx.Data()[len(tx.Data())-32:])
//			from := *tx.From()
//			data[blkHash] = append(data[blkHash], from)
//		}
//	}
//	header = chain.GetHeader(header.ParentHash, prevCheckpoint)
//	masternodes := posv.GetMasternodesFromCheckpointHeader(header)
//
//	for i := startBlockNumber; i <= endBlockNumber; i++ {
//		if i%common.MergeSignRange == 0 || !chain.Config().IsTIP2019(big.NewInt(int64(i))) {
//			addrs := data[mapBlkHash[i]]
//			// Filter duplicate address.
//			if len(addrs) > 0 {
//				addrSigners := make(map[common.Address]bool)
//				for _, masternode := range masternodes {
//					for _, addr := range addrs {
//						if addr == masternode {
//							if _, ok := addrSigners[addr]; !ok {
//								addrSigners[addr] = true
//							}
//							break
//						}
//					}
//				}
//
//				for addr := range addrSigners {
//					_, exist := signers[addr]
//					if exist {
//						signers[addr].Sign++
//					} else {
//						signers[addr] = &rewardLog{1, new(big.Int)}
//					}
//					*totalSigner++
//				}
//			}
//		}
//	}
//
//	log.Info("Calculate reward at checkpoint", "startBlock", startBlockNumber, "endBlock", endBlockNumber)
//
//	return signers, nil
//}

// Calculate reward for signers.
//func CalculateRewardForSigner(chainReward *big.Int, signers map[common.Address]*rewardLog, totalSigner uint64) (map[common.Address]*big.Int, error) {
//	resultSigners := make(map[common.Address]*big.Int)
//	// Add reward for signers.
//	if totalSigner > 0 {
//		for signer, rLog := range signers {
//			// Add reward for signer.
//			calcReward := new(big.Int)
//			calcReward.Div(chainReward, new(big.Int).SetUint64(totalSigner))
//			calcReward.Mul(calcReward, new(big.Int).SetUint64(rLog.Sign))
//			rLog.Reward = calcReward
//
//			resultSigners[signer] = calcReward
//		}
//	}
//	jsonSigners, err := json.Marshal(signers)
//	if err != nil {
//		log.Error("Fail to parse json signers", "error", err)
//		return nil, err
//	}
//	log.Info("Signers data", "signers", string(jsonSigners), "totalSigner", totalSigner, "totalReward", chainReward)
//
//	return resultSigners, nil
//}

// GetCandidatesOwnerBySigner Get candidate owner by address.
func GetCandidatesOwnerBySigner(stateDB *state.StateDB, signerAddr common.Address) common.Address {
	owner := state.GetCandidateOwner(stateDB, signerAddr)
	return owner
}

//func CalculateRewardForHolders(foundationWalletAddr common.Address, state *state.StateDB, signer common.Address, calcReward *big.Int, blockNumber uint64) (error, map[common.Address]*big.Int) {
//	rewards, err := GetRewardBalancesRate(foundationWalletAddr, state, signer, calcReward, blockNumber)
//	if err != nil {
//		return err, nil
//	}
//	return nil, rewards
//}

//func GetRewardBalancesRate(foundationWalletAddr common.Address, stateDB *state.StateDB, masterAddr common.Address, totalReward *big.Int, blockNumber uint64) (map[common.Address]*big.Int, error) {
//	owner := GetCandidatesOwnerBySigner(stateDB, masterAddr)
//	balances := make(map[common.Address]*big.Int)
//	rewardMaster := new(big.Int).Mul(totalReward, new(big.Int).SetInt64(common.RewardMasterPercent))
//	rewardMaster = new(big.Int).Div(rewardMaster, new(big.Int).SetInt64(100))
//	balances[owner] = rewardMaster
//	// Get voters for masternode.
//	voters := state.GetVoters(stateDB, masterAddr)
//
//	if len(voters) > 0 {
//		totalVoterReward := new(big.Int).Mul(totalReward, new(big.Int).SetUint64(common.RewardVoterPercent))
//		totalVoterReward = new(big.Int).Div(totalVoterReward, new(big.Int).SetUint64(100))
//		totalCap := new(big.Int)
//		// Get voters capacities.
//		voterCaps := make(map[common.Address]*big.Int)
//		for _, voteAddr := range voters {
//			if _, ok := voterCaps[voteAddr]; ok && common.TIP2019Block.Uint64() <= blockNumber {
//				continue
//			}
//			voterCap := state.GetVoterCap(stateDB, masterAddr, voteAddr)
//			totalCap.Add(totalCap, voterCap)
//			voterCaps[voteAddr] = voterCap
//		}
//		if totalCap.Cmp(new(big.Int).SetInt64(0)) > 0 {
//			for addr, voteCap := range voterCaps {
//				// Only valid voter has cap > 0.
//				if voteCap.Cmp(new(big.Int).SetInt64(0)) > 0 {
//					rcap := new(big.Int).Mul(totalVoterReward, voteCap)
//					rcap = new(big.Int).Div(rcap, totalCap)
//					if balances[addr] != nil {
//						balances[addr].Add(balances[addr], rcap)
//					} else {
//						balances[addr] = rcap
//					}
//				}
//			}
//		}
//	}
//
//	foundationReward := new(big.Int).Mul(totalReward, new(big.Int).SetInt64(common.RewardFoundationPercent))
//	foundationReward = new(big.Int).Div(foundationReward, new(big.Int).SetInt64(100))
//	balances[foundationWalletAddr] = foundationReward
//
//	jsonHolders, err := json.Marshal(balances)
//	if err != nil {
//		log.Error("Fail to parse json holders", "error", err)
//		return nil, err
//	}
//	log.Trace("Holders reward", "holders", string(jsonHolders), "masternode", masterAddr.String())
//
//	return balances, nil
//}

// NewSlice Dynamic generate array sequence of numbers.
func NewSlice(start int64, end int64, step int64) []int64 {
	s := make([]int64, end-start)
	for i := range s {
		s[i] = start
		start += step
	}

	return s
}

// Shuffle array.
func Shuffle(slice []int64) []int64 {
	newSlice := make([]int64, len(slice))
	copy(newSlice, slice)

	for i := 0; i < len(slice)-1; i++ {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(len(newSlice))
		x := newSlice[i]
		newSlice[i] = newSlice[randIndex]
		newSlice[randIndex] = x
	}

	return newSlice
}

// Encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error("Fail to encrypt", "err", err)
		return ""
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cryptoRand.Reader, iv); err != nil {
		log.Error("Fail to encrypt iv", "err", err)
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// Decrypt from base64 to decrypted string
func Decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error("Fail to decrypt", "err", err)
		return ""
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		log.Error("ciphertext too short")
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

// RandStringByte Generate random string.
func RandStringByte(n int) []byte {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

// Helper function check string is numeric.
func isInt(strNumber string) bool {
	if _, err := strconv.Atoi(strNumber); err == nil {
		return true
	} else {
		return false
	}
}
