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
	Bin: "0x6060604052341561000f57600080fd5b6104658061001e6000396000f300606060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063284180fc1461006757806334d38600146100f5578063d442d6cc1461014f578063e11f5ba2146101a4575b600080fd5b341561007257600080fd5b61009e600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506101cb565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100e15780820151818401526020810190506100c6565b505050509050019250505060405180910390f35b341561010057600080fd5b61014d60048080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190505061026b565b005b341561015a57600080fd5b610186600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506102f5565b60405180826000191660001916815260200191505060405180910390f35b34156101af57600080fd5b6101c960048080356000191690602001909190505061033e565b005b6101d36103ad565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561025f57602002820191906000526020600020905b81546000191681526020019060010190808311610247575b50505050509050919050565b60006103844381151561027a57fe5b069050610320811015151561028e57600080fd5b6103528110151561029e57600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906102f09291906103c1565b505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006103844381151561034d57fe5b069050610352811015151561036157600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081600019169055505050565b602060405190810160405280600081525090565b828054828255906000526020600020908101928215610403579160200282015b828111156104025782518290600019169055916020019190600101906103e1565b5b5090506104109190610414565b5090565b61043691905b8082111561043257600081600090555060010161041a565b5090565b905600a165627a7a72305820ce55dbcb900884ce63baa508bb06e3a909d95cc0826e9d36752e69fb908b3df10029",
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
