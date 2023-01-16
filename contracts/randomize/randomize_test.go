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

package randomize

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/juncachain/juncachain/common/hexutil"

	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/crypto"
)

var (
	initBal, _ = new(big.Int).SetString("1000000000000000000", 10)
	epocNumber = int64(12)
	key, _     = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr       = crypto.PubkeyToAddress(key.PublicKey)
	byte0      = make([][32]byte, epocNumber)
	acc1Key, _ = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	acc1Addr   = crypto.PubkeyToAddress(acc1Key.PublicKey)
)

func TestRandomize(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: initBal},
		acc1Addr: {Balance: initBal},
	}, 10000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	transactOpts.GasLimit = 1000000

	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	randomizeAddress, randomize, err := DeployRandomize(transactOpts, contractBackend)
	t.Log("contract address", randomizeAddress.String())
	if err != nil {
		t.Fatalf("can't deploy root registry: %v", err)
	}
	contractBackend.Commit()

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, _ := contractBackend.CodeAt(ctx, randomizeAddress, nil)
	t.Log("contract code", hexutil.Encode(code))
	f := func(key, val common.Hash) bool {
		t.Log(key.Hex(), val.Hex())
		return true
	}

	head, _ = contractBackend.HeaderByNumber(context.Background(), nil)
	stateDB, err := contractBackend.Blockchain().StateAt(head.Root)
	if err != nil {
		t.Fatalf("can't get stateDB by root: %v err:%v", head.Root.Hex(), err)
	}
	_ = stateDB.ForEachStorage(randomizeAddress, f)

	s, err := randomize.SetSecret(byte0)
	if err != nil {
		t.Fatalf("can't set secret: %v", err)
	}
	t.Log("tx data", s)
	contractBackend.Commit()
}

func TestSendTxRandomizeSecretAndOpening(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: initBal},
		acc1Addr: {Balance: initBal},
	}, 10000000)
	backend.Commit()
	signer := types.HomesteadSigner{}
	ctx := context.Background()

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(acc1Key, big.NewInt(1337))
	transactOpts.GasLimit = 4200000
	epocNumber := uint64(900)

	head, _ := backend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	randomizeAddr, randomizeContract, err := DeployRandomize(transactOpts, backend)
	if err != nil {
		t.Fatalf("Can't deploy randomize SC: %v", err)
	}
	backend.Commit()

	randomizeKeyValue := contracts.RandStringByte(32)

	for i := 1; i <= 900; i++ {
		nonce := uint64(i)
		switch i {
		case 800:
			tx, err := contracts.BuildTxSecretRandomize(nonce, gasPrice, randomizeAddr, epocNumber, randomizeKeyValue)
			if err != nil {
				t.Fatalf("Can't create tx randomize secret: %v", err)
			}
			tx, err = types.SignTx(tx, signer, acc1Key)
			if err != nil {
				t.Fatalf("Can't sign tx randomize secret: %v", err)
			}
			err = backend.SendTransaction(ctx, tx)
			if err != nil {
				t.Fatalf("Can't send tx for create randomize secret: %v", err)
			}
		case 850:
			// Set opening.
			tx, err := contracts.BuildTxOpeningRandomize(nonce, gasPrice, randomizeAddr, randomizeKeyValue)
			if err != nil {
				t.Fatalf("Can't create tx randomize opening: %v", err)
			}
			tx, err = types.SignTx(tx, signer, acc1Key)
			if err != nil {
				t.Fatalf("Can't sign tx randomize opening: %v", err)
			}
			err = backend.SendTransaction(ctx, tx)
			if err != nil {
				t.Fatalf("Can't send tx for create randomize opening: %v", err)
			}

		case 900:
			// Get randomize secret from SC.
			secrets, err := randomizeContract.GetSecret(acc1Addr)
			if err != nil {
				t.Error("Fail get secrets from randomize", err)
			}
			if len(secrets) <= 0 {
				t.Error("Empty get secrets from SC", err)
			}
			// Decrypt randomize from SC.
			opening, err := randomizeContract.GetOpening(acc1Addr)
			if err != nil {
				t.Fatalf("Can't get secret from SC: %v", err)
			}
			randomize, err := contracts.DecryptRandomizeFromSecretsAndOpening(secrets, opening)
			t.Log("randomize", randomize)
			if err != nil {
				t.Error("Can't decrypt secret and opening", err)
			}
		default:
			tx, err := types.SignTx(types.NewTransaction(nonce, common.Address{}, new(big.Int), 21000, gasPrice, nil), signer, acc1Key)
			if err != nil {
				t.Fatalf("Can't sign tx randomize: %v", err)
			}
			err = backend.SendTransaction(ctx, tx)
			if err != nil {
				t.Fatalf("Can't send tx for create randomize: %v", err)
			}
		}
		backend.Commit()
	}
}
