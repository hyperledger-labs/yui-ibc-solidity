// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockclient

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

// IbcLightclientsMockV1HeaderData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsMockV1HeaderData struct {
	Height    HeightData
	Timestamp uint64
}

// MockclientMetaData contains all meta data concerning the Mockclient contract.
var MockclientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcHandler_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"consensusStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStatus\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClient.ClientStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestampAtHeight\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"protoConsensusState\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routeUpdateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"protoClientMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setStatus\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumILightClient.ClientStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIbcLightclientsMockV1Header.Data\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"heights\",\"type\":\"tuple[]\",\"internalType\":\"structHeight.Data[]\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ClientStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidClientState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeader\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrefix\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActiveClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedProtoAnyTypeURL\",\"inputs\":[{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// MockclientABI is the input ABI used to generate the binding from.
// Deprecated: Use MockclientMetaData.ABI instead.
var MockclientABI = MockclientMetaData.ABI

// Mockclient is an auto generated Go binding around an Ethereum contract.
type Mockclient struct {
	MockclientCaller     // Read-only binding to the contract
	MockclientTransactor // Write-only binding to the contract
	MockclientFilterer   // Log filterer for contract events
}

// MockclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockclientSession struct {
	Contract     *Mockclient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockclientCallerSession struct {
	Contract *MockclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockclientTransactorSession struct {
	Contract     *MockclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockclientRaw struct {
	Contract *Mockclient // Generic contract binding to access the raw methods on
}

// MockclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockclientCallerRaw struct {
	Contract *MockclientCaller // Generic read-only contract binding to access the raw methods on
}

// MockclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockclientTransactorRaw struct {
	Contract *MockclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockclient creates a new instance of Mockclient, bound to a specific deployed contract.
func NewMockclient(address common.Address, backend bind.ContractBackend) (*Mockclient, error) {
	contract, err := bindMockclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mockclient{MockclientCaller: MockclientCaller{contract: contract}, MockclientTransactor: MockclientTransactor{contract: contract}, MockclientFilterer: MockclientFilterer{contract: contract}}, nil
}

// NewMockclientCaller creates a new read-only instance of Mockclient, bound to a specific deployed contract.
func NewMockclientCaller(address common.Address, caller bind.ContractCaller) (*MockclientCaller, error) {
	contract, err := bindMockclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockclientCaller{contract: contract}, nil
}

// NewMockclientTransactor creates a new write-only instance of Mockclient, bound to a specific deployed contract.
func NewMockclientTransactor(address common.Address, transactor bind.ContractTransactor) (*MockclientTransactor, error) {
	contract, err := bindMockclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockclientTransactor{contract: contract}, nil
}

// NewMockclientFilterer creates a new log filterer instance of Mockclient, bound to a specific deployed contract.
func NewMockclientFilterer(address common.Address, filterer bind.ContractFilterer) (*MockclientFilterer, error) {
	contract, err := bindMockclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockclientFilterer{contract: contract}, nil
}

// bindMockclient binds a generic wrapper to an already deployed contract.
func bindMockclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockclient *MockclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockclient.Contract.MockclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockclient *MockclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockclient.Contract.MockclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockclient *MockclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockclient.Contract.MockclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockclient *MockclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockclient *MockclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockclient *MockclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockclient.Contract.contract.Transact(opts, method, params...)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Mockclient *MockclientCaller) GetClientState(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getClientState", clientId)

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
func (_Mockclient *MockclientSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Mockclient.Contract.GetClientState(&_Mockclient.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes clientStateBytes, bool)
func (_Mockclient *MockclientCallerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Mockclient.Contract.GetClientState(&_Mockclient.CallOpts, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Mockclient *MockclientCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height HeightData) ([]byte, bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getConsensusState", clientId, height)

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
func (_Mockclient *MockclientSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Mockclient.Contract.GetConsensusState(&_Mockclient.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Mockclient *MockclientCallerSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Mockclient.Contract.GetConsensusState(&_Mockclient.CallOpts, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Mockclient *MockclientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (HeightData, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getLatestHeight", clientId)

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Mockclient *MockclientSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Mockclient.Contract.GetLatestHeight(&_Mockclient.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64))
func (_Mockclient *MockclientCallerSession) GetLatestHeight(clientId string) (HeightData, error) {
	return _Mockclient.Contract.GetLatestHeight(&_Mockclient.CallOpts, clientId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string clientId) view returns(uint8)
func (_Mockclient *MockclientCaller) GetStatus(opts *bind.CallOpts, clientId string) (uint8, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getStatus", clientId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string clientId) view returns(uint8)
func (_Mockclient *MockclientSession) GetStatus(clientId string) (uint8, error) {
	return _Mockclient.Contract.GetStatus(&_Mockclient.CallOpts, clientId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string clientId) view returns(uint8)
func (_Mockclient *MockclientCallerSession) GetStatus(clientId string) (uint8, error) {
	return _Mockclient.Contract.GetStatus(&_Mockclient.CallOpts, clientId)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Mockclient *MockclientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height HeightData) (uint64, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Mockclient *MockclientSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Mockclient.Contract.GetTimestampAtHeight(&_Mockclient.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64)
func (_Mockclient *MockclientCallerSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, error) {
	return _Mockclient.Contract.GetTimestampAtHeight(&_Mockclient.CallOpts, clientId, height)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockclient *MockclientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockclient *MockclientSession) Owner() (common.Address, error) {
	return _Mockclient.Contract.Owner(&_Mockclient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockclient *MockclientCallerSession) Owner() (common.Address, error) {
	return _Mockclient.Contract.Owner(&_Mockclient.CallOpts)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Mockclient *MockclientCaller) RouteUpdateClient(opts *bind.CallOpts, clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "routeUpdateClient", clientId, protoClientMessage)

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
func (_Mockclient *MockclientSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Mockclient.Contract.RouteUpdateClient(&_Mockclient.CallOpts, clientId, protoClientMessage)
}

// RouteUpdateClient is a free data retrieval call binding the contract method 0xf13a62f9.
//
// Solidity: function routeUpdateClient(string clientId, bytes protoClientMessage) pure returns(bytes4, bytes)
func (_Mockclient *MockclientCallerSession) RouteUpdateClient(clientId string, protoClientMessage []byte) ([4]byte, []byte, error) {
	return _Mockclient.Contract.RouteUpdateClient(&_Mockclient.CallOpts, clientId, protoClientMessage)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Mockclient *MockclientCaller) VerifyMembership(opts *bind.CallOpts, clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "verifyMembership", clientId, height, arg2, arg3, proof, prefix, path, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Mockclient *MockclientSession) VerifyMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Mockclient.Contract.VerifyMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, prefix, path, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path, bytes value) view returns(bool)
func (_Mockclient *MockclientCallerSession) VerifyMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte, value []byte) (bool, error) {
	return _Mockclient.Contract.VerifyMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, prefix, path, value)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Mockclient *MockclientCaller) VerifyNonMembership(opts *bind.CallOpts, clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "verifyNonMembership", clientId, height, arg2, arg3, proof, prefix, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Mockclient *MockclientSession) VerifyNonMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Mockclient.Contract.VerifyNonMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, prefix, path)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x999fbbb3.
//
// Solidity: function verifyNonMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes prefix, bytes path) view returns(bool)
func (_Mockclient *MockclientCallerSession) VerifyNonMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, prefix []byte, path []byte) (bool, error) {
	return _Mockclient.Contract.VerifyNonMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, prefix, path)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Mockclient *MockclientTransactor) InitializeClient(opts *bind.TransactOpts, clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "initializeClient", clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Mockclient *MockclientSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.InitializeClient(&_Mockclient.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// InitializeClient is a paid mutator transaction binding the contract method 0xfe66819f.
//
// Solidity: function initializeClient(string clientId, bytes protoClientState, bytes protoConsensusState) returns((uint64,uint64) height)
func (_Mockclient *MockclientTransactorSession) InitializeClient(clientId string, protoClientState []byte, protoConsensusState []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.InitializeClient(&_Mockclient.TransactOpts, clientId, protoClientState, protoConsensusState)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockclient *MockclientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockclient *MockclientSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mockclient.Contract.RenounceOwnership(&_Mockclient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockclient *MockclientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mockclient.Contract.RenounceOwnership(&_Mockclient.TransactOpts)
}

// SetStatus is a paid mutator transaction binding the contract method 0xdeaab7a9.
//
// Solidity: function setStatus(string clientId, uint8 status) returns()
func (_Mockclient *MockclientTransactor) SetStatus(opts *bind.TransactOpts, clientId string, status uint8) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "setStatus", clientId, status)
}

// SetStatus is a paid mutator transaction binding the contract method 0xdeaab7a9.
//
// Solidity: function setStatus(string clientId, uint8 status) returns()
func (_Mockclient *MockclientSession) SetStatus(clientId string, status uint8) (*types.Transaction, error) {
	return _Mockclient.Contract.SetStatus(&_Mockclient.TransactOpts, clientId, status)
}

// SetStatus is a paid mutator transaction binding the contract method 0xdeaab7a9.
//
// Solidity: function setStatus(string clientId, uint8 status) returns()
func (_Mockclient *MockclientTransactorSession) SetStatus(clientId string, status uint8) (*types.Transaction, error) {
	return _Mockclient.Contract.SetStatus(&_Mockclient.TransactOpts, clientId, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockclient *MockclientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockclient *MockclientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mockclient.Contract.TransferOwnership(&_Mockclient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockclient *MockclientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mockclient.Contract.TransferOwnership(&_Mockclient.TransactOpts, newOwner)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x883a9998.
//
// Solidity: function updateClient(string clientId, ((uint64,uint64),uint64) header) returns((uint64,uint64)[] heights)
func (_Mockclient *MockclientTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, header IbcLightclientsMockV1HeaderData) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "updateClient", clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x883a9998.
//
// Solidity: function updateClient(string clientId, ((uint64,uint64),uint64) header) returns((uint64,uint64)[] heights)
func (_Mockclient *MockclientSession) UpdateClient(clientId string, header IbcLightclientsMockV1HeaderData) (*types.Transaction, error) {
	return _Mockclient.Contract.UpdateClient(&_Mockclient.TransactOpts, clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x883a9998.
//
// Solidity: function updateClient(string clientId, ((uint64,uint64),uint64) header) returns((uint64,uint64)[] heights)
func (_Mockclient *MockclientTransactorSession) UpdateClient(clientId string, header IbcLightclientsMockV1HeaderData) (*types.Transaction, error) {
	return _Mockclient.Contract.UpdateClient(&_Mockclient.TransactOpts, clientId, header)
}

// MockclientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mockclient contract.
type MockclientOwnershipTransferredIterator struct {
	Event *MockclientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockclientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockclientOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockclientOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockclientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockclientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockclientOwnershipTransferred represents a OwnershipTransferred event raised by the Mockclient contract.
type MockclientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mockclient *MockclientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MockclientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mockclient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MockclientOwnershipTransferredIterator{contract: _Mockclient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mockclient *MockclientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockclientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mockclient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockclientOwnershipTransferred)
				if err := _Mockclient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mockclient *MockclientFilterer) ParseOwnershipTransferred(log types.Log) (*MockclientOwnershipTransferred, error) {
	event := new(MockclientOwnershipTransferred)
	if err := _Mockclient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
