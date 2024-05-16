package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/chains"
)

type ETHClient struct {
	*ethclient.Client
	rpcClient *rpc.Client
	erepo     ErrorsRepository
	option    option
}

type Option func(*option)

type option struct {
	retryOpts []retry.Option
}

func DefaultOption() *option {
	return &option{
		retryOpts: []retry.Option{
			retry.Delay(1 * time.Second),
			retry.Attempts(10),
		},
	}
}

func WithRetryOption(rops ...retry.Option) Option {
	return func(opt *option) {
		opt.retryOpts = rops
	}
}

func NewETHClient(endpoint string, erepo ErrorsRepository, opts ...Option) (*ETHClient, error) {
	rpcClient, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	opt := DefaultOption()
	for _, o := range opts {
		o(opt)
	}
	return &ETHClient{
		rpcClient: rpcClient,
		Client:    ethclient.NewClient(rpcClient),
		erepo:     erepo,
		option:    *opt,
	}, nil
}

func (cl *ETHClient) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*gethtypes.Receipt, bool, error) {
	var r *Receipt
	if err := cl.rpcClient.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash); err != nil {
		return r.GetGethReceipt(), true, err
	}
	if r == nil {
		return nil, true, ethereum.NotFound
	} else if r.Status == gethtypes.ReceiptStatusSuccessful {
		return r.GetGethReceipt(), false, nil
	} else if r.HasRevertReason() {
		e, args, err := cl.erepo.ParseError(r.RevertReason)
		if err == nil {
			return r.GetGethReceipt(), false, fmt.Errorf("revert-reason=\"%v\" args=\"%v\"", e.String(), args)
		} else {
			return r.GetGethReceipt(), false, fmt.Errorf("raw-revert-reason=\"%x\" parse-err=\"%v\"", []byte(r.RevertReason), err)
		}
	} else {
		return r.GetGethReceipt(), false, fmt.Errorf("failed to execute a transaction: %v", r)
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
		cl.option.retryOpts...,
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (cl *ETHClient) DetectBesuConsensusType() (chains.ConsensusType, error) {
	var res []string
	if err := cl.rpcClient.CallContext(context.TODO(), &res, "ibft_getValidatorsByBlockNumber", "latest"); err == nil {
		return chains.IBFT2, nil
	} else {
		if err.Error() != "Method not enabled" {
			return chains.Unspecified, err
		}
	}
	if err := cl.rpcClient.CallContext(context.TODO(), &res, "qbft_getValidatorsByBlockNumber", "latest"); err == nil {
		return chains.QBFT, nil
	} else {
		if err.Error() != "Method not enabled" {
			return chains.Unspecified, err
		}
	}
	return chains.Unspecified, errors.New("failed to detect consensus type")
}

type ErrorsRepository struct {
	errs map[[4]byte]abi.Error
}

func NewErrorsRepository() ErrorsRepository {
	er := ErrorsRepository{
		errs: make(map[[4]byte]abi.Error),
	}
	for _, e := range builtInErrors() {
		if err := er.Add(e); err != nil {
			panic(err)
		}
	}
	return er
}

func (r *ErrorsRepository) Add(e0 abi.Error) error {
	var sel [4]byte
	copy(sel[:], e0.ID[:4])
	if e1, ok := r.errs[sel]; ok {
		if e1.Sig == e0.Sig {
			return nil
		}
		return fmt.Errorf("error selector collision: sel=%x e0=%v e1=%v", sel, e0, e1)
	}
	r.errs[sel] = e0
	return nil
}

func (r *ErrorsRepository) GetError(sel [4]byte) (abi.Error, bool) {
	e, ok := r.errs[sel]
	return e, ok
}

func (r *ErrorsRepository) ParseError(bz []byte) (abi.Error, interface{}, error) {
	if len(bz) < 4 {
		return abi.Error{}, nil, fmt.Errorf("invalid error data: %v", bz)
	}
	var sel [4]byte
	copy(sel[:], bz[:4])
	e, ok := r.GetError(sel)
	if !ok {
		return abi.Error{}, nil, fmt.Errorf("unknown error: sel=%x", sel)
	}
	v, err := e.Unpack(bz)
	return e, v, err
}

// builtInErrors returns a list of solidity built-in errors.
// The list includes "Error(string)" and "Panic(uint256)" errors.
func builtInErrors() []abi.Error {
	var errors []abi.Error
	strT, err := abi.NewType("string", "", nil)
	if err != nil {
		panic(err)
	}
	errors = append(errors, abi.NewError("Error", abi.Arguments{{Type: strT}}))

	uint256T, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}
	errors = append(errors, abi.NewError("Panic", abi.Arguments{{Type: uint256T}}))

	return errors
}
