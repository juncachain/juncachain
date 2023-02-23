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
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"setSecret\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getOpening\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_opening\",\"type\":\"bytes32\"}],\"name\":\"setOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x6060604052341561000f57600080fd5b6102fe8061001e6000396000f3006060604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663284180fc811461006657806334d38600146100d8578063d442d6cc14610129578063e11f5ba21461015a575b600080fd5b341561007157600080fd5b610085600160a060020a0360043516610170565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100c45780820151838201526020016100ac565b505050509050019250505060405180910390f35b34156100e357600080fd5b61012760046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506101f395505050505050565b005b341561013457600080fd5b610148600160a060020a0360043516610220565b60405190815260200160405180910390f35b341561016557600080fd5b61012760043561023b565b610178610256565b60008083600160a060020a0316600160a060020a031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156101e757602002820191906000526020600020905b815481526001909101906020018083116101d2575b50505050509050919050565b600160a060020a033316600090815260208190526040902081805161021c929160200190610268565b5050565b600160a060020a031660009081526001602052604090205490565b600160a060020a033316600090815260016020526040902055565b60206040519081016040526000815290565b8280548282559060005260206000209081019282156102a5579160200282015b828111156102a55782518255602090920191600190910190610288565b506102b19291506102b5565b5090565b6102cf91905b808211156102b157600081556001016102bb565b905600a165627a7a72305820ff08e3dd306701ce4608e351d3672afca3f7fd7f9eada549e55c2eff1797e9800029",
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
