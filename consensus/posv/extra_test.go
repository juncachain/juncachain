// Package posv Copyright (c) 2023 Juncachain
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

package posv

import (
	"bytes"
	"math/big"
	"sort"
	"testing"

	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
)

func TestEncodeDecode(t *testing.T) {
	var extra = Extra{
		Vanity: [32]byte{0},
		Epoch: Epoch{
			Checkpoint: 235,
			M1s:        nil,
			M2s:        nil,
			Reward:     nil,
			Penalties:  nil,
		},
		Signature: [65]byte{2},
	}
	vanity := "0xc37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487befefef"
	signature := "0xc37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487bef2501c37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487bef250102"
	copy(extra.Vanity[:], hexutil.MustDecode(vanity))
	copy(extra.Signature[:], hexutil.MustDecode(signature))
	extra.Epoch.M1s = append(extra.Epoch.M1s, MasterNode{
		Address: common.BigToAddress(big.NewInt(1)),
		Stake:   big.NewInt(50000),
	})
	extra.Epoch.Penalties = append(extra.Epoch.Penalties, Penalty{
		Address: common.BigToAddress(big.NewInt(2)),
		Miss:    big.NewInt(1),
	})
	b := extra.ToBytes()

	var newExtra Extra
	if err := newExtra.FromBytes(b); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(newExtra.Vanity[:], extra.Vanity[:]) {
		t.Fatal("Vanity not equal")
	}
	if !bytes.Equal(newExtra.Signature[:], extra.Signature[:]) {
		t.Fatal("Signature not equal")
	}
	if new(big.Int).SetBytes(newExtra.Epoch.M1s[0].Address.Bytes()).Cmp(big.NewInt(1)) != 0 {
		t.Fatal("Validators not equal")
	}
	if new(big.Int).SetBytes(newExtra.Epoch.Penalties[0].Address.Bytes()).Cmp(big.NewInt(2)) != 0 {
		t.Fatal("Penalties not equal")
	}
}

func TestM1(t *testing.T) {
	var epoch Epoch
	epoch.Checkpoint = 0
	epoch.M1s = append(epoch.M1s, MasterNode{
		Address: common.BigToAddress(new(big.Int).SetInt64(1)),
		Stake:   nil,
	})
	epoch.M1s = append(epoch.M1s, MasterNode{
		Address: common.BigToAddress(new(big.Int).SetInt64(2)),
		Stake:   nil,
	})
	epoch.M1s = append(epoch.M1s, MasterNode{
		Address: common.BigToAddress(new(big.Int).SetInt64(3)),
		Stake:   nil,
	})

	for i := 1; i <= 20; i++ {
		m1 := epoch.M1(9, uint64(i))
		nextSealBlock := epoch.M1NextTurn(9, uint64(i), m1)
		t.Log("Number:", i, " M1:", m1, " nextSealBlock:", nextSealBlock)
		if m1 != epoch.M1s[(i-1)%epoch.M1Length()].Address {
			t.Fatal("not in order")
		}
	}

	for i := 1; i < 20; i++ {
		m1 := common.BigToAddress(new(big.Int).SetInt64(1))
		nextSealBlock := epoch.M1NextTurn(9, uint64(i), m1)
		t.Log("Number:", i, " M1:", m1, " nextSealBlock:", nextSealBlock)
	}
}

func TestM2(t *testing.T) {
	var epoch Epoch
	epoch.Checkpoint = 0
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(1)))
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(2)))
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(3)))

	for i := 1; i <= 20; i++ {
		m2s := epoch.M2(9, uint64(i))
		if len(m2s) != 2 {
			t.Fatal("got no m2")
		}
		if m2s[0] != epoch.M2s[(i)%epoch.M2Length()] {
			t.Fatal("The first M2 not in order")
		}
		if m2s[1] != epoch.M2s[(i+1)%epoch.M2Length()] {
			t.Fatal("The second M2 not in order")
		}
	}
}

func TestSortMasterNodes(t *testing.T) {
	var masternodes MasterNodes
	masternodes = append(masternodes, MasterNode{
		Address: common.BigToAddress(big.NewInt(1)),
		Stake:   big.NewInt(1),
	})
	masternodes = append(masternodes, MasterNode{
		Address: common.BigToAddress(big.NewInt(3)),
		Stake:   big.NewInt(3),
	})
	masternodes = append(masternodes, MasterNode{
		Address: common.BigToAddress(big.NewInt(2)),
		Stake:   big.NewInt(2),
	})
	masternodes = append(masternodes, MasterNode{
		Address: common.BigToAddress(big.NewInt(5)),
		Stake:   big.NewInt(5),
	})
	masternodes = append(masternodes, MasterNode{
		Address: common.BigToAddress(big.NewInt(4)),
		Stake:   big.NewInt(5),
	})
	sort.Sort(masternodes)
	for i, v := range masternodes {
		t.Log(v.Address, v.Stake)
		if i == 0 && v.Address != common.BigToAddress(big.NewInt(4)) {
			t.Fatal("sort failed")
		}
		if i == 1 && v.Address != common.BigToAddress(big.NewInt(5)) {
			t.Fatal("sort failed")
		}
		if i == 2 && v.Address != common.BigToAddress(big.NewInt(3)) {
			t.Fatal("sort failed")
		}
		if i == 3 && v.Address != common.BigToAddress(big.NewInt(2)) {
			t.Fatal("sort failed")
		}
		if i == 4 && v.Address != common.BigToAddress(big.NewInt(1)) {
			t.Fatal("sort failed")
		}
	}
}
