// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transactionreg

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

// TransactionregABI is the input ABI used to generate the binding from.
const TransactionregABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transactionAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"TransactionRegister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transactionAddress\",\"type\":\"address\"}],\"name\":\"getTransactionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transactionAddress\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransactionregBin is the compiled bytecode used for deploying new contracts.
var TransactionregBin = "0x608060405234801561001057600080fd5b50600160009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610757806100846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806316ae6dde146100465780632801617e146100775780639507d39a14610093575b600080fd5b610060600480360381019061005b919061047e565b6100c4565b60405161006e92919061055d565b60405180910390f35b610091600480360381019061008c919061047e565b6101a9565b005b6100ad60048036038101906100a891906105b9565b610332565b6040516100bb9291906105f5565b60405180910390f35b6000606060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156101505760006040518060400160405280600781526020017f756e666f756e6400000000000000000000000000000000000000000000000000815250915091506101a4565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460405180602001604052806000815250915091505b915091565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461022e57606f60ff167fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc316040516102219061064b565b60405180910390a261032f565b6001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600180805490506102a4919061069a565b9050806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167f4a4aff16c381ad8929f8b834c8723c495b0ab0fc2de27b0e3786ef201496a33d60405160405180910390a3505b50565b6000606060008314156103635760006040518060600160405280602481526020016106fe6024913991509150610416565b60018080549050610374919061069a565b8311156103bc5760006040518060400160405280600781526020017f756e666f756e640000000000000000000000000000000000000000000000000081525091509150610416565b6000600184815481106103d2576103d16106ce565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806040518060200160405280600081525092509250505b915091565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061044b82610420565b9050919050565b61045b81610440565b811461046657600080fd5b50565b60008135905061047881610452565b92915050565b6000602082840312156104945761049361041b565b5b60006104a284828501610469565b91505092915050565b6000819050919050565b6104be816104ab565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104fe5780820151818401526020810190506104e3565b8381111561050d576000848401525b50505050565b6000601f19601f8301169050919050565b600061052f826104c4565b61053981856104cf565b93506105498185602086016104e0565b61055281610513565b840191505092915050565b600060408201905061057260008301856104b5565b81810360208301526105848184610524565b90509392505050565b610596816104ab565b81146105a157600080fd5b50565b6000813590506105b38161058d565b92915050565b6000602082840312156105cf576105ce61041b565b5b60006105dd848285016105a4565b91505092915050565b6105ef81610440565b82525050565b600060408201905061060a60008301856105e6565b818103602083015261061c8184610524565b90509392505050565b50565b60006106356000836104cf565b915061064082610625565b600082019050919050565b6000602082019050818103600083015261066481610628565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106a5826104ab565b91506106b0836104ab565b9250828210156106c3576106c261066b565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe7472616e73616374696f6e4964206d7573742062652067726561746572207468616e2030a2646970667358221220b49d7f9f9e19fb6f1b9012cccd8cad903136c02becaa42925d9231e110e0e55b64736f6c634300080b0033"
var TransactionregSMBin = "0x"

// DeployTransactionreg deploys a new contract, binding an instance of Transactionreg to it.
func DeployTransactionreg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transactionreg, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionregABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransactionregSMBin)
	} else {
		bytecode = common.FromHex(TransactionregBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TransactionregABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transactionreg{TransactionregCaller: TransactionregCaller{contract: contract}, TransactionregTransactor: TransactionregTransactor{contract: contract}, TransactionregFilterer: TransactionregFilterer{contract: contract}}, nil
}

func AsyncDeployTransactionreg(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionregABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransactionregSMBin)
	} else {
		bytecode = common.FromHex(TransactionregBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TransactionregABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transactionreg is an auto generated Go binding around a Solidity contract.
type Transactionreg struct {
	TransactionregCaller     // Read-only binding to the contract
	TransactionregTransactor // Write-only binding to the contract
	TransactionregFilterer   // Log filterer for contract events
}

// TransactionregCaller is an auto generated read-only Go binding around a Solidity contract.
type TransactionregCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionregTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransactionregTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionregFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransactionregFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionregSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransactionregSession struct {
	Contract     *Transactionreg   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionregCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransactionregCallerSession struct {
	Contract *TransactionregCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransactionregTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransactionregTransactorSession struct {
	Contract     *TransactionregTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransactionregRaw is an auto generated low-level Go binding around a Solidity contract.
type TransactionregRaw struct {
	Contract *Transactionreg // Generic contract binding to access the raw methods on
}

// TransactionregCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransactionregCallerRaw struct {
	Contract *TransactionregCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionregTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransactionregTransactorRaw struct {
	Contract *TransactionregTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionreg creates a new instance of Transactionreg, bound to a specific deployed contract.
func NewTransactionreg(address common.Address, backend bind.ContractBackend) (*Transactionreg, error) {
	contract, err := bindTransactionreg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transactionreg{TransactionregCaller: TransactionregCaller{contract: contract}, TransactionregTransactor: TransactionregTransactor{contract: contract}, TransactionregFilterer: TransactionregFilterer{contract: contract}}, nil
}

// NewTransactionregCaller creates a new read-only instance of Transactionreg, bound to a specific deployed contract.
func NewTransactionregCaller(address common.Address, caller bind.ContractCaller) (*TransactionregCaller, error) {
	contract, err := bindTransactionreg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionregCaller{contract: contract}, nil
}

// NewTransactionregTransactor creates a new write-only instance of Transactionreg, bound to a specific deployed contract.
func NewTransactionregTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionregTransactor, error) {
	contract, err := bindTransactionreg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionregTransactor{contract: contract}, nil
}

// NewTransactionregFilterer creates a new log filterer instance of Transactionreg, bound to a specific deployed contract.
func NewTransactionregFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionregFilterer, error) {
	contract, err := bindTransactionreg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionregFilterer{contract: contract}, nil
}

// bindTransactionreg binds a generic wrapper to an already deployed contract.
func bindTransactionreg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionregABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionreg *TransactionregRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactionreg.Contract.TransactionregCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionreg *TransactionregRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.TransactionregTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionreg *TransactionregRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.TransactionregTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionreg *TransactionregCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactionreg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionreg *TransactionregTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionreg *TransactionregTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transactionId) constant returns(address, string)
func (_Transactionreg *TransactionregCaller) Get(opts *bind.CallOpts, transactionId *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transactionreg.contract.Call(opts, out, "get", transactionId)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transactionId) constant returns(address, string)
func (_Transactionreg *TransactionregSession) Get(transactionId *big.Int) (common.Address, string, error) {
	return _Transactionreg.Contract.Get(&_Transactionreg.CallOpts, transactionId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transactionId) constant returns(address, string)
func (_Transactionreg *TransactionregCallerSession) Get(transactionId *big.Int) (common.Address, string, error) {
	return _Transactionreg.Contract.Get(&_Transactionreg.CallOpts, transactionId)
}

// GetTransactionId is a free data retrieval call binding the contract method 0x16ae6dde.
//
// Solidity: function getTransactionId(address transactionAddress) constant returns(uint256, string)
func (_Transactionreg *TransactionregCaller) GetTransactionId(opts *bind.CallOpts, transactionAddress common.Address) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transactionreg.contract.Call(opts, out, "getTransactionId", transactionAddress)
	return *ret0, *ret1, err
}

// GetTransactionId is a free data retrieval call binding the contract method 0x16ae6dde.
//
// Solidity: function getTransactionId(address transactionAddress) constant returns(uint256, string)
func (_Transactionreg *TransactionregSession) GetTransactionId(transactionAddress common.Address) (*big.Int, string, error) {
	return _Transactionreg.Contract.GetTransactionId(&_Transactionreg.CallOpts, transactionAddress)
}

// GetTransactionId is a free data retrieval call binding the contract method 0x16ae6dde.
//
// Solidity: function getTransactionId(address transactionAddress) constant returns(uint256, string)
func (_Transactionreg *TransactionregCallerSession) GetTransactionId(transactionAddress common.Address) (*big.Int, string, error) {
	return _Transactionreg.Contract.GetTransactionId(&_Transactionreg.CallOpts, transactionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transactionAddress) returns()
func (_Transactionreg *TransactionregTransactor) Set(opts *bind.TransactOpts, transactionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transactionreg.contract.TransactWithResult(opts, out, "set", transactionAddress)
	return transaction, receipt, err
}

func (_Transactionreg *TransactionregTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, transactionAddress common.Address) (*types.Transaction, error) {
	return _Transactionreg.contract.AsyncTransact(opts, handler, "set", transactionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transactionAddress) returns()
func (_Transactionreg *TransactionregSession) Set(transactionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.Set(&_Transactionreg.TransactOpts, transactionAddress)
}

func (_Transactionreg *TransactionregSession) AsyncSet(handler func(*types.Receipt, error), transactionAddress common.Address) (*types.Transaction, error) {
	return _Transactionreg.Contract.AsyncSet(handler, &_Transactionreg.TransactOpts, transactionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transactionAddress) returns()
func (_Transactionreg *TransactionregTransactorSession) Set(transactionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transactionreg.Contract.Set(&_Transactionreg.TransactOpts, transactionAddress)
}

func (_Transactionreg *TransactionregTransactorSession) AsyncSet(handler func(*types.Receipt, error), transactionAddress common.Address) (*types.Transaction, error) {
	return _Transactionreg.Contract.AsyncSet(handler, &_Transactionreg.TransactOpts, transactionAddress)
}

// TransactionregTransactionRegister represents a TransactionRegister event raised by the Transactionreg contract.
type TransactionregTransactionRegister struct {
	TransactionAddr common.Address
	TransactionId   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// WatchTransactionRegister is a free log subscription operation binding the contract event 0x4a4aff16c381ad8929f8b834c8723c495b0ab0fc2de27b0e3786ef201496a33d.
//
// Solidity: event TransactionRegister(address indexed transactionAddr, uint256 indexed transactionId)
func (_Transactionreg *TransactionregFilterer) WatchTransactionRegister(fromBlock *int64, handler func(int, []types.Log), transactionAddr common.Address, transactionId *big.Int) (string, error) {
	return _Transactionreg.contract.WatchLogs(fromBlock, handler, "TransactionRegister", transactionAddr, transactionId)
}

func (_Transactionreg *TransactionregFilterer) WatchAllTransactionRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transactionreg.contract.WatchLogs(fromBlock, handler, "TransactionRegister")
}

// ParseTransactionRegister is a log parse operation binding the contract event 0x4a4aff16c381ad8929f8b834c8723c495b0ab0fc2de27b0e3786ef201496a33d.
//
// Solidity: event TransactionRegister(address indexed transactionAddr, uint256 indexed transactionId)
func (_Transactionreg *TransactionregFilterer) ParseTransactionRegister(log types.Log) (*TransactionregTransactionRegister, error) {
	event := new(TransactionregTransactionRegister)
	if err := _Transactionreg.contract.UnpackLog(event, "TransactionRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransactionRegister is a free log subscription operation binding the contract event 0x4a4aff16c381ad8929f8b834c8723c495b0ab0fc2de27b0e3786ef201496a33d.
//
// Solidity: event TransactionRegister(address indexed transactionAddr, uint256 indexed transactionId)
func (_Transactionreg *TransactionregSession) WatchTransactionRegister(fromBlock *int64, handler func(int, []types.Log), transactionAddr common.Address, transactionId *big.Int) (string, error) {
	return _Transactionreg.Contract.WatchTransactionRegister(fromBlock, handler, transactionAddr, transactionId)
}

func (_Transactionreg *TransactionregSession) WatchAllTransactionRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transactionreg.Contract.WatchAllTransactionRegister(fromBlock, handler)
}

// ParseTransactionRegister is a log parse operation binding the contract event 0x4a4aff16c381ad8929f8b834c8723c495b0ab0fc2de27b0e3786ef201496a33d.
//
// Solidity: event TransactionRegister(address indexed transactionAddr, uint256 indexed transactionId)
func (_Transactionreg *TransactionregSession) ParseTransactionRegister(log types.Log) (*TransactionregTransactionRegister, error) {
	return _Transactionreg.Contract.ParseTransactionRegister(log)
}
