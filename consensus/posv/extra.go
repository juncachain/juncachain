package posv

import (
	"bytes"
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

func (ms MasterNodes) M1(epochLength uint64, number uint64) common.Address {
	if len(ms) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	var index = 0
	if mod != 0 {
		index = int(mod) % len(ms)
	}
	return ms[index].Address
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
	Checkpoint  uint64           `json:"checkpoint"`
	MasterNodes MasterNodes      `json:"masterNodes"` // max 150
	Validators  MasterNodes      `json:"validators"`  // max 99
	Penalties   []common.Address `json:"penalties"`
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

func (e *Epoch) Signer(epochLength uint64, number uint64) common.Address {
	if len(e.Validators) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	return e.Validators[int(mod)%len(e.Validators)].Address
}

func (e *Epoch) Signers() []common.Address {
	var signers []common.Address
	for _, v := range e.Validators {
		signers = append(signers, v.Address)
	}
	return signers
}
