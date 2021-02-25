package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xC5de6fc9374BA11cF47Cc1aECF94aBE620821BAa"
	IBCClientAddress = "0x3Ac60839D86731884d28C709A1d2f63a3728a882"
	IBCConnectionAddress = "0xc2351FdD53F418F788BC4801401b4C4484b49711"
	IBCChannelAddress = "0xA22BA4045bc50Fa615Ce9E373BeDe56782197BeE"
	IBCRoutingModuleAddress = "0x9E75ED48E7AB483Cf810EF84D7DAb844c11B1644"
	IBFT2ClientAddress = "0x31DEd0C2E93BB82760DCb11ddb6AfCF5045Dd371"
	SimpleTokenModuleAddress = "0x23ef151a3C438033ddAB620452bC035EA966ceaF"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetIBCClientAddress() common.Address {
	return common.HexToAddress(IBCClientAddress)
}

func (contractConfig) GetIBCConnectionAddress() common.Address {
	return common.HexToAddress(IBCConnectionAddress)
}

func (contractConfig) GetIBCChannelAddress() common.Address {
	return common.HexToAddress(IBCChannelAddress)
}

func (contractConfig) GetIBCRoutingModuleAddress() common.Address {
	return common.HexToAddress(IBCRoutingModuleAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
