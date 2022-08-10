package relay

import (
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
)

func connectionEndToPB(conn ibchost.ConnectionEndData) connectiontypes.ConnectionEnd {
	connpb := connectiontypes.ConnectionEnd{
		ClientId:    conn.ClientId,
		Versions:    []*connectiontypes.Version{},
		State:       connectiontypes.State(conn.State),
		DelayPeriod: conn.DelayPeriod,
		Counterparty: connectiontypes.Counterparty{
			ClientId:     conn.Counterparty.ClientId,
			ConnectionId: conn.Counterparty.ConnectionId,
			Prefix:       commitmenttypes.MerklePrefix(conn.Counterparty.Prefix),
		},
	}
	for _, v := range conn.Versions {
		ver := connectiontypes.Version(v)
		connpb.Versions = append(connpb.Versions, &ver)
	}
	return connpb
}

func channelToPB(chann ibchost.ChannelData) channeltypes.Channel {
	return channeltypes.Channel{
		State:          channeltypes.State(chann.State),
		Ordering:       channeltypes.Order(chann.Ordering),
		Counterparty:   channeltypes.Counterparty(chann.Counterparty),
		ConnectionHops: chann.ConnectionHops,
		Version:        chann.Version,
	}
}

func pbToHandlerHeight(height exported.Height) ibchandler.HeightData {
	return ibchandler.HeightData{
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
	}
}

func pbToHostHeight(height exported.Height) ibchost.HeightData {
	return ibchost.HeightData{
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
	}
}
