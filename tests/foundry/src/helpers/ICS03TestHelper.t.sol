// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IIBCConnection} from "../../../../contracts/core/03-connection/IIBCConnection.sol";
import "./MockClientTestHelper.t.sol";

abstract contract ICS03TestHelper is IBCTestHelper {
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
        Version.Data[] memory defaultVersions = getConnectionVersions();
        assertEq(versions.length, defaultVersions.length);
        for (uint256 i = 0; i < versions.length; i++) {
            assertEq(versions[i].identifier, defaultVersions[i].identifier);
            assertEq(versions[i].features.length, defaultVersions[i].features.length);
            for (uint256 j = 0; j < versions[i].features.length; j++) {
                assertEq(versions[i].features[j], defaultVersions[i].features[j]);
            }
        }
    }

    function orderedIBCVersion() internal pure returns (Version.Data memory) {
        Version.Data memory version =
            Version.Data({identifier: IBCConnectionLib.IBC_VERSION_IDENTIFIER, features: new string[](1)});
        version.features[0] = IBCConnectionLib.ORDER_ORDERED;
        return version;
    }

    function unorderedIBCVersion() internal pure returns (Version.Data memory) {
        Version.Data memory version =
            Version.Data({identifier: IBCConnectionLib.IBC_VERSION_IDENTIFIER, features: new string[](1)});
        version.features[0] = IBCConnectionLib.ORDER_UNORDERED;
        return version;
    }

    function setConnection(
        TestableIBCHandler handler,
        string memory clientId,
        string memory connectionId,
        string memory counterpartyClientId,
        string memory counterpartyConnectionId,
        Version.Data memory version
    ) internal {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = version;
        handler.setConnection(
            connectionId,
            ConnectionEnd.Data({
                client_id: clientId,
                state: ConnectionEnd.State.STATE_OPEN,
                delay_period: 0,
                versions: versions,
                counterparty: Counterparty.Data({
                    client_id: counterpartyClientId,
                    connection_id: counterpartyConnectionId,
                    prefix: MerklePrefix.Data({key_prefix: DEFAULT_COMMITMENT_PREFIX})
                })
            })
        );
    }
}

abstract contract ICS03HandshakeTestHelper is ICS03TestHelper {
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
        ConnectionHandshakeStep step,
        Height.Data memory proofHeight
    ) internal returns (string memory connectionId, string memory counterpartyConnectionId) {
        Height.Data memory consensusHeight = H(1);
        connectionId = handler.connectionOpenInit(msgConnectionOpenInit(clientId, counterpartyClientId));
        if (step == ConnectionHandshakeStep.INIT) {
            return (connectionId, "");
        }
        counterpartyConnectionId = counterpartyHandler.connectionOpenTry(
            msgConnectionOpenTry(
                counterpartyClientId,
                clientId,
                connectionId,
                getConnectionEnd(handler, connectionId, ConnectionEnd.State.STATE_INIT),
                handler.getCommitmentPrefix(),
                proofHeight,
                consensusHeight
            )
        );
        if (step == ConnectionHandshakeStep.TRY) {
            return (connectionId, counterpartyConnectionId);
        }
        handler.connectionOpenAck(
            msgConnectionOpenAck(
                connectionId,
                counterpartyConnectionId,
                counterpartyClientId,
                getConnectionEnd(counterpartyHandler, counterpartyConnectionId, ConnectionEnd.State.STATE_TRYOPEN),
                counterpartyHandler.getCommitmentPrefix(),
                proofHeight,
                consensusHeight
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
                handler.getCommitmentPrefix(),
                proofHeight
            )
        );
        return (connectionId, counterpartyConnectionId);
    }

    function msgConnectionOpenInit(string memory clientId, string memory counterpartyClientId)
        internal
        pure
        virtual
        returns (IIBCConnection.MsgConnectionOpenInit memory)
    {
        return IIBCConnection.MsgConnectionOpenInit({
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
        ConnectionEnd.Data memory counterpartyConnection,
        bytes memory counterpartyPrefix,
        Height.Data memory proofHeight,
        Height.Data memory consensusHeight
    ) internal view virtual returns (IIBCConnection.MsgConnectionOpenTry memory) {
        return IIBCConnection.MsgConnectionOpenTry({
            clientId: clientId,
            delayPeriod: 0,
            counterparty: Counterparty.Data({
                connection_id: counterpartyConnectionId,
                client_id: counterpartyClientId,
                prefix: MerklePrefix.Data({key_prefix: counterpartyPrefix})
            }),
            clientStateBytes: getSelfClientState(proofHeight),
            counterpartyVersions: getConnectionVersions(),
            proofHeight: proofHeight,
            proofInit: proveConnectionState(
                proofHeight, counterpartyPrefix, counterpartyConnectionId, counterpartyConnection
            ),
            proofClient: proveClientState(
                proofHeight, counterpartyPrefix, counterpartyClientId, getSelfClientState(proofHeight)
            ),
            proofConsensus: proveConsensusState(
                proofHeight,
                counterpartyPrefix,
                counterpartyClientId,
                consensusHeight,
                getSelfConsensusState(consensusHeight)
            ),
            consensusHeight: consensusHeight,
            hostConsensusStateProof: getSelfConsensusState(consensusHeight)
        });
    }

    function msgConnectionOpenAck(
        string memory connectionId,
        string memory counterpartyConnectionId,
        string memory counterpartyClientId,
        ConnectionEnd.Data memory counterpartyConnection,
        bytes memory counterpartyPrefix,
        Height.Data memory proofHeight,
        Height.Data memory consensusHeight
    ) internal view virtual returns (IIBCConnection.MsgConnectionOpenAck memory) {
        bytes memory selfClientState = getSelfClientState(proofHeight);
        bytes memory selfConsensusState = getSelfConsensusState(consensusHeight);
        return IIBCConnection.MsgConnectionOpenAck({
            connectionId: connectionId,
            counterpartyConnectionId: counterpartyConnectionId,
            version: getConnectionVersions()[0],
            clientStateBytes: selfClientState,
            proofHeight: proofHeight,
            proofTry: proveConnectionState(
                proofHeight, counterpartyPrefix, counterpartyConnectionId, counterpartyConnection
            ),
            proofClient: proveClientState(proofHeight, counterpartyPrefix, counterpartyClientId, selfClientState),
            proofConsensus: proveConsensusState(
                proofHeight, counterpartyPrefix, counterpartyClientId, H(1), selfConsensusState
            ),
            consensusHeight: consensusHeight,
            hostConsensusStateProof: selfConsensusState
        });
    }

    function msgConnectionOpenConfirm(
        string memory connectionId,
        string memory counterpartyConnectionId,
        ConnectionEnd.Data memory counterpartyConnection,
        bytes memory counterpartyPrefix,
        Height.Data memory proofHeight
    ) internal view virtual returns (IIBCConnection.MsgConnectionOpenConfirm memory) {
        return IIBCConnection.MsgConnectionOpenConfirm({
            connectionId: connectionId,
            proofHeight: proofHeight,
            proofAck: proveConnectionState(
                proofHeight, counterpartyPrefix, counterpartyConnectionId, counterpartyConnection
            )
        });
    }

    function getSelfClientState(Height.Data memory height) internal view virtual returns (bytes memory);

    function getSelfConsensusState(Height.Data memory height) internal view virtual returns (bytes memory);

    function proveConnectionState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory connectionId,
        ConnectionEnd.Data memory connection
    ) internal view virtual returns (bytes memory);

    function proveClientState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory clientId,
        bytes memory clientState
    ) internal view virtual returns (bytes memory);

    function proveConsensusState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory clientId,
        Height.Data memory consensusHeight,
        bytes memory consensusState
    ) internal view virtual returns (bytes memory);
}

abstract contract ICS03HandshakeMockClientTestHelper is ICS03HandshakeTestHelper, MockClientTestHelper {
    function getSelfClientState(Height.Data memory height) internal view virtual override returns (bytes memory) {
        return mockClientState(height.revision_number, height.revision_height);
    }

    function getSelfConsensusState(Height.Data memory) internal view virtual override returns (bytes memory) {
        return mockConsensusState(uint64(getBlockTimestampNano()));
    }

    function proveClientState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory clientId,
        bytes memory clientState
    ) internal view virtual override returns (bytes memory) {
        return genMockProof(proofHeight, prefix, IBCCommitment.clientStatePath(clientId), clientState);
    }

    function proveConsensusState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory clientId,
        Height.Data memory consensusHeight,
        bytes memory consensusState
    ) internal view virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight,
            prefix,
            IBCCommitment.consensusStatePath(clientId, consensusHeight.revision_number, consensusHeight.revision_height),
            consensusState
        );
    }

    function proveConnectionState(
        Height.Data memory proofHeight,
        bytes memory prefix,
        string memory connectionId,
        ConnectionEnd.Data memory connection
    ) internal view virtual override returns (bytes memory) {
        return genMockProof(
            proofHeight, prefix, IBCCommitment.connectionPath(connectionId), ConnectionEnd.encode(connection)
        );
    }
}
