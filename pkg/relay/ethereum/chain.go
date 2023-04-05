package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	committypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/wallet"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Chain struct {
	config ChainConfig

	pathEnd          *core.PathEnd
	homePath         string
	chainID          *big.Int
	codec            codec.ProtoCodecMarshaler
	msgEventListener core.MsgEventListener

	relayerPrvKey *ecdsa.PrivateKey
	client        *client.ETHClient
	ibcHandler    *ibchandler.Ibchandler
}

var _ core.Chain = (*Chain)(nil)

func NewChain(config ChainConfig) (*Chain, error) {
	id := big.NewInt(config.EthChainId)
	client, err := client.NewETHClient(config.RpcAddr)
	if err != nil {
		return nil, err
	}
	key, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(config.HdwMnemonic, config.HdwPath)
	if err != nil {
		return nil, err
	}
	ibcHandler, err := ibchandler.NewIbchandler(config.IBCAddress(), client)
	if err != nil {
		return nil, err
	}
	return &Chain{
		config:        config,
		client:        client,
		relayerPrvKey: key,
		chainID:       id,

		ibcHandler: ibcHandler,
	}, nil
}

// Config returns ChainConfig
func (c *Chain) Config() ChainConfig {
	return c.config
}

// Init ...
func (c *Chain) Init(homePath string, timeout time.Duration, codec codec.ProtoCodecMarshaler, debug bool) error {
	c.homePath = homePath
	c.codec = codec
	return nil
}

// SetupForRelay ...
func (c *Chain) SetupForRelay(ctx context.Context) error {
	return nil
}

// ChainID returns ID of the chain
func (c *Chain) ChainID() string {
	return c.config.ChainId
}

// GetLatestHeight gets the chain for the latest height and returns it
func (c *Chain) LatestHeight() (ibcexported.Height, error) {
	bn, err := c.client.BlockNumber(context.TODO())
	if err != nil {
		return nil, err
	}
	return clienttypes.NewHeight(0, bn), nil
}

// GetAddress returns the address of relayer
func (c *Chain) GetAddress() (sdk.AccAddress, error) {
	addr := make([]byte, 20)
	return addr, nil
}

// Marshaler returns the marshaler
func (c *Chain) Codec() codec.ProtoCodecMarshaler {
	return c.codec
}

// Client returns the RPC client for ethereum
func (c *Chain) Client() *client.ETHClient {
	return c.client
}

// SetRelayInfo sets source's path and counterparty's info to the chain
func (c *Chain) SetRelayInfo(p *core.PathEnd, _ *core.ProvableChain, _ *core.PathEnd) error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("path on chain %s failed to set: %w", c.ChainID(), err)
	}
	c.pathEnd = p
	return nil
}

func (c *Chain) Path() *core.PathEnd {
	return c.pathEnd
}

// RegisterMsgEventListener registers a given EventListener to the chain
func (c *Chain) RegisterMsgEventListener(listener core.MsgEventListener) {
	c.msgEventListener = listener
}

// QueryClientConsensusState retrevies the latest consensus state for a client in state at a given height
func (c *Chain) QueryClientConsensusState(ctx core.QueryContext, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	s, found, err := c.ibcHandler.GetConsensusState(c.callOptsFromQueryContext(ctx), c.pathEnd.ClientID, pbToHostHeight(dstClientConsHeight))
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("client consensus not found: %v", c.pathEnd.ClientID)
	}
	var consensusState ibcexported.ConsensusState
	if err := c.Codec().UnmarshalInterface(s, &consensusState); err != nil {
		return nil, err
	}
	any, err := clienttypes.PackConsensusState(consensusState)
	if err != nil {
		return nil, err
	}
	return clienttypes.NewQueryConsensusStateResponse(any, nil, ctx.Height().(clienttypes.Height)), nil
}

// QueryClientState returns the client state of dst chain
// height represents the height of dst chain
func (c *Chain) QueryClientState(ctx core.QueryContext) (*clienttypes.QueryClientStateResponse, error) {
	s, found, err := c.ibcHandler.GetClientState(c.callOptsFromQueryContext(ctx), c.pathEnd.ClientID)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("client not found: %v", c.pathEnd.ClientID)
	}
	var clientState ibcexported.ClientState
	if err := c.Codec().UnmarshalInterface(s, &clientState); err != nil {
		return nil, err
	}
	any, err := clienttypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}
	return clienttypes.NewQueryClientStateResponse(any, nil, ctx.Height().(clienttypes.Height)), nil
}

var emptyConnRes = conntypes.NewQueryConnectionResponse(
	conntypes.NewConnectionEnd(
		conntypes.UNINITIALIZED,
		"client",
		conntypes.NewCounterparty(
			"client",
			"connection",
			committypes.NewMerklePrefix([]byte{}),
		),
		[]*conntypes.Version{},
		0,
	),
	[]byte{},
	clienttypes.NewHeight(0, 0),
)

// QueryConnection returns the remote end of a given connection
func (c *Chain) QueryConnection(ctx core.QueryContext) (*conntypes.QueryConnectionResponse, error) {
	conn, found, err := c.ibcHandler.GetConnection(c.callOptsFromQueryContext(ctx), c.pathEnd.ConnectionID)
	if err != nil {
		return nil, err
	} else if !found {
		return emptyConnRes, nil
	}
	return conntypes.NewQueryConnectionResponse(connectionEndToPB(conn), nil, ctx.Height().(clienttypes.Height)), nil
}

var emptyChannelRes = chantypes.NewQueryChannelResponse(
	chantypes.NewChannel(
		chantypes.UNINITIALIZED,
		chantypes.UNORDERED,
		chantypes.NewCounterparty(
			"port",
			"channel",
		),
		[]string{},
		"version",
	),
	[]byte{},
	clienttypes.NewHeight(0, 0),
)

// QueryChannel returns the channel associated with a channelID
func (c *Chain) QueryChannel(ctx core.QueryContext) (chanRes *chantypes.QueryChannelResponse, err error) {
	chann, found, err := c.ibcHandler.GetChannel(c.callOptsFromQueryContext(ctx), c.pathEnd.PortID, c.pathEnd.ChannelID)
	if err != nil {
		return nil, err
	} else if !found {
		return emptyChannelRes, nil
	}
	return chantypes.NewQueryChannelResponse(channelToPB(chann), nil, ctx.Height().(clienttypes.Height)), nil
}

// QueryPacketCommitment returns the packet commitment corresponding to a given sequence
func (c *Chain) QueryPacketCommitment(ctx core.QueryContext, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	packet, err := c.QueryPacket(ctx, seq)
	if err != nil {
		return nil, err
	}
	commitment := chantypes.CommitPacket(c.Codec(), packet)
	return chantypes.NewQueryPacketCommitmentResponse(commitment, nil, ctx.Height().(clienttypes.Height)), nil
}

// QueryPacketAcknowledgementCommitment returns the acknowledgement corresponding to a given sequence
func (c *Chain) QueryPacketAcknowledgementCommitment(ctx core.QueryContext, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	ack, err := c.QueryPacketAcknowledgement(ctx, seq)
	if err != nil {
		return nil, err
	}
	commitment := chantypes.CommitAcknowledgement(ack)
	return chantypes.NewQueryPacketAcknowledgementResponse(commitment, nil, ctx.Height().(clienttypes.Height)), nil
}

// NOTE: The current implementation returns all packets, including those for that acknowledgement has already received.
// QueryPacketCommitments returns an array of packet commitments
func (c *Chain) QueryPacketCommitments(ctx core.QueryContext, offset uint64, limit uint64) (comRes *chantypes.QueryPacketCommitmentsResponse, err error) {
	// WARNING: It may be slow to use in the production. Instead of it, it might be better to use an external event indexer to get all packet commitments.
	packets, err := c.getAllPackets(ctx, c.pathEnd.PortID, c.pathEnd.ChannelID)
	if err != nil {
		return nil, err
	}
	var res chantypes.QueryPacketCommitmentsResponse
	for _, p := range packets {
		ps := chantypes.NewPacketState(c.pathEnd.PortID, c.pathEnd.ChannelID, p.Sequence, chantypes.CommitPacket(c.Codec(), p))
		res.Commitments = append(res.Commitments, &ps)
	}
	res.Height = ctx.Height().(clienttypes.Height)
	return &res, nil
}

// QueryUnrecievedPackets returns a list of unrelayed packet commitments
func (c *Chain) QueryUnrecievedPackets(ctx core.QueryContext, seqs []uint64) ([]uint64, error) {
	var ret []uint64
	for _, seq := range seqs {
		found, err := c.ibcHandler.HasPacketReceipt(c.callOptsFromQueryContext(ctx), c.pathEnd.PortID, c.pathEnd.ChannelID, seq)
		if err != nil {
			return nil, err
		} else if !found {
			ret = append(ret, seq)
		}
	}
	return ret, nil
}

// QueryPacketAcknowledgementCommitments returns an array of packet acks
func (c *Chain) QueryPacketAcknowledgementCommitments(ctx core.QueryContext, offset uint64, limit uint64) (comRes *chantypes.QueryPacketAcknowledgementsResponse, err error) {
	// WARNING: It may be slow to use in the production. Instead of it, it might be better to use an external event indexer to get all packet acknowledgements.
	acks, err := c.getAllAcknowledgements(ctx, c.pathEnd.PortID, c.pathEnd.ChannelID)
	if err != nil {
		return nil, err
	}
	var res chantypes.QueryPacketAcknowledgementsResponse
	for _, a := range acks {
		ps := chantypes.NewPacketState(c.pathEnd.PortID, c.pathEnd.ChannelID, a.Sequence, chantypes.CommitAcknowledgement(a.Data))
		res.Acknowledgements = append(res.Acknowledgements, &ps)
	}
	return &res, nil
}

// QueryUnrecievedAcknowledgements returns a list of unrelayed packet acks
func (c *Chain) QueryUnrecievedAcknowledgements(ctx core.QueryContext, seqs []uint64) ([]uint64, error) {
	var ret []uint64
	for _, seq := range seqs {
		_, found, err := c.ibcHandler.GetHashedPacketCommitment(c.callOptsFromQueryContext(ctx), c.pathEnd.PortID, c.pathEnd.ChannelID, seq)
		if err != nil {
			return nil, err
		} else if found {
			ret = append(ret, seq)
		}
	}
	return ret, nil
}

// QueryPacket returns the packet corresponding to a sequence
func (c *Chain) QueryPacket(ctx core.QueryContext, sequence uint64) (*chantypes.Packet, error) {
	return c.findPacket(ctx, c.pathEnd.PortID, c.pathEnd.ChannelID, sequence)
}

// QueryPacketAcknowledgement returns the acknowledgement corresponding to a sequence
func (c *Chain) QueryPacketAcknowledgement(ctx core.QueryContext, sequence uint64) ([]byte, error) {
	return c.findAcknowledgement(ctx, c.pathEnd.PortID, c.pathEnd.ChannelID, sequence)
}

// QueryBalance returns the amount of coins in the relayer account
func (c *Chain) QueryBalance(ctx core.QueryContext, address sdk.AccAddress) (sdk.Coins, error) {
	panic("not supported")
}

// QueryDenomTraces returns all the denom traces from a given chain
func (c *Chain) QueryDenomTraces(ctx core.QueryContext, offset uint64, limit uint64) (*transfertypes.QueryDenomTracesResponse, error) {
	panic("not supported")
}

func (c *Chain) callOptsFromQueryContext(ctx core.QueryContext) *bind.CallOpts {
	return c.CallOpts(ctx.Context(), int64(ctx.Height().GetRevisionHeight()))
}
