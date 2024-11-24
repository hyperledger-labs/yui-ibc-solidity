// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";
import {IBCStore} from "./IBCStore.sol";
import {IIBCHostErrors} from "./IIBCHostErrors.sol";

contract IBCHost is IBCStore, IIBCHostErrors {
    /// @dev It represents the prefix of the commitment proof(https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments#prefix).
    ///      In ibc-solidity, the prefix is not required, but for compatibility with ibc-go this must be a non-empty value.
    bytes internal constant DEFAULT_COMMITMENT_PREFIX = bytes("ibc");
    /// @dev It represents the default revision number.
    uint64 internal constant DEFAULT_REVISION_NUMBER = 0;

    /**
     * @dev hostTimestamp returns the current timestamp(Unix time in nanoseconds) of the host chain.
     */
    function hostTimestamp() internal view returns (uint64) {
        return uint64(block.timestamp) * 1e9;
    }

    /**
     * @dev hostHeight returns the current height of the host chain.
     */
    function hostHeight() internal view returns (Height.Data memory) {
        return Height.Data({revision_number: _getRevisionNumber(), revision_height: uint64(block.number)});
    }

    /**
     * @dev checkAndGetClient returns the client implementation for the given client ID.
     */
    function checkAndGetClient(string memory clientId) internal view returns (ILightClient) {
        address clientImpl = getClientStorage()[clientId].clientImpl;
        if (clientImpl == address(0)) {
            revert IIBCClientErrors.IBCClientClientNotFound(clientId);
        }
        return ILightClient(clientImpl);
    }

    /**
     * @dev _getCommitmentPrefix returns the prefix of the commitment proof.
     */
    function _getCommitmentPrefix() internal view virtual returns (bytes memory) {
        return DEFAULT_COMMITMENT_PREFIX;
    }

    /**
     * @dev _getRevisionNumber returns the revision number of the host chain.
     */
    function _getRevisionNumber() internal view virtual returns (uint64) {
        return DEFAULT_REVISION_NUMBER;
    }
}
