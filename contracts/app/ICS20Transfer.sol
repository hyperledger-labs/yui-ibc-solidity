pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/types/Channel.sol";
import "../core/IBCModule.sol";
import "../core/IBCHandler.sol";
import "../core/IBCHost.sol";
import "../core/types/App.sol";
import "../lib/strings.sol";
import "../lib/Bytes.sol";
import "./IICS20Bank.sol";
import "openzeppelin-solidity/contracts/utils/Context.sol";

contract ICS20Transfer is Context, IModuleCallbacks {
    using strings for *;
    using Bytes for *;

    IBCHandler ibcHandler;
    IBCHost ibcHost;
    IICS20Bank bank;

    mapping(string => address) channelEscrowAddresses;

    constructor(IBCHost host_, IBCHandler ibcHandler_, IICS20Bank bank_) public {
        ibcHost = host_;
        ibcHandler = ibcHandler_;
        bank = bank_;
    }

    function sendTransfer(
        string calldata denom,
        uint64 amount,
        address receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external {
        if (!denom.toSlice().startsWith(makeDenomPrefix(sourcePort, sourceChannel))) { // sender is source chain
            bank.transferFrom(_msgSender(), getEscrowAddress(sourceChannel), denom, amount);
        } else {
            bank.burn(_msgSender(), denom, amount);
        }

        sendPacket(
            FungibleTokenPacketData.Data({
                denom: denom,
                amount: amount,
                sender: abi.encodePacked(_msgSender()),
                receiver: abi.encodePacked(receiver)
            }),
            sourcePort,
            sourceChannel,
            timeoutHeight
        );
    }

    function sendPacket(FungibleTokenPacketData.Data memory data, string memory sourcePort, string memory sourceChannel, uint64 timeoutHeight) internal {
        (Channel.Data memory channel, bool found) = ibcHost.getChannel(sourcePort, sourceChannel);
        require(found, "channel not found");
        ibcHandler.sendPacket(Packet.Data({
            sequence: ibcHost.getNextSequenceSend(sourcePort, sourceChannel),
            source_port: sourcePort,
            source_channel: sourceChannel,
            destination_port: channel.counterparty.port_id,
            destination_channel: channel.counterparty.channel_id,
            data: FungibleTokenPacketData.encode(data),
            timeout_height: Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            timeout_timestamp: 0
        }));
    }

    function getEscrowAddress(string memory sourceChannel) internal view returns (address) {
        address escrow = channelEscrowAddresses[sourceChannel];
        require(escrow != address(0));
        return escrow;
    }

    function newAcknowledgement(bool success) internal pure returns (bytes memory) {
        bytes memory acknowledgement = new bytes(1);
        if (success) {
            acknowledgement[0] = 0x01;
        } else {
            acknowledgement[0] = 0x00;
        }
        return acknowledgement;
    }
    
    function isSuccessAcknowledgement(bytes memory acknowledgement) internal pure returns (bool) {
        require(acknowledgement.length == 1);
        return acknowledgement[0] == 0x01;
    }

    function refundTokens(FungibleTokenPacketData.Data memory data, string memory sourcePort, string memory sourceChannel) internal {
        if (!data.denom.toSlice().startsWith(makeDenomPrefix(sourcePort, sourceChannel))) { // sender was source chain
            bank.transferFrom(getEscrowAddress(sourceChannel), data.sender.toAddress(), data.denom, data.amount);
        } else {
            bank.mint(data.sender.toAddress(), data.denom, data.amount);
        }
    }

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata packet) external override returns (bytes memory acknowledgement) {
        FungibleTokenPacketData.Data memory data = FungibleTokenPacketData.decode(packet.data);
        strings.slice memory denom = data.denom.toSlice();
        strings.slice memory trimedDenom = data.denom.toSlice().beyond(
            makeDenomPrefix(packet.source_port, packet.source_channel)
        );
        if (!denom.equals(trimedDenom)) { // receiver is source chain
            try bank.transferFrom(getEscrowAddress(packet.destination_channel), data.receiver.toAddress(), trimedDenom.toString(), data.amount) {
                return newAcknowledgement(true);
            } catch (bytes memory) {
                return newAcknowledgement(false);
            }
        } else {
            string memory prefixedDenom = makeDenomPrefix(packet.destination_port, packet.destination_channel).concat(denom);
            try bank.mint(data.receiver.toAddress(), prefixedDenom, data.amount) {
                return newAcknowledgement(true);
            } catch (bytes memory) {
                return newAcknowledgement(false);
            }
        }
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement) external override {
        if (!isSuccessAcknowledgement(acknowledgement)) {
            refundTokens(FungibleTokenPacketData.decode(packet.data), packet.source_port, packet.source_channel);
        }
    }

    function onChanOpenInit(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version) external override {
        // TODO authenticate a capability
        channelEscrowAddresses[channelId] = address(this);
    }

    function onChanOpenTry(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version, string calldata counterpartyVersion) external override {
        // TODO authenticate a capability
        channelEscrowAddresses[channelId] = address(this);
    }

    function onChanOpenAck(string calldata portId, string calldata channelId, string calldata counterpartyVersion) external override {}

    function onChanOpenConfirm(string calldata portId, string calldata channelId) external override {}

    /// Helper functions ///

    function makeDenomPrefix(string memory port, string memory channel) internal pure returns (strings.slice memory) {
        return port.toSlice()
            .concat("/".toSlice()).toSlice()
            .concat(channel.toSlice()).toSlice()
            .concat("/".toSlice()).toSlice();
    }
}
