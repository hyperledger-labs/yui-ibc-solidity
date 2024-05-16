package client

import (
	"math/big"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

// client type
const (
	BesuQBFTClient = "hb-qbft"
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
