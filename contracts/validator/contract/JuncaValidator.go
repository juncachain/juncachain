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

// JuncaValidatorMetaData contains all meta data concerning the JuncaValidator contract.
var JuncaValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_candidates\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_firstOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260006004553480156200001657600080fd5b50604051620015c5380380620015c58339810160408190526200003991620003e4565b60058590556006849055600783905560088290556009819055875160045560005b88518110156200033e5760038982815181106200007b576200007b62000583565b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b039092169190911790558a518992908c9085908110620000d057620000d062000583565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060018060008b848151811062000139576200013962000583565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160146101000a81548160ff02191690831515021790555087818151811062000190576200019062000583565b6020026020010151600160008b8481518110620001b157620001b162000583565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010181905550600554600160008b8481518110620001fc57620001fc62000583565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206002016000896001600160a01b03166001600160a01b0316815260200190815260200160002081905550600260008a838151811062000267576200026762000583565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600090812080546001808201835591835292822090920180546001600160a01b031916938b16939093179092556005548b519092908c9085908110620002d657620002d662000583565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206002016000896001600160a01b03166001600160a01b03168152602001908152602001600020819055508080620003359062000559565b9150506200005a565b505050505050505050620005af565b80516001600160a01b03811681146200036557600080fd5b919050565b600082601f8301126200037c57600080fd5b81516020620003956200038f8362000533565b62000500565b80838252828201915082860187848660051b8901011115620003b657600080fd5b60005b85811015620003d757815184529284019290840190600101620003b9565b5090979650505050505050565b600080600080600080600080610100898b0312156200040257600080fd5b88516001600160401b03808211156200041a57600080fd5b818b0191508b601f8301126200042f57600080fd5b8151620004406200038f8262000533565b80828252602082019150602085018f60208560051b88010111156200046457600080fd5b600095505b8386101562000493576200047d816200034d565b8352600195909501946020928301920162000469565b5060208e0151909c509350505080821115620004ae57600080fd5b50620004bd8b828c016200036a565b975050620004ce60408a016200034d565b9550606089015194506080890151935060a0890151925060c0890151915060e089015190509295985092959890939650565b604051601f8201601f191681016001600160401b03811182821017156200052b576200052b62000599565b604052919050565b60006001600160401b038211156200054f576200054f62000599565b5060051b60200190565b60006000198214156200057c57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b61100680620005bf6000396000f3fe60806040526004361061011f5760003560e01c80636dd7d8ea116100a0578063d09f1ab411610064578063d09f1ab41461037e578063d161c76714610394578063d51b9e93146103aa578063d55b7dff146103fa578063f8ac9dd51461041057600080fd5b80636dd7d8ea146102e6578063a9a981a3146102f9578063a9ff959e1461030f578063ae6e43f514610325578063b642facd1461034557600080fd5b80632f9c4bba116100e75780632f9c4bba146101e8578063302b68721461020a5780633477ee2e14610254578063441a3e701461028c57806358e7525f146102ac57600080fd5b8063012679511461012457806302aa9be21461013957806306a49fce1461015957806315febd68146101845780632d15cc04146101c8575b600080fd5b610137610132366004610e22565b610426565b005b34801561014557600080fd5b50610137610154366004610e70565b6105e4565b34801561016557600080fd5b5061016e6107c1565b60405161017b9190610ed5565b60405180910390f35b34801561019057600080fd5b506101ba61019f366004610e9a565b33600090815260208181526040808320938352929052205490565b60405190815260200161017b565b3480156101d457600080fd5b5061016e6101e3366004610e22565b610823565b3480156101f457600080fd5b506101fd610899565b60405161017b9190610f22565b34801561021657600080fd5b506101ba610225366004610e3d565b6001600160a01b0391821660009081526001602090815260408083209390941682526002909201909152205490565b34801561026057600080fd5b5061027461026f366004610e9a565b6108fa565b6040516001600160a01b03909116815260200161017b565b34801561029857600080fd5b506101376102a7366004610eb3565b610924565b3480156102b857600080fd5b506101ba6102c7366004610e22565b6001600160a01b03166000908152600160208190526040909120015490565b6101376102f4366004610e22565b610a4f565b34801561030557600080fd5b506101ba60045481565b34801561031b57600080fd5b506101ba60095481565b34801561033157600080fd5b50610137610340366004610e22565b610bbd565b34801561035157600080fd5b50610274610360366004610e22565b6001600160a01b039081166000908152600160205260409020541690565b34801561038a57600080fd5b506101ba60075481565b3480156103a057600080fd5b506101ba60085481565b3480156103b657600080fd5b506103ea6103c5366004610e22565b6001600160a01b0316600090815260016020526040902054600160a01b900460ff1690565b604051901515815260200161017b565b34801561040657600080fd5b506101ba60055481565b34801561041c57600080fd5b506101ba60065481565b60055434101561043557600080fd5b6001600160a01b0381166000908152600160205260409020548190600160a01b900460ff161561046457600080fd5b6001600160a01b03821660009081526001602081905260408220015461048a9034610de7565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0387166001600160a01b03199091168117909155600090815260208281526040808320805460ff60a01b19339081166001600160a81b031990921691909117600160a01b1782559481018690559383526002909301905220549091506105279034610de7565b6001600160a01b03841660009081526001602081815260408084203385526002019091529091209190915560045461055e91610de7565b6004556001600160a01b038316600081815260026020908152604080832080546001810182559084529282902090920180546001600160a01b0319163390811790915582519081529081019290925234908201527f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c19060600160405180910390a1505050565b6001600160a01b03821660009081526001602090815260408083203384526002019091529020548290829081111561061b57600080fd5b6001600160a01b038281166000908152600160205260409020541633141561067c576005546001600160a01b03831660009081526001602090815260408083203384526002019091529020546106719083610dfa565b101561067c57600080fd5b6001600160a01b038416600090815260016020819052604090912001546106a39084610dfa565b6001600160a01b0385166000908152600160208181526040808420928301949094553383526002909101905220546106db9084610dfa565b6001600160a01b03851660009081526001602090815260408083203384526002019091528120919091556009546107129043610de7565b336000908152602081815260408083208484529091529020549091506107389085610de7565b336000818152602081815260408083208684528083528184209590955582825260019485018054958601815583529181902090930184905580519182526001600160a01b038816928201929092529081018590527faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2906060015b60405180910390a15050505050565b6060600380548060200260200160405190810160405280929190818152602001828054801561081957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107fb575b5050505050905090565b6001600160a01b03811660009081526002602090815260409182902080548351818402810184019094528084526060939283018282801561088d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161086f575b50505050509050919050565b336000908152602081815260409182902060010180548351818402810184019094528084526060939283018282801561081957602002820191906000526020600020905b8154815260200190600101908083116108dd575050505050905090565b6003818154811061090a57600080fd5b6000918252602090912001546001600160a01b0316905081565b81816000821161093357600080fd5b8143101561094057600080fd5b3360009081526020818152604080832085845290915290205461096257600080fd5b33600090815260208190526040902060010180548391908390811061098957610989610fba565b90600052602060002001541461099e57600080fd5b336000818152602081815260408083208884528083529083208054908490559383529190526001018054859081106109d8576109d8610fba565b60009182526020822001819055604051339183156108fc02918491818181858888f19350505050158015610a10573d6000803e3d6000fd5b5060408051338152602081018790529081018290527ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568906060016107b2565b600654341015610a5e57600080fd5b6001600160a01b0381166000908152600160205260409020548190600160a01b900460ff16610a8c57600080fd5b6001600160a01b03821660009081526001602081905260409091200154610ab39034610de7565b6001600160a01b038316600090815260016020818152604080842092830194909455338352600290910190522054610b1e576001600160a01b03821660009081526002602090815260408220805460018101825590835291200180546001600160a01b031916331790555b6001600160a01b0382166000908152600160209081526040808320338452600201909152902054610b4f9034610de7565b6001600160a01b0383166000818152600160209081526040808320338085526002909101835292819020949094558351918252810191909152348183015290517f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc9181900360600190a15050565b6001600160a01b038082166000908152600160205260409020548291163314610be557600080fd5b6001600160a01b0382166000908152600160205260409020548290600160a01b900460ff16610c1357600080fd5b6001600160a01b0383166000908152600160208190526040909120805460ff60a01b19169055600454610c4591610dfa565b60045560005b600354811015610cd057836001600160a01b031660038281548110610c7257610c72610fba565b6000918252602090912001546001600160a01b03161415610cbe5760038181548110610ca057610ca0610fba565b600091825260209091200180546001600160a01b0319169055610cd0565b80610cc881610f89565b915050610c4b565b506001600160a01b038316600081815260016020818152604080842033855260028101835290842054949093528190520154610d0c9082610dfa565b6001600160a01b03851660009081526001602081815260408084209283019490945533835260029091019052908120819055600854610d4b9043610de7565b33600090815260208181526040808320848452909152902054909150610d719083610de7565b336000818152602081815260408083208684528083528184209590955582825260019485018054958601815583529181902090930184905580519182526001600160a01b038816928201929092527f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d391016107b2565b6000610df38284610f5a565b9392505050565b6000610df38284610f72565b80356001600160a01b0381168114610e1d57600080fd5b919050565b600060208284031215610e3457600080fd5b610df382610e06565b60008060408385031215610e5057600080fd5b610e5983610e06565b9150610e6760208401610e06565b90509250929050565b60008060408385031215610e8357600080fd5b610e8c83610e06565b946020939093013593505050565b600060208284031215610eac57600080fd5b5035919050565b60008060408385031215610ec657600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015610f165783516001600160a01b031683529284019291840191600101610ef1565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610f1657835183529284019291840191600101610f3e565b60008219821115610f6d57610f6d610fa4565b500190565b600082821015610f8457610f84610fa4565b500390565b6000600019821415610f9d57610f9d610fa4565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fdfea2646970667358221220206de0541cf1366dd39a20f53e9a4070ac6e4fdb92b3e2485adf43ecd92d2a5e64736f6c63430008070033",
}

// JuncaValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use JuncaValidatorMetaData.ABI instead.
var JuncaValidatorABI = JuncaValidatorMetaData.ABI

// JuncaValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JuncaValidatorMetaData.Bin instead.
var JuncaValidatorBin = JuncaValidatorMetaData.Bin

// DeployJuncaValidator deploys a new Ethereum contract, binding an instance of JuncaValidator to it.
func DeployJuncaValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *JuncaValidator, error) {
	parsed, err := JuncaValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JuncaValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JuncaValidator{JuncaValidatorCaller: JuncaValidatorCaller{contract: contract}, JuncaValidatorTransactor: JuncaValidatorTransactor{contract: contract}, JuncaValidatorFilterer: JuncaValidatorFilterer{contract: contract}}, nil
}

// JuncaValidator is an auto generated Go binding around an Ethereum contract.
type JuncaValidator struct {
	JuncaValidatorCaller     // Read-only binding to the contract
	JuncaValidatorTransactor // Write-only binding to the contract
	JuncaValidatorFilterer   // Log filterer for contract events
}

// JuncaValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JuncaValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JuncaValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JuncaValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuncaValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JuncaValidatorSession struct {
	Contract     *JuncaValidator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JuncaValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JuncaValidatorCallerSession struct {
	Contract *JuncaValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// JuncaValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JuncaValidatorTransactorSession struct {
	Contract     *JuncaValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// JuncaValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JuncaValidatorRaw struct {
	Contract *JuncaValidator // Generic contract binding to access the raw methods on
}

// JuncaValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JuncaValidatorCallerRaw struct {
	Contract *JuncaValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// JuncaValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JuncaValidatorTransactorRaw struct {
	Contract *JuncaValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJuncaValidator creates a new instance of JuncaValidator, bound to a specific deployed contract.
func NewJuncaValidator(address common.Address, backend bind.ContractBackend) (*JuncaValidator, error) {
	contract, err := bindJuncaValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JuncaValidator{JuncaValidatorCaller: JuncaValidatorCaller{contract: contract}, JuncaValidatorTransactor: JuncaValidatorTransactor{contract: contract}, JuncaValidatorFilterer: JuncaValidatorFilterer{contract: contract}}, nil
}

// NewJuncaValidatorCaller creates a new read-only instance of JuncaValidator, bound to a specific deployed contract.
func NewJuncaValidatorCaller(address common.Address, caller bind.ContractCaller) (*JuncaValidatorCaller, error) {
	contract, err := bindJuncaValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorCaller{contract: contract}, nil
}

// NewJuncaValidatorTransactor creates a new write-only instance of JuncaValidator, bound to a specific deployed contract.
func NewJuncaValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JuncaValidatorTransactor, error) {
	contract, err := bindJuncaValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorTransactor{contract: contract}, nil
}

// NewJuncaValidatorFilterer creates a new log filterer instance of JuncaValidator, bound to a specific deployed contract.
func NewJuncaValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JuncaValidatorFilterer, error) {
	contract, err := bindJuncaValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorFilterer{contract: contract}, nil
}

// bindJuncaValidator binds a generic wrapper to an already deployed contract.
func bindJuncaValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JuncaValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuncaValidator *JuncaValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuncaValidator.Contract.JuncaValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuncaValidator *JuncaValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuncaValidator.Contract.JuncaValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuncaValidator *JuncaValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuncaValidator.Contract.JuncaValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuncaValidator *JuncaValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuncaValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuncaValidator *JuncaValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuncaValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuncaValidator *JuncaValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuncaValidator.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "candidateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) CandidateCount() (*big.Int, error) {
	return _JuncaValidator.Contract.CandidateCount(&_JuncaValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) CandidateCount() (*big.Int, error) {
	return _JuncaValidator.Contract.CandidateCount(&_JuncaValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "candidateWithdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _JuncaValidator.Contract.CandidateWithdrawDelay(&_JuncaValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _JuncaValidator.Contract.CandidateWithdrawDelay(&_JuncaValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_JuncaValidator *JuncaValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "candidates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_JuncaValidator *JuncaValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _JuncaValidator.Contract.Candidates(&_JuncaValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_JuncaValidator *JuncaValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _JuncaValidator.Contract.Candidates(&_JuncaValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(address _candidate) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getCandidateCap", _candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(address _candidate) view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _JuncaValidator.Contract.GetCandidateCap(&_JuncaValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(address _candidate) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _JuncaValidator.Contract.GetCandidateCap(&_JuncaValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_JuncaValidator *JuncaValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getCandidateOwner", _candidate)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_JuncaValidator *JuncaValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _JuncaValidator.Contract.GetCandidateOwner(&_JuncaValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_JuncaValidator *JuncaValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _JuncaValidator.Contract.GetCandidateOwner(&_JuncaValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns(address[])
func (_JuncaValidator *JuncaValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getCandidates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns(address[])
func (_JuncaValidator *JuncaValidatorSession) GetCandidates() ([]common.Address, error) {
	return _JuncaValidator.Contract.GetCandidates(&_JuncaValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns(address[])
func (_JuncaValidator *JuncaValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _JuncaValidator.Contract.GetCandidates(&_JuncaValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(address _candidate, address _voter) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getVoterCap", _candidate, _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(address _candidate, address _voter) view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _JuncaValidator.Contract.GetVoterCap(&_JuncaValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(address _candidate, address _voter) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _JuncaValidator.Contract.GetVoterCap(&_JuncaValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[])
func (_JuncaValidator *JuncaValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getVoters", _candidate)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[])
func (_JuncaValidator *JuncaValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _JuncaValidator.Contract.GetVoters(&_JuncaValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[])
func (_JuncaValidator *JuncaValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _JuncaValidator.Contract.GetVoters(&_JuncaValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() view returns(uint256[])
func (_JuncaValidator *JuncaValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getWithdrawBlockNumbers")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() view returns(uint256[])
func (_JuncaValidator *JuncaValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _JuncaValidator.Contract.GetWithdrawBlockNumbers(&_JuncaValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() view returns(uint256[])
func (_JuncaValidator *JuncaValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _JuncaValidator.Contract.GetWithdrawBlockNumbers(&_JuncaValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 _blockNumber) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "getWithdrawCap", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 _blockNumber) view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _JuncaValidator.Contract.GetWithdrawCap(&_JuncaValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 _blockNumber) view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _JuncaValidator.Contract.GetWithdrawCap(&_JuncaValidator.CallOpts, _blockNumber)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_JuncaValidator *JuncaValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "isCandidate", _candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_JuncaValidator *JuncaValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _JuncaValidator.Contract.IsCandidate(&_JuncaValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_JuncaValidator *JuncaValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _JuncaValidator.Contract.IsCandidate(&_JuncaValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "maxValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _JuncaValidator.Contract.MaxValidatorNumber(&_JuncaValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _JuncaValidator.Contract.MaxValidatorNumber(&_JuncaValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "minCandidateCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _JuncaValidator.Contract.MinCandidateCap(&_JuncaValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _JuncaValidator.Contract.MinCandidateCap(&_JuncaValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "minVoterCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) MinVoterCap() (*big.Int, error) {
	return _JuncaValidator.Contract.MinVoterCap(&_JuncaValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) MinVoterCap() (*big.Int, error) {
	return _JuncaValidator.Contract.MinVoterCap(&_JuncaValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuncaValidator.contract.Call(opts, &out, "voterWithdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
	return _JuncaValidator.Contract.VoterWithdrawDelay(&_JuncaValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() view returns(uint256)
func (_JuncaValidator *JuncaValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
	return _JuncaValidator.Contract.VoterWithdrawDelay(&_JuncaValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Propose(&_JuncaValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Propose(&_JuncaValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_JuncaValidator *JuncaValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_JuncaValidator *JuncaValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Resign(&_JuncaValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_JuncaValidator *JuncaValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Resign(&_JuncaValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address _candidate, uint256 _cap) returns()
func (_JuncaValidator *JuncaValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address _candidate, uint256 _cap) returns()
func (_JuncaValidator *JuncaValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Unvote(&_JuncaValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address _candidate, uint256 _cap) returns()
func (_JuncaValidator *JuncaValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Unvote(&_JuncaValidator.TransactOpts, _candidate, _cap)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Vote(&_JuncaValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _candidate) payable returns()
func (_JuncaValidator *JuncaValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Vote(&_JuncaValidator.TransactOpts, _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _blockNumber, uint256 _index) returns()
func (_JuncaValidator *JuncaValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _blockNumber, uint256 _index) returns()
func (_JuncaValidator *JuncaValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Withdraw(&_JuncaValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _blockNumber, uint256 _index) returns()
func (_JuncaValidator *JuncaValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JuncaValidator.Contract.Withdraw(&_JuncaValidator.TransactOpts, _blockNumber, _index)
}

// JuncaValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the JuncaValidator contract.
type JuncaValidatorProposeIterator struct {
	Event *JuncaValidatorPropose // Event containing the contract specifics and raw log

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
func (it *JuncaValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuncaValidatorPropose)
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
		it.Event = new(JuncaValidatorPropose)
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
func (it *JuncaValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuncaValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuncaValidatorPropose represents a Propose event raised by the JuncaValidator contract.
type JuncaValidatorPropose struct {
	Owner     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(address _owner, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*JuncaValidatorProposeIterator, error) {

	logs, sub, err := _JuncaValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorProposeIterator{contract: _JuncaValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(address _owner, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *JuncaValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _JuncaValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuncaValidatorPropose)
				if err := _JuncaValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// ParsePropose is a log parse operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(address _owner, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) ParsePropose(log types.Log) (*JuncaValidatorPropose, error) {
	event := new(JuncaValidatorPropose)
	if err := _JuncaValidator.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuncaValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the JuncaValidator contract.
type JuncaValidatorResignIterator struct {
	Event *JuncaValidatorResign // Event containing the contract specifics and raw log

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
func (it *JuncaValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuncaValidatorResign)
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
		it.Event = new(JuncaValidatorResign)
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
func (it *JuncaValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuncaValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuncaValidatorResign represents a Resign event raised by the JuncaValidator contract.
type JuncaValidatorResign struct {
	Owner     common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(address _owner, address _candidate)
func (_JuncaValidator *JuncaValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*JuncaValidatorResignIterator, error) {

	logs, sub, err := _JuncaValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorResignIterator{contract: _JuncaValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(address _owner, address _candidate)
func (_JuncaValidator *JuncaValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *JuncaValidatorResign) (event.Subscription, error) {

	logs, sub, err := _JuncaValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuncaValidatorResign)
				if err := _JuncaValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// ParseResign is a log parse operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(address _owner, address _candidate)
func (_JuncaValidator *JuncaValidatorFilterer) ParseResign(log types.Log) (*JuncaValidatorResign, error) {
	event := new(JuncaValidatorResign)
	if err := _JuncaValidator.contract.UnpackLog(event, "Resign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuncaValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the JuncaValidator contract.
type JuncaValidatorUnvoteIterator struct {
	Event *JuncaValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *JuncaValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuncaValidatorUnvote)
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
		it.Event = new(JuncaValidatorUnvote)
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
func (it *JuncaValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuncaValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuncaValidatorUnvote represents a Unvote event raised by the JuncaValidator contract.
type JuncaValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*JuncaValidatorUnvoteIterator, error) {

	logs, sub, err := _JuncaValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorUnvoteIterator{contract: _JuncaValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *JuncaValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _JuncaValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuncaValidatorUnvote)
				if err := _JuncaValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// ParseUnvote is a log parse operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) ParseUnvote(log types.Log) (*JuncaValidatorUnvote, error) {
	event := new(JuncaValidatorUnvote)
	if err := _JuncaValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuncaValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the JuncaValidator contract.
type JuncaValidatorVoteIterator struct {
	Event *JuncaValidatorVote // Event containing the contract specifics and raw log

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
func (it *JuncaValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuncaValidatorVote)
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
		it.Event = new(JuncaValidatorVote)
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
func (it *JuncaValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuncaValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuncaValidatorVote represents a Vote event raised by the JuncaValidator contract.
type JuncaValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*JuncaValidatorVoteIterator, error) {

	logs, sub, err := _JuncaValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorVoteIterator{contract: _JuncaValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *JuncaValidatorVote) (event.Subscription, error) {

	logs, sub, err := _JuncaValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuncaValidatorVote)
				if err := _JuncaValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address _voter, address _candidate, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) ParseVote(log types.Log) (*JuncaValidatorVote, error) {
	event := new(JuncaValidatorVote)
	if err := _JuncaValidator.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuncaValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the JuncaValidator contract.
type JuncaValidatorWithdrawIterator struct {
	Event *JuncaValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *JuncaValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuncaValidatorWithdraw)
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
		it.Event = new(JuncaValidatorWithdraw)
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
func (it *JuncaValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuncaValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuncaValidatorWithdraw represents a Withdraw event raised by the JuncaValidator contract.
type JuncaValidatorWithdraw struct {
	Owner       common.Address
	BlockNumber *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address _owner, uint256 _blockNumber, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*JuncaValidatorWithdrawIterator, error) {

	logs, sub, err := _JuncaValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &JuncaValidatorWithdrawIterator{contract: _JuncaValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address _owner, uint256 _blockNumber, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *JuncaValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _JuncaValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuncaValidatorWithdraw)
				if err := _JuncaValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address _owner, uint256 _blockNumber, uint256 _cap)
func (_JuncaValidator *JuncaValidatorFilterer) ParseWithdraw(log types.Log) (*JuncaValidatorWithdraw, error) {
	event := new(JuncaValidatorWithdraw)
	if err := _JuncaValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
