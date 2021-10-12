pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";
import "./IBCConnection.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";

library IBCChannelHandshake {

    function channelOpenInit(
        IBCHost host,
        IBCMsgs.MsgChannelOpenInit memory msg_
    ) public returns (string memory) {
        host.onlyIBCModule();
        ConnectionEnd.Data memory connection;
        bool found;
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        (connection, found) = host.getConnection(msg_.channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.versions.length == 1, "single version must be negotiated on connection before opening channel");
        require(msg_.channel.state == Channel.State.STATE_INIT, "channel state must STATE_INIT");

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        string memory channelId = host.generateChannelIdentifier();
        host.setChannel(msg_.portId, channelId, msg_.channel);
        host.setNextSequenceSend(msg_.portId, channelId, 1);
        host.setNextSequenceRecv(msg_.portId, channelId, 1);
        host.setNextSequenceAck(msg_.portId, channelId, 1);

        return channelId;
    }

    function channelOpenTry(
        IBCHost host,
        IBCMsgs.MsgChannelOpenTry memory msg_
    ) public returns (string memory) {
        host.onlyIBCModule();
        ConnectionEnd.Data memory connection;
        bool found;
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        (connection, found) = host.getConnection(msg_.channel.connection_hops[0]);
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
            connection_hops: getCounterpartyHops(host, msg_.channel),
            version: msg_.counterpartyVersion
        });
        require(IBCConnection.verifyChannelState(host, connection, msg_.proofHeight, msg_.proofInit, msg_.channel.counterparty.port_id, msg_.channel.counterparty.channel_id, Channel.encode(expectedChannel)), "failed to verify channel state");

        string memory channelId = host.generateChannelIdentifier();
        host.setChannel(msg_.portId, channelId, msg_.channel);
        host.setNextSequenceSend(msg_.portId, channelId, 1);
        host.setNextSequenceRecv(msg_.portId, channelId, 1);
        host.setNextSequenceAck(msg_.portId, channelId, 1);

        return channelId;
    }

    function channelOpenAck(
        IBCHost host,
        IBCMsgs.MsgChannelOpenAck memory msg_
    ) public {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = host.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_INIT || channel.state == Channel.State.STATE_TRYOPEN, "invalid channel state");

        // TODO authenticates a port binding

        (connection, found) = host.getConnection(channel.connection_hops[0]);
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
            connection_hops: getCounterpartyHops(host, channel),
            version: msg_.counterpartyVersion
        });
        require(IBCConnection.verifyChannelState(host, connection, msg_.proofHeight, msg_.proofTry, channel.counterparty.port_id, msg_.counterpartyChannelId, Channel.encode(expectedChannel)), "failed to verify channel state");
        channel.state = Channel.State.STATE_OPEN;
        channel.version = msg_.counterpartyVersion;
        channel.counterparty.channel_id = msg_.counterpartyChannelId;
        host.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function channelOpenConfirm(
        IBCHost host,
        IBCMsgs.MsgChannelOpenConfirm memory msg_
    ) public {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = host.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state == Channel.State.STATE_TRYOPEN, "channel state is not TRYOPEN");

        // TODO authenticates a port binding

        (connection, found) = host.getConnection(channel.connection_hops[0]);
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
            connection_hops: getCounterpartyHops(host, channel),
            version: channel.version
        });
        require(IBCConnection.verifyChannelState(host, connection, msg_.proofHeight, msg_.proofAck, channel.counterparty.port_id, channel.counterparty.channel_id, Channel.encode(expectedChannel)), "failed to verify channel state");
        channel.state = Channel.State.STATE_OPEN;
        host.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function channelCloseInit(
        IBCHost host,
        IBCMsgs.MsgChannelCloseInit memory msg_
    ) public {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = host.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state != Channel.State.STATE_CLOSED, "channel state is already CLOSED");

        // TODO authenticates a port binding

        (connection, found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        channel.state = Channel.State.STATE_CLOSED;
        host.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function channelCloseConfirm(
        IBCHost host,
        IBCMsgs.MsgChannelCloseConfirm memory msg_
    ) public {
        host.onlyIBCModule();
        Channel.Data memory channel;
        ConnectionEnd.Data memory connection;
        bool found;

        (channel, found) = host.getChannel(msg_.portId, msg_.channelId);
        require(found, "channel not found");
        require(channel.state != Channel.State.STATE_CLOSED, "channel state is already CLOSED");

        // TODO authenticates a port binding

        (connection, found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.state == ConnectionEnd.State.STATE_OPEN, "connection state is not OPEN");

        ChannelCounterparty.Data memory expectedCounterparty = ChannelCounterparty.Data({
            port_id: msg_.portId,
            channel_id: msg_.channelId
        });
        Channel.Data memory expectedChannel = Channel.Data({
            state: Channel.State.STATE_CLOSED,
            ordering: channel.ordering,
            counterparty: expectedCounterparty,
            connection_hops: getCounterpartyHops(host, channel),
            version: channel.version
        });
        require(IBCConnection.verifyChannelState(host, connection, msg_.proofHeight, msg_.proofInit, channel.counterparty.port_id, channel.counterparty.channel_id, Channel.encode(expectedChannel)), "failed to verify channel state");
        channel.state = Channel.State.STATE_CLOSED;
        host.setChannel(msg_.portId, msg_.channelId, channel);
    }

    function getCounterpartyHops(IBCHost host, Channel.Data memory channel) internal view returns (string[] memory hops) {
        require(channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = host.getConnection(channel.connection_hops[0]);
        require(found, "connection not found");
        hops = new string[](1);
        hops[0] = connection.counterparty.connection_id;
        return hops;
    }

}
