package contract

import (
	"context"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common/hexutil"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
	"math/big"
	"testing"
)

var (
	initBal, _ = new(big.Int).SetString("1000000000000000000", 10)
	key, _     = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr       = crypto.PubkeyToAddress(key.PublicKey)
)

func TestDeployJRC21PresetMinter(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr: {Balance: initBal},
	}, 3000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	transactOpts.GasLimit = 3000000
	contractAddr, _, _, err := DeployJRC21PresetMinter(transactOpts, contractBackend, "example token",
		"EPT", uint8(8), addr, big.NewInt(100))
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
