// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

interface IIBCMockErrors {
    /// @dev An error indicating that the version is unexpected
    /// @param actual Actual version
    /// @param expected Expected version
    error IBCMockUnexpectedVersion(string actual, string expected);

    /// @dev An error indicating that the packet acknowledgement is unexpected
    error IBCMockUnexpectedAcknowledgement(bytes actual, bytes expected);

    /// @dev An error indicating that the packet is unexpected
    error IBCMockUnexpectedPacket(bytes actual);
}
