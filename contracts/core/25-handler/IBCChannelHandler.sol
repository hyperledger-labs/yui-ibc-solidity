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
        channelId = abi.decode(
            ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenInit.selector, msg_)),
            (string)
        );
        IIBCModule module = lookupModuleByPort(msg_.portId);
        version = module.onChanOpenInit(
            IIBCModule.MsgOnChanOpenInit({
                order: msg_.channel.ordering,
                connectionHops: msg_.channel.connection_hops,
                portId: msg_.portId,
                channelId: channelId,
                counterparty: msg_.channel.counterparty,
                version: msg_.channel.version
            })
        );
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        ibcChannel.functionDelegateCall(
            abi.encodeWithSelector(
                IIBCChannelHandshake.writeChannel.selector,
                msg_.portId,
                channelId,
                msg_.channel.state,
                msg_.channel.ordering,
                msg_.channel.counterparty,
                msg_.channel.connection_hops,
                version
            )
        );
        emit GeneratedChannelIdentifier(channelId);
        return (channelId, version);
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_)
        external
        returns (string memory channelId, string memory version)
    {
        channelId = abi.decode(
            ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenTry.selector, msg_)),
            (string)
        );
        IIBCModule module = lookupModuleByPort(msg_.portId);
        version = module.onChanOpenTry(
            IIBCModule.MsgOnChanOpenTry({
                order: msg_.channel.ordering,
                connectionHops: msg_.channel.connection_hops,
                portId: msg_.portId,
                channelId: channelId,
                counterparty: msg_.channel.counterparty,
                counterpartyVersion: msg_.counterpartyVersion
            })
        );
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        ibcChannel.functionDelegateCall(
            abi.encodeWithSelector(
                IIBCChannelHandshake.writeChannel.selector,
                msg_.portId,
                channelId,
                msg_.channel.state,
                msg_.channel.ordering,
                msg_.channel.counterparty,
                msg_.channel.connection_hops,
                version
            )
        );
        emit GeneratedChannelIdentifier(channelId);
        return (channelId, version);
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenAck.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanOpenAck(
            IIBCModule.MsgOnChanOpenAck({
                portId: msg_.portId,
                channelId: msg_.channelId,
                counterpartyVersion: msg_.counterpartyVersion
            })
        );
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelOpenConfirm.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanOpenConfirm(
            IIBCModule.MsgOnChanOpenConfirm({portId: msg_.portId, channelId: msg_.channelId})
        );
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelCloseInit.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanCloseInit(
            IIBCModule.MsgOnChanCloseInit({portId: msg_.portId, channelId: msg_.channelId})
        );
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) external {
        ibcChannel.functionDelegateCall(abi.encodeWithSelector(IIBCChannelHandshake.channelCloseConfirm.selector, msg_));
        lookupModuleByPort(msg_.portId).onChanCloseConfirm(
            IIBCModule.MsgOnChanCloseConfirm({portId: msg_.portId, channelId: msg_.channelId})
        );
    }
}
