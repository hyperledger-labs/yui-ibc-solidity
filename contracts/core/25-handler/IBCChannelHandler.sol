// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "../25-handler/IBCMsgs.sol";
import "../24-host/IBCHost.sol";
import "../04-channel/IIBCChannel.sol";
import "../05-port/IIBCModule.sol";
import "../05-port/ModuleManager.sol";

/**
 * @dev IBCChannelHandler is a contract that calls a contract that implements `IIBCChannelHandshake` with delegatecall.
 */
abstract contract IBCChannelHandler is ModuleManager {
    using Address for address;

    // IBC Channel contract address
    address immutable ibcChannel;

    event GeneratedChannelIdentifier(string);

    constructor(address _ibcChannel) {
        require(Address.isContract(_ibcChannel));
        ibcChannel = _ibcChannel;
    }

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit calldata msg_)
        external
        returns (string memory channelId, string memory version)
    {
        bytes memory res =
            ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenInit.selector, msg_));
        channelId = abi.decode(res, (string));

        IIBCModule module = lookupModuleByPort(msg_.portId);
        version = module.onChanOpenInit(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version
        );
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        emit GeneratedChannelIdentifier(channelId);
        return (channelId, version);
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_)
        external
        returns (string memory channelId, string memory version)
    {
        bytes memory res =
            ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenTry.selector, msg_));
        channelId = abi.decode(res, (string));
        IIBCModule module = lookupModuleByPort(msg_.portId);
        version = module.onChanOpenTry(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.counterpartyVersion
        );
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        emit GeneratedChannelIdentifier(channelId);
        return (channelId, version);
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenAck.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanOpenAck(msg_.portId, msg_.channelId, msg_.counterpartyVersion);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenConfirm.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanOpenConfirm(msg_.portId, msg_.channelId);
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelCloseInit.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanCloseInit(msg_.portId, msg_.channelId);
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelCloseConfirm.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanCloseConfirm(msg_.portId, msg_.channelId);
    }
}
