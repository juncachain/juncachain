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

// JuncaRandomizeMetaData contains all meta data concerning the JuncaRandomize contract.
var JuncaRandomizeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getOpening\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getSecret\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opening\",\"type\":\"bytes32\"}],\"name\":\"setOpening\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"setSecret\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610370806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063284180fc1461005157806334d386001461007a578063d442d6cc1461008f578063e11f5ba2146100c6575b600080fd5b61006461005f3660046101d2565b6100e6565b60405161007191906102e0565b60405180910390f35b61008d610088366004610202565b610150565b005b6100b861009d3660046101d2565b6001600160a01b031660009081526001602052604090205490565b604051908152602001610071565b61008d6100d43660046102c7565b33600090815260016020526040902055565b6001600160a01b0381166000908152602081815260409182902080548351818402810184019094528084526060939283018282801561014457602002820191906000526020600020905b815481526020019060010190808311610130575b50505050509050919050565b33600090815260208181526040909120825161016e92840190610172565b5050565b8280548282559060005260206000209081019282156101ad579160200282015b828111156101ad578251825591602001919060010190610192565b506101b99291506101bd565b5090565b5b808211156101b957600081556001016101be565b6000602082840312156101e457600080fd5b81356001600160a01b03811681146101fb57600080fd5b9392505050565b6000602080838503121561021557600080fd5b823567ffffffffffffffff8082111561022d57600080fd5b818501915085601f83011261024157600080fd5b81358181111561025357610253610324565b8060051b604051601f19603f8301168101818110858211171561027857610278610324565b604052828152858101935084860182860187018a101561029757600080fd5b600095505b838610156102ba57803585526001959095019493860193860161029c565b5098975050505050505050565b6000602082840312156102d957600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015610318578351835292840192918401916001016102fc565b50909695505050505050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220c9cde02a1b9524fcb4937eed1414d15ef8713905c22e66789511e8430545211164736f6c63430008070033",
}

// JuncaRandomizeABI is the input ABI used to generate the binding from.
// Deprecated: Use JuncaRandomizeMetaData.ABI instead.
var JuncaRandomizeABI = JuncaRandomizeMetaData.ABI

// JuncaRandomizeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JuncaRandomizeMetaData.Bin instead.
var JuncaRandomizeBin = JuncaRandomizeMetaData.Bin

// DeployJuncaRandomize deploys a new Ethereum contract, binding an instance of JuncaRandomize to it.
func DeployJuncaRandomize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JuncaRandomize, error) {
	parsed, err := JuncaRandomizeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JuncaRandomizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JuncaRandomize{JuncaRandomizeCaller: JuncaRandomizeCaller{contract: contract}, JuncaRandomizeTransactor: JuncaRandomizeTransactor{contract: contract}, JuncaRandomizeFilterer: JuncaRandomizeFilterer{contract: contract}}, nil
}

// JuncaRandomize is an auto generated Go binding around an Ethereum contract.
type JuncaRandomize struct {
	JuncaRandomizeCaller     // Read-only binding to the contract
	JuncaRandomizeTransactor // Write-only binding to the contract
	JuncaRandomizeFilterer   // Log filterer for contract events
}

// JuncaRandomizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type JuncaRandomizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaRandomizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JuncaRandomizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaRandomizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JuncaRandomizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaRandomizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JuncaRandomizeSession struct {
	Contract     *JuncaRandomize   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JuncaRandomizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JuncaRandomizeCallerSession struct {
	Contract *JuncaRandomizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// JuncaRandomizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JuncaRandomizeTransactorSession struct {
	Contract     *JuncaRandomizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// JuncaRandomizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type JuncaRandomizeRaw struct {
	Contract *JuncaRandomize // Generic contract binding to access the raw methods on
}

// JuncaRandomizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JuncaRandomizeCallerRaw struct {
	Contract *JuncaRandomizeCaller // Generic read-only contract binding to access the raw methods on
}

// JuncaRandomizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JuncaRandomizeTransactorRaw struct {
	Contract *JuncaRandomizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJuncaRandomize creates a new instance of JuncaRandomize, bound to a specific deployed contract.
func NewJuncaRandomize(address common.Address, backend bind.ContractBackend) (*JuncaRandomize, error) {
	contract, err := bindJuncaRandomize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JuncaRandomize{JuncaRandomizeCaller: JuncaRandomizeCaller{contract: contract}, JuncaRandomizeTransactor: JuncaRandomizeTransactor{contract: contract}, JuncaRandomizeFilterer: JuncaRandomizeFilterer{contract: contract}}, nil
}

// NewJuncaRandomizeCaller creates a new read-only instance of JuncaRandomize, bound to a specific deployed contract.
func NewJuncaRandomizeCaller(address common.Address, caller bind.ContractCaller) (*JuncaRandomizeCaller, error) {
	contract, err := bindJuncaRandomize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JuncaRandomizeCaller{contract: contract}, nil
}

// NewJuncaRandomizeTransactor creates a new write-only instance of JuncaRandomize, bound to a specific deployed contract.
func NewJuncaRandomizeTransactor(address common.Address, transactor bind.ContractTransactor) (*JuncaRandomizeTransactor, error) {
	contract, err := bindJuncaRandomize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JuncaRandomizeTransactor{contract: contract}, nil
}

// NewJuncaRandomizeFilterer creates a new log filterer instance of JuncaRandomize, bound to a specific deployed contract.
func NewJuncaRandomizeFilterer(address common.Address, filterer bind.ContractFilterer) (*JuncaRandomizeFilterer, error) {
	contract, err := bindJuncaRandomize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JuncaRandomizeFilterer{contract: contract}, nil
}

// bindJuncaRandomize binds a generic wrapper to an already deployed contract.
func bindJuncaRandomize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JuncaRandomizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuncaRandomize *JuncaRandomizeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuncaRandomize.Contract.JuncaRandomizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuncaRandomize *JuncaRandomizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.JuncaRandomizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuncaRandomize *JuncaRandomizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.JuncaRandomizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuncaRandomize *JuncaRandomizeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuncaRandomize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuncaRandomize *JuncaRandomizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuncaRandomize *JuncaRandomizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.contract.Transact(opts, method, params...)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(address _validator) view returns(bytes32)
func (_JuncaRandomize *JuncaRandomizeCaller) GetOpening(opts *bind.CallOpts, _validator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _JuncaRandomize.contract.Call(opts, &out, "getOpening", _validator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(address _validator) view returns(bytes32)
func (_JuncaRandomize *JuncaRandomizeSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _JuncaRandomize.Contract.GetOpening(&_JuncaRandomize.CallOpts, _validator)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(address _validator) view returns(bytes32)
func (_JuncaRandomize *JuncaRandomizeCallerSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _JuncaRandomize.Contract.GetOpening(&_JuncaRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(address _validator) view returns(bytes32[])
func (_JuncaRandomize *JuncaRandomizeCaller) GetSecret(opts *bind.CallOpts, _validator common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _JuncaRandomize.contract.Call(opts, &out, "getSecret", _validator)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(address _validator) view returns(bytes32[])
func (_JuncaRandomize *JuncaRandomizeSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _JuncaRandomize.Contract.GetSecret(&_JuncaRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(address _validator) view returns(bytes32[])
func (_JuncaRandomize *JuncaRandomizeCallerSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _JuncaRandomize.Contract.GetSecret(&_JuncaRandomize.CallOpts, _validator)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(bytes32 _opening) returns()
func (_JuncaRandomize *JuncaRandomizeTransactor) SetOpening(opts *bind.TransactOpts, _opening [32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.contract.Transact(opts, "setOpening", _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(bytes32 _opening) returns()
func (_JuncaRandomize *JuncaRandomizeSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.SetOpening(&_JuncaRandomize.TransactOpts, _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(bytes32 _opening) returns()
func (_JuncaRandomize *JuncaRandomizeTransactorSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.SetOpening(&_JuncaRandomize.TransactOpts, _opening)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(bytes32[] _secret) returns()
func (_JuncaRandomize *JuncaRandomizeTransactor) SetSecret(opts *bind.TransactOpts, _secret [][32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.contract.Transact(opts, "setSecret", _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(bytes32[] _secret) returns()
func (_JuncaRandomize *JuncaRandomizeSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.SetSecret(&_JuncaRandomize.TransactOpts, _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(bytes32[] _secret) returns()
func (_JuncaRandomize *JuncaRandomizeTransactorSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _JuncaRandomize.Contract.SetSecret(&_JuncaRandomize.TransactOpts, _secret)
}
