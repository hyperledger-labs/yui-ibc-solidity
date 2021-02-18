package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ProvableStoreAddress = "0x2d5d1eD64FC938204A898DF18C5343D84b57073d"
	IBCClientAddress     = "0x22EbbE299c54Efb345a051F852633b800ffA535a"
	IBCConnectionAddress = "0x1DeF13a624E14dd99b87653770bcf4B5A353534F"
	IBCChannelAddress = "0x729b4C12a6688A5601eE458068c532D44c04e97B"
	IBCHandlerAddress = "0xA0Bf8114B147695EfeC268b379FB74D262076C87"
	SimpleTokenModuleAddress = "0x997acb6FDab36e6946daFfAbFC35B4e2EF243745"
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
