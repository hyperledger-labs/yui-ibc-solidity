// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcclient

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

// IBCMsgsMsgCreateClient is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgCreateClient struct {
	ClientId            string
	ClientType          string
	Height              uint64
	ClientStateBytes    []byte
	ConsensusStateBytes []byte
}

// IBCMsgsMsgUpdateClient is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgUpdateClient struct {
	ClientId string
	Header   []byte
}

// IbcclientABI is the input ABI used to generate the binding from.
const IbcclientABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"store\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgCreateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgUpdateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"}],\"name\":\"validateSelfClient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"contractIClient\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"registerClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClient\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Ibcclient is an auto generated Go binding around an Ethereum contract.
type Ibcclient struct {
	IbcclientCaller     // Read-only binding to the contract
	IbcclientTransactor // Write-only binding to the contract
	IbcclientFilterer   // Log filterer for contract events
}

// IbcclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcclientSession struct {
	Contract     *Ibcclient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcclientCallerSession struct {
	Contract *IbcclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IbcclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcclientTransactorSession struct {
	Contract     *IbcclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IbcclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcclientRaw struct {
	Contract *Ibcclient // Generic contract binding to access the raw methods on
}

// IbcclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcclientCallerRaw struct {
	Contract *IbcclientCaller // Generic read-only contract binding to access the raw methods on
}

// IbcclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcclientTransactorRaw struct {
	Contract *IbcclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcclient creates a new instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclient(address common.Address, backend bind.ContractBackend) (*Ibcclient, error) {
	contract, err := bindIbcclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcclient{IbcclientCaller: IbcclientCaller{contract: contract}, IbcclientTransactor: IbcclientTransactor{contract: contract}, IbcclientFilterer: IbcclientFilterer{contract: contract}}, nil
}

// NewIbcclientCaller creates a new read-only instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientCaller(address common.Address, caller bind.ContractCaller) (*IbcclientCaller, error) {
	contract, err := bindIbcclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcclientCaller{contract: contract}, nil
}

// NewIbcclientTransactor creates a new write-only instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcclientTransactor, error) {
	contract, err := bindIbcclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcclientTransactor{contract: contract}, nil
}

// NewIbcclientFilterer creates a new log filterer instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcclientFilterer, error) {
	contract, err := bindIbcclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcclientFilterer{contract: contract}, nil
}

// bindIbcclient binds a generic wrapper to an already deployed contract.
func bindIbcclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcclient *IbcclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcclient.Contract.IbcclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcclient *IbcclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcclient.Contract.IbcclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcclient *IbcclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcclient.Contract.IbcclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcclient *IbcclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcclient *IbcclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcclient *IbcclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcclient.Contract.contract.Transact(opts, method, params...)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibcclient *IbcclientCaller) GetClient(opts *bind.CallOpts, clientId string) (common.Address, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "getClient", clientId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibcclient *IbcclientSession) GetClient(clientId string) (common.Address, error) {
	return _Ibcclient.Contract.GetClient(&_Ibcclient.CallOpts, clientId)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibcclient *IbcclientCallerSession) GetClient(clientId string) (common.Address, error) {
	return _Ibcclient.Contract.GetClient(&_Ibcclient.CallOpts, clientId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcclient *IbcclientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcclient *IbcclientSession) Owner() (common.Address, error) {
	return _Ibcclient.Contract.Owner(&_Ibcclient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcclient *IbcclientCallerSession) Owner() (common.Address, error) {
	return _Ibcclient.Contract.Owner(&_Ibcclient.CallOpts)
}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xd08d8423.
//
// Solidity: function validateSelfClient(bytes clientStateBytes) view returns(bool)
func (_Ibcclient *IbcclientCaller) ValidateSelfClient(opts *bind.CallOpts, clientStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "validateSelfClient", clientStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xd08d8423.
//
// Solidity: function validateSelfClient(bytes clientStateBytes) view returns(bool)
func (_Ibcclient *IbcclientSession) ValidateSelfClient(clientStateBytes []byte) (bool, error) {
	return _Ibcclient.Contract.ValidateSelfClient(&_Ibcclient.CallOpts, clientStateBytes)
}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xd08d8423.
//
// Solidity: function validateSelfClient(bytes clientStateBytes) view returns(bool)
func (_Ibcclient *IbcclientCallerSession) ValidateSelfClient(clientStateBytes []byte) (bool, error) {
	return _Ibcclient.Contract.ValidateSelfClient(&_Ibcclient.CallOpts, clientStateBytes)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcclient *IbcclientTransactor) CreateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "createClient", msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcclient *IbcclientSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcclient.Contract.CreateClient(&_Ibcclient.TransactOpts, msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcclient *IbcclientTransactorSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcclient.Contract.CreateClient(&_Ibcclient.TransactOpts, msg_)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcclient *IbcclientTransactor) RegisterClient(opts *bind.TransactOpts, clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "registerClient", clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcclient *IbcclientSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcclient.Contract.RegisterClient(&_Ibcclient.TransactOpts, clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcclient *IbcclientTransactorSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcclient.Contract.RegisterClient(&_Ibcclient.TransactOpts, clientType, client)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcclient *IbcclientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcclient *IbcclientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ibcclient.Contract.TransferOwnership(&_Ibcclient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcclient *IbcclientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ibcclient.Contract.TransferOwnership(&_Ibcclient.TransactOpts, newOwner)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcclient *IbcclientTransactor) UpdateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "updateClient", msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcclient *IbcclientSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcclient.Contract.UpdateClient(&_Ibcclient.TransactOpts, msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcclient *IbcclientTransactorSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcclient.Contract.UpdateClient(&_Ibcclient.TransactOpts, msg_)
}
