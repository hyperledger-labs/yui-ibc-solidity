// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Upgrade} from "../../proto/Channel.sol";

abstract contract IBCStore {
    // keccak256(abi.encode(uint256(keccak256("ibc.commitment")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant COMMITMENT_STORAGE_LOCATION =
        0x1ee222554989dda120e26ecacf756fe1235cd8d726706b57517715dde4f0c900;

    // keccak256(abi.encode(uint256(keccak256("ibc.host")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant HOST_STORAGE_LOCATION = 0x74277c96171a830beeb656543654929b9b37cec88976b4c31924799951550500;

    // keccak256(abi.encode(uint256(keccak256("ibc.client")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant CLIENT_STORAGE_LOCATION =
        0x521e6acb905d37b69880078e1a941104ad5d8bcb8c5cf52f1d5f47d31739d500;

    // keccak256(abi.encode(uint256(keccak256("ibc.connection")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant CONNECTION_STORAGE_LOCATION =
        0x9ef02a9acd7179d999aa130fa65a34ac06dd2f1bae667ae0fb55000408793800;

    // keccak256(abi.encode(uint256(keccak256("ibc.channel")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant CHANNEL_STORAGE_LOCATION =
        0x1952ed347963c5b7b1856335782fc9c26716d4219254baf3dfc6b26981b2dc00;

    /// @custom:storage-location erc7201:ibc.commitment
    struct CommitmentStorage {
        mapping(bytes32 => bytes32) commitments;
    }

    /// @custom:storage-location erc7201:ibc.host
    struct HostStorage {
        mapping(string => address) clientRegistry;
        mapping(string => address) portCapabilities;
        mapping(string => mapping(string => address)) channelCapabilities;
        uint64 nextClientSequence;
        uint64 nextConnectionSequence;
        uint64 nextChannelSequence;
        uint64 expectedTimePerBlock;
    }

    /// @custom:storage-location erc7201:ibc.client
    struct ClientStorage {
        string clientType;
        address clientImpl;
    }

    /// @custom:storage-location erc7201:ibc.connection
    struct ConnectionStorage {
        ConnectionEnd.Data connection;
    }

    struct RecvStartSequence {
        uint64 sequence;
        uint64 prevSequence;
    }

    /// @custom:storage-location erc7201:ibc.channel
    struct ChannelStorage {
        Channel.Data channel;
        uint64 nextSequenceSend;
        uint64 nextSequenceRecv;
        uint64 nextSequenceAck;
        Upgrade.Data upgrade;
        uint64 latestErrorReceiptSequence;
        RecvStartSequence recvStartSequence;
        uint64 ackStartSequence;
    }

    /**
     * @dev getCommitments returns the commitment storage
     */
    function getCommitments() internal pure returns (mapping(bytes32 => bytes32) storage $) {
        assembly {
            // this is safe because first field of CommitmentStorage is a commitments mapping
            $.slot := COMMITMENT_STORAGE_LOCATION
        }
    }

    /**
     * @dev getHostStorage returns the host storage
     */
    function getHostStorage() internal pure returns (HostStorage storage $) {
        assembly {
            $.slot := HOST_STORAGE_LOCATION
        }
    }

    /**
     * @dev getClientStorage returns the client storage
     */
    function getClientStorage() internal pure returns (mapping(string => ClientStorage) storage $) {
        assembly {
            $.slot := CLIENT_STORAGE_LOCATION
        }
    }

    /**
     * @dev getConnectionStorage returns the connection storage
     */
    function getConnectionStorage() internal pure returns (mapping(string => ConnectionStorage) storage $) {
        assembly {
            $.slot := CONNECTION_STORAGE_LOCATION
        }
    }

    /**
     * @dev getChannelStorage returns the channel storage
     */
    function getChannelStorage()
        internal
        pure
        returns (mapping(string => mapping(string => ChannelStorage)) storage $)
    {
        assembly {
            $.slot := CHANNEL_STORAGE_LOCATION
        }
    }
}
