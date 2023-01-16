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
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]",
	Bin: "0x6060604052600060045534156200001557600080fd5b6040516200264d3803806200264d83398101604052808051820191906020018051820191906020018051906020019091908051906020019091908051906020019091908051906020019091908051906020019091908051906020019091905050600085600581905550846006819055508360078190555082600881905550816009819055508851600481905550600090505b8851811015620003ac5760038054806001018281620000c79190620003bb565b916000526020600020900160008b84815181101515620000e357fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506060604051908101604052808873ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200189838151811015156200016e57fe5b90602001906020020151815250600160008b848151811015156200018e57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155905050600260008a838151811015156200025957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281620002b19190620003bb565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600554600160008b848151811015156200031657fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080600101915050620000a7565b50505050505050505062000412565b815481835581811511620003e557818360005260206000209182019101620003e49190620003ea565b5b505050565b6200040f91905b808211156200040b576000816000905550600101620003f1565b5090565b90565b61222b80620004226000396000f300606060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063012679511461011757806302aa9be21461014557806306a49fce1461018757806315febd68146101f15780632d15cc04146102285780632f9c4bba146102b6578063302b6872146103205780633477ee2e1461038c578063441a3e70146103ef57806358e7525f1461041b5780636dd7d8ea14610468578063a9a981a314610496578063a9ff959e146104bf578063ae6e43f5146104e8578063b642facd14610521578063d09f1ab41461059a578063d161c767146105c3578063d51b9e93146105ec578063d55b7dff1461063d578063f8ac9dd514610666575b600080fd5b610143600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061068f565b005b341561015057600080fd5b610185600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610b16565b005b341561019257600080fd5b61019a611071565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101dd5780820151818401526020810190506101c2565b505050509050019250505060405180910390f35b34156101fc57600080fd5b6102126004808035906020019091905050611105565b6040518082815260200191505060405180910390f35b341561023357600080fd5b61025f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611161565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102a2578082015181840152602081019050610287565b505050509050019250505060405180910390f35b34156102c157600080fd5b6102c9611234565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561030c5780820151818401526020810190506102f1565b505050509050019250505060405180910390f35b341561032b57600080fd5b610376600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506112d1565b6040518082815260200191505060405180910390f35b341561039757600080fd5b6103ad600480803590602001909190505061135b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103fa57600080fd5b610419600480803590602001909190803590602001909190505061139a565b005b341561042657600080fd5b610452600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611646565b6040518082815260200191505060405180910390f35b610494600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611692565b005b34156104a157600080fd5b6104a9611a7b565b6040518082815260200191505060405180910390f35b34156104ca57600080fd5b6104d2611a81565b6040518082815260200191505060405180910390f35b34156104f357600080fd5b61051f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611a87565b005b341561052c57600080fd5b610558600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612046565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105a557600080fd5b6105ad6120b2565b6040518082815260200191505060405180910390f35b34156105ce57600080fd5b6105d66120b8565b6040518082815260200191505060405180910390f35b34156105f757600080fd5b610623600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506120be565b604051808215151515815260200191505060405180910390f35b341561064857600080fd5b610650612117565b6040518082815260200191505060405180910390f35b341561067157600080fd5b61067961211d565b6040518082815260200191505060405180910390f35b600060055434101515156106a257600080fd5b81600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515156106ff57600080fd5b61075434600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461212390919063ffffffff16565b91506003805480600101828161076a919061215a565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506060604051908101604052803373ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200183815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff0219169083151502179055506040820151816001015590505061093334600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461212390919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506109cc600160045461212390919063ffffffff16565b600481905550600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281610a23919061215a565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1338434604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b6000828280600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610ba857600080fd5b3373ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ce157600554610cd382600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461214190919063ffffffff16565b10151515610ce057600080fd5b5b610d3684600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461214190919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550610e0e84600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461214190919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ea64360095461212390919063ffffffff16565b9250610f0d846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008681526020019081526020016000205461212390919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018054806001018281610fb69190612186565b9160005260206000209001600085909190915055507faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338686604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b6110796121b2565b60038054806020026020016040519081016040528092919081815260200182805480156110fb57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116110b1575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000838152602001908152602001600020549050919050565b6111696121b2565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561122857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116111de575b50505050509050919050565b61123c6121c6565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018054806020026020016040519081016040528092919081815260200182805480156112c757602002820191906000526020600020905b8154815260200190600101908083116112b3575b5050505050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60038181548110151561136a57fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600082826000821115156113ad57600080fd5b8143101515156113bc57600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008481526020019081526020016000205411151561141d57600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018281548110151561146c57fe5b90600052602060002090015414151561148457600080fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008681526020019081526020016000205492506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020600090556000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018481548110151561157d57fe5b9060005260206000209001600090553373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f1935050505015156115cc57600080fd5b7ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568338685604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a15050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050919050565b60065434101515156116a357600080fd5b80600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff1615156116ff57600080fd5b61175434600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461212390919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156118c357600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281611873919061215a565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b61195534600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461212390919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc338334604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b60045481565b60095481565b6000806000833373ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515611b2957600080fd5b84600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515611b8557600080fd5b6000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff021916908315150217905550611bf6600160045461214190919063ffffffff16565b600481905550600094505b600380549050851015611ccb578573ffffffffffffffffffffffffffffffffffffffff16600386815481101515611c3457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611cbe57600385815481101515611c8b57fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055611ccb565b8480600101955050611c01565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549350611da284600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461214190919063ffffffff16565b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611e824360085461212390919063ffffffff16565b9250611ee9846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008681526020019081526020016000205461212390919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018054806001018281611f929190612186565b9160005260206000209001600085909190915055507f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33387604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60075481565b60085481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff169050919050565b60055481565b60065481565b600080828401905083811015151561213757fe5b8091505092915050565b600082821115151561214f57fe5b818303905092915050565b8154818355818115116121815781836000526020600020918201910161218091906121da565b5b505050565b8154818355818115116121ad578183600052602060002091820191016121ac91906121da565b5b505050565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b6121fc91905b808211156121f85760008160009055506001016121e0565b5090565b905600a165627a7a7230582057741f577a1575b5daff87f6c7ec34ed30d1cd52d11e75082031eaad5941c06e0029",
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
