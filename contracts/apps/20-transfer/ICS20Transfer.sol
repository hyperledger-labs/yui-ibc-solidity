// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../commons/IBCAppBase.sol";
import "../../core/05-port/IIBCModule.sol";
import "../../core/25-handler/IBCHandler.sol";
import "../../proto/Channel.sol";
import "./ICS20Packet.sol";
import "solidity-stringutils/src/strings.sol";
import "solidity-bytes-utils/contracts/BytesLib.sol";

abstract contract ICS20Transfer is IBCAppBase {
    using strings for *;
    using BytesLib for bytes;

    mapping(string => address) channelEscrowAddresses;

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata packet, address)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        ICS20Packet.PacketData memory data = ICS20Packet.unmarshalJSON(packet.data);
        strings.slice memory denom = data.denom.toSlice();
        strings.slice memory trimedDenom =
            data.denom.toSlice().beyond(_makeDenomPrefix(packet.source_port, packet.source_channel));
        if (!denom.equals(trimedDenom)) {
            // receiver is source chain
            return _newAcknowledgement(
                _transferFrom(
                    _getEscrowAddress(packet.destination_channel),
                    bytes(data.receiver).toAddress(0),
                    trimedDenom.toString(),
                    data.amount
                )
            );
        } else {
            string memory prefixedDenom =
                _makeDenomPrefix(packet.destination_port, packet.destination_channel).concat(denom);
            return _newAcknowledgement(_mint(bytes(data.receiver).toAddress(0), prefixedDenom, data.amount));
        }
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement, address)
        external
        virtual
        override
        onlyIBC
    {
        if (!_isSuccessAcknowledgement(acknowledgement)) {
            _refundTokens(ICS20Packet.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
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
        _refundTokens(ICS20Packet.unmarshalJSON(packet.data), packet.source_port, packet.source_channel);
    }

    /// Internal functions ///

    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount)
        internal
        virtual
        returns (bool);

    function _mint(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    function _burn(address account, string memory denom, uint256 amount) internal virtual returns (bool);

    function _sendPacket(
        ICS20Packet.PacketData memory data,
        string memory sourcePort,
        string memory sourceChannel,
        uint64 timeoutHeight
    ) internal virtual {
        IBCHandler(ibcAddress()).sendPacket(
            sourcePort,
            sourceChannel,
            Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            0,
            ICS20Packet.marshalUnsafeJSON(data)
        );
    }

    function _getEscrowAddress(string memory sourceChannel) internal view virtual returns (address) {
        address escrow = channelEscrowAddresses[sourceChannel];
        require(escrow != address(0));
        return escrow;
    }

    function _newAcknowledgement(bool success) internal pure virtual returns (bytes memory) {
        if (success) {
            return ICS20Packet.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
        } else {
            return ICS20Packet.FAILED_ACKNOWLEDGEMENT_JSON;
        }
    }

    function _isSuccessAcknowledgement(bytes memory acknowledgement) internal pure virtual returns (bool) {
        return keccak256(acknowledgement) == keccak256(ICS20Packet.SUCCESSFUL_ACKNOWLEDGEMENT_JSON);
    }

    function _refundTokens(ICS20Packet.PacketData memory data, string memory sourcePort, string memory sourceChannel)
        internal
        virtual
    {
        if (!data.denom.toSlice().startsWith(_makeDenomPrefix(sourcePort, sourceChannel))) {
            // sender was source chain
            require(
                _transferFrom(
                    _getEscrowAddress(sourceChannel), bytes(data.sender).toAddress(0), data.denom, data.amount
                )
            );
        } else {
            require(_mint(bytes(data.sender).toAddress(0), data.denom, data.amount));
        }
    }

    /// Helper functions ///

    function _makeDenomPrefix(string memory port, string memory channel)
        internal
        pure
        virtual
        returns (strings.slice memory)
    {
        return port.toSlice().concat("/".toSlice()).toSlice().concat(channel.toSlice()).toSlice().concat("/".toSlice())
            .toSlice();
    }
}
