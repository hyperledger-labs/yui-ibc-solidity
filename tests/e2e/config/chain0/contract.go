package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xc45b69894Bd1dD735779927dd841563ECB612C46"
	IBCClientAddress = "0x550844A8A809a7247f45952DE6E93A9aa5b4081C"
	IBCConnectionAddress = "0x1f75edf929d2037d2018cD33f669764b9B706c62"
	IBCChannelAddress = "0xC2644568Cae8A12720E657cc611fa0AcFA958a95"
	IBCRoutingModuleAddress = "0xB40d873AEB1fb78c92D9Bc57CCCdC323543C9203"
	IBCIdentifierAddress = "0x1FeE5DA76FDdb2D1e99cf47ce35225023542eCd2"
	IBFT2ClientAddress = "0x41eF8F233812Ed317F833879Ebf86B0cdD302DDD"
	SimpleTokenModuleAddress = "0x2FB9F1D2310D520e9D116b4752Db3c582ecB972c"
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
