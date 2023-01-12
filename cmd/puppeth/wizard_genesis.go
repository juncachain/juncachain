// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/accounts/abi/bind/backends"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/consensus/posv"
	blockSignerContract "github.com/juncachain/juncachain/contracts/blocksigner"
	randomizeContract "github.com/juncachain/juncachain/contracts/randomize"
	validatorContract "github.com/juncachain/juncachain/contracts/validator"
	"github.com/juncachain/juncachain/core"
	"github.com/juncachain/juncachain/crypto"
	"github.com/juncachain/juncachain/log"
	"github.com/juncachain/juncachain/params"
	"github.com/juncachain/juncachain/rlp"
)

// makeGenesis creates a new genesis struct based on some user input.
func (w *wizard) makeGenesis() {
	// Construct a default genesis block
	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(524288),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock:      big.NewInt(0),
			EIP150Block:         big.NewInt(0),
			EIP155Block:         big.NewInt(0),
			EIP158Block:         big.NewInt(0),
			ByzantiumBlock:      big.NewInt(0),
			ConstantinopleBlock: big.NewInt(0),
			PetersburgBlock:     big.NewInt(0),
			IstanbulBlock:       big.NewInt(0),
		},
	}
	// Figure out which consensus engine to choose
	fmt.Println()
	fmt.Println("Which consensus engine to use? (default = clique)")
	fmt.Println(" 1. Ethash - proof-of-work")
	fmt.Println(" 2. Clique - proof-of-authority")
	fmt.Println(" 3. PoSV   - proof-of-stake voting")

	choice := w.read()
	switch {
	case choice == "1":
		// In case of ethash, we're pretty much done
		genesis.Config.Ethash = new(params.EthashConfig)
		genesis.ExtraData = make([]byte, 32)

	case choice == "" || choice == "2":
		// In the case of clique, configure the consensus parameters
		genesis.Difficulty = big.NewInt(1)
		genesis.Config.Clique = &params.CliqueConfig{
			Period: 15,
			Epoch:  30000,
		}
		fmt.Println()
		fmt.Println("How many seconds should blocks take? (default = 15)")
		genesis.Config.Clique.Period = uint64(w.readDefaultInt(15))

		// We also need the initial list of signers
		fmt.Println()
		fmt.Println("Which accounts are allowed to seal? (mandatory at least one)")

		var signers []common.Address
		for {
			if address := w.readAddress(); address != nil {
				signers = append(signers, *address)
				continue
			}
			if len(signers) > 0 {
				break
			}
		}
		// Sort the signers and embed into the extra-data section
		for i := 0; i < len(signers); i++ {
			for j := i + 1; j < len(signers); j++ {
				if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
					signers[i], signers[j] = signers[j], signers[i]
				}
			}
		}
		genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
		for i, signer := range signers {
			copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
		}
	case choice == "3":
		// In the case of posv, configure the consensus parameters
		genesis.Difficulty = big.NewInt(1)
		base := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		genesis.Config.Posv = &params.PoSVConfig{
			Period:    2,
			Epoch:     900,
			MinStaked: new(big.Int).Mul(big.NewInt(50000), base),
		}
		fmt.Println()
		fmt.Println("How many seconds should blocks take? (default = 2)")
		genesis.Config.Posv.Period = uint64(w.readDefaultInt(2))

		fmt.Println()
		fmt.Println("How many min tokens should be staked? (default = 50000)")
		genesis.Config.Posv.MinStaked = new(big.Int).Mul(big.NewInt(int64(w.readDefaultInt(50000))), base)

		// We also need the initial list of signers
		fmt.Println()
		fmt.Println("Which accounts are allowed to seal? (mandatory at least one)")

		var signers posv.MasterNodes
		for {
			if address := w.readAddress(); address != nil {
				signers = append(signers, posv.MasterNode{
					Address: *address,
					Stake:   genesis.Config.Posv.MinStaked,
				})
				continue
			}
			if len(signers) > 0 {
				break
			}
		}
		// Sort the signers and embed into the extra-data section
		//for i := 0; i < len(signers); i++ {
		//	for j := i + 1; j < len(signers); j++ {
		//		if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
		//			signers[i], signers[j] = signers[j], signers[i]
		//		}
		//	}
		//}
		sort.Sort(signers)
		genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
		for i, signer := range signers {
			copy(genesis.ExtraData[32+i*common.AddressLength:], signer.Address[:])
		}
		if err := deployValidatorContract(signers, genesis.Config.Posv.MinStaked, genesis.Alloc); err != nil {
			log.Crit("Error on deployValidatorContract", "err", err)
		}
		if err := deployBlockSignerContract(genesis.Config.Posv.Epoch, genesis.Alloc); err != nil {
			log.Crit("Error on deployBlockSignerContract", "err", err)
		}
		if err := deployRandomizeContract(genesis.Alloc); err != nil {
			log.Crit("Error on deployRandomizeContract", "err", err)
		}
	default:
		log.Crit("Invalid consensus engine choice", "choice", choice)
	}
	// Consensus all set, just ask for initial funds and go
	fmt.Println()
	fmt.Println("Which accounts should be pre-funded? (advisable at least one)")
	for {
		// Read the address of the account to fund
		if address := w.readAddress(); address != nil {
			genesis.Alloc[*address] = core.GenesisAccount{
				Balance: new(big.Int).Lsh(big.NewInt(1), 256-7), // 2^256 / 128 (allow many pre-funds without balance overflows)
			}
			continue
		}
		break
	}
	fmt.Println()
	fmt.Println("Should the precompile-addresses (0x1 .. 0xff) be pre-funded with 1 wei? (advisable yes)")
	if w.readDefaultYesNo(true) {
		// Add a batch of precompile balances to avoid them getting deleted
		for i := int64(0); i < 256; i++ {
			genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
		}
	}
	// Query the user for some custom extras
	fmt.Println()
	fmt.Println("Specify your chain/network ID if you want an explicit one (default = random)")
	genesis.Config.ChainID = new(big.Int).SetUint64(uint64(w.readDefaultInt(rand.Intn(65536))))

	// All done, store the genesis and flush to disk
	log.Info("Configured new genesis block")

	w.conf.Genesis = genesis
	w.conf.flush()
}

// importGenesis imports a Geth genesis spec into puppeth.
func (w *wizard) importGenesis() {
	// Request the genesis JSON spec URL from the user
	fmt.Println()
	fmt.Println("Where's the genesis file? (local file or http/https url)")
	url := w.readURL()

	// Convert the various allowed URLs to a reader stream
	var reader io.Reader

	switch url.Scheme {
	case "http", "https":
		// Remote web URL, retrieve it via an HTTP client
		res, err := http.Get(url.String())
		if err != nil {
			log.Error("Failed to retrieve remote genesis", "err", err)
			return
		}
		defer res.Body.Close()
		reader = res.Body

	case "":
		// Schemaless URL, interpret as a local file
		file, err := os.Open(url.String())
		if err != nil {
			log.Error("Failed to open local genesis", "err", err)
			return
		}
		defer file.Close()
		reader = file

	default:
		log.Error("Unsupported genesis URL scheme", "scheme", url.Scheme)
		return
	}
	// Parse the genesis file and inject it successful
	var genesis core.Genesis
	if err := json.NewDecoder(reader).Decode(&genesis); err != nil {
		log.Error("Invalid genesis spec", "err", err)
		return
	}
	log.Info("Imported genesis block")

	w.conf.Genesis = &genesis
	w.conf.flush()
}

// manageGenesis permits the modification of chain configuration parameters in
// a genesis config and the export of the entire genesis spec.
func (w *wizard) manageGenesis() {
	// Figure out whether to modify or export the genesis
	fmt.Println()
	fmt.Println(" 1. Modify existing configurations")
	fmt.Println(" 2. Export genesis configurations")
	fmt.Println(" 3. Remove genesis configuration")

	choice := w.read()
	switch choice {
	case "1":
		// Fork rule updating requested, iterate over each fork
		fmt.Println()
		fmt.Printf("Which block should Homestead come into effect? (default = %v)\n", w.conf.Genesis.Config.HomesteadBlock)
		w.conf.Genesis.Config.HomesteadBlock = w.readDefaultBigInt(w.conf.Genesis.Config.HomesteadBlock)

		fmt.Println()
		fmt.Printf("Which block should EIP150 (Tangerine Whistle) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP150Block)
		w.conf.Genesis.Config.EIP150Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP150Block)

		fmt.Println()
		fmt.Printf("Which block should EIP155 (Spurious Dragon) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP155Block)
		w.conf.Genesis.Config.EIP155Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP155Block)

		fmt.Println()
		fmt.Printf("Which block should EIP158/161 (also Spurious Dragon) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP158Block)
		w.conf.Genesis.Config.EIP158Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP158Block)

		fmt.Println()
		fmt.Printf("Which block should Byzantium come into effect? (default = %v)\n", w.conf.Genesis.Config.ByzantiumBlock)
		w.conf.Genesis.Config.ByzantiumBlock = w.readDefaultBigInt(w.conf.Genesis.Config.ByzantiumBlock)

		fmt.Println()
		fmt.Printf("Which block should Constantinople come into effect? (default = %v)\n", w.conf.Genesis.Config.ConstantinopleBlock)
		w.conf.Genesis.Config.ConstantinopleBlock = w.readDefaultBigInt(w.conf.Genesis.Config.ConstantinopleBlock)
		if w.conf.Genesis.Config.PetersburgBlock == nil {
			w.conf.Genesis.Config.PetersburgBlock = w.conf.Genesis.Config.ConstantinopleBlock
		}
		fmt.Println()
		fmt.Printf("Which block should Petersburg come into effect? (default = %v)\n", w.conf.Genesis.Config.PetersburgBlock)
		w.conf.Genesis.Config.PetersburgBlock = w.readDefaultBigInt(w.conf.Genesis.Config.PetersburgBlock)

		fmt.Println()
		fmt.Printf("Which block should Istanbul come into effect? (default = %v)\n", w.conf.Genesis.Config.IstanbulBlock)
		w.conf.Genesis.Config.IstanbulBlock = w.readDefaultBigInt(w.conf.Genesis.Config.IstanbulBlock)

		fmt.Println()
		fmt.Printf("Which block should Berlin come into effect? (default = %v)\n", w.conf.Genesis.Config.BerlinBlock)
		w.conf.Genesis.Config.BerlinBlock = w.readDefaultBigInt(w.conf.Genesis.Config.BerlinBlock)

		fmt.Println()
		fmt.Printf("Which block should London come into effect? (default = %v)\n", w.conf.Genesis.Config.LondonBlock)
		w.conf.Genesis.Config.LondonBlock = w.readDefaultBigInt(w.conf.Genesis.Config.LondonBlock)

		out, _ := json.MarshalIndent(w.conf.Genesis.Config, "", "  ")
		fmt.Printf("Chain configuration updated:\n\n%s\n", out)

		w.conf.flush()

	case "2":
		// Save whatever genesis configuration we currently have
		fmt.Println()
		fmt.Printf("Which folder to save the genesis spec into? (default = current)\n")
		fmt.Printf("  Will create %s.json\n", w.network)

		folder := w.readDefaultString(".")
		if err := os.MkdirAll(folder, 0755); err != nil {
			log.Error("Failed to create spec folder", "folder", folder, "err", err)
			return
		}
		out, _ := json.MarshalIndent(w.conf.Genesis, "", "  ")

		// Export the native genesis spec used by puppeth and Geth
		gethJson := filepath.Join(folder, fmt.Sprintf("%s.json", w.network))
		if err := os.WriteFile(gethJson, out, 0644); err != nil {
			log.Error("Failed to save genesis file", "err", err)
			return
		}
		log.Info("Saved native genesis chain spec", "path", gethJson)

	case "3":
		// Make sure we don't have any services running
		if len(w.conf.servers()) > 0 {
			log.Error("Genesis reset requires all services and servers torn down")
			return
		}
		log.Info("Genesis block destroyed")

		w.conf.Genesis = nil
		w.conf.flush()
	default:
		log.Error("That's not something I can do")
		return
	}
}

func deployValidatorContract(masters posv.MasterNodes, stakeCap *big.Int, genesisAlloc core.GenesisAlloc) error {
	if len(masters) == 0 {
		return nil
	}
	pKey, _ := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr := crypto.PubkeyToAddress(pKey.PublicKey)
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000000)}}, 10000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(pKey, big.NewInt(1337))
	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	var validatorCaps []*big.Int
	var signers []common.Address
	for i := 0; i < len(masters); i++ {
		validatorCaps = append(validatorCaps, stakeCap)
		signers = append(signers, masters[i].Address)
	}
	validatorAddress, _, err := validatorContract.DeployValidator(transactOpts, contractBackend, signers, validatorCaps, signers[0])
	if err != nil {
		log.Error("Can't deploy root registry", "err", err)
		return err
	}
	contractBackend.Commit()

	storage := make(map[common.Hash]common.Hash)
	f := func(key, val common.Hash) bool {
		decode := []byte{}
		trim := bytes.TrimLeft(val.Bytes(), "\x00")
		_ = rlp.DecodeBytes(trim, &decode)
		storage[key] = common.BytesToHash(decode)
		log.Info("DecodeBytes", "value", val.String(), "decode", storage[key].String())
		return true
	}

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, _ := contractBackend.CodeAt(ctx, validatorAddress, nil)

	head, _ = contractBackend.HeaderByNumber(context.Background(), nil)
	stateDB, err := contractBackend.Blockchain().StateAt(head.Root)
	if err != nil {
		log.Error("can't get stateDB", "root", head.Root.Hex(), "err", err)
		return err
	}
	if err := stateDB.ForEachStorage(validatorAddress, f); err != nil {
		log.Error("can't foreach validator", "validatorAddress", validatorAddress.Hex())
		return err
	}

	genesisAlloc[common.HexToAddress(params.JuncaValidator)] = core.GenesisAccount{
		Balance: stakeCap.Mul(stakeCap, big.NewInt(int64(len(validatorCaps)))),
		Code:    code,
		Storage: storage,
	}
	return nil
}

func deployBlockSignerContract(epochNumber uint64, genesisAlloc core.GenesisAlloc) error {
	pKey, _ := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr := crypto.PubkeyToAddress(pKey.PublicKey)
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000000)}}, 10000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(pKey, big.NewInt(1337))
	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	blockSignerAddress, _, err := blockSignerContract.DeployBlockSigner(transactOpts, contractBackend, big.NewInt(int64(epochNumber)))
	if err != nil {
		log.Error("Can't deploy root registry", "err", err)
		return err
	}
	contractBackend.Commit()

	storage := make(map[common.Hash]common.Hash)
	f := func(key, val common.Hash) bool {
		decode := []byte{}
		trim := bytes.TrimLeft(val.Bytes(), "\x00")
		_ = rlp.DecodeBytes(trim, &decode)
		storage[key] = common.BytesToHash(decode)
		log.Info("DecodeBytes", "value", val.String(), "decode", storage[key].String())
		return true
	}

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, _ := contractBackend.CodeAt(ctx, blockSignerAddress, nil)

	head, _ = contractBackend.HeaderByNumber(context.Background(), nil)
	stateDB, err := contractBackend.Blockchain().StateAt(head.Root)
	if err != nil {
		log.Error("can't get stateDB", "root", head.Root.Hex(), "err", err)
		return err
	}
	if err := stateDB.ForEachStorage(blockSignerAddress, f); err != nil {
		log.Error("can't foreach blockSigner", "blockSignerAddress", blockSignerAddress.Hex())
		return err
	}

	genesisAlloc[common.HexToAddress(params.JuncaBlockSigner)] = core.GenesisAccount{
		Balance: big.NewInt(0),
		Code:    code,
		Storage: storage,
	}
	return nil
}

func deployRandomizeContract(genesisAlloc core.GenesisAlloc) error {
	pKey, _ := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr := crypto.PubkeyToAddress(pKey.PublicKey)
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000000)}}, 10000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(pKey, big.NewInt(1337))
	head, _ := contractBackend.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))
	transactOpts.GasPrice = gasPrice

	randomizeAddress, _, err := randomizeContract.DeployRandomize(transactOpts, contractBackend)
	if err != nil {
		log.Error("Can't deploy root registry", "err", err)
		return err
	}
	contractBackend.Commit()

	storage := make(map[common.Hash]common.Hash)
	f := func(key, val common.Hash) bool {
		decode := []byte{}
		trim := bytes.TrimLeft(val.Bytes(), "\x00")
		_ = rlp.DecodeBytes(trim, &decode)
		storage[key] = common.BytesToHash(decode)
		log.Info("DecodeBytes", "value", val.String(), "decode", storage[key].String())
		return true
	}

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, _ := contractBackend.CodeAt(ctx, randomizeAddress, nil)

	head, _ = contractBackend.HeaderByNumber(context.Background(), nil)
	stateDB, err := contractBackend.Blockchain().StateAt(head.Root)
	if err != nil {
		log.Error("can't get stateDB", "root", head.Root.Hex(), "err", err)
		return err
	}
	if err := stateDB.ForEachStorage(randomizeAddress, f); err != nil {
		log.Error("can't foreach randomize", "randomizeAddress", randomizeAddress.Hex())
		return err
	}

	genesisAlloc[common.HexToAddress(params.JuncaRandomize)] = core.GenesisAccount{
		Balance: big.NewInt(0),
		Code:    code,
		Storage: storage,
	}
	return nil
}
