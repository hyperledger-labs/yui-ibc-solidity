package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
	IBCClientAddress     = "0xDf919000A9A2533Fb000866654ce06D565DDDb97"
	IBCConnectionAddress = "0xc0ba8346289ec43cd3f68E5EBf0a3169B1d14a2d"
	IBCChannelAddress = "0xA2f3403490466E33dcF0d74cAfc1DE0BeE0f47B4"
	IBCRoutingModuleAddress = "0x5987561e4396FC977AceFdB8DC2745305c53543a"
	SimpleTokenModuleAddress = "0x77197bBc495CEd8b342E210A4209A3F51B900BC4"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetProvableStoreAddress() common.Address {
	return common.HexToAddress(ProvableStoreAddress)
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
