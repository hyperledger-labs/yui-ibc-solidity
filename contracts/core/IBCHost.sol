pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCStore.sol";

// IBCHost implements ics-024
abstract contract IBCHost {
    IBCStore ibcStore;

    constructor(IBCStore s) public {
        ibcStore = s;
    }
}
