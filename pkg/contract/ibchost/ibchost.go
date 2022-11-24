// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibchost

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

// IbchostABI is the input ABI used to generate the binding from.
const IbchostABI = "[]"

// Ibchost is an auto generated Go binding around an Ethereum contract.
type Ibchost struct {
	IbchostCaller     // Read-only binding to the contract
	IbchostTransactor // Write-only binding to the contract
	IbchostFilterer   // Log filterer for contract events
}

// IbchostCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbchostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbchostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbchostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbchostSession struct {
	Contract     *Ibchost          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbchostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbchostCallerSession struct {
	Contract *IbchostCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IbchostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbchostTransactorSession struct {
	Contract     *IbchostTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IbchostRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbchostRaw struct {
	Contract *Ibchost // Generic contract binding to access the raw methods on
}

// IbchostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbchostCallerRaw struct {
	Contract *IbchostCaller // Generic read-only contract binding to access the raw methods on
}

// IbchostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbchostTransactorRaw struct {
	Contract *IbchostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbchost creates a new instance of Ibchost, bound to a specific deployed contract.
func NewIbchost(address common.Address, backend bind.ContractBackend) (*Ibchost, error) {
	contract, err := bindIbchost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibchost{IbchostCaller: IbchostCaller{contract: contract}, IbchostTransactor: IbchostTransactor{contract: contract}, IbchostFilterer: IbchostFilterer{contract: contract}}, nil
}

// NewIbchostCaller creates a new read-only instance of Ibchost, bound to a specific deployed contract.
func NewIbchostCaller(address common.Address, caller bind.ContractCaller) (*IbchostCaller, error) {
	contract, err := bindIbchost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbchostCaller{contract: contract}, nil
}

// NewIbchostTransactor creates a new write-only instance of Ibchost, bound to a specific deployed contract.
func NewIbchostTransactor(address common.Address, transactor bind.ContractTransactor) (*IbchostTransactor, error) {
	contract, err := bindIbchost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbchostTransactor{contract: contract}, nil
}

// NewIbchostFilterer creates a new log filterer instance of Ibchost, bound to a specific deployed contract.
func NewIbchostFilterer(address common.Address, filterer bind.ContractFilterer) (*IbchostFilterer, error) {
	contract, err := bindIbchost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbchostFilterer{contract: contract}, nil
}

// bindIbchost binds a generic wrapper to an already deployed contract.
func bindIbchost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbchostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchost *IbchostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchost.Contract.IbchostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchost *IbchostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.Contract.IbchostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchost *IbchostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchost.Contract.IbchostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchost *IbchostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchost *IbchostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchost *IbchostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchost.Contract.contract.Transact(opts, method, params...)
}
