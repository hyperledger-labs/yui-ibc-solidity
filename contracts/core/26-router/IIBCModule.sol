// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Channel, ChannelCounterparty} from "../../proto/Channel.sol";
import {Packet} from "../04-channel/IIBCChannel.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @notice IIBCModuleInitializer is an interface that defines the callbacks that modules must define as specified in ICS-26
 * @dev IBCModules registered with IBCModuleManager via `bindPort` must implement this interface.
 * IERC165's `supportsInterface` must return true for the interface ID of this interface
 */
interface IIBCModuleInitializer is IERC165 {
    struct MsgOnChanOpenInit {
        Channel.Order order;
        string[] connectionHops;
        string portId;
        string channelId;
        ChannelCounterparty.Data counterparty;
        string version;
    }

    struct MsgOnChanOpenTry {
        Channel.Order order;
        string[] connectionHops;
        string portId;
        string channelId;
        ChannelCounterparty.Data counterparty;
        string counterpartyVersion;
    }

    /**
     * @dev onChanOpenInit will verify that the relayer-chosen parameters
     * are valid and perform any custom INIT logic.
     * It may return an error if the chosen parameters are invalid
     * in which case the handshake is aborted.
     * If the provided version string is non-empty, OnChanOpenInit should return
     * the version string if valid or an error if the provided version is invalid.
     * If the version string is empty, OnChanOpenInit is expected to
     * return a default version string representing the version(s) it supports.
     * If there is no default version string for the application,
     * it should return an error if provided version is empty string.
     * onChanOpenInit can also create a new contract instance corresponding to the channel and return the address.
     * Otherwise, it should return the address of this contract.
     */
    function onChanOpenInit(MsgOnChanOpenInit calldata msg_) external returns (address, string memory);

    /**
     * @dev onChanOpenTry will verify the relayer-chosen parameters along with the
     * counterparty-chosen version string and perform custom TRY logic.
     * If the relayer-chosen parameters are invalid, the callback must return
     * an error to abort the handshake. If the counterparty-chosen version is not
     * compatible with this modules supported versions, the callback must return
     * an error to abort the handshake. If the versions are compatible, the try callback
     * must select the final version string and return it to core IBC.
     * onChanOpenTry may also perform custom initialization logic.
     * onChanOpenTry can also create a new contract instance corresponding to the channel and return the address.
     * Otherwise, it should return the address of this contract.
     */
    function onChanOpenTry(MsgOnChanOpenTry calldata msg_) external returns (address, string memory);
}

/**
 * @notice IIBCModule is an interface that defines all the callbacks that modules must define as specified in ICS-26
 * @dev IBC Modules returned by `onChanOpenInit` and `onChanOpenTry` must implement this interface.
 * IERC165's `supportsInterface` must return true for the interface ID of this interface, and the interface ID of IIBCModuleUpgrade if the module supports the channel upgrade
 */
interface IIBCModule is IIBCModuleInitializer {
    struct MsgOnChanOpenAck {
        string portId;
        string channelId;
        string counterpartyVersion;
    }

    struct MsgOnChanOpenConfirm {
        string portId;
        string channelId;
    }

    struct MsgOnChanCloseInit {
        string portId;
        string channelId;
    }

    struct MsgOnChanCloseConfirm {
        string portId;
        string channelId;
    }

    /**
     * @dev OnChanOpenAck will error if the counterparty selected version string
     * is invalid to abort the handshake. It may also perform custom ACK logic.
     */
    function onChanOpenAck(MsgOnChanOpenAck calldata msd_) external;

    /**
     * @dev OnChanOpenConfirm will perform custom CONFIRM logic and may error to abort the handshake.
     */
    function onChanOpenConfirm(MsgOnChanOpenConfirm calldata msg_) external;

    /**
     * @dev OnChanCloseInit will perform custom CLOSE_INIT logic and may error to abort the handshake.
     * NOTE: If the application does not allow the channel to be closed, this function must revert.
     */
    function onChanCloseInit(MsgOnChanCloseInit calldata msg_) external;

    /**
     * @dev OnChanCloseConfirm will perform custom CLOSE_CONFIRM logic and may error to abort the handshake.
     */
    function onChanCloseConfirm(MsgOnChanCloseConfirm calldata msg_) external;

    /**
     * @dev OnRecvPacket must return an acknowledgement that implements the Acknowledgement interface.
     * In the case of an asynchronous acknowledgement, nil should be returned.
     * If the acknowledgement returned is successful, the state changes on callback are written,
     * otherwise the application state changes are discarded. In either case the packet is received
     * and the acknowledgement is written (in synchronous cases).
     */
    function onRecvPacket(Packet calldata, address relayer) external returns (bytes memory);

    /**
     * @dev onAcknowledgementPacket is called when a packet sent by this module has been acknowledged.
     */
    function onAcknowledgementPacket(Packet calldata, bytes calldata acknowledgement, address relayer) external;

    /**
     * @dev onTimeoutPacket is called when a packet sent by this module has timed-out (such that it will not be received on the destination chain).
     */
    function onTimeoutPacket(Packet calldata, address relayer) external;
}
