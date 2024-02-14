// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../../contracts/core/02-client/IBCClient.sol";
import "../../../contracts/core/03-connection/IBCConnectionSelfStateNoValidation.sol";
import "../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import "../../../contracts/core/04-channel/IBCChannelPacketSendRecv.sol";
import "../../../contracts/core/04-channel/IBCChannelPacketTimeout.sol";
import "../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../contracts/proto/MockClient.sol";
import "../../../contracts/proto/Connection.sol";
import "../../../contracts/proto/Channel.sol";
import "../../../contracts/apps/mock/IBCMockApp.sol";
import "../../../contracts/clients/MockClient.sol";
import "./helpers/TestableIBCHandler.t.sol";
import "./helpers/IBCTestHelper.t.sol";
import "./helpers/MockClientTestHelper.t.sol";

contract TestICS02 is Test, MockClientTestHelper {
    function testRegisterClient() public {
        TestableIBCHandler handler = defaultIBCHandler();
        MockClient mockClient = new MockClient(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        handler.registerClient("test", mockClient);
    }

    function testRegisterClientDuplicatedClientType() public {
        TestableIBCHandler handler = defaultIBCHandler();
        MockClient mockClient = new MockClient(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        vm.expectRevert("clientType already exists");
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
    }

    function testRegisterClientInvalidClientType() public {
        TestableIBCHandler handler = defaultIBCHandler();
        vm.expectRevert("invalid client address");
        handler.registerClient(MOCK_CLIENT_TYPE, ILightClient(address(0)));

        MockClient mockClient = new MockClient(address(handler));
        vm.expectRevert("invalid clientType");
        handler.registerClient("", mockClient);

        vm.expectRevert("invalid clientType");
        handler.registerClient("-mock", mockClient);

        vm.expectRevert("invalid clientType");
        handler.registerClient("mock-", mockClient);
    }

    function testCreateClient() public {
        (TestableIBCHandler handler, MockClient mockClient) = ibcHandlerMockClient();
        {
            string memory clientId = handler.createClient(msgCreateMockClient(1));
            assertEq(clientId, mockClientId(0));
            assertEq(handler.getClientType(clientId), MOCK_CLIENT_TYPE);
            assertEq(handler.getClient(clientId), address(mockClient));
            assertFalse(handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId)) == bytes32(0));
            assertFalse(handler.getCommitment(IBCCommitment.consensusStateCommitmentKey(clientId, 0, 1)) == bytes32(0));
        }
        {
            string memory clientId = handler.createClient(msgCreateMockClient(100));
            assertEq(clientId, mockClientId(1));
            assertEq(handler.getClientType(clientId), MOCK_CLIENT_TYPE);
            assertEq(handler.getClient(clientId), address(mockClient));
            assertFalse(handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId)) == bytes32(0));
            assertFalse(
                handler.getCommitment(IBCCommitment.consensusStateCommitmentKey(clientId, 0, 100)) == bytes32(0)
            );
        }
    }

    function testInvalidCreateClient() public {
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.clientType = "";
            vm.expectRevert("unregistered client type");
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.clientType = "06-solomachine";
            vm.expectRevert("unregistered client type");
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.protoClientState = abi.encodePacked(msg_.protoClientState, hex"00");
            vm.expectRevert();
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.protoConsensusState = abi.encodePacked(msg_.protoConsensusState, hex"00");
            vm.expectRevert();
            handler.createClient(msg_);
        }
    }

    function testUpdateClient() public {
        bytes32 prevClientStateCommitment;
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        string memory clientId = handler.createClient(msgCreateMockClient(1));
        prevClientStateCommitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));

        {
            handler.updateClient(msgUpdateMockClient(clientId, 2));
            bytes32 commitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));
            assertTrue(
                commitment != prevClientStateCommitment && commitment != bytes32(0), "commitment should be updated"
            );
            prevClientStateCommitment = commitment;
        }
        {
            handler.updateClient(msgUpdateMockClient(clientId, 3));
            bytes32 commitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));
            assertTrue(
                commitment != prevClientStateCommitment && commitment != bytes32(0), "commitment should be updated"
            );
            prevClientStateCommitment = commitment;
        }
    }

    function testInvalidUpdateClient() public {
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        string memory clientId = handler.createClient(msgCreateMockClient(1));
        assertEq(clientId, mockClientId(0));
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.clientId = "";
            vm.expectRevert();
            handler.updateClient(msg_);
        }
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.clientId = mockClientId(1);
            vm.expectRevert();
            handler.updateClient(msg_);
        }
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.protoClientMessage = abi.encodePacked(msg_.protoClientMessage, hex"00");
            vm.expectRevert();
            handler.updateClient(msg_);
        }
    }
}
