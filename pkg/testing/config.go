package testing

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type ContractConfig struct {
	ibcHandlerAddress        common.Address
	erc1967ProxyAddress      common.Address
	ICS20TransferBankAddress common.Address
	ICS20BankAddress         common.Address
	ERC20TokenAddress        common.Address
	IBCMockAppAddress        common.Address
}

func (cc *ContractConfig) Validate() error {
	var zero common.Address
	if cc.GetIBCHandlerAddress() == zero {
		if cc.IsUpgradeable() {
			return errors.New("ERC1967ProxyAddress is empty")
		} else {
			return errors.New("IBCHandlerAddress is empty")
		}
	} else if cc.ICS20TransferBankAddress == zero {
		return errors.New("ICS20TransferBankAddress is empty")
	} else if cc.ICS20BankAddress == zero {
		return errors.New("ICS20BankAddress is empty")
	} else if cc.ERC20TokenAddress == zero {
		return errors.New("ERC20TokenAddress is empty")
	} else if cc.IBCMockAppAddress == zero {
		return errors.New("IBCMockAppAddress is empty")
	} else {
		return nil
	}
}

func (cc *ContractConfig) GetIBCHandlerAddress() common.Address {
	if cc.IsUpgradeable() {
		return cc.erc1967ProxyAddress
	}
	return cc.ibcHandlerAddress
}

func (cc *ContractConfig) IsUpgradeable() bool {
	return os.Getenv("TEST_UPGRADEABLE") == "true"
}

type BroadcastLog struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionType string         `json:"transactionType"`
	ContractName    string         `json:"contractName"`
	ContractAddress common.Address `json:"contractAddress"`
}

func buildContractConfigFromBroadcastLog(path string) (*ContractConfig, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cc ContractConfig

	var log BroadcastLog
	if err := json.Unmarshal(bz, &log); err != nil {
		return nil, err
	}
	for _, tx := range log.Transactions {
		if tx.TransactionType != "CREATE" {
			continue
		}
		switch tx.ContractName {
		case "OwnableIBCHandler":
			cc.ibcHandlerAddress = tx.ContractAddress
		case "ERC1967Proxy":
			cc.erc1967ProxyAddress = tx.ContractAddress
		case "ICS20TransferBank":
			cc.ICS20TransferBankAddress = tx.ContractAddress
		case "ICS20Bank":
			cc.ICS20BankAddress = tx.ContractAddress
		case "ERC20Token":
			cc.ERC20TokenAddress = tx.ContractAddress
		case "IBCMockApp":
			cc.IBCMockAppAddress = tx.ContractAddress
		}
	}
	if err := cc.Validate(); err != nil {
		return nil, err
	}
	return &cc, nil
}
