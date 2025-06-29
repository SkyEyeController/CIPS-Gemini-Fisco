// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivericationprotocol

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

// IvericationprotocolABI is the input ABI used to generate the binding from.
const IvericationprotocolABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"prepare\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ivericationprotocol is an auto generated Go binding around a Solidity contract.
type Ivericationprotocol struct {
	IvericationprotocolCaller     // Read-only binding to the contract
	IvericationprotocolTransactor // Write-only binding to the contract
	IvericationprotocolFilterer   // Log filterer for contract events
}

// IvericationprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type IvericationprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IvericationprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type IvericationprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IvericationprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IvericationprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IvericationprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IvericationprotocolSession struct {
	Contract     *Ivericationprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IvericationprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IvericationprotocolCallerSession struct {
	Contract *IvericationprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IvericationprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IvericationprotocolTransactorSession struct {
	Contract     *IvericationprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IvericationprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type IvericationprotocolRaw struct {
	Contract *Ivericationprotocol // Generic contract binding to access the raw methods on
}

// IvericationprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IvericationprotocolCallerRaw struct {
	Contract *IvericationprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// IvericationprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IvericationprotocolTransactorRaw struct {
	Contract *IvericationprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIvericationprotocol creates a new instance of Ivericationprotocol, bound to a specific deployed contract.
func NewIvericationprotocol(address common.Address, backend bind.ContractBackend) (*Ivericationprotocol, error) {
	contract, err := bindIvericationprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ivericationprotocol{IvericationprotocolCaller: IvericationprotocolCaller{contract: contract}, IvericationprotocolTransactor: IvericationprotocolTransactor{contract: contract}, IvericationprotocolFilterer: IvericationprotocolFilterer{contract: contract}}, nil
}

// NewIvericationprotocolCaller creates a new read-only instance of Ivericationprotocol, bound to a specific deployed contract.
func NewIvericationprotocolCaller(address common.Address, caller bind.ContractCaller) (*IvericationprotocolCaller, error) {
	contract, err := bindIvericationprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IvericationprotocolCaller{contract: contract}, nil
}

// NewIvericationprotocolTransactor creates a new write-only instance of Ivericationprotocol, bound to a specific deployed contract.
func NewIvericationprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*IvericationprotocolTransactor, error) {
	contract, err := bindIvericationprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IvericationprotocolTransactor{contract: contract}, nil
}

// NewIvericationprotocolFilterer creates a new log filterer instance of Ivericationprotocol, bound to a specific deployed contract.
func NewIvericationprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*IvericationprotocolFilterer, error) {
	contract, err := bindIvericationprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IvericationprotocolFilterer{contract: contract}, nil
}

// bindIvericationprotocol binds a generic wrapper to an already deployed contract.
func bindIvericationprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IvericationprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ivericationprotocol *IvericationprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ivericationprotocol.Contract.IvericationprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ivericationprotocol *IvericationprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.IvericationprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ivericationprotocol *IvericationprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.IvericationprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ivericationprotocol *IvericationprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ivericationprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ivericationprotocol *IvericationprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ivericationprotocol *IvericationprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolTransactor) Prepare(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Ivericationprotocol.contract.TransactWithResult(opts, out, "prepare", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Ivericationprotocol *IvericationprotocolTransactor) AsyncPrepare(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.contract.AsyncTransact(opts, handler, "prepare", ccMsg)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolSession) Prepare(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Prepare(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolSession) AsyncPrepare(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncPrepare(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolTransactorSession) Prepare(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Prepare(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolTransactorSession) AsyncPrepare(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncPrepare(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Ivericationprotocol.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Ivericationprotocol *IvericationprotocolTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Response(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncResponse(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Ivericationprotocol *IvericationprotocolTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Response(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncResponse(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolTransactor) Update(opts *bind.TransactOpts, data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Ivericationprotocol.contract.TransactWithResult(opts, out, "update", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Ivericationprotocol *IvericationprotocolTransactor) AsyncUpdate(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Ivericationprotocol.contract.AsyncTransact(opts, handler, "update", data)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolSession) Update(data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Update(&_Ivericationprotocol.TransactOpts, data)
}

func (_Ivericationprotocol *IvericationprotocolSession) AsyncUpdate(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncUpdate(handler, &_Ivericationprotocol.TransactOpts, data)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolTransactorSession) Update(data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Update(&_Ivericationprotocol.TransactOpts, data)
}

func (_Ivericationprotocol *IvericationprotocolTransactorSession) AsyncUpdate(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncUpdate(handler, &_Ivericationprotocol.TransactOpts, data)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolTransactor) Verify(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Ivericationprotocol.contract.TransactWithResult(opts, out, "verify", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Ivericationprotocol *IvericationprotocolTransactor) AsyncVerify(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.contract.AsyncTransact(opts, handler, "verify", ccMsg)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolSession) Verify(ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Verify(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolSession) AsyncVerify(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncVerify(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Ivericationprotocol *IvericationprotocolTransactorSession) Verify(ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Ivericationprotocol.Contract.Verify(&_Ivericationprotocol.TransactOpts, ccMsg)
}

func (_Ivericationprotocol *IvericationprotocolTransactorSession) AsyncVerify(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Ivericationprotocol.Contract.AsyncVerify(handler, &_Ivericationprotocol.TransactOpts, ccMsg)
}
