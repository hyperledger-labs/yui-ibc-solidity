// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../../contracts/core/IBCHandler.sol";
import "../../../contracts/core/IBCClient.sol";
import "../../../contracts/core/IBCConnection.sol";
import "../../../contracts/core/IBCChannel.sol";
import "../../../contracts/core/IBCCommitment.sol";
import "../../../contracts/clients/MockClient.sol";
import "../../../contracts/proto/MockClient.sol";
import "../../../contracts/proto/Connection.sol";
import "../../../contracts/proto/Channel.sol";

import "./TestableIBCHandler.sol";
import "./MockApp.sol";

// TODO split setup code into other contracts
contract IBCTest is Test {
    TestableIBCHandler handler;
    MockClient mockClient;
    MockApp mockApp;

    string private constant mockClientType = "mock-client";
    string private constant portId = "mock";
    bytes32 private testPacketCommitment;

    function setUp() public {
        address ibcClient = address(new IBCClient());
        address ibcConnection = address(new IBCConnection());
        address ibcChannel = address(new IBCChannel());
        handler = new TestableIBCHandler(ibcClient, ibcConnection, ibcChannel);

        mockClient = new MockClient(address(handler));
        handler.registerClient(mockClientType, mockClient);

        setUpMockClient();
        setUpConnection();
        setUpChannel();
        setUpMockApp();
    }

    function setUpMockClient() internal {
        createMockClient(1);
    }

    function setUpConnection() internal {
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
        handler.setConnection("connection-0", connection);
        handler.setNextConnectionSequence(1);
    }

    function setUpChannel() internal {
        string[] memory hops = new string[](1);
        hops[0] = "connection-0";
        Channel.Data memory channel = Channel.Data({
            state: Channel.State.STATE_OPEN,
            ordering: Channel.Order.ORDER_UNORDERED,
            counterparty: ChannelCounterparty.Data({port_id: portId, channel_id: "channel-0"}),
            connection_hops: hops,
            version: "1"
        });
        handler.setChannel(portId, "channel-0", channel);
        handler.setNextChannelSequence(1);
        handler.setNextSequenceSend(portId, "channel-0", 1);
        handler.setNextSequenceRecv(portId, "channel-0", 1);
        handler.setNextSequenceAck(portId, "channel-0", 1);

        testPacketCommitment = makePacketCommitment(getPacket());
    }

    function setUpMockApp() internal {
        mockApp = new MockApp();
        handler.bindPort(portId, address(mockApp));
        handler.claimCapabilityDirectly(handler.channelCapabilityPath(portId, "channel-0"), address(mockApp));
        handler.claimCapabilityDirectly(handler.channelCapabilityPath(portId, "channel-0"), address(this));
    }

    /* test cases */

    function testConnectionOpenInit() public {
        IBCMsgs.MsgConnectionOpenInit memory msg_ = IBCMsgs.MsgConnectionOpenInit({
            clientId: "mock-client-1",
            counterparty: Counterparty.Data({
                client_id: "mock-client-1",
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: bytes("ibc")})
            }),
            delayPeriod: 0
        });
        string memory connectionId = handler.connectionOpenInit(msg_);
        assertEq(connectionId, "connection-1");
    }

    /* gas benchmarks */

    function testBenchmarkCreateMockClient() public {
        createMockClient(1);
    }

    function testBenchmarkUpdateMockClient() public {
        updateMockClient(2);
    }

    function testBenchmarkSendPacket() public {
        Packet.Data memory packet = getPacket();
        handler.sendPacket(packet);
    }

    event MockRecv(bool ok);

    function testBenchmarkRecvPacket() public {
        Packet.Data memory packet = getPacket();
        vm.expectEmit(false, false, false, true);
        emit MockRecv(true);
        handler.recvPacket(
            IBCMsgs.MsgPacketRecv({
                packet: packet,
                proof: abi.encodePacked(sha256(abi.encodePacked(testPacketCommitment))),
                proofHeight: Height.Data({revision_number: 0, revision_height: 1})
            })
        );
    }

    /* internal functions */

    function createMockClient(uint64 revision_height) internal {
        handler.createClient(
            IBCMsgs.MsgCreateClient({
                clientType: mockClientType,
                height: Height.Data({revision_number: 0, revision_height: revision_height}),
                clientStateBytes: wrapAnyMockClientState(
                    IbcLightclientsMockV1ClientState.Data({
                        latest_height: Height.Data({revision_number: 0, revision_height: revision_height})
                    })
                    ),
                consensusStateBytes: wrapAnyMockConsensusState(
                    IbcLightclientsMockV1ConsensusState.Data({timestamp: uint64(block.timestamp)})
                    )
            })
        );
    }

    function updateMockClient(uint64 next_revision_height) internal {
        handler.updateClient(
            IBCMsgs.MsgUpdateClient({
                clientId: "mock-client-0",
                clientMessage: wrapAnyMockHeader(
                    IbcLightclientsMockV1Header.Data({
                        height: Height.Data({revision_number: 0, revision_height: next_revision_height}),
                        timestamp: uint64(block.timestamp)
                    })
                    )
            })
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

    function getConnectionVersions() internal pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        string[] memory features = new string[](2);
        features[0] = "ORDER_ORDERED";
        features[1] = "ORDER_UNORDERED";
        versions[0] = Version.Data({identifier: "1", features: features});
        return versions;
    }

    function getPacket() internal pure returns (Packet.Data memory packet) {
        return Packet.Data({
            sequence: 1,
            source_port: portId,
            source_channel: "channel-0",
            destination_port: portId,
            destination_channel: "channel-0",
            data: bytes("{\"amount\": \"100\"}"),
            timeout_height: Height.Data({revision_number: 0, revision_height: 100}),
            timeout_timestamp: 0
        });
    }

    function makePacketCommitment(Packet.Data memory packet) internal pure returns (bytes32) {
        return sha256(
            abi.encodePacked(
                packet.timeout_timestamp,
                packet.timeout_height.revision_number,
                packet.timeout_height.revision_height,
                sha256(packet.data)
            )
        );
    }
}
