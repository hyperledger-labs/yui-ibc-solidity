package testing

import (
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibcclient"
	clienttypes "github.com/datachainlab/ibc-solidity/pkg/ibc/client"
)

type MsgCreateClient struct {
	ClientState    *clienttypes.ClientState
	ConsensusState *clienttypes.ConsensusState
}

func NewMsgCreateClient(clientState *clienttypes.ClientState, consensusState *clienttypes.ConsensusState) MsgCreateClient {
	return MsgCreateClient{
		ClientState:    clientState,
		ConsensusState: consensusState,
	}
}

func (msg MsgCreateClient) ClientStateData() ibcclient.ClientStateData {
	return ibcclient.ClientStateData{
		ChainId:              msg.ClientState.ChainId,
		ProvableStoreAddress: msg.ClientState.ProvableStoreAddress,
		LatestHeight:         msg.ClientState.LatestHeight,
	}
}

func (msg MsgCreateClient) ConsensusStateData() ibcclient.ConsensusStateData {
	return ibcclient.ConsensusStateData{
		Timestamp:  msg.ConsensusState.Timestamp,
		Root:       msg.ConsensusState.Root,
		Validators: msg.ConsensusState.Validators,
	}
}

type MsgUpdateClient struct {
	Header ibcclient.IBCClientHeader
}

func NewMsgUpdateClient(header ibcclient.IBCClientHeader) MsgUpdateClient {
	return MsgUpdateClient{Header: header}
}

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
