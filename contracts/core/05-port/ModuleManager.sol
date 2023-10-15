// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "./IIBCModule.sol";

/**
 * @dev ModuleManager is an abstract contract that provides the functions defined in [ICS 5](https://github.com/cosmos/ibc/tree/main/spec/core/ics-005-port-allocation) and [ICS 26](https://github.com/cosmos/ibc/blob/main/spec/core/ics-005-port-module/README.md).
 */
abstract contract ModuleManager {
    /**
     * @dev bindPort binds to an unallocated port, failing if the port has already been allocated.
     */
    function bindPort(string calldata portId, address moduleAddress) public virtual {
        require(validatePortIdentifier(bytes(portId)), "invalid portId");
        require(moduleAddress != address(this) && Address.isContract(moduleAddress), "invalid moduleAddress");
        claimCapability(portCapabilityPath(portId), moduleAddress);
    }

    /**
     * @dev lookupModuleByPort will return the IBCModule along with the capability associated with a given portID
     */
    function lookupModuleByPort(string memory portId) internal view virtual returns (IIBCModule) {
        (address[] storage modules, bool found) = lookupModules(portCapabilityPath(portId));
        require(found);
        return IIBCModule(modules[0]);
    }

    /**
     * @dev lookupModuleByChannel will return the IBCModule along with the capability associated with a given channel defined by its portID and channelID
     */
    function lookupModuleByChannel(string memory portId, string memory channelId)
        internal
        view
        virtual
        returns (IIBCModule)
    {
        (address[] storage modules, bool found) = lookupModules(channelCapabilityPath(portId, channelId));
        require(found);
        return IIBCModule(modules[0]);
    }

    /**
     * @dev portCapabilityPath returns the path under which owner module address associated with a port should be stored.
     */
    function portCapabilityPath(string memory portId) public pure returns (bytes memory) {
        return abi.encodePacked(portId);
    }

    /**
     * @dev channelCapabilityPath returns the path under which module address associated with a port and channel should be stored.
     */
    function channelCapabilityPath(string memory portId, string memory channelId) public pure returns (bytes memory) {
        return abi.encodePacked(portId, "/", channelId);
    }

    /**
     * @dev claimCapability allows the IBC app module to claim a capability that core IBC passes to it
     */
    function claimCapability(bytes memory name, address addr) internal virtual;

    /**
     * @dev authenticateCapability attempts to authenticate a given name from a caller.
     * It allows for a caller to check that a capability does in fact correspond to a particular name.
     */
    function authenticateCapability(bytes memory name) internal view virtual returns (bool);

    /**
     * @dev lookupModule will return the IBCModule address bound to a given name.
     * Currently, the function returns only one module.
     */
    function lookupModules(bytes memory name) internal view virtual returns (address[] storage, bool);

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
