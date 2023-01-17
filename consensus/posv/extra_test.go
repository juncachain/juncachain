package posv

import (
	"bytes"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/common/hexutil"
	"math/big"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	var extra = Extra{
		Vanity: [32]byte{0},
		Epoch: Epoch{
			Checkpoint:  235,
			MasterNodes: nil,
			Validators:  nil,
			Penalties:   nil,
		},
		Signature: [65]byte{2},
	}
	vanity := "0xc37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487befefef"
	signature := "0xc37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487bef2501c37a7a0937c1e6bc7364909242c47ac3ae9ab92be957978da6f0eb487bef250102"
	copy(extra.Vanity[:], hexutil.MustDecode(vanity))
	copy(extra.Signature[:], hexutil.MustDecode(signature))
	extra.Epoch.Validators = append(extra.Epoch.Validators, MasterNode{
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
	if new(big.Int).SetBytes(newExtra.Epoch.Validators[0].Address.Bytes()).Cmp(big.NewInt(1)) != 0 {
		t.Fatal("Validators not equal")
	}
	if new(big.Int).SetBytes(newExtra.Epoch.Penalties[0].Bytes()).Cmp(big.NewInt(2)) != 0 {
		t.Fatal("Penalties not equal")
	}
}
