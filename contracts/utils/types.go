// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

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

// TypesABI is the input ABI used to generate the binding from.
const TypesABI = "[]"

// TypesBin is the compiled bytecode used for deploying new contracts.
var TypesBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203ff6a6a7a8c539cd1d0d4243ebc9bcc08f8f0e764295b3eecd553d96f0ec060c64736f6c634300080b0033"
var TypesSMBin = "0x"

// DeployTypes deploys a new contract, binding an instance of Types to it.
func DeployTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Types, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TypesSMBin)
	} else {
		bytecode = common.FromHex(TypesBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TypesABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Types{TypesCaller: TypesCaller{contract: contract}, TypesTransactor: TypesTransactor{contract: contract}, TypesFilterer: TypesFilterer{contract: contract}}, nil
}

func AsyncDeployTypes(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TypesSMBin)
	} else {
		bytecode = common.FromHex(TypesBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TypesABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Types is an auto generated Go binding around a Solidity contract.
type Types struct {
	TypesCaller     // Read-only binding to the contract
	TypesTransactor // Write-only binding to the contract
	TypesFilterer   // Log filterer for contract events
}

// TypesCaller is an auto generated read-only Go binding around a Solidity contract.
type TypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesTransactor is an auto generated write-only Go binding around a Solidity contract.
type TypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TypesSession struct {
	Contract     *Types            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TypesCallerSession struct {
	Contract *TypesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TypesTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TypesTransactorSession struct {
	Contract     *TypesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesRaw is an auto generated low-level Go binding around a Solidity contract.
type TypesRaw struct {
	Contract *Types // Generic contract binding to access the raw methods on
}

// TypesCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TypesCallerRaw struct {
	Contract *TypesCaller // Generic read-only contract binding to access the raw methods on
}

// TypesTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TypesTransactorRaw struct {
	Contract *TypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypes creates a new instance of Types, bound to a specific deployed contract.
func NewTypes(address common.Address, backend bind.ContractBackend) (*Types, error) {
	contract, err := bindTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Types{TypesCaller: TypesCaller{contract: contract}, TypesTransactor: TypesTransactor{contract: contract}, TypesFilterer: TypesFilterer{contract: contract}}, nil
}

// NewTypesCaller creates a new read-only instance of Types, bound to a specific deployed contract.
func NewTypesCaller(address common.Address, caller bind.ContractCaller) (*TypesCaller, error) {
	contract, err := bindTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypesCaller{contract: contract}, nil
}

// NewTypesTransactor creates a new write-only instance of Types, bound to a specific deployed contract.
func NewTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*TypesTransactor, error) {
	contract, err := bindTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypesTransactor{contract: contract}, nil
}

// NewTypesFilterer creates a new log filterer instance of Types, bound to a specific deployed contract.
func NewTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*TypesFilterer, error) {
	contract, err := bindTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypesFilterer{contract: contract}, nil
}

// bindTypes binds a generic wrapper to an already deployed contract.
func bindTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.TypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Types.Contract.TypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Types.Contract.TypesTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Types.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Types.Contract.contract.TransactWithResult(opts, result, method, params...)
}
