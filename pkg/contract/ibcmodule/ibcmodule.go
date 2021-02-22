// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcmodule

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

// IBCMsgsMsgChannelOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgChannelOpenAck struct {
	PortId                string
	ChannelId             string
	CounterpartyVersion   string
	CounterpartyChannelId string
	ProofTry              []byte
	ProofHeight           uint64
}

// IBCMsgsMsgChannelOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgChannelOpenConfirm struct {
	PortId      string
	ChannelId   string
	ProofAck    []byte
	ProofHeight uint64
}

// IBCMsgsMsgChannelOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgChannelOpenInit struct {
	PortId    string
	ChannelId string
	Channel   ChannelData
}

// IBCMsgsMsgChannelOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgChannelOpenTry struct {
	PortId              string
	ChannelId           string
	Channel             ChannelData
	CounterpartyVersion string
	ProofInit           []byte
	ProofHeight         uint64
}

// IBCMsgsMsgConnectionOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgConnectionOpenAck struct {
	ConnectionId             string
	ClientStateBytes         []byte
	Version                  VersionData
	CounterpartyConnectionID string
	ProofTry                 []byte
	ProofClient              []byte
	ProofConsensus           []byte
	ProofHeight              uint64
	ConsensusHeight          uint64
}

// IBCMsgsMsgConnectionOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgConnectionOpenConfirm struct {
	ConnectionId string
	ProofAck     []byte
	ProofHeight  uint64
}

// IBCMsgsMsgConnectionOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgConnectionOpenInit struct {
	ClientId     string
	ConnectionId string
	Counterparty CounterpartyData
	DelayPeriod  uint64
}

// IBCMsgsMsgConnectionOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgConnectionOpenTry struct {
	ConnectionId         string
	Counterparty         CounterpartyData
	DelayPeriod          uint64
	ClientId             string
	ClientStateBytes     []byte
	CounterpartyVersions []VersionData
	ProofInit            []byte
	ProofClient          []byte
	ProofConsensus       []byte
	ProofHeight          uint64
	ConsensusHeight      uint64
}

// IBCMsgsMsgCreateClient is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgCreateClient struct {
	ClientId            string
	ClientType          string
	Height              uint64
	ClientStateBytes    []byte
	ConsensusStateBytes []byte
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

// IBCMsgsMsgUpdateClient is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgUpdateClient struct {
	ClientId string
	Header   []byte
}

// IBCRoutingModuleModule is an auto generated low-level Go binding around an user-defined struct.
type IBCRoutingModuleModule struct {
	Callbacks common.Address
	Exists    bool
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

// IbcmoduleABI is the input ABI used to generate the binding from.
const IbcmoduleABI = "[{\"inputs\":[{\"internalType\":\"contractIBCStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBFT2Client\",\"name\":\"ibft2Client\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data\",\"name\":\"version\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyConnectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"counterpartyVersions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgCreateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"lookupModule\",\"outputs\":[{\"components\":[{\"internalType\":\"contractCallbacksI\",\"name\":\"callbacks\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIBCRoutingModule.Module\",\"name\":\"module\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"contractIClient\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"registerClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgUpdateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketAcknowledgement\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"acknowledgePacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"handlePacketRecvWithoutVerification\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ibcmodule is an auto generated Go binding around an Ethereum contract.
type Ibcmodule struct {
	IbcmoduleCaller     // Read-only binding to the contract
	IbcmoduleTransactor // Write-only binding to the contract
	IbcmoduleFilterer   // Log filterer for contract events
}

// IbcmoduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcmoduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmoduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcmoduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmoduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcmoduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcmoduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcmoduleSession struct {
	Contract     *Ibcmodule        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcmoduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcmoduleCallerSession struct {
	Contract *IbcmoduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IbcmoduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcmoduleTransactorSession struct {
	Contract     *IbcmoduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IbcmoduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcmoduleRaw struct {
	Contract *Ibcmodule // Generic contract binding to access the raw methods on
}

// IbcmoduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcmoduleCallerRaw struct {
	Contract *IbcmoduleCaller // Generic read-only contract binding to access the raw methods on
}

// IbcmoduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcmoduleTransactorRaw struct {
	Contract *IbcmoduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcmodule creates a new instance of Ibcmodule, bound to a specific deployed contract.
func NewIbcmodule(address common.Address, backend bind.ContractBackend) (*Ibcmodule, error) {
	contract, err := bindIbcmodule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcmodule{IbcmoduleCaller: IbcmoduleCaller{contract: contract}, IbcmoduleTransactor: IbcmoduleTransactor{contract: contract}, IbcmoduleFilterer: IbcmoduleFilterer{contract: contract}}, nil
}

// NewIbcmoduleCaller creates a new read-only instance of Ibcmodule, bound to a specific deployed contract.
func NewIbcmoduleCaller(address common.Address, caller bind.ContractCaller) (*IbcmoduleCaller, error) {
	contract, err := bindIbcmodule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcmoduleCaller{contract: contract}, nil
}

// NewIbcmoduleTransactor creates a new write-only instance of Ibcmodule, bound to a specific deployed contract.
func NewIbcmoduleTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcmoduleTransactor, error) {
	contract, err := bindIbcmodule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcmoduleTransactor{contract: contract}, nil
}

// NewIbcmoduleFilterer creates a new log filterer instance of Ibcmodule, bound to a specific deployed contract.
func NewIbcmoduleFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcmoduleFilterer, error) {
	contract, err := bindIbcmodule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcmoduleFilterer{contract: contract}, nil
}

// bindIbcmodule binds a generic wrapper to an already deployed contract.
func bindIbcmodule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcmoduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcmodule *IbcmoduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcmodule.Contract.IbcmoduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcmodule *IbcmoduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcmodule.Contract.IbcmoduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcmodule *IbcmoduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcmodule.Contract.IbcmoduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcmodule *IbcmoduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcmodule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcmodule *IbcmoduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcmodule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcmodule *IbcmoduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcmodule.Contract.contract.Transact(opts, method, params...)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibcmodule *IbcmoduleCaller) LookupModule(opts *bind.CallOpts, portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	var out []interface{}
	err := _Ibcmodule.contract.Call(opts, &out, "lookupModule", portId)

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
func (_Ibcmodule *IbcmoduleSession) LookupModule(portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	return _Ibcmodule.Contract.LookupModule(&_Ibcmodule.CallOpts, portId)
}

// LookupModule is a free data retrieval call binding the contract method 0x557c9bb5.
//
// Solidity: function lookupModule(string portId) view returns((address,bool) module, bool found)
func (_Ibcmodule *IbcmoduleCallerSession) LookupModule(portId string) (struct {
	Module IBCRoutingModuleModule
	Found  bool
}, error) {
	return _Ibcmodule.Contract.LookupModule(&_Ibcmodule.CallOpts, portId)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) AcknowledgePacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "acknowledgePacket", msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcmodule.Contract.AcknowledgePacket(&_Ibcmodule.TransactOpts, msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibcmodule.Contract.AcknowledgePacket(&_Ibcmodule.TransactOpts, msg_)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcmodule *IbcmoduleTransactor) BindPort(opts *bind.TransactOpts, portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "bindPort", portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcmodule *IbcmoduleSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcmodule.Contract.BindPort(&_Ibcmodule.TransactOpts, portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibcmodule.Contract.BindPort(&_Ibcmodule.TransactOpts, portId, moduleAddress)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) ChannelOpenAck(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "channelOpenAck", msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) ChannelOpenAck(msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenAck(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) ChannelOpenAck(msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenAck(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "channelOpenConfirm", msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) ChannelOpenConfirm(msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenConfirm(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) ChannelOpenConfirm(msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenConfirm(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactor) ChannelOpenInit(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "channelOpenInit", msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcmodule *IbcmoduleSession) ChannelOpenInit(msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenInit(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactorSession) ChannelOpenInit(msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenInit(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactor) ChannelOpenTry(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "channelOpenTry", msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleSession) ChannelOpenTry(msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenTry(&_Ibcmodule.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactorSession) ChannelOpenTry(msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ChannelOpenTry(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) ConnectionOpenAck(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "connectionOpenAck", msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) ConnectionOpenAck(msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenAck(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) ConnectionOpenAck(msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenAck(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) ConnectionOpenConfirm(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "connectionOpenConfirm", msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) ConnectionOpenConfirm(msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenConfirm(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) ConnectionOpenConfirm(msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenConfirm(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe0d887da.
//
// Solidity: function connectionOpenInit((string,string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactor) ConnectionOpenInit(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "connectionOpenInit", msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe0d887da.
//
// Solidity: function connectionOpenInit((string,string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleSession) ConnectionOpenInit(msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenInit(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xe0d887da.
//
// Solidity: function connectionOpenInit((string,string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactorSession) ConnectionOpenInit(msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenInit(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactor) ConnectionOpenTry(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "connectionOpenTry", msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleSession) ConnectionOpenTry(msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenTry(&_Ibcmodule.TransactOpts, msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibcmodule *IbcmoduleTransactorSession) ConnectionOpenTry(msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibcmodule.Contract.ConnectionOpenTry(&_Ibcmodule.TransactOpts, msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) CreateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "createClient", msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcmodule.Contract.CreateClient(&_Ibcmodule.TransactOpts, msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xbfa9c864.
//
// Solidity: function createClient((string,string,uint64,bytes,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibcmodule.Contract.CreateClient(&_Ibcmodule.TransactOpts, msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleTransactor) HandlePacketRecvWithoutVerification(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "handlePacketRecvWithoutVerification", msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.Contract.HandlePacketRecvWithoutVerification(&_Ibcmodule.TransactOpts, msg_)
}

// HandlePacketRecvWithoutVerification is a paid mutator transaction binding the contract method 0x9ee76fac.
//
// Solidity: function handlePacketRecvWithoutVerification(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleTransactorSession) HandlePacketRecvWithoutVerification(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.Contract.HandlePacketRecvWithoutVerification(&_Ibcmodule.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleTransactor) RecvPacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "recvPacket", msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.Contract.RecvPacket(&_Ibcmodule.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes)
func (_Ibcmodule *IbcmoduleTransactorSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibcmodule.Contract.RecvPacket(&_Ibcmodule.TransactOpts, msg_)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcmodule *IbcmoduleTransactor) RegisterClient(opts *bind.TransactOpts, clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "registerClient", clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcmodule *IbcmoduleSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcmodule.Contract.RegisterClient(&_Ibcmodule.TransactOpts, clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibcmodule.Contract.RegisterClient(&_Ibcmodule.TransactOpts, clientType, client)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcmodule *IbcmoduleTransactor) SendPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "sendPacket", packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcmodule *IbcmoduleSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibcmodule.Contract.SendPacket(&_Ibcmodule.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibcmodule.Contract.SendPacket(&_Ibcmodule.TransactOpts, packet)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactor) UpdateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcmodule.contract.Transact(opts, "updateClient", msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcmodule.Contract.UpdateClient(&_Ibcmodule.TransactOpts, msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibcmodule *IbcmoduleTransactorSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibcmodule.Contract.UpdateClient(&_Ibcmodule.TransactOpts, msg_)
}
