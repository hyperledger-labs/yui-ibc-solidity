package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xe1B4773F25Fb5109E04CE623c71d4693777cf908"
	IBCClientAddress = "0x18680e36c779B7d08d47f38769668a264086300D"
	IBCConnectionAddress = "0x67335fC1B11f450953c317e77423c8162903f0F2"
	IBCChannelAddress = "0x9B72A7f6824758552812D267C99ef7492778Df0f"
	IBCRoutingModuleAddress = "0x57f8dD374BBe73724a123b6834f8D6e1aD5efC78"
	IBCIdentifierAddress = "0x5200801e44F7099158226E4740A84b62b1204EDD"
	IBFT2ClientAddress = "0x1Be5c5629Fe3037aF0a1a4355687DE694a6d55Bc"
	SimpleTokenModuleAddress = "0x38b1a3755e2b674cA7141c40DA97bD3A25831560"
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

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
