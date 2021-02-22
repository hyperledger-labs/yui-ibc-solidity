package testing

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
