package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type StateProof struct {
	Balance     big.Int
	CodeHash    [32]byte
	Nonce       uint64
	StorageHash [32]byte

	AccountProofRLP []byte
	StorageProofRLP [][]byte
}

func (cl ETHClient) GetProof(address common.Address, storageKeys [][]byte, blockNumber *big.Int) (*StateProof, error) {
	bz, err := cl.getProof(address, storageKeys, "0x"+blockNumber.Text(16))
	if err != nil {
		return nil, err
	}
	var proof struct {
		Balance      string   `json:"balance"`
		CodeHash     string   `json:"codeHash"`
		Nonce        string   `json:"nonce"`
		StorageHash  string   `json:"storageHash"`
		AccountProof []string `json:"accountProof"`
		StorageProof []struct {
			Proof []string `json:"proof"`
		} `json:"storageProof"`
	}
	if err := json.Unmarshal(bz, &proof); err != nil {
		return nil, err
	}

	var balance big.Int
	{
		bz, err := decodeHexString(proof.Balance)
		if err != nil {
			return nil, err
		}
		balance.SetBytes(bz)
	}
	var codeHash [32]byte
	{
		bz, err := decodeHexString(proof.CodeHash)
		if err != nil {
			return nil, err
		}
		copy(codeHash[:], bz)
	}
	var nonce big.Int
	{
		bz, err := decodeHexString(proof.Nonce)
		if err != nil {
			return nil, err
		}
		nonce.SetBytes(bz)
	}
	var storageHash [32]byte
	{
		bz, err := decodeHexString(proof.StorageHash)
		if err != nil {
			return nil, err
		}
		copy(storageHash[:], bz)
	}

	var encodedProof = StateProof{
		Balance:     balance,
		CodeHash:    codeHash,
		Nonce:       nonce.Uint64(),
		StorageHash: storageHash,
	}
	encodedProof.AccountProofRLP, err = encodeRLP(proof.AccountProof)
	if err != nil {
		return nil, err
	}
	for _, p := range proof.StorageProof {
		bz, err := encodeRLP(p.Proof)
		if err != nil {
			return nil, err
		}
		encodedProof.StorageProofRLP = append(encodedProof.StorageProofRLP, bz)
	}
	return &encodedProof, nil
}

func (cl ETHClient) getProof(address common.Address, storageKeys [][]byte, blockNumber string) ([]byte, error) {
	hashes := []common.Hash{}
	for _, k := range storageKeys {
		var h common.Hash
		if err := h.UnmarshalText(k); err != nil {
			return nil, err
		}
		hashes = append(hashes, h)
	}
	var msg json.RawMessage
	if err := cl.rpcClient.Call(&msg, "eth_getProof", address, hashes, blockNumber); err != nil {
		return nil, err
	}
	return msg, nil
}

func encodeRLP(proof []string) ([]byte, error) {
	var target [][][]byte
	for _, p := range proof {
		bz, err := decodeHexString(p)
		if err != nil {
			return nil, err
		}
		var val [][]byte
		if err := rlp.DecodeBytes(bz, &val); err != nil {
			return nil, err
		}
		target = append(target, val)
	}
	bz, err := rlp.EncodeToBytes(target)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func decodeHexString(s string) ([]byte, error) {
	if !strings.HasPrefix(s, "0x") {
		return nil, fmt.Errorf("missing prefix '0x': %v", s)
	}
	s = s[2:]
	if len(s)%2 != 0 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}
