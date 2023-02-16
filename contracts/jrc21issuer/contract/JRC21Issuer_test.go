package contract

import (
	"context"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/core"
	"math/big"
	"testing"
)

func TestDeployJRC21Issuer(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr: {Balance: initBal},
	}, 30000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	transactOpts.GasLimit = 30000000
	contractAddr, _, _, err := DeployJRC21Issuer(transactOpts, contractBackend, big.NewInt(100))
	if err != nil {
		t.Fatal(err)
	}
	contractBackend.Commit()
	code, err := contractBackend.CodeAt(context.Background(), contractAddr, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("contract code", hexutil.Encode(code))
}
