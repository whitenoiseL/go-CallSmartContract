// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Try

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TryABI is the input ABI used to generate the binding from.
const TryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"setId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"ItemSet\",\"type\":\"event\"}]"

// TryBin is the compiled bytecode used for deploying new contracts.
const TryBin = `608060405234801561001057600080fd5b506101fe806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680635d1ca6311461005c578063af640d0f1461009b578063d1371961146100da575b600080fd5b34801561006857600080fd5b50610071610111565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156100a757600080fd5b506100b061012e565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156100e657600080fd5b5061010f600480360381019080803567ffffffffffffffff169060200190929190505050610147565b005b60008060009054906101000a900467ffffffffffffffff16905090565b6000809054906101000a900467ffffffffffffffff1681565b806000806101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fe7fd42cfef56f1caa721e298bd9ae0bc5ee36e3efb5f960a9f4ad0770a6ffc1e6000809054906101000a900467ffffffffffffffff16604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390a1505600a165627a7a7230582088de163ef7af42ff08fafb2ed5c0729070d495b92c35f6a5d9e80304b473383c0029`

// DeployTry deploys a new Ethereum contract, binding an instance of Try to it.
func DeployTry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Try, error) {
	parsed, err := abi.JSON(strings.NewReader(TryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Try{TryCaller: TryCaller{contract: contract}, TryTransactor: TryTransactor{contract: contract}, TryFilterer: TryFilterer{contract: contract}}, nil
}

// Try is an auto generated Go binding around an Ethereum contract.
type Try struct {
	TryCaller     // Read-only binding to the contract
	TryTransactor // Write-only binding to the contract
	TryFilterer   // Log filterer for contract events
}

// TryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrySession struct {
	Contract     *Try              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TryCallerSession struct {
	Contract *TryCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TryTransactorSession struct {
	Contract     *TryTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TryRaw struct {
	Contract *Try // Generic contract binding to access the raw methods on
}

// TryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TryCallerRaw struct {
	Contract *TryCaller // Generic read-only contract binding to access the raw methods on
}

// TryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TryTransactorRaw struct {
	Contract *TryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTry creates a new instance of Try, bound to a specific deployed contract.
func NewTry(address common.Address, backend bind.ContractBackend) (*Try, error) {
	contract, err := bindTry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Try{TryCaller: TryCaller{contract: contract}, TryTransactor: TryTransactor{contract: contract}, TryFilterer: TryFilterer{contract: contract}}, nil
}

// NewTryCaller creates a new read-only instance of Try, bound to a specific deployed contract.
func NewTryCaller(address common.Address, caller bind.ContractCaller) (*TryCaller, error) {
	contract, err := bindTry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TryCaller{contract: contract}, nil
}

// NewTryTransactor creates a new write-only instance of Try, bound to a specific deployed contract.
func NewTryTransactor(address common.Address, transactor bind.ContractTransactor) (*TryTransactor, error) {
	contract, err := bindTry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TryTransactor{contract: contract}, nil
}

// NewTryFilterer creates a new log filterer instance of Try, bound to a specific deployed contract.
func NewTryFilterer(address common.Address, filterer bind.ContractFilterer) (*TryFilterer, error) {
	contract, err := bindTry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TryFilterer{contract: contract}, nil
}

// bindTry binds a generic wrapper to an already deployed contract.
func bindTry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Try *TryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Try.Contract.TryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Try *TryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Try.Contract.TryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Try *TryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Try.Contract.TryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Try *TryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Try.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Try *TryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Try.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Try *TryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Try.Contract.contract.Transact(opts, method, params...)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(uint64)
func (_Try *TryCaller) GetId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Try.contract.Call(opts, out, "getId")
	return *ret0, err
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(uint64)
func (_Try *TrySession) GetId() (uint64, error) {
	return _Try.Contract.GetId(&_Try.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(uint64)
func (_Try *TryCallerSession) GetId() (uint64, error) {
	return _Try.Contract.GetId(&_Try.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint64)
func (_Try *TryCaller) Id(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Try.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint64)
func (_Try *TrySession) Id() (uint64, error) {
	return _Try.Contract.Id(&_Try.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint64)
func (_Try *TryCallerSession) Id() (uint64, error) {
	return _Try.Contract.Id(&_Try.CallOpts)
}

// SetId is a paid mutator transaction binding the contract method 0xd1371961.
//
// Solidity: function setId(_id uint64) returns()
func (_Try *TryTransactor) SetId(opts *bind.TransactOpts, _id uint64) (*types.Transaction, error) {
	return _Try.contract.Transact(opts, "setId", _id)
}

// SetId is a paid mutator transaction binding the contract method 0xd1371961.
//
// Solidity: function setId(_id uint64) returns()
func (_Try *TrySession) SetId(_id uint64) (*types.Transaction, error) {
	return _Try.Contract.SetId(&_Try.TransactOpts, _id)
}

// SetId is a paid mutator transaction binding the contract method 0xd1371961.
//
// Solidity: function setId(_id uint64) returns()
func (_Try *TryTransactorSession) SetId(_id uint64) (*types.Transaction, error) {
	return _Try.Contract.SetId(&_Try.TransactOpts, _id)
}

// TryItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Try contract.
type TryItemSetIterator struct {
	Event *TryItemSet // Event containing the contract specifics and raw log

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
func (it *TryItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TryItemSet)
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
		it.Event = new(TryItemSet)
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
func (it *TryItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TryItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TryItemSet represents a ItemSet event raised by the Try contract.
type TryItemSet struct {
	Id  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe7fd42cfef56f1caa721e298bd9ae0bc5ee36e3efb5f960a9f4ad0770a6ffc1e.
//
// Solidity: e ItemSet(id uint64)
func (_Try *TryFilterer) FilterItemSet(opts *bind.FilterOpts) (*TryItemSetIterator, error) {

	logs, sub, err := _Try.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &TryItemSetIterator{contract: _Try.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe7fd42cfef56f1caa721e298bd9ae0bc5ee36e3efb5f960a9f4ad0770a6ffc1e.
//
// Solidity: e ItemSet(id uint64)
func (_Try *TryFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *TryItemSet) (event.Subscription, error) {

	logs, sub, err := _Try.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TryItemSet)
				if err := _Try.contract.UnpackLog(event, "ItemSet", log); err != nil {
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
