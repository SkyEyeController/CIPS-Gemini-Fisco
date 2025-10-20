// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datastorage

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

// DatastorageABI is the input ABI used to generate the binding from.
const DatastorageABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DatastorageBin is the compiled bytecode used for deploying new contracts.
var DatastorageBin = "0x608060405234801561001057600080fd5b5061060a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806358537ca914610046578063ae55c88814610076578063af22068e146100a6575b600080fd5b610060600480360381019061005b919061038b565b6100c2565b60405161006d91906103ef565b60405180910390f35b610090600480360381019061008b919061038b565b610117565b60405161009d9190610423565b60405180910390f35b6100c060048036038101906100bb919061046a565b61015e565b005b600080826040516020016100d69190610540565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16915050919050565b6000808260405160200161012b9190610540565b60405160208183030381529060405280519060200120905060008082815260200190815260200160002054915050919050565b6000826040516020016101719190610540565b6040516020818303038152906040528051906020012090506001600082815260200190815260200160002060009054906101000a900460ff16156101ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e1906105b4565b60405180910390fd5b8160008083815260200190815260200160002081905550600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102988261024f565b810181811067ffffffffffffffff821117156102b7576102b6610260565b5b80604052505050565b60006102ca610231565b90506102d6828261028f565b919050565b600067ffffffffffffffff8211156102f6576102f5610260565b5b6102ff8261024f565b9050602081019050919050565b82818337600083830152505050565b600061032e610329846102db565b6102c0565b90508281526020810184848401111561034a5761034961024a565b5b61035584828561030c565b509392505050565b600082601f83011261037257610371610245565b5b813561038284826020860161031b565b91505092915050565b6000602082840312156103a1576103a061023b565b5b600082013567ffffffffffffffff8111156103bf576103be610240565b5b6103cb8482850161035d565b91505092915050565b60008115159050919050565b6103e9816103d4565b82525050565b600060208201905061040460008301846103e0565b92915050565b6000819050919050565b61041d8161040a565b82525050565b60006020820190506104386000830184610414565b92915050565b6104478161040a565b811461045257600080fd5b50565b6000813590506104648161043e565b92915050565b600080604083850312156104815761048061023b565b5b600083013567ffffffffffffffff81111561049f5761049e610240565b5b6104ab8582860161035d565b92505060206104bc85828601610455565b9150509250929050565b600081519050919050565b600081905092915050565b60005b838110156104fa5780820151818401526020810190506104df565b83811115610509576000848401525b50505050565b600061051a826104c6565b61052481856104d1565b93506105348185602086016104dc565b80840191505092915050565b600061054c828461050f565b915081905092915050565b600082825260208201905092915050565b7f4b657920616c7265616479206578697374730000000000000000000000000000600082015250565b600061059e601283610557565b91506105a982610568565b602082019050919050565b600060208201905081810360008301526105cd81610591565b905091905056fea2646970667358221220ecf0c17a7e3c6564c3a46ad1bcf00f7f2a78f4e747384ac4ea810d5fee7b003664736f6c634300080b0033"
var DatastorageSMBin = "0x"

// DeployDatastorage deploys a new contract, binding an instance of Datastorage to it.
func DeployDatastorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Datastorage, error) {
	parsed, err := abi.JSON(strings.NewReader(DatastorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DatastorageSMBin)
	} else {
		bytecode = common.FromHex(DatastorageBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, DatastorageABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Datastorage{DatastorageCaller: DatastorageCaller{contract: contract}, DatastorageTransactor: DatastorageTransactor{contract: contract}, DatastorageFilterer: DatastorageFilterer{contract: contract}}, nil
}

func AsyncDeployDatastorage(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(DatastorageABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DatastorageSMBin)
	} else {
		bytecode = common.FromHex(DatastorageBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, DatastorageABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Datastorage is an auto generated Go binding around a Solidity contract.
type Datastorage struct {
	DatastorageCaller     // Read-only binding to the contract
	DatastorageTransactor // Write-only binding to the contract
	DatastorageFilterer   // Log filterer for contract events
}

// DatastorageCaller is an auto generated read-only Go binding around a Solidity contract.
type DatastorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageTransactor is an auto generated write-only Go binding around a Solidity contract.
type DatastorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type DatastorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type DatastorageSession struct {
	Contract     *Datastorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatastorageCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type DatastorageCallerSession struct {
	Contract *DatastorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DatastorageTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type DatastorageTransactorSession struct {
	Contract     *DatastorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DatastorageRaw is an auto generated low-level Go binding around a Solidity contract.
type DatastorageRaw struct {
	Contract *Datastorage // Generic contract binding to access the raw methods on
}

// DatastorageCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type DatastorageCallerRaw struct {
	Contract *DatastorageCaller // Generic read-only contract binding to access the raw methods on
}

// DatastorageTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type DatastorageTransactorRaw struct {
	Contract *DatastorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatastorage creates a new instance of Datastorage, bound to a specific deployed contract.
func NewDatastorage(address common.Address, backend bind.ContractBackend) (*Datastorage, error) {
	contract, err := bindDatastorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datastorage{DatastorageCaller: DatastorageCaller{contract: contract}, DatastorageTransactor: DatastorageTransactor{contract: contract}, DatastorageFilterer: DatastorageFilterer{contract: contract}}, nil
}

// NewDatastorageCaller creates a new read-only instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageCaller(address common.Address, caller bind.ContractCaller) (*DatastorageCaller, error) {
	contract, err := bindDatastorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatastorageCaller{contract: contract}, nil
}

// NewDatastorageTransactor creates a new write-only instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DatastorageTransactor, error) {
	contract, err := bindDatastorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatastorageTransactor{contract: contract}, nil
}

// NewDatastorageFilterer creates a new log filterer instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DatastorageFilterer, error) {
	contract, err := bindDatastorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatastorageFilterer{contract: contract}, nil
}

// bindDatastorage binds a generic wrapper to an already deployed contract.
func bindDatastorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatastorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastorage *DatastorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datastorage.Contract.DatastorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastorage *DatastorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.DatastorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastorage *DatastorageRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.DatastorageTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastorage *DatastorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datastorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastorage *DatastorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastorage *DatastorageTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Datastorage *DatastorageCaller) GetData(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datastorage.contract.Call(opts, out, "getData", key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Datastorage *DatastorageSession) GetData(key string) (*big.Int, error) {
	return _Datastorage.Contract.GetData(&_Datastorage.CallOpts, key)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string key) constant returns(uint256)
func (_Datastorage *DatastorageCallerSession) GetData(key string) (*big.Int, error) {
	return _Datastorage.Contract.GetData(&_Datastorage.CallOpts, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Datastorage *DatastorageCaller) KeyExists(opts *bind.CallOpts, key string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Datastorage.contract.Call(opts, out, "keyExists", key)
	return *ret0, err
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Datastorage *DatastorageSession) KeyExists(key string) (bool, error) {
	return _Datastorage.Contract.KeyExists(&_Datastorage.CallOpts, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x58537ca9.
//
// Solidity: function keyExists(string key) constant returns(bool)
func (_Datastorage *DatastorageCallerSession) KeyExists(key string) (bool, error) {
	return _Datastorage.Contract.KeyExists(&_Datastorage.CallOpts, key)
}

// SetData is a paid mutator transaction binding the contract method 0xaf22068e.
//
// Solidity: function setData(string key, uint256 value) returns()
func (_Datastorage *DatastorageTransactor) SetData(opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Datastorage.contract.TransactWithResult(opts, out, "setData", key, value)
	return transaction, receipt, err
}

func (_Datastorage *DatastorageTransactor) AsyncSetData(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, error) {
	return _Datastorage.contract.AsyncTransact(opts, handler, "setData", key, value)
}

// SetData is a paid mutator transaction binding the contract method 0xaf22068e.
//
// Solidity: function setData(string key, uint256 value) returns()
func (_Datastorage *DatastorageSession) SetData(key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.SetData(&_Datastorage.TransactOpts, key, value)
}

func (_Datastorage *DatastorageSession) AsyncSetData(handler func(*types.Receipt, error), key string, value *big.Int) (*types.Transaction, error) {
	return _Datastorage.Contract.AsyncSetData(handler, &_Datastorage.TransactOpts, key, value)
}

// SetData is a paid mutator transaction binding the contract method 0xaf22068e.
//
// Solidity: function setData(string key, uint256 value) returns()
func (_Datastorage *DatastorageTransactorSession) SetData(key string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Datastorage.Contract.SetData(&_Datastorage.TransactOpts, key, value)
}

func (_Datastorage *DatastorageTransactorSession) AsyncSetData(handler func(*types.Receipt, error), key string, value *big.Int) (*types.Transaction, error) {
	return _Datastorage.Contract.AsyncSetData(handler, &_Datastorage.TransactOpts, key, value)
}
