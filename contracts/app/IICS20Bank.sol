pragma solidity ^0.6.8;

interface IICS20Bank {
    function transferFrom(address from, address to, string calldata id, uint256 amount) external;
    function mint(address account, string calldata id, uint256 amount) external;
    function burn(address account, string calldata id, uint256 amount) external;
}
