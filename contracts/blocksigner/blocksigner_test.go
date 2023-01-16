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

package blocksigner

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestBlockSigner(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000000)}}, 0)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(key, contractBackend.Blockchain().Config().ChainID)
	if err != nil {
		t.Fatal(err)
	}

	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))

	transactOpts.GasPrice = gasPrice
	blockSignerAddress, blockSigner, err := DeployBlockSigner(transactOpts, contractBackend, big.NewInt(99))
	if err != nil {
		t.Fatalf("can't deploy root registry: %v", err)
	}
	contractBackend.Commit()

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, _ := contractBackend.CodeAt(ctx, blockSignerAddress, nil)
	t.Log("contract code", hexutil.Encode(code))

	head, _ = contractBackend.HeaderByNumber(context.Background(), nil)
	f := func(key, val common.Hash) bool {
		t.Log(key.Hex(), val.Hex())
		return true
	}
	stateDB, err := contractBackend.Blockchain().StateAt(head.Root)
	if err != nil {
		t.Fatalf("can't get stateDB by root: %v err:%v", head.Root.Hex(), err)
	}
	_ = stateDB.ForEachStorage(blockSignerAddress, f)

	byte0 := randomHash()

	// Test sign.
	tx, err := blockSigner.Sign(big.NewInt(2), byte0)
	if err != nil {
		t.Fatalf("can't sign: %v", err)
	}
	contractBackend.Commit()
	t.Log("tx", tx)

	signers, err := blockSigner.GetSigners(byte0)
	if err != nil {
		t.Fatalf("can't get candidates: %v", err)
	}
	for _, it := range signers {
		t.Log("signer", it.String())
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
