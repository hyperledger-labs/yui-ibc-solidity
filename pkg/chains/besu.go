package chains

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type ParsedHeader struct {
	Base *gethtypes.Header

	Vanity     [32]byte
	Validators []common.Address
	Vote       interface{}
	Round      [4]byte
	Seals      [][]byte
}

func ParseHeader(header *gethtypes.Header) (*ParsedHeader, error) {
	parsed := ParsedHeader{Base: header}

	r := bytes.NewReader(header.Extra)
	stream := rlp.NewStream(r, uint64(len(header.Extra)))
	if _, err := stream.List(); err != nil {
		return nil, err
	}
	if err := stream.Decode(&parsed.Vanity); err != nil {
		return nil, err
	}
	if err := stream.Decode(&parsed.Validators); err != nil {
		return nil, err
	}
	if err := stream.Decode(&parsed.Vote); err != nil {
		return nil, err
	}
	if err := stream.Decode(&parsed.Round); err != nil {
		return nil, err
	}
	if err := stream.Decode(&parsed.Seals); err != nil {
		return nil, err
	}
	if err := stream.ListEnd(); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func (h ParsedHeader) GetSealingHeaderBytes() ([]byte, error) {
	newHeader := *h.Base
	extra, err := rlp.EncodeToBytes([]interface{}{
		h.Vanity, h.Validators, h.Vote, h.Round,
	})
	if err != nil {
		return nil, err
	}
	newHeader.Extra = extra
	return rlp.EncodeToBytes(newHeader)
}

func (h ParsedHeader) GetChainHeaderBytes() ([]byte, error) {
	newHeader := *h.Base
	extra, err := rlp.EncodeToBytes([]interface{}{
		h.Vanity, h.Validators, h.Vote,
	})
	if err != nil {
		return nil, err
	}
	newHeader.Extra = extra
	return rlp.EncodeToBytes(newHeader)
}

func (h ParsedHeader) ValidateAndGetCommitSeals() ([][]byte, error) {
	header, err := h.GetSealingHeaderBytes()
	if err != nil {
		return nil, err
	}
	vals, err := RecoverCommitterAddressesVals(crypto.Keccak256(header), h.Seals)
	if err != nil {
		return nil, err
	}
	var newSeals [][]byte
	count := 0
	for _, val := range h.Validators {
		if seal, ok := vals[val]; ok {
			count++
			newSeals = append(newSeals, seal)
		} else {
			newSeals = append(newSeals, nil)
		}
	}
	if threshold := len(h.Validators) * 2 / 3; count > threshold {
		return newSeals, nil
	} else {
		return nil, fmt.Errorf("insufficient voting: %v > %v", count, threshold)
	}
}
