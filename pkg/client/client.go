package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type ETHClient struct {
	*ethclient.Client
	rpcClient *rpc.Client
}

func NewETHClient(endpoint string) (*ETHClient, error) {
	rpcClient, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return &ETHClient{
		rpcClient: rpcClient,
		Client:    ethclient.NewClient(rpcClient),
	}, nil
}

func (cl *ETHClient) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *Receipt
	err := cl.rpcClient.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		} else if len(r.RevertReason_) > 0 {
			return nil, fmt.Errorf("revert: %v", r.RevertReason())
		}
	}
	return &r.Receipt, err
}

func (cl *ETHClient) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction) (*types.Receipt, error) {
	var receipt *types.Receipt
	err := retry.Do(
		func() error {
			rc, err := cl.GetTransactionReceipt(ctx, tx.Hash())
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

type Receipt struct {
	types.Receipt
	RevertReason_ []byte `json:"revertReason,omitempty"`
}

func (rc Receipt) RevertReason() string {
	reason, err := parseRevertReason(rc.RevertReason_)
	if err != nil {
		panic(err)
	}
	return reason
}

// A format of revertReason is:
// 4byte: Function selector for Error(string)
// 32byte: Data offset
// 32byte: String length
// Remains: String Data
func parseRevertReason(bz []byte) (string, error) {
	if l := len(bz); l == 0 {
		return "", nil
	} else if l < 68 {
		return "", fmt.Errorf("invalid length")
	}

	size := &big.Int{}
	size.SetBytes(bz[36:68])
	return string(bz[68 : 68+size.Int64()]), nil
}
