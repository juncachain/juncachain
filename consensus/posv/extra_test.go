package posv

import (
	"bytes"
	"fmt"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"math/big"
	"testing"
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
	extra.Epoch.Penalties = append(extra.Epoch.Penalties, common.BigToAddress(big.NewInt(2)))
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
	if new(big.Int).SetBytes(newExtra.Epoch.Penalties[0].Bytes()).Cmp(big.NewInt(2)) != 0 {
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
		fmt.Println(i, " ", m1)
	}
}

func TestM2(t *testing.T) {
	var epoch Epoch
	epoch.Checkpoint = 0
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(1)))
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(2)))
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(3)))
	epoch.M2s = append(epoch.M2s, common.BigToAddress(new(big.Int).SetInt64(4)))

	for i := 1; i <= 20; i++ {
		var str = fmt.Sprintf("block:%d", i)
		for _, v := range epoch.M2(9, uint64(i)) {
			str = str + fmt.Sprintf(" %s ", v.Hex())
		}
		fmt.Println(str)
	}
}
