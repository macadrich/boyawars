// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package biddingwars

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

// BiddingwarsMetaData contains all meta data concerning the Biddingwars contract.
var BiddingwarsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"GameStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEndTime\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WinnerPaid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"commissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"gameActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[],\"name\":\"endGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BiddingwarsABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingwarsMetaData.ABI instead.
var BiddingwarsABI = BiddingwarsMetaData.ABI

// Biddingwars is an auto generated Go binding around an Ethereum contract.
type Biddingwars struct {
	BiddingwarsCaller     // Read-only binding to the contract
	BiddingwarsTransactor // Write-only binding to the contract
	BiddingwarsFilterer   // Log filterer for contract events
}

// BiddingwarsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingwarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingwarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingwarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingwarsSession struct {
	Contract     *Biddingwars      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingwarsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingwarsCallerSession struct {
	Contract *BiddingwarsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BiddingwarsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingwarsTransactorSession struct {
	Contract     *BiddingwarsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BiddingwarsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingwarsRaw struct {
	Contract *Biddingwars // Generic contract binding to access the raw methods on
}

// BiddingwarsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingwarsCallerRaw struct {
	Contract *BiddingwarsCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingwarsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingwarsTransactorRaw struct {
	Contract *BiddingwarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiddingwars creates a new instance of Biddingwars, bound to a specific deployed contract.
func NewBiddingwars(address common.Address, backend bind.ContractBackend) (*Biddingwars, error) {
	contract, err := bindBiddingwars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Biddingwars{BiddingwarsCaller: BiddingwarsCaller{contract: contract}, BiddingwarsTransactor: BiddingwarsTransactor{contract: contract}, BiddingwarsFilterer: BiddingwarsFilterer{contract: contract}}, nil
}

// NewBiddingwarsCaller creates a new read-only instance of Biddingwars, bound to a specific deployed contract.
func NewBiddingwarsCaller(address common.Address, caller bind.ContractCaller) (*BiddingwarsCaller, error) {
	contract, err := bindBiddingwars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarsCaller{contract: contract}, nil
}

// NewBiddingwarsTransactor creates a new write-only instance of Biddingwars, bound to a specific deployed contract.
func NewBiddingwarsTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingwarsTransactor, error) {
	contract, err := bindBiddingwars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarsTransactor{contract: contract}, nil
}

// NewBiddingwarsFilterer creates a new log filterer instance of Biddingwars, bound to a specific deployed contract.
func NewBiddingwarsFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingwarsFilterer, error) {
	contract, err := bindBiddingwars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingwarsFilterer{contract: contract}, nil
}

// bindBiddingwars binds a generic wrapper to an already deployed contract.
func bindBiddingwars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BiddingwarsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwars *BiddingwarsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwars.Contract.BiddingwarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwars *BiddingwarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.Contract.BiddingwarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwars *BiddingwarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwars.Contract.BiddingwarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwars *BiddingwarsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwars *BiddingwarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwars *BiddingwarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwars.Contract.contract.Transact(opts, method, params...)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Biddingwars *BiddingwarsCaller) CommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "commissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Biddingwars *BiddingwarsSession) CommissionRate() (*big.Int, error) {
	return _Biddingwars.Contract.CommissionRate(&_Biddingwars.CallOpts)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Biddingwars *BiddingwarsCallerSession) CommissionRate() (*big.Int, error) {
	return _Biddingwars.Contract.CommissionRate(&_Biddingwars.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Biddingwars *BiddingwarsCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Biddingwars *BiddingwarsSession) EndTime() (*big.Int, error) {
	return _Biddingwars.Contract.EndTime(&_Biddingwars.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Biddingwars *BiddingwarsCallerSession) EndTime() (*big.Int, error) {
	return _Biddingwars.Contract.EndTime(&_Biddingwars.CallOpts)
}

// GameActive is a free data retrieval call binding the contract method 0xf020044f.
//
// Solidity: function gameActive() view returns(bool)
func (_Biddingwars *BiddingwarsCaller) GameActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "gameActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameActive is a free data retrieval call binding the contract method 0xf020044f.
//
// Solidity: function gameActive() view returns(bool)
func (_Biddingwars *BiddingwarsSession) GameActive() (bool, error) {
	return _Biddingwars.Contract.GameActive(&_Biddingwars.CallOpts)
}

// GameActive is a free data retrieval call binding the contract method 0xf020044f.
//
// Solidity: function gameActive() view returns(bool)
func (_Biddingwars *BiddingwarsCallerSession) GameActive() (bool, error) {
	return _Biddingwars.Contract.GameActive(&_Biddingwars.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwars *BiddingwarsCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwars *BiddingwarsSession) HighestBid() (*big.Int, error) {
	return _Biddingwars.Contract.HighestBid(&_Biddingwars.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwars *BiddingwarsCallerSession) HighestBid() (*big.Int, error) {
	return _Biddingwars.Contract.HighestBid(&_Biddingwars.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwars *BiddingwarsCaller) LastBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "lastBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwars *BiddingwarsSession) LastBidder() (common.Address, error) {
	return _Biddingwars.Contract.LastBidder(&_Biddingwars.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwars *BiddingwarsCallerSession) LastBidder() (common.Address, error) {
	return _Biddingwars.Contract.LastBidder(&_Biddingwars.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwars *BiddingwarsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwars.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwars *BiddingwarsSession) Owner() (common.Address, error) {
	return _Biddingwars.Contract.Owner(&_Biddingwars.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwars *BiddingwarsCallerSession) Owner() (common.Address, error) {
	return _Biddingwars.Contract.Owner(&_Biddingwars.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwars *BiddingwarsTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwars *BiddingwarsSession) Bid() (*types.Transaction, error) {
	return _Biddingwars.Contract.Bid(&_Biddingwars.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwars *BiddingwarsTransactorSession) Bid() (*types.Transaction, error) {
	return _Biddingwars.Contract.Bid(&_Biddingwars.TransactOpts)
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Biddingwars *BiddingwarsTransactor) EndGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.contract.Transact(opts, "endGame")
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Biddingwars *BiddingwarsSession) EndGame() (*types.Transaction, error) {
	return _Biddingwars.Contract.EndGame(&_Biddingwars.TransactOpts)
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Biddingwars *BiddingwarsTransactorSession) EndGame() (*types.Transaction, error) {
	return _Biddingwars.Contract.EndGame(&_Biddingwars.TransactOpts)
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Biddingwars *BiddingwarsTransactor) StartGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.contract.Transact(opts, "startGame")
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Biddingwars *BiddingwarsSession) StartGame() (*types.Transaction, error) {
	return _Biddingwars.Contract.StartGame(&_Biddingwars.TransactOpts)
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Biddingwars *BiddingwarsTransactorSession) StartGame() (*types.Transaction, error) {
	return _Biddingwars.Contract.StartGame(&_Biddingwars.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Biddingwars *BiddingwarsTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwars.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Biddingwars *BiddingwarsSession) WithdrawFunds() (*types.Transaction, error) {
	return _Biddingwars.Contract.WithdrawFunds(&_Biddingwars.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Biddingwars *BiddingwarsTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _Biddingwars.Contract.WithdrawFunds(&_Biddingwars.TransactOpts)
}

// BiddingwarsGameStartedIterator is returned from FilterGameStarted and is used to iterate over the raw logs and unpacked data for GameStarted events raised by the Biddingwars contract.
type BiddingwarsGameStartedIterator struct {
	Event *BiddingwarsGameStarted // Event containing the contract specifics and raw log

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
func (it *BiddingwarsGameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarsGameStarted)
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
		it.Event = new(BiddingwarsGameStarted)
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
func (it *BiddingwarsGameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarsGameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarsGameStarted represents a GameStarted event raised by the Biddingwars contract.
type BiddingwarsGameStarted struct {
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameStarted is a free log retrieval operation binding the contract event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 endTime)
func (_Biddingwars *BiddingwarsFilterer) FilterGameStarted(opts *bind.FilterOpts) (*BiddingwarsGameStartedIterator, error) {

	logs, sub, err := _Biddingwars.contract.FilterLogs(opts, "GameStarted")
	if err != nil {
		return nil, err
	}
	return &BiddingwarsGameStartedIterator{contract: _Biddingwars.contract, event: "GameStarted", logs: logs, sub: sub}, nil
}

// WatchGameStarted is a free log subscription operation binding the contract event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 endTime)
func (_Biddingwars *BiddingwarsFilterer) WatchGameStarted(opts *bind.WatchOpts, sink chan<- *BiddingwarsGameStarted) (event.Subscription, error) {

	logs, sub, err := _Biddingwars.contract.WatchLogs(opts, "GameStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarsGameStarted)
				if err := _Biddingwars.contract.UnpackLog(event, "GameStarted", log); err != nil {
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

// ParseGameStarted is a log parse operation binding the contract event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 endTime)
func (_Biddingwars *BiddingwarsFilterer) ParseGameStarted(log types.Log) (*BiddingwarsGameStarted, error) {
	event := new(BiddingwarsGameStarted)
	if err := _Biddingwars.contract.UnpackLog(event, "GameStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarsNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Biddingwars contract.
type BiddingwarsNewBidIterator struct {
	Event *BiddingwarsNewBid // Event containing the contract specifics and raw log

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
func (it *BiddingwarsNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarsNewBid)
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
		it.Event = new(BiddingwarsNewBid)
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
func (it *BiddingwarsNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarsNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarsNewBid represents a NewBid event raised by the Biddingwars contract.
type BiddingwarsNewBid struct {
	Bidder     common.Address
	Amount     *big.Int
	NewEndTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 newEndTime)
func (_Biddingwars *BiddingwarsFilterer) FilterNewBid(opts *bind.FilterOpts) (*BiddingwarsNewBidIterator, error) {

	logs, sub, err := _Biddingwars.contract.FilterLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return &BiddingwarsNewBidIterator{contract: _Biddingwars.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 newEndTime)
func (_Biddingwars *BiddingwarsFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *BiddingwarsNewBid) (event.Subscription, error) {

	logs, sub, err := _Biddingwars.contract.WatchLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarsNewBid)
				if err := _Biddingwars.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 newEndTime)
func (_Biddingwars *BiddingwarsFilterer) ParseNewBid(log types.Log) (*BiddingwarsNewBid, error) {
	event := new(BiddingwarsNewBid)
	if err := _Biddingwars.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarsWinnerPaidIterator is returned from FilterWinnerPaid and is used to iterate over the raw logs and unpacked data for WinnerPaid events raised by the Biddingwars contract.
type BiddingwarsWinnerPaidIterator struct {
	Event *BiddingwarsWinnerPaid // Event containing the contract specifics and raw log

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
func (it *BiddingwarsWinnerPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarsWinnerPaid)
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
		it.Event = new(BiddingwarsWinnerPaid)
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
func (it *BiddingwarsWinnerPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarsWinnerPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarsWinnerPaid represents a WinnerPaid event raised by the Biddingwars contract.
type BiddingwarsWinnerPaid struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinnerPaid is a free log retrieval operation binding the contract event 0x8cbbe5cd65720098fc8ce6e99a5deb232085117dd486475b49cb11604b528f30.
//
// Solidity: event WinnerPaid(address winner, uint256 amount)
func (_Biddingwars *BiddingwarsFilterer) FilterWinnerPaid(opts *bind.FilterOpts) (*BiddingwarsWinnerPaidIterator, error) {

	logs, sub, err := _Biddingwars.contract.FilterLogs(opts, "WinnerPaid")
	if err != nil {
		return nil, err
	}
	return &BiddingwarsWinnerPaidIterator{contract: _Biddingwars.contract, event: "WinnerPaid", logs: logs, sub: sub}, nil
}

// WatchWinnerPaid is a free log subscription operation binding the contract event 0x8cbbe5cd65720098fc8ce6e99a5deb232085117dd486475b49cb11604b528f30.
//
// Solidity: event WinnerPaid(address winner, uint256 amount)
func (_Biddingwars *BiddingwarsFilterer) WatchWinnerPaid(opts *bind.WatchOpts, sink chan<- *BiddingwarsWinnerPaid) (event.Subscription, error) {

	logs, sub, err := _Biddingwars.contract.WatchLogs(opts, "WinnerPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarsWinnerPaid)
				if err := _Biddingwars.contract.UnpackLog(event, "WinnerPaid", log); err != nil {
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

// ParseWinnerPaid is a log parse operation binding the contract event 0x8cbbe5cd65720098fc8ce6e99a5deb232085117dd486475b49cb11604b528f30.
//
// Solidity: event WinnerPaid(address winner, uint256 amount)
func (_Biddingwars *BiddingwarsFilterer) ParseWinnerPaid(log types.Log) (*BiddingwarsWinnerPaid, error) {
	event := new(BiddingwarsWinnerPaid)
	if err := _Biddingwars.contract.UnpackLog(event, "WinnerPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}