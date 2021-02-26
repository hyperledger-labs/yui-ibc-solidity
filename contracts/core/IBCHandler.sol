pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCClient.sol";
import "./IBCChannel.sol";
import "./IBCMsgs.sol";
import "./types/Channel.sol";
import "../lib/IBCIdentifier.sol";

contract IBCHandler {

    address owner;
    IBCHost host;

    constructor(IBCHost host_) public {
        owner = msg.sender;
        host = host_;
    }

    function getHostAddress() external view returns (address) {
        return address(host);
    }

    function registerClient(string calldata clientType, IClient client) external {
        require(msg.sender == owner);
        return IBCClient.registerClient(host, clientType, client);
    }

    /// Handler interface implementations ///

    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external {
        return IBCClient.createClient(host, msg_);
    }

    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        return IBCClient.updateClient(host, msg_);
    }

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit memory msg_) public returns (string memory) {
        return IBCConnection.connectionOpenInit(host, msg_);
    }

    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry memory msg_) public returns (string memory) {
        return IBCConnection.connectionOpenTry(host, msg_);
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck memory msg_) public {
        return IBCConnection.connectionOpenAck(host, msg_);
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm memory msg_) public {
        return IBCConnection.connectionOpenConfirm(host, msg_);
    }

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit memory msg_) public returns (string memory) {
        string memory channelId = IBCChannel.channelOpenInit(host, msg_);
        CallbacksI module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenInit(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            msg_.channelId,
            msg_.channel.counterparty,
            msg_.channel.version
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, msg_.channelId), address(module));
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry memory msg_) public returns (string memory) {
        string memory channelId = IBCChannel.channelOpenTry(host, msg_);
        CallbacksI module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenTry(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            msg_.channelId,
            msg_.channel.counterparty,
            msg_.channel.version,
            msg_.counterpartyVersion
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, msg_.channelId), address(module));
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck memory msg_) public {
        return IBCChannel.channelOpenAck(host, msg_);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm memory msg_) public {
        return IBCChannel.channelOpenConfirm(host, msg_);
    }

    function sendPacket(Packet.Data calldata packet) external {
        require(host.authenticateCapability(
            IBCIdentifier.channelCapabilityPath(packet.source_port, packet.source_channel),
            msg.sender
        ));
        IBCChannel.sendPacket(host, packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory) {
        CallbacksI module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        bytes memory acknowledgement = module.onRecvPacket(msg_.packet);
        IBCChannel.recvPacket(host, msg_);
        if (acknowledgement.length > 0) {
            IBCChannel.writeAcknowledgement(host, msg_.packet, acknowledgement);
        }
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        CallbacksI module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        module.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        IBCChannel.acknowledgePacket(host, msg_);
    }

    function bindPort(string memory portId, address moduleAddress) public {
        onlyOwner();
        host.claimCapability(IBCIdentifier.portCapabilityPath(portId), moduleAddress);
    }

    function lookupModuleByPortId(string memory portId) internal view returns (CallbacksI) {
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.portCapabilityPath(portId));
        require(found);
        return CallbacksI(module);
    }

    function lookupModuleByChannel(string memory portId, string memory channelId) internal view returns (CallbacksI) {
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.channelCapabilityPath(portId, channelId));
        require(found);
        return CallbacksI(module);
    }

    function onlyOwner() internal view {
        require(msg.sender == owner);
    }
}

interface CallbacksI {
    function onChanOpenInit(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version) external;
    function onChanOpenTry(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version, string calldata counterpartyVersion) external;

    function onRecvPacket(Packet.Data calldata) external returns(bytes memory);
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement) external;
}
