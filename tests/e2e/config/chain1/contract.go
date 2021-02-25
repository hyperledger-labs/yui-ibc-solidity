package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCStoreAddress = "0x603C100aA0C46795D76d398766838E833Ee39182"
	IBCClientAddress = "0x4bd62E979ED735aaDec4d481b3331200BC18f1C5"
	IBCConnectionAddress = "0x4Fccc6578Bc322Fcd114E020FE248c58D2bC5Ec9"
	IBCChannelAddress = "0x228C8C14A1E1476C4D90432E345906E5C3FceEE3"
	IBCRoutingModuleAddress = "0x1374280eaddD68e3378C9852c644fabB91B7Ebc2"
	IBCIdentifierAddress = "0x3b42Ab5df4a6efD0b10bCb47c6863A56B539C264"
	IBFT2ClientAddress = "0x1019A0622E397147402fFDBa5a65a4046286a233"
	SimpleTokenModuleAddress = "0xe9C26a6280135B69f3383708458d7277EE2fCfeE"
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
