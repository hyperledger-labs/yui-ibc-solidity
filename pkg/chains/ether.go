package chains

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverCommitterAddressesVals(headerHash []byte, seals [][]byte) (map[common.Address][]byte, error) {
	vals := make(map[common.Address][]byte)
	for _, seal := range seals {
		addr, err := ECRecoverAddress(headerHash, seal[:])
		if err != nil {
			return nil, err
		}
		vals[addr] = seal[:]
	}
	return vals, nil
}

func ECRecoverAddress(hash, sig []byte) (common.Address, error) {
	pub, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}
