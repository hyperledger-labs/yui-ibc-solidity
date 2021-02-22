pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./types/Channel.sol";
import "./IBCMsgs.sol";

interface IHandler {
    function createClient(IBCMsgs.MsgCreateClient calldata msg_) external;
    function updateClient(IBCMsgs.MsgUpdateClient calldata msg_) external;

    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_) external returns (string memory);
    function connectionOpenTry(
        IBCMsgs.MsgConnectionOpenTry calldata msg_
    ) external returns (string memory);
    function connectionOpenAck(
        IBCMsgs.MsgConnectionOpenAck calldata msg_
    ) external;
    function connectionOpenConfirm(
        IBCMsgs.MsgConnectionOpenConfirm calldata msg_
    ) external;

    function channelOpenInit(
        IBCMsgs.MsgChannelOpenInit calldata msg_
    ) external returns (string memory);
    function channelOpenTry(
        IBCMsgs.MsgChannelOpenTry calldata msg_
    ) external returns (string memory);
    function channelOpenAck(
        IBCMsgs.MsgChannelOpenAck calldata msg_
    ) external;
    function channelOpenConfirm(
        IBCMsgs.MsgChannelOpenConfirm calldata msg_
    ) external;
    
    function sendPacket(Packet.Data calldata packet) external;
    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external returns (bytes memory);
    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external;
}
