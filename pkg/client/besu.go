package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewBesuClient(endpoint string, clientType string) (*Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{
		endpoint:   endpoint,
		clientType: clientType,
		conn:       conn,
		ETHClient:  besuClient{Client: ethclient.NewClient(conn), rpcClient: conn},
	}, nil
}

type besuClient struct {
	*ethclient.Client
	rpcClient *rpc.Client
}

func (cl besuClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (Receipt, error) {
	var r *besuReceipt
	err := cl.rpcClient.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

type besuReceipt struct {
	ethReceipt
	RevertReason_ []byte `json:"revertReason"`
}

func (rc besuReceipt) RevertReason() string {
	reason, err := parseRevertReason(rc.RevertReason_)
	if err != nil {
		panic(err)
	}
	return reason
}

// MarshalJSON marshals as JSON.
func (r besuReceipt) MarshalJSON() ([]byte, error) {
	type receipt struct {
		PostState         hexutil.Bytes  `json:"root"`
		Status            hexutil.Uint64 `json:"status"`
		CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             types.Bloom    `json:"logsBloom"         gencodec:"required"`
		Logs              []*types.Log   `json:"logs"              gencodec:"required"`
		TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   common.Address `json:"contractAddress"`
		GasUsed           hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		BlockHash         common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big   `json:"blockNumber,omitempty"`
		TransactionIndex  hexutil.Uint   `json:"transactionIndex"`
		RevertReason      hexutil.Bytes  `json:"revertReason"`
	}
	var enc receipt
	enc.PostState = r.PostState_
	enc.Status = hexutil.Uint64(r.Status_)
	enc.CumulativeGasUsed = hexutil.Uint64(r.CumulativeGasUsed_)
	enc.Bloom = r.Bloom_
	enc.Logs = r.Logs_
	enc.TxHash = r.TxHash_
	enc.ContractAddress = r.ContractAddress_
	enc.GasUsed = hexutil.Uint64(r.GasUsed_)
	enc.BlockHash = r.BlockHash_
	enc.BlockNumber = (*hexutil.Big)(r.BlockNumber_)
	enc.TransactionIndex = hexutil.Uint(r.TransactionIndex_)
	enc.RevertReason = r.RevertReason_
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *besuReceipt) UnmarshalJSON(input []byte) error {
	type receipt struct {
		PostState         *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             *types.Bloom    `json:"logsBloom"         gencodec:"required"`
		Logs              []*types.Log    `json:"logs"              gencodec:"required"`
		TxHash            *common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   *common.Address `json:"contractAddress"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		BlockHash         *common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
		RevertReason      *hexutil.Bytes  `json:"revertReason"`
	}
	var dec receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.PostState != nil {
		r.PostState_ = *dec.PostState
	}
	if dec.Status != nil {
		r.Status_ = uint64(*dec.Status)
	}
	if dec.CumulativeGasUsed == nil {
		return errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	}
	r.CumulativeGasUsed_ = uint64(*dec.CumulativeGasUsed)
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom_ = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs_ = dec.Logs
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash_ = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress_ = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed_ = uint64(*dec.GasUsed)
	if dec.BlockHash != nil {
		r.BlockHash_ = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		r.BlockNumber_ = (*big.Int)(dec.BlockNumber)
	}
	if dec.TransactionIndex != nil {
		r.TransactionIndex_ = uint(*dec.TransactionIndex)
	}
	if dec.RevertReason != nil {
		r.RevertReason_ = *dec.RevertReason
	}
	return nil
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
