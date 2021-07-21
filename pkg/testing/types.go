package testing

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/channel"
	connectiontypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/connection"
)

// TestConnection is a testing helper struct to keep track of the connectionID, source clientID,
// counterparty clientID, and the next channel version used in creating and interacting with a
// connection.
type TestConnection struct {
	ID                   string
	ClientID             string
	CounterpartyClientID string
	NextChannelVersion   string
	Channels             []TestChannel
}

// TestChannel is a testing helper struct to keep track of the portID and channelID
// used in creating and interacting with a channel. The clientID and counterparty
// client ID are also tracked to cut down on querying and argument passing.
type TestChannel struct {
	PortID               string
	ID                   string
	ClientID             string
	CounterpartyClientID string
	Version              string
}

func connectionEndToPB(conn ibchost.ConnectionEndData) *connectiontypes.ConnectionEnd {
	connpb := &connectiontypes.ConnectionEnd{
		ClientId:    conn.ClientId,
		Versions:    []*connectiontypes.Version{},
		State:       connectiontypes.ConnectionEnd_State(conn.State),
		DelayPeriod: conn.DelayPeriod,
		Counterparty: &connectiontypes.Counterparty{
			ClientId:     conn.Counterparty.ClientId,
			ConnectionId: conn.Counterparty.ConnectionId,
			Prefix:       (*connectiontypes.MerklePrefix)(&conn.Counterparty.Prefix),
		},
	}
	for _, v := range conn.Versions {
		ver := connectiontypes.Version(v)
		connpb.Versions = append(connpb.Versions, &ver)
	}
	return connpb
}

func channelToPB(ch ibchost.ChannelData) *channeltypes.Channel {
	return &channeltypes.Channel{
		State:          channeltypes.Channel_State(ch.State),
		Ordering:       channeltypes.Channel_Order(ch.Ordering),
		Counterparty:   channeltypes.Channel_Counterparty(ch.Counterparty),
		ConnectionHops: ch.ConnectionHops,
		Version:        ch.Version,
	}
}

// uint64ToBigEndian - marshals uint64 to a bigendian byte slice so it can be sorted
func uint64ToBigEndian(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// commitPacket returns the packet commitment bytes. The commitment consists of:
// sha256_hash(timeout_timestamp + timeout_height.RevisionNumber + timeout_height.RevisionHeight + sha256_hash(data))
// from a given packet. This results in a fixed length preimage.
// NOTE: uint64ToBigEndian sets the uint64 to a slice of length 8.
func commitPacket(packet channeltypes.Packet) []byte {
	timeoutHeight := packet.TimeoutHeight

	buf := uint64ToBigEndian(packet.TimeoutTimestamp)

	revisionNumber := uint64ToBigEndian(timeoutHeight.GetRevisionNumber())
	buf = append(buf, revisionNumber...)

	revisionHeight := uint64ToBigEndian(timeoutHeight.GetRevisionHeight())
	buf = append(buf, revisionHeight...)

	dataHash := sha256.Sum256(packet.Data)
	buf = append(buf, dataHash[:]...)

	hash := sha256.Sum256(buf)
	return hash[:]
}

// commitAcknowledgement returns the hash of commitment bytes
func commitAcknowledgement(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func PackAny(msg proto.Message) (*types.Any, error) {
	var any types.Any
	any.TypeUrl = "/" + proto.MessageName(msg)

	bz, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	any.Value = bz
	return &any, nil
}

func UnpackAny(bz []byte) (*types.Any, error) {
	var any types.Any
	if err := proto.Unmarshal(bz, &any); err != nil {
		return nil, err
	}
	return &any, nil
}

func MarshalWithAny(msg proto.Message) ([]byte, error) {
	any, err := PackAny(msg)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(any)
}

func UnmarshalWithAny(bz []byte, msg proto.Message) error {
	any, err := UnpackAny(bz)
	if err != nil {
		return err
	}
	if t := "/" + proto.MessageName(msg); any.TypeUrl != t {
		return fmt.Errorf("expected %v, but got %v", t, any.TypeUrl)
	}
	return proto.Unmarshal(any.Value, msg)
}
