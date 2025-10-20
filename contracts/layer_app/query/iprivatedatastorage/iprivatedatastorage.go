// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iprivatedatastorage

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

// IprivatedatastorageABI is the input ABI used to generate the binding from.
const IprivatedatastorageABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getDataRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Iprivatedatastorage is an auto generated Go binding around a Solidity contract.
type Iprivatedatastorage struct {
	IprivatedatastorageCaller     // Read-only binding to the contract
	IprivatedatastorageTransactor // Write-only binding to the contract
	IprivatedatastorageFilterer   // Log filterer for contract events
}

// IprivatedatastorageCaller is an auto generated read-only Go binding around a Solidity contract.
type IprivatedatastorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IprivatedatastorageTransactor is an auto generated write-only Go binding around a Solidity contract.
type IprivatedatastorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IprivatedatastorageFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IprivatedatastorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IprivatedatastorageSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IprivatedatastorageSession struct {
	Contract     *Iprivatedatastorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IprivatedatastorageCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IprivatedatastorageCallerSession struct {
	Contract *IprivatedatastorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IprivatedatastorageTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IprivatedatastorageTransactorSession struct {
	Contract     *IprivatedatastorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IprivatedatastorageRaw is an auto generated low-level Go binding around a Solidity contract.
type IprivatedatastorageRaw struct {
	Contract *Iprivatedatastorage // Generic contract binding to access the raw methods on
}

// IprivatedatastorageCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IprivatedatastorageCallerRaw struct {
	Contract *IprivatedatastorageCaller // Generic read-only contract binding to access the raw methods on
}

// IprivatedatastorageTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IprivatedatastorageTransactorRaw struct {
	Contract *IprivatedatastorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIprivatedatastorage creates a new instance of Iprivatedatastorage, bound to a specific deployed contract.
func NewIprivatedatastorage(address common.Address, backend bind.ContractBackend) (*Iprivatedatastorage, error) {
	contract, err := bindIprivatedatastorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iprivatedatastorage{IprivatedatastorageCaller: IprivatedatastorageCaller{contract: contract}, IprivatedatastorageTransactor: IprivatedatastorageTransactor{contract: contract}, IprivatedatastorageFilterer: IprivatedatastorageFilterer{contract: contract}}, nil
}

// NewIprivatedatastorageCaller creates a new read-only instance of Iprivatedatastorage, bound to a specific deployed contract.
func NewIprivatedatastorageCaller(address common.Address, caller bind.ContractCaller) (*IprivatedatastorageCaller, error) {
	contract, err := bindIprivatedatastorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IprivatedatastorageCaller{contract: contract}, nil
}

// NewIprivatedatastorageTransactor creates a new write-only instance of Iprivatedatastorage, bound to a specific deployed contract.
func NewIprivatedatastorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IprivatedatastorageTransactor, error) {
	contract, err := bindIprivatedatastorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IprivatedatastorageTransactor{contract: contract}, nil
}

// NewIprivatedatastorageFilterer creates a new log filterer instance of Iprivatedatastorage, bound to a specific deployed contract.
func NewIprivatedatastorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IprivatedatastorageFilterer, error) {
	contract, err := bindIprivatedatastorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IprivatedatastorageFilterer{contract: contract}, nil
}

// bindIprivatedatastorage binds a generic wrapper to an already deployed contract.
func bindIprivatedatastorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IprivatedatastorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iprivatedatastorage *IprivatedatastorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iprivatedatastorage.Contract.IprivatedatastorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iprivatedatastorage *IprivatedatastorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Iprivatedatastorage.Contract.IprivatedatastorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iprivatedatastorage *IprivatedatastorageRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Iprivatedatastorage.Contract.IprivatedatastorageTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iprivatedatastorage *IprivatedatastorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iprivatedatastorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iprivatedatastorage *IprivatedatastorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Iprivatedatastorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iprivatedatastorage *IprivatedatastorageTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Iprivatedatastorage.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetDataRecord is a free data retrieval call binding the contract method 0x828b6ee3.
//
// Solidity: function getDataRecord(bytes32 dataHash) constant returns(bytes)
func (_Iprivatedatastorage *IprivatedatastorageCaller) GetDataRecord(opts *bind.CallOpts, dataHash [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Iprivatedatastorage.contract.Call(opts, out, "getDataRecord", dataHash)
	return *ret0, err
}

// GetDataRecord is a free data retrieval call binding the contract method 0x828b6ee3.
//
// Solidity: function getDataRecord(bytes32 dataHash) constant returns(bytes)
func (_Iprivatedatastorage *IprivatedatastorageSession) GetDataRecord(dataHash [32]byte) ([]byte, error) {
	return _Iprivatedatastorage.Contract.GetDataRecord(&_Iprivatedatastorage.CallOpts, dataHash)
}

// GetDataRecord is a free data retrieval call binding the contract method 0x828b6ee3.
//
// Solidity: function getDataRecord(bytes32 dataHash) constant returns(bytes)
func (_Iprivatedatastorage *IprivatedatastorageCallerSession) GetDataRecord(dataHash [32]byte) ([]byte, error) {
	return _Iprivatedatastorage.Contract.GetDataRecord(&_Iprivatedatastorage.CallOpts, dataHash)
}
