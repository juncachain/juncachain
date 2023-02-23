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
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts/blocksigner"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/crypto"
)

var (
	acc1Key, _ = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	acc2Key, _ = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	acc3Key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	acc4Key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee04aefe388d1e14474d32c45c72ce7b7a")
	acc1Addr   = crypto.PubkeyToAddress(acc1Key.PublicKey)
	acc2Addr   = crypto.PubkeyToAddress(acc2Key.PublicKey)
	acc3Addr   = crypto.PubkeyToAddress(acc3Key.PublicKey)
	acc4Addr   = crypto.PubkeyToAddress(acc4Key.PublicKey)
)

func getCommonBackend() *backends.SimulatedBackend {
	genesis := core.GenesisAlloc{
		acc1Addr: {Balance: big.NewInt(100000000000000000)},
		acc2Addr: {Balance: big.NewInt(100000000000000000)},
		acc3Addr: {Balance: big.NewInt(100000000000000000)},
		acc4Addr: {Balance: big.NewInt(100000000000000000)}}
	backend := backends.NewSimulatedBackend(genesis, 10000000)
	backend.Commit()

	return backend
}

func TestSendTxSign(t *testing.T) {
	accounts := []common.Address{acc2Addr, acc3Addr, acc4Addr}
	keys := []*ecdsa.PrivateKey{acc2Key, acc3Key, acc4Key}
	backend := getCommonBackend()
	signer := types.NewEIP155Signer(big.NewInt(1337))
	ctx := context.Background()

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(acc1Key, big.NewInt(1337))

	head, _ := backend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	blockSignerAddr, blockSigner, err := blocksigner.DeployBlockSigner(transactOpts, backend, big.NewInt(99))
	if err != nil {
		t.Fatalf("Can't get block signer: %v", err)
	}
	backend.Commit()
	t.Log("blockSignerAddr:", blockSignerAddr.Hex())

	nonces := make(map[*ecdsa.PrivateKey]int)
	oldBlocks := make(map[common.Hash]common.Address)

	signTx := func(ctx context.Context, backend *backends.SimulatedBackend, signer types.Signer, nonces map[*ecdsa.PrivateKey]int, accKey *ecdsa.PrivateKey, blockNumber *big.Int, blockHash common.Hash) *types.Transaction {
		tx, err := types.SignTx(BuildTxSignBlockSigner(blockNumber, blockHash, uint64(nonces[accKey]), gasPrice, blockSignerAddr), signer, accKey)
		if err != nil {
			t.Fatalf("Can't sign tx: %v", err)
		}
		backend.SendTransaction(ctx, tx)
		backend.Commit()
		nonces[accKey]++

		return tx
	}

	// Tx sign for signer.
	signCount := int64(0)
	blockHashes := make([]common.Hash, 10)
	for i := int64(0); i < 10; i++ {
		blockHash := randomHash()
		blockHashes[i] = blockHash
		randIndex := rand.Intn(len(keys))
		accKey := keys[randIndex]
		signTx(ctx, backend, signer, nonces, accKey, new(big.Int).SetInt64(i), blockHash)
		oldBlocks[blockHash] = accounts[randIndex]
		signCount++

		// Tx sign for validators.
		for _, key := range keys {
			if key != accKey {
				signTx(ctx, backend, signer, nonces, key, new(big.Int).SetInt64(i), blockHash)
				signCount++
			}
		}
	}

	for _, blockHash := range blockHashes {
		signers, err := blockSigner.GetSigners(blockHash)
		if err != nil {
			t.Fatalf("Can't get signers: %v", err)
		}

		if signers[0].String() != oldBlocks[blockHash].String() {
			t.Errorf("Tx sign for block signer not match %v - %v", signers[0].String(), oldBlocks[blockHash].String())
		}

		if len(signers) != len(keys) {
			t.Error("Tx sign for block validators not match")
		}
	}
}

// Generate random string.
func randomHash() common.Hash {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var b common.Hash
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func TestEncryptDecrypt(t *testing.T) {
	//byteInteger := common.LeftPadBytes([]byte(new(big.Int).SetInt64(4).String()), 32)
	randomByte := RandStringByte(32)
	encrypt := Encrypt(randomByte, new(big.Int).SetInt64(4).String())
	decrypt := Decrypt(randomByte, encrypt)
	t.Log("Encrypt", encrypt, "Test", string(randomByte), "Decrypt", decrypt, "trim", string(bytes.TrimLeft([]byte(decrypt), "\x00")))
}

// Unit test for
func TestGenM2FromRandomize(t *testing.T) {
	var a []int64
	var masters []common.Address
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		a = append(a, int64(rand.Intn(9999)))
		masters = append(masters, common.BigToAddress(new(big.Int).SetInt64(int64(i))))
	}
	m2s, err := GenerateM2FromRandomize(a, masters)
	t.Log("randomize", m2s, "len", len(m2s))
	if err != nil {
		t.Error("Fail to test gen m2 for randomize.", err)
	}
	for _, v := range m2s {
		fmt.Println(v)
	}
}
