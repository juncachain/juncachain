// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package posv

import (
	"math/big"
	"os"
	"testing"

	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/core/rawdb"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/params"
)

var (
	_engine *PoSV
)

func TestMain(m *testing.M) {
	cfg := &params.PoSVConfig{
		Period:      2,
		Epoch:       900,
		MinStaked:   nil,
		Reward:      new(big.Int).Mul(big.NewInt(150), big.NewInt(params.Ether)),
		TotalReward: new(big.Int).Mul(big.NewInt(10000000), big.NewInt(params.Ether)),
		InstanceDir: "",
	}
	_engine = New(cfg, rawdb.NewMemoryDatabase())

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSealHash(t *testing.T) {
	have := SealHash(&types.Header{
		Difficulty: new(big.Int),
		Number:     new(big.Int),
		Extra:      make([]byte, 32+65),
		BaseFee:    new(big.Int),
	})
	want := common.HexToHash("0xbd3d1fa43fbc4c5bfcc91b179ec92e2861df3654de60468beb908ff805359e8f")
	if have != want {
		t.Errorf("have %x, want %x", have, want)
	}
}

func TestCalcRewards(t *testing.T) {
	base := uint64(1)
	for epoch := uint64(1); epoch < 184667; epoch += 17520 {
		t.Log(_engine.CalcReward(900*epoch + base))
	}

	for number := uint64(1); number < 166200300; number += 17520 * 900 {
		t.Log(_engine.CalcReward(number))
	}

	t.Log(166200300, _engine.CalcReward(166200300))
	t.Log(166200301, _engine.CalcReward(166200301))
	t.Log(166201201, _engine.CalcReward(166201201))
	if _engine.CalcReward(166201201).Cmp(big.NewInt(0)) != 0 {
		t.Fatal("not equal")
	}
}

func TestLastEpoch(t *testing.T) {
	if _engine.LastEpoch(1) != 0 {
		t.Fatal("got bad checkpoint number")
	}
	if _engine.LastEpoch(900) != 0 {
		t.Fatal("got bad checkpoint number")
	}
	if _engine.LastEpoch(901) != 900 {
		t.Fatal("got bad checkpoint number")
	}
}
