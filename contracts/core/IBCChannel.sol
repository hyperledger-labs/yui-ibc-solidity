pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";
import "./ProvableStore.sol";
import "./IBCClient.sol";
import "./IBCConnection.sol";

contract IBCChannel {
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

    function getCounterpartyHops(Channel.Data memory channel) internal view returns (string[] memory hops) {
        require(channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        hops = new string[](1);
        hops[0] = connection.counterparty.connection_id;
        return hops;
    }

}
