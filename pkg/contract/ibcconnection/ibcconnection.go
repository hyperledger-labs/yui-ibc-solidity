// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcconnection

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

// ConnectionEndData is an auto generated low-level Go binding around an user-defined struct.
type ConnectionEndData struct {
	ClientId     string
	Versions     []VersionData
	State        uint8
	DelayPeriod  uint64
	Counterparty CounterpartyData
}

// CounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type CounterpartyData struct {
	ClientId     string
	ConnectionId string
	Prefix       MerklePrefixData
}

// IBCConnectionMsgConnectionOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IBCConnectionMsgConnectionOpenAck struct {
	ConnectionId             string
	ClientState              ClientStateData
	Version                  VersionData
	CounterpartyConnectionID string
	ProofTry                 []byte
	ProofClient              []byte
	ProofConsensus           []byte
	ProofHeight              uint64
	ConsensusHeight          uint64
}

// IBCConnectionMsgConnectionOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IBCConnectionMsgConnectionOpenConfirm struct {
	ConnectionId string
	ProofAck     []byte
	ProofHeight  uint64
}

// IBCConnectionMsgConnectionOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCConnectionMsgConnectionOpenTry struct {
	ConnectionId         string
	Counterparty         CounterpartyData
	DelayPeriod          uint64
	ClientId             string
	ClientState          ClientStateData
	CounterpartyVersions []VersionData
	ProofInit            []byte
	ProofClient          []byte
	ProofConsensus       []byte
	ProofHeight          uint64
	ConsensusHeight      uint64
}

// MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type MerklePrefixData struct {
	KeyPrefix []byte
}

// VersionData is an auto generated low-level Go binding around an user-defined struct.
type VersionData struct {
	Identifier string
	Features   []string
}

// IbcconnectionABI is the input ABI used to generate the binding from.
const IbcconnectionABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBCClient\",\"name\":\"client_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"}],\"name\":\"connectionOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"counterpartyVersions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCConnection.MsgConnectionOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data\",\"name\":\"version\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyConnectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCConnection.MsgConnectionOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCConnection.MsgConnectionOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"counterpartyConnection\",\"type\":\"tuple\"}],\"name\":\"verifyConnectionState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channelBytes\",\"type\":\"bytes\"}],\"name\":\"verifyChannelState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"ackCommitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketAcknowledgement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Ibcconnection is an auto generated Go binding around an Ethereum contract.
type Ibcconnection struct {
	IbcconnectionCaller     // Read-only binding to the contract
	IbcconnectionTransactor // Write-only binding to the contract
	IbcconnectionFilterer   // Log filterer for contract events
}

// IbcconnectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcconnectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcconnectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcconnectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcconnectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcconnectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcconnectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcconnectionSession struct {
	Contract     *Ibcconnection    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcconnectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcconnectionCallerSession struct {
	Contract *IbcconnectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbcconnectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcconnectionTransactorSession struct {
	Contract     *IbcconnectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbcconnectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcconnectionRaw struct {
	Contract *Ibcconnection // Generic contract binding to access the raw methods on
}

// IbcconnectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcconnectionCallerRaw struct {
	Contract *IbcconnectionCaller // Generic read-only contract binding to access the raw methods on
}

// IbcconnectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcconnectionTransactorRaw struct {
	Contract *IbcconnectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcconnection creates a new instance of Ibcconnection, bound to a specific deployed contract.
func NewIbcconnection(address common.Address, backend bind.ContractBackend) (*Ibcconnection, error) {
	contract, err := bindIbcconnection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcconnection{IbcconnectionCaller: IbcconnectionCaller{contract: contract}, IbcconnectionTransactor: IbcconnectionTransactor{contract: contract}, IbcconnectionFilterer: IbcconnectionFilterer{contract: contract}}, nil
}

// NewIbcconnectionCaller creates a new read-only instance of Ibcconnection, bound to a specific deployed contract.
func NewIbcconnectionCaller(address common.Address, caller bind.ContractCaller) (*IbcconnectionCaller, error) {
	contract, err := bindIbcconnection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcconnectionCaller{contract: contract}, nil
}

// NewIbcconnectionTransactor creates a new write-only instance of Ibcconnection, bound to a specific deployed contract.
func NewIbcconnectionTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcconnectionTransactor, error) {
	contract, err := bindIbcconnection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcconnectionTransactor{contract: contract}, nil
}

// NewIbcconnectionFilterer creates a new log filterer instance of Ibcconnection, bound to a specific deployed contract.
func NewIbcconnectionFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcconnectionFilterer, error) {
	contract, err := bindIbcconnection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcconnectionFilterer{contract: contract}, nil
}

// bindIbcconnection binds a generic wrapper to an already deployed contract.
func bindIbcconnection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcconnectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcconnection *IbcconnectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcconnection.Contract.IbcconnectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcconnection *IbcconnectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcconnection.Contract.IbcconnectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcconnection *IbcconnectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcconnection.Contract.IbcconnectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcconnection *IbcconnectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcconnection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcconnection *IbcconnectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcconnection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcconnection *IbcconnectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcconnection.Contract.contract.Transact(opts, method, params...)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0x5ebb4492.
//
// Solidity: function verifyChannelState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCaller) VerifyChannelState(opts *bind.CallOpts, connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibcconnection.contract.Call(opts, &out, "verifyChannelState", connection, height, proof, portId, channelId, channelBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyChannelState is a free data retrieval call binding the contract method 0x5ebb4492.
//
// Solidity: function verifyChannelState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionSession) VerifyChannelState(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyChannelState(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, channelBytes)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0x5ebb4492.
//
// Solidity: function verifyChannelState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCallerSession) VerifyChannelState(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyChannelState(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, channelBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa1165649.
//
// Solidity: function verifyConnectionState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) counterpartyConnection) view returns(bool)
func (_Ibcconnection *IbcconnectionCaller) VerifyConnectionState(opts *bind.CallOpts, connection ConnectionEndData, height uint64, proof []byte, connectionId string, counterpartyConnection ConnectionEndData) (bool, error) {
	var out []interface{}
	err := _Ibcconnection.contract.Call(opts, &out, "verifyConnectionState", connection, height, proof, connectionId, counterpartyConnection)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa1165649.
//
// Solidity: function verifyConnectionState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) counterpartyConnection) view returns(bool)
func (_Ibcconnection *IbcconnectionSession) VerifyConnectionState(connection ConnectionEndData, height uint64, proof []byte, connectionId string, counterpartyConnection ConnectionEndData) (bool, error) {
	return _Ibcconnection.Contract.VerifyConnectionState(&_Ibcconnection.CallOpts, connection, height, proof, connectionId, counterpartyConnection)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa1165649.
//
// Solidity: function verifyConnectionState((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) counterpartyConnection) view returns(bool)
func (_Ibcconnection *IbcconnectionCallerSession) VerifyConnectionState(connection ConnectionEndData, height uint64, proof []byte, connectionId string, counterpartyConnection ConnectionEndData) (bool, error) {
	return _Ibcconnection.Contract.VerifyConnectionState(&_Ibcconnection.CallOpts, connection, height, proof, connectionId, counterpartyConnection)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0xcbae9a10.
//
// Solidity: function verifyPacketAcknowledgement((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCaller) VerifyPacketAcknowledgement(opts *bind.CallOpts, connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Ibcconnection.contract.Call(opts, &out, "verifyPacketAcknowledgement", connection, height, proof, portId, channelId, sequence, ackCommitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0xcbae9a10.
//
// Solidity: function verifyPacketAcknowledgement((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionSession) VerifyPacketAcknowledgement(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyPacketAcknowledgement(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, sequence, ackCommitmentBytes)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0xcbae9a10.
//
// Solidity: function verifyPacketAcknowledgement((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 ackCommitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCallerSession) VerifyPacketAcknowledgement(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, ackCommitmentBytes [32]byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyPacketAcknowledgement(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, sequence, ackCommitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x6142e898.
//
// Solidity: function verifyPacketCommitment((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCaller) VerifyPacketCommitment(opts *bind.CallOpts, connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Ibcconnection.contract.Call(opts, &out, "verifyPacketCommitment", connection, height, proof, portId, channelId, sequence, commitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x6142e898.
//
// Solidity: function verifyPacketCommitment((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionSession) VerifyPacketCommitment(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyPacketCommitment(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x6142e898.
//
// Solidity: function verifyPacketCommitment((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, uint64 height, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Ibcconnection *IbcconnectionCallerSession) VerifyPacketCommitment(connection ConnectionEndData, height uint64, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Ibcconnection.Contract.VerifyPacketCommitment(&_Ibcconnection.CallOpts, connection, height, proof, portId, channelId, sequence, commitmentBytes)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xee7e215f.
//
// Solidity: function connectionOpenAck((string,(string,bytes,uint64),(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionTransactor) ConnectionOpenAck(opts *bind.TransactOpts, msg_ IBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcconnection.contract.Transact(opts, "connectionOpenAck", msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xee7e215f.
//
// Solidity: function connectionOpenAck((string,(string,bytes,uint64),(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionSession) ConnectionOpenAck(msg_ IBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenAck(&_Ibcconnection.TransactOpts, msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xee7e215f.
//
// Solidity: function connectionOpenAck((string,(string,bytes,uint64),(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionTransactorSession) ConnectionOpenAck(msg_ IBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenAck(&_Ibcconnection.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionTransactor) ConnectionOpenConfirm(opts *bind.TransactOpts, msg_ IBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcconnection.contract.Transact(opts, "connectionOpenConfirm", msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionSession) ConnectionOpenConfirm(msg_ IBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenConfirm(&_Ibcconnection.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcconnection *IbcconnectionTransactorSession) ConnectionOpenConfirm(msg_ IBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenConfirm(&_Ibcconnection.TransactOpts, msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe9f9caf7.
//
// Solidity: function connectionOpenInit(string clientId, string connectionId, (string,string,(bytes)) counterparty, uint64 delayPeriod) returns(string)
func (_Ibcconnection *IbcconnectionTransactor) ConnectionOpenInit(opts *bind.TransactOpts, clientId string, connectionId string, counterparty CounterpartyData, delayPeriod uint64) (*types.Transaction, error) {
	return _Ibcconnection.contract.Transact(opts, "connectionOpenInit", clientId, connectionId, counterparty, delayPeriod)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe9f9caf7.
//
// Solidity: function connectionOpenInit(string clientId, string connectionId, (string,string,(bytes)) counterparty, uint64 delayPeriod) returns(string)
func (_Ibcconnection *IbcconnectionSession) ConnectionOpenInit(clientId string, connectionId string, counterparty CounterpartyData, delayPeriod uint64) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenInit(&_Ibcconnection.TransactOpts, clientId, connectionId, counterparty, delayPeriod)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe9f9caf7.
//
// Solidity: function connectionOpenInit(string clientId, string connectionId, (string,string,(bytes)) counterparty, uint64 delayPeriod) returns(string)
func (_Ibcconnection *IbcconnectionTransactorSession) ConnectionOpenInit(clientId string, connectionId string, counterparty CounterpartyData, delayPeriod uint64) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenInit(&_Ibcconnection.TransactOpts, clientId, connectionId, counterparty, delayPeriod)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0xfdcacbb4.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,(string,bytes,uint64),(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcconnection *IbcconnectionTransactor) ConnectionOpenTry(opts *bind.TransactOpts, msg_ IBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcconnection.contract.Transact(opts, "connectionOpenTry", msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0xfdcacbb4.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,(string,bytes,uint64),(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcconnection *IbcconnectionSession) ConnectionOpenTry(msg_ IBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenTry(&_Ibcconnection.TransactOpts, msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0xfdcacbb4.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,(string,bytes,uint64),(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcconnection *IbcconnectionTransactorSession) ConnectionOpenTry(msg_ IBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcconnection.Contract.ConnectionOpenTry(&_Ibcconnection.TransactOpts, msg_)
}
