// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Height} from "../../proto/Client.sol";
import {IICS20Errors} from "./IICS20Errors.sol";

library ICS20Lib {
    bytes internal constant SUCCESSFUL_ACKNOWLEDGEMENT_JSON = bytes('{"result":"AQ=="}');
    bytes internal constant FAILED_ACKNOWLEDGEMENT_JSON = bytes('{"error":"failed"}');
    bytes32 internal constant KECCAK256_SUCCESSFUL_ACKNOWLEDGEMENT_JSON = keccak256(SUCCESSFUL_ACKNOWLEDGEMENT_JSON);

    uint256 private constant CHAR_DOUBLE_QUOTE = 0x22; // '"'
    uint256 private constant CHAR_SLASH = 0x2f; // "/"
    uint256 private constant CHAR_BACKSLASH = 0x5c;
    uint256 private constant CHAR_F = 0x66; // "f"
    uint256 private constant CHAR_R = 0x72; // "r"
    uint256 private constant CHAR_N = 0x6e; // "n"
    uint256 private constant CHAR_B = 0x62; // "b"
    uint256 private constant CHAR_T = 0x74; // "t"
    uint256 private constant CHAR_CLOSING_BRACE = 0x7d; // "}"
    uint256 private constant CHAR_M = 0x6d; // "m"

    bytes16 private constant HEX_DIGITS = "0123456789abcdef";

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

    /**
     * @dev Either `height` or `timestampNanos` must be set.
     */
    struct Timeout {
        Height.Data height;
        uint64 timestampNanos;
    }

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
     * @param escapedDenom is the denom string escaped.
     * @param amount is the amount field.
     * @param escapedSender is the sender string escaped.
     * @param escapedReceiver is the receiver string escaped.
     * @param escapedMemo is the memo string escaped.
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
     * @param escapedDenom is the denom string escaped.
     * @param amount is the amount field.
     * @param escapedSender is the sender string escaped.
     * @param escapedReceiver is the receiver string escaped.
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
     * @param bz the JSON bytes to unmarshal. It must be either of the following JSON formats. It is assumed that string fields are escaped.
     * 1. {"amount":"<uint256>","denom":"<string>","memo":"<string>","receiver":"<string>","sender":"<string>"}
     * 2. {"amount":"<uint256>","denom":"<string>","receiver":"<string>","sender":"<string>"}
     */
    function unmarshalJSON(bytes calldata bz) internal pure returns (PacketData memory) {
        PacketData memory pd;
        uint256 pos = 0;

        unchecked {
            // SAFETY: `pos` never overflow because it is always less than `bz.length`.
            if (bytes32(bz[pos:pos + 11]) != bytes32('{"amount":"')) {
                revert IICS20Errors.ICS20JSONUnexpectedBytes(pos, bytes32('{"amount":"'), bytes32(bz[pos:pos + 11]));
            }
            (pd.amount, pos) = parseUint256String(bz, pos + 11);
            if (bytes32(bz[pos:pos + 10]) != bytes32(',"denom":"')) {
                revert IICS20Errors.ICS20JSONUnexpectedBytes(pos, bytes32(',"denom":"'), bytes32(bz[pos:pos + 10]));
            }
            (pd.denom, pos) = parseString(bz, pos + 10);

            if (uint256(uint8(bz[pos + 2])) == CHAR_M) {
                if (bytes32(bz[pos:pos + 9]) != bytes32(',"memo":"')) {
                    revert IICS20Errors.ICS20JSONUnexpectedBytes(pos, bytes32(',"memo":"'), bytes32(bz[pos:pos + 9]));
                }
                (pd.memo, pos) = parseString(bz, pos + 9);
            }

            if (bytes32(bz[pos:pos + 13]) != bytes32(',"receiver":"')) {
                revert IICS20Errors.ICS20JSONUnexpectedBytes(pos, bytes32(',"receiver":"'), bytes32(bz[pos:pos + 13]));
            }
            (pd.receiver, pos) = parseString(bz, pos + 13);

            if (bytes32(bz[pos:pos + 11]) != bytes32(',"sender":"')) {
                revert IICS20Errors.ICS20JSONUnexpectedBytes(pos, bytes32(',"sender":"'), bytes32(bz[pos:pos + 11]));
            }
            (pd.sender, pos) = parseString(bz, pos + 11);

            if (pos != bz.length - 1 || uint256(uint8(bz[pos])) != CHAR_CLOSING_BRACE) {
                revert IICS20Errors.ICS20JSONClosingBraceNotFound(pos, bz[pos]);
            }
        }

        return pd;
    }

    /**
     * @dev timeout returns a Timeout struct with the given height.
     */
    function timeout(uint64 revisionNumber, uint64 revisionHeight) internal pure returns (Timeout memory) {
        return Timeout({
            height: Height.Data({revision_number: revisionNumber, revision_height: revisionHeight}),
            timestampNanos: 0
        });
    }

    /**
     * @dev timeout returns a Timeout struct with the given timestamp.
     */
    function timeout(uint64 timestampNanos) internal pure returns (Timeout memory) {
        return Timeout({height: Height.Data({revision_number: 0, revision_height: 0}), timestampNanos: timestampNanos});
    }

    /**
     * @dev parseUint256String parses `bz` from a position `pos` to produce a uint256 value.
     * The parse will stop parsing when it encounters a non-digit character.
     * @param bz the byte array to parse.
     * @param pos the position to start parsing.
     * @return ret the parsed uint256 value.
     * @return pos the new position after parsing.
     */
    function parseUint256String(bytes calldata bz, uint256 pos) internal pure returns (uint256, uint256) {
        uint256 ret = 0;
        uint256 bzLen = bz.length;
        for (; pos < bzLen; pos++) {
            uint256 c = uint256(uint8(bz[pos]));
            if (c < 48 || c > 57) {
                break;
            }
            unchecked {
                // SAFETY: we assume that the amount is uint256, so `ret` never overflows.
                ret = ret * 10 + (c - 48);
            }
        }
        if (uint256(uint8(bz[pos])) != CHAR_DOUBLE_QUOTE) {
            revert IICS20Errors.ICS20JSONStringClosingDoubleQuoteNotFound(pos, bz[pos]);
        }
        unchecked {
            // SAFETY: `pos` is always less than `bz.length`.
            return (ret, pos + 1);
        }
    }

    /**
     * @dev parseString parses `bz` from a position `pos` to produce a string.
     * @param bz the byte array to parse.
     * @param pos the position to start parsing.
     * @return parsedStr the parsed string.
     * @return position the new position after parsing.
     */
    function parseString(bytes calldata bz, uint256 pos) internal pure returns (string memory, uint256) {
        uint256 bzLen = bz.length;
        unchecked {
            // SAFETY: i + 1 <= bzLen <= type(uint256).max
            for (uint256 i = pos; i < bzLen; i++) {
                uint256 c = uint256(uint8(bz[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return (string(bz[pos:i]), i + 1);
                } else if (c == CHAR_BACKSLASH && i + 1 < bzLen) {
                    i++;
                    c = uint256(uint8(bz[i]));
                    if (
                        c != CHAR_DOUBLE_QUOTE && c != CHAR_SLASH && c != CHAR_BACKSLASH && c != CHAR_F && c != CHAR_R
                            && c != CHAR_N && c != CHAR_B && c != CHAR_T
                    ) {
                        revert IICS20Errors.ICS20JSONInvalidEscape(i, bz[i]);
                    }
                }
            }
        }
        revert IICS20Errors.ICS20JSONStringUnclosed(bz, pos);
    }

    /**
     * @dev isEscapedJSONString checks if a string is escaped JSON.
     */
    function isEscapedJSONString(string calldata s) internal pure returns (bool) {
        bytes memory bz = bytes(s);
        uint256 bzLen = bz.length;
        for (uint256 i = 0; i < bzLen; i++) {
            unchecked {
                uint256 c = uint256(uint8(bz[i]));
                if (c == CHAR_DOUBLE_QUOTE) {
                    return false;
                } else if (c == CHAR_BACKSLASH && i + 1 < bzLen) {
                    // SAFETY: i + 1 <= bzLen <= type(uint256).max
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

    /**
     * @dev isEscapeNeededString checks if a given string needs to be escaped.
     * @param bz the byte array to check.
     */
    function isEscapeNeededString(bytes memory bz) internal pure returns (bool) {
        uint256 bzLen = bz.length;
        for (uint256 i = 0; i < bzLen; i++) {
            uint256 c = uint256(uint8(bz[i]));
            if (c == CHAR_DOUBLE_QUOTE) {
                return true;
            }
        }
        return false;
    }

    /**
     * @dev addressToHexString converts an address to a hex string.
     * @param addr the address to convert.
     * @return the hex string.
     */
    function addressToHexString(address addr) internal pure returns (string memory) {
        uint256 localValue = uint256(uint160(addr));
        bytes memory buffer = new bytes(42);
        buffer[0] = "0";
        buffer[1] = "x";
        unchecked {
            // SAFETY: `i` is always greater than or equal to 1.
            for (uint256 i = 41; i >= 2; --i) {
                buffer[i] = HEX_DIGITS[localValue & 0xf];
                localValue >>= 4;
            }
        }
        return string(buffer);
    }

    /**
     * @dev hexStringToAddress converts a hex string to an address.
     * @param addrHexString the hex string to convert. It must be 42 characters long and start with "0x".
     * @return the address and a boolean indicating whether the conversion was successful.
     */
    function hexStringToAddress(string memory addrHexString) internal pure returns (address, bool) {
        bytes memory addrBytes = bytes(addrHexString);
        if (addrBytes.length != 42) {
            return (address(0), false);
        } else if (addrBytes[0] != "0" || addrBytes[1] != "x") {
            return (address(0), false);
        }
        uint256 addr = 0;
        for (uint256 i = 2; i < 42; i++) {
            uint256 c = uint256(uint8(addrBytes[i]));
            unchecked {
                // SAFETY: we assume that the address is a valid ethereum addrress, so `addr` never overflows.
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

    /**
     * @dev slice returns a slice of the original bytes from `start` to `start + length`.
     *      This is a copy from https://github.com/GNSPS/solidity-bytes-utils/blob/v0.8.0/contracts/BytesLib.sol
     */
    function slice(bytes memory _bytes, uint256 _start, uint256 _length) internal pure returns (bytes memory) {
        if (_length + 31 < _length) {
            revert IICS20Errors.ICS20BytesSliceOverflow(_length);
        } else if (_start + _length > _bytes.length) {
            revert IICS20Errors.ICS20BytesSliceOutOfBounds(_bytes.length, _start, _start + _length);
        }

        bytes memory tempBytes;

        assembly {
            switch iszero(_length)
            case 0 {
                // Get a location of some free memory and store it in tempBytes as
                // Solidity does for memory variables.
                tempBytes := mload(0x40)

                // The first word of the slice result is potentially a partial
                // word read from the original array. To read it, we calculate
                // the length of that partial word and start copying that many
                // bytes into the array. The first word we copy will start with
                // data we don't care about, but the last `lengthmod` bytes will
                // land at the beginning of the contents of the new array. When
                // we're done copying, we overwrite the full first word with
                // the actual length of the slice.
                let lengthmod := and(_length, 31)

                // The multiplication in the next line is necessary
                // because when slicing multiples of 32 bytes (lengthmod == 0)
                // the following copy loop was copying the origin's length
                // and then ending prematurely not copying everything it should.
                let mc := add(add(tempBytes, lengthmod), mul(0x20, iszero(lengthmod)))
                let end := add(mc, _length)

                for {
                    // The multiplication in the next line has the same exact purpose
                    // as the one above.
                    let cc := add(add(add(_bytes, lengthmod), mul(0x20, iszero(lengthmod))), _start)
                } lt(mc, end) {
                    mc := add(mc, 0x20)
                    cc := add(cc, 0x20)
                } { mstore(mc, mload(cc)) }

                mstore(tempBytes, _length)

                //update free-memory pointer
                //allocating the array padded to 32 bytes like the compiler does now
                mstore(0x40, and(add(mc, 31), not(31)))
            }
            //if we want a zero-length slice let's just return a zero-length array
            default {
                tempBytes := mload(0x40)
                //zero out the 32 bytes slice we are about to return
                //we need to do it because Solidity does not garbage collect
                mstore(tempBytes, 0)

                mstore(0x40, add(tempBytes, 0x20))
            }
        }

        return tempBytes;
    }

    /**
     * @dev equal returns true if two byte arrays are equal.
     */
    function equal(bytes memory a, bytes memory b) internal pure returns (bool) {
        return keccak256(a) == keccak256(b);
    }

    /**
     * @dev denomPrefix returns the prefix of the denomination.
     */
    function denomPrefix(string memory port, string memory channel) internal pure returns (bytes memory) {
        return abi.encodePacked(port, "/", channel, "/");
    }

    /**
     * @dev denom returns the denomination string.
     */
    function denom(string memory port, string memory channel, address tokenContract)
        internal
        pure
        returns (string memory)
    {
        return denom(port, channel, addressToHexString(tokenContract));
    }

    /**
     * @dev denom returns the denomination string.
     */
    function denom(string memory port, string memory channel, string memory baseDenom)
        internal
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(denomPrefix(port, channel), baseDenom));
    }
}
