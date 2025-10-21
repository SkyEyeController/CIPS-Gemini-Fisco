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
const PrivatedataqueryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"queryEncryptedData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tag\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PrivatedataqueryBin is the compiled bytecode used for deploying new contracts.
var PrivatedataqueryBin = "0x608060405234801561001057600080fd5b5061048a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636ee05a4f14610030575b600080fd5b61004a600480360381019061004591906102a0565b610061565b604051610058929190610315565b60405180910390f35b6000808373ffffffffffffffffffffffffffffffffffffffff166374d2f180846040518263ffffffff1660e01b815260040161009d91906103c6565b6040805180830381865afa1580156100b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dd9190610414565b915091509250929050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610127826100fc565b9050919050565b6101378161011c565b811461014257600080fd5b50565b6000813590506101548161012e565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101ad82610164565b810181811067ffffffffffffffff821117156101cc576101cb610175565b5b80604052505050565b60006101df6100e8565b90506101eb82826101a4565b919050565b600067ffffffffffffffff82111561020b5761020a610175565b5b61021482610164565b9050602081019050919050565b82818337600083830152505050565b600061024361023e846101f0565b6101d5565b90508281526020810184848401111561025f5761025e61015f565b5b61026a848285610221565b509392505050565b600082601f8301126102875761028661015a565b5b8135610297848260208601610230565b91505092915050565b600080604083850312156102b7576102b66100f2565b5b60006102c585828601610145565b925050602083013567ffffffffffffffff8111156102e6576102e56100f7565b5b6102f285828601610272565b9150509250929050565b6000819050919050565b61030f816102fc565b82525050565b600060408201905061032a6000830185610306565b6103376020830184610306565b9392505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561037857808201518184015260208101905061035d565b83811115610387576000848401525b50505050565b60006103988261033e565b6103a28185610349565b93506103b281856020860161035a565b6103bb81610164565b840191505092915050565b600060208201905081810360008301526103e0818461038d565b905092915050565b6103f1816102fc565b81146103fc57600080fd5b50565b60008151905061040e816103e8565b92915050565b6000806040838503121561042b5761042a6100f2565b5b6000610439858286016103ff565b925050602061044a858286016103ff565b915050925092905056fea2646970667358221220edc622e37763ae66ab0820a80b91fa83fdb0d959c3c71d0d4bc2ce393e54a47e64736f6c634300080b0033"
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

// QueryEncryptedData is a free data retrieval call binding the contract method 0x6ee05a4f.
//
// Solidity: function queryEncryptedData(address storageContract, string key) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataqueryCaller) QueryEncryptedData(opts *bind.CallOpts, storageContract common.Address, key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	ret := new(struct {
		Value *big.Int
		Tag   *big.Int
	})
	out := ret
	err := _Privatedataquery.contract.Call(opts, out, "queryEncryptedData", storageContract, key)
	return *ret, err
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x6ee05a4f.
//
// Solidity: function queryEncryptedData(address storageContract, string key) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataquerySession) QueryEncryptedData(storageContract common.Address, key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, key)
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x6ee05a4f.
//
// Solidity: function queryEncryptedData(address storageContract, string key) constant returns(uint256 value, uint256 tag)
func (_Privatedataquery *PrivatedataqueryCallerSession) QueryEncryptedData(storageContract common.Address, key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, key)
}
