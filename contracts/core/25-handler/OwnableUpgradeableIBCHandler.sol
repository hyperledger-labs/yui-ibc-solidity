// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCConnection} from "../03-connection/IIBCConnection.sol";
import {
    IIBCChannelHandshake, IIBCChannelPacketSendRecv, IIBCChannelPacketTimeout
} from "../04-channel/IIBCChannel.sol";
import {
    IIBCChannelUpgradeInitTryAck,
    IIBCChannelUpgradeConfirmOpenTimeoutCancel
} from "../04-channel/IIBCChannelUpgrade.sol";
import {IIBCModuleInitializer} from "../26-router/IIBCModule.sol";
import {IBCHandler} from "./IBCHandler.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {ContextUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract OwnableUpgradeableIBCHandler is IBCHandler, UUPSUpgradeable, OwnableUpgradeable {
    /**
     * @dev The arguments of constructor must satisfy the followings:
     * @param ibcClient_ is the address of a contract that implements `IIBCClient`.
     * @param ibcConnection_ is the address of a contract that implements `IIBCConnection`.
     * @param ibcChannelHandshake_ is the address of a contract that implements `IIBCChannelHandshake`.
     * @param ibcChannelPacketSendRecv_ is the address of a contract that implements `IIBCChannelPacketSendRecv`.
     * @param ibcChannelPacketTimeout_ is the address of a contract that implements `IIBCChannelPacketTimeout`.
     * @param ibcChannelUpgradeInitTryAck_ is the address of a contract that implements `IIBCChannelUpgrade`.
     * @param ibcChannelUpgradeConfirmOpenTimeoutCancel_ is the address of a contract that implements `IIBCChannelUpgrade`.
     * @custom:oz-upgrades-unsafe-allow constructor
     */
    constructor(
        IIBCClient ibcClient_,
        IIBCConnection ibcConnection_,
        IIBCChannelHandshake ibcChannelHandshake_,
        IIBCChannelPacketSendRecv ibcChannelPacketSendRecv_,
        IIBCChannelPacketTimeout ibcChannelPacketTimeout_,
        IIBCChannelUpgradeInitTryAck ibcChannelUpgradeInitTryAck_,
        IIBCChannelUpgradeConfirmOpenTimeoutCancel ibcChannelUpgradeConfirmOpenTimeoutCancel_
    )
        IBCHandler(
            ibcClient_,
            ibcConnection_,
            ibcChannelHandshake_,
            ibcChannelPacketSendRecv_,
            ibcChannelPacketTimeout_,
            ibcChannelUpgradeInitTryAck_,
            ibcChannelUpgradeConfirmOpenTimeoutCancel_
        )
    {}

    function initialize() public virtual initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
    }

    function registerClient(string calldata clientType, ILightClient client) public virtual onlyOwner {
        super._registerClient(clientType, client);
    }

    function bindPort(string calldata portId, IIBCModuleInitializer moduleAddress) public virtual onlyOwner {
        super._bindPort(portId, moduleAddress);
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) public virtual onlyOwner {
        super._setExpectedTimePerBlock(expectedTimePerBlock_);
    }

    function _msgSender() internal view virtual override(Context, ContextUpgradeable) returns (address) {
        return ContextUpgradeable._msgSender();
    }

    function _msgData() internal view virtual override(Context, ContextUpgradeable) returns (bytes calldata) {
        return ContextUpgradeable._msgData();
    }

    function _contextSuffixLength() internal view virtual override(Context, ContextUpgradeable) returns (uint256) {
        return ContextUpgradeable._contextSuffixLength();
    }

    function _authorizeUpgrade(address newImplementation) internal virtual override onlyOwner {}
}
