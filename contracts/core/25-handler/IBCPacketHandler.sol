// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Context.sol";
import "../25-handler/IBCMsgs.sol";
import "../24-host/IBCHost.sol";
import "../04-channel/IIBCChannel.sol";
import "../05-port/ModuleManager.sol";
import "../05-port/IIBCModule.sol";

/**
 * @dev IBCPacketHandler is a contract that calls a contract that implements `IIBCPacket` with delegatecall.
 */
abstract contract IBCPacketHandler is Context, ModuleManager {
    using Address for address;

    // IBC Packet contract address
    address immutable ibcPacket;

    // Events
    event SendPacket(
        uint64 sequence,
        string sourcePort,
        string sourceChannel,
        Height.Data timeoutHeight,
        uint64 timeoutTimestamp,
        bytes data
    );
    event RecvPacket(Packet.Data packet);
    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );
    event AcknowledgePacket(Packet.Data packet, bytes acknowledgement);
    event TimeoutPacket(Packet.Data packet);

    constructor(address _ibcPacket) {
        require(Address.isContract(_ibcPacket));
        ibcPacket = _ibcPacket;
    }

    function sendPacket(
        string calldata sourcePort,
        string calldata sourceChannel,
        Height.Data calldata timeoutHeight,
        uint64 timeoutTimestamp,
        bytes calldata data
    ) external returns (uint64) {
        require(authenticateCapability(channelCapabilityPath(sourcePort, sourceChannel)));
        uint64 sequence = abi.decode(
            ibcPacket.functionDelegateCall(
                abi.encodeWithSelector(
                    IIBCPacket.sendPacket.selector, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data
                )
            ),
            (uint64)
        );
        emit SendPacket(sequence, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data);
        return sequence;
    }

    function recvPacket(IBCMsgs.MsgPacketRecv calldata msg_) external {
        IIBCModule module = lookupModuleByChannel(msg_.packet.destination_port, msg_.packet.destination_channel);
        ibcPacket.functionDelegateCall(abi.encodeWithSelector(IIBCPacket.recvPacket.selector, msg_));
        bytes memory acknowledgement = module.onRecvPacket(msg_.packet, _msgSender());
        if (acknowledgement.length > 0) {
            ibcPacket.functionDelegateCall(
                abi.encodeWithSelector(
                    IIBCPacket.writeAcknowledgement.selector,
                    msg_.packet.destination_port,
                    msg_.packet.destination_channel,
                    msg_.packet.sequence,
                    acknowledgement
                )
            );
            emit WriteAcknowledgement(
                msg_.packet.destination_port, msg_.packet.destination_channel, msg_.packet.sequence, acknowledgement
            );
        }
        emit RecvPacket(msg_.packet);
    }

    function writeAcknowledgement(
        string calldata destinationPortId,
        string calldata destinationChannel,
        uint64 sequence,
        bytes calldata acknowledgement
    ) external {
        require(authenticateCapability(channelCapabilityPath(destinationPortId, destinationChannel)));
        ibcPacket.functionDelegateCall(
            abi.encodeWithSelector(
                IIBCPacket.writeAcknowledgement.selector,
                destinationPortId,
                destinationChannel,
                sequence,
                acknowledgement
            )
        );
        emit WriteAcknowledgement(destinationPortId, destinationChannel, sequence, acknowledgement);
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement calldata msg_) external {
        IIBCModule module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        ibcPacket.functionDelegateCall(abi.encodeWithSelector(IIBCPacket.acknowledgePacket.selector, msg_));
        module.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement, _msgSender());
        emit AcknowledgePacket(msg_.packet, msg_.acknowledgement);
    }

    function timeoutPacket(IBCMsgs.MsgTimeoutPacket calldata msg_) external {
        IIBCModule module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        ibcPacket.functionDelegateCall(abi.encodeWithSelector(IIBCPacket.timeoutPacket.selector, msg_));
        module.onTimeoutPacket(msg_.packet, _msgSender());
        emit TimeoutPacket(msg_.packet);
    }

    function timeoutOnClose(IBCMsgs.MsgTimeoutOnClose calldata msg_) external {
        IIBCModule module = lookupModuleByChannel(msg_.packet.source_port, msg_.packet.source_channel);
        ibcPacket.functionDelegateCall(abi.encodeWithSelector(IIBCPacket.timeoutOnClose.selector, msg_));
        module.onTimeoutPacket(msg_.packet, _msgSender());
        emit TimeoutPacket(msg_.packet);
    }
}
