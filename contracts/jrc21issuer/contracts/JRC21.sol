// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts@4.4.0/token/ERC20/extensions/draft-ERC20Permit.sol";
import "@openzeppelin/contracts@4.4.0/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts@4.4.0/token/ERC20/extensions/ERC20Pausable.sol";
import "@openzeppelin/contracts@4.4.0/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts@4.4.0/access/AccessControlEnumerable.sol";
import "@openzeppelin/contracts@4.4.0/utils/math/SafeMath.sol";

contract JRC21Issuer{
    using SafeMath for uint256;
    uint256 public _gas;
    address[] public _tokens;
    mapping(address => uint256) public _gasLefts;
    mapping(address => address) public _issuers;

	event Issue(address indexed issuer, address indexed token, uint256 value);
	event Charge(address indexed supporter, address indexed token, uint256 value);
	event Fee(address indexed sender, address indexed token, address indexed issuer, uint256 value);

    constructor(uint256 value){
        _gas = value;
    }

    function issueJRC21PresetFixed(
        string memory name,
        string memory symbol,
        address initialAccount,
        uint256 initialBalance,
        uint256 minFee
    ) public payable {
        JRC21PresetFixed _token = new JRC21PresetFixed(
            name,
            symbol,
            initialAccount,
            initialBalance,
            address(this),
            minFee);
        address token = address(_token);
        _tokens.push(token);
        _issuers[token] = msg.sender;
        _gasLefts[token] = _gasLefts[token].add(msg.value);
        emit Issue(msg.sender, token, msg.value);
    }

    function issueJRC21PresetMinter(
        string memory name,
        string memory symbol,
        uint256 minFee
    ) public payable {
        JRC21PresetMinter _token = new JRC21PresetMinter(
            name,
            symbol,
            address(this),
            minFee);
        address token = address(_token);
        _tokens.push(token);
        _issuers[token] = msg.sender;
        _gasLefts[token] = _gasLefts[token].add(msg.value);
        emit Issue(msg.sender, token, msg.value);
    }

    function charge(address token) public payable {
        require(msg.value > 0,"Cannot rechage 0");
        _gasLefts[token] = _gasLefts[token].add(msg.value);
        emit Charge(msg.sender, token, msg.value);
    }

    function transferWithPermit(address token,address recipient,uint256 value,uint256 fee,uint deadline, uint8 v, bytes32 r, bytes32 s)external{
        require(_gasLefts[token] > _gas,"JRC21 gas is insufficient");
        require(IJRC21(token).estimateFee(value) == fee,"Bad JRC21 fee");
        uint256 total = value + fee;
        IERC20Permit(token).permit(msg.sender,address(this),total,deadline,v,r,s);
        _gasLefts[token] = _gasLefts[token].sub(_gas);
        block.coinbase.transfer(_gas);
        SafeERC20.safeTransferFrom(IERC20(token),msg.sender,recipient,value);
        SafeERC20.safeTransferFrom(IERC20(token),msg.sender,_issuers[token],fee);

        emit Fee(msg.sender, token, _issuers[token], fee);
    }

	function transfer(address token,address recipient,uint256 value,uint256 fee)external{
        require(_gasLefts[token] > _gas,"JRC21 gas is insufficient");
        require(IJRC21(token).estimateFee(value) == fee,"Bad JRC21 fee");
        _gasLefts[token] = _gasLefts[token].sub(_gas);
        block.coinbase.transfer(_gas);
        SafeERC20.safeTransferFrom(IERC20(token),msg.sender,recipient,value);
        SafeERC20.safeTransferFrom(IERC20(token),msg.sender,_issuers[token],fee);

        emit Fee(msg.sender, token, _issuers[token], fee);
    }
}

interface IJRC21{
    function issuer() external view returns(address);
    function estimateFee(uint256 value) external view returns (uint256);
    function setMinFee(uint256 value)external;
}

contract JRC21 is AccessControlEnumerable{
    using SafeMath for uint256;
    uint256 public _minFee;
    address public _issuer;

    constructor(
        address issuer_,
        uint256 minFee_
    ) {
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _issuer = issuer_;
        _minFee = minFee_;
    }

    function getChainId() external view returns (uint256) {
        return block.chainid;
    }

    function estimateFee(uint256 value) external view returns (uint256){
        return value.mul(0).add(_minFee);
    }

    function setMinFee(uint256 value)external{
        require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "JRC21: must have admin role to set");
        _minFee = value;
    }

    function issuer() external view returns(address){
        return _issuer;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlEnumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}

contract JRC21PresetFixed is JRC21,ERC20,ERC20Permit,ERC20Pausable{
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    constructor(
        string memory name,
        string memory symbol,
        address initialAccount,
        uint256 initialBalance,
        address issuer,
        uint256 minFee
    ) ERC20(name, symbol) ERC20Permit(name) JRC21(issuer,minFee) {
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _setupRole(PAUSER_ROLE, _msgSender());
        _mint(initialAccount, initialBalance);
    }

    function pause() public virtual {
        require(hasRole(PAUSER_ROLE, _msgSender()), "JRC21PresetFixed: must have pauser role to pause");
        _pause();
    }

    function unpause() public virtual {
        require(hasRole(PAUSER_ROLE, _msgSender()), "JRC21PresetFixed: must have pauser role to unpause");
        _unpause();
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override(ERC20, ERC20Pausable) {
        super._beforeTokenTransfer(from, to, amount);
    }
}

contract JRC21PresetMinter is JRC21,ERC20,ERC20Permit,ERC20Burnable,ERC20Pausable{
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    constructor(
        string memory name,
        string memory symbol,
        address issuer,
        uint256 minFee
    ) ERC20(name, symbol) ERC20Permit(name) JRC21(issuer,minFee) {
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _setupRole(MINTER_ROLE, _msgSender());
        _setupRole(PAUSER_ROLE, _msgSender());
    }

    function mint(address to, uint256 amount) public virtual {
        require(hasRole(MINTER_ROLE, _msgSender()), "JRC21PresetMinter: must have minter role to mint");
        _mint(to, amount);
    }

    function pause() public virtual {
        require(hasRole(PAUSER_ROLE, _msgSender()), "JRC21PresetMinter: must have pauser role to pause");
        _pause();
    }

    function unpause() public virtual {
        require(hasRole(PAUSER_ROLE, _msgSender()), "JRC21PresetMinter: must have pauser role to unpause");
        _unpause();
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override(ERC20, ERC20Pausable) {
        super._beforeTokenTransfer(from, to, amount);
    }
}
