package testing

import (
	"context"
	"math/big"
)

func (c *Coordinator) Approve(
	ctx context.Context,
	chain *Chain,
	senderIndex uint32,
	amount uint64,
) error {
	return chain.WaitIfNoError(ctx, "ERC20::Approve")(
		chain.ERC20.Approve(chain.TxOpts(ctx, senderIndex), chain.ContractConfig.ICS20TransferAddress, big.NewInt(int64(amount))),
	)
}
