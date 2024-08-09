// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import "../../../../contracts/core/02-client/IBCClient.sol";
import "../../../../contracts/core/03-connection/IBCConnectionSelfStateNoValidation.sol";
import "../../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import "../../../../contracts/core/04-channel/IBCChannelPacketSendRecv.sol";
import "../../../../contracts/core/04-channel/IBCChannelPacketTimeout.sol";
import "../../../../contracts/core/04-channel/IBCChannelUpgrade.sol";
import "../../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../../contracts/proto/MockClient.sol";
import "../../../../contracts/proto/Connection.sol";
import "../../../../contracts/proto/Channel.sol";
import "../../../../contracts/apps/mock/IBCMockApp.sol";
import "../../../../contracts/clients/mock/MockClient.sol";
import {TestableIBCHandler} from "./TestableIBCHandler.t.sol";

abstract contract IBCTestHelper is Test {
    bytes internal constant DEFAULT_COMMITMENT_PREFIX = bytes("ibc");

    uint256 internal prevBlockNumber;
    uint256 internal prevBlockTimestamp;

    function defaultIBCHandler() internal returns (TestableIBCHandler) {
        return new TestableIBCHandler(
            new IBCClient(),
            new IBCConnectionSelfStateNoValidation(),
            new IBCChannelHandshake(),
            new IBCChannelPacketSendRecv(),
            new IBCChannelPacketTimeout(),
            new IBCChannelUpgradeInitTryAck(),
            new IBCChannelUpgradeConfirmOpenTimeoutCancel()
        );
    }

    // solhint-disable func-name-mixedcase
    function H(uint256 revisionNumber, uint256 revisionHeight) internal pure returns (Height.Data memory) {
        return Height.Data({revision_number: uint64(revisionNumber), revision_height: uint64(revisionHeight)});
    }

    // solhint-disable func-name-mixedcase
    function H(uint256 revisionHeight) internal pure returns (Height.Data memory) {
        return H(0, revisionHeight);
    }

    function getBlockNumber() internal view returns (uint256) {
        return vm.getBlockNumber();
    }

    function getBlockNumber(int256 offset) internal view returns (uint256) {
        uint256 blockNumber = vm.getBlockNumber();
        require(int256(blockNumber) + offset >= 0, "getBlockNumber: negative block number");
        return blockNumber + uint256(offset);
    }

    function getBlockTimestampNano() internal view returns (uint64) {
        return getBlockTimestampNano(0);
    }

    function getBlockTimestampNano(int256 offsetSecs) internal view returns (uint64) {
        uint256 timestamp = vm.getBlockTimestamp();
        require(int256(timestamp) + offsetSecs >= 0, "getTimestamp: negative timestamp");
        return uint64(uint256((int256(timestamp) + offsetSecs) * 1e9));
    }

    function roll(uint256 bn) internal {
        require(prevBlockNumber == 0, "roll: already rolled");
        prevBlockNumber = vm.getBlockNumber();
        vm.roll(bn);
    }

    function unroll() internal {
        require(prevBlockNumber != 0, "unroll: not rolled");
        vm.roll(prevBlockNumber);
        prevBlockNumber = 0;
    }

    function warp(uint256 seconds_) internal {
        require(prevBlockTimestamp == 0, "warp: already warped");
        prevBlockTimestamp = vm.getBlockTimestamp();
        vm.warp(seconds_);
    }

    function unwarp() internal {
        require(prevBlockTimestamp != 0, "unwarp: not warped");
        vm.warp(prevBlockTimestamp);
        prevBlockTimestamp = 0;
    }

    function rollAndWarp(uint256 bn, uint256 seconds_) internal {
        roll(bn);
        warp(seconds_);
    }

    function unrollAndUnwarp() internal {
        unroll();
        unwarp();
    }
}
