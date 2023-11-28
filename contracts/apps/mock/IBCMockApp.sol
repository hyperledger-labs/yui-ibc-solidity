// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import {Height} from "../../proto/Client.sol";
import {Packet} from "../../proto/Channel.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {IBCMockLib} from "./IBCMockLib.sol";

contract IBCMockApp is IBCAppBase {
    string public constant MOCKAPP_VERSION = "mockapp-1";

    IIBCHandler immutable ibcHandler;

    constructor(IIBCHandler ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }

    function sendPacket(
        bytes memory message,
        string calldata sourcePort,
        string calldata sourceChannel,
        Height.Data calldata timeoutHeight,
        uint64 timeoutTimestamp
    ) external returns (uint64) {
        return ibcHandler.sendPacket(sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, message);
    }

    function onRecvPacket(Packet.Data calldata packet, address)
        external
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        if (keccak256(packet.data) == keccak256(IBCMockLib.MOCK_PACKET_DATA)) {
            return IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
        } else if (keccak256(packet.data) == keccak256(IBCMockLib.MOCK_ASYNC_PACKET_DATA)) {
            return bytes("");
        } else {
            return IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON;
        }
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement, address)
        external
        virtual
        override
        onlyIBC
    {
        if (keccak256(packet.data) == keccak256(IBCMockLib.MOCK_PACKET_DATA)) {
            require(keccak256(acknowledgement) == keccak256(IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON));
        } else if (keccak256(packet.data) == keccak256(IBCMockLib.MOCK_ASYNC_PACKET_DATA)) {
            require(acknowledgement.length == 0);
        } else {
            require(keccak256(acknowledgement) == keccak256(IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON));
        }
    }

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        require(
            bytes(msg_.version).length == 0 || keccak256(bytes(msg_.version)) == keccak256(bytes(MOCKAPP_VERSION)),
            "version mismatch"
        );
        return MOCKAPP_VERSION;
    }

    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        require(keccak256(bytes(msg_.counterpartyVersion)) == keccak256(bytes(MOCKAPP_VERSION)), "version mismatch");
        return MOCKAPP_VERSION;
    }
}
