// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";

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

    /**
     * @dev routeUpdateClient returns the lc contract address and the calldata to the receiving function of the client message.
     *      Light client contract may encode a client message as other encoding scheme(e.g. ethereum ABI)
     */
    function routeUpdateClient(MsgUpdateClient calldata msg_) external view returns (address, bytes4, bytes memory);

    /**
     * @dev updateClientCommitments updates the commitments of the light client's states corresponding to the given heights.
     */
    function updateClientCommitments(string calldata clientId, Height.Data[] calldata heights) external;
}
