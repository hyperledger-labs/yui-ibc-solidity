// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IClient.sol";
import "./IBCHeight.sol";
import "../proto/Client.sol";
import {
    IbcLightclientsIbft2V1ClientState as ClientState,
    IbcLightclientsIbft2V1ConsensusState as ConsensusState,
    IbcLightclientsIbft2V1Header as Header
} from "../proto/IBFT2.sol";
import {GoogleProtobufAny as Any} from "../proto/GoogleProtobufAny.sol";
import "../lib/ECRecovery.sol";
import "../lib/Bytes.sol";
import "../lib/TrieProofs.sol";
import "../lib/RLP.sol";

// please see docs/ibft2-light-client.md for client spec
contract IBFT2Client is IClient {
    using TrieProofs for bytes;
    using RLP for RLP.RLPItem;
    using RLP for bytes;
    using Bytes for bytes;
    using IBCHeight for Height.Data;

    uint256 private constant COMMITMENT_SLOT = 0;
    uint8 private constant ACCOUNT_STORAGE_ROOT_INDEX = 2;
    bytes32 private constant HEADER_TYPE_URL_HASH = keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.Header"));
    bytes32 private constant CLIENT_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ClientState"));
    bytes32 private constant CONSENSUS_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ConsensusState"));

    mapping(string => ClientState.Data) internal clientStates;
    mapping(string => mapping(uint128 => ConsensusState.Data)) internal consensusStates;
    mapping(string => mapping(uint128 => uint256)) internal processedTimes;
    mapping(string => mapping(uint128 => uint256)) internal processedHeights;

    struct ParsedBesuHeader {
        Header.Data base;
        Height.Data height;
        bytes32 stateRoot;
        uint64 time;
        RLP.RLPItem[] validators;
    }

    struct Fraction {
        uint64 numerator;
        uint64 denominator;
    }

    function createClient(
        string calldata clientId,
        Height.Data calldata height,
        bytes calldata clientStateBytes,
        bytes calldata consensusStateBytes
    ) external override returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;

        (clientState, ok) = unmarshalClientState(clientStateBytes);
        if (!ok) {
            return (clientStateCommitment, updates, false);
        }
        (consensusState, ok) = unmarshalConsensusState(consensusStateBytes);
        if (!ok) {
            return (clientStateCommitment, updates, false);
        }
        clientStates[clientId] = clientState;
        consensusStates[clientId][height.toUint128()] = consensusState;
        updates = new ConsensusStateUpdates[](1);
        updates[0] = ConsensusStateUpdates({consensusStateCommitment: keccak256(consensusStateBytes), height: height});
        return (keccak256(clientStateBytes), updates, true);
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        external
        view
        override
        returns (uint64, bool)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        return (consensusState.timestamp, consensusState.timestamp != 0);
    }

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) external view override returns (Height.Data memory, bool) {
        ClientState.Data storage clientState = clientStates[clientId];
        return (clientState.latest_height, clientState.latest_height.revision_height != 0);
    }

    /**
     * @dev verifyClientMessageAndUpdateState is intended to perform the followings:
     * 1. verify a given client message(e.g. header)
     * 2. check misbehaviour such like duplicate block height
     * 3. if misbehaviour is found, update state accordingly and return
     * 4. update state(s) with the client message
     * 5. persist the state(s) on the host
     */
    function verifyClientMessageAndUpdateState(string calldata clientId, bytes calldata clientMessageBytes)
        external
        override
        returns (bytes32 clientStateCommitment, ConsensusStateUpdates[] memory updates, bool ok)
    {
        Header.Data memory header;
        bytes[] memory validators;
        uint128 newHeight;

        /* Validation */

        ClientState.Data storage clientState = clientStates[clientId];
        assert(clientState.ibc_store_address.length != 0);

        // TODO add misbehaviour check support
        (header, ok) = unmarshalHeader(clientMessageBytes);
        require(ok, "header is invalid");
        // check if the provided client message is valid
        ParsedBesuHeader memory parsedHeader = parseBesuHeader(header);
        require(parsedHeader.height.gt(header.trusted_height), "header height <= consensus state height");
        newHeight = parsedHeader.height.toUint128();

        /* State verification */

        ConsensusState.Data storage consensusState = consensusStates[clientId][header.trusted_height.toUint128()];
        assert(consensusState.timestamp != 0);
        (validators, ok) = verify(consensusState.validators, parsedHeader);
        require(ok, "failed to verify the header");

        /* Update states */

        if (parsedHeader.height.gt(clientState.latest_height)) {
            clientState.latest_height = parsedHeader.height;
        }
        // if client message is verified and there is no misbehaviour, update state
        consensusState = consensusStates[clientId][newHeight];
        consensusState.timestamp = parsedHeader.time;
        consensusState.root = abi.encodePacked(
            verifyStorageProof(
                Bytes.toAddress(clientState.ibc_store_address), parsedHeader.stateRoot, header.account_state_proof
            )
        );
        consensusState.validators = validators;

        /* Make updates message */
        updates = new ConsensusStateUpdates[](1);
        updates[0] = ConsensusStateUpdates({
            consensusStateCommitment: keccak256(marshalConsensusState(consensusState)),
            height: parsedHeader.height
        });

        processedTimes[clientId][newHeight] = block.timestamp;
        processedHeights[clientId][newHeight] = block.number;

        return (keccak256(marshalClientState(clientState)), updates, true);
    }

    /**
     * @dev verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes calldata proof,
        bytes memory prefix,
        bytes calldata path,
        bytes calldata value
    ) external view override returns (bool) {
        {
            ClientState.Data storage clientState = clientStates[clientId];
            assert(clientState.ibc_store_address.length != 0);

            if (!validateArgs(clientState, height, prefix, proof)) {
                return false;
            }
            if (
                (delayTimePeriod != 0 || delayBlockPeriod != 0)
                    && !validateDelayPeriod(clientId, height, delayTimePeriod, delayBlockPeriod)
            ) {
                return false;
            }
        }

        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        assert(consensusState.timestamp != 0);
        return verifyMembership( // TODO
            proof,
            consensusState.root.toBytes32(),
            keccak256(abi.encodePacked(keccak256(path), COMMITMENT_SLOT)),
            keccak256(value)
        );
    }

    function marshalClientState(ClientState.Data storage clientState) internal pure returns (bytes memory) {
        Any.Data memory anyClientState;
        anyClientState.type_url = "/ibc.lightclients.ibft2.v1.ClientState";
        anyClientState.value = ClientState.encode(clientState);
        return Any.encode(anyClientState);
    }

    function marshalConsensusState(ConsensusState.Data storage consensusState) internal pure returns (bytes memory) {
        Any.Data memory anyConsensusState;
        anyConsensusState.type_url = "/ibc.lightclients.ibft2.v1.ConsensusState";
        anyConsensusState.value = ConsensusState.encode(consensusState);
        return Any.encode(anyConsensusState);
    }

    function unmarshalHeader(bytes memory bz) internal pure returns (Header.Data memory header, bool ok) {
        Any.Data memory anyHeader = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyHeader.type_url)) != HEADER_TYPE_URL_HASH) {
            return (header, false);
        }
        return (Header.decode(anyHeader.value), true);
    }

    function unmarshalClientState(bytes memory bz)
        internal
        pure
        returns (ClientState.Data memory clientState, bool ok)
    {
        Any.Data memory anyClientState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyClientState.type_url)) != CLIENT_STATE_TYPE_URL_HASH) {
            return (clientState, false);
        }
        return (ClientState.decode(anyClientState.value), true);
    }

    function unmarshalConsensusState(bytes memory bz)
        internal
        pure
        returns (ConsensusState.Data memory consensusState, bool ok)
    {
        Any.Data memory anyConsensusState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyConsensusState.type_url)) != CONSENSUS_STATE_TYPE_URL_HASH) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(anyConsensusState.value), true);
    }

    /// Validity predicate ///

    /**
     * @dev verify verifies untrusted header
     * @param trustedVals trusted validators
     * @param untrustedHeader untrusted header
     */
    function verify(bytes[] memory trustedVals, ParsedBesuHeader memory untrustedHeader)
        internal
        pure
        returns (bytes[] memory validators, bool ok)
    {
        bytes32 blkHash = keccak256(untrustedHeader.base.besu_header_rlp);

        if (
            !verifyCommitSealsTrusting(
                trustedVals, untrustedHeader.base.seals, blkHash, Fraction({numerator: 1, denominator: 3})
            )
        ) {
            return (validators, false);
        }

        return verifyCommitSeals(untrustedHeader.validators, untrustedHeader.base.seals, blkHash);
    }

    /**
     * @dev verifyCommitSealsTrusting verifies that trustLevel of the validator set signed this commit.
     * @param trustedVals trusted validators
     * @param seals commit seals for untrusted block header
     * @param blkHash the hash of untrusted block
     * @param trustLevel new header can be trusted if at least one correct validator signed it
     */
    function verifyCommitSealsTrusting(
        bytes[] memory trustedVals,
        bytes[] memory seals,
        bytes32 blkHash,
        Fraction memory trustLevel
    ) internal pure returns (bool) {
        uint8 success = 0;
        bool[] memory marked = new bool[](trustedVals.length);
        for (uint256 i = 0; i < seals.length; i++) {
            if (seals[i].length == 0) {
                continue;
            }
            address signer = ECRecovery.recover(blkHash, seals[i]);
            for (uint256 j = 0; j < trustedVals.length; j++) {
                if (!marked[j] && trustedVals[j].toAddress() == signer) {
                    success++;
                    marked[j] = true;
                }
            }
        }
        return success >= trustedVals.length * trustLevel.numerator / trustLevel.denominator;
    }

    /**
     * @dev verifyCommitSeals verifies the seals with untrustedVals. The order of seals must match the order of untrustedVals.
     * @param untrustedVals validators of untrusted block header
     * @param seals commit seals for untrusted block header
     * @param blkHash the hash of untrusted block
     */
    function verifyCommitSeals(RLP.RLPItem[] memory untrustedVals, bytes[] memory seals, bytes32 blkHash)
        internal
        pure
        returns (bytes[] memory, bool)
    {
        bytes[] memory validators = new bytes[](untrustedVals.length);
        uint8 success = 0;
        for (uint256 i = 0; i < seals.length; i++) {
            validators[i] = untrustedVals[i].toBytes();
            if (seals[i].length == 0) {
                continue;
            } else if (validators[i].toAddress() == ECRecovery.recover(blkHash, seals[i])) {
                success++;
            }
        }
        return (validators, success > untrustedVals.length * 2 / 3);
    }

    /// helper functions ///

    function validateArgs(
        ClientState.Data storage cs,
        Height.Data calldata height,
        bytes memory prefix,
        bytes calldata proof
    ) internal view returns (bool) {
        if (cs.latest_height.lt(height)) {
            return false;
        } else if (prefix.length == 0) {
            return false;
        } else if (proof.length == 0) {
            return false;
        }
        return true;
    }

    function validateDelayPeriod(
        string memory clientId,
        Height.Data calldata height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks
    ) private view returns (bool) {
        uint128 heightU128 = height.toUint128();
        uint64 currentTime = uint64(block.timestamp * 1000 * 1000 * 1000);
        uint64 validTime = uint64(processedTimes[clientId][heightU128]) * 1000 * 1000 * 1000 + delayPeriodTime;
        if (currentTime < validTime) {
            return false;
        }
        uint64 currentHeight = uint64(block.number);
        uint64 validHeight = uint64(processedHeights[clientId][heightU128]) + delayPeriodBlocks;
        if (currentHeight < validHeight) {
            return false;
        }
        return true;
    }

    function verifyMembership(bytes calldata proof, bytes32 root, bytes32 slot, bytes32 expectedValue)
        internal
        pure
        returns (bool)
    {
        bytes32 path = keccak256(abi.encodePacked(slot));
        bytes memory dataHash = proof.verify(root, path); // reverts if proof is invalid
        return expectedValue == dataHash.toRLPItem().toBytes().toBytes32();
    }

    function parseBesuHeader(Header.Data memory header) internal pure returns (ParsedBesuHeader memory) {
        ParsedBesuHeader memory parsedHeader;

        parsedHeader.base = header;
        RLP.RLPItem[] memory items = header.besu_header_rlp.toRLPItem().toList();
        parsedHeader.stateRoot = items[3].toBytes().toBytes32();
        parsedHeader.height = Height.Data({revision_number: 0, revision_height: uint64(items[8].toUint())});

        require(items.length == 15, "items length must be 15");
        parsedHeader.time = uint64(items[11].toUint());
        items = items[12].toBytes().toRLPItem().toList();
        require(items.length == 4, "extra length must be 4");

        parsedHeader.validators = items[1].toList();
        return parsedHeader;
    }

    function verifyStorageProof(address account, bytes32 stateRoot, bytes memory accountStateProof)
        internal
        pure
        returns (bytes32)
    {
        bytes32 proofPath = keccak256(abi.encodePacked(account));
        bytes memory accountRLP = accountStateProof.verify(stateRoot, proofPath); // reverts if proof is invalid
        return bytes32(accountRLP.toRLPItem().toList()[ACCOUNT_STORAGE_ROOT_INDEX].toUint());
    }

    /* State accessors */

    function getClientState(string calldata clientId) external view returns (ClientState.Data memory, bool) {
        ClientState.Data memory clientState = clientStates[clientId];
        return (clientState, clientState.latest_height.revision_height != 0);
    }

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (ConsensusState.Data memory, bool)
    {
        ConsensusState.Data memory consensusState = consensusStates[clientId][height.toUint128()];
        return (consensusState, consensusState.timestamp != 0);
    }
}
