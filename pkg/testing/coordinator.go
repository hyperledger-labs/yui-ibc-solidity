package testing

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Coordinator struct {
	t      *testing.T
	chains []*Chain
}

func NewCoordinator(t *testing.T, chains ...*Chain) Coordinator {
	for _, chain := range chains {
		chain.UpdateHeader()
	}
	return Coordinator{t: t, chains: chains}
}

func (c Coordinator) GetChain(idx int) *Chain {
	return c.chains[idx]
}

// SetupClients is a helper function to create clients on both chains. It assumes the
// caller does not anticipate any errors.
func (coord *Coordinator) SetupClients(
	ctx context.Context,
	chainA, chainB *Chain,
	clientType string,
) (string, string) {

	clientA, err := coord.CreateClient(ctx, chainA, chainB, clientType)
	require.NoError(coord.t, err)

	clientB, err := coord.CreateClient(ctx, chainB, chainA, clientType)
	require.NoError(coord.t, err)

	return clientA, clientB
}

// SetupClientConnections is a helper function to create clients and the appropriate
// connections on both the source and counterparty chain. It assumes the caller does not
// anticipate any errors.
func (coord *Coordinator) SetupClientConnections(
	ctx context.Context,
	chainA, chainB *Chain,
	clientType string,
) (string, string, *TestConnection, *TestConnection) {

	clientA, clientB := coord.SetupClients(ctx, chainA, chainB, clientType)

	connA, connB := coord.CreateConnection(ctx, chainA, chainB, clientA, clientB)

	return clientA, clientB, connA, connB
}

func (coord *Coordinator) UpdateHeaders() {
	for _, c := range coord.chains {
		c.UpdateHeader()
	}
}

func (c Coordinator) CreateClient(
	ctx context.Context,
	source, counterparty *Chain,
	clientType string,
) (clientID string, err error) {
	clientID = source.NewClientID(clientType)
	switch clientType {
	case BesuIBFT2Client:
		err = source.CreateBesuClient(ctx, counterparty, clientID)

	default:
		err = fmt.Errorf("client type %s is not supported", clientType)
	}

	if err != nil {
		return "", err
	}

	return clientID, nil
}

func (c Coordinator) UpdateClient(
	ctx context.Context,
	source, counterparty *Chain,
	clientID string,
	clientType string,
) error {
	var err error
	switch clientType {
	case BesuIBFT2Client:
		err = source.UpdateBesuClient(ctx, counterparty, clientID)
	default:
		err = fmt.Errorf("client type %s is not supported", clientType)
	}
	if err != nil {
		return err
	}
	return nil
}

// CreateConnection constructs and executes connection handshake messages in order to create
// OPEN channels on chainA and chainB. The connection information of for chainA and chainB
// are returned within a TestConnection struct. The function expects the connections to be
// successfully opened otherwise testing will fail.
func (c *Coordinator) CreateConnection(
	ctx context.Context,
	chainA, chainB *Chain,
	clientA, clientB string,
) (*TestConnection, *TestConnection) {

	connA, connB, err := c.ConnOpenInit(ctx, chainA, chainB, clientA, clientB)
	require.NoError(c.t, err)

	require.NoError(c.t, c.ConnOpenTry(ctx, chainB, chainA, connB, connA))
	require.NoError(c.t, c.ConnOpenAck(ctx, chainA, chainB, connA, connB))
	require.NoError(c.t, c.ConnOpenConfirm(ctx, chainB, chainA, connB, connA))

	return connA, connB
}

// ConnOpenInit initializes a connection on the source chain with the state INIT
// using the OpenInit handshake call.
//
// NOTE: The counterparty testing connection will be created even if it is not created in the
// application state.
func (c Coordinator) ConnOpenInit(
	ctx context.Context,
	source, counterparty *Chain,
	clientID, counterpartyClientID string,
) (*TestConnection, *TestConnection, error) {

	sourceConnection := source.AddTestConnection(clientID, counterpartyClientID)
	counterpartyConnection := counterparty.AddTestConnection(counterpartyClientID, clientID)

	// initialize connection on source
	if err := source.ConnectionOpenInit(ctx, counterparty, sourceConnection, counterpartyConnection); err != nil {
		return sourceConnection, counterpartyConnection, err
	}

	source.UpdateHeader()

	// update source client on counterparty connection
	if err := c.UpdateClient(
		ctx,
		counterparty, source,
		counterpartyClientID,
		BesuIBFT2Client,
	); err != nil {
		return sourceConnection, counterpartyConnection, err
	}

	return sourceConnection, counterpartyConnection, nil
}

// ConnOpenTry initializes a connection on the source chain with the state TRYOPEN
// using the OpenTry handshake call.
func (c *Coordinator) ConnOpenTry(
	ctx context.Context,
	source, counterparty *Chain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {

	if err := source.ConnectionOpenTry(ctx, counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}

	source.UpdateHeader()

	return c.UpdateClient(
		ctx,
		counterparty, source,
		counterpartyConnection.ClientID,
		BesuIBFT2Client,
	)
}

// ConnOpenAck initializes a connection on the source chain with the state OPEN
// using the OpenAck handshake call.
func (c *Coordinator) ConnOpenAck(
	ctx context.Context,
	source, counterparty *Chain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {
	// set OPEN connection on source using OpenAck
	if err := source.ConnectionOpenAck(ctx, counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}

	source.UpdateHeader()

	// update source client on counterparty connection
	return c.UpdateClient(
		ctx,
		counterparty, source,
		counterpartyConnection.ClientID,
		BesuIBFT2Client,
	)
}

// ConnOpenConfirm initializes a connection on the source chain with the state OPEN
// using the OpenConfirm handshake call.
func (c *Coordinator) ConnOpenConfirm(
	ctx context.Context,
	source, counterparty *Chain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {
	if err := source.ConnectionOpenConfirm(ctx, counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}

	source.UpdateHeader()

	// update source client on counterparty connection
	return c.UpdateClient(
		ctx,
		counterparty, source,
		counterpartyConnection.ClientID,
		BesuIBFT2Client,
	)
}
