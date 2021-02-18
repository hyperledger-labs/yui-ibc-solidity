pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";

/*
IBCMsgs defines Datagrams in ics-026.
*/
library IBCMsgs {
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
