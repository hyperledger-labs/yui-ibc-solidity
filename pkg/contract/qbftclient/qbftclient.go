// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qbftclient

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcLightclientsQbftV1HeaderData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsQbftV1HeaderData struct {
	BesuHeaderRlp     []byte
	Seals             [][]byte
	TrustedHeight     HeightData
	AccountStateProof []byte
}

// QbftclientMetaData contains all meta data concerning the Qbftclient contract.
var QbftclientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcHandler_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"consensusStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestInfo\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"latestTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumILightClient.ClientStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClient.ClientStatus\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTimestampAtHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ibcHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"protoConsensusState\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routeUpdateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIbcLightclientsQbftV1Header.Data\",\"components\":[{\"name\":\"besu_header_rlp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"seals\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"trusted_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"account_state_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"heights\",\"type\":\"tuple[]\",\"internalType\":\"structHeight.Data[]\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"delayTimePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delayBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"delayTimePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delayBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"EmptyValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientTrustedValidatorsSeals\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsuffientUntrustedValidatorsSeals\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClientStateLatestHeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusStateRootLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidECDSASignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidIBCAddressLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddressLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LightClientClientStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"LightClientConsensusStateExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LightClientConsensusStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"type\":\"error\",\"name\":\"LightClientInvalidCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LightClientNotActiveClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"LightClientUnexpectedProtoAnyTypeURL\",\"inputs\":[{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UnexpectedEthereumHeaderFormat\",\"inputs\":[{\"name\":\"itemsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnexpectedExtraDataFormat\",\"inputs\":[{\"name\":\"itemsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// QbftclientABI is the input ABI used to generate the binding from.
// Deprecated: Use QbftclientMetaData.ABI instead.
var QbftclientABI = QbftclientMetaData.ABI

// Qbftclient is an auto generated Go binding around an Ethereum contract.
type Qbftclient struct {
	QbftclientCaller     // Read-only binding to the contract
	QbftclientTransactor // Write-only binding to the contract
	QbftclientFilterer   // Log filterer for contract events
}

// QbftclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type QbftclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QbftclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QbftclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QbftclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QbftclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QbftclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QbftclientSession struct {
	Contract     *Qbftclient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QbftclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QbftclientCallerSession struct {
	Contract *QbftclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// QbftclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QbftclientTransactorSession struct {
	Contract     *QbftclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// QbftclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type QbftclientRaw struct {
	Contract *Qbftclient // Generic contract binding to access the raw methods on
}

// QbftclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QbftclientCallerRaw struct {
	Contract *QbftclientCaller // Generic read-only contract binding to access the raw methods on
}

// QbftclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QbftclientTransactorRaw struct {
	Contract *QbftclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQbftclient creates a new instance of Qbftclient, bound to a specific deployed contract.
func NewQbftclient(address common.Address, backend bind.ContractBackend) (*Qbftclient, error) {
	contract, err := bindQbftclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Qbftclient{QbftclientCaller: QbftclientCaller{contract: contract}, QbftclientTransactor: QbftclientTransactor{contract: contract}, QbftclientFilterer: QbftclientFilterer{contract: contract}}, nil
}

// NewQbftclientCaller creates a new read-only instance of Qbftclient, bound to a specific deployed contract.
func NewQbftclientCaller(address common.Address, caller bind.ContractCaller) (*QbftclientCaller, error) {
	contract, err := bindQbftclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QbftclientCaller{contract: contract}, nil
}

// NewQbftclientTransactor creates a new write-only instance of Qbftclient, bound to a specific deployed contract.
func NewQbftclientTransactor(address common.Address, transactor bind.ContractTransactor) (*QbftclientTransactor, error) {
	contract, err := bindQbftclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QbftclientTransactor{contract: contract}, nil
}

// NewQbftclientFilterer creates a new log filterer instance of Qbftclient, bound to a specific deployed contract.
func NewQbftclientFilterer(address common.Address, filterer bind.ContractFilterer) (*QbftclientFilterer, error) {
	contract, err := bindQbftclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QbftclientFilterer{contract: contract}, nil
}

// bindQbftclient binds a generic wrapper to an already deployed contract.
func bindQbftclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QbftclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qbftclient *QbftclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qbftclient.Contract.QbftclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qbftclient *QbftclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qbftclient.Contract.QbftclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qbftclient *QbftclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qbftclient.Contract.QbftclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qbftclient *QbftclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qbftclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qbftclient *QbftclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qbftclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qbftclient *QbftclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qbftclient.Contract.contract.Transact(opts, method, params...)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Qbftclient *QbftclientCaller) GetClientState(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Qbftclient *QbftclientSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Qbftclient.Contract.GetClientState(&_Qbftclient.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Qbftclient *QbftclientCallerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Qbftclient.Contract.GetClientState(&_Qbftclient.CallOpts, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Qbftclient *QbftclientCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height HeightData) ([]byte, bool, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Qbftclient *QbftclientSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Qbftclient.Contract.GetConsensusState(&_Qbftclient.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Qbftclient *QbftclientCallerSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Qbftclient.Contract.GetConsensusState(&_Qbftclient.CallOpts, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Qbftclient *QbftclientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (HeightData, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getLatestHeight", clientId)

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Qbftclient *QbftclientSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Qbftclient.Contract.GetLatestHeight(&_Qbftclient.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Qbftclient *QbftclientCallerSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Qbftclient.Contract.GetLatestHeight(&_Qbftclient.CallOpts, clientId)
}

// GetLatestInfo is a free data retrieval call binding the contract method 0xa5906897.
//
// Solidity: function getLatestInfo(string clientId) view returns((uint64,uint64) latestHeight, uint64 latestTimestamp, uint8 status)
func (_Qbftclient *QbftclientCaller) GetLatestInfo(opts *bind.CallOpts, clientId string) (struct {
	LatestHeight    HeightData
	LatestTimestamp uint64
	Status          uint8
}, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getLatestInfo", clientId)

	outstruct := new(struct {
		LatestHeight    HeightData
		LatestTimestamp uint64
		Status          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LatestHeight = *abi.ConvertType(out[0], new(HeightData)).(*HeightData)
	outstruct.LatestTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetLatestInfo is a free data retrieval call binding the contract method 0xa5906897.
//
// Solidity: function getLatestInfo(string clientId) view returns((uint64,uint64) latestHeight, uint64 latestTimestamp, uint8 status)
func (_Qbftclient *QbftclientSession) GetLatestInfo(clientId string) (struct {
	LatestHeight    HeightData
	LatestTimestamp uint64
	Status          uint8
}, error) {
	return _Qbftclient.Contract.GetLatestInfo(&_Qbftclient.CallOpts, clientId)
}

// GetLatestInfo is a free data retrieval call binding the contract method 0xa5906897.
//
// Solidity: function getLatestInfo(string clientId) view returns((uint64,uint64) latestHeight, uint64 latestTimestamp, uint8 status)
func (_Qbftclient *QbftclientCallerSession) GetLatestInfo(clientId string) (struct {
	LatestHeight    HeightData
	LatestTimestamp uint64
	Status          uint8
}, error) {
	return _Qbftclient.Contract.GetLatestInfo(&_Qbftclient.CallOpts, clientId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Qbftclient *QbftclientCaller) GetStatus(opts *bind.CallOpts, arg0 string) (uint8, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Qbftclient *QbftclientSession) GetStatus(arg0 string) (uint8, error) {
	return _Qbftclient.Contract.GetStatus(&_Qbftclient.CallOpts, arg0)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Qbftclient *QbftclientCallerSession) GetStatus(arg0 string) (uint8, error) {
	return _Qbftclient.Contract.GetStatus(&_Qbftclient.CallOpts, arg0)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Qbftclient *QbftclientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height HeightData) (uint64, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Qbftclient *QbftclientSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Qbftclient.Contract.GetTimestampAtHeight(&_Qbftclient.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Qbftclient *QbftclientCallerSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Qbftclient.Contract.GetTimestampAtHeight(&_Qbftclient.CallOpts, clientId, height)
}

// IbcHandler is a free data retrieval call binding the contract method 0x2dc1bd40.
//
// Solidity: function ibcHandler() view returns(address)
func (_Qbftclient *QbftclientCaller) IbcHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "ibcHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcHandler is a free data retrieval call binding the contract method 0x2dc1bd40.
//
// Solidity: function ibcHandler() view returns(address)
func (_Qbftclient *QbftclientSession) IbcHandler() (common.Address, error) {
	return _Qbftclient.Contract.IbcHandler(&_Qbftclient.CallOpts)
}

// IbcHandler is a free data retrieval call binding the contract method 0x2dc1bd40.
//
// Solidity: function ibcHandler() view returns(address)
func (_Qbftclient *QbftclientCallerSession) IbcHandler() (common.Address, error) {
	return _Qbftclient.Contract.IbcHandler(&_Qbftclient.CallOpts)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Qbftclient *QbftclientCaller) RouteUpdateClient(opts *bind.CallOpts, clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "routeUpdateClient", clientId, protoClientMessage)

	if err != nil {
		return *new([4]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Qbftclient *QbftclientSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Qbftclient.Contract.RouteUpdateClient(&_Qbftclient.CallOpts, clientId, protoClientMessage)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Qbftclient *QbftclientCallerSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Qbftclient.Contract.RouteUpdateClient(&_Qbftclient.CallOpts, clientId, protoClientMessage)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Qbftclient *QbftclientCaller) VerifyMembership(opts *bind.CallOpts, clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "verifyMembership", clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Qbftclient *QbftclientSession) VerifyMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Qbftclient.Contract.VerifyMembership(&_Qbftclient.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Qbftclient *QbftclientCallerSession) VerifyMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Qbftclient.Contract.VerifyMembership(&_Qbftclient.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Qbftclient *QbftclientCaller) VerifyNonMembership(opts *bind.CallOpts, clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	var out []interface{}
	err := _Qbftclient.contract.Call(opts, &out, "verifyNonMembership", clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Qbftclient *QbftclientSession) VerifyNonMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Qbftclient.Contract.VerifyNonMembership(&_Qbftclient.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Qbftclient *QbftclientCallerSession) VerifyNonMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Qbftclient.Contract.VerifyNonMembership(&_Qbftclient.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Qbftclient *QbftclientTransactor) InitializeClient(opts *bind.TransactOpts, clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Qbftclient.contract.Transact(opts, "initializeClient", clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Qbftclient *QbftclientSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Qbftclient.Contract.InitializeClient(&_Qbftclient.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Qbftclient *QbftclientTransactorSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Qbftclient.Contract.InitializeClient(&_Qbftclient.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Qbftclient *QbftclientTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, header IbcLightclientsQbftV1HeaderData) (*types.Transaction, error) {
	return _Qbftclient.contract.Transact(opts, "updateClient", clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Qbftclient *QbftclientSession) UpdateClient(clientId string, header IbcLightclientsQbftV1HeaderData) (*types.Transaction, error) {
	return _Qbftclient.Contract.UpdateClient(&_Qbftclient.TransactOpts, clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Qbftclient *QbftclientTransactorSession) UpdateClient(clientId string, header IbcLightclientsQbftV1HeaderData) (*types.Transaction, error) {
	return _Qbftclient.Contract.UpdateClient(&_Qbftclient.TransactOpts, clientId, header)
}
