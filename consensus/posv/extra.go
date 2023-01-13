package posv

import (
	"errors"

	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/rlp"
)

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
	Validators []common.Address `json:"validators"`
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

func (e *Epoch) Validator(epochLength uint64, number uint64) common.Address {
	if len(e.Validators) == 0 {
		return common.Address{}
	}
	mod := number % epochLength
	return e.Validators[int(mod)%len(e.Validators)]
}
