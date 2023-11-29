// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {IBCConnection} from "./IBCConnection.sol";

/**
 * @dev IBCConnectionSelfStateNoValidation is an IBCConnection that does not validate the self client state in the connection handshake.
 */
contract IBCConnectionSelfStateNoValidation is IBCConnection {
    /**
     * @dev validateSelfClient always returns true
     */
    function validateSelfClient(bytes calldata) public view override returns (bool) {
        this; // this is a trick that suppresses "Warning: Function state mutability can be restricted to pure"
        return true;
    }
}
