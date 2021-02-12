package contract

import (
	"context"
	"math/big"

	"github.com/datachainlab/ibc-solidity/pkg/chains"
	"github.com/ethereum/go-ethereum/common"
)

func (cl Client) GetContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (*ContractState, error) {
	var state ContractState
	block, err := cl.Client.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	proof, err := cl.GetETHProof(address, storageKeys, block.Number())
	if err != nil {
		return nil, err
	}
	state.ETHProof = proof
	state.ParsedHeader, err = chains.ParseHeader(block.Header())
	if err != nil {
		return nil, err
	}
	state.CommitSeals, err = state.ParsedHeader.ValidateAndGetCommitSeals()
	if err != nil {
		return nil, err
	}
	return &state, nil
}

type ContractState struct {
	ParsedHeader *chains.ParsedHeader
	ETHProof     *ETHProof
	CommitSeals  [][]byte
}

func (cs ContractState) ChainHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetChainHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs ContractState) SealingHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetSealingHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs ContractState) AccountProofRLP() []byte {
	return cs.ETHProof.AccountProofRLP
}

func (cs ContractState) StorageProofRLP(idx int) []byte {
	return cs.ETHProof.StorageProofRLP[idx]
}

func (cs ContractState) GetCommitSeals() [][]byte {
	return cs.CommitSeals
}
