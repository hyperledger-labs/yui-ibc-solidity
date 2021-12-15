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

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcidentifierABI is the input ABI used to generate the binding from.
const IbcidentifierABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"consensusCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"consensusStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"portCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

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

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ChannelCommitmentKey(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "channelCommitmentKey", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ChannelCommitmentSlot(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "channelCommitmentSlot", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ChannelCommitmentSlot(portId string, channelId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ChannelCommitmentSlot(portId string, channelId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ClientCommitmentKey(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "clientCommitmentKey", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ClientCommitmentKey(&_Ibcidentifier.CallOpts, clientId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ClientCommitmentKey(&_Ibcidentifier.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ClientStateCommitmentSlot(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "clientStateCommitmentSlot", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ClientStateCommitmentSlot(&_Ibcidentifier.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ClientStateCommitmentSlot(&_Ibcidentifier.CallOpts, clientId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ConnectionCommitmentKey(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "connectionCommitmentKey", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentKey(&_Ibcidentifier.CallOpts, connectionId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentKey(&_Ibcidentifier.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ConnectionCommitmentSlot(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "connectionCommitmentSlot", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentSlot(&_Ibcidentifier.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentSlot(&_Ibcidentifier.CallOpts, connectionId)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ConsensusCommitmentKey(opts *bind.CallOpts, clientId string, height HeightData) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "consensusCommitmentKey", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConsensusCommitmentKey(clientId string, height HeightData) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConsensusCommitmentKey(&_Ibcidentifier.CallOpts, clientId, height)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ConsensusCommitmentKey(clientId string, height HeightData) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConsensusCommitmentKey(&_Ibcidentifier.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) ConsensusStateCommitmentSlot(opts *bind.CallOpts, clientId string, height HeightData) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "consensusStateCommitmentSlot", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConsensusStateCommitmentSlot(clientId string, height HeightData) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConsensusStateCommitmentSlot(&_Ibcidentifier.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, (uint64,uint64) height) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) ConsensusStateCommitmentSlot(clientId string, height HeightData) ([32]byte, error) {
	return _Ibcidentifier.Contract.ConsensusStateCommitmentSlot(&_Ibcidentifier.CallOpts, clientId, height)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) PacketAcknowledgementCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetAcknowledgementCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) PacketAcknowledgementCommitmentSlot(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetAcknowledgementCommitmentSlot", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) PacketCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentKey(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCaller) PacketCommitmentSlot(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcidentifier.contract.Call(opts, &out, "packetCommitmentSlot", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcidentifier *IbcidentifierCallerSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcidentifier.Contract.PacketCommitmentSlot(&_Ibcidentifier.CallOpts, portId, channelId, sequence)
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
