pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCMsgs.sol";
import "./IBCChannel.sol";
import "./IBCStore.sol";

/*
IBCHandler implements ics-025 and ics-026
*/
contract IBCHandler {
    IBCStore ibcStore;
    IBCChannel ibcchannel;

    mapping(string => Module) modules;

    constructor(IBCStore store, IBCChannel ibcchannel_) public {
        ibcStore = store;
        ibcchannel = ibcchannel_;
    }

    /// Msg(Datagram) handlers ///

    function handlePacketRecv(IBCMsgs.MsgPacketRecv memory msg_) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        bytes memory acknowledgement = module.callbacks.onRecvPacket(msg_.packet);
        ibcchannel.recvPacket(msg_);
        if (acknowledgement.length > 0) {
            ibcchannel.writeAcknowledgement(msg_.packet, acknowledgement);
        }
    }

    function handlePacketAcknowledgement(IBCMsgs.MsgPacketAcknowledgement memory msg_) public {
        (Module memory module, bool found) = lookupModule(msg_.packet.source_port);
        require(found, "module not found");
        module.callbacks.onAcknowledgementPacket(msg_.packet, msg_.acknowledgement);
        ibcchannel.acknowledgePacket(msg_);
    }

    // WARNING: This function **must be** removed in production
    function handlePacketRecvWithoutVerification(IBCMsgs.MsgPacketRecv memory msg_) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        return module.callbacks.onRecvPacket(msg_.packet);
    }

    /// Module manager ///

    struct Module {
        CallbacksI callbacks;
        bool exists;
    }

    // TODO apply ACL to this
    function bindPort(string memory portId, address moduleAddress) public {
        require(!modules[portId].exists, "the portId is already used by other module");
        modules[portId] = Module({callbacks: CallbacksI(moduleAddress), exists: true});
    }

    function lookupModule(string memory portId) public view returns (Module memory module, bool found) {
        if (!modules[portId].exists) {
            return (module, false);
        }
        return (modules[portId], true);
    }
}

interface CallbacksI {
    function onRecvPacket(Packet.Data calldata) external returns(bytes memory);
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement) external;
}
