package e2e

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/contract"
	ibcapp "github.com/datachainlab/ibc-solidity/pkg/ibc/app"
	"github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"
	testchain0 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain0"
	testchain1 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain1"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
)

const mnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ChainTestSuite struct {
	suite.Suite

	coordinator ibctesting.Coordinator
	chainA      *ibctesting.Chain
	chainB      *ibctesting.Chain
}

func (suite *ChainTestSuite) SetupTest() {
	chainClientA, err := contract.CreateClient("http://127.0.0.1:8645")
	suite.Require().NoError(err)

	chainClientB, err := contract.CreateClient("http://127.0.0.1:8745")
	suite.Require().NoError(err)

	ibcID := uint64(time.Now().UnixNano())
	suite.chainA = ibctesting.NewChain(suite.T(), 2018, *chainClientA, testchain0.Contract, mnemonicPhrase, ibcID)
	suite.chainB = ibctesting.NewChain(suite.T(), 3018, *chainClientB, testchain1.Contract, mnemonicPhrase, ibcID)
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite ChainTestSuite) TestStringUtils() {
	ctx := context.Background()
	addr, err := suite.chainA.ICS20Transfer.AddressToString(
		suite.chainA.CallOpts(ctx),
		suite.chainA.CallOpts(ctx).From,
	)
	suite.Require().NoError(err)
	fmt.Println(suite.chainA.CallOpts(ctx).From.String(), addr)

	parsed, err := suite.chainA.ICS20Transfer.ParseAddr(
		suite.chainA.CallOpts(ctx),
		addr,
	)
	suite.Require().NoError(err)
	fmt.Println(suite.chainA.CallOpts(ctx).From.String(), parsed.String())
}

func (suite ChainTestSuite) TestChannel() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, ibctesting.BesuIBFT2Client)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channel.UNORDERED)
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.SimpleToken.Approve(chainA.TxOpts(ctx), chainA.ContractConfig.GetICS20TransferAddress(), big.NewInt(100)),
	))
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransferWithTokenContract(
			chainA.TxOpts(ctx),
			chainA.ContractConfig.GetSimpleTokenAddress(),
			big.NewInt(100),
			chainB.CallOpts(ctx).From,
			chanA.PortID, chanA.ID,
			uint64(chainA.LastHeader().Base.Number.Int64())+1000,
		),
	))
	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB, ibctesting.BesuIBFT2Client))

	seq, err := suite.chainA.IBCHost.GetNextSequenceSend(chainA.CallOpts(ctx), chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	packet, err := chainA.IBCHost.GetPacket(chainA.CallOpts(ctx), chanA.PortID, chanA.ID, seq-1)
	suite.Require().NoError(err)

	var pd ibcapp.FungibleTokenPacketData
	suite.Require().NoError(proto.Unmarshal(packet.Data, &pd))
	transferPacket := channel.NewPacket(packet.Data, packet.Sequence, packet.SourcePort, packet.SourceChannel, packet.DestinationPort, packet.DestinationChannel, channeltypes.Height(packet.TimeoutHeight), packet.TimeoutTimestamp)
	fmt.Println(transferPacket.String())
	fmt.Println(pd.String())

	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, transferPacket))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, transferPacket, []byte{1}))
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ChainTestSuite))
}
