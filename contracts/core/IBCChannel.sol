pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../lib/strings.sol";
import "./types/Channel.sol";
import "./types/Client.sol";
import "./ProvableStore.sol";
import "./IBCClient.sol";
import "./IBCConnection.sol";
import "./IBCMsgs.sol";

contract IBCChannel {
    using strings for *;

    ProvableStore provableStore;
    IBCClient ibcclient;
    IBCConnection ibcconnection;

    // types
    struct MsgChannelOpenInit {
        string portId;
        string channelId;
        Channel.Data channel;
    }

    struct MsgChannelOpenTry {
        string portId;
        string channelId;
        Channel.Data channel;
        string counterpartyVersion;
        bytes proofInit;
        uint64 proofHeight;
    }

    struct MsgChannelOpenAck {
        string portId;
        string channelId;
        string counterpartyVersion;
        string counterpartyChannelId;
        bytes proofTry;
        uint64 proofHeight;
    }

    struct MsgChannelOpenConfirm {
        string portId;
        string channelId;
        bytes proofAck;
        uint64 proofHeight;
    }

    constructor(ProvableStore store, IBCClient ibcclient_, IBCConnection ibcconnection_) public {
        provableStore = store;
        ibcclient = ibcclient_;
        ibcconnection = ibcconnection_;
    }

    function channelOpenInit(
        MsgChannelOpenInit memory msg_
    ) public returns (string memory) {
        require(!provableStore.hasChannel(msg_.portId, msg_.channelId), "channel already exists");
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(msg_.channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.versions.length == 1, "single version must be negotiated on connection before opening channel");
        require(msg_.channel.state == Channel.State.STATE_INIT, "channel state must STATE_INIT");

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        provableStore.setChannel(msg_.portId, msg_.channelId, msg_.channel);
        provableStore.setNextSequenceSend(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceRecv(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceAck(msg_.portId, msg_.channelId, 1);

        return msg_.channelId;
    }

    function channelOpenTry(
        MsgChannelOpenTry memory msg_
    ) public returns (string memory) {
        require(!provableStore.hasChannel(msg_.portId, msg_.channelId), "channel already exists");
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(msg_.channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.versions.length == 1, "single version must be negotiated on connection before opening channel");
        require(msg_.channel.state == Channel.State.STATE_TRYOPEN, "channel state must be STATE_TRYOPEN");

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        ChannelCounterparty.Data memory expectedCounterparty = ChannelCounterparty.Data({
            port_id: msg_.portId,
            channel_id: ""
        });
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_INIT,
            ordering: msg_.channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(msg_.channel),
            version: msg_.counterpartyVersion
        });
        require(ibcconnection.verifyChannelState(connection, msg_.proofHeight, msg_.proofInit, msg_.channel.counterparty.port_id, msg_.channel.counterparty.channel_id, Channel.encode(expectedChannel)), "failed to verify channel state");

        provableStore.setChannel(msg_.portId, msg_.channelId, msg_.channel);
        provableStore.setNextSequenceSend(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceRecv(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceAck(msg_.portId, msg_.channelId, 1);        

        return msg_.channelId;
    }

    function channelOpenAck(
        MsgChannelOpenAck memory msg_
    ) public {
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = provableStore.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_INIT || channel.state == Channel.State.STATE_TRYOPEN, "invalid channel state");

        // TODO authenticates a port binding

        (connection, found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        ChannelCounterparty.Data memory expectedCounterparty = ChannelCounterparty.Data({
            port_id: msg_.portId,
            channel_id: msg_.channelId
        });
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_TRYOPEN,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(channel),
            version: msg_.counterpartyVersion
        });
        require(ibcconnection.verifyChannelState(connection, msg_.proofHeight, msg_.proofTry, channel.counterparty.port_id, msg_.counterpartyChannelId, Channel.encode(expectedChannel)), "failed to verify channel state");
        channel.state = Channel.State.STATE_OPEN;
        channel.version = msg_.counterpartyVersion;
        channel.counterparty.channel_id = msg_.counterpartyChannelId;
        provableStore.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function channelOpenConfirm(
        MsgChannelOpenConfirm memory msg_
    ) public {
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = provableStore.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_TRYOPEN, "channel state is not TRYOPEN");

        // TODO authenticates a port binding

        (connection, found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        ChannelCounterparty.Data memory expectedCounterparty = ChannelCounterparty.Data({
            port_id: msg_.portId,
            channel_id: msg_.channelId
        });
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_OPEN,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(channel),
            version: channel.version
        });
        require(ibcconnection.verifyChannelState(connection, msg_.proofHeight, msg_.proofAck, channel.counterparty.port_id, channel.counterparty.channel_id, Channel.encode(expectedChannel)), "failed to verify channel state");
        channel.state = Channel.State.STATE_OPEN;
        provableStore.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function sendPacket(Packet.Data memory packet) public {
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        ClientState.Data memory clientState;
        uint64 latestTimestamp;
        uint64 nextSequenceSend;
        bool found;

        (channel, found) = provableStore.getChannel(packet.source_port, packet.source_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");
        require(packet.destination_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet destination port doesn't match the counterparty's port");
        require(packet.destination_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet destination channel doesn't match the counterparty's channel");
        (connection, found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        (clientState, found) = provableStore.getClientState(connection.client_id);
        require(found, "clientState not found");

        require(packet.timeout_height.revision_height == 0 || clientState.latest_height < packet.timeout_height.revision_height, "receiving chain block height >= packet timeout height");
        (latestTimestamp, found) = ibcclient.getTimestampAtHeight(connection.client_id, clientState.latest_height);
        require(found, "consensusState not found");
        require(packet.timeout_timestamp == 0 || latestTimestamp < packet.timeout_timestamp, "receiving chain block timestamp >= packet timeout timestamp");

        nextSequenceSend = provableStore.getNextSequenceSend(packet.source_port, packet.source_channel);
        require(nextSequenceSend > 0, "sequenceSend not found");
        require(packet.sequence == nextSequenceSend, "packet sequence ≠ next send sequence");

        nextSequenceSend++;
        provableStore.setNextSequenceSend(packet.source_port, packet.source_channel, nextSequenceSend);
        provableStore.setPacketCommitment(packet.source_port, packet.source_channel, packet.sequence, packet);

        // TODO emit an event that includes a packet
    }

    function recvPacket(IBCMsgs.MsgPacketRecv memory msg_) public {
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;
        (channel, found) = provableStore.getChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        // TODO
        // Authenticate capability to ensure caller has authority to receive packet on this channel

        require(msg_.packet.source_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet source port doesn't match the counterparty's port");
        require(msg_.packet.source_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet source channel doesn't match the counterparty's channel");

        (connection, found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        require(msg_.packet.timeout_height.revision_height == 0 || block.number < msg_.packet.timeout_height.revision_height, "block height >= packet timeout height");
        require(msg_.packet.timeout_timestamp == 0 || block.timestamp < msg_.packet.timeout_timestamp, "block timestamp >= packet timeout timestamp");

        bytes32 commitment = provableStore.makePacketCommitment(msg_.packet);
        require(ibcconnection.verifyPacketCommitment(connection, msg_.proofHeight, msg_.proof, msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence, commitment), "failed to verify packet commitment");

        if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            require(!provableStore.hasPacketReceipt(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence), "packet sequence already has been received");
            provableStore.setPacketReceipt(msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence);
        } else if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            uint64 nextSequenceRecv = provableStore.getNextSequenceRecv(msg_.packet.destination_port, msg_.packet.destination_channel);
            require(nextSequenceRecv > 0 && nextSequenceRecv == msg_.packet.sequence, "packet sequence ≠ next receive sequence");
            provableStore.setNextSequenceRecv(msg_.packet.destination_port, msg_.packet.destination_channel, nextSequenceRecv+1);
        } else {
            revert("unknown ordering type");
        }
    }

    // WriteAcknowledgement writes the packet execution acknowledgement to the state,
    // which will be verified by the counterparty chain using AcknowledgePacket.
    function writeAcknowledgement(Packet.Data memory packet, bytes memory acknowledgement) public {
        Channel.Data memory channel;
        bytes32 ackHash;
        bool found;
        (channel, found) = provableStore.getChannel(packet.destination_port, packet.destination_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        (ackHash, found) = provableStore.getPacketAcknowledgementCommitment(packet.destination_port, packet.destination_channel, packet.sequence);
        require(!found, "acknowledgement for packet already exists");

        require(acknowledgement.length > 0, "acknowledgement cannot be empty");
        provableStore.setPacketAcknowledgementCommitment(packet.destination_port, packet.destination_channel, packet.sequence, acknowledgement);
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement memory msg_) public {
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bytes32 commitment;
        uint64 nextSequenceAck;
        bool found;
        (channel, found) = provableStore.getChannel(msg_.packet.source_port, msg_.packet.source_channel);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_OPEN, "channel state must be OPEN");

        require(msg_.packet.destination_port.toSlice().equals(channel.counterparty.port_id.toSlice()), "packet destination port doesn't match the counterparty's port");
        require(msg_.packet.destination_channel.toSlice().equals(channel.counterparty.channel_id.toSlice()), "packet destination channel doesn't match the counterparty's channel");

        (connection, found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        (commitment, found) = provableStore.getPacketCommitment(msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence);
        require(found, "packet commitment not found");

        require(commitment == provableStore.makePacketCommitment(msg_.packet), "commitment bytes are not equal");

        require(ibcconnection.verifyPacketAcknowledgement(connection, msg_.proofHeight, msg_.proof, msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, provableStore.makePacketAcknowledgementCommitment(msg_.acknowledgement)), "failed to verify packet acknowledgement commitment");

        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            nextSequenceAck = provableStore.getNextSequenceAck(msg_.packet.source_port, msg_.packet.source_channel);
            require(nextSequenceAck == 0, "sequence ack not found");
            require(msg_.packet.sequence == nextSequenceAck, "packet sequence ≠ next ack sequence");
            nextSequenceAck++;
            provableStore.setNextSequenceAck(msg_.packet.source_port, msg_.packet.source_channel, nextSequenceAck);
        }

        provableStore.deletePacketCommitment(msg_.packet.source_port, msg_.packet.source_channel, msg_.packet.sequence);
    }

    function getCounterpartyHops(Channel.Data memory channel) internal view returns (string[] memory hops) {
        require(channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        hops = new string[](1);
        hops[0] = connection.counterparty.connection_id;
        return hops;
    }
}
