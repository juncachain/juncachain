// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/juncachain/juncachain"
	"github.com/juncachain/juncachain/accounts/abi"
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/core/types"
	"github.com/juncachain/juncachain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlockSignerMetaData contains all meta data concerning the BlockSigner contract.
var BlockSignerMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockSigners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"getSigners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epochNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"Sign\",\"type\":\"event\"}]",
	Bin: "0x6060604052341561000f57600080fd5b604051602080610469833981016040528080516002555050610433806100366000396000f30060606040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663046a42908114610071578063e341eaa4146100b3578063e7ec6aef146100ce578063f4145a8314610137578063f4f911db1461015c575b600080fd5b341561007c57600080fd5b61008a600435602435610175565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100be57600080fd5b6100cc6004356024356101b9565b005b34156100d957600080fd5b6100e46004356102dc565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561012357808201518382015260200161010b565b505050509050019250505060405180910390f35b341561014257600080fd5b61014a610361565b60405190815260200160405180910390f35b341561016757600080fd5b61014a600435602435610367565b60006020528160005260406000208181548110151561019057fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b43829010156101c757600080fd5b600280546101dd9184910263ffffffff61039516565b4311156101e957600080fd5b600082815260016020819052604090912080549091810161020a83826103ab565b5060009182526020808320919091018390558282528190526040902080546001810161023683826103ab565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19163373ffffffffffffffffffffffffffffffffffffffff8116919091179091557f62855fa22e051687c32ac285857751f6d3f2c100c72756d8d30cb7ecb1f64f5490838360405173ffffffffffffffffffffffffffffffffffffffff909316835260208301919091526040808301919091526060909101905180910390a15050565b6102e46103d4565b6000828152602081815260409182902080549092909182810201905190810160405280929190818152602001828054801561035557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032a575b50505050509050919050565b60025481565b60016020528160005260406000208181548110151561038257fe5b6000918252602090912001549150829050565b6000828201838110156103a457fe5b9392505050565b8154818355818115116103cf576000838152602090206103cf9181019083016103e6565b505050565b60206040519081016040526000815290565b61040491905b8082111561040057600081556001016103ec565b5090565b905600a165627a7a723058209138a3c8bb763d1f616d5d2ddd160d0f24c37bf277edd5bc358e277394e5f94e0029",
}

// BlockSignerABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockSignerMetaData.ABI instead.
var BlockSignerABI = BlockSignerMetaData.ABI

// BlockSignerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockSignerMetaData.Bin instead.
var BlockSignerBin = BlockSignerMetaData.Bin

// DeployBlockSigner deploys a new Ethereum contract, binding an instance of BlockSigner to it.
func DeployBlockSigner(auth *bind.TransactOpts, backend bind.ContractBackend, _epochNumber *big.Int) (common.Address, *types.Transaction, *BlockSigner, error) {
	parsed, err := BlockSignerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockSignerBin), backend, _epochNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockSigner{BlockSignerCaller: BlockSignerCaller{contract: contract}, BlockSignerTransactor: BlockSignerTransactor{contract: contract}, BlockSignerFilterer: BlockSignerFilterer{contract: contract}}, nil
}

// BlockSigner is an auto generated Go binding around an Ethereum contract.
type BlockSigner struct {
	BlockSignerCaller     // Read-only binding to the contract
	BlockSignerTransactor // Write-only binding to the contract
	BlockSignerFilterer   // Log filterer for contract events
}

// BlockSignerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockSignerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSignerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockSignerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSignerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockSignerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSignerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockSignerSession struct {
	Contract     *BlockSigner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockSignerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockSignerCallerSession struct {
	Contract *BlockSignerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BlockSignerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockSignerTransactorSession struct {
	Contract     *BlockSignerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BlockSignerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockSignerRaw struct {
	Contract *BlockSigner // Generic contract binding to access the raw methods on
}

// BlockSignerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockSignerCallerRaw struct {
	Contract *BlockSignerCaller // Generic read-only contract binding to access the raw methods on
}

// BlockSignerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockSignerTransactorRaw struct {
	Contract *BlockSignerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockSigner creates a new instance of BlockSigner, bound to a specific deployed contract.
func NewBlockSigner(address common.Address, backend bind.ContractBackend) (*BlockSigner, error) {
	contract, err := bindBlockSigner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockSigner{BlockSignerCaller: BlockSignerCaller{contract: contract}, BlockSignerTransactor: BlockSignerTransactor{contract: contract}, BlockSignerFilterer: BlockSignerFilterer{contract: contract}}, nil
}

// NewBlockSignerCaller creates a new read-only instance of BlockSigner, bound to a specific deployed contract.
func NewBlockSignerCaller(address common.Address, caller bind.ContractCaller) (*BlockSignerCaller, error) {
	contract, err := bindBlockSigner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockSignerCaller{contract: contract}, nil
}

// NewBlockSignerTransactor creates a new write-only instance of BlockSigner, bound to a specific deployed contract.
func NewBlockSignerTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockSignerTransactor, error) {
	contract, err := bindBlockSigner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockSignerTransactor{contract: contract}, nil
}

// NewBlockSignerFilterer creates a new log filterer instance of BlockSigner, bound to a specific deployed contract.
func NewBlockSignerFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockSignerFilterer, error) {
	contract, err := bindBlockSigner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockSignerFilterer{contract: contract}, nil
}

// bindBlockSigner binds a generic wrapper to an already deployed contract.
func bindBlockSigner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockSignerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockSigner *BlockSignerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockSigner.Contract.BlockSignerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockSigner *BlockSignerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockSigner.Contract.BlockSignerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockSigner *BlockSignerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockSigner.Contract.BlockSignerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockSigner *BlockSignerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockSigner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockSigner *BlockSignerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockSigner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockSigner *BlockSignerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockSigner.Contract.contract.Transact(opts, method, params...)
}

// BlockSigners is a free data retrieval call binding the contract method 0x046a4290.
//
// Solidity: function blockSigners(bytes32 , uint256 ) view returns(address)
func (_BlockSigner *BlockSignerCaller) BlockSigners(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BlockSigner.contract.Call(opts, &out, "blockSigners", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockSigners is a free data retrieval call binding the contract method 0x046a4290.
//
// Solidity: function blockSigners(bytes32 , uint256 ) view returns(address)
func (_BlockSigner *BlockSignerSession) BlockSigners(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _BlockSigner.Contract.BlockSigners(&_BlockSigner.CallOpts, arg0, arg1)
}

// BlockSigners is a free data retrieval call binding the contract method 0x046a4290.
//
// Solidity: function blockSigners(bytes32 , uint256 ) view returns(address)
func (_BlockSigner *BlockSignerCallerSession) BlockSigners(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _BlockSigner.Contract.BlockSigners(&_BlockSigner.CallOpts, arg0, arg1)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks(uint256 , uint256 ) view returns(bytes32)
func (_BlockSigner *BlockSignerCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockSigner.contract.Call(opts, &out, "blocks", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks(uint256 , uint256 ) view returns(bytes32)
func (_BlockSigner *BlockSignerSession) Blocks(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _BlockSigner.Contract.Blocks(&_BlockSigner.CallOpts, arg0, arg1)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks(uint256 , uint256 ) view returns(bytes32)
func (_BlockSigner *BlockSignerCallerSession) Blocks(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _BlockSigner.Contract.Blocks(&_BlockSigner.CallOpts, arg0, arg1)
}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_BlockSigner *BlockSignerCaller) EpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockSigner.contract.Call(opts, &out, "epochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_BlockSigner *BlockSignerSession) EpochNumber() (*big.Int, error) {
	return _BlockSigner.Contract.EpochNumber(&_BlockSigner.CallOpts)
}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_BlockSigner *BlockSignerCallerSession) EpochNumber() (*big.Int, error) {
	return _BlockSigner.Contract.EpochNumber(&_BlockSigner.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0xe7ec6aef.
//
// Solidity: function getSigners(bytes32 _blockHash) view returns(address[])
func (_BlockSigner *BlockSignerCaller) GetSigners(opts *bind.CallOpts, _blockHash [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BlockSigner.contract.Call(opts, &out, "getSigners", _blockHash)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0xe7ec6aef.
//
// Solidity: function getSigners(bytes32 _blockHash) view returns(address[])
func (_BlockSigner *BlockSignerSession) GetSigners(_blockHash [32]byte) ([]common.Address, error) {
	return _BlockSigner.Contract.GetSigners(&_BlockSigner.CallOpts, _blockHash)
}

// GetSigners is a free data retrieval call binding the contract method 0xe7ec6aef.
//
// Solidity: function getSigners(bytes32 _blockHash) view returns(address[])
func (_BlockSigner *BlockSignerCallerSession) GetSigners(_blockHash [32]byte) ([]common.Address, error) {
	return _BlockSigner.Contract.GetSigners(&_BlockSigner.CallOpts, _blockHash)
}

// Sign is a paid mutator transaction binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_BlockSigner *BlockSignerTransactor) Sign(opts *bind.TransactOpts, _blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _BlockSigner.contract.Transact(opts, "sign", _blockNumber, _blockHash)
}

// Sign is a paid mutator transaction binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_BlockSigner *BlockSignerSession) Sign(_blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _BlockSigner.Contract.Sign(&_BlockSigner.TransactOpts, _blockNumber, _blockHash)
}

// Sign is a paid mutator transaction binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_BlockSigner *BlockSignerTransactorSession) Sign(_blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _BlockSigner.Contract.Sign(&_BlockSigner.TransactOpts, _blockNumber, _blockHash)
}

// BlockSignerSignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the BlockSigner contract.
type BlockSignerSignIterator struct {
	Event *BlockSignerSign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockSignerSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockSignerSign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockSignerSign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockSignerSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockSignerSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockSignerSign represents a Sign event raised by the BlockSigner contract.
type BlockSignerSign struct {
	Signer      common.Address
	BlockNumber *big.Int
	BlockHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0x62855fa22e051687c32ac285857751f6d3f2c100c72756d8d30cb7ecb1f64f54.
//
// Solidity: event Sign(address _signer, uint256 _blockNumber, bytes32 _blockHash)
func (_BlockSigner *BlockSignerFilterer) FilterSign(opts *bind.FilterOpts) (*BlockSignerSignIterator, error) {

	logs, sub, err := _BlockSigner.contract.FilterLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return &BlockSignerSignIterator{contract: _BlockSigner.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0x62855fa22e051687c32ac285857751f6d3f2c100c72756d8d30cb7ecb1f64f54.
//
// Solidity: event Sign(address _signer, uint256 _blockNumber, bytes32 _blockHash)
func (_BlockSigner *BlockSignerFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *BlockSignerSign) (event.Subscription, error) {

	logs, sub, err := _BlockSigner.contract.WatchLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockSignerSign)
				if err := _BlockSigner.contract.UnpackLog(event, "Sign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSign is a log parse operation binding the contract event 0x62855fa22e051687c32ac285857751f6d3f2c100c72756d8d30cb7ecb1f64f54.
//
// Solidity: event Sign(address _signer, uint256 _blockNumber, bytes32 _blockHash)
func (_BlockSigner *BlockSignerFilterer) ParseSign(log types.Log) (*BlockSignerSign, error) {
	event := new(BlockSignerSign)
	if err := _BlockSigner.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
