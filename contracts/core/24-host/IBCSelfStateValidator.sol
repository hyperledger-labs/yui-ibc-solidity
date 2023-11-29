// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

/**
 * @dev IBCSelfStateValidator is an interface that validates the self client state in the connection handshake.
 */
abstract contract IBCSelfStateValidator {
    /**
     * @dev validateSelfClient validates the client parameters for a client of the host chain.
     *
     * NOTE: Developers can override this function to support an arbitrary EVM chain.
     */
    function validateSelfClient(bytes calldata clientState) public view virtual returns (bool);
}
