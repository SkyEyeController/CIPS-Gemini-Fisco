// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package calculator

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

// CalculatorABI is the input ABI used to generate the binding from.
const CalculatorABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difference\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"product\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quotient\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// CalculatorBin is the compiled bytecode used for deploying new contracts.
var CalculatorBin = "0x608060405234801561001057600080fd5b50610492806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638dfa436314610030575b600080fd5b61004a6004803603810190610045919061016c565b610063565b60405161005a94939291906101bb565b60405180910390f35b6000806000808486610075919061022f565b9350848610156100ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b1906102e2565b60405180910390fd5b84866100c69190610302565b925084866100d49190610336565b9150600085141561011a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610111906103dc565b60405180910390fd5b8486610126919061042b565b905092959194509250565b600080fd5b6000819050919050565b61014981610136565b811461015457600080fd5b50565b60008135905061016681610140565b92915050565b6000806040838503121561018357610182610131565b5b600061019185828601610157565b92505060206101a285828601610157565b9150509250929050565b6101b581610136565b82525050565b60006080820190506101d060008301876101ac565b6101dd60208301866101ac565b6101ea60408301856101ac565b6101f760608301846101ac565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061023a82610136565b915061024583610136565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561027a57610279610200565b5b828201905092915050565b600082825260208201905092915050565b7f5375627472616374696f6e20756e646572666c6f770000000000000000000000600082015250565b60006102cc601583610285565b91506102d782610296565b602082019050919050565b600060208201905081810360008301526102fb816102bf565b9050919050565b600061030d82610136565b915061031883610136565b92508282101561032b5761032a610200565b5b828203905092915050565b600061034182610136565b915061034c83610136565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561038557610384610200565b5b828202905092915050565b7f4469766973696f6e206279207a65726f00000000000000000000000000000000600082015250565b60006103c6601083610285565b91506103d182610390565b602082019050919050565b600060208201905081810360008301526103f5816103b9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061043682610136565b915061044183610136565b925082610451576104506103fc565b5b82820490509291505056fea264697066735822122095afec3b7efda8f271a74316f3a732b3ecfadc275e322521bf9889567adfcc8f64736f6c634300080b0033"
var CalculatorSMBin = "0x"

// DeployCalculator deploys a new contract, binding an instance of Calculator to it.
func DeployCalculator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Calculator, error) {
	parsed, err := abi.JSON(strings.NewReader(CalculatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(CalculatorSMBin)
	} else {
		bytecode = common.FromHex(CalculatorBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, CalculatorABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Calculator{CalculatorCaller: CalculatorCaller{contract: contract}, CalculatorTransactor: CalculatorTransactor{contract: contract}, CalculatorFilterer: CalculatorFilterer{contract: contract}}, nil
}

func AsyncDeployCalculator(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(CalculatorABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(CalculatorSMBin)
	} else {
		bytecode = common.FromHex(CalculatorBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, CalculatorABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Calculator is an auto generated Go binding around a Solidity contract.
type Calculator struct {
	CalculatorCaller     // Read-only binding to the contract
	CalculatorTransactor // Write-only binding to the contract
	CalculatorFilterer   // Log filterer for contract events
}

// CalculatorCaller is an auto generated read-only Go binding around a Solidity contract.
type CalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorTransactor is an auto generated write-only Go binding around a Solidity contract.
type CalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CalculatorSession struct {
	Contract     *Calculator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalculatorCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CalculatorCallerSession struct {
	Contract *CalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CalculatorTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CalculatorTransactorSession struct {
	Contract     *CalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CalculatorRaw is an auto generated low-level Go binding around a Solidity contract.
type CalculatorRaw struct {
	Contract *Calculator // Generic contract binding to access the raw methods on
}

// CalculatorCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CalculatorCallerRaw struct {
	Contract *CalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// CalculatorTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CalculatorTransactorRaw struct {
	Contract *CalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalculator creates a new instance of Calculator, bound to a specific deployed contract.
func NewCalculator(address common.Address, backend bind.ContractBackend) (*Calculator, error) {
	contract, err := bindCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calculator{CalculatorCaller: CalculatorCaller{contract: contract}, CalculatorTransactor: CalculatorTransactor{contract: contract}, CalculatorFilterer: CalculatorFilterer{contract: contract}}, nil
}

// NewCalculatorCaller creates a new read-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorCaller(address common.Address, caller bind.ContractCaller) (*CalculatorCaller, error) {
	contract, err := bindCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorCaller{contract: contract}, nil
}

// NewCalculatorTransactor creates a new write-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CalculatorTransactor, error) {
	contract, err := bindCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorTransactor{contract: contract}, nil
}

// NewCalculatorFilterer creates a new log filterer instance of Calculator, bound to a specific deployed contract.
func NewCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CalculatorFilterer, error) {
	contract, err := bindCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalculatorFilterer{contract: contract}, nil
}

// bindCalculator binds a generic wrapper to an already deployed contract.
func bindCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalculatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.CalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Calculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Calculator.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Calculate is a free data retrieval call binding the contract method 0x8dfa4363.
//
// Solidity: function calculate(uint256 a, uint256 b) constant returns(uint256 sum, uint256 difference, uint256 product, uint256 quotient)
func (_Calculator *CalculatorCaller) Calculate(opts *bind.CallOpts, a *big.Int, b *big.Int) (struct {
	Sum        *big.Int
	Difference *big.Int
	Product    *big.Int
	Quotient   *big.Int
}, error) {
	ret := new(struct {
		Sum        *big.Int
		Difference *big.Int
		Product    *big.Int
		Quotient   *big.Int
	})
	out := ret
	err := _Calculator.contract.Call(opts, out, "calculate", a, b)
	return *ret, err
}

// Calculate is a free data retrieval call binding the contract method 0x8dfa4363.
//
// Solidity: function calculate(uint256 a, uint256 b) constant returns(uint256 sum, uint256 difference, uint256 product, uint256 quotient)
func (_Calculator *CalculatorSession) Calculate(a *big.Int, b *big.Int) (struct {
	Sum        *big.Int
	Difference *big.Int
	Product    *big.Int
	Quotient   *big.Int
}, error) {
	return _Calculator.Contract.Calculate(&_Calculator.CallOpts, a, b)
}

// Calculate is a free data retrieval call binding the contract method 0x8dfa4363.
//
// Solidity: function calculate(uint256 a, uint256 b) constant returns(uint256 sum, uint256 difference, uint256 product, uint256 quotient)
func (_Calculator *CalculatorCallerSession) Calculate(a *big.Int, b *big.Int) (struct {
	Sum        *big.Int
	Difference *big.Int
	Product    *big.Int
	Quotient   *big.Int
}, error) {
	return _Calculator.Contract.Calculate(&_Calculator.CallOpts, a, b)
}
