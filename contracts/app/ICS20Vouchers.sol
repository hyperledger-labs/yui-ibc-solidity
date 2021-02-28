pragma solidity ^0.6.8;

import "openzeppelin-solidity/contracts/utils/Context.sol";
import "openzeppelin-solidity/contracts/access/AccessControl.sol";
import "./IICS20Vouchers.sol";

contract ICS20Vouchers is Context, AccessControl, IICS20Vouchers {

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    // Mapping from token ID to account balances
    mapping(bytes => mapping(address => uint256)) private _balances;

    constructor() public {
        _setupRole(ADMIN_ROLE, _msgSender());
    }

    function balanceOf(address account, bytes calldata id) external view returns (uint256) {
        require(account != address(0), "ERC1155: balance query for the zero address");
        return _balances[id][account];
    }

    function setOperator(address operator) public {
        require(hasRole(ADMIN_ROLE, _msgSender()), "must have admin role to set new operator");
        _setupRole(OPERATOR_ROLE, operator);
    }

    function transferFrom(address from, address to, bytes calldata id, uint256 amount) override external {
        require(to != address(0), "ERC1155: transfer to the zero address");
        require(
            from == _msgSender() || hasRole(OPERATOR_ROLE, from),
            "ERC1155: caller is not owner nor approved"
        );

        uint256 fromBalance = _balances[id][from];
        require(fromBalance >= amount, "Vouchers: insufficient balance for transfer");
        _balances[id][from] = fromBalance - amount;
        _balances[id][to] += amount;
    }

    function mint(address account, bytes calldata id, uint256 amount) override external {
        require(hasRole(OPERATOR_ROLE, _msgSender()), "Vouchers: must have minter role to mint");
        _balances[id][account] += amount;
    }

    function burnFrom(address account, bytes calldata id, uint256 amount) override external {
        require(hasRole(OPERATOR_ROLE, _msgSender()), "Vouchers: must have minter role to mint");
        uint256 accountBalance = _balances[id][account];
        require(accountBalance >= amount, "Vouchers: burn amount exceeds balance");
        _balances[id][account] = accountBalance - amount;
    }

}
