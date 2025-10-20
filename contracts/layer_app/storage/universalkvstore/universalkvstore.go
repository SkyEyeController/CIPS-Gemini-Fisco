// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalkvstore

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// UniversalkvstoreABI is the input ABI used to generate the binding from.
const UniversalkvstoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ValueSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniversalkvstoreBin is the compiled bytecode used for deploying new contracts.
var UniversalkvstoreBin = "0x608060405234801561001057600080fd5b506105dc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063693ec85e1461003b578063e942b5161461006b575b600080fd5b61005560048036038101906100509190610313565b610087565b60405161006291906103f9565b60405180910390f35b6100856004803603810190610080919061041b565b61015a565b005b60606000838360405160200161009e9291906104db565b60405160208183030381529060405280519060200120905060008082815260200190815260200160002080546100d390610523565b80601f01602080910402602001604051908101604052809291908181526020018280546100ff90610523565b801561014c5780601f106101215761010080835404028352916020019161014c565b820191906000526020600020905b81548152906001019060200180831161012f57829003601f168201915b505050505091505092915050565b6000848460405160200161016f9291906104db565b604051602081830303815290604052805190602001209050828260008084815260200190815260200160002091906101a8929190610201565b5084846040516101b99291906104db565b60405180910390207fc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f584846040516101f2929190610582565b60405180910390a25050505050565b82805461020d90610523565b90600052602060002090601f01602090048101928261022f5760008555610276565b82601f1061024857803560ff1916838001178555610276565b82800160010185558215610276579182015b8281111561027557823582559160200191906001019061025a565b5b5090506102839190610287565b5090565b5b808211156102a0576000816000905550600101610288565b5090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126102d3576102d26102ae565b5b8235905067ffffffffffffffff8111156102f0576102ef6102b3565b5b60208301915083600182028301111561030c5761030b6102b8565b5b9250929050565b6000806020838503121561032a576103296102a4565b5b600083013567ffffffffffffffff811115610348576103476102a9565b5b610354858286016102bd565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561039a57808201518184015260208101905061037f565b838111156103a9576000848401525b50505050565b6000601f19601f8301169050919050565b60006103cb82610360565b6103d5818561036b565b93506103e581856020860161037c565b6103ee816103af565b840191505092915050565b6000602082019050818103600083015261041381846103c0565b905092915050565b60008060008060408587031215610435576104346102a4565b5b600085013567ffffffffffffffff811115610453576104526102a9565b5b61045f878288016102bd565b9450945050602085013567ffffffffffffffff811115610482576104816102a9565b5b61048e878288016102bd565b925092505092959194509250565b600081905092915050565b82818337600083830152505050565b60006104c2838561049c565b93506104cf8385846104a7565b82840190509392505050565b60006104e88284866104b6565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061053b57607f821691505b6020821081141561054f5761054e6104f4565b5b50919050565b6000610561838561036b565b935061056e8385846104a7565b610577836103af565b840190509392505050565b6000602082019050818103600083015261059d818486610555565b9050939250505056fea2646970667358221220914b386ac33f6e9931fd89bdd5284ad174e86faab36b06544b4d3aa28880646f64736f6c634300080b0033"
var UniversalkvstoreSMBin = "0x"

// DeployUniversalkvstore deploys a new contract, binding an instance of Universalkvstore to it.
func DeployUniversalkvstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Universalkvstore, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalkvstoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(UniversalkvstoreSMBin)
	} else {
		bytecode = common.FromHex(UniversalkvstoreBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, UniversalkvstoreABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Universalkvstore{UniversalkvstoreCaller: UniversalkvstoreCaller{contract: contract}, UniversalkvstoreTransactor: UniversalkvstoreTransactor{contract: contract}, UniversalkvstoreFilterer: UniversalkvstoreFilterer{contract: contract}}, nil
}

func AsyncDeployUniversalkvstore(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalkvstoreABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(UniversalkvstoreSMBin)
	} else {
		bytecode = common.FromHex(UniversalkvstoreBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, UniversalkvstoreABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Universalkvstore is an auto generated Go binding around a Solidity contract.
type Universalkvstore struct {
	UniversalkvstoreCaller     // Read-only binding to the contract
	UniversalkvstoreTransactor // Write-only binding to the contract
	UniversalkvstoreFilterer   // Log filterer for contract events
}

// UniversalkvstoreCaller is an auto generated read-only Go binding around a Solidity contract.
type UniversalkvstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalkvstoreTransactor is an auto generated write-only Go binding around a Solidity contract.
type UniversalkvstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalkvstoreFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type UniversalkvstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalkvstoreSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type UniversalkvstoreSession struct {
	Contract     *Universalkvstore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniversalkvstoreCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type UniversalkvstoreCallerSession struct {
	Contract *UniversalkvstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniversalkvstoreTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type UniversalkvstoreTransactorSession struct {
	Contract     *UniversalkvstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniversalkvstoreRaw is an auto generated low-level Go binding around a Solidity contract.
type UniversalkvstoreRaw struct {
	Contract *Universalkvstore // Generic contract binding to access the raw methods on
}

// UniversalkvstoreCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type UniversalkvstoreCallerRaw struct {
	Contract *UniversalkvstoreCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalkvstoreTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type UniversalkvstoreTransactorRaw struct {
	Contract *UniversalkvstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalkvstore creates a new instance of Universalkvstore, bound to a specific deployed contract.
func NewUniversalkvstore(address common.Address, backend bind.ContractBackend) (*Universalkvstore, error) {
	contract, err := bindUniversalkvstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Universalkvstore{UniversalkvstoreCaller: UniversalkvstoreCaller{contract: contract}, UniversalkvstoreTransactor: UniversalkvstoreTransactor{contract: contract}, UniversalkvstoreFilterer: UniversalkvstoreFilterer{contract: contract}}, nil
}

// NewUniversalkvstoreCaller creates a new read-only instance of Universalkvstore, bound to a specific deployed contract.
func NewUniversalkvstoreCaller(address common.Address, caller bind.ContractCaller) (*UniversalkvstoreCaller, error) {
	contract, err := bindUniversalkvstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalkvstoreCaller{contract: contract}, nil
}

// NewUniversalkvstoreTransactor creates a new write-only instance of Universalkvstore, bound to a specific deployed contract.
func NewUniversalkvstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalkvstoreTransactor, error) {
	contract, err := bindUniversalkvstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalkvstoreTransactor{contract: contract}, nil
}

// NewUniversalkvstoreFilterer creates a new log filterer instance of Universalkvstore, bound to a specific deployed contract.
func NewUniversalkvstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalkvstoreFilterer, error) {
	contract, err := bindUniversalkvstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalkvstoreFilterer{contract: contract}, nil
}

// bindUniversalkvstore binds a generic wrapper to an already deployed contract.
func bindUniversalkvstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalkvstoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalkvstore *UniversalkvstoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Universalkvstore.Contract.UniversalkvstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalkvstore *UniversalkvstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.UniversalkvstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalkvstore *UniversalkvstoreRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.UniversalkvstoreTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalkvstore *UniversalkvstoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Universalkvstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalkvstore *UniversalkvstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalkvstore *UniversalkvstoreTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_Universalkvstore *UniversalkvstoreCaller) Get(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Universalkvstore.contract.Call(opts, out, "get", key)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_Universalkvstore *UniversalkvstoreSession) Get(key string) (string, error) {
	return _Universalkvstore.Contract.Get(&_Universalkvstore.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_Universalkvstore *UniversalkvstoreCallerSession) Get(key string) (string, error) {
	return _Universalkvstore.Contract.Get(&_Universalkvstore.CallOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Universalkvstore *UniversalkvstoreTransactor) Set(opts *bind.TransactOpts, key string, value string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Universalkvstore.contract.TransactWithResult(opts, out, "set", key, value)
	return transaction, receipt, err
}

func (_Universalkvstore *UniversalkvstoreTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Universalkvstore.contract.AsyncTransact(opts, handler, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Universalkvstore *UniversalkvstoreSession) Set(key string, value string) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.Set(&_Universalkvstore.TransactOpts, key, value)
}

func (_Universalkvstore *UniversalkvstoreSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _Universalkvstore.Contract.AsyncSet(handler, &_Universalkvstore.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Universalkvstore *UniversalkvstoreTransactorSession) Set(key string, value string) (*types.Transaction, *types.Receipt, error) {
	return _Universalkvstore.Contract.Set(&_Universalkvstore.TransactOpts, key, value)
}

func (_Universalkvstore *UniversalkvstoreTransactorSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _Universalkvstore.Contract.AsyncSet(handler, &_Universalkvstore.TransactOpts, key, value)
}

// UniversalkvstoreValueSet represents a ValueSet event raised by the Universalkvstore contract.
type UniversalkvstoreValueSet struct {
	Key   common.Hash
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchValueSet is a free log subscription operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_Universalkvstore *UniversalkvstoreFilterer) WatchValueSet(fromBlock *int64, handler func(int, []types.Log), key string) (string, error) {
	return _Universalkvstore.contract.WatchLogs(fromBlock, handler, "ValueSet", key)
}

func (_Universalkvstore *UniversalkvstoreFilterer) WatchAllValueSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Universalkvstore.contract.WatchLogs(fromBlock, handler, "ValueSet")
}

// ParseValueSet is a log parse operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_Universalkvstore *UniversalkvstoreFilterer) ParseValueSet(log types.Log) (*UniversalkvstoreValueSet, error) {
	event := new(UniversalkvstoreValueSet)
	if err := _Universalkvstore.contract.UnpackLog(event, "ValueSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchValueSet is a free log subscription operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_Universalkvstore *UniversalkvstoreSession) WatchValueSet(fromBlock *int64, handler func(int, []types.Log), key string) (string, error) {
	return _Universalkvstore.Contract.WatchValueSet(fromBlock, handler, key)
}

func (_Universalkvstore *UniversalkvstoreSession) WatchAllValueSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Universalkvstore.Contract.WatchAllValueSet(fromBlock, handler)
}

// ParseValueSet is a log parse operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_Universalkvstore *UniversalkvstoreSession) ParseValueSet(log types.Log) (*UniversalkvstoreValueSet, error) {
	return _Universalkvstore.Contract.ParseValueSet(log)
}
