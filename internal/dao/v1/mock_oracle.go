// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v1

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MockOracleMetaData contains all meta data concerning the MockOracle contract.
var MockOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"price\",\"type\":\"bytes32\"}],\"name\":\"PriceUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MockOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use MockOracleMetaData.ABI instead.
var MockOracleABI = MockOracleMetaData.ABI

// MockOracle is an auto generated Go binding around an Ethereum contract.
type MockOracle struct {
	MockOracleCaller     // Read-only binding to the contract
	MockOracleTransactor // Write-only binding to the contract
	MockOracleFilterer   // Log filterer for contract events
}

// MockOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockOracleSession struct {
	Contract     *MockOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockOracleCallerSession struct {
	Contract *MockOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockOracleTransactorSession struct {
	Contract     *MockOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockOracleRaw struct {
	Contract *MockOracle // Generic contract binding to access the raw methods on
}

// MockOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockOracleCallerRaw struct {
	Contract *MockOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MockOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockOracleTransactorRaw struct {
	Contract *MockOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockOracle creates a new instance of MockOracle, bound to a specific deployed contract.
func NewMockOracle(address common.Address, backend bind.ContractBackend) (*MockOracle, error) {
	contract, err := bindMockOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockOracle{MockOracleCaller: MockOracleCaller{contract: contract}, MockOracleTransactor: MockOracleTransactor{contract: contract}, MockOracleFilterer: MockOracleFilterer{contract: contract}}, nil
}

// NewMockOracleCaller creates a new read-only instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleCaller(address common.Address, caller bind.ContractCaller) (*MockOracleCaller, error) {
	contract, err := bindMockOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleCaller{contract: contract}, nil
}

// NewMockOracleTransactor creates a new write-only instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MockOracleTransactor, error) {
	contract, err := bindMockOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleTransactor{contract: contract}, nil
}

// NewMockOracleFilterer creates a new log filterer instance of MockOracle, bound to a specific deployed contract.
func NewMockOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MockOracleFilterer, error) {
	contract, err := bindMockOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockOracleFilterer{contract: contract}, nil
}

// bindMockOracle binds a generic wrapper to an already deployed contract.
func bindMockOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracle *MockOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracle.Contract.MockOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracle *MockOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracle.Contract.MockOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracle *MockOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracle.Contract.MockOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracle *MockOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracle *MockOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracle *MockOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracle.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockOracle *MockOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockOracle *MockOracleSession) Owner() (common.Address, error) {
	return _MockOracle.Contract.Owner(&_MockOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockOracle *MockOracleCallerSession) Owner() (common.Address, error) {
	return _MockOracle.Contract.Owner(&_MockOracle.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_MockOracle *MockOracleCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var out []interface{}
	err := _MockOracle.contract.Call(opts, &out, "peek")

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_MockOracle *MockOracleSession) Peek() ([32]byte, bool, error) {
	return _MockOracle.Contract.Peek(&_MockOracle.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_MockOracle *MockOracleCallerSession) Peek() ([32]byte, bool, error) {
	return _MockOracle.Contract.Peek(&_MockOracle.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(bytes32)
func (_MockOracle *MockOracleCaller) Price(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockOracle.contract.Call(opts, &out, "price")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(bytes32)
func (_MockOracle *MockOracleSession) Price() ([32]byte, error) {
	return _MockOracle.Contract.Price(&_MockOracle.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(bytes32)
func (_MockOracle *MockOracleCallerSession) Price() ([32]byte, error) {
	return _MockOracle.Contract.Price(&_MockOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockOracle *MockOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockOracle *MockOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _MockOracle.Contract.RenounceOwnership(&_MockOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockOracle *MockOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MockOracle.Contract.RenounceOwnership(&_MockOracle.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_MockOracle *MockOracleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_MockOracle *MockOracleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _MockOracle.Contract.SetPrice(&_MockOracle.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_MockOracle *MockOracleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _MockOracle.Contract.SetPrice(&_MockOracle.TransactOpts, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockOracle *MockOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MockOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockOracle *MockOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MockOracle.Contract.TransferOwnership(&_MockOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockOracle *MockOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MockOracle.Contract.TransferOwnership(&_MockOracle.TransactOpts, newOwner)
}

// MockOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MockOracle contract.
type MockOracleOwnershipTransferredIterator struct {
	Event *MockOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MockOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOracleOwnershipTransferred)
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
		it.Event = new(MockOracleOwnershipTransferred)
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
func (it *MockOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOracleOwnershipTransferred represents a OwnershipTransferred event raised by the MockOracle contract.
type MockOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MockOracle *MockOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MockOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MockOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MockOracleOwnershipTransferredIterator{contract: _MockOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MockOracle *MockOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MockOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOracleOwnershipTransferred)
				if err := _MockOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MockOracle *MockOracleFilterer) ParseOwnershipTransferred(log types.Log) (*MockOracleOwnershipTransferred, error) {
	event := new(MockOracleOwnershipTransferred)
	if err := _MockOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockOraclePriceUpdateIterator is returned from FilterPriceUpdate and is used to iterate over the raw logs and unpacked data for PriceUpdate events raised by the MockOracle contract.
type MockOraclePriceUpdateIterator struct {
	Event *MockOraclePriceUpdate // Event containing the contract specifics and raw log

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
func (it *MockOraclePriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOraclePriceUpdate)
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
		it.Event = new(MockOraclePriceUpdate)
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
func (it *MockOraclePriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOraclePriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOraclePriceUpdate represents a PriceUpdate event raised by the MockOracle contract.
type MockOraclePriceUpdate struct {
	Price [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdate is a free log retrieval operation binding the contract event 0x3500614d792a368ab3c20285be59acb332a20632feb94ce03975d7ed7682c8d0.
//
// Solidity: event PriceUpdate(bytes32 price)
func (_MockOracle *MockOracleFilterer) FilterPriceUpdate(opts *bind.FilterOpts) (*MockOraclePriceUpdateIterator, error) {

	logs, sub, err := _MockOracle.contract.FilterLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return &MockOraclePriceUpdateIterator{contract: _MockOracle.contract, event: "PriceUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceUpdate is a free log subscription operation binding the contract event 0x3500614d792a368ab3c20285be59acb332a20632feb94ce03975d7ed7682c8d0.
//
// Solidity: event PriceUpdate(bytes32 price)
func (_MockOracle *MockOracleFilterer) WatchPriceUpdate(opts *bind.WatchOpts, sink chan<- *MockOraclePriceUpdate) (event.Subscription, error) {

	logs, sub, err := _MockOracle.contract.WatchLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOraclePriceUpdate)
				if err := _MockOracle.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
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

// ParsePriceUpdate is a log parse operation binding the contract event 0x3500614d792a368ab3c20285be59acb332a20632feb94ce03975d7ed7682c8d0.
//
// Solidity: event PriceUpdate(bytes32 price)
func (_MockOracle *MockOracleFilterer) ParsePriceUpdate(log types.Log) (*MockOraclePriceUpdate, error) {
	event := new(MockOraclePriceUpdate)
	if err := _MockOracle.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
