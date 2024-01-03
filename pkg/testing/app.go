package testing

import (
	"context"
	"fmt"
	"math/big"
	"strings"
)

func (c *Coordinator) ApproveAndDepositToken(
	ctx context.Context,
	chain *Chain,
	senderIndex uint32,
	amount uint64,
	receiverIndex uint32,
) error {
	berforeBalance, err := chain.ERC20.BalanceOf(chain.CallOpts(ctx, senderIndex), chain.CallOpts(ctx, senderIndex).From)
	if err != nil {
		return err
	}
	beforeBankBalance, err := chain.ICS20Bank.BalanceOf(
		chain.CallOpts(ctx, receiverIndex),
		chain.CallOpts(ctx, receiverIndex).From,
		strings.ToLower(chain.ContractConfig.ERC20TokenAddress.String()))
	if err != nil {
		return err
	}

	err = chain.WaitIfNoError(ctx, "ERC20::Approve")(
		chain.ERC20.Approve(chain.TxOpts(ctx, senderIndex), chain.ContractConfig.ICS20BankAddress, big.NewInt(int64(amount))),
	)
	if err != nil {
		return err
	}
	err = chain.WaitIfNoError(ctx, "ERC20::Deposit")(chain.ICS20Bank.Deposit(
		chain.TxOpts(ctx, senderIndex),
		chain.ContractConfig.ERC20TokenAddress,
		big.NewInt(int64(amount)),
		chain.CallOpts(ctx, receiverIndex).From,
	))
	if err != nil {
		return err
	}

	// ensure that the balance is reduced
	afterBalance, err := chain.ERC20.BalanceOf(chain.CallOpts(ctx, senderIndex), chain.CallOpts(ctx, senderIndex).From)
	if err != nil {
		return err
	}
	if berforeBalance.Int64()-int64(amount) != afterBalance.Int64() {
		return fmt.Errorf("balance is not reduced")
	}

	// ensure that the bank balance is increased
	afterBankBalance, err := chain.ICS20Bank.BalanceOf(
		chain.CallOpts(ctx, receiverIndex),
		chain.CallOpts(ctx, receiverIndex).From,
		strings.ToLower(chain.ContractConfig.ERC20TokenAddress.String()))
	if err != nil {
		return err
	}
	if beforeBankBalance.Int64()+int64(amount) != afterBankBalance.Int64() {
		return fmt.Errorf("bank balance is not increased")
	}

	return nil
}
