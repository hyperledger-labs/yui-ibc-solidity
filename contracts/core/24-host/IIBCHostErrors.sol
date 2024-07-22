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

    /// @param portId port identifier
    error IBCHostModulePortNotFound(string portId);

    /// @param portId port identifier
    /// @param channelId channel identifier
    error IBCHostModuleChannelNotFound(string portId, string channelId);

    error IBCHostModuleDoesNotSupportERC165();

    /// @param interfaceId expected interface identifier
    error IBCHostModuleDoesNotSupportIIBCModule(bytes4 interfaceId);

    /// @param module module contract address
    error IBCHostModuleDoesNotSupportIIBCModuleUpgrade(address module);

    /// @param portId port identifier
    error IBCHostPortCapabilityAlreadyClaimed(string portId);

    /// @param portId port identifier
    /// @param channelId channel identifier
    error IBCHostChannelCapabilityAlreadyClaimed(string portId, string channelId);

    /// @param portId port identifier
    /// @param channelId channel identifier
    /// @param caller caller address
    error IBCHostFailedAuthenticateChannelCapability(string portId, string channelId, address caller);
}
