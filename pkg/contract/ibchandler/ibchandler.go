// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibchandler

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

// IBCHandlerModule is an auto generated low-level Go binding around an user-defined struct.
type IBCHandlerModule struct {
	Callbacks common.Address
	Exists    bool
}

// IBCMsgsMsgPacketAcknowledgement is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgPacketAcknowledgement struct {
	Packet          PacketData
	Acknowledgement []byte
	Proof           []byte
	ProofHeight     uint64
}

// IBCMsgsMsgPacketRecv is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgPacketRecv struct {
	Packet      PacketData
	Proof       []byte
	ProofHeight uint64
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

// IbchandlerABI is the input ABI used to generate the binding from.
const IbchandlerABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBCChannel\",\"name\":\"ibcchannel_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"handlePacketRecv\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketAcknowledgement\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"handlePacketAcknowledgement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"handlePacketRecvWithoutVerification\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"lookupModule\",\"outputs\":[{\"components\":[{\"internalType\":\"contractCallbacksI\",\"name\":\"callbacks\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIBCHandler.Module\",\"name\":\"module\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Ibchandler is an auto generated Go binding around an Ethereum contract.
type Ibchandler struct {
	IbchandlerCaller     // Read-only binding to the contract
	IbchandlerTransactor // Write-only binding to the contract
	IbchandlerFilterer   // Log filterer for contract events
}

// IbchandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbchandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbchandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbchandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbchandlerSession struct {
	Contract     *Ibchandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbchandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbchandlerCallerSession struct {
	Contract *IbchandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IbchandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbchandlerTransactorSession struct {
	Contract     *IbchandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbchandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbchandlerRaw struct {
	Contract *Ibchandler // Generic contract binding to access the raw methods on
}

// IbchandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbchandlerCallerRaw struct {
	Contract *IbchandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IbchandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbchandlerTransactorRaw struct {
	Contract *IbchandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbchandler creates a new instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandler(address common.Address, backend bind.ContractBackend) (*Ibchandler, error) {
	contract, err := bindIbchandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibchandler{IbchandlerCaller: IbchandlerCaller{contract: contract}, IbchandlerTransactor: IbchandlerTransactor{contract: contract}, IbchandlerFilterer: IbchandlerFilterer{contract: contract}}, nil
}

// NewIbchandlerCaller creates a new read-only instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerCaller(address common.Address, caller bind.ContractCaller) (*IbchandlerCaller, error) {
	contract, err := bindIbchandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbchandlerCaller{contract: contract}, nil
}

// NewIbchandlerTransactor creates a new write-only instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IbchandlerTransactor, error) {
	contract, err := bindIbchandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbchandlerTransactor{contract: contract}, nil
}

// NewIbchandlerFilterer creates a new log filterer instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IbchandlerFilterer, error) {
	contract, err := bindIbchandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbchandlerFilterer{contract: contract}, nil
}

// bindIbchandler binds a generic wrapper to an already deployed contract.
func bindIbchandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbchandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchandler *IbchandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchandler.Contract.IbchandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchandler *IbchandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchandler.Contract.IbchandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchandler *IbchandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchandler.Contract.IbchandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchandler *IbchandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchandler *IbchandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchandler *IbchandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchandler.Contract.contract.Transact(opts, method, params...)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibchandler *IbchandlerCaller) LookupModule(opts *bind.CallOpts, portId string) (struct {
	Module IBCHandlerModule
	Found  bool
}, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "lookupModule", portId)

	outstruct := new(struct {
		Module IBCHandlerModule
		Found  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Module = out[0].(IBCHandlerModule)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibchandler *IbchandlerSession) LookupModule(portId string) (struct {
	Module IBCHandlerModule
	Found  bool
}, error) {
	return _Ibchandler.Contract.LookupModule(&_Ibchandler.CallOpts, portId)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibchandler *IbchandlerCallerSession) LookupModule(portId string) (struct {
	Module IBCHandlerModule
	Found  bool
}, error) {
	return _Ibchandler.Contract.LookupModule(&_Ibchandler.CallOpts, portId)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerTransactor) BindPort(opts *bind.TransactOpts, portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "bindPort", portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.BindPort(&_Ibchandler.TransactOpts, portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerTransactorSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.BindPort(&_Ibchandler.TransactOpts, portId, moduleAddress)
}

// HandlePacketAcknowledgement is a paid mutator transaction binding the contract method 0x7b566424.
//
// Solidity: function handlePacketAcknowledgement(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) HandlePacketAcknowledgement(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "handlePacketAcknowledgement", msg_)
}

// HandlePacketAcknowledgement is a paid mutator transaction binding the contract method 0x7b566424.
//
// Solidity: function handlePacketAcknowledgement(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) HandlePacketAcknowledgement(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketAcknowledgement(&_Ibchandler.TransactOpts, msg_)
}

// HandlePacketAcknowledgement is a paid mutator transaction binding the contract method 0x7b566424.
//
// Solidity: function handlePacketAcknowledgement(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) HandlePacketAcknowledgement(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketAcknowledgement(&_Ibchandler.TransactOpts, msg_)
}

// HandlePacketRecv is a paid mutator transaction binding the contract method 0xad1ad8fe.
//
// Solidity: function handlePacketRecv(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerTransactor) HandlePacketRecv(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "handlePacketRecv", msg_)
}

// HandlePacketRecv is a paid mutator transaction binding the contract method 0xad1ad8fe.
//
// Solidity: function handlePacketRecv(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerSession) HandlePacketRecv(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketRecv(&_Ibchandler.TransactOpts, msg_)
}

// HandlePacketRecv is a paid mutator transaction binding the contract method 0xad1ad8fe.
//
// Solidity: function handlePacketRecv(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerTransactorSession) HandlePacketRecv(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketRecv(&_Ibchandler.TransactOpts, msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerTransactor) HandlePacketRecvWithoutVerification(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "handlePacketRecvWithoutVerification", msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketRecvWithoutVerification(&_Ibchandler.TransactOpts, msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibchandler *IbchandlerTransactorSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.HandlePacketRecvWithoutVerification(&_Ibchandler.TransactOpts, msg_)
}
