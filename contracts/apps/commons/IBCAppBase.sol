// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Context.sol";
import "../../core/05-port/IIBCModule.sol";

import "../../core/04-channel/IIBCChannel.sol";

/**
 * @dev Base contract of the IBC App protocol
 */
abstract contract IBCAppBaseImpl is Context, IIBCModule, IICS04Wrapper {
    IICS04Wrapper private immutable ics04Wrapper;

    /**
     * @dev Throws if called by any account other than the IBC contract.
     */
    modifier onlyIBC() {
        _checkIBC();
        _;
    }

    constructor(IICS04Wrapper ics04Wrapper_) {
        ics04Wrapper = ics04Wrapper_;
    }

    modifier onlyWrapped() {
        require(
            ibcAddress() == _msgSender() || address(this) == _msgSender(),
            "IBCAppBase: caller is not the IBC contract nor this contract(middeware)"
        );
        _;
    }

    /**
     * @dev Returns the address of the IBC contract.
     */
    function ibcAddress() public view virtual returns (address) {
        return address(ics04Wrapper);
    }

    /**
     * @dev Throws if the sender is not the IBC contract.
     */
    function _checkIBC() internal view virtual {
        require(ibcAddress() == _msgSender(), "IBCAppBase: caller is not the IBC contract");
    }

    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external virtual onlyWrapped {
        ics04Wrapper.writeAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function sendPacket(
        string calldata sourcePort,
        string calldata sourceChannel,
        Height.Data calldata timeoutHeight,
        uint64 timeoutTimestamp,
        bytes calldata data
    ) external virtual onlyWrapped returns (uint64) {
        return ics04Wrapper.sendPacket(sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data);
    }
}

abstract contract IBCAppBaseOnRecvPacket is IBCAppBaseImpl {
    function onRecvPacket(Packet.Data memory, address)
        public
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement, bool success)
    {}
}

abstract contract IBCAppBase is IBCAppBaseImpl {
    constructor(IICS04Wrapper ics04Wrapper_) IBCAppBaseImpl(ics04Wrapper_) {}

    /**
     * @dev See IIBCModule-onChanOpenInit
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {}

    /**
     * @dev See IIBCModule-onChanOpenTry
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata)
        external
        virtual
        override
        onlyIBC
        returns (string memory)
    {}

    /**
     * @dev See IIBCModule-onChanOpenAck
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenAck(IIBCModule.MsgOnChanOpenAck calldata) external virtual override onlyIBC {}

    /**
     * @dev See IIBCModule-onChanOpenConfirm
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenConfirm(IIBCModule.MsgOnChanOpenConfirm calldata) external virtual override onlyIBC {}

    /**
     * @dev See IIBCModule-onChanCloseInit
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata) external virtual override onlyIBC {}

    /**
     * @dev See IIBCModule-onChanCloseConfirm
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanCloseConfirm(IIBCModule.MsgOnChanCloseConfirm calldata) external virtual override onlyIBC {}

    /**
     * @dev See IIBCModule-onRecvPacket
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onRecvPacket(Packet.Data memory, address)
        public
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement, bool success)
    {}

    /**
     * @dev See IIBCModule-onAcknowledgementPacket
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata, address) external virtual override onlyIBC {}

    /**
     * @dev See IIBCModule-onTimeoutPacket
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onTimeoutPacket(Packet.Data calldata, address relayer) external virtual onlyIBC {}
}
