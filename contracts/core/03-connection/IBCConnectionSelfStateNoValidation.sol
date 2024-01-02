// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {Height} from "../../proto/Client.sol";
import {IBCConnection} from "./IBCConnection.sol";

/**
 * @dev IBCConnectionSelfStateNoValidation is an IBCConnection that does not validate the self client state in the connection handshake.
 */
contract IBCConnectionSelfStateNoValidation is IBCConnection {
    /**
     * @dev validateSelfClient always returns true
     */
    function validateSelfClient(bytes calldata) public pure override returns (bool) {
        return true;
    }

    /**
     * @dev getSelfConsensusState gets the consensus state of the host chain.
     *
     * NOTE: Developers can override this function to support an arbitrary EVM chain.
     */
    function getSelfConsensusState(Height.Data calldata, bytes calldata hostConsensusStateProof)
        public
        pure
        override
        returns (bytes memory)
    {
        require(hostConsensusStateProof.length != 0, "hostConsensusStateProof cannot be empty");
        return hostConsensusStateProof;
    }
}
