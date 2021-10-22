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

// ConnectionEndData is an auto generated low-level Go binding around an user-defined struct.
type ConnectionEndData struct {
	ClientId     string
	Versions     []VersionData
	State        uint8
	Counterparty CounterpartyData
	DelayPeriod  uint64
}

// CounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type CounterpartyData struct {
	ClientId     string
	ConnectionId string
	Prefix       MerklePrefixData
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type MerklePrefixData struct {
	KeyPrefix []byte
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

// VersionData is an auto generated low-level Go binding around an user-defined struct.
type VersionData struct {
	Identifier string
	Features   []string
}

// IbchostABI is the input ABI used to generate the binding from.
const IbchostABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedChannelIdentifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedClientIdentifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedConnectionIdentifier\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcModule_\",\"type\":\"address\"}],\"name\":\"setIBCModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyIBCModule\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"clientImpl\",\"type\":\"address\"}],\"name\":\"setClientImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"}],\"name\":\"getClientImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"}],\"name\":\"setClientType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"}],\"name\":\"setClientState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"setConsensusState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"processedTime\",\"type\":\"uint256\"}],\"name\":\"setProcessedTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getProcessedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"processedHeight\",\"type\":\"uint256\"}],\"name\":\"setProcessedHeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getProcessedHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"}],\"name\":\"setConnection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getConnection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"name\":\"setChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getChannel\",\"outputs\":[{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceRecv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceRecv\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceAck\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"setPacketCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"deletePacketCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"getPacketCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"makePacketCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"setPacketAcknowledgementCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"getPacketAcknowledgementCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"makePacketAcknowledgementCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setPacketReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"hasPacketReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"claimCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"authenticateCapability\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"getModuleOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"}],\"name\":\"generateClientIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generateConnectionIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generateChannelIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AuthenticateCapability is a free data retrieval call binding the contract method 0x2d46858f.
//
// Solidity: function authenticateCapability(bytes name, address addr) view returns(bool)
func (_Ibchost *IbchostCaller) AuthenticateCapability(opts *bind.CallOpts, name []byte, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "authenticateCapability", name, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthenticateCapability is a free data retrieval call binding the contract method 0x2d46858f.
//
// Solidity: function authenticateCapability(bytes name, address addr) view returns(bool)
func (_Ibchost *IbchostSession) AuthenticateCapability(name []byte, addr common.Address) (bool, error) {
	return _Ibchost.Contract.AuthenticateCapability(&_Ibchost.CallOpts, name, addr)
}

// AuthenticateCapability is a free data retrieval call binding the contract method 0x2d46858f.
//
// Solidity: function authenticateCapability(bytes name, address addr) view returns(bool)
func (_Ibchost *IbchostCallerSession) AuthenticateCapability(name []byte, addr common.Address) (bool, error) {
	return _Ibchost.Contract.AuthenticateCapability(&_Ibchost.CallOpts, name, addr)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Ibchost *IbchostCaller) GetChannel(opts *bind.CallOpts, portId string, channelId string) (ChannelData, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getChannel", portId, channelId)

	if err != nil {
		return *new(ChannelData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ChannelData)).(*ChannelData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Ibchost *IbchostSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibchost.Contract.GetChannel(&_Ibchost.CallOpts, portId, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Ibchost *IbchostCallerSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibchost.Contract.GetChannel(&_Ibchost.CallOpts, portId, channelId)
}

// GetClientImpl is a free data retrieval call binding the contract method 0xbfe7aa66.
//
// Solidity: function getClientImpl(string clientType) view returns(address, bool)
func (_Ibchost *IbchostCaller) GetClientImpl(opts *bind.CallOpts, clientType string) (common.Address, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getClientImpl", clientType)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientImpl is a free data retrieval call binding the contract method 0xbfe7aa66.
//
// Solidity: function getClientImpl(string clientType) view returns(address, bool)
func (_Ibchost *IbchostSession) GetClientImpl(clientType string) (common.Address, bool, error) {
	return _Ibchost.Contract.GetClientImpl(&_Ibchost.CallOpts, clientType)
}

// GetClientImpl is a free data retrieval call binding the contract method 0xbfe7aa66.
//
// Solidity: function getClientImpl(string clientType) view returns(address, bool)
func (_Ibchost *IbchostCallerSession) GetClientImpl(clientType string) (common.Address, bool, error) {
	return _Ibchost.Contract.GetClientImpl(&_Ibchost.CallOpts, clientType)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchost *IbchostCaller) GetClientState(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchost *IbchostSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibchost.Contract.GetClientState(&_Ibchost.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchost *IbchostCallerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibchost.Contract.GetClientState(&_Ibchost.CallOpts, clientId)
}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchost *IbchostCaller) GetClientType(opts *bind.CallOpts, clientId string) (string, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getClientType", clientId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchost *IbchostSession) GetClientType(clientId string) (string, error) {
	return _Ibchost.Contract.GetClientType(&_Ibchost.CallOpts, clientId)
}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchost *IbchostCallerSession) GetClientType(clientId string) (string, error) {
	return _Ibchost.Contract.GetClientType(&_Ibchost.CallOpts, clientId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection, bool)
func (_Ibchost *IbchostCaller) GetConnection(opts *bind.CallOpts, connectionId string) (ConnectionEndData, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getConnection", connectionId)

	if err != nil {
		return *new(ConnectionEndData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection, bool)
func (_Ibchost *IbchostSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibchost.Contract.GetConnection(&_Ibchost.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection, bool)
func (_Ibchost *IbchostCallerSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibchost.Contract.GetConnection(&_Ibchost.CallOpts, connectionId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns(bytes, bool)
func (_Ibchost *IbchostCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height uint64) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns(bytes, bool)
func (_Ibchost *IbchostSession) GetConsensusState(clientId string, height uint64) ([]byte, bool, error) {
	return _Ibchost.Contract.GetConsensusState(&_Ibchost.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns(bytes, bool)
func (_Ibchost *IbchostCallerSession) GetConsensusState(clientId string, height uint64) ([]byte, bool, error) {
	return _Ibchost.Contract.GetConsensusState(&_Ibchost.CallOpts, clientId, height)
}

// GetModuleOwner is a free data retrieval call binding the contract method 0xace6cfb7.
//
// Solidity: function getModuleOwner(bytes name) view returns(address, bool)
func (_Ibchost *IbchostCaller) GetModuleOwner(opts *bind.CallOpts, name []byte) (common.Address, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getModuleOwner", name)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetModuleOwner is a free data retrieval call binding the contract method 0xace6cfb7.
//
// Solidity: function getModuleOwner(bytes name) view returns(address, bool)
func (_Ibchost *IbchostSession) GetModuleOwner(name []byte) (common.Address, bool, error) {
	return _Ibchost.Contract.GetModuleOwner(&_Ibchost.CallOpts, name)
}

// GetModuleOwner is a free data retrieval call binding the contract method 0xace6cfb7.
//
// Solidity: function getModuleOwner(bytes name) view returns(address, bool)
func (_Ibchost *IbchostCallerSession) GetModuleOwner(name []byte) (common.Address, bool, error) {
	return _Ibchost.Contract.GetModuleOwner(&_Ibchost.CallOpts, name)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCaller) GetNextSequenceAck(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getNextSequenceAck", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceAck(&_Ibchost.CallOpts, portId, channelId)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCallerSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceAck(&_Ibchost.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCaller) GetNextSequenceRecv(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getNextSequenceRecv", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceRecv(&_Ibchost.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCallerSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceRecv(&_Ibchost.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCaller) GetNextSequenceSend(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getNextSequenceSend", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceSend(&_Ibchost.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchost *IbchostCallerSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibchost.Contract.GetNextSequenceSend(&_Ibchost.CallOpts, portId, channelId)
}

// GetPacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x71f56c59.
//
// Solidity: function getPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostCaller) GetPacketAcknowledgementCommitment(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getPacketAcknowledgementCommitment", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetPacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x71f56c59.
//
// Solidity: function getPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostSession) GetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibchost.Contract.GetPacketAcknowledgementCommitment(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// GetPacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x71f56c59.
//
// Solidity: function getPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostCallerSession) GetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibchost.Contract.GetPacketAcknowledgementCommitment(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// GetPacketCommitment is a free data retrieval call binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostCaller) GetPacketCommitment(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getPacketCommitment", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetPacketCommitment is a free data retrieval call binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostSession) GetPacketCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibchost.Contract.GetPacketCommitment(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// GetPacketCommitment is a free data retrieval call binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibchost *IbchostCallerSession) GetPacketCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibchost.Contract.GetPacketCommitment(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// GetProcessedHeight is a free data retrieval call binding the contract method 0x36f57592.
//
// Solidity: function getProcessedHeight(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostCaller) GetProcessedHeight(opts *bind.CallOpts, clientId string, height uint64) (*big.Int, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getProcessedHeight", clientId, height)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetProcessedHeight is a free data retrieval call binding the contract method 0x36f57592.
//
// Solidity: function getProcessedHeight(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostSession) GetProcessedHeight(clientId string, height uint64) (*big.Int, bool, error) {
	return _Ibchost.Contract.GetProcessedHeight(&_Ibchost.CallOpts, clientId, height)
}

// GetProcessedHeight is a free data retrieval call binding the contract method 0x36f57592.
//
// Solidity: function getProcessedHeight(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostCallerSession) GetProcessedHeight(clientId string, height uint64) (*big.Int, bool, error) {
	return _Ibchost.Contract.GetProcessedHeight(&_Ibchost.CallOpts, clientId, height)
}

// GetProcessedTime is a free data retrieval call binding the contract method 0x2f80e3ad.
//
// Solidity: function getProcessedTime(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostCaller) GetProcessedTime(opts *bind.CallOpts, clientId string, height uint64) (*big.Int, bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "getProcessedTime", clientId, height)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetProcessedTime is a free data retrieval call binding the contract method 0x2f80e3ad.
//
// Solidity: function getProcessedTime(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostSession) GetProcessedTime(clientId string, height uint64) (*big.Int, bool, error) {
	return _Ibchost.Contract.GetProcessedTime(&_Ibchost.CallOpts, clientId, height)
}

// GetProcessedTime is a free data retrieval call binding the contract method 0x2f80e3ad.
//
// Solidity: function getProcessedTime(string clientId, uint64 height) view returns(uint256, bool)
func (_Ibchost *IbchostCallerSession) GetProcessedTime(clientId string, height uint64) (*big.Int, bool, error) {
	return _Ibchost.Contract.GetProcessedTime(&_Ibchost.CallOpts, clientId, height)
}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibchost *IbchostCaller) HasPacketReceipt(opts *bind.CallOpts, portId string, channelId string, sequence uint64) (bool, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "hasPacketReceipt", portId, channelId, sequence)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibchost *IbchostSession) HasPacketReceipt(portId string, channelId string, sequence uint64) (bool, error) {
	return _Ibchost.Contract.HasPacketReceipt(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibchost *IbchostCallerSession) HasPacketReceipt(portId string, channelId string, sequence uint64) (bool, error) {
	return _Ibchost.Contract.HasPacketReceipt(&_Ibchost.CallOpts, portId, channelId, sequence)
}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) pure returns(bytes32)
func (_Ibchost *IbchostCaller) MakePacketAcknowledgementCommitment(opts *bind.CallOpts, acknowledgement []byte) ([32]byte, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "makePacketAcknowledgementCommitment", acknowledgement)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) pure returns(bytes32)
func (_Ibchost *IbchostSession) MakePacketAcknowledgementCommitment(acknowledgement []byte) ([32]byte, error) {
	return _Ibchost.Contract.MakePacketAcknowledgementCommitment(&_Ibchost.CallOpts, acknowledgement)
}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) pure returns(bytes32)
func (_Ibchost *IbchostCallerSession) MakePacketAcknowledgementCommitment(acknowledgement []byte) ([32]byte, error) {
	return _Ibchost.Contract.MakePacketAcknowledgementCommitment(&_Ibchost.CallOpts, acknowledgement)
}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) pure returns(bytes32)
func (_Ibchost *IbchostCaller) MakePacketCommitment(opts *bind.CallOpts, packet PacketData) ([32]byte, error) {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "makePacketCommitment", packet)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) pure returns(bytes32)
func (_Ibchost *IbchostSession) MakePacketCommitment(packet PacketData) ([32]byte, error) {
	return _Ibchost.Contract.MakePacketCommitment(&_Ibchost.CallOpts, packet)
}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) pure returns(bytes32)
func (_Ibchost *IbchostCallerSession) MakePacketCommitment(packet PacketData) ([32]byte, error) {
	return _Ibchost.Contract.MakePacketCommitment(&_Ibchost.CallOpts, packet)
}

// OnlyIBCModule is a free data retrieval call binding the contract method 0x1649a4f7.
//
// Solidity: function onlyIBCModule() view returns()
func (_Ibchost *IbchostCaller) OnlyIBCModule(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Ibchost.contract.Call(opts, &out, "onlyIBCModule")

	if err != nil {
		return err
	}

	return err

}

// OnlyIBCModule is a free data retrieval call binding the contract method 0x1649a4f7.
//
// Solidity: function onlyIBCModule() view returns()
func (_Ibchost *IbchostSession) OnlyIBCModule() error {
	return _Ibchost.Contract.OnlyIBCModule(&_Ibchost.CallOpts)
}

// OnlyIBCModule is a free data retrieval call binding the contract method 0x1649a4f7.
//
// Solidity: function onlyIBCModule() view returns()
func (_Ibchost *IbchostCallerSession) OnlyIBCModule() error {
	return _Ibchost.Contract.OnlyIBCModule(&_Ibchost.CallOpts)
}

// ClaimCapability is a paid mutator transaction binding the contract method 0xbe3d6e92.
//
// Solidity: function claimCapability(bytes name, address addr) returns()
func (_Ibchost *IbchostTransactor) ClaimCapability(opts *bind.TransactOpts, name []byte, addr common.Address) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "claimCapability", name, addr)
}

// ClaimCapability is a paid mutator transaction binding the contract method 0xbe3d6e92.
//
// Solidity: function claimCapability(bytes name, address addr) returns()
func (_Ibchost *IbchostSession) ClaimCapability(name []byte, addr common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.ClaimCapability(&_Ibchost.TransactOpts, name, addr)
}

// ClaimCapability is a paid mutator transaction binding the contract method 0xbe3d6e92.
//
// Solidity: function claimCapability(bytes name, address addr) returns()
func (_Ibchost *IbchostTransactorSession) ClaimCapability(name []byte, addr common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.ClaimCapability(&_Ibchost.TransactOpts, name, addr)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactor) DeletePacketCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "deletePacketCommitment", portId, channelId, sequence)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostSession) DeletePacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.DeletePacketCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactorSession) DeletePacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.DeletePacketCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// GenerateChannelIdentifier is a paid mutator transaction binding the contract method 0x53ccd1c8.
//
// Solidity: function generateChannelIdentifier() returns(string)
func (_Ibchost *IbchostTransactor) GenerateChannelIdentifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "generateChannelIdentifier")
}

// GenerateChannelIdentifier is a paid mutator transaction binding the contract method 0x53ccd1c8.
//
// Solidity: function generateChannelIdentifier() returns(string)
func (_Ibchost *IbchostSession) GenerateChannelIdentifier() (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateChannelIdentifier(&_Ibchost.TransactOpts)
}

// GenerateChannelIdentifier is a paid mutator transaction binding the contract method 0x53ccd1c8.
//
// Solidity: function generateChannelIdentifier() returns(string)
func (_Ibchost *IbchostTransactorSession) GenerateChannelIdentifier() (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateChannelIdentifier(&_Ibchost.TransactOpts)
}

// GenerateClientIdentifier is a paid mutator transaction binding the contract method 0x22c99536.
//
// Solidity: function generateClientIdentifier(string clientType) returns(string)
func (_Ibchost *IbchostTransactor) GenerateClientIdentifier(opts *bind.TransactOpts, clientType string) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "generateClientIdentifier", clientType)
}

// GenerateClientIdentifier is a paid mutator transaction binding the contract method 0x22c99536.
//
// Solidity: function generateClientIdentifier(string clientType) returns(string)
func (_Ibchost *IbchostSession) GenerateClientIdentifier(clientType string) (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateClientIdentifier(&_Ibchost.TransactOpts, clientType)
}

// GenerateClientIdentifier is a paid mutator transaction binding the contract method 0x22c99536.
//
// Solidity: function generateClientIdentifier(string clientType) returns(string)
func (_Ibchost *IbchostTransactorSession) GenerateClientIdentifier(clientType string) (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateClientIdentifier(&_Ibchost.TransactOpts, clientType)
}

// GenerateConnectionIdentifier is a paid mutator transaction binding the contract method 0xd9e73f50.
//
// Solidity: function generateConnectionIdentifier() returns(string)
func (_Ibchost *IbchostTransactor) GenerateConnectionIdentifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "generateConnectionIdentifier")
}

// GenerateConnectionIdentifier is a paid mutator transaction binding the contract method 0xd9e73f50.
//
// Solidity: function generateConnectionIdentifier() returns(string)
func (_Ibchost *IbchostSession) GenerateConnectionIdentifier() (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateConnectionIdentifier(&_Ibchost.TransactOpts)
}

// GenerateConnectionIdentifier is a paid mutator transaction binding the contract method 0xd9e73f50.
//
// Solidity: function generateConnectionIdentifier() returns(string)
func (_Ibchost *IbchostTransactorSession) GenerateConnectionIdentifier() (*types.Transaction, error) {
	return _Ibchost.Contract.GenerateConnectionIdentifier(&_Ibchost.TransactOpts)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibchost *IbchostTransactor) SetChannel(opts *bind.TransactOpts, portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setChannel", portId, channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibchost *IbchostSession) SetChannel(portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetChannel(&_Ibchost.TransactOpts, portId, channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibchost *IbchostTransactorSession) SetChannel(portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetChannel(&_Ibchost.TransactOpts, portId, channelId, channel)
}

// SetClientImpl is a paid mutator transaction binding the contract method 0x9d0b7650.
//
// Solidity: function setClientImpl(string clientType, address clientImpl) returns()
func (_Ibchost *IbchostTransactor) SetClientImpl(opts *bind.TransactOpts, clientType string, clientImpl common.Address) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setClientImpl", clientType, clientImpl)
}

// SetClientImpl is a paid mutator transaction binding the contract method 0x9d0b7650.
//
// Solidity: function setClientImpl(string clientType, address clientImpl) returns()
func (_Ibchost *IbchostSession) SetClientImpl(clientType string, clientImpl common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientImpl(&_Ibchost.TransactOpts, clientType, clientImpl)
}

// SetClientImpl is a paid mutator transaction binding the contract method 0x9d0b7650.
//
// Solidity: function setClientImpl(string clientType, address clientImpl) returns()
func (_Ibchost *IbchostTransactorSession) SetClientImpl(clientType string, clientImpl common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientImpl(&_Ibchost.TransactOpts, clientType, clientImpl)
}

// SetClientState is a paid mutator transaction binding the contract method 0x0838f56d.
//
// Solidity: function setClientState(string clientId, bytes clientStateBytes) returns()
func (_Ibchost *IbchostTransactor) SetClientState(opts *bind.TransactOpts, clientId string, clientStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setClientState", clientId, clientStateBytes)
}

// SetClientState is a paid mutator transaction binding the contract method 0x0838f56d.
//
// Solidity: function setClientState(string clientId, bytes clientStateBytes) returns()
func (_Ibchost *IbchostSession) SetClientState(clientId string, clientStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientState(&_Ibchost.TransactOpts, clientId, clientStateBytes)
}

// SetClientState is a paid mutator transaction binding the contract method 0x0838f56d.
//
// Solidity: function setClientState(string clientId, bytes clientStateBytes) returns()
func (_Ibchost *IbchostTransactorSession) SetClientState(clientId string, clientStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientState(&_Ibchost.TransactOpts, clientId, clientStateBytes)
}

// SetClientType is a paid mutator transaction binding the contract method 0x78338a17.
//
// Solidity: function setClientType(string clientId, string clientType) returns()
func (_Ibchost *IbchostTransactor) SetClientType(opts *bind.TransactOpts, clientId string, clientType string) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setClientType", clientId, clientType)
}

// SetClientType is a paid mutator transaction binding the contract method 0x78338a17.
//
// Solidity: function setClientType(string clientId, string clientType) returns()
func (_Ibchost *IbchostSession) SetClientType(clientId string, clientType string) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientType(&_Ibchost.TransactOpts, clientId, clientType)
}

// SetClientType is a paid mutator transaction binding the contract method 0x78338a17.
//
// Solidity: function setClientType(string clientId, string clientType) returns()
func (_Ibchost *IbchostTransactorSession) SetClientType(clientId string, clientType string) (*types.Transaction, error) {
	return _Ibchost.Contract.SetClientType(&_Ibchost.TransactOpts, clientId, clientType)
}

// SetConnection is a paid mutator transaction binding the contract method 0x2dbe1450.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection) returns()
func (_Ibchost *IbchostTransactor) SetConnection(opts *bind.TransactOpts, connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setConnection", connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x2dbe1450.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection) returns()
func (_Ibchost *IbchostSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetConnection(&_Ibchost.TransactOpts, connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x2dbe1450.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,(string,string,(bytes)),uint64) connection) returns()
func (_Ibchost *IbchostTransactorSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetConnection(&_Ibchost.TransactOpts, connectionId, connection)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xff618292.
//
// Solidity: function setConsensusState(string clientId, uint64 height, bytes consensusStateBytes) returns()
func (_Ibchost *IbchostTransactor) SetConsensusState(opts *bind.TransactOpts, clientId string, height uint64, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setConsensusState", clientId, height, consensusStateBytes)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xff618292.
//
// Solidity: function setConsensusState(string clientId, uint64 height, bytes consensusStateBytes) returns()
func (_Ibchost *IbchostSession) SetConsensusState(clientId string, height uint64, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetConsensusState(&_Ibchost.TransactOpts, clientId, height, consensusStateBytes)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xff618292.
//
// Solidity: function setConsensusState(string clientId, uint64 height, bytes consensusStateBytes) returns()
func (_Ibchost *IbchostTransactorSession) SetConsensusState(clientId string, height uint64, consensusStateBytes []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetConsensusState(&_Ibchost.TransactOpts, clientId, height, consensusStateBytes)
}

// SetIBCModule is a paid mutator transaction binding the contract method 0x1e742f32.
//
// Solidity: function setIBCModule(address ibcModule_) returns()
func (_Ibchost *IbchostTransactor) SetIBCModule(opts *bind.TransactOpts, ibcModule_ common.Address) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setIBCModule", ibcModule_)
}

// SetIBCModule is a paid mutator transaction binding the contract method 0x1e742f32.
//
// Solidity: function setIBCModule(address ibcModule_) returns()
func (_Ibchost *IbchostSession) SetIBCModule(ibcModule_ common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.SetIBCModule(&_Ibchost.TransactOpts, ibcModule_)
}

// SetIBCModule is a paid mutator transaction binding the contract method 0x1e742f32.
//
// Solidity: function setIBCModule(address ibcModule_) returns()
func (_Ibchost *IbchostTransactorSession) SetIBCModule(ibcModule_ common.Address) (*types.Transaction, error) {
	return _Ibchost.Contract.SetIBCModule(&_Ibchost.TransactOpts, ibcModule_)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactor) SetNextSequenceAck(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setNextSequenceAck", portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceAck(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactorSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceAck(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactor) SetNextSequenceRecv(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setNextSequenceRecv", portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceRecv(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactorSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceRecv(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactor) SetNextSequenceSend(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setNextSequenceSend", portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceSend(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactorSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetNextSequenceSend(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchost *IbchostTransactor) SetPacketAcknowledgementCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setPacketAcknowledgementCommitment", portId, channelId, sequence, acknowledgement)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchost *IbchostSession) SetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketAcknowledgementCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence, acknowledgement)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchost *IbchostTransactorSession) SetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketAcknowledgementCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence, acknowledgement)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchost *IbchostTransactor) SetPacketCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setPacketCommitment", portId, channelId, sequence, packet)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchost *IbchostSession) SetPacketCommitment(portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence, packet)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchost *IbchostTransactorSession) SetPacketCommitment(portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketCommitment(&_Ibchost.TransactOpts, portId, channelId, sequence, packet)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactor) SetPacketReceipt(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setPacketReceipt", portId, channelId, sequence)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostSession) SetPacketReceipt(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketReceipt(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibchost *IbchostTransactorSession) SetPacketReceipt(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibchost.Contract.SetPacketReceipt(&_Ibchost.TransactOpts, portId, channelId, sequence)
}

// SetProcessedHeight is a paid mutator transaction binding the contract method 0x1ea8dfe1.
//
// Solidity: function setProcessedHeight(string clientId, uint64 height, uint256 processedHeight) returns()
func (_Ibchost *IbchostTransactor) SetProcessedHeight(opts *bind.TransactOpts, clientId string, height uint64, processedHeight *big.Int) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setProcessedHeight", clientId, height, processedHeight)
}

// SetProcessedHeight is a paid mutator transaction binding the contract method 0x1ea8dfe1.
//
// Solidity: function setProcessedHeight(string clientId, uint64 height, uint256 processedHeight) returns()
func (_Ibchost *IbchostSession) SetProcessedHeight(clientId string, height uint64, processedHeight *big.Int) (*types.Transaction, error) {
	return _Ibchost.Contract.SetProcessedHeight(&_Ibchost.TransactOpts, clientId, height, processedHeight)
}

// SetProcessedHeight is a paid mutator transaction binding the contract method 0x1ea8dfe1.
//
// Solidity: function setProcessedHeight(string clientId, uint64 height, uint256 processedHeight) returns()
func (_Ibchost *IbchostTransactorSession) SetProcessedHeight(clientId string, height uint64, processedHeight *big.Int) (*types.Transaction, error) {
	return _Ibchost.Contract.SetProcessedHeight(&_Ibchost.TransactOpts, clientId, height, processedHeight)
}

// SetProcessedTime is a paid mutator transaction binding the contract method 0xfcc2fb00.
//
// Solidity: function setProcessedTime(string clientId, uint64 height, uint256 processedTime) returns()
func (_Ibchost *IbchostTransactor) SetProcessedTime(opts *bind.TransactOpts, clientId string, height uint64, processedTime *big.Int) (*types.Transaction, error) {
	return _Ibchost.contract.Transact(opts, "setProcessedTime", clientId, height, processedTime)
}

// SetProcessedTime is a paid mutator transaction binding the contract method 0xfcc2fb00.
//
// Solidity: function setProcessedTime(string clientId, uint64 height, uint256 processedTime) returns()
func (_Ibchost *IbchostSession) SetProcessedTime(clientId string, height uint64, processedTime *big.Int) (*types.Transaction, error) {
	return _Ibchost.Contract.SetProcessedTime(&_Ibchost.TransactOpts, clientId, height, processedTime)
}

// SetProcessedTime is a paid mutator transaction binding the contract method 0xfcc2fb00.
//
// Solidity: function setProcessedTime(string clientId, uint64 height, uint256 processedTime) returns()
func (_Ibchost *IbchostTransactorSession) SetProcessedTime(clientId string, height uint64, processedTime *big.Int) (*types.Transaction, error) {
	return _Ibchost.Contract.SetProcessedTime(&_Ibchost.TransactOpts, clientId, height, processedTime)
}

// IbchostGeneratedChannelIdentifierIterator is returned from FilterGeneratedChannelIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedChannelIdentifier events raised by the Ibchost contract.
type IbchostGeneratedChannelIdentifierIterator struct {
	Event *IbchostGeneratedChannelIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedChannelIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedChannelIdentifier)
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
		it.Event = new(IbchostGeneratedChannelIdentifier)
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
func (it *IbchostGeneratedChannelIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedChannelIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedChannelIdentifier represents a GeneratedChannelIdentifier event raised by the Ibchost contract.
type IbchostGeneratedChannelIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedChannelIdentifier is a free log retrieval operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedChannelIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedChannelIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedChannelIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedChannelIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedChannelIdentifier is a free log subscription operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedChannelIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedChannelIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedChannelIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
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

// ParseGeneratedChannelIdentifier is a log parse operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedChannelIdentifier(log types.Log) (*IbchostGeneratedChannelIdentifier, error) {
	event := new(IbchostGeneratedChannelIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchostGeneratedClientIdentifierIterator is returned from FilterGeneratedClientIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedClientIdentifier events raised by the Ibchost contract.
type IbchostGeneratedClientIdentifierIterator struct {
	Event *IbchostGeneratedClientIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedClientIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedClientIdentifier)
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
		it.Event = new(IbchostGeneratedClientIdentifier)
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
func (it *IbchostGeneratedClientIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedClientIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedClientIdentifier represents a GeneratedClientIdentifier event raised by the Ibchost contract.
type IbchostGeneratedClientIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedClientIdentifier is a free log retrieval operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedClientIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedClientIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedClientIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedClientIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedClientIdentifier is a free log subscription operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedClientIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedClientIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedClientIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
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

// ParseGeneratedClientIdentifier is a log parse operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedClientIdentifier(log types.Log) (*IbchostGeneratedClientIdentifier, error) {
	event := new(IbchostGeneratedClientIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchostGeneratedConnectionIdentifierIterator is returned from FilterGeneratedConnectionIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedConnectionIdentifier events raised by the Ibchost contract.
type IbchostGeneratedConnectionIdentifierIterator struct {
	Event *IbchostGeneratedConnectionIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedConnectionIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedConnectionIdentifier)
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
		it.Event = new(IbchostGeneratedConnectionIdentifier)
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
func (it *IbchostGeneratedConnectionIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedConnectionIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedConnectionIdentifier represents a GeneratedConnectionIdentifier event raised by the Ibchost contract.
type IbchostGeneratedConnectionIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedConnectionIdentifier is a free log retrieval operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedConnectionIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedConnectionIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedConnectionIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedConnectionIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedConnectionIdentifier is a free log subscription operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedConnectionIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedConnectionIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedConnectionIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
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

// ParseGeneratedConnectionIdentifier is a log parse operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedConnectionIdentifier(log types.Log) (*IbchostGeneratedConnectionIdentifier, error) {
	event := new(IbchostGeneratedConnectionIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
