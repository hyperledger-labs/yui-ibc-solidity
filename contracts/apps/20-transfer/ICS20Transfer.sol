// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../commons/IBCAppBase.sol";
import "../../core/05-port/IIBCModule.sol";
import "../../core/25-handler/IBCHandler.sol";
import "../../proto/Channel.sol";
import "./ICS20Lib.sol";
import "solidity-stringutils/src/strings.sol";
import "solidity-bytes-utils/contracts/BytesLib.sol";

abstract contract ICS20Transfer is IBCAppBase {
    using strings for *;
    using BytesLib for bytes;

    mapping(string => address) channelEscrowAddresses;

    function onRecvPacket(Packet.Data calldata packet, address)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        ICS20Lib.PacketData memory data = ICS20Lib.unmarshalJSON(packet.data);
        strings.slice memory denom = data.denom.toSlice();
        strings.slice memory trimedDenom =
            data.denom.toSlice().beyond(_makeDenomPrefix(packet.source_port, packet.source_channel));
        bool success;
        address receiver;
        (receiver, success) = _decodeReceiver(data.receiver);
        if (!success) {
            return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
        }
        if (!denom.equals(trimedDenom)) {
            // receiver is source chain
            success = _transferFrom(
                _getEscrowAddress(packet.destination_channel), receiver, trimedDenom.toString(), data.amount
            );
        } else {
            string memory prefixedDenom =
                _makeDenomPrefix(packet.destination_port, packet.destination_channel).concat(denom);
            success = _mint(receiver, prefixedDenom, data.amount);
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
        if (keccak256(acknowledgement) != keccak256(ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON)) {
            _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
        }
    }

    function onChanOpenInit(
        Channel.Order,
        string[] calldata,
        string calldata,
        string calldata channelId,
        ChannelCounterparty.Data calldata,
        string calldata
    ) external virtual override onlyIBC {
        channelEscrowAddresses[channelId] = address(this);
    }

    function onChanOpenTry(
        Channel.Order,
        string[] calldata,
        string calldata,
        string calldata channelId,
        ChannelCounterparty.Data calldata,
        string calldata,
        string calldata
    ) external virtual override onlyIBC {
        channelEscrowAddresses[channelId] = address(this);
    }

    function onTimeoutPacket(Packet.Data calldata packet, address) external virtual override onlyIBC {
        _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
    }

    function _getEscrowAddress(string memory sourceChannel) internal view virtual returns (address) {
        address escrow = channelEscrowAddresses[sourceChannel];
        require(escrow != address(0));
        return escrow;
    }

    function _refundTokens(ICS20Lib.PacketData memory data, string memory sourcePort, string memory sourceChannel)
        internal
        virtual
    {
        if (!data.denom.toSlice().startsWith(_makeDenomPrefix(sourcePort, sourceChannel))) {
            // sender was source chain
            require(
                _transferFrom(_getEscrowAddress(sourceChannel), _decodeSender(data.sender), data.denom, data.amount)
            );
        } else {
            require(_mint(_decodeSender(data.sender), data.denom, data.amount));
        }
    }

    function _makeDenomPrefix(string memory port, string memory channel)
        internal
        pure
        virtual
        returns (strings.slice memory)
    {
        return string(abi.encodePacked(port, "/", channel, "/")).toSlice();
    }

    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount)
        internal
        virtual
        returns (bool);

    function _mint(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    function _burn(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    function _encodeSender(address sender) internal pure virtual returns (string memory);

    function _decodeSender(string memory sender) internal pure virtual returns (address);

    // @dev `receiver` may be a invalid address.
    function _decodeReceiver(string memory receiver) internal pure virtual returns (address, bool);
}
