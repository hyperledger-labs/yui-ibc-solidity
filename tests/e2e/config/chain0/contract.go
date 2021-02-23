package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xA05d3D10aB5aB40f5e7751411a4ff975e6Ecc97e"
	IBCModuleAddress = "0xA59f1b02B85212bBCBEB8A060aa4C595ADDFA25e"
	IBFT2ClientAddress = "0x747296FC9d600e4Ce2156dE3aeE8aa75bf2E459a"
	SimpleTokenModuleAddress = "0xF5bd99Ee3fc9eCD0fd7511218AB89161dE50cba5"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBCModuleAddress() common.Address {
	return common.HexToAddress(IBCModuleAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
