package posv

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/rlp"
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

type Extra struct {
	Vanity    [extraVanity]byte `json:"vanity"`
	Epoch     Epoch             `json:"epoch"`
	Signature [extraSeal]byte   `json:"signature"`
}

func (e *Extra) FromBytes(b []byte) error {
	if len(b) < extraVanity+extraSeal {
		return errors.New("extra-data 32 byte vanity prefix or 65 byte signature suffix missing")
	}
	copy(e.Vanity[:], b[:extraVanity])
	copy(e.Signature[:], b[len(b)-extraSeal:])
	if err := e.Epoch.FromBytes(b[extraVanity : len(b)-extraSeal]); err != nil {
		return err
	}
	return nil
}

func (e *Extra) ToBytes() []byte {
	eb := e.Epoch.ToBytes()
	b := make([]byte, len(eb)+extraVanity+extraSeal)
	copy(b[:], e.Vanity[:])
	copy(b[extraVanity:len(b)-extraSeal], eb)
	copy(b[len(b)-extraSeal:], e.Signature[:])
	return b
}

type Epoch struct {
	Checkpoint uint64           `json:"checkpoint"`
	M1s        MasterNodes      `json:"m1s"` // max 99
	M2s        []common.Address `json:"m2s"` // max 150
	Reward     *big.Int         `json:"reward"`
	Penalties  []common.Address `json:"penalties"`
}

func (e *Epoch) FromBytes(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return rlp.DecodeBytes(b, e)
}

func (e *Epoch) ToBytes() []byte {
	b, err := rlp.EncodeToBytes(e)
	if err != nil {
		log.Crit("Epoch ToBytes", "error", err)
	}
	return b
}

func (e *Epoch) IsM1(address common.Address) bool {
	for _, v := range e.M1s {
		if v.Address == address {
			return true
		}
	}
	return false
}

func (e *Epoch) M1(epochLength uint64, number uint64) common.Address {
	if len(e.M1s) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	index := int(mod)%len(e.M1s) - 1
	if index < 0 {
		index = len(e.M1s) - 1
	}
	return e.M1s[index].Address
}

func (e *Epoch) NextTurn(epochLength uint64, number uint64, m1 common.Address) uint64 {
	for i := number; i < number+epochLength; i++ {
		if e.M1(epochLength, i) == m1 {
			return i
		}
	}
	return number + epochLength
}

func (e *Epoch) M1Slice() []common.Address {
	var m1s []common.Address
	for _, v := range e.M1s {
		m1s = append(m1s, v.Address)
	}
	return m1s
}

func (e *Epoch) M1Index(signer common.Address) int {
	for i, v := range e.M1s {
		if signer == v.Address {
			return i
		}
	}
	return -1
}

func (e *Epoch) M1Length() int {
	return len(e.M1s)
}

func (e *Epoch) M2(epochLength uint64, number uint64) common.Address {
	if len(e.M2s) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	return e.M2s[int(mod)%len(e.M1s)]
}

func (e *Epoch) M2Slice() []common.Address {
	var m2s []common.Address
	for _, v := range e.M2s {
		m2s = append(m2s, v)
	}
	return m2s
}

func (e *Epoch) M2Length() int {
	return len(e.M2s)
}

func (e *Epoch) String() string {
	bz, _ := json.Marshal(e)
	return string(bz)
}
