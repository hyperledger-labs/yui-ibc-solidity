// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

library IBCHostLib {
    /**
     * @dev validatePortIdentifier validates a port identifier string
     * check if the string consist of characters in one of the following categories only:
     * - Alphanumeric
     * - `.`, `_`, `+`, `-`, `#`
     * - `[`, `]`, `<`, `>`
     *
     * The validation is based on the ibc-go implementation:
     * https://github.com/cosmos/ibc-go/blob/b0ed0b412ea75e66091cc02746c95f9e6cf4445d/modules/core/24-host/validate.go#L86
     */
    function validatePortIdentifier(bytes memory portId) internal pure returns (bool) {
        uint256 portIdLength = portId.length;
        if (portIdLength < 2 || portIdLength > 128) {
            return false;
        }
        unchecked {
            for (uint256 i = 0; i < portIdLength; i++) {
                uint256 c = uint256(uint8(portId[i]));
                if (
                    // a-z
                    !(c >= 0x61 && c <= 0x7A)
                    // 0-9
                    && !(c >= 0x30 && c <= 0x39)
                    // A-Z
                    && !(c >= 0x41 && c <= 0x5A)
                    // ".", "_", "+", "-"
                    && !(c == 0x2E || c == 0x5F || c == 0x2B || c == 0x2D)
                    // "#", "[", "]", "<", ">"
                    && !(c == 0x23 || c == 0x5B || c == 0x5D || c == 0x3C || c == 0x3E)
                ) {
                    return false;
                }
            }
        }
        return true;
    }
}
