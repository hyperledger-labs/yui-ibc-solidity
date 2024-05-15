// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";

interface ILightClientErrors {
    /// @param caller the caller of the function
    error InvalidCaller(address caller);
    error NotActiveClient(string clientId);
    /// @param clientId client identifier
    error ClientStateNotFound(string clientId);
    /// @param clientId client identifier
    /// @param height consensus height
    error ConsensusStateNotFound(string clientId, Height.Data height);
    error ConsensusStateExpired();
    /// @param url type url of the any
    error UnexpectedProtoAnyTypeURL(string url);
}
