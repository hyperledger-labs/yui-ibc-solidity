package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	endpoint string

	conn *rpc.Client
	ETHClient
}

func (cl Client) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction) (Receipt, error) {
	var receipt Receipt
	err := retry.Do(
		func() error {
			rc, err := cl.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return err
			}
			receipt = rc
			return nil
		},
		retry.Delay(1*time.Second),
		retry.Attempts(10),
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

type ETHClient interface {
	bind.ContractBackend
	BlockByNumber(ctx context.Context, bn *big.Int) (*gethtypes.Block, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (Receipt, error)
}

type Receipt interface {
	PostState() []byte
	Status() uint64
	CumulativeGasUsed() uint64
	Bloom() gethtypes.Bloom
	Logs() []*gethtypes.Log
	TxHash() common.Hash
	ContractAddress() common.Address
	GasUsed() uint64
	BlockHash() common.Hash
	BlockNumber() *big.Int
	TransactionIndex() uint
	RevertReason() string
}

type GenTxOpts func(ctx context.Context) *bind.TransactOpts

func MakeGenTxOpts(chainID *big.Int, prv *ecdsa.PrivateKey) GenTxOpts {
	signer := gethtypes.NewEIP155Signer(chainID)
	addr := gethcrypto.PubkeyToAddress(prv.PublicKey)
	return func(ctx context.Context) *bind.TransactOpts {
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
}
