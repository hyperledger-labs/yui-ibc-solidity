// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

interface IIBCModuleManager {
    /// @notice Claim a capability for a port
    /// @param portId the port identifier
    /// @param module the module claiming the capability
    event IBCModuleManagerPortCapabilityClaimed(string portId, address module);
    /// @notice Emitted when a module claims a capability for a channel
    /// @param portId the port identifier
    /// @param channelId the channel identifier
    /// @param module the module claiming the capability
    event IBCModuleManagerChannelCapabilityClaimed(string portId, string channelId, address module);
}
