// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

library IBCClientLib {
    /**
     * @dev validateClientType validates the client type
     *      - clientType must be non-empty
     *      - clientType must be between 7 and 62 characters long
     *        - This is because the length of the client ID is 9-64 characters long in the ICS-24, and the client ID is composed of the client type and the counter suffix (minimum 2 characters long).
     *      - clientType must be in the form of `^[a-z][a-z0-9-]*[a-z0-9]$`
     */
    function validateClientType(bytes memory clientTypeBytes) internal pure returns (bool) {
        uint256 bytesLength = clientTypeBytes.length;
        if (bytesLength == 0) {
            return false;
        }
        if (bytesLength < 7 || bytesLength > 62) {
            return false;
        }
        for (uint256 i = 0; i < bytesLength; i++) {
            uint256 c = uint256(uint8(clientTypeBytes[i]));
            if (0x61 <= c && c <= 0x7a) {
                // a-z
                continue;
            } else if (c == 0x2d) {
                // hyphen cannot be the first or last character
                unchecked {
                    // SAFETY: `bytesLength` is greater than 0
                    if (i == 0 || i == bytesLength - 1) {
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

    /**
     * @dev validateClientId validates the client ID
     *      NOTE: The client ID must be composed of the client type is validated by `validateClientType` and the counter suffix.
     *      - clientId must be between 9 and 64 characters long
     */
    function validateClientId(bytes memory clientIdBytes) internal pure returns (bool) {
        uint256 bytesLength = clientIdBytes.length;
        if (bytesLength < 9 || bytesLength > 64) {
            return false;
        }
        return true;
    }
}
