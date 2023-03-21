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

package juncaswap

import (
	"context"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts/WJGC"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
	"math/big"
	"testing"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestGetCodeHash(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000000)}}, 10000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))

	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	WJGC, _, err := WJGC.DeployWrappedJGC(transactOpts, contractBackend)
	if err != nil {
		t.Fatal(err)
	}

	_, factory, err := DeployJuncaswapFactory(transactOpts, contractBackend, WJGC)
	if err != nil {
		t.Fatal(err)
	}
	contractBackend.Commit()
	codeHash, err := factory.GetCodeHash()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(common.BytesToHash(codeHash[:]).Hex())
}
