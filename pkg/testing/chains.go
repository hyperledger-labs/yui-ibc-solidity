package testing

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/client"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibchandler"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibchost"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ibcidentifier"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ics20bank"
	"github.com/datachainlab/ibc-solidity/pkg/contract/ics20transferbank"
	"github.com/datachainlab/ibc-solidity/pkg/contract/simpletoken"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	ibcclient "github.com/datachainlab/ibc-solidity/pkg/ibc/client"
	ibft2clienttypes "github.com/datachainlab/ibc-solidity/pkg/ibc/client/ibft2"
	mockclienttypes "github.com/datachainlab/ibc-solidity/pkg/ibc/client/mock"
	"github.com/gogo/protobuf/proto"

	"github.com/datachainlab/ibc-solidity/pkg/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

const (
	DefaultChannelVersion        = "ics20-1"
	DefaultDelayPeriod    uint64 = 0
	DefaultPrefix                = "ibc"
	TransferPort                 = "transfer"
)

type Chain struct {
	t *testing.T

	// Core Modules
	client        client.Client
	IBCHandler    ibchandler.Ibchandler
	IBCHost       ibchost.Ibchost
	IBCIdentifier ibcidentifier.Ibcidentifier

	// App Modules
	SimpleToken   simpletoken.Simpletoken
	ICS20Transfer ics20transferbank.Ics20transferbank
	ICS20Bank     ics20bank.Ics20bank

	chainID int64

	ContractConfig ContractConfig

	key0 *ecdsa.PrivateKey

	// State
	LastContractState client.ContractState

	// IBC specific helpers
	ClientIDs   []string          // ClientID's used on this chain
	Connections []*TestConnection // track connectionID's created for this chain
	IBCID       uint64
}

type ContractConfig interface {
	GetIBCHostAddress() common.Address
	GetIBCHandlerAddress() common.Address
	GetIBCIdentifierAddress() common.Address
	GetIBFT2ClientAddress() common.Address
	GetMockClientAddress() common.Address

	GetSimpleTokenAddress() common.Address
	GetICS20TransferBankAddress() common.Address
	GetICS20BankAddress() common.Address
}

func NewChain(t *testing.T, chainID int64, client client.Client, config ContractConfig, mnemonicPhrase string, ibcID uint64) *Chain {
	key0, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(mnemonicPhrase, "m/44'/60'/0'/0/0")
	if err != nil {
		t.Error(err)
	}

	ibcHost, err := ibchost.NewIbchost(config.GetIBCHostAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ibcHandler, err := ibchandler.NewIbchandler(config.GetIBCHandlerAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ibcIdentifier, err := ibcidentifier.NewIbcidentifier(config.GetIBCIdentifierAddress(), client)
	if err != nil {
		t.Error(err)
	}
	simpletoken, err := simpletoken.NewSimpletoken(config.GetSimpleTokenAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ics20transfer, err := ics20transferbank.NewIcs20transferbank(config.GetICS20TransferBankAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ics20bank, err := ics20bank.NewIcs20bank(config.GetICS20BankAddress(), client)
	if err != nil {
		t.Error(err)
	}

	return &Chain{
		t:              t,
		client:         client,
		chainID:        chainID,
		ContractConfig: config,
		key0:           key0,
		IBCID:          ibcID,

		IBCHost:       *ibcHost,
		IBCHandler:    *ibcHandler,
		IBCIdentifier: *ibcIdentifier,
		SimpleToken:   *simpletoken,
		ICS20Transfer: *ics20transfer,
		ICS20Bank:     *ics20bank,
	}
}

func (chain *Chain) Client() client.Client {
	return chain.client
}

func (chain *Chain) ClientType() string {
	return chain.client.ClientType()
}

func (chain *Chain) TxOpts(ctx context.Context) *bind.TransactOpts {
	return client.MakeGenTxOpts(big.NewInt(chain.chainID), chain.key0)(ctx)
}

func (chain *Chain) CallOpts(ctx context.Context) *bind.CallOpts {
	opts := chain.TxOpts(ctx)
	return &bind.CallOpts{
		From:    opts.From,
		Context: opts.Context,
	}
}

func (chain *Chain) ChainID() int64 {
	return chain.chainID
}

func (chain *Chain) ChainIDString() string {
	return fmt.Sprint(chain.chainID)
}

func (chain *Chain) GetCommitmentPrefix() []byte {
	return []byte(DefaultPrefix)
}

func (chain *Chain) GetIBFT2ClientState(clientID string) *ibft2clienttypes.ClientState {
	ctx := context.Background()
	bz, found, err := chain.IBCHost.GetClientState(chain.CallOpts(ctx), clientID)
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("clientState not found")
	}
	var cs ibft2clienttypes.ClientState
	if err := proto.Unmarshal(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetMockClientState(clientID string) *mockclienttypes.ClientState {
	ctx := context.Background()
	bz, found, err := chain.IBCHost.GetClientState(chain.CallOpts(ctx), clientID)
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("clientState not found")
	}
	var cs mockclienttypes.ClientState
	if err := proto.Unmarshal(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetContractState(counterparty *Chain, counterpartyClientID string, storageKeys [][]byte) (client.ContractState, error) {
	var height uint64

	switch counterparty.ClientType() {
	case ibcclient.MockClient:
		height = counterparty.GetMockClientState(counterpartyClientID).LatestHeight
	case ibcclient.BesuIBFT2Client:
		height = counterparty.GetIBFT2ClientState(counterpartyClientID).LatestHeight
	default:
		return nil, fmt.Errorf("unknown client type: '%v'", counterparty.ClientType())
	}

	return chain.client.GetContractState(
		context.Background(),
		chain.ContractConfig.GetIBCHostAddress(),
		storageKeys,
		big.NewInt(int64(height)),
	)
}

func (chain *Chain) Init() error {
	ctx := context.Background()
	if err := chain.WaitIfNoError(ctx)(
		chain.IBCHost.SetIBCModule(
			chain.TxOpts(ctx),
			chain.ContractConfig.GetIBCHandlerAddress(),
		),
	); err != nil {
		return err
	}

	if name, err := chain.IBCIdentifier.PortCapabilityPath(chain.CallOpts(ctx), TransferPort); err != nil {
		return err
	} else if _, found, err := chain.IBCHost.GetModuleOwner(chain.CallOpts(ctx), name); err != nil {
		return err
	} else if !found {
		if err := chain.WaitIfNoError(ctx)(
			chain.IBCHandler.BindPort(chain.TxOpts(ctx), TransferPort, chain.ContractConfig.GetICS20TransferBankAddress()),
		); err != nil {
			return err
		}
	}

	if _, found, err := chain.IBCHost.GetClientImpl(chain.CallOpts(ctx), ibcclient.BesuIBFT2Client); err != nil {
		return err
	} else if !found {
		if err := chain.WaitIfNoError(ctx)(
			chain.IBCHandler.RegisterClient(
				chain.TxOpts(ctx),
				ibcclient.BesuIBFT2Client,
				chain.ContractConfig.GetIBFT2ClientAddress(),
			),
		); err != nil {
			return err
		}
	}

	if _, found, err := chain.IBCHost.GetClientImpl(chain.CallOpts(ctx), ibcclient.MockClient); err != nil {
		return err
	} else if !found {
		if err := chain.WaitIfNoError(ctx)(
			chain.IBCHandler.RegisterClient(
				chain.TxOpts(ctx),
				ibcclient.MockClient,
				chain.ContractConfig.GetMockClientAddress(),
			),
		); err != nil {
			return err
		}
	}

	if err := chain.WaitIfNoError(ctx)(
		chain.ICS20Bank.SetOperator(chain.TxOpts(ctx), chain.ContractConfig.GetICS20TransferBankAddress()),
	); err != nil {
		return err
	}

	return nil
}

func (chain *Chain) ConstructMockMsgCreateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgCreateClient {
	clientState := mockclienttypes.ClientState{
		LatestHeight: counterparty.LastHeader().Number.Uint64(),
	}
	consensusState := mockclienttypes.ConsensusState{
		Timestamp: counterparty.LastHeader().Time,
	}
	clientStateBytes, err := proto.Marshal(&clientState)
	if err != nil {
		panic(err)
	}
	consensusStateBytes, err := proto.Marshal(&consensusState)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgCreateClient{
		ClientId:            clientID,
		ClientType:          ibcclient.MockClient,
		Height:              clientState.LatestHeight,
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	}
}

func (chain *Chain) ConstructIBFT2MsgCreateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgCreateClient {
	clientState := ibft2clienttypes.ClientState{
		ChainId:         counterparty.ChainIDString(),
		IbcStoreAddress: counterparty.ContractConfig.GetIBCHostAddress().Bytes(),
		LatestHeight:    counterparty.LastHeader().Number.Uint64(),
	}
	consensusState := ibft2clienttypes.ConsensusState{
		Timestamp:  counterparty.LastHeader().Time,
		Root:       counterparty.LastHeader().Root.Bytes(),
		Validators: counterparty.LastContractState.(client.IBFT2ContractState).Validators(),
	}
	clientStateBytes, err := proto.Marshal(&clientState)
	if err != nil {
		panic(err)
	}
	consensusStateBytes, err := proto.Marshal(&consensusState)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgCreateClient{
		ClientId:            clientID,
		ClientType:          ibcclient.BesuIBFT2Client,
		Height:              clientState.LatestHeight,
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	}
}

func (chain *Chain) ConstructMockMsgUpdateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgUpdateClient {
	cs := counterparty.LastContractState.(client.ETHContractState)
	bz, err := rlp.EncodeToBytes(cs.Header())
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: clientID,
		Header:   bz,
	}
}

func (chain *Chain) ConstructIBFT2MsgUpdateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgUpdateClient {
	trustedHeight := chain.GetIBFT2ClientState(clientID).LatestHeight
	cs := counterparty.LastContractState.(client.IBFT2ContractState)
	var header = ibft2clienttypes.Header{
		BesuHeaderRlp:     cs.SealingHeaderRLP(),
		Seals:             cs.CommitSeals,
		TrustedHeight:     trustedHeight,
		AccountStateProof: cs.ETHProof().AccountProofRLP,
	}
	headerBytes, err := proto.Marshal(&header)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: clientID,
		Header:   headerBytes,
	}
}

func (chain *Chain) UpdateHeader() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		state, err := chain.client.GetContractState(ctx, chain.ContractConfig.GetIBCHostAddress(), nil, nil)
		if err != nil {
			panic(err)
		}
		if chain.LastContractState == nil || state.Header().Number.Cmp(chain.LastHeader().Number) == 1 {
			chain.LastContractState = state
			return
		} else {
			continue
		}
	}
}

func (chain *Chain) CreateMockClient(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructMockMsgCreateClient(counterparty, clientID)
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.CreateClient(chain.TxOpts(ctx), msg),
	)
}

func (chain *Chain) UpdateMockClient(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructMockMsgUpdateClient(counterparty, clientID)
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.UpdateClient(chain.TxOpts(ctx), msg),
	)
}

func (chain *Chain) CreateIBFT2Client(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructIBFT2MsgCreateClient(counterparty, clientID)
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.CreateClient(chain.TxOpts(ctx), msg),
	)
}

func (chain *Chain) UpdateIBFT2Client(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructIBFT2MsgUpdateClient(counterparty, clientID)
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.UpdateClient(chain.TxOpts(ctx), msg),
	)
}

func (chain *Chain) ConnectionOpenInit(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) error {
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ConnectionOpenInit(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgConnectionOpenInit{
				ClientId:     connection.ClientID,
				ConnectionId: connection.ID,
				Counterparty: ibchandler.CounterpartyData{
					ClientId:     connection.CounterpartyClientID,
					ConnectionId: "",
					Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
				},
				DelayPeriod: DefaultDelayPeriod,
			},
		),
	)
}

func (chain *Chain) ConnectionOpenTry(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) error {
	proof, err := counterparty.QueryProof(chain, connection.ClientID, chain.ConnectionStateCommitmentSlot(counterpartyConnection.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ConnectionOpenTry(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgConnectionOpenTry{
				ConnectionId: connection.ID,
				Counterparty: ibchandler.CounterpartyData{
					ClientId:     counterpartyConnection.ClientID,
					ConnectionId: counterpartyConnection.ID,
					Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
				},
				DelayPeriod: DefaultDelayPeriod,
				ClientId:    connection.ClientID,
				// ClientState: ibcconnection.ClientStateData{}, // TODO set chain's clientState
				CounterpartyVersions: []ibchandler.VersionData{
					{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
				},
				ProofHeight: proof.Height,
				ProofInit:   proof.Data,
			},
		),
	)
}

// ConnectionOpenAck will construct and execute a MsgConnectionOpenAck.
func (chain *Chain) ConnectionOpenAck(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proof, err := counterparty.QueryProof(chain, connection.ClientID, chain.ConnectionStateCommitmentSlot(counterpartyConnection.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ConnectionOpenAck(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgConnectionOpenAck{
				ConnectionId:             connection.ID,
				CounterpartyConnectionID: counterpartyConnection.ID,
				// clientState
				Version:     ibchandler.VersionData{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
				ProofTry:    proof.Data,
				ProofHeight: proof.Height,
			},
		),
	)
}

func (chain *Chain) ConnectionOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proof, err := counterparty.QueryProof(chain, connection.ClientID, chain.ConnectionStateCommitmentSlot(counterpartyConnection.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ConnectionOpenConfirm(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgConnectionOpenConfirm{
				ConnectionId: connection.ID,
				ProofAck:     proof.Data,
				ProofHeight:  proof.Height,
			},
		),
	)
}

func (chain *Chain) ChannelOpenInit(
	ctx context.Context,
	ch, counterparty TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) error {
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ChannelOpenInit(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgChannelOpenInit{
				ChannelId: ch.ID,
				PortId:    ch.PortID,
				Channel: ibchandler.ChannelData{
					State:    uint8(channeltypes.INIT),
					Ordering: uint8(order),
					Counterparty: ibchandler.ChannelCounterpartyData{
						PortId:    counterparty.PortID,
						ChannelId: "",
					},
					ConnectionHops: []string{connectionID},
					Version:        ch.Version,
				},
			},
		),
	)
}

func (chain *Chain) ChannelOpenTry(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.ChannelStateCommitmentSlot(counterpartyCh.PortID, counterpartyCh.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ChannelOpenTry(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgChannelOpenTry{
				PortId:    ch.PortID,
				ChannelId: ch.ID,
				Channel: ibchandler.ChannelData{
					State:    uint8(channeltypes.TRYOPEN),
					Ordering: uint8(order),
					Counterparty: ibchandler.ChannelCounterpartyData{
						PortId:    counterpartyCh.PortID,
						ChannelId: counterpartyCh.ID,
					},
					ConnectionHops: []string{connectionID},
					Version:        ch.Version,
				},
				CounterpartyVersion: counterpartyCh.Version,
				ProofInit:           proof.Data,
				ProofHeight:         proof.Height,
			},
		),
	)
}

func (chain *Chain) ChannelOpenAck(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.ChannelStateCommitmentSlot(counterpartyCh.PortID, counterpartyCh.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ChannelOpenAck(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgChannelOpenAck{
				PortId:                ch.PortID,
				ChannelId:             ch.ID,
				CounterpartyVersion:   counterpartyCh.Version,
				CounterpartyChannelId: counterpartyCh.ID,
				ProofTry:              proof.Data,
				ProofHeight:           proof.Height,
			},
		),
	)
}

func (chain *Chain) ChannelOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.ChannelStateCommitmentSlot(counterpartyCh.PortID, counterpartyCh.ID))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.ChannelOpenConfirm(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgChannelOpenConfirm{
				PortId:      ch.PortID,
				ChannelId:   ch.ID,
				ProofAck:    proof.Data,
				ProofHeight: proof.Height,
			},
		),
	)
}

func (chain *Chain) SendPacket(
	ctx context.Context,
	packet channeltypes.Packet,
) error {
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.SendPacket(
			chain.TxOpts(ctx),
			packetToCallData(packet),
		),
	)
}

func (chain *Chain) RecvPacket(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet channeltypes.Packet,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.PacketCommitmentSlot(packet.SourcePort, packet.SourceChannel, packet.Sequence))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.RecvPacket(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgPacketRecv{
				Packet:      packetToCallData(packet),
				Proof:       proof.Data,
				ProofHeight: proof.Height,
			},
		),
	)
}

func (chain *Chain) HandlePacketRecv(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet channeltypes.Packet,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.PacketCommitmentSlot(packet.SourcePort, packet.SourceChannel, packet.Sequence))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.RecvPacket(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgPacketRecv{
				Packet: ibchandler.PacketData{
					Sequence:           packet.Sequence,
					SourcePort:         packet.SourcePort,
					SourceChannel:      packet.SourceChannel,
					DestinationPort:    packet.DestinationPort,
					DestinationChannel: packet.DestinationChannel,
					Data:               packet.Data,
					TimeoutHeight:      ibchandler.HeightData(packet.TimeoutHeight),
					TimeoutTimestamp:   packet.TimeoutTimestamp,
				},
				Proof:       proof.Data,
				ProofHeight: proof.Height,
			},
		),
	)
}

func (chain *Chain) HandlePacketAcknowledgement(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.PacketAcknowledgementCommitmentSlot(packet.DestinationPort, packet.DestinationChannel, packet.Sequence))
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.AcknowledgePacket(
			chain.TxOpts(ctx),
			ibchandler.IBCMsgsMsgPacketAcknowledgement{
				Packet: ibchandler.PacketData{
					Sequence:           packet.Sequence,
					SourcePort:         packet.SourcePort,
					SourceChannel:      packet.SourceChannel,
					DestinationPort:    packet.DestinationPort,
					DestinationChannel: packet.DestinationChannel,
					Data:               packet.Data,
					TimeoutHeight:      ibchandler.HeightData(packet.TimeoutHeight),
					TimeoutTimestamp:   packet.TimeoutTimestamp,
				},
				Acknowledgement: acknowledgement,
				Proof:           proof.Data,
				ProofHeight:     proof.Height,
			},
		),
	)
}

func (chain *Chain) GetLastSentPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
) (*channeltypes.Packet, error) {
	seq, err := chain.IBCHost.GetNextSequenceSend(chain.CallOpts(ctx), sourcePortID, sourceChannel)
	if err != nil {
		return nil, err
	}
	return chain.FindPacket(ctx, sourcePortID, sourceChannel, seq-1)
}

func (chain *Chain) FindPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
	sequence uint64,
) (*channeltypes.Packet, error) {
	iter, err := chain.IBCHandler.FilterSendPacket(&bind.FilterOpts{
		Start:   0,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	for iter.Next() {
		p := iter.Event.Packet
		if p.SourcePort == sourcePortID && p.SourceChannel == sourceChannel && p.Sequence == sequence {
			return &channeltypes.Packet{
				Sequence:           p.Sequence,
				SourcePort:         p.SourcePort,
				SourceChannel:      p.SourceChannel,
				DestinationPort:    p.DestinationPort,
				DestinationChannel: p.DestinationChannel,
				Data:               p.Data,
				TimeoutHeight:      channeltypes.Height(p.TimeoutHeight),
				TimeoutTimestamp:   p.TimeoutTimestamp,
			}, nil
		}
	}
	return nil, fmt.Errorf("packet not found: sourcePortID=%v sourceChannel=%v sequence=%v", sourcePortID, sourceChannel, sequence)
}

func packetToCallData(packet channeltypes.Packet) ibchandler.PacketData {
	return ibchandler.PacketData{
		Sequence:           packet.Sequence,
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               packet.Data,
		TimeoutHeight:      ibchandler.HeightData(packet.TimeoutHeight),
		TimeoutTimestamp:   packet.TimeoutTimestamp,
	}
}

// Slot calculator

func (chain *Chain) ConnectionStateCommitmentSlot(connectionID string) string {
	key, err := chain.IBCIdentifier.ConnectionCommitmentSlot(chain.CallOpts(context.Background()), connectionID)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) ChannelStateCommitmentSlot(portID, channelID string) string {
	key, err := chain.IBCIdentifier.ChannelCommitmentSlot(chain.CallOpts(context.Background()), portID, channelID)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) PacketCommitmentSlot(portID, channelID string, sequence uint64) string {
	key, err := chain.IBCIdentifier.PacketCommitmentSlot(chain.CallOpts(context.Background()), portID, channelID, sequence)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) PacketAcknowledgementCommitmentSlot(portID, channelID string, sequence uint64) string {
	key, err := chain.IBCIdentifier.PacketAcknowledgementCommitmentSlot(chain.CallOpts(context.Background()), portID, channelID, sequence)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

// Querier

type Proof struct {
	Height uint64
	Data   []byte
}

func (chain *Chain) QueryProof(counterparty *Chain, counterpartyClientID string, storageKey string) (*Proof, error) {
	if !strings.HasPrefix(storageKey, "0x") {
		return nil, fmt.Errorf("storageKey must be hex string")
	}
	s, err := chain.GetContractState(counterparty, counterpartyClientID, [][]byte{[]byte(storageKey)})
	if err != nil {
		return nil, err
	}
	return &Proof{Height: s.Header().Number.Uint64(), Data: s.ETHProof().StorageProofRLP[0]}, nil
}

func (chain *Chain) LastHeader() *gethtypes.Header {
	return chain.LastContractState.Header()
}

func (chain *Chain) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction) error {
	rc, err := chain.Client().WaitForReceiptAndGet(ctx, tx)
	if err != nil {
		return err
	}
	if rc.Status() == 1 {
		return nil
	} else {
		return fmt.Errorf("failed to call transaction: err='%v' rc='%v' reason='%v'", err, rc, rc.RevertReason())
	}
}

func (chain *Chain) WaitIfNoError(ctx context.Context) func(tx *gethtypes.Transaction, err error) error {
	return func(tx *gethtypes.Transaction, err error) error {
		if err != nil {
			return err
		}
		if err := chain.WaitForReceiptAndGet(ctx, tx); err != nil {
			return err
		}
		return nil
	}
}

// NewClientID appends a new clientID string in the format:
// ClientFor<counterparty-chain-id><index>
func (chain *Chain) NewClientID(clientType string) string {
	clientID := fmt.Sprintf("%s-%s-%v-%v", clientType, strconv.Itoa(len(chain.ClientIDs)), chain.chainID, chain.IBCID)
	chain.ClientIDs = append(chain.ClientIDs, clientID)
	return clientID
}

// AddTestConnection appends a new TestConnection which contains references
// to the connection id, client id and counterparty client id.
func (chain *Chain) AddTestConnection(clientID, counterpartyClientID string) *TestConnection {
	conn := chain.ConstructNextTestConnection(clientID, counterpartyClientID)

	chain.Connections = append(chain.Connections, conn)
	return conn
}

// ConstructNextTestConnection constructs the next test connection to be
// created given a clientID and counterparty clientID. The connection id
// format: <chainID>-conn<index>
func (chain *Chain) ConstructNextTestConnection(clientID, counterpartyClientID string) *TestConnection {
	connectionID := fmt.Sprintf("connection-%v-%v-%v", uint64(len(chain.Connections)), chain.chainID, chain.IBCID)
	return &TestConnection{
		ID:                   connectionID,
		ClientID:             clientID,
		NextChannelVersion:   DefaultChannelVersion,
		CounterpartyClientID: counterpartyClientID,
	}
}

// AddTestChannel appends a new TestChannel which contains references to the port and channel ID
// used for channel creation and interaction. See 'NextTestChannel' for channel ID naming format.
func (chain *Chain) AddTestChannel(conn *TestConnection, portID string) TestChannel {
	channel := chain.NextTestChannel(conn, portID)
	conn.Channels = append(conn.Channels, channel)
	return channel
}

// NextTestChannel returns the next test channel to be created on this connection, but does not
// add it to the list of created channels. This function is expected to be used when the caller
// has not created the associated channel in app state, but would still like to refer to the
// non-existent channel usually to test for its non-existence.
//
// channel ID format: <connectionid>-chan<channel-index>
//
// The port is passed in by the caller.
func (chain *Chain) NextTestChannel(conn *TestConnection, portID string) TestChannel {
	channelID := fmt.Sprintf("channel-%v-%v", chain.chainID, chain.IBCID)
	return TestChannel{
		PortID:               portID,
		ID:                   channelID,
		ClientID:             conn.ClientID,
		CounterpartyClientID: conn.CounterpartyClientID,
		Version:              conn.NextChannelVersion,
	}
}
