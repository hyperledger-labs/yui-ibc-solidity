// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Height} from "../../proto/Client.sol";
import {Packet} from "../../proto/Channel.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {IBCMockLib} from "./IBCMockLib.sol";

contract IBCMockApp is IBCAppBase, Ownable {
    string public constant MOCKAPP_VERSION = "mockapp-1";

    IIBCHandler immutable ibcHandler;

    bool public closeChannelAllowed = true;

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

    function writeAcknowledgement(string calldata destinationPort, string calldata destinationChannel, uint64 sequence)
        external
    {
        ibcHandler.writeAcknowledgement(
            destinationPort, destinationChannel, sequence, IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON
        );
    }

    function allowCloseChannel(bool allow) public onlyOwner {
        closeChannelAllowed = allow;
    }

    function onRecvPacket(Packet.Data calldata packet, address)
        external
        view
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        if (equals(packet.data, IBCMockLib.MOCK_PACKET_DATA)) {
            return IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
        } else if (equals(packet.data, IBCMockLib.MOCK_ASYNC_PACKET_DATA)) {
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
        if (equals(packet.data, IBCMockLib.MOCK_PACKET_DATA)) {
            require(equals(acknowledgement, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON), "invalid ack");
        } else if (equals(packet.data, IBCMockLib.MOCK_ASYNC_PACKET_DATA)) {
            require(equals(acknowledgement, IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON), "invalid async ack");
        } else if (equals(packet.data, IBCMockLib.MOCK_FAIL_PACKET_DATA)) {
            require(equals(acknowledgement, IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON), "invalid failed ack");
        } else {
            revert("invalid packet data");
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

    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata) external virtual override onlyIBC {
        require(closeChannelAllowed, "close not allowed");
    }

    function onChanCloseConfirm(IIBCModule.MsgOnChanCloseConfirm calldata) external virtual override onlyIBC {
        require(closeChannelAllowed, "close not allowed");
    }

    function onTimeoutPacket(Packet.Data calldata, address) external view override onlyIBC {
        require(closeChannelAllowed, "timeout not allowed");
    }

    function equals(bytes calldata a, bytes memory b) private pure returns (bool) {
        return keccak256(a) == keccak256(b);
    }
}
