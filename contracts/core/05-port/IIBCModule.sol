// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../proto/Channel.sol";

// IIBCModule defines an interface that implements all the callbacks
// that modules must define as specified in ICS-26
interface IIBCModule {
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
     */
    function onChanOpenInit(MsgOnChanOpenInit calldata msg_) external returns (string memory);

    /**
     * @dev onChanOpenTry will verify the relayer-chosen parameters along with the
     * counterparty-chosen version string and perform custom TRY logic.
     * If the relayer-chosen parameters are invalid, the callback must return
     * an error to abort the handshake. If the counterparty-chosen version is not
     * compatible with this modules supported versions, the callback must return
     * an error to abort the handshake. If the versions are compatible, the try callback
     * must select the final version string and return it to core IBC.
     * OnChanOpenTry may also perform custom initialization logic
     */
    function onChanOpenTry(MsgOnChanOpenTry calldata msg_) external returns (string memory);

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
    function onRecvPacket(Packet.Data calldata, address relayer) external returns (bytes memory);

    /**
     * @dev onAcknowledgementPacket is called when a packet sent by this module has been acknowledged.
     */
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement, address relayer) external;

    /**
     * @dev onTimeoutPacket is called when a packet sent by this module has timed-out (such that it will not be received on the destination chain).
     */
    function onTimeoutPacket(Packet.Data calldata, address relayer) external;
}
