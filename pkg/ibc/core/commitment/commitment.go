package commitment

import (
	"encoding/hex"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/ethereum/go-ethereum/crypto"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/client"
)

var IBCCommitmentsSlot [32]byte

func init() {
	ibcCommitmentsSlotStr := "1ee222554989dda120e26ecacf756fe1235cd8d726706b57517715dde4f0c900"
	bz, err := hex.DecodeString(ibcCommitmentsSlotStr)
	if err != nil {
		panic(err)
	}
	copy(IBCCommitmentsSlot[:], bz)
}

// Slot calculator

func ClientStateCommitmentSlot(clientID string) string {
	return CalculateCommitmentSlot(host.FullClientStateKey(clientID))
}

func ConsensusStateCommitmentSlot(clientID string, height ibcclient.Height) string {
	return CalculateCommitmentSlot(host.FullConsensusStateKey(clientID, clienttypes.NewHeight(height.RevisionNumber, height.RevisionHeight)))
}

func ConnectionStateCommitmentSlot(connectionID string) string {
	return CalculateCommitmentSlot(host.ConnectionKey(connectionID))
}

func ChannelStateCommitmentSlot(portID, channelID string) string {
	return CalculateCommitmentSlot(host.ChannelKey(portID, channelID))
}

func PacketCommitmentSlot(portID, channelID string, sequence uint64) string {
	return CalculateCommitmentSlot(host.PacketCommitmentKey(portID, channelID, sequence))
}

func PacketAcknowledgementCommitmentSlot(portID, channelID string, sequence uint64) string {
	return CalculateCommitmentSlot(host.PacketAcknowledgementKey(portID, channelID, sequence))
}

func PacketReceiptCommitmentSlot(portID, channelID string, sequence uint64) string {
	return CalculateCommitmentSlot(host.PacketReceiptKey(portID, channelID, sequence))
}

func NextSequenceRecvCommitmentSlot(portID, channelID string) string {
	return CalculateCommitmentSlot(host.NextSequenceRecvKey(portID, channelID))
}

func CalculateCommitmentSlot(path []byte) string {
	return crypto.Keccak256Hash(crypto.Keccak256Hash(path).Bytes(), IBCCommitmentsSlot[:]).Hex()
}
