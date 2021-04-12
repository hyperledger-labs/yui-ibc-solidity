package tests

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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
	chainClient, err := client.NewETHClient("http://127.0.0.1:8545", clienttypes.MockClient)
	suite.Require().NoError(err)

	suite.chainA = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), 2018, *chainClient, consts.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite *ContractTestSuite) TestChannel() {
	ctx := context.Background()

	const (
		relayer         = ibctesting.RelayerKeyIndex // the key-index of relayer on chain
		deployer        = ibctesting.RelayerKeyIndex // the key-index of contract deployer on chain
		alice    uint32 = 1                          // the key-index of alice on chain
		bob      uint32 = 2                          // the key-index of bob on chain
	)

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	/// Tests for Transfer module ///

	balance0, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
	suite.Require().NoError(err)
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.SimpleToken.Approve(chainA.TxOpts(ctx, deployer), chainA.ContractConfig.GetICS20BankAddress(), big.NewInt(100)),
	))

	// deposit a simple token to the bank
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.ICS20Bank.Deposit(
		chainA.TxOpts(ctx, deployer),
		chainA.ContractConfig.GetSimpleTokenAddress(),
		big.NewInt(100),
		chainA.CallOpts(ctx, alice).From,
	)))

	// ensure that the balance is reduced
	balance1, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
	suite.Require().NoError(err)
	suite.Require().Equal(balance0.Int64()-100, balance1.Int64())

	baseDenom := strings.ToLower(chainA.ContractConfig.GetSimpleTokenAddress().String())

	bankA, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, alice).From, baseDenom)
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(bankA.Int64(), int64(100))

	// try to transfer the token to chainB
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransfer(
			chainA.TxOpts(ctx, alice),
			baseDenom,
			100,
			chainB.CallOpts(ctx, bob).From,
			chanA.PortID, chanA.ID,
			uint64(chainA.LastHeader().Number.Int64())+1000,
		),
	))
	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB))

	// ensure that escrow has correct balance
	escrowBalance, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, alice), chainA.ContractConfig.GetICS20TransferBankAddress(), baseDenom)
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(escrowBalance.Int64(), int64(100))

	// relay the packet
	transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, *transferPacket))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *transferPacket, []byte{1}))

	// ensure that chainB has correct balance
	expectedDenom := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, baseDenom)
	balance, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx, relayer), chainB.CallOpts(ctx, bob).From, expectedDenom)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(100), balance.Int64())

	// try to transfer the token to chainA
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.ICS20Transfer.SendTransfer(
			chainB.TxOpts(ctx, bob),
			expectedDenom,
			100,
			chainA.CallOpts(ctx, alice).From,
			chanB.PortID,
			chanB.ID,
			uint64(chainB.LastHeader().Number.Int64())+1000,
		),
	))
	chainB.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))

	// relay the packet
	transferPacket, err = chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainA, chainB, chanA, chanB, *transferPacket))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *transferPacket, []byte{1}))

	// withdraw tokens from the bank
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Bank.Withdraw(
			chainA.TxOpts(ctx, alice),
			chainA.ContractConfig.GetSimpleTokenAddress(),
			big.NewInt(100),
			chainA.CallOpts(ctx, deployer).From,
		)))

	// ensure that token balance equals original value
	balanceA2, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
	suite.Require().NoError(err)
	suite.Require().Equal(balance0.Int64(), balanceA2.Int64())
}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
