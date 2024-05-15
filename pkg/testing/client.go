package testing

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/chains"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
)

type LightClient struct {
	client        *client.ETHClient
	consensusType chains.ConsensusType
}

func NewLightClient(cl *client.ETHClient, consensusType chains.ConsensusType) *LightClient {
	return &LightClient{client: cl, consensusType: consensusType}
}

func (lc LightClient) GenerateInputData(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (*LightClientInputData, error) {
	var state LightClientInputData
	block, err := lc.client.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	proof, err := lc.client.GetProof(address, storageKeys, block.Number())
	if err != nil {
		return nil, err
	}
	state.StateProof = proof
	state.ParsedHeader, err = chains.ParseHeader(block.Header())
	if err != nil {
		return nil, err
	}
	state.CommitSeals, err = state.ParsedHeader.ValidateAndGetCommitSeals(lc.consensusType)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

type LightClientInputData struct {
	ParsedHeader *chains.ParsedHeader
	StateProof   *client.StateProof
	CommitSeals  [][]byte
}

func (cs LightClientInputData) Header() *gethtypes.Header {
	return cs.ParsedHeader.Base
}

func (cs LightClientInputData) MembershipProof() *client.StateProof {
	return cs.StateProof
}

func (cs LightClientInputData) SealingHeaderRLP(consensusType chains.ConsensusType) []byte {
	bz, err := cs.ParsedHeader.GetSealingHeaderBytes(consensusType)
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs LightClientInputData) GetCommitSeals() [][]byte {
	return cs.CommitSeals
}

func (cs LightClientInputData) Validators() [][]byte {
	var addrs [][]byte
	for _, val := range cs.ParsedHeader.Validators {
		addrs = append(addrs, val.Bytes())
	}
	return addrs
}
