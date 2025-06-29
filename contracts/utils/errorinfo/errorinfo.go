// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package errorinfo

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

// ErrorinfoABI is the input ABI used to generate the binding from.
const ErrorinfoABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ExecuteError\",\"type\":\"event\"}]"

// ErrorinfoBin is the compiled bytecode used for deploying new contracts.
var ErrorinfoBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aab0503e86c8ed8183c5ca904c78f7aefc2a68fbb306529eda8edaa003e3d28464736f6c634300080b0033"
var ErrorinfoSMBin = "0x"

// DeployErrorinfo deploys a new contract, binding an instance of Errorinfo to it.
func DeployErrorinfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Errorinfo, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorinfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(ErrorinfoSMBin)
	} else {
		bytecode = common.FromHex(ErrorinfoBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, ErrorinfoABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Errorinfo{ErrorinfoCaller: ErrorinfoCaller{contract: contract}, ErrorinfoTransactor: ErrorinfoTransactor{contract: contract}, ErrorinfoFilterer: ErrorinfoFilterer{contract: contract}}, nil
}

func AsyncDeployErrorinfo(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorinfoABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(ErrorinfoSMBin)
	} else {
		bytecode = common.FromHex(ErrorinfoBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, ErrorinfoABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Errorinfo is an auto generated Go binding around a Solidity contract.
type Errorinfo struct {
	ErrorinfoCaller     // Read-only binding to the contract
	ErrorinfoTransactor // Write-only binding to the contract
	ErrorinfoFilterer   // Log filterer for contract events
}

// ErrorinfoCaller is an auto generated read-only Go binding around a Solidity contract.
type ErrorinfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorinfoTransactor is an auto generated write-only Go binding around a Solidity contract.
type ErrorinfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorinfoFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ErrorinfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorinfoSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ErrorinfoSession struct {
	Contract     *Errorinfo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorinfoCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ErrorinfoCallerSession struct {
	Contract *ErrorinfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ErrorinfoTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ErrorinfoTransactorSession struct {
	Contract     *ErrorinfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ErrorinfoRaw is an auto generated low-level Go binding around a Solidity contract.
type ErrorinfoRaw struct {
	Contract *Errorinfo // Generic contract binding to access the raw methods on
}

// ErrorinfoCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ErrorinfoCallerRaw struct {
	Contract *ErrorinfoCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorinfoTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ErrorinfoTransactorRaw struct {
	Contract *ErrorinfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrorinfo creates a new instance of Errorinfo, bound to a specific deployed contract.
func NewErrorinfo(address common.Address, backend bind.ContractBackend) (*Errorinfo, error) {
	contract, err := bindErrorinfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Errorinfo{ErrorinfoCaller: ErrorinfoCaller{contract: contract}, ErrorinfoTransactor: ErrorinfoTransactor{contract: contract}, ErrorinfoFilterer: ErrorinfoFilterer{contract: contract}}, nil
}

// NewErrorinfoCaller creates a new read-only instance of Errorinfo, bound to a specific deployed contract.
func NewErrorinfoCaller(address common.Address, caller bind.ContractCaller) (*ErrorinfoCaller, error) {
	contract, err := bindErrorinfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorinfoCaller{contract: contract}, nil
}

// NewErrorinfoTransactor creates a new write-only instance of Errorinfo, bound to a specific deployed contract.
func NewErrorinfoTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorinfoTransactor, error) {
	contract, err := bindErrorinfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorinfoTransactor{contract: contract}, nil
}

// NewErrorinfoFilterer creates a new log filterer instance of Errorinfo, bound to a specific deployed contract.
func NewErrorinfoFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorinfoFilterer, error) {
	contract, err := bindErrorinfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorinfoFilterer{contract: contract}, nil
}

// bindErrorinfo binds a generic wrapper to an already deployed contract.
func bindErrorinfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorinfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errorinfo *ErrorinfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Errorinfo.Contract.ErrorinfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errorinfo *ErrorinfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Errorinfo.Contract.ErrorinfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errorinfo *ErrorinfoRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Errorinfo.Contract.ErrorinfoTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errorinfo *ErrorinfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Errorinfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errorinfo *ErrorinfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Errorinfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errorinfo *ErrorinfoTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Errorinfo.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// ErrorinfoExecuteError represents a ExecuteError event raised by the Errorinfo contract.
type ErrorinfoExecuteError struct {
	Code    *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchExecuteError is a free log subscription operation binding the contract event 0xc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc31.
//
// Solidity: event ExecuteError(uint256 indexed code, string message)
func (_Errorinfo *ErrorinfoFilterer) WatchExecuteError(fromBlock *int64, handler func(int, []types.Log), code *big.Int) (string, error) {
	return _Errorinfo.contract.WatchLogs(fromBlock, handler, "ExecuteError", code)
}

func (_Errorinfo *ErrorinfoFilterer) WatchAllExecuteError(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Errorinfo.contract.WatchLogs(fromBlock, handler, "ExecuteError")
}

// ParseExecuteError is a log parse operation binding the contract event 0xc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc31.
//
// Solidity: event ExecuteError(uint256 indexed code, string message)
func (_Errorinfo *ErrorinfoFilterer) ParseExecuteError(log types.Log) (*ErrorinfoExecuteError, error) {
	event := new(ErrorinfoExecuteError)
	if err := _Errorinfo.contract.UnpackLog(event, "ExecuteError", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchExecuteError is a free log subscription operation binding the contract event 0xc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc31.
//
// Solidity: event ExecuteError(uint256 indexed code, string message)
func (_Errorinfo *ErrorinfoSession) WatchExecuteError(fromBlock *int64, handler func(int, []types.Log), code *big.Int) (string, error) {
	return _Errorinfo.Contract.WatchExecuteError(fromBlock, handler, code)
}

func (_Errorinfo *ErrorinfoSession) WatchAllExecuteError(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Errorinfo.Contract.WatchAllExecuteError(fromBlock, handler)
}

// ParseExecuteError is a log parse operation binding the contract event 0xc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc31.
//
// Solidity: event ExecuteError(uint256 indexed code, string message)
func (_Errorinfo *ErrorinfoSession) ParseExecuteError(log types.Log) (*ErrorinfoExecuteError, error) {
	return _Errorinfo.Contract.ParseExecuteError(log)
}
