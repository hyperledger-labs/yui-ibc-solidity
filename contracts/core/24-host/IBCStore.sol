// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel} from "../../proto/Channel.sol";

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
}
