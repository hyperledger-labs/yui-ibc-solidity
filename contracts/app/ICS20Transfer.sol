pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/types/Channel.sol";
import "../core/IBCModule.sol";
import "../core/IBCHandler.sol";
import "../core/IBCHost.sol";
import "../core/types/App.sol";
import "./IICS20Vouchers.sol";
import "../lib/strings.sol";
import "../lib/Bytes.sol";
import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/utils/Address.sol";

contract ICS20Transfer is IModuleCallbacks {
    using Address for address;
    using strings for *;
    using Bytes for *;

    IBCHandler ibcHandler;
    IBCHost ibcHost;
    IICS20Vouchers bank;

    constructor(IBCHost host_, IBCHandler ibcHandler_, IICS20Vouchers bank_) public {
        ibcHost = host_;
        ibcHandler = ibcHandler_;
        bank = bank_;
    }

    mapping(string => address) channelEscrowAddresses;

    function sendTransferWithTokenContract(
        address tokenContract,
        uint256 amount,
        address receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight        
    ) external {
        require(tokenContract.isContract());
        (Channel.Data memory channel, bool found) = ibcHost.getChannel(sourcePort, sourceChannel);
        require(found, "channel not found");

        IERC20(tokenContract).transferFrom(msg.sender, address(this), amount);

        bytes memory data = FungibleTokenPacketData.encode(FungibleTokenPacketData.Data({
            denom: addressToString(tokenContract),
            amount: uint64(amount), // TODO fix type
            sender: abi.encodePacked(msg.sender),
            receiver: abi.encodePacked(receiver)
        }));
        ibcHandler.sendPacket(Packet.Data({
            sequence: ibcHost.getNextSequenceSend(sourcePort, sourceChannel),
            source_port: sourcePort,
            source_channel: sourceChannel,
            destination_port: channel.counterparty.port_id,
            destination_channel: channel.counterparty.channel_id,
            data: data,
            timeout_height: Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            timeout_timestamp: 0
        }));
    }

    function sendTransfer(
        string calldata denom,
        uint256 amount,
        address receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external {
        (Channel.Data memory channel, bool found) = ibcHost.getChannel(sourcePort, sourceChannel);
        require(found, "channel not found");

        bool source = !denom.toSlice().startsWith(
            sourcePort.toSlice().concat(
            "/".toSlice()
            ).toSlice().concat(sourceChannel.toSlice()).toSlice()
        );

        if (source) {
            bank.transferFrom(msg.sender, getEscrowAddress(sourceChannel), bytes(denom), amount);
        } else {
            bank.burnFrom(msg.sender, bytes(denom), amount);
        }

        bytes memory data = FungibleTokenPacketData.encode(FungibleTokenPacketData.Data({
            denom: denom,
            amount: uint64(amount), // TODO fix type
            sender: abi.encodePacked(msg.sender),
            receiver: abi.encodePacked(receiver)
        }));
        ibcHandler.sendPacket(Packet.Data({
            sequence: ibcHost.getNextSequenceSend(sourcePort, sourceChannel),
            source_port: sourcePort,
            source_channel: sourceChannel,
            destination_port: channel.counterparty.port_id,
            destination_channel: channel.counterparty.channel_id,
            data: data,
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

    // function refundTokens(Packet.Data memory packet) internal {
    //     // TODO implements
    // }

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata packet) external override returns (bytes memory acknowledgement) {
        FungibleTokenPacketData.Data memory data = FungibleTokenPacketData.decode(packet.data);
        strings.slice memory denom = data.denom.toSlice();
        strings.slice memory trimedDenom = data.denom.toSlice().beyond(
            packet.source_port.toSlice().concat(
            "/".toSlice()
            ).toSlice().concat(packet.source_channel.toSlice()).toSlice()
        );
        if (!denom.equals(trimedDenom)) { // receiver is source chain
            // TODO try and catch
            bank.transferFrom(getEscrowAddress(packet.destination_channel), data.receiver.toAddress(), bytes(trimedDenom.toString()), data.amount);
            return newAcknowledgement(true);
        } else {
            string memory prefixedDenom = packet.destination_port.toSlice().concat(
            "/".toSlice()
            ).toSlice().concat(packet.destination_channel.toSlice()).toSlice().concat(denom);
            // TODO try and catch
            bank.mint(data.receiver.toAddress(), bytes(prefixedDenom), data.amount);
            return newAcknowledgement(true);
        }
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement) external override {
        if (!isSuccessAcknowledgement(acknowledgement)) {
            // refundTokens(packet);
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

    function addressToString(address _address) internal pure returns (string memory _uintAsString) {
        uint _i = uint256(_address);
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (_i != 0) {
            bstr[k--] = byte(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);
    }
}
