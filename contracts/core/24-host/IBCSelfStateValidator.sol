// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";

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

    /**
     * @dev getSelfConsensusState gets the consensus state of the host chain.
     *
     * NOTE: Developers can override this function to support an arbitrary EVM chain.
     */
    function getSelfConsensusState(Height.Data calldata consensusHeight, bytes calldata hostConsensusStateProof)
        public
        view
        virtual
        returns (bytes memory);
}
