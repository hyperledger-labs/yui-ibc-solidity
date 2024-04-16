// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Upgrade} from "../../proto/Channel.sol";

abstract contract IBCStore {
    struct RecvStartSequence {
        uint64 sequence;
        uint64 prevSequence;
    }

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
    mapping(string => address) internal portCapabilities;
    mapping(string => mapping(string => address)) internal channelCapabilities;

    mapping(string => mapping(string => Upgrade.Data)) internal upgrades;
    mapping(string => mapping(string => uint64)) internal latestErrorReceiptSequences;
    mapping(string => mapping(string => RecvStartSequence)) internal recvStartSequences;
    mapping(string => mapping(string => uint64)) internal ackStartSequences;

    // Host parameters
    uint64 internal expectedTimePerBlock;

    // Sequences for identifier
    uint64 internal nextClientSequence;
    uint64 internal nextConnectionSequence;
    uint64 internal nextChannelSequence;
}
