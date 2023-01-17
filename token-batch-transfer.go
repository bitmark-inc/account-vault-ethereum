// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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
	_ = abi.ConvertType
)

// TokenBatchTransferMetaData contains all meta data concerning the TokenBatchTransfer contract.
var TokenBatchTransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transferOperator\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokenHolders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenBatchTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBatchTransferMetaData.ABI instead.
var TokenBatchTransferABI = TokenBatchTransferMetaData.ABI

// TokenBatchTransfer is an auto generated Go binding around an Ethereum contract.
type TokenBatchTransfer struct {
	TokenBatchTransferCaller     // Read-only binding to the contract
	TokenBatchTransferTransactor // Write-only binding to the contract
	TokenBatchTransferFilterer   // Log filterer for contract events
}

// TokenBatchTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBatchTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBatchTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBatchTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBatchTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBatchTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBatchTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBatchTransferSession struct {
	Contract     *TokenBatchTransfer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenBatchTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBatchTransferCallerSession struct {
	Contract *TokenBatchTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenBatchTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBatchTransferTransactorSession struct {
	Contract     *TokenBatchTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenBatchTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBatchTransferRaw struct {
	Contract *TokenBatchTransfer // Generic contract binding to access the raw methods on
}

// TokenBatchTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBatchTransferCallerRaw struct {
	Contract *TokenBatchTransferCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBatchTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBatchTransferTransactorRaw struct {
	Contract *TokenBatchTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBatchTransfer creates a new instance of TokenBatchTransfer, bound to a specific deployed contract.
func NewTokenBatchTransfer(address common.Address, backend bind.ContractBackend) (*TokenBatchTransfer, error) {
	contract, err := bindTokenBatchTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransfer{TokenBatchTransferCaller: TokenBatchTransferCaller{contract: contract}, TokenBatchTransferTransactor: TokenBatchTransferTransactor{contract: contract}, TokenBatchTransferFilterer: TokenBatchTransferFilterer{contract: contract}}, nil
}

// NewTokenBatchTransferCaller creates a new read-only instance of TokenBatchTransfer, bound to a specific deployed contract.
func NewTokenBatchTransferCaller(address common.Address, caller bind.ContractCaller) (*TokenBatchTransferCaller, error) {
	contract, err := bindTokenBatchTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferCaller{contract: contract}, nil
}

// NewTokenBatchTransferTransactor creates a new write-only instance of TokenBatchTransfer, bound to a specific deployed contract.
func NewTokenBatchTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBatchTransferTransactor, error) {
	contract, err := bindTokenBatchTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferTransactor{contract: contract}, nil
}

// NewTokenBatchTransferFilterer creates a new log filterer instance of TokenBatchTransfer, bound to a specific deployed contract.
func NewTokenBatchTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBatchTransferFilterer, error) {
	contract, err := bindTokenBatchTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferFilterer{contract: contract}, nil
}

// bindTokenBatchTransfer binds a generic wrapper to an already deployed contract.
func bindTokenBatchTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenBatchTransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBatchTransfer *TokenBatchTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBatchTransfer.Contract.TokenBatchTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBatchTransfer *TokenBatchTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.TokenBatchTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBatchTransfer *TokenBatchTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.TokenBatchTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBatchTransfer *TokenBatchTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBatchTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBatchTransfer *TokenBatchTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBatchTransfer *TokenBatchTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBatchTransfer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferSession) Owner() (common.Address, error) {
	return _TokenBatchTransfer.Contract.Owner(&_TokenBatchTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCallerSession) Owner() (common.Address, error) {
	return _TokenBatchTransfer.Contract.Owner(&_TokenBatchTransfer.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBatchTransfer.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferSession) Token() (common.Address, error) {
	return _TokenBatchTransfer.Contract.Token(&_TokenBatchTransfer.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCallerSession) Token() (common.Address, error) {
	return _TokenBatchTransfer.Contract.Token(&_TokenBatchTransfer.CallOpts)
}

// TransferOperator is a free data retrieval call binding the contract method 0xd0835895.
//
// Solidity: function transferOperator() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCaller) TransferOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBatchTransfer.contract.Call(opts, &out, "transferOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferOperator is a free data retrieval call binding the contract method 0xd0835895.
//
// Solidity: function transferOperator() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferSession) TransferOperator() (common.Address, error) {
	return _TokenBatchTransfer.Contract.TransferOperator(&_TokenBatchTransfer.CallOpts)
}

// TransferOperator is a free data retrieval call binding the contract method 0xd0835895.
//
// Solidity: function transferOperator() view returns(address)
func (_TokenBatchTransfer *TokenBatchTransferCallerSession) TransferOperator() (common.Address, error) {
	return _TokenBatchTransfer.Contract.TransferOperator(&_TokenBatchTransfer.CallOpts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x4885b254.
//
// Solidity: function batchTransferFrom(address tokenOwner, address[] tokenHolders, uint256[] amounts) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactor) BatchTransferFrom(opts *bind.TransactOpts, tokenOwner common.Address, tokenHolders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.contract.Transact(opts, "batchTransferFrom", tokenOwner, tokenHolders, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x4885b254.
//
// Solidity: function batchTransferFrom(address tokenOwner, address[] tokenHolders, uint256[] amounts) returns()
func (_TokenBatchTransfer *TokenBatchTransferSession) BatchTransferFrom(tokenOwner common.Address, tokenHolders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.BatchTransferFrom(&_TokenBatchTransfer.TransactOpts, tokenOwner, tokenHolders, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x4885b254.
//
// Solidity: function batchTransferFrom(address tokenOwner, address[] tokenHolders, uint256[] amounts) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactorSession) BatchTransferFrom(tokenOwner common.Address, tokenHolders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.BatchTransferFrom(&_TokenBatchTransfer.TransactOpts, tokenOwner, tokenHolders, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBatchTransfer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBatchTransfer *TokenBatchTransferSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.RenounceOwnership(&_TokenBatchTransfer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.RenounceOwnership(&_TokenBatchTransfer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBatchTransfer *TokenBatchTransferSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.TransferOwnership(&_TokenBatchTransfer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.TransferOwnership(&_TokenBatchTransfer.TransactOpts, newOwner)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address newOperator) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactor) UpdateOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.contract.Transact(opts, "updateOperator", newOperator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address newOperator) returns()
func (_TokenBatchTransfer *TokenBatchTransferSession) UpdateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.UpdateOperator(&_TokenBatchTransfer.TransactOpts, newOperator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address newOperator) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactorSession) UpdateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.UpdateOperator(&_TokenBatchTransfer.TransactOpts, newOperator)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 value) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactor) WithdrawToken(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.contract.Transact(opts, "withdrawToken", value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 value) returns()
func (_TokenBatchTransfer *TokenBatchTransferSession) WithdrawToken(value *big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.WithdrawToken(&_TokenBatchTransfer.TransactOpts, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 value) returns()
func (_TokenBatchTransfer *TokenBatchTransferTransactorSession) WithdrawToken(value *big.Int) (*types.Transaction, error) {
	return _TokenBatchTransfer.Contract.WithdrawToken(&_TokenBatchTransfer.TransactOpts, value)
}

// TokenBatchTransferNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the TokenBatchTransfer contract.
type TokenBatchTransferNewOperatorIterator struct {
	Event *TokenBatchTransferNewOperator // Event containing the contract specifics and raw log

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
func (it *TokenBatchTransferNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBatchTransferNewOperator)
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
		it.Event = new(TokenBatchTransferNewOperator)
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
func (it *TokenBatchTransferNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBatchTransferNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBatchTransferNewOperator represents a NewOperator event raised by the TokenBatchTransfer contract.
type TokenBatchTransferNewOperator struct {
	TransferOperator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address transferOperator)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) FilterNewOperator(opts *bind.FilterOpts) (*TokenBatchTransferNewOperatorIterator, error) {

	logs, sub, err := _TokenBatchTransfer.contract.FilterLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferNewOperatorIterator{contract: _TokenBatchTransfer.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address transferOperator)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *TokenBatchTransferNewOperator) (event.Subscription, error) {

	logs, sub, err := _TokenBatchTransfer.contract.WatchLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBatchTransferNewOperator)
				if err := _TokenBatchTransfer.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address transferOperator)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) ParseNewOperator(log types.Log) (*TokenBatchTransferNewOperator, error) {
	event := new(TokenBatchTransferNewOperator)
	if err := _TokenBatchTransfer.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBatchTransferOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenBatchTransfer contract.
type TokenBatchTransferOwnershipTransferredIterator struct {
	Event *TokenBatchTransferOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenBatchTransferOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBatchTransferOwnershipTransferred)
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
		it.Event = new(TokenBatchTransferOwnershipTransferred)
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
func (it *TokenBatchTransferOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBatchTransferOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBatchTransferOwnershipTransferred represents a OwnershipTransferred event raised by the TokenBatchTransfer contract.
type TokenBatchTransferOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenBatchTransferOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBatchTransfer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferOwnershipTransferredIterator{contract: _TokenBatchTransfer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenBatchTransferOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBatchTransfer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBatchTransferOwnershipTransferred)
				if err := _TokenBatchTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenBatchTransfer *TokenBatchTransferFilterer) ParseOwnershipTransferred(log types.Log) (*TokenBatchTransferOwnershipTransferred, error) {
	event := new(TokenBatchTransferOwnershipTransferred)
	if err := _TokenBatchTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBatchTransferWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the TokenBatchTransfer contract.
type TokenBatchTransferWithdrawTokenIterator struct {
	Event *TokenBatchTransferWithdrawToken // Event containing the contract specifics and raw log

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
func (it *TokenBatchTransferWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBatchTransferWithdrawToken)
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
		it.Event = new(TokenBatchTransferWithdrawToken)
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
func (it *TokenBatchTransferWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBatchTransferWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBatchTransferWithdrawToken represents a WithdrawToken event raised by the TokenBatchTransfer contract.
type TokenBatchTransferWithdrawToken struct {
	Owner       common.Address
	StakeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed owner, uint256 stakeAmount)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) FilterWithdrawToken(opts *bind.FilterOpts, owner []common.Address) (*TokenBatchTransferWithdrawTokenIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenBatchTransfer.contract.FilterLogs(opts, "WithdrawToken", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBatchTransferWithdrawTokenIterator{contract: _TokenBatchTransfer.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed owner, uint256 stakeAmount)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *TokenBatchTransferWithdrawToken, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenBatchTransfer.contract.WatchLogs(opts, "WithdrawToken", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBatchTransferWithdrawToken)
				if err := _TokenBatchTransfer.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed owner, uint256 stakeAmount)
func (_TokenBatchTransfer *TokenBatchTransferFilterer) ParseWithdrawToken(log types.Log) (*TokenBatchTransferWithdrawToken, error) {
	event := new(TokenBatchTransferWithdrawToken)
	if err := _TokenBatchTransfer.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
