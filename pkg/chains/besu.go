package chains

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type ConsensusType uint8

const (
	Unspecified ConsensusType = iota
	IBFT2
	QBFT
)

func (c ConsensusType) String() string {
	switch c {
	case IBFT2:
		return "IBFT2"
	case QBFT:
		return "QBFT"
	default:
		return "Unspecified"
	}
}

type ParsedHeader struct {
	Base *gethtypes.Header

	Vanity     []byte
	Validators []common.Address
	Vote       interface{}
	Round      []byte
	Seals      [][]byte
}

func ParseHeader(header *gethtypes.Header) (*ParsedHeader, error) {
	parsed := ParsedHeader{Base: header}

	r := bytes.NewReader(header.Extra)
	stream := rlp.NewStream(r, 0)
	if _, err := stream.List(); err != nil {
		return nil, fmt.Errorf("failed to decode list: %w", err)
	}
	if err := stream.Decode(&parsed.Vanity); err != nil {
		return nil, fmt.Errorf("failed to decode vanity: %w", err)
	}
	if err := stream.Decode(&parsed.Validators); err != nil {
		return nil, fmt.Errorf("failed to decode validators: %w", err)
	}
	if err := stream.Decode(&parsed.Vote); err != nil {
		return nil, fmt.Errorf("failed to decode vote: %w", err)
	}
	if err := stream.Decode(&parsed.Round); err != nil {
		return nil, fmt.Errorf("failed to decode round: %w", err)
	}
	if err := stream.Decode(&parsed.Seals); err != nil {
		return nil, fmt.Errorf("failed to decode seals: %w", err)
	}
	if err := stream.ListEnd(); err != nil {
		return nil, fmt.Errorf("failed to decode list end: %w", err)
	}

	return &parsed, nil
}

func (h ParsedHeader) GetSealingHeaderBytes(consensusType ConsensusType) ([]byte, error) {
	newHeader := *h.Base
	// IBFT2: {Vanity, Validators, Vote, Round}
	// QBFT: {Vanity, Validators, Vote, Round, Empty-Seals}
	if consensusType == IBFT2 {
		extra, err := rlp.EncodeToBytes([]interface{}{
			h.Vanity, h.Validators, h.Vote, h.Round,
		})
		if err != nil {
			return nil, err
		}
		newHeader.Extra = extra
	} else if consensusType == QBFT {
		extra, err := rlp.EncodeToBytes([]interface{}{
			h.Vanity, h.Validators, h.Vote, h.Round, [][]byte{},
		})
		if err != nil {
			return nil, err
		}
		newHeader.Extra = extra
	} else {
		return nil, fmt.Errorf("unsupported consensus type: %v", consensusType)
	}
	return rlp.EncodeToBytes(&newHeader)
}

func (h ParsedHeader) ValidateAndGetCommitSeals(consensusType ConsensusType) ([][]byte, error) {
	header, err := h.GetSealingHeaderBytes(consensusType)
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
