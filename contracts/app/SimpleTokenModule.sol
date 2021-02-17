pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/types/Channel.sol";
import "../core/IBCRoutingModule.sol";

contract SimpleTokenModule {
    /// Storages ///

    // Token storages
    mapping (address => uint256) _balances;

    // Module storages
    IBCRoutingModule ibcRoutingModule;

    /// Constructor ///

    constructor(IBCRoutingModule ibcRoutingModule_) public {
        ibcRoutingModule = ibcRoutingModule_;
        ibcRoutingModule.bindPort("bank", address(this));

        _balances[msg.sender] = 10000;
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

    modifier onlyRoutingModule (){
        require(msg.sender == address(ibcRoutingModule));
        _;
    }

    function onRecvPacket(Packet.Data calldata packet) onlyRoutingModule external returns (bytes memory) {
        // TODO implements
        return packet.data;
    }
}
