// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";

interface IIBCClientErrors {
    /// @param clientType the client type
    error IBCClientUnregisteredClientType(string clientType);

    /// @param clientId the client identifier
    error IBCClientClientNotFound(string clientId);

    /// @param clientId the client identifier
    /// @param consensusHeight the consensus height
    error IBCClientConsensusStateNotFound(string clientId, Height.Data consensusHeight);

    /// @param clientId the client identifier
    error IBCClientNotActiveClient(string clientId);

    /// @param selector the function selector
    /// @param args the calldata
    error IBCClientFailedUpdateClient(bytes4 selector, bytes args);

    /// @param commitmentKey the commitment key
    /// @param commitment the commitment
    /// @param prev the previous commitment
    error IBCClientInconsistentConsensusStateCommitment(bytes32 commitmentKey, bytes32 commitment, bytes32 prev);
}
