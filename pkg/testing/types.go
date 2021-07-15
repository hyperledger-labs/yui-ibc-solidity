package testing

import (
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibchost"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	connectiontypes "github.com/datachainlab/ibc-solidity/pkg/ibc/connection"
)

// TestConnection is a testing helper struct to keep track of the connectionID, source clientID,
// counterparty clientID, and the next channel version used in creating and interacting with a
// connection.
type TestConnection struct {
	ID                   string
	ClientID             string
	CounterpartyClientID string
	NextChannelVersion   string
	Channels             []TestChannel
}

// TestChannel is a testing helper struct to keep track of the portID and channelID
// used in creating and interacting with a channel. The clientID and counterparty
// client ID are also tracked to cut down on querying and argument passing.
type TestChannel struct {
	PortID               string
	ID                   string
	ClientID             string
	CounterpartyClientID string
	Version              string
}

func connectionEndToPB(conn ibchost.ConnectionEndData) *connectiontypes.ConnectionEnd {
	connpb := &connectiontypes.ConnectionEnd{
		ClientId:    conn.ClientId,
		Versions:    []*connectiontypes.Version{},
		State:       connectiontypes.ConnectionEnd_State(conn.State),
		DelayPeriod: conn.DelayPeriod,
		Counterparty: &connectiontypes.Counterparty{
			ClientId:     conn.Counterparty.ClientId,
			ConnectionId: conn.Counterparty.ConnectionId,
			Prefix:       (*connectiontypes.MerklePrefix)(&conn.Counterparty.Prefix),
		},
	}
	for _, v := range conn.Versions {
		ver := connectiontypes.Version(v)
		connpb.Versions = append(connpb.Versions, &ver)
	}
	return connpb
}

func channelToPB(ch ibchost.ChannelData) *channeltypes.Channel {
	return &channeltypes.Channel{
		State:          channeltypes.Channel_State(ch.State),
		Ordering:       channeltypes.Channel_Order(ch.Ordering),
		Counterparty:   channeltypes.Channel_Counterparty(ch.Counterparty),
		ConnectionHops: ch.ConnectionHops,
		Version:        ch.Version,
	}
}
