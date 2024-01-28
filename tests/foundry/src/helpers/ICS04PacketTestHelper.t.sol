// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Vm} from "forge-std/Test.sol";
import {Height} from "../../../../contracts/proto/Client.sol";
import {Channel} from "../../../../contracts/proto/Channel.sol";
import {
    Packet,
    IIBCChannelRecvPacket,
    IIBCChannelAcknowledgePacket,
    IIBCChannelPacketTimeout
} from "../../../../contracts/core/04-channel/IIBCChannel.sol";
import {IBCChannelLib} from "../../../../contracts/core/04-channel/IBCChannelLib.sol";
import {IIBCHandler} from "../../../../contracts/core/25-handler/IIBCHandler.sol";
import {IBCCommitment} from "../../../../contracts/core/24-host/IBCCommitment.sol";
import {ICS04HandshakeTestHelper} from "./ICS04HandshakeTestHelper.t.sol";
import {MockClientTestHelper} from "./MockClientTestHelper.t.sol";
import {IBCCommitmentTestHelper} from "./IBCCommitmentTestHelper.sol";

abstract contract ICS04PacketTestHelper is ICS04HandshakeTestHelper {
    function validateRecvPacketPostState(IIBCHandler h, ChannelInfo memory dst, uint64 nextSequenceRecv) internal {
        (Channel.Data memory channel, bool ok) = h.getChannel(dst.portId, dst.channelId);
        assertEq(ok, true);
        if (channel.ordering == Channel.Order.ORDER_ORDERED) {
            assertEq(h.getNextSequenceRecv(dst.portId, dst.channelId), nextSequenceRecv);
            assertEq(
                h.getCommitment(IBCCommitment.nextSequenceRecvCommitmentKey(dst.portId, dst.channelId)),
                keccak256(abi.encodePacked(bytes8(nextSequenceRecv)))
            );
        } else if (channel.ordering == Channel.Order.ORDER_UNORDERED) {
            assertTrue(
                h.getPacketReceipt(dst.portId, dst.channelId, nextSequenceRecv - 1)
                    == IBCChannelLib.PacketReceipt.SUCCESSFUL
            );
        } else {
            revert("invalid channel order");
        }
    }

    function createPacket(
        ChannelInfo memory src,
        ChannelInfo memory dst,
        uint64 sequence,
        bytes memory data,
        Height.Data memory timeoutHeight,
        uint64 timeoutTimestamp
    ) internal pure returns (Packet memory) {
        return Packet({
            sequence: sequence,
            sourcePort: src.portId,
            sourceChannel: src.channelId,
            destinationPort: dst.portId,
            destinationChannel: dst.channelId,
            data: data,
            timeoutHeight: timeoutHeight,
            timeoutTimestamp: timeoutTimestamp
        });
    }

    function msgPacketRecv(Packet memory packet, Height.Data memory proofHeight)
        internal
        pure
        returns (IIBCChannelRecvPacket.MsgPacketRecv memory)
    {
        return IIBCChannelRecvPacket.MsgPacketRecv({
            packet: packet,
            proof: provePacketCommitment(packet, proofHeight),
            proofHeight: proofHeight
        });
    }

    function msgPacketAcknowledgement(
        Packet memory packet,
        bytes memory acknowledgement,
        Height.Data memory proofHeight
    ) internal pure returns (IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement memory) {
        return IIBCChannelAcknowledgePacket.MsgPacketAcknowledgement({
            packet: packet,
            acknowledgement: acknowledgement,
            proof: proveAcknowledgementCommitment(
                packet.destinationPort, packet.destinationChannel, packet.sequence, acknowledgement, proofHeight
                ),
            proofHeight: proofHeight
        });
    }

    function msgTimeoutPacket(Channel.Order ordering, Packet memory packet, Height.Data memory proofHeight)
        internal
        pure
        returns (IIBCChannelPacketTimeout.MsgTimeoutPacket memory)
    {
        if (ordering == Channel.Order.ORDER_ORDERED) {
            return IIBCChannelPacketTimeout.MsgTimeoutPacket({
                packet: packet,
                proof: proveNextSequenceRecv(
                    packet.destinationPort, packet.destinationChannel, packet.sequence, proofHeight
                    ),
                proofHeight: proofHeight,
                nextSequenceRecv: packet.sequence
            });
        } else if (ordering == Channel.Order.ORDER_UNORDERED) {
            return IIBCChannelPacketTimeout.MsgTimeoutPacket({
                packet: packet,
                proof: provePacketReceiptAbsence(
                    packet.destinationPort, packet.destinationChannel, packet.sequence, proofHeight
                    ),
                proofHeight: proofHeight,
                nextSequenceRecv: 0
            });
        } else {
            revert("unknown ordering");
        }
    }

    function msgTimeoutOnClose(
        IIBCHandler cpH,
        Channel.Order ordering,
        Packet memory packet,
        Height.Data memory proofHeight
    ) internal view returns (IIBCChannelPacketTimeout.MsgTimeoutOnClose memory) {
        (Channel.Data memory channel, bool ok) = cpH.getChannel(packet.destinationPort, packet.destinationChannel);
        require(ok, "channel not found");
        if (ordering == Channel.Order.ORDER_ORDERED) {
            return IIBCChannelPacketTimeout.MsgTimeoutOnClose({
                packet: packet,
                proofUnreceived: proveNextSequenceRecv(
                    packet.destinationPort, packet.destinationChannel, packet.sequence, proofHeight
                    ),
                proofClose: proveChannelState(proofHeight, packet.destinationPort, packet.destinationChannel, channel),
                proofHeight: proofHeight,
                nextSequenceRecv: packet.sequence,
                counterpartyUpgradeSequence: 0
            });
        } else if (ordering == Channel.Order.ORDER_UNORDERED) {
            return IIBCChannelPacketTimeout.MsgTimeoutOnClose({
                packet: packet,
                proofUnreceived: provePacketReceiptAbsence(
                    packet.destinationPort, packet.destinationChannel, packet.sequence, proofHeight
                    ),
                proofClose: proveChannelState(proofHeight, packet.destinationPort, packet.destinationChannel, channel),
                proofHeight: proofHeight,
                nextSequenceRecv: 0,
                counterpartyUpgradeSequence: 0
            });
        } else {
            revert("unknown ordering");
        }
    }

    function provePacketCommitment(Packet memory packet, Height.Data memory proofHeight)
        internal
        pure
        virtual
        returns (bytes memory);

    function proveAcknowledgementCommitment(
        string memory destPort,
        string memory destChannel,
        uint64 sequence,
        bytes memory acknowledgement,
        Height.Data memory proofHeight
    ) internal pure virtual returns (bytes memory);

    function provePacketReceiptAbsence(
        string memory destPort,
        string memory destChannel,
        uint64 sequence,
        Height.Data memory proofHeight
    ) internal pure virtual returns (bytes memory);

    function proveNextSequenceRecv(
        string memory portId,
        string memory channelId,
        uint64 sequence,
        Height.Data memory proofHeight
    ) internal pure virtual returns (bytes memory);
}

abstract contract ICS04PacketMockClientTestHelper is ICS04PacketTestHelper, MockClientTestHelper {
    function provePacketCommitment(Packet memory packet, Height.Data memory proofHeight)
        internal
        pure
        virtual
        override
        returns (bytes memory)
    {
        return genMockProof(
            proofHeight,
            DEFAULT_COMMITMENT_PREFIX,
            IBCCommitmentTestHelper.packetCommitmentPath(packet.sourcePort, packet.sourceChannel, packet.sequence),
            abi.encodePacked(
                sha256(
                    abi.encodePacked(
                        packet.timeoutTimestamp,
                        packet.timeoutHeight.revision_number,
                        packet.timeoutHeight.revision_height,
                        sha256(packet.data)
                    )
                )
            )
        );
    }

    function proveAcknowledgementCommitment(
        string memory destPort,
        string memory destChannel,
        uint64 sequence,
        bytes memory acknowledgement,
        Height.Data memory proofHeight
    ) internal pure virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight,
            DEFAULT_COMMITMENT_PREFIX,
            IBCCommitmentTestHelper.packetAcknowledgementCommitmentPath(destPort, destChannel, sequence),
            abi.encodePacked(sha256(acknowledgement))
        );
    }

    function provePacketReceiptAbsence(
        string memory destPort,
        string memory destChannel,
        uint64 sequence,
        Height.Data memory proofHeight
    ) internal pure virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight,
            DEFAULT_COMMITMENT_PREFIX,
            IBCCommitmentTestHelper.packetReceiptCommitmentPath(destPort, destChannel, sequence),
            bytes("")
        );
    }

    function proveNextSequenceRecv(
        string memory portId,
        string memory channelId,
        uint64 sequence,
        Height.Data memory proofHeight
    ) internal pure virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight,
            DEFAULT_COMMITMENT_PREFIX,
            IBCCommitmentTestHelper.nextSequenceRecvCommitmentPath(portId, channelId),
            abi.encodePacked(bytes8(sequence))
        );
    }
}

abstract contract ICS04PacketEventTestHelper {
    struct WriteAcknolwedgement {
        string destinationPortId;
        string destinationChannel;
        uint64 sequence;
        bytes acknowledgement;
    }

    event SendPacket(
        uint64 sequence,
        string sourcePort,
        string sourceChannel,
        Height.Data timeoutHeight,
        uint64 timeoutTimestamp,
        bytes data
    );

    event WriteAcknowledgement(
        string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement
    );

    event RecvPacket(Packet packet);

    function getLastSentPacket(
        IIBCHandler handler,
        string memory sourcePort,
        string memory sourceChannel,
        Vm.Log[] memory logs
    ) internal view returns (Packet memory) {
        for (uint256 i = logs.length; i > 0; i--) {
            if (logs[i - 1].emitter == address(handler)) {
                (Packet memory p, bool ok) = tryDecodeSendPacketEvent(logs[i - 1]);
                if (ok) {
                    if (
                        keccak256(abi.encodePacked(p.sourcePort)) != keccak256(abi.encodePacked(sourcePort))
                            || keccak256(abi.encodePacked(p.sourceChannel)) != keccak256(abi.encodePacked(sourceChannel))
                    ) {
                        continue;
                    }
                    Channel.Data memory c;
                    (c, ok) = handler.getChannel(p.sourcePort, p.sourceChannel);
                    require(ok, "channel not found");
                    p.destinationPort = c.counterparty.port_id;
                    p.destinationChannel = c.counterparty.channel_id;
                    return p;
                }
            }
        }
        revert("no packet sent");
    }

    function getLastWrittenAcknowledgement(IIBCHandler handler, Vm.Log[] memory logs)
        internal
        pure
        returns (WriteAcknolwedgement memory)
    {
        for (uint256 i = logs.length; i > 0; i--) {
            if (logs[i - 1].emitter == address(handler)) {
                (WriteAcknolwedgement memory wa, bool ok) = tryDecodeWriteAcknowledgementEvent(logs[i - 1]);
                if (ok) {
                    return wa;
                }
            }
        }
        revert("no acknowledgement written");
    }

    function findWrittenAcknowledgement(IIBCHandler handler, Vm.Log[] memory logs)
        internal
        pure
        returns (WriteAcknolwedgement[] memory)
    {
        WriteAcknolwedgement[] memory acks = new WriteAcknolwedgement[](logs.length);
        uint256 count = 0;
        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].emitter == address(handler)) {
                (WriteAcknolwedgement memory wa, bool ok) = tryDecodeWriteAcknowledgementEvent(logs[i]);
                if (ok) {
                    acks[count] = wa;
                    count++;
                }
            }
        }
        assembly {
            mstore(acks, count)
        }
        return acks;
    }

    function getLastRecvPacket(IIBCHandler handler, Vm.Log[] memory logs) internal pure returns (Packet memory) {
        for (uint256 i = logs.length; i > 0; i--) {
            if (logs[i - 1].emitter == address(handler)) {
                (Packet memory p, bool ok) = tryDecodeRecvPacketEvent(logs[i - 1]);
                if (ok) {
                    return p;
                }
            }
        }
        revert("no packet received");
    }

    function tryDecodeSendPacketEvent(Vm.Log memory log) internal pure returns (Packet memory p, bool) {
        if (log.topics[0] != SendPacket.selector) {
            return (p, false);
        }
        return (decodeSendPacketEvent(log.data), true);
    }

    function decodeSendPacketEvent(bytes memory data) internal pure returns (Packet memory) {
        (
            uint64 sequence,
            string memory sourcePort,
            string memory sourceChannel,
            Height.Data memory timeoutHeight,
            uint64 timeoutTimestamp,
            bytes memory packetData
        ) = abi.decode(data, (uint64, string, string, Height.Data, uint64, bytes));
        return Packet({
            sequence: sequence,
            sourcePort: sourcePort,
            sourceChannel: sourceChannel,
            destinationPort: "",
            destinationChannel: "",
            timeoutHeight: timeoutHeight,
            timeoutTimestamp: timeoutTimestamp,
            data: packetData
        });
    }

    function tryDecodeWriteAcknowledgementEvent(Vm.Log memory log)
        internal
        pure
        returns (WriteAcknolwedgement memory wa, bool)
    {
        if (log.topics[0] != WriteAcknowledgement.selector) {
            return (wa, false);
        }
        return (decodeWriteAcknowledgementEvent(log.data), true);
    }

    function decodeWriteAcknowledgementEvent(bytes memory data)
        internal
        pure
        returns (WriteAcknolwedgement memory wa)
    {
        (wa.destinationPortId, wa.destinationChannel, wa.sequence, wa.acknowledgement) =
            abi.decode(data, (string, string, uint64, bytes));
        return wa;
    }

    function tryDecodeRecvPacketEvent(Vm.Log memory log) internal pure returns (Packet memory p, bool) {
        if (log.topics[0] != RecvPacket.selector) {
            return (p, false);
        }
        return (decodeRecvPacketEvent(log.data), true);
    }

    function decodeRecvPacketEvent(bytes memory data) internal pure returns (Packet memory) {
        return abi.decode(data, (Packet));
    }
}
