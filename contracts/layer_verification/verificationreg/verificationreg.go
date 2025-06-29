// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verificationreg

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

// VerificationregABI is the input ABI used to generate the binding from.
const VerificationregABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verificationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"verificationId\",\"type\":\"uint256\"}],\"name\":\"VerificationRegister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verificationAddress\",\"type\":\"address\"}],\"name\":\"getVerificationId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verificationAddress\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VerificationregBin is the compiled bytecode used for deploying new contracts.
var VerificationregBin = "0x608060405234801561001057600080fd5b50600160009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610753806100846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632801617e146100465780639507d39a14610062578063d90692f014610093575b600080fd5b610060600480360381019061005b9190610479565b6100c4565b005b61007c600480360381019061007791906104dc565b610248565b60405161008a9291906105b1565b60405180910390f35b6100ad60048036038101906100a89190610479565b610331565b6040516100bb9291906105f0565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461014557608360ff167fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3160405161013c90610646565b60405180910390a25b6001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600180805490506101bb9190610695565b9050806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167f2cfa3c239b72ce7bc95cffab0b41d30d465b2d7dacf1a17eacd2a3264a75591060405160405180910390a35050565b6000606060008314156102795760006040518060600160405280602581526020016106f9602591399150915061032c565b6001808054905061028a9190610695565b8311156102d25760006040518060400160405280600781526020017f756e666f756e64000000000000000000000000000000000000000000000000008152509150915061032c565b6000600184815481106102e8576102e76106c9565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806040518060200160405280600081525092509250505b915091565b6000606060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156103bd5760006040518060400160405280600781526020017f756e666f756e640000000000000000000000000000000000000000000000000081525091509150610411565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460405180602001604052806000815250915091505b915091565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104468261041b565b9050919050565b6104568161043b565b811461046157600080fd5b50565b6000813590506104738161044d565b92915050565b60006020828403121561048f5761048e610416565b5b600061049d84828501610464565b91505092915050565b6000819050919050565b6104b9816104a6565b81146104c457600080fd5b50565b6000813590506104d6816104b0565b92915050565b6000602082840312156104f2576104f1610416565b5b6000610500848285016104c7565b91505092915050565b6105128161043b565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610552578082015181840152602081019050610537565b83811115610561576000848401525b50505050565b6000601f19601f8301169050919050565b600061058382610518565b61058d8185610523565b935061059d818560208601610534565b6105a681610567565b840191505092915050565b60006040820190506105c66000830185610509565b81810360208301526105d88184610578565b90509392505050565b6105ea816104a6565b82525050565b600060408201905061060560008301856105e1565b81810360208301526106178184610578565b90509392505050565b50565b6000610630600083610523565b915061063b82610620565b600082019050919050565b6000602082019050818103600083015261065f81610623565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106a0826104a6565b91506106ab836104a6565b9250828210156106be576106bd610666565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe766572696669636174696f6e4964206d7573742062652067726561746572207468616e2030a2646970667358221220c1987538532a050ac873a2cd15395278ac0ba6ff40f78f47b692eef5669d392c64736f6c634300080b0033"
var VerificationregSMBin = "0x"

// DeployVerificationreg deploys a new contract, binding an instance of Verificationreg to it.
func DeployVerificationreg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Verificationreg, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationregABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(VerificationregSMBin)
	} else {
		bytecode = common.FromHex(VerificationregBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, VerificationregABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Verificationreg{VerificationregCaller: VerificationregCaller{contract: contract}, VerificationregTransactor: VerificationregTransactor{contract: contract}, VerificationregFilterer: VerificationregFilterer{contract: contract}}, nil
}

func AsyncDeployVerificationreg(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationregABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(VerificationregSMBin)
	} else {
		bytecode = common.FromHex(VerificationregBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, VerificationregABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Verificationreg is an auto generated Go binding around a Solidity contract.
type Verificationreg struct {
	VerificationregCaller     // Read-only binding to the contract
	VerificationregTransactor // Write-only binding to the contract
	VerificationregFilterer   // Log filterer for contract events
}

// VerificationregCaller is an auto generated read-only Go binding around a Solidity contract.
type VerificationregCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationregTransactor is an auto generated write-only Go binding around a Solidity contract.
type VerificationregTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationregFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type VerificationregFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationregSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type VerificationregSession struct {
	Contract     *Verificationreg  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerificationregCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type VerificationregCallerSession struct {
	Contract *VerificationregCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VerificationregTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type VerificationregTransactorSession struct {
	Contract     *VerificationregTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VerificationregRaw is an auto generated low-level Go binding around a Solidity contract.
type VerificationregRaw struct {
	Contract *Verificationreg // Generic contract binding to access the raw methods on
}

// VerificationregCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type VerificationregCallerRaw struct {
	Contract *VerificationregCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationregTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type VerificationregTransactorRaw struct {
	Contract *VerificationregTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationreg creates a new instance of Verificationreg, bound to a specific deployed contract.
func NewVerificationreg(address common.Address, backend bind.ContractBackend) (*Verificationreg, error) {
	contract, err := bindVerificationreg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verificationreg{VerificationregCaller: VerificationregCaller{contract: contract}, VerificationregTransactor: VerificationregTransactor{contract: contract}, VerificationregFilterer: VerificationregFilterer{contract: contract}}, nil
}

// NewVerificationregCaller creates a new read-only instance of Verificationreg, bound to a specific deployed contract.
func NewVerificationregCaller(address common.Address, caller bind.ContractCaller) (*VerificationregCaller, error) {
	contract, err := bindVerificationreg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationregCaller{contract: contract}, nil
}

// NewVerificationregTransactor creates a new write-only instance of Verificationreg, bound to a specific deployed contract.
func NewVerificationregTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationregTransactor, error) {
	contract, err := bindVerificationreg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationregTransactor{contract: contract}, nil
}

// NewVerificationregFilterer creates a new log filterer instance of Verificationreg, bound to a specific deployed contract.
func NewVerificationregFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationregFilterer, error) {
	contract, err := bindVerificationreg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationregFilterer{contract: contract}, nil
}

// bindVerificationreg binds a generic wrapper to an already deployed contract.
func bindVerificationreg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationregABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verificationreg *VerificationregRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verificationreg.Contract.VerificationregCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verificationreg *VerificationregRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.VerificationregTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verificationreg *VerificationregRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.VerificationregTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verificationreg *VerificationregCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verificationreg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verificationreg *VerificationregTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verificationreg *VerificationregTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 verificationId) constant returns(address, string)
func (_Verificationreg *VerificationregCaller) Get(opts *bind.CallOpts, verificationId *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Verificationreg.contract.Call(opts, out, "get", verificationId)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 verificationId) constant returns(address, string)
func (_Verificationreg *VerificationregSession) Get(verificationId *big.Int) (common.Address, string, error) {
	return _Verificationreg.Contract.Get(&_Verificationreg.CallOpts, verificationId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 verificationId) constant returns(address, string)
func (_Verificationreg *VerificationregCallerSession) Get(verificationId *big.Int) (common.Address, string, error) {
	return _Verificationreg.Contract.Get(&_Verificationreg.CallOpts, verificationId)
}

// GetVerificationId is a free data retrieval call binding the contract method 0xd90692f0.
//
// Solidity: function getVerificationId(address verificationAddress) constant returns(uint256, string)
func (_Verificationreg *VerificationregCaller) GetVerificationId(opts *bind.CallOpts, verificationAddress common.Address) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Verificationreg.contract.Call(opts, out, "getVerificationId", verificationAddress)
	return *ret0, *ret1, err
}

// GetVerificationId is a free data retrieval call binding the contract method 0xd90692f0.
//
// Solidity: function getVerificationId(address verificationAddress) constant returns(uint256, string)
func (_Verificationreg *VerificationregSession) GetVerificationId(verificationAddress common.Address) (*big.Int, string, error) {
	return _Verificationreg.Contract.GetVerificationId(&_Verificationreg.CallOpts, verificationAddress)
}

// GetVerificationId is a free data retrieval call binding the contract method 0xd90692f0.
//
// Solidity: function getVerificationId(address verificationAddress) constant returns(uint256, string)
func (_Verificationreg *VerificationregCallerSession) GetVerificationId(verificationAddress common.Address) (*big.Int, string, error) {
	return _Verificationreg.Contract.GetVerificationId(&_Verificationreg.CallOpts, verificationAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address verificationAddress) returns()
func (_Verificationreg *VerificationregTransactor) Set(opts *bind.TransactOpts, verificationAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Verificationreg.contract.TransactWithResult(opts, out, "set", verificationAddress)
	return transaction, receipt, err
}

func (_Verificationreg *VerificationregTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, verificationAddress common.Address) (*types.Transaction, error) {
	return _Verificationreg.contract.AsyncTransact(opts, handler, "set", verificationAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address verificationAddress) returns()
func (_Verificationreg *VerificationregSession) Set(verificationAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.Set(&_Verificationreg.TransactOpts, verificationAddress)
}

func (_Verificationreg *VerificationregSession) AsyncSet(handler func(*types.Receipt, error), verificationAddress common.Address) (*types.Transaction, error) {
	return _Verificationreg.Contract.AsyncSet(handler, &_Verificationreg.TransactOpts, verificationAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address verificationAddress) returns()
func (_Verificationreg *VerificationregTransactorSession) Set(verificationAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Verificationreg.Contract.Set(&_Verificationreg.TransactOpts, verificationAddress)
}

func (_Verificationreg *VerificationregTransactorSession) AsyncSet(handler func(*types.Receipt, error), verificationAddress common.Address) (*types.Transaction, error) {
	return _Verificationreg.Contract.AsyncSet(handler, &_Verificationreg.TransactOpts, verificationAddress)
}

// VerificationregVerificationRegister represents a VerificationRegister event raised by the Verificationreg contract.
type VerificationregVerificationRegister struct {
	VerificationAddress common.Address
	VerificationId      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// WatchVerificationRegister is a free log subscription operation binding the contract event 0x2cfa3c239b72ce7bc95cffab0b41d30d465b2d7dacf1a17eacd2a3264a755910.
//
// Solidity: event VerificationRegister(address indexed verificationAddress, uint256 indexed verificationId)
func (_Verificationreg *VerificationregFilterer) WatchVerificationRegister(fromBlock *int64, handler func(int, []types.Log), verificationAddress common.Address, verificationId *big.Int) (string, error) {
	return _Verificationreg.contract.WatchLogs(fromBlock, handler, "VerificationRegister", verificationAddress, verificationId)
}

func (_Verificationreg *VerificationregFilterer) WatchAllVerificationRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Verificationreg.contract.WatchLogs(fromBlock, handler, "VerificationRegister")
}

// ParseVerificationRegister is a log parse operation binding the contract event 0x2cfa3c239b72ce7bc95cffab0b41d30d465b2d7dacf1a17eacd2a3264a755910.
//
// Solidity: event VerificationRegister(address indexed verificationAddress, uint256 indexed verificationId)
func (_Verificationreg *VerificationregFilterer) ParseVerificationRegister(log types.Log) (*VerificationregVerificationRegister, error) {
	event := new(VerificationregVerificationRegister)
	if err := _Verificationreg.contract.UnpackLog(event, "VerificationRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchVerificationRegister is a free log subscription operation binding the contract event 0x2cfa3c239b72ce7bc95cffab0b41d30d465b2d7dacf1a17eacd2a3264a755910.
//
// Solidity: event VerificationRegister(address indexed verificationAddress, uint256 indexed verificationId)
func (_Verificationreg *VerificationregSession) WatchVerificationRegister(fromBlock *int64, handler func(int, []types.Log), verificationAddress common.Address, verificationId *big.Int) (string, error) {
	return _Verificationreg.Contract.WatchVerificationRegister(fromBlock, handler, verificationAddress, verificationId)
}

func (_Verificationreg *VerificationregSession) WatchAllVerificationRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Verificationreg.Contract.WatchAllVerificationRegister(fromBlock, handler)
}

// ParseVerificationRegister is a log parse operation binding the contract event 0x2cfa3c239b72ce7bc95cffab0b41d30d465b2d7dacf1a17eacd2a3264a755910.
//
// Solidity: event VerificationRegister(address indexed verificationAddress, uint256 indexed verificationId)
func (_Verificationreg *VerificationregSession) ParseVerificationRegister(log types.Log) (*VerificationregVerificationRegister, error) {
	return _Verificationreg.Contract.ParseVerificationRegister(log)
}
