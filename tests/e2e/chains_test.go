package e2e

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/contract"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"
	testchain0 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain0"
	testchain1 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain1"
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

func (suite ChainTestSuite) TestChannel() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, ibctesting.BesuIBFT2Client)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channeltypes.UNORDERED)

	/// Tests for Transfer module ///

	balanceA0, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From)
	suite.Require().NoError(err)
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.SimpleToken.Approve(chainA.TxOpts(ctx), chainA.ContractConfig.GetICS20BankAddress(), big.NewInt(100)),
	))

	// deposit a simple token to the bank
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.ICS20Bank.Deposit(
		chainA.TxOpts(ctx),
		chainA.ContractConfig.GetSimpleTokenAddress(),
		big.NewInt(100),
		chainA.CallOpts(ctx).From,
	)))

	// ensure that the balance is reduced
	balanceA1, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From)
	suite.Require().NoError(err)
	suite.Require().Equal(balanceA0.Int64()-100, balanceA1.Int64())

	baseDenom := strings.ToLower(chainA.ContractConfig.GetSimpleTokenAddress().String())

	bankA, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From, baseDenom)
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(bankA.Int64(), int64(100))

	// try to transfer the token to chainB
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Transfer.SendTransfer(
			chainA.TxOpts(ctx),
			baseDenom,
			100,
			chainB.CallOpts(ctx).From,
			chanA.PortID, chanA.ID,
			uint64(chainA.LastHeader().Base.Number.Int64())+1000,
		),
	))
	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB, ibctesting.BesuIBFT2Client))

	// ensure that escrow has correct balance
	escrowBalance, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx), chainA.ContractConfig.GetICS20TransferBankAddress(), baseDenom)
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(escrowBalance.Int64(), int64(100))

	// relay the packet
	transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, *transferPacket))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *transferPacket, []byte{1}))

	// ensure that chainB has correct balance
	expectedDenom := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, baseDenom)
	balance, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx), chainB.CallOpts(ctx).From, expectedDenom)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(100), balance.Int64())

	// try to transfer the token to chainA
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.ICS20Transfer.SendTransfer(
			chainB.TxOpts(ctx),
			expectedDenom,
			100,
			chainA.CallOpts(ctx).From,
			chanB.PortID,
			chanB.ID,
			uint64(chainB.LastHeader().Base.Number.Int64())+1000,
		),
	))
	chainB.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA, ibctesting.BesuIBFT2Client))

	// relay the packet
	transferPacket, err = chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainA, chainB, chanA, chanB, *transferPacket))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *transferPacket, []byte{1}))

	// withdraw tokens from the bank
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.ICS20Bank.Withdraw(
			chainA.TxOpts(ctx),
			chainA.ContractConfig.GetSimpleTokenAddress(),
			big.NewInt(100),
			chainA.CallOpts(ctx).From,
		)))

	// ensure that token balance equals original value
	balanceA2, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From)
	suite.Require().NoError(err)
	suite.Require().Equal(balanceA0.Int64(), balanceA2.Int64())
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ChainTestSuite))
}
