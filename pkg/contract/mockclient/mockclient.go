// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockclient

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ConsensusStateUpdates is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateUpdates struct {
	ConsensusStateCommitment [32]byte
	Height                   HeightData
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcLightclientsMockV1ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsMockV1ClientStateData struct {
	LatestHeight HeightData
}

// MockclientABI is the input ABI used to generate the binding from.
const MockclientABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcModule_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"clientStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"consensusStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"internalType\":\"structConsensusStateUpdates[]\",\"name\":\"updates\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getTimestampAtHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientMessageBytes\",\"type\":\"bytes\"}],\"name\":\"updateClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"clientStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"consensusStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"internalType\":\"structConsensusStateUpdates[]\",\"name\":\"updates\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"latest_height\",\"type\":\"tuple\"}],\"internalType\":\"structIbcLightclientsMockV1ClientState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

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
	parsed, err := abi.JSON(strings.NewReader(MockclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
// Solidity: function getClientState(string clientId) view returns(((uint64,uint64)), bool)
func (_Mockclient *MockclientCaller) GetClientState(opts *bind.CallOpts, clientId string) (IbcLightclientsMockV1ClientStateData, bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new(IbcLightclientsMockV1ClientStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(IbcLightclientsMockV1ClientStateData)).(*IbcLightclientsMockV1ClientStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(((uint64,uint64)), bool)
func (_Mockclient *MockclientSession) GetClientState(clientId string) (IbcLightclientsMockV1ClientStateData, bool, error) {
	return _Mockclient.Contract.GetClientState(&_Mockclient.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(((uint64,uint64)), bool)
func (_Mockclient *MockclientCallerSession) GetClientState(clientId string) (IbcLightclientsMockV1ClientStateData, bool, error) {
	return _Mockclient.Contract.GetClientState(&_Mockclient.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64), bool)
func (_Mockclient *MockclientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (HeightData, bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getLatestHeight", clientId)

	if err != nil {
		return *new(HeightData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64), bool)
func (_Mockclient *MockclientSession) GetLatestHeight(clientId string) (HeightData, bool, error) {
	return _Mockclient.Contract.GetLatestHeight(&_Mockclient.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64), bool)
func (_Mockclient *MockclientCallerSession) GetLatestHeight(clientId string) (HeightData, bool, error) {
	return _Mockclient.Contract.GetLatestHeight(&_Mockclient.CallOpts, clientId)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Mockclient *MockclientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height HeightData) (uint64, bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Mockclient *MockclientSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, bool, error) {
	return _Mockclient.Contract.GetTimestampAtHeight(&_Mockclient.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Mockclient *MockclientCallerSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, bool, error) {
	return _Mockclient.Contract.GetTimestampAtHeight(&_Mockclient.CallOpts, clientId, height)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes , bytes , bytes value) view returns(bool)
func (_Mockclient *MockclientCaller) VerifyMembership(opts *bind.CallOpts, clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, arg5 []byte, arg6 []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _Mockclient.contract.Call(opts, &out, "verifyMembership", clientId, height, arg2, arg3, proof, arg5, arg6, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes , bytes , bytes value) view returns(bool)
func (_Mockclient *MockclientSession) VerifyMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, arg5 []byte, arg6 []byte, value []byte) (bool, error) {
	return _Mockclient.Contract.VerifyMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, arg5, arg6, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xf9bb5a51.
//
// Solidity: function verifyMembership(string clientId, (uint64,uint64) height, uint64 , uint64 , bytes proof, bytes , bytes , bytes value) view returns(bool)
func (_Mockclient *MockclientCallerSession) VerifyMembership(clientId string, height HeightData, arg2 uint64, arg3 uint64, proof []byte, arg5 []byte, arg6 []byte, value []byte) (bool, error) {
	return _Mockclient.Contract.VerifyMembership(&_Mockclient.CallOpts, clientId, height, arg2, arg3, proof, arg5, arg6, value)
}

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientTransactor) CreateClient(opts *bind.TransactOpts, clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "createClient", clientId, height, clientStateBytes, consensusStateBytes)
}

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientSession) CreateClient(clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.CreateClient(&_Mockclient.TransactOpts, clientId, height, clientStateBytes, consensusStateBytes)
}

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientTransactorSession) CreateClient(clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.CreateClient(&_Mockclient.TransactOpts, clientId, height, clientStateBytes, consensusStateBytes)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Mockclient.contract.Transact(opts, "updateClient", clientId, clientMessageBytes)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientSession) UpdateClient(clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.UpdateClient(&_Mockclient.TransactOpts, clientId, clientMessageBytes)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Mockclient *MockclientTransactorSession) UpdateClient(clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Mockclient.Contract.UpdateClient(&_Mockclient.TransactOpts, clientId, clientMessageBytes)
}
