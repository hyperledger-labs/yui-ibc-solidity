package contract

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	web3 "github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/jsonrpc"
)

type ETHProof struct {
	AccountProofRLP []byte
	StorageProofRLP [][]byte
}

func (cl Client) GetETHProof(address common.Address, storageKeys [][]byte, blockNumber *big.Int) (*ETHProof, error) {
	bz, err := cl.getProof(address, storageKeys, "0x"+blockNumber.Text(16))
	if err != nil {
		return nil, err
	}
	var proof struct {
		AccountProof []string `json:"accountProof"`
		StorageProof []struct {
			Proof []string `json:"proof"`
		} `json:"storageProof"`
	}
	if err := json.Unmarshal(bz, &proof); err != nil {
		return nil, err
	}

	var encodedProof ETHProof
	encodedProof.AccountProofRLP, err = encodeRLP(proof.AccountProof)
	if err != nil {
		return nil, err
	}
	for _, p := range proof.StorageProof {
		bz, err := encodeRLP(p.Proof)
		if err != nil {
			panic(err)
		}
		encodedProof.StorageProofRLP = append(encodedProof.StorageProofRLP, bz)
	}

	return &encodedProof, nil
}

func (cl Client) getProof(address common.Address, storageKeys [][]byte, blockNumber string) ([]byte, error) {
	jc, err := jsonrpc.NewClient(cl.endpoint)
	if err != nil {
		return nil, err
	}
	defer jc.Close()
	hashes := []web3.Hash{}
	for _, k := range storageKeys {
		var h web3.Hash
		err = h.UnmarshalText(k)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, h)
	}
	var msg json.RawMessage
	if err := jc.Call("eth_getProof", &msg, address, hashes, blockNumber); err != nil {
		return nil, err
	}
	return msg, nil
}

func encodeRLP(proof []string) ([]byte, error) {
	var target [][][]byte
	for _, p := range proof {
		bz, err := hex.DecodeString(p[2:])
		if err != nil {
			panic(err)
		}
		var val [][]byte
		if err := rlp.DecodeBytes(bz, &val); err != nil {
			panic(err)
		}
		target = append(target, val)
	}
	bz, err := rlp.EncodeToBytes(target)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
