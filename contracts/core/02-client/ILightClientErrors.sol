// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";

interface ILightClientErrors {
    /// @param caller the caller of the function
    error LightClientInvalidCaller(address caller);
    error LightClientNotActiveClient(string clientId);
    /// @param clientId client identifier
    error LightClientClientStateNotFound(string clientId);
    /// @param clientId client identifier
    /// @param height consensus height
    error LightClientConsensusStateNotFound(string clientId, Height.Data height);
    error LightClientConsensusStateExpired();
    /// @param url type url of the any
    error LightClientUnexpectedProtoAnyTypeURL(string url);
}
