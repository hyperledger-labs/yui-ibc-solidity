// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

interface IIBCHostErrors {
    /// @param clientType client type
    error IBCHostInvalidClientType(string clientType);

    /// @param clientType client type
    error IBCHostClientTypeAlreadyExists(string clientType);

    /// @param clientId the client identifier
    error IBCHostClientNotFound(string clientId);

    /// @param portId port identifier
    error IBCHostInvalidPortIdentifier(string portId);

    /// @param lcAddress light client contract address
    error IBCHostInvalidLightClientAddress(address lcAddress);

    /// @param moduleAddress module contract address
    error IBCHostInvalidModuleAddress(address moduleAddress);

    /// @param name module name
    error IBCHostModuleNotFound(string name);

    /// @param name capability name
    error IBCHostCapabilityAlreadyClaimed(string name);

    /// @param name capability name
    /// @param caller caller address
    error IBCHostFailedAuthenticateCapability(string name, address caller);
}
