// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";

interface IIBCConnectionErrors {
    error IBCConnectionAlreadyConnectionExists();

    /// @param counterpartyConnectionId counterparty connection identifier
    error IBCConnectionInvalidCounterpartyConnectionIdentifier(string counterpartyConnectionId);

    error IBCConnectionEmptyConnectionCounterpartyVersions();

    error IBCConnectionNoMatchingVersionFound();

    error IBCConnectionVersionsAlreadySet();

    error IBCConnectionIBCVersionNotSupported();

    error IBCConnectionVersionIdentifierNotEmpty();

    error IBCConnectionInvalidSelfClientState();

    error IBCConnectionInvalidHostConsensusStateProof();

    /// @param state connection state
    error IBCConnectionUnexpectedConnectionState(ConnectionEnd.State state);

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param value value
    /// @param proof proof
    /// @param height proof height
    error IBCConnectionFailedVerifyConnectionState(
        string clientId, bytes path, bytes value, bytes proof, Height.Data height
    );

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param value value
    /// @param proof proof
    /// @param height proof height
    error IBCConnectionFailedVerifyClientState(
        string clientId, bytes path, bytes value, bytes proof, Height.Data height
    );

    /// @param clientId client identifier
    /// @param path commitment path
    /// @param value value
    /// @param proof proof
    /// @param height proof height
    error IBCConnectionFailedVerifyClientConsensusState(
        string clientId, bytes path, bytes value, bytes proof, Height.Data height
    );
}
