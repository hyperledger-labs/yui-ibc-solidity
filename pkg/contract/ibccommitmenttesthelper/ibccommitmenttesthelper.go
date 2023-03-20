// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibccommitmenttesthelper

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

// IbccommitmenttesthelperMetaData contains all meta data concerning the Ibccommitmenttesthelper contract.
var IbccommitmenttesthelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"revisionNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revisionHeight\",\"type\":\"uint64\"}],\"name\":\"consensusStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetReceiptCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"nextSequenceRecvCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]",
}

// IbccommitmenttesthelperABI is the input ABI used to generate the binding from.
// Deprecated: Use IbccommitmenttesthelperMetaData.ABI instead.
var IbccommitmenttesthelperABI = IbccommitmenttesthelperMetaData.ABI

// Ibccommitmenttesthelper is an auto generated Go binding around an Ethereum contract.
type Ibccommitmenttesthelper struct {
	IbccommitmenttesthelperCaller     // Read-only binding to the contract
	IbccommitmenttesthelperTransactor // Write-only binding to the contract
	IbccommitmenttesthelperFilterer   // Log filterer for contract events
}

// IbccommitmenttesthelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbccommitmenttesthelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmenttesthelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbccommitmenttesthelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmenttesthelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbccommitmenttesthelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmenttesthelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbccommitmenttesthelperSession struct {
	Contract     *Ibccommitmenttesthelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbccommitmenttesthelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbccommitmenttesthelperCallerSession struct {
	Contract *IbccommitmenttesthelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IbccommitmenttesthelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbccommitmenttesthelperTransactorSession struct {
	Contract     *IbccommitmenttesthelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IbccommitmenttesthelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbccommitmenttesthelperRaw struct {
	Contract *Ibccommitmenttesthelper // Generic contract binding to access the raw methods on
}

// IbccommitmenttesthelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbccommitmenttesthelperCallerRaw struct {
	Contract *IbccommitmenttesthelperCaller // Generic read-only contract binding to access the raw methods on
}

// IbccommitmenttesthelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbccommitmenttesthelperTransactorRaw struct {
	Contract *IbccommitmenttesthelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbccommitmenttesthelper creates a new instance of Ibccommitmenttesthelper, bound to a specific deployed contract.
func NewIbccommitmenttesthelper(address common.Address, backend bind.ContractBackend) (*Ibccommitmenttesthelper, error) {
	contract, err := bindIbccommitmenttesthelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibccommitmenttesthelper{IbccommitmenttesthelperCaller: IbccommitmenttesthelperCaller{contract: contract}, IbccommitmenttesthelperTransactor: IbccommitmenttesthelperTransactor{contract: contract}, IbccommitmenttesthelperFilterer: IbccommitmenttesthelperFilterer{contract: contract}}, nil
}

// NewIbccommitmenttesthelperCaller creates a new read-only instance of Ibccommitmenttesthelper, bound to a specific deployed contract.
func NewIbccommitmenttesthelperCaller(address common.Address, caller bind.ContractCaller) (*IbccommitmenttesthelperCaller, error) {
	contract, err := bindIbccommitmenttesthelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbccommitmenttesthelperCaller{contract: contract}, nil
}

// NewIbccommitmenttesthelperTransactor creates a new write-only instance of Ibccommitmenttesthelper, bound to a specific deployed contract.
func NewIbccommitmenttesthelperTransactor(address common.Address, transactor bind.ContractTransactor) (*IbccommitmenttesthelperTransactor, error) {
	contract, err := bindIbccommitmenttesthelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbccommitmenttesthelperTransactor{contract: contract}, nil
}

// NewIbccommitmenttesthelperFilterer creates a new log filterer instance of Ibccommitmenttesthelper, bound to a specific deployed contract.
func NewIbccommitmenttesthelperFilterer(address common.Address, filterer bind.ContractFilterer) (*IbccommitmenttesthelperFilterer, error) {
	contract, err := bindIbccommitmenttesthelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbccommitmenttesthelperFilterer{contract: contract}, nil
}

// bindIbccommitmenttesthelper binds a generic wrapper to an already deployed contract.
func bindIbccommitmenttesthelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbccommitmenttesthelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibccommitmenttesthelper.Contract.IbccommitmenttesthelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibccommitmenttesthelper.Contract.IbccommitmenttesthelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibccommitmenttesthelper.Contract.IbccommitmenttesthelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibccommitmenttesthelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibccommitmenttesthelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibccommitmenttesthelper.Contract.contract.Transact(opts, method, params...)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) ChannelPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "channelPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ChannelPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ChannelPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) ClientStatePath(opts *bind.CallOpts, clientId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "clientStatePath", clientId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ClientStatePath(&_Ibccommitmenttesthelper.CallOpts, clientId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ClientStatePath(&_Ibccommitmenttesthelper.CallOpts, clientId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) ConnectionPath(opts *bind.CallOpts, connectionId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "connectionPath", connectionId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ConnectionPath(&_Ibccommitmenttesthelper.CallOpts, connectionId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ConnectionPath(&_Ibccommitmenttesthelper.CallOpts, connectionId)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) ConsensusStatePath(opts *bind.CallOpts, clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "consensusStatePath", clientId, revisionNumber, revisionHeight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ConsensusStatePath(&_Ibccommitmenttesthelper.CallOpts, clientId, revisionNumber, revisionHeight)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.ConsensusStatePath(&_Ibccommitmenttesthelper.CallOpts, clientId, revisionNumber, revisionHeight)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) NextSequenceRecvCommitmentPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "nextSequenceRecvCommitmentPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.NextSequenceRecvCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.NextSequenceRecvCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) PacketAcknowledgementCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "packetAcknowledgementCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketAcknowledgementCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketAcknowledgementCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) PacketCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "packetCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCaller) PacketReceiptCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitmenttesthelper.contract.Call(opts, &out, "packetReceiptCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketReceiptCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitmenttesthelper *IbccommitmenttesthelperCallerSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitmenttesthelper.Contract.PacketReceiptCommitmentPath(&_Ibccommitmenttesthelper.CallOpts, portId, channelId, sequence)
}
