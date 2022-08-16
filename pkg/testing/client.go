package testing

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/chains"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
)

type LightClient struct {
	client     *client.ETHClient
	clientType string
}

func NewLightClient(cl *client.ETHClient, clientType string) *LightClient {
	return &LightClient{client: cl, clientType: clientType}
}

type LightClientState interface {
	Header() *gethtypes.Header
	Proof() *client.StateProof
}

func (lc LightClient) ClientType() string {
	return lc.clientType
}

func (lc LightClient) GetState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (LightClientState, error) {
	switch lc.clientType {
	case ibcclient.BesuIBFT2Client:
		return lc.GetIBFT2State(ctx, address, storageKeys, bn)
	case ibcclient.MockClient:
		return lc.GetMockContractState(ctx, address, storageKeys, bn)
	default:
		panic(fmt.Sprintf("unknown client type '%v'", lc.clientType))
	}
}

func (lc LightClient) GetMockContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (LightClientState, error) {
	block, err := lc.client.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	// this is dummy
	proof := &client.StateProof{
		StorageProofRLP: make([][]byte, len(storageKeys)),
	}
	return ETHState{header: block.Header(), StateProof: proof}, nil
}

func (lc LightClient) GetIBFT2State(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (LightClientState, error) {
	var state IBFT2State
	block, err := lc.client.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	proof, err := lc.client.GetStateProof(address, storageKeys, block.Number())
	if err != nil {
		return nil, err
	}
	state.StateProof = proof
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

type ETHState struct {
	header     *gethtypes.Header
	StateProof *client.StateProof
}

var _ LightClientState = (*ETHState)(nil)

func (cs ETHState) Header() *gethtypes.Header {
	return cs.header
}

func (cs ETHState) Proof() *client.StateProof {
	return cs.StateProof
}

type IBFT2State struct {
	ParsedHeader *chains.ParsedHeader
	StateProof   *client.StateProof
	CommitSeals  [][]byte
}

func (cs IBFT2State) Header() *gethtypes.Header {
	return cs.ParsedHeader.Base
}

func (cs IBFT2State) Proof() *client.StateProof {
	return cs.StateProof
}

func (cs IBFT2State) ChainHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetChainHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs IBFT2State) SealingHeaderRLP() []byte {
	bz, err := cs.ParsedHeader.GetSealingHeaderBytes()
	if err != nil {
		panic(err)
	}
	return bz
}

func (cs IBFT2State) GetCommitSeals() [][]byte {
	return cs.CommitSeals
}

func (cs IBFT2State) Validators() [][]byte {
	var addrs [][]byte
	for _, val := range cs.ParsedHeader.Validators {
		addrs = append(addrs, val.Bytes())
	}
	return addrs
}
