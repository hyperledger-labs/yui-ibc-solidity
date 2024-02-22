// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibft2client

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

// IbcLightclientsIbft2V1HeaderData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsIbft2V1HeaderData struct {
	BesuHeaderRlp     []byte
	Seals             [][]byte
	TrustedHeight     HeightData
	AccountStateProof []byte
}

// Ibft2clientMetaData contains all meta data concerning the Ibft2client contract.
var Ibft2clientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcHandler_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"consensusStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClient.ClientStatus\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTimestampAtHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"protoConsensusState\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routeUpdateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIbcLightclientsIbft2V1Header.Data\",\"components\":[{\"name\":\"besu_header_rlp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"seals\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"trusted_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"account_state_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"heights\",\"type\":\"tuple[]\",\"internalType\":\"structHeight.Data[]\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"delayTimePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delayBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"delayTimePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delayBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ClientStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"type\":\"error\",\"name\":\"EmptyValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientTrustedValidatorsSeals\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsuffientUntrustedValidatorsSeals\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConsensusStateRootLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidECDSASignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidIBCAddressLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddressLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActiveClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UnexpectedEthereumHeaderFormat\",\"inputs\":[{\"name\":\"itemsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnexpectedExtraDataFormat\",\"inputs\":[{\"name\":\"itemsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnexpectedProtoAnyTypeURL\",\"inputs\":[{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// Ibft2clientABI is the input ABI used to generate the binding from.
// Deprecated: Use Ibft2clientMetaData.ABI instead.
var Ibft2clientABI = Ibft2clientMetaData.ABI

// Ibft2client is an auto generated Go binding around an Ethereum contract.
type Ibft2client struct {
	Ibft2clientCaller     // Read-only binding to the contract
	Ibft2clientTransactor // Write-only binding to the contract
	Ibft2clientFilterer   // Log filterer for contract events
}

// Ibft2clientCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ibft2clientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibft2clientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ibft2clientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibft2clientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ibft2clientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ibft2clientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ibft2clientSession struct {
	Contract     *Ibft2client      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ibft2clientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ibft2clientCallerSession struct {
	Contract *Ibft2clientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Ibft2clientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ibft2clientTransactorSession struct {
	Contract     *Ibft2clientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Ibft2clientRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ibft2clientRaw struct {
	Contract *Ibft2client // Generic contract binding to access the raw methods on
}

// Ibft2clientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ibft2clientCallerRaw struct {
	Contract *Ibft2clientCaller // Generic read-only contract binding to access the raw methods on
}

// Ibft2clientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ibft2clientTransactorRaw struct {
	Contract *Ibft2clientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbft2client creates a new instance of Ibft2client, bound to a specific deployed contract.
func NewIbft2client(address common.Address, backend bind.ContractBackend) (*Ibft2client, error) {
	contract, err := bindIbft2client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibft2client{Ibft2clientCaller: Ibft2clientCaller{contract: contract}, Ibft2clientTransactor: Ibft2clientTransactor{contract: contract}, Ibft2clientFilterer: Ibft2clientFilterer{contract: contract}}, nil
}

// NewIbft2clientCaller creates a new read-only instance of Ibft2client, bound to a specific deployed contract.
func NewIbft2clientCaller(address common.Address, caller bind.ContractCaller) (*Ibft2clientCaller, error) {
	contract, err := bindIbft2client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ibft2clientCaller{contract: contract}, nil
}

// NewIbft2clientTransactor creates a new write-only instance of Ibft2client, bound to a specific deployed contract.
func NewIbft2clientTransactor(address common.Address, transactor bind.ContractTransactor) (*Ibft2clientTransactor, error) {
	contract, err := bindIbft2client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ibft2clientTransactor{contract: contract}, nil
}

// NewIbft2clientFilterer creates a new log filterer instance of Ibft2client, bound to a specific deployed contract.
func NewIbft2clientFilterer(address common.Address, filterer bind.ContractFilterer) (*Ibft2clientFilterer, error) {
	contract, err := bindIbft2client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ibft2clientFilterer{contract: contract}, nil
}

// bindIbft2client binds a generic wrapper to an already deployed contract.
func bindIbft2client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ibft2clientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibft2client *Ibft2clientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibft2client.Contract.Ibft2clientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibft2client *Ibft2clientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibft2client.Contract.Ibft2clientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibft2client *Ibft2clientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibft2client.Contract.Ibft2clientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibft2client *Ibft2clientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibft2client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibft2client *Ibft2clientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibft2client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibft2client *Ibft2clientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibft2client.Contract.contract.Transact(opts, method, params...)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Ibft2client *Ibft2clientCaller) GetClientState(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getClientState", clientId)

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
func (_Ibft2client *Ibft2clientSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Ibft2client *Ibft2clientCallerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Ibft2client *Ibft2clientCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height HeightData) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getConsensusState", clientId, height)

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
func (_Ibft2client *Ibft2clientSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Ibft2client *Ibft2clientCallerSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Ibft2client *Ibft2clientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (HeightData, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getLatestHeight", clientId)

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Ibft2client *Ibft2clientSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Ibft2client *Ibft2clientCallerSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Ibft2client *Ibft2clientCaller) GetStatus(opts *bind.CallOpts, arg0 string) (uint8, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Ibft2client *Ibft2clientSession) GetStatus(arg0 string) (uint8, error) {
	return _Ibft2client.Contract.GetStatus(&_Ibft2client.CallOpts, arg0)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string ) pure returns(uint8)
func (_Ibft2client *Ibft2clientCallerSession) GetStatus(arg0 string) (uint8, error) {
	return _Ibft2client.Contract.GetStatus(&_Ibft2client.CallOpts, arg0)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Ibft2client *Ibft2clientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height HeightData) (uint64, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Ibft2client *Ibft2clientSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Ibft2client *Ibft2clientCallerSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Ibft2client *Ibft2clientCaller) RouteUpdateClient(opts *bind.CallOpts, clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "routeUpdateClient", clientId, protoClientMessage)

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
func (_Ibft2client *Ibft2clientSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Ibft2client.Contract.RouteUpdateClient(&_Ibft2client.CallOpts, clientId, protoClientMessage)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Ibft2client *Ibft2clientCallerSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Ibft2client.Contract.RouteUpdateClient(&_Ibft2client.CallOpts, clientId, protoClientMessage)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyMembership(opts *bind.CallOpts, clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyMembership", clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyMembership(&_Ibft2client.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyMembership(&_Ibft2client.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path, value)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyNonMembership(opts *bind.CallOpts, clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyNonMembership", clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyNonMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyNonMembership(&_Ibft2client.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 delayTimePeriod, uint64 delayBlockPeriod, bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyNonMembership(clientId string, height HeightData, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyNonMembership(&_Ibft2client.CallOpts, clientId, height, delayTimePeriod, delayBlockPeriod, proof, prefix, path)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Ibft2client *Ibft2clientTransactor) InitializeClient(opts *bind.TransactOpts, clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Ibft2client.contract.Transact(opts, "initializeClient", clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Ibft2client *Ibft2clientSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.InitializeClient(&_Ibft2client.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Ibft2client *Ibft2clientTransactorSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.InitializeClient(&_Ibft2client.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Ibft2client *Ibft2clientTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, header IbcLightclientsIbft2V1HeaderData) (*types.Transaction, error) {
	return _Ibft2client.contract.Transact(opts, "updateClient", clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Ibft2client *Ibft2clientSession) UpdateClient(clientId string, header IbcLightclientsIbft2V1HeaderData) (*types.Transaction, error) {
	return _Ibft2client.Contract.UpdateClient(&_Ibft2client.TransactOpts, clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xa4f1ec28.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],(uint64,uint64),bytes) header) returns((uint64,uint64)[] heights)
func (_Ibft2client *Ibft2clientTransactorSession) UpdateClient(clientId string, header IbcLightclientsIbft2V1HeaderData) (*types.Transaction, error) {
	return _Ibft2client.Contract.UpdateClient(&_Ibft2client.TransactOpts, clientId, header)
}
