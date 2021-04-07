package tests

import (
	"context"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/client"
	"github.com/datachainlab/ibc-solidity/pkg/consts"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	clienttypes "github.com/datachainlab/ibc-solidity/pkg/ibc/client"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"

	"github.com/stretchr/testify/suite"
)

const mnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ContractTestSuite struct {
	suite.Suite

	coordinator ibctesting.Coordinator
	chainA      *ibctesting.Chain
	chainB      *ibctesting.Chain
}

func (suite *ContractTestSuite) SetupTest() {
	chainClient, err := client.NewETHClient("http://127.0.0.1:8545", clienttypes.MockClient)
	suite.Require().NoError(err)

	suite.chainA = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite *ContractTestSuite) TestClient() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)
	_, _ = chanA, chanB
}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
