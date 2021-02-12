package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/tyler-smith/go-bip39"
)

type HDPathLevel struct {
	Purpose  uint32
	CoinType uint32
	Account  uint32
	Change   uint32
	Index    uint32
}

func ParseHDPathLevel(path string) (*HDPathLevel, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 6 {
		return nil, errors.New("invalid path level")
	}
	if parts[0] != "m" {
		return nil, fmt.Errorf("prefix should be 'm'")
	}
	parts = parts[1:]

	var err error
	hp := new(HDPathLevel)
	for _, idx := range []int{0, 1, 2} {
		if !checkValidApostrophe(parts[idx]) {
			return nil, fmt.Errorf("missing apostrophe: %v", parts[idx])
		}
	}
	hp.Purpose, err = strToUint32(parts[0][:len(parts[0])-1])
	if err != nil {
		return nil, err
	}
	hp.CoinType, err = strToUint32(parts[1][:len(parts[1])-1])
	if err != nil {
		return nil, err
	}
	hp.Account, err = strToUint32(parts[2][:len(parts[2])-1])
	if err != nil {
		return nil, err
	}
	hp.Change, err = strToUint32(parts[3])
	if err != nil {
		return nil, err
	}
	hp.Index, err = strToUint32(parts[4])
	if err != nil {
		return nil, err
	}
	if err := hp.Validate(); err != nil {
		return nil, err
	}
	return hp, nil
}

func checkValidApostrophe(s string) bool {
	return len(s) > 0 && s[len(s)-1] == '\''
}

func strToUint32(s string) (uint32, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func (hp *HDPathLevel) Validate() error {
	if hp.Purpose != 44 {
		return errors.New("purpose should be 44")
	}
	if hp.Change != 0 && hp.Change != 1 {
		return errors.New("change should be 0 or 1")
	}
	return nil
}

func (hp *HDPathLevel) String() string {
	return fmt.Sprintf("m/%v'/%v'/%v'/%v/%v", hp.Purpose, hp.CoinType, hp.Account, hp.Change, hp.Index)
}

func GetPrvKeyFromHDWallet(seed []byte, hp *HDPathLevel) (*ecdsa.PrivateKey, error) {
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	// This gives the path: m/{purpose}'
	acc, err := masterKey.Child(hdkeychain.HardenedKeyStart + hp.Purpose)
	if err != nil {
		return nil, err
	}
	// This gives the path: m/{purpose}'/{coin_type}'
	acc, err = acc.Child(hdkeychain.HardenedKeyStart + hp.CoinType)
	if err != nil {
		return nil, err
	}
	// This gives the path: m/{purpose}'/{coin_type}'/{account}'
	acc, err = acc.Child(hdkeychain.HardenedKeyStart + hp.Account)
	if err != nil {
		return nil, err
	}
	// This gives the path: m/{purpose}'/{coin_type}'/{account}'/{change}
	acc, err = acc.Child(0 + hp.Change)
	if err != nil {
		return nil, err
	}
	// This gives the path: m/{purpose}'/{coin_type}'/{account}'/{change}/{index}
	acc, err = acc.Child(hp.Index)
	if err != nil {
		return nil, err
	}
	btcecPrivKey, err := acc.ECPrivKey()
	if err != nil {
		return nil, err
	}
	return btcecPrivKey.ToECDSA(), nil
}

func GetPrvKeyFromMnemonicAndHDWPath(mnemonic, path string) (*ecdsa.PrivateKey, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}
	hp, err := ParseHDPathLevel(path)
	if err != nil {
		return nil, err
	}
	seed := bip39.NewSeed(mnemonic, "")
	return GetPrvKeyFromHDWallet(seed, hp)
}
