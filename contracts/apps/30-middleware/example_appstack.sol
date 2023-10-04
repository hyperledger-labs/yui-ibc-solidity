pragma solidity ^0.8.13;

import "forge-std/Test.sol";

library Lib {
    struct Packet {
        bytes data;
    }

    function log(string memory message, Packet memory packet) internal view {
        console2.log(string(abi.encodePacked(message, "/", packet.data)));
    }
}

abstract contract AppBase {
    address immutable ibcAddress;

    modifier onlyIBC() {
        require(msg.sender == ibcAddress || msg.sender == address(this), "caller must be self or ibc");
        _;
    }

    constructor(address ibcAddress_) {
        ibcAddress = ibcAddress_;
    }

    function onRecvPacket(Lib.Packet calldata) public virtual returns (uint256) {
        revert("must be overried");
    }

    function sendPacket(Lib.Packet memory packet) internal virtual {
        Lib.log("Base: send packet via IBC Handler", packet);
        // IBCHandler(ibcAddress).sendPacket(packet);
    }
}

contract App is AppBase {
    constructor(address ibcAddress_) AppBase(ibcAddress_) {}

    function onRecvPacket(Lib.Packet calldata packet) public virtual override onlyIBC returns (uint256) {
        Lib.log("App: recv packet", packet);
        return 1;
    }

    function sendTransfer(string calldata message) public {
        console2.log("App: sendTransfer called");
        sendPacket(Lib.Packet(bytes(message)));
    }
}

abstract contract FeeMiddlewareBase is AppBase {}

abstract contract FeeMiddlewarePacketSender is FeeMiddlewareBase {
    function sendPacket(Lib.Packet memory packet) internal virtual override {
        Lib.log("FeeMiddlewarePacketSender: send packet", packet);
        super.sendPacket(packet);
    }
}

abstract contract FeeMiddlewarePacketReceiver is FeeMiddlewareBase {
    function onRecvPacket(Lib.Packet calldata packet) public virtual override returns (uint256) {
        Lib.log("FeeMiddlewarePacketReceiver: recv packet", packet);
        uint256 res = super.onRecvPacket(packet);
        console2.log("FeeMiddlewarePacketReceiver: after recv packet");
        return res;
    }
}

// this is helpful in a case using a single middleware
abstract contract FeeMiddleware is FeeMiddlewarePacketSender, FeeMiddlewarePacketReceiver {
    function sendPacket(Lib.Packet memory packet) internal virtual override(AppBase, FeeMiddlewarePacketSender) {
        super.sendPacket(packet);
    }

    function onRecvPacket(Lib.Packet calldata packet)
        public
        virtual
        override(AppBase, FeeMiddlewarePacketReceiver)
        returns (uint256)
    {
        return super.onRecvPacket(packet);
    }
}

abstract contract HookMiddlewareBase is AppBase {
    string public dataSuffix;

    constructor(string memory dataSuffix_) {
        dataSuffix = dataSuffix_;
    }
}

abstract contract HookMiddlewarePacketSender is HookMiddlewareBase {
    function sendPacket(Lib.Packet memory packet) internal virtual override {
        packet.data = bytes.concat(packet.data, bytes(dataSuffix));
        Lib.log("HookMiddlewarePacketSender: send packet", packet);
        super.sendPacket(packet);
    }
}

abstract contract HookMiddlewarePacketReceiver is HookMiddlewareBase {
    function onRecvPacket(Lib.Packet calldata packet) public virtual override returns (uint256) {
        Lib.log("HookMiddlewarePacketReceiver: recv packet", packet);
        require(
            packet.data.length > bytes(dataSuffix).length
                && keccak256(packet.data[packet.data.length - bytes(dataSuffix).length:]) == keccak256(bytes(dataSuffix)),
            "invalid suffix"
        );
        uint256 res = super.onRecvPacket(packet);
        console2.log("HookMiddlewarePacketReceiver: after recv packet");
        return res;
    }
}

// this is helpful in a case using a single middleware
abstract contract HookMiddleware is HookMiddlewarePacketSender, HookMiddlewarePacketReceiver {
    function sendPacket(Lib.Packet memory packet) internal virtual override(AppBase, HookMiddlewarePacketSender) {
        super.sendPacket(packet);
    }

    function onRecvPacket(Lib.Packet calldata packet)
        public
        virtual
        override(AppBase, HookMiddlewarePacketReceiver)
        returns (uint256)
    {
        return super.onRecvPacket(packet);
    }
}

contract HookFeeMiddlewaredApp is
    App,
    HookMiddlewarePacketSender,
    FeeMiddlewarePacketSender,
    FeeMiddlewarePacketReceiver,
    HookMiddlewarePacketReceiver
{
    constructor(address ibcAddress_, string memory dataSuffix) App(ibcAddress_) HookMiddlewareBase(dataSuffix) {}

    function sendPacket(Lib.Packet memory packet)
        internal
        virtual
        override(AppBase, HookMiddlewarePacketSender, FeeMiddlewarePacketSender)
    {
        // this function is called by app's sendTransfer
        // fee -> hook -> base
        // NOTE: `base`'s sendPacket just send a packet to handler
        super.sendPacket(packet);
    }

    function onRecvPacket(Lib.Packet calldata packet)
        public
        virtual
        override(AppBase, App, HookMiddlewarePacketReceiver, FeeMiddlewarePacketReceiver)
        onlyIBC
        returns (uint256)
    {
        // hook -> fee -> app
        return super.onRecvPacket(packet);
    }
}

contract MiddlewareTest is Test {
    function test_example_appstack() public {
        // send: app.sendTransfer -> fee.send -> hook.send -> base.send(ibc handler)
        // recv: hook.recv -> fee.recv -> app.recv
        HookFeeMiddlewaredApp app = new HookFeeMiddlewaredApp(address(this), "_hooked");
        app.sendTransfer("hello");
        console2.log("=================");
        app.onRecvPacket(Lib.Packet("hello_hooked"));
    }
}
