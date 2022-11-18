// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCClient.sol";
import "./IBCConnection.sol";
import "./IBCChannel.sol";
import "./IBCModule.sol";
import "./IBCMsgs.sol";
import "./IBCIdentifier.sol";
import "./IBCHost.sol";

contract IBCHandler is IBCHost {
    address immutable owner;

    /* Event definitions */

    event SendPacket(Packet.Data packet);
    event RecvPacket(Packet.Data packet);
    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );
    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);

    constructor(address ibcClientAddress_, address ibcConnectionAddress_, address ibcChannelAddress_) {
        owner = msg.sender;
        ibcClientAddress = ibcClientAddress_;
        ibcConnectionAddress = ibcConnectionAddress_;
        ibcChannelAddress = ibcChannelAddress_;
    }

    /* Client/Module/Settings accessors */

    function registerClient(string calldata clientType, IClient client) external {
        onlyOwner();
        require(address(clientRegistry[clientType]) == address(0), "clientImpl already exists");
        clientRegistry[clientType] = address(client);
    }

    function bindPort(string memory portId, address moduleAddress) public {
        onlyOwner();
        claimCapability(IBCIdentifier.portCapabilityPath(portId), moduleAddress);
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external {
        // TODO: consider better authn/authz for this operation
        onlyOwner();
        expectedTimePerBlock = expectedTimePerBlock_;
    }

    /* Handshake interface */

    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external {
        (bool success,) = ibcClientAddress.delegatecall(abi.encodeWithSelector(IBCClient.createClient.selector, msg_));
        require(success);
    }

    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        (bool success,) = ibcClientAddress.delegatecall(abi.encodeWithSelector(IBCClient.updateClient.selector, msg_));
        require(success);
    }

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_) public returns (string memory) {
        (bool success, bytes memory res) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IBCConnection.connectionOpenInit.selector, msg_));
        require(success);
        return abi.decode(res, (string));
    }

    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry calldata msg_) public returns (string memory) {
        (bool success, bytes memory res) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IBCConnection.connectionOpenTry.selector, msg_));
        require(success);
        return abi.decode(res, (string));
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck calldata msg_) public {
        (bool success,) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IBCConnection.connectionOpenAck.selector, msg_));
        require(success);
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm calldata msg_) public {
        (bool success,) = ibcConnectionAddress.delegatecall(
            abi.encodeWithSelector(IBCConnection.connectionOpenConfirm.selector, msg_)
        );
        require(success);
    }

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit calldata msg_) public returns (string memory) {
        (bool success, bytes memory res) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelOpenInit.selector, msg_));
        require(success);
        string memory channelId = abi.decode(res, (string));

        IModuleCallbacks module = lookupModuleByPortId(msg_.portId);
        module.onChanOpenInit(
            msg_.channel.ordering,
            msg_.channel.connection_hops,
            msg_.portId,
            channelId,
            msg_.channel.counterparty,
            msg_.channel.version
        );
        claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_) public returns (string memory) {
        string memory channelId;
        {
            // avoid "Stack too deep" error
            (bool success, bytes memory res) =
                ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelOpenTry.selector, msg_));
            require(success);
            channelId = abi.decode(res, (string));
        }
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
        claimCapability(IBCIdentifier.channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) public {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelOpenAck.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanOpenAck(msg_.portId, msg_.channelId, msg_.counterpartyVersion);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) public {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelOpenConfirm.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanOpenConfirm(msg_.portId, msg_.channelId);
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) public {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelCloseInit.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanCloseInit(msg_.portId, msg_.channelId);
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) public {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.channelCloseConfirm.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanCloseConfirm(msg_.portId, msg_.channelId);
    }

    /* Packet handlers */

    function sendPacket(Packet.Data calldata packet) external {
        require(
            authenticateCapability(
                IBCIdentifier.channelCapabilityPath(packet.source_port, packet.source_channel), msg.sender
            )
        );
        (bool success,) = ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.sendPacket.selector, packet));
        require(success);
        emit SendPacket(packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory acknowledgement) {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        acknowledgement = module.onRecvPacket(msg_.packet, msg.sender);
        (bool success,) = ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.recvPacket.selector, msg_));
        require(success);
        if (acknowledgement.length > 0) {
            (success,) = ibcChannelAddress.delegatecall(
                abi.encodeWithSelector(
                    IBCChannel.writeAcknowledgement.selector,
                    msg_.packet.destination_port,
                    msg_.packet.destination_channel,
                    msg_.packet.sequence,
                    acknowledgement
                )
            );
            require(success);
            emit WriteAcknowledgement(
                msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement
                );
        }
        emit RecvPacket(msg_.packet);
        return acknowledgement;
    }

    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external {
        require(
            authenticateCapability(
                IBCIdentifier.channelCapabilityPath(destinationPortId, destinationChannel), msg.sender
            )
        );
        (bool success,) = ibcChannelAddress.delegatecall(
            abi.encodeWithSelector(
                IBCChannel.writeAcknowledgement.selector,
                destinationPortId,
                destinationChannel,
                sequence,
                acknowledgement
            )
        );
        require(success);
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        module.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement, msg.sender);
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IBCChannel.acknowledgePacket.selector, msg_));
        require(success);
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
    }

    /* Internal functions */

    function lookupModuleByPortId(string memory portId) internal view returns (IModuleCallbacks) {
        (address module, bool found) = getModuleOwner(IBCIdentifier.portCapabilityPath(portId));
        require(found);
        return IModuleCallbacks(module);
    }

    function lookupModuleByChannel(string memory portId, string memory channelId)
        internal
        view
        returns (IModuleCallbacks)
    {
        (address module, bool found) = getModuleOwner(IBCIdentifier.channelCapabilityPath(portId, channelId));
        require(found);
        return IModuleCallbacks(module);
    }

    function onlyOwner() internal view {
        require(owner == msg.sender);
    }
}
