pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCMsgs.sol";
import "./IBCChannel.sol";
import "./ProvableStore.sol";

/*
IBCHandler implements ics-025 and ics-026
*/
contract IBCHandler {
    ProvableStore provableStore;
    IBCChannel ibcchannel;

    mapping(string => Module) modules;

    struct Module {
        CallbacksI callbacks;
        bool exists;
    }

    constructor(ProvableStore store, IBCChannel ibcchannel_) public {
        provableStore = store;
        ibcchannel = ibcchannel_;
    }

    /// Module manager ///

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

    /// Datagram handlers ///

    function handlePacketRecv(IBCMsgs.MsgPacketRecv memory datagram) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(datagram.packet.destination_port);
        require(found, "module not found");
        bytes memory acknowledgement = module.callbacks.onRecvPacket(datagram.packet);
        ibcchannel.recvPacket(datagram.packet, datagram.proof, datagram.proofHeight);
        if (acknowledgement.length > 0) {
            ibcchannel.writeAcknowledgement(datagram.packet, acknowledgement);
        }
    }

    function handlePacketAcknowledgement(IBCMsgs.MsgPacketAcknowledgement memory datagram) public {
        (Module memory module, bool found) = lookupModule(datagram.packet.source_port);
        require(found, "module not found");
        module.callbacks.onAcknowledgementPacket(datagram.packet, datagram.acknowledgement);
        ibcchannel.acknowledgePacket(datagram.packet, datagram.acknowledgement, datagram.proof, datagram.proofHeight);
    }

    // WARNING: This function **must be** removed in production
    function handlePacketRecvWithoutVerification(IBCMsgs.MsgPacketRecv memory datagram) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(datagram.packet.destination_port);
        require(found, "module not found");
        return module.callbacks.onRecvPacket(datagram.packet);
    }
}

interface CallbacksI {
    function onRecvPacket(Packet.Data calldata) external returns(bytes memory);
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement) external;
}
