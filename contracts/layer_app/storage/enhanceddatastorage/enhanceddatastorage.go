// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package enhanceddatastorage

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

// EnhanceddatastorageABI is the input ABI used to generate the binding from.
const EnhanceddatastorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tag\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTag\",\"type\":\"uint256\"}],\"name\":\"TagUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getDataRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tag\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cipherValue\",\"type\":\"uint256\"}],\"name\":\"setEncryptedData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPlainData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newTag\",\"type\":\"uint256\"}],\"name\":\"updateTag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EnhanceddatastorageBin is the compiled bytecode used for deploying new contracts.
var EnhanceddatastorageBin = "0x608060405234801561001057600080fd5b50610bc2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630c8832b914610067578063140d1de61461008357806358537ca91461009f57806374d2f180146100cf578063a5cfa75f14610100578063ae55c8881461011c575b600080fd5b610081600480360381019061007c91906107c2565b61014c565b005b61009d600480360381019061009891906107c2565b610280565b005b6100b960048036038101906100b4919061081e565b6103b4565b6040516100c69190610882565b60405180910390f35b6100e960048036038101906100e4919061081e565b610409565b6040516100f79291906108ac565b60405180910390f35b61011a600480360381019061011591906107c2565b6104c2565b005b6101366004803603810190610131919061081e565b6105e8565b60405161014391906108d5565b60405180910390f35b60008260405160200161015f919061096a565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16156101d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101cf906109de565b60405180910390fd5b604051806040016040528083815260200160008152506000808381526020019081526020016000206000820151816000015560208201518160010155905050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff60006040516102739190610a43565b60405180910390a2505050565b600082604051602001610293919061096a565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff161561030c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610303906109de565b60405180910390fd5b604051806040016040528083815260200160018152506000808381526020019081526020016000206000820151816000015560208201518160010155905050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff60016040516103a79190610a99565b60405180910390a2505050565b600080826040516020016103c8919061096a565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16915050919050565b60008060008360405160200161041f919061096a565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90610b00565b60405180910390fd5b6000806000838152602001908152602001600020905080600001548160010154935093505050915091565b6001811115610506576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fd90610b6c565b60405180910390fd5b600082604051602001610519919061096a565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058890610b00565b60405180910390fd5b8160008083815260200190815260200160002060010181905550807fc7926db7aa801ec5bdf0df192a46f92ac57384f2c47ff5c30b3dcee541ba8b36836040516105db91906108d5565b60405180910390a2505050565b600080826040516020016105fc919061096a565b60405160208183030381529060405280519060200120905060008082815260200190815260200160002060000154915050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61069982610650565b810181811067ffffffffffffffff821117156106b8576106b7610661565b5b80604052505050565b60006106cb610632565b90506106d78282610690565b919050565b600067ffffffffffffffff8211156106f7576106f6610661565b5b61070082610650565b9050602081019050919050565b82818337600083830152505050565b600061072f61072a846106dc565b6106c1565b90508281526020810184848401111561074b5761074a61064b565b5b61075684828561070d565b509392505050565b600082601f83011261077357610772610646565b5b813561078384826020860161071c565b91505092915050565b6000819050919050565b61079f8161078c565b81146107aa57600080fd5b50565b6000813590506107bc81610796565b92915050565b600080604083850312156107d9576107d861063c565b5b600083013567ffffffffffffffff8111156107f7576107f6610641565b5b6108038582860161075e565b9250506020610814858286016107ad565b9150509250929050565b6000602082840312156108345761083361063c565b5b600082013567ffffffffffffffff81111561085257610851610641565b5b61085e8482850161075e565b91505092915050565b60008115159050919050565b61087c81610867565b82525050565b60006020820190506108976000830184610873565b92915050565b6108a68161078c565b82525050565b60006040820190506108c1600083018561089d565b6108ce602083018461089d565b9392505050565b60006020820190506108ea600083018461089d565b92915050565b600081519050919050565b600081905092915050565b60005b83811015610924578082015181840152602081019050610909565b83811115610933576000848401525b50505050565b6000610944826108f0565b61094e81856108fb565b935061095e818560208601610906565b80840191505092915050565b60006109768284610939565b915081905092915050565b600082825260208201905092915050565b7f4b657920616c7265616479206578697374730000000000000000000000000000600082015250565b60006109c8601283610981565b91506109d382610992565b602082019050919050565b600060208201905081810360008301526109f7816109bb565b9050919050565b6000819050919050565b6000819050919050565b6000610a2d610a28610a23846109fe565b610a08565b61078c565b9050919050565b610a3d81610a12565b82525050565b6000602082019050610a586000830184610a34565b92915050565b6000819050919050565b6000610a83610a7e610a7984610a5e565b610a08565b61078c565b9050919050565b610a9381610a68565b82525050565b6000602082019050610aae6000830184610a8a565b92915050565b7f4b6579206e6f7420657869737400000000000000000000000000000000000000600082015250565b6000610aea600d83610981565b9150610af582610ab4565b602082019050919050565b60006020820190508181036000830152610b1981610add565b9050919050565b7f496e76616c6964207461672076616c7565000000000000000000000000000000600082015250565b6000610b56601183610981565b9150610b6182610b20565b602082019050919050565b60006020820190508181036000830152610b8581610b49565b905091905056fea264697066735822122050d1ce8e786e28646491c848530115f2908903467ed2910f35ed9b7a2c2bb56864736f6c634300080b0033"
var EnhanceddatastorageSMBin = "0x"

// DeployEnhanceddatastorage deploys a new contract, binding an instance of Enhanceddatastorage to it.
func DeployEnhanceddatastorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Enhanceddatastorage, error) {
	parsed, err := abi.JSON(strings.NewReader(EnhanceddatastorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(EnhanceddatastorageSMBin)
	} else {
		bytecode = common.FromHex(EnhanceddatastorageBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, EnhanceddatastorageABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Enhanceddatastorage{EnhanceddatastorageCaller: EnhanceddatastorageCaller{contract: contract}, EnhanceddatastorageTransactor: EnhanceddatastorageTransactor{contract: contract}, EnhanceddatastorageFilterer: EnhanceddatastorageFilterer{contract: contract}}, nil
}

func AsyncDeployEnhanceddatastorage(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(EnhanceddatastorageABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(EnhanceddatastorageSMBin)
	} else {
		bytecode = common.FromHex(EnhanceddatastorageBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, EnhanceddatastorageABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Enhanceddatastorage is an auto generated Go binding around a Solidity contract.
type Enhanceddatastorage struct {
	EnhanceddatastorageCaller     // Read-only binding to the contract
	EnhanceddatastorageTransactor // Write-only binding to the contract
	EnhanceddatastorageFilterer   // Log filterer for contract events
}

// EnhanceddatastorageCaller is an auto generated read-only Go binding around a Solidity contract.
type EnhanceddatastorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnhanceddatastorageTransactor is an auto generated write-only Go binding around a Solidity contract.
type EnhanceddatastorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnhanceddatastorageFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EnhanceddatastorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnhanceddatastorageSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EnhanceddatastorageSession struct {
	Contract     *Enhanceddatastorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EnhanceddatastorageCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EnhanceddatastorageCallerSession struct {
	Contract *EnhanceddatastorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EnhanceddatastorageTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EnhanceddatastorageTransactorSession struct {
	Contract     *EnhanceddatastorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EnhanceddatastorageRaw is an auto generated low-level Go binding around a Solidity contract.
type EnhanceddatastorageRaw struct {
	Contract *Enhanceddatastorage // Generic contract binding to access the raw methods on
}

// EnhanceddatastorageCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EnhanceddatastorageCallerRaw struct {
	Contract *EnhanceddatastorageCaller // Generic read-only contract binding to access the raw methods on
}

// EnhanceddatastorageTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EnhanceddatastorageTransactorRaw struct {
	Contract *EnhanceddatastorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnhanceddatastorage creates a new instance of Enhanceddatastorage, bound to a specific deployed contract.
func NewEnhanceddatastorage(address common.Address, backend bind.ContractBackend) (*Enhanceddatastorage, error) {
	contract, err := bindEnhanceddatastorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enhanceddatastorage{EnhanceddatastorageCaller: EnhanceddatastorageCaller{contract: contract}, EnhanceddatastorageTransactor: EnhanceddatastorageTransactor{contract: contract}, EnhanceddatastorageFilterer: EnhanceddatastorageFilterer{contract: contract}}, nil
}

// NewEnhanceddatastorageCaller creates a new read-only instance of Enhanceddatastorage, bound to a specific deployed contract.
func NewEnhanceddatastorageCaller(address common.Address, caller bind.ContractCaller) (*EnhanceddatastorageCaller, error) {
	contract, err := bindEnhanceddatastorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnhanceddatastorageCaller{contract: contract}, nil
}

// NewEnhanceddatastorageTransactor creates a new write-only instance of Enhanceddatastorage, bound to a specific deployed contract.
func NewEnhanceddatastorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EnhanceddatastorageTransactor, error) {
	contract, err := bindEnhanceddatastorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnhanceddatastorageTransactor{contract: contract}, nil
}

// NewEnhanceddatastorageFilterer creates a new log filterer instance of Enhanceddatastorage, bound to a specific deployed contract.
func NewEnhanceddatastorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EnhanceddatastorageFilterer, error) {
	contract, err := bindEnhanceddatastorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnhanceddatastorageFilterer{contract: contract}, nil
}

// bindEnhanceddatastorage binds a generic wrapper to an already deployed contract.
func bindEnhanceddatastorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnhanceddatastorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enhanceddatastorage *EnhanceddatastorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enhanceddatastorage.Contract.EnhanceddatastorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enhanceddatastorage *EnhanceddatastorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.EnhanceddatastorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enhanceddatastorage *EnhanceddatastorageRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.EnhanceddatastorageTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enhanceddatastorage *EnhanceddatastorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enhanceddatastorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enhanceddatastorage *EnhanceddatastorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enhanceddatastorage *EnhanceddatastorageTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Enhanceddatastorage *EnhanceddatastorageCaller) GetData(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Enhanceddatastorage.contract.Call(opts, out, "getData", key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Enhanceddatastorage *EnhanceddatastorageSession) GetData(key string) (*big.Int, error) {
	return _Enhanceddatastorage.Contract.GetData(&_Enhanceddatastorage.CallOpts, key)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Enhanceddatastorage *EnhanceddatastorageCallerSession) GetData(key string) (*big.Int, error) {
	return _Enhanceddatastorage.Contract.GetData(&_Enhanceddatastorage.CallOpts, key)
}

// GetDataRecord is a free data retrieval call binding the contract method 0x74d2f180.
//
// Solidity: function getDataRecord(string key) constant returns(uint256 value, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageCaller) GetDataRecord(opts *bind.CallOpts, key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	ret := new(struct {
		Value *big.Int
		Tag   *big.Int
	})
	out := ret
	err := _Enhanceddatastorage.contract.Call(opts, out, "getDataRecord", key)
	return *ret, err
}

// GetDataRecord is a free data retrieval call binding the contract method 0x74d2f180.
//
// Solidity: function getDataRecord(string key) constant returns(uint256 value, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageSession) GetDataRecord(key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Enhanceddatastorage.Contract.GetDataRecord(&_Enhanceddatastorage.CallOpts, key)
}

// GetDataRecord is a free data retrieval call binding the contract method 0x74d2f180.
//
// Solidity: function getDataRecord(string key) constant returns(uint256 value, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageCallerSession) GetDataRecord(key string) (struct {
	Value *big.Int
	Tag   *big.Int
}, error) {
	return _Enhanceddatastorage.Contract.GetDataRecord(&_Enhanceddatastorage.CallOpts, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Enhanceddatastorage *EnhanceddatastorageCaller) KeyExists(opts *bind.CallOpts, key string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Enhanceddatastorage.contract.Call(opts, out, "keyExists", key)
	return *ret0, err
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Enhanceddatastorage *EnhanceddatastorageSession) KeyExists(key string) (bool, error) {
	return _Enhanceddatastorage.Contract.KeyExists(&_Enhanceddatastorage.CallOpts, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Enhanceddatastorage *EnhanceddatastorageCallerSession) KeyExists(key string) (bool, error) {
	return _Enhanceddatastorage.Contract.KeyExists(&_Enhanceddatastorage.CallOpts, key)
}

// SetEncryptedData is a paid mutator transaction binding the contract method 0x140d1de6.
//
// Solidity: function setEncryptedData(string key, uint256 cipherValue) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactor) SetEncryptedData(opts *bind.TransactOpts, key string, cipherValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Enhanceddatastorage.contract.TransactWithResult(opts, out, "setEncryptedData", key, cipherValue)
	return transaction, receipt, err
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactor) AsyncSetEncryptedData(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, cipherValue *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.contract.AsyncTransact(opts, handler, "setEncryptedData", key, cipherValue)
}

// SetEncryptedData is a paid mutator transaction binding the contract method 0x140d1de6.
//
// Solidity: function setEncryptedData(string key, uint256 cipherValue) returns()
func (_Enhanceddatastorage *EnhanceddatastorageSession) SetEncryptedData(key string, cipherValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.SetEncryptedData(&_Enhanceddatastorage.TransactOpts, key, cipherValue)
}

func (_Enhanceddatastorage *EnhanceddatastorageSession) AsyncSetEncryptedData(handler func(*types.Receipt, error), key string, cipherValue *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncSetEncryptedData(handler, &_Enhanceddatastorage.TransactOpts, key, cipherValue)
}

// SetEncryptedData is a paid mutator transaction binding the contract method 0x140d1de6.
//
// Solidity: function setEncryptedData(string key, uint256 cipherValue) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) SetEncryptedData(key string, cipherValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.SetEncryptedData(&_Enhanceddatastorage.TransactOpts, key, cipherValue)
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) AsyncSetEncryptedData(handler func(*types.Receipt, error), key string, cipherValue *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncSetEncryptedData(handler, &_Enhanceddatastorage.TransactOpts, key, cipherValue)
}

// SetPlainData is a paid mutator transaction binding the contract method 0x0c8832b9.
//
// Solidity: function setPlainData(string key, uint256 value) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactor) SetPlainData(opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Enhanceddatastorage.contract.TransactWithResult(opts, out, "setPlainData", key, value)
	return transaction, receipt, err
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactor) AsyncSetPlainData(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.contract.AsyncTransact(opts, handler, "setPlainData", key, value)
}

// SetPlainData is a paid mutator transaction binding the contract method 0x0c8832b9.
//
// Solidity: function setPlainData(string key, uint256 value) returns()
func (_Enhanceddatastorage *EnhanceddatastorageSession) SetPlainData(key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.SetPlainData(&_Enhanceddatastorage.TransactOpts, key, value)
}

func (_Enhanceddatastorage *EnhanceddatastorageSession) AsyncSetPlainData(handler func(*types.Receipt, error), key string, value *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncSetPlainData(handler, &_Enhanceddatastorage.TransactOpts, key, value)
}

// SetPlainData is a paid mutator transaction binding the contract method 0x0c8832b9.
//
// Solidity: function setPlainData(string key, uint256 value) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) SetPlainData(key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.SetPlainData(&_Enhanceddatastorage.TransactOpts, key, value)
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) AsyncSetPlainData(handler func(*types.Receipt, error), key string, value *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncSetPlainData(handler, &_Enhanceddatastorage.TransactOpts, key, value)
}

// UpdateTag is a paid mutator transaction binding the contract method 0xa5cfa75f.
//
// Solidity: function updateTag(string key, uint256 newTag) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactor) UpdateTag(opts *bind.TransactOpts, key string, newTag *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Enhanceddatastorage.contract.TransactWithResult(opts, out, "updateTag", key, newTag)
	return transaction, receipt, err
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactor) AsyncUpdateTag(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, newTag *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.contract.AsyncTransact(opts, handler, "updateTag", key, newTag)
}

// UpdateTag is a paid mutator transaction binding the contract method 0xa5cfa75f.
//
// Solidity: function updateTag(string key, uint256 newTag) returns()
func (_Enhanceddatastorage *EnhanceddatastorageSession) UpdateTag(key string, newTag *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.UpdateTag(&_Enhanceddatastorage.TransactOpts, key, newTag)
}

func (_Enhanceddatastorage *EnhanceddatastorageSession) AsyncUpdateTag(handler func(*types.Receipt, error), key string, newTag *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncUpdateTag(handler, &_Enhanceddatastorage.TransactOpts, key, newTag)
}

// UpdateTag is a paid mutator transaction binding the contract method 0xa5cfa75f.
//
// Solidity: function updateTag(string key, uint256 newTag) returns()
func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) UpdateTag(key string, newTag *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Enhanceddatastorage.Contract.UpdateTag(&_Enhanceddatastorage.TransactOpts, key, newTag)
}

func (_Enhanceddatastorage *EnhanceddatastorageTransactorSession) AsyncUpdateTag(handler func(*types.Receipt, error), key string, newTag *big.Int) (*types.Transaction, error) {
	return _Enhanceddatastorage.Contract.AsyncUpdateTag(handler, &_Enhanceddatastorage.TransactOpts, key, newTag)
}

// EnhanceddatastorageDataStored represents a DataStored event raised by the Enhanceddatastorage contract.
type EnhanceddatastorageDataStored struct {
	KeyHash [32]byte
	Tag     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchDataStored is a free log subscription operation binding the contract event 0x0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff.
//
// Solidity: event DataStored(bytes32 indexed keyHash, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageFilterer) WatchDataStored(fromBlock *int64, handler func(int, []types.Log), keyHash [32]byte) (string, error) {
	return _Enhanceddatastorage.contract.WatchLogs(fromBlock, handler, "DataStored", keyHash)
}

func (_Enhanceddatastorage *EnhanceddatastorageFilterer) WatchAllDataStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Enhanceddatastorage.contract.WatchLogs(fromBlock, handler, "DataStored")
}

// ParseDataStored is a log parse operation binding the contract event 0x0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff.
//
// Solidity: event DataStored(bytes32 indexed keyHash, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageFilterer) ParseDataStored(log types.Log) (*EnhanceddatastorageDataStored, error) {
	event := new(EnhanceddatastorageDataStored)
	if err := _Enhanceddatastorage.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0x0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff.
//
// Solidity: event DataStored(bytes32 indexed keyHash, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageSession) WatchDataStored(fromBlock *int64, handler func(int, []types.Log), keyHash [32]byte) (string, error) {
	return _Enhanceddatastorage.Contract.WatchDataStored(fromBlock, handler, keyHash)
}

func (_Enhanceddatastorage *EnhanceddatastorageSession) WatchAllDataStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Enhanceddatastorage.Contract.WatchAllDataStored(fromBlock, handler)
}

// ParseDataStored is a log parse operation binding the contract event 0x0b44a4099a1a4710e3d84492d3ab62eec832c496798a031e7a66d0d0e89cdbff.
//
// Solidity: event DataStored(bytes32 indexed keyHash, uint256 tag)
func (_Enhanceddatastorage *EnhanceddatastorageSession) ParseDataStored(log types.Log) (*EnhanceddatastorageDataStored, error) {
	return _Enhanceddatastorage.Contract.ParseDataStored(log)
}

// EnhanceddatastorageTagUpdated represents a TagUpdated event raised by the Enhanceddatastorage contract.
type EnhanceddatastorageTagUpdated struct {
	KeyHash [32]byte
	NewTag  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchTagUpdated is a free log subscription operation binding the contract event 0xc7926db7aa801ec5bdf0df192a46f92ac57384f2c47ff5c30b3dcee541ba8b36.
//
// Solidity: event TagUpdated(bytes32 indexed keyHash, uint256 newTag)
func (_Enhanceddatastorage *EnhanceddatastorageFilterer) WatchTagUpdated(fromBlock *int64, handler func(int, []types.Log), keyHash [32]byte) (string, error) {
	return _Enhanceddatastorage.contract.WatchLogs(fromBlock, handler, "TagUpdated", keyHash)
}

func (_Enhanceddatastorage *EnhanceddatastorageFilterer) WatchAllTagUpdated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Enhanceddatastorage.contract.WatchLogs(fromBlock, handler, "TagUpdated")
}

// ParseTagUpdated is a log parse operation binding the contract event 0xc7926db7aa801ec5bdf0df192a46f92ac57384f2c47ff5c30b3dcee541ba8b36.
//
// Solidity: event TagUpdated(bytes32 indexed keyHash, uint256 newTag)
func (_Enhanceddatastorage *EnhanceddatastorageFilterer) ParseTagUpdated(log types.Log) (*EnhanceddatastorageTagUpdated, error) {
	event := new(EnhanceddatastorageTagUpdated)
	if err := _Enhanceddatastorage.contract.UnpackLog(event, "TagUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTagUpdated is a free log subscription operation binding the contract event 0xc7926db7aa801ec5bdf0df192a46f92ac57384f2c47ff5c30b3dcee541ba8b36.
//
// Solidity: event TagUpdated(bytes32 indexed keyHash, uint256 newTag)
func (_Enhanceddatastorage *EnhanceddatastorageSession) WatchTagUpdated(fromBlock *int64, handler func(int, []types.Log), keyHash [32]byte) (string, error) {
	return _Enhanceddatastorage.Contract.WatchTagUpdated(fromBlock, handler, keyHash)
}

func (_Enhanceddatastorage *EnhanceddatastorageSession) WatchAllTagUpdated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Enhanceddatastorage.Contract.WatchAllTagUpdated(fromBlock, handler)
}

// ParseTagUpdated is a log parse operation binding the contract event 0xc7926db7aa801ec5bdf0df192a46f92ac57384f2c47ff5c30b3dcee541ba8b36.
//
// Solidity: event TagUpdated(bytes32 indexed keyHash, uint256 newTag)
func (_Enhanceddatastorage *EnhanceddatastorageSession) ParseTagUpdated(log types.Log) (*EnhanceddatastorageTagUpdated, error) {
	return _Enhanceddatastorage.Contract.ParseTagUpdated(log)
}
