package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x1f75edf929d2037d2018cD33f669764b9B706c62"
	IBCClientAddress = "0xB40d873AEB1fb78c92D9Bc57CCCdC323543C9203"
	IBCConnectionAddress = "0x2FB9F1D2310D520e9D116b4752Db3c582ecB972c"
	IBCChannelAddress = "0x07392E81183f86CD2B66f427c1E8Cc21eca2fB31"
	IBCRoutingModuleAddress = "0xbD55811C336F26d6F19F8E0cD140AB8b37bCbe83"
	IBCIdentifierAddress = "0x41eF8F233812Ed317F833879Ebf86B0cdD302DDD"
	IBFT2ClientAddress = "0xC2644568Cae8A12720E657cc611fa0AcFA958a95"
	SimpleTokenModuleAddress = "0xA95251102a0D5C21642319eC70F5b84d49A69d66"
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
