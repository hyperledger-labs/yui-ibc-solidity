// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcroutingmodule

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

// IBCRoutingModuleModule is an auto generated low-level Go binding around an user-defined struct.
type IBCRoutingModuleModule struct {
	Callbacks common.Address
	Exists    bool
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

// IbcroutingmoduleABI is the input ABI used to generate the binding from.
const IbcroutingmoduleABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBCChannel\",\"name\":\"ibcchannel_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"lookupModule\",\"outputs\":[{\"components\":[{\"internalType\":\"contractCallbacksI\",\"name\":\"callbacks\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIBCRoutingModule.Module\",\"name\":\"module\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketAcknowledgement\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"acknowledgePacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"handlePacketRecvWithoutVerification\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ibcroutingmodule is an auto generated Go binding around an Ethereum contract.
type Ibcroutingmodule struct {
	IbcroutingmoduleCaller     // Read-only binding to the contract
	IbcroutingmoduleTransactor // Write-only binding to the contract
	IbcroutingmoduleFilterer   // Log filterer for contract events
}

// IbcroutingmoduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcroutingmoduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcroutingmoduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcroutingmoduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcroutingmoduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcroutingmoduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcroutingmoduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcroutingmoduleSession struct {
	Contract     *Ibcroutingmodule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcroutingmoduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcroutingmoduleCallerSession struct {
	Contract *IbcroutingmoduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IbcroutingmoduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcroutingmoduleTransactorSession struct {
	Contract     *IbcroutingmoduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IbcroutingmoduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcroutingmoduleRaw struct {
	Contract *Ibcroutingmodule // Generic contract binding to access the raw methods on
}

// IbcroutingmoduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcroutingmoduleCallerRaw struct {
	Contract *IbcroutingmoduleCaller // Generic read-only contract binding to access the raw methods on
}

// IbcroutingmoduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcroutingmoduleTransactorRaw struct {
	Contract *IbcroutingmoduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcroutingmodule creates a new instance of Ibcroutingmodule, bound to a specific deployed contract.
func NewIbcroutingmodule(address common.Address, backend bind.ContractBackend) (*Ibcroutingmodule, error) {
	contract, err := bindIbcroutingmodule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcroutingmodule{IbcroutingmoduleCaller: IbcroutingmoduleCaller{contract: contract}, IbcroutingmoduleTransactor: IbcroutingmoduleTransactor{contract: contract}, IbcroutingmoduleFilterer: IbcroutingmoduleFilterer{contract: contract}}, nil
}

// NewIbcroutingmoduleCaller creates a new read-only instance of Ibcroutingmodule, bound to a specific deployed contract.
func NewIbcroutingmoduleCaller(address common.Address, caller bind.ContractCaller) (*IbcroutingmoduleCaller, error) {
	contract, err := bindIbcroutingmodule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcroutingmoduleCaller{contract: contract}, nil
}

// NewIbcroutingmoduleTransactor creates a new write-only instance of Ibcroutingmodule, bound to a specific deployed contract.
func NewIbcroutingmoduleTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcroutingmoduleTransactor, error) {
	contract, err := bindIbcroutingmodule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcroutingmoduleTransactor{contract: contract}, nil
}

// NewIbcroutingmoduleFilterer creates a new log filterer instance of Ibcroutingmodule, bound to a specific deployed contract.
func NewIbcroutingmoduleFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcroutingmoduleFilterer, error) {
	contract, err := bindIbcroutingmodule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcroutingmoduleFilterer{contract: contract}, nil
}

// bindIbcroutingmodule binds a generic wrapper to an already deployed contract.
func bindIbcroutingmodule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcroutingmoduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcroutingmodule *IbcroutingmoduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcroutingmodule.Contract.IbcroutingmoduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcroutingmodule *IbcroutingmoduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.IbcroutingmoduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcroutingmodule *IbcroutingmoduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.IbcroutingmoduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcroutingmodule *IbcroutingmoduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcroutingmodule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcroutingmodule *IbcroutingmoduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcroutingmodule *IbcroutingmoduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.contract.Transact(opts, method, params...)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibcroutingmodule *IbcroutingmoduleCaller) LookupModule(opts *bind.CallOpts, portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	var out []interface{}
	err := _Ibcroutingmodule.contract.Call(opts, &out, "lookupModule", portId)

	outstruct := new(struct {
		Module IBCRoutingModuleModule
		Found  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Module = out[0].(IBCRoutingModuleModule)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibcroutingmodule *IbcroutingmoduleSession) LookupModule(portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	return _Ibcroutingmodule.Contract.LookupModule(&_Ibcroutingmodule.CallOpts, portId)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibcroutingmodule *IbcroutingmoduleCallerSession) LookupModule(portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	return _Ibcroutingmodule.Contract.LookupModule(&_Ibcroutingmodule.CallOpts, portId)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcroutingmodule *IbcroutingmoduleTransactor) AcknowledgePacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcroutingmodule.contract.Transact(opts, "acknowledgePacket", msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcroutingmodule *IbcroutingmoduleSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.AcknowledgePacket(&_Ibcroutingmodule.TransactOpts, msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcroutingmodule *IbcroutingmoduleTransactorSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.AcknowledgePacket(&_Ibcroutingmodule.TransactOpts, msg_)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcroutingmodule *IbcroutingmoduleTransactor) BindPort(opts *bind.TransactOpts, portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcroutingmodule.contract.Transact(opts, "bindPort", portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcroutingmodule *IbcroutingmoduleSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.BindPort(&_Ibcroutingmodule.TransactOpts, portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcroutingmodule *IbcroutingmoduleTransactorSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.BindPort(&_Ibcroutingmodule.TransactOpts, portId, moduleAddress)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleTransactor) HandlePacketRecvWithoutVerification(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.contract.Transact(opts, "handlePacketRecvWithoutVerification", msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.HandlePacketRecvWithoutVerification(&_Ibcroutingmodule.TransactOpts, msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleTransactorSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.HandlePacketRecvWithoutVerification(&_Ibcroutingmodule.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleTransactor) RecvPacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.contract.Transact(opts, "recvPacket", msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.RecvPacket(&_Ibcroutingmodule.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcroutingmodule *IbcroutingmoduleTransactorSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcroutingmodule.Contract.RecvPacket(&_Ibcroutingmodule.TransactOpts, msg_)
}
