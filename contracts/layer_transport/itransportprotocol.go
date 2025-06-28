// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itransportprotocol

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

// ItransportprotocolABI is the input ABI used to generate the binding from.
const ItransportprotocolABI = "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"receiveIn\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"sendOut\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Itransportprotocol is an auto generated Go binding around a Solidity contract.
type Itransportprotocol struct {
	ItransportprotocolCaller     // Read-only binding to the contract
	ItransportprotocolTransactor // Write-only binding to the contract
	ItransportprotocolFilterer   // Log filterer for contract events
}

// ItransportprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type ItransportprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransportprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type ItransportprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransportprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ItransportprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransportprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ItransportprotocolSession struct {
	Contract     *Itransportprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ItransportprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ItransportprotocolCallerSession struct {
	Contract *ItransportprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ItransportprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ItransportprotocolTransactorSession struct {
	Contract     *ItransportprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ItransportprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type ItransportprotocolRaw struct {
	Contract *Itransportprotocol // Generic contract binding to access the raw methods on
}

// ItransportprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ItransportprotocolCallerRaw struct {
	Contract *ItransportprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// ItransportprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ItransportprotocolTransactorRaw struct {
	Contract *ItransportprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItransportprotocol creates a new instance of Itransportprotocol, bound to a specific deployed contract.
func NewItransportprotocol(address common.Address, backend bind.ContractBackend) (*Itransportprotocol, error) {
	contract, err := bindItransportprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itransportprotocol{ItransportprotocolCaller: ItransportprotocolCaller{contract: contract}, ItransportprotocolTransactor: ItransportprotocolTransactor{contract: contract}, ItransportprotocolFilterer: ItransportprotocolFilterer{contract: contract}}, nil
}

// NewItransportprotocolCaller creates a new read-only instance of Itransportprotocol, bound to a specific deployed contract.
func NewItransportprotocolCaller(address common.Address, caller bind.ContractCaller) (*ItransportprotocolCaller, error) {
	contract, err := bindItransportprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItransportprotocolCaller{contract: contract}, nil
}

// NewItransportprotocolTransactor creates a new write-only instance of Itransportprotocol, bound to a specific deployed contract.
func NewItransportprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ItransportprotocolTransactor, error) {
	contract, err := bindItransportprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItransportprotocolTransactor{contract: contract}, nil
}

// NewItransportprotocolFilterer creates a new log filterer instance of Itransportprotocol, bound to a specific deployed contract.
func NewItransportprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ItransportprotocolFilterer, error) {
	contract, err := bindItransportprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItransportprotocolFilterer{contract: contract}, nil
}

// bindItransportprotocol binds a generic wrapper to an already deployed contract.
func bindItransportprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItransportprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransportprotocol *ItransportprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransportprotocol.Contract.ItransportprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransportprotocol *ItransportprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.ItransportprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransportprotocol *ItransportprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.ItransportprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransportprotocol *ItransportprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransportprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransportprotocol *ItransportprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransportprotocol *ItransportprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactor) Acknowledge(opts *bind.TransactOpts, data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransportprotocol.contract.TransactWithResult(opts, out, "acknowledge", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransportprotocol *ItransportprotocolTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.contract.AsyncTransact(opts, handler, "acknowledge", data)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolSession) Acknowledge(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.Acknowledge(&_Itransportprotocol.TransactOpts, data)
}

func (_Itransportprotocol *ItransportprotocolSession) AsyncAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncAcknowledge(handler, &_Itransportprotocol.TransactOpts, data)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactorSession) Acknowledge(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.Acknowledge(&_Itransportprotocol.TransactOpts, data)
}

func (_Itransportprotocol *ItransportprotocolTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncAcknowledge(handler, &_Itransportprotocol.TransactOpts, data)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactor) ReceiveIn(opts *bind.TransactOpts, data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransportprotocol.contract.TransactWithResult(opts, out, "receiveIn", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransportprotocol *ItransportprotocolTransactor) AsyncReceiveIn(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.contract.AsyncTransact(opts, handler, "receiveIn", data)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolSession) ReceiveIn(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.ReceiveIn(&_Itransportprotocol.TransactOpts, data)
}

func (_Itransportprotocol *ItransportprotocolSession) AsyncReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncReceiveIn(handler, &_Itransportprotocol.TransactOpts, data)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactorSession) ReceiveIn(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.ReceiveIn(&_Itransportprotocol.TransactOpts, data)
}

func (_Itransportprotocol *ItransportprotocolTransactorSession) AsyncReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncReceiveIn(handler, &_Itransportprotocol.TransactOpts, data)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransportprotocol.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransportprotocol *ItransportprotocolTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.Response(&_Itransportprotocol.TransactOpts, ccMsg)
}

func (_Itransportprotocol *ItransportprotocolSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncResponse(handler, &_Itransportprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.Response(&_Itransportprotocol.TransactOpts, ccMsg)
}

func (_Itransportprotocol *ItransportprotocolTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncResponse(handler, &_Itransportprotocol.TransactOpts, ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactor) SendOut(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransportprotocol.contract.TransactWithResult(opts, out, "sendOut", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransportprotocol *ItransportprotocolTransactor) AsyncSendOut(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.contract.AsyncTransact(opts, handler, "sendOut", ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolSession) SendOut(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.SendOut(&_Itransportprotocol.TransactOpts, ccMsg)
}

func (_Itransportprotocol *ItransportprotocolSession) AsyncSendOut(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncSendOut(handler, &_Itransportprotocol.TransactOpts, ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransportprotocol *ItransportprotocolTransactorSession) SendOut(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransportprotocol.Contract.SendOut(&_Itransportprotocol.TransactOpts, ccMsg)
}

func (_Itransportprotocol *ItransportprotocolTransactorSession) AsyncSendOut(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransportprotocol.Contract.AsyncSendOut(handler, &_Itransportprotocol.TransactOpts, ccMsg)
}
