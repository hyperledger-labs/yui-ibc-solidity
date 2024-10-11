// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

interface IICS20Errors {
    /// @param version Version string
    error ICS20UnexpectedVersion(string version);
    /// @param actual port
    /// @param expected port
    error ICS20UnexpectedPort(string actual, string expected);
    /// @param tokenContract Address of the token contract
    error ICS20InvalidTokenContract(address tokenContract);
    /// @param tokenContract Address of the token contract
    /// @param from Address from which tokens are being transferred
    /// @param to Address to which tokens are being transferred
    /// @param amount Amount of tokens being transferred
    error ICS20FailedERC20Transfer(address tokenContract, address from, address to, uint256 amount);
    /// @param tokenContract Address of the token contract
    /// @param sender Address whose tokens are being transferred
    /// @param from Address from which tokens are being transferred
    /// @param to Address to which tokens are being transferred
    /// @param amount Amount of tokens being transferred
    error ICS20FailedERC20TransferFrom(address tokenContract, address sender, address from, address to, uint256 amount);
    /// @param sender Address whose tokens are being transferred
    /// @param balance Current balance for the interacting account
    /// @param needed Minimum amount required to perform a transfer
    error ICS20InsufficientBalance(address sender, uint256 balance, uint256 needed);
    /// @param receiver Address to which tokens are being transferred
    error ICS20InvalidReceiver(address receiver);
    /// @param sender Address of the sender
    error ICS20InvalidSenderAddress(string sender);
    /// @param receiver Address of the receiver
    error ICS20InvalidReceiverAddress(string receiver);
    /// @param position position in packet data bytes
    /// @param expected expected bytes
    /// @param actual actual bytes
    error ICS20JSONUnexpectedBytes(uint256 position, bytes32 expected, bytes32 actual);
    /// @param position position in packet data bytes
    /// @param actual actual value
    error ICS20JSONClosingBraceNotFound(uint256 position, bytes1 actual);
    /// @param position position in packet data bytes
    /// @param actual actual value
    error ICS20JSONStringClosingDoubleQuoteNotFound(uint256 position, bytes1 actual);
    /// @param bz json string value
    /// @param position position in packet data bytes
    error ICS20JSONStringUnclosed(bytes bz, uint256 position);
    /// @param position position in packet data bytes
    /// @param actual actual value
    error ICS20JSONInvalidEscape(uint256 position, bytes1 actual);
    /// @param length length of the slice
    error ICS20BytesSliceOverflow(uint256 length);
    /// @param length length of the bytes
    /// @param start start index
    /// @param end end index
    error ICS20BytesSliceOutOfBounds(uint256 length, uint256 start, uint256 end);
    /// @param escrow Address of the escrow contract
    /// @param sender Address of the sender
    /// @param denom Denomination of the tokens
    /// @param amount Amount of tokens
    error ICS20FailedRefund(address escrow, address sender, string denom, uint256 amount);
}
