pragma solidity ^0.6.8;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

contract SimpleToken is ERC20 {
    constructor(
        string memory name,
        string memory symbol,
        uint256 initSupply
    )
        ERC20(name, symbol)
        public
    {
        _mint(msg.sender, initSupply);
    }
}
