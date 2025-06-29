// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iapp

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

// IappABI is the input ABI used to generate the binding from.
const IappABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Iapp is an auto generated Go binding around a Solidity contract.
type Iapp struct {
	IappCaller     // Read-only binding to the contract
	IappTransactor // Write-only binding to the contract
	IappFilterer   // Log filterer for contract events
}

// IappCaller is an auto generated read-only Go binding around a Solidity contract.
type IappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IappTransactor is an auto generated write-only Go binding around a Solidity contract.
type IappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IappFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IappSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IappSession struct {
	Contract     *Iapp             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IappCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IappCallerSession struct {
	Contract *IappCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IappTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IappTransactorSession struct {
	Contract     *IappTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IappRaw is an auto generated low-level Go binding around a Solidity contract.
type IappRaw struct {
	Contract *Iapp // Generic contract binding to access the raw methods on
}

// IappCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IappCallerRaw struct {
	Contract *IappCaller // Generic read-only contract binding to access the raw methods on
}

// IappTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IappTransactorRaw struct {
	Contract *IappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIapp creates a new instance of Iapp, bound to a specific deployed contract.
func NewIapp(address common.Address, backend bind.ContractBackend) (*Iapp, error) {
	contract, err := bindIapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iapp{IappCaller: IappCaller{contract: contract}, IappTransactor: IappTransactor{contract: contract}, IappFilterer: IappFilterer{contract: contract}}, nil
}

// NewIappCaller creates a new read-only instance of Iapp, bound to a specific deployed contract.
func NewIappCaller(address common.Address, caller bind.ContractCaller) (*IappCaller, error) {
	contract, err := bindIapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IappCaller{contract: contract}, nil
}

// NewIappTransactor creates a new write-only instance of Iapp, bound to a specific deployed contract.
func NewIappTransactor(address common.Address, transactor bind.ContractTransactor) (*IappTransactor, error) {
	contract, err := bindIapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IappTransactor{contract: contract}, nil
}

// NewIappFilterer creates a new log filterer instance of Iapp, bound to a specific deployed contract.
func NewIappFilterer(address common.Address, filterer bind.ContractFilterer) (*IappFilterer, error) {
	contract, err := bindIapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IappFilterer{contract: contract}, nil
}

// bindIapp binds a generic wrapper to an already deployed contract.
func bindIapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iapp *IappRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iapp.Contract.IappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iapp *IappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.IappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iapp *IappRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.IappTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iapp *IappCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Iapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iapp *IappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iapp *IappTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactor) Acknowledge(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Iapp.contract.TransactWithResult(opts, out, "acknowledge", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Iapp *IappTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.contract.AsyncTransact(opts, handler, "acknowledge", ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.Acknowledge(&_Iapp.TransactOpts, ccMsg)
}

func (_Iapp *IappSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncAcknowledge(handler, &_Iapp.TransactOpts, ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactorSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.Acknowledge(&_Iapp.TransactOpts, ccMsg)
}

func (_Iapp *IappTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncAcknowledge(handler, &_Iapp.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactor) ReceiveMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Iapp.contract.TransactWithResult(opts, out, "receive_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Iapp *IappTransactor) AsyncReceiveMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.contract.AsyncTransact(opts, handler, "receive_msg", ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.ReceiveMsg(&_Iapp.TransactOpts, ccMsg)
}

func (_Iapp *IappSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncReceiveMsg(handler, &_Iapp.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactorSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.ReceiveMsg(&_Iapp.TransactOpts, ccMsg)
}

func (_Iapp *IappTransactorSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncReceiveMsg(handler, &_Iapp.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactor) SendMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Iapp.contract.TransactWithResult(opts, out, "send_msg", ccMsg, payload)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Iapp *IappTransactor) AsyncSendMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Iapp.contract.AsyncTransact(opts, handler, "send_msg", ccMsg, payload)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Iapp *IappSession) SendMsg(ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.SendMsg(&_Iapp.TransactOpts, ccMsg, payload)
}

func (_Iapp *IappSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncSendMsg(handler, &_Iapp.TransactOpts, ccMsg, payload)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Iapp *IappTransactorSession) SendMsg(ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Iapp.Contract.SendMsg(&_Iapp.TransactOpts, ccMsg, payload)
}

func (_Iapp *IappTransactorSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Iapp.Contract.AsyncSendMsg(handler, &_Iapp.TransactOpts, ccMsg, payload)
}
