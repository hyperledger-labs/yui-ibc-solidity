// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../proto/Client.sol";

library IBCHeight {
    function toUint128(Height.Data memory self) internal pure returns (uint128) {
        return (uint128(self.revision_number) << 64) | uint128(self.revision_height);
    }

    function fromUint128(uint128 height) internal pure returns (Height.Data memory) {
        return Height.Data({revision_number: uint64(height >> 64), revision_height: uint64(height)});
    }

    function isZero(Height.Data memory self) internal pure returns (bool) {
        return self.revision_number == 0 && self.revision_height == 0;
    }

    function lt(Height.Data memory self, Height.Data memory other) internal pure returns (bool) {
        return self.revision_number < other.revision_number
            || (self.revision_number == other.revision_number && self.revision_height < other.revision_height);
    }

    function lte(Height.Data memory self, Height.Data memory other) internal pure returns (bool) {
        return self.revision_number < other.revision_number
            || (self.revision_number == other.revision_number && self.revision_height <= other.revision_height);
    }

    function eq(Height.Data memory self, Height.Data memory other) internal pure returns (bool) {
        return self.revision_number == other.revision_number && self.revision_height == other.revision_height;
    }

    function gt(Height.Data memory self, Height.Data memory other) internal pure returns (bool) {
        return self.revision_number > other.revision_number
            || (self.revision_number == other.revision_number && self.revision_height > other.revision_height);
    }

    function gte(Height.Data memory self, Height.Data memory other) internal pure returns (bool) {
        return self.revision_number > other.revision_number
            || (self.revision_number == other.revision_number && self.revision_height >= other.revision_height);
    }
}
