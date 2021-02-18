// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpletokenmodule

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

// PacketData is an auto generated low-level Go binding around an user-defined struct.
type PacketData struct {
	Sequence           uint64
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      HeightData
	TimeoutTimestamp   uint64
}

// SimpletokenmoduleABI is the input ABI used to generate the binding from.
const SimpletokenmoduleABI = "[{\"inputs\":[{\"internalType\":\"contractProvableStore\",\"name\":\"store_\",\"type\":\"address\"},{\"internalType\":\"contractIBCRoutingModule\",\"name\":\"ibcRoutingModule_\",\"type\":\"address\"},{\"internalType\":\"contractIBCChannel\",\"name\":\"ibcChannel_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeoutHeight\",\"type\":\"uint64\"}],\"name\":\"crossTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Simpletokenmodule is an auto generated Go binding around an Ethereum contract.
type Simpletokenmodule struct {
	SimpletokenmoduleCaller     // Read-only binding to the contract
	SimpletokenmoduleTransactor // Write-only binding to the contract
	SimpletokenmoduleFilterer   // Log filterer for contract events
}

// SimpletokenmoduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpletokenmoduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpletokenmoduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpletokenmoduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpletokenmoduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpletokenmoduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpletokenmoduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpletokenmoduleSession struct {
	Contract     *Simpletokenmodule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SimpletokenmoduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpletokenmoduleCallerSession struct {
	Contract *SimpletokenmoduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SimpletokenmoduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpletokenmoduleTransactorSession struct {
	Contract     *SimpletokenmoduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SimpletokenmoduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpletokenmoduleRaw struct {
	Contract *Simpletokenmodule // Generic contract binding to access the raw methods on
}

// SimpletokenmoduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpletokenmoduleCallerRaw struct {
	Contract *SimpletokenmoduleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpletokenmoduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpletokenmoduleTransactorRaw struct {
	Contract *SimpletokenmoduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpletokenmodule creates a new instance of Simpletokenmodule, bound to a specific deployed contract.
func NewSimpletokenmodule(address common.Address, backend bind.ContractBackend) (*Simpletokenmodule, error) {
	contract, err := bindSimpletokenmodule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simpletokenmodule{SimpletokenmoduleCaller: SimpletokenmoduleCaller{contract: contract}, SimpletokenmoduleTransactor: SimpletokenmoduleTransactor{contract: contract}, SimpletokenmoduleFilterer: SimpletokenmoduleFilterer{contract: contract}}, nil
}

// NewSimpletokenmoduleCaller creates a new read-only instance of Simpletokenmodule, bound to a specific deployed contract.
func NewSimpletokenmoduleCaller(address common.Address, caller bind.ContractCaller) (*SimpletokenmoduleCaller, error) {
	contract, err := bindSimpletokenmodule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpletokenmoduleCaller{contract: contract}, nil
}

// NewSimpletokenmoduleTransactor creates a new write-only instance of Simpletokenmodule, bound to a specific deployed contract.
func NewSimpletokenmoduleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpletokenmoduleTransactor, error) {
	contract, err := bindSimpletokenmodule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpletokenmoduleTransactor{contract: contract}, nil
}

// NewSimpletokenmoduleFilterer creates a new log filterer instance of Simpletokenmodule, bound to a specific deployed contract.
func NewSimpletokenmoduleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpletokenmoduleFilterer, error) {
	contract, err := bindSimpletokenmodule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpletokenmoduleFilterer{contract: contract}, nil
}

// bindSimpletokenmodule binds a generic wrapper to an already deployed contract.
func bindSimpletokenmodule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpletokenmoduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpletokenmodule *SimpletokenmoduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpletokenmodule.Contract.SimpletokenmoduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpletokenmodule *SimpletokenmoduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.SimpletokenmoduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpletokenmodule *SimpletokenmoduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.SimpletokenmoduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpletokenmodule *SimpletokenmoduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpletokenmodule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpletokenmodule *SimpletokenmoduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpletokenmodule *SimpletokenmoduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Simpletokenmodule *SimpletokenmoduleCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Simpletokenmodule.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Simpletokenmodule *SimpletokenmoduleSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Simpletokenmodule.Contract.BalanceOf(&_Simpletokenmodule.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Simpletokenmodule *SimpletokenmoduleCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Simpletokenmodule.Contract.BalanceOf(&_Simpletokenmodule.CallOpts, account)
}

// CrossTransfer is a paid mutator transaction binding the contract method 0x3e011797.
//
// Solidity: function crossTransfer(string sourcePort, string sourceChannel, address recipient, uint64 amount, uint64 timeoutHeight) returns()
func (_Simpletokenmodule *SimpletokenmoduleTransactor) CrossTransfer(opts *bind.TransactOpts, sourcePort string, sourceChannel string, recipient common.Address, amount uint64, timeoutHeight uint64) (*types.Transaction, error) {
	return _Simpletokenmodule.contract.Transact(opts, "crossTransfer", sourcePort, sourceChannel, recipient, amount, timeoutHeight)
}

// CrossTransfer is a paid mutator transaction binding the contract method 0x3e011797.
//
// Solidity: function crossTransfer(string sourcePort, string sourceChannel, address recipient, uint64 amount, uint64 timeoutHeight) returns()
func (_Simpletokenmodule *SimpletokenmoduleSession) CrossTransfer(sourcePort string, sourceChannel string, recipient common.Address, amount uint64, timeoutHeight uint64) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.CrossTransfer(&_Simpletokenmodule.TransactOpts, sourcePort, sourceChannel, recipient, amount, timeoutHeight)
}

// CrossTransfer is a paid mutator transaction binding the contract method 0x3e011797.
//
// Solidity: function crossTransfer(string sourcePort, string sourceChannel, address recipient, uint64 amount, uint64 timeoutHeight) returns()
func (_Simpletokenmodule *SimpletokenmoduleTransactorSession) CrossTransfer(sourcePort string, sourceChannel string, recipient common.Address, amount uint64, timeoutHeight uint64) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.CrossTransfer(&_Simpletokenmodule.TransactOpts, sourcePort, sourceChannel, recipient, amount, timeoutHeight)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Simpletokenmodule *SimpletokenmoduleTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Simpletokenmodule.contract.Transact(opts, "onAcknowledgementPacket", packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Simpletokenmodule *SimpletokenmoduleSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.OnAcknowledgementPacket(&_Simpletokenmodule.TransactOpts, packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Simpletokenmodule *SimpletokenmoduleTransactorSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.OnAcknowledgementPacket(&_Simpletokenmodule.TransactOpts, packet, acknowledgement)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Simpletokenmodule *SimpletokenmoduleTransactor) OnRecvPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Simpletokenmodule.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Simpletokenmodule *SimpletokenmoduleSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.OnRecvPacket(&_Simpletokenmodule.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Simpletokenmodule *SimpletokenmoduleTransactorSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.OnRecvPacket(&_Simpletokenmodule.TransactOpts, packet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Simpletokenmodule *SimpletokenmoduleTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpletokenmodule.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Simpletokenmodule *SimpletokenmoduleSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.Transfer(&_Simpletokenmodule.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Simpletokenmodule *SimpletokenmoduleTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpletokenmodule.Contract.Transfer(&_Simpletokenmodule.TransactOpts, recipient, amount)
}
