// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../../contracts/core/IBCHost.sol";
import "../../../contracts/core/IBCHandler.sol";
import "../../../contracts/core/IBCIdentifier.sol";
import "../../../contracts/core/MockClient.sol";
import "../../../contracts/core/types/MockClient.sol";
import "../../../contracts/core/types/Connection.sol";
import "../../../contracts/core/types/Channel.sol";

// TODO split setup code into other contracts
contract IBCTest {
    IBCHost host;
    IBCHandler handler;
    MockClient mockClient;

    string private constant mockClientType = "mock-client";
    string private constant portId = "transfer";

    function setUp() public {
        mockClient = new MockClient();
        host = new IBCHost();
        handler = new IBCHandler(host);
        host.setIBCModule(address(handler));
        handler.registerClient(mockClientType, mockClient);

        setUpMockClient();
        setUpConnection();
        setUpChannel();
    }

    function setUpMockClient() internal {
        handler.createClient(IBCMsgs.MsgCreateClient({
            clientType: mockClientType,
            height: Height.Data({revision_number: 0, revision_height: 1}),
            clientStateBytes: wrapMockClientState(IbcLightclientsMockV1ClientState.Data({latest_height: Height.Data({revision_number: 0, revision_height: 1})})),
            consensusStateBytes: wrapMockConsensusState(IbcLightclientsMockV1ConsensusState.Data({timestamp: uint64(block.timestamp)}))
        }));
    }

    function wrapMockClientState(IbcLightclientsMockV1ClientState.Data memory clientState) internal returns (bytes memory) {
        Any.Data memory anyClientState;
        anyClientState.type_url = "/ibc.lightclients.mock.v1.ClientState";
        anyClientState.value = IbcLightclientsMockV1ClientState.encode(clientState);
        return Any.encode(anyClientState);
    }

    function wrapMockConsensusState(IbcLightclientsMockV1ConsensusState.Data memory consensusState) internal returns (bytes memory) {
        Any.Data memory anyConsensusState;
        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = IbcLightclientsMockV1ConsensusState.encode(consensusState);
        return Any.encode(anyConsensusState);
    }

    function setUpConnection() internal {
        address prev = host.getIBCModule();
        host.setIBCModule(address(this));
        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: "mock-client-0",
            versions: getConnectionVersions(),
            state: ConnectionEnd.State.STATE_OPEN,
            delay_period: 0,
            counterparty: Counterparty.Data({
                client_id: "mock-client-0",
                connection_id: "connection-0",
                prefix: MerklePrefix.Data({key_prefix: bytes("ibc")})
            })
        });
        host.setConnection("connection-0", connection);
        host.setIBCModule(prev);
    }

    function getConnectionVersions() internal pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        string[] memory features = new string[](2);
        features[0] = "ORDER_ORDERED";
        features[1] = "ORDER_UNORDERED";
        versions[0] = Version.Data({
            identifier: "1",
            features: features
        });
        return versions;
    }

    function setUpChannel() internal {
        address prev = host.getIBCModule();
        host.setIBCModule(address(this));
        string[] memory hops = new string[](1);
        hops[0] = "connection-0";
        Channel.Data memory channel = Channel.Data({
            state: Channel.State.STATE_OPEN,
            ordering: Channel.Order.ORDER_UNORDERED,
            counterparty: ChannelCounterparty.Data({
                port_id: portId,
                channel_id: "channel-0"
            }),
            connection_hops: hops,
            version: "1"
        });
        host.setChannel(portId, "channel-0", channel);
        host.claimCapability(IBCIdentifier.channelCapabilityPath(portId, "channel-0"), address(this));
        host.setNextSequenceSend(portId, "channel-0", 1);
        host.setNextSequenceRecv(portId, "channel-0", 1);
        host.setNextSequenceAck(portId, "channel-0", 1);
        host.setIBCModule(prev);
    }

    function testClientRegistration() public view {
        (address client, bool found) = host.getClientImpl(mockClientType);
        assert(found);
        assert(address(mockClient) == client);
    }

    function testSendPacket() public {
        Packet.Data memory packet = Packet.Data({
            sequence: 1,
            source_port: portId,
            source_channel: "channel-0",
            destination_port: portId,
            destination_channel: "channel-0",
            data: bytes("{\"amount\": \"100\"}"),
            timeout_height: Height.Data({revision_number: 0, revision_height: 100}),
            timeout_timestamp: 0
        });
        handler.sendPacket(packet);
    }
}
