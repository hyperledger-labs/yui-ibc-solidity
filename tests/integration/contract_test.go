package tests

import (
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/consts"
	"github.com/datachainlab/ibc-solidity/pkg/contract"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"

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

// func (suite *ContractTestSuite) TestHandlePacketRecv() {
// 	ctx := context.Background()

// 	portId := "transfer"
// 	suite.T().Log(portId, suite.chain.ContractConfig.GetSimpleTokenModuleAddress())
// 	data := []byte("data0")
// 	suite.Require().NoError(
// 		suite.chain.WaitIfNoError(ctx)(
// 			suite.chain.IBCModule.HandlePacketRecvWithoutVerification(
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
