// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcmockapp

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

// IIBCModuleMsgOnChanCloseConfirm is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanCloseConfirm struct {
	PortId    string
	ChannelId string
}

// IIBCModuleMsgOnChanCloseInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanCloseInit struct {
	PortId    string
	ChannelId string
}

// IIBCModuleMsgOnChanOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanOpenAck struct {
	PortId              string
	ChannelId           string
	CounterpartyVersion string
}

// IIBCModuleMsgOnChanOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanOpenConfirm struct {
	PortId    string
	ChannelId string
}

// IIBCModuleMsgOnChanOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanOpenInit struct {
	Order          uint8
	ConnectionHops []string
	PortId         string
	ChannelId      string
	Counterparty   ChannelCounterpartyData
	Version        string
}

// IIBCModuleMsgOnChanOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleMsgOnChanOpenTry struct {
	Order               uint8
	ConnectionHops      []string
	PortId              string
	ChannelId           string
	Counterparty        ChannelCounterpartyData
	CounterpartyVersion string
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

// IbcmockappMetaData contains all meta data concerning the Ibcmockapp contract.
var IbcmockappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBCHandler\",\"name\":\"ibcHandler_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MOCKAPP_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanCloseConfirm\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"onChanCloseConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanCloseInit\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"onChanCloseInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanOpenAck\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"onChanOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanOpenConfirm\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"onChanOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"order\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"connectionHops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"onChanOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"order\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"connectionHops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"}],\"internalType\":\"structIIBCModule.MsgOnChanOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"onChanOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"onTimeoutPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeoutTimestamp\",\"type\":\"uint64\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IbcmockappABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcmockappMetaData.ABI instead.
var IbcmockappABI = IbcmockappMetaData.ABI

// Ibcmockapp is an auto generated Go binding around an Ethereum contract.
type Ibcmockapp struct {
	IbcmockappCaller     // Read-only binding to the contract
	IbcmockappTransactor // Write-only binding to the contract
	IbcmockappFilterer   // Log filterer for contract events
}

// IbcmockappCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcmockappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmockappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcmockappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmockappFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcmockappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmockappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcmockappSession struct {
	Contract     *Ibcmockapp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcmockappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcmockappCallerSession struct {
	Contract *IbcmockappCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IbcmockappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcmockappTransactorSession struct {
	Contract     *IbcmockappTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbcmockappRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcmockappRaw struct {
	Contract *Ibcmockapp // Generic contract binding to access the raw methods on
}

// IbcmockappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcmockappCallerRaw struct {
	Contract *IbcmockappCaller // Generic read-only contract binding to access the raw methods on
}

// IbcmockappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcmockappTransactorRaw struct {
	Contract *IbcmockappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcmockapp creates a new instance of Ibcmockapp, bound to a specific deployed contract.
func NewIbcmockapp(address common.Address, backend bind.ContractBackend) (*Ibcmockapp, error) {
	contract, err := bindIbcmockapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcmockapp{IbcmockappCaller: IbcmockappCaller{contract: contract}, IbcmockappTransactor: IbcmockappTransactor{contract: contract}, IbcmockappFilterer: IbcmockappFilterer{contract: contract}}, nil
}

// NewIbcmockappCaller creates a new read-only instance of Ibcmockapp, bound to a specific deployed contract.
func NewIbcmockappCaller(address common.Address, caller bind.ContractCaller) (*IbcmockappCaller, error) {
	contract, err := bindIbcmockapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcmockappCaller{contract: contract}, nil
}

// NewIbcmockappTransactor creates a new write-only instance of Ibcmockapp, bound to a specific deployed contract.
func NewIbcmockappTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcmockappTransactor, error) {
	contract, err := bindIbcmockapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcmockappTransactor{contract: contract}, nil
}

// NewIbcmockappFilterer creates a new log filterer instance of Ibcmockapp, bound to a specific deployed contract.
func NewIbcmockappFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcmockappFilterer, error) {
	contract, err := bindIbcmockapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcmockappFilterer{contract: contract}, nil
}

// bindIbcmockapp binds a generic wrapper to an already deployed contract.
func bindIbcmockapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcmockappMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcmockapp *IbcmockappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcmockapp.Contract.IbcmockappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcmockapp *IbcmockappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.IbcmockappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcmockapp *IbcmockappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.IbcmockappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcmockapp *IbcmockappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcmockapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcmockapp *IbcmockappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcmockapp *IbcmockappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.contract.Transact(opts, method, params...)
}

// MOCKAPPVERSION is a free data retrieval call binding the contract method 0x7ef15d45.
//
// Solidity: function MOCKAPP_VERSION() view returns(string)
func (_Ibcmockapp *IbcmockappCaller) MOCKAPPVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "MOCKAPP_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MOCKAPPVERSION is a free data retrieval call binding the contract method 0x7ef15d45.
//
// Solidity: function MOCKAPP_VERSION() view returns(string)
func (_Ibcmockapp *IbcmockappSession) MOCKAPPVERSION() (string, error) {
	return _Ibcmockapp.Contract.MOCKAPPVERSION(&_Ibcmockapp.CallOpts)
}

// MOCKAPPVERSION is a free data retrieval call binding the contract method 0x7ef15d45.
//
// Solidity: function MOCKAPP_VERSION() view returns(string)
func (_Ibcmockapp *IbcmockappCallerSession) MOCKAPPVERSION() (string, error) {
	return _Ibcmockapp.Contract.MOCKAPPVERSION(&_Ibcmockapp.CallOpts)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ibcmockapp *IbcmockappCaller) IbcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "ibcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ibcmockapp *IbcmockappSession) IbcAddress() (common.Address, error) {
	return _Ibcmockapp.Contract.IbcAddress(&_Ibcmockapp.CallOpts)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ibcmockapp *IbcmockappCallerSession) IbcAddress() (common.Address, error) {
	return _Ibcmockapp.Contract.IbcAddress(&_Ibcmockapp.CallOpts)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet PacketData, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onAcknowledgementPacket", packet, acknowledgement, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ibcmockapp *IbcmockappSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnAcknowledgementPacket(&_Ibcmockapp.TransactOpts, packet, acknowledgement, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnAcknowledgementPacket(&_Ibcmockapp.TransactOpts, packet, acknowledgement, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanCloseConfirm", arg0)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappSession) OnChanCloseConfirm(arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanCloseConfirm(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanCloseConfirm(arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanCloseConfirm(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnChanCloseInit(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanCloseInit", arg0)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) ) returns()
func (_Ibcmockapp *IbcmockappSession) OnChanCloseInit(arg0 IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanCloseInit(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanCloseInit(arg0 IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanCloseInit(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnChanOpenAck(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanOpenAck", arg0)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) ) returns()
func (_Ibcmockapp *IbcmockappSession) OnChanOpenAck(arg0 IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenAck(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanOpenAck(arg0 IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenAck(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanOpenConfirm", arg0)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappSession) OnChanOpenConfirm(arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenConfirm(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanOpenConfirm(arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenConfirm(&_Ibcmockapp.TransactOpts, arg0)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappTransactor) OnChanOpenInit(opts *bind.TransactOpts, msg_ IIBCModuleMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanOpenInit", msg_)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappSession) OnChanOpenInit(msg_ IIBCModuleMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenInit(&_Ibcmockapp.TransactOpts, msg_)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanOpenInit(msg_ IIBCModuleMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenInit(&_Ibcmockapp.TransactOpts, msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappTransactor) OnChanOpenTry(opts *bind.TransactOpts, msg_ IIBCModuleMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onChanOpenTry", msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappSession) OnChanOpenTry(msg_ IIBCModuleMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenTry(&_Ibcmockapp.TransactOpts, msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(string)
func (_Ibcmockapp *IbcmockappTransactorSession) OnChanOpenTry(msg_ IIBCModuleMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnChanOpenTry(&_Ibcmockapp.TransactOpts, msg_)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappTransactor) OnRecvPacket(opts *bind.TransactOpts, packet PacketData, arg1 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onRecvPacket", packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappSession) OnRecvPacket(packet PacketData, arg1 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnRecvPacket(&_Ibcmockapp.TransactOpts, packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappTransactorSession) OnRecvPacket(packet PacketData, arg1 common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnRecvPacket(&_Ibcmockapp.TransactOpts, packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address relayer) returns()
func (_Ibcmockapp *IbcmockappTransactor) OnTimeoutPacket(opts *bind.TransactOpts, arg0 PacketData, relayer common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "onTimeoutPacket", arg0, relayer)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address relayer) returns()
func (_Ibcmockapp *IbcmockappSession) OnTimeoutPacket(arg0 PacketData, relayer common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnTimeoutPacket(&_Ibcmockapp.TransactOpts, arg0, relayer)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address relayer) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) OnTimeoutPacket(arg0 PacketData, relayer common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.OnTimeoutPacket(&_Ibcmockapp.TransactOpts, arg0, relayer)
}

// SendPacket is a paid mutator transaction binding the contract method 0x98e195f9.
//
// Solidity: function sendPacket(string message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns()
func (_Ibcmockapp *IbcmockappTransactor) SendPacket(opts *bind.TransactOpts, message string, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "sendPacket", message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0x98e195f9.
//
// Solidity: function sendPacket(string message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns()
func (_Ibcmockapp *IbcmockappSession) SendPacket(message string, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.SendPacket(&_Ibcmockapp.TransactOpts, message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0x98e195f9.
//
// Solidity: function sendPacket(string message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) SendPacket(message string, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.SendPacket(&_Ibcmockapp.TransactOpts, message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}
