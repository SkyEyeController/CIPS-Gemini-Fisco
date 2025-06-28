// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itransactionprotocol

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

// TypesCrosschainMessage is an auto generated low-level Go binding around an user-defined struct.
type TypesCrosschainMessage struct {
	SrcChainId          *big.Int
	DstChainId          *big.Int
	Seq                 *big.Int
	SrcAppId            *big.Int
	DstAppId            *big.Int
	PayloadReq          [][]byte
	PayloadResp         [][]byte
	TransactionTypeId   *big.Int
	TransactionPayload  [][]byte
	TransmissionTypeId  *big.Int
	TransmissionPayload [][]byte
	VerificationTypeId  *big.Int
	VerificationPayload [][]byte
	TransportTypeId     *big.Int
	TransportPayload    [][]byte
	HashReq             [32]byte
	HashResp            [32]byte
	Ack                 bool
}

// ItransactionprotocolABI is the input ABI used to generate the binding from.
const ItransactionprotocolABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"flag\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"work\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Itransactionprotocol is an auto generated Go binding around a Solidity contract.
type Itransactionprotocol struct {
	ItransactionprotocolCaller     // Read-only binding to the contract
	ItransactionprotocolTransactor // Write-only binding to the contract
	ItransactionprotocolFilterer   // Log filterer for contract events
}

// ItransactionprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type ItransactionprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransactionprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type ItransactionprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransactionprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ItransactionprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransactionprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ItransactionprotocolSession struct {
	Contract     *Itransactionprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ItransactionprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ItransactionprotocolCallerSession struct {
	Contract *ItransactionprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ItransactionprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ItransactionprotocolTransactorSession struct {
	Contract     *ItransactionprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ItransactionprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type ItransactionprotocolRaw struct {
	Contract *Itransactionprotocol // Generic contract binding to access the raw methods on
}

// ItransactionprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ItransactionprotocolCallerRaw struct {
	Contract *ItransactionprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// ItransactionprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ItransactionprotocolTransactorRaw struct {
	Contract *ItransactionprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItransactionprotocol creates a new instance of Itransactionprotocol, bound to a specific deployed contract.
func NewItransactionprotocol(address common.Address, backend bind.ContractBackend) (*Itransactionprotocol, error) {
	contract, err := bindItransactionprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itransactionprotocol{ItransactionprotocolCaller: ItransactionprotocolCaller{contract: contract}, ItransactionprotocolTransactor: ItransactionprotocolTransactor{contract: contract}, ItransactionprotocolFilterer: ItransactionprotocolFilterer{contract: contract}}, nil
}

// NewItransactionprotocolCaller creates a new read-only instance of Itransactionprotocol, bound to a specific deployed contract.
func NewItransactionprotocolCaller(address common.Address, caller bind.ContractCaller) (*ItransactionprotocolCaller, error) {
	contract, err := bindItransactionprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItransactionprotocolCaller{contract: contract}, nil
}

// NewItransactionprotocolTransactor creates a new write-only instance of Itransactionprotocol, bound to a specific deployed contract.
func NewItransactionprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ItransactionprotocolTransactor, error) {
	contract, err := bindItransactionprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItransactionprotocolTransactor{contract: contract}, nil
}

// NewItransactionprotocolFilterer creates a new log filterer instance of Itransactionprotocol, bound to a specific deployed contract.
func NewItransactionprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ItransactionprotocolFilterer, error) {
	contract, err := bindItransactionprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItransactionprotocolFilterer{contract: contract}, nil
}

// bindItransactionprotocol binds a generic wrapper to an already deployed contract.
func bindItransactionprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItransactionprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransactionprotocol *ItransactionprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransactionprotocol.Contract.ItransactionprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransactionprotocol *ItransactionprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.ItransactionprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransactionprotocol *ItransactionprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.ItransactionprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransactionprotocol *ItransactionprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransactionprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransactionprotocol *ItransactionprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransactionprotocol *ItransactionprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Itransactionprotocol *ItransactionprotocolTransactor) Work(opts *bind.TransactOpts, flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransactionprotocol.contract.TransactWithResult(opts, out, "work", flag, ccMsg, payload)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransactionprotocol *ItransactionprotocolTransactor) AsyncWork(handler func(*types.Receipt, error), opts *bind.TransactOpts, flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Itransactionprotocol.contract.AsyncTransact(opts, handler, "work", flag, ccMsg, payload)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Itransactionprotocol *ItransactionprotocolSession) Work(flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.Work(&_Itransactionprotocol.TransactOpts, flag, ccMsg, payload)
}

func (_Itransactionprotocol *ItransactionprotocolSession) AsyncWork(handler func(*types.Receipt, error), flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Itransactionprotocol.Contract.AsyncWork(handler, &_Itransactionprotocol.TransactOpts, flag, ccMsg, payload)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Itransactionprotocol *ItransactionprotocolTransactorSession) Work(flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransactionprotocol.Contract.Work(&_Itransactionprotocol.TransactOpts, flag, ccMsg, payload)
}

func (_Itransactionprotocol *ItransactionprotocolTransactorSession) AsyncWork(handler func(*types.Receipt, error), flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Itransactionprotocol.Contract.AsyncWork(handler, &_Itransactionprotocol.TransactOpts, flag, ccMsg, payload)
}
