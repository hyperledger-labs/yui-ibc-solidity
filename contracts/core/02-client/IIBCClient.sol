// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";

interface IIBCClient {
    // --------------------- Data Structure --------------------- //

    struct MsgCreateClient {
        string clientType;
        bytes protoClientState;
        bytes protoConsensusState;
    }

    struct MsgUpdateClient {
        string clientId;
        bytes protoClientMessage;
    }

    // --------------------- Events --------------------- //

    /// @notice Emitted when a client identifier is generated
    /// @param clientId client identifier
    event GeneratedClientIdentifier(string clientId);

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
     *      Check ADR-001 for details.
     */
    function routeUpdateClient(MsgUpdateClient calldata msg_) external view returns (address, bytes4, bytes memory);

    /**
     * @dev updateClientCommitments updates the commitments of the light client's states corresponding to the given heights.
     *      Check ADR-001 for details.
     */
    function updateClientCommitments(string calldata clientId, Height.Data[] calldata heights) external;
}
