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

        if (!denom.toSlice().startsWith(makeDenomPrefix(sourcePort, sourceChannel))) { // sender chain is source
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
            makeDenomPrefix(packet.source_port, packet.source_channel)
        );
        if (!denom.equals(trimedDenom)) { // receiver is source chain
            // TODO try and catch
            if (trimedDenom.len() == 42) {
                IERC20(parseAddr(trimedDenom.toString())).transfer(data.receiver.toAddress(), data.amount);
            } else {
                bank.transferFrom(getEscrowAddress(packet.destination_channel), data.receiver.toAddress(), bytes(trimedDenom.toString()), data.amount);
            }
            return newAcknowledgement(true);
        } else {
            string memory prefixedDenom = makeDenomPrefix(packet.destination_port, packet.destination_channel).concat(denom);
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

    function makeDenomPrefix(string memory port, string memory channel) internal pure returns (strings.slice memory) {
        return port.toSlice()
            .concat("/".toSlice()).toSlice()
            .concat(channel.toSlice()).toSlice()
            .concat("/".toSlice()).toSlice();
    }

    function addressToString(address _address) internal pure returns(string memory) {
        bytes memory alphabet = "0123456789abcdef";
        bytes20 data = bytes20(_address);

        bytes memory str = new bytes(42);
        str[0] = "0";
        str[1] = "x";
        for (uint i = 0; i < 20; i++) {
            str[2+i*2] = alphabet[uint(uint8(data[i] >> 4))];
            str[2+1+i*2] = alphabet[uint(uint8(data[i] & 0x0f))];
        }
        return string(str);
    }

    // a copy from https://github.com/provable-things/ethereum-api/blob/161552ebd4f77090d86482cff8c863cf903c6f5f/oraclizeAPI_0.6.sol
    function parseAddr(string memory _a) internal pure returns (address _parsedAddress) {
        bytes memory tmp = bytes(_a);
        uint160 iaddr = 0;
        uint160 b1;
        uint160 b2;
        for (uint i = 2; i < 2 + 2 * 20; i += 2) {
            iaddr *= 256;
            b1 = uint160(uint8(tmp[i]));
            b2 = uint160(uint8(tmp[i + 1]));
            if ((b1 >= 97) && (b1 <= 102)) {
                b1 -= 87;
            } else if ((b1 >= 65) && (b1 <= 70)) {
                b1 -= 55;
            } else if ((b1 >= 48) && (b1 <= 57)) {
                b1 -= 48;
            }
            if ((b2 >= 97) && (b2 <= 102)) {
                b2 -= 87;
            } else if ((b2 >= 65) && (b2 <= 70)) {
                b2 -= 55;
            } else if ((b2 >= 48) && (b2 <= 57)) {
                b2 -= 48;
            }
            iaddr += (b1 * 16 + b2);
        }
        return address(iaddr);
    }
}
