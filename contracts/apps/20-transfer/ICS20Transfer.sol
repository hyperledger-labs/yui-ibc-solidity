// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {Packet} from "../../core/04-channel/IIBCChannel.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {Channel} from "../../proto/Channel.sol";
import {ICS20Lib} from "./ICS20Lib.sol";
import {IICS20Errors} from "./IICS20Errors.sol";

abstract contract ICS20Transfer is IBCAppBase, IICS20Errors {
    string public constant ICS20_VERSION = "ics20-1";

    mapping(string => address) channelEscrowAddresses;

    function onRecvPacket(Packet calldata packet, address)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        ICS20Lib.PacketData memory data = ICS20Lib.unmarshalJSON(packet.data);
        bool success;
        address receiver;
        (receiver, success) = _decodeReceiver(data.receiver);
        if (!success) {
            return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
        }

        bytes memory denomPrefix = _getDenomPrefix(packet.sourcePort, packet.sourceChannel);
        bytes memory denom = bytes(data.denom);
        if (
            denom.length >= denomPrefix.length
                && ICS20Lib.equal(ICS20Lib.slice(denom, 0, denomPrefix.length), denomPrefix)
        ) {
            // sender chain is not the source, unescrow tokens
            bytes memory unprefixedDenom = ICS20Lib.slice(denom, denomPrefix.length, denom.length - denomPrefix.length);
            success = _tryTransferFrom(
                _getEscrowAddress(packet.destinationChannel), receiver, string(unprefixedDenom), data.amount
            );
        } else {
            // sender chain is the source, mint vouchers

            // ensure denom is not required to be escaped
            if (ICS20Lib.isEscapeNeededString(denom)) {
                success = false;
            } else {
                success = _tryMint(
                    receiver,
                    string(abi.encodePacked(_getDenomPrefix(packet.destinationPort, packet.destinationChannel), denom)),
                    data.amount
                );
            }
        }
        if (success) {
            return ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
        } else {
            return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
        }
    }

    function onAcknowledgementPacket(Packet calldata packet, bytes calldata acknowledgement, address)
        external
        virtual
        override
        onlyIBC
    {
        if (keccak256(acknowledgement) != ICS20Lib.KECCAK256_SUCCESSFUL_ACKNOWLEDGEMENT_JSON) {
            _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.sourcePort, packet.sourceChannel);
        }
    }

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        if (msg_.order != Channel.Order.ORDER_UNORDERED) {
            revert IBCModuleChannelOrderNotAllowed(msg_.portId, msg_.channelId, msg_.order);
        }
        bytes memory versionBytes = bytes(msg_.version);
        if (versionBytes.length != 0 && keccak256(versionBytes) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.version);
        }
        channelEscrowAddresses[msg_.channelId] = address(this);
        return ICS20_VERSION;
    }

    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        if (msg_.order != Channel.Order.ORDER_UNORDERED) {
            revert IBCModuleChannelOrderNotAllowed(msg_.portId, msg_.channelId, msg_.order);
        }
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.counterpartyVersion);
        }
        channelEscrowAddresses[msg_.channelId] = address(this);
        return ICS20_VERSION;
    }

    function onChanOpenAck(IIBCModule.MsgOnChanOpenAck calldata msg_) external virtual override onlyIBC {
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.counterpartyVersion);
        }
    }

    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata msg_) external virtual override onlyIBC {
        revert IBCModuleChannelCloseNotAllowed(msg_.portId, msg_.channelId);
    }

    function onTimeoutPacket(Packet calldata packet, address) external virtual override onlyIBC {
        _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.sourcePort, packet.sourceChannel);
    }

    function _getEscrowAddress(string memory sourceChannel) internal view virtual returns (address) {
        address escrow = channelEscrowAddresses[sourceChannel];
        if (escrow == address(0)) {
            revert ICS20EscrowAddressNotFound(sourceChannel);
        }
        return escrow;
    }

    function _refundTokens(ICS20Lib.PacketData memory data, string calldata sourcePort, string calldata sourceChannel)
        internal
        virtual
    {
        bytes memory denomPrefix = _getDenomPrefix(sourcePort, sourceChannel);
        bytes memory denom = bytes(data.denom);
        if (
            denom.length >= denomPrefix.length
                && ICS20Lib.equal(ICS20Lib.slice(denom, 0, denomPrefix.length), denomPrefix)
        ) {
            _mint(_decodeSender(data.sender), data.denom, data.amount);
        } else {
            // sender was source chain
            _transferFrom(_getEscrowAddress(sourceChannel), _decodeSender(data.sender), data.denom, data.amount);
        }
    }

    function _getDenomPrefix(string calldata port, string calldata channel) internal pure returns (bytes memory) {
        return abi.encodePacked(port, "/", channel, "/");
    }

    /**
     * @dev _transferFrom transfers tokens from `sender` to `receiver` in the bank.
     */
    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount) internal virtual;

    /**
     * @dev _burn burns tokens from `account` in the bank.
     */
    function _burn(address account, string memory denom, uint256 amount) internal virtual;

    /**
     * @dev _mint mints tokens to `account` in the bank.
     */
    function _mint(address account, string memory denom, uint256 amount) internal virtual;

    /**
     * @dev _tryTransferFrom transfers tokens from `sender` to `receiver` in the bank.
     *      If the transfer fails, it returns false and does not revert.
     */
    function _tryTransferFrom(address sender, address receiver, string memory denom, uint256 amount)
        internal
        virtual
        returns (bool);

    /**
     * @dev _tryMint mints tokens to `account` in the bank.
     *      If the mint fails, it returns false and does not revert.
     */
    function _tryMint(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    /**
     * @dev _tryBurn burns tokens from `account` in the bank.
     *      If the burn fails, it returns false and does not revert.
     */
    function _tryBurn(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    /**
     * @dev _encodeSender encodes an address to a hex string.
     *      The encoded sender is used as `sender` field in the packet data.
     */
    function _encodeSender(address sender) internal pure virtual returns (string memory) {
        return ICS20Lib.addressToHexString(sender);
    }

    /**
     * @dev _decodeSender decodes a hex string to an address.
     *      `sender` must be a valid address format.
     */
    function _decodeSender(string memory sender) internal pure virtual returns (address) {
        (address addr, bool ok) = ICS20Lib.hexStringToAddress(sender);
        if (!ok) {
            revert ICS20InvalidSenderAddress(sender);
        }
        return addr;
    }

    /**
     * @dev _decodeSender decodes a hex string to an address.
     *       `receiver` may be an invalid address format.
     */
    function _decodeReceiver(string memory receiver) internal pure virtual returns (address, bool) {
        return ICS20Lib.hexStringToAddress(receiver);
    }
}
