package jrc21issuer

import (
	"context"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/contracts/jrc21issuer/contract"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
	"math/big"
	"testing"
)

var (
	key, _         = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr           = crypto.PubkeyToAddress(key.PublicKey)
	initBal, _     = new(big.Int).SetString("1000000000000000000000", 10)
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

	transactOpts.GasLimit = gasLimit
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
