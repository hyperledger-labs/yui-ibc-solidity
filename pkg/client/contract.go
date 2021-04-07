package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/datachainlab/ibc-solidity/pkg/chains"
	ibcclient "github.com/datachainlab/ibc-solidity/pkg/ibc/client"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ContractState interface {
	Header() *gethtypes.Header
	ETHProof() *ETHProof
}

func (cl Client) GetContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (ContractState, error) {
	switch cl.clientType {
	case ibcclient.BesuIBFT2Client:
		return cl.GetIBFT2ContractState(ctx, address, storageKeys, bn)
	case ibcclient.MockClient:
		return cl.GetETHContractState(ctx, address, storageKeys, bn)
	default:
		panic(fmt.Sprintf("unknown client type '%v'", cl.clientType))
	}
}

func (cl Client) GetETHContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (ContractState, error) {
	block, err := cl.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	proof, err := cl.GetETHProof(address, storageKeys, block.Number())
	if err != nil {
		return nil, err
	}
	return ETHContractState{header: block.Header(), ethProof: proof}, nil
}

func (cl Client) GetIBFT2ContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (ContractState, error) {
	var state IBFT2ContractState
	block, err := cl.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	proof, err := cl.GetETHProof(address, storageKeys, block.Number())
	if err != nil {
		return nil, err
	}
	state.ethProof = proof
	state.ParsedHeader, err = chains.ParseHeader(block.Header())
	if err != nil {
		return nil, err
	}
	state.CommitSeals, err = state.ParsedHeader.ValidateAndGetCommitSeals()
	if err != nil {
		return nil, err
	}
	return state, nil
}

type ETHContractState struct {
	header   *gethtypes.Header
	ethProof *ETHProof
}

func (cs ETHContractState) Header() *gethtypes.Header {
	return cs.header
}

func (cs ETHContractState) ETHProof() *ETHProof {
	return cs.ethProof
}

type IBFT2ContractState struct {
	ParsedHeader *chains.ParsedHeader
	ethProof     *ETHProof
	CommitSeals  [][]byte
}

func (cs IBFT2ContractState) Header() *gethtypes.Header {
	return cs.ParsedHeader.Base
}

func (cs IBFT2ContractState) ETHProof() *ETHProof {
	return cs.ethProof
}

func (cs IBFT2ContractState) ChainHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetChainHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs IBFT2ContractState) SealingHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetSealingHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs IBFT2ContractState) GetCommitSeals() [][]byte {
	return cs.CommitSeals
}

func (cs IBFT2ContractState) Validators() [][]byte {
	var addrs [][]byte
	for _, val := range cs.ParsedHeader.Validators {
		addrs = append(addrs, val.Bytes())
	}
	return addrs
}
