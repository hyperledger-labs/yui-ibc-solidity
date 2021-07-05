pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Connection.sol";
import "./types/Channel.sol";

/*
IBCMsgs defines Datagrams in ics-026.
*/
library IBCMsgs {

    /// Client ///

    struct MsgCreateClient {
        string clientType;
        uint64 height;
        bytes clientStateBytes;
        bytes consensusStateBytes;
    }

    struct MsgUpdateClient {
        string clientId;
        bytes header;
    }

    /// Connection handshake ///

    struct MsgConnectionOpenInit {
        string clientId;
        Counterparty.Data counterparty;
        uint64 delayPeriod;
    }

    struct MsgConnectionOpenTry {
        string previousConnectionId;
        Counterparty.Data counterparty; // counterpartyConnectionIdentifier, counterpartyPrefix and counterpartyClientIdentifier
        uint64 delayPeriod;
        string clientId; // clientID of chainA
        bytes clientStateBytes; // clientState that chainA has for chainB
        Version.Data[] counterpartyVersions; // supported versions of chain A
        bytes proofInit; // proof that chainA stored connectionEnd in state (on ConnOpenInit)
        bytes proofClient; // proof that chainA stored a light client of chainB
        bytes proofConsensus; // proof that chainA stored chainB's consensus state at consensus height
        uint64 proofHeight; // height at which relayer constructs proof of A storing connectionEnd in state
        uint64 consensusHeight; // latest height of chain B which chain A has stored in its chain B client
    }

    struct MsgConnectionOpenAck {
        string connectionId;
        bytes clientStateBytes; // client state for chainA on chainB
        Version.Data version; // version that ChainB chose in ConnOpenTry
        string counterpartyConnectionID;
        bytes proofTry; // proof that connectionEnd was added to ChainB state in ConnOpenTry
        bytes proofClient; // proof of client state on chainB for chainA
        bytes proofConsensus; // proof that chainB has stored ConsensusState of chainA on its client
        uint64 proofHeight; // height that relayer constructed proofTry
        uint64 consensusHeight; // latest height of chainA that chainB has stored on its chainA client
    }

    struct MsgConnectionOpenConfirm {
        string connectionId;
        bytes proofAck;
        uint64 proofHeight;
    }

    /// Channel handshake ///

    struct MsgChannelOpenInit {
        string portId;
        Channel.Data channel;
    }

    struct MsgChannelOpenTry {
        string portId;
        string previousChannelId;
        Channel.Data channel;
        string counterpartyVersion;
        bytes proofInit;
        uint64 proofHeight;
    }

    struct MsgChannelOpenAck {
        string portId;
        string channelId;
        string counterpartyVersion;
        string counterpartyChannelId;
        bytes proofTry;
        uint64 proofHeight;
    }

    struct MsgChannelOpenConfirm {
        string portId;
        string channelId;
        bytes proofAck;
        uint64 proofHeight;
    }

    /// Packet ///

    struct MsgPacketRecv {
        Packet.Data packet;
        bytes proof;
        uint64 proofHeight;
    }

    struct MsgPacketAcknowledgement {
        Packet.Data packet;
        bytes acknowledgement;
        bytes proof;
        uint64 proofHeight;
    }
}
