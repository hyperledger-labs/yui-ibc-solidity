// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../../contracts/proto/Channel.sol";
import "../../../contracts/core/05-port/IIBCModule.sol";
import "../../../contracts/core/25-handler/IBCHandler.sol";
import "../../../contracts/core/24-host/IBCHost.sol";
import "../../../contracts/lib/Bytes.sol";
import "@openzeppelin/contracts/utils/Context.sol";

contract MockApp is IIBCModule {
    event MockRecv(bool ok);

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata, address) external virtual override returns (bytes memory) {
        emit MockRecv(true);
        return bytes("1");
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement, address relayer)
        external
        virtual
        override
    {}

    function onChanOpenInit(
        Channel.Order,
        string[] calldata,
        string calldata,
        string calldata channelId,
        ChannelCounterparty.Data calldata,
        string calldata
    ) external virtual override {}

    function onChanOpenTry(
        Channel.Order,
        string[] calldata,
        string calldata,
        string calldata channelId,
        ChannelCounterparty.Data calldata,
        string calldata,
        string calldata
    ) external virtual override {}

    function onChanOpenAck(string calldata portId, string calldata channelId, string calldata counterpartyVersion)
        external
        virtual
        override
    {}

    function onChanOpenConfirm(string calldata portId, string calldata channelId) external virtual override {}

    function onChanCloseInit(string calldata portId, string calldata channelId) external virtual override {}

    function onChanCloseConfirm(string calldata portId, string calldata channelId) external virtual override {}
}
