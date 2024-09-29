// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ShortString, ShortStrings} from "@openzeppelin/contracts/utils/ShortStrings.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Height} from "../../proto/Client.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {Packet} from "../../core/04-channel/IIBCChannel.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {IBCMockLib} from "./IBCMockLib.sol";
import {IIBCMockErrors} from "./IIBCMockErrors.sol";
import {IBCAppInitializerBase} from "../commons/IBCAppBase.sol";

contract IBCMockAppFactory is IIBCMockErrors, IBCAppInitializerBase {
    string public constant MOCKAPP_VERSION = "mockapp-1";

    IIBCHandler immutable ibcHandler;

    mapping(string channelId => IBCMockApp2) internal apps;

    constructor(IIBCHandler ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        if (bytes(msg_.version).length != 0 && keccak256(bytes(msg_.version)) != keccak256(bytes(MOCKAPP_VERSION))) {
            revert IBCMockUnexpectedVersion(msg_.version, MOCKAPP_VERSION);
        }
        apps[msg_.channelId] = new IBCMockApp2(ibcHandler, msg_.portId, msg_.channelId);
        return (address(apps[msg_.channelId]), MOCKAPP_VERSION);
    }

    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata msg_)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(MOCKAPP_VERSION))) {
            revert IBCMockUnexpectedVersion(msg_.counterpartyVersion, MOCKAPP_VERSION);
        }
        apps[msg_.channelId] = new IBCMockApp2(ibcHandler, msg_.portId, msg_.channelId);
        return (address(apps[msg_.channelId]), MOCKAPP_VERSION);
    }

    function lookupApp(string calldata channelId) external view returns (IBCMockApp2) {
        return apps[channelId];
    }
}

contract IBCMockApp2 is IBCAppBase, IIBCMockErrors, Ownable {
    using ShortStrings for string;
    using ShortStrings for ShortString;

    string public constant MOCKAPP_VERSION = "mockapp-1";

    IIBCHandler immutable ibcHandler;
    ShortString immutable port;
    ShortString immutable channel;

    bool public closeChannelAllowed = true;

    constructor(IIBCHandler ibcHandler_, string memory port_, string memory channel_) Ownable(msg.sender) {
        ibcHandler = ibcHandler_;
        port = port_.toShortString();
        channel = channel_.toShortString();
    }

    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }

    function sendPacket(bytes memory message, Height.Data calldata timeoutHeight, uint64 timeoutTimestamp)
        external
        returns (uint64)
    {
        return ibcHandler.sendPacket(port.toString(), channel.toString(), timeoutHeight, timeoutTimestamp, message);
    }

    function writeAcknowledgement(uint64 sequence) external {
        ibcHandler.writeAcknowledgement(
            port.toString(), channel.toString(), sequence, IBCMockLib.SUCCESSFUL_ASYNC_ACKNOWLEDGEMENT_JSON
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

    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        revert("not supported");
    }

    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        revert("not supported");
    }

    function equals(bytes calldata a, bytes memory b) internal pure returns (bool) {
        return keccak256(a) == keccak256(b);
    }
}
