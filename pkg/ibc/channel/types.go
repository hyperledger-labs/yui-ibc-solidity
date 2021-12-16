package channel

import client "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"

// NewPacket creates a new Packet instance.
func NewPacket(
	data []byte,
	sequence uint64, sourcePort, sourceChannel,
	destinationPort, destinationChannel string,
	timeoutHeight client.Height, timeoutTimestamp uint64,
) Packet {
	return Packet{
		Data:               data,
		Sequence:           sequence,
		SourcePort:         sourcePort,
		SourceChannel:      sourceChannel,
		DestinationPort:    destinationPort,
		DestinationChannel: destinationChannel,
		TimeoutHeight:      timeoutHeight,
		TimeoutTimestamp:   timeoutTimestamp,
	}
}
