// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package privatedataquery

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

// PrivatedataqueryABI is the input ABI used to generate the binding from.
const PrivatedataqueryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"queryEncryptedData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tag\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PrivatedataqueryBin is the compiled bytecode used for deploying new contracts.
var PrivatedataqueryBin = "0x608060405234801561001057600080fd5b506102cf806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637c41bda214610030575b600080fd5b61004a60048036038101906100459190610181565b610061565b6040516100589291906101da565b60405180910390f35b6000808373ffffffffffffffffffffffffffffffffffffffff1663828b6ee3846040518263ffffffff1660e01b815260040161009d9190610212565b6040805180830381865afa1580156100b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dd9190610259565b915091509250929050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610118826100ed565b9050919050565b6101288161010d565b811461013357600080fd5b50565b6000813590506101458161011f565b92915050565b6000819050919050565b61015e8161014b565b811461016957600080fd5b50565b60008135905061017b81610155565b92915050565b60008060408385031215610198576101976100e8565b5b60006101a685828601610136565b92505060206101b78582860161016c565b9150509250929050565b6000819050919050565b6101d4816101c1565b82525050565b60006040820190506101ef60008301856101cb565b6101fc60208301846101cb565b9392505050565b61020c8161014b565b82525050565b60006020820190506102276000830184610203565b92915050565b610236816101c1565b811461024157600080fd5b50565b6000815190506102538161022d565b92915050565b600080604083850312156102705761026f6100e8565b5b600061027e85828601610244565b925050602061028f85828601610244565b915050925092905056fea2646970667358221220a3eaa9d70749968c76e178d0922679c94bdb281a9ae0ed1a7c3aea9907bbf26b64736f6c634300080b0033"
var PrivatedataquerySMBin = "0x"

// DeployPrivatedataquery deploys a new contract, binding an instance of Privatedataquery to it.
func DeployPrivatedataquery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Privatedataquery, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatedataqueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(PrivatedataquerySMBin)
	} else {
		bytecode = common.FromHex(PrivatedataqueryBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, PrivatedataqueryABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Privatedataquery{PrivatedataqueryCaller: PrivatedataqueryCaller{contract: contract}, PrivatedataqueryTransactor: PrivatedataqueryTransactor{contract: contract}, PrivatedataqueryFilterer: PrivatedataqueryFilterer{contract: contract}}, nil
}

func AsyncDeployPrivatedataquery(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatedataqueryABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(PrivatedataquerySMBin)
	} else {
		bytecode = common.FromHex(PrivatedataqueryBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, PrivatedataqueryABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Privatedataquery is an auto generated Go binding around a Solidity contract.
type Privatedataquery struct {
	PrivatedataqueryCaller     // Read-only binding to the contract
	PrivatedataqueryTransactor // Write-only binding to the contract
	PrivatedataqueryFilterer   // Log filterer for contract events
}

// PrivatedataqueryCaller is an auto generated read-only Go binding around a Solidity contract.
type PrivatedataqueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatedataqueryTransactor is an auto generated write-only Go binding around a Solidity contract.
type PrivatedataqueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatedataqueryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type PrivatedataqueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatedataquerySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type PrivatedataquerySession struct {
	Contract     *Privatedataquery // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivatedataqueryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type PrivatedataqueryCallerSession struct {
	Contract *PrivatedataqueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PrivatedataqueryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type PrivatedataqueryTransactorSession struct {
	Contract     *PrivatedataqueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PrivatedataqueryRaw is an auto generated low-level Go binding around a Solidity contract.
type PrivatedataqueryRaw struct {
	Contract *Privatedataquery // Generic contract binding to access the raw methods on
}

// PrivatedataqueryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type PrivatedataqueryCallerRaw struct {
	Contract *PrivatedataqueryCaller // Generic read-only contract binding to access the raw methods on
}

// PrivatedataqueryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type PrivatedataqueryTransactorRaw struct {
	Contract *PrivatedataqueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivatedataquery creates a new instance of Privatedataquery, bound to a specific deployed contract.
func NewPrivatedataquery(address common.Address, backend bind.ContractBackend) (*Privatedataquery, error) {
	contract, err := bindPrivatedataquery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Privatedataquery{PrivatedataqueryCaller: PrivatedataqueryCaller{contract: contract}, PrivatedataqueryTransactor: PrivatedataqueryTransactor{contract: contract}, PrivatedataqueryFilterer: PrivatedataqueryFilterer{contract: contract}}, nil
}

// NewPrivatedataqueryCaller creates a new read-only instance of Privatedataquery, bound to a specific deployed contract.
func NewPrivatedataqueryCaller(address common.Address, caller bind.ContractCaller) (*PrivatedataqueryCaller, error) {
	contract, err := bindPrivatedataquery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatedataqueryCaller{contract: contract}, nil
}

// NewPrivatedataqueryTransactor creates a new write-only instance of Privatedataquery, bound to a specific deployed contract.
func NewPrivatedataqueryTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivatedataqueryTransactor, error) {
	contract, err := bindPrivatedataquery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatedataqueryTransactor{contract: contract}, nil
}

// NewPrivatedataqueryFilterer creates a new log filterer instance of Privatedataquery, bound to a specific deployed contract.
func NewPrivatedataqueryFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivatedataqueryFilterer, error) {
	contract, err := bindPrivatedataquery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivatedataqueryFilterer{contract: contract}, nil
}

// bindPrivatedataquery binds a generic wrapper to an already deployed contract.
func bindPrivatedataquery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatedataqueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Privatedataquery *PrivatedataqueryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Privatedataquery.Contract.PrivatedataqueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Privatedataquery *PrivatedataqueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Privatedataquery.Contract.PrivatedataqueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Privatedataquery *PrivatedataqueryRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Privatedataquery.Contract.PrivatedataqueryTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Privatedataquery *PrivatedataqueryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Privatedataquery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Privatedataquery *PrivatedataqueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Privatedataquery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Privatedataquery *PrivatedataqueryTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Privatedataquery.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x7c41bda2.
//
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataqueryCaller) QueryEncryptedData(opts *bind.CallOpts, storageContract common.Address, dataHash [32]byte) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	ret := new(struct {
		Value *big.Int
		Tag   *big.Int
	})
	out := ret
	err := _Privatedataquery.contract.Call(opts, out, "queryEncryptedData", storageContract, dataHash)
	return *ret, err
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x7c41bda2.
//
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataquerySession) QueryEncryptedData(storageContract common.Address, dataHash [32]byte) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, dataHash)
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x7c41bda2.
//
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataqueryCallerSession) QueryEncryptedData(storageContract common.Address, dataHash [32]byte) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, dataHash)
}
