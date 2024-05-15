// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../../../contracts/proto/Client.sol";
import {IBCHeight} from "../../../../contracts/core/02-client/IBCHeight.sol";
import {MockClient} from "../../../../contracts/clients/mock/MockClient.sol";
import "./IBCTestHelper.t.sol";

abstract contract MockClientTestHelper is IBCTestHelper {
    using IBCHeight for Height.Data;

    string internal constant MOCK_CLIENT_TYPE = "mock-client";

    function ibcHandlerMockClient() internal returns (TestableIBCHandler, MockClient) {
        TestableIBCHandler ibcHandler = defaultIBCHandler();
        MockClient mockClient = new MockClient(address(ibcHandler));
        ibcHandler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        return (ibcHandler, mockClient);
    }

    function createMockClient(TestableIBCHandler handler, uint64 revisionHeight) internal returns (string memory) {
        return createMockClient(handler, revisionHeight, 1);
    }

    function createMockClient(TestableIBCHandler handler, uint64 revisionHeight, uint64 times)
        internal
        returns (string memory)
    {
        string memory clientId;
        for (uint64 i = 0; i < times; i++) {
            clientId = handler.createClient(msgCreateMockClient(revisionHeight));
        }
        return clientId;
    }

    function mockClientId(uint64 sequence) internal pure returns (string memory) {
        return string(abi.encodePacked(MOCK_CLIENT_TYPE, "-", Strings.toString(sequence)));
    }

    function mockClientState(uint64 revisionNumber, uint64 revisionHeight) internal pure returns (bytes memory) {
        return wrapAnyMockClientState(
            IbcLightclientsMockV1ClientState.Data({latest_height: H(revisionNumber, revisionHeight)})
        );
    }

    function mockConsensusState(uint64 timestamp) internal pure returns (bytes memory) {
        return wrapAnyMockConsensusState(IbcLightclientsMockV1ConsensusState.Data({timestamp: timestamp}));
    }

    function msgCreateMockClient(uint64 revisionHeight) internal view returns (IIBCClient.MsgCreateClient memory) {
        return msgCreateMockClient(0, revisionHeight);
    }

    function msgCreateMockClient(uint64 revisionNumber, uint64 revisionHeight)
        internal
        view
        returns (IIBCClient.MsgCreateClient memory)
    {
        return IIBCClient.MsgCreateClient({
            clientType: MOCK_CLIENT_TYPE,
            protoClientState: mockClientState(revisionNumber, revisionHeight),
            protoConsensusState: mockConsensusState(uint64(block.timestamp * 1e9))
        });
    }

    function msgUpdateMockClient(string memory clientId, uint64 nextRevisionHeight)
        internal
        view
        returns (IIBCClient.MsgUpdateClient memory)
    {
        return IIBCClient.MsgUpdateClient({
            clientId: clientId,
            protoClientMessage: wrapAnyMockHeader(mockClientHeader(nextRevisionHeight))
        });
    }

    function mockClientHeader(uint64 nextRevisionHeight)
        internal
        view
        returns (IbcLightclientsMockV1Header.Data memory)
    {
        return
            IbcLightclientsMockV1Header.Data({height: H(nextRevisionHeight), timestamp: uint64(block.timestamp * 1e9)});
    }

    function genMockProof(Height.Data memory proofHeight, bytes memory prefix, bytes memory path, bytes memory value)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(
            sha256(abi.encodePacked(proofHeight.toUint128(), sha256(prefix), sha256(path), sha256(value)))
        );
    }

    function wrapAnyMockHeader(IbcLightclientsMockV1Header.Data memory header) internal pure returns (bytes memory) {
        Any.Data memory anyHeader;
        anyHeader.type_url = "/ibc.lightclients.mock.v1.Header";
        anyHeader.value = IbcLightclientsMockV1Header.encode(header);
        return Any.encode(anyHeader);
    }

    function wrapAnyMockClientState(IbcLightclientsMockV1ClientState.Data memory clientState)
        internal
        pure
        returns (bytes memory)
    {
        Any.Data memory anyClientState;
        anyClientState.type_url = "/ibc.lightclients.mock.v1.ClientState";
        anyClientState.value = IbcLightclientsMockV1ClientState.encode(clientState);
        return Any.encode(anyClientState);
    }

    function wrapAnyMockConsensusState(IbcLightclientsMockV1ConsensusState.Data memory consensusState)
        internal
        pure
        returns (bytes memory)
    {
        Any.Data memory anyConsensusState;
        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = IbcLightclientsMockV1ConsensusState.encode(consensusState);
        return Any.encode(anyConsensusState);
    }

    function getTimestamp(ILightClient client, string memory clientId, int64 diff) internal view returns (uint64) {
        (, uint64 timestamp) = getClientLatestInfo(client, clientId);
        return uint64(int64(timestamp) + diff);
    }

    function getHeight(ILightClient client, string memory clientId, int64 diff)
        internal
        view
        returns (Height.Data memory)
    {
        (Height.Data memory latestHeight,) = getClientLatestInfo(client, clientId);
        return Height.Data({
            revision_number: latestHeight.revision_number,
            revision_height: uint64(int64(latestHeight.revision_height) + diff)
        });
    }

    function getClientLatestInfo(ILightClient client, string memory clientId)
        internal
        view
        returns (Height.Data memory, uint64)
    {
        Height.Data memory latestHeight = client.getLatestHeight(clientId);
        uint64 timestamp = client.getTimestampAtHeight(clientId, latestHeight);
        return (latestHeight, timestamp);
    }
}
