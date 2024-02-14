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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcHandler_\",\"type\":\"address\",\"internalType\":\"contractIIBCHandler\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MOCKAPP_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowCloseChannel\",\"inputs\":[{\"name\":\"allow\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeChannelAllowed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ibcAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanCloseConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanCloseInit\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenAck\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenInit\",\"components\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenTry\",\"components\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeAcknowledgement\",\"inputs\":[{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
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

// CloseChannelAllowed is a free data retrieval call binding the contract method 0x5a42d326.
//
// Solidity: function closeChannelAllowed() view returns(bool)
func (_Ibcmockapp *IbcmockappCaller) CloseChannelAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "closeChannelAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CloseChannelAllowed is a free data retrieval call binding the contract method 0x5a42d326.
//
// Solidity: function closeChannelAllowed() view returns(bool)
func (_Ibcmockapp *IbcmockappSession) CloseChannelAllowed() (bool, error) {
	return _Ibcmockapp.Contract.CloseChannelAllowed(&_Ibcmockapp.CallOpts)
}

// CloseChannelAllowed is a free data retrieval call binding the contract method 0x5a42d326.
//
// Solidity: function closeChannelAllowed() view returns(bool)
func (_Ibcmockapp *IbcmockappCallerSession) CloseChannelAllowed() (bool, error) {
	return _Ibcmockapp.Contract.CloseChannelAllowed(&_Ibcmockapp.CallOpts)
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

// OnRecvPacket is a free data retrieval call binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) view returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappCaller) OnRecvPacket(opts *bind.CallOpts, packet PacketData, arg1 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "onRecvPacket", packet, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// OnRecvPacket is a free data retrieval call binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) view returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappSession) OnRecvPacket(packet PacketData, arg1 common.Address) ([]byte, error) {
	return _Ibcmockapp.Contract.OnRecvPacket(&_Ibcmockapp.CallOpts, packet, arg1)
}

// OnRecvPacket is a free data retrieval call binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) view returns(bytes acknowledgement)
func (_Ibcmockapp *IbcmockappCallerSession) OnRecvPacket(packet PacketData, arg1 common.Address) ([]byte, error) {
	return _Ibcmockapp.Contract.OnRecvPacket(&_Ibcmockapp.CallOpts, packet, arg1)
}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address ) view returns()
func (_Ibcmockapp *IbcmockappCaller) OnTimeoutPacket(opts *bind.CallOpts, arg0 PacketData, arg1 common.Address) error {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "onTimeoutPacket", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address ) view returns()
func (_Ibcmockapp *IbcmockappSession) OnTimeoutPacket(arg0 PacketData, arg1 common.Address) error {
	return _Ibcmockapp.Contract.OnTimeoutPacket(&_Ibcmockapp.CallOpts, arg0, arg1)
}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) , address ) view returns()
func (_Ibcmockapp *IbcmockappCallerSession) OnTimeoutPacket(arg0 PacketData, arg1 common.Address) error {
	return _Ibcmockapp.Contract.OnTimeoutPacket(&_Ibcmockapp.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcmockapp *IbcmockappCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ibcmockapp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcmockapp *IbcmockappSession) Owner() (common.Address, error) {
	return _Ibcmockapp.Contract.Owner(&_Ibcmockapp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ibcmockapp *IbcmockappCallerSession) Owner() (common.Address, error) {
	return _Ibcmockapp.Contract.Owner(&_Ibcmockapp.CallOpts)
}

// AllowCloseChannel is a paid mutator transaction binding the contract method 0x9cc3ca73.
//
// Solidity: function allowCloseChannel(bool allow) returns()
func (_Ibcmockapp *IbcmockappTransactor) AllowCloseChannel(opts *bind.TransactOpts, allow bool) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "allowCloseChannel", allow)
}

// AllowCloseChannel is a paid mutator transaction binding the contract method 0x9cc3ca73.
//
// Solidity: function allowCloseChannel(bool allow) returns()
func (_Ibcmockapp *IbcmockappSession) AllowCloseChannel(allow bool) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.AllowCloseChannel(&_Ibcmockapp.TransactOpts, allow)
}

// AllowCloseChannel is a paid mutator transaction binding the contract method 0x9cc3ca73.
//
// Solidity: function allowCloseChannel(bool allow) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) AllowCloseChannel(allow bool) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.AllowCloseChannel(&_Ibcmockapp.TransactOpts, allow)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ibcmockapp *IbcmockappTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ibcmockapp *IbcmockappSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ibcmockapp.Contract.RenounceOwnership(&_Ibcmockapp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ibcmockapp *IbcmockappTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ibcmockapp.Contract.RenounceOwnership(&_Ibcmockapp.TransactOpts)
}

// SendPacket is a paid mutator transaction binding the contract method 0xd9df74fd.
//
// Solidity: function sendPacket(bytes message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns(uint64)
func (_Ibcmockapp *IbcmockappTransactor) SendPacket(opts *bind.TransactOpts, message []byte, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "sendPacket", message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xd9df74fd.
//
// Solidity: function sendPacket(bytes message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns(uint64)
func (_Ibcmockapp *IbcmockappSession) SendPacket(message []byte, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.SendPacket(&_Ibcmockapp.TransactOpts, message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xd9df74fd.
//
// Solidity: function sendPacket(bytes message, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp) returns(uint64)
func (_Ibcmockapp *IbcmockappTransactorSession) SendPacket(message []byte, sourcePort string, sourceChannel string, timeoutHeight HeightData, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.SendPacket(&_Ibcmockapp.TransactOpts, message, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcmockapp *IbcmockappTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcmockapp *IbcmockappSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.TransferOwnership(&_Ibcmockapp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.TransferOwnership(&_Ibcmockapp.TransactOpts, newOwner)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0x62892d7a.
//
// Solidity: function writeAcknowledgement(string destinationPort, string destinationChannel, uint64 sequence) returns()
func (_Ibcmockapp *IbcmockappTransactor) WriteAcknowledgement(opts *bind.TransactOpts, destinationPort string, destinationChannel string, sequence uint64) (*types.Transaction, error) {
	return _Ibcmockapp.contract.Transact(opts, "writeAcknowledgement", destinationPort, destinationChannel, sequence)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0x62892d7a.
//
// Solidity: function writeAcknowledgement(string destinationPort, string destinationChannel, uint64 sequence) returns()
func (_Ibcmockapp *IbcmockappSession) WriteAcknowledgement(destinationPort string, destinationChannel string, sequence uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.WriteAcknowledgement(&_Ibcmockapp.TransactOpts, destinationPort, destinationChannel, sequence)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0x62892d7a.
//
// Solidity: function writeAcknowledgement(string destinationPort, string destinationChannel, uint64 sequence) returns()
func (_Ibcmockapp *IbcmockappTransactorSession) WriteAcknowledgement(destinationPort string, destinationChannel string, sequence uint64) (*types.Transaction, error) {
	return _Ibcmockapp.Contract.WriteAcknowledgement(&_Ibcmockapp.TransactOpts, destinationPort, destinationChannel, sequence)
}

// IbcmockappOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ibcmockapp contract.
type IbcmockappOwnershipTransferredIterator struct {
	Event *IbcmockappOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbcmockappOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcmockappOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbcmockappOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbcmockappOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcmockappOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcmockappOwnershipTransferred represents a OwnershipTransferred event raised by the Ibcmockapp contract.
type IbcmockappOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ibcmockapp *IbcmockappFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IbcmockappOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ibcmockapp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IbcmockappOwnershipTransferredIterator{contract: _Ibcmockapp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ibcmockapp *IbcmockappFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IbcmockappOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ibcmockapp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcmockappOwnershipTransferred)
				if err := _Ibcmockapp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ibcmockapp *IbcmockappFilterer) ParseOwnershipTransferred(log types.Log) (*IbcmockappOwnershipTransferred, error) {
	event := new(IbcmockappOwnershipTransferred)
	if err := _Ibcmockapp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
