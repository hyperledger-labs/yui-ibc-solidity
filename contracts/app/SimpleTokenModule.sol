pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/types/Channel.sol";

contract SimpleTokenModule {
    // Token storages
    mapping (address => uint256) _balances;

    // Module storages
    address ibcRoutingModuleAddress;

    constructor(address ibcRoutingModuleAddress_) public {
        ibcRoutingModuleAddress = ibcRoutingModuleAddress_;
        _balances[msg.sender] = 10000;
    }

    modifier onlyRoutingModule (){
        require(msg.sender == ibcRoutingModuleAddress);
        _;
    }

    /// Token implementations ///

    function transfer(address recipient, uint256 amount) public returns (bool) {
        uint256 senderBalance = _balances[msg.sender];
        require(senderBalance >= amount, "transfer amount exceeds balance");
        _balances[msg.sender] = senderBalance - amount;
        _balances[recipient] += amount;
        return true;
    }

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    /// Module implementations ///

    function onRecvPacket(Packet.Data calldata packet) onlyRoutingModule external returns (bytes memory) {
        // TODO implements
        return packet.data;
    }
}
