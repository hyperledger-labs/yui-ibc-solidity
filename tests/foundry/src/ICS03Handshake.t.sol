// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import "./helpers/ICS03TestHelper.t.sol";

contract TestICS03Handshake is ICS03HandshakeMockClientTestHelper {
    TestableIBCHandler handler;
    TestableIBCHandler counterpartyHandler;

    function setUp() public {
        (TestableIBCHandler _handler,) = ibcHandlerMockClient();
        (TestableIBCHandler _counterpartyHandler,) = ibcHandlerMockClient();
        handler = _handler;
        counterpartyHandler = _counterpartyHandler;
    }

    function testConnOpenInit() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        {
            string memory connectionId =
                handler.connectionOpenInit(msgConnectionOpenInit(clientId, counterpartyClientId));
            assertEq(connectionId, genConnectionId(0));
            ConnectionEnd.Data memory connection =
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_INIT);
            assertEq(connection.client_id, clientId);
            matchDefaultConnectionVersions(connection.versions);
        }
        {
            IIBCConnection.MsgConnectionOpenInit memory msg_ = msgConnectionOpenInit(clientId, counterpartyClientId);
            msg_.delayPeriod = 1;
            string memory connectionId = handler.connectionOpenInit(msg_);
            assertEq(connectionId, genConnectionId(1));
            ConnectionEnd.Data memory connection =
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_INIT);
            assertEq(connection.client_id, clientId);
            assertEq(connection.delay_period, uint64(1));
            matchDefaultConnectionVersions(connection.versions);
        }
    }

    function testInvalidConnOpenInit() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        {
            // counterparty connection identifier must be empty
            IIBCConnection.MsgConnectionOpenInit memory msg_ = msgConnectionOpenInit(clientId, counterpartyClientId);
            msg_.counterparty.connection_id = genConnectionId(0);
            vm.expectRevert();
            handler.connectionOpenInit(msg_);
        }
    }

    function testConnOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        {
            (, string memory connectionId) = handshakeConnection(
                counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.TRY, H(1)
            );
            assertEq(connectionId, genConnectionId(0));
        }
        {
            (, string memory connectionId) = handshakeConnection(
                counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.TRY, H(1)
            );
            assertEq(connectionId, genConnectionId(1));
        }
    }

    function testInvalidConnOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory counterpartyConnectionId,) = handshakeConnection(
            counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.INIT, H(1)
        );
        bytes memory prefix = handler.getCommitmentPrefix();
        ConnectionEnd.Data memory counterpartyConnection = ConnectionEnd.Data({
            client_id: counterpartyClientId,
            versions: getConnectionVersions(),
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: 0,
            counterparty: Counterparty.Data({
                client_id: clientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
            })
        });
        {
            // commitment prefix is invalid
            IIBCConnection.MsgConnectionOpenTry memory msg_ = msgConnectionOpenTry(
                clientId,
                counterpartyClientId,
                counterpartyConnectionId,
                counterpartyConnection,
                bytes("ibcibc"),
                H(1),
                H(1)
            );
            vm.expectRevert();
            handler.connectionOpenTry(msg_);
        }
        {
            // proof height not found
            IIBCConnection.MsgConnectionOpenTry memory msg_ = msgConnectionOpenTry(
                clientId, counterpartyClientId, counterpartyConnectionId, counterpartyConnection, prefix, H(2), H(1)
            );
            vm.expectRevert();
            handler.connectionOpenTry(msg_);
        }
        {
            // invalid connection state
            IIBCConnection.MsgConnectionOpenTry memory msg_ = msgConnectionOpenTry(
                clientId, counterpartyClientId, counterpartyConnectionId, counterpartyConnection, prefix, H(1), H(1)
            );
            msg_.proofInit = proveConnectionState(
                H(1),
                prefix,
                genConnectionId(0),
                ConnectionEnd.Data({
                    client_id: counterpartyClientId,
                    versions: getConnectionVersions(),
                    state: ConnectionEnd.State.STATE_OPEN,
                    delay_period: 0,
                    counterparty: Counterparty.Data({
                        client_id: clientId,
                        connection_id: "",
                        prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
                    })
                })
            );
            vm.expectRevert();
            handler.connectionOpenTry(msg_);
        }
    }

    function testConnOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);

        handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.ACK, H(1)
        );
    }

    function testInvalidConnOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory connectionId, string memory counterpartyConnectionId) = handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.TRY, H(1)
        );
        bytes memory counterpartyPrefix = counterpartyHandler.getCommitmentPrefix();

        {
            // unexpected counterparty's connection state
            ConnectionEnd.State[2] memory invalidStates =
                [ConnectionEnd.State.STATE_INIT, ConnectionEnd.State.STATE_OPEN];
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            for (uint256 i = 0; i < invalidStates.length; i++) {
                counterpartyConnection.state = invalidStates[i];
                IIBCConnection.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                    connectionId,
                    counterpartyConnectionId,
                    counterpartyClientId,
                    counterpartyConnection,
                    counterpartyPrefix,
                    H(1),
                    H(1)
                );
                vm.expectRevert();
                handler.connectionOpenAck(msg_);
            }
        }
        {
            // invalid proof height
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            IIBCConnection.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                connectionId,
                counterpartyConnectionId,
                counterpartyClientId,
                counterpartyConnection,
                counterpartyPrefix,
                H(2),
                H(1)
            );
            vm.expectRevert();
            handler.connectionOpenAck(msg_);
        }
        {
            // invalid commitment prefix
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            IIBCConnection.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                connectionId,
                counterpartyConnectionId,
                counterpartyClientId,
                counterpartyConnection,
                bytes("ibcibc"),
                H(1),
                H(1)
            );
            vm.expectRevert();
            handler.connectionOpenAck(msg_);
        }
    }

    function testConnOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.CONFIRM, H(1)
        );
    }

    function testInvalidConnOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory connectionId, string memory counterpartyConnectionId) = handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.ACK, H(1)
        );

        {
            // invalid proof height
            IIBCConnection.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                counterpartyConnectionId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN),
                DEFAULT_COMMITMENT_PREFIX,
                H(2)
            );
            vm.expectRevert();
            counterpartyHandler.connectionOpenConfirm(msg_);
        }
        {
            // invalid commitment prefix
            IIBCConnection.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                counterpartyConnectionId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN),
                bytes("ibcibc"),
                H(1)
            );
            vm.expectRevert();
            counterpartyHandler.connectionOpenConfirm(msg_);
        }
        {
            // invalid connection state
            ConnectionEnd.State[2] memory invalidStates =
                [ConnectionEnd.State.STATE_INIT, ConnectionEnd.State.STATE_TRYOPEN];
            ConnectionEnd.Data memory connection =
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN);
            for (uint256 i = 0; i < invalidStates.length; i++) {
                connection.state = invalidStates[i];
                IIBCConnection.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                    counterpartyConnectionId, connectionId, connection, DEFAULT_COMMITMENT_PREFIX, H(1)
                );
                vm.expectRevert();
                counterpartyHandler.connectionOpenConfirm(msg_);
            }
        }
    }
}
