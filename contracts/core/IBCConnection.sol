// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";
import "../proto/Connection.sol";
import "./IBCClient.sol";
import "./IBCMsgs.sol";
import "./IBCHost.sol";

library IBCConnection {
    string constant commitmentPrefix = "ibc";

    /* Public functions */

    // ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
    // is returned.
    function connectionOpenInit(IBCHost host, IBCMsgs.MsgConnectionOpenInit memory msg_)
        public
        returns (string memory)
    {
        host.onlyIBCModule();
        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: getVersions(),
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });
        string memory connectionId = host.generateConnectionIdentifier();
        host.setConnection(connectionId, connection);
        return connectionId;
    }

    // ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
    // code is executed on chain B).
    function connectionOpenTry(IBCHost host, IBCMsgs.MsgConnectionOpenTry memory msg_) public returns (string memory) {
        host.onlyIBCModule();
        require(IBCClient.validateSelfClient(host, msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        ConnectionEnd.Data memory connection = ConnectionEnd.Data({
            client_id: msg_.clientId,
            versions: getVersions(),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: msg_.delayPeriod,
            counterparty: msg_.counterparty
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: msg_.counterparty.client_id,
            versions: msg_.counterpartyVersions,
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: Counterparty.Data({
                client_id: msg_.clientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
            })
        });

        require(
            IBCClient.verifyConnectionState(
                host, connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            IBCClient.verifyClientState(host, connection, msg_.proofHeight, msg_.proofClient, msg_.clientStateBytes),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        string memory connectionId = host.generateConnectionIdentifier();
        host.setConnection(connectionId, connection);
        return connectionId;
    }

    function connectionOpenAck(IBCHost host, IBCMsgs.MsgConnectionOpenAck memory msg_) public {
        host.onlyIBCModule();
        (ConnectionEnd.Data memory connection, bool found) = host.getConnection(msg_.connectionId);
        require(found, "connection not found");

        if (connection.state != ConnectionEnd.State.STATE_INIT && connection.state != ConnectionEnd.State.STATE_TRYOPEN)
        {
            revert("connection state is not INIT or TRYOPEN");
        } else if (connection.state == ConnectionEnd.State.STATE_INIT && !isSupportedVersion(msg_.version)) {
            revert("connection state is in INIT but the provided version is not supported");
        } else if (
            connection.state == ConnectionEnd.State.STATE_TRYOPEN
                && (connection.versions.length != 1 || !isEqualVersion(connection.versions[0], msg_.version))
        ) {
            revert(
                "connection state is in TRYOPEN but the provided version is not set in the previous connection versions"
            );
        }

        require(IBCClient.validateSelfClient(host, msg_.clientStateBytes), "failed to validate self client state");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: makeVersionArray(msg_.version),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            IBCClient.verifyConnectionState(
                host, connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionID, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            IBCClient.verifyClientState(host, connection, msg_.proofHeight, msg_.proofClient, msg_.clientStateBytes),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        connection.state = ConnectionEnd.State.STATE_OPEN;
        connection.versions = expectedConnection.versions;
        connection.counterparty.connection_id = msg_.counterpartyConnectionID;
        host.setConnection(msg_.connectionId, connection);
    }

    function connectionOpenConfirm(IBCHost host, IBCMsgs.MsgConnectionOpenConfirm memory msg_) public {
        host.onlyIBCModule();
        (ConnectionEnd.Data memory connection, bool found) = host.getConnection(msg_.connectionId);
        require(found, "connection not found");

        require(connection.state == ConnectionEnd.State.STATE_TRYOPEN, "connection state is not TRYOPEN");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: bytes(commitmentPrefix)})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: connection.versions,
            state: ConnectionEnd.State.STATE_OPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            IBCClient.verifyConnectionState(
                host,
                connection,
                msg_.proofHeight,
                msg_.proofAck,
                connection.counterparty.connection_id,
                expectedConnection
            ),
            "failed to verify connection state"
        );

        connection.state = ConnectionEnd.State.STATE_OPEN;
        host.setConnection(msg_.connectionId, connection);
    }

    // Internal functions

    function getVersions() internal pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        string[] memory features = new string[](2);
        features[0] = "ORDER_ORDERED";
        features[1] = "ORDER_UNORDERED";
        versions[0] = Version.Data({identifier: "1", features: features});
        return versions;
    }

    // TODO implements
    function isSupportedVersion(Version.Data memory) internal pure returns (bool) {
        return true;
    }

    function isEqualVersion(Version.Data memory a, Version.Data memory b) internal pure returns (bool) {
        return keccak256(Version.encode(a)) == keccak256(Version.encode(b));
    }

    function makeVersionArray(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }
}
