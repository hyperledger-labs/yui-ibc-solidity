pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCClient.sol";
import "./IBCChannel.sol";
import "./IBCModule.sol";
import "./IBCMsgs.sol";
import "./IBCIdentifier.sol";
import "./types/Channel.sol";

contract IBCHandler {

    address owner;
    IBCHost host;

    /// Event definitions ///
    event SendPacket(Packet.Data packet);
    event RecvPacket(Packet.Data packet);
    event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement);
    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);

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
        IModuleCallbacks module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenInit(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry memory msg_) public returns (string memory) {
        string memory channelId = IBCChannel.channelOpenTry(host, msg_);
        IModuleCallbacks module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenTry(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version,
            msg_.counterpartyVersion
        );
        host.claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck memory msg_) public {
        IBCChannel.channelOpenAck(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanOpenAck(msg_.portId, msg_.channelId, msg_.counterpartyVersion);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm memory msg_) public {
        IBCChannel.channelOpenConfirm(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanOpenConfirm(msg_.portId, msg_.channelId);
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit memory msg_) public {
        IBCChannel.channelCloseInit(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanCloseInit(msg_.portId, msg_.channelId);
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm memory msg_) public {
        IBCChannel.channelCloseConfirm(host, msg_);
        lookupModuleByPortId(msg_.portId).onChanCloseConfirm(msg_.portId, msg_.channelId);
    }

    function sendPacket(Packet.Data calldata packet) external {
        require(host.authenticateCapability(
            IBCIdentifier.channelCapabilityPath(packet.source_port, packet.source_channel),
            msg.sender
        ));
        IBCChannel.sendPacket(host, packet);
        emit SendPacket(packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory acknowledgement) {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        acknowledgement = module.onRecvPacket(msg_.packet);
        IBCChannel.recvPacket(host, msg_);
        if (acknowledgement.length > 0) {
            IBCChannel.writeAcknowledgement(host, msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement);
            emit WriteAcknowledgement(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement);
        }
        emit RecvPacket(msg_.packet);
        return acknowledgement;
    }

    function writeAcknowledgement(string calldata destinationPortId, string calldata destinationChannel, uint64 sequence, bytes calldata acknowledgement) external {
        require(host.authenticateCapability(
            IBCIdentifier.channelCapabilityPath(destinationPortId, destinationChannel),
            msg.sender
        ));
        IBCChannel.writeAcknowledgement(host, destinationPortId, destinationChannel, sequence, acknowledgement);
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        module.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        IBCChannel.acknowledgePacket(host, msg_);
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
    }

    function bindPort(string memory portId, address moduleAddress) public {
        onlyOwner();
        host.claimCapability(IBCIdentifier.portCapabilityPath(portId), moduleAddress);
    }

    function lookupModuleByPortId(string memory portId) internal view returns (IModuleCallbacks) {
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.portCapabilityPath(portId));
        require(found);
        return IModuleCallbacks(module);
    }

    function lookupModuleByChannel(string memory portId, string memory channelId) internal view returns (IModuleCallbacks) {
        (address module, bool found) = host.getModuleOwner(IBCIdentifier.channelCapabilityPath(portId, channelId));
        require(found);
        return IModuleCallbacks(module);
    }

    function onlyOwner() internal view {
        require(msg.sender == owner);
    }
}
