// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {Channel} from "../../proto/Channel.sol";
import {IIBCModule} from "./IIBCModule.sol";
import {IIBCModuleUpgrade} from "./IIBCModuleUpgrade.sol";

/**
 * @dev IBCModuleManager is a contract that provides the functions defined in [ICS 5](https://github.com/cosmos/ibc/tree/main/spec/core/ics-005-port-allocation) and [ICS 26](https://github.com/cosmos/ibc/tree/main/spec/core/ics-026-routing-module).
 */
contract IBCModuleManager is IBCHost, Context {
    /**
     * @dev claimPortCapability allows the IBC app module to claim a capability that core IBC passes to it
     */
    function claimPortCapability(string calldata portId, address addr) internal {
        if (portCapabilities[portId] != address(0)) {
            revert IBCHostPortCapabilityAlreadyClaimed(portId);
        }
        portCapabilities[portId] = addr;
    }

    /**
     * @dev claimChannelCapability allows the IBC app module to claim a capability that core IBC passes to it
     */
    function claimChannelCapability(string calldata portId, string memory channelId, address addr) internal {
        if (channelCapabilities[portId][channelId] != address(0)) {
            revert IBCHostChannelCapabilityAlreadyClaimed(portId, channelId);
        }
        channelCapabilities[portId][channelId] = addr;
    }

    /**
     * @dev authenticateChannelCapability attempts to authenticate a given name from a caller.
     * It allows for a caller to check that a capability does in fact correspond to a particular name.
     */
    function authenticateChannelCapability(string calldata portId, string calldata channelId) internal view {
        address msgSender = _msgSender();
        if (channelCapabilities[portId][channelId] != msgSender) {
            revert IBCHostFailedAuthenticateChannelCapability(portId, channelId, msgSender);
        }
    }

    /**
     * @dev lookupModuleByPort will return the IBCModule along with the capability associated with a given portID
     * If the module is not found, it will revert
     */
    function lookupModuleByPort(string calldata portId) internal view virtual returns (IIBCModule) {
        address module = portCapabilities[portId];
        if (module == address(0)) {
            revert IBCHostModulePortNotFound(portId);
        }
        return IIBCModule(module);
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
        address module = channelCapabilities[portId][channelId];
        if (module == address(0)) {
            revert IBCHostModuleChannelNotFound(portId, channelId);
        }
        return IIBCModule(module);
    }

    /**
     * @dev unsafeLookupUpgradableModuleByPort will return the IBCModule corresponding to the portID
     * It will revert if the module is not found
     *
     * Since the function does not check if the module supports the `IIBCModuleUpgrade` interface via ERC-165, it is unsafe but cheaper in gas cost than `lookupUpgradableModuleByPort`
     */
    function unsafeLookupUpgradableModuleByPort(string calldata portId) internal view returns (IIBCModuleUpgrade) {
        return IIBCModuleUpgrade(address(lookupModuleByPort(portId)));
    }

    /**
     * @dev lookupUpgradableModuleByPort will return the IBCModule corresponding to the portID
     * It will revert if the module does not support the `IIBCModuleUpgrade` interface or the module is not found
     */
    function lookupUpgradableModuleByPort(string calldata portId) internal view returns (IIBCModuleUpgrade) {
        IIBCModule module = lookupModuleByPort(portId);
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
            if (nextSequenceSends[portId][channelId] == nextSequenceAcks[portId][channelId]) {
                return true;
            }
        }
        return unsafeLookupUpgradableModuleByPort(portId).canTransitionToFlushComplete(
            portId, channelId, upgradeSequence, _msgSender()
        );
    }

    /**
     * @dev validatePortIdentifier validates a port identifier string
     *     check if the string consist of characters in one of the following categories only:
     *     - Alphanumeric
     *     - `.`, `_`, `+`, `-`, `#`
     *     - `[`, `]`, `<`, `>`
     */
    function validatePortIdentifier(bytes memory portId) internal pure returns (bool) {
        if (portId.length < 2 || portId.length > 128) {
            return false;
        }
        unchecked {
            for (uint256 i = 0; i < portId.length; i++) {
                uint256 c = uint256(uint8(portId[i]));
                if (
                    // a-z
                    (c >= 0x61 && c <= 0x7A)
                    // 0-9
                    || (c >= 0x30 && c <= 0x39)
                    // A-Z
                    || (c >= 0x41 && c <= 0x5A)
                    // ".", "_", "+", "-"
                    || (c == 0x2E || c == 0x5F || c == 0x2B || c == 0x2D)
                    // "#", "[", "]", "<", ">"
                    || (c == 0x23 || c == 0x5B || c == 0x5D || c == 0x3C || c == 0x3E)
                ) {
                    continue;
                }
            }
        }
        return true;
    }
}
