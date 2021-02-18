package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x56CF24731e6DaAC0499C52bdd61Fae800bB7A173"
	IBCClientAddress     = "0xfed5D6a133522d9616C3F5e4916B27E27F026013"
	IBCConnectionAddress = "0x2d5d1eD64FC938204A898DF18C5343D84b57073d"
	IBCChannelAddress = "0x22EbbE299c54Efb345a051F852633b800ffA535a"
	IBCHandlerAddress = "0x1DeF13a624E14dd99b87653770bcf4B5A353534F"
	SimpleTokenModuleAddress = "0x729b4C12a6688A5601eE458068c532D44c04e97B"
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

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetSimpleTokenModuleAddress() common.Address {
	return common.HexToAddress(SimpleTokenModuleAddress)
}
