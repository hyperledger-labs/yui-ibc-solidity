// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcidentifier

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

// IbcidentifierABI is the input ABI used to generate the binding from.
const IbcidentifierABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"revisionNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revisionHeight\",\"type\":\"uint64\"}],\"name\":\"consensusStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetReceiptCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"nextSequenceRecvCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"portCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

// Ibcidentifier is an auto generated Go binding around an Ethereum contract.
type Ibcidentifier struct {
	IbcidentifierCaller     // Read-only binding to the contract
	IbcidentifierTransactor // Write-only binding to the contract
	IbcidentifierFilterer   // Log filterer for contract events
}

// IbcidentifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcidentifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcidentifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcidentifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcidentifierSession struct {
	Contract     *Ibcidentifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcidentifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcidentifierCallerSession struct {
	Contract *IbcidentifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbcidentifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcidentifierTransactorSession struct {
	Contract     *IbcidentifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbcidentifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcidentifierRaw struct {
	Contract *Ibcidentifier // Generic contract binding to access the raw methods on
}

// IbcidentifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcidentifierCallerRaw struct {
	Contract *IbcidentifierCaller // Generic read-only contract binding to access the raw methods on
}

// IbcidentifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcidentifierTransactorRaw struct {
	Contract *IbcidentifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcidentifier creates a new instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifier(address common.Address, backend bind.ContractBackend) (*Ibcidentifier, error) {
	contract, err := bindIbcidentifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcidentifier{IbcidentifierCaller: IbcidentifierCaller{contract: contract}, IbcidentifierTransactor: IbcidentifierTransactor{contract: contract}, IbcidentifierFilterer: IbcidentifierFilterer{contract: contract}}, nil
}

// NewIbcidentifierCaller creates a new read-only instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierCaller(address common.Address, caller bind.ContractCaller) (*IbcidentifierCaller, error) {
	contract, err := bindIbcidentifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierCaller{contract: contract}, nil
}

// NewIbcidentifierTransactor creates a new write-only instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcidentifierTransactor, error) {
	contract, err := bindIbcidentifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierTransactor{contract: contract}, nil
}

// NewIbcidentifierFilterer creates a new log filterer instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcidentifierFilterer, error) {
	contract, err := bindIbcidentifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierFilterer{contract: contract}, nil
}

// bindIbcidentifier binds a generic wrapper to an already deployed contract.
func bindIbcidentifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcidentifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcidentifier *IbcidentifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcidentifier.Contract.IbcidentifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcidentifier *IbcidentifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.IbcidentifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcidentifier *IbcidentifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.IbcidentifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcidentifier *IbcidentifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcidentifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcidentifier *IbcidentifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcidentifier *IbcidentifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.contract.Transact(opts, method, params...)
}

// ChannelCapabilityPath is a free data retrieval call binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) ChannelCapabilityPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "channelCapabilityPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ChannelCapabilityPath is a free data retrieval call binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ChannelCapabilityPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ChannelCapabilityPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelCapabilityPath is a free data retrieval call binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) ChannelCapabilityPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ChannelCapabilityPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) ChannelPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "channelPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ChannelPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ChannelPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) ClientStatePath(opts *bind.CallOpts, clientId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "clientStatePath", clientId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ClientStatePath(&_Ibcidentifier.CallOpts, clientId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ClientStatePath(&_Ibcidentifier.CallOpts, clientId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) ConnectionPath(opts *bind.CallOpts, connectionId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "connectionPath", connectionId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ConnectionPath(&_Ibcidentifier.CallOpts, connectionId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibcidentifier.Contract.ConnectionPath(&_Ibcidentifier.CallOpts, connectionId)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) ConsensusStatePath(opts *bind.CallOpts, clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "consensusStatePath", clientId, revisionNumber, revisionHeight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.ConsensusStatePath(&_Ibcidentifier.CallOpts, clientId, revisionNumber, revisionHeight)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.ConsensusStatePath(&_Ibcidentifier.CallOpts, clientId, revisionNumber, revisionHeight)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) NextSequenceRecvCommitmentPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "nextSequenceRecvCommitmentPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.NextSequenceRecvCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibcidentifier.Contract.NextSequenceRecvCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) PacketAcknowledgementCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetAcknowledgementCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) PacketCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) PacketReceiptCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetReceiptCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketReceiptCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibcidentifier.Contract.PacketReceiptCommitmentPath(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PortCapabilityPath is a free data retrieval call binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCaller) PortCapabilityPath(opts *bind.CallOpts, portId string) ([]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "portCapabilityPath", portId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PortCapabilityPath is a free data retrieval call binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) PortCapabilityPath(portId string) ([]byte, error) {
	return _Ibcidentifier.Contract.PortCapabilityPath(&_Ibcidentifier.CallOpts, portId)
}

// PortCapabilityPath is a free data retrieval call binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) pure returns(bytes)
func (_Ibcidentifier *IbcidentifierCallerSession) PortCapabilityPath(portId string) ([]byte, error) {
	return _Ibcidentifier.Contract.PortCapabilityPath(&_Ibcidentifier.CallOpts, portId)
}
