package tests

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibcmockapp"
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

	suite.chainA = ibctesting.NewChain(suite.T(), ethClient, ibctesting.NewLightClient(ethClient, clienttypes.MockClient), false)
	suite.chainB = ibctesting.NewChain(suite.T(), ethClient, ibctesting.NewLightClient(ethClient, clienttypes.MockClient), false)
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

		// packet receipt
		path, err = suite.chainA.IBCCommitment.PacketReceiptCommitmentPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testPortID, testChannelID, testSequence)
		require.NoError(err)
		require.Equal(host.PacketReceiptKey(testPortID, testChannelID, testSequence), path)

		// next sequence receive
		path, err = suite.chainA.IBCCommitment.NextSequenceRecvCommitmentPath(suite.chainA.CallOpts(ctx, ibctesting.RelayerKeyIndex), testPortID, testChannelID)
		require.NoError(err)
		require.Equal(host.NextSequenceRecvKey(testPortID, testChannelID), path)
	})
}

func (suite *ContractTestSuite) TestICS20() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED, ibctesting.ICS20Version)

	/// Tests for Transfer module ///

	denomA := strings.ToLower(chainA.ContractConfig.ERC20TokenAddress.String())
	balanceA0, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
	suite.Require().NoError(err)
	// Case: transfer tokens from chainA to chainB
	{
		suite.Require().NoError(suite.coordinator.ApproveAndDepositToken(ctx, chainA, deployer, 100, alice))

		// try to transfer the token to chainB
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(
			chainA.ICS20Transfer.SendTransfer(
				chainA.TxOpts(ctx, alice),
				denomA,
				big.NewInt(100),
				addressToHexString(chainB.CallOpts(ctx, bob).From),
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
		suite.coordinator.RelayLastSentPacket(ctx, chainA, chainB, chanA, chanB, func(b []byte) {
			var data transfertypes.FungibleTokenPacketData
			suite.Require().NoError(transfertypes.ModuleCdc.UnmarshalJSON(b, &data))
			suite.Require().NoError(data.ValidateBasic())
			suite.Require().Equal(denomA, data.Denom)
			suite.Require().Equal("100", data.Amount)
			suite.Require().Equal(addressToHexString(chainA.CallOpts(ctx, alice).From), data.Sender)
			suite.Require().Equal(addressToHexString(chainB.CallOpts(ctx, bob).From), data.Receiver)
			suite.Require().Equal("", data.Memo)
			suite.Require().Equal(data.GetBytes(), b)
		}, func(b []byte) {
			var ack ibcchanneltypes.Acknowledgement
			suite.Require().NoError(transfertypes.ModuleCdc.UnmarshalJSON(b, &ack))
			suite.Require().NoError(ack.ValidateBasic())
			suite.Require().True(ack.Success())
			suite.Require().Equal(ibcchanneltypes.NewResultAcknowledgement([]byte{byte(1)}).Acknowledgement(), b)
		})
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
				big.NewInt(100),
				addressToHexString(chainA.CallOpts(ctx, alice).From),
				chanB.PortID,
				chanB.ID,
				uint64(chainA.LastHeader().Number.Int64())+1000,
			),
		))

		suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))

		// relay the packet
		suite.coordinator.RelayLastSentPacket(ctx, chainB, chainA, chanB, chanA, func(b []byte) {
			var data transfertypes.FungibleTokenPacketData
			suite.Require().NoError(transfertypes.ModuleCdc.UnmarshalJSON(b, &data))
			suite.Require().NoError(data.ValidateBasic())
			suite.Require().Equal(denomB, data.Denom)
			suite.Require().Equal("100", data.Amount)
			suite.Require().Equal(addressToHexString(chainB.CallOpts(ctx, bob).From), data.Sender)
			suite.Require().Equal(addressToHexString(chainA.CallOpts(ctx, alice).From), data.Receiver)
			suite.Require().Equal("", data.Memo)
			suite.Require().Equal(data.GetBytes(), b)
		}, func(b []byte) {
			var ack ibcchanneltypes.Acknowledgement
			suite.Require().NoError(transfertypes.ModuleCdc.UnmarshalJSON(b, &ack))
			suite.Require().NoError(ack.ValidateBasic())
			suite.Require().True(ack.Success())
			suite.Require().Equal(ibcchanneltypes.NewResultAcknowledgement([]byte{byte(1)}).Acknowledgement(), b)
		})

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
}

func (suite *ContractTestSuite) TestTimeoutAndClose() {
	ctx := context.Background()
	coordinator := suite.coordinator
	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)

	// Case: timeoutOnClose on ordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.ORDERED, ibctesting.MockAppVersion)
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.IBCMockApp.SendPacket(
			chainA.TxOpts(ctx, alice),
			ibctesting.MockPacketData,
			chanA.PortID, chanA.ID,
			ibcmockapp.HeightData{RevisionNumber: 0, RevisionHeight: uint64(chainB.LastHeader().Number.Int64()) + 1000},
			0,
		)))
		packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(coordinator.ChanCloseInit(ctx, chainB, chainA, chanB))
		suite.Require().NoError(chainA.TimeoutOnClose(ctx, *packet, chainB, chanA, chanB))
		chainA.EnsureChannelState(ctx, chanA.PortID, chanA.ID, channeltypes.CLOSED)
	}

	// Case: timeoutOnClose on unordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.UNORDERED, ibctesting.MockAppVersion)
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.IBCMockApp.SendPacket(
			chainA.TxOpts(ctx, alice),
			ibctesting.MockPacketData,
			chanA.PortID, chanA.ID,
			ibcmockapp.HeightData{RevisionNumber: 0, RevisionHeight: uint64(chainB.LastHeader().Number.Int64()) + 1000},
			0,
		)))
		packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(coordinator.ChanCloseInit(ctx, chainB, chainA, chanB))
		suite.Require().NoError(chainA.TimeoutOnClose(ctx, *packet, chainB, chanA, chanB))
		chainA.EnsureChannelState(ctx, chanA.PortID, chanA.ID, channeltypes.CLOSED)
	}

	// Case: timeout packet on ordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.ORDERED, ibctesting.MockAppVersion)
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.IBCMockApp.SendPacket(
			chainA.TxOpts(ctx, alice),
			ibctesting.MockPacketData,
			chanA.PortID, chanA.ID,
			ibcmockapp.HeightData{RevisionNumber: 0, RevisionHeight: uint64(chainB.LastHeader().Number.Int64()) + 1},
			0,
		)))
		packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)

		// should fail to timeout packet because the timeout height is not reached
		suite.Require().Error(chainA.TimeoutPacket(ctx, *packet, chainB, chanA, chanB))

		suite.Require().NoError(chainB.AdvanceBlockNumber(ctx, uint64(chainB.LastHeader().Number.Int64())+1))

		// then, update the client to reach the timeout height
		suite.Require().NoError(coordinator.UpdateClient(ctx, chainA, chainB, clientA))

		suite.Require().NoError(chainA.EnsurePacketCommitmentExistence(ctx, true, packet.SourcePort, packet.SourceChannel, packet.Sequence))
		suite.Require().NoError(chainA.TimeoutPacket(ctx, *packet, chainB, chanA, chanB))
		// confirm that the packet commitment is deleted
		suite.Require().NoError(chainA.EnsurePacketCommitmentExistence(ctx, false, packet.SourcePort, packet.SourceChannel, packet.Sequence))
		chainA.EnsureChannelState(ctx, chanA.PortID, chanA.ID, channeltypes.CLOSED)
	}

	// Case: timeout packet on unordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.UNORDERED, ibctesting.MockAppVersion)
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.IBCMockApp.SendPacket(
			chainA.TxOpts(ctx, alice),
			ibctesting.MockPacketData,
			chanA.PortID, chanA.ID,
			ibcmockapp.HeightData{RevisionNumber: 0, RevisionHeight: uint64(chainB.LastHeader().Number.Int64()) + 1},
			0,
		)))
		packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)

		// should fail to timeout packet because the timeout height is not reached
		suite.Require().Error(chainA.TimeoutPacket(ctx, *packet, chainB, chanA, chanB))

		suite.Require().NoError(chainB.AdvanceBlockNumber(ctx, uint64(chainB.LastHeader().Number.Int64())+1))

		// then, update the client to reach the timeout height
		suite.Require().NoError(coordinator.UpdateClient(ctx, chainA, chainB, clientA))

		suite.Require().NoError(chainA.EnsurePacketCommitmentExistence(ctx, true, packet.SourcePort, packet.SourceChannel, packet.Sequence))
		suite.Require().NoError(chainA.TimeoutPacket(ctx, *packet, chainB, chanA, chanB))
		// confirm that the packet commitment is deleted
		suite.Require().NoError(chainA.EnsurePacketCommitmentExistence(ctx, false, packet.SourcePort, packet.SourceChannel, packet.Sequence))
		chainA.EnsureChannelState(ctx, chanA.PortID, chanA.ID, channeltypes.OPEN)
	}

	// Case: close channel on ordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.ORDERED, ibctesting.MockAppVersion)
		coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
	}

	// Case: close channel on unordered channel
	{
		chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.MockPort, ibctesting.MockPort, channeltypes.UNORDERED, ibctesting.MockAppVersion)
		coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
	}
}

func addressToHexString(addr common.Address) string {
	return strings.ToLower(addr.String())
}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
