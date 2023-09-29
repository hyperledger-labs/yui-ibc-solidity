// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";

library ICS20Packet {
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

    function marshalUnsafeJSON(PacketData memory data) internal pure returns (bytes memory) {
        return abi.encodePacked(
            '{"amount":"',
            Strings.toString(data.amount),
            '","denom":"',
            data.denom,
            '","memo":"',
            data.memo,
            '","receiver":"',
            data.receiver,
            '","sender":"',
            data.sender,
            '"}'
        );
    }

    function unmarshalJSON(bytes memory bz) internal pure returns (PacketData memory) {
        PacketData memory pd;
        uint256 pos = 0;

        (pd.amount, pos) = parseUint256String(bz, consumeString(bz, pos, '{"amount":"'));
        (pd.denom, pos) = parseString(bz, consumeString(bz, pos, ',"denom":"'));
        if (uint256(uint8(bz[pos + 2])) == CHAR_M) {
            (pd.memo, pos) = parseString(bz, consumeString(bz, pos, ',"memo":"'));
        }
        (pd.receiver, pos) = parseString(bz, consumeString(bz, pos, ',"receiver":"'));
        (pd.sender, pos) = parseString(bz, consumeString(bz, pos, ',"sender":"'));
        require(pos == bz.length - 1 && uint256(uint8(bz[pos])) == CHAR_CLOSING_BRACE, "closing brace");

        return pd;
    }

    function consumeString(bytes memory self, uint256 offset, bytes memory sub) internal pure returns (uint256) {
        unchecked {
            uint256 len = sub.length;
            require(offset + len <= self.length);

            bytes memory ret = new bytes(len);
            uint256 dest;
            uint256 src;

            assembly {
                dest := add(ret, 32)
                src := add(add(self, 32), offset)
            }
            memcpy(dest, src, len);
            require(bytes32(ret) == bytes32(sub), "mismatch");
            return offset + len;
        }
    }

    function parseUint256String(bytes memory s, uint256 pos) internal pure returns (uint256, uint256) {
        uint256 ret = 0;
        unchecked {
            for (; pos < s.length; pos++) {
                uint256 c = uint256(uint8(s[pos]));
                if (c < 48 || c > 57) {
                    break;
                }
                ret = ret * 10 + (c - 48);
            }
            require(pos < s.length && uint256(uint8(s[pos])) == CHAR_DOUBLE_QUOTE, "unterminated string");
            return (ret, pos + 1);
        }
    }

    function parseString(bytes memory s, uint256 pos) internal pure returns (string memory, uint256) {
        unchecked {
            for (uint256 i = pos; i < s.length; i++) {
                uint256 c = uint256(uint8(s[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return (string(substring(s, pos, i - pos)), i + 1);
                } else if (c == CHAR_BACKSLASH && i + 1 < s.length) {
                    i++;
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

    function parseStringWithoutEscape(bytes memory s, uint256 pos) internal pure returns (string memory, uint256) {
        unchecked {
            for (uint256 i = pos; i < s.length; i++) {
                if (uint256(uint8(s[i])) == CHAR_DOUBLE_QUOTE) {
                    return (string(substring(s, pos, i - pos)), i + 1);
                }
            }
        }
        revert("unterminated string");
    }

    function substring(bytes memory self, uint256 offset, uint256 len) internal pure returns (bytes memory) {
        require(offset + len <= self.length);

        bytes memory ret = new bytes(len);
        uint256 dest;
        uint256 src;

        assembly {
            dest := add(ret, 32)
            src := add(add(self, 32), offset)
        }
        memcpy(dest, src, len);

        return ret;
    }

    function memcpy(uint256 dest, uint256 src, uint256 len) private pure {
        // Copy word-length chunks while possible
        for (; len >= 32; len -= 32) {
            assembly {
                mstore(dest, mload(src))
            }
            dest += 32;
            src += 32;
        }

        // Copy remaining bytes
        unchecked {
            uint256 mask = (256 ** (32 - len)) - 1;
            assembly {
                let srcpart := and(mload(src), not(mask))
                let destpart := and(mload(dest), mask)
                mstore(dest, or(destpart, srcpart))
            }
        }
    }
}
