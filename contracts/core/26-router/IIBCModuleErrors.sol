// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Channel} from "../../proto/Channel.sol";

interface IIBCModuleErrors {
    /// @dev An error indicating that the sender is not the IBC contract
    /// @param sender Address of the sender
    error IBCModuleInvalidSender(address sender);

    /// @dev An error indicating that the channel ordering is not allowed
    /// @param portId Port identifier
    /// @param channelId Channel identifier
    /// @param order Channel ordering type
    error IBCModuleChannelOrderNotAllowed(string portId, string channelId, Channel.Order order);

    /// @dev An error indicating that the channel close is not allowed
    error IBCModuleChannelCloseNotAllowed(string portId, string channelId);
}
