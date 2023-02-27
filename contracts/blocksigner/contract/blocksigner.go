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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516103fd3803806103fd83398101604081905261002f91610037565b600255610050565b60006020828403121561004957600080fd5b5051919050565b61039e8061005f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063046a42901461005c578063e341eaa41461008c578063e7ec6aef146100a1578063f4145a83146100c1578063f4f911db146100d8575b600080fd5b61006f61006a3660046102ac565b6100eb565b6040516001600160a01b0390911681526020015b60405180910390f35b61009f61009a3660046102ac565b610123565b005b6100b46100af366004610293565b6101e5565b60405161008391906102ce565b6100ca60025481565b604051908152602001610083565b6100ca6100e63660046102ac565b61024f565b6000602052816000526040600020818154811061010757600080fd5b6000918252602090912001546001600160a01b03169150829050565b8143101561013057600080fd5b61014960025460026101429190610333565b8390610280565b43111561015557600080fd5b600082815260016020818152604080842080548085018255908552828520018590558484528382528084208054938401815584529281902090910180546001600160a01b0319163390811790915582519081529081018490529081018290527f62855fa22e051687c32ac285857751f6d3f2c100c72756d8d30cb7ecb1f64f549060600160405180910390a15050565b6000818152602081815260409182902080548351818402810184019094528084526060939283018282801561024357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610225575b50505050509050919050565b6001602052816000526040600020818154811061026b57600080fd5b90600052602060002001600091509150505481565b600061028c828461031b565b9392505050565b6000602082840312156102a557600080fd5b5035919050565b600080604083850312156102bf57600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b8181101561030f5783516001600160a01b0316835292840192918401916001016102ea565b50909695505050505050565b6000821982111561032e5761032e610352565b500190565b600081600019048311821515161561034d5761034d610352565b500290565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220b8fde478a1077d4acee8fe83760962497fb1c2774a77ec02302d31a9a6c1e43a64736f6c63430008070033",
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
