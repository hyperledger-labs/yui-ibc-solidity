// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./ICS02.t.sol";

abstract contract TestICS03Helper is TestIBCBase, TestMockClientHelper {
    function genConnectionId(uint64 sequence) internal pure returns (string memory) {
        return string(abi.encodePacked("connection-", Strings.toString(sequence)));
    }

    function getConnectionEnd(TestableIBCHandler handler, string memory connectionId, ConnectionEnd.State expectedState)
        internal
        returns (ConnectionEnd.Data memory)
    {
        (ConnectionEnd.Data memory connection, bool ok) = handler.getConnection(connectionId);
        assertTrue(ok);
        assertTrue(connection.state == expectedState);
        return connection;
    }

    function getConnectionVersions() internal pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        string[] memory features = new string[](2);
        features[0] = "ORDER_ORDERED";
        features[1] = "ORDER_UNORDERED";
        versions[0] = Version.Data({identifier: "1", features: features});
        return versions;
    }

    function matchDefaultConnectionVersions(Version.Data[] memory versions) internal {
        assertEq(versions.length, 1);
        assertEq(versions[0].identifier, "1");
        assertEq(versions[0].features.length, 2);
        assertEq(versions[0].features[0], "ORDER_ORDERED");
        assertEq(versions[0].features[1], "ORDER_UNORDERED");
    }

    function msgConnectionOpenInit(string memory clientId, string memory counterpartyClientId)
        internal
        view
        returns (IBCMsgs.MsgConnectionOpenInit memory)
    {
        return IBCMsgs.MsgConnectionOpenInit({
            clientId: clientId,
            delayPeriod: 0,
            counterparty: Counterparty.Data({
                client_id: counterpartyClientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
            })
        });
    }

    function msgConnectionOpenTry(
        string memory clientId,
        string memory counterpartyClientId,
        string memory counterpartyConnectionId,
        bytes memory counterpartyPrefix
    ) internal view returns (IBCMsgs.MsgConnectionOpenTry memory) {
        return IBCMsgs.MsgConnectionOpenTry({
            clientId: clientId,
            delayPeriod: 0,
            counterparty: Counterparty.Data({
                connection_id: counterpartyConnectionId,
                client_id: counterpartyClientId,
                prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
            }),
            clientStateBytes: mockClientState(0, 1),
            counterpartyVersions: getConnectionVersions(),
            proofHeight: Height.Data({revision_number: 0, revision_height: 1}),
            proofInit: genMockConnectionStateProof(
                counterpartyPrefix,
                counterpartyConnectionId,
                ConnectionEnd.Data({
                    client_id: counterpartyClientId,
                    versions: getConnectionVersions(),
                    state: ConnectionEnd.State.STATE_INIT,
                    delay_period: 0,
                    counterparty: Counterparty.Data({
                        client_id: clientId,
                        connection_id: "",
                        prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
                    })
                })
                ),
            proofClient: genMockClientStateProof(counterpartyClientId, 0, 1),
            proofConsensus: genMockConsensusStateProof(counterpartyClientId, 0, 1, uint64(block.timestamp * 1e9)),
            consensusHeight: Height.Data({revision_number: 0, revision_height: 1})
        });
    }

    function msgConnectionOpenTry(
        string memory clientId,
        string memory counterpartyClientId,
        string memory counterpartyConnectionId
    ) internal view returns (IBCMsgs.MsgConnectionOpenTry memory) {
        return msgConnectionOpenTry(clientId, counterpartyClientId, counterpartyConnectionId, DEFAULT_COMMITMENT_PREFIX);
    }

    function msgConnectionOpenAck(
        string memory connectionId,
        string memory counterpartyConnectionId,
        string memory counterpartyClientId,
        ConnectionEnd.Data memory counterpartyConnection
    ) internal view returns (IBCMsgs.MsgConnectionOpenAck memory) {
        return IBCMsgs.MsgConnectionOpenAck({
            connectionId: connectionId,
            counterpartyConnectionId: counterpartyConnectionId,
            version: getConnectionVersions()[0],
            clientStateBytes: mockClientState(0, 1),
            proofHeight: Height.Data({revision_number: 0, revision_height: 1}),
            proofTry: genMockConnectionStateProof(counterpartyConnectionId, counterpartyConnection),
            proofClient: genMockClientStateProof(counterpartyClientId, 0, 1),
            proofConsensus: genMockConsensusStateProof(counterpartyClientId, 0, 1, uint64(block.timestamp * 1e9)),
            consensusHeight: Height.Data({revision_number: 0, revision_height: 1})
        });
    }

    function msgConnectionOpenAck(
        string memory connectionId,
        string memory counterpartyConnectionId,
        string memory counterpartyClientId,
        ConnectionEnd.Data memory counterpartyConnection,
        bytes memory prefix
    ) internal view returns (IBCMsgs.MsgConnectionOpenAck memory) {
        return IBCMsgs.MsgConnectionOpenAck({
            connectionId: connectionId,
            counterpartyConnectionId: counterpartyConnectionId,
            version: getConnectionVersions()[0],
            clientStateBytes: mockClientState(0, 1),
            proofHeight: Height.Data({revision_number: 0, revision_height: 1}),
            proofTry: genMockConnectionStateProof(prefix, counterpartyConnectionId, counterpartyConnection),
            proofClient: genMockClientStateProof(counterpartyClientId, 0, 1),
            proofConsensus: genMockConsensusStateProof(counterpartyClientId, 0, 1, uint64(block.timestamp * 1e9)),
            consensusHeight: Height.Data({revision_number: 0, revision_height: 1})
        });
    }

    function msgConnectionOpenConfirm(
        string memory connectionId,
        string memory counterpartyConnectionId,
        ConnectionEnd.Data memory counterpartyConnection,
        bytes memory prefix
    ) internal view returns (IBCMsgs.MsgConnectionOpenConfirm memory) {
        return IBCMsgs.MsgConnectionOpenConfirm({
            connectionId: connectionId,
            proofHeight: Height.Data({revision_number: 0, revision_height: 1}),
            proofAck: genMockConnectionStateProof(prefix, counterpartyConnectionId, counterpartyConnection)
        });
    }

    enum ConnectionHandshakeStep {
        INIT,
        TRY,
        ACK,
        CONFIRM
    }

    function handshakeConnection(
        TestableIBCHandler handler,
        TestableIBCHandler counterpartyHandler,
        string memory clientId,
        string memory counterpartyClientId,
        ConnectionHandshakeStep step
    ) internal returns (string memory connectionId, string memory counterpartyConnectionId) {
        connectionId = handler.connectionOpenInit(msgConnectionOpenInit(clientId, counterpartyClientId));
        if (step == ConnectionHandshakeStep.INIT) {
            return (connectionId, "");
        }
        counterpartyConnectionId =
            counterpartyHandler.connectionOpenTry(msgConnectionOpenTry(counterpartyClientId, clientId, connectionId));
        if (step == ConnectionHandshakeStep.TRY) {
            return (connectionId, counterpartyConnectionId);
        }
        handler.connectionOpenAck(
            msgConnectionOpenAck(
                connectionId,
                counterpartyConnectionId,
                counterpartyClientId,
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN)
            )
        );
        if (step == ConnectionHandshakeStep.ACK) {
            return (connectionId, counterpartyConnectionId);
        }
        counterpartyHandler.connectionOpenConfirm(
            msgConnectionOpenConfirm(
                counterpartyConnectionId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN),
                DEFAULT_COMMITMENT_PREFIX
            )
        );
        return (connectionId, counterpartyConnectionId);
    }
}

contract TestICS03 is TestIBCBase, TestMockClientHelper, TestICS03Helper {
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
            IBCMsgs.MsgConnectionOpenInit memory msg_ = msgConnectionOpenInit(clientId, counterpartyClientId);
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
            IBCMsgs.MsgConnectionOpenInit memory msg_ = msgConnectionOpenInit(clientId, counterpartyClientId);
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
                counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.TRY
            );
            assertEq(connectionId, genConnectionId(0));
        }
        {
            (, string memory connectionId) = handshakeConnection(
                counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.TRY
            );
            assertEq(connectionId, genConnectionId(1));
        }
    }

    function testInvalidConnOpenTry() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory counterpartyConnectionId,) = handshakeConnection(
            counterpartyHandler, handler, counterpartyClientId, clientId, ConnectionHandshakeStep.INIT
        );
        {
            // commitment prefix is invalid
            IBCMsgs.MsgConnectionOpenTry memory msg_ =
                msgConnectionOpenTry(clientId, counterpartyClientId, counterpartyConnectionId);
            msg_.proofInit = genMockConnectionStateProof(
                bytes("ibcibc"),
                genConnectionId(0),
                ConnectionEnd.Data({
                    client_id: counterpartyClientId,
                    versions: getConnectionVersions(),
                    state: ConnectionEnd.State.STATE_INIT,
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
        {
            // proof height not found
            IBCMsgs.MsgConnectionOpenTry memory msg_ =
                msgConnectionOpenTry(clientId, counterpartyClientId, counterpartyConnectionId);
            msg_.proofHeight = Height.Data({revision_number: 0, revision_height: 2});
            vm.expectRevert();
            handler.connectionOpenTry(msg_);
        }
        {
            // invalid connection state
            IBCMsgs.MsgConnectionOpenTry memory msg_ =
                msgConnectionOpenTry(clientId, counterpartyClientId, counterpartyConnectionId);
            msg_.proofInit = genMockConnectionStateProof(
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

        // TODO test for incompatible connection versions
        // blocked by https://github.com/hyperledger-labs/yui-ibc-solidity/issues/25
    }

    function testConnOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);

        handshakeConnection(handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.ACK);
    }

    function testInvalidConnOpenAck() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory connectionId, string memory counterpartyConnectionId) = handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.TRY
        );

        {
            // unexpected counterparty's connection state
            ConnectionEnd.State[2] memory invalidStates =
                [ConnectionEnd.State.STATE_INIT, ConnectionEnd.State.STATE_OPEN];
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            for (uint256 i = 0; i < invalidStates.length; i++) {
                counterpartyConnection.state = invalidStates[i];
                IBCMsgs.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                    connectionId, counterpartyConnectionId, counterpartyClientId, counterpartyConnection
                );
                vm.expectRevert();
                handler.connectionOpenAck(msg_);
            }
        }
        {
            // invalid proof height
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            IBCMsgs.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                connectionId, counterpartyConnectionId, counterpartyClientId, counterpartyConnection
            );
            msg_.proofHeight = Height.Data({revision_number: 0, revision_height: 2});
            vm.expectRevert();
            handler.connectionOpenAck(msg_);
        }
        {
            // invalid commitment prefix
            ConnectionEnd.Data memory counterpartyConnection =
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN);
            IBCMsgs.MsgConnectionOpenAck memory msg_ = msgConnectionOpenAck(
                connectionId, counterpartyConnectionId, counterpartyClientId, counterpartyConnection, bytes("ibcibc")
            );
            vm.expectRevert();
            handler.connectionOpenAck(msg_);
        }
    }

    function testConnOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.CONFIRM
        );
    }

    function testInvalidConnOpenConfirm() public {
        string memory clientId = createMockClient(handler, 1);
        string memory counterpartyClientId = createMockClient(counterpartyHandler, 1, 2);
        (string memory connectionId, string memory counterpartyConnectionId) = handshakeConnection(
            handler, counterpartyHandler, clientId, counterpartyClientId, ConnectionHandshakeStep.ACK
        );

        {
            // invalid proof height
            IBCMsgs.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                counterpartyConnectionId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN),
                DEFAULT_COMMITMENT_PREFIX
            );
            msg_.proofHeight = Height.Data({revision_number: 0, revision_height: 2});
            vm.expectRevert();
            counterpartyHandler.connectionOpenConfirm(msg_);
        }
        {
            // invalid commitment prefix
            IBCMsgs.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                counterpartyConnectionId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_OPEN),
                bytes("ibcibc")
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
                IBCMsgs.MsgConnectionOpenConfirm memory msg_ = msgConnectionOpenConfirm(
                    counterpartyConnectionId, connectionId, connection, DEFAULT_COMMITMENT_PREFIX
                );
                vm.expectRevert();
                counterpartyHandler.connectionOpenConfirm(msg_);
            }
        }
    }
}
