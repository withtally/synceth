// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ExampleMetaData contains all meta data concerning the Example contract.
var ExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exampleValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600660808190526532ba3433b2b760d11b60a0908152610027916000919061003a565b5034801561003457600080fd5b5061010e565b828054610046906100d3565b90600052602060002090601f01602090048101928261006857600085556100ae565b82601f1061008157805160ff19168380011785556100ae565b828001600101855582156100ae579182015b828111156100ae578251825591602001919060010190610093565b506100ba9291506100be565b5090565b5b808211156100ba57600081556001016100bf565b600181811c908216806100e757607f821691505b6020821081141561010857634e487b7160e01b600052602260045260246000fd5b50919050565b6101a68061011d6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d8f7c4cf14610030575b600080fd5b61003861004e565b60405161004591906100e0565b60405180910390f35b60606000805461005d90610135565b80601f016020809104026020016040519081016040528092919081815260200182805461008990610135565b80156100d65780601f106100ab576101008083540402835291602001916100d6565b820191906000526020600020905b8154815290600101906020018083116100b957829003601f168201915b5050505050905090565b600060208083528351808285015260005b8181101561010d578581018301518582016040015282016100f1565b8181111561011f576000604083870101525b50601f01601f1916929092016040019392505050565b600181811c9082168061014957607f821691505b6020821081141561016a57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220d432acfc601a55b38872c4482f4ff23a49d4b5ed1caa9f6ba25dbf74dff0072364736f6c634300080a0033",
}

// ExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleMetaData.ABI instead.
var ExampleABI = ExampleMetaData.ABI

// ExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleMetaData.Bin instead.
var ExampleBin = ExampleMetaData.Bin

// DeployExample deploys a new Ethereum contract, binding an instance of Example to it.
func DeployExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Example, error) {
	parsed, err := ExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Example{ExampleCaller: ExampleCaller{contract: contract}, ExampleTransactor: ExampleTransactor{contract: contract}, ExampleFilterer: ExampleFilterer{contract: contract}}, nil
}

// Example is an auto generated Go binding around an Ethereum contract.
type Example struct {
	ExampleCaller     // Read-only binding to the contract
	ExampleTransactor // Write-only binding to the contract
	ExampleFilterer   // Log filterer for contract events
}

// ExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleSession struct {
	Contract     *Example          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleCallerSession struct {
	Contract *ExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleTransactorSession struct {
	Contract     *ExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRaw struct {
	Contract *Example // Generic contract binding to access the raw methods on
}

// ExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleCallerRaw struct {
	Contract *ExampleCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleTransactorRaw struct {
	Contract *ExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExample creates a new instance of Example, bound to a specific deployed contract.
func NewExample(address common.Address, backend bind.ContractBackend) (*Example, error) {
	contract, err := bindExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Example{ExampleCaller: ExampleCaller{contract: contract}, ExampleTransactor: ExampleTransactor{contract: contract}, ExampleFilterer: ExampleFilterer{contract: contract}}, nil
}

// NewExampleCaller creates a new read-only instance of Example, bound to a specific deployed contract.
func NewExampleCaller(address common.Address, caller bind.ContractCaller) (*ExampleCaller, error) {
	contract, err := bindExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCaller{contract: contract}, nil
}

// NewExampleTransactor creates a new write-only instance of Example, bound to a specific deployed contract.
func NewExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleTransactor, error) {
	contract, err := bindExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleTransactor{contract: contract}, nil
}

// NewExampleFilterer creates a new log filterer instance of Example, bound to a specific deployed contract.
func NewExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleFilterer, error) {
	contract, err := bindExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleFilterer{contract: contract}, nil
}

// bindExample binds a generic wrapper to an already deployed contract.
func bindExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.ExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.contract.Transact(opts, method, params...)
}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_Example *ExampleCaller) ExampleValue(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Example.contract.Call(opts, &out, "exampleValue")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_Example *ExampleSession) ExampleValue() (string, error) {
	return _Example.Contract.ExampleValue(&_Example.CallOpts)
}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_Example *ExampleCallerSession) ExampleValue() (string, error) {
	return _Example.Contract.ExampleValue(&_Example.CallOpts)
}

// ExampleExampleEventIterator is returned from FilterExampleEvent and is used to iterate over the raw logs and unpacked data for ExampleEvent events raised by the Example contract.
type ExampleExampleEventIterator struct {
	Event *ExampleExampleEvent // Event containing the contract specifics and raw log

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
func (it *ExampleExampleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleExampleEvent)
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
		it.Event = new(ExampleExampleEvent)
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
func (it *ExampleExampleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleExampleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleExampleEvent represents a ExampleEvent event raised by the Example contract.
type ExampleExampleEvent struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExampleEvent is a free log retrieval operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_Example *ExampleFilterer) FilterExampleEvent(opts *bind.FilterOpts) (*ExampleExampleEventIterator, error) {

	logs, sub, err := _Example.contract.FilterLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return &ExampleExampleEventIterator{contract: _Example.contract, event: "ExampleEvent", logs: logs, sub: sub}, nil
}

// WatchExampleEvent is a free log subscription operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_Example *ExampleFilterer) WatchExampleEvent(opts *bind.WatchOpts, sink chan<- *ExampleExampleEvent) (event.Subscription, error) {

	logs, sub, err := _Example.contract.WatchLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleExampleEvent)
				if err := _Example.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
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

// ParseExampleEvent is a log parse operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_Example *ExampleFilterer) ParseExampleEvent(log types.Log) (*ExampleExampleEvent, error) {
	event := new(ExampleExampleEvent)
	if err := _Example.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FakeExampleMetaData contains all meta data concerning the FakeExample contract.
var FakeExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exampleValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"fakeEmitExampleEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"retstringExampleValue0\",\"type\":\"string\"}],\"name\":\"fakeSetExampleValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506103e68061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063010ac92514610043578063b8290c1a14610058578063d8f7c4cf1461006b575b5f80fd5b610056610051366004610175565b610089565b005b610056610066366004610175565b6100c3565b6100736100d2565b6040516100809190610220565b60405180910390f35b7fb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70816040516100b89190610220565b60405180910390a150565b5f6100ce82826102f0565b5050565b60605f80546100e09061026c565b80601f016020809104026020016040519081016040528092919081815260200182805461010c9061026c565b80156101575780601f1061012e57610100808354040283529160200191610157565b820191905f5260205f20905b81548152906001019060200180831161013a57829003601f168201915b5050505050905090565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610185575f80fd5b813567ffffffffffffffff8082111561019c575f80fd5b818401915084601f8301126101af575f80fd5b8135818111156101c1576101c1610161565b604051601f8201601f19908116603f011681019083821181831017156101e9576101e9610161565b81604052828152876020848701011115610201575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f602080835283518060208501525f5b8181101561024c57858101830151858201604001528201610230565b505f604082860101526040601f19601f8301168501019250505092915050565b600181811c9082168061028057607f821691505b60208210810361029e57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102eb57805f5260205f20601f840160051c810160208510156102c95750805b601f840160051c820191505b818110156102e8575f81556001016102d5565b50505b505050565b815167ffffffffffffffff81111561030a5761030a610161565b61031e81610318845461026c565b846102a4565b602080601f831160018114610351575f841561033a5750858301515b5f19600386901b1c1916600185901b1785556103a8565b5f85815260208120601f198616915b8281101561037f57888601518255948401946001909101908401610360565b508582101561039c57878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212204a60b5797056e6f2d3fe036fc0602e31f043cefe6854e74cffc0093546f1924164736f6c63430008170033",
}

// FakeExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use FakeExampleMetaData.ABI instead.
var FakeExampleABI = FakeExampleMetaData.ABI

// FakeExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FakeExampleMetaData.Bin instead.
var FakeExampleBin = FakeExampleMetaData.Bin

// DeployFakeExample deploys a new Ethereum contract, binding an instance of FakeExample to it.
func DeployFakeExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FakeExample, error) {
	parsed, err := FakeExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FakeExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FakeExample{FakeExampleCaller: FakeExampleCaller{contract: contract}, FakeExampleTransactor: FakeExampleTransactor{contract: contract}, FakeExampleFilterer: FakeExampleFilterer{contract: contract}}, nil
}

// FakeExample is an auto generated Go binding around an Ethereum contract.
type FakeExample struct {
	FakeExampleCaller     // Read-only binding to the contract
	FakeExampleTransactor // Write-only binding to the contract
	FakeExampleFilterer   // Log filterer for contract events
}

// FakeExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FakeExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FakeExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FakeExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FakeExampleSession struct {
	Contract     *FakeExample      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FakeExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FakeExampleCallerSession struct {
	Contract *FakeExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FakeExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FakeExampleTransactorSession struct {
	Contract     *FakeExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FakeExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type FakeExampleRaw struct {
	Contract *FakeExample // Generic contract binding to access the raw methods on
}

// FakeExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FakeExampleCallerRaw struct {
	Contract *FakeExampleCaller // Generic read-only contract binding to access the raw methods on
}

// FakeExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FakeExampleTransactorRaw struct {
	Contract *FakeExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFakeExample creates a new instance of FakeExample, bound to a specific deployed contract.
func NewFakeExample(address common.Address, backend bind.ContractBackend) (*FakeExample, error) {
	contract, err := bindFakeExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FakeExample{FakeExampleCaller: FakeExampleCaller{contract: contract}, FakeExampleTransactor: FakeExampleTransactor{contract: contract}, FakeExampleFilterer: FakeExampleFilterer{contract: contract}}, nil
}

// NewFakeExampleCaller creates a new read-only instance of FakeExample, bound to a specific deployed contract.
func NewFakeExampleCaller(address common.Address, caller bind.ContractCaller) (*FakeExampleCaller, error) {
	contract, err := bindFakeExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FakeExampleCaller{contract: contract}, nil
}

// NewFakeExampleTransactor creates a new write-only instance of FakeExample, bound to a specific deployed contract.
func NewFakeExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*FakeExampleTransactor, error) {
	contract, err := bindFakeExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FakeExampleTransactor{contract: contract}, nil
}

// NewFakeExampleFilterer creates a new log filterer instance of FakeExample, bound to a specific deployed contract.
func NewFakeExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*FakeExampleFilterer, error) {
	contract, err := bindFakeExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FakeExampleFilterer{contract: contract}, nil
}

// bindFakeExample binds a generic wrapper to an already deployed contract.
func bindFakeExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FakeExampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FakeExample *FakeExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FakeExample.Contract.FakeExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FakeExample *FakeExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FakeExample *FakeExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FakeExample *FakeExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FakeExample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FakeExample *FakeExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FakeExample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FakeExample *FakeExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FakeExample.Contract.contract.Transact(opts, method, params...)
}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_FakeExample *FakeExampleCaller) ExampleValue(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FakeExample.contract.Call(opts, &out, "exampleValue")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_FakeExample *FakeExampleSession) ExampleValue() (string, error) {
	return _FakeExample.Contract.ExampleValue(&_FakeExample.CallOpts)
}

// ExampleValue is a free data retrieval call binding the contract method 0xd8f7c4cf.
//
// Solidity: function exampleValue() view returns(string)
func (_FakeExample *FakeExampleCallerSession) ExampleValue() (string, error) {
	return _FakeExample.Contract.ExampleValue(&_FakeExample.CallOpts)
}

// FakeEmitExampleEvent is a paid mutator transaction binding the contract method 0x010ac925.
//
// Solidity: function fakeEmitExampleEvent(string value) returns()
func (_FakeExample *FakeExampleTransactor) FakeEmitExampleEvent(opts *bind.TransactOpts, value string) (*types.Transaction, error) {
	return _FakeExample.contract.Transact(opts, "fakeEmitExampleEvent", value)
}

// FakeEmitExampleEvent is a paid mutator transaction binding the contract method 0x010ac925.
//
// Solidity: function fakeEmitExampleEvent(string value) returns()
func (_FakeExample *FakeExampleSession) FakeEmitExampleEvent(value string) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeEmitExampleEvent(&_FakeExample.TransactOpts, value)
}

// FakeEmitExampleEvent is a paid mutator transaction binding the contract method 0x010ac925.
//
// Solidity: function fakeEmitExampleEvent(string value) returns()
func (_FakeExample *FakeExampleTransactorSession) FakeEmitExampleEvent(value string) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeEmitExampleEvent(&_FakeExample.TransactOpts, value)
}

// FakeSetExampleValue is a paid mutator transaction binding the contract method 0xb8290c1a.
//
// Solidity: function fakeSetExampleValue(string retstringExampleValue0) returns()
func (_FakeExample *FakeExampleTransactor) FakeSetExampleValue(opts *bind.TransactOpts, retstringExampleValue0 string) (*types.Transaction, error) {
	return _FakeExample.contract.Transact(opts, "fakeSetExampleValue", retstringExampleValue0)
}

// FakeSetExampleValue is a paid mutator transaction binding the contract method 0xb8290c1a.
//
// Solidity: function fakeSetExampleValue(string retstringExampleValue0) returns()
func (_FakeExample *FakeExampleSession) FakeSetExampleValue(retstringExampleValue0 string) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeSetExampleValue(&_FakeExample.TransactOpts, retstringExampleValue0)
}

// FakeSetExampleValue is a paid mutator transaction binding the contract method 0xb8290c1a.
//
// Solidity: function fakeSetExampleValue(string retstringExampleValue0) returns()
func (_FakeExample *FakeExampleTransactorSession) FakeSetExampleValue(retstringExampleValue0 string) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeSetExampleValue(&_FakeExample.TransactOpts, retstringExampleValue0)
}

// FakeExampleExampleEventIterator is returned from FilterExampleEvent and is used to iterate over the raw logs and unpacked data for ExampleEvent events raised by the FakeExample contract.
type FakeExampleExampleEventIterator struct {
	Event *FakeExampleExampleEvent // Event containing the contract specifics and raw log

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
func (it *FakeExampleExampleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FakeExampleExampleEvent)
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
		it.Event = new(FakeExampleExampleEvent)
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
func (it *FakeExampleExampleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FakeExampleExampleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FakeExampleExampleEvent represents a ExampleEvent event raised by the FakeExample contract.
type FakeExampleExampleEvent struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExampleEvent is a free log retrieval operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_FakeExample *FakeExampleFilterer) FilterExampleEvent(opts *bind.FilterOpts) (*FakeExampleExampleEventIterator, error) {

	logs, sub, err := _FakeExample.contract.FilterLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return &FakeExampleExampleEventIterator{contract: _FakeExample.contract, event: "ExampleEvent", logs: logs, sub: sub}, nil
}

// WatchExampleEvent is a free log subscription operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_FakeExample *FakeExampleFilterer) WatchExampleEvent(opts *bind.WatchOpts, sink chan<- *FakeExampleExampleEvent) (event.Subscription, error) {

	logs, sub, err := _FakeExample.contract.WatchLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FakeExampleExampleEvent)
				if err := _FakeExample.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
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

// ParseExampleEvent is a log parse operation binding the contract event 0xb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70.
//
// Solidity: event ExampleEvent(string value)
func (_FakeExample *FakeExampleFilterer) ParseExampleEvent(log types.Log) (*FakeExampleExampleEvent, error) {
	event := new(FakeExampleExampleEvent)
	if err := _FakeExample.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
