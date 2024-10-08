// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {Packet} from "../../core/04-channel/IIBCChannel.sol";
import {IIBCModule, IIBCModuleInitializer} from "../../core/26-router/IIBCModule.sol";
import {IIBCModuleErrors} from "../../core/26-router/IIBCModuleErrors.sol";

abstract contract AppBase is Context, IERC165, IIBCModuleErrors {
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
     * @dev Returns true if this contract implements the interface defined by
     * `interfaceId`. See the corresponding
     * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[ERC section]
     * to learn more about how these ids are created.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual returns (bool) {
        return interfaceId == type(IERC165).interfaceId || interfaceId == this.ibcAddress.selector;
    }
}

abstract contract IBCAppInitializerBase is AppBase, IIBCModuleInitializer {
    /**
     * @dev See {IERC165-supportsInterface}
     *
     * NOTE: This must return true if the `interfaceId` is equal to the `IIBCModule` interface.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(IERC165, AppBase) returns (bool) {
        return interfaceId == type(IIBCModuleInitializer).interfaceId || super.supportsInterface(interfaceId);
    }
}

/**
 * @dev Base contract of the IBC App protocol
 */
abstract contract IBCAppBase is AppBase, IIBCModule {
    /**
     * @dev See {IIBCModule-onChanOpenInit}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        return (address(this), "");
    }

    /**
     * @dev See {IIBCModule-onChanOpenTry}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata)
        external
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        return (address(this), "");
    }

    /**
     * @dev See {IIBCModule-onChanOpenAck}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenAck(IIBCModule.MsgOnChanOpenAck calldata) external virtual override onlyIBC {}

    /**
     * @dev See {IIBCModule-onChanOpenConfirm}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanOpenConfirm(IIBCModule.MsgOnChanOpenConfirm calldata) external virtual override onlyIBC {}

    /**
     * @dev See {IIBCModule-onChanCloseInit}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata) external virtual override onlyIBC {}

    /**
     * @dev See {IIBCModule-onChanCloseConfirm}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onChanCloseConfirm(IIBCModule.MsgOnChanCloseConfirm calldata) external virtual override onlyIBC {}

    /**
     * @dev See {IIBCModule-onRecvPacket}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onRecvPacket(Packet calldata, address)
        external
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {}

    /**
     * @dev See {IIBCModule-onAcknowledgementPacket}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onAcknowledgementPacket(Packet calldata, bytes calldata, address) external virtual override onlyIBC {}

    /**
     * @dev See {IIBCModule-onTimeoutPacket}
     *
     * NOTE: You should apply an `onlyIBC` modifier to the function if a derived contract overrides it.
     */
    function onTimeoutPacket(Packet calldata, address relayer) external virtual override onlyIBC {}

    /**
     * @dev See {IERC165-supportsInterface}
     *
     * NOTE: This must return true if the `interfaceId` is equal to the `IIBCModule` interface.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(IERC165, AppBase) returns (bool) {
        return interfaceId == type(IIBCModule).interfaceId || super.supportsInterface(interfaceId);
    }
}
