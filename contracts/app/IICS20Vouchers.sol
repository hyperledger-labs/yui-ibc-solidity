pragma solidity ^0.6.8;

import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

interface IICS20Vouchers {
    function transferFrom(address from, address to, bytes calldata id, uint256 amount) external;
    function mint(address to, bytes calldata id, uint256 amount) external;
    function burnFrom(address from, bytes calldata id, uint256 amount) external;
}
