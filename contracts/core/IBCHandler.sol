pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCClient.sol";
import "./IBCChannel.sol";
import "./IBCMsgs.sol";
import "./types/Channel.sol";
import "../lib/IBCIdentifier.sol";

contract IBCHandler {
    // TODO move it into IBCStore
    mapping(string => Module) modules;

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
        // TODO call the binded module's callback `channelOpenInit`
        return IBCChannel.channelOpenInit(host, msg_);
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry memory msg_) public returns (string memory) {
        // TODO call the binded module's callback `channelOpenTry`
        return IBCChannel.channelOpenTry(host, msg_);
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck memory msg_) public {
        return IBCChannel.channelOpenAck(host, msg_);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm memory msg_) public {
        return IBCChannel.channelOpenConfirm(host, msg_);
    }

    function sendPacket(Packet.Data calldata packet) external {
        // require(host.authenticateCapability(
        //     IBCIdentifier.channelCapabilityPath(packet.source_port, packet.source_channel),
        //     msg.sender
        // ));
        IBCChannel.sendPacket(host, packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        bytes memory acknowledgement = module.callbacks.onRecvPacket(msg_.packet);
        IBCChannel.recvPacket(host, msg_);
        if (acknowledgement.length > 0) {
            IBCChannel.writeAcknowledgement(host, msg_.packet, acknowledgement);
        }
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        (Module memory module, bool found) = lookupModule(msg_.packet.source_port);
        require(found, "module not found");
        module.callbacks.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        IBCChannel.acknowledgePacket(host, msg_);
    }

    // WARNING: This function **must be** removed in production
    function handlePacketRecvWithoutVerification(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        return module.callbacks.onRecvPacket(msg_.packet);
    }

    /// Module manager ///

    struct Module {
        CallbacksI callbacks;
        bool exists;
    }

    // TODO apply ACL to this
    function bindPort(string memory portId, address moduleAddress) public {
        require(!modules[portId].exists, "the portId is already used by other module");
        modules[portId] = Module({callbacks: CallbacksI(moduleAddress), exists: true});
    }

    function lookupModule(string memory portId) public view returns (Module memory module, bool found) {
        if (!modules[portId].exists) {
            return (module, false);
        }
        return (modules[portId], true);
    }
}

interface CallbacksI {
    function onChanOpenInit(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version) external;
    function onChanOpenTry(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version, string calldata counterpartyVersion) external;

    function onRecvPacket(Packet.Data calldata) external returns(bytes memory);
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement) external;
}
