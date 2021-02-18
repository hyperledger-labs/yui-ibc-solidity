package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0xB9F2435339644FDF60E35bF9d02F6C190a7a1930"
	IBCClientAddress     = "0xC5777A1ac9A446aBDae4d4D71a330ebF77705a46"
	IBCConnectionAddress = "0x912911A4D64C3917E78C19531E4b0df9Db534938"
	IBCChannelAddress = "0xBF346b5BC386c7C3378688286406B08E9327d312"
	IBCHandlerAddress = "0x4DB8e6C8BdE4c9AFCEDb590C5446c965c073BED8"
	SimpleTokenModuleAddress = "0xeB50cA91c99ceD8EDa25D362AdE560df9661Bc31"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCStoreAddress() common.Address {
	return common.HexToAddress(IBCStoreAddress)
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
