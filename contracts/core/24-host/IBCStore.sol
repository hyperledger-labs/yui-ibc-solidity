// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../proto/Connection.sol";
import "../../proto/Channel.sol";
import "../02-client/ILightClient.sol";

abstract contract IBCStore {
    // Commitments
    // keccak256(IBC-compatible-store-path) => keccak256(IBC-compatible-commitment)
    mapping(bytes32 => bytes32) internal commitments;

    // Store
    mapping(string => address) internal clientRegistry; // clientType => clientImpl
    mapping(string => string) internal clientTypes; // clientID => clientType
    mapping(string => address) internal clientImpls; // clientID => clientImpl
    mapping(string => ConnectionEnd.Data) internal connections;
    mapping(string => mapping(string => Channel.Data)) internal channels;
    mapping(string => mapping(string => uint64)) internal nextSequenceSends;
    mapping(string => mapping(string => uint64)) internal nextSequenceRecvs;
    mapping(string => mapping(string => uint64)) internal nextSequenceAcks;
    mapping(bytes => address[]) internal capabilities;

    // Host parameters
    uint64 internal expectedTimePerBlock;

    // Sequences for identifier
    uint64 internal nextClientSequence;
    uint64 internal nextConnectionSequence;
    uint64 internal nextChannelSequence;

    // It represents the prefix of the commitment proof(https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments#prefix).
    // In ibc-solidity, the prefix is not required, but for compatibility with ibc-go this must be a non-empty value.
    bytes internal constant DEFAULT_COMMITMENT_PREFIX = bytes("ibc");

    /**
     * @dev getCommitmentPrefix returns the prefix of the commitment proof.
     */
    function getCommitmentPrefix() public pure virtual returns (bytes memory) {
        return DEFAULT_COMMITMENT_PREFIX;
    }

    // Storage accessors

    /**
     * @dev checkAndGetClient returns the client implementation for the given client ID.
     */
    function checkAndGetClient(string memory clientId) internal view returns (ILightClient) {
        address clientImpl = clientImpls[clientId];
        require(clientImpl != address(0));
        return ILightClient(clientImpl);
    }
}
