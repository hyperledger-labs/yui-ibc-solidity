// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20vouchers

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

// Ics20vouchersABI is the input ABI used to generate the binding from.
const Ics20vouchersABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ics20vouchers is an auto generated Go binding around an Ethereum contract.
type Ics20vouchers struct {
	Ics20vouchersCaller     // Read-only binding to the contract
	Ics20vouchersTransactor // Write-only binding to the contract
	Ics20vouchersFilterer   // Log filterer for contract events
}

// Ics20vouchersCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ics20vouchersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20vouchersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ics20vouchersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20vouchersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ics20vouchersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics20vouchersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ics20vouchersSession struct {
	Contract     *Ics20vouchers    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics20vouchersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ics20vouchersCallerSession struct {
	Contract *Ics20vouchersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Ics20vouchersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ics20vouchersTransactorSession struct {
	Contract     *Ics20vouchersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Ics20vouchersRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ics20vouchersRaw struct {
	Contract *Ics20vouchers // Generic contract binding to access the raw methods on
}

// Ics20vouchersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ics20vouchersCallerRaw struct {
	Contract *Ics20vouchersCaller // Generic read-only contract binding to access the raw methods on
}

// Ics20vouchersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ics20vouchersTransactorRaw struct {
	Contract *Ics20vouchersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIcs20vouchers creates a new instance of Ics20vouchers, bound to a specific deployed contract.
func NewIcs20vouchers(address common.Address, backend bind.ContractBackend) (*Ics20vouchers, error) {
	contract, err := bindIcs20vouchers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchers{Ics20vouchersCaller: Ics20vouchersCaller{contract: contract}, Ics20vouchersTransactor: Ics20vouchersTransactor{contract: contract}, Ics20vouchersFilterer: Ics20vouchersFilterer{contract: contract}}, nil
}

// NewIcs20vouchersCaller creates a new read-only instance of Ics20vouchers, bound to a specific deployed contract.
func NewIcs20vouchersCaller(address common.Address, caller bind.ContractCaller) (*Ics20vouchersCaller, error) {
	contract, err := bindIcs20vouchers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersCaller{contract: contract}, nil
}

// NewIcs20vouchersTransactor creates a new write-only instance of Ics20vouchers, bound to a specific deployed contract.
func NewIcs20vouchersTransactor(address common.Address, transactor bind.ContractTransactor) (*Ics20vouchersTransactor, error) {
	contract, err := bindIcs20vouchers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersTransactor{contract: contract}, nil
}

// NewIcs20vouchersFilterer creates a new log filterer instance of Ics20vouchers, bound to a specific deployed contract.
func NewIcs20vouchersFilterer(address common.Address, filterer bind.ContractFilterer) (*Ics20vouchersFilterer, error) {
	contract, err := bindIcs20vouchers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersFilterer{contract: contract}, nil
}

// bindIcs20vouchers binds a generic wrapper to an already deployed contract.
func bindIcs20vouchers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ics20vouchersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics20vouchers *Ics20vouchersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics20vouchers.Contract.Ics20vouchersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics20vouchers *Ics20vouchersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.Ics20vouchersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics20vouchers *Ics20vouchersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.Ics20vouchersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics20vouchers *Ics20vouchersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics20vouchers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics20vouchers *Ics20vouchersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics20vouchers *Ics20vouchersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersSession) ADMINROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.ADMINROLE(&_Ics20vouchers.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCallerSession) ADMINROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.ADMINROLE(&_Ics20vouchers.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.DEFAULTADMINROLE(&_Ics20vouchers.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.DEFAULTADMINROLE(&_Ics20vouchers.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersSession) OPERATORROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.OPERATORROLE(&_Ics20vouchers.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Ics20vouchers.Contract.OPERATORROLE(&_Ics20vouchers.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x8606f905.
//
// Solidity: function balanceOf(address account, bytes id) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id []byte) (*big.Int, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x8606f905.
//
// Solidity: function balanceOf(address account, bytes id) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersSession) BalanceOf(account common.Address, id []byte) (*big.Int, error) {
	return _Ics20vouchers.Contract.BalanceOf(&_Ics20vouchers.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x8606f905.
//
// Solidity: function balanceOf(address account, bytes id) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersCallerSession) BalanceOf(account common.Address, id []byte) (*big.Int, error) {
	return _Ics20vouchers.Contract.BalanceOf(&_Ics20vouchers.CallOpts, account, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ics20vouchers.Contract.GetRoleAdmin(&_Ics20vouchers.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ics20vouchers *Ics20vouchersCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ics20vouchers.Contract.GetRoleAdmin(&_Ics20vouchers.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ics20vouchers *Ics20vouchersCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ics20vouchers *Ics20vouchersSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ics20vouchers.Contract.GetRoleMember(&_Ics20vouchers.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ics20vouchers *Ics20vouchersCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ics20vouchers.Contract.GetRoleMember(&_Ics20vouchers.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ics20vouchers.Contract.GetRoleMemberCount(&_Ics20vouchers.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ics20vouchers *Ics20vouchersCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ics20vouchers.Contract.GetRoleMemberCount(&_Ics20vouchers.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ics20vouchers *Ics20vouchersCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ics20vouchers.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ics20vouchers *Ics20vouchersSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ics20vouchers.Contract.HasRole(&_Ics20vouchers.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ics20vouchers *Ics20vouchersCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ics20vouchers.Contract.HasRole(&_Ics20vouchers.CallOpts, role, account)
}

// BurnFrom is a paid mutator transaction binding the contract method 0xc8c7a33b.
//
// Solidity: function burnFrom(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "burnFrom", account, id, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0xc8c7a33b.
//
// Solidity: function burnFrom(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersSession) BurnFrom(account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.BurnFrom(&_Ics20vouchers.TransactOpts, account, id, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0xc8c7a33b.
//
// Solidity: function burnFrom(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) BurnFrom(account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.BurnFrom(&_Ics20vouchers.TransactOpts, account, id, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.GrantRole(&_Ics20vouchers.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.GrantRole(&_Ics20vouchers.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x775ed63d.
//
// Solidity: function mint(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) Mint(opts *bind.TransactOpts, account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "mint", account, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x775ed63d.
//
// Solidity: function mint(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersSession) Mint(account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.Mint(&_Ics20vouchers.TransactOpts, account, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x775ed63d.
//
// Solidity: function mint(address account, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) Mint(account common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.Mint(&_Ics20vouchers.TransactOpts, account, id, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.RenounceRole(&_Ics20vouchers.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.RenounceRole(&_Ics20vouchers.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.RevokeRole(&_Ics20vouchers.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.RevokeRole(&_Ics20vouchers.TransactOpts, role, account)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Ics20vouchers *Ics20vouchersSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.SetOperator(&_Ics20vouchers.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.SetOperator(&_Ics20vouchers.TransactOpts, operator)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f61389a.
//
// Solidity: function transferFrom(address from, address to, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.contract.Transact(opts, "transferFrom", from, to, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f61389a.
//
// Solidity: function transferFrom(address from, address to, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersSession) TransferFrom(from common.Address, to common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.TransferFrom(&_Ics20vouchers.TransactOpts, from, to, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f61389a.
//
// Solidity: function transferFrom(address from, address to, bytes id, uint256 amount) returns()
func (_Ics20vouchers *Ics20vouchersTransactorSession) TransferFrom(from common.Address, to common.Address, id []byte, amount *big.Int) (*types.Transaction, error) {
	return _Ics20vouchers.Contract.TransferFrom(&_Ics20vouchers.TransactOpts, from, to, id, amount)
}

// Ics20vouchersRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Ics20vouchers contract.
type Ics20vouchersRoleAdminChangedIterator struct {
	Event *Ics20vouchersRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics20vouchersRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics20vouchersRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics20vouchersRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics20vouchersRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics20vouchersRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics20vouchersRoleAdminChanged represents a RoleAdminChanged event raised by the Ics20vouchers contract.
type Ics20vouchersRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ics20vouchers *Ics20vouchersFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Ics20vouchersRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ics20vouchers.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersRoleAdminChangedIterator{contract: _Ics20vouchers.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ics20vouchers *Ics20vouchersFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Ics20vouchersRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ics20vouchers.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics20vouchersRoleAdminChanged)
				if err := _Ics20vouchers.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ics20vouchers *Ics20vouchersFilterer) ParseRoleAdminChanged(log types.Log) (*Ics20vouchersRoleAdminChanged, error) {
	event := new(Ics20vouchersRoleAdminChanged)
	if err := _Ics20vouchers.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ics20vouchersRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Ics20vouchers contract.
type Ics20vouchersRoleGrantedIterator struct {
	Event *Ics20vouchersRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics20vouchersRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics20vouchersRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics20vouchersRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics20vouchersRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics20vouchersRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics20vouchersRoleGranted represents a RoleGranted event raised by the Ics20vouchers contract.
type Ics20vouchersRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Ics20vouchersRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ics20vouchers.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersRoleGrantedIterator{contract: _Ics20vouchers.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Ics20vouchersRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ics20vouchers.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics20vouchersRoleGranted)
				if err := _Ics20vouchers.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) ParseRoleGranted(log types.Log) (*Ics20vouchersRoleGranted, error) {
	event := new(Ics20vouchersRoleGranted)
	if err := _Ics20vouchers.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ics20vouchersRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Ics20vouchers contract.
type Ics20vouchersRoleRevokedIterator struct {
	Event *Ics20vouchersRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics20vouchersRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics20vouchersRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics20vouchersRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics20vouchersRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics20vouchersRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics20vouchersRoleRevoked represents a RoleRevoked event raised by the Ics20vouchers contract.
type Ics20vouchersRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Ics20vouchersRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ics20vouchers.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Ics20vouchersRoleRevokedIterator{contract: _Ics20vouchers.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Ics20vouchersRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ics20vouchers.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics20vouchersRoleRevoked)
				if err := _Ics20vouchers.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ics20vouchers *Ics20vouchersFilterer) ParseRoleRevoked(log types.Log) (*Ics20vouchersRoleRevoked, error) {
	event := new(Ics20vouchersRoleRevoked)
	if err := _Ics20vouchers.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
