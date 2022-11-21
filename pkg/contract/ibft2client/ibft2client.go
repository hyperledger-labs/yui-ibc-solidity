// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibft2client

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

// IbcLightclientsIbft2V1ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsIbft2V1ClientStateData struct {
	ChainId         string
	IbcStoreAddress []byte
	LatestHeight    HeightData
}

// IbcLightclientsIbft2V1ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsIbft2V1ConsensusStateData struct {
	Timestamp  uint64
	Root       []byte
	Validators [][]byte
}

// Ibft2clientABI is the input ABI used to generate the binding from.
const Ibft2clientABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"clientStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"consensusStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"internalType\":\"structConsensusStateUpdates[]\",\"name\":\"updates\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getTimestampAtHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientMessageBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientMessageAndUpdateState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"clientStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"consensusStateCommitment\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"internalType\":\"structConsensusStateUpdates[]\",\"name\":\"updates\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayTimePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"delayBlockPeriod\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"latest_height\",\"type\":\"tuple\"}],\"internalType\":\"structIbcLightclientsIbft2V1ClientState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structIbcLightclientsIbft2V1ConsensusState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

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
	parsed, err := abi.JSON(strings.NewReader(Ibft2clientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
// Solidity: function getClientState(string clientId) view returns((string,bytes,(uint64,uint64)), bool)
func (_Ibft2client *Ibft2clientCaller) GetClientState(opts *bind.CallOpts, clientId string) (IbcLightclientsIbft2V1ClientStateData, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new(IbcLightclientsIbft2V1ClientStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(IbcLightclientsIbft2V1ClientStateData)).(*IbcLightclientsIbft2V1ClientStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,(uint64,uint64)), bool)
func (_Ibft2client *Ibft2clientSession) GetClientState(clientId string) (IbcLightclientsIbft2V1ClientStateData, bool, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,(uint64,uint64)), bool)
func (_Ibft2client *Ibft2clientCallerSession) GetClientState(clientId string) (IbcLightclientsIbft2V1ClientStateData, bool, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns((uint64,bytes,bytes[]), bool)
func (_Ibft2client *Ibft2clientCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height HeightData) (IbcLightclientsIbft2V1ConsensusStateData, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new(IbcLightclientsIbft2V1ConsensusStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(IbcLightclientsIbft2V1ConsensusStateData)).(*IbcLightclientsIbft2V1ConsensusStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns((uint64,bytes,bytes[]), bool)
func (_Ibft2client *Ibft2clientSession) GetConsensusState(clientId string, height HeightData) (IbcLightclientsIbft2V1ConsensusStateData, bool, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns((uint64,bytes,bytes[]), bool)
func (_Ibft2client *Ibft2clientCallerSession) GetConsensusState(clientId string, height HeightData) (IbcLightclientsIbft2V1ConsensusStateData, bool, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64), bool)
func (_Ibft2client *Ibft2clientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (HeightData, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getLatestHeight", clientId)

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
func (_Ibft2client *Ibft2clientSession) GetLatestHeight(clientId string) (HeightData, bool, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns((uint64,uint64), bool)
func (_Ibft2client *Ibft2clientCallerSession) GetLatestHeight(clientId string) (HeightData, bool, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height HeightData) (uint64, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

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
func (_Ibft2client *Ibft2clientSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, bool, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0x4b0bbdc4.
//
// Solidity: function getTimestampAtHeight(string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCallerSession) GetTimestampAtHeight(clientId string, height HeightData) (uint64, bool, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
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

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientTransactor) CreateClient(opts *bind.TransactOpts, clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.contract.Transact(opts, "createClient", clientId, height, clientStateBytes, consensusStateBytes)
}

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientSession) CreateClient(clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.CreateClient(&_Ibft2client.TransactOpts, clientId, height, clientStateBytes, consensusStateBytes)
}

// CreateClient is a paid mutator transaction binding the contract method 0x1f9a5400.
//
// Solidity: function createClient(string clientId, (uint64,uint64) height, bytes clientStateBytes, bytes consensusStateBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientTransactorSession) CreateClient(clientId string, height HeightData, clientStateBytes []byte, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.CreateClient(&_Ibft2client.TransactOpts, clientId, height, clientStateBytes, consensusStateBytes)
}

// VerifyClientMessageAndUpdateState is a paid mutator transaction binding the contract method 0xf9ec3614.
//
// Solidity: function verifyClientMessageAndUpdateState(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientTransactor) VerifyClientMessageAndUpdateState(opts *bind.TransactOpts, clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.contract.Transact(opts, "verifyClientMessageAndUpdateState", clientId, clientMessageBytes)
}

// VerifyClientMessageAndUpdateState is a paid mutator transaction binding the contract method 0xf9ec3614.
//
// Solidity: function verifyClientMessageAndUpdateState(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientSession) VerifyClientMessageAndUpdateState(clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.VerifyClientMessageAndUpdateState(&_Ibft2client.TransactOpts, clientId, clientMessageBytes)
}

// VerifyClientMessageAndUpdateState is a paid mutator transaction binding the contract method 0xf9ec3614.
//
// Solidity: function verifyClientMessageAndUpdateState(string clientId, bytes clientMessageBytes) returns(bytes32 clientStateCommitment, (bytes32,(uint64,uint64))[] updates, bool ok)
func (_Ibft2client *Ibft2clientTransactorSession) VerifyClientMessageAndUpdateState(clientId string, clientMessageBytes []byte) (*types.Transaction, error) {
	return _Ibft2client.Contract.VerifyClientMessageAndUpdateState(&_Ibft2client.TransactOpts, clientId, clientMessageBytes)
}
