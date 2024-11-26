// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IBCClientLib} from "../02-client/IBCClientLib.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCHostLib} from "./IBCHostLib.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCModuleInitializer} from "../26-router/IIBCModule.sol";
import {IIBCHostConfigurator} from "./IIBCHostConfigurator.sol";

/**
 * @dev IBCHostConfigurator is a contract that provides the host configuration.
 */
abstract contract IBCHostConfigurator is IIBCHostConfigurator, IBCModuleManager {
    function _setExpectedTimePerBlock(uint64 expectedTimePerBlock_) internal virtual {
        getHostStorage().expectedTimePerBlock = expectedTimePerBlock_;
    }

    function _registerClient(string calldata clientType, ILightClient client) internal virtual {
        HostStorage storage hostStorage = getHostStorage();
        if (!IBCClientLib.validateClientType(bytes(clientType))) {
            revert IBCHostInvalidClientType(clientType);
        } else if (address(hostStorage.clientRegistry[clientType]) != address(0)) {
            revert IBCHostClientTypeAlreadyExists(clientType);
        }
        if (address(client) == address(0) || address(client) == address(this)) {
            revert IBCHostInvalidLightClientAddress(address(client));
        }
        hostStorage.clientRegistry[clientType] = address(client);
    }

    function _bindPort(string calldata portId, IIBCModuleInitializer module) internal virtual {
        if (!IBCHostLib.validatePortIdentifier(bytes(portId))) {
            revert IBCHostInvalidPortIdentifier(portId);
        }
        claimPortCapability(portId, address(module));
    }
}
