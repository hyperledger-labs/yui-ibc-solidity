pragma solidity ^0.8.9;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import {
    IbcLightclientsIbft2V1ClientState as ClientState,
    IbcLightclientsIbft2V1ConsensusState as ConsensusState,
    IbcLightclientsIbft2V1Header as Header
} from "./types/IBFT2.sol";
import {GoogleProtobufAny as Any} from "./types/GoogleProtobufAny.sol";
import "../lib/ECRecovery.sol";
import "../lib/Bytes.sol";
import "../lib/TrieProofs.sol";
import "../lib/RLP.sol";
import "./IBCIdentifier.sol";

// please see docs/ibft2-light-client.md for client spec
contract IBFT2Client is IClient {
    using TrieProofs for bytes;
    using RLP for RLP.RLPItem;
    using RLP for bytes;
    using Bytes for bytes;

    struct protoTypes {
        bytes32 clientState;
        bytes32 consensusState;
        bytes32 header;
    }

    protoTypes pts;

    constructor() public {
        // TODO The typeUrl should be defined in types/IBFT2Client.sol
        // The schema of typeUrl follows cosmos/cosmos-sdk/codec/types/any.go
        pts = protoTypes({
            clientState: keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ClientState")),
            consensusState: keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.ConsensusState")),
            header: keccak256(abi.encodePacked("/ibc.lightclients.ibft2.v1.Header"))
        });
    }

    struct ParsedBesuHeader {
        Header.Data base;
        uint64 height;
        bytes32 stateRoot;
        uint64 time;
        RLP.RLPItem[] validators;
    }

    struct Fraction {
        uint64 numerator;
        uint64 denominator;
    }

    uint8 private constant ACCOUNT_STORAGE_ROOT_INDEX = 2;

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        uint64 height
    ) public override view returns (uint64, bool) {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, height);
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (0, false);
        }
        return (clientState.latest_height, true);
    }

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId, 
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public override view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height) {
        Header.Data memory header;
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bytes[] memory validators;
        bool ok;

        (clientState, ok) = unmarshalClientState(clientStateBytes);
        require(ok, "client state is invalid");

        (header, ok) = unmarshalHeader(headerBytes);
        require(ok, "header is invalid");

        (consensusState, ok) = getConsensusState(host, clientId, header.trusted_height);
        require(ok, "consensusState not found");

        //// check validity ////
        ParsedBesuHeader memory parsedHeader = parseBesuHeader(header);
        require(parsedHeader.height > header.trusted_height, "header height <= consensus state height");
        (validators, ok) = verify(consensusState, parsedHeader);
        require(ok, "failed to verify the header");

        //// update ////
        consensusState.timestamp = parsedHeader.time;
        consensusState.root = abi.encodePacked(
            verifyStorageProof(Bytes.toAddress(clientState.ibc_store_address), parsedHeader.stateRoot, header.account_state_proof)
        );
        consensusState.validators = validators;

        if (parsedHeader.height > clientState.latest_height) {
            clientState.latest_height = parsedHeader.height;
        }
        return (marshalClientState(clientState), marshalConsensusState(consensusState), parsedHeader.height);
    }

    function marshalClientState(ClientState.Data memory clientState) internal view returns (bytes memory) {
        Any.Data memory anyClientState;
        anyClientState.type_url = "/ibc.lightclients.ibft2.v1.ClientState";
        anyClientState.value = ClientState.encode(clientState);
        return Any.encode(anyClientState);
    }

    function marshalConsensusState(ConsensusState.Data memory consensusState) internal view returns (bytes memory) {
        Any.Data memory anyConsensusState;
        anyConsensusState.type_url = "/ibc.lightclients.ibft2.v1.ConsensusState";
        anyConsensusState.value = ConsensusState.encode(consensusState);
        return Any.encode(anyConsensusState);
    }

    function unmarshalHeader(bytes memory bz) internal view returns (Header.Data memory header, bool ok) {
        Any.Data memory anyHeader = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyHeader.type_url)) != pts.header) {
            return (header, false);
        }
        return (Header.decode(anyHeader.value), true);
    }

    function unmarshalClientState(bytes memory bz) internal view returns (ClientState.Data memory clientState, bool ok) {
        Any.Data memory anyClientState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyClientState.type_url)) != pts.clientState) {
            return (clientState, false);
        }
        return (ClientState.decode(anyClientState.value), true);
    }

    function unmarshalConsensusState(bytes memory bz) internal view returns (ConsensusState.Data memory consensusState, bool ok) {
        Any.Data memory anyConsensusState = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyConsensusState.type_url)) != pts.consensusState) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(anyConsensusState.value), true);
    }

    /// Validity predicate ///

    /**
     * @dev verify verifies untrusted header
     * @param consensusState consensusState corresponding to trusted height
     * @param untrustedHeader untrusted header
     */
    function verify(ConsensusState.Data memory consensusState, ParsedBesuHeader memory untrustedHeader) internal pure returns (bytes[] memory validators, bool ok) {
        bytes32 blkHash = keccak256(untrustedHeader.base.besu_header_rlp);

        if (!verifyCommitSealsTrusting(consensusState.validators, untrustedHeader.base.seals, blkHash, Fraction({numerator: 1, denominator: 3}))) {
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
    function verifyCommitSealsTrusting(bytes[] memory trustedVals, bytes[] memory seals, bytes32 blkHash, Fraction memory trustLevel) internal pure returns (bool) {
        uint8 success = 0;
        bool[] memory marked = new bool[](trustedVals.length);
        for (uint i = 0; i < seals.length; i++) {
            if (seals[i].length == 0) {
                continue;
            }
            address signer = ECRecovery.recover(blkHash, seals[i]);
            for (uint j = 0; j < trustedVals.length; j++) {
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
    function verifyCommitSeals(RLP.RLPItem[] memory untrustedVals, bytes[] memory seals, bytes32 blkHash) internal pure returns (bytes[] memory, bool) {
        bytes[] memory validators = new bytes[](untrustedVals.length);
        uint8 success = 0;
        for (uint i = 0; i < seals.length; i++) {
            validators[i] = untrustedVals[i].toBytes();
            if (seals[i].length == 0) {
                continue;
            } else if (validators[i].toAddress() == ECRecovery.recover(blkHash, seals[i])) {
                success++;
            }
        }
        return (validators, success > untrustedVals.length * 2 / 3);
    }

    /// State verification functions ///

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes
    ) public override returns (bool) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bool found;

        (clientState, found) = getClientState(host, clientId);
        if (!found) {
            return false;
        }
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        (consensusState, found) = getConsensusState(host, clientId, height);
        if (!found) {
            return false;
        }
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.clientStateCommitmentSlot(counterpartyClientIdentifier), keccak256(clientStateBytes));
    }

    function verifyClientConsensusState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        string memory counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes // serialized with pb
    ) public override returns (bool) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bool found;

        (clientState, found) = getClientState(host, clientId);
        if (!found) {
            return false;
        }
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        (consensusState, found) = getConsensusState(host, clientId, height);
        if (!found) {
            return false;
        }
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.consensusStateCommitmentSlot(counterpartyClientIdentifier, consensusHeight), keccak256(consensusStateBytes));
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public override returns (bool) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bool found;

        (clientState, found) = getClientState(host, clientId);
        if (!found) {
            return false;
        }
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        (consensusState, found) = getConsensusState(host, clientId, height);
        if (!found) {
            return false;
        }
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.connectionCommitmentSlot(connectionId), keccak256(connectionBytes));
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes // serialized with pb
    ) public override returns (bool) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bool found;

        (clientState, found) = getClientState(host, clientId);
        if (!found) {
            return false;
        }
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        (consensusState, found) = getConsensusState(host, clientId, height);
        if (!found) {
            return false;
        }
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.channelCommitmentSlot(portId, channelId), keccak256(channelBytes));
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        uint64 height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public override returns (bool) {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        bool found;

        (clientState, found) = getClientState(host, clientId);
        if (!found) {
            return false;
        }
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        if (!validateDelayPeriod(host, clientId, height, delayPeriodTime, delayPeriodBlocks)) {
            return false;
        }
        (consensusState, found) = getConsensusState(host, clientId, height);
        if (!found) {
            return false;
        }
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.packetCommitmentSlot(portId, channelId, sequence), commitmentBytes);
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        uint64 height,
        uint64 delayPeriodTime,
        uint64 delayPeriodBlocks,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes memory acknowledgement
    ) public override returns (bool) {
        ClientState.Data memory clientState = mustGetClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        if (!validateDelayPeriod(host, clientId, height, delayPeriodTime, delayPeriodBlocks)) {
            return false;
        }
        bytes32 stateRoot = mustGetConsensusState(host, clientId, height).root.toBytes32();
        bytes32 ackCommitmentSlot = IBCIdentifier.packetAcknowledgementCommitmentSlot(portId, channelId, sequence);
        bytes32 ackCommitment = host.makePacketAcknowledgementCommitment(acknowledgement);
        return verifyMembership(proof, stateRoot, prefix, ackCommitmentSlot, ackCommitment);
    }

    /// helper functions ///

    function getClientState(IBCHost host, string memory clientId) public view returns (ClientState.Data memory clientState, bool found) {
        bytes memory clientStateBytes;
        (clientStateBytes, found) = host.getClientState(clientId);
        if (!found) {
            return (clientState, false);
        }
        return (ClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getConsensusState(IBCHost host, string memory clientId, uint64 height) public view returns (ConsensusState.Data memory consensusState, bool found) {
        bytes memory consensusStateBytes;
        (consensusStateBytes, found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (consensusState, false);
        }
        return (ConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    function validateArgs(ClientState.Data memory cs, uint64 height, bytes memory prefix, bytes memory proof) internal pure returns (bool) {
        if (cs.latest_height < height) {
            return false;
        } else if (prefix.length == 0) {
            return false;
        } else if (proof.length == 0) {
            return false;
        }
        return true;
    }

    function validateDelayPeriod(IBCHost host, string memory clientId, uint64 height, uint64 delayPeriodTime, uint64 delayPeriodBlocks) private view returns (bool) {
        uint64 currentTime = uint64(block.timestamp * 1000 * 1000 * 1000);
        uint64 validTime = mustGetProcessedTime(host, clientId, height) + delayPeriodTime;
        if (currentTime < validTime) {
            return false;
        }
        uint64 currentHeight = uint64(block.number);
        uint64 validHeight = mustGetProcessedHeight(host, clientId, height) + delayPeriodBlocks;
        if (currentHeight < validHeight) {
            return false;
        }
        return true;
    }

    // NOTE: this is a workaround to avoid the error `Stack too deep` in caller side
    function mustGetClientState(IBCHost host, string memory clientId) internal view returns (ClientState.Data memory) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        require(found, "client state not found");
        return clientState;
    }

    // NOTE: this is a workaround to avoid the error `Stack too deep` in caller side
    function mustGetConsensusState(IBCHost host, string memory clientId, uint64 height) internal view returns (ConsensusState.Data memory) {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, height);
        require(found, "consensus state not found");
        return consensusState;
    }

    function mustGetProcessedTime(IBCHost host, string memory clientId, uint64 height) internal view returns (uint64) {
        (uint256 processedTime, bool found) = host.getProcessedTime(clientId, height);
        require(found, "processed time not found");
        return uint64(processedTime) * 1000 * 1000 * 1000;
    }

    function mustGetProcessedHeight(IBCHost host, string memory clientId, uint64 height) internal view returns (uint64) {
        (uint256 processedHeight, bool found) = host.getProcessedHeight(clientId, height);
        require(found, "processed height not found");
        return uint64(processedHeight);
    }

    function verifyMembership(
        bytes memory proof,
        bytes32 root,
        bytes memory prefix,
        bytes32 slot,
        bytes32 expectedValue
    ) internal pure returns (bool) {
        bytes32 path = keccak256(abi.encodePacked(slot));
        bytes memory dataHash = proof.verify(root, path); // reverts if proof is invalid
        return expectedValue == dataHash.toRLPItem().toBytes().toBytes32();
    }

    function parseBesuHeader(Header.Data memory header) internal pure returns (ParsedBesuHeader memory) {
        ParsedBesuHeader memory parsedHeader;

        parsedHeader.base = header;
        RLP.RLPItem[] memory items = header.besu_header_rlp.toRLPItem().toList();
        parsedHeader.stateRoot = items[3].toBytes().toBytes32();
        parsedHeader.height = uint64(items[8].toUint());

        require(items.length == 15, "items length must be 15");
        parsedHeader.time = uint64(items[11].toUint());
        items = items[12].toBytes().toRLPItem().toList();
        require(items.length == 4, "extra length must be 4");

        parsedHeader.validators = items[1].toList();
        return parsedHeader;
    }

    function verifyStorageProof(address account, bytes32 stateRoot, bytes memory accountStateProof) internal pure returns (bytes32) {
        bytes32 proofPath = keccak256(abi.encodePacked(account));
        bytes memory accountRLP = accountStateProof.verify(stateRoot, proofPath); // reverts if proof is invalid
        return bytes32(accountRLP.toRLPItem().toList()[ACCOUNT_STORAGE_ROOT_INDEX].toUint());
    }
}
