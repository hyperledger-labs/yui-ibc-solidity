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
	PortId  string
	Channel ChannelData
}

// IBCMsgsMsgChannelOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgChannelOpenTry struct {
	PortId              string
	PreviousChannelId   string
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
	Counterparty CounterpartyData
	DelayPeriod  uint64
}

// IBCMsgsMsgConnectionOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgConnectionOpenTry struct {
	PreviousConnectionId string
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

// IbchandlerABI is the input ABI used to generate the binding from.
const IbchandlerABI = "[{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"AcknowledgePacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"RecvPacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"SendPacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationPortId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"WriteAcknowledgement\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getHostAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"contractIClient\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"registerClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientType\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgCreateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"internalType\":\"structIBCMsgs.MsgUpdateClient\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"previousConnectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"delayPeriod\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"counterpartyVersions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data\",\"name\":\"version\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyConnectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofClient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofConsensus\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgConnectionOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"connectionOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"previousChannelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofInit\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenTry\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenTry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyChannelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofTry\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenAck\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proofAck\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgChannelOpenConfirm\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketRecv\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"destinationPortId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"writeAcknowledgement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"proofHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBCMsgs.MsgPacketAcknowledgement\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"acknowledgePacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// GetHostAddress is a free data retrieval call binding the contract method 0xeb238c77.
//
// Solidity: function getHostAddress() view returns(address)
func (_Ibchandler *IbchandlerCaller) GetHostAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getHostAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHostAddress is a free data retrieval call binding the contract method 0xeb238c77.
//
// Solidity: function getHostAddress() view returns(address)
func (_Ibchandler *IbchandlerSession) GetHostAddress() (common.Address, error) {
	return _Ibchandler.Contract.GetHostAddress(&_Ibchandler.CallOpts)
}

// GetHostAddress is a free data retrieval call binding the contract method 0xeb238c77.
//
// Solidity: function getHostAddress() view returns(address)
func (_Ibchandler *IbchandlerCallerSession) GetHostAddress() (common.Address, error) {
	return _Ibchandler.Contract.GetHostAddress(&_Ibchandler.CallOpts)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) AcknowledgePacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "acknowledgePacket", msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.AcknowledgePacket(&_Ibchandler.TransactOpts, msg_)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xfa044e8f.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) AcknowledgePacket(msg_ IBCMsgsMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.AcknowledgePacket(&_Ibchandler.TransactOpts, msg_)
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

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelOpenAck(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenAck", msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) ChannelOpenAck(msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenAck(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0xe46ea828.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenAck(msg_ IBCMsgsMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenAck(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenConfirm", msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) ChannelOpenConfirm(msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenConfirm(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x9e6d4959.
//
// Solidity: function channelOpenConfirm((string,string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenConfirm(msg_ IBCMsgsMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenConfirm(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactor) ChannelOpenInit(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenInit", msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibchandler *IbchandlerSession) ChannelOpenInit(msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenInit(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenInit(msg_ IBCMsgsMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenInit(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactor) ChannelOpenTry(opts *bind.TransactOpts, msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenTry", msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerSession) ChannelOpenTry(msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenTry(&_Ibchandler.TransactOpts, msg_)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x56a5dc5a.
//
// Solidity: function channelOpenTry((string,string,(uint8,uint8,(string,string),string[],string),string,bytes,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenTry(msg_ IBCMsgsMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenTry(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenAck(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenAck", msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) ConnectionOpenAck(msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenAck(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0x6cf60640.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,uint64,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenAck(msg_ IBCMsgsMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenAck(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenConfirm(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenConfirm", msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerSession) ConnectionOpenConfirm(msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenConfirm(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0xec0cda87.
//
// Solidity: function connectionOpenConfirm((string,bytes,uint64) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenConfirm(msg_ IBCMsgsMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenConfirm(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0x01c6400f.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenInit(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenInit", msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0x01c6400f.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerSession) ConnectionOpenInit(msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenInit(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0x01c6400f.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenInit(msg_ IBCMsgsMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenInit(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenTry(opts *bind.TransactOpts, msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenTry", msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerSession) ConnectionOpenTry(msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenTry(&_Ibchandler.TransactOpts, msg_)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x147aec7f.
//
// Solidity: function connectionOpenTry((string,(string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,uint64,uint64) msg_) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenTry(msg_ IBCMsgsMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenTry(&_Ibchandler.TransactOpts, msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xe2b5827d.
//
// Solidity: function createClient((string,uint64,bytes,bytes) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) CreateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "createClient", msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xe2b5827d.
//
// Solidity: function createClient((string,uint64,bytes,bytes) msg_) returns()
func (_Ibchandler *IbchandlerSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.CreateClient(&_Ibchandler.TransactOpts, msg_)
}

// CreateClient is a paid mutator transaction binding the contract method 0xe2b5827d.
//
// Solidity: function createClient((string,uint64,bytes,bytes) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) CreateClient(msg_ IBCMsgsMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.CreateClient(&_Ibchandler.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes acknowledgement)
func (_Ibchandler *IbchandlerTransactor) RecvPacket(opts *bind.TransactOpts, msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "recvPacket", msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes acknowledgement)
func (_Ibchandler *IbchandlerSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.RecvPacket(&_Ibchandler.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0xa3af5cf3.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,uint64) msg_) returns(bytes acknowledgement)
func (_Ibchandler *IbchandlerTransactorSession) RecvPacket(msg_ IBCMsgsMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.RecvPacket(&_Ibchandler.TransactOpts, msg_)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerTransactor) RegisterClient(opts *bind.TransactOpts, clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "registerClient", clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.RegisterClient(&_Ibchandler.TransactOpts, clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerTransactorSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.RegisterClient(&_Ibchandler.TransactOpts, clientType, client)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchandler *IbchandlerTransactor) SendPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "sendPacket", packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchandler *IbchandlerSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibchandler.Contract.SendPacket(&_Ibchandler.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x40835e44.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibchandler *IbchandlerTransactorSession) SendPacket(packet PacketData) (*types.Transaction, error) {
	return _Ibchandler.Contract.SendPacket(&_Ibchandler.TransactOpts, packet)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibchandler *IbchandlerTransactor) UpdateClient(opts *bind.TransactOpts, msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "updateClient", msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibchandler *IbchandlerSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.UpdateClient(&_Ibchandler.TransactOpts, msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) msg_) returns()
func (_Ibchandler *IbchandlerTransactorSession) UpdateClient(msg_ IBCMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.UpdateClient(&_Ibchandler.TransactOpts, msg_)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchandler *IbchandlerTransactor) WriteAcknowledgement(opts *bind.TransactOpts, destinationPortId string, destinationChannel string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "writeAcknowledgement", destinationPortId, destinationChannel, sequence, acknowledgement)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchandler *IbchandlerSession) WriteAcknowledgement(destinationPortId string, destinationChannel string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.WriteAcknowledgement(&_Ibchandler.TransactOpts, destinationPortId, destinationChannel, sequence, acknowledgement)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement) returns()
func (_Ibchandler *IbchandlerTransactorSession) WriteAcknowledgement(destinationPortId string, destinationChannel string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.WriteAcknowledgement(&_Ibchandler.TransactOpts, destinationPortId, destinationChannel, sequence, acknowledgement)
}

// IbchandlerAcknowledgePacketIterator is returned from FilterAcknowledgePacket and is used to iterate over the raw logs and unpacked data for AcknowledgePacket events raised by the Ibchandler contract.
type IbchandlerAcknowledgePacketIterator struct {
	Event *IbchandlerAcknowledgePacket // Event containing the contract specifics and raw log

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
func (it *IbchandlerAcknowledgePacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerAcknowledgePacket)
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
		it.Event = new(IbchandlerAcknowledgePacket)
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
func (it *IbchandlerAcknowledgePacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerAcknowledgePacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerAcknowledgePacket represents a AcknowledgePacket event raised by the Ibchandler contract.
type IbchandlerAcknowledgePacket struct {
	Packet          PacketData
	Acknowledgement []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgePacket is a free log retrieval operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) FilterAcknowledgePacket(opts *bind.FilterOpts) (*IbchandlerAcknowledgePacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "AcknowledgePacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerAcknowledgePacketIterator{contract: _Ibchandler.contract, event: "AcknowledgePacket", logs: logs, sub: sub}, nil
}

// WatchAcknowledgePacket is a free log subscription operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) WatchAcknowledgePacket(opts *bind.WatchOpts, sink chan<- *IbchandlerAcknowledgePacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "AcknowledgePacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerAcknowledgePacket)
				if err := _Ibchandler.contract.UnpackLog(event, "AcknowledgePacket", log); err != nil {
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

// ParseAcknowledgePacket is a log parse operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) ParseAcknowledgePacket(log types.Log) (*IbchandlerAcknowledgePacket, error) {
	event := new(IbchandlerAcknowledgePacket)
	if err := _Ibchandler.contract.UnpackLog(event, "AcknowledgePacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerRecvPacketIterator is returned from FilterRecvPacket and is used to iterate over the raw logs and unpacked data for RecvPacket events raised by the Ibchandler contract.
type IbchandlerRecvPacketIterator struct {
	Event *IbchandlerRecvPacket // Event containing the contract specifics and raw log

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
func (it *IbchandlerRecvPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerRecvPacket)
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
		it.Event = new(IbchandlerRecvPacket)
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
func (it *IbchandlerRecvPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerRecvPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerRecvPacket represents a RecvPacket event raised by the Ibchandler contract.
type IbchandlerRecvPacket struct {
	Packet          PacketData
	Acknowledgement []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRecvPacket is a free log retrieval operation binding the contract event 0x5ba247606c41195940027cd60b1ff8d3c6567b0d5b8fa3baa205b671d1c4c8f8.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) FilterRecvPacket(opts *bind.FilterOpts) (*IbchandlerRecvPacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "RecvPacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerRecvPacketIterator{contract: _Ibchandler.contract, event: "RecvPacket", logs: logs, sub: sub}, nil
}

// WatchRecvPacket is a free log subscription operation binding the contract event 0x5ba247606c41195940027cd60b1ff8d3c6567b0d5b8fa3baa205b671d1c4c8f8.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) WatchRecvPacket(opts *bind.WatchOpts, sink chan<- *IbchandlerRecvPacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "RecvPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerRecvPacket)
				if err := _Ibchandler.contract.UnpackLog(event, "RecvPacket", log); err != nil {
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

// ParseRecvPacket is a log parse operation binding the contract event 0x5ba247606c41195940027cd60b1ff8d3c6567b0d5b8fa3baa205b671d1c4c8f8.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) ParseRecvPacket(log types.Log) (*IbchandlerRecvPacket, error) {
	event := new(IbchandlerRecvPacket)
	if err := _Ibchandler.contract.UnpackLog(event, "RecvPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the Ibchandler contract.
type IbchandlerSendPacketIterator struct {
	Event *IbchandlerSendPacket // Event containing the contract specifics and raw log

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
func (it *IbchandlerSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerSendPacket)
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
		it.Event = new(IbchandlerSendPacket)
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
func (it *IbchandlerSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerSendPacket represents a SendPacket event raised by the Ibchandler contract.
type IbchandlerSendPacket struct {
	Packet PacketData
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendPacket is a free log retrieval operation binding the contract event 0xe701f25bda8992b211749f81adb9a8ea6e8cf8a3c9f2e29ed496e6c5f059154c.
//
// Solidity: event SendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) FilterSendPacket(opts *bind.FilterOpts) (*IbchandlerSendPacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "SendPacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerSendPacketIterator{contract: _Ibchandler.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0xe701f25bda8992b211749f81adb9a8ea6e8cf8a3c9f2e29ed496e6c5f059154c.
//
// Solidity: event SendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *IbchandlerSendPacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "SendPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerSendPacket)
				if err := _Ibchandler.contract.UnpackLog(event, "SendPacket", log); err != nil {
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

// ParseSendPacket is a log parse operation binding the contract event 0xe701f25bda8992b211749f81adb9a8ea6e8cf8a3c9f2e29ed496e6c5f059154c.
//
// Solidity: event SendPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) ParseSendPacket(log types.Log) (*IbchandlerSendPacket, error) {
	event := new(IbchandlerSendPacket)
	if err := _Ibchandler.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerWriteAcknowledgementIterator is returned from FilterWriteAcknowledgement and is used to iterate over the raw logs and unpacked data for WriteAcknowledgement events raised by the Ibchandler contract.
type IbchandlerWriteAcknowledgementIterator struct {
	Event *IbchandlerWriteAcknowledgement // Event containing the contract specifics and raw log

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
func (it *IbchandlerWriteAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerWriteAcknowledgement)
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
		it.Event = new(IbchandlerWriteAcknowledgement)
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
func (it *IbchandlerWriteAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerWriteAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerWriteAcknowledgement represents a WriteAcknowledgement event raised by the Ibchandler contract.
type IbchandlerWriteAcknowledgement struct {
	DestinationPortId  string
	DestinationChannel string
	Sequence           uint64
	Acknowledgement    []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWriteAcknowledgement is a free log retrieval operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) FilterWriteAcknowledgement(opts *bind.FilterOpts) (*IbchandlerWriteAcknowledgementIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "WriteAcknowledgement")
	if err != nil {
		return nil, err
	}
	return &IbchandlerWriteAcknowledgementIterator{contract: _Ibchandler.contract, event: "WriteAcknowledgement", logs: logs, sub: sub}, nil
}

// WatchWriteAcknowledgement is a free log subscription operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) WatchWriteAcknowledgement(opts *bind.WatchOpts, sink chan<- *IbchandlerWriteAcknowledgement) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "WriteAcknowledgement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerWriteAcknowledgement)
				if err := _Ibchandler.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
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

// ParseWriteAcknowledgement is a log parse operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) ParseWriteAcknowledgement(log types.Log) (*IbchandlerWriteAcknowledgement, error) {
	event := new(IbchandlerWriteAcknowledgement)
	if err := _Ibchandler.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
