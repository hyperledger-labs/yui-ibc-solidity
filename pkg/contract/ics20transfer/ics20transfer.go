// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20transfer

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

// ICS20LibTimeout is an auto generated low-level Go binding around an user-defined struct.
type ICS20LibTimeout struct {
	Height         HeightData
	TimestampNanos uint64
}

// IIBCModuleInitializerMsgOnChanOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleInitializerMsgOnChanOpenInit struct {
	Order          uint8
	ConnectionHops []string
	PortId         string
	ChannelId      string
	Counterparty   ChannelCounterpartyData
	Version        string
}

// IIBCModuleInitializerMsgOnChanOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleInitializerMsgOnChanOpenTry struct {
	Order               uint8
	ConnectionHops      []string
	PortId              string
	ChannelId           string
	Counterparty        ChannelCounterpartyData
	CounterpartyVersion string
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

// Packet is an auto generated low-level Go binding around an user-defined struct.
type Packet struct {
	Sequence           uint64
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      HeightData
	TimeoutTimestamp   uint64
}

// Ics20transferMetaData contains all meta data concerning the Ics20transfer contract.
var Ics20transferMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcHandler_\",\"type\":\"address\",\"internalType\":\"contractIIBCHandler\"},{\"name\":\"port_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ICS20_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositSendTransfer\",\"inputs\":[{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeout\",\"type\":\"tuple\",\"internalType\":\"structICS20Lib.Timeout\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timestampNanos\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encodeAddress\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getVoucherEscrow\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ibcAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanCloseConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanCloseInit\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenAck\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.MsgOnChanOpenConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModuleInitializer.MsgOnChanOpenInit\",\"components\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCModuleInitializer.MsgOnChanOpenTry\",\"components\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTransfer\",\"inputs\":[{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeout\",\"type\":\"tuple\",\"internalType\":\"structICS20Lib.Timeout\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timestampNanos\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"IBCModuleChannelCloseNotAllowed\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCModuleChannelOrderNotAllowed\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"}]},{\"type\":\"error\",\"name\":\"IBCModuleInvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20BytesSliceOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20BytesSliceOverflow\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20FailedERC20Transfer\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20FailedERC20TransferFrom\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20FailedRefund\",\"inputs\":[{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidReceiverAddress\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidSenderAddress\",\"inputs\":[{\"name\":\"sender\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidTokenContract\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20JSONClosingBraceNotFound\",\"inputs\":[{\"name\":\"position\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"}]},{\"type\":\"error\",\"name\":\"ICS20JSONInvalidEscape\",\"inputs\":[{\"name\":\"position\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"}]},{\"type\":\"error\",\"name\":\"ICS20JSONStringClosingDoubleQuoteNotFound\",\"inputs\":[{\"name\":\"position\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"}]},{\"type\":\"error\",\"name\":\"ICS20JSONStringUnclosed\",\"inputs\":[{\"name\":\"bz\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"position\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20JSONUnexpectedBytes\",\"inputs\":[{\"name\":\"position\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ICS20UnexpectedPort\",\"inputs\":[{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20UnexpectedVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// Ics20transferABI is the input ABI used to generate the binding from.
// Deprecated: Use Ics20transferMetaData.ABI instead.
var Ics20transferABI = Ics20transferMetaData.ABI

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
	parsed, err := Ics20transferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ICS20VERSION is a free data retrieval call binding the contract method 0x025183eb.
//
// Solidity: function ICS20_VERSION() view returns(string)
func (_Ics20transfer *Ics20transferCaller) ICS20VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "ICS20_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ICS20VERSION is a free data retrieval call binding the contract method 0x025183eb.
//
// Solidity: function ICS20_VERSION() view returns(string)
func (_Ics20transfer *Ics20transferSession) ICS20VERSION() (string, error) {
	return _Ics20transfer.Contract.ICS20VERSION(&_Ics20transfer.CallOpts)
}

// ICS20VERSION is a free data retrieval call binding the contract method 0x025183eb.
//
// Solidity: function ICS20_VERSION() view returns(string)
func (_Ics20transfer *Ics20transferCallerSession) ICS20VERSION() (string, error) {
	return _Ics20transfer.Contract.ICS20VERSION(&_Ics20transfer.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string denom) view returns(uint256)
func (_Ics20transfer *Ics20transferCaller) BalanceOf(opts *bind.CallOpts, account common.Address, denom string) (*big.Int, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "balanceOf", account, denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string denom) view returns(uint256)
func (_Ics20transfer *Ics20transferSession) BalanceOf(account common.Address, denom string) (*big.Int, error) {
	return _Ics20transfer.Contract.BalanceOf(&_Ics20transfer.CallOpts, account, denom)
}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string denom) view returns(uint256)
func (_Ics20transfer *Ics20transferCallerSession) BalanceOf(account common.Address, denom string) (*big.Int, error) {
	return _Ics20transfer.Contract.BalanceOf(&_Ics20transfer.CallOpts, account, denom)
}

// EncodeAddress is a free data retrieval call binding the contract method 0xd32b1bea.
//
// Solidity: function encodeAddress(address sender) pure returns(string)
func (_Ics20transfer *Ics20transferCaller) EncodeAddress(opts *bind.CallOpts, sender common.Address) (string, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "encodeAddress", sender)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncodeAddress is a free data retrieval call binding the contract method 0xd32b1bea.
//
// Solidity: function encodeAddress(address sender) pure returns(string)
func (_Ics20transfer *Ics20transferSession) EncodeAddress(sender common.Address) (string, error) {
	return _Ics20transfer.Contract.EncodeAddress(&_Ics20transfer.CallOpts, sender)
}

// EncodeAddress is a free data retrieval call binding the contract method 0xd32b1bea.
//
// Solidity: function encodeAddress(address sender) pure returns(string)
func (_Ics20transfer *Ics20transferCallerSession) EncodeAddress(sender common.Address) (string, error) {
	return _Ics20transfer.Contract.EncodeAddress(&_Ics20transfer.CallOpts, sender)
}

// GetVoucherEscrow is a free data retrieval call binding the contract method 0x39c9c070.
//
// Solidity: function getVoucherEscrow(string channelId) view returns(address)
func (_Ics20transfer *Ics20transferCaller) GetVoucherEscrow(opts *bind.CallOpts, channelId string) (common.Address, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "getVoucherEscrow", channelId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoucherEscrow is a free data retrieval call binding the contract method 0x39c9c070.
//
// Solidity: function getVoucherEscrow(string channelId) view returns(address)
func (_Ics20transfer *Ics20transferSession) GetVoucherEscrow(channelId string) (common.Address, error) {
	return _Ics20transfer.Contract.GetVoucherEscrow(&_Ics20transfer.CallOpts, channelId)
}

// GetVoucherEscrow is a free data retrieval call binding the contract method 0x39c9c070.
//
// Solidity: function getVoucherEscrow(string channelId) view returns(address)
func (_Ics20transfer *Ics20transferCallerSession) GetVoucherEscrow(channelId string) (common.Address, error) {
	return _Ics20transfer.Contract.GetVoucherEscrow(&_Ics20transfer.CallOpts, channelId)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ics20transfer *Ics20transferCaller) IbcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "ibcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ics20transfer *Ics20transferSession) IbcAddress() (common.Address, error) {
	return _Ics20transfer.Contract.IbcAddress(&_Ics20transfer.CallOpts)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_Ics20transfer *Ics20transferCallerSession) IbcAddress() (common.Address, error) {
	return _Ics20transfer.Contract.IbcAddress(&_Ics20transfer.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics20transfer *Ics20transferCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ics20transfer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics20transfer *Ics20transferSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ics20transfer.Contract.SupportsInterface(&_Ics20transfer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics20transfer *Ics20transferCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ics20transfer.Contract.SupportsInterface(&_Ics20transfer.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactor) Deposit(opts *bind.TransactOpts, to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "deposit", to, tokenContract, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferSession) Deposit(to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Deposit(&_Ics20transfer.TransactOpts, to, tokenContract, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactorSession) Deposit(to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Deposit(&_Ics20transfer.TransactOpts, to, tokenContract, amount)
}

// DepositSendTransfer is a paid mutator transaction binding the contract method 0x8e8f8ca3.
//
// Solidity: function depositSendTransfer(string sourceChannel, address tokenContract, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferTransactor) DepositSendTransfer(opts *bind.TransactOpts, sourceChannel string, tokenContract common.Address, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "depositSendTransfer", sourceChannel, tokenContract, amount, receiver, timeout)
}

// DepositSendTransfer is a paid mutator transaction binding the contract method 0x8e8f8ca3.
//
// Solidity: function depositSendTransfer(string sourceChannel, address tokenContract, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferSession) DepositSendTransfer(sourceChannel string, tokenContract common.Address, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.Contract.DepositSendTransfer(&_Ics20transfer.TransactOpts, sourceChannel, tokenContract, amount, receiver, timeout)
}

// DepositSendTransfer is a paid mutator transaction binding the contract method 0x8e8f8ca3.
//
// Solidity: function depositSendTransfer(string sourceChannel, address tokenContract, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferTransactorSession) DepositSendTransfer(sourceChannel string, tokenContract common.Address, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.Contract.DepositSendTransfer(&_Ics20transfer.TransactOpts, sourceChannel, tokenContract, amount, receiver, timeout)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ics20transfer *Ics20transferTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet Packet, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onAcknowledgementPacket", packet, acknowledgement, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ics20transfer *Ics20transferSession) OnAcknowledgementPacket(packet Packet, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnAcknowledgementPacket(&_Ics20transfer.TransactOpts, packet, acknowledgement, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xfb8b532e.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement, address ) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnAcknowledgementPacket(packet Packet, acknowledgement []byte, arg2 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnAcknowledgementPacket(&_Ics20transfer.TransactOpts, packet, acknowledgement, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanCloseConfirm", arg0)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferSession) OnChanCloseConfirm(arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanCloseConfirm(&_Ics20transfer.TransactOpts, arg0)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x38c858bc.
//
// Solidity: function onChanCloseConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanCloseConfirm(arg0 IIBCModuleMsgOnChanCloseConfirm) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanCloseConfirm(&_Ics20transfer.TransactOpts, arg0)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) msg_) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanCloseInit(opts *bind.TransactOpts, msg_ IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanCloseInit", msg_)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) msg_) returns()
func (_Ics20transfer *Ics20transferSession) OnChanCloseInit(msg_ IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanCloseInit(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x3c7df3fb.
//
// Solidity: function onChanCloseInit((string,string) msg_) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanCloseInit(msg_ IIBCModuleMsgOnChanCloseInit) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanCloseInit(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) msg_) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenAck(opts *bind.TransactOpts, msg_ IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenAck", msg_)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) msg_) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenAck(msg_ IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenAck(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x12f6ff6f.
//
// Solidity: function onChanOpenAck((string,string,string) msg_) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenAck(msg_ IIBCModuleMsgOnChanOpenAck) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenAck(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenConfirm", arg0)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferSession) OnChanOpenConfirm(arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenConfirm(&_Ics20transfer.TransactOpts, arg0)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0x81b174dc.
//
// Solidity: function onChanOpenConfirm((string,string) ) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenConfirm(arg0 IIBCModuleMsgOnChanOpenConfirm) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenConfirm(&_Ics20transfer.TransactOpts, arg0)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenInit(opts *bind.TransactOpts, msg_ IIBCModuleInitializerMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenInit", msg_)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferSession) OnChanOpenInit(msg_ IIBCModuleInitializerMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenInit(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x0b7b4ccb.
//
// Solidity: function onChanOpenInit((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenInit(msg_ IIBCModuleInitializerMsgOnChanOpenInit) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenInit(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferTransactor) OnChanOpenTry(opts *bind.TransactOpts, msg_ IIBCModuleInitializerMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onChanOpenTry", msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferSession) OnChanOpenTry(msg_ IIBCModuleInitializerMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenTry(&_Ics20transfer.TransactOpts, msg_)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0xa7a61e66.
//
// Solidity: function onChanOpenTry((uint8,string[],string,string,(string,string),string) msg_) returns(address, string)
func (_Ics20transfer *Ics20transferTransactorSession) OnChanOpenTry(msg_ IIBCModuleInitializerMsgOnChanOpenTry) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnChanOpenTry(&_Ics20transfer.TransactOpts, msg_)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferTransactor) OnRecvPacket(opts *bind.TransactOpts, packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onRecvPacket", packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferSession) OnRecvPacket(packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnRecvPacket(&_Ics20transfer.TransactOpts, packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2301c6f5.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns(bytes acknowledgement)
func (_Ics20transfer *Ics20transferTransactorSession) OnRecvPacket(packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnRecvPacket(&_Ics20transfer.TransactOpts, packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns()
func (_Ics20transfer *Ics20transferTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "onTimeoutPacket", packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns()
func (_Ics20transfer *Ics20transferSession) OnTimeoutPacket(packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnTimeoutPacket(&_Ics20transfer.TransactOpts, packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x52c7157d.
//
// Solidity: function onTimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, address ) returns()
func (_Ics20transfer *Ics20transferTransactorSession) OnTimeoutPacket(packet Packet, arg1 common.Address) (*types.Transaction, error) {
	return _Ics20transfer.Contract.OnTimeoutPacket(&_Ics20transfer.TransactOpts, packet, arg1)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xb3d1f4a3.
//
// Solidity: function sendTransfer(string sourceChannel, string denom, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferTransactor) SendTransfer(opts *bind.TransactOpts, sourceChannel string, denom string, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "sendTransfer", sourceChannel, denom, amount, receiver, timeout)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xb3d1f4a3.
//
// Solidity: function sendTransfer(string sourceChannel, string denom, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferSession) SendTransfer(sourceChannel string, denom string, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransfer(&_Ics20transfer.TransactOpts, sourceChannel, denom, amount, receiver, timeout)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xb3d1f4a3.
//
// Solidity: function sendTransfer(string sourceChannel, string denom, uint256 amount, string receiver, ((uint64,uint64),uint64) timeout) returns(uint64)
func (_Ics20transfer *Ics20transferTransactorSession) SendTransfer(sourceChannel string, denom string, amount *big.Int, receiver string, timeout ICS20LibTimeout) (*types.Transaction, error) {
	return _Ics20transfer.Contract.SendTransfer(&_Ics20transfer.TransactOpts, sourceChannel, denom, amount, receiver, timeout)
}

// Transfer is a paid mutator transaction binding the contract method 0xfff3a01b.
//
// Solidity: function transfer(address to, string denom, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactor) Transfer(opts *bind.TransactOpts, to common.Address, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "transfer", to, denom, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xfff3a01b.
//
// Solidity: function transfer(address to, string denom, uint256 amount) returns()
func (_Ics20transfer *Ics20transferSession) Transfer(to common.Address, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Transfer(&_Ics20transfer.TransactOpts, to, denom, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xfff3a01b.
//
// Solidity: function transfer(address to, string denom, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactorSession) Transfer(to common.Address, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Transfer(&_Ics20transfer.TransactOpts, to, denom, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.contract.Transact(opts, "withdraw", to, tokenContract, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferSession) Withdraw(to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Withdraw(&_Ics20transfer.TransactOpts, to, tokenContract, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address tokenContract, uint256 amount) returns()
func (_Ics20transfer *Ics20transferTransactorSession) Withdraw(to common.Address, tokenContract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ics20transfer.Contract.Withdraw(&_Ics20transfer.TransactOpts, to, tokenContract, amount)
}
