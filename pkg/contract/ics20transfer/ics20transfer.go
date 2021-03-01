// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20transfer

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

// ChannelCounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type ChannelCounterpartyData struct {
	PortId    string
	ChannelId string
}

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

// Ics20transferABI is the input ABI used to generate the binding from.
const Ics20transferABI = "[{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host_\",\"type\":\"address\"},{\"internalType\":\"contractIBCHandler\",\"name\":\"ibcHandler_\",\"type\":\"address\"},{\"internalType\":\"contractIICS20Vouchers\",\"name\":\"bank_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timeoutHeight\",\"type\":\"uint64\"}],\"name\":\"sendTransferWithTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timeoutHeight\",\"type\":\"uint64\"}],\"name\":\"sendTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"connectionHops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"onChanOpenInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"connectionHops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"}],\"name\":\"onChanOpenTry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"}],\"name\":\"onChanOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"onChanOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ics20transfer is an auto generated Go binding around an Ethereum contract.
type Ics20transfer struct {
	Ics20transferCaller     // Read-only binding to the contract
	Ics20transferTransactor // Write-only binding to the contract
	Ics20transferFilterer   // Log filterer for contract events
}

// Ics20transferCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ics20transferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20transferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ics20transferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20transferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ics20transferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20transferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ics20transferSession struct {
	Contract     *Ics20transfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics20transferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ics20transferCallerSession struct {
	Contract *Ics20transferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Ics20transferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ics20transferTransactorSession struct {
	Contract     *Ics20transferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Ics20transferRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ics20transferRaw struct {
	Contract *Ics20transfer // Generic contract binding to access the raw methods on
}

// Ics20transferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ics20transferCallerRaw struct {
	Contract *Ics20transferCaller // Generic read-only contract binding to access the raw methods on
}

// Ics20transferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ics20transferTransactorRaw struct {
	Contract *Ics20transferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIcs20transfer creates a new instance of Ics20transfer, bound to a specific deployed contract.
func NewIcs20transfer(address common.Address, backend bind.ContractBackend) (*Ics20transfer, error) {
	contract, err := bindIcs20transfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ics20transfer{Ics20transferCaller: Ics20transferCaller{contract: contract}, Ics20transferTransactor: Ics20transferTransactor{contract: contract}, Ics20transferFilterer: Ics20transferFilterer{contract: contract}}, nil
}

// NewIcs20transferCaller creates a new read-only instance of Ics20transfer, bound to a specific deployed contract.
func NewIcs20transferCaller(address common.Address, caller bind.ContractCaller) (*Ics20transferCaller, error) {
	contract, err := bindIcs20transfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ics20transferCaller{contract: contract}, nil
}

// NewIcs20transferTransactor creates a new write-only instance of Ics20transfer, bound to a specific deployed contract.
func NewIcs20transferTransactor(address common.Address, transactor bind.ContractTransactor) (*Ics20transferTransactor, error) {
	contract, err := bindIcs20transfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ics20transferTransactor{contract: contract}, nil
}

// NewIcs20transferFilterer creates a new log filterer instance of Ics20transfer, bound to a specific deployed contract.
func NewIcs20transferFilterer(address common.Address, filterer bind.ContractFilterer) (*Ics20transferFilterer, error) {
	contract, err := bindIcs20transfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ics20transferFilterer{contract: contract}, nil
}

// bindIcs20transfer binds a generic wrapper to an already deployed contract.
func bindIcs20transfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ics20transferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics20transfer *Ics20transferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics20transfer.Contract.Ics20transferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics20transfer *Ics20transferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Ics20transferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics20transfer *Ics20transferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Ics20transferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics20transfer *Ics20transferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics20transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics20transfer *Ics20transferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics20transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics20transfer *Ics20transferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics20transfer.Contract.contract.Transact(opts, method, params...)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Ics20transfer *Ics20transferTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onAcknowledgementPacket", packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Ics20transfer *Ics20transferSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnAcknowledgementPacket(&_Ics20transfer.TransactOpts, packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnAcknowledgementPacket(&_Ics20transfer.TransactOpts, packet, acknowledgement)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenAck(opts *bind.TransactOpts, portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenAck", portId, channelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenAck(portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenAck(&_Ics20transfer.TransactOpts, portId, channelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenAck(portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenAck(&_Ics20transfer.TransactOpts, portId, channelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenConfirm", portId, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenConfirm(&_Ics20transfer.TransactOpts, portId, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenConfirm(&_Ics20transfer.TransactOpts, portId, channelId)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenInit(opts *bind.TransactOpts, arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenInit", arg0, connectionHops, portId, channelId, counterparty, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenInit(arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenInit(&_Ics20transfer.TransactOpts, arg0, connectionHops, portId, channelId, counterparty, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenInit(arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenInit(&_Ics20transfer.TransactOpts, arg0, connectionHops, portId, channelId, counterparty, version)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenTry(opts *bind.TransactOpts, arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenTry", arg0, connectionHops, portId, channelId, counterparty, version, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenTry(arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenTry(&_Ics20transfer.TransactOpts, arg0, connectionHops, portId, channelId, counterparty, version, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] connectionHops, string portId, string channelId, (string,string) counterparty, string version, string counterpartyVersion) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenTry(arg0 uint8, connectionHops []string, portId string, channelId string, counterparty ChannelCounterpartyData, version string, counterpartyVersion string) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenTry(&_Ics20transfer.TransactOpts, arg0, connectionHops, portId, channelId, counterparty, version, counterpartyVersion)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferTransactor) OnRecvPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnRecvPacket(&_Ics20transfer.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferTransactorSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnRecvPacket(&_Ics20transfer.TransactOpts, packet)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xf1614af7.
//
// Solidity: function sendTransfer(string denom, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferTransactor) SendTransfer(opts *bind.TransactOpts, denom string, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "sendTransfer", denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xf1614af7.
//
// Solidity: function sendTransfer(string denom, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferSession) SendTransfer(denom string, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransfer(&_Ics20transfer.TransactOpts, denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xf1614af7.
//
// Solidity: function sendTransfer(string denom, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferTransactorSession) SendTransfer(denom string, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransfer(&_Ics20transfer.TransactOpts, denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransferWithTokenContract is a paid mutator transaction binding the contract method 0x3adcf18d.
//
// Solidity: function sendTransferWithTokenContract(address tokenContract, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferTransactor) SendTransferWithTokenContract(opts *bind.TransactOpts, tokenContract common.Address, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "sendTransferWithTokenContract", tokenContract, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransferWithTokenContract is a paid mutator transaction binding the contract method 0x3adcf18d.
//
// Solidity: function sendTransferWithTokenContract(address tokenContract, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferSession) SendTransferWithTokenContract(tokenContract common.Address, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransferWithTokenContract(&_Ics20transfer.TransactOpts, tokenContract, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransferWithTokenContract is a paid mutator transaction binding the contract method 0x3adcf18d.
//
// Solidity: function sendTransferWithTokenContract(address tokenContract, uint256 amount, address receiver, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Ics20transfer *Ics20transferTransactorSession) SendTransferWithTokenContract(tokenContract common.Address, amount *big.Int, receiver common.Address, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransferWithTokenContract(&_Ics20transfer.TransactOpts, tokenContract, amount, receiver, sourcePort, sourceChannel, timeoutHeight)
}
