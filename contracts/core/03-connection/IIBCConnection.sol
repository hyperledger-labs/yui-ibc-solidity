// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {Version, Counterparty} from "../../proto/Connection.sol";

interface IIBCConnection {
    // --------------------- Data Structure --------------------- //

    /**
     * @dev MsgConnectionOpenInit defines the msg sent by an account on Chain A to initialize a connection with Chain B.
     * @param clientId of Chain B
     * @param counterparty is composed of clientID of Chain A on Chain B and the prefix of Chain B
     * @param version chosen by an account on Chain A
     * @param delayPeriod is set to the connection delay time period
     */
    struct MsgConnectionOpenInit {
        string clientId;
        Counterparty.Data counterparty;
        Version.Data version;
        uint64 delayPeriod;
    }

    /**
     * @dev MsgConnectionOpenTry defines a msg sent by an account on Chain B to try to open a connection on Chain B.
     * @param counterparty is composed of clientID of Chain A on Chain B and the prefix of Chain B
     * @param delayPeriod chosen by Chain A in `connectionOpenInit`
     * @param clientId of Chain A
     * @param clientStateBytes clientState that Chain A has for Chain B
     * @param counterpartyVersions supported versions of Chain A
     * @param proofInit proof that Chain A stored connectionEnd in state (on connectionOpenInit)
     * @param proofClient proof that Chain A stored a light client of Chain B
     * @param proofConsensus proof that Chain A stored Chain B's consensus state at consensus height
     * @param proofHeight height at which relayer constructs proof of A storing connectionEnd in state
     * @param consensusHeight latest height of Chain B which Chain A has stored in its Chain B client
     * @param hostConsensusStateProof proof data for the consensus state of Chain B
     */
    struct MsgConnectionOpenTry {
        Counterparty.Data counterparty;
        uint64 delayPeriod;
        string clientId;
        bytes clientStateBytes;
        Version.Data[] counterpartyVersions;
        bytes proofInit;
        bytes proofClient;
        bytes proofConsensus;
        Height.Data proofHeight;
        Height.Data consensusHeight;
        bytes hostConsensusStateProof;
    }

    /**
     * @dev MsgConnectionOpenAck defines a msg sent by an account on Chain A to acknowledge the change of connection state to TRYOPEN on Chain B.
     * @param connectionId identifier of the connection generated at `connectionOpenInit`
     * @param clientStateBytes clientState that Chain B has for Chain A
     * @param version chosen by Chain B in `connectionOpenTry`
     * @param counterpartyConnectionId identifier of the connection on Chain B
     * @param proofTry proof that Chain B stored connectionEnd in state
     * @param proofClient proof that Chain B stored a light client of Chain A
     * @param proofConsensus proof that Chain B stored Chain A's consensus state at consensus height
     * @param proofHeight height at which relayer constructed proof of B storing connectionEnd in state
     * @param consensusHeight latest height of Chain A which Chain B has stored in its Chain A client
     * @param hostConsensusStateProof proof data for the consensus state of Chain A
     */
    struct MsgConnectionOpenAck {
        string connectionId;
        bytes clientStateBytes;
        Version.Data version;
        string counterpartyConnectionId;
        bytes proofTry;
        bytes proofClient;
        bytes proofConsensus;
        Height.Data proofHeight;
        Height.Data consensusHeight;
        bytes hostConsensusStateProof;
    }

    /**
     * @dev MsgConnectionOpenConfirm defines a msg sent by an account on Chain B to acknowledge the change of connection state to OPEN on Chain A.
     * @param connectionId identifier of the connection generated at `connectionOpenTry`
     * @param proofAck proof that Chain A stored connectionEnd in state
     * @param proofHeight height at which relayer constructed proof of A storing connectionEnd in state
     */
    struct MsgConnectionOpenConfirm {
        string connectionId;
        bytes proofAck;
        Height.Data proofHeight;
    }

    // --------------------- Events --------------------- //

    /// @notice Emitted when a connection identifier is generated
    /// @param connectionId connection identifier
    event GeneratedConnectionIdentifier(string connectionId);

    /**
     * @dev connectionOpenInit initialises a connection attempt on Chain A. The generated connection identifier
     * is returned.
     */
    function connectionOpenInit(MsgConnectionOpenInit calldata msg_) external returns (string memory connectionId);

    /**
     * @dev connectionOpenTry relays notice of a connection attempt on Chain A to Chain B (this
     * code is executed on Chain B).
     */
    function connectionOpenTry(MsgConnectionOpenTry calldata msg_) external returns (string memory);

    /**
     * @dev connectionOpenAck relays acceptance of a connection open attempt from Chain B back
     * to Chain A (this code is executed on Chain A).
     */
    function connectionOpenAck(MsgConnectionOpenAck calldata msg_) external;

    /**
     * @dev connectionOpenConfirm confirms opening of a connection on Chain A to Chain B, after
     * which the connection is open on both chains (this code is executed on Chain B).
     */
    function connectionOpenConfirm(MsgConnectionOpenConfirm calldata msg_) external;
}
