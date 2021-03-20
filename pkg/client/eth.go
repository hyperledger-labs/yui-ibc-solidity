package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewETHClient(endpoint string) (*Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{
		endpoint:  endpoint,
		conn:      conn,
		ETHClient: ethClient{Client: ethclient.NewClient(conn)},
	}, nil
}

type ethReceipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	PostState_         []byte       `json:"root"`
	Status_            uint64       `json:"status"`
	CumulativeGasUsed_ uint64       `json:"cumulativeGasUsed" gencodec:"required"`
	Bloom_             types.Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs_              []*types.Log `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash_          common.Hash    `json:"transactionHash" gencodec:"required"`
	ContractAddress_ common.Address `json:"contractAddress"`
	GasUsed_         uint64         `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash_        common.Hash `json:"blockHash,omitempty"`
	BlockNumber_      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex_ uint        `json:"transactionIndex"`
}

func (rc ethReceipt) PostState() []byte {
	return rc.PostState_
}

func (rc ethReceipt) Status() uint64 {
	return rc.Status_
}

func (rc ethReceipt) CumulativeGasUsed() uint64 {
	return rc.CumulativeGasUsed_
}

func (rc ethReceipt) Bloom() types.Bloom {
	return rc.Bloom_
}

func (rc ethReceipt) Logs() []*types.Log {
	return rc.Logs_
}

func (rc ethReceipt) TxHash() common.Hash {
	return rc.TxHash_
}

func (rc ethReceipt) ContractAddress() common.Address {
	return rc.ContractAddress_
}

func (rc ethReceipt) GasUsed() uint64 {
	return rc.GasUsed_
}

func (rc ethReceipt) BlockHash() common.Hash {
	return rc.BlockHash_
}

func (rc ethReceipt) BlockNumber() *big.Int {
	return rc.BlockNumber_
}

func (rc ethReceipt) TransactionIndex() uint {
	return rc.TransactionIndex_
}

func (rc ethReceipt) RevertReason() string {
	return ""
}

type ethClient struct {
	*ethclient.Client
}

func (cl ethClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (Receipt, error) {
	rc, err := cl.Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return ethReceipt{
		PostState_:         rc.PostState,
		Status_:            rc.Status,
		CumulativeGasUsed_: rc.CumulativeGasUsed,
		Bloom_:             rc.Bloom,
		Logs_:              rc.Logs,
		TxHash_:            rc.TxHash,
		ContractAddress_:   rc.ContractAddress,
		GasUsed_:           rc.GasUsed,
		BlockHash_:         rc.BlockHash,
		BlockNumber_:       rc.BlockNumber,
		TransactionIndex_:  rc.TransactionIndex,
	}, nil
}
