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

// JRC21PresetFixedMetaData contains all meta data concerning the JRC21PresetFixed contract.
var JRC21PresetFixedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"initialAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9610140523480156200003757600080fd5b506040516200253c3803806200253c8339810160408190526200005a9162000644565b6040805180820190915260018152603160f81b6020820152879081908189868662000087600083620001d6565b600380546001600160a01b0319166001600160a01b0393909316929092179091556002558151620000c0906007906020850190620004ca565b508051620000d6906008906020840190620004ca565b5050825160208085019190912083518483012060e08290526101008190524660a0818152604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81880181905281830187905260608201869052608082019490945230818401528151808203909301835260c0019052805194019390932091935091906080523060601b60c052610120525050600a805461ffff191661010060ff8b1602179055506200019191506000905083620001d6565b620001bd7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a83620001d6565b620001c98484620001e6565b505050505050506200077a565b620001e28282620002dd565b5050565b6001600160a01b038216620002425760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b620002506000838362000320565b806006600082825462000264919062000700565b90915550506001600160a01b038216600090815260046020526040812080548392906200029390849062000700565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b620002f482826200033860201b62000b4e1760201c565b60008281526001602090815260409091206200031b91839062000bd2620003d8821b17901c565b505050565b6200031b838383620003f860201b62000be71760201c565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16620001e2576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620003943390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000620003ef836001600160a01b03841662000478565b90505b92915050565b620004108383836200031b60201b6200062b1760201c565b600a5460ff16156200031b5760405162461bcd60e51b815260206004820152602a60248201527f45524332305061757361626c653a20746f6b656e207472616e736665722077686044820152691a5b19481c185d5cd95960b21b606482015260840162000239565b6000818152600183016020526040812054620004c157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003f2565b506000620003f2565b828054620004d89062000727565b90600052602060002090601f016020900481019282620004fc576000855562000547565b82601f106200051757805160ff191683800117855562000547565b8280016001018555821562000547579182015b82811115620005475782518255916020019190600101906200052a565b506200055592915062000559565b5090565b5b808211156200055557600081556001016200055a565b80516001600160a01b03811681146200058857600080fd5b919050565b600082601f8301126200059f57600080fd5b81516001600160401b0380821115620005bc57620005bc62000764565b604051601f8301601f19908116603f01168101908282118183101715620005e757620005e762000764565b816040528381526020925086838588010111156200060457600080fd5b600091505b8382101562000628578582018301518183018401529082019062000609565b838211156200063a5760008385830101525b9695505050505050565b600080600080600080600060e0888a0312156200066057600080fd5b87516001600160401b03808211156200067857600080fd5b620006868b838c016200058d565b985060208a01519150808211156200069d57600080fd5b50620006ac8a828b016200058d565b965050604088015160ff81168114620006c457600080fd5b9450620006d46060890162000570565b935060808801519250620006eb60a0890162000570565b915060c0880151905092959891949750929550565b600082198211156200072257634e487b7160e01b600052601160045260246000fd5b500190565b600181811c908216806200073c57607f821691505b602082108114156200075e57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b60805160a05160c05160601c60e051610100516101205161014051611d64620007d86000396000610a1801526000611099015260006110e8015260006110c30152600061101c01526000611046015260006110700152611d646000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80633f4ba83a1161010f578063a217fddf116100a2578063d505accf11610071578063d505accf146103fd578063d547741f14610410578063dd62ed3e14610423578063e63ab1e91461045c57600080fd5b8063a217fddf146103bc578063a457c2d7146103c4578063a9059cbb146103d7578063ca15c873146103ea57600080fd5b80638456cb59116100de5780638456cb59146103865780639010d07c1461038e57806391d14854146103a157806395d89b41146103b457600080fd5b80633f4ba83a146103375780635c975abb1461033f57806370a082311461034a5780637ecebe001461037357600080fd5b8063248a9ca3116101875780633408e470116101565780633408e470146103035780633644e5151461030957806336568abe14610311578063395093511461032457600080fd5b8063248a9ca31461029b5780632f2ff15d146102be578063313ce567146102d357806331ac9920146102f057600080fd5b8063127e8e4d116101c3578063127e8e4d1461023a57806318160ddd1461025b5780631d1438481461026357806323b872dd1461028857600080fd5b806301ffc9a7146101ea57806306fdde0314610212578063095ea7b314610227575b600080fd5b6101fd6101f8366004611b28565b610483565b60405190151581526020015b60405180910390f35b61021a610494565b6040516102099190611bc7565b6101fd610235366004611aa0565b610526565b61024d610248366004611aca565b61053c565b604051908152602001610209565b60065461024d565b6003546001600160a01b03165b6040516001600160a01b039091168152602001610209565b6101fd6102963660046119f1565b610556565b61024d6102a9366004611aca565b60009081526020819052604090206001015490565b6102d16102cc366004611ae3565b610605565b005b600a54610100900460ff1660405160ff9091168152602001610209565b6102d16102fe366004611aca565b610630565b4661024d565b61024d610697565b6102d161031f366004611ae3565b6106a6565b6101fd610332366004611aa0565b610724565b6102d1610760565b600a5460ff166101fd565b61024d6103583660046119a3565b6001600160a01b031660009081526004602052604090205490565b61024d6103813660046119a3565b6107fb565b6102d1610819565b61027061039c366004611b06565b6108b0565b6101fd6103af366004611ae3565b6108cf565b61021a6108f8565b61024d600081565b6101fd6103d2366004611aa0565b610907565b6101fd6103e5366004611aa0565b6109a0565b61024d6103f8366004611aca565b6109ad565b6102d161040b366004611a2d565b6109c4565b6102d161041e366004611ae3565b610b28565b61024d6104313660046119be565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205490565b61024d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b600061048e82610c4d565b92915050565b6060600780546104a390611c8b565b80601f01602080910402602001604051908101604052809291908181526020018280546104cf90611c8b565b801561051c5780601f106104f15761010080835404028352916020019161051c565b820191906000526020600020905b8154815290600101906020018083116104ff57829003601f168201915b5050505050905090565b6000610533338484610c72565b50600192915050565b60025460009061048e906105508484610d96565b90610da2565b6000610563848484610dae565b6001600160a01b0384166000908152600560209081526040808320338452909152902054828110156105ed5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6105fa8533858403610c72565b506001949350505050565b6000828152602081905260409020600101546106218133610f89565b61062b8383610fed565b505050565b61063b6000336108cf565b6106925760405162461bcd60e51b815260206004820152602260248201527f4a524332313a206d75737420686176652061646d696e20726f6c6520746f2073604482015261195d60f21b60648201526084016105e4565b600255565b60006106a161100f565b905090565b6001600160a01b03811633146107165760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016105e4565b6107208282611136565b5050565b3360008181526005602090815260408083206001600160a01b0387168452909152812054909161053391859061075b908690611bfa565b610c72565b61078a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a336108cf565b6107f15760405162461bcd60e51b815260206004820152603260248201527f4a5243323150726573657446697865643a206d75737420686176652070617573604482015271657220726f6c6520746f20756e706175736560701b60648201526084016105e4565b6107f9611158565b565b6001600160a01b03811660009081526009602052604081205461048e565b6108437f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a336108cf565b6108a85760405162461bcd60e51b815260206004820152603060248201527f4a5243323150726573657446697865643a206d7573742068617665207061757360448201526f657220726f6c6520746f20706175736560801b60648201526084016105e4565b6107f96111eb565b60008281526001602052604081206108c89083611266565b9392505050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b6060600880546104a390611c8b565b3360009081526005602090815260408083206001600160a01b0386168452909152812054828110156109895760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016105e4565b6109963385858403610c72565b5060019392505050565b6000610533338484610dae565b600081815260016020526040812061048e90611272565b83421115610a145760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016105e4565b60007f0000000000000000000000000000000000000000000000000000000000000000888888610a438c61127c565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610a9e826112a4565b90506000610aae828787876112f2565b9050896001600160a01b0316816001600160a01b031614610b115760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016105e4565b610b1c8a8a8a610c72565b50505050505050505050565b600082815260208190526040902060010154610b448133610f89565b61062b8383611136565b610b5882826108cf565b610720576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055610b8e3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006108c8836001600160a01b03841661131a565b600a5460ff161561062b5760405162461bcd60e51b815260206004820152602a60248201527f45524332305061757361626c653a20746f6b656e207472616e736665722077686044820152691a5b19481c185d5cd95960b21b60648201526084016105e4565b60006001600160e01b03198216635a05180f60e01b148061048e575061048e82611369565b6001600160a01b038316610cd45760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016105e4565b6001600160a01b038216610d355760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016105e4565b6001600160a01b0383811660008181526005602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006108c88284611c12565b60006108c88284611bfa565b6001600160a01b038316610e125760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016105e4565b6001600160a01b038216610e745760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016105e4565b610e7f83838361139e565b6001600160a01b03831660009081526004602052604090205481811015610ef75760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016105e4565b6001600160a01b03808516600090815260046020526040808220858503905591851681529081208054849290610f2e908490611bfa565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f7a91815260200190565b60405180910390a35b50505050565b610f9382826108cf565b61072057610fab816001600160a01b031660146113a9565b610fb68360206113a9565b604051602001610fc7929190611b52565b60408051601f198184030181529082905262461bcd60e51b82526105e491600401611bc7565b610ff78282610b4e565b600082815260016020526040902061062b9082610bd2565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561106857507f000000000000000000000000000000000000000000000000000000000000000046145b1561109257507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6111408282611545565b600082815260016020526040902061062b90826115aa565b600a5460ff166111a15760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016105e4565b600a805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600a5460ff16156112315760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016105e4565b600a805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111ce3390565b60006108c883836115bf565b600061048e825490565b6001600160a01b03811660009081526009602052604090208054600181018255905b50919050565b600061048e6112b161100f565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6000806000611303878787876115e9565b91509150611310816116d6565b5095945050505050565b60008181526001830160205260408120546113615750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561048e565b50600061048e565b60006001600160e01b03198216637965db0b60e01b148061048e57506301ffc9a760e01b6001600160e01b031983161461048e565b61062b838383610be7565b606060006113b8836002611c12565b6113c3906002611bfa565b67ffffffffffffffff8111156113db576113db611d18565b6040519080825280601f01601f191660200182016040528015611405576020820181803683370190505b509050600360fc1b8160008151811061142057611420611d02565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061144f5761144f611d02565b60200101906001600160f81b031916908160001a9053506000611473846002611c12565b61147e906001611bfa565b90505b60018111156114f6576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106114b2576114b2611d02565b1a60f81b8282815181106114c8576114c8611d02565b60200101906001600160f81b031916908160001a90535060049490941c936114ef81611c74565b9050611481565b5083156108c85760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016105e4565b61154f82826108cf565b15610720576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60006108c8836001600160a01b038416611894565b60008260000182815481106115d6576115d6611d02565b9060005260206000200154905092915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561162057506000905060036116cd565b8460ff16601b1415801561163857508460ff16601c14155b1561164957506000905060046116cd565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561169d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166116c6576000600192509250506116cd565b9150600090505b94509492505050565b60008160048111156116ea576116ea611cd6565b14156116f35750565b600181600481111561170757611707611cd6565b14156117555760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105e4565b600281600481111561176957611769611cd6565b14156117b75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105e4565b60038160048111156117cb576117cb611cd6565b14156118245760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016105e4565b600481600481111561183857611838611cd6565b14156118915760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016105e4565b50565b6000818152600183016020526040812054801561197d5760006118b8600183611c31565b85549091506000906118cc90600190611c31565b90508181146119315760008660000182815481106118ec576118ec611d02565b906000526020600020015490508087600001848154811061190f5761190f611d02565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061194257611942611cec565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061048e565b600091505061048e565b80356001600160a01b038116811461199e57600080fd5b919050565b6000602082840312156119b557600080fd5b6108c882611987565b600080604083850312156119d157600080fd5b6119da83611987565b91506119e860208401611987565b90509250929050565b600080600060608486031215611a0657600080fd5b611a0f84611987565b9250611a1d60208501611987565b9150604084013590509250925092565b600080600080600080600060e0888a031215611a4857600080fd5b611a5188611987565b9650611a5f60208901611987565b95506040880135945060608801359350608088013560ff81168114611a8357600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611ab357600080fd5b611abc83611987565b946020939093013593505050565b600060208284031215611adc57600080fd5b5035919050565b60008060408385031215611af657600080fd5b823591506119e860208401611987565b60008060408385031215611b1957600080fd5b50508035926020909101359150565b600060208284031215611b3a57600080fd5b81356001600160e01b0319811681146108c857600080fd5b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611b8a816017850160208801611c48565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611bbb816028840160208801611c48565b01602801949350505050565b6020815260008251806020840152611be6816040850160208701611c48565b601f01601f19169190910160400192915050565b60008219821115611c0d57611c0d611cc0565b500190565b6000816000190483118215151615611c2c57611c2c611cc0565b500290565b600082821015611c4357611c43611cc0565b500390565b60005b83811015611c63578181015183820152602001611c4b565b83811115610f835750506000910152565b600081611c8357611c83611cc0565b506000190190565b600181811c90821680611c9f57607f821691505b6020821081141561129e57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220935befabc5929bd837e474657f9dd49a116f31b589106a7917c05e0d45923a0364736f6c63430008070033",
}

// JRC21PresetFixedABI is the input ABI used to generate the binding from.
// Deprecated: Use JRC21PresetFixedMetaData.ABI instead.
var JRC21PresetFixedABI = JRC21PresetFixedMetaData.ABI

// JRC21PresetFixedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JRC21PresetFixedMetaData.Bin instead.
var JRC21PresetFixedBin = JRC21PresetFixedMetaData.Bin

// DeployJRC21PresetFixed deploys a new Ethereum contract, binding an instance of JRC21PresetFixed to it.
func DeployJRC21PresetFixed(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8, initialAccount common.Address, initialBalance *big.Int, issuer common.Address, minFee *big.Int) (common.Address, *types.Transaction, *JRC21PresetFixed, error) {
	parsed, err := JRC21PresetFixedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JRC21PresetFixedBin), backend, name, symbol, decimals_, initialAccount, initialBalance, issuer, minFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JRC21PresetFixed{JRC21PresetFixedCaller: JRC21PresetFixedCaller{contract: contract}, JRC21PresetFixedTransactor: JRC21PresetFixedTransactor{contract: contract}, JRC21PresetFixedFilterer: JRC21PresetFixedFilterer{contract: contract}}, nil
}

// JRC21PresetFixed is an auto generated Go binding around an Ethereum contract.
type JRC21PresetFixed struct {
	JRC21PresetFixedCaller     // Read-only binding to the contract
	JRC21PresetFixedTransactor // Write-only binding to the contract
	JRC21PresetFixedFilterer   // Log filterer for contract events
}

// JRC21PresetFixedCaller is an auto generated read-only Go binding around an Ethereum contract.
type JRC21PresetFixedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JRC21PresetFixedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JRC21PresetFixedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JRC21PresetFixedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JRC21PresetFixedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JRC21PresetFixedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JRC21PresetFixedSession struct {
	Contract     *JRC21PresetFixed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JRC21PresetFixedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JRC21PresetFixedCallerSession struct {
	Contract *JRC21PresetFixedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// JRC21PresetFixedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JRC21PresetFixedTransactorSession struct {
	Contract     *JRC21PresetFixedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// JRC21PresetFixedRaw is an auto generated low-level Go binding around an Ethereum contract.
type JRC21PresetFixedRaw struct {
	Contract *JRC21PresetFixed // Generic contract binding to access the raw methods on
}

// JRC21PresetFixedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JRC21PresetFixedCallerRaw struct {
	Contract *JRC21PresetFixedCaller // Generic read-only contract binding to access the raw methods on
}

// JRC21PresetFixedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JRC21PresetFixedTransactorRaw struct {
	Contract *JRC21PresetFixedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJRC21PresetFixed creates a new instance of JRC21PresetFixed, bound to a specific deployed contract.
func NewJRC21PresetFixed(address common.Address, backend bind.ContractBackend) (*JRC21PresetFixed, error) {
	contract, err := bindJRC21PresetFixed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixed{JRC21PresetFixedCaller: JRC21PresetFixedCaller{contract: contract}, JRC21PresetFixedTransactor: JRC21PresetFixedTransactor{contract: contract}, JRC21PresetFixedFilterer: JRC21PresetFixedFilterer{contract: contract}}, nil
}

// NewJRC21PresetFixedCaller creates a new read-only instance of JRC21PresetFixed, bound to a specific deployed contract.
func NewJRC21PresetFixedCaller(address common.Address, caller bind.ContractCaller) (*JRC21PresetFixedCaller, error) {
	contract, err := bindJRC21PresetFixed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedCaller{contract: contract}, nil
}

// NewJRC21PresetFixedTransactor creates a new write-only instance of JRC21PresetFixed, bound to a specific deployed contract.
func NewJRC21PresetFixedTransactor(address common.Address, transactor bind.ContractTransactor) (*JRC21PresetFixedTransactor, error) {
	contract, err := bindJRC21PresetFixed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedTransactor{contract: contract}, nil
}

// NewJRC21PresetFixedFilterer creates a new log filterer instance of JRC21PresetFixed, bound to a specific deployed contract.
func NewJRC21PresetFixedFilterer(address common.Address, filterer bind.ContractFilterer) (*JRC21PresetFixedFilterer, error) {
	contract, err := bindJRC21PresetFixed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedFilterer{contract: contract}, nil
}

// bindJRC21PresetFixed binds a generic wrapper to an already deployed contract.
func bindJRC21PresetFixed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JRC21PresetFixedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JRC21PresetFixed *JRC21PresetFixedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JRC21PresetFixed.Contract.JRC21PresetFixedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JRC21PresetFixed *JRC21PresetFixedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.JRC21PresetFixedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JRC21PresetFixed *JRC21PresetFixedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.JRC21PresetFixedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JRC21PresetFixed *JRC21PresetFixedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JRC21PresetFixed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JRC21PresetFixed *JRC21PresetFixedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JRC21PresetFixed *JRC21PresetFixedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.DEFAULTADMINROLE(&_JRC21PresetFixed.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.DEFAULTADMINROLE(&_JRC21PresetFixed.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.DOMAINSEPARATOR(&_JRC21PresetFixed.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.DOMAINSEPARATOR(&_JRC21PresetFixed.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedSession) PAUSERROLE() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.PAUSERROLE(&_JRC21PresetFixed.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) PAUSERROLE() ([32]byte, error) {
	return _JRC21PresetFixed.Contract.PAUSERROLE(&_JRC21PresetFixed.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.Allowance(&_JRC21PresetFixed.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.Allowance(&_JRC21PresetFixed.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.BalanceOf(&_JRC21PresetFixed.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.BalanceOf(&_JRC21PresetFixed.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Decimals() (uint8, error) {
	return _JRC21PresetFixed.Contract.Decimals(&_JRC21PresetFixed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Decimals() (uint8, error) {
	return _JRC21PresetFixed.Contract.Decimals(&_JRC21PresetFixed.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "estimateFee", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.EstimateFee(&_JRC21PresetFixed.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.EstimateFee(&_JRC21PresetFixed.CallOpts, value)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) GetChainId() (*big.Int, error) {
	return _JRC21PresetFixed.Contract.GetChainId(&_JRC21PresetFixed.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) GetChainId() (*big.Int, error) {
	return _JRC21PresetFixed.Contract.GetChainId(&_JRC21PresetFixed.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _JRC21PresetFixed.Contract.GetRoleAdmin(&_JRC21PresetFixed.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _JRC21PresetFixed.Contract.GetRoleAdmin(&_JRC21PresetFixed.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _JRC21PresetFixed.Contract.GetRoleMember(&_JRC21PresetFixed.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _JRC21PresetFixed.Contract.GetRoleMember(&_JRC21PresetFixed.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.GetRoleMemberCount(&_JRC21PresetFixed.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.GetRoleMemberCount(&_JRC21PresetFixed.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _JRC21PresetFixed.Contract.HasRole(&_JRC21PresetFixed.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _JRC21PresetFixed.Contract.HasRole(&_JRC21PresetFixed.CallOpts, role, account)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "issuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Issuer() (common.Address, error) {
	return _JRC21PresetFixed.Contract.Issuer(&_JRC21PresetFixed.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Issuer() (common.Address, error) {
	return _JRC21PresetFixed.Contract.Issuer(&_JRC21PresetFixed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Name() (string, error) {
	return _JRC21PresetFixed.Contract.Name(&_JRC21PresetFixed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Name() (string, error) {
	return _JRC21PresetFixed.Contract.Name(&_JRC21PresetFixed.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Nonces(owner common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.Nonces(&_JRC21PresetFixed.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _JRC21PresetFixed.Contract.Nonces(&_JRC21PresetFixed.CallOpts, owner)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Paused() (bool, error) {
	return _JRC21PresetFixed.Contract.Paused(&_JRC21PresetFixed.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Paused() (bool, error) {
	return _JRC21PresetFixed.Contract.Paused(&_JRC21PresetFixed.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JRC21PresetFixed.Contract.SupportsInterface(&_JRC21PresetFixed.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JRC21PresetFixed.Contract.SupportsInterface(&_JRC21PresetFixed.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Symbol() (string, error) {
	return _JRC21PresetFixed.Contract.Symbol(&_JRC21PresetFixed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) Symbol() (string, error) {
	return _JRC21PresetFixed.Contract.Symbol(&_JRC21PresetFixed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JRC21PresetFixed.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedSession) TotalSupply() (*big.Int, error) {
	return _JRC21PresetFixed.Contract.TotalSupply(&_JRC21PresetFixed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JRC21PresetFixed *JRC21PresetFixedCallerSession) TotalSupply() (*big.Int, error) {
	return _JRC21PresetFixed.Contract.TotalSupply(&_JRC21PresetFixed.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Approve(&_JRC21PresetFixed.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Approve(&_JRC21PresetFixed.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.DecreaseAllowance(&_JRC21PresetFixed.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.DecreaseAllowance(&_JRC21PresetFixed.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.GrantRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.GrantRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.IncreaseAllowance(&_JRC21PresetFixed.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.IncreaseAllowance(&_JRC21PresetFixed.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) Pause() (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Pause(&_JRC21PresetFixed.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) Pause() (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Pause(&_JRC21PresetFixed.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Permit(&_JRC21PresetFixed.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Permit(&_JRC21PresetFixed.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.RenounceRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.RenounceRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.RevokeRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.RevokeRole(&_JRC21PresetFixed.TransactOpts, role, account)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(uint256 value) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) SetMinFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "setMinFee", value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(uint256 value) returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.SetMinFee(&_JRC21PresetFixed.TransactOpts, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(uint256 value) returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.SetMinFee(&_JRC21PresetFixed.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Transfer(&_JRC21PresetFixed.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Transfer(&_JRC21PresetFixed.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.TransferFrom(&_JRC21PresetFixed.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.TransferFrom(&_JRC21PresetFixed.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JRC21PresetFixed.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedSession) Unpause() (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Unpause(&_JRC21PresetFixed.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JRC21PresetFixed *JRC21PresetFixedTransactorSession) Unpause() (*types.Transaction, error) {
	return _JRC21PresetFixed.Contract.Unpause(&_JRC21PresetFixed.TransactOpts)
}

// JRC21PresetFixedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedApprovalIterator struct {
	Event *JRC21PresetFixedApproval // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedApproval)
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
		it.Event = new(JRC21PresetFixedApproval)
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
func (it *JRC21PresetFixedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedApproval represents a Approval event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JRC21PresetFixedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedApprovalIterator{contract: _JRC21PresetFixed.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedApproval)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseApproval(log types.Log) (*JRC21PresetFixedApproval, error) {
	event := new(JRC21PresetFixedApproval)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedPausedIterator struct {
	Event *JRC21PresetFixedPaused // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedPaused)
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
		it.Event = new(JRC21PresetFixedPaused)
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
func (it *JRC21PresetFixedPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedPaused represents a Paused event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterPaused(opts *bind.FilterOpts) (*JRC21PresetFixedPausedIterator, error) {

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedPausedIterator{contract: _JRC21PresetFixed.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedPaused) (event.Subscription, error) {

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedPaused)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParsePaused(log types.Log) (*JRC21PresetFixedPaused, error) {
	event := new(JRC21PresetFixedPaused)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleAdminChangedIterator struct {
	Event *JRC21PresetFixedRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedRoleAdminChanged)
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
		it.Event = new(JRC21PresetFixedRoleAdminChanged)
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
func (it *JRC21PresetFixedRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedRoleAdminChanged represents a RoleAdminChanged event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*JRC21PresetFixedRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedRoleAdminChangedIterator{contract: _JRC21PresetFixed.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedRoleAdminChanged)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseRoleAdminChanged(log types.Log) (*JRC21PresetFixedRoleAdminChanged, error) {
	event := new(JRC21PresetFixedRoleAdminChanged)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleGrantedIterator struct {
	Event *JRC21PresetFixedRoleGranted // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedRoleGranted)
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
		it.Event = new(JRC21PresetFixedRoleGranted)
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
func (it *JRC21PresetFixedRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedRoleGranted represents a RoleGranted event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JRC21PresetFixedRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedRoleGrantedIterator{contract: _JRC21PresetFixed.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedRoleGranted)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseRoleGranted(log types.Log) (*JRC21PresetFixedRoleGranted, error) {
	event := new(JRC21PresetFixedRoleGranted)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleRevokedIterator struct {
	Event *JRC21PresetFixedRoleRevoked // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedRoleRevoked)
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
		it.Event = new(JRC21PresetFixedRoleRevoked)
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
func (it *JRC21PresetFixedRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedRoleRevoked represents a RoleRevoked event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JRC21PresetFixedRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedRoleRevokedIterator{contract: _JRC21PresetFixed.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedRoleRevoked)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseRoleRevoked(log types.Log) (*JRC21PresetFixedRoleRevoked, error) {
	event := new(JRC21PresetFixedRoleRevoked)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedTransferIterator struct {
	Event *JRC21PresetFixedTransfer // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedTransfer)
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
		it.Event = new(JRC21PresetFixedTransfer)
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
func (it *JRC21PresetFixedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedTransfer represents a Transfer event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JRC21PresetFixedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedTransferIterator{contract: _JRC21PresetFixed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedTransfer)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseTransfer(log types.Log) (*JRC21PresetFixedTransfer, error) {
	event := new(JRC21PresetFixedTransfer)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JRC21PresetFixedUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the JRC21PresetFixed contract.
type JRC21PresetFixedUnpausedIterator struct {
	Event *JRC21PresetFixedUnpaused // Event containing the contract specifics and raw log

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
func (it *JRC21PresetFixedUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JRC21PresetFixedUnpaused)
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
		it.Event = new(JRC21PresetFixedUnpaused)
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
func (it *JRC21PresetFixedUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JRC21PresetFixedUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JRC21PresetFixedUnpaused represents a Unpaused event raised by the JRC21PresetFixed contract.
type JRC21PresetFixedUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) FilterUnpaused(opts *bind.FilterOpts) (*JRC21PresetFixedUnpausedIterator, error) {

	logs, sub, err := _JRC21PresetFixed.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &JRC21PresetFixedUnpausedIterator{contract: _JRC21PresetFixed.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *JRC21PresetFixedUnpaused) (event.Subscription, error) {

	logs, sub, err := _JRC21PresetFixed.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JRC21PresetFixedUnpaused)
				if err := _JRC21PresetFixed.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JRC21PresetFixed *JRC21PresetFixedFilterer) ParseUnpaused(log types.Log) (*JRC21PresetFixedUnpaused, error) {
	event := new(JRC21PresetFixedUnpaused)
	if err := _JRC21PresetFixed.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
