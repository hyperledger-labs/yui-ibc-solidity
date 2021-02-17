package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/consts"
	"github.com/datachainlab/ibc-solidity/pkg/contract"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibcroutingmodule"
	"github.com/datachainlab/ibc-solidity/pkg/contract/provablestore"
	connectiontypes "github.com/datachainlab/ibc-solidity/pkg/ibc/connection"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"
	"github.com/gogo/protobuf/proto"

	"github.com/stretchr/testify/suite"
)

const mnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ContractTestSuite struct {
	suite.Suite

	chain *ibctesting.Chain
}

func (suite *ContractTestSuite) SetupTest() {
	chainClient, err := contract.CreateClient("http://127.0.0.1:8545")
	suite.Require().NoError(err)

	suite.chain = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase)
}

func (suite *ContractTestSuite) TestConnection() {
	ctx := context.Background()
	connectionID := fmt.Sprintf("connection-%v", time.Now().Unix())
	connEnd := provablestore.ConnectionEndData{
		ClientId: "client",
		Versions: []provablestore.VersionData{
			{Identifier: "id", Features: []string{"a", "b"}},
		},
		// State: 1,
		// DelayPeriod: 1,
		// Counterparty: provablestore.CounterpartyData{
		// 	// ClientId:     "cpclient",
		// 	// ConnectionId: "cpconnection",
		// 	// Prefix: provablestore.MerklePrefixData{
		// 	// 	// KeyPrefix: []byte("ibc"),
		// 	// },
		// },
	}
	tx, err := suite.chain.ProvableStore.SetConnection(
		suite.chain.TxOpts(ctx),
		connectionID,
		connEnd,
	)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.chain.WaitForReceiptAndGet(ctx, tx))

	actual, found, err := suite.chain.ProvableStore.GetConnection(
		suite.chain.CallOpts(ctx),
		connectionID,
	)
	suite.Require().NoError(err)
	suite.Require().True(found)
	suite.Equal(connEnd.Versions[0], actual.Versions[0])

	connEndPB := &connectiontypes.ConnectionEnd{
		ClientId: "client",
		Versions: []*connectiontypes.Version{
			{Identifier: "id", Features: []string{"a", "b"}},
		},
		// State: ibctypes.State_STATE_INIT,
		// DelayPeriod: 2,
		Counterparty: &connectiontypes.Counterparty{
			// ClientId:     "cpclient",
			// ConnectionId: "cpconnection",
			Prefix: &connectiontypes.MerklePrefix{
				// KeyPrefix: []byte("ibc"),
			},
		},
	}

	expectedBytes, err := proto.Marshal(connEndPB)
	suite.Require().NoError(err)
	actualBytes, found, err := suite.chain.ProvableStore.GetConnectionBytes(
		suite.chain.CallOpts(ctx),
		connectionID,
	)
	suite.Require().NoError(err)
	suite.Require().True(found)
	suite.Equal(expectedBytes, actualBytes)

	var ret connectiontypes.ConnectionEnd
	suite.Require().NoError(proto.Unmarshal(actualBytes, &ret))
	suite.True(proto.Equal(connEndPB, &ret))
}

func (suite *ContractTestSuite) TestHandlePacketRecv() {
	ctx := context.Background()

	portId := fmt.Sprintf("port-%v", time.Now().Unix())
	suite.T().Log(portId, suite.chain.ContractConfig.GetSimpleTokenModuleAddress())

	suite.Require().NoError(
		suite.chain.WaitIfNoError(ctx)(
			suite.chain.IBCRoutingModule.BindPort(
				suite.chain.TxOpts(ctx),
				portId,
				suite.chain.ContractConfig.GetSimpleTokenModuleAddress(),
			),
		),
	)
	data := []byte("data0")
	suite.Require().NoError(
		suite.chain.WaitIfNoError(ctx)(
			suite.chain.IBCRoutingModule.HandlePacketRecvWithoutVerification(
				suite.chain.TxOpts(ctx),
				ibcroutingmodule.IBCRoutingModulePacketRecv{
					Packet: ibcroutingmodule.PacketData{
						DestinationPort: portId,
						Data:            data,
					},
				},
			),
		),
	)
}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
