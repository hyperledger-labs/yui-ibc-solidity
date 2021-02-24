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

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	ChainId         string
	IbcStoreAddress []byte
	LatestHeight    uint64
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Timestamp  uint64
	Root       []byte
	Validators [][]byte
}

// Ibft2clientABI is the input ABI used to generate the binding from.
const Ibft2clientABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getTimestampAtHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerBytes\",\"type\":\"bytes\"}],\"name\":\"checkHeaderAndUpdateState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newClientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newConsensusStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientConsensusState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"verifyConnectionState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channelBytes\",\"type\":\"bytes\"}],\"name\":\"verifyChannelState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"ackCommitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketAcknowledgement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

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

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xc8abd8d3.
//
// Solidity: function checkHeaderAndUpdateState(string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Ibft2client *Ibft2clientCaller) CheckHeaderAndUpdateState(opts *bind.CallOpts, clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "checkHeaderAndUpdateState", clientId, clientStateBytes, headerBytes)

	outstruct := new(struct {
		NewClientStateBytes    []byte
		NewConsensusStateBytes []byte
		Height                 uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewClientStateBytes = out[0].([]byte)
	outstruct.NewConsensusStateBytes = out[1].([]byte)
	outstruct.Height = out[2].(uint64)

	return *outstruct, err

}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xc8abd8d3.
//
// Solidity: function checkHeaderAndUpdateState(string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Ibft2client *Ibft2clientSession) CheckHeaderAndUpdateState(clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	return _Ibft2client.Contract.CheckHeaderAndUpdateState(&_Ibft2client.CallOpts, clientId, clientStateBytes, headerBytes)
}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xc8abd8d3.
//
// Solidity: function checkHeaderAndUpdateState(string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Ibft2client *Ibft2clientCallerSession) CheckHeaderAndUpdateState(clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	return _Ibft2client.Contract.CheckHeaderAndUpdateState(&_Ibft2client.CallOpts, clientId, clientStateBytes, headerBytes)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64) clientState)
func (_Ibft2client *Ibft2clientCaller) GetClientState(opts *bind.CallOpts, clientId string) (ClientStateData, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new(ClientStateData), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientStateData)).(*ClientStateData)

	return out0, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64) clientState)
func (_Ibft2client *Ibft2clientSession) GetClientState(clientId string) (ClientStateData, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64) clientState)
func (_Ibft2client *Ibft2clientCallerSession) GetClientState(clientId string) (ClientStateData, error) {
	return _Ibft2client.Contract.GetClientState(&_Ibft2client.CallOpts, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]))
func (_Ibft2client *Ibft2clientCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height uint64) (ConsensusStateData, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new(ConsensusStateData), err
	}

	out0 := *abi.ConvertType(out[0], new(ConsensusStateData)).(*ConsensusStateData)

	return out0, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]))
func (_Ibft2client *Ibft2clientSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]))
func (_Ibft2client *Ibft2clientCallerSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, error) {
	return _Ibft2client.Contract.GetConsensusState(&_Ibft2client.CallOpts, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCaller) GetLatestHeight(opts *bind.CallOpts, clientId string) (uint64, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getLatestHeight", clientId)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientSession) GetLatestHeight(clientId string) (uint64, bool, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string clientId) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCallerSession) GetLatestHeight(clientId string) (uint64, bool, error) {
	return _Ibft2client.Contract.GetLatestHeight(&_Ibft2client.CallOpts, clientId)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xfa74f2f4.
//
// Solidity: function getTimestampAtHeight(string clientId, uint64 height) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCaller) GetTimestampAtHeight(opts *bind.CallOpts, clientId string, height uint64) (uint64, bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "getTimestampAtHeight", clientId, height)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xfa74f2f4.
//
// Solidity: function getTimestampAtHeight(string clientId, uint64 height) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientSession) GetTimestampAtHeight(clientId string, height uint64) (uint64, bool, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xfa74f2f4.
//
// Solidity: function getTimestampAtHeight(string clientId, uint64 height) view returns(uint64, bool)
func (_Ibft2client *Ibft2clientCallerSession) GetTimestampAtHeight(clientId string, height uint64) (uint64, bool, error) {
	return _Ibft2client.Contract.GetTimestampAtHeight(&_Ibft2client.CallOpts, clientId, height)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0xf391a09f.
//
// Solidity: function verifyChannelState(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyChannelState(opts *bind.CallOpts, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyChannelState", clientId, height, prefix, proof, portId, channelId, channelBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyChannelState is a free data retrieval call binding the contract method 0xf391a09f.
//
// Solidity: function verifyChannelState(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyChannelState(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyChannelState(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0xf391a09f.
//
// Solidity: function verifyChannelState(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyChannelState(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyChannelState(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xa3fc0907.
//
// Solidity: function verifyClientConsensusState(string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyClientConsensusState(opts *bind.CallOpts, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyClientConsensusState", clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xa3fc0907.
//
// Solidity: function verifyClientConsensusState(string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyClientConsensusState(clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyClientConsensusState(&_Ibft2client.CallOpts, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xa3fc0907.
//
// Solidity: function verifyClientConsensusState(string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyClientConsensusState(clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyClientConsensusState(&_Ibft2client.CallOpts, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x19db89f4.
//
// Solidity: function verifyClientState(string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyClientState(opts *bind.CallOpts, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyClientState", clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientState is a free data retrieval call binding the contract method 0x19db89f4.
//
// Solidity: function verifyClientState(string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyClientState(clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyClientState(&_Ibft2client.CallOpts, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x19db89f4.
//
// Solidity: function verifyClientState(string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyClientState(clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyClientState(&_Ibft2client.CallOpts, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x358f307a.
//
// Solidity: function verifyConnectionState(string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyConnectionState(opts *bind.CallOpts, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyConnectionState", clientId, height, prefix, proof, connectionId, connectionBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x358f307a.
//
// Solidity: function verifyConnectionState(string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyConnectionState(clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyConnectionState(&_Ibft2client.CallOpts, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x358f307a.
//
// Solidity: function verifyConnectionState(string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyConnectionState(clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Ibft2client.Contract.VerifyConnectionState(&_Ibft2client.CallOpts, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x3017ad22.
//
// Solidity: function verifyPacketAcknowledgement(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyPacketAcknowledgement(opts *bind.CallOpts, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyPacketAcknowledgement", clientId, height, prefix, proof, portId, channelId, sequence, ackCommitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x3017ad22.
//
// Solidity: function verifyPacketAcknowledgement(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyPacketAcknowledgement(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	return _Ibft2client.Contract.VerifyPacketAcknowledgement(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, sequence, ackCommitmentBytes)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x3017ad22.
//
// Solidity: function verifyPacketAcknowledgement(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyPacketAcknowledgement(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	return _Ibft2client.Contract.VerifyPacketAcknowledgement(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, sequence, ackCommitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xb1aeea32.
//
// Solidity: function verifyPacketCommitment(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCaller) VerifyPacketCommitment(opts *bind.CallOpts, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Ibft2client.contract.Call(opts, &out, "verifyPacketCommitment", clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xb1aeea32.
//
// Solidity: function verifyPacketCommitment(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientSession) VerifyPacketCommitment(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Ibft2client.Contract.VerifyPacketCommitment(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xb1aeea32.
//
// Solidity: function verifyPacketCommitment(string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibft2client *Ibft2clientCallerSession) VerifyPacketCommitment(clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Ibft2client.Contract.VerifyPacketCommitment(&_Ibft2client.CallOpts, clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)
}
