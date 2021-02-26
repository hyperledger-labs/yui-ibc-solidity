pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IClient.sol";
import "./IBCHost.sol";
import "./IBCMsgs.sol";
import "./types/IBFT2.sol";
import "../lib/ECRecovery.sol";
import "../lib/Bytes.sol";
import "../lib/TrieProofs.sol";
import "../lib/RLP.sol";
import "../lib/IBCIdentifier.sol";

contract IBFT2Client is IClient {
    using TrieProofs for bytes;
    using RLP for RLP.RLPItem;
    using RLP for bytes;
    using Bytes for bytes;

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
    function getTimestampAtHeight(IBCHost host, string memory clientId, uint64 height) public override view returns (uint64, bool) {
        (bytes memory consensusStateBytes, bool found) = host.getConsensusState(clientId, height);
        if (!found) {
            return (0, false);
        }
        return (ConsensusState.decode(consensusStateBytes).timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
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
        ConsensusState.Data memory consensusState;
        bytes memory consensusStateBytes;
        bytes[] memory validators;
        bool ok;
        ClientState.Data memory clientState = ClientState.decode(clientStateBytes);

        header = Header.decode(headerBytes);
        (consensusStateBytes, ok) = host.getConsensusState(clientId, header.trusted_height);
        require(ok, "consensusState not found");
        consensusState = ConsensusState.decode(consensusStateBytes);

        //// check validity ////
        ParsedBesuHeader memory parsedHeader = parseBesuHeader(header);
        require(parsedHeader.height > header.trusted_height, "header height ≤ consensus state height");
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
        return (ClientState.encode(clientState), ConsensusState.encode(consensusState), parsedHeader.height);
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
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        ConsensusState.Data memory consensusState = getConsensusState(host, clientId, height);
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
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        ConsensusState.Data memory consensusState = getConsensusState(host, clientId, height);
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
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        ConsensusState.Data memory consensusState = getConsensusState(host, clientId, height);
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
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        ConsensusState.Data memory consensusState = getConsensusState(host, clientId, height);
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, IBCIdentifier.channelCommitmentSlot(portId, channelId), keccak256(channelBytes));
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        return verifyMembership(proof, getConsensusState(host, clientId, height).root.toBytes32(), prefix, IBCIdentifier.packetCommitmentSlot(portId, channelId, sequence), commitmentBytes);
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 ackCommitmentBytes
    ) public override view returns (bool) {
        ClientState.Data memory clientState = getClientState(host, clientId);
        if (!validateArgs(clientState, height, prefix, proof)) {
            return false;
        }
        return verifyMembership(proof, getConsensusState(host, clientId, height).root.toBytes32(), prefix, IBCIdentifier.packetAcknowledgementCommitmentSlot(portId, channelId, sequence), ackCommitmentBytes);
    }

    /// helper functions ///

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

    function getClientState(IBCHost host, string memory clientId) public view returns (ClientState.Data memory clientState) {
        (bytes memory clientStateBytes, bool found) = host.getClientState(clientId);
        require(found, "client state not found");
        return ClientState.decode(clientStateBytes);
    }

    function getConsensusState(IBCHost host, string memory clientId, uint64 height) public view returns (ConsensusState.Data memory) {
        (bytes memory consensusStateBytes, bool found) = host.getConsensusState(clientId, height);
        require(found, "clientState not found");
        return ConsensusState.decode(consensusStateBytes);
    }

    function verifyMembership(
        bytes memory proof,
        bytes32 root,
        bytes memory prefix,
        bytes32 slot,
        bytes32 expectedValue
    ) internal pure returns (bool) {
        uint256 slotNum = toUint256(abi.encodePacked(slot), 0);
        bytes32 path = keccak256(abi.encodePacked(slotNum));
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

    function toUint256(bytes memory _bytes, uint256 _start) internal pure returns (uint256) {
        require(_start + 32 >= _start, "toUint256_overflow");
        require(_bytes.length >= _start + 32, "toUint256_outOfBounds");
        uint256 tempUint;

        assembly {
            tempUint := mload(add(add(_bytes, 0x20), _start))
        }

        return tempUint;
    }
}
