// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

library IBCClientLib {
    /**
     * @dev validateClientType validates the client type
     *   - clientType must be non-empty
     *   - clientType must be in the form of `^[a-z][a-z0-9-]*[a-z0-9]$`
     */
    function validateClientType(bytes memory clientTypeBytes) internal pure returns (bool) {
        uint256 byteLength = clientTypeBytes.length;
        if (byteLength == 0) {
            return false;
        }
        for (uint256 i = 0; i < byteLength; i++) {
            uint256 c = uint256(uint8(clientTypeBytes[i]));
            if (0x61 <= c && c <= 0x7a) {
                // a-z
                continue;
            } else if (c == 0x2d) {
                // hyphen cannot be the first or last character
                unchecked {
                    // SAFETY: `byteLength` is greater than 0
                    if (i == 0 || i == byteLength - 1) {
                        return false;
                    }
                }
                continue;
            } else if (0x30 <= c && c <= 0x39) {
                // 0-9
                continue;
            } else {
                return false;
            }
        }
        return true;
    }
}
