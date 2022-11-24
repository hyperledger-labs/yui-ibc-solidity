// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCModule.sol";
import "./IBCMsgs.sol";
import "./IBCCommitment.sol";
import "./IBCHost.sol";
import "./IIBCClient.sol";
import "./IIBCConnection.sol";
import "./IIBCChannel.sol";

contract IBCHandler is IBCHost {
    address public owner;

    /* Event definitions */

    event SendPacket(Packet.Data packet);
    event RecvPacket(Packet.Data packet);
    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );
    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);

    constructor(IIBCClient ibcClient, IIBCConnection ibcConnection, IIBCChannel ibcChannel) {
        owner = msg.sender;
        ibcClientAddress = address(ibcClient);
        ibcConnectionAddress = address(ibcConnection);
        ibcChannelAddress = address(ibcChannel);
    }

    /* Client/Module/Settings accessors */

    function registerClient(string calldata clientType, IClient client) external {
        onlyOwner();
        require(address(clientRegistry[clientType]) == address(0), "clientImpl already exists");
        clientRegistry[clientType] = address(client);
    }

    function bindPort(string calldata portId, address moduleAddress) external {
        onlyOwner();
        claimCapability(portCapabilityPath(portId), moduleAddress);
    }

    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external {
        // TODO: consider better authn/authz for this operation
        onlyOwner();
        expectedTimePerBlock = expectedTimePerBlock_;
    }

    /* Handshake interface */

    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external {
        (bool success,) = ibcClientAddress.delegatecall(abi.encodeWithSelector(IIBCClient.createClient.selector, msg_));
        require(success);
    }

    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external {
        (bool success,) = ibcClientAddress.delegatecall(abi.encodeWithSelector(IIBCClient.updateClient.selector, msg_));
        require(success);
    }

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_) external returns (string memory) {
        (bool success, bytes memory res) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IIBCConnection.connectionOpenInit.selector, msg_));
        require(success);
        return abi.decode(res, (string));
    }

    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry calldata msg_) external returns (string memory) {
        (bool success, bytes memory res) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IIBCConnection.connectionOpenTry.selector, msg_));
        require(success);
        return abi.decode(res, (string));
    }

    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck calldata msg_) external {
        (bool success,) =
            ibcConnectionAddress.delegatecall(abi.encodeWithSelector(IIBCConnection.connectionOpenAck.selector, msg_));
        require(success);
    }

    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm calldata msg_) external {
        (bool success,) = ibcConnectionAddress.delegatecall(
            abi.encodeWithSelector(IIBCConnection.connectionOpenConfirm.selector, msg_)
        );
        require(success);
    }

    function channelOpenInit(IBCMsgs.MsgChannelOpenInit calldata msg_) external returns (string memory) {
        (bool success, bytes memory res) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelOpenInit.selector, msg_));
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
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenTry(IBCMsgs.MsgChannelOpenTry calldata msg_) external returns (string memory) {
        string memory channelId;
        {
            // avoid "Stack too deep" error
            (bool success, bytes memory res) =
                ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelOpenTry.selector, msg_));
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
        claimCapability(channelCapabilityPath(msg_.portId, channelId), address(module));
        return channelId;
    }

    function channelOpenAck(IBCMsgs.MsgChannelOpenAck calldata msg_) external {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelOpenAck.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanOpenAck(msg_.portId, msg_.channelId, msg_.counterpartyVersion);
    }

    function channelOpenConfirm(IBCMsgs.MsgChannelOpenConfirm calldata msg_) external {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelOpenConfirm.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanOpenConfirm(msg_.portId, msg_.channelId);
    }

    function channelCloseInit(IBCMsgs.MsgChannelCloseInit calldata msg_) external {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelCloseInit.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanCloseInit(msg_.portId, msg_.channelId);
    }

    function channelCloseConfirm(IBCMsgs.MsgChannelCloseConfirm calldata msg_) external {
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.channelCloseConfirm.selector, msg_));
        require(success);
        lookupModuleByPortId(msg_.portId).onChanCloseConfirm(msg_.portId, msg_.channelId);
    }

    /* Packet handlers */

    function sendPacket(Packet.Data calldata packet) external {
        require(authenticateCapability(channelCapabilityPath(packet.source_port, packet.source_channel), msg.sender));
        (bool success,) =
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.sendPacket.selector, packet));
        require(success);
        emit SendPacket(packet);
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory acknowledgement) {
        IModuleCallbacks module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        acknowledgement = module.onRecvPacket(msg_.packet, msg.sender);
        (bool success,) = ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.recvPacket.selector, msg_));
        require(success);
        if (acknowledgement.length > 0) {
            (success,) = ibcChannelAddress.delegatecall(
                abi.encodeWithSelector(
                    IIBCChannel.writeAcknowledgement.selector,
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
        require(authenticateCapability(channelCapabilityPath(destinationPortId, destinationChannel), msg.sender));
        (bool success,) = ibcChannelAddress.delegatecall(
            abi.encodeWithSelector(
                IIBCChannel.writeAcknowledgement.selector,
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
            ibcChannelAddress.delegatecall(abi.encodeWithSelector(IIBCChannel.acknowledgePacket.selector, msg_));
        require(success);
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
    }

    /* Capabilities */

    function portCapabilityPath(string memory portId) public pure returns (bytes memory) {
        return abi.encodePacked(portId);
    }

    function channelCapabilityPath(string memory portId, string memory channelId) public pure returns (bytes memory) {
        return abi.encodePacked(portId, "/", channelId);
    }

    /* Internal functions */

    function lookupModuleByPortId(string memory portId) internal view returns (IModuleCallbacks) {
        (address module, bool found) = getModuleOwner(portCapabilityPath(portId));
        require(found);
        return IModuleCallbacks(module);
    }

    function lookupModuleByChannel(string memory portId, string memory channelId)
        internal
        view
        returns (IModuleCallbacks)
    {
        (address module, bool found) = getModuleOwner(channelCapabilityPath(portId, channelId));
        require(found);
        return IModuleCallbacks(module);
    }

    function onlyOwner() internal view {
        require(owner == msg.sender);
    }

    /* State accessors */

    function getClientState(string calldata clientId) external view returns (bytes memory, bool) {
        return getClient(clientId).getClientState(clientId);
    }

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (bytes memory consensusStateBytes, bool)
    {
        return getClient(clientId).getConsensusState(clientId, height);
    }

    function getConnection(string calldata connectionId) external view returns (ConnectionEnd.Data memory, bool) {
        ConnectionEnd.Data storage connection = connections[connectionId];
        return (connection, connection.state != ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getChannel(string calldata portId, string calldata channelId)
        external
        view
        returns (Channel.Data memory, bool)
    {
        Channel.Data storage channel = channels[portId][channelId];
        return (channel, channel.state != Channel.State.STATE_UNINITIALIZED_UNSPECIFIED);
    }

    function getPacketCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32, bool)
    {
        bytes32 commitment = commitments[keccak256(IBCCommitment.packetCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function getPacketAcknowledgementCommitment(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bytes32, bool)
    {
        bytes32 commitment =
            commitments[keccak256(IBCCommitment.packetAcknowledgementCommitmentPath(portId, channelId, sequence))];
        return (commitment, commitment != bytes32(0));
    }

    function hasPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (bool)
    {
        return packetReceipts[portId][channelId][sequence] == 1;
    }

    function getNextSequenceSend(string calldata portId, string calldata channelId) external view returns (uint64) {
        return nextSequenceSends[portId][channelId];
    }

    function getExpectedTimePerBlock() external view returns (uint64) {
        return expectedTimePerBlock;
    }
}
