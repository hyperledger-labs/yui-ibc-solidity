// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {
    IIBCChannelHandshake, IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout
} from "../04-channel/IIBCChannel.sol";
import {IIBCChannelUpgrade} from "../04-channel/IIBCChannelUpgrade.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";
import {IBCHandler} from "./IBCHandler.sol";

/**
 * @dev OwnableIBCHandler is a contract that implements [ICS-25](https://github.com/cosmos/ibc/tree/main/spec/core/ics-025-handler-interface) and Ownable for host configuration.
 */
contract OwnableIBCHandler is IBCHandler, Ownable {
    /**
     * @dev The arguments of constructor must satisfy the followings:
     * @param ibcClient_ is the address of a contract that implements `IIBCClient`.
     * @param ibcConnection_ is the address of a contract that implements `IIBCConnection`.
     * @param ibcChannelHandshake_ is the address of a contract that implements `IIBCChannelHandshake`.
     * @param ibcChannelPacketSendRecv_ is the address of a contract that implements `IIBCChannelPacketSendRecv`.
     * @param ibcChannelPacketTimeout_ is the address of a contract that implements `IIBCChannelPacketTimeout`.
     * @param ibcChannelUpgrade_ is the address of a contract that implements `IIBCChannelUpgrade`.
     */
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgrade ibcChannelUpgrade_
    )
        IBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgrade_
        )
        Ownable(msg.sender)
    {}

    function registerClient(string calldata clientType, ILightClient client) public onlyOwner {
        super._registerClient(clientType, client);
    }

    function bindPort(string calldata portId, IIBCModule moduleAddress) public onlyOwner {
        super._bindPort(portId, moduleAddress);
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) public onlyOwner {
        super._setExpectedTimePerBlock(expectedTimePerBlock_);
    }
}
