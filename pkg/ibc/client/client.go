package client

import (
	"math/big"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

// client type
const (
	// IBFT2 Client
	BesuIBFT2Client = "hyperledger-besu-ibft2"
	// NOTE: The mock client is only intended for use in development such as ganache.
	MockClient = "mock-client"
)

func NewHeightFromBN(n *big.Int) Height {
	return Height{
		RevisionNumber: 0,
		RevisionHeight: n.Uint64(),
	}
}

func (h *Height) ToBN() *big.Int {
	if h.RevisionNumber != 0 {
		panic("revision number must be zero")
	}
	return big.NewInt(int64(h.RevisionHeight))
}

func (h *Height) ToCallData() ibchandler.HeightData {
	return ibchandler.HeightData{
		RevisionNumber: h.RevisionNumber,
		RevisionHeight: h.RevisionHeight,
	}
}
