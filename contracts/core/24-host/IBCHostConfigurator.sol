// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ERC165Checker} from "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import {IBCClientLib} from "../02-client/IBCClientLib.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";
import {IIBCHostConfigurator} from "./IIBCHostConfigurator.sol";

/**
 * @dev IBCHostConfigurator is a contract that provides the host configuration.
 */
abstract contract IBCHostConfigurator is IIBCHostConfigurator, IBCModuleManager {
    function _setExpectedTimePerBlock(uint64 expectedTimePerBlock_) internal virtual {
        expectedTimePerBlock = expectedTimePerBlock_;
    }

    function _registerClient(string calldata clientType, ILightClient client) internal virtual {
        if (!IBCClientLib.validateClientType(bytes(clientType))) {
            revert IBCHostInvalidClientType(clientType);
        } else if (address(clientRegistry[clientType]) != address(0)) {
            revert IBCHostClientTypeAlreadyExists(clientType);
        }
        if (address(client) == address(0) || address(client) == address(this)) {
            revert IBCHostInvalidLightClientAddress(address(client));
        }
        clientRegistry[clientType] = address(client);
    }

    function _bindPort(string calldata portId, IIBCModule module) internal virtual {
        address moduleAddress = address(module);
        if (!validatePortIdentifier(bytes(portId))) {
            revert IBCHostInvalidPortIdentifier(portId);
        }
        if (moduleAddress == address(0) || moduleAddress == address(this)) {
            revert IBCHostInvalidModuleAddress(moduleAddress);
        }
        if (!ERC165Checker.supportsERC165(moduleAddress)) {
            revert IBCHostModuleDoesNotSupportERC165();
        }
        if (!ERC165Checker.supportsERC165InterfaceUnchecked(moduleAddress, type(IIBCModule).interfaceId)) {
            revert IBCHostModuleDoesNotSupportIIBCModule(type(IIBCModule).interfaceId);
        }
        claimPortCapability(portId, moduleAddress);
    }
}
