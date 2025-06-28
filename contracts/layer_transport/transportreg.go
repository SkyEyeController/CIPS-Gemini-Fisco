// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transportreg

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

// TransportregABI is the input ABI used to generate the binding from.
const TransportregABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transportAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transportId\",\"type\":\"uint256\"}],\"name\":\"TransportRegister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transportId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transportAddress\",\"type\":\"address\"}],\"name\":\"getTransportId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transportAddress\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransportregBin is the compiled bytecode used for deploying new contracts.
var TransportregBin = "0x608060405234801561001057600080fd5b50600160009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610750806100846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632801617e146100465780639507d39a14610062578063ffe53d8414610093575b600080fd5b610060600480360381019061005b9190610479565b6100c4565b005b61007c600480360381019061007791906104dc565b610248565b60405161008a9291906105b1565b60405180910390f35b6100ad60048036038101906100a89190610479565b610331565b6040516100bb9291906105f0565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461014557608d60ff167fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3160405161013c90610646565b60405180910390a25b6001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600180805490506101bb9190610695565b9050806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167ff66b366deb7e49c4c4edd9354a338d82014ff4e26f0994aec60632641378700560405160405180910390a35050565b6000606060008314156102795760006040518060600160405280602281526020016106f9602291399150915061032c565b6001808054905061028a9190610695565b8311156102d25760006040518060400160405280600781526020017f756e666f756e64000000000000000000000000000000000000000000000000008152509150915061032c565b6000600184815481106102e8576102e76106c9565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806040518060200160405280600081525092509250505b915091565b6000606060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156103bd5760006040518060400160405280600781526020017f756e666f756e640000000000000000000000000000000000000000000000000081525091509150610411565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460405180602001604052806000815250915091505b915091565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104468261041b565b9050919050565b6104568161043b565b811461046157600080fd5b50565b6000813590506104738161044d565b92915050565b60006020828403121561048f5761048e610416565b5b600061049d84828501610464565b91505092915050565b6000819050919050565b6104b9816104a6565b81146104c457600080fd5b50565b6000813590506104d6816104b0565b92915050565b6000602082840312156104f2576104f1610416565b5b6000610500848285016104c7565b91505092915050565b6105128161043b565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610552578082015181840152602081019050610537565b83811115610561576000848401525b50505050565b6000601f19601f8301169050919050565b600061058382610518565b61058d8185610523565b935061059d818560208601610534565b6105a681610567565b840191505092915050565b60006040820190506105c66000830185610509565b81810360208301526105d88184610578565b90509392505050565b6105ea816104a6565b82525050565b600060408201905061060560008301856105e1565b81810360208301526106178184610578565b90509392505050565b50565b6000610630600083610523565b915061063b82610620565b600082019050919050565b6000602082019050818103600083015261065f81610623565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106a0826104a6565b91506106ab836104a6565b9250828210156106be576106bd610666565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe7472616e73706f72744964206d7573742062652067726561746572207468616e2030a2646970667358221220c817055acdad08606645ffa9e823a9e8f37b6f4bc763396a33497a7a5502118164736f6c634300080b0033"
var TransportregSMBin = "0x"

// DeployTransportreg deploys a new contract, binding an instance of Transportreg to it.
func DeployTransportreg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transportreg, error) {
	parsed, err := abi.JSON(strings.NewReader(TransportregABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransportregSMBin)
	} else {
		bytecode = common.FromHex(TransportregBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TransportregABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transportreg{TransportregCaller: TransportregCaller{contract: contract}, TransportregTransactor: TransportregTransactor{contract: contract}, TransportregFilterer: TransportregFilterer{contract: contract}}, nil
}

func AsyncDeployTransportreg(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransportregABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransportregSMBin)
	} else {
		bytecode = common.FromHex(TransportregBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TransportregABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transportreg is an auto generated Go binding around a Solidity contract.
type Transportreg struct {
	TransportregCaller     // Read-only binding to the contract
	TransportregTransactor // Write-only binding to the contract
	TransportregFilterer   // Log filterer for contract events
}

// TransportregCaller is an auto generated read-only Go binding around a Solidity contract.
type TransportregCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportregTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransportregTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportregFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransportregFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportregSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransportregSession struct {
	Contract     *Transportreg     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransportregCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransportregCallerSession struct {
	Contract *TransportregCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransportregTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransportregTransactorSession struct {
	Contract     *TransportregTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransportregRaw is an auto generated low-level Go binding around a Solidity contract.
type TransportregRaw struct {
	Contract *Transportreg // Generic contract binding to access the raw methods on
}

// TransportregCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransportregCallerRaw struct {
	Contract *TransportregCaller // Generic read-only contract binding to access the raw methods on
}

// TransportregTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransportregTransactorRaw struct {
	Contract *TransportregTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransportreg creates a new instance of Transportreg, bound to a specific deployed contract.
func NewTransportreg(address common.Address, backend bind.ContractBackend) (*Transportreg, error) {
	contract, err := bindTransportreg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transportreg{TransportregCaller: TransportregCaller{contract: contract}, TransportregTransactor: TransportregTransactor{contract: contract}, TransportregFilterer: TransportregFilterer{contract: contract}}, nil
}

// NewTransportregCaller creates a new read-only instance of Transportreg, bound to a specific deployed contract.
func NewTransportregCaller(address common.Address, caller bind.ContractCaller) (*TransportregCaller, error) {
	contract, err := bindTransportreg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransportregCaller{contract: contract}, nil
}

// NewTransportregTransactor creates a new write-only instance of Transportreg, bound to a specific deployed contract.
func NewTransportregTransactor(address common.Address, transactor bind.ContractTransactor) (*TransportregTransactor, error) {
	contract, err := bindTransportreg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransportregTransactor{contract: contract}, nil
}

// NewTransportregFilterer creates a new log filterer instance of Transportreg, bound to a specific deployed contract.
func NewTransportregFilterer(address common.Address, filterer bind.ContractFilterer) (*TransportregFilterer, error) {
	contract, err := bindTransportreg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransportregFilterer{contract: contract}, nil
}

// bindTransportreg binds a generic wrapper to an already deployed contract.
func bindTransportreg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransportregABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transportreg *TransportregRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transportreg.Contract.TransportregCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transportreg *TransportregRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.TransportregTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transportreg *TransportregRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.TransportregTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transportreg *TransportregCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transportreg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transportreg *TransportregTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transportreg *TransportregTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transportId) constant returns(address, string)
func (_Transportreg *TransportregCaller) Get(opts *bind.CallOpts, transportId *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transportreg.contract.Call(opts, out, "get", transportId)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transportId) constant returns(address, string)
func (_Transportreg *TransportregSession) Get(transportId *big.Int) (common.Address, string, error) {
	return _Transportreg.Contract.Get(&_Transportreg.CallOpts, transportId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transportId) constant returns(address, string)
func (_Transportreg *TransportregCallerSession) Get(transportId *big.Int) (common.Address, string, error) {
	return _Transportreg.Contract.Get(&_Transportreg.CallOpts, transportId)
}

// GetTransportId is a free data retrieval call binding the contract method 0xffe53d84.
//
// Solidity: function getTransportId(address transportAddress) constant returns(uint256, string)
func (_Transportreg *TransportregCaller) GetTransportId(opts *bind.CallOpts, transportAddress common.Address) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transportreg.contract.Call(opts, out, "getTransportId", transportAddress)
	return *ret0, *ret1, err
}

// GetTransportId is a free data retrieval call binding the contract method 0xffe53d84.
//
// Solidity: function getTransportId(address transportAddress) constant returns(uint256, string)
func (_Transportreg *TransportregSession) GetTransportId(transportAddress common.Address) (*big.Int, string, error) {
	return _Transportreg.Contract.GetTransportId(&_Transportreg.CallOpts, transportAddress)
}

// GetTransportId is a free data retrieval call binding the contract method 0xffe53d84.
//
// Solidity: function getTransportId(address transportAddress) constant returns(uint256, string)
func (_Transportreg *TransportregCallerSession) GetTransportId(transportAddress common.Address) (*big.Int, string, error) {
	return _Transportreg.Contract.GetTransportId(&_Transportreg.CallOpts, transportAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transportAddress) returns()
func (_Transportreg *TransportregTransactor) Set(opts *bind.TransactOpts, transportAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transportreg.contract.TransactWithResult(opts, out, "set", transportAddress)
	return transaction, receipt, err
}

func (_Transportreg *TransportregTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, transportAddress common.Address) (*types.Transaction, error) {
	return _Transportreg.contract.AsyncTransact(opts, handler, "set", transportAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transportAddress) returns()
func (_Transportreg *TransportregSession) Set(transportAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.Set(&_Transportreg.TransactOpts, transportAddress)
}

func (_Transportreg *TransportregSession) AsyncSet(handler func(*types.Receipt, error), transportAddress common.Address) (*types.Transaction, error) {
	return _Transportreg.Contract.AsyncSet(handler, &_Transportreg.TransactOpts, transportAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transportAddress) returns()
func (_Transportreg *TransportregTransactorSession) Set(transportAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transportreg.Contract.Set(&_Transportreg.TransactOpts, transportAddress)
}

func (_Transportreg *TransportregTransactorSession) AsyncSet(handler func(*types.Receipt, error), transportAddress common.Address) (*types.Transaction, error) {
	return _Transportreg.Contract.AsyncSet(handler, &_Transportreg.TransactOpts, transportAddress)
}

// TransportregTransportRegister represents a TransportRegister event raised by the Transportreg contract.
type TransportregTransportRegister struct {
	TransportAddress common.Address
	TransportId      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// WatchTransportRegister is a free log subscription operation binding the contract event 0xf66b366deb7e49c4c4edd9354a338d82014ff4e26f0994aec606326413787005.
//
// Solidity: event TransportRegister(address indexed transportAddress, uint256 indexed transportId)
func (_Transportreg *TransportregFilterer) WatchTransportRegister(fromBlock *int64, handler func(int, []types.Log), transportAddress common.Address, transportId *big.Int) (string, error) {
	return _Transportreg.contract.WatchLogs(fromBlock, handler, "TransportRegister", transportAddress, transportId)
}

func (_Transportreg *TransportregFilterer) WatchAllTransportRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transportreg.contract.WatchLogs(fromBlock, handler, "TransportRegister")
}

// ParseTransportRegister is a log parse operation binding the contract event 0xf66b366deb7e49c4c4edd9354a338d82014ff4e26f0994aec606326413787005.
//
// Solidity: event TransportRegister(address indexed transportAddress, uint256 indexed transportId)
func (_Transportreg *TransportregFilterer) ParseTransportRegister(log types.Log) (*TransportregTransportRegister, error) {
	event := new(TransportregTransportRegister)
	if err := _Transportreg.contract.UnpackLog(event, "TransportRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransportRegister is a free log subscription operation binding the contract event 0xf66b366deb7e49c4c4edd9354a338d82014ff4e26f0994aec606326413787005.
//
// Solidity: event TransportRegister(address indexed transportAddress, uint256 indexed transportId)
func (_Transportreg *TransportregSession) WatchTransportRegister(fromBlock *int64, handler func(int, []types.Log), transportAddress common.Address, transportId *big.Int) (string, error) {
	return _Transportreg.Contract.WatchTransportRegister(fromBlock, handler, transportAddress, transportId)
}

func (_Transportreg *TransportregSession) WatchAllTransportRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transportreg.Contract.WatchAllTransportRegister(fromBlock, handler)
}

// ParseTransportRegister is a log parse operation binding the contract event 0xf66b366deb7e49c4c4edd9354a338d82014ff4e26f0994aec606326413787005.
//
// Solidity: event TransportRegister(address indexed transportAddress, uint256 indexed transportId)
func (_Transportreg *TransportregSession) ParseTransportRegister(log types.Log) (*TransportregTransportRegister, error) {
	return _Transportreg.Contract.ParseTransportRegister(log)
}
