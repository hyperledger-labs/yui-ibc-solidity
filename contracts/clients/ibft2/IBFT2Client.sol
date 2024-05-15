// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ILightClient} from "../../core/02-client/ILightClient.sol";
import {ILightClientErrors} from "../../core/02-client/ILightClientErrors.sol";
import {IBCHeight} from "../../core/02-client/IBCHeight.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {Height} from "../../proto/Client.sol";
import {
    IbcLightclientsIbft2V1ClientState as ClientState,
    IbcLightclientsIbft2V1ConsensusState as ConsensusState,
    IbcLightclientsIbft2V1Header as Header
} from "../../proto/IBFT2.sol";
import {GoogleProtobufAny as Any} from "../../proto/GoogleProtobufAny.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {RLPReader} from "solidity-rlp/contracts/RLPReader.sol";
import {MPTProof} from "solidity-mpt/src/MPTProof.sol";

/// @notice please see docs/ibft2-light-client.md for client spec
contract IBFT2Client is ILightClient, ILightClientErrors {
    using MPTProof for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;
    using IBCHeight for Height.Data;

    struct ParsedBesuHeader {
        Header.Data base;
        Height.Data height;
        bytes32 stateRoot;
        uint64 time;
        RLPReader.RLPItem[] validators;
    }

    /// @dev An error indicating that the IBC address of the initial client state is invalid
    error InvalidIBCAddressLength();
    /// @dev An error indicating that the initial consensus state root length is invalid
    error InvalidConsensusStateRootLength();
    /// @dev An error indicating that the latest height of the client state is invalid
    error InvalidClientStateLatestHeight();
    /// @dev An error indicating that the validator set is empty
    error EmptyValidators();
    /// @dev An error indicating that a validator address length is invalid
    error InvalidValidatorAddressLength();
    /// @param itemsLength length of items in the header
    error UnexpectedEthereumHeaderFormat(uint256 itemsLength);
    /// @param itemsLength length of items in the extra data
    error UnexpectedExtraDataFormat(uint256 itemsLength);
    /// @param actual actual number of valid seals
    /// @param threshold threshold of seals
    error InsufficientTrustedValidatorsSeals(uint256 actual, uint256 threshold);
    /// @param actual actual number of valid seals
    /// @param threshold threshold of seals
    error InsuffientUntrustedValidatorsSeals(uint256 actual, uint256 threshold);
    /// @param length length of the signature
    error InvalidECDSASignatureLength(uint256 length);

    string internal constant HEADER_TYPE_URL = "/ibc.lightclients.ibft2.v1.Header";
    string internal constant CLIENT_STATE_TYPE_URL = "/ibc.lightclients.ibft2.v1.ClientState";
    string internal constant CONSENSUS_STATE_TYPE_URL = "/ibc.lightclients.ibft2.v1.ConsensusState";

    bytes32 internal constant HEADER_TYPE_URL_HASH = keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.Header"));
    bytes32 internal constant CLIENT_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ClientState"));
    bytes32 internal constant CONSENSUS_STATE_TYPE_URL_HASH =
        keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ConsensusState"));

    uint256 internal constant COMMITMENT_SLOT = 0;
    uint8 internal constant ACCOUNT_STORAGE_ROOT_INDEX = 2;

    address public immutable ibcHandler;

    mapping(string => ClientState.Data) internal clientStates;
    mapping(string => mapping(uint128 => ConsensusState.Data)) internal consensusStates;
    mapping(string => mapping(uint128 => uint256)) internal processedTimes;
    mapping(string => mapping(uint128 => uint256)) internal processedHeights;

    modifier onlyIBC() {
        if (msg.sender != ibcHandler) {
            revert LightClientInvalidCaller(msg.sender);
        }
        _;
    }

    constructor(address ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    /**
     * @dev initializeClient initializes a new client with the given state
     */
    function initializeClient(
        string calldata clientId,
        bytes calldata protoClientState,
        bytes calldata protoConsensusState
    ) public override onlyIBC returns (Height.Data memory height) {
        ClientState.Data memory clientState = ClientState.decode(decode(protoClientState, CLIENT_STATE_TYPE_URL_HASH));
        ConsensusState.Data memory consensusState =
            ConsensusState.decode(decode(protoConsensusState, CONSENSUS_STATE_TYPE_URL_HASH));
        if (clientState.ibc_store_address.length != 20) {
            revert InvalidIBCAddressLength();
        }
        if (clientState.latest_height.revision_height == 0) {
            revert InvalidClientStateLatestHeight();
        }
        if (consensusState.root.length != 32) {
            revert InvalidConsensusStateRootLength();
        }
        if (consensusState.validators.length == 0) {
            revert EmptyValidators();
        }
        clientStates[clientId] = clientState;
        consensusStates[clientId][clientState.latest_height.toUint128()] = consensusState;
        return clientState.latest_height;
    }

    /**
     * @dev routeUpdateClient returns the calldata to the receiving function of the client message.
     *      The light client encodes a client message as ethereum ABI.
     */
    function routeUpdateClient(string calldata clientId, bytes calldata protoClientMessage)
        public
        pure
        virtual
        override
        returns (bytes4, bytes memory)
    {
        Header.Data memory header = Header.decode(decode(protoClientMessage, HEADER_TYPE_URL_HASH));
        return (this.updateClient.selector, abi.encode(clientId, header));
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     *      The timestamp is nanoseconds since unix epoch.
     */
    function getTimestampAtHeight(string calldata clientId, Height.Data calldata height)
        public
        view
        override
        returns (uint64)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            revert LightClientConsensusStateNotFound(clientId, height);
        }
        // ConsensState.timestamp is seconds since unix epoch, so need to convert it to nanoseconds
        return consensusState.timestamp * 1e9;
    }

    /**
     * @dev getLatestHeight returns the latest height of the client state corresponding to `clientId`.
     */
    function getLatestHeight(string calldata clientId) public view override returns (Height.Data memory) {
        ClientState.Data storage clientState = clientStates[clientId];
        if (clientState.latest_height.revision_height == 0) {
            revert LightClientClientStateNotFound(clientId);
        }
        return clientState.latest_height;
    }

    /**
     * @dev getStatus returns the status of the client corresponding to `clientId`.
     */
    function getStatus(string calldata) public pure override returns (ILightClient.ClientStatus) {
        return ILightClient.ClientStatus.Active;
    }

    /**
     * @dev getLatestInfo returns the latest height, timestamp and status of the client corresponding to `clientId`.
     */
    function getLatestInfo(string calldata clientId)
        public
        view
        returns (Height.Data memory latestHeight, uint64 latestTimestamp, ClientStatus status)
    {
        latestHeight = getLatestHeight(clientId);
        latestTimestamp = consensusStates[clientId][latestHeight.toUint128()].timestamp;
        status = ILightClient.ClientStatus.Active;
    }

    /**
     * @dev updateClient updates the client with the given header
     */
    function updateClient(string calldata clientId, Header.Data calldata header)
        public
        returns (Height.Data[] memory heights)
    {
        ClientState.Data storage clientState = clientStates[clientId];
        assert(clientState.ibc_store_address.length != 0);

        ParsedBesuHeader memory parsedHeader = parseBesuHeader(header);
        uint128 newHeight = parsedHeader.height.toUint128();

        ConsensusState.Data storage trustedConsensusState = consensusStates[clientId][header.trusted_height.toUint128()];
        if (trustedConsensusState.timestamp == 0) {
            revert LightClientConsensusStateNotFound(clientId, header.trusted_height);
        }
        if (
            clientState.trusting_period > 0
                && trustedConsensusState.timestamp + clientState.trusting_period <= block.timestamp
        ) {
            revert LightClientConsensusStateExpired();
        }

        bytes[] memory validators = verify(trustedConsensusState.validators, parsedHeader);
        if (validators.length == 0) {
            revert EmptyValidators();
        }
        if (parsedHeader.height.gt(clientState.latest_height)) {
            clientState.latest_height = parsedHeader.height;
        }
        ConsensusState.Data storage consensusState = consensusStates[clientId][newHeight];
        consensusState.timestamp = parsedHeader.time;
        consensusState.root = abi.encodePacked(
            verifyStorageProof(
                address(bytes20(clientState.ibc_store_address)), parsedHeader.stateRoot, header.account_state_proof
            )
        );
        consensusState.validators = validators;

        processedTimes[clientId][newHeight] = block.timestamp;
        processedHeights[clientId][newHeight] = block.number;

        heights = new Height.Data[](1);
        heights[0] = parsedHeader.height;
        return heights;
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
        bytes memory path,
        bytes calldata value
    ) public view override returns (bool) {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            revert LightClientConsensusStateNotFound(clientId, height);
        }
        if (!validateArgsAndDelayPeriod(clientId, height, delayTimePeriod, delayBlockPeriod, prefix)) {
            return false;
        }
        return verifyMembership(
            proof,
            bytes32(consensusState.root),
            keccak256(abi.encodePacked(keccak256(path), COMMITMENT_SLOT)),
            keccak256(value)
        );
    }

    /**
     * @dev verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
     * The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
     */
    function verifyNonMembership(
        string calldata clientId,
        Height.Data calldata height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes calldata proof,
        bytes calldata prefix,
        bytes calldata path
    ) public view override returns (bool) {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            revert LightClientConsensusStateNotFound(clientId, height);
        }
        if (!validateArgsAndDelayPeriod(clientId, height, delayTimePeriod, delayBlockPeriod, prefix)) {
            return false;
        }
        return verifyNonMembership(
            proof, bytes32(consensusState.root), keccak256(abi.encodePacked(keccak256(path), COMMITMENT_SLOT))
        );
    }

    /**
     * @dev getClientState returns the clientState corresponding to `clientId`.
     *      If it's not found, the function returns false.
     */
    function getClientState(string calldata clientId) public view returns (bytes memory clientStateBytes, bool) {
        ClientState.Data storage clientState = clientStates[clientId];
        if (clientState.latest_height.revision_height == 0) {
            return (clientStateBytes, false);
        }
        return (Any.encode(Any.Data({type_url: CLIENT_STATE_TYPE_URL, value: ClientState.encode(clientState)})), true);
    }

    /**
     * @dev getConsensusState returns the consensusState corresponding to `clientId` and `height`.
     *      If it's not found, the function returns false.
     */
    function getConsensusState(string calldata clientId, Height.Data calldata height)
        public
        view
        returns (bytes memory consensusStateBytes, bool)
    {
        ConsensusState.Data storage consensusState = consensusStates[clientId][height.toUint128()];
        if (consensusState.timestamp == 0) {
            return (consensusStateBytes, false);
        }
        return (
            Any.encode(Any.Data({type_url: CONSENSUS_STATE_TYPE_URL, value: ConsensusState.encode(consensusState)})),
            true
        );
    }

    // --------------------- Signature verification --------------------- //

    /**
     * @dev verify verifies untrusted header
     * @param trustedVals trusted validators
     * @param untrustedHeader untrusted header
     */
    function verify(bytes[] storage trustedVals, ParsedBesuHeader memory untrustedHeader)
        internal
        view
        returns (bytes[] memory)
    {
        bytes32 blkHash = keccak256(untrustedHeader.base.besu_header_rlp);
        verifyCommitSealsTrusting(trustedVals, untrustedHeader.base.seals, blkHash);
        return verifyCommitSeals(untrustedHeader.validators, untrustedHeader.base.seals, blkHash);
    }

    /**
     * @dev verifyCommitSealsTrusting verifies that trustLevel of the validator set signed this commit.
     * @param trustedVals trusted validators
     * @param seals commit seals for untrusted block header
     * @param blkHash the hash of untrusted block
     */
    function verifyCommitSealsTrusting(bytes[] storage trustedVals, bytes[] memory seals, bytes32 blkHash)
        internal
        view
    {
        uint256 success = 0;
        uint256 valsNum = trustedVals.length;
        bool[] memory marked = new bool[](valsNum);
        for (uint256 i = 0; i < seals.length; i++) {
            if (seals[i].length == 0) {
                continue;
            }
            address signer = ecdsaRecover(blkHash, seals[i]);
            for (uint256 j = 0; j < valsNum; j++) {
                if (trustedVals[j].length != 20) {
                    revert InvalidValidatorAddressLength();
                }
                if (!marked[j] && address(bytes20(trustedVals[j])) == signer) {
                    success++;
                    marked[j] = true;
                }
            }
        }
        if (success < valsNum / 3) {
            revert InsufficientTrustedValidatorsSeals(success, valsNum / 3);
        }
    }

    /**
     * @dev verifyCommitSeals verifies the seals with untrustedVals. The order of seals must match the order of untrustedVals.
     * @param untrustedVals validators of untrusted block header
     * @param seals commit seals for untrusted block header
     * @param blkHash the hash of untrusted block
     */
    function verifyCommitSeals(RLPReader.RLPItem[] memory untrustedVals, bytes[] memory seals, bytes32 blkHash)
        internal
        pure
        returns (bytes[] memory)
    {
        uint256 valsNum = untrustedVals.length;
        bytes[] memory validators = new bytes[](valsNum);
        uint256 success = 0;
        for (uint256 i = 0; i < seals.length; i++) {
            validators[i] = untrustedVals[i].toBytes();
            if (validators[i].length != 20) {
                revert InvalidValidatorAddressLength();
            }
            if (seals[i].length == 0) {
                continue;
            } else if (address(bytes20(validators[i])) == ecdsaRecover(blkHash, seals[i])) {
                success++;
            }
        }
        if (success < valsNum * 2 / 3) {
            revert InsuffientUntrustedValidatorsSeals(success, valsNum * 2 / 3);
        }
        return validators;
    }

    // --------------------- Internal functions --------------------- //

    function validateDelayPeriod(
        string memory clientId,
        Height.Data memory height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks
    ) internal view returns (bool) {
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

    function validateArgsAndDelayPeriod(
        string memory clientId,
        Height.Data memory height,
        uint64 delayTimePeriod,
        uint64 delayBlockPeriod,
        bytes memory prefix
    ) internal view returns (bool) {
        ClientState.Data storage clientState = clientStates[clientId];
        if (clientState.latest_height.lt(height)) {
            return false;
        }
        if (
            (delayTimePeriod != 0 || delayBlockPeriod != 0)
                && !validateDelayPeriod(clientId, height, delayTimePeriod, delayBlockPeriod)
        ) {
            return false;
        }

        return keccak256(IIBCHandler(ibcHandler).getCommitmentPrefix()) == keccak256(prefix);
    }

    function verifyMembership(bytes calldata proof, bytes32 root, bytes32 slot, bytes32 expectedValue)
        internal
        pure
        returns (bool)
    {
        bytes32 path = keccak256(abi.encodePacked(slot));
        bytes memory dataHash = proof.verifyRLPProof(root, path); // reverts if proof is invalid
        return expectedValue == bytes32(dataHash.toRlpItem().toUint());
    }

    function verifyNonMembership(bytes calldata proof, bytes32 root, bytes32 slot) internal pure returns (bool) {
        bytes32 path = keccak256(abi.encodePacked(slot));
        bytes memory dataHash = proof.verifyRLPProof(root, path); // reverts if proof is invalid
        return dataHash.length == 0;
    }

    function parseBesuHeader(Header.Data memory header) internal pure returns (ParsedBesuHeader memory) {
        ParsedBesuHeader memory parsedHeader;

        parsedHeader.base = header;
        RLPReader.RLPItem[] memory items = header.besu_header_rlp.toRlpItem().toList();
        parsedHeader.stateRoot = bytes32(items[3].toUint());
        parsedHeader.height = Height.Data({revision_number: 0, revision_height: uint64(items[8].toUint())});

        if (items.length != 15) {
            revert UnexpectedEthereumHeaderFormat(items.length);
        }
        parsedHeader.time = uint64(items[11].toUint());
        items = items[12].toBytes().toRlpItem().toList();
        if (items.length != 4) {
            revert UnexpectedExtraDataFormat(items.length);
        }
        parsedHeader.validators = items[1].toList();
        return parsedHeader;
    }

    function verifyStorageProof(address account, bytes32 stateRoot, bytes memory accountStateProof)
        internal
        pure
        returns (bytes32)
    {
        bytes32 proofPath = keccak256(abi.encodePacked(account));
        bytes memory accountRLP = accountStateProof.verifyRLPProof(stateRoot, proofPath); // reverts if proof is invalid
        return bytes32(accountRLP.toRlpItem().toList()[ACCOUNT_STORAGE_ROOT_INDEX].toUint());
    }

    function ecdsaRecover(bytes32 hash, bytes memory sig) private pure returns (address) {
        if (sig.length != 65) {
            revert InvalidECDSASignatureLength(sig.length);
        } else if (uint8(sig[64]) < 27) {
            sig[64] = bytes1(uint8(sig[64]) + 27);
        }
        (address signer, ECDSA.RecoverError e,) = ECDSA.tryRecover(hash, sig);
        if (e != ECDSA.RecoverError.NoError) {
            return address(0);
        }
        return signer;
    }

    function decode(bytes calldata bz, bytes32 hashTypeUrl) internal pure returns (bytes memory) {
        Any.Data memory any = Any.decode(bz);
        if (keccak256(abi.encodePacked(any.type_url)) != hashTypeUrl) {
            revert LightClientUnexpectedProtoAnyTypeURL(any.type_url);
        }
        return any.value;
    }
}
