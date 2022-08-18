package ibft2

import (
	"context"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client/ibft2"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/relay/ethereum"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Prover struct {
	chain  *ethereum.Chain
	config ProverConfig
}

var _ core.ProverI = (*Prover)(nil)

func NewProver(chain *ethereum.Chain, config ProverConfig) (*Prover, error) {
	return &Prover{chain: chain, config: config}, nil
}

// Init initializes the chain
func (pr *Prover) Init(homePath string, timeout time.Duration, codec codec.ProtoCodecMarshaler, debug bool) error {
	return nil
}

// SetRelayInfo sets source's path and counterparty's info to the chain
func (pr *Prover) SetRelayInfo(path *core.PathEnd, counterparty *core.ProvableChain, counterpartyPath *core.PathEnd) error {
	return nil
}

// SetupForRelay performs chain-specific setup before starting the relay
func (pr *Prover) SetupForRelay(ctx context.Context) error {
	return nil
}

// GetChainID returns the chain ID
func (pr *Prover) GetChainID() string {
	return pr.chain.ChainID()
}

// QueryLatestHeader returns the latest header from the chain
func (pr *Prover) QueryLatestHeader() (out core.HeaderI, err error) {
	res, err := pr.chain.QueryClientState(0)
	if err != nil {
		return nil, err
	}
	cs, err := clienttypes.UnpackClientState(res.ClientState)
	if err != nil {
		return nil, err
	}

	h, err := pr.chain.Client().HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	headerRLP, err := rlp.EncodeToBytes(h)
	if err != nil {
		return nil, err
	}
	stateProof, err := pr.QueryStateProof(h.Number)
	if err != nil {
		return nil, err
	}

	return &ibft2.Header{
		BesuHeaderRlp:     headerRLP,
		TrustedHeight:     ibcclient.Height(cs.GetLatestHeight().(clienttypes.Height)),
		AccountStateProof: stateProof.AccountProofRLP,
	}, nil
}

func (pr *Prover) QueryStateProof(bn *big.Int) (*client.StateProof, error) {
	return pr.chain.Client().GetStateProof(
		pr.chain.Config().IBCHostAddress(),
		nil,
		bn,
	)
}

// GetLatestLightHeight returns the latest height on the light client
func (pr *Prover) GetLatestLightHeight() (int64, error) {
	panic("not implemented") // TODO: Implement
}

// CreateMsgCreateClient creates a CreateClientMsg to this chain
func (pr *Prover) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	panic("not implemented") // TODO: Implement
}

// SetupHeader creates a new header based on a given header
func (pr *Prover) SetupHeader(dst core.LightClientIBCQueryierI, baseSrcHeader core.HeaderI) (core.HeaderI, error) {
	panic("not implemented") // TODO: Implement
}

// UpdateLightWithHeader updates a header on the light client and returns the header and height corresponding to the chain
func (pr *Prover) UpdateLightWithHeader() (header core.HeaderI, provableHeight int64, queryableHeight int64, err error) {
	panic("not implemented") // TODO: Implement
}

// QueryClientConsensusState returns the ClientConsensusState and its proof
func (pr *Prover) QueryClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryClientStateWithProof returns the ClientState and its proof
func (pr *Prover) QueryClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryConnectionWithProof returns the Connection and its proof
func (pr *Prover) QueryConnectionWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryChannelWithProof returns the Channel and its proof
func (pr *Prover) QueryChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	panic("not implemented") // TODO: Implement
}

// QueryPacketCommitmentWithProof returns the packet commitment and its proof
func (pr *Prover) QueryPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	panic("not implemented") // TODO: Implement
}

// QueryPacketAcknowledgementCommitmentWithProof returns the packet acknowledgement commitment and its proof
func (pr *Prover) QueryPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	panic("not implemented") // TODO: Implement
}
