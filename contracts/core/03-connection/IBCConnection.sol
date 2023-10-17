// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";
import "../../proto/Client.sol";
import "../../proto/Connection.sol";
import "../25-handler/IBCMsgs.sol";
import "../24-host/IBCStore.sol";
import "../24-host/IBCCommitment.sol";
import "../03-connection/IIBCConnection.sol";

/**
 * @dev IBCConnection is a contract that implements [ICS-3](https://github.com/cosmos/ibc/tree/main/spec/core/ics-003-connection-semantics).
 */
contract IBCConnection is IBCStore, IIBCConnectionHandshake {
    /* Handshake functions */

    /**
     * @dev connectionOpenInit initialises a connection attempt on chain A. The generated connection identifier
     * is returned.
     */
    function connectionOpenInit(IBCMsgs.MsgConnectionOpenInit calldata msg_)
        external
        override
        returns (string memory)
    {
        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        require(connection.state == ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED, "connectionId already exists");
        // ensure the client exists
        checkAndGetClient(msg_.clientId);
        require(bytes(msg_.counterparty.connection_id).length == 0, "counterparty connectionId must be empty");
        connection.client_id = msg_.clientId;

        if (msg_.version.features.length > 0) {
            require(
                IBCConnectionLib.isSupportedVersion(getCompatibleVersions(), msg_.version),
                "the selected version is not supported"
            );
            connection.versions.push(msg_.version);
        } else {
            IBCConnectionLib.setSupportedVersions(getCompatibleVersions(), connection.versions);
        }

        connection.state = ConnectionEnd.State.STATE_INIT;
        connection.delay_period = msg_.delayPeriod;
        connection.counterparty = msg_.counterparty;
        updateConnectionCommitment(connectionId);
        return connectionId;
    }

    /**
     * @dev connectionOpenTry relays notice of a connection attempt on chain A to chain B (this
     * code is executed on chain B).
     */
    function connectionOpenTry(IBCMsgs.MsgConnectionOpenTry calldata msg_) external override returns (string memory) {
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");
        require(msg_.counterpartyVersions.length > 0, "counterpartyVersions length must be greater than 0");

        string memory connectionId = generateConnectionIdentifier();
        ConnectionEnd.Data storage connection = connections[connectionId];
        require(connection.state == ConnectionEnd.State.STATE_UNINITIALIZED_UNSPECIFIED, "connectionId already exists");
        // ensure the client exists
        checkAndGetClient(msg_.clientId);

        connection.versions.push(IBCConnectionLib.pickVersion(getCompatibleVersions(), msg_.counterpartyVersions));
        connection.client_id = msg_.clientId;
        connection.state = ConnectionEnd.State.STATE_TRYOPEN;
        connection.delay_period = msg_.delayPeriod;
        connection.counterparty = msg_.counterparty;

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: msg_.counterparty.client_id,
            versions: msg_.counterpartyVersions,
            state: ConnectionEnd.State.STATE_INIT,
            delay_period: msg_.delayPeriod,
            counterparty: Counterparty.Data({
                client_id: msg_.clientId,
                connection_id: "",
                prefix: MerklePrefix.Data({key_prefix: getCommitmentPrefix()})
            })
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofInit, msg_.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            verifyClientState(
                connection,
                msg_.proofHeight,
                IBCCommitment.clientStatePath(connection.counterparty.client_id),
                msg_.proofClient,
                msg_.clientStateBytes
            ),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        updateConnectionCommitment(connectionId);
        return connectionId;
    }

    /**
     * @dev connectionOpenAck relays acceptance of a connection open attempt from chain B back
     * to chain A (this code is executed on chain A).
     */
    function connectionOpenAck(IBCMsgs.MsgConnectionOpenAck calldata msg_) external override {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        require(connection.state == ConnectionEnd.State.STATE_INIT, "connection state is not INIT");
        require(
            IBCConnectionLib.isSupportedVersion(connection.versions, msg_.version),
            "the counterparty selected version is not supported by versions selected on INIT"
        );
        require(validateSelfClient(msg_.clientStateBytes), "failed to validate self client state");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: getCommitmentPrefix()})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: IBCConnectionLib.newVersions(msg_.version),
            state: ConnectionEnd.State.STATE_TRYOPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofTry, msg_.counterpartyConnectionId, expectedConnection
            ),
            "failed to verify connection state"
        );
        require(
            verifyClientState(
                connection,
                msg_.proofHeight,
                IBCCommitment.clientStatePath(connection.counterparty.client_id),
                msg_.proofClient,
                msg_.clientStateBytes
            ),
            "failed to verify clientState"
        );
        // TODO we should also verify a consensus state

        connection.state = ConnectionEnd.State.STATE_OPEN;
        connection.counterparty.connection_id = msg_.counterpartyConnectionId;
        IBCConnectionLib.copyVersions(expectedConnection.versions, connection.versions);
        updateConnectionCommitment(msg_.connectionId);
    }

    /**
     * @dev connectionOpenConfirm confirms opening of a connection on chain A to chain B, after
     * which the connection is open on both chains (this code is executed on chain B).
     */
    function connectionOpenConfirm(IBCMsgs.MsgConnectionOpenConfirm calldata msg_) external override {
        ConnectionEnd.Data storage connection = connections[msg_.connectionId];
        require(connection.state == ConnectionEnd.State.STATE_TRYOPEN, "connection state is not TRYOPEN");

        Counterparty.Data memory expectedCounterparty = Counterparty.Data({
            client_id: connection.client_id,
            connection_id: msg_.connectionId,
            prefix: MerklePrefix.Data({key_prefix: getCommitmentPrefix()})
        });

        ConnectionEnd.Data memory expectedConnection = ConnectionEnd.Data({
            client_id: connection.counterparty.client_id,
            versions: connection.versions,
            state: ConnectionEnd.State.STATE_OPEN,
            delay_period: connection.delay_period,
            counterparty: expectedCounterparty
        });

        require(
            verifyConnectionState(
                connection, msg_.proofHeight, msg_.proofAck, connection.counterparty.connection_id, expectedConnection
            ),
            "failed to verify connection state"
        );

        connection.state = ConnectionEnd.State.STATE_OPEN;
        updateConnectionCommitment(msg_.connectionId);
    }

    function updateConnectionCommitment(string memory connectionId) private {
        commitments[IBCCommitment.connectionCommitmentKey(connectionId)] =
            keccak256(ConnectionEnd.encode(connections[connectionId]));
    }

    /* Verification functions */

    function verifyClientState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory path,
        bytes memory proof,
        bytes memory clientStateBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id, height, 0, 0, proof, connection.counterparty.prefix.key_prefix, path, clientStateBytes
        );
    }

    function verifyClientConsensusState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        Height.Data memory consensusHeight,
        bytes memory proof,
        bytes memory consensusStateBytes
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCCommitment.consensusStatePath(
                connection.counterparty.client_id, consensusHeight.revision_number, consensusHeight.revision_height
            ),
            consensusStateBytes
        );
    }

    function verifyConnectionState(
        ConnectionEnd.Data storage connection,
        Height.Data memory height,
        bytes memory proof,
        string memory connectionId,
        ConnectionEnd.Data memory counterpartyConnection
    ) private returns (bool) {
        return checkAndGetClient(connection.client_id).verifyMembership(
            connection.client_id,
            height,
            0,
            0,
            proof,
            connection.counterparty.prefix.key_prefix,
            IBCCommitment.connectionPath(connectionId),
            ConnectionEnd.encode(counterpartyConnection)
        );
    }

    /* Internal functions */

    function generateConnectionIdentifier() private returns (string memory) {
        string memory identifier = string(abi.encodePacked("connection-", Strings.toString(nextConnectionSequence)));
        nextConnectionSequence++;
        return identifier;
    }

    /**
     * @dev validateSelfClient validates the client parameters for a client of the host chain.
     *
     * NOTE: Developers can override this function to support an arbitrary EVM chain.
     */
    function validateSelfClient(bytes memory) internal view virtual returns (bool) {
        this; // this is a trick that suppresses "Warning: Function state mutability can be restricted to pure"
        return true;
    }

    /**
     * @dev getCompatibleVersions returns the supported versions of the host chain.
     */
    function getCompatibleVersions() internal pure virtual returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = IBCConnectionLib.defaultIBCVersion();
        return versions;
    }
}

library IBCConnectionLib {
    string public constant IBC_VERSION_IDENTIFIER = "1";
    string public constant ORDER_ORDERED = "ORDER_ORDERED";
    string public constant ORDER_UNORDERED = "ORDER_UNORDERED";

    /**
     * @dev defaultIBCVersion returns the latest supported version of IBC used in connection version negotiation
     */
    function defaultIBCVersion() internal pure returns (Version.Data memory) {
        Version.Data memory version = Version.Data({identifier: IBC_VERSION_IDENTIFIER, features: new string[](2)});
        version.features[0] = ORDER_ORDERED;
        version.features[1] = ORDER_UNORDERED;
        return version;
    }

    /**
     * @dev setSupportedVersions sets the supported versions to a given array.
     *
     * NOTE: `versions` must be an empty array
     */
    function setSupportedVersions(Version.Data[] memory supportedVersions, Version.Data[] storage dst) internal {
        require(dst.length == 0, "versions must be empty");
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            dst.push(supportedVersions[i]);
        }
    }

    /**
     * @dev isSupportedVersion returns true if the proposed version has a matching version
     * identifier and its entire feature set is supported or the version identifier
     * supports an empty feature set.
     */
    function isSupportedVersion(Version.Data[] memory supportedVersions, Version.Data memory version)
        internal
        pure
        returns (bool)
    {
        (Version.Data memory supportedVersion, bool found) = findSupportedVersion(version, supportedVersions);
        if (!found) {
            return false;
        }
        return verifyProposedVersion(supportedVersion, version);
    }

    /**
     * @dev verifyProposedVersion verifies that the entire feature set in the
     * proposed version is supported by this chain. If the feature set is
     * empty it verifies that this is allowed for the specified version
     * identifier.
     */
    function verifyProposedVersion(Version.Data memory supportedVersion, Version.Data memory proposedVersion)
        internal
        pure
        returns (bool)
    {
        if (
            keccak256(abi.encodePacked(proposedVersion.identifier))
                != keccak256(abi.encodePacked(supportedVersion.identifier))
        ) {
            return false;
        }
        if (proposedVersion.features.length == 0) {
            return false;
        }
        for (uint256 i = 0; i < proposedVersion.features.length; i++) {
            if (!contains(proposedVersion.features[i], supportedVersion.features)) {
                return false;
            }
        }
        return true;
    }

    /**
     * @dev findSupportedVersion returns the version with a matching version identifier
     * if it exists. The returned boolean is true if the version is found and
     * false otherwise.
     */
    function findSupportedVersion(Version.Data memory version, Version.Data[] memory supportedVersions)
        internal
        pure
        returns (Version.Data memory supportedVersion, bool found)
    {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            supportedVersion = supportedVersions[i];
            if (
                keccak256(abi.encodePacked(supportedVersion.identifier))
                    == keccak256(abi.encodePacked(version.identifier))
            ) {
                return (supportedVersion, true);
            }
        }
        return (supportedVersion, false);
    }

    /**
     * @dev pickVersion iterates over the descending ordered set of compatible IBC
     * versions and selects the first version with a version identifier that is
     * supported by the counterparty. The returned version contains a feature
     * set with the intersection of the features supported by the source and
     * counterparty chains. If the feature set intersection is nil and this is
     * not allowed for the chosen version identifier then the search for a
     * compatible version continues. This function is called in the ConnOpenTry
     * handshake procedure.
     *
     * CONTRACT: pickVersion must only provide a version that is in the
     * intersection of the supported versions and the counterparty versions.
     */
    function pickVersion(Version.Data[] memory supportedVersions, Version.Data[] memory counterpartyVersions)
        internal
        pure
        returns (Version.Data memory)
    {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            Version.Data memory supportedVersion = supportedVersions[i];
            (Version.Data memory counterpartyVersion, bool found) =
                findSupportedVersion(supportedVersion, counterpartyVersions);
            if (!found) {
                continue;
            }
            string[] memory featureSet =
                getFeatureSetIntersection(supportedVersion.features, counterpartyVersion.features);
            if (featureSet.length > 0) {
                return Version.Data({identifier: supportedVersion.identifier, features: featureSet});
            }
        }
        revert("matching counterparty version not found");
    }

    /**
     * @dev copyVersions copies `src` to `dst`
     */
    function copyVersions(Version.Data[] memory src, Version.Data[] storage dst) internal {
        require(dst.length == src.length, "length mismatch");
        for (uint256 i = 0; i < src.length; i++) {
            copyVersion(src[i], dst[i]);
        }
    }

    /**
     * @dev newVersions returns a new array with a given version
     */
    function newVersions(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }

    /**
     * @dev verifySupportedFeature takes in a version and feature string and returns
     * true if the feature is supported by the version and false otherwise.
     */
    function verifySupportedFeature(Version.Data memory version, string memory feature) internal pure returns (bool) {
        bytes32 hashedFeature = keccak256(bytes(feature));
        for (uint256 i = 0; i < version.features.length; i++) {
            if (keccak256(bytes(version.features[i])) == hashedFeature) {
                return true;
            }
        }
        return false;
    }

    function getFeatureSetIntersection(string[] memory sourceFeatureSet, string[] memory counterpartyFeatureSet)
        private
        pure
        returns (string[] memory)
    {
        string[] memory featureSet = new string[](sourceFeatureSet.length);
        uint256 featureSetLength = 0;
        for (uint256 i = 0; i < sourceFeatureSet.length; i++) {
            if (contains(sourceFeatureSet[i], counterpartyFeatureSet)) {
                featureSet[featureSetLength] = sourceFeatureSet[i];
                featureSetLength++;
            }
        }
        string[] memory ret = new string[](featureSetLength);
        for (uint256 i = 0; i < featureSetLength; i++) {
            ret[i] = featureSet[i];
        }
        return ret;
    }

    function copyVersion(Version.Data memory src, Version.Data storage dst) private {
        dst.identifier = src.identifier;
        for (uint256 i = 0; i < src.features.length; i++) {
            dst.features[i] = src.features[i];
        }
    }

    function contains(string memory elem, string[] memory set) private pure returns (bool) {
        bytes32 hashedElem = keccak256(bytes(elem));
        for (uint256 i = 0; i < set.length; i++) {
            if (keccak256(bytes(set[i])) == hashedElem) {
                return true;
            }
        }
        return false;
    }
}
