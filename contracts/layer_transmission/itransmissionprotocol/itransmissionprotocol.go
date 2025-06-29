// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itransmissionprotocol

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

// ItransmissionprotocolABI is the input ABI used to generate the binding from.
const ItransmissionprotocolABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Itransmissionprotocol is an auto generated Go binding around a Solidity contract.
type Itransmissionprotocol struct {
	ItransmissionprotocolCaller     // Read-only binding to the contract
	ItransmissionprotocolTransactor // Write-only binding to the contract
	ItransmissionprotocolFilterer   // Log filterer for contract events
}

// ItransmissionprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type ItransmissionprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransmissionprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type ItransmissionprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransmissionprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ItransmissionprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItransmissionprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ItransmissionprotocolSession struct {
	Contract     *Itransmissionprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ItransmissionprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ItransmissionprotocolCallerSession struct {
	Contract *ItransmissionprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ItransmissionprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ItransmissionprotocolTransactorSession struct {
	Contract     *ItransmissionprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ItransmissionprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type ItransmissionprotocolRaw struct {
	Contract *Itransmissionprotocol // Generic contract binding to access the raw methods on
}

// ItransmissionprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ItransmissionprotocolCallerRaw struct {
	Contract *ItransmissionprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// ItransmissionprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ItransmissionprotocolTransactorRaw struct {
	Contract *ItransmissionprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItransmissionprotocol creates a new instance of Itransmissionprotocol, bound to a specific deployed contract.
func NewItransmissionprotocol(address common.Address, backend bind.ContractBackend) (*Itransmissionprotocol, error) {
	contract, err := bindItransmissionprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itransmissionprotocol{ItransmissionprotocolCaller: ItransmissionprotocolCaller{contract: contract}, ItransmissionprotocolTransactor: ItransmissionprotocolTransactor{contract: contract}, ItransmissionprotocolFilterer: ItransmissionprotocolFilterer{contract: contract}}, nil
}

// NewItransmissionprotocolCaller creates a new read-only instance of Itransmissionprotocol, bound to a specific deployed contract.
func NewItransmissionprotocolCaller(address common.Address, caller bind.ContractCaller) (*ItransmissionprotocolCaller, error) {
	contract, err := bindItransmissionprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItransmissionprotocolCaller{contract: contract}, nil
}

// NewItransmissionprotocolTransactor creates a new write-only instance of Itransmissionprotocol, bound to a specific deployed contract.
func NewItransmissionprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ItransmissionprotocolTransactor, error) {
	contract, err := bindItransmissionprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItransmissionprotocolTransactor{contract: contract}, nil
}

// NewItransmissionprotocolFilterer creates a new log filterer instance of Itransmissionprotocol, bound to a specific deployed contract.
func NewItransmissionprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ItransmissionprotocolFilterer, error) {
	contract, err := bindItransmissionprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItransmissionprotocolFilterer{contract: contract}, nil
}

// bindItransmissionprotocol binds a generic wrapper to an already deployed contract.
func bindItransmissionprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItransmissionprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransmissionprotocol *ItransmissionprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransmissionprotocol.Contract.ItransmissionprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransmissionprotocol *ItransmissionprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.ItransmissionprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransmissionprotocol *ItransmissionprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.ItransmissionprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itransmissionprotocol *ItransmissionprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Itransmissionprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itransmissionprotocol *ItransmissionprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itransmissionprotocol *ItransmissionprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactor) Acknowledge(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransmissionprotocol.contract.TransactWithResult(opts, out, "acknowledge", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.contract.AsyncTransact(opts, handler, "acknowledge", ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.Acknowledge(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncAcknowledge(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.Acknowledge(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncAcknowledge(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactor) ReceiveMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransmissionprotocol.contract.TransactWithResult(opts, out, "receive_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactor) AsyncReceiveMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.contract.AsyncTransact(opts, handler, "receive_msg", ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.ReceiveMsg(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncReceiveMsg(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.ReceiveMsg(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncReceiveMsg(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransmissionprotocol.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.Response(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncResponse(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.Response(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncResponse(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactor) SendMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Itransmissionprotocol.contract.TransactWithResult(opts, out, "send_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactor) AsyncSendMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.contract.AsyncTransact(opts, handler, "send_msg", ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolSession) SendMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.SendMsg(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncSendMsg(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) SendMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Itransmissionprotocol.Contract.SendMsg(&_Itransmissionprotocol.TransactOpts, ccMsg)
}

func (_Itransmissionprotocol *ItransmissionprotocolTransactorSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Itransmissionprotocol.Contract.AsyncSendMsg(handler, &_Itransmissionprotocol.TransactOpts, ccMsg)
}
