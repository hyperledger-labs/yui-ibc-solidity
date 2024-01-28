// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../../contracts/core/02-client/IBCClient.sol";
import "../../../contracts/core/03-connection/IBCConnectionSelfStateNoValidation.sol";
import "../../../contracts/core/04-channel/IIBCChannel.sol";
import "../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import "../../../contracts/core/04-channel/IBCChannelPacketSendRecv.sol";
import "../../../contracts/core/04-channel/IBCChannelPacketTimeout.sol";
import "../../../contracts/core/04-channel/IBCChannelUpgrade.sol";
import "../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../contracts/clients/MockClient.sol";
import "../../../contracts/proto/MockClient.sol";
import "../../../contracts/proto/Connection.sol";
import "../../../contracts/proto/Channel.sol";
import "../../../contracts/apps/mock/IBCMockApp.sol";
import "./helpers/TestableIBCHandler.t.sol";
import "./helpers/IBCCommitmentTestHelper.sol";

// TODO split setup code into other contracts
contract IBCTest is Test {
    using IBCHeight for Height.Data;

    TestableIBCHandler handler;
    MockClient mockClient;
    IBCMockApp mockApp;

    string private constant MOCK_CLIENT_TYPE = "mock-client";
    string private constant MOCK_PORT_ID = "mock";
    bytes32 private testPacketCommitmentProof;

    function setUp() public {
        handler = new TestableIBCHandler(
            new IBCClient(),
            new IBCConnectionSelfStateNoValidation(),
            new IBCChannelHandshake(),
            new IBCChannelPacketSendRecv(),
            new IBCChannelPacketTimeout(),
            new IBCChannelUpgrade()
        );

        mockClient = new MockClient(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);

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
            counterparty: ChannelCounterparty.Data({port_id: MOCK_PORT_ID, channel_id: "channel-0"}),
            connection_hops: hops,
            version: "1",
            upgrade_sequence: 0
        });
        handler.setChannel(MOCK_PORT_ID, "channel-0", channel);
        handler.setNextChannelSequence(1);
        handler.setNextSequenceSend(MOCK_PORT_ID, "channel-0", 1);
        handler.setNextSequenceRecv(MOCK_PORT_ID, "channel-0", 1);
        handler.setNextSequenceAck(MOCK_PORT_ID, "channel-0", 1);

        testPacketCommitmentProof = makeMockClientPacketCommitmentProof(
            createPacket(0, 100), Height.Data({revision_number: 0, revision_height: 1})
        );
    }

    function setUpMockApp() internal {
        mockApp = new IBCMockApp(handler);
        handler.bindPort(MOCK_PORT_ID, mockApp);
        handler.setChannelCapability(MOCK_PORT_ID, "channel-0", address(mockApp));
    }

    /* test cases */

    function testToUint128(Height.Data memory height) public pure {
        Height.Data memory actual = IBCHeight.fromUint128(IBCHeight.toUint128(height));
        assert(height.eq(actual));
    }

    /* gas benchmarks */

    function testBenchmarkCreateMockClient() public {
        createMockClient(1);
    }

    function testBenchmarkUpdateMockClient() public {
        updateMockClient(2);
    }

    function testBenchmarkLCUpdateMockClient() public {
        updateLCMockClient(2);
    }

    function testBenchmarkSendPacket() public {
        handler.setChannelCapability(MOCK_PORT_ID, "channel-0", address(this));
        Packet memory packet = createPacket(0, 100);
        handler.sendPacket(
            packet.sourcePort, packet.sourceChannel, packet.timeoutHeight, packet.timeoutTimestamp, packet.data
        );
    }

    function testBenchmarkRecvPacket() public {
        Packet memory packet = createPacket(0, 100);
        handler.recvPacket(
            IIBCChannelRecvPacket.MsgPacketRecv({
                packet: packet,
                proof: abi.encodePacked(testPacketCommitmentProof),
                proofHeight: Height.Data({revision_number: 0, revision_height: 1})
            })
        );
    }

    /* internal functions */

    function createMockClient(uint64 revisionHeight) internal {
        handler.createClient(
            IIBCClient.MsgCreateClient({
                clientType: MOCK_CLIENT_TYPE,
                protoClientState: wrapAnyMockClientState(
                    IbcLightclientsMockV1ClientState.Data({
                        latest_height: Height.Data({revision_number: 0, revision_height: revisionHeight})
                    })
                    ),
                protoConsensusState: wrapAnyMockConsensusState(
                    IbcLightclientsMockV1ConsensusState.Data({timestamp: uint64(block.timestamp * 1e9)})
                    )
            })
        );
    }

    function updateMockClient(uint64 nextRevisionHeight) internal {
        handler.updateClient(
            IIBCClient.MsgUpdateClient({
                clientId: "mock-client-0",
                protoClientMessage: wrapAnyMockHeader(
                    IbcLightclientsMockV1Header.Data({
                        height: Height.Data({revision_number: 0, revision_height: nextRevisionHeight}),
                        timestamp: uint64(block.timestamp * 1e9)
                    })
                    )
            })
        );
    }

    function updateLCMockClient(uint64 nextRevisionHeight) internal {
        mockClient.updateClient(
            "mock-client-0",
            IbcLightclientsMockV1Header.Data({
                height: Height.Data({revision_number: 0, revision_height: nextRevisionHeight}),
                timestamp: uint64(block.timestamp * 1e9)
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

    function createPacket(uint64 revisionNumber, uint64 revisionHeight) internal pure returns (Packet memory packet) {
        return Packet({
            sequence: 1,
            sourcePort: MOCK_PORT_ID,
            sourceChannel: "channel-0",
            destinationPort: MOCK_PORT_ID,
            destinationChannel: "channel-0",
            data: bytes("{\"amount\": \"100\"}"),
            timeoutHeight: Height.Data({revision_number: revisionNumber, revision_height: revisionHeight}),
            timeoutTimestamp: 0
        });
    }

    function makeMockClientPacketCommitmentProof(Packet memory packet, Height.Data memory proofHeight)
        internal
        pure
        returns (bytes32)
    {
        bytes32 value = sha256(
            abi.encodePacked(
                packet.timeoutTimestamp,
                packet.timeoutHeight.revision_number,
                packet.timeoutHeight.revision_height,
                sha256(packet.data)
            )
        );
        return sha256(
            abi.encodePacked(
                proofHeight.toUint128(),
                sha256("ibc"),
                sha256(
                    IBCCommitmentTestHelper.packetCommitmentPath(
                        packet.sourcePort, packet.sourceChannel, packet.sequence
                    )
                ),
                sha256(abi.encodePacked(value))
            )
        );
    }
}
