// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IIBCHandler} from "../core/25-handler/IIBCHandler.sol";
import {LocalhostClient, LocalhostClientLib} from "../clients/LocalhostClient.sol";
import {Version, Counterparty, MerklePrefix} from "../proto/Connection.sol";
import {IBCConnectionLib} from "../core/03-connection/IBCConnectionLib.sol";
import {IIBCClient} from "../core/02-client/IIBCClient.sol";
import {IIBCClientErrors} from "../core/02-client/IIBCClientErrors.sol";
import {IIBCConnection} from "../core/03-connection/IIBCConnection.sol";
import {Height} from "../proto/Client.sol";
import {Channel, ChannelCounterparty} from "../proto/Channel.sol";
import {IbcLightclientsLocalhostV2ClientState as ClientState} from "../proto/Localhost.sol";
import {GoogleProtobufAny as Any} from "../proto/GoogleProtobufAny.sol";
import {IIBCChannelHandshake} from "../core/04-channel/IIBCChannel.sol";

/**
 * @title LocalhostHelper
 * @notice Helper functions for creating a client, connections, and channels with the localhost client
 */
library LocalhostHelper {
    /// @dev Msg for creating a connection
    struct MsgCreateConnection {
        Version.Data version;
        uint64 delayPeriod;
    }

    /// @dev Msg for creating a channel
    struct MsgCreateChannel {
        string connectionId0;
        string connectionId1;
        string portId0;
        string portId1;
        string version;
        Channel.Order ordering;
    }

    /**
     * @dev Get the localhost client
     */
    function getLocalhostClient(IIBCHandler ibcHandler) internal view returns (LocalhostClient) {
        address localhost = ibcHandler.getClientByType(LocalhostClientLib.CLIENT_TYPE);
        if (localhost == address(0)) {
            revert IIBCClientErrors.IBCClientUnregisteredClientType(LocalhostClientLib.CLIENT_TYPE);
        }
        return LocalhostClient(localhost);
    }

    /**
     * @dev Register the localhost client with the handler
     *      This function should be called only once per handler
     */
    function registerLocalhostClient(IIBCHandler ibcHandler) internal {
        LocalhostClient localhost = new LocalhostClient(address(ibcHandler));
        ibcHandler.registerClient(LocalhostClientLib.CLIENT_TYPE, localhost);
    }

    /**
     * @dev Create a localhost client
     */
    function createLocalhostClient(IIBCHandler ibcHandler) internal {
        ibcHandler.createClient(
            IIBCClient.MsgCreateClient({
                clientType: LocalhostClientLib.CLIENT_TYPE,
                protoClientState: Any.encode(
                    Any.Data({
                        type_url: LocalhostClientLib.CLIENT_STATE_TYPE_URL,
                        value: ClientState.encode(
                            ClientState.Data({
                                latest_height: Height.Data({revision_number: 0, revision_height: uint64(block.number)})
                            })
                            )
                    })
                    ),
                protoConsensusState: LocalhostClientLib.sentinelConsensusState()
            })
        );
    }

    /**
     * @dev Update the localhost client
     */
    function updateLocalhostClient(IIBCHandler ibcHandler) internal {
        ibcHandler.updateClient(
            IIBCClient.MsgUpdateClient({
                clientId: LocalhostClientLib.CLIENT_ID,
                protoClientMessage: Any.encode(
                    Any.Data({
                        type_url: LocalhostClientLib.CLIENT_STATE_TYPE_URL,
                        value: ClientState.encode(
                            ClientState.Data({
                                latest_height: Height.Data({revision_number: 0, revision_height: uint64(block.number)})
                            })
                            )
                    })
                    )
            })
        );
    }

    /**
     * @dev Create a localhost connection
     */
    function createLocalhostConnection(IIBCHandler ibcHandler)
        internal
        returns (string memory connectionId0, string memory connectionId1)
    {
        return createLocalhostConnection(ibcHandler, defaultMsgCreateConnection());
    }

    /**
     * @dev Create a localhost connection with the localhost client
     */
    function createLocalhostConnection(IIBCHandler ibcHandler, MsgCreateConnection memory msg_)
        internal
        returns (string memory connectionId0, string memory connectionId1)
    {
        // ensure the localhost client is created
        getLocalhostClient(ibcHandler);

        connectionId0 = ibcHandler.connectionOpenInit(
            IIBCConnection.MsgConnectionOpenInit({
                clientId: LocalhostClientLib.CLIENT_ID,
                counterparty: Counterparty.Data({
                    client_id: LocalhostClientLib.CLIENT_ID,
                    connection_id: "",
                    prefix: MerklePrefix.Data({key_prefix: ibcHandler.getCommitmentPrefix()})
                }),
                version: msg_.version,
                delayPeriod: msg_.delayPeriod
            })
        );
        connectionId1 = ibcHandler.connectionOpenTry(
            IIBCConnection.MsgConnectionOpenTry({
                counterparty: Counterparty.Data({
                    client_id: LocalhostClientLib.CLIENT_ID,
                    connection_id: connectionId0,
                    prefix: MerklePrefix.Data({key_prefix: ibcHandler.getCommitmentPrefix()})
                }),
                delayPeriod: msg_.delayPeriod,
                clientId: LocalhostClientLib.CLIENT_ID,
                clientStateBytes: Any.encode(
                    Any.Data({
                        type_url: LocalhostClientLib.CLIENT_STATE_TYPE_URL,
                        value: ClientState.encode(
                            ClientState.Data({
                                latest_height: Height.Data({revision_number: 0, revision_height: uint64(block.number)})
                            })
                            )
                    })
                    ),
                counterpartyVersions: getConnectionVersions(),
                proofInit: LocalhostClientLib.sentinelProof(),
                proofClient: LocalhostClientLib.sentinelProof(),
                proofConsensus: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0}),
                consensusHeight: Height.Data({revision_number: 0, revision_height: uint64(block.number)}),
                hostConsensusStateProof: LocalhostClientLib.sentinelConsensusState()
            })
        );
        ibcHandler.connectionOpenAck(
            IIBCConnection.MsgConnectionOpenAck({
                connectionId: connectionId0,
                counterpartyConnectionId: connectionId1,
                version: IBCConnectionLib.defaultIBCVersion(),
                clientStateBytes: Any.encode(
                    Any.Data({
                        type_url: LocalhostClientLib.CLIENT_STATE_TYPE_URL,
                        value: ClientState.encode(
                            ClientState.Data({
                                latest_height: Height.Data({revision_number: 0, revision_height: uint64(block.number)})
                            })
                            )
                    })
                    ),
                proofTry: LocalhostClientLib.sentinelProof(),
                proofClient: LocalhostClientLib.sentinelProof(),
                proofConsensus: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0}),
                consensusHeight: Height.Data({revision_number: 0, revision_height: uint64(block.number)}),
                hostConsensusStateProof: LocalhostClientLib.sentinelConsensusState()
            })
        );
        ibcHandler.connectionOpenConfirm(
            IIBCConnection.MsgConnectionOpenConfirm({
                connectionId: connectionId1,
                proofAck: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0})
            })
        );
    }

    /**
     * @dev Create a localhost channel with the localhost client
     */
    function createLocalhostChannel(IIBCHandler ibcHandler, MsgCreateChannel memory msg_)
        internal
        returns (string memory channelId0, string memory channelId1)
    {
        string memory version0;
        string memory version1;
        (channelId0, version0) = ibcHandler.channelOpenInit(
            IIBCChannelHandshake.MsgChannelOpenInit({
                portId: msg_.portId0,
                channel: Channel.Data({
                    state: Channel.State.STATE_INIT,
                    ordering: msg_.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: msg_.portId1, channel_id: ""}),
                    connection_hops: newConnectionHops(msg_.connectionId0),
                    version: msg_.version,
                    upgrade_sequence: 0
                })
            })
        );
        (channelId1, version1) = ibcHandler.channelOpenTry(
            IIBCChannelHandshake.MsgChannelOpenTry({
                portId: msg_.portId1,
                channel: Channel.Data({
                    state: Channel.State.STATE_TRYOPEN,
                    ordering: msg_.ordering,
                    counterparty: ChannelCounterparty.Data({port_id: msg_.portId0, channel_id: channelId0}),
                    connection_hops: newConnectionHops(msg_.connectionId1),
                    version: msg_.version,
                    upgrade_sequence: 0
                }),
                counterpartyVersion: version0,
                proofInit: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0})
            })
        );
        ibcHandler.channelOpenAck(
            IIBCChannelHandshake.MsgChannelOpenAck({
                portId: msg_.portId0,
                channelId: channelId0,
                counterpartyVersion: version1,
                counterpartyChannelId: channelId1,
                proofTry: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0})
            })
        );
        ibcHandler.channelOpenConfirm(
            IIBCChannelHandshake.MsgChannelOpenConfirm({
                portId: msg_.portId1,
                channelId: channelId1,
                proofAck: LocalhostClientLib.sentinelProof(),
                proofHeight: Height.Data({revision_number: 0, revision_height: 0})
            })
        );
    }

    function defaultMsgCreateConnection() internal pure returns (MsgCreateConnection memory) {
        return MsgCreateConnection({version: IBCConnectionLib.defaultIBCVersion(), delayPeriod: 0});
    }

    function getConnectionVersions() private pure returns (Version.Data[] memory) {
        Version.Data[] memory versions = new Version.Data[](1);
        versions[0] = IBCConnectionLib.defaultIBCVersion();
        return versions;
    }

    function newConnectionHops(string memory connectionId) private pure returns (string[] memory) {
        string[] memory connectionHops = new string[](1);
        connectionHops[0] = connectionId;
        return connectionHops;
    }
}
