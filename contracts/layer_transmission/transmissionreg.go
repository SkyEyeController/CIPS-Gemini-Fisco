// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transmissionreg

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

// TransmissionregABI is the input ABI used to generate the binding from.
const TransmissionregABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmissionAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transmissionId\",\"type\":\"uint256\"}],\"name\":\"TransmissionRegister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transmissionId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmissionAddress\",\"type\":\"address\"}],\"name\":\"getTransmissionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmissionAddress\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransmissionregBin is the compiled bytecode used for deploying new contracts.
var TransmissionregBin = "0x608060405234801561001057600080fd5b50600160009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610758806100846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806322b465da146100465780632801617e146100775780639507d39a14610093575b600080fd5b610060600480360381019061005b919061047e565b6100c4565b60405161006e92919061055d565b60405180910390f35b610091600480360381019061008c919061047e565b6101a9565b005b6100ad60048036038101906100a891906105b9565b610332565b6040516100bb9291906105f5565b60405180910390f35b6000606060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156101505760006040518060400160405280600781526020017f756e666f756e6400000000000000000000000000000000000000000000000000815250915091506101a4565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460405180602001604052806000815250915091505b915091565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541461022e57607960ff167fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc316040516102219061064b565b60405180910390a261032f565b6001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600180805490506102a4919061069a565b9050806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167fe30cda668b9625a246c64a56e1c97468bfead7190a2cd65206a468611665bcb760405160405180910390a3505b50565b6000606060008314156103635760006040518060600160405280602581526020016106fe6025913991509150610416565b60018080549050610374919061069a565b8311156103bc5760006040518060400160405280600781526020017f756e666f756e640000000000000000000000000000000000000000000000000081525091509150610416565b6000600184815481106103d2576103d16106ce565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806040518060200160405280600081525092509250505b915091565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061044b82610420565b9050919050565b61045b81610440565b811461046657600080fd5b50565b60008135905061047881610452565b92915050565b6000602082840312156104945761049361041b565b5b60006104a284828501610469565b91505092915050565b6000819050919050565b6104be816104ab565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104fe5780820151818401526020810190506104e3565b8381111561050d576000848401525b50505050565b6000601f19601f8301169050919050565b600061052f826104c4565b61053981856104cf565b93506105498185602086016104e0565b61055281610513565b840191505092915050565b600060408201905061057260008301856104b5565b81810360208301526105848184610524565b90509392505050565b610596816104ab565b81146105a157600080fd5b50565b6000813590506105b38161058d565b92915050565b6000602082840312156105cf576105ce61041b565b5b60006105dd848285016105a4565b91505092915050565b6105ef81610440565b82525050565b600060408201905061060a60008301856105e6565b818103602083015261061c8184610524565b90509392505050565b50565b60006106356000836104cf565b915061064082610625565b600082019050919050565b6000602082019050818103600083015261066481610628565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106a5826104ab565b91506106b0836104ab565b9250828210156106c3576106c261066b565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe7472616e736d697373696f6e4964206d7573742062652067726561746572207468616e2030a2646970667358221220e3deace597551f0ce79a3f6fdda5cc209d3a0fd942fefe036f7f3089f884e29064736f6c634300080b0033"
var TransmissionregSMBin = "0x"

// DeployTransmissionreg deploys a new contract, binding an instance of Transmissionreg to it.
func DeployTransmissionreg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transmissionreg, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionregABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransmissionregSMBin)
	} else {
		bytecode = common.FromHex(TransmissionregBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TransmissionregABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transmissionreg{TransmissionregCaller: TransmissionregCaller{contract: contract}, TransmissionregTransactor: TransmissionregTransactor{contract: contract}, TransmissionregFilterer: TransmissionregFilterer{contract: contract}}, nil
}

func AsyncDeployTransmissionreg(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionregABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransmissionregSMBin)
	} else {
		bytecode = common.FromHex(TransmissionregBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TransmissionregABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transmissionreg is an auto generated Go binding around a Solidity contract.
type Transmissionreg struct {
	TransmissionregCaller     // Read-only binding to the contract
	TransmissionregTransactor // Write-only binding to the contract
	TransmissionregFilterer   // Log filterer for contract events
}

// TransmissionregCaller is an auto generated read-only Go binding around a Solidity contract.
type TransmissionregCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionregTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransmissionregTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionregFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransmissionregFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionregSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransmissionregSession struct {
	Contract     *Transmissionreg  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransmissionregCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransmissionregCallerSession struct {
	Contract *TransmissionregCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TransmissionregTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransmissionregTransactorSession struct {
	Contract     *TransmissionregTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TransmissionregRaw is an auto generated low-level Go binding around a Solidity contract.
type TransmissionregRaw struct {
	Contract *Transmissionreg // Generic contract binding to access the raw methods on
}

// TransmissionregCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransmissionregCallerRaw struct {
	Contract *TransmissionregCaller // Generic read-only contract binding to access the raw methods on
}

// TransmissionregTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransmissionregTransactorRaw struct {
	Contract *TransmissionregTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransmissionreg creates a new instance of Transmissionreg, bound to a specific deployed contract.
func NewTransmissionreg(address common.Address, backend bind.ContractBackend) (*Transmissionreg, error) {
	contract, err := bindTransmissionreg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transmissionreg{TransmissionregCaller: TransmissionregCaller{contract: contract}, TransmissionregTransactor: TransmissionregTransactor{contract: contract}, TransmissionregFilterer: TransmissionregFilterer{contract: contract}}, nil
}

// NewTransmissionregCaller creates a new read-only instance of Transmissionreg, bound to a specific deployed contract.
func NewTransmissionregCaller(address common.Address, caller bind.ContractCaller) (*TransmissionregCaller, error) {
	contract, err := bindTransmissionreg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransmissionregCaller{contract: contract}, nil
}

// NewTransmissionregTransactor creates a new write-only instance of Transmissionreg, bound to a specific deployed contract.
func NewTransmissionregTransactor(address common.Address, transactor bind.ContractTransactor) (*TransmissionregTransactor, error) {
	contract, err := bindTransmissionreg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransmissionregTransactor{contract: contract}, nil
}

// NewTransmissionregFilterer creates a new log filterer instance of Transmissionreg, bound to a specific deployed contract.
func NewTransmissionregFilterer(address common.Address, filterer bind.ContractFilterer) (*TransmissionregFilterer, error) {
	contract, err := bindTransmissionreg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransmissionregFilterer{contract: contract}, nil
}

// bindTransmissionreg binds a generic wrapper to an already deployed contract.
func bindTransmissionreg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionregABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transmissionreg *TransmissionregRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transmissionreg.Contract.TransmissionregCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transmissionreg *TransmissionregRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.TransmissionregTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transmissionreg *TransmissionregRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.TransmissionregTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transmissionreg *TransmissionregCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transmissionreg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transmissionreg *TransmissionregTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transmissionreg *TransmissionregTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transmissionId) constant returns(address, string)
func (_Transmissionreg *TransmissionregCaller) Get(opts *bind.CallOpts, transmissionId *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transmissionreg.contract.Call(opts, out, "get", transmissionId)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transmissionId) constant returns(address, string)
func (_Transmissionreg *TransmissionregSession) Get(transmissionId *big.Int) (common.Address, string, error) {
	return _Transmissionreg.Contract.Get(&_Transmissionreg.CallOpts, transmissionId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 transmissionId) constant returns(address, string)
func (_Transmissionreg *TransmissionregCallerSession) Get(transmissionId *big.Int) (common.Address, string, error) {
	return _Transmissionreg.Contract.Get(&_Transmissionreg.CallOpts, transmissionId)
}

// GetTransmissionId is a free data retrieval call binding the contract method 0x22b465da.
//
// Solidity: function getTransmissionId(address transmissionAddress) constant returns(uint256, string)
func (_Transmissionreg *TransmissionregCaller) GetTransmissionId(opts *bind.CallOpts, transmissionAddress common.Address) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transmissionreg.contract.Call(opts, out, "getTransmissionId", transmissionAddress)
	return *ret0, *ret1, err
}

// GetTransmissionId is a free data retrieval call binding the contract method 0x22b465da.
//
// Solidity: function getTransmissionId(address transmissionAddress) constant returns(uint256, string)
func (_Transmissionreg *TransmissionregSession) GetTransmissionId(transmissionAddress common.Address) (*big.Int, string, error) {
	return _Transmissionreg.Contract.GetTransmissionId(&_Transmissionreg.CallOpts, transmissionAddress)
}

// GetTransmissionId is a free data retrieval call binding the contract method 0x22b465da.
//
// Solidity: function getTransmissionId(address transmissionAddress) constant returns(uint256, string)
func (_Transmissionreg *TransmissionregCallerSession) GetTransmissionId(transmissionAddress common.Address) (*big.Int, string, error) {
	return _Transmissionreg.Contract.GetTransmissionId(&_Transmissionreg.CallOpts, transmissionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transmissionAddress) returns()
func (_Transmissionreg *TransmissionregTransactor) Set(opts *bind.TransactOpts, transmissionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transmissionreg.contract.TransactWithResult(opts, out, "set", transmissionAddress)
	return transaction, receipt, err
}

func (_Transmissionreg *TransmissionregTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, transmissionAddress common.Address) (*types.Transaction, error) {
	return _Transmissionreg.contract.AsyncTransact(opts, handler, "set", transmissionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transmissionAddress) returns()
func (_Transmissionreg *TransmissionregSession) Set(transmissionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.Set(&_Transmissionreg.TransactOpts, transmissionAddress)
}

func (_Transmissionreg *TransmissionregSession) AsyncSet(handler func(*types.Receipt, error), transmissionAddress common.Address) (*types.Transaction, error) {
	return _Transmissionreg.Contract.AsyncSet(handler, &_Transmissionreg.TransactOpts, transmissionAddress)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address transmissionAddress) returns()
func (_Transmissionreg *TransmissionregTransactorSession) Set(transmissionAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionreg.Contract.Set(&_Transmissionreg.TransactOpts, transmissionAddress)
}

func (_Transmissionreg *TransmissionregTransactorSession) AsyncSet(handler func(*types.Receipt, error), transmissionAddress common.Address) (*types.Transaction, error) {
	return _Transmissionreg.Contract.AsyncSet(handler, &_Transmissionreg.TransactOpts, transmissionAddress)
}

// TransmissionregTransmissionRegister represents a TransmissionRegister event raised by the Transmissionreg contract.
type TransmissionregTransmissionRegister struct {
	TransmissionAddress common.Address
	TransmissionId      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// WatchTransmissionRegister is a free log subscription operation binding the contract event 0xe30cda668b9625a246c64a56e1c97468bfead7190a2cd65206a468611665bcb7.
//
// Solidity: event TransmissionRegister(address indexed transmissionAddress, uint256 indexed transmissionId)
func (_Transmissionreg *TransmissionregFilterer) WatchTransmissionRegister(fromBlock *int64, handler func(int, []types.Log), transmissionAddress common.Address, transmissionId *big.Int) (string, error) {
	return _Transmissionreg.contract.WatchLogs(fromBlock, handler, "TransmissionRegister", transmissionAddress, transmissionId)
}

func (_Transmissionreg *TransmissionregFilterer) WatchAllTransmissionRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transmissionreg.contract.WatchLogs(fromBlock, handler, "TransmissionRegister")
}

// ParseTransmissionRegister is a log parse operation binding the contract event 0xe30cda668b9625a246c64a56e1c97468bfead7190a2cd65206a468611665bcb7.
//
// Solidity: event TransmissionRegister(address indexed transmissionAddress, uint256 indexed transmissionId)
func (_Transmissionreg *TransmissionregFilterer) ParseTransmissionRegister(log types.Log) (*TransmissionregTransmissionRegister, error) {
	event := new(TransmissionregTransmissionRegister)
	if err := _Transmissionreg.contract.UnpackLog(event, "TransmissionRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransmissionRegister is a free log subscription operation binding the contract event 0xe30cda668b9625a246c64a56e1c97468bfead7190a2cd65206a468611665bcb7.
//
// Solidity: event TransmissionRegister(address indexed transmissionAddress, uint256 indexed transmissionId)
func (_Transmissionreg *TransmissionregSession) WatchTransmissionRegister(fromBlock *int64, handler func(int, []types.Log), transmissionAddress common.Address, transmissionId *big.Int) (string, error) {
	return _Transmissionreg.Contract.WatchTransmissionRegister(fromBlock, handler, transmissionAddress, transmissionId)
}

func (_Transmissionreg *TransmissionregSession) WatchAllTransmissionRegister(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Transmissionreg.Contract.WatchAllTransmissionRegister(fromBlock, handler)
}

// ParseTransmissionRegister is a log parse operation binding the contract event 0xe30cda668b9625a246c64a56e1c97468bfead7190a2cd65206a468611665bcb7.
//
// Solidity: event TransmissionRegister(address indexed transmissionAddress, uint256 indexed transmissionId)
func (_Transmissionreg *TransmissionregSession) ParseTransmissionRegister(log types.Log) (*TransmissionregTransmissionRegister, error) {
	return _Transmissionreg.Contract.ParseTransmissionRegister(log)
}
