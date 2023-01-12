package posv

import (
	"bytes"
	"math/big"

	"github.com/juncachain/juncachain/common"
)

type MasterNode struct {
	Address common.Address
	Stake   *big.Int
}

type MasterNodes []MasterNode

func (ms MasterNodes) Len() int {
	return len(ms)
}

func (ms MasterNodes) Less(i, j int) bool {
	if ms[i].Stake.Cmp(ms[i].Stake) > 0 {
		return true
	}
	return bytes.Compare(ms[i].Address[:], ms[j].Address[:]) > 0
}

func (ms MasterNodes) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms MasterNodes) M1(epochLength uint64, number uint64) common.Address {
	if len(ms) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	var index = 0
	if mod != 0 {
		index = len(ms) % int(mod)
	}
	return ms[index].Address
}
