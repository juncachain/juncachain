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

package jrc21issuer

import (
	"context"
	"math/big"
	"testing"

	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/contracts/jrc21issuer/contract"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
	"github.com/juncachain/juncachain/params"
)

var (
	key, _         = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr           = crypto.PubkeyToAddress(key.PublicKey)
	initBal        = new(big.Int).Mul(big.NewInt(5000000), big.NewInt(params.Ether))
	gasLimit       = uint64(30000000)
	minFee         = new(big.Int).SetInt64(100)
	tokenSupply, _ = new(big.Int).SetString("100000000", 10)
)

func TestJRC21Issuer(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: initBal}}, gasLimit)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}

	jrc21IssuerAddress, jrc21Issuer, err := DeployJRC21Issuer(transactOpts, contractBackend, minFee)
	t.Log("contract address", jrc21IssuerAddress.String())
	if err != nil {
		t.Fatalf("can't DeployJRC21Issuer: %v", err)
	}
	root := contractBackend.Commit()

	code, err := contractBackend.CodeAt(context.Background(), jrc21IssuerAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("contract code", hexutil.Encode(code))

	stateDB, err := contractBackend.Blockchain().State()
	if err != nil {
		t.Fatalf("can't get stateDB by root: %v err:%v", root.Hex(), err)
	}
	t.Log(string(stateDB.Dump(nil)))
	f := func(key, val common.Hash) bool {
		t.Log(key.Hex(), val.Hex())
		return true
	}
	_ = stateDB.ForEachStorage(jrc21IssuerAddress, f)

	gas, err := jrc21Issuer.Gas()
	if err != nil {
		t.Fatalf("can't get Gas: %v", err)
	}
	if gas.Cmp(minFee) != 0 {
		t.Fatalf("gas not equal: %v", gas)
	}

	// Test issue.
	transactOpts, err = bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}
	transactOpts.Value = new(big.Int).Mul(big.NewInt(50000), big.NewInt(params.Ether))
	jrc21Issuer, err = NewJRC21Issuer(transactOpts, jrc21IssuerAddress, contractBackend)
	if err != nil {
		t.Fatal(err)
	}
	_, err = jrc21Issuer.IssueJRC21PresetFixed("example JRC21 token", "EJT", 8,
		addr, tokenSupply, new(big.Int).SetInt64(10000000))
	if err != nil {
		t.Fatalf("can't sign: %v", err)
	}
	contractBackend.Commit()

	exampleTokenAddr, err := jrc21Issuer.Tokens(big.NewInt(0))
	if err != nil {
		t.Fatalf("can't get Tokens: %v", err)
	}

	t.Log("exampleTokenAddr", exampleTokenAddr.Hex())

	jrc21PresetFixed, err := contract.NewJRC21PresetFixedCaller(exampleTokenAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't sign: %v", err)
	}
	bal, err := jrc21PresetFixed.BalanceOf(nil, addr)
	if err != nil {
		t.Fatalf("can't sign: %v", err)
	}
	if bal.Cmp(tokenSupply) != 0 {
		t.Fatal("balance is not equal")
	}
}
