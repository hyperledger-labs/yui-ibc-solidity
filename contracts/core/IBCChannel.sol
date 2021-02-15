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
        string channelId;
        string portId;
        Channel.Data channel;
    }

    constructor(ProvableStore store, IBCClient ibcclient_, IBCConnection ibcconnection_) public {
        provableStore = store;
        ibcclient = ibcclient_;
        ibcconnection = ibcconnection_;
    }

    function channelOpenInit(
        MsgChannelOpenInit memory msg_
    ) public returns (string memory) {
        require(msg_.channel.connection_hops.length == 1, "connection_hops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(msg_.channel.connection_hops[0]);
        require(found, "connection not found");
        require(connection.versions.length == 1, "single version must be negotiated on connection before opening channel");

        // TODO verifySupportedFeature

        // TODO authenticates a port binding

        provableStore.setChannel(msg_.channelId, msg_.channel);
        provableStore.setNextSequenceSend(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceRecv(msg_.portId, msg_.channelId, 1);
        provableStore.setNextSequenceAck(msg_.portId, msg_.channelId, 1);

        return msg_.channelId;
    }

}
