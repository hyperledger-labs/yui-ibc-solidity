package client

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Receipt definition from geth v1.11.6
// Modified to add RevertReason field
// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8            `json:"type,omitempty"`
	PostState         []byte           `json:"root"`
	Status            uint64           `json:"status"`
	CumulativeGasUsed uint64           `json:"cumulativeGasUsed" gencodec:"required"`
	Bloom             gethtypes.Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs              []*gethtypes.Log `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
	ContractAddress   common.Address `json:"contractAddress"`
	GasUsed           uint64         `json:"gasUsed" gencodec:"required"`
	EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint        `json:"transactionIndex"`

	RevertReason hexutil.Bytes `json:"revertReason,omitempty"`
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		Type              *hexutil.Uint64  `json:"type,omitempty"`
		PostState         *hexutil.Bytes   `json:"root"`
		Status            *hexutil.Uint64  `json:"status"`
		CumulativeGasUsed *hexutil.Uint64  `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             *gethtypes.Bloom `json:"logsBloom"         gencodec:"required"`
		Logs              []*gethtypes.Log `json:"logs"              gencodec:"required"`
		TxHash            *common.Hash     `json:"transactionHash" gencodec:"required"`
		ContractAddress   *common.Address  `json:"contractAddress"`
		GasUsed           *hexutil.Uint64  `json:"gasUsed" gencodec:"required"`
		EffectiveGasPrice *hexutil.Big     `json:"effectiveGasPrice,omitempty"`
		BlockHash         *common.Hash     `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big     `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint    `json:"transactionIndex"`
		RevertReason      *hexutil.Bytes   `json:"revertReason,omitempty"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type != nil {
		r.Type = uint8(*dec.Type)
	}
	if dec.PostState != nil {
		r.PostState = *dec.PostState
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.CumulativeGasUsed == nil {
		return errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	}
	r.CumulativeGasUsed = uint64(*dec.CumulativeGasUsed)
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*dec.GasUsed)
	if dec.EffectiveGasPrice != nil {
		r.EffectiveGasPrice = (*big.Int)(dec.EffectiveGasPrice)
	}
	if dec.BlockHash != nil {
		r.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		r.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.TransactionIndex != nil {
		r.TransactionIndex = uint(*dec.TransactionIndex)
	}
	if dec.RevertReason != nil {
		r.RevertReason = *dec.RevertReason
	}
	return nil
}

func (rc Receipt) GetGethReceipt() *gethtypes.Receipt {
	return &gethtypes.Receipt{
		Type:              rc.Type,
		PostState:         rc.PostState,
		Status:            rc.Status,
		CumulativeGasUsed: rc.CumulativeGasUsed,
		Bloom:             rc.Bloom,
		Logs:              rc.Logs,
		TxHash:            rc.TxHash,
		ContractAddress:   rc.ContractAddress,
		GasUsed:           rc.GasUsed,
		EffectiveGasPrice: rc.EffectiveGasPrice,
		BlockHash:         rc.BlockHash,
		BlockNumber:       rc.BlockNumber,
		TransactionIndex:  rc.TransactionIndex,
	}
}

func (rc Receipt) HasRevertReason() bool {
	return len(rc.RevertReason) > 0
}
