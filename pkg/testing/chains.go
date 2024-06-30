package testing

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/chains"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/erc20"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibcmockapp"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20bank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20transferbank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/qbftclient"
	qbftclienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/clients/qbft"
	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/channel"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/core/commitment"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/wallet"
)

const (
	ICS20Version   = "ics20-1"
	MockAppVersion = "mockapp-1"
	TransferPort   = "transfer"
	MockPort       = "mock"

	RelayerKeyIndex uint32 = 0

	MockPacketData      = "mock packet data"
	MockFailPacketData  = "mock failed packet data"
	MockAsyncPacketData = "mock async packet data"
)

var (
	DefaultPrefix = "ibc"

	BlockTime            uint64 = 1000 * 1000 * 1000 // 1[sec]
	DefaultDelayPeriod   uint64 = 0                  // sec
	DefaultTrustPeriod   uint64 = 60                 // sec
	DefaultMaxClockDrift uint64 = 30                 // sec
)

var (
	abiIBCHandler abi.ABI
	abiSendPacket,
	abiWriteAcknowledgement,
	abiGeneratedClientIdentifier,
	abiGeneratedConnectionIdentifier,
	abiGeneratedChannelIdentifier abi.Event
)

// IBCErrorsRepository is a repository of custom errors defined in the ibc-solidity contracts.
var IBCErrorsRepository client.ErrorsRepository

func init() {
	var err error
	abiIBCHandler, err = abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		panic(err)
	}
	abiSendPacket = abiIBCHandler.Events["SendPacket"]
	abiWriteAcknowledgement = abiIBCHandler.Events["WriteAcknowledgement"]
	abiGeneratedClientIdentifier = abiIBCHandler.Events["GeneratedClientIdentifier"]
	abiGeneratedConnectionIdentifier = abiIBCHandler.Events["GeneratedConnectionIdentifier"]
	abiGeneratedChannelIdentifier = abiIBCHandler.Events["GeneratedChannelIdentifier"]

	abiICS20Bank, err := abi.JSON(strings.NewReader(ics20bank.Ics20bankABI))
	if err != nil {
		panic(err)
	}
	abiICS20TransferBank, err := abi.JSON(strings.NewReader(ics20transferbank.Ics20transferbankABI))
	if err != nil {
		panic(err)
	}
	abiQBFTClient, err := abi.JSON(strings.NewReader(qbftclient.QbftclientABI))
	if err != nil {
		panic(err)
	}
	IBCErrorsRepository = client.NewErrorsRepository()
	addErrorsToRepository(abiIBCHandler.Errors, IBCErrorsRepository)
	addErrorsToRepository(abiICS20Bank.Errors, IBCErrorsRepository)
	addErrorsToRepository(abiICS20TransferBank.Errors, IBCErrorsRepository)
	addErrorsToRepository(abiQBFTClient.Errors, IBCErrorsRepository)
}

type Chain struct {
	t *testing.T

	chainID       *big.Int
	client        *client.ETHClient
	lc            *LightClient
	delayPeriod   uint64 // nano second
	lcAddr        common.Address
	consensusType chains.ConsensusType

	mnemonic string
	keys     map[uint32]*ecdsa.PrivateKey

	// isAutoMining is true if the chain generates new blocks automatically
	isAutoMining bool

	// startBlockNumber is the block number when the chain instance is created
	// each event query should specify this as `FromBlock`
	startBlockNumber *big.Int

	ContractConfig ContractConfig

	// Core Modules
	IBCHandler ibchandler.Ibchandler

	// App Modules
	ERC20         erc20.Erc20
	ICS20Transfer ics20transferbank.Ics20transferbank
	ICS20Bank     ics20bank.Ics20bank
	IBCMockApp    ibcmockapp.Ibcmockapp

	// Input data for light client
	LatestLCInputData *LightClientInputData

	// IBC specific helpers
	ClientIDs   []string          // ClientID's used on this chain
	Connections []*TestConnection // track connectionID's created for this chain
}

func NewChain(t *testing.T, client *client.ETHClient, lc *LightClient, isAutoMining bool) *Chain {
	mnemonic := os.Getenv("TEST_MNEMONIC")
	if mnemonic == "" {
		t.Fatal("environ variable 'TEST_MNEMONIC' is empty")
	}
	logDir := os.Getenv("TEST_BROADCAST_LOG_DIR")
	if logDir == "" {
		t.Fatal("environ variable 'TEST_BROADCAST_LOG_DIR' is empty")
	}
	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	config, err := buildContractConfigFromBroadcastLog(filepath.Join(logDir, chainID.String(), "run-latest.json"))
	if err != nil {
		t.Fatal(err)
	}
	ibcHandler, err := ibchandler.NewIbchandler(config.IBCHandlerAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	ibcMockApp, err := ibcmockapp.NewIbcmockapp(config.IBCMockAppAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	erc20_, err := erc20.NewErc20(config.ERC20TokenAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	ics20transfer, err := ics20transferbank.NewIcs20transferbank(config.ICS20TransferBankAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	ics20bank, err := ics20bank.NewIcs20bank(config.ICS20BankAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	startBlockNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	chain := &Chain{
		t:             t,
		client:        client,
		chainID:       chainID,
		lc:            lc,
		consensusType: lc.consensusType,
		delayPeriod:   DefaultDelayPeriod,

		mnemonic:         mnemonic,
		ContractConfig:   *config,
		keys:             make(map[uint32]*ecdsa.PrivateKey),
		isAutoMining:     isAutoMining,
		startBlockNumber: big.NewInt(int64(startBlockNumber)),

		IBCHandler: *ibcHandler,

		ERC20:         *erc20_,
		ICS20Transfer: *ics20transfer,
		ICS20Bank:     *ics20bank,
		IBCMockApp:    *ibcMockApp,
	}

	lcAddr, err := ibcHandler.GetClientByType(chain.CallOpts(context.TODO(), RelayerKeyIndex), ibcclient.BesuQBFTClient)
	if err != nil {
		t.Fatal(err)
	}
	chain.lcAddr = lcAddr
	return chain
}

func (chain *Chain) SetDelayPeriod(delayPeriod uint64) {
	chain.delayPeriod = delayPeriod
}

func (chain *Chain) GetDelayPeriod() uint64 {
	return chain.delayPeriod
}

func (chain *Chain) Client() *client.ETHClient {
	return chain.client
}

func (chain *Chain) TxOpts(ctx context.Context, index uint32) *bind.TransactOpts {
	return makeGenTxOpts(chain.chainID, chain.prvKey(index))(ctx)
}

func (chain *Chain) CallOpts(ctx context.Context, index uint32) *bind.CallOpts {
	opts := chain.TxOpts(ctx, index)
	return &bind.CallOpts{
		From:    opts.From,
		Context: opts.Context,
	}
}

func (chain *Chain) prvKey(index uint32) *ecdsa.PrivateKey {
	key, ok := chain.keys[index]
	if ok {
		return key
	}
	key, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(chain.mnemonic, fmt.Sprintf("m/44'/60'/0'/0/%v", index))
	if err != nil {
		panic(err)
	}
	chain.keys[index] = key
	return key
}

func (chain *Chain) ChainIDU256() []byte {
	var chainID [32]byte
	chain.chainID.FillBytes(chainID[:])
	return chainID[:]
}

func (chain *Chain) GetCommitmentPrefix() []byte {
	return []byte(DefaultPrefix)
}

func (chain *Chain) GetQBFTClientState(clientID string) *qbftclienttypes.ClientState {
	ctx := context.Background()
	bz, found, err := chain.IBCHandler.GetClientState(chain.CallOpts(ctx, RelayerKeyIndex), clientID)
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("clientState not found")
	}
	var cs qbftclienttypes.ClientState
	if err := UnmarshalWithAny(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetQBFTConsensusState(clientID string, height ibcclient.Height) *qbftclienttypes.ConsensusState {
	ctx := context.Background()
	bz, found, err := chain.IBCHandler.GetConsensusState(chain.CallOpts(ctx, RelayerKeyIndex), clientID, ibchandler.HeightData(height))
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("consensusState not found")
	}
	var cs qbftclienttypes.ConsensusState
	if err := UnmarshalWithAny(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetLightClientInputData(counterparty *Chain, counterpartyClientID string, storageKeys [][]byte, height *big.Int) (*LightClientInputData, error) {
	if height == nil {
		height = counterparty.GetQBFTClientState(counterpartyClientID).LatestHeight.ToBN()
	}
	return chain.lc.GenerateInputData(
		context.Background(),
		chain.ContractConfig.IBCHandlerAddress,
		storageKeys,
		height,
	)
}

func (chain *Chain) ConstructQBFTMsgCreateClient(counterparty *Chain) ibchandler.IIBCClientMsgCreateClient {
	clientState := qbftclienttypes.ClientState{
		ChainId:         counterparty.ChainIDU256(),
		IbcStoreAddress: counterparty.ContractConfig.IBCHandlerAddress.Bytes(),
		LatestHeight:    ibcclient.NewHeightFromBN(counterparty.LastHeader().Number),
		TrustingPeriod:  DefaultTrustPeriod,
		MaxClockDrift:   DefaultMaxClockDrift,
	}
	consensusState := qbftclienttypes.ConsensusState{
		Timestamp:  counterparty.LastHeader().Time,
		Root:       counterparty.LastHeader().Root.Bytes(),
		Validators: counterparty.LatestLCInputData.Validators(),
	}
	clientStateBytes, err := MarshalWithAny(&clientState)
	if err != nil {
		panic(err)
	}
	consensusStateBytes, err := MarshalWithAny(&consensusState)
	if err != nil {
		panic(err)
	}
	return ibchandler.IIBCClientMsgCreateClient{
		ClientType:          ibcclient.BesuQBFTClient,
		ProtoClientState:    clientStateBytes,
		ProtoConsensusState: consensusStateBytes,
	}
}

func (chain *Chain) ConstructQBFTMsgUpdateClient(counterparty *Chain, clientID string) ibchandler.IIBCClientMsgUpdateClient {
	trustedHeight := chain.GetQBFTClientState(clientID).LatestHeight
	cs := counterparty.LatestLCInputData
	var header = qbftclienttypes.Header{
		BesuHeaderRlp:     cs.SealingHeaderRLP(chain.consensusType),
		Seals:             cs.CommitSeals,
		TrustedHeight:     trustedHeight,
		AccountStateProof: cs.MembershipProof().AccountProofRLP,
	}
	bz, err := MarshalWithAny(&header)
	if err != nil {
		panic(err)
	}
	return ibchandler.IIBCClientMsgUpdateClient{
		ClientId:           clientID,
		ProtoClientMessage: bz,
	}
}

func (chain *Chain) UpdateLCInputData() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		data, err := chain.lc.GenerateInputData(ctx, chain.ContractConfig.IBCHandlerAddress, nil, nil)
		if err != nil {
			panic(err)
		}
		if chain.LatestLCInputData == nil || data.Header().Number.Cmp(chain.LastHeader().Number) == 1 {
			chain.LatestLCInputData = data
			return
		} else {
			continue
		}
	}
}

func (chain *Chain) CreateQBFTClient(ctx context.Context, counterparty *Chain) (string, error) {
	msg := chain.ConstructQBFTMsgCreateClient(counterparty)
	if err := chain.WaitIfNoError(ctx, "IBCHandler::CreateClient")(
		chain.IBCHandler.CreateClient(chain.TxOpts(ctx, RelayerKeyIndex), msg),
	); err != nil {
		return "", fmt.Errorf("CreateQBFTClient: %w", err)
	}
	return chain.GetLastGeneratedClientID(ctx)
}

func (chain *Chain) UpdateQBFTClient(ctx context.Context, counterparty *Chain, clientID string, updateCommitment bool) error {
	msg := chain.ConstructQBFTMsgUpdateClient(counterparty, clientID)
	return chain.updateClient(ctx, msg, updateCommitment)
}

func (chain *Chain) updateClient(ctx context.Context, msg ibchandler.IIBCClientMsgUpdateClient, updateCommitment bool) error {
	if updateCommitment {
		return chain.WaitIfNoError(ctx, "IBCHandler::UpdateClient")(
			chain.IBCHandler.UpdateClient(chain.TxOpts(ctx, RelayerKeyIndex), msg),
		)
	} else {
		lcAddr, fnID, args, err := chain.IBCHandler.RouteUpdateClient(chain.CallOpts(ctx, RelayerKeyIndex), msg)
		if err != nil {
			return err
		}
		if lcAddr != chain.lcAddr {
			return fmt.Errorf("invalid light client address: expected=%v actual=%v", chain.lcAddr, lcAddr)
		}
		calldata := append(fnID[:], args...)
		return chain.WaitIfNoError(ctx, "IBCHandler::UpdateClient")(
			bind.NewBoundContract(lcAddr, abi.ABI{}, chain.client, chain.client, chain.client).RawTransact(chain.TxOpts(ctx, RelayerKeyIndex), calldata),
		)
	}
}

func (chain *Chain) ConnectionOpenInit(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) (string, error) {
	if err := chain.WaitIfNoError(ctx, "IBCHandler::ConnectionOpenInit")(
		chain.IBCHandler.ConnectionOpenInit(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCConnectionMsgConnectionOpenInit{
				ClientId: connection.ClientID,
				Counterparty: ibchandler.CounterpartyData{
					ClientId:     connection.CounterpartyClientID,
					ConnectionId: "",
					Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
				},
				Version:     ibchandler.VersionData{},
				DelayPeriod: chain.delayPeriod,
			},
		),
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedConnectionID(ctx)
}

func (chain *Chain) ConnectionOpenTry(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) (string, error) {
	proofConnection, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return "", err
	}
	clientStateBytes, latestHeight, proofClient, err := counterparty.QueryClientStateProof(chain, connection.ClientID, counterpartyConnection.ClientID, proofConnection.Height.ToBN())
	if err != nil {
		return "", err
	}
	consensusStateBytes, proofConsensus, err := counterparty.QueryConsensusStateProof(chain, connection.ClientID, counterpartyConnection.ClientID, latestHeight, proofConnection.Height.ToBN())
	if err != nil {
		return "", err
	}
	if err := chain.WaitIfNoError(ctx, "IBCHandler::ConnectionOpenTry")(
		chain.IBCHandler.ConnectionOpenTry(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCConnectionMsgConnectionOpenTry{
				Counterparty: ibchandler.CounterpartyData{
					ClientId:     counterpartyConnection.ClientID,
					ConnectionId: counterpartyConnection.ID,
					Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
				},
				DelayPeriod:      chain.delayPeriod,
				ClientId:         connection.ClientID,
				ClientStateBytes: clientStateBytes,
				CounterpartyVersions: []ibchandler.VersionData{
					{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
				},
				ProofHeight:             proofConnection.Height.ToCallData(),
				ConsensusHeight:         latestHeight.ToCallData(),
				ProofInit:               proofConnection.Data,
				ProofClient:             proofClient.Data,
				ProofConsensus:          proofConsensus.Data,
				HostConsensusStateProof: consensusStateBytes,
			},
		),
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedConnectionID(ctx)
}

// ConnectionOpenAck will construct and execute a MsgConnectionOpenAck.
func (chain *Chain) ConnectionOpenAck(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proofConnection, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return err
	}
	clientStateBytes, latestHeight, proofClient, err := counterparty.QueryClientStateProof(chain, connection.ClientID, counterpartyConnection.ClientID, proofConnection.Height.ToBN())
	if err != nil {
		return err
	}
	consensusStateBytes, proofConsensus, err := counterparty.QueryConsensusStateProof(chain, connection.ClientID, counterpartyConnection.ClientID, latestHeight, proofConnection.Height.ToBN())
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::ConnectionOpenAck")(
		chain.IBCHandler.ConnectionOpenAck(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCConnectionMsgConnectionOpenAck{
				ConnectionId:             connection.ID,
				CounterpartyConnectionId: counterpartyConnection.ID,
				ClientStateBytes:         clientStateBytes,
				Version:                  ibchandler.VersionData{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
				ProofHeight:              proofConnection.Height.ToCallData(),
				ConsensusHeight:          latestHeight.ToCallData(),
				ProofTry:                 proofConnection.Data,
				ProofClient:              proofClient.Data,
				ProofConsensus:           proofConsensus.Data,
				HostConsensusStateProof:  consensusStateBytes,
			},
		),
	)
}

func (chain *Chain) ConnectionOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proof, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::ConnectionOpenConfirm")(
		chain.IBCHandler.ConnectionOpenConfirm(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCConnectionMsgConnectionOpenConfirm{
				ConnectionId: connection.ID,
				ProofAck:     proof.Data,
				ProofHeight:  proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) ChannelOpenInit(
	ctx context.Context,
	ch, counterparty TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) (string, error) {
	if err := chain.WaitIfNoError(ctx, "IBCHandler::ChannelOpenInit")(
		chain.IBCHandler.ChannelOpenInit(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelOpenInit{
				PortId: ch.PortID,
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
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedChannelID(ctx)
}

func (chain *Chain) ChannelOpenTry(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) (string, error) {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return "", err
	}
	if err := chain.WaitIfNoError(ctx, "IBCHandler::ChannelOpenTry")(
		chain.IBCHandler.ChannelOpenTry(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelOpenTry{
				PortId: ch.PortID,
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
				ProofHeight:         proof.Height.ToCallData(),
			},
		),
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedChannelID(ctx)
}

func (chain *Chain) ChannelOpenAck(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::ChannelOpenAck")(
		chain.IBCHandler.ChannelOpenAck(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelOpenAck{
				PortId:                ch.PortID,
				ChannelId:             ch.ID,
				CounterpartyVersion:   counterpartyCh.Version,
				CounterpartyChannelId: counterpartyCh.ID,
				ProofTry:              proof.Data,
				ProofHeight:           proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) ChannelOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::ChannelOpenConfirm")(
		chain.IBCHandler.ChannelOpenConfirm(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelOpenConfirm{
				PortId:      ch.PortID,
				ChannelId:   ch.ID,
				ProofAck:    proof.Data,
				ProofHeight: proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) ChannelCloseInit(
	ctx context.Context,
	ch TestChannel,
) error {
	return chain.WaitIfNoError(ctx, "IBCHandler::ChannelCloseInit")(
		chain.IBCHandler.ChannelCloseInit(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelCloseInit{
				PortId:    ch.PortID,
				ChannelId: ch.ID,
			},
		),
	)
}

func (chain *Chain) ChannelCloseConfirm(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::ChannelCloseConfirm")(
		chain.IBCHandler.ChannelCloseConfirm(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelHandshakeMsgChannelCloseConfirm{
				PortId:      ch.PortID,
				ChannelId:   ch.ID,
				ProofInit:   proof.Data,
				ProofHeight: proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) SendPacket(
	ctx context.Context,
	packet ibchandler.Packet,
) error {
	return chain.WaitIfNoError(ctx, "IBCHandler::SendPacket")(
		chain.IBCHandler.SendPacket(
			chain.TxOpts(ctx, RelayerKeyIndex),
			packet.SourcePort,
			packet.SourceChannel,
			ibchandler.HeightData(packet.TimeoutHeight),
			packet.TimeoutTimestamp,
			packet.Data,
		),
	)
}

func (chain *Chain) HandlePacketRecv(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet ibchandler.Packet,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, commitment.PacketCommitmentSlot(packet.SourcePort, packet.SourceChannel, packet.Sequence), nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::RecvPacket")(
		chain.IBCHandler.RecvPacket(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelRecvPacketMsgPacketRecv{
				Packet:      PacketToCallData(packet),
				Proof:       proof.Data,
				ProofHeight: proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) HandlePacketAcknowledgement(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet ibchandler.Packet,
	acknowledgement []byte,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, commitment.PacketAcknowledgementCommitmentSlot(packet.DestinationPort, packet.DestinationChannel, packet.Sequence), nil)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::AcknowledgePacket")(
		chain.IBCHandler.AcknowledgePacket(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelAcknowledgePacketMsgPacketAcknowledgement{
				Packet:          PacketToCallData(packet),
				Acknowledgement: acknowledgement,
				Proof:           proof.Data,
				ProofHeight:     proof.Height.ToCallData(),
			},
		),
	)
}

func (chain *Chain) TimeoutPacket(
	ctx context.Context,
	packet ibchandler.Packet,
	counterparty *Chain,
	channel TestChannel,
	counterpartyChannel TestChannel,
) error {
	counterpartyCh, found, err := counterparty.IBCHandler.GetChannel(
		counterparty.CallOpts(ctx, RelayerKeyIndex),
		packet.DestinationPort,
		packet.DestinationChannel,
	)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("channel not found: port=%v channel=%v", packet.DestinationPort, packet.DestinationChannel)
	}
	if counterpartyCh.State == uint8(channeltypes.CLOSED) {
		return fmt.Errorf("channel is closed: port=%v channel=%v", packet.DestinationPort, packet.DestinationChannel)
	}

	var proof []byte
	proofHeight := counterparty.LatestLCInputData.Header().Number
	if counterpartyCh.Ordering == uint8(channeltypes.ORDERED) {
		p, err := counterparty.QueryNextSequenceRecvProof(chain, counterpartyChannel.ClientID, counterpartyChannel, proofHeight)
		if err != nil {
			return err
		}
		proof = p.Data
	} else {
		rc, err := counterparty.IBCHandler.GetPacketReceipt(
			counterparty.CallOpts(ctx, RelayerKeyIndex),
			packet.DestinationPort, packet.DestinationChannel, packet.Sequence,
		)
		if err != nil {
			return err
		}
		if rc == 1 {
			return fmt.Errorf("packet receipt exists: port=%v channel=%v sequence=%v", packet.DestinationPort, packet.DestinationChannel, packet.Sequence)
		}
		p, err := counterparty.QueryPacketReceiptProof(chain, channel.ClientID, packet, proofHeight)
		if err != nil {
			return err
		}
		proof = p.Data
	}

	nextSequenceRecv, err := counterparty.IBCHandler.GetNextSequenceRecv(
		counterparty.CallOpts(ctx, RelayerKeyIndex),
		counterpartyChannel.PortID, counterpartyChannel.ID,
	)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::TimeoutPacket")(
		chain.IBCHandler.TimeoutPacket(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelPacketTimeoutMsgTimeoutPacket{
				Packet: PacketToCallData(packet),
				Proof:  proof,
				ProofHeight: ibchandler.HeightData{
					RevisionNumber: 0,
					RevisionHeight: proofHeight.Uint64(),
				},
				NextSequenceRecv: nextSequenceRecv,
			},
		),
	)
}

func (chain *Chain) TimeoutOnClose(
	ctx context.Context,
	packet ibchandler.Packet,
	counterparty *Chain,
	channel TestChannel,
	counterpartyChannel TestChannel,
) error {
	counterpartyCh, found, err := counterparty.IBCHandler.GetChannel(
		counterparty.CallOpts(ctx, RelayerKeyIndex),
		packet.DestinationPort,
		packet.DestinationChannel,
	)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("channel not found: port=%v channel=%v", packet.DestinationPort, packet.DestinationChannel)
	}
	if counterpartyCh.State != uint8(channeltypes.CLOSED) {
		return fmt.Errorf("channel is not closed: port=%v channel=%v", packet.DestinationPort, packet.DestinationChannel)
	}

	proofHeight := counterparty.LatestLCInputData.Header().Number

	p, err := counterparty.QueryChannelProof(chain, channel.ClientID, counterpartyChannel, proofHeight)
	if err != nil {
		return err
	}
	proofClose := p.Data

	var proofUnreceived []byte
	if counterpartyCh.Ordering == uint8(channeltypes.ORDERED) {
		proof, err := counterparty.QueryNextSequenceRecvProof(chain, counterpartyChannel.ClientID, counterpartyChannel, proofHeight)
		if err != nil {
			return err
		}
		proofUnreceived = proof.Data
	} else {
		rc, err := counterparty.IBCHandler.GetPacketReceipt(
			counterparty.CallOpts(ctx, RelayerKeyIndex),
			packet.DestinationPort, packet.DestinationChannel, packet.Sequence,
		)
		if err != nil {
			return err
		}
		if rc == 1 {
			return fmt.Errorf("packet receipt exists: port=%v channel=%v sequence=%v", packet.DestinationPort, packet.DestinationChannel, packet.Sequence)
		}
		p, err := counterparty.QueryPacketReceiptProof(chain, channel.ClientID, packet, proofHeight)
		if err != nil {
			return err
		}
		proofUnreceived = p.Data
	}

	nextSequenceRecv, err := counterparty.IBCHandler.GetNextSequenceRecv(
		counterparty.CallOpts(ctx, RelayerKeyIndex),
		counterpartyChannel.PortID, counterpartyChannel.ID,
	)
	if err != nil {
		return err
	}
	return chain.WaitIfNoError(ctx, "IBCHandler::TimeoutOnClose")(
		chain.IBCHandler.TimeoutOnClose(
			chain.TxOpts(ctx, RelayerKeyIndex),
			ibchandler.IIBCChannelPacketTimeoutMsgTimeoutOnClose{
				Packet:          PacketToCallData(packet),
				ProofClose:      proofClose,
				ProofUnreceived: proofUnreceived,
				ProofHeight: ibchandler.HeightData{
					RevisionNumber: 0,
					RevisionHeight: proofHeight.Uint64(),
				},
				NextSequenceRecv: nextSequenceRecv,
			},
		),
	)
}

func (chain *Chain) SetExpectedTimePerBlock(
	ctx context.Context,
	callerIndex uint32,
	duration uint64,
) error {
	err := chain.WaitIfNoError(ctx, "IBCHandler::SetExpectedTimePerBlock")(
		chain.IBCHandler.SetExpectedTimePerBlock(
			chain.TxOpts(ctx, callerIndex),
			duration,
		),
	)
	if err != nil {
		return err
	}
	actual, err := chain.IBCHandler.GetExpectedTimePerBlock(
		chain.CallOpts(ctx, callerIndex),
	)
	if err != nil {
		return err
	}
	if actual != duration {
		return fmt.Errorf("expected=%v actual=%v", duration, actual)
	}
	return nil
}

func (chain *Chain) GetLastGeneratedClientID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedClientIdentifier)
}

func (chain *Chain) GetLastGeneratedConnectionID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedConnectionIdentifier)
}

func (chain *Chain) GetLastGeneratedChannelID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedChannelIdentifier)
}

func (chain *Chain) getLastID(ctx context.Context, event abi.Event) (string, error) {
	query := ethereum.FilterQuery{
		FromBlock: chain.startBlockNumber,
		Addresses: []common.Address{
			chain.ContractConfig.IBCHandlerAddress,
		},
		Topics: [][]common.Hash{{
			event.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx, query)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", errors.New("no items")
	}
	log := logs[len(logs)-1]
	values, err := event.Inputs.Unpack(log.Data)
	if err != nil {
		return "", err
	}
	return values[0].(string), nil
}

func (chain *Chain) GetLastSentPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
) (*ibchandler.Packet, error) {
	seq, err := chain.IBCHandler.GetNextSequenceSend(chain.CallOpts(ctx, RelayerKeyIndex), sourcePortID, sourceChannel)
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
) (*ibchandler.Packet, error) {
	channel, found, err := chain.IBCHandler.GetChannel(chain.CallOpts(ctx, RelayerKeyIndex), sourcePortID, sourceChannel)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("channel not found: sourcePortID=%v sourceChannel=%v", sourcePortID, sourceChannel)
	}

	query := ethereum.FilterQuery{
		FromBlock: chain.startBlockNumber,
		Addresses: []common.Address{
			chain.ContractConfig.IBCHandlerAddress,
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, log := range logs {
		var sendPacket ibchandler.IbchandlerSendPacket
		if err := abiIBCHandler.UnpackIntoInterface(&sendPacket, "SendPacket", log.Data); err != nil {
			return nil, err
		}
		if sendPacket.SourcePort != sourcePortID || sendPacket.SourceChannel != sourceChannel || sendPacket.Sequence != sequence {
			continue
		}
		return &ibchandler.Packet{
			Sequence:           sendPacket.Sequence,
			SourcePort:         sendPacket.SourcePort,
			SourceChannel:      sendPacket.SourceChannel,
			DestinationPort:    channel.Counterparty.PortId,
			DestinationChannel: channel.Counterparty.ChannelId,
			Data:               sendPacket.Data,
			TimeoutHeight:      ibchandler.HeightData{RevisionNumber: sendPacket.TimeoutHeight.RevisionNumber, RevisionHeight: sendPacket.TimeoutHeight.RevisionHeight},
			TimeoutTimestamp:   sendPacket.TimeoutTimestamp,
		}, nil
	}

	return nil, fmt.Errorf("packet not found: sourcePortID=%v sourceChannel=%v sequence=%v", sourcePortID, sourceChannel, sequence)
}

func (chain *Chain) FindAcknowledgement(
	ctx context.Context,
	portID string,
	channelID string,
	sequence uint64,
) ([]byte, error) {
	query := ethereum.FilterQuery{
		FromBlock: chain.startBlockNumber,
		Addresses: []common.Address{
			chain.ContractConfig.IBCHandlerAddress,
		},
		Topics: [][]common.Hash{{
			abiWriteAcknowledgement.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		var writeAcknowledgement ibchandler.IbchandlerWriteAcknowledgement
		if err := abiIBCHandler.UnpackIntoInterface(&writeAcknowledgement, "WriteAcknowledgement", log.Data); err != nil {
			return nil, err
		}
		if writeAcknowledgement.DestinationPortId != portID || writeAcknowledgement.DestinationChannel != channelID || writeAcknowledgement.Sequence != sequence {
			continue
		}
		return writeAcknowledgement.Acknowledgement, nil
	}
	return nil, fmt.Errorf("acknowledgement not found: portID=%v channelID=%v sequence=%v", portID, channelID, sequence)
}

func (chain *Chain) AdvanceBlockNumber(
	ctx context.Context,
	toBlockNumber uint64,
) error {
	for {
		blockNumber, err := chain.client.BlockNumber(ctx)
		if err != nil {
			return err
		}
		if blockNumber >= toBlockNumber {
			return nil
		}
		if chain.isAutoMining {
			time.Sleep(100 * time.Millisecond)
		} else {
			// execute a meaningless transaction to increase the block height
			err := chain.WaitIfNoError(ctx, "ERC20::Approve")(
				chain.ERC20.Approve(chain.TxOpts(ctx, RelayerKeyIndex), chain.ContractConfig.ICS20BankAddress, big.NewInt(0)),
			)
			if err != nil {
				return err
			}
		}
	}
}

func (chain *Chain) EnsureChannelState(
	ctx context.Context,
	portID string,
	channelID string,
	state channeltypes.Channel_State,
) error {
	channels, found, err := chain.IBCHandler.GetChannel(chain.CallOpts(ctx, RelayerKeyIndex), portID, channelID)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("channel not found")
	}
	if channels.State != uint8(state) {
		return fmt.Errorf("unexpected channel state: expected=%v actual=%v", state, channels.State)
	}
	return nil
}

func (chain *Chain) EnsurePacketCommitmentExistence(
	ctx context.Context,
	exists bool,
	portID string,
	channelID string,
	sequence uint64,
) error {
	commitment, err := chain.IBCHandler.GetCommitment(chain.CallOpts(ctx, RelayerKeyIndex), crypto.Keccak256Hash(host.PacketCommitmentKey(portID, channelID, sequence)))
	if err != nil {
		return err
	}
	if exists && commitment == [32]byte{} {
		return fmt.Errorf("packet commitment not found")
	} else if !exists && commitment != [32]byte{} {
		return fmt.Errorf("packet commitment found")
	}
	return nil
}

func PacketToCallData(packet ibchandler.Packet) ibchandler.Packet {
	return ibchandler.Packet{
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

// Querier

type Proof struct {
	Height ibcclient.Height
	Data   []byte
}

func (chain *Chain) QueryProof(counterparty *Chain, counterpartyClientID string, storageKey string, height *big.Int) (*Proof, error) {
	if !strings.HasPrefix(storageKey, "0x") {
		return nil, fmt.Errorf("storageKey must be hex string")
	}
	s, err := chain.GetLightClientInputData(counterparty, counterpartyClientID, [][]byte{[]byte(storageKey)}, height)
	if err != nil {
		return nil, err
	}
	return &Proof{
		Height: ibcclient.NewHeightFromBN(s.Header().Number),
		Data:   s.MembershipProof().StorageProofRLP[0],
	}, nil
}

func (chain *Chain) QueryClientStateProof(counterparty *Chain, clientID, counterpartyClientID string, height *big.Int) ([]byte, ibcclient.Height, *Proof, error) {
	cs, found, err := chain.IBCHandler.GetClientState(chain.CallOpts(context.Background(), RelayerKeyIndex), clientID)
	if err != nil {
		return nil, ibcclient.Height{}, nil, err
	} else if !found {
		return nil, ibcclient.Height{}, nil, fmt.Errorf("client not found: %v", counterpartyClientID)
	}
	proof, err := chain.QueryProof(counterparty, counterpartyClientID, commitment.ClientStateCommitmentSlot(clientID), height)
	if err != nil {
		return nil, ibcclient.Height{}, nil, err
	}
	return cs, chain.GetQBFTClientState(clientID).LatestHeight, proof, nil
}

func (chain *Chain) QueryConsensusStateProof(counterparty *Chain, clientID, counterpartyClientID string, consensusHeight ibcclient.Height, height *big.Int) ([]byte, *Proof, error) {
	cons, found, err := chain.IBCHandler.GetConsensusState(chain.CallOpts(context.Background(), RelayerKeyIndex), clientID, ibchandler.HeightData(consensusHeight))
	if err != nil {
		return nil, nil, err
	} else if !found {
		return nil, nil, fmt.Errorf("consensus state not found: %v", consensusHeight)
	}
	proof, err := chain.QueryProof(counterparty, counterpartyClientID, commitment.ConsensusStateCommitmentSlot(clientID, consensusHeight), height)
	if err != nil {
		return nil, nil, err
	}
	return cons, proof, nil
}

func (chain *Chain) QueryConnectionProof(counterparty *Chain, counterpartyClientID string, connectionID string, height *big.Int) (*Proof, error) {
	return chain.QueryProof(counterparty, counterpartyClientID, commitment.ConnectionStateCommitmentSlot(connectionID), height)
}

func (chain *Chain) QueryChannelProof(counterparty *Chain, counterpartyClientID string, channel TestChannel, height *big.Int) (*Proof, error) {
	return chain.QueryProof(counterparty, counterpartyClientID, commitment.ChannelStateCommitmentSlot(channel.PortID, channel.ID), height)
}

func (chain *Chain) QueryPacketReceiptProof(counterparty *Chain, counterpartyClientID string, packetFromCounterparty ibchandler.Packet, counterpartyHeight *big.Int) (*Proof, error) {
	return chain.QueryProof(counterparty, counterpartyClientID, commitment.PacketReceiptCommitmentSlot(
		packetFromCounterparty.DestinationPort,
		packetFromCounterparty.DestinationChannel,
		packetFromCounterparty.Sequence,
	), counterpartyHeight)
}

func (chain *Chain) QueryNextSequenceRecvProof(counterparty *Chain, counterpartyClientID string, channel TestChannel, height *big.Int) (*Proof, error) {
	return chain.QueryProof(counterparty, counterpartyClientID, commitment.NextSequenceRecvCommitmentSlot(channel.PortID, channel.ID), height)
}

func (chain *Chain) generateMockClientProof(height ibcclient.Height, path string, value []byte) []byte {
	var heightBz [16]byte
	binary.BigEndian.PutUint64(heightBz[:8], height.RevisionNumber)
	binary.BigEndian.PutUint64(heightBz[8:], height.RevisionHeight)
	hashPrefix := sha256.Sum256([]byte("ibc"))
	hashPath := sha256.Sum256([]byte(path))
	hashValue := sha256.Sum256(value)

	hash := append(heightBz[:], hashPrefix[:]...)
	hash = append(hash, hashPath[:]...)
	hash = append(hash, hashValue[:]...)
	h := sha256.Sum256(hash)
	return h[:]
}

func (chain *Chain) LastHeader() *gethtypes.Header {
	return chain.LatestLCInputData.Header()
}

func (chain *Chain) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction, txName string) error {
	rc, err := chain.Client().WaitForReceiptAndGet(ctx, tx)
	if err != nil {
		return err
	}
	if rc.Status == 1 {
		chain.t.Logf("tx=%v gasUsed=%v", txName, rc.GasUsed)
		return nil
	} else {
		return fmt.Errorf("failed to call transaction: tx=%v err='%v' rc='%v'", txName, err, rc)
	}
}

func (chain *Chain) WaitIfNoError(ctx context.Context, txName string) func(tx *gethtypes.Transaction, err error) error {
	return func(tx *gethtypes.Transaction, err error) error {
		if err != nil {
			return fmt.Errorf("failed to call transaction: tx=%v err='%v'", txName, err)
		}
		return chain.WaitForReceiptAndGet(ctx, tx, txName)
	}
}

// AddTestConnection appends a new TestConnection which contains references
// to the connection id, client id and counterparty client id.
func (chain *Chain) AddTestConnection(clientID, counterpartyClientID string) *TestConnection {
	conn := chain.ConstructNextTestConnection(clientID, counterpartyClientID)

	chain.Connections = append(chain.Connections, conn)
	return conn
}

// ConstructNextTestConnection constructs the next test connection to be
// created given a clientID and counterparty clientID.
func (chain *Chain) ConstructNextTestConnection(clientID, counterpartyClientID string) *TestConnection {
	return &TestConnection{
		ID:                   "",
		ClientID:             clientID,
		CounterpartyClientID: counterpartyClientID,
	}
}

// AddTestChannel appends a new TestChannel which contains references to the port and channel ID
// used for channel creation and interaction. See 'NextTestChannel' for channel ID naming format.
func (chain *Chain) AddTestChannel(conn *TestConnection, portID string, version string) TestChannel {
	channel := chain.NextTestChannel(conn, portID, version)
	conn.Channels = append(conn.Channels, channel)
	return channel
}

// NextTestChannel returns the next test channel to be created on this connection, but does not
// add it to the list of created channels. This function is expected to be used when the caller
// has not created the associated channel in app state, but would still like to refer to the
// non-existent channel usually to test for its non-existence.
//
// The port is passed in by the caller.
func (chain *Chain) NextTestChannel(conn *TestConnection, portID string, version string) TestChannel {
	return TestChannel{
		PortID:               portID,
		ID:                   "",
		ClientID:             conn.ClientID,
		CounterpartyClientID: conn.CounterpartyClientID,
		Version:              version,
	}
}

func makeGenTxOpts(chainID *big.Int, prv *ecdsa.PrivateKey) func(ctx context.Context) *bind.TransactOpts {
	signer := gethtypes.LatestSignerForChainID(chainID)
	addr := gethcrypto.PubkeyToAddress(prv.PublicKey)
	return func(ctx context.Context) *bind.TransactOpts {
		return &bind.TransactOpts{
			From: addr,
			// Set non-zero value to avoid call `estimateGas`
			// This allows we can extract the revert reason from the transaction receipt if the transaction fails.
			GasLimit: 6382056,
			Signer: func(address common.Address, tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
				if address != addr {
					return nil, errors.New("not authorized to sign this account")
				}
				signature, err := gethcrypto.Sign(signer.Hash(tx).Bytes(), prv)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
		}
	}
}

func addErrorsToRepository(errors map[string]abi.Error, repository client.ErrorsRepository) {
	for _, e := range errors {
		if err := repository.Add(e); err != nil {
			panic(err)
		}
	}
}
