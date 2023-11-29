// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";

interface IIBCHostConfigurator {
    /**
     * @dev setExpectedTimePerBlock sets expected time per block.
     */
    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external;

    /**
     * @dev registerClient registers a new client type into the client registry
     */
    function registerClient(string calldata clientType, ILightClient client) external;

    /**
     * @dev bindPort binds to an unallocated port, failing if the port has already been allocated.
     */
    function bindPort(string calldata portId, IIBCModule moduleAddress) external;
}
