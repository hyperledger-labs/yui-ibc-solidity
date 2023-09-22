package tests

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/channel"
	clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/client"
	ibctesting "github.com/hyperledger-labs/yui-ibc-solidity/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	relayer         = ibctesting.RelayerKeyIndex // the key-index of relayer on chain
	deployer        = ibctesting.RelayerKeyIndex // the key-index of contract deployer on chain
	alice    uint32 = 1                          // the key-index of alice on chain
	bob      uint32 = 2                          // the key-index of bob on chain
)

/*
NOTE: This test is intended to be run on ganache. Therefore, we are using MockClient instead of IBFT2Client.
*/
type ContractTestSuite struct {
	suite.Suite

	coordinator ibctesting.Coordinator
	chainA      *ibctesting.Chain
	chainB      *ibctesting.Chain
}

func (suite *ContractTestSuite) SetupTest() {
	ethClient, err := client.NewETHClient("http://127.0.0.1:8545")
	suite.Require().NoError(err)

	suite.chainA = ibctesting.NewChain(suite.T(), ethClient, ibctesting.NewLightClient(ethClient, clienttypes.MockClient))
	suite.chainB = ibctesting.NewChain(suite.T(), ethClient, ibctesting.NewLightClient(ethClient, clienttypes.MockClient))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite *ContractTestSuite) TestIBCCompatibility() {
	suite.T().Run("commitment path", func(t *testing.T) {
		const (
			testClientID = "tendermint-0"

			testConnectionID = "connection-0"
			testPortID       = "port-0"
			testChannelID    = "channel-0"
		)
		require := require.New(t)
		ctx := context.Background()

		// clientState
		path, err := suite.chainA.IBCCommitment.ClientStatePath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testClientID)
		require.NoError(err)
		require.Equal(host.FullClientStateKey(testClientID), path)

		// consensusState
		var cases = []uint64{0, 1, 10, 100}
		for _, n := range cases {
			for _, h := range cases {
				testHeight := ibcclienttypes.NewHeight(n, h)
				path, err := suite.chainA.IBCCommitment.ConsensusStatePath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testClientID, testHeight.RevisionNumber, testHeight.RevisionHeight)
				require.NoError(err)
				require.Equal(host.FullConsensusStateKey(testClientID, testHeight), path)
			}
		}
		// connectionState
		path, err = suite.chainA.IBCCommitment.ConnectionPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testConnectionID)
		require.NoError(err)
		require.Equal(host.ConnectionKey(testConnectionID), path)

		// channelState
		path, err = suite.chainA.IBCCommitment.ChannelPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testPortID, testChannelID)
		require.NoError(err)
		require.Equal(host.ChannelKey(testPortID, testChannelID), path)

		// packetCommitment
		var testSequence uint64 = 1
		path, err = suite.chainA.IBCCommitment.PacketCommitmentPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testPortID, testChannelID, testSequence)
		require.NoError(err)
		require.Equal(host.PacketCommitmentKey(testPortID, testChannelID, testSequence), path)

		// acknowledgementCommitment
		path, err = suite.chainA.IBCCommitment.PacketAcknowledgementCommitmentPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testPortID, testChannelID, testSequence)
		require.NoError(err)
		require.Equal(host.PacketAcknowledgementKey(testPortID, testChannelID, testSequence), path)
	})
}

func (suite *ContractTestSuite) TestPacketRelay() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	/// Tests for Transfer module ///

	denomA := strings.ToLower(chainA.ContractConfig.ERC20TokenAddress.String())
	balanceA0, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
	suite.Require().NoError(err)
	// Case: transfer tokens from chainA to chainB
	{
		suite.Require().NoError(chainA.ApproveAndDepositToken(ctx, deployer, 100, alice))

		// ensure that the balance is reduced
		balanceA1, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
		suite.Require().NoError(err)
		suite.Require().Equal(balanceA0.Int64()-100, balanceA1.Int64())

		bankA, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, alice).From, denomA)
		suite.Require().NoError(err)
		suite.Require().GreaterOrEqual(bankA.Int64(), int64(100))

		// try to transfer the token to chainB
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(
			chainA.ICS20Transfer.SendTransfer(
				chainA.TxOpts(ctx, alice),
				denomA,
				100,
				chainB.CallOpts(ctx, bob).From,
				chanA.PortID, chanA.ID,
				uint64(chainB.LastHeader().Number.Int64())+1000,
			),
		))

		suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB))

		// ensure that escrow has correct balance
		escrowBalance, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, alice), chainA.ContractConfig.ICS20TransferBankAddress, denomA)
		suite.Require().NoError(err)
		suite.Require().GreaterOrEqual(escrowBalance.Int64(), int64(100))

		// relay the packet
		transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, *transferPacket))
		suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *transferPacket, []byte{1}))
	}

	denomB := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, denomA)
	// Case: transfer tokens from chainB to chainA
	{
		// ensure that chainB has correct balance
		balance, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx, relayer), chainB.CallOpts(ctx, bob).From, denomB)
		suite.Require().NoError(err)
		suite.Require().Equal(int64(100), balance.Int64())

		// try to transfer the token to chainA
		suite.Require().NoError(chainB.WaitIfNoError(ctx)(
			chainB.ICS20Transfer.SendTransfer(
				chainB.TxOpts(ctx, bob),
				denomB,
				100,
				chainA.CallOpts(ctx, alice).From,
				chanB.PortID,
				chanB.ID,
				uint64(chainA.LastHeader().Number.Int64())+1000,
			),
		))

		suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))

		// relay the packet
		transferPacket, err := chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainA, chainB, chanA, chanB, *transferPacket))
		suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *transferPacket, []byte{1}))

		// withdraw tokens from the bank
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(
			chainA.ICS20Bank.Withdraw(
				chainA.TxOpts(ctx, alice),
				chainA.ContractConfig.ERC20TokenAddress,
				big.NewInt(100),
				chainA.CallOpts(ctx, deployer).From,
			)))

		// ensure that token balance equals original value
		balanceA2, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
		suite.Require().NoError(err)
		suite.Require().Equal(balanceA0.Int64(), balanceA2.Int64())
	}

	// Case: close channel
	{
		suite.coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
		// confirm that the channel is CLOSED on chain A
		chanData, ok, err := chainA.IBCHandler.GetChannel(chainA.CallOpts(ctx, relayer), chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().True(ok)
		suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)
		// confirm that the channel is CLOSED on chain B
		chanData, ok, err = chainB.IBCHandler.GetChannel(chainB.CallOpts(ctx, relayer), chanB.PortID, chanB.ID)
		suite.Require().NoError(err)
		suite.Require().True(ok)
		suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)
	}
}

func (suite *ContractTestSuite) TestTimeoutPacket() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, _ := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	denomA := strings.ToLower(chainA.ContractConfig.ERC20TokenAddress.String())

	suite.Require().NoError(chainA.ApproveAndDepositToken(ctx, deployer, 100, alice))

	// try to transfer the token to chainB
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransfer(
			chainA.TxOpts(ctx, alice),
			denomA,
			100,
			chainB.CallOpts(ctx, bob).From,
			chanA.PortID, chanA.ID,
			uint64(chainB.LastHeader().Number.Int64())+1,
		),
	))
	transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)

	// should fail to timeout packet because the timeout height is not reached
	suite.Require().Error(chainA.TimeoutPacket(ctx, *transferPacket, chainB))

	// execute a meaningless transaction to increase the block height
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainB.ERC20.Approve(chainA.TxOpts(ctx, deployer), chainB.ContractConfig.ICS20BankAddress, big.NewInt(0)),
	))
	// then, update the client to reach the timeout height
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))

	_, found, err := chainA.IBCHandler.GetHashedPacketCommitment(chainA.CallOpts(ctx, relayer), transferPacket.SourcePort, transferPacket.SourceChannel, transferPacket.Sequence)
	suite.Require().NoError(err)
	suite.Require().True(found)
	suite.Require().NoError(chainA.TimeoutPacket(ctx, *transferPacket, chainB))

	// confirm that the channel is OPEN on chain A
	chanData, ok, err := chainA.IBCHandler.GetChannel(chainA.CallOpts(ctx, relayer), chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	suite.Require().True(ok)
	suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.OPEN)

	// confirm that the packet commitment is deleted
	_, found, err = chainA.IBCHandler.GetHashedPacketCommitment(chainA.CallOpts(ctx, relayer), transferPacket.SourcePort, transferPacket.SourceChannel, transferPacket.Sequence)
	suite.Require().NoError(err)
	suite.Require().False(found)
}

func (suite *ContractTestSuite) TestTimeoutOnClose() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	suite.Require().NoError(chainA.ApproveAndDepositToken(ctx, deployer, 100, alice))
	denomA := strings.ToLower(chainA.ContractConfig.ERC20TokenAddress.String())

	// try to transfer the token to chainB
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransfer(
			chainA.TxOpts(ctx, alice),
			denomA,
			100,
			chainB.CallOpts(ctx, bob).From,
			chanA.PortID, chanA.ID,
			uint64(chainB.LastHeader().Number.Int64())+1000,
		),
	))

	transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)

	suite.Require().NoError(suite.coordinator.ChanCloseInit(ctx, chainB, chainA, chanB))
	suite.Require().NoError(suite.chainA.TimeoutOnClose(ctx, *transferPacket, chainB))
}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
