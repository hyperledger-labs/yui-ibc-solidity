// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ERC165Checker} from "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {Channel} from "../../proto/Channel.sol";
import {IIBCModule, IIBCModuleInitializer} from "./IIBCModule.sol";
import {IIBCModuleUpgrade} from "./IIBCModuleUpgrade.sol";
import {IIBCModuleManager} from "./IIBCModuleManager.sol";

/**
 * @dev IBCModuleManager is a contract that provides the functions defined in [ICS 5](https://github.com/cosmos/ibc/tree/main/spec/core/ics-005-port-allocation) and [ICS 26](https://github.com/cosmos/ibc/tree/main/spec/core/ics-026-routing-module).
 */
contract IBCModuleManager is Context, IBCHost, IIBCModuleManager {
    /**
     * @dev claimPortCapability allows the IBC app module to claim a capability that core IBC passes to it
     */
    function claimPortCapability(string calldata portId, address module) internal {
        HostStorage storage hostStorage = getHostStorage();
        if (hostStorage.portCapabilities[portId] != address(0)) {
            revert IBCHostPortCapabilityAlreadyClaimed(portId);
        }
        if (module == address(0) || module == address(this)) {
            revert IBCHostInvalidModuleAddress(module);
        }
        if (!ERC165Checker.supportsERC165(module)) {
            revert IBCHostModuleDoesNotSupportERC165();
        }
        if (
            !(
                ERC165Checker.supportsERC165InterfaceUnchecked(module, type(IIBCModule).interfaceId)
                    || ERC165Checker.supportsERC165InterfaceUnchecked(module, type(IIBCModuleInitializer).interfaceId)
            )
        ) {
            revert IBCHostModuleDoesNotSupportIIBCModuleInitializer(module, type(IIBCModuleInitializer).interfaceId);
        }
        hostStorage.portCapabilities[portId] = module;
        emit IBCModuleManagerPortCapabilityClaimed(portId, module);
    }

    /**
     * @dev claimChannelCapability allows the IBC app module to claim a capability that core IBC passes to it
     */
    function claimChannelCapability(string calldata portId, string memory channelId, address module) internal {
        HostStorage storage hostStorage = getHostStorage();
        if (hostStorage.channelCapabilities[portId][channelId] != address(0)) {
            revert IBCHostChannelCapabilityAlreadyClaimed(portId, channelId);
        }
        if (module == address(0)) {
            revert IBCHostInvalidModuleAddress(module);
        }
        if (!ERC165Checker.supportsERC165(module)) {
            revert IBCHostModuleDoesNotSupportERC165();
        }
        if (!ERC165Checker.supportsERC165InterfaceUnchecked(module, type(IIBCModule).interfaceId)) {
            revert IBCHostModuleDoesNotSupportIIBCModule(module, type(IIBCModule).interfaceId);
        }
        hostStorage.channelCapabilities[portId][channelId] = module;
        emit IBCModuleManagerChannelCapabilityClaimed(portId, channelId, module);
    }

    /**
     * @dev authenticateChannelCapability attempts to authenticate a given name from a caller.
     * It allows for a caller to check that a capability does in fact correspond to a particular name.
     */
    function authenticateChannelCapability(string calldata portId, string calldata channelId) internal view {
        address msgSender = _msgSender();
        if (getHostStorage().channelCapabilities[portId][channelId] != msgSender) {
            revert IBCHostFailedAuthenticateChannelCapability(portId, channelId, msgSender);
        }
    }

    /**
     * @dev lookupModuleByPort will return the IBCModule along with the capability associated with a given portID
     * If the module is not found, it will revert
     */
    function lookupModuleByPort(string calldata portId) internal view virtual returns (IIBCModuleInitializer) {
        address module = getHostStorage().portCapabilities[portId];
        if (module == address(0)) {
            revert IBCHostModulePortNotFound(portId);
        }
        return IIBCModuleInitializer(module);
    }

    /**
     * @dev lookupModuleByChannel will return the IBCModule along with the capability associated with a given channel defined by its portID and channelID
     * If the module is not found, it will revert
     */
    function lookupModuleByChannel(string calldata portId, string calldata channelId)
        internal
        view
        virtual
        returns (IIBCModule)
    {
        address module = getHostStorage().channelCapabilities[portId][channelId];
        if (module == address(0)) {
            revert IBCHostModuleChannelNotFound(portId, channelId);
        }
        return IIBCModule(module);
    }

    /**
     * @dev lookupUpgradableModuleByPortUnchecked will return the IBCModule corresponding to the portID
     * It will revert if the module is not found
     *
     * Since the function does not check if the module supports the `IIBCModuleUpgrade` interface via ERC-165, it is unsafe but cheaper in gas cost than `lookupUpgradableModuleByPort`
     */
    function lookupUpgradableModuleUnchecked(string calldata portId, string calldata channelId)
        internal
        view
        returns (IIBCModuleUpgrade)
    {
        return IIBCModuleUpgrade(address(lookupModuleByChannel(portId, channelId)));
    }

    /**
     * @dev lookupUpgradableModule will return the IBCModule corresponding to the portID and channelID
     * It will revert if the module does not support the `IIBCModuleUpgrade` interface or the module is not found
     */
    function lookupUpgradableModule(string calldata portId, string calldata channelId)
        internal
        view
        returns (IIBCModuleUpgrade)
    {
        IIBCModule module = lookupModuleByChannel(portId, channelId);
        if (!module.supportsInterface(type(IIBCModuleUpgrade).interfaceId)) {
            revert IBCHostModuleDoesNotSupportIIBCModuleUpgrade(address(module));
        }
        return IIBCModuleUpgrade(address(module));
    }

    /**
     * @dev canTransitionToFlushComplete checks if the module can transition to flush complete at the given upgrade sequence
     * If the module is not found, it will revert
     */
    function canTransitionToFlushComplete(
        Channel.Order ordering,
        string calldata portId,
        string calldata channelId,
        uint64 upgradeSequence
    ) internal view virtual returns (bool) {
        if (ordering == Channel.Order.ORDER_ORDERED) {
            ChannelStorage storage channelStorage = getChannelStorage()[portId][channelId];
            if (channelStorage.nextSequenceSend == channelStorage.nextSequenceAck) {
                return true;
            }
        }
        return lookupUpgradableModuleUnchecked(portId, channelId).canTransitionToFlushComplete(
            portId, channelId, upgradeSequence, _msgSender()
        );
    }
}
