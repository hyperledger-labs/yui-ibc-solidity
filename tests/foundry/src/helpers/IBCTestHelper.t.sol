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
import "../../../../contracts/clients/MockClient.sol";
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
            new IBCChannelUpgrade()
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

    function roll(uint256 bn) internal {
        require(prevBlockNumber == 0, "roll: already rolled");
        prevBlockNumber = block.number;
        vm.roll(bn);
    }

    function unroll() internal {
        require(prevBlockNumber != 0, "unroll: not rolled");
        vm.roll(prevBlockNumber);
        prevBlockNumber = 0;
    }

    function warp(uint256 seconds_) internal {
        require(prevBlockTimestamp == 0, "warp: already warped");
        prevBlockTimestamp = block.timestamp;
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
