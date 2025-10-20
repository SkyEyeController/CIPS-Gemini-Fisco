// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homomorphiccalculator

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

// HomomorphiccalculatorABI is the input ABI used to generate the binding from.
const HomomorphiccalculatorABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_g\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"}],\"name\":\"homomorphicAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalarK\",\"type\":\"uint256\"}],\"name\":\"homomorphicMultiply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"g\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nSquared\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HomomorphiccalculatorBin is the compiled bytecode used for deploying new contracts.
var HomomorphiccalculatorBin = "0x608060405234801561001057600080fd5b5060405161050a38038061050a8339818101604052810190610032919061009c565b816000800181905550806000600101819055508182610051919061010b565b6000600201819055505050610165565b600080fd5b6000819050919050565b61007981610066565b811461008457600080fd5b50565b60008151905061009681610070565b92915050565b600080604083850312156100b3576100b2610061565b5b60006100c185828601610087565b92505060206100d285828601610087565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061011682610066565b915061012183610066565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561015a576101596100dc565b5b828202905092915050565b610396806101746000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633943380c1461004657806359e18a0514610066578063baf38aae14610096575b600080fd5b61004e6100c6565b60405161005d939291906101b4565b60405180910390f35b610080600480360381019061007b919061021c565b6100de565b60405161008d919061025c565b60405180910390f35b6100b060048036038101906100ab919061021c565b610103565b6040516100bd919061025c565b60405180910390f35b60008060000154908060010154908060020154905083565b6000806002015482846100f191906102a6565b6100fb919061032f565b905092915050565b6000610115838360006002015461011d565b905092915050565b600080600190508285610130919061032f565b94505b600084111561019057600160028561014b919061032f565b141561016b5782858261015e91906102a6565b610168919061032f565b90505b600184901c935082858661017f91906102a6565b610189919061032f565b9450610133565b809150509392505050565b6000819050919050565b6101ae8161019b565b82525050565b60006060820190506101c960008301866101a5565b6101d660208301856101a5565b6101e360408301846101a5565b949350505050565b600080fd5b6101f98161019b565b811461020457600080fd5b50565b600081359050610216816101f0565b92915050565b60008060408385031215610233576102326101eb565b5b600061024185828601610207565b925050602061025285828601610207565b9150509250929050565b600060208201905061027160008301846101a5565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006102b18261019b565b91506102bc8361019b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156102f5576102f4610277565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061033a8261019b565b91506103458361019b565b92508261035557610354610300565b5b82820690509291505056fea2646970667358221220c2b40740f9f95acf8912c69c4fd8380f057dc5142736b225b4bc39a73b59e88364736f6c634300080b0033"
var HomomorphiccalculatorSMBin = "0x"

// DeployHomomorphiccalculator deploys a new contract, binding an instance of Homomorphiccalculator to it.
func DeployHomomorphiccalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _n *big.Int, _g *big.Int) (common.Address, *types.Receipt, *Homomorphiccalculator, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphiccalculatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(HomomorphiccalculatorSMBin)
	} else {
		bytecode = common.FromHex(HomomorphiccalculatorBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, HomomorphiccalculatorABI, backend, _n, _g)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Homomorphiccalculator{HomomorphiccalculatorCaller: HomomorphiccalculatorCaller{contract: contract}, HomomorphiccalculatorTransactor: HomomorphiccalculatorTransactor{contract: contract}, HomomorphiccalculatorFilterer: HomomorphiccalculatorFilterer{contract: contract}}, nil
}

func AsyncDeployHomomorphiccalculator(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _n *big.Int, _g *big.Int) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphiccalculatorABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(HomomorphiccalculatorSMBin)
	} else {
		bytecode = common.FromHex(HomomorphiccalculatorBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, HomomorphiccalculatorABI, backend, _n, _g)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Homomorphiccalculator is an auto generated Go binding around a Solidity contract.
type Homomorphiccalculator struct {
	HomomorphiccalculatorCaller     // Read-only binding to the contract
	HomomorphiccalculatorTransactor // Write-only binding to the contract
	HomomorphiccalculatorFilterer   // Log filterer for contract events
}

// HomomorphiccalculatorCaller is an auto generated read-only Go binding around a Solidity contract.
type HomomorphiccalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphiccalculatorTransactor is an auto generated write-only Go binding around a Solidity contract.
type HomomorphiccalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphiccalculatorFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type HomomorphiccalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphiccalculatorSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type HomomorphiccalculatorSession struct {
	Contract     *Homomorphiccalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HomomorphiccalculatorCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type HomomorphiccalculatorCallerSession struct {
	Contract *HomomorphiccalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// HomomorphiccalculatorTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type HomomorphiccalculatorTransactorSession struct {
	Contract     *HomomorphiccalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// HomomorphiccalculatorRaw is an auto generated low-level Go binding around a Solidity contract.
type HomomorphiccalculatorRaw struct {
	Contract *Homomorphiccalculator // Generic contract binding to access the raw methods on
}

// HomomorphiccalculatorCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type HomomorphiccalculatorCallerRaw struct {
	Contract *HomomorphiccalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// HomomorphiccalculatorTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type HomomorphiccalculatorTransactorRaw struct {
	Contract *HomomorphiccalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHomomorphiccalculator creates a new instance of Homomorphiccalculator, bound to a specific deployed contract.
func NewHomomorphiccalculator(address common.Address, backend bind.ContractBackend) (*Homomorphiccalculator, error) {
	contract, err := bindHomomorphiccalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Homomorphiccalculator{HomomorphiccalculatorCaller: HomomorphiccalculatorCaller{contract: contract}, HomomorphiccalculatorTransactor: HomomorphiccalculatorTransactor{contract: contract}, HomomorphiccalculatorFilterer: HomomorphiccalculatorFilterer{contract: contract}}, nil
}

// NewHomomorphiccalculatorCaller creates a new read-only instance of Homomorphiccalculator, bound to a specific deployed contract.
func NewHomomorphiccalculatorCaller(address common.Address, caller bind.ContractCaller) (*HomomorphiccalculatorCaller, error) {
	contract, err := bindHomomorphiccalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomomorphiccalculatorCaller{contract: contract}, nil
}

// NewHomomorphiccalculatorTransactor creates a new write-only instance of Homomorphiccalculator, bound to a specific deployed contract.
func NewHomomorphiccalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*HomomorphiccalculatorTransactor, error) {
	contract, err := bindHomomorphiccalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomomorphiccalculatorTransactor{contract: contract}, nil
}

// NewHomomorphiccalculatorFilterer creates a new log filterer instance of Homomorphiccalculator, bound to a specific deployed contract.
func NewHomomorphiccalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*HomomorphiccalculatorFilterer, error) {
	contract, err := bindHomomorphiccalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomomorphiccalculatorFilterer{contract: contract}, nil
}

// bindHomomorphiccalculator binds a generic wrapper to an already deployed contract.
func bindHomomorphiccalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphiccalculatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Homomorphiccalculator *HomomorphiccalculatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Homomorphiccalculator.Contract.HomomorphiccalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Homomorphiccalculator *HomomorphiccalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Homomorphiccalculator.Contract.HomomorphiccalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Homomorphiccalculator *HomomorphiccalculatorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Homomorphiccalculator.Contract.HomomorphiccalculatorTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Homomorphiccalculator *HomomorphiccalculatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Homomorphiccalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Homomorphiccalculator *HomomorphiccalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Homomorphiccalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Homomorphiccalculator *HomomorphiccalculatorTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Homomorphiccalculator.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// HomomorphicAdd is a free data retrieval call binding the contract method 0x59e18a05.
//
// Solidity: function homomorphicAdd(uint256 encryptedA, uint256 encryptedB) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorCaller) HomomorphicAdd(opts *bind.CallOpts, encryptedA *big.Int, encryptedB *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Homomorphiccalculator.contract.Call(opts, out, "homomorphicAdd", encryptedA, encryptedB)
	return *ret0, err
}

// HomomorphicAdd is a free data retrieval call binding the contract method 0x59e18a05.
//
// Solidity: function homomorphicAdd(uint256 encryptedA, uint256 encryptedB) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorSession) HomomorphicAdd(encryptedA *big.Int, encryptedB *big.Int) (*big.Int, error) {
	return _Homomorphiccalculator.Contract.HomomorphicAdd(&_Homomorphiccalculator.CallOpts, encryptedA, encryptedB)
}

// HomomorphicAdd is a free data retrieval call binding the contract method 0x59e18a05.
//
// Solidity: function homomorphicAdd(uint256 encryptedA, uint256 encryptedB) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorCallerSession) HomomorphicAdd(encryptedA *big.Int, encryptedB *big.Int) (*big.Int, error) {
	return _Homomorphiccalculator.Contract.HomomorphicAdd(&_Homomorphiccalculator.CallOpts, encryptedA, encryptedB)
}

// HomomorphicMultiply is a free data retrieval call binding the contract method 0xbaf38aae.
//
// Solidity: function homomorphicMultiply(uint256 encryptedA, uint256 scalarK) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorCaller) HomomorphicMultiply(opts *bind.CallOpts, encryptedA *big.Int, scalarK *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Homomorphiccalculator.contract.Call(opts, out, "homomorphicMultiply", encryptedA, scalarK)
	return *ret0, err
}

// HomomorphicMultiply is a free data retrieval call binding the contract method 0xbaf38aae.
//
// Solidity: function homomorphicMultiply(uint256 encryptedA, uint256 scalarK) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorSession) HomomorphicMultiply(encryptedA *big.Int, scalarK *big.Int) (*big.Int, error) {
	return _Homomorphiccalculator.Contract.HomomorphicMultiply(&_Homomorphiccalculator.CallOpts, encryptedA, scalarK)
}

// HomomorphicMultiply is a free data retrieval call binding the contract method 0xbaf38aae.
//
// Solidity: function homomorphicMultiply(uint256 encryptedA, uint256 scalarK) constant returns(uint256)
func (_Homomorphiccalculator *HomomorphiccalculatorCallerSession) HomomorphicMultiply(encryptedA *big.Int, scalarK *big.Int) (*big.Int, error) {
	return _Homomorphiccalculator.Contract.HomomorphicMultiply(&_Homomorphiccalculator.CallOpts, encryptedA, scalarK)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(uint256 n, uint256 g, uint256 nSquared)
func (_Homomorphiccalculator *HomomorphiccalculatorCaller) Key(opts *bind.CallOpts) (struct {
	N        *big.Int
	G        *big.Int
	NSquared *big.Int
}, error) {
	ret := new(struct {
		N        *big.Int
		G        *big.Int
		NSquared *big.Int
	})
	out := ret
	err := _Homomorphiccalculator.contract.Call(opts, out, "key")
	return *ret, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(uint256 n, uint256 g, uint256 nSquared)
func (_Homomorphiccalculator *HomomorphiccalculatorSession) Key() (struct {
	N        *big.Int
	G        *big.Int
	NSquared *big.Int
}, error) {
	return _Homomorphiccalculator.Contract.Key(&_Homomorphiccalculator.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(uint256 n, uint256 g, uint256 nSquared)
func (_Homomorphiccalculator *HomomorphiccalculatorCallerSession) Key() (struct {
	N        *big.Int
	G        *big.Int
	NSquared *big.Int
}, error) {
	return _Homomorphiccalculator.Contract.Key(&_Homomorphiccalculator.CallOpts)
}
