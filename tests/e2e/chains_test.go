package e2e

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/channel"
	clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/client"
	ibctesting "github.com/hyperledger-labs/yui-ibc-solidity/pkg/testing"
	"github.com/stretchr/testify/suite"
)

const (
	relayer          = ibctesting.RelayerKeyIndex // the key-index of relayer on both chains
	deployerA        = ibctesting.RelayerKeyIndex // the key-index of contract deployer on chain A
	deployerB        = ibctesting.RelayerKeyIndex // the key-index of contract deployer on chain B
	aliceA    uint32 = 1                          // the key-index of alice on chain A
	bobB      uint32 = 2                          // the key-index of alice on chain B

	delayPeriodExtensionA = 5
	delayPeriodExtensionB = 10
)

type ChainTestSuite struct {
	suite.Suite
}

func (suite *ChainTestSuite) SetupTest() {}

func (suite *ChainTestSuite) TestPacketRelay() {
	ctx := context.Background()

	ethClA, err := client.NewETHClient("http://127.0.0.1:8645")
	suite.Require().NoError(err)
	ethClB, err := client.NewETHClient("http://127.0.0.1:8745")
	suite.Require().NoError(err)

	chainA := ibctesting.NewChain(suite.T(), ethClA, ibctesting.NewLightClient(ethClA, clienttypes.BesuIBFT2Client), true)
	chainA.SetDelayPeriod(3 * ibctesting.BlockTime)
	chainB := ibctesting.NewChain(suite.T(), ethClB, ibctesting.NewLightClient(ethClB, clienttypes.BesuIBFT2Client), true)
	chainB.SetDelayPeriod(3 * ibctesting.BlockTime)
	coordinator := ibctesting.NewCoordinator(suite.T(), chainA, chainB)

	clientA, clientB := coordinator.SetupClients(ctx, chainA, chainB, clienttypes.BesuIBFT2Client)
	connA, connB := coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	/// Tests for Transfer module ///

	beforeBalanceA, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployerA).From)
	suite.Require().NoError(err)
	suite.Require().NoError(
		coordinator.ApproveAndDepositToken(ctx, chainA, deployerA, 100, aliceA),
	)

	baseDenom := strings.ToLower(chainA.ContractConfig.ERC20TokenAddress.String())

	// set expectedTimePerBlock = block time on chainA
	suite.Require().NoError(chainA.SetExpectedTimePerBlock(ctx, deployerA, ibctesting.BlockTime))
	// set expectedTimePerBlock = 0 on chainB
	suite.Require().NoError(chainB.SetExpectedTimePerBlock(ctx, deployerB, 0))

	// try to transfer the token to chainB
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransfer(
			chainA.TxOpts(ctx, aliceA),
			baseDenom,
			100,
			chainB.CallOpts(ctx, bobB).From,
			chanA.PortID, chanA.ID,
			uint64(chainB.LastHeader().Number.Int64())+1000,
		),
	))
	suite.Require().NoError(coordinator.UpdateClient(ctx, chainB, chainA, clientB))

	// ensure that escrow has correct balance
	escrowBalance, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.ContractConfig.ICS20TransferBankAddress, baseDenom)
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(escrowBalance.Int64(), int64(100))

	// relay the packet
	coordinator.RelayLastSentPacketWithDelay(ctx, chainA, chainB, chanA, chanB, 1, 1)

	// ensure that chainB has correct balance
	expectedDenom := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, baseDenom)
	balance, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx, relayer), chainB.CallOpts(ctx, bobB).From, expectedDenom)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(100), balance.Int64())

	// make delay period 10 times longer on chainA
	suite.Require().NoError(
		chainA.SetExpectedTimePerBlock(ctx, deployerA, ibctesting.BlockTime/delayPeriodExtensionA),
	)

	// make delay period 20 times longer on chainB
	suite.Require().NoError(
		chainB.SetExpectedTimePerBlock(ctx, deployerB, ibctesting.BlockTime/delayPeriodExtensionB),
	)

	// try to transfer the token to chainA
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.ICS20Transfer.SendTransfer(
			chainB.TxOpts(ctx, bobB),
			expectedDenom,
			100,
			chainA.CallOpts(ctx, aliceA).From,
			chanB.PortID,
			chanB.ID,
			uint64(chainA.LastHeader().Number.Int64())+1000,
		),
	))

	suite.Require().NoError(coordinator.UpdateClient(ctx, chainA, chainB, clientA))

	// relay the packet
	coordinator.RelayLastSentPacketWithDelay(ctx, chainB, chainA, chanB, chanA, delayPeriodExtensionB, delayPeriodExtensionA)

	// withdraw tokens from the bank
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Bank.Withdraw(
			chainA.TxOpts(ctx, aliceA),
			chainA.ContractConfig.ERC20TokenAddress,
			big.NewInt(100),
			chainA.CallOpts(ctx, deployerA).From,
		)))

	// ensure that token balance equals original value
	afterBalanceA, err := chainA.ERC20.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployerA).From)
	suite.Require().NoError(err)
	suite.Require().Equal(beforeBalanceA.Int64(), afterBalanceA.Int64())

	// close channel
	coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ChainTestSuite))
}
