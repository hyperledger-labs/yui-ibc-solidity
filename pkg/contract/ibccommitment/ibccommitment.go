// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibccommitment

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

// IbccommitmentABI is the input ABI used to generate the binding from.
const IbccommitmentABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"revisionNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revisionHeight\",\"type\":\"uint64\"}],\"name\":\"consensusStatePath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetReceiptCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"nextSequenceRecvCommitmentPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStateCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"revisionNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revisionHeight\",\"type\":\"uint64\"}],\"name\":\"consensusStateCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetReceiptCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"nextSequenceRecvCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

// Ibccommitment is an auto generated Go binding around an Ethereum contract.
type Ibccommitment struct {
	IbccommitmentCaller     // Read-only binding to the contract
	IbccommitmentTransactor // Write-only binding to the contract
	IbccommitmentFilterer   // Log filterer for contract events
}

// IbccommitmentCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbccommitmentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbccommitmentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbccommitmentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbccommitmentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbccommitmentSession struct {
	Contract     *Ibccommitment    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbccommitmentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbccommitmentCallerSession struct {
	Contract *IbccommitmentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbccommitmentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbccommitmentTransactorSession struct {
	Contract     *IbccommitmentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbccommitmentRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbccommitmentRaw struct {
	Contract *Ibccommitment // Generic contract binding to access the raw methods on
}

// IbccommitmentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbccommitmentCallerRaw struct {
	Contract *IbccommitmentCaller // Generic read-only contract binding to access the raw methods on
}

// IbccommitmentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbccommitmentTransactorRaw struct {
	Contract *IbccommitmentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbccommitment creates a new instance of Ibccommitment, bound to a specific deployed contract.
func NewIbccommitment(address common.Address, backend bind.ContractBackend) (*Ibccommitment, error) {
	contract, err := bindIbccommitment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibccommitment{IbccommitmentCaller: IbccommitmentCaller{contract: contract}, IbccommitmentTransactor: IbccommitmentTransactor{contract: contract}, IbccommitmentFilterer: IbccommitmentFilterer{contract: contract}}, nil
}

// NewIbccommitmentCaller creates a new read-only instance of Ibccommitment, bound to a specific deployed contract.
func NewIbccommitmentCaller(address common.Address, caller bind.ContractCaller) (*IbccommitmentCaller, error) {
	contract, err := bindIbccommitment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbccommitmentCaller{contract: contract}, nil
}

// NewIbccommitmentTransactor creates a new write-only instance of Ibccommitment, bound to a specific deployed contract.
func NewIbccommitmentTransactor(address common.Address, transactor bind.ContractTransactor) (*IbccommitmentTransactor, error) {
	contract, err := bindIbccommitment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbccommitmentTransactor{contract: contract}, nil
}

// NewIbccommitmentFilterer creates a new log filterer instance of Ibccommitment, bound to a specific deployed contract.
func NewIbccommitmentFilterer(address common.Address, filterer bind.ContractFilterer) (*IbccommitmentFilterer, error) {
	contract, err := bindIbccommitment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbccommitmentFilterer{contract: contract}, nil
}

// bindIbccommitment binds a generic wrapper to an already deployed contract.
func bindIbccommitment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbccommitmentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibccommitment *IbccommitmentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibccommitment.Contract.IbccommitmentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibccommitment *IbccommitmentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibccommitment.Contract.IbccommitmentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibccommitment *IbccommitmentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibccommitment.Contract.IbccommitmentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibccommitment *IbccommitmentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibccommitment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibccommitment *IbccommitmentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibccommitment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibccommitment *IbccommitmentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibccommitment.Contract.contract.Transact(opts, method, params...)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) ChannelCommitmentKey(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "channelCommitmentKey", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ChannelCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ChannelCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) ChannelPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "channelPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitment.Contract.ChannelPath(&_Ibccommitment.CallOpts, portId, channelId)
}

// ChannelPath is a free data retrieval call binding the contract method 0xf0381c21.
//
// Solidity: function channelPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) ChannelPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitment.Contract.ChannelPath(&_Ibccommitment.CallOpts, portId, channelId)
}

// ClientStateCommitmentKey is a free data retrieval call binding the contract method 0x05a1cdfa.
//
// Solidity: function clientStateCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) ClientStateCommitmentKey(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "clientStateCommitmentKey", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientStateCommitmentKey is a free data retrieval call binding the contract method 0x05a1cdfa.
//
// Solidity: function clientStateCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) ClientStateCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ClientStateCommitmentKey(&_Ibccommitment.CallOpts, clientId)
}

// ClientStateCommitmentKey is a free data retrieval call binding the contract method 0x05a1cdfa.
//
// Solidity: function clientStateCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) ClientStateCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ClientStateCommitmentKey(&_Ibccommitment.CallOpts, clientId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) ClientStatePath(opts *bind.CallOpts, clientId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "clientStatePath", clientId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibccommitment.Contract.ClientStatePath(&_Ibccommitment.CallOpts, clientId)
}

// ClientStatePath is a free data retrieval call binding the contract method 0xd3b78910.
//
// Solidity: function clientStatePath(string clientId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) ClientStatePath(clientId string) ([]byte, error) {
	return _Ibccommitment.Contract.ClientStatePath(&_Ibccommitment.CallOpts, clientId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) ConnectionCommitmentKey(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "connectionCommitmentKey", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ConnectionCommitmentKey(&_Ibccommitment.CallOpts, connectionId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibccommitment.Contract.ConnectionCommitmentKey(&_Ibccommitment.CallOpts, connectionId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) ConnectionPath(opts *bind.CallOpts, connectionId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "connectionPath", connectionId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibccommitment.Contract.ConnectionPath(&_Ibccommitment.CallOpts, connectionId)
}

// ConnectionPath is a free data retrieval call binding the contract method 0x5fe6557a.
//
// Solidity: function connectionPath(string connectionId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) ConnectionPath(connectionId string) ([]byte, error) {
	return _Ibccommitment.Contract.ConnectionPath(&_Ibccommitment.CallOpts, connectionId)
}

// ConsensusStateCommitmentKey is a free data retrieval call binding the contract method 0x0aa37e16.
//
// Solidity: function consensusStateCommitmentKey(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) ConsensusStateCommitmentKey(opts *bind.CallOpts, clientId string, revisionNumber uint64, revisionHeight uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "consensusStateCommitmentKey", clientId, revisionNumber, revisionHeight)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusStateCommitmentKey is a free data retrieval call binding the contract method 0x0aa37e16.
//
// Solidity: function consensusStateCommitmentKey(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) ConsensusStateCommitmentKey(clientId string, revisionNumber uint64, revisionHeight uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.ConsensusStateCommitmentKey(&_Ibccommitment.CallOpts, clientId, revisionNumber, revisionHeight)
}

// ConsensusStateCommitmentKey is a free data retrieval call binding the contract method 0x0aa37e16.
//
// Solidity: function consensusStateCommitmentKey(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) ConsensusStateCommitmentKey(clientId string, revisionNumber uint64, revisionHeight uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.ConsensusStateCommitmentKey(&_Ibccommitment.CallOpts, clientId, revisionNumber, revisionHeight)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) ConsensusStatePath(opts *bind.CallOpts, clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "consensusStatePath", clientId, revisionNumber, revisionHeight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibccommitment.Contract.ConsensusStatePath(&_Ibccommitment.CallOpts, clientId, revisionNumber, revisionHeight)
}

// ConsensusStatePath is a free data retrieval call binding the contract method 0x6dfb92e3.
//
// Solidity: function consensusStatePath(string clientId, uint64 revisionNumber, uint64 revisionHeight) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) ConsensusStatePath(clientId string, revisionNumber uint64, revisionHeight uint64) ([]byte, error) {
	return _Ibccommitment.Contract.ConsensusStatePath(&_Ibccommitment.CallOpts, clientId, revisionNumber, revisionHeight)
}

// NextSequenceRecvCommitmentKey is a free data retrieval call binding the contract method 0xfceb03ab.
//
// Solidity: function nextSequenceRecvCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) NextSequenceRecvCommitmentKey(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "nextSequenceRecvCommitmentKey", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextSequenceRecvCommitmentKey is a free data retrieval call binding the contract method 0xfceb03ab.
//
// Solidity: function nextSequenceRecvCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) NextSequenceRecvCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibccommitment.Contract.NextSequenceRecvCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId)
}

// NextSequenceRecvCommitmentKey is a free data retrieval call binding the contract method 0xfceb03ab.
//
// Solidity: function nextSequenceRecvCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) NextSequenceRecvCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibccommitment.Contract.NextSequenceRecvCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) NextSequenceRecvCommitmentPath(opts *bind.CallOpts, portId string, channelId string) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "nextSequenceRecvCommitmentPath", portId, channelId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitment.Contract.NextSequenceRecvCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId)
}

// NextSequenceRecvCommitmentPath is a free data retrieval call binding the contract method 0x5209b2dd.
//
// Solidity: function nextSequenceRecvCommitmentPath(string portId, string channelId) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) NextSequenceRecvCommitmentPath(portId string, channelId string) ([]byte, error) {
	return _Ibccommitment.Contract.NextSequenceRecvCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) PacketAcknowledgementCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetAcknowledgementCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketAcknowledgementCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketAcknowledgementCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) PacketAcknowledgementCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetAcknowledgementCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketAcknowledgementCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentPath is a free data retrieval call binding the contract method 0xc53bb5b3.
//
// Solidity: function packetAcknowledgementCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) PacketAcknowledgementCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketAcknowledgementCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) PacketCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) PacketCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentPath is a free data retrieval call binding the contract method 0x20eda6aa.
//
// Solidity: function packetCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) PacketCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentKey is a free data retrieval call binding the contract method 0x83c28eac.
//
// Solidity: function packetReceiptCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCaller) PacketReceiptCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetReceiptCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketReceiptCommitmentKey is a free data retrieval call binding the contract method 0x83c28eac.
//
// Solidity: function packetReceiptCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentSession) PacketReceiptCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketReceiptCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentKey is a free data retrieval call binding the contract method 0x83c28eac.
//
// Solidity: function packetReceiptCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibccommitment *IbccommitmentCallerSession) PacketReceiptCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibccommitment.Contract.PacketReceiptCommitmentKey(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCaller) PacketReceiptCommitmentPath(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([]byte, error) {
	var out []interface{}
	err := _Ibccommitment.contract.Call(opts, &out, "packetReceiptCommitmentPath", portId, channelId, sequence)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketReceiptCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}

// PacketReceiptCommitmentPath is a free data retrieval call binding the contract method 0x11118772.
//
// Solidity: function packetReceiptCommitmentPath(string portId, string channelId, uint64 sequence) pure returns(bytes)
func (_Ibccommitment *IbccommitmentCallerSession) PacketReceiptCommitmentPath(portId string, channelId string, sequence uint64) ([]byte, error) {
	return _Ibccommitment.Contract.PacketReceiptCommitmentPath(&_Ibccommitment.CallOpts, portId, channelId, sequence)
}
