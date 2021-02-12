pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";
import "./ProvableStore.sol";
import "./IBCClient.sol";
import "./IBCConnection.sol";

contract IBCChannel {
    ProvableStore provableStore;
    IBCClient client;
    IBCConnection connection;

    constructor(ProvableStore store, IBCClient client_, IBCConnection connection_) public {
        provableStore = store;
        client = client_;
        connection = connection_;
    }

    function channelOpenInit(
        string memory channelId,
        Channel.Order order,
        string[] memory connectionHops,
        string memory portId,
        ChannelCounterparty.Data memory counterparty,
        string memory version
    ) public returns (string memory) {
        require(connectionHops.length == 1, "connectionHops length must be 1");
        (ConnectionEnd.Data memory connection, bool found) = provableStore.getConnection(connectionHops[0]);
        require(found, "connection not found");

        return channelId;
    }
}
