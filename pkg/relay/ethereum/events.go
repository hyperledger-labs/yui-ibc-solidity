package ethereum

import (
	"fmt"
	"math/big"
	"strings"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var (
	abiSendPacket,
	abiWriteAcknowledgement abi.Event
)

func init() {
	parsedHandlerABI, err := abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		panic(err)
	}
	abiSendPacket = parsedHandlerABI.Events["SendPacket"]
	abiWriteAcknowledgement = parsedHandlerABI.Events["WriteAcknowledgement"]
}

func (chain *Chain) findPacket(
	ctx core.QueryContext,
	sourcePortID string,
	sourceChannel string,
	sequence uint64,
) (*chantypes.Packet, error) {
	channel, found, err := chain.ibcHandler.GetChannel(
		chain.callOptsFromQueryContext(ctx),
		sourcePortID, sourceChannel,
	)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("channel not found: sourcePortID=%v sourceChannel=%v", sourcePortID, sourceChannel)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   new(big.Int).SetUint64(ctx.Height().GetRevisionHeight()),
		Addresses: []common.Address{
			chain.config.IBCAddress(),
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	for _, log := range logs {
		if values, err := abiSendPacket.Inputs.Unpack(log.Data); err != nil {
			return nil, err
		} else {
			if l := len(values); l != 6 {
				return nil, fmt.Errorf("unexpected values length: expected=%v actual=%v", 6, l)
			}
			pSequence := values[0].(uint64)
			pSourcePortID := values[1].(string)
			pSourceChannel := values[2].(string)
			pTimeoutHeight := values[3].(struct {
				RevisionNumber uint64 "json:\"revision_number\""
				RevisionHeight uint64 "json:\"revision_height\""
			})
			pTimeoutTimestamp := values[4].(uint64)
			pData := values[5].([]uint8)

			if pSequence == sequence && pSourcePortID == sourcePortID && pSourceChannel == sourceChannel {
				return &channeltypes.Packet{
					Sequence:           pSequence,
					SourcePort:         pSourcePortID,
					SourceChannel:      pSourceChannel,
					DestinationPort:    channel.Counterparty.PortId,
					DestinationChannel: channel.Counterparty.ChannelId,
					Data:               pData,
					TimeoutHeight:      clienttypes.Height{RevisionNumber: pTimeoutHeight.RevisionNumber, RevisionHeight: pTimeoutHeight.RevisionHeight},
					TimeoutTimestamp:   pTimeoutTimestamp,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("packet not found: sourcePortID=%v sourceChannel=%v sequence=%v", sourcePortID, sourceChannel, sequence)
}

// getAllPackets returns all packets from events
func (chain *Chain) getAllPackets(
	ctx core.QueryContext,
	sourcePortID string,
	sourceChannel string,
) ([]*chantypes.Packet, error) {
	channel, found, err := chain.ibcHandler.GetChannel(
		chain.callOptsFromQueryContext(ctx),
		sourcePortID, sourceChannel,
	)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("channel not found: sourcePortID=%v sourceChannel=%v", sourcePortID, sourceChannel)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   new(big.Int).SetUint64(ctx.Height().GetRevisionHeight()),
		Addresses: []common.Address{
			chain.config.IBCAddress(),
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	var packets []*chantypes.Packet
	for _, log := range logs {
		if values, err := abiSendPacket.Inputs.Unpack(log.Data); err != nil {
			return nil, err
		} else {
			if l := len(values); l != 6 {
				return nil, fmt.Errorf("unexpected values length: expected=%v actual=%v", 6, l)
			}
			pSequence := values[0].(uint64)
			pSourcePortID := values[1].(string)
			pSourceChannel := values[2].(string)
			pTimeoutHeight := values[3].(struct {
				RevisionNumber uint64 "json:\"revision_number\""
				RevisionHeight uint64 "json:\"revision_height\""
			})
			pTimeoutTimestamp := values[4].(uint64)
			pData := values[5].([]uint8)

			if pSourcePortID == sourcePortID && pSourceChannel == sourceChannel {
				packets = append(packets, &channeltypes.Packet{
					Sequence:           pSequence,
					SourcePort:         pSourcePortID,
					SourceChannel:      pSourceChannel,
					DestinationPort:    channel.Counterparty.PortId,
					DestinationChannel: channel.Counterparty.ChannelId,
					Data:               pData,
					TimeoutHeight:      clienttypes.Height{RevisionNumber: pTimeoutHeight.RevisionNumber, RevisionHeight: pTimeoutHeight.RevisionHeight},
					TimeoutTimestamp:   pTimeoutTimestamp,
				})
			}
		}
	}
	return packets, nil
}

func (chain *Chain) findAcknowledgement(
	ctx core.QueryContext,
	dstPortID string,
	dstChannel string,
	sequence uint64,
) ([]byte, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   new(big.Int).SetUint64(ctx.Height().GetRevisionHeight()),
		Addresses: []common.Address{
			chain.config.IBCAddress(),
		},
		Topics: [][]common.Hash{{
			abiWriteAcknowledgement.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	for _, log := range logs {
		if values, err := abiWriteAcknowledgement.Inputs.Unpack(log.Data); err != nil {
			return nil, err
		} else {
			if len(values) != 4 {
				return nil, fmt.Errorf("unexpected values: %v", values)
			}
			if dstPortID == values[0].(string) && dstChannel == values[1].(string) && sequence == values[2].(uint64) {
				return values[3].([]byte), nil
			}
		}
	}

	return nil, fmt.Errorf("ack not found: dstPortID=%v dstChannel=%v sequence=%v", dstPortID, dstChannel, sequence)
}

type PacketAcknowledgement struct {
	Sequence uint64
	Data     []byte
}

func (chain *Chain) getAllAcknowledgements(
	ctx core.QueryContext,
	dstPortID string,
	dstChannel string,
) ([]PacketAcknowledgement, error) {
	var acks []PacketAcknowledgement
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   new(big.Int).SetUint64(ctx.Height().GetRevisionHeight()),
		Addresses: []common.Address{
			chain.config.IBCAddress(),
		},
		Topics: [][]common.Hash{{
			abiWriteAcknowledgement.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx.Context(), query)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		if values, err := abiWriteAcknowledgement.Inputs.Unpack(log.Data); err != nil {
			return nil, err
		} else {
			if len(values) != 4 {
				return nil, fmt.Errorf("unexpected values: %v", values)
			}
			if dstPortID == values[0].(string) && dstChannel == values[1].(string) {
				acks = append(acks, PacketAcknowledgement{
					Sequence: values[2].(uint64),
					Data:     values[3].([]byte),
				})
			}
		}
	}
	return acks, nil
}
