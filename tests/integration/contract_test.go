package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/consts"
	"github.com/datachainlab/ibc-solidity/pkg/contract"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibcstore"
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

	suite.chain = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
}

func (suite *ContractTestSuite) TestConnectionSerialization() {
	ctx := context.Background()
	connectionID := fmt.Sprintf("connection-%v", time.Now().Unix())
	connEnd := ibcstore.ConnectionEndData{
		ClientId: "client",
		Versions: []ibcstore.VersionData{
			{Identifier: "id", Features: []string{"a", "b"}},
		},
		// State: 1,
		// DelayPeriod: 1,
		// Counterparty: ibcstore.CounterpartyData{
		// 	// ClientId:     "cpclient",
		// 	// ConnectionId: "cpconnection",
		// 	// Prefix: ibcstore.MerklePrefixData{
		// 	// 	// KeyPrefix: []byte("ibc"),
		// 	// },
		// },
	}
	tx, err := suite.chain.IBCStore.SetConnection(
		suite.chain.TxOpts(ctx),
		connectionID,
		connEnd,
	)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.chain.WaitForReceiptAndGet(ctx, tx))

	actual, found, err := suite.chain.IBCStore.GetConnection(
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
	actualBytes, found, err := suite.chain.IBCStore.GetConnectionBytes(
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

// func (suite *ContractTestSuite) TestHandlePacketRecv() {
// 	ctx := context.Background()

// 	portId := "transfer"
// 	suite.T().Log(portId, suite.chain.ContractConfig.GetSimpleTokenModuleAddress())
// 	data := []byte("data0")
// 	suite.Require().NoError(
// 		suite.chain.WaitIfNoError(ctx)(
// 			suite.chain.IBCHandler.HandlePacketRecvWithoutVerification(
// 				suite.chain.TxOpts(ctx),
// 				ibchandler.IBCMsgsMsgPacketRecv{
// 					Packet: ibchandler.PacketData{
// 						DestinationPort: portId,
// 						Data:            data,
// 					},
// 				},
// 			),
// 		),
// 	)
// }

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
