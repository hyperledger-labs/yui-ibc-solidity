// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";

import "../../proto/Channel.sol";
import "../../core/04-channel/IIBCChannel.sol";
import "../20-transfer/ICS20Lib.sol";
import "../20-transfer/IICS20Bank.sol";
import "../20-transfer/ICS20TransferBank.sol";

import "solidity-bytes-utils/contracts/BytesLib.sol";

library Osmosis {
    using BytesLib for bytes;

    /// it is all sorted jsons, see https://github.com/cosmos/ibc-go/blob/main/modules/core/04-channel/types/acknowledgement.go
    function marshalError(string memory error) internal pure returns (bytes memory) {
        return abi.encodePacked('{"error":"', error, '"}');
    }

    function marshalResult(bytes memory result) internal pure returns (bytes memory) {
        return abi.encodePacked('{"result":"', result, '"}');
    }

    /// ContractAck is the response to be stored when a evm hook is executed
    function marshalContractAck(bytes memory contract_result, bytes memory ibc_ack)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked('{"contract_result":"', contract_result, '","ibc_ack":"', ibc_ack, '"}');
    }

    function isSuccess(bytes calldata acknowledgement) internal pure returns (bool) {
        // https://github.com/osmosis-labs/osmosis/blob/main/osmoutils/ibc.go
        return (bytes32(acknowledgement.slice(0, 10)) == bytes32('{"error":"')) && acknowledgement.length > 12;
    }

    // here we should overide send, and grab callback for tracking
    // and override ack/timeout and check if there is callback

    function deriveIntermediateSender(string memory channel, string memory sender) internal pure returns (bytes32) {
        return bytes32(keccak256(abi.encodePacked(channel, sender)));
    }
}

abstract contract OsmosisHookMiddlewarePacketReceiver is IBCAppBaseOnRecvPacket {
    using BytesLib for bytes;

    /// passed to executor so it can transfer denoms toward contract invoked
    IICS20Bank internal immutable ics20Bank;

    constructor(IICS20Bank ics20Bank_) {
        ics20Bank = ics20Bank_;
    }

    /// @notice maps unique sender to contract address
    mapping(bytes32 => address) public channelSenderToContract;

    function onRecvPacket(Packet.Data memory packet, address relayer)
        public
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement, bool success)
    {
        bool isIcs20;
        ICS20Lib.PacketData memory data;
        (isIcs20, data) = ICS20Lib.tryUnmarshalJSON(packet.data);

        if (!isIcs20) {
            return super.onRecvPacket(packet, relayer);
        }

        bool isEvmRouted;
        address contractAddress;
        bytes memory ethAbi;
        (isEvmRouted, contractAddress, ethAbi) = _validataAndParseMemo(data.memo);
        if (!isEvmRouted) {
            return super.onRecvPacket(packet, relayer);
        }

        // please read Cosmos SDK hooks about what is going on here
        bytes32 intermediateSender = Osmosis.deriveIntermediateSender(packet.destination_channel, data.sender);

        address executor = channelSenderToContract[intermediateSender];

        if (executor == address(0)) {
            executor = address(new IbcOsmosisHookExecutor{salt: intermediateSender}(ics20Bank));
            channelSenderToContract[intermediateSender] = executor;
        }

        require(
            keccak256(abi.encodePacked(data.receiver)) == keccak256(abi.encodePacked(Strings.toHexString(executor))),
            "Sender must use predictable address derived from channel and sender used as salt so that sender always sends to well known app stack"
        );

        bytes memory newData = ICS20Lib.marshalJSON(data.denom, data.amount, data.sender, data.receiver, "");

        Packet.Data memory newPacket = Packet.Data(
            packet.sequence,
            packet.source_port,
            packet.source_channel,
            packet.destination_port,
            packet.destination_channel,
            newData,
            packet.timeout_height,
            packet.timeout_timestamp
        );

        (bytes memory appAcknowledgement, bool success) = super.onRecvPacket(newPacket, relayer);

        if (!success) {
            return (appAcknowledgement, success);
        }

        bytes memory result;
        (success, result) = IIbcOsmosisHookExecutor(executor).execute(data.denom, data.amount, contractAddress, ethAbi);

        if (!success) {
            /// not really sure how to port next:
            /// https://github.com/osmosis-labs/osmosis/blob/main/x/ibc-hooks/types/errors.go
            /// ErrWasmError           = errorsmod.Register("wasm-hooks", 6, "wasm error")
            /// Error: fmt.Sprintf("ABCI code: %d: %s", code, ackErrorString),
            /// but pretty sure that `error` in ACK works properly,
            /// so best approximation for now
            return (Osmosis.marshalError("evm error"), success);
        }

        // so here osmosis does something fo async packets, but as we do not have them, do nothing for now

        return (marshalResult(result, acknowledgement), true);
    }

    function marshalResult(bytes memory result, bytes memory acknowledgement) internal pure returns (bytes memory) {
        return Osmosis.marshalResult(Osmosis.marshalContractAck(result, acknowledgement));
    }

    function _validataAndParseMemo(string memory memo)
        private
        pure
        returns (bool isEvmRouted, address contractAddress, bytes memory msg_)
    {
        // sender must have this one (same order of properties)
        // ```json
        // {
        //   "evm": {
        //     "contract": "0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
        //     "msg": "0xc87b56dda752230262935940d907f047a9f86bb5ee6aa33511fc86db33fea6cc"
        //   }
        // }
        // ```
        // minify
        // ```json
        // {"evm":{"contract":"0x71C7656EC7ab88b098defB751B7401B5f6d8976F","msg":"0xc87b56dda752230262935940d907f047a9f86bb5ee6aa33511fc86db33fea6cc"}}
        // ```
        // and escape
        // ```
        // {\"evm\":{\"contract\":\"0x71C7656EC7ab88b098defB751B7401B5f6d8976F\",\"msg\":\"0xc87b56dda752230262935940d907f047a9f86bb5ee6aa33511fc86db33fea6cc\"}}
        // ```
        // we do not change any of words, because
        // 1. allow existing indexers/senders to operate more or less
        // 2. contract address unique identifiers what VM in domain to call anyway
        // 3. if osmosis will replace json with pb, than we can change it

        bytes memory bz = bytes(memo);
        uint256 pos = 0;
        unchecked {
            if (bytes32(bz.slice(pos, pos + 25)) != bytes32('{"evm":{"contract":"')) {
                contractAddress = abi.decode(bz.slice(pos + 25, pos + 25 + 42), (address));
                pos += 42;
            } else {
                return (false, address(0), "");
            }

            if (bytes32(bz.slice(pos, pos + 13)) != bytes32('","msg":"')) {
                msg_ = abi.decode(bz.slice(pos + 13, bz.length - 4), (bytes));
            } else {
                return (false, address(0), "");
            }
        }

        return (true, contractAddress, msg_);
    }
}

/// Here is short explnation in design difference of Go vs Solidity version:
/// Go does not support interitnace and overrides.
/// So each time one want to change behavior in middleware, it have to impl 100% of middleware or fork code.
/// So they invented inherirtance with override hooks (hook is just  single method interface per method).
/// Evidently Solidity does not need that.
/// Additionally, Go code hard to upgrade (as it is native), hooks are easier.
/// There is no such problem with Solidity, so implementation is just usual middleware override.
contract OsmosisHookICS20AppStack is ICS20TransferBank, OsmosisHookMiddlewarePacketReceiver {
    using BytesLib for bytes;

    constructor(IICS04Wrapper ics04Wrapper_, IICS20Bank ics20Bank_)
        ICS20TransferBank(this, ics20Bank_)
        OsmosisHookMiddlewarePacketReceiver(ics20Bank_)
    {}

    function onRecvPacket(Packet.Data memory packet, address relayer)
        public
        virtual
        override(ICS20Transfer, OsmosisHookMiddlewarePacketReceiver)
        onlyIBC
        returns (bytes memory acknowledgement, bool success)
    {
        return super.onRecvPacket(packet, relayer);
    }
}

interface IIbcOsmosisHookExecutor {
    function execute(string memory denom, uint256 amount, address contractAddress, bytes memory msg_)
        external
        returns (bool, bytes memory);
}

/// use used to impersonate sender per channel
contract IbcOsmosisHookExecutor is IIbcOsmosisHookExecutor {
    IICS20Bank private ics20Bank;

    constructor(IICS20Bank ics20Bank_) {
        ics20Bank = ics20Bank_;
    }

    function execute(string memory denom, uint256 amount, address contractAddress, bytes memory msg_)
        external
        returns (bool, bytes memory)
    {
        ics20Bank.transferFrom(address(this), contractAddress, denom, amount);
        // osmosis hooks does not specify how to handle gaeys and gas token at all
        // this to be added as soon as they implemente (or port from official IBC hooks)
        // currently assuming that it is IBC port/channel spec is up to set proper gas enought to write ack
        return contractAddress.call{gas: gasleft()}(msg_);
    }
}

// allow to notifiy sender about packet execution success final as defined by original sender
// should never fail, be fast, just for tracking
// called by `IbcOsmosisHook` from IBC app stack
interface IIbcOsmosisHookCallback {
    // @parameter success - true in casse of both transfer and contract execution success, falsse - if both reverted
    function onAcknowledgementPacket(string calldata port, string calldata channnel, uint64 sequence, bool success)
        external;

    function onTimeoutPacket(string calldata port, string calldata channnel, uint64 sequence) external;
}
