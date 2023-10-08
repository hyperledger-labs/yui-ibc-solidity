// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Context.sol";
import "../../core/05-port/IIBCModule.sol";
import "../../core/04-channel/IIBCChannel.sol";
import "../commons/IBCAppBase.sol";

/**
 * IBC host callbacacaks are delegated to underlying IBCAppBase module.
 * IBCAppBase must be initialized with address of IIBCMiddleware as `ibcAddress`.
 * IBCAppBase calls
 */
abstract contract IIBCMiddleware is Context, IBCAppBase, IICS04Wrapper {
    // middleware has acccess to an underlying application which may be wrapped
    // by more middleware.
    IBCAppBase private ibcModule;

    // middleware has access to ICS4Wrapper which may be core IBC Channel Handler
    // or a higher-level middleware that wraps this middleware.
    IICS04Wrapper private ics04Wrapper;

    /// either raw IBC handler or other middleware
    address private ibcHandler;

    constructor(IBCAppBase ibcModule_, IICS04Wrapper ics04Wrapper_, address ibcHandler_) {
        ibcModule = ibcModule_;
        ics04Wrapper = ics04Wrapper_;
        ibcHandler = ibcHandler_;
    }

    function ibcAddress() public view virtual override returns (address) {
        return ibcHandler;
    }

    modifier onlyWrapped() {
        require(
            _msgSender() == address(ibcModule) || _msgSender() == address(this),
            "IIBCMiddleware: caller is not the IBCAppBase or IICS04Wrapper"
        );
        _;
    }

    function onChanOpenInit(
        Channel.Order order,
        string[] calldata connectionHops,
        string calldata portId,
        string calldata channelId,
        ChannelCounterparty.Data calldata counterparty,
        string calldata version
    ) external virtual override onlyIBC {
        ibcModule.onChanOpenInit(order, connectionHops, portId, channelId, counterparty, version);
    }

    function onChanOpenTry(
        Channel.Order order,
        string[] calldata connectionHops,
        string calldata portId,
        string calldata channelId,
        ChannelCounterparty.Data calldata counterparty,
        string calldata version,
        string calldata counterpartyVersion
    ) external virtual override onlyIBC {
        ibcModule.onChanOpenTry(order, connectionHops, portId, channelId, counterparty, version, counterpartyVersion);
    }

    function onChanOpenConfirm(string calldata portId, string calldata channelId) external virtual override onlyIBC {
        ibcModule.onChanOpenConfirm(portId, channelId);
    }

    function onChanCloseInit(string calldata portId, string calldata channelId) external virtual override onlyIBC {
        ibcModule.onChanCloseInit(portId, channelId);
    }

    function onChanCloseConfirm(string calldata portId, string calldata channelId) external virtual override onlyIBC {
        ibcModule.onChanCloseConfirm(portId, channelId);
    }

    function onRecvPacket(Packet.Data calldata packet, address relayer)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        return ibcModule.onRecvPacket(packet, relayer);
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement, address relayer)
        external
        virtual
        override
        onlyIBC
    {
        return ibcModule.onAcknowledgementPacket(packet, acknowledgement, relayer);
    }

    function onTimeoutPacket(Packet.Data calldata packet, address relayer) external virtual override onlyIBC {
        return ibcModule.onTimeoutPacket(packet, relayer);
    }

    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external virtual override onlyWrapped {
        ics04Wrapper.writeAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function sendPacket(
        string calldata sourcePort,
        string calldata sourceChannel,
        Height.Data calldata timeoutHeight,
        uint64 timeoutTimestamp,
        bytes calldata data
    ) external virtual override onlyWrapped returns (uint64) {
        return ics04Wrapper.sendPacket(sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data);
    }
}
