// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcchannel

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

// ChannelData is an auto generated low-level Go binding around an user-defined struct.
type ChannelData struct {
	State          uint8
	Ordering       uint8
	Counterparty   ChannelCounterpartyData
	ConnectionHops []string
	Version        string
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IBCChannelMsgChannelOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IBCChannelMsgChannelOpenAck struct {
	PortId                string
	ChannelId             string
	CounterpartyVersion   string
	CounterpartyChannelId string
	ProofTry              []byte
	ProofHeight           uint64
}

// IBCChannelMsgChannelOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IBCChannelMsgChannelOpenConfirm struct {
	PortId      string
	ChannelId   string
	ProofAck    []byte
	ProofHeight uint64
}

// IBCChannelMsgChannelOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IBCChannelMsgChannelOpenInit struct {
	PortId    string
	ChannelId string
	Channel   ChannelData
}

// IBCChannelMsgChannelOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCChannelMsgChannelOpenTry struct {
	PortId              string
	ChannelId           string
	Channel             ChannelData
	CounterpartyVersion string
	ProofInit           []byte
	ProofHeight         uint64
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

// IbcchannelABI is the input ABI used to generate the binding from.
const IbcchannelABI = "[{\"inputs\":[{\"internalType\":\"contractProvableStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBCClient\",\"name\":\"ibcclient_\",\"type\":\"address\"},{\"internalType\":\"contractIBCConnection\",\"name\":\"ibcconnection_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"internalType\":\"structIBCChannel.MsgChannelOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCChannel.MsgChannelOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCChannel.MsgChannelOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCChannel.MsgChannelOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"name\":\"recvPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ibcchannel is an auto generated Go binding around an Ethereum contract.
type Ibcchannel struct {
	IbcchannelCaller     // Read-only binding to the contract
	IbcchannelTransactor // Write-only binding to the contract
	IbcchannelFilterer   // Log filterer for contract events
}

// IbcchannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcchannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcchannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcchannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcchannelSession struct {
	Contract     *Ibcchannel       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcchannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcchannelCallerSession struct {
	Contract *IbcchannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IbcchannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcchannelTransactorSession struct {
	Contract     *IbcchannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbcchannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcchannelRaw struct {
	Contract *Ibcchannel // Generic contract binding to access the raw methods on
}

// IbcchannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcchannelCallerRaw struct {
	Contract *IbcchannelCaller // Generic read-only contract binding to access the raw methods on
}

// IbcchannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcchannelTransactorRaw struct {
	Contract *IbcchannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcchannel creates a new instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannel(address common.Address, backend bind.ContractBackend) (*Ibcchannel, error) {
	contract, err := bindIbcchannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcchannel{IbcchannelCaller: IbcchannelCaller{contract: contract}, IbcchannelTransactor: IbcchannelTransactor{contract: contract}, IbcchannelFilterer: IbcchannelFilterer{contract: contract}}, nil
}

// NewIbcchannelCaller creates a new read-only instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelCaller(address common.Address, caller bind.ContractCaller) (*IbcchannelCaller, error) {
	contract, err := bindIbcchannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcchannelCaller{contract: contract}, nil
}

// NewIbcchannelTransactor creates a new write-only instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcchannelTransactor, error) {
	contract, err := bindIbcchannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcchannelTransactor{contract: contract}, nil
}

// NewIbcchannelFilterer creates a new log filterer instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcchannelFilterer, error) {
	contract, err := bindIbcchannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcchannelFilterer{contract: contract}, nil
}

// bindIbcchannel binds a generic wrapper to an already deployed contract.
func bindIbcchannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcchannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcchannel *IbcchannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcchannel.Contract.IbcchannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcchannel *IbcchannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcchannel.Contract.IbcchannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcchannel *IbcchannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcchannel.Contract.IbcchannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcchannel *IbcchannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcchannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcchannel *IbcchannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcchannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcchannel *IbcchannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcchannel.Contract.contract.Transact(opts, method, params...)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelTransactor) ChannelOpenAck(opts *bind.TransactOpts, msg_ IBCChannelMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "channelOpenAck", msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelSession) ChannelOpenAck(msg_ IBCChannelMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenAck(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelTransactorSession) ChannelOpenAck(msg_ IBCChannelMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenAck(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, msg_ IBCChannelMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "channelOpenConfirm", msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelSession) ChannelOpenConfirm(msg_ IBCChannelMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenConfirm(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcchannel *IbcchannelTransactorSession) ChannelOpenConfirm(msg_ IBCChannelMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenConfirm(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactor) ChannelOpenInit(opts *bind.TransactOpts, msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "channelOpenInit", msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelSession) ChannelOpenInit(msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenInit(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactorSession) ChannelOpenInit(msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenInit(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactor) ChannelOpenTry(opts *bind.TransactOpts, msg_ IBCChannelMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "channelOpenTry", msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcchannel *IbcchannelSession) ChannelOpenTry(msg_ IBCChannelMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenTry(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactorSession) ChannelOpenTry(msg_ IBCChannelMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenTry(&_Ibcchannel.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x7524937e.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes proof, uint64 proofHeight) returns()
func (_Ibcchannel *IbcchannelTransactor) RecvPacket(opts *bind.TransactOpts, packet PacketData, proof []byte, proofHeight uint64) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "recvPacket", packet, proof, proofHeight)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x7524937e.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes proof, uint64 proofHeight) returns()
func (_Ibcchannel *IbcchannelSession) RecvPacket(packet PacketData, proof []byte, proofHeight uint64) (*types.Transaction, error) {
	return _Ibcchannel.Contract.RecvPacket(&_Ibcchannel.TransactOpts, packet, proof, proofHeight)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x7524937e.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes proof, uint64 proofHeight) returns()
func (_Ibcchannel *IbcchannelTransactorSession) RecvPacket(packet PacketData, proof []byte, proofHeight uint64) (*types.Transaction, error) {
	return _Ibcchannel.Contract.RecvPacket(&_Ibcchannel.TransactOpts, packet, proof, proofHeight)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcchannel *IbcchannelTransactor) SendPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "sendPacket", packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcchannel *IbcchannelSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibcchannel.Contract.SendPacket(&_Ibcchannel.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcchannel *IbcchannelTransactorSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibcchannel.Contract.SendPacket(&_Ibcchannel.TransactOpts, packet)
}
