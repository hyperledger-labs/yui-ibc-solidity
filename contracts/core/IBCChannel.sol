pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./ProvableStore.sol";
import "./IBCClient.sol";

contract IBCChannel {
    ProvableStore provableStore;
    IBCClient client;

    constructor(ProvableStore store, IBCClient client_) public {
        provableStore = store;
        client = client_;
    }
}
