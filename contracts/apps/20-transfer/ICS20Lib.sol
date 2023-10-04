// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";

library ICS20Lib {
    /**
     * @dev PacketData is defined in [ICS-20](https://github.com/cosmos/ibc/tree/main/spec/app/ics-020-fungible-token-transfer).
     */
    struct PacketData {
        string denom;
        string sender;
        string receiver;
        uint256 amount;
        string memo;
    }

    bytes public constant SUCCESSFUL_ACKNOWLEDGEMENT_JSON = bytes('{"result":"AQ=="}');
    bytes public constant FAILED_ACKNOWLEDGEMENT_JSON = bytes('{"error":"failed"}');

    uint256 private constant CHAR_DOUBLE_QUOTE = 0x22;
    uint256 private constant CHAR_SLASH = 0x2f;
    uint256 private constant CHAR_BACKSLASH = 0x5c;
    uint256 private constant CHAR_F = 0x66;
    uint256 private constant CHAR_R = 0x72;
    uint256 private constant CHAR_N = 0x6e;
    uint256 private constant CHAR_B = 0x62;
    uint256 private constant CHAR_T = 0x74;
    uint256 private constant CHAR_CLOSING_BRACE = 0x7d;
    uint256 private constant CHAR_M = 0x6d;

    bytes16 private constant HEX_DIGITS = "0123456789abcdef";

    /**
     * @dev marshalUnsafeJSON marshals PacketData into JSON bytes without escaping.
     *      `memo` field is omitted if it is empty.
     */
    function marshalUnsafeJSON(PacketData memory data) internal pure returns (bytes memory) {
        if (bytes(data.memo).length == 0) {
            return marshalJSON(data.denom, data.amount, data.sender, data.receiver);
        } else {
            return marshalJSON(data.denom, data.amount, data.sender, data.receiver, data.memo);
        }
    }

    /**
     * @dev marshalJSON marshals PacketData into JSON bytes with escaping.
     */
    function marshalJSON(
        string memory escapedDenom,
        uint256 amount,
        string memory escapedSender,
        string memory escapedReceiver,
        string memory escapedMemo
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            '{"amount":"',
            Strings.toString(amount),
            '","denom":"',
            escapedDenom,
            '","memo":"',
            escapedMemo,
            '","receiver":"',
            escapedReceiver,
            '","sender":"',
            escapedSender,
            '"}'
        );
    }

    /**
     * @dev marshalJSON marshals PacketData into JSON bytes with escaping.
     */
    function marshalJSON(
        string memory escapedDenom,
        uint256 amount,
        string memory escapedSender,
        string memory escapedReceiver
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            '{"amount":"',
            Strings.toString(amount),
            '","denom":"',
            escapedDenom,
            '","receiver":"',
            escapedReceiver,
            '","sender":"',
            escapedSender,
            '"}'
        );
    }

    /**
     * @dev unmarshalJSON unmarshals JSON bytes into PacketData.
     */
    function unmarshalJSON(bytes calldata bz) internal pure returns (PacketData memory) {
        PacketData memory pd;
        uint256 pos = 0;

        unchecked {
            require(bytes32(bz[pos:pos + 11]) == bytes32('{"amount":"'), "amount");
            (pd.amount, pos) = parseUint256String(bz, pos + 11);

            require(bytes32(bz[pos:pos + 10]) == bytes32(',"denom":"'), "denom");
            (pd.denom, pos) = parseString(bz, pos + 10);

            if (uint256(uint8(bz[pos + 2])) == CHAR_M) {
                require(bytes32(bz[pos:pos + 9]) == bytes32(',"memo":"'), "memo");
                (pd.memo, pos) = parseString(bz, pos + 9);
            }

            require(bytes32(bz[pos:pos + 13]) == bytes32(',"receiver":"'), "receiver");
            (pd.receiver, pos) = parseString(bz, pos + 13);

            require(bytes32(bz[pos:pos + 11]) == bytes32(',"sender":"'), "sender");
            (pd.sender, pos) = parseString(bz, pos + 11);

            require(pos == bz.length - 1 && uint256(uint8(bz[pos])) == CHAR_CLOSING_BRACE, "closing brace");
        }

        return pd;
    }

    /**
     * @dev parseUint256String parses `bz` from a position `pos` to produce a uint256.
     */
    function parseUint256String(bytes calldata bz, uint256 pos) internal pure returns (uint256, uint256) {
        uint256 ret = 0;
        unchecked {
            for (; pos < bz.length; pos++) {
                uint256 c = uint256(uint8(bz[pos]));
                if (c < 48 || c > 57) {
                    break;
                }
                ret = ret * 10 + (c - 48);
            }
            require(pos < bz.length && uint256(uint8(bz[pos])) == CHAR_DOUBLE_QUOTE, "unterminated string");
            return (ret, pos + 1);
        }
    }

    /**
     * @dev parseString parses `bz` from a position `pos` to produce a string.
     */
    function parseString(bytes calldata bz, uint256 pos) internal pure returns (string memory, uint256) {
        unchecked {
            for (uint256 i = pos; i < bz.length; i++) {
                uint256 c = uint256(uint8(bz[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return (string(bz[pos:i]), i + 1);
                } else if (c == CHAR_BACKSLASH && i + 1 < bz.length) {
                    i++;
                    c = uint256(uint8(bz[i]));
                    require(
                        c == CHAR_DOUBLE_QUOTE || c == CHAR_SLASH || c == CHAR_BACKSLASH || c == CHAR_F || c == CHAR_R
                            || c == CHAR_N || c == CHAR_B || c == CHAR_T,
                        "invalid escape"
                    );
                }
            }
        }
        revert("unterminated string");
    }

    function isEscapedJSONString(string calldata s) internal pure returns (bool) {
        bytes memory bz = bytes(s);
        unchecked {
            for (uint256 i = 0; i < bz.length; i++) {
                uint256 c = uint256(uint8(bz[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return false;
                } else if (c == CHAR_BACKSLASH && i + 1 < bz.length) {
                    i++;
                    c = uint256(uint8(bz[i]));
                    if (
                        c != CHAR_DOUBLE_QUOTE && c != CHAR_SLASH && c != CHAR_BACKSLASH && c != CHAR_F && c != CHAR_R
                            && c != CHAR_N && c != CHAR_B && c != CHAR_T
                    ) {
                        return false;
                    }
                }
            }
        }
        return true;
    }

    function isEscapeNeededString(bytes memory bz) internal pure returns (bool) {
        unchecked {
            for (uint256 i = 0; i < bz.length; i++) {
                uint256 c = uint256(uint8(bz[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return true;
                }
            }
        }
        return false;
    }

    /**
     * @dev addressToHexString converts an address to a hex string.
     */
    function addressToHexString(address addr) internal pure returns (string memory) {
        uint256 localValue = uint256(uint160(addr));
        bytes memory buffer = new bytes(40);
        unchecked {
            for (int256 i = 39; i >= 0; --i) {
                buffer[uint256(i)] = HEX_DIGITS[localValue & 0xf];
                localValue >>= 4;
            }
        }
        if (localValue != 0) {
            revert("insufficient hex length");
        }
        return string(buffer);
    }

    /**
     * @dev hexStringToAddress converts a hex string to an address.
     */
    function hexStringToAddress(string memory addrHexString) internal pure returns (address, bool) {
        bytes memory addrBytes = bytes(addrHexString);
        if (addrBytes.length != 40) {
            return (address(0), false);
        }
        uint256 addr = 0;
        unchecked {
            for (uint256 i = 0; i < 40; i++) {
                uint256 c = uint256(uint8(addrBytes[i]));
                if (c >= 48 && c <= 57) {
                    addr = addr * 16 + (c - 48);
                } else if (c >= 97 && c <= 102) {
                    addr = addr * 16 + (c - 87);
                } else if (c >= 65 && c <= 70) {
                    addr = addr * 16 + (c - 55);
                } else {
                    return (address(0), false);
                }
            }
        }
        return (address(uint160(addr)), true);
    }
}
