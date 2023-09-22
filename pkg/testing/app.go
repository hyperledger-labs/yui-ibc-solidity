package testing

import (
	"context"
	"math/big"
)

func (chain *Chain) ApproveAndDepositToken(
	ctx context.Context,
	senderIndex uint32,
	amount uint64,
	receiverIndex uint32,
) error {
	err := chain.WaitIfNoError(ctx)(
		chain.ERC20.Approve(chain.TxOpts(ctx, senderIndex), chain.ContractConfig.ICS20BankAddress, big.NewInt(int64(amount))),
	)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(chain.ICS20Bank.Deposit(
		chain.TxOpts(ctx, senderIndex),
		chain.ContractConfig.ERC20TokenAddress,
		big.NewInt(int64(amount)),
		chain.CallOpts(ctx, receiverIndex).From,
	))
}
