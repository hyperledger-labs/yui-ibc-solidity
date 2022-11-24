// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IBCMsgs.sol";

interface IIBCClient {
    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external;

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external;
}
