// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {Channel, Packet} from "../../proto/Channel.sol";
import {ICS20Lib} from "./ICS20Lib.sol";

abstract contract ICS20Transfer is IBCAppBase {
    using BytesLib for bytes;

    string public constant ICS20_VERSION = "ics20-1";

    mapping(string => address) channelEscrowAddresses;

    function onRecvPacket(Packet.Data calldata packet, address)
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

        bytes memory denomPrefix = _getDenomPrefix(packet.source_port, packet.source_channel);
        bytes memory denom = bytes(data.denom);
        if (denom.length >= denomPrefix.length && denom.slice(0, denomPrefix.length).equal(denomPrefix)) {
            // sender chain is not the source, unescrow tokens
            bytes memory unprefixedDenom = denom.slice(denomPrefix.length, denom.length - denomPrefix.length);
            success = _transferFrom(
                _getEscrowAddress(packet.destination_channel), receiver, string(unprefixedDenom), data.amount
            );
        } else {
            // sender chain is the source, mint vouchers

            // ensure denom is not required to be escaped
            if (ICS20Lib.isEscapeNeededString(denom)) {
                success = false;
            } else {
                success = _mint(
                    receiver,
                    string(
                        abi.encodePacked(_getDenomPrefix(packet.destination_port, packet.destination_channel), denom)
                    ),
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

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement, address)
        external
        virtual
        override
        onlyIBC
    {
        if (keccak256(acknowledgement) != ICS20Lib.KECCAK256_SUCCESSFUL_ACKNOWLEDGEMENT_JSON) {
            _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
        }
    }

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        require(msg_.order == Channel.Order.ORDER_UNORDERED, "must be unordered");
        bytes memory versionBytes = bytes(msg_.version);
        require(versionBytes.length == 0 || keccak256(versionBytes) == keccak256(bytes(ICS20_VERSION)));
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
        require(msg_.order == Channel.Order.ORDER_UNORDERED, "must be unordered");
        require(keccak256(bytes(msg_.counterpartyVersion)) == keccak256(bytes(ICS20_VERSION)));
        channelEscrowAddresses[msg_.channelId] = address(this);
        return ICS20_VERSION;
    }

    function onChanOpenAck(IIBCModule.MsgOnChanOpenAck calldata msg_) external virtual override onlyIBC {
        require(keccak256(bytes(msg_.counterpartyVersion)) == keccak256(bytes(ICS20_VERSION)));
    }

    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata) external virtual override onlyIBC {
        revert("not allowed");
    }

    function onTimeoutPacket(Packet.Data calldata packet, address) external virtual override onlyIBC {
        _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
    }

    function _getEscrowAddress(string memory sourceChannel) internal view virtual returns (address) {
        address escrow = channelEscrowAddresses[sourceChannel];
        require(escrow != address(0));
        return escrow;
    }

    function _refundTokens(ICS20Lib.PacketData memory data, string calldata sourcePort, string calldata sourceChannel)
        internal
        virtual
    {
        bytes memory denomPrefix = _getDenomPrefix(sourcePort, sourceChannel);
        bytes memory denom = bytes(data.denom);
        if (denom.length >= denomPrefix.length && denom.slice(0, denomPrefix.length).equal(denomPrefix)) {
            require(_mint(_decodeSender(data.sender), data.denom, data.amount));
        } else {
            // sender was source chain
            require(
                _transferFrom(_getEscrowAddress(sourceChannel), _decodeSender(data.sender), data.denom, data.amount)
            );
        }
    }

    function _getDenomPrefix(string calldata port, string calldata channel) internal pure returns (bytes memory) {
        return abi.encodePacked(port, "/", channel, "/");
    }

    /**
     * @dev _transferFrom transfers tokens from `sender` to `receiver` in the bank.
     */
    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount)
        internal
        virtual
        returns (bool);

    /**
     * @dev _mint mints tokens to `account` in the bank.
     */
    function _mint(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    /**
     * @dev _burn burns tokens from `account` in the bank.
     */
    function _burn(address account, string memory denom, uint256 amount) internal virtual returns (bool);

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
        require(ok, "invalid address");
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
