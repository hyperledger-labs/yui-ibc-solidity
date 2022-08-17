package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func (cl *ETHClient) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (rc *gethtypes.Receipt, recoverable bool, err error) {
	var r *Receipt
	if err := cl.rpcClient.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash); err != nil {
		return &r.Receipt, true, err
	}
	if r == nil {
		return nil, true, ethereum.NotFound
	} else if r.Status == gethtypes.ReceiptStatusSuccessful {
		return &r.Receipt, false, nil
	} else if r.HasRevertReason() {
		reason, err := r.GetRevertReason()
		return &r.Receipt, false, fmt.Errorf("revert: %v(parse-err=%v)", reason, err)
	} else {
		return &r.Receipt, false, fmt.Errorf("failed to execute a transaction: %v", r)
	}
}

func (cl *ETHClient) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction) (*gethtypes.Receipt, error) {
	var receipt *gethtypes.Receipt
	err := retry.Do(
		func() error {
			rc, recoverable, err := cl.GetTransactionReceipt(ctx, tx.Hash())
			if err != nil {
				if recoverable {
					return err
				} else {
					return retry.Unrecoverable(err)
				}
			}
			receipt = rc
			return nil
		},
		// TODO make these configurable
		retry.Delay(1*time.Second),
		retry.Attempts(10),
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

type Receipt struct {
	gethtypes.Receipt
	RevertReason []byte `json:"revertReason,omitempty"`
}

func (rc Receipt) HasRevertReason() bool {
	return len(rc.RevertReason) > 0
}

func (rc Receipt) GetRevertReason() (string, error) {
	return parseRevertReason(rc.RevertReason)
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
