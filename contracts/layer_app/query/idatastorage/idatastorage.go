// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idatastorage

import (
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

// IdatastorageABI is the input ABI used to generate the binding from.
const IdatastorageABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Idatastorage is an auto generated Go binding around a Solidity contract.
type Idatastorage struct {
	IdatastorageCaller     // Read-only binding to the contract
	IdatastorageTransactor // Write-only binding to the contract
	IdatastorageFilterer   // Log filterer for contract events
}

// IdatastorageCaller is an auto generated read-only Go binding around a Solidity contract.
type IdatastorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdatastorageTransactor is an auto generated write-only Go binding around a Solidity contract.
type IdatastorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdatastorageFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IdatastorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdatastorageSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IdatastorageSession struct {
	Contract     *Idatastorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdatastorageCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IdatastorageCallerSession struct {
	Contract *IdatastorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IdatastorageTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IdatastorageTransactorSession struct {
	Contract     *IdatastorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IdatastorageRaw is an auto generated low-level Go binding around a Solidity contract.
type IdatastorageRaw struct {
	Contract *Idatastorage // Generic contract binding to access the raw methods on
}

// IdatastorageCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IdatastorageCallerRaw struct {
	Contract *IdatastorageCaller // Generic read-only contract binding to access the raw methods on
}

// IdatastorageTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IdatastorageTransactorRaw struct {
	Contract *IdatastorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdatastorage creates a new instance of Idatastorage, bound to a specific deployed contract.
func NewIdatastorage(address common.Address, backend bind.ContractBackend) (*Idatastorage, error) {
	contract, err := bindIdatastorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Idatastorage{IdatastorageCaller: IdatastorageCaller{contract: contract}, IdatastorageTransactor: IdatastorageTransactor{contract: contract}, IdatastorageFilterer: IdatastorageFilterer{contract: contract}}, nil
}

// NewIdatastorageCaller creates a new read-only instance of Idatastorage, bound to a specific deployed contract.
func NewIdatastorageCaller(address common.Address, caller bind.ContractCaller) (*IdatastorageCaller, error) {
	contract, err := bindIdatastorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdatastorageCaller{contract: contract}, nil
}

// NewIdatastorageTransactor creates a new write-only instance of Idatastorage, bound to a specific deployed contract.
func NewIdatastorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IdatastorageTransactor, error) {
	contract, err := bindIdatastorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdatastorageTransactor{contract: contract}, nil
}

// NewIdatastorageFilterer creates a new log filterer instance of Idatastorage, bound to a specific deployed contract.
func NewIdatastorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IdatastorageFilterer, error) {
	contract, err := bindIdatastorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdatastorageFilterer{contract: contract}, nil
}

// bindIdatastorage binds a generic wrapper to an already deployed contract.
func bindIdatastorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdatastorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idatastorage *IdatastorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Idatastorage.Contract.IdatastorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idatastorage *IdatastorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Idatastorage.Contract.IdatastorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idatastorage *IdatastorageRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Idatastorage.Contract.IdatastorageTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idatastorage *IdatastorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Idatastorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idatastorage *IdatastorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Idatastorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idatastorage *IdatastorageTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Idatastorage.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Idatastorage *IdatastorageCaller) GetData(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Idatastorage.contract.Call(opts, out, "getData", key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Idatastorage *IdatastorageSession) GetData(key string) (*big.Int, error) {
	return _Idatastorage.Contract.GetData(&_Idatastorage.CallOpts, key)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Idatastorage *IdatastorageCallerSession) GetData(key string) (*big.Int, error) {
	return _Idatastorage.Contract.GetData(&_Idatastorage.CallOpts, key)
}
