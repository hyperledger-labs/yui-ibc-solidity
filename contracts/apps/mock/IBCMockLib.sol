// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

library IBCMockLib {
    bytes public constant MOCK_PACKET_DATA = bytes("mock packet data");
    bytes public constant MOCK_FAIL_PACKET_DATA = bytes("mock failed packet data");
    bytes public constant MOCK_ASYNC_PACKET_DATA = bytes("mock async packet data");

    bytes public constant SUCCESSFUL_ACKNOWLEDGEMENT_JSON = bytes('{"result":"bW9jayBhY2tub3dsZWRnZW1lbnQ="}');
    bytes public constant FAILED_ACKNOWLEDGEMENT_JSON = bytes('{"error":"mock failed acknowledgement"}');
}
