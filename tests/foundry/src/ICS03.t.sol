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
        versions[0] = IBCConnectionLib.defaultIBCVersion();
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
            }),
            version: getConnectionVersions()[0]
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

contract TestICS03Handshake is TestIBCBase, TestMockClientHelper, TestICS03Helper {
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

contract TestICS03Version is Test, TestICS03Helper {
    function testIsSupportedVersion() public {
        Version.Data[] memory versions = getConnectionVersions();
        assertTrue(IBCConnectionLib.isSupportedVersion(versions, versions[0]));
        Version.Data memory version = Version.Data({identifier: "", features: new string[](0)});
        assertFalse(IBCConnectionLib.isSupportedVersion(versions, version));
        version = Version.Data({identifier: "1", features: new string[](1)});
        version.features[0] = "ORDER_DAG";
        assertFalse(IBCConnectionLib.isSupportedVersion(versions, version));
    }

    function testFindSupportedVersion() public {
        // "valid supported version"
        {
            Version.Data[] memory versions = getConnectionVersions();
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, versions[0].identifier);
        }
        // "empty (invalid) version"
        {
            Version.Data[] memory versions = getConnectionVersions();
            (, bool found) = IBCConnectionLib.findSupportedVersion(
                Version.Data({identifier: "", features: new string[](0)}), versions
            );
            assertFalse(found);
        }
        // "empty supported versions"
        {
            Version.Data[] memory versions = new Version.Data[](0);
            (, bool found) = IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertFalse(found);
        }
        // "desired version is last"
        {
            Version.Data[] memory versions = new Version.Data[](4);
            versions[0] = Version.Data({identifier: "1.1", features: new string[](0)});
            versions[1] = Version.Data({identifier: "2", features: new string[](1)});
            versions[1].features[0] = "ORDER_UNORDERED";
            versions[2] = Version.Data({identifier: "3", features: new string[](0)});
            versions[3] = IBCConnectionLib.defaultIBCVersion();
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, versions[3].identifier);
        }
        // "desired version identifier with different feature set"
        {
            Version.Data[] memory versions = new Version.Data[](1);
            versions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            versions[0].features[0] = "ORDER_DAG";
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "version not supported"
        {
            Version.Data[] memory versions = new Version.Data[](1);
            versions[0] = Version.Data({identifier: "2", features: new string[](1)});
            versions[0].features[0] = "ORDER_DAG";
            (, bool found) = IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertFalse(found);
        }
    }

    function testPickVersion() public {
        // "valid default ibc version"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = getConnectionVersions();
            Version.Data memory v = IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "valid version in counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](3);
            counterpartyVersions[0] = Version.Data({identifier: "version1", features: new string[](0)});
            counterpartyVersions[1] = Version.Data({identifier: "2.0.0", features: new string[](1)});
            counterpartyVersions[1].features[0] = "ORDER_UNORDERED-ZK";
            counterpartyVersions[2] = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory v = IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "valid identifier match but empty feature set not allowed"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](2);
            counterpartyVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](3)});
            counterpartyVersions[0].features[0] = "DAG";
            counterpartyVersions[0].features[1] = "ORDERED-ZK";
            counterpartyVersions[0].features[2] = "UNORDERED-zk";
            counterpartyVersions[1] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](0)});
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "empty counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](0);
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "non-matching counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](2);
            counterpartyVersions[0] = Version.Data({identifier: "2.0.0", features: new string[](0)});
            counterpartyVersions[1] = Version.Data({identifier: "", features: new string[](0)});
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "non-matching counterparty versions (uses ordered channels only) contained in supported versions (uses unordered channels only)"
        {
            Version.Data[] memory supportedVersions = new Version.Data[](1);
            supportedVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            supportedVersions[0].features[0] = "ORDER_UNORDERED";
            Version.Data[] memory counterpartyVersions = new Version.Data[](1);
            counterpartyVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            counterpartyVersions[0].features[0] = "ORDER_ORDERED";
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
    }

    function testVerifyProposedVersion() public {
        // "entire feature set supported"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](3)});
            supportedVersion.features[0] = "ORDER_ORDERED";
            supportedVersion.features[1] = "ORDER_UNORDERED";
            supportedVersion.features[2] = "ORDER_DAG";
            assertTrue(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "empty feature sets not supported"
        {
            Version.Data memory proposedVersion = Version.Data({identifier: "1", features: new string[](0)});
            Version.Data memory supportedVersion = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "one feature missing"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](2)});
            supportedVersion.features[0] = "ORDER_UNORDERED";
            supportedVersion.features[1] = "ORDER_DAG";
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "both features missing"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](1)});
            supportedVersion.features[0] = "ORDER_DAG";
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "identifiers do not match"
        {
            Version.Data memory proposedVersion = Version.Data({identifier: "2", features: new string[](2)});
            proposedVersion.features[0] = "ORDER_UNORDERED";
            proposedVersion.features[1] = "ORDER_ORDERED";
            Version.Data memory supportedVersion = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
    }

    function testVerifySupportedFeature() public {
        // "check ORDERED supported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertTrue(IBCConnectionLib.verifySupportedFeature(version, "ORDER_ORDERED"));
        }
        // "check UNORDERED supported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertTrue(IBCConnectionLib.verifySupportedFeature(version, "ORDER_UNORDERED"));
        }
        // "check DAG unsupported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifySupportedFeature(version, "ORDER_DAG"));
        }
        // "check empty feature set returns false"
        {
            Version.Data memory version = Version.Data({identifier: "1", features: new string[](0)});
            assertFalse(IBCConnectionLib.verifySupportedFeature(version, "ORDER_ORDERED"));
        }
    }
}
