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
const PrivatedataqueryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"queryEncryptedData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PrivatedataqueryBin is the compiled bytecode used for deploying new contracts.
var PrivatedataqueryBin = "0x608060405234801561001057600080fd5b5061045b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637c41bda214610030575b600080fd5b61004a60048036038101906100459190610191565b610060565b604051610057919061026a565b60405180910390f35b60608273ffffffffffffffffffffffffffffffffffffffff1663828b6ee3836040518263ffffffff1660e01b815260040161009b919061029b565b600060405180830381865afa1580156100b8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906100e191906103dc565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610128826100fd565b9050919050565b6101388161011d565b811461014357600080fd5b50565b6000813590506101558161012f565b92915050565b6000819050919050565b61016e8161015b565b811461017957600080fd5b50565b60008135905061018b81610165565b92915050565b600080604083850312156101a8576101a76100f3565b5b60006101b685828601610146565b92505060206101c78582860161017c565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561020b5780820151818401526020810190506101f0565b8381111561021a576000848401525b50505050565b6000601f19601f8301169050919050565b600061023c826101d1565b61024681856101dc565b93506102568185602086016101ed565b61025f81610220565b840191505092915050565b600060208201905081810360008301526102848184610231565b905092915050565b6102958161015b565b82525050565b60006020820190506102b0600083018461028c565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102f882610220565b810181811067ffffffffffffffff82111715610317576103166102c0565b5b80604052505050565b600061032a6100e9565b905061033682826102ef565b919050565b600067ffffffffffffffff821115610356576103556102c0565b5b61035f82610220565b9050602081019050919050565b600061037f61037a8461033b565b610320565b90508281526020810184848401111561039b5761039a6102bb565b5b6103a68482856101ed565b509392505050565b600082601f8301126103c3576103c26102b6565b5b81516103d384826020860161036c565b91505092915050565b6000602082840312156103f2576103f16100f3565b5b600082015167ffffffffffffffff8111156104105761040f6100f8565b5b61041c848285016103ae565b9150509291505056fea264697066735822122017cd966de2da8499ecb0cc25be1aedd38ea1f3d1c9832440b7d58276efbaada564736f6c634300080b0033"
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
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(bytes)
func (_Privatedataquery *PrivatedataqueryCaller) QueryEncryptedData(opts *bind.CallOpts, storageContract common.Address, dataHash [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Privatedataquery.contract.Call(opts, out, "queryEncryptedData", storageContract, dataHash)
	return *ret0, err
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x7c41bda2.
//
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(bytes)
func (_Privatedataquery *PrivatedataquerySession) QueryEncryptedData(storageContract common.Address, dataHash [32]byte) ([]byte, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, dataHash)
}

// QueryEncryptedData is a free data retrieval call binding the contract method 0x7c41bda2.
//
// Solidity: function queryEncryptedData(address storageContract, bytes32 dataHash) constant returns(bytes)
func (_Privatedataquery *PrivatedataqueryCallerSession) QueryEncryptedData(storageContract common.Address, dataHash [32]byte) ([]byte, error) {
	return _Privatedataquery.Contract.QueryEncryptedData(&_Privatedataquery.CallOpts, storageContract, dataHash)
}
