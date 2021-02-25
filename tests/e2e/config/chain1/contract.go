package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xA629b31f3F0C8B48E395014b6283b7B616dfdb8D"
	IBCClientAddress = "0x56CF24731e6DaAC0499C52bdd61Fae800bB7A173"
	IBCConnectionAddress = "0xfed5D6a133522d9616C3F5e4916B27E27F026013"
	IBCChannelAddress = "0x2d5d1eD64FC938204A898DF18C5343D84b57073d"
	IBCRoutingModuleAddress = "0x22EbbE299c54Efb345a051F852633b800ffA535a"
	IBCIdentifierAddress = "0x3d906FFc0D2389A5162058897cDd54C96F07c496"
	IBFT2ClientAddress = "0x029EAfA3eED79299F635E2A6F1607875FF7fA983"
	SimpleTokenModuleAddress = "0x1DeF13a624E14dd99b87653770bcf4B5A353534F"
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
