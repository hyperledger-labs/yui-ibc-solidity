// SPDX-License-Identifier: Apache-2.0
/**
 * This file is copied from solidity-mpt library.
 * https://github.com/ibc-solidity/solidity-mpt/blob/d157b5fd0aafb0b1c23bc3a3eb5f5bc04b3fd0a3/src/MPTProof.sol
 */
pragma solidity ^0.8.0;

import {RLPReader} from "./RLPReader.sol";

library MPTProof {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    /// @dev Verifies a Merkle-Patricia-Trie proof.
    ///      If the proof proves the inclusion of some key-value pair in the
    ///      trie, the value is returned. Otherwise, i.e. if the proof proves
    ///      the exclusion of a key from the trie, an empty byte array is
    ///      returned.
    /// @param rlpProof is the stack of MPT nodes (starting with the root) that
    ///        need to be traversed during verification. It's encoded with RLP.
    /// @param rootHash is the Keccak-256 hash of the root node of the MPT.
    /// @param mptKey is a trie key of the node whose
    ///        inclusion/exclusion we are proving.
    /// @return value whose inclusion is proved or an empty byte array for
    ///         a proof of exclusion
    function verifyRLPProof(bytes memory rlpProof, bytes32 rootHash, bytes32 mptKey)
        internal
        pure
        returns (bytes memory value)
    {
        bytes memory key = new bytes(32);
        assembly {
            mstore(add(key, 0x20), mptKey)
        }
        return verify(rlpProof.toRlpItem().toList(), rootHash, decodeNibbles(key, 0));
    }

    /// @dev Verifies a Merkle-Patricia-Trie proof.
    ///      If the proof proves the inclusion of some key-value pair in the
    ///      trie, the value is returned. Otherwise, i.e. if the proof proves
    ///      the exclusion of a key from the trie, an empty byte array is
    ///      returned.
    /// @param proof is the stack of MPT nodes (starting with the root) that
    ///        need to be traversed during verification.
    /// @param rootHash is the Keccak-256 hash of the root node of the MPT.
    /// @param mptKeyNibbles is the key (consisting of nibbles) of the node whose
    ///        inclusion/exclusion we are proving.
    /// @return value whose inclusion is proved or an empty byte array for
    ///         a proof of exclusion
    function verify(RLPReader.RLPItem[] memory proof, bytes32 rootHash, bytes memory mptKeyNibbles)
        internal
        pure
        returns (bytes memory value)
    {
        uint256 mptKeyOffset = 0;
        bytes32 nodeHashHash;
        RLPReader.RLPItem[] memory node;
        RLPReader.RLPItem memory rlpValue;

        if (proof.length == 0) {
            // Root hash of empty Merkle-Patricia-Trie
            require(rootHash == 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421);
            return new bytes(0);
        }

        // Traverse stack of nodes starting at root.
        for (uint256 i = 0; i < proof.length; i++) {
            // The root node is hashed with Keccak-256 ...
            if (i == 0 && rootHash != proof[i].rlpBytesKeccak256()) {
                revert();
            }
            // ... whereas all other nodes are hashed with the MPT
            // hash function.
            if (i != 0 && nodeHashHash != mptHashHash(proof[i])) {
                revert();
            }
            // We verified that proof[i] has the correct hash, so we
            // may safely decode it.
            node = proof[i].toList();

            if (node.length == 2) {
                // Extension or Leaf node

                bool isLeaf;
                bytes memory nodeKey;
                (isLeaf, nodeKey) = merklePatriciaCompactDecode(node[0].toBytes());

                uint256 prefixLength = sharedPrefixLength(mptKeyOffset, mptKeyNibbles, nodeKey);
                mptKeyOffset += prefixLength;

                if (prefixLength < nodeKey.length) {
                    // Proof claims divergent extension or leaf. (Only
                    // relevant for proofs of exclusion.)
                    // An Extension/Leaf node is divergent iff it "skips" over
                    // the point at which a Branch node should have been had the
                    // excluded key been included in the trie.
                    // Example: Imagine a proof of exclusion for path [1, 4],
                    // where the current node is a Leaf node with
                    // path [1, 3, 3, 7]. For [1, 4] to be included, there
                    // should have been a Branch node at [1] with a child
                    // at 3 and a child at 4.

                    // Sanity check
                    if (i < proof.length - 1) {
                        // divergent node must come last in proof
                        revert();
                    }

                    return new bytes(0);
                }

                if (isLeaf) {
                    // Sanity check
                    if (i < proof.length - 1) {
                        // leaf node must come last in proof
                        revert();
                    }

                    if (mptKeyOffset < mptKeyNibbles.length) {
                        return new bytes(0);
                    }

                    rlpValue = node[1];
                    return rlpValue.toBytes();
                } else {
                    // extension
                    // Sanity check
                    if (i == proof.length - 1) {
                        // shouldn't be at last level
                        revert();
                    }

                    if (!node[1].isList()) {
                        // rlp(child) was at least 32 bytes. node[1] contains
                        // Keccak256(rlp(child)).
                        nodeHashHash = node[1].payloadKeccak256();
                    } else {
                        // rlp(child) was at less than 32 bytes. node[1] contains
                        // rlp(child).
                        nodeHashHash = node[1].rlpBytesKeccak256();
                    }
                }
            } else if (node.length == 17) {
                // Branch node

                if (mptKeyOffset != mptKeyNibbles.length) {
                    // we haven't consumed the entire path, so we need to look at a child
                    uint256 nibble = uint256(uint8(mptKeyNibbles[mptKeyOffset]));
                    mptKeyOffset += 1;
                    if (nibble >= 16) {
                        // each element of the path has to be a nibble
                        revert();
                    }

                    if (isEmptyBytesequence(node[nibble])) {
                        // Sanity
                        if (i != proof.length - 1) {
                            // leaf node should be at last level
                            revert();
                        }

                        return new bytes(0);
                    } else if (!node[nibble].isList()) {
                        nodeHashHash = node[nibble].payloadKeccak256();
                    } else {
                        nodeHashHash = node[nibble].rlpBytesKeccak256();
                    }
                } else {
                    // we have consumed the entire mptKey, so we need to look at what's contained in this node.

                    // Sanity
                    if (i != proof.length - 1) {
                        // should be at last level
                        revert();
                    }

                    return node[16].toBytes();
                }
            } else {
                revert("invalid node length");
            }
        }
        // unreachable here
        revert();
    }

    function isEmptyBytesequence(RLPReader.RLPItem memory item) internal pure returns (bool) {
        if (item.len != 1) {
            return false;
        }
        uint8 b;
        uint256 memPtr = item.memPtr;
        assembly {
            b := byte(0, mload(memPtr))
        }
        return b == 0x80; /* empty byte string */
    }

    function decodeNibbles(bytes memory bz, uint256 offset) internal pure returns (bytes memory nibbles) {
        uint256 length = bz.length * 2;
        require(bz.length > 0 && offset <= length);

        nibbles = new bytes(length - offset);
        uint256 i = offset;
        if (offset & 1 == 1) {
            nibbles[0] = bytes1((uint8(bz[offset / 2]) >> 0) & 0xF);
            i++;
        }
        unchecked {
            for (; i < length - 1; i += 2) {
                nibbles[i - offset] = bytes1((uint8(bz[i / 2]) >> 4) & 0xF);
                nibbles[i - offset + 1] = bytes1((uint8(bz[i / 2]) >> 0) & 0xF);
            }
        }
    }

    function merklePatriciaCompactDecode(bytes memory bz) internal pure returns (bool isLeaf, bytes memory nibbles) {
        require(bz.length > 0);
        uint256 first_nibble = uint8(bz[0]) >> 4 & 0xF;
        uint256 offset = 0;
        if (first_nibble == 0) {
            offset = 2;
            isLeaf = false;
        } else if (first_nibble == 1) {
            offset = 1;
            isLeaf = false;
        } else if (first_nibble == 2) {
            offset = 2;
            isLeaf = true;
        } else if (first_nibble == 3) {
            offset = 1;
            isLeaf = true;
        } else {
            // Not supposed to happen!
            revert();
        }
        return (isLeaf, decodeNibbles(bz, offset));
    }

    function sharedPrefixLength(uint256 xsOffset, bytes memory xs, bytes memory ys) internal pure returns (uint256) {
        uint256 i = 0;
        for (; i + xsOffset < xs.length && i < ys.length; i++) {
            if (xs[i + xsOffset] != ys[i]) {
                return i;
            }
        }
        return i;
    }

    /// @dev Computes the hash of the Merkle-Patricia-Trie hash of the input.
    ///      Merkle-Patricia-Tries use a weird "hash function" that outputs
    ///      *variable-length* hashes: If the input is shorter than 32 bytes,
    ///      the MPT hash is the input. Otherwise, the MPT hash is the
    ///      Keccak-256 hash of the input.
    ///      The easiest way to compare variable-length byte sequences is
    ///      to compare their Keccak-256 hashes.
    /// @param input The byte sequence to be hashed.
    /// @return Keccak-256(MPT-hash(input))
    function mptHashHash(RLPReader.RLPItem memory input) internal pure returns (bytes32) {
        if (input.len < 32) {
            return input.rlpBytesKeccak256();
        } else {
            return keccak256(abi.encodePacked(input.rlpBytesKeccak256()));
        }
    }
}
