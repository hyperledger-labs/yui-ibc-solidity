pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCStore.sol";
import "./types/Client.sol";
import "../lib/ECRecovery.sol";
import "../lib/Bytes.sol";
import "../lib/TrieProofs.sol";
import "../lib/RLP.sol";

contract IBCClient {
    using TrieProofs for bytes;
    using RLP for RLP.RLPItem;
    using RLP for bytes;
    using Bytes for bytes;

    uint8 private constant ACCOUNT_STORAGE_ROOT_INDEX = 2;

    IBCStore ibcStore;

    constructor(IBCStore s) public {
        ibcStore = s;
    }

    struct Header {
        bytes besuHeaderRLPBytes;
        bytes[] seals;
        uint64 trustedHeight;
        bytes accountStateProof;
    }

    struct ParsedBesuHeader {
        Header base;
        uint64 height;
        bytes32 stateRoot;
        uint64 time;
        RLP.RLPItem[] validators;
    }

    struct Fraction {
        uint64 numerator;
        uint64 denominator;
    }

    Fraction defaultTrustLevel = Fraction({numerator: 1, denominator: 3});

    /* Public functions */

    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(string memory clientId, ClientState.Data memory clientState, ConsensusState.Data memory consensusState) public {
        require(!ibcStore.hasClientState(clientId));

        ibcStore.setClientState(clientId, clientState);
        ibcStore.setConsensusState(clientId, clientState.latest_height, consensusState);
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(string memory clientId, Header memory header) public {
        ClientState.Data memory clientState;
        ConsensusState.Data memory consensusState;
        uint64 height;
        bool found;
    
        (clientState, found) = ibcStore.getClientState(clientId);
        require(found, "clientState not found");

        (clientState, consensusState, height) = checkHeaderAndUpdateState(clientId, clientState, header);
    
        //// persist states ////
        ibcStore.setClientState(clientId, clientState);
        ibcStore.setConsensusState(clientId, height, consensusState);
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(string memory clientId, uint64 height) public view returns (uint64, bool) {
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    /* Internal functions */

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(string memory clientId, ClientState.Data memory clientState, Header memory header) internal view returns (ClientState.Data memory, ConsensusState.Data memory, uint64) {
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, header.trustedHeight);
        require(found, "consensusState not found");

        //// check validity ////
        ParsedBesuHeader memory parsedHeader = parseBesuHeader(header);
        require(parsedHeader.height > parsedHeader.base.trustedHeight, "header height ≤ consensus state height");
        (bytes[] memory validators, bool ok) = verify(consensusState, parsedHeader);
        require(ok, "failed to verify the header");
        bytes32 accountStorageRoot = verifyStorageProof(Bytes.toAddress(clientState.ibc_store_address), parsedHeader.stateRoot, header.accountStateProof);

        //// update ////
        consensusState.timestamp = parsedHeader.time;
        consensusState.root = abi.encodePacked(accountStorageRoot);
        consensusState.validators = validators;

        if (parsedHeader.height > clientState.latest_height) {
            clientState.latest_height = parsedHeader.height;
        }
    
        return (clientState, consensusState, parsedHeader.height);
    }

    /**
     * @dev verify verifies untrusted header
     * @param consensusState consensusState corresponding to trusted height
     * @param untrustedHeader untrusted header
     */
    function verify(ConsensusState.Data memory consensusState, ParsedBesuHeader memory untrustedHeader) internal view returns (bytes[] memory validators, bool ok) {
        bytes32 blkHash = keccak256(untrustedHeader.base.besuHeaderRLPBytes);

        if (!verifyCommitSealsTrusting(consensusState.validators, untrustedHeader.base.seals, blkHash, defaultTrustLevel)) {
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

    function verifyStorageProof(address account, bytes32 stateRoot, bytes memory accountStateProof) internal pure returns (bytes32) {
        bytes32 proofPath = keccak256(abi.encodePacked(account));
        bytes memory accountRLP = accountStateProof.verify(stateRoot, proofPath); // reverts if proof is invalid
        return bytes32(accountRLP.toRLPItem().toList()[ACCOUNT_STORAGE_ROOT_INDEX].toUint());
    }

    function parseBesuHeader(Header memory header) internal pure returns (ParsedBesuHeader memory) {
        ParsedBesuHeader memory parsedHeader;

        parsedHeader.base = header;
        RLP.RLPItem[] memory items = header.besuHeaderRLPBytes.toRLPItem().toList();
        parsedHeader.stateRoot = items[3].toBytes().toBytes32();
        parsedHeader.height = uint64(items[8].toUint());

        require(items.length == 15, "items length must be 15");
        parsedHeader.time = uint64(items[11].toUint());
        items = items[12].toBytes().toRLPItem().toList();
        require(items.length == 4, "extra length must be 4");

        parsedHeader.validators = items[1].toList();
        return parsedHeader;
    }

    // TODO implements
    function validateSelfClient(
        ClientState.Data memory self
    ) public pure returns (bool) {
        return true;
    }

    // TODO implements
    function getSelfConsensusState(
        uint64 height
    ) public pure returns (ConsensusState.Data memory, bool) {
        ConsensusState.Data memory cs;
        return (cs, false);
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

    function validateArgs(ClientState.Data memory cs, uint64 height, bytes memory prefix, bytes memory proof) internal view returns (bool) {
        if (cs.latest_height < height) {
            return false;
        } else if (prefix.length == 0) {
            return false;
        } else if (proof.length == 0) {
            return false;
        }
        return true;
    }

    /* State verification functions */

    function verifyClientState(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        ClientState.Data memory target
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            return false;
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.clientStateCommitmentSlot(counterpartyClientIdentifier), keccak256(ClientState.encode(target)));
    }

    function verifyClientConsensusState(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        string memory counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes // serialized with pb
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            return false;
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.consensusStateCommitmentSlot(counterpartyClientIdentifier, consensusHeight), keccak256(consensusStateBytes));
    }

    function verifyConnectionState(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            return false;
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.connectionCommitmentSlot(connectionId), keccak256(connectionBytes));
    }

    function verifyChannelState(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes // serialized with pb
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            revert("fail");
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.channelCommitmentSlot(portId, channelId), keccak256(channelBytes));
    }

    function verifyPacketCommitment(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            revert("fail");
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.packetCommitmentSlot(portId, channelId, sequence), commitmentBytes);
    }

    function verifyPacketAcknowledgement(
        ClientState.Data memory self,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 ackCommitmentBytes
    ) public view returns (bool) {
        if (!validateArgs(self, height, prefix, proof)) {
            revert("fail");
        }
        (ConsensusState.Data memory consensusState, bool found) = ibcStore.getConsensusState(clientId, height);
        require(found, "consensusState not found");
        return verifyMembership(proof, consensusState.root.toBytes32(), prefix, ibcStore.packetAcknowledgementCommitmentSlot(portId, channelId, sequence), ackCommitmentBytes);
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
