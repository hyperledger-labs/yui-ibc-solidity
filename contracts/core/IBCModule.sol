pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCMsgs.sol";
import "./IBCChannel.sol";
import "./IBCStore.sol";
import "./IBCRoutingModule.sol";
import "./IHandler.sol";
import "./IBFT2Client.sol";

/*
IBCModule implements ics-025 and ics-026
*/
contract IBCModule is IHandler, IBCRoutingModule, IBCChannel {

    constructor(IBCStore store, IBFT2Client ibft2Client) IBCHost(store) public {
        registerClient("ibft2", ibft2Client);
    }

    /* Packet Handler */

    function recvPacket(IBCMsgs.MsgPacketRecv memory msg_) public override returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        bytes memory acknowledgement = module.callbacks.onRecvPacket(msg_.packet);
        _recvPacket(msg_);
        if (acknowledgement.length > 0) {
            writeAcknowledgement(msg_.packet, acknowledgement);
        }
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement memory msg_) public override {
        (Module memory module, bool found) = lookupModule(msg_.packet.source_port);
        require(found, "module not found");
        module.callbacks.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        _acknowledgePacket(msg_);
    }

    // WARNING: This function **must be** removed in production
    function handlePacketRecvWithoutVerification(IBCMsgs.MsgPacketRecv memory msg_) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        return module.callbacks.onRecvPacket(msg_.packet);
    }
}
