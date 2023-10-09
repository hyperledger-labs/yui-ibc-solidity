// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "./IICS20Bank.sol";

contract ICS20Bank is Context, AccessControl, IICS20Bank {
    using Address for address;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    // Mapping from token ID to account balances
    mapping(string => mapping(address => uint256)) private _balances;

    constructor() {
        _setupRole(ADMIN_ROLE, _msgSender());
    }

    function setOperator(address operator) public virtual {
        require(hasRole(ADMIN_ROLE, _msgSender()), "must have admin role to set new operator");
        _setupRole(OPERATOR_ROLE, operator);
    }

    function balanceOf(address account, string calldata denom) external view virtual returns (uint256) {
        require(account != address(0), "ICS20Bank: balance query for the zero address");
        return _balances[denom][account];
    }

    function transferFrom(address from, address to, string calldata denom, uint256 amount) external virtual override {
        require(to != address(0), "ICS20Bank: transfer to the zero address");
        require(
            from == _msgSender() || hasRole(OPERATOR_ROLE, _msgSender()), "ICS20Bank: caller is not owner nor approved"
        );
        uint256 fromBalance = _balances[denom][from];
        require(fromBalance >= amount, "ICS20Bank: insufficient balance for transfer");
        unchecked {
            _balances[denom][from] = fromBalance - amount;
        }
        _balances[denom][to] += amount;
    }

    function mint(address account, string calldata denom, uint256 amount) external virtual override {
        require(hasRole(OPERATOR_ROLE, _msgSender()), "ICS20Bank: must have minter role to mint");
        _mint(account, denom, amount);
    }

    function burn(address account, string calldata denom, uint256 amount) external virtual override {
        require(hasRole(OPERATOR_ROLE, _msgSender()), "ICS20Bank: must have minter role to mint");
        _burn(account, denom, amount);
    }

    function addressToDenom(address tokenContract) public pure virtual override returns (string memory) {
        return Strings.toHexString(tokenContract);
    }

    function deposit(address tokenContract, uint256 amount, address receiver) external virtual override {
        require(tokenContract.isContract());
        require(IERC20(tokenContract).transferFrom(_msgSender(), address(this), amount));
        _mint(receiver, addressToDenom(tokenContract), amount);
    }

    function withdraw(address tokenContract, uint256 amount, address receiver) external virtual override {
        require(tokenContract.isContract());
        _burn(_msgSender(), addressToDenom(tokenContract), amount);
        require(IERC20(tokenContract).transfer(receiver, amount));
    }

    function _mint(address account, string memory denom, uint256 amount) internal virtual {
        _balances[denom][account] += amount;
    }

    function _burn(address account, string memory denom, uint256 amount) internal virtual {
        uint256 accountBalance = _balances[denom][account];
        require(accountBalance >= amount, "ICS20Bank: burn amount exceeds balance");
        unchecked {
            _balances[denom][account] = accountBalance - amount;
        }
    }
}
