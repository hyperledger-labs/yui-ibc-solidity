// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcstore

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

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	ChainId         string
	IbcStoreAddress []byte
	LatestHeight    uint64
}

// ConnectionEndData is an auto generated low-level Go binding around an user-defined struct.
type ConnectionEndData struct {
	ClientId     string
	Versions     []VersionData
	State        uint8
	DelayPeriod  uint64
	Counterparty CounterpartyData
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Timestamp  uint64
	Root       []byte
	Validators [][]byte
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

// IbcstoreABI is the input ABI used to generate the binding from.
const IbcstoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"consensusCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"consensusStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"setClientState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"ibc_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"hasClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"}],\"name\":\"setConsensusState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"}],\"name\":\"setConnection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getConnection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"name\":\"setChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getChannel\",\"outputs\":[{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"hasChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceRecv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceRecv\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getNextSequenceAck\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"getPacket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"setPacketCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"deletePacketCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"getPacketCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"makePacketCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"setPacketAcknowledgementCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"getPacketAcknowledgementCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"makePacketAcknowledgementCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setPacketReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"hasPacketReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientStateBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getConnectionBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"parseConnectionBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Ibcstore is an auto generated Go binding around an Ethereum contract.
type Ibcstore struct {
	IbcstoreCaller     // Read-only binding to the contract
	IbcstoreTransactor // Write-only binding to the contract
	IbcstoreFilterer   // Log filterer for contract events
}

// IbcstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcstoreSession struct {
	Contract     *Ibcstore         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcstoreCallerSession struct {
	Contract *IbcstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IbcstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcstoreTransactorSession struct {
	Contract     *IbcstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IbcstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcstoreRaw struct {
	Contract *Ibcstore // Generic contract binding to access the raw methods on
}

// IbcstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcstoreCallerRaw struct {
	Contract *IbcstoreCaller // Generic read-only contract binding to access the raw methods on
}

// IbcstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcstoreTransactorRaw struct {
	Contract *IbcstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcstore creates a new instance of Ibcstore, bound to a specific deployed contract.
func NewIbcstore(address common.Address, backend bind.ContractBackend) (*Ibcstore, error) {
	contract, err := bindIbcstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcstore{IbcstoreCaller: IbcstoreCaller{contract: contract}, IbcstoreTransactor: IbcstoreTransactor{contract: contract}, IbcstoreFilterer: IbcstoreFilterer{contract: contract}}, nil
}

// NewIbcstoreCaller creates a new read-only instance of Ibcstore, bound to a specific deployed contract.
func NewIbcstoreCaller(address common.Address, caller bind.ContractCaller) (*IbcstoreCaller, error) {
	contract, err := bindIbcstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcstoreCaller{contract: contract}, nil
}

// NewIbcstoreTransactor creates a new write-only instance of Ibcstore, bound to a specific deployed contract.
func NewIbcstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcstoreTransactor, error) {
	contract, err := bindIbcstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcstoreTransactor{contract: contract}, nil
}

// NewIbcstoreFilterer creates a new log filterer instance of Ibcstore, bound to a specific deployed contract.
func NewIbcstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcstoreFilterer, error) {
	contract, err := bindIbcstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcstoreFilterer{contract: contract}, nil
}

// bindIbcstore binds a generic wrapper to an already deployed contract.
func bindIbcstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcstoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcstore *IbcstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcstore.Contract.IbcstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcstore *IbcstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcstore.Contract.IbcstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcstore *IbcstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcstore.Contract.IbcstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcstore *IbcstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcstore *IbcstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcstore *IbcstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcstore.Contract.contract.Transact(opts, method, params...)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ChannelCommitmentKey(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "channelCommitmentKey", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibcstore.Contract.ChannelCommitmentKey(&_Ibcstore.CallOpts, portId, channelId)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ChannelCommitmentKey(portId string, channelId string) ([32]byte, error) {
	return _Ibcstore.Contract.ChannelCommitmentKey(&_Ibcstore.CallOpts, portId, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ChannelCommitmentSlot(opts *bind.CallOpts, portId string, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "channelCommitmentSlot", portId, channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ChannelCommitmentSlot(portId string, channelId string) ([32]byte, error) {
	return _Ibcstore.Contract.ChannelCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ChannelCommitmentSlot(portId string, channelId string) ([32]byte, error) {
	return _Ibcstore.Contract.ChannelCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ClientCommitmentKey(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "clientCommitmentKey", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibcstore.Contract.ClientCommitmentKey(&_Ibcstore.CallOpts, clientId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Ibcstore.Contract.ClientCommitmentKey(&_Ibcstore.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ClientStateCommitmentSlot(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "clientStateCommitmentSlot", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Ibcstore.Contract.ClientStateCommitmentSlot(&_Ibcstore.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Ibcstore.Contract.ClientStateCommitmentSlot(&_Ibcstore.CallOpts, clientId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ConnectionCommitmentKey(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "connectionCommitmentKey", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.ConnectionCommitmentKey(&_Ibcstore.CallOpts, connectionId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.ConnectionCommitmentKey(&_Ibcstore.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ConnectionCommitmentSlot(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "connectionCommitmentSlot", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.ConnectionCommitmentSlot(&_Ibcstore.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.ConnectionCommitmentSlot(&_Ibcstore.CallOpts, connectionId)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ConsensusCommitmentKey(opts *bind.CallOpts, clientId string, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "consensusCommitmentKey", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ConsensusCommitmentKey(clientId string, height uint64) ([32]byte, error) {
	return _Ibcstore.Contract.ConsensusCommitmentKey(&_Ibcstore.CallOpts, clientId, height)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ConsensusCommitmentKey(clientId string, height uint64) ([32]byte, error) {
	return _Ibcstore.Contract.ConsensusCommitmentKey(&_Ibcstore.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) ConsensusStateCommitmentSlot(opts *bind.CallOpts, clientId string, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "consensusStateCommitmentSlot", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) ConsensusStateCommitmentSlot(clientId string, height uint64) ([32]byte, error) {
	return _Ibcstore.Contract.ConsensusStateCommitmentSlot(&_Ibcstore.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) ConsensusStateCommitmentSlot(clientId string, height uint64) ([32]byte, error) {
	return _Ibcstore.Contract.ConsensusStateCommitmentSlot(&_Ibcstore.CallOpts, clientId, height)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Ibcstore *IbcstoreCaller) GetChannel(opts *bind.CallOpts, portId string, channelId string) (ChannelData, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getChannel", portId, channelId)

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
func (_Ibcstore *IbcstoreSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibcstore.Contract.GetChannel(&_Ibcstore.CallOpts, portId, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Ibcstore *IbcstoreCallerSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibcstore.Contract.GetChannel(&_Ibcstore.CallOpts, portId, channelId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Ibcstore *IbcstoreCaller) GetClientState(opts *bind.CallOpts, clientId string) (ClientStateData, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new(ClientStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientStateData)).(*ClientStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Ibcstore *IbcstoreSession) GetClientState(clientId string) (ClientStateData, bool, error) {
	return _Ibcstore.Contract.GetClientState(&_Ibcstore.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Ibcstore *IbcstoreCallerSession) GetClientState(clientId string) (ClientStateData, bool, error) {
	return _Ibcstore.Contract.GetClientState(&_Ibcstore.CallOpts, clientId)
}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreCaller) GetClientStateBytes(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getClientStateBytes", clientId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreSession) GetClientStateBytes(clientId string) ([]byte, bool, error) {
	return _Ibcstore.Contract.GetClientStateBytes(&_Ibcstore.CallOpts, clientId)
}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreCallerSession) GetClientStateBytes(clientId string) ([]byte, bool, error) {
	return _Ibcstore.Contract.GetClientStateBytes(&_Ibcstore.CallOpts, clientId)
}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Ibcstore *IbcstoreCaller) GetCommitment(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getCommitment", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Ibcstore *IbcstoreSession) GetCommitment(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.GetCommitment(&_Ibcstore.CallOpts, connectionId)
}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) GetCommitment(connectionId string) ([32]byte, error) {
	return _Ibcstore.Contract.GetCommitment(&_Ibcstore.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Ibcstore *IbcstoreCaller) GetConnection(opts *bind.CallOpts, connectionId string) (ConnectionEndData, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getConnection", connectionId)

	if err != nil {
		return *new(ConnectionEndData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Ibcstore *IbcstoreSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibcstore.Contract.GetConnection(&_Ibcstore.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Ibcstore *IbcstoreCallerSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibcstore.Contract.GetConnection(&_Ibcstore.CallOpts, connectionId)
}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreCaller) GetConnectionBytes(opts *bind.CallOpts, connectionId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getConnectionBytes", connectionId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreSession) GetConnectionBytes(connectionId string) ([]byte, bool, error) {
	return _Ibcstore.Contract.GetConnectionBytes(&_Ibcstore.CallOpts, connectionId)
}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Ibcstore *IbcstoreCallerSession) GetConnectionBytes(connectionId string) ([]byte, bool, error) {
	return _Ibcstore.Contract.GetConnectionBytes(&_Ibcstore.CallOpts, connectionId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Ibcstore *IbcstoreCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height uint64) (ConsensusStateData, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new(ConsensusStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConsensusStateData)).(*ConsensusStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Ibcstore *IbcstoreSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, bool, error) {
	return _Ibcstore.Contract.GetConsensusState(&_Ibcstore.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Ibcstore *IbcstoreCallerSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, bool, error) {
	return _Ibcstore.Contract.GetConsensusState(&_Ibcstore.CallOpts, clientId, height)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCaller) GetNextSequenceAck(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getNextSequenceAck", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceAck(&_Ibcstore.CallOpts, portId, channelId)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCallerSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceAck(&_Ibcstore.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCaller) GetNextSequenceRecv(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getNextSequenceRecv", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceRecv(&_Ibcstore.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCallerSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceRecv(&_Ibcstore.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCaller) GetNextSequenceSend(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getNextSequenceSend", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceSend(&_Ibcstore.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibcstore *IbcstoreCallerSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibcstore.Contract.GetNextSequenceSend(&_Ibcstore.CallOpts, portId, channelId)
}

// GetPacket is a free data retrieval call binding the contract method 0x162e3377.
//
// Solidity: function getPacket(string portId, string channelId, uint64 sequence) view returns((uint64,string,string,string,string,bytes,(uint64,uint64),uint64))
func (_Ibcstore *IbcstoreCaller) GetPacket(opts *bind.CallOpts, portId string, channelId string, sequence uint64) (PacketData, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getPacket", portId, channelId, sequence)

	if err != nil {
		return *new(PacketData), err
	}

	out0 := *abi.ConvertType(out[0], new(PacketData)).(*PacketData)

	return out0, err

}

// GetPacket is a free data retrieval call binding the contract method 0x162e3377.
//
// Solidity: function getPacket(string portId, string channelId, uint64 sequence) view returns((uint64,string,string,string,string,bytes,(uint64,uint64),uint64))
func (_Ibcstore *IbcstoreSession) GetPacket(portId string, channelId string, sequence uint64) (PacketData, error) {
	return _Ibcstore.Contract.GetPacket(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// GetPacket is a free data retrieval call binding the contract method 0x162e3377.
//
// Solidity: function getPacket(string portId, string channelId, uint64 sequence) view returns((uint64,string,string,string,string,bytes,(uint64,uint64),uint64))
func (_Ibcstore *IbcstoreCallerSession) GetPacket(portId string, channelId string, sequence uint64) (PacketData, error) {
	return _Ibcstore.Contract.GetPacket(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// GetPacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x71f56c59.
//
// Solidity: function getPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibcstore *IbcstoreCaller) GetPacketAcknowledgementCommitment(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "getPacketAcknowledgementCommitment", portId, channelId, sequence)

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
func (_Ibcstore *IbcstoreSession) GetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibcstore.Contract.GetPacketAcknowledgementCommitment(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// GetPacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x71f56c59.
//
// Solidity: function getPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence) view returns(bytes32, bool)
func (_Ibcstore *IbcstoreCallerSession) GetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64) ([32]byte, bool, error) {
	return _Ibcstore.Contract.GetPacketAcknowledgementCommitment(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// HasChannel is a free data retrieval call binding the contract method 0x4381fc29.
//
// Solidity: function hasChannel(string portId, string channelId) view returns(bool)
func (_Ibcstore *IbcstoreCaller) HasChannel(opts *bind.CallOpts, portId string, channelId string) (bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "hasChannel", portId, channelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasChannel is a free data retrieval call binding the contract method 0x4381fc29.
//
// Solidity: function hasChannel(string portId, string channelId) view returns(bool)
func (_Ibcstore *IbcstoreSession) HasChannel(portId string, channelId string) (bool, error) {
	return _Ibcstore.Contract.HasChannel(&_Ibcstore.CallOpts, portId, channelId)
}

// HasChannel is a free data retrieval call binding the contract method 0x4381fc29.
//
// Solidity: function hasChannel(string portId, string channelId) view returns(bool)
func (_Ibcstore *IbcstoreCallerSession) HasChannel(portId string, channelId string) (bool, error) {
	return _Ibcstore.Contract.HasChannel(&_Ibcstore.CallOpts, portId, channelId)
}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Ibcstore *IbcstoreCaller) HasClientState(opts *bind.CallOpts, clientId string) (bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "hasClientState", clientId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Ibcstore *IbcstoreSession) HasClientState(clientId string) (bool, error) {
	return _Ibcstore.Contract.HasClientState(&_Ibcstore.CallOpts, clientId)
}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Ibcstore *IbcstoreCallerSession) HasClientState(clientId string) (bool, error) {
	return _Ibcstore.Contract.HasClientState(&_Ibcstore.CallOpts, clientId)
}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibcstore *IbcstoreCaller) HasPacketReceipt(opts *bind.CallOpts, portId string, channelId string, sequence uint64) (bool, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "hasPacketReceipt", portId, channelId, sequence)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibcstore *IbcstoreSession) HasPacketReceipt(portId string, channelId string, sequence uint64) (bool, error) {
	return _Ibcstore.Contract.HasPacketReceipt(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// HasPacketReceipt is a free data retrieval call binding the contract method 0x5a9afac3.
//
// Solidity: function hasPacketReceipt(string portId, string channelId, uint64 sequence) view returns(bool)
func (_Ibcstore *IbcstoreCallerSession) HasPacketReceipt(portId string, channelId string, sequence uint64) (bool, error) {
	return _Ibcstore.Contract.HasPacketReceipt(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) view returns(bytes32)
func (_Ibcstore *IbcstoreCaller) MakePacketAcknowledgementCommitment(opts *bind.CallOpts, acknowledgement []byte) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "makePacketAcknowledgementCommitment", acknowledgement)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) view returns(bytes32)
func (_Ibcstore *IbcstoreSession) MakePacketAcknowledgementCommitment(acknowledgement []byte) ([32]byte, error) {
	return _Ibcstore.Contract.MakePacketAcknowledgementCommitment(&_Ibcstore.CallOpts, acknowledgement)
}

// MakePacketAcknowledgementCommitment is a free data retrieval call binding the contract method 0x8a889658.
//
// Solidity: function makePacketAcknowledgementCommitment(bytes acknowledgement) view returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) MakePacketAcknowledgementCommitment(acknowledgement []byte) ([32]byte, error) {
	return _Ibcstore.Contract.MakePacketAcknowledgementCommitment(&_Ibcstore.CallOpts, acknowledgement)
}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) view returns(bytes32)
func (_Ibcstore *IbcstoreCaller) MakePacketCommitment(opts *bind.CallOpts, packet PacketData) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "makePacketCommitment", packet)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) view returns(bytes32)
func (_Ibcstore *IbcstoreSession) MakePacketCommitment(packet PacketData) ([32]byte, error) {
	return _Ibcstore.Contract.MakePacketCommitment(&_Ibcstore.CallOpts, packet)
}

// MakePacketCommitment is a free data retrieval call binding the contract method 0x12a68750.
//
// Solidity: function makePacketCommitment((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) view returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) MakePacketCommitment(packet PacketData) ([32]byte, error) {
	return _Ibcstore.Contract.MakePacketCommitment(&_Ibcstore.CallOpts, packet)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) PacketAcknowledgementCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "packetAcknowledgementCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketAcknowledgementCommitmentKey(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentKey is a free data retrieval call binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketAcknowledgementCommitmentKey(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) PacketAcknowledgementCommitmentSlot(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "packetAcknowledgementCommitmentSlot", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a free data retrieval call binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) PacketCommitmentKey(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "packetCommitmentKey", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketCommitmentKey(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a free data retrieval call binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketCommitmentKey(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCaller) PacketCommitmentSlot(opts *bind.CallOpts, portId string, channelId string, sequence uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "packetCommitmentSlot", portId, channelId, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a free data retrieval call binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) pure returns(bytes32)
func (_Ibcstore *IbcstoreCallerSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) ([32]byte, error) {
	return _Ibcstore.Contract.PacketCommitmentSlot(&_Ibcstore.CallOpts, portId, channelId, sequence)
}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Ibcstore *IbcstoreCaller) ParseConnectionBytes(opts *bind.CallOpts, connectionBytes []byte) (ConnectionEndData, error) {
	var out []interface{}
	err := _Ibcstore.contract.Call(opts, &out, "parseConnectionBytes", connectionBytes)

	if err != nil {
		return *new(ConnectionEndData), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)

	return out0, err

}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Ibcstore *IbcstoreSession) ParseConnectionBytes(connectionBytes []byte) (ConnectionEndData, error) {
	return _Ibcstore.Contract.ParseConnectionBytes(&_Ibcstore.CallOpts, connectionBytes)
}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Ibcstore *IbcstoreCallerSession) ParseConnectionBytes(connectionBytes []byte) (ConnectionEndData, error) {
	return _Ibcstore.Contract.ParseConnectionBytes(&_Ibcstore.CallOpts, connectionBytes)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactor) DeletePacketCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "deletePacketCommitment", portId, channelId, sequence)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreSession) DeletePacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.DeletePacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// DeletePacketCommitment is a paid mutator transaction binding the contract method 0x94dcb4c6.
//
// Solidity: function deletePacketCommitment(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactorSession) DeletePacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.DeletePacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// GetPacketCommitment is a paid mutator transaction binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) returns(bytes32, bool)
func (_Ibcstore *IbcstoreTransactor) GetPacketCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "getPacketCommitment", portId, channelId, sequence)
}

// GetPacketCommitment is a paid mutator transaction binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) returns(bytes32, bool)
func (_Ibcstore *IbcstoreSession) GetPacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.GetPacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// GetPacketCommitment is a paid mutator transaction binding the contract method 0x61fc5e7b.
//
// Solidity: function getPacketCommitment(string portId, string channelId, uint64 sequence) returns(bytes32, bool)
func (_Ibcstore *IbcstoreTransactorSession) GetPacketCommitment(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.GetPacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibcstore *IbcstoreTransactor) SetChannel(opts *bind.TransactOpts, portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setChannel", portId, channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibcstore *IbcstoreSession) SetChannel(portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetChannel(&_Ibcstore.TransactOpts, portId, channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x3564d550.
//
// Solidity: function setChannel(string portId, string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetChannel(portId string, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetChannel(&_Ibcstore.TransactOpts, portId, channelId, channel)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Ibcstore *IbcstoreTransactor) SetClientState(opts *bind.TransactOpts, clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setClientState", clientId, data)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Ibcstore *IbcstoreSession) SetClientState(clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetClientState(&_Ibcstore.TransactOpts, clientId, data)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetClientState(clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetClientState(&_Ibcstore.TransactOpts, clientId, data)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Ibcstore *IbcstoreTransactor) SetConnection(opts *bind.TransactOpts, connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setConnection", connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Ibcstore *IbcstoreSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetConnection(&_Ibcstore.TransactOpts, connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetConnection(&_Ibcstore.TransactOpts, connectionId, connection)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcstore *IbcstoreTransactor) SetConsensusState(opts *bind.TransactOpts, clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setConsensusState", clientId, height, consensusState)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcstore *IbcstoreSession) SetConsensusState(clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetConsensusState(&_Ibcstore.TransactOpts, clientId, height, consensusState)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetConsensusState(clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetConsensusState(&_Ibcstore.TransactOpts, clientId, height, consensusState)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactor) SetNextSequenceAck(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setNextSequenceAck", portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceAck(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceAck(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactor) SetNextSequenceRecv(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setNextSequenceRecv", portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceRecv(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceRecv(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactor) SetNextSequenceSend(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setNextSequenceSend", portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceSend(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetNextSequenceSend(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibcstore *IbcstoreTransactor) SetPacketAcknowledgementCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setPacketAcknowledgementCommitment", portId, channelId, sequence, acknowledgement)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibcstore *IbcstoreSession) SetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketAcknowledgementCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence, acknowledgement)
}

// SetPacketAcknowledgementCommitment is a paid mutator transaction binding the contract method 0xf81cc9e1.
//
// Solidity: function setPacketAcknowledgementCommitment(string portId, string channelId, uint64 sequence, bytes acknowledgement) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetPacketAcknowledgementCommitment(portId string, channelId string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketAcknowledgementCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence, acknowledgement)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcstore *IbcstoreTransactor) SetPacketCommitment(opts *bind.TransactOpts, portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setPacketCommitment", portId, channelId, sequence, packet)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcstore *IbcstoreSession) SetPacketCommitment(portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence, packet)
}

// SetPacketCommitment is a paid mutator transaction binding the contract method 0xb7ccdc57.
//
// Solidity: function setPacketCommitment(string portId, string channelId, uint64 sequence, (uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetPacketCommitment(portId string, channelId string, sequence uint64, packet PacketData) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketCommitment(&_Ibcstore.TransactOpts, portId, channelId, sequence, packet)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactor) SetPacketReceipt(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.contract.Transact(opts, "setPacketReceipt", portId, channelId, sequence)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreSession) SetPacketReceipt(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketReceipt(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}

// SetPacketReceipt is a paid mutator transaction binding the contract method 0xf2a47da3.
//
// Solidity: function setPacketReceipt(string portId, string channelId, uint64 sequence) returns()
func (_Ibcstore *IbcstoreTransactorSession) SetPacketReceipt(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcstore.Contract.SetPacketReceipt(&_Ibcstore.TransactOpts, portId, channelId, sequence)
}
