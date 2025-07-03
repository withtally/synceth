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

// FakeExampleWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type FakeExampleWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// ExampleMetaData contains all meta data concerning the Example contract.
var ExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exampleValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"WithdrawalQueued\",\"type\":\"event\"}]",
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

// ExampleWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the Example contract.
type ExampleWithdrawalQueuedIterator struct {
	Event *ExampleWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *ExampleWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleWithdrawalQueued)
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
		it.Event = new(ExampleWithdrawalQueued)
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
func (it *ExampleWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleWithdrawalQueued represents a WithdrawalQueued event raised by the Example contract.
type ExampleWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationManagerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_Example *ExampleFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*ExampleWithdrawalQueuedIterator, error) {

	logs, sub, err := _Example.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ExampleWithdrawalQueuedIterator{contract: _Example.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_Example *ExampleFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ExampleWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _Example.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleWithdrawalQueued)
				if err := _Example.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_Example *ExampleFilterer) ParseWithdrawalQueued(log types.Log) (*ExampleWithdrawalQueued, error) {
	event := new(ExampleWithdrawalQueued)
	if err := _Example.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FakeExampleMetaData contains all meta data concerning the FakeExample contract.
var FakeExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structFakeExample.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"WithdrawalQueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exampleValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"fakeEmitExampleEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structFakeExample.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"fakeEmitWithdrawalQueued\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"retstringExampleValue0\",\"type\":\"string\"}],\"name\":\"fakeSetExampleValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610720806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063010ac9251461005157806399c9c6dc14610066578063b8290c1a14610079578063d8f7c4cf1461008c575b600080fd5b61006461005f366004610448565b6100aa565b005b610064610074366004610360565b6100e4565b610064610087366004610448565b610121565b610094610138565b6040516100a191906105f8565b60405180910390f35b7fb74a38eb2ebca56512a2bb0283f335555a4a4dac46ab998d65fd76f9027dca70816040516100d991906105f8565b60405180910390a150565b7f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9828260405161011592919061055e565b60405180910390a15050565b80516101349060009060208401906101ca565b5050565b60606000805461014790610699565b80601f016020809104026020016040519081016040528092919081815260200182805461017390610699565b80156101c05780601f10610195576101008083540402835291602001916101c0565b820191906000526020600020905b8154815290600101906020018083116101a357829003601f168201915b5050505050905090565b8280546101d690610699565b90600052602060002090601f0160209004810192826101f8576000855561023e565b82601f1061021157805160ff191683800117855561023e565b8280016001018555821561023e579182015b8281111561023e578251825591602001919060010190610223565b5061024a92915061024e565b5090565b5b8082111561024a576000815560010161024f565b80356001600160a01b038116811461027a57600080fd5b919050565b600082601f83011261028f578081fd5b813560206102a461029f83610675565b61064b565b82815281810190858301838502870184018810156102c0578586fd5b855b858110156102e5576102d382610263565b845292840192908401906001016102c2565b5090979650505050505050565b600082601f830112610302578081fd5b8135602061031261029f83610675565b828152818101908583018385028701840188101561032e578586fd5b855b858110156102e557813584529284019290840190600101610330565b803563ffffffff8116811461027a57600080fd5b60008060408385031215610372578182fd5b82359150602083013567ffffffffffffffff80821115610390578283fd5b9084019060e082870312156103a3578283fd5b6103ad60e061064b565b6103b683610263565b81526103c460208401610263565b60208201526103d560408401610263565b6040820152606083013560608201526103f06080840161034c565b608082015260a083013582811115610406578485fd5b6104128882860161027f565b60a08301525060c083013582811115610429578485fd5b610435888286016102f2565b60c0830152508093505050509250929050565b6000602080838503121561045a578182fd5b823567ffffffffffffffff80821115610471578384fd5b818501915085601f830112610484578384fd5b813581811115610496576104966106d4565b6104a8601f8201601f1916850161064b565b915080825286848285010111156104bd578485fd5b80848401858401378101909201929092529392505050565b6001600160a01b03169052565b6000815180845260208085019450808401835b8381101561051a5781516001600160a01b0316875295820195908201906001016104f5565b509495945050505050565b6000815180845260208085019450808401835b8381101561051a57815187529582019590820190600101610538565b63ffffffff169052565b8281526040602080830182905283516001600160a01b03908116838501529084015116606083015282015160009061059960808401826104d5565b50606083015160a083015260808301516105b660c0840182610554565b5060a083015160e0808401526105d06101208401826104e2565b905060c0840151603f19848303016101008501526105ee8282610525565b9695505050505050565b6000602080835283518082850152825b8181101561062457858101830151858201604001528201610608565b818111156106355783604083870101525b50601f01601f1916929092016040019392505050565b60405181810167ffffffffffffffff8111828210171561066d5761066d6106d4565b604052919050565b600067ffffffffffffffff82111561068f5761068f6106d4565b5060209081020190565b6002810460018216806106ad57607f821691505b602082108114156106ce57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220fcdb9333f0670b79e7099984f4e1eef1d3573869f20b9812f6608f3c9d3c9cff64736f6c63430008000033",
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

// FakeEmitWithdrawalQueued is a paid mutator transaction binding the contract method 0x99c9c6dc.
//
// Solidity: function fakeEmitWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal) returns()
func (_FakeExample *FakeExampleTransactor) FakeEmitWithdrawalQueued(opts *bind.TransactOpts, withdrawalRoot [32]byte, withdrawal FakeExampleWithdrawal) (*types.Transaction, error) {
	return _FakeExample.contract.Transact(opts, "fakeEmitWithdrawalQueued", withdrawalRoot, withdrawal)
}

// FakeEmitWithdrawalQueued is a paid mutator transaction binding the contract method 0x99c9c6dc.
//
// Solidity: function fakeEmitWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal) returns()
func (_FakeExample *FakeExampleSession) FakeEmitWithdrawalQueued(withdrawalRoot [32]byte, withdrawal FakeExampleWithdrawal) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeEmitWithdrawalQueued(&_FakeExample.TransactOpts, withdrawalRoot, withdrawal)
}

// FakeEmitWithdrawalQueued is a paid mutator transaction binding the contract method 0x99c9c6dc.
//
// Solidity: function fakeEmitWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal) returns()
func (_FakeExample *FakeExampleTransactorSession) FakeEmitWithdrawalQueued(withdrawalRoot [32]byte, withdrawal FakeExampleWithdrawal) (*types.Transaction, error) {
	return _FakeExample.Contract.FakeEmitWithdrawalQueued(&_FakeExample.TransactOpts, withdrawalRoot, withdrawal)
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

// FakeExampleWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the FakeExample contract.
type FakeExampleWithdrawalQueuedIterator struct {
	Event *FakeExampleWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *FakeExampleWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FakeExampleWithdrawalQueued)
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
		it.Event = new(FakeExampleWithdrawalQueued)
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
func (it *FakeExampleWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FakeExampleWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FakeExampleWithdrawalQueued represents a WithdrawalQueued event raised by the FakeExample contract.
type FakeExampleWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     FakeExampleWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_FakeExample *FakeExampleFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*FakeExampleWithdrawalQueuedIterator, error) {

	logs, sub, err := _FakeExample.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &FakeExampleWithdrawalQueuedIterator{contract: _FakeExample.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_FakeExample *FakeExampleFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *FakeExampleWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _FakeExample.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FakeExampleWithdrawalQueued)
				if err := _FakeExample.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_FakeExample *FakeExampleFilterer) ParseWithdrawalQueued(log types.Log) (*FakeExampleWithdrawalQueued, error) {
	event := new(FakeExampleWithdrawalQueued)
	if err := _FakeExample.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
