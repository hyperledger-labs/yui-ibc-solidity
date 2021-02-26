pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/types/Channel.sol";
import "../core/IBCChannel.sol";
import "../core/IBCHandler.sol";
import "../core/IBCHost.sol";
import "../core/IBCModule.sol";
import "../core/types/App.sol";
import "../lib/IBCIdentifier.sol";
import "../lib/Bytes.sol";

contract SimpleTokenModule is IModuleCallbacks {
    using Bytes for bytes;

    /// Storages ///

    // Token storages
    mapping (address => uint256) _balances;
    uint256 lockedBalance;

    // Module storages
    IBCHandler ibcHandler;
    IBCHost ibcHost;

    /// Constructor ///

    constructor(IBCHost host_, IBCHandler ibcHandler_) public {
        ibcHost = host_;
        ibcHandler = ibcHandler_;

        _balances[msg.sender] = 10000;
    }

    /// Token API ///

    function transfer(address recipient, uint256 amount) public returns (bool) {
        return _transfer(msg.sender, recipient, amount);
    }

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function crossTransfer(string memory sourcePort, string memory sourceChannel, address recipient, uint64 amount, uint64 timeoutHeight) public {
        // ensure that the sender has sufficient balance
        require(balanceOf(msg.sender) >= amount, "insufficient balance");

        (Channel.Data memory channel, bool found) = ibcHost.getChannel(sourcePort, sourceChannel);
        require(found, "channel not found");

        bytes memory data = FungibleTokenPacketData.encode(FungibleTokenPacketData.Data({
            denom: "default",
            amount: amount,
            sender: abi.encodePacked(msg.sender),
            receiver: abi.encodePacked(recipient)
        }));
        Packet.Data memory packet = Packet.Data({
            sequence: ibcHost.getNextSequenceSend(sourcePort, sourceChannel),
            source_port: sourcePort,
            source_channel: sourceChannel,
            destination_port: channel.counterparty.port_id,
            destination_channel: channel.counterparty.channel_id,
            data: data,
            timeout_height: Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            timeout_timestamp: 0
        });
        ibcHandler.sendPacket(packet);
        lock(msg.sender, amount);
    }

    /// Internal functions ///

    function _transfer(address sender, address recipient, uint256 amount) internal returns (bool) {
        uint256 senderBalance = _balances[sender];
        require(senderBalance >= amount, "transfer amount exceeds balance");
        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;
        return true;
    }

    function mint(address recipient, uint256 amount) internal {
        _balances[recipient] += amount;
    }

    function lock(address sender, uint256 amount) internal {
        uint256 senderBalance = _balances[sender];
        require(senderBalance >= amount, "burn amount exceeds balance");
        _balances[sender] = senderBalance - amount;
        lockedBalance += amount;
    }

    function burn(uint256 amount) internal {
        require(lockedBalance >= amount);
        lockedBalance -= amount;
    }

    /// Module implementations ///

    modifier onlyIBCModule (){
        require(msg.sender == address(ibcHandler));
        _;
    }

    function onRecvPacket(Packet.Data calldata packet) onlyIBCModule external override returns (bytes memory acknowledgement) {
        FungibleTokenPacketData.Data memory data = FungibleTokenPacketData.decode(packet.data);
        mint(data.receiver.toAddress(), data.amount);
        acknowledgement = new bytes(1);
        acknowledgement[0] = 0x01;
        return acknowledgement;
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement) onlyIBCModule external override {
        FungibleTokenPacketData.Data memory data = FungibleTokenPacketData.decode(packet.data);
        // if acknowledgement indicates an error, refund the tokens to sender
        if (acknowledgement.length == 1 && acknowledgement[0] == 0x01) {
            burn(data.amount);
        } else {
            burn(data.amount);
            mint(data.sender.toAddress(), data.amount);
        }
    }

    function onChanOpenInit(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version) external override {}
    function onChanOpenTry(Channel.Order, string[] calldata connectionHops, string calldata portId, string calldata channelId, ChannelCounterparty.Data calldata counterparty, string calldata version, string calldata counterpartyVersion) external override {}
    function onChanOpenAck(string calldata portId, string calldata channelId, string calldata counterpartyVersion) external override {}
    function onChanOpenConfirm(string calldata portId, string calldata channelId) external override {}
}
