// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package appreg

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

// AppregABI is the input ABI used to generate the binding from.
const AppregABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"appAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"appId\",\"type\":\"uint256\"}],\"name\":\"AppRegister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"}],\"name\":\"getAppId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AppregBin is the compiled bytecode used for deploying new contracts.
var AppregBin = "0x608060405234801561001057600080fd5b50600160009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610744806100846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632801617e146100465780632ba9b839146100625780639507d39a14610093575b600080fd5b610060600480360381019061005b919061048f565b6100c4565b005b61007c6004803603810190610077919061048f565b61024d565b60405161008a92919061056e565b60405180910390f35b6100ad60048036038101906100a891906105ca565b610332565b6040516100bb929190610606565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461014957606560ff167fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3160405161013c9061065c565b60405180910390a261024a565b6001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600180805490506101bf91906106ab565b9050806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167fed1405e8771772ffb0c1eb9504f162ad218f9a16e9e4f83a4d1f3142c89e5d3a60405160405180910390a3505b50565b6000606060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156102d95760006040518060400160405280600781526020017f756e666f756e64000000000000000000000000000000000000000000000000008152509150915061032d565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460405180602001604052806000815250915091505b915091565b6000606060008314156103805760006040518060400160405280601c81526020017f6170704964206d7573742062652067726561746572207468616e20300000000081525091509150610427565b60018054905083106103cd5760006040518060400160405280600781526020017f756e666f756e640000000000000000000000000000000000000000000000000081525091509150610427565b6000600184815481106103e3576103e26106df565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806040518060200160405280600081525092509250505b915091565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061045c82610431565b9050919050565b61046c81610451565b811461047757600080fd5b50565b60008135905061048981610463565b92915050565b6000602082840312156104a5576104a461042c565b5b60006104b38482850161047a565b91505092915050565b6000819050919050565b6104cf816104bc565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561050f5780820151818401526020810190506104f4565b8381111561051e576000848401525b50505050565b6000601f19601f8301169050919050565b6000610540826104d5565b61054a81856104e0565b935061055a8185602086016104f1565b61056381610524565b840191505092915050565b600060408201905061058360008301856104c6565b81810360208301526105958184610535565b90509392505050565b6105a7816104bc565b81146105b257600080fd5b50565b6000813590506105c48161059e565b92915050565b6000602082840312156105e0576105df61042c565b5b60006105ee848285016105b5565b91505092915050565b61060081610451565b82525050565b600060408201905061061b60008301856105f7565b818103602083015261062d8184610535565b90509392505050565b50565b60006106466000836104e0565b915061065182610636565b600082019050919050565b6000602082019050818103600083015261067581610639565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106b6826104bc565b91506106c1836104bc565b9250828210156106d4576106d361067c565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220eea1ae12f28e1faf07c5bc05247ed40f4927c1e1a456341d4fed9f4582bb01b164736f6c634300080b0033"
var AppregSMBin = "0x"

// DeployAppreg deploys a new contract, binding an instance of Appreg to it.
func DeployAppreg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Appreg, error) {
	parsed, err := abi.JSON(strings.NewReader(AppregABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(AppregSMBin)
	} else {
		bytecode = common.FromHex(AppregBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, AppregABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Appreg{AppregCaller: AppregCaller{contract: contract}, AppregTransactor: AppregTransactor{contract: contract}, AppregFilterer: AppregFilterer{contract: contract}}, nil
}

func AsyncDeployAppreg(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(AppregABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(AppregSMBin)
	} else {
		bytecode = common.FromHex(AppregBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, AppregABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Appreg is an auto generated Go binding around a Solidity contract.
type Appreg struct {
	AppregCaller     // Read-only binding to the contract
	AppregTransactor // Write-only binding to the contract
	AppregFilterer   // Log filterer for contract events
}

// AppregCaller is an auto generated read-only Go binding around a Solidity contract.
type AppregCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppregTransactor is an auto generated write-only Go binding around a Solidity contract.
type AppregTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppregFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AppregFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppregSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AppregSession struct {
	Contract     *Appreg           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppregCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AppregCallerSession struct {
	Contract *AppregCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppregTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AppregTransactorSession struct {
	Contract     *AppregTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppregRaw is an auto generated low-level Go binding around a Solidity contract.
type AppregRaw struct {
	Contract *Appreg // Generic contract binding to access the raw methods on
}

// AppregCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AppregCallerRaw struct {
	Contract *AppregCaller // Generic read-only contract binding to access the raw methods on
}

// AppregTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AppregTransactorRaw struct {
	Contract *AppregTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppreg creates a new instance of Appreg, bound to a specific deployed contract.
func NewAppreg(address common.Address, backend bind.ContractBackend) (*Appreg, error) {
	contract, err := bindAppreg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Appreg{AppregCaller: AppregCaller{contract: contract}, AppregTransactor: AppregTransactor{contract: contract}, AppregFilterer: AppregFilterer{contract: contract}}, nil
}

// NewAppregCaller creates a new read-only instance of Appreg, bound to a specific deployed contract.
func NewAppregCaller(address common.Address, caller bind.ContractCaller) (*AppregCaller, error) {
	contract, err := bindAppreg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppregCaller{contract: contract}, nil
}

// NewAppregTransactor creates a new write-only instance of Appreg, bound to a specific deployed contract.
func NewAppregTransactor(address common.Address, transactor bind.ContractTransactor) (*AppregTransactor, error) {
	contract, err := bindAppreg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppregTransactor{contract: contract}, nil
}

// NewAppregFilterer creates a new log filterer instance of Appreg, bound to a specific deployed contract.
func NewAppregFilterer(address common.Address, filterer bind.ContractFilterer) (*AppregFilterer, error) {
	contract, err := bindAppreg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppregFilterer{contract: contract}, nil
}

// bindAppreg binds a generic wrapper to an already deployed contract.
func bindAppreg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppregABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Appreg *AppregRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Appreg.Contract.AppregCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Appreg *AppregRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.AppregTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Appreg *AppregRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.AppregTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Appreg *AppregCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Appreg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Appreg *AppregTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Appreg *AppregTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 appId) constant returns(address, string)
func (_Appreg *AppregCaller) Get(opts *bind.CallOpts, appId *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Appreg.contract.Call(opts, out, "get", appId)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 appId) constant returns(address, string)
func (_Appreg *AppregSession) Get(appId *big.Int) (common.Address, string, error) {
	return _Appreg.Contract.Get(&_Appreg.CallOpts, appId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 appId) constant returns(address, string)
func (_Appreg *AppregCallerSession) Get(appId *big.Int) (common.Address, string, error) {
	return _Appreg.Contract.Get(&_Appreg.CallOpts, appId)
}

// GetAppId is a free data retrieval call binding the contract method 0x2ba9b839.
//
// Solidity: function getAppId(address appAddress) constant returns(uint256, string)
func (_Appreg *AppregCaller) GetAppId(opts *bind.CallOpts, appAddress common.Address) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Appreg.contract.Call(opts, out, "getAppId", appAddress)
	return *ret0, *ret1, err
}

// GetAppId is a free data retrieval call binding the contract method 0x2ba9b839.
//
// Solidity: function getAppId(address appAddress) constant returns(uint256, string)
func (_Appreg *AppregSession) GetAppId(appAddress common.Address) (*big.Int, string, error) {
	return _Appreg.Contract.GetAppId(&_Appreg.CallOpts, appAddress)
}

// GetAppId is a free data retrieval call binding the contract method 0x2ba9b839.
//
// Solidity: function getAppId(address appAddress) constant returns(uint256, string)
func (_Appreg *AppregCallerSession) GetAppId(appAddress common.Address) (*big.Int, string, error) {
	return _Appreg.Contract.GetAppId(&_Appreg.CallOpts, appAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address appAddress) returns()
func (_Appreg *AppregTransactor) Set(opts *bind.TransactOpts, appAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Appreg.contract.TransactWithResult(opts, out, "set", appAddress)
	return transaction, receipt, err
}

func (_Appreg *AppregTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, appAddress common.Address) (*types.Transaction, error) {
	return _Appreg.contract.AsyncTransact(opts, handler, "set", appAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address appAddress) returns()
func (_Appreg *AppregSession) Set(appAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.Set(&_Appreg.TransactOpts, appAddress)
}

func (_Appreg *AppregSession) AsyncSet(handler func(*types.Receipt, error), appAddress common.Address) (*types.Transaction, error) {
	return _Appreg.Contract.AsyncSet(handler, &_Appreg.TransactOpts, appAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address appAddress) returns()
func (_Appreg *AppregTransactorSession) Set(appAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Appreg.Contract.Set(&_Appreg.TransactOpts, appAddress)
}

func (_Appreg *AppregTransactorSession) AsyncSet(handler func(*types.Receipt, error), appAddress common.Address) (*types.Transaction, error) {
	return _Appreg.Contract.AsyncSet(handler, &_Appreg.TransactOpts, appAddress)
}

// AppregAppRegister represents a AppRegister event raised by the Appreg contract.
type AppregAppRegister struct {
	AppAddr common.Address
	AppId   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchAppRegister is a free log subscription operation binding the contract event 0xed1405e8771772ffb0c1eb9504f162ad218f9a16e9e4f83a4d1f3142c89e5d3a.
//
// Solidity: event AppRegister(address indexed appAddr, uint256 indexed appId)
func (_Appreg *AppregFilterer) WatchAppRegister(fromBlock *int64, handler func(int, []types.Log), appAddr common.Address, appId *big.Int) (string, error) {
	return _Appreg.contract.WatchLogs(fromBlock, handler, "AppRegister", appAddr, appId)
}

func (_Appreg *AppregFilterer) WatchAllAppRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Appreg.contract.WatchLogs(fromBlock, handler, "AppRegister")
}

// ParseAppRegister is a log parse operation binding the contract event 0xed1405e8771772ffb0c1eb9504f162ad218f9a16e9e4f83a4d1f3142c89e5d3a.
//
// Solidity: event AppRegister(address indexed appAddr, uint256 indexed appId)
func (_Appreg *AppregFilterer) ParseAppRegister(log types.Log) (*AppregAppRegister, error) {
	event := new(AppregAppRegister)
	if err := _Appreg.contract.UnpackLog(event, "AppRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchAppRegister is a free log subscription operation binding the contract event 0xed1405e8771772ffb0c1eb9504f162ad218f9a16e9e4f83a4d1f3142c89e5d3a.
//
// Solidity: event AppRegister(address indexed appAddr, uint256 indexed appId)
func (_Appreg *AppregSession) WatchAppRegister(fromBlock *int64, handler func(int, []types.Log), appAddr common.Address, appId *big.Int) (string, error) {
	return _Appreg.Contract.WatchAppRegister(fromBlock, handler, appAddr, appId)
}

func (_Appreg *AppregSession) WatchAllAppRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Appreg.Contract.WatchAllAppRegister(fromBlock, handler)
}

// ParseAppRegister is a log parse operation binding the contract event 0xed1405e8771772ffb0c1eb9504f162ad218f9a16e9e4f83a4d1f3142c89e5d3a.
//
// Solidity: event AppRegister(address indexed appAddr, uint256 indexed appId)
func (_Appreg *AppregSession) ParseAppRegister(log types.Log) (*AppregAppRegister, error) {
	return _Appreg.Contract.ParseAppRegister(log)
}
