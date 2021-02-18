pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Client.sol";
import "./types/Connection.sol";
import "./types/Channel.sol";

/*
IBCMsgs defines Datagrams in ics-026.
*/
library IBCMsgs {

    /// Client ///

    struct MsgCreateClient {
        string clientId;
        ClientState.Data clientState;
        ConsensusState.Data consensusState;
    }

    struct MsgUpdateClient {
        string clientId;
        Header header;
    }

    struct Header {
        bytes besuHeaderRLPBytes;
        bytes[] seals;
        uint64 trustedHeight;
        bytes accountStateProof;
    }

    /// Connection handshake ///

    struct MsgConnectionOpenInit {
        string clientId;
        string connectionId;
        Counterparty.Data counterparty;
        uint64 delayPeriod;
    }

    struct MsgConnectionOpenTry {
        string connectionId;
        Counterparty.Data counterparty; // counterpartyConnectionIdentifier, counterpartyPrefix and counterpartyClientIdentifier
        uint64 delayPeriod;
        string clientId; // clientID of chainA
        ClientState.Data clientState; // clientState that chainA has for chainB
        Version.Data[] counterpartyVersions; // supported versions of chain A
        bytes proofInit; // proof that chainA stored connectionEnd in state (on ConnOpenInit)
        bytes proofClient; // proof that chainA stored a light client of chainB
        bytes proofConsensus; // proof that chainA stored chainB's consensus state at consensus height
        uint64 proofHeight; // height at which relayer constructs proof of A storing connectionEnd in state
        uint64 consensusHeight; // latest height of chain B which chain A has stored in its chain B client
    }

    struct MsgConnectionOpenAck {
        string connectionId;
        ClientState.Data clientState; // client state for chainA on chainB
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
        string channelId;
        Channel.Data channel;
    }

    struct MsgChannelOpenTry {
        string portId;
        string channelId;
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
