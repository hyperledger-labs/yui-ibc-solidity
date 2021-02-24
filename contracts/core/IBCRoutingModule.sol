pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./IBCChannel.sol";
import "./IBCMsgs.sol";
import "./types/Channel.sol";

contract IBCRoutingModule is IBCHost {
    // TODO move it into IBCStore
    mapping(string => Module) modules;

    IBCChannel ibcchannel;

    constructor(IBCStore store, IBCChannel ibcchannel_) IBCHost(store) public {
        ibcchannel = ibcchannel_;
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

    /// Packet Handler ///

    function recvPacket(IBCMsgs.MsgPacketRecv memory msg_) public returns (bytes memory) {
        (Module memory module, bool found) = lookupModule(msg_.packet.destination_port);
        require(found, "module not found");
        bytes memory acknowledgement = module.callbacks.onRecvPacket(msg_.packet);
        ibcchannel.recvPacket(msg_);
        if (acknowledgement.length > 0) {
            ibcchannel.writeAcknowledgement(msg_.packet, acknowledgement);
        }
    }

    function acknowledgePacket(IBCMsgs.MsgPacketAcknowledgement memory msg_) public {
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
}

interface CallbacksI {
    function onRecvPacket(Packet.Data calldata) external returns(bytes memory);
    function onAcknowledgementPacket(Packet.Data calldata, bytes calldata acknowledgement) external;
}
