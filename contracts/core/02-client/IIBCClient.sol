// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

interface IIBCClient {
    struct MsgCreateClient {
        string clientType;
        bytes clientStateBytes;
        bytes consensusStateBytes;
    }

    struct MsgUpdateClient {
        string clientId;
        bytes clientMessage;
    }

    event GeneratedClientIdentifier(string);

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(MsgCreateClient calldata msg_) external returns (string memory clientId);

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(MsgUpdateClient calldata msg_) external;
}
