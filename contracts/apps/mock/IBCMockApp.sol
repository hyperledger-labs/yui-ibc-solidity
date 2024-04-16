// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Height} from "../../proto/Client.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {Packet} from "../../core/04-channel/IIBCChannel.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {IBCMockLib} from "./IBCMockLib.sol";
import {IIBCMockErrors} from "./IIBCMockErrors.sol";

contract IBCMockApp is IBCAppBase, IIBCMockErrors, Ownable {
    string public constant MOCKAPP_VERSION = "mockapp-1";

    IIBCHandler immutable ibcHandler;

    bool public closeChannelAllowed = true;

    constructor(IIBCHandler ibcHandler_) Ownable(msg.sender) {
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

    function onRecvPacket(Packet calldata packet, address)
        external
        view
        virtual
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

    function onAcknowledgementPacket(Packet calldata packet, bytes calldata acknowledgement, address)
        external
        virtual
        override
        onlyIBC
    {
        if (equals(packet.data, IBCMockLib.MOCK_PACKET_DATA)) {
            if (!equals(acknowledgement, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON)) {
                revert IBCMockUnexpectedAcknowledgement(acknowledgement, IBCMockLib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON);
            }
        } else if (equals(packet.data, IBCMockLib.MOCK_ASYNC_PACKET_DATA)) {
            if (!equals(acknowledgement, IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON)) {
                revert IBCMockUnexpectedAcknowledgement(
                    acknowledgement, IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON
                );
            }
        } else if (equals(packet.data, IBCMockLib.MOCK_FAIL_PACKET_DATA)) {
            if (!equals(acknowledgement, IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON)) {
                revert IBCMockUnexpectedAcknowledgement(acknowledgement, IBCMockLib.FAILED_ACKNOWLEDGEMENT_JSON);
            }
        } else {
            revert IBCMockUnexpectedPacket(packet.data);
        }
    }

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        if (bytes(msg_.version).length != 0 && keccak256(bytes(msg_.version)) != keccak256(bytes(MOCKAPP_VERSION))) {
            revert IBCMockUnexpectedVersion(msg_.version, MOCKAPP_VERSION);
        }
        return MOCKAPP_VERSION;
    }

    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(MOCKAPP_VERSION))) {
            revert IBCMockUnexpectedVersion(msg_.counterpartyVersion, MOCKAPP_VERSION);
        }
        return MOCKAPP_VERSION;
    }

    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata msg_) external virtual override onlyIBC {
        if (!closeChannelAllowed) {
            revert IBCModuleChannelCloseNotAllowed(msg_.portId, msg_.channelId);
        }
    }

    function onChanCloseConfirm(IIBCModule.MsgOnChanCloseConfirm calldata msg_) external virtual override onlyIBC {
        if (!closeChannelAllowed) {
            revert IBCModuleChannelCloseNotAllowed(msg_.portId, msg_.channelId);
        }
    }

    function onTimeoutPacket(Packet calldata packet, address) external view virtual override onlyIBC {
        if (!closeChannelAllowed) {
            revert IBCModuleChannelCloseNotAllowed(packet.sourcePort, packet.sourceChannel);
        }
    }

    function equals(bytes calldata a, bytes memory b) internal pure returns (bool) {
        return keccak256(a) == keccak256(b);
    }
}
