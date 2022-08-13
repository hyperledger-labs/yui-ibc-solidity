package ethereum

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewETHClient(endpoint string) (*ethclient.Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(conn), nil
}

func (chain *Chain) CallOpts(ctx context.Context, height int64) *bind.CallOpts {
	opts := &bind.CallOpts{
		From:    gethcrypto.PubkeyToAddress(chain.relayerPrvKey.PublicKey),
		Context: ctx,
	}
	if height > 0 {
		opts.BlockNumber = big.NewInt(height)
	}
	return opts
}

func (chain *Chain) TxOpts(ctx context.Context) *bind.TransactOpts {
	signer := gethtypes.NewEIP155Signer(chain.chainID)
	prv := chain.relayerPrvKey
	addr := gethcrypto.PubkeyToAddress(prv.PublicKey)
	return &bind.TransactOpts{
		From:     addr,
		GasLimit: 6382056,
		Signer: func(address common.Address, tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
			if address != addr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := gethcrypto.Sign(signer.Hash(tx).Bytes(), prv)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
}
