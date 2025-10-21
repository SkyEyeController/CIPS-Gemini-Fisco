// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataquery

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

// DataqueryABI is the input ABI used to generate the binding from.
const DataqueryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataqueryBin is the compiled bytecode used for deploying new contracts.
var DataqueryBin = "0x608060405234801561001057600080fd5b50610465806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063de0cb0ad14610030575b600080fd5b61004a6004803603810190610045919061029c565b610060565b6040516100579190610311565b60405180910390f35b60008273ffffffffffffffffffffffffffffffffffffffff1663ae55c888836040518263ffffffff1660e01b815260040161009b91906103b4565b602060405180830381865afa1580156100b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dc9190610402565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610123826100f8565b9050919050565b61013381610118565b811461013e57600080fd5b50565b6000813590506101508161012a565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101a982610160565b810181811067ffffffffffffffff821117156101c8576101c7610171565b5b80604052505050565b60006101db6100e4565b90506101e782826101a0565b919050565b600067ffffffffffffffff82111561020757610206610171565b5b61021082610160565b9050602081019050919050565b82818337600083830152505050565b600061023f61023a846101ec565b6101d1565b90508281526020810184848401111561025b5761025a61015b565b5b61026684828561021d565b509392505050565b600082601f83011261028357610282610156565b5b813561029384826020860161022c565b91505092915050565b600080604083850312156102b3576102b26100ee565b5b60006102c185828601610141565b925050602083013567ffffffffffffffff8111156102e2576102e16100f3565b5b6102ee8582860161026e565b9150509250929050565b6000819050919050565b61030b816102f8565b82525050565b60006020820190506103266000830184610302565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561036657808201518184015260208101905061034b565b83811115610375576000848401525b50505050565b60006103868261032c565b6103908185610337565b93506103a0818560208601610348565b6103a981610160565b840191505092915050565b600060208201905081810360008301526103ce818461037b565b905092915050565b6103df816102f8565b81146103ea57600080fd5b50565b6000815190506103fc816103d6565b92915050565b600060208284031215610418576104176100ee565b5b6000610426848285016103ed565b9150509291505056fea2646970667358221220bdc1ca1d0dcec4b5c17f8678cdc6467f262da7ed3350902afe472725a24aa8b764736f6c634300080b0033"
var DataquerySMBin = "0x"

// DeployDataquery deploys a new contract, binding an instance of Dataquery to it.
func DeployDataquery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Dataquery, error) {
	parsed, err := abi.JSON(strings.NewReader(DataqueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DataquerySMBin)
	} else {
		bytecode = common.FromHex(DataqueryBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, DataqueryABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Dataquery{DataqueryCaller: DataqueryCaller{contract: contract}, DataqueryTransactor: DataqueryTransactor{contract: contract}, DataqueryFilterer: DataqueryFilterer{contract: contract}}, nil
}

func AsyncDeployDataquery(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(DataqueryABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DataquerySMBin)
	} else {
		bytecode = common.FromHex(DataqueryBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, DataqueryABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Dataquery is an auto generated Go binding around a Solidity contract.
type Dataquery struct {
	DataqueryCaller     // Read-only binding to the contract
	DataqueryTransactor // Write-only binding to the contract
	DataqueryFilterer   // Log filterer for contract events
}

// DataqueryCaller is an auto generated read-only Go binding around a Solidity contract.
type DataqueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataqueryTransactor is an auto generated write-only Go binding around a Solidity contract.
type DataqueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataqueryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type DataqueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataquerySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type DataquerySession struct {
	Contract     *Dataquery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataqueryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type DataqueryCallerSession struct {
	Contract *DataqueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DataqueryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type DataqueryTransactorSession struct {
	Contract     *DataqueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DataqueryRaw is an auto generated low-level Go binding around a Solidity contract.
type DataqueryRaw struct {
	Contract *Dataquery // Generic contract binding to access the raw methods on
}

// DataqueryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type DataqueryCallerRaw struct {
	Contract *DataqueryCaller // Generic read-only contract binding to access the raw methods on
}

// DataqueryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type DataqueryTransactorRaw struct {
	Contract *DataqueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataquery creates a new instance of Dataquery, bound to a specific deployed contract.
func NewDataquery(address common.Address, backend bind.ContractBackend) (*Dataquery, error) {
	contract, err := bindDataquery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dataquery{DataqueryCaller: DataqueryCaller{contract: contract}, DataqueryTransactor: DataqueryTransactor{contract: contract}, DataqueryFilterer: DataqueryFilterer{contract: contract}}, nil
}

// NewDataqueryCaller creates a new read-only instance of Dataquery, bound to a specific deployed contract.
func NewDataqueryCaller(address common.Address, caller bind.ContractCaller) (*DataqueryCaller, error) {
	contract, err := bindDataquery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataqueryCaller{contract: contract}, nil
}

// NewDataqueryTransactor creates a new write-only instance of Dataquery, bound to a specific deployed contract.
func NewDataqueryTransactor(address common.Address, transactor bind.ContractTransactor) (*DataqueryTransactor, error) {
	contract, err := bindDataquery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataqueryTransactor{contract: contract}, nil
}

// NewDataqueryFilterer creates a new log filterer instance of Dataquery, bound to a specific deployed contract.
func NewDataqueryFilterer(address common.Address, filterer bind.ContractFilterer) (*DataqueryFilterer, error) {
	contract, err := bindDataquery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataqueryFilterer{contract: contract}, nil
}

// bindDataquery binds a generic wrapper to an already deployed contract.
func bindDataquery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataqueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dataquery *DataqueryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dataquery.Contract.DataqueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dataquery *DataqueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Dataquery.Contract.DataqueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dataquery *DataqueryRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Dataquery.Contract.DataqueryTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dataquery *DataqueryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dataquery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dataquery *DataqueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Dataquery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dataquery *DataqueryTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Dataquery.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0xde0cb0ad.
//
// Solidity: function getValue(address contractAddress, string key) constant returns(uint256)
func (_Dataquery *DataqueryCaller) GetValue(opts *bind.CallOpts, contractAddress common.Address, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dataquery.contract.Call(opts, out, "getValue", contractAddress, key)
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0xde0cb0ad.
//
// Solidity: function getValue(address contractAddress, string key) constant returns(uint256)
func (_Dataquery *DataquerySession) GetValue(contractAddress common.Address, key string) (*big.Int, error) {
	return _Dataquery.Contract.GetValue(&_Dataquery.CallOpts, contractAddress, key)
}

// GetValue is a free data retrieval call binding the contract method 0xde0cb0ad.
//
// Solidity: function getValue(address contractAddress, string key) constant returns(uint256)
func (_Dataquery *DataqueryCallerSession) GetValue(contractAddress common.Address, key string) (*big.Int, error) {
	return _Dataquery.Contract.GetValue(&_Dataquery.CallOpts, contractAddress, key)
}
