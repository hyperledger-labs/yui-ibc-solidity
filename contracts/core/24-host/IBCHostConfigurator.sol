// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {IBCClientLib} from "../02-client/IBCClientLib.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";
import {IIBCHostConfigurator} from "./IIBCHostConfigurator.sol";

/**
 * @dev IBCHostConfigurator is a contract that provides the host configuration.
 */
abstract contract IBCHostConfigurator is IBCModuleManager, IIBCHostConfigurator {
    function _setExpectedTimePerBlock(uint64 expectedTimePerBlock_) internal virtual {
        expectedTimePerBlock = expectedTimePerBlock_;
    }

    function _registerClient(string calldata clientType, ILightClient client) internal virtual {
        require(IBCClientLib.validateClientType(bytes(clientType)), "invalid clientType");
        require(address(clientRegistry[clientType]) == address(0), "clientType already exists");
        require(address(client) != address(this) && Address.isContract(address(client)), "invalid client address");
        clientRegistry[clientType] = address(client);
    }

    function _bindPort(string calldata portId, IIBCModule moduleAddress) internal virtual {
        require(validatePortIdentifier(bytes(portId)), "invalid portId");
        require(
            address(moduleAddress) != address(this) && Address.isContract(address(moduleAddress)),
            "invalid moduleAddress"
        );
        claimCapability(portCapabilityPath(portId), address(moduleAddress));
    }
}
