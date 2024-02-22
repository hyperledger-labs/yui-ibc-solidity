// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {Packet} from "../../proto/Channel.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {IIBCModuleErrors} from "../../core/26-router/IIBCModuleErrors.sol";

/**
 * @dev Base contract of the IBC App protocol
 */
abstract contract IBCAppBase is Context, IIBCModule, IIBCModuleErrors {
    /**
     * @dev Throws if called by any account other than the IBC contract.
     */
    modifier onlyIBC() {
        _checkIBC();
        _;
    }

    /**
     * @dev Returns the address of the IBC contract.
     */
    function ibcAddress() public view virtual returns (address);

    /**
     * @dev Throws if the sender is not the IBC contract.
     */
    function _checkIBC() internal view virtual {
        if (ibcAddress() != _msgSender()) {
            revert IBCModuleInvalidSender(_msgSender());
        }
    }

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
    function onRecvPacket(Packet.Data calldata, address)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
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
