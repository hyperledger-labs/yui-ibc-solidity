package relay

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

// SendMsgs sends msgs to the chain
func (c *Chain) SendMsgs(msgs []sdk.Msg) ([]byte, error) {
	for _, msg := range msgs {
		var (
			tx  *gethtypes.Transaction
			err error
		)
		ctx := context.Background()
		opts := c.TxOpts(ctx)
		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:
			tx, err = c.TxCreateClient(opts, msg)
		case *clienttypes.MsgUpdateClient:
			tx, err = c.TxUpdateClient(opts, msg)
		case *conntypes.MsgConnectionOpenInit:
			tx, err = c.TxConnectionOpenInit(opts, msg)
		case *conntypes.MsgConnectionOpenTry:
			tx, err = c.TxConnectionOpenTry(opts, msg)
		case *conntypes.MsgConnectionOpenAck:
			tx, err = c.TxConnectionOpenAck(opts, msg)
		case *conntypes.MsgConnectionOpenConfirm:
			tx, err = c.TxConnectionOpenConfirm(opts, msg)
		case *chantypes.MsgChannelOpenInit:
			tx, err = c.TxChannelOpenInit(opts, msg)
		case *chantypes.MsgChannelOpenTry:
			tx, err = c.TxChannelOpenTry(opts, msg)
		case *chantypes.MsgChannelOpenAck:
			tx, err = c.TxChannelOpenAck(opts, msg)
		case *chantypes.MsgChannelOpenConfirm:
			tx, err = c.TxChannelOpenConfirm(opts, msg)
		case *chantypes.MsgRecvPacket:
			tx, err = c.TxRecvPacket(opts, msg)
		case *chantypes.MsgAcknowledgement:
			tx, err = c.TxAcknowledgement(opts, msg)
		// case *transfertypes.MsgTransfer:
		// 	err = c.client.transfer(msg)
		default:
			panic("illegal msg type")
		}
		if err != nil {
			return nil, err
		}
		if err := c.TxSync(ctx, tx); err != nil {
			return nil, err
		}
		if c.msgEventListener != nil {
			if err := c.msgEventListener.OnSentMsg([]sdk.Msg{msg}); err != nil {
				log.Println("failed to OnSendMsg call", "msg", msg, "err", err)
			}
		}
	}
	return nil, nil
}

// Send sends msgs to the chain and logging a result of it
// It returns a boolean value whether the result is success
func (c *Chain) Send(msgs []sdk.Msg) bool {
	_, err := c.SendMsgs(msgs)
	if err != nil {
		log.Println("ethereum: failed to send:", err)
	}
	return err == nil
}

func (c *Chain) TxCreateClient(opts *bind.TransactOpts, msg *clienttypes.MsgCreateClient) (*gethtypes.Transaction, error) {
	var clientState exported.ClientState
	if err := c.codec.UnpackAny(msg.ClientState, &clientState); err != nil {
		return nil, err
	}
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	consensusStateBytes, err := proto.Marshal(msg.ConsensusState)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.CreateClient(opts, ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          clientState.ClientType(),
		Height:              pbToHandlerHeight(clientState.GetLatestHeight()),
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	})
}

func (c *Chain) TxUpdateClient(opts *bind.TransactOpts, msg *clienttypes.MsgUpdateClient) (*gethtypes.Transaction, error) {
	headerBytes, err := proto.Marshal(msg.Header)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.UpdateClient(opts, ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: msg.ClientId,
		Header:   headerBytes,
	})
}

func (c *Chain) TxConnectionOpenInit(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenInit) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ConnectionOpenInit(opts, ibchandler.IBCMsgsMsgConnectionOpenInit{
		ClientId: msg.ClientId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod: msg.DelayPeriod,
	})
}

func (c *Chain) TxConnectionOpenTry(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenTry) (*gethtypes.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	var versions []ibchandler.VersionData
	for _, v := range msg.CounterpartyVersions {
		versions = append(versions, ibchandler.VersionData(*v))
	}
	return c.ibcHandler.ConnectionOpenTry(opts, ibchandler.IBCMsgsMsgConnectionOpenTry{
		PreviousConnectionId: msg.PreviousConnectionId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod:          msg.DelayPeriod,
		ClientId:             msg.ClientId,
		ClientStateBytes:     clientStateBytes,
		CounterpartyVersions: versions,
		ProofInit:            msg.ProofInit,
		ProofClient:          msg.ProofClient,
		ProofConsensus:       msg.ProofConsensus,
		ProofHeight:          pbToHandlerHeight(msg.ProofHeight),
		ConsensusHeight:      pbToHandlerHeight(msg.ConsensusHeight),
	})
}

func (c *Chain) TxConnectionOpenAck(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenAck) (*gethtypes.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.ConnectionOpenAck(opts, ibchandler.IBCMsgsMsgConnectionOpenAck{
		ConnectionId:     msg.ConnectionId,
		ClientStateBytes: clientStateBytes,
		Version: ibchandler.VersionData{
			Identifier: msg.Version.Identifier,
			Features:   msg.Version.Features,
		},
		CounterpartyConnectionID: msg.CounterpartyConnectionId,
		ProofTry:                 msg.ProofTry,
		ProofClient:              msg.ProofClient,
		ProofConsensus:           msg.ProofConsensus,
		ProofHeight:              pbToHandlerHeight(msg.ProofHeight),
		ConsensusHeight:          pbToHandlerHeight(msg.ConsensusHeight),
	})
}

func (c *Chain) TxConnectionOpenConfirm(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenConfirm) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ConnectionOpenConfirm(opts, ibchandler.IBCMsgsMsgConnectionOpenConfirm{
		ConnectionId: msg.ConnectionId,
		ProofAck:     msg.ProofAck,
		ProofHeight:  pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) TxChannelOpenInit(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenInit) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ChannelOpenInit(opts, ibchandler.IBCMsgsMsgChannelOpenInit{
		PortId: msg.PortId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
	})
}

func (c *Chain) TxChannelOpenTry(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenTry) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ChannelOpenTry(opts, ibchandler.IBCMsgsMsgChannelOpenTry{
		PortId:            msg.PortId,
		PreviousChannelId: msg.PreviousChannelId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
		CounterpartyVersion: msg.CounterpartyVersion,
		ProofInit:           msg.ProofInit,
		ProofHeight:         pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) TxChannelOpenAck(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenAck) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ChannelOpenAck(opts, ibchandler.IBCMsgsMsgChannelOpenAck{
		PortId:                msg.PortId,
		ChannelId:             msg.ChannelId,
		CounterpartyVersion:   msg.CounterpartyVersion,
		CounterpartyChannelId: msg.CounterpartyChannelId,
		ProofTry:              msg.ProofTry,
		ProofHeight:           pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) TxChannelOpenConfirm(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenConfirm) (*gethtypes.Transaction, error) {
	return c.ibcHandler.ChannelOpenConfirm(opts, ibchandler.IBCMsgsMsgChannelOpenConfirm{
		PortId:      msg.PortId,
		ChannelId:   msg.ChannelId,
		ProofAck:    msg.ProofAck,
		ProofHeight: pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) TxRecvPacket(opts *bind.TransactOpts, msg *chantypes.MsgRecvPacket) (*gethtypes.Transaction, error) {
	return c.ibcHandler.RecvPacket(opts, ibchandler.IBCMsgsMsgPacketRecv{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Proof:       msg.ProofCommitment,
		ProofHeight: pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) TxAcknowledgement(opts *bind.TransactOpts, msg *chantypes.MsgAcknowledgement) (*gethtypes.Transaction, error) {
	return c.ibcHandler.AcknowledgePacket(opts, ibchandler.IBCMsgsMsgPacketAcknowledgement{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Acknowledgement: msg.Acknowledgement,
		Proof:           msg.ProofAcked,
		ProofHeight:     pbToHandlerHeight(msg.ProofHeight),
	})
}

func (chain *Chain) TxSync(ctx context.Context, tx *gethtypes.Transaction) error {
	var receipt *gethtypes.Receipt
	err := retry.Do(
		func() error {
			rc, err := chain.client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return err
			}
			receipt = rc
			return nil
		},
		// TODO make these configurable
		retry.Delay(1*time.Second),
		retry.Attempts(10),
	)
	if err != nil {
		return err
	}
	if receipt.Status == gethtypes.ReceiptStatusSuccessful {
		return nil
	} else {
		return fmt.Errorf("failed to call transaction: err='%v' rc='%v'", err, receipt)
	}
}

func (chain *Chain) TxSyncIfNoError(ctx context.Context) func(tx *gethtypes.Transaction, err error) error {
	return func(tx *gethtypes.Transaction, err error) error {
		if err != nil {
			return err
		}
		return chain.TxSync(ctx, tx)
	}
}
