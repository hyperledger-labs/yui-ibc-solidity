package testing

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type ContractConfig struct {
	IBCHandlerAddress              common.Address
	ICS20TransferBankAddress       common.Address
	ICS20BankAddress               common.Address
	IBCCommitmentTestHelperAddress common.Address
	SimpleTokenAddress             common.Address
}

func (cc *ContractConfig) Validate() error {
	var zero common.Address
	if cc.IBCHandlerAddress == zero {
		return errors.New("IBCHandlerAddress is empty")
	} else if cc.ICS20TransferBankAddress == zero {
		return errors.New("ICS20TransferBankAddress is empty")
	} else if cc.ICS20BankAddress == zero {
		return errors.New("ICS20BankAddress is empty")
	} else if cc.IBCCommitmentTestHelperAddress == zero {
		return errors.New("IBCCommitmentTestHelperAddress is empty")
	} else if cc.SimpleTokenAddress == zero {
		return errors.New("SimpleTokenAddress is empty")
	} else {
		return nil
	}
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
			cc.IBCHandlerAddress = tx.ContractAddress
		case "ICS20TransferBank":
			cc.ICS20TransferBankAddress = tx.ContractAddress
		case "ICS20Bank":
			cc.ICS20BankAddress = tx.ContractAddress
		case "IBCCommitmentTestHelper":
			cc.IBCCommitmentTestHelperAddress = tx.ContractAddress
		case "SimpleToken":
			cc.SimpleTokenAddress = tx.ContractAddress
		}
	}
	if err := cc.Validate(); err != nil {
		return nil, err
	}
	return &cc, nil
}
