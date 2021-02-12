// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcclient

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

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	ChainId              string
	ProvableStoreAddress []byte
	LatestHeight         uint64
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Timestamp  uint64
	Root       []byte
	Validators [][]byte
}

// IBCClientHeader is an auto generated low-level Go binding around an user-defined struct.
type IBCClientHeader struct {
	BesuHeaderRLPBytes []byte
	Seals              [][]byte
	TrustedHeight      uint64
	AccountStateProof  []byte
}

// IbcclientABI is the input ABI used to generate the binding from.
const IbcclientABI = "[{\"inputs\":[{\"internalType\":\"contractProvableStore\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"besuHeaderRLPBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"seals\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"trustedHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"accountStateProof\",\"type\":\"bytes\"}],\"internalType\":\"structIBCClient.Header\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"self\",\"type\":\"tuple\"}],\"name\":\"validateSelfClient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getSelfConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ignore0\",\"type\":\"bytes32\"}],\"name\":\"verifyMembershipAndGetLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"self\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"target\",\"type\":\"tuple\"}],\"name\":\"verifyClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"self\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientConsensusState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"self\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"verifyConnectionState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"self\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"verifyConnectionStateAndGet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Ibcclient is an auto generated Go binding around an Ethereum contract.
type Ibcclient struct {
	IbcclientCaller     // Read-only binding to the contract
	IbcclientTransactor // Write-only binding to the contract
	IbcclientFilterer   // Log filterer for contract events
}

// IbcclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcclientSession struct {
	Contract     *Ibcclient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcclientCallerSession struct {
	Contract *IbcclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IbcclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcclientTransactorSession struct {
	Contract     *IbcclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IbcclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcclientRaw struct {
	Contract *Ibcclient // Generic contract binding to access the raw methods on
}

// IbcclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcclientCallerRaw struct {
	Contract *IbcclientCaller // Generic read-only contract binding to access the raw methods on
}

// IbcclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcclientTransactorRaw struct {
	Contract *IbcclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcclient creates a new instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclient(address common.Address, backend bind.ContractBackend) (*Ibcclient, error) {
	contract, err := bindIbcclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcclient{IbcclientCaller: IbcclientCaller{contract: contract}, IbcclientTransactor: IbcclientTransactor{contract: contract}, IbcclientFilterer: IbcclientFilterer{contract: contract}}, nil
}

// NewIbcclientCaller creates a new read-only instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientCaller(address common.Address, caller bind.ContractCaller) (*IbcclientCaller, error) {
	contract, err := bindIbcclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcclientCaller{contract: contract}, nil
}

// NewIbcclientTransactor creates a new write-only instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcclientTransactor, error) {
	contract, err := bindIbcclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcclientTransactor{contract: contract}, nil
}

// NewIbcclientFilterer creates a new log filterer instance of Ibcclient, bound to a specific deployed contract.
func NewIbcclientFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcclientFilterer, error) {
	contract, err := bindIbcclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcclientFilterer{contract: contract}, nil
}

// bindIbcclient binds a generic wrapper to an already deployed contract.
func bindIbcclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcclient *IbcclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcclient.Contract.IbcclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcclient *IbcclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcclient.Contract.IbcclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcclient *IbcclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcclient.Contract.IbcclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcclient *IbcclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcclient *IbcclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcclient *IbcclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcclient.Contract.contract.Transact(opts, method, params...)
}

// GetSelfConsensusState is a free data retrieval call binding the contract method 0x90fc575c.
//
// Solidity: function getSelfConsensusState(uint64 height) pure returns((uint64,bytes,bytes[]), bool)
func (_Ibcclient *IbcclientCaller) GetSelfConsensusState(opts *bind.CallOpts, height uint64) (ConsensusStateData, bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "getSelfConsensusState", height)

	if err != nil {
		return *new(ConsensusStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConsensusStateData)).(*ConsensusStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetSelfConsensusState is a free data retrieval call binding the contract method 0x90fc575c.
//
// Solidity: function getSelfConsensusState(uint64 height) pure returns((uint64,bytes,bytes[]), bool)
func (_Ibcclient *IbcclientSession) GetSelfConsensusState(height uint64) (ConsensusStateData, bool, error) {
	return _Ibcclient.Contract.GetSelfConsensusState(&_Ibcclient.CallOpts, height)
}

// GetSelfConsensusState is a free data retrieval call binding the contract method 0x90fc575c.
//
// Solidity: function getSelfConsensusState(uint64 height) pure returns((uint64,bytes,bytes[]), bool)
func (_Ibcclient *IbcclientCallerSession) GetSelfConsensusState(height uint64) (ConsensusStateData, bool, error) {
	return _Ibcclient.Contract.GetSelfConsensusState(&_Ibcclient.CallOpts, height)
}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xebe9db81.
//
// Solidity: function validateSelfClient((string,bytes,uint64) self) pure returns(bool)
func (_Ibcclient *IbcclientCaller) ValidateSelfClient(opts *bind.CallOpts, self ClientStateData) (bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "validateSelfClient", self)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xebe9db81.
//
// Solidity: function validateSelfClient((string,bytes,uint64) self) pure returns(bool)
func (_Ibcclient *IbcclientSession) ValidateSelfClient(self ClientStateData) (bool, error) {
	return _Ibcclient.Contract.ValidateSelfClient(&_Ibcclient.CallOpts, self)
}

// ValidateSelfClient is a free data retrieval call binding the contract method 0xebe9db81.
//
// Solidity: function validateSelfClient((string,bytes,uint64) self) pure returns(bool)
func (_Ibcclient *IbcclientCallerSession) ValidateSelfClient(self ClientStateData) (bool, error) {
	return _Ibcclient.Contract.ValidateSelfClient(&_Ibcclient.CallOpts, self)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x91c2dbfe.
//
// Solidity: function verifyClientConsensusState((string,bytes,uint64) self, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibcclient *IbcclientCaller) VerifyClientConsensusState(opts *bind.CallOpts, self ClientStateData, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "verifyClientConsensusState", self, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x91c2dbfe.
//
// Solidity: function verifyClientConsensusState((string,bytes,uint64) self, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibcclient *IbcclientSession) VerifyClientConsensusState(self ClientStateData, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Ibcclient.Contract.VerifyClientConsensusState(&_Ibcclient.CallOpts, self, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x91c2dbfe.
//
// Solidity: function verifyClientConsensusState((string,bytes,uint64) self, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Ibcclient *IbcclientCallerSession) VerifyClientConsensusState(self ClientStateData, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Ibcclient.Contract.VerifyClientConsensusState(&_Ibcclient.CallOpts, self, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x445d2c67.
//
// Solidity: function verifyClientState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, (string,bytes,uint64) target) view returns(bool)
func (_Ibcclient *IbcclientCaller) VerifyClientState(opts *bind.CallOpts, self ClientStateData, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, target ClientStateData) (bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "verifyClientState", self, clientId, height, prefix, counterpartyClientIdentifier, proof, target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientState is a free data retrieval call binding the contract method 0x445d2c67.
//
// Solidity: function verifyClientState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, (string,bytes,uint64) target) view returns(bool)
func (_Ibcclient *IbcclientSession) VerifyClientState(self ClientStateData, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, target ClientStateData) (bool, error) {
	return _Ibcclient.Contract.VerifyClientState(&_Ibcclient.CallOpts, self, clientId, height, prefix, counterpartyClientIdentifier, proof, target)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x445d2c67.
//
// Solidity: function verifyClientState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, (string,bytes,uint64) target) view returns(bool)
func (_Ibcclient *IbcclientCallerSession) VerifyClientState(self ClientStateData, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, target ClientStateData) (bool, error) {
	return _Ibcclient.Contract.VerifyClientState(&_Ibcclient.CallOpts, self, clientId, height, prefix, counterpartyClientIdentifier, proof, target)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x75ea08cc.
//
// Solidity: function verifyConnectionState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibcclient *IbcclientCaller) VerifyConnectionState(opts *bind.CallOpts, self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "verifyConnectionState", self, clientId, height, prefix, proof, connectionId, connectionBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x75ea08cc.
//
// Solidity: function verifyConnectionState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibcclient *IbcclientSession) VerifyConnectionState(self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Ibcclient.Contract.VerifyConnectionState(&_Ibcclient.CallOpts, self, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0x75ea08cc.
//
// Solidity: function verifyConnectionState((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Ibcclient *IbcclientCallerSession) VerifyConnectionState(self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Ibcclient.Contract.VerifyConnectionState(&_Ibcclient.CallOpts, self, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionStateAndGet is a free data retrieval call binding the contract method 0x0cfcb909.
//
// Solidity: function verifyConnectionStateAndGet((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bytes32, bytes32)
func (_Ibcclient *IbcclientCaller) VerifyConnectionStateAndGet(opts *bind.CallOpts, self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "verifyConnectionStateAndGet", self, clientId, height, prefix, proof, connectionId, connectionBytes)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// VerifyConnectionStateAndGet is a free data retrieval call binding the contract method 0x0cfcb909.
//
// Solidity: function verifyConnectionStateAndGet((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bytes32, bytes32)
func (_Ibcclient *IbcclientSession) VerifyConnectionStateAndGet(self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) ([32]byte, [32]byte, error) {
	return _Ibcclient.Contract.VerifyConnectionStateAndGet(&_Ibcclient.CallOpts, self, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionStateAndGet is a free data retrieval call binding the contract method 0x0cfcb909.
//
// Solidity: function verifyConnectionStateAndGet((string,bytes,uint64) self, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bytes32, bytes32)
func (_Ibcclient *IbcclientCallerSession) VerifyConnectionStateAndGet(self ClientStateData, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) ([32]byte, [32]byte, error) {
	return _Ibcclient.Contract.VerifyConnectionStateAndGet(&_Ibcclient.CallOpts, self, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyMembershipAndGetLeaf is a free data retrieval call binding the contract method 0x9715513a.
//
// Solidity: function verifyMembershipAndGetLeaf(bytes proof, bytes32 root, bytes prefix, bytes32 slot, bytes32 ignore0) view returns(bytes32)
func (_Ibcclient *IbcclientCaller) VerifyMembershipAndGetLeaf(opts *bind.CallOpts, proof []byte, root [32]byte, prefix []byte, slot [32]byte, ignore0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ibcclient.contract.Call(opts, &out, "verifyMembershipAndGetLeaf", proof, root, prefix, slot, ignore0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyMembershipAndGetLeaf is a free data retrieval call binding the contract method 0x9715513a.
//
// Solidity: function verifyMembershipAndGetLeaf(bytes proof, bytes32 root, bytes prefix, bytes32 slot, bytes32 ignore0) view returns(bytes32)
func (_Ibcclient *IbcclientSession) VerifyMembershipAndGetLeaf(proof []byte, root [32]byte, prefix []byte, slot [32]byte, ignore0 [32]byte) ([32]byte, error) {
	return _Ibcclient.Contract.VerifyMembershipAndGetLeaf(&_Ibcclient.CallOpts, proof, root, prefix, slot, ignore0)
}

// VerifyMembershipAndGetLeaf is a free data retrieval call binding the contract method 0x9715513a.
//
// Solidity: function verifyMembershipAndGetLeaf(bytes proof, bytes32 root, bytes prefix, bytes32 slot, bytes32 ignore0) view returns(bytes32)
func (_Ibcclient *IbcclientCallerSession) VerifyMembershipAndGetLeaf(proof []byte, root [32]byte, prefix []byte, slot [32]byte, ignore0 [32]byte) ([32]byte, error) {
	return _Ibcclient.Contract.VerifyMembershipAndGetLeaf(&_Ibcclient.CallOpts, proof, root, prefix, slot, ignore0)
}

// CreateClient is a paid mutator transaction binding the contract method 0x5f189e56.
//
// Solidity: function createClient(string clientId, (string,bytes,uint64) clientState, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcclient *IbcclientTransactor) CreateClient(opts *bind.TransactOpts, clientId string, clientState ClientStateData, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "createClient", clientId, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x5f189e56.
//
// Solidity: function createClient(string clientId, (string,bytes,uint64) clientState, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcclient *IbcclientSession) CreateClient(clientId string, clientState ClientStateData, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcclient.Contract.CreateClient(&_Ibcclient.TransactOpts, clientId, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x5f189e56.
//
// Solidity: function createClient(string clientId, (string,bytes,uint64) clientState, (uint64,bytes,bytes[]) consensusState) returns()
func (_Ibcclient *IbcclientTransactorSession) CreateClient(clientId string, clientState ClientStateData, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Ibcclient.Contract.CreateClient(&_Ibcclient.TransactOpts, clientId, clientState, consensusState)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xe3ad0ac2.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],uint64,bytes) header) returns()
func (_Ibcclient *IbcclientTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, header IBCClientHeader) (*types.Transaction, error) {
	return _Ibcclient.contract.Transact(opts, "updateClient", clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xe3ad0ac2.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],uint64,bytes) header) returns()
func (_Ibcclient *IbcclientSession) UpdateClient(clientId string, header IBCClientHeader) (*types.Transaction, error) {
	return _Ibcclient.Contract.UpdateClient(&_Ibcclient.TransactOpts, clientId, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xe3ad0ac2.
//
// Solidity: function updateClient(string clientId, (bytes,bytes[],uint64,bytes) header) returns()
func (_Ibcclient *IbcclientTransactorSession) UpdateClient(clientId string, header IBCClientHeader) (*types.Transaction, error) {
	return _Ibcclient.Contract.UpdateClient(&_Ibcclient.TransactOpts, clientId, header)
}
