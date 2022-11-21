// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibchost

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

// IbchostABI is the input ABI used to generate the binding from.
const IbchostABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedChannelIdentifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedClientIdentifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"GeneratedConnectionIdentifier\",\"type\":\"event\"}]"

// Ibchost is an auto generated Go binding around an Ethereum contract.
type Ibchost struct {
	IbchostCaller     // Read-only binding to the contract
	IbchostTransactor // Write-only binding to the contract
	IbchostFilterer   // Log filterer for contract events
}

// IbchostCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbchostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbchostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbchostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbchostSession struct {
	Contract     *Ibchost          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbchostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbchostCallerSession struct {
	Contract *IbchostCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IbchostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbchostTransactorSession struct {
	Contract     *IbchostTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IbchostRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbchostRaw struct {
	Contract *Ibchost // Generic contract binding to access the raw methods on
}

// IbchostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbchostCallerRaw struct {
	Contract *IbchostCaller // Generic read-only contract binding to access the raw methods on
}

// IbchostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbchostTransactorRaw struct {
	Contract *IbchostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbchost creates a new instance of Ibchost, bound to a specific deployed contract.
func NewIbchost(address common.Address, backend bind.ContractBackend) (*Ibchost, error) {
	contract, err := bindIbchost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibchost{IbchostCaller: IbchostCaller{contract: contract}, IbchostTransactor: IbchostTransactor{contract: contract}, IbchostFilterer: IbchostFilterer{contract: contract}}, nil
}

// NewIbchostCaller creates a new read-only instance of Ibchost, bound to a specific deployed contract.
func NewIbchostCaller(address common.Address, caller bind.ContractCaller) (*IbchostCaller, error) {
	contract, err := bindIbchost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbchostCaller{contract: contract}, nil
}

// NewIbchostTransactor creates a new write-only instance of Ibchost, bound to a specific deployed contract.
func NewIbchostTransactor(address common.Address, transactor bind.ContractTransactor) (*IbchostTransactor, error) {
	contract, err := bindIbchost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbchostTransactor{contract: contract}, nil
}

// NewIbchostFilterer creates a new log filterer instance of Ibchost, bound to a specific deployed contract.
func NewIbchostFilterer(address common.Address, filterer bind.ContractFilterer) (*IbchostFilterer, error) {
	contract, err := bindIbchost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbchostFilterer{contract: contract}, nil
}

// bindIbchost binds a generic wrapper to an already deployed contract.
func bindIbchost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbchostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchost *IbchostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchost.Contract.IbchostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchost *IbchostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.Contract.IbchostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchost *IbchostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchost.Contract.IbchostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchost *IbchostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchost *IbchostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchost *IbchostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchost.Contract.contract.Transact(opts, method, params...)
}

// IbchostGeneratedChannelIdentifierIterator is returned from FilterGeneratedChannelIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedChannelIdentifier events raised by the Ibchost contract.
type IbchostGeneratedChannelIdentifierIterator struct {
	Event *IbchostGeneratedChannelIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedChannelIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedChannelIdentifier)
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
		it.Event = new(IbchostGeneratedChannelIdentifier)
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
func (it *IbchostGeneratedChannelIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedChannelIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedChannelIdentifier represents a GeneratedChannelIdentifier event raised by the Ibchost contract.
type IbchostGeneratedChannelIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedChannelIdentifier is a free log retrieval operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedChannelIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedChannelIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedChannelIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedChannelIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedChannelIdentifier is a free log subscription operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedChannelIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedChannelIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedChannelIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
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

// ParseGeneratedChannelIdentifier is a log parse operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedChannelIdentifier(log types.Log) (*IbchostGeneratedChannelIdentifier, error) {
	event := new(IbchostGeneratedChannelIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchostGeneratedClientIdentifierIterator is returned from FilterGeneratedClientIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedClientIdentifier events raised by the Ibchost contract.
type IbchostGeneratedClientIdentifierIterator struct {
	Event *IbchostGeneratedClientIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedClientIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedClientIdentifier)
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
		it.Event = new(IbchostGeneratedClientIdentifier)
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
func (it *IbchostGeneratedClientIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedClientIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedClientIdentifier represents a GeneratedClientIdentifier event raised by the Ibchost contract.
type IbchostGeneratedClientIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedClientIdentifier is a free log retrieval operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedClientIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedClientIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedClientIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedClientIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedClientIdentifier is a free log subscription operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedClientIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedClientIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedClientIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
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

// ParseGeneratedClientIdentifier is a log parse operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedClientIdentifier(log types.Log) (*IbchostGeneratedClientIdentifier, error) {
	event := new(IbchostGeneratedClientIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchostGeneratedConnectionIdentifierIterator is returned from FilterGeneratedConnectionIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedConnectionIdentifier events raised by the Ibchost contract.
type IbchostGeneratedConnectionIdentifierIterator struct {
	Event *IbchostGeneratedConnectionIdentifier // Event containing the contract specifics and raw log

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
func (it *IbchostGeneratedConnectionIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchostGeneratedConnectionIdentifier)
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
		it.Event = new(IbchostGeneratedConnectionIdentifier)
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
func (it *IbchostGeneratedConnectionIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchostGeneratedConnectionIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchostGeneratedConnectionIdentifier represents a GeneratedConnectionIdentifier event raised by the Ibchost contract.
type IbchostGeneratedConnectionIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedConnectionIdentifier is a free log retrieval operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) FilterGeneratedConnectionIdentifier(opts *bind.FilterOpts) (*IbchostGeneratedConnectionIdentifierIterator, error) {

	logs, sub, err := _Ibchost.contract.FilterLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchostGeneratedConnectionIdentifierIterator{contract: _Ibchost.contract, event: "GeneratedConnectionIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedConnectionIdentifier is a free log subscription operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) WatchGeneratedConnectionIdentifier(opts *bind.WatchOpts, sink chan<- *IbchostGeneratedConnectionIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchost.contract.WatchLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchostGeneratedConnectionIdentifier)
				if err := _Ibchost.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
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

// ParseGeneratedConnectionIdentifier is a log parse operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchost *IbchostFilterer) ParseGeneratedConnectionIdentifier(log types.Log) (*IbchostGeneratedConnectionIdentifier, error) {
	event := new(IbchostGeneratedConnectionIdentifier)
	if err := _Ibchost.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
