// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";

interface IIBCHostConfigurator {
    /**
     * @dev setExpectedTimePerBlock sets expected time per block.
     * Typically this function should be called by an authority like an IBC contract owner or govenance.
     */
    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) external;

    /**
     * @dev registerClient registers a new client type into the client registry
     * Typically this function should be called by an authority like an IBC contract owner or govenance.
     * The authority should verify the light client contract is a valid implementation as follows:
     * - The contract implements ILightClient
     * - To avoid reentrancy attack, the contract never performs `call` to the IBC contract directly or indirectly in the `verifyMembership` and the `verifyNonMembership`
     */
    function registerClient(string calldata clientType, ILightClient client) external;

    /**
     * @dev bindPort binds to an unallocated port, failing if the port has already been allocated.
     * Typically this function should be called by an authority like an IBC contract owner or govenance.
     * The authority should verify the light client contract is a valid implementation as follows:
     * - The contract implements IIBCModule
     */
    function bindPort(string calldata portId, IIBCModule moduleAddress) external;
}
