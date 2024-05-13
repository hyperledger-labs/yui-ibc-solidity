// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ICS20Lib} from "./ICS20Lib.sol";
import {IICS20Bank} from "./IICS20Bank.sol";
import {IICS20Errors} from "./IICS20Errors.sol";

contract ICS20Bank is Context, AccessControl, IICS20Bank, IICS20Errors {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    // Mapping from token ID to account balances
    mapping(string => mapping(address => uint256)) internal _balances;

    constructor() {
        _grantRole(ADMIN_ROLE, _msgSender());
    }

    function setOperator(address operator) public virtual {
        if (!hasRole(ADMIN_ROLE, _msgSender())) {
            revert ICS20BankNotAdminRole(_msgSender());
        }
        _grantRole(OPERATOR_ROLE, operator);
    }

    function balanceOf(address account, string calldata denom) public view virtual returns (uint256) {
        return _balances[denom][account];
    }

    function transferFrom(address from, address to, string calldata denom, uint256 amount) public virtual override {
        if (to == address(0)) {
            revert ICS20InvalidReceiver(to);
        } else if (from != _msgSender() && !hasRole(OPERATOR_ROLE, _msgSender())) {
            revert ICS20InvalidSender(from);
        }
        uint256 fromBalance = _balances[denom][from];
        if (fromBalance < amount) {
            revert ICS20InsufficientBalance(from, fromBalance, amount);
        }
        unchecked {
            _balances[denom][from] = fromBalance - amount;
        }
        _balances[denom][to] += amount;
    }

    function mint(address account, string calldata denom, uint256 amount) public virtual override {
        if (!hasRole(OPERATOR_ROLE, _msgSender())) {
            revert ICS20BankNotMintRole(_msgSender());
        }
        _mint(account, denom, amount);
    }

    function burn(address account, string calldata denom, uint256 amount) public virtual override {
        if (!hasRole(OPERATOR_ROLE, _msgSender())) {
            revert ICS20BankNotBurnRole(_msgSender());
        }
        _burn(account, denom, amount);
    }

    function deposit(address tokenContract, uint256 amount, address receiver) public virtual override {
        if (tokenContract == address(0)) {
            revert ICS20InvalidTokenContract(tokenContract);
        }
        if (!IERC20(tokenContract).transferFrom(_msgSender(), address(this), amount)) {
            revert ICS20FailedERC20Transfer(tokenContract, _msgSender(), address(this), amount);
        }
        _mint(receiver, addressToDenom(tokenContract), amount);
    }

    function withdraw(address tokenContract, uint256 amount, address receiver) public virtual override {
        if (tokenContract == address(0)) {
            revert ICS20InvalidTokenContract(tokenContract);
        }
        _burn(_msgSender(), addressToDenom(tokenContract), amount);
        if (!IERC20(tokenContract).transfer(receiver, amount)) {
            revert ICS20FailedERC20TransferFrom(tokenContract, _msgSender(), address(this), receiver, amount);
        }
    }

    function addressToDenom(address tokenContract) public pure virtual override returns (string memory) {
        return ICS20Lib.addressToHexString(tokenContract);
    }

    function _mint(address account, string memory denom, uint256 amount) internal {
        _balances[denom][account] += amount;
    }

    function _burn(address account, string memory denom, uint256 amount) internal {
        uint256 accountBalance = _balances[denom][account];
        if (accountBalance < amount) {
            revert ICS20InsufficientBalance(account, accountBalance, amount);
        }
        unchecked {
            _balances[denom][account] = accountBalance - amount;
        }
    }
}
