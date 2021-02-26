package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xA4A16feb39a5D11a1066601B6EDF369a84BBa2B2"
	IBCHandlerAddress = "0xEFC749e34Bf28D450F3B76BE9E91705dcB3C6484"
	IBCIdentifierAddress = "0x5Ea1AA48f0C4f470F5aAa12D233aA6e7b5D348Fd"
	IBFT2ClientAddress = "0x09199A80877b8481e92293dE5635dbC0DcB73127"
	SimpleTokenModuleAddress = "0x36193E778f7A586d335cDF0E5077ACf7B923a856"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
