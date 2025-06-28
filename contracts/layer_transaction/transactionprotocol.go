// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transactionprotocol

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

// TransactionprotocolABI is the input ABI used to generate the binding from.
const TransactionprotocolABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"flag\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"work\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransactionprotocolBin is the compiled bytecode used for deploying new contracts.
var TransactionprotocolBin = "0x608060405234801561001057600080fd5b50610af9806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806390f248a714610030575b600080fd5b61004a60048036038101906100459190610690565b610061565b604051610058929190610a8c565b60405180910390f35b610069610088565b6060836040518060200160405280600081525091509150935093915050565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600060ff82169050919050565b6101458161012f565b811461015057600080fd5b50565b6000813590506101628161013c565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101b68261016d565b810181811067ffffffffffffffff821117156101d5576101d461017e565b5b80604052505050565b60006101e861011b565b90506101f482826101ad565b919050565b600080fd5b6000819050919050565b610211816101fe565b811461021c57600080fd5b50565b60008135905061022e81610208565b92915050565b600080fd5b600067ffffffffffffffff8211156102545761025361017e565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561028a5761028961017e565b5b6102938261016d565b9050602081019050919050565b82818337600083830152505050565b60006102c26102bd8461026f565b6101de565b9050828152602081018484840111156102de576102dd61026a565b5b6102e98482856102a0565b509392505050565b600082601f83011261030657610305610234565b5b81356103168482602086016102af565b91505092915050565b600061033261032d84610239565b6101de565b9050808382526020820190506020840283018581111561035557610354610265565b5b835b8181101561039c57803567ffffffffffffffff81111561037a57610379610234565b5b80860161038789826102f1565b85526020850194505050602081019050610357565b5050509392505050565b600082601f8301126103bb576103ba610234565b5b81356103cb84826020860161031f565b91505092915050565b6000819050919050565b6103e7816103d4565b81146103f257600080fd5b50565b600081359050610404816103de565b92915050565b60008115159050919050565b61041f8161040a565b811461042a57600080fd5b50565b60008135905061043c81610416565b92915050565b6000610240828403121561045957610458610168565b5b6104646102406101de565b905060006104748482850161021f565b60008301525060206104888482850161021f565b602083015250604061049c8482850161021f565b60408301525060606104b08482850161021f565b60608301525060806104c48482850161021f565b60808301525060a082013567ffffffffffffffff8111156104e8576104e76101f9565b5b6104f4848285016103a6565b60a08301525060c082013567ffffffffffffffff811115610518576105176101f9565b5b610524848285016103a6565b60c08301525060e06105388482850161021f565b60e08301525061010082013567ffffffffffffffff81111561055d5761055c6101f9565b5b610569848285016103a6565b6101008301525061012061057f8482850161021f565b6101208301525061014082013567ffffffffffffffff8111156105a5576105a46101f9565b5b6105b1848285016103a6565b610140830152506101606105c78482850161021f565b6101608301525061018082013567ffffffffffffffff8111156105ed576105ec6101f9565b5b6105f9848285016103a6565b610180830152506101a061060f8482850161021f565b6101a0830152506101c082013567ffffffffffffffff811115610635576106346101f9565b5b610641848285016103a6565b6101c0830152506101e0610657848285016103f5565b6101e08301525061020061066d848285016103f5565b610200830152506102206106838482850161042d565b6102208301525092915050565b6000806000606084860312156106a9576106a8610125565b5b60006106b786828701610153565b935050602084013567ffffffffffffffff8111156106d8576106d761012a565b5b6106e486828701610442565b925050604084013567ffffffffffffffff8111156107055761070461012a565b5b610711868287016103a6565b9150509250925092565b610724816101fe565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610790578082015181840152602081019050610775565b8381111561079f576000848401525b50505050565b60006107b082610756565b6107ba8185610761565b93506107ca818560208601610772565b6107d38161016d565b840191505092915050565b60006107ea83836107a5565b905092915050565b6000602082019050919050565b600061080a8261072a565b6108148185610735565b93508360208202850161082685610746565b8060005b85811015610862578484038952815161084385826107de565b945061084e836107f2565b925060208a0199505060018101905061082a565b50829750879550505050505092915050565b61087d816103d4565b82525050565b61088c8161040a565b82525050565b6000610240830160008301516108ab600086018261071b565b5060208301516108be602086018261071b565b5060408301516108d1604086018261071b565b5060608301516108e4606086018261071b565b5060808301516108f7608086018261071b565b5060a083015184820360a086015261090f82826107ff565b91505060c083015184820360c086015261092982826107ff565b91505060e083015161093e60e086018261071b565b5061010083015184820361010086015261095882826107ff565b91505061012083015161096f61012086018261071b565b5061014083015184820361014086015261098982826107ff565b9150506101608301516109a061016086018261071b565b506101808301518482036101808601526109ba82826107ff565b9150506101a08301516109d16101a086018261071b565b506101c08301518482036101c08601526109eb82826107ff565b9150506101e0830151610a026101e0860182610874565b50610200830151610a17610200860182610874565b50610220830151610a2c610220860182610883565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000610a5e82610a37565b610a688185610a42565b9350610a78818560208601610772565b610a818161016d565b840191505092915050565b60006040820190508181036000830152610aa68185610892565b90508181036020830152610aba8184610a53565b9050939250505056fea2646970667358221220a7f8a51c3526c8f94ceef75ddfd01673c58f019585d7fb9812154144bee8ef0a64736f6c634300080b0033"
var TransactionprotocolSMBin = "0x"

// DeployTransactionprotocol deploys a new contract, binding an instance of Transactionprotocol to it.
func DeployTransactionprotocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transactionprotocol, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionprotocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransactionprotocolSMBin)
	} else {
		bytecode = common.FromHex(TransactionprotocolBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TransactionprotocolABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transactionprotocol{TransactionprotocolCaller: TransactionprotocolCaller{contract: contract}, TransactionprotocolTransactor: TransactionprotocolTransactor{contract: contract}, TransactionprotocolFilterer: TransactionprotocolFilterer{contract: contract}}, nil
}

func AsyncDeployTransactionprotocol(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionprotocolABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransactionprotocolSMBin)
	} else {
		bytecode = common.FromHex(TransactionprotocolBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TransactionprotocolABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transactionprotocol is an auto generated Go binding around a Solidity contract.
type Transactionprotocol struct {
	TransactionprotocolCaller     // Read-only binding to the contract
	TransactionprotocolTransactor // Write-only binding to the contract
	TransactionprotocolFilterer   // Log filterer for contract events
}

// TransactionprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type TransactionprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransactionprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransactionprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransactionprotocolSession struct {
	Contract     *Transactionprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TransactionprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransactionprotocolCallerSession struct {
	Contract *TransactionprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TransactionprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransactionprotocolTransactorSession struct {
	Contract     *TransactionprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TransactionprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type TransactionprotocolRaw struct {
	Contract *Transactionprotocol // Generic contract binding to access the raw methods on
}

// TransactionprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransactionprotocolCallerRaw struct {
	Contract *TransactionprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransactionprotocolTransactorRaw struct {
	Contract *TransactionprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionprotocol creates a new instance of Transactionprotocol, bound to a specific deployed contract.
func NewTransactionprotocol(address common.Address, backend bind.ContractBackend) (*Transactionprotocol, error) {
	contract, err := bindTransactionprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transactionprotocol{TransactionprotocolCaller: TransactionprotocolCaller{contract: contract}, TransactionprotocolTransactor: TransactionprotocolTransactor{contract: contract}, TransactionprotocolFilterer: TransactionprotocolFilterer{contract: contract}}, nil
}

// NewTransactionprotocolCaller creates a new read-only instance of Transactionprotocol, bound to a specific deployed contract.
func NewTransactionprotocolCaller(address common.Address, caller bind.ContractCaller) (*TransactionprotocolCaller, error) {
	contract, err := bindTransactionprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionprotocolCaller{contract: contract}, nil
}

// NewTransactionprotocolTransactor creates a new write-only instance of Transactionprotocol, bound to a specific deployed contract.
func NewTransactionprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionprotocolTransactor, error) {
	contract, err := bindTransactionprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionprotocolTransactor{contract: contract}, nil
}

// NewTransactionprotocolFilterer creates a new log filterer instance of Transactionprotocol, bound to a specific deployed contract.
func NewTransactionprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionprotocolFilterer, error) {
	contract, err := bindTransactionprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionprotocolFilterer{contract: contract}, nil
}

// bindTransactionprotocol binds a generic wrapper to an already deployed contract.
func bindTransactionprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionprotocol *TransactionprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactionprotocol.Contract.TransactionprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionprotocol *TransactionprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.TransactionprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionprotocol *TransactionprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.TransactionprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionprotocol *TransactionprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactionprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionprotocol *TransactionprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionprotocol *TransactionprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Transactionprotocol *TransactionprotocolTransactor) Work(opts *bind.TransactOpts, flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Transactionprotocol.contract.TransactWithResult(opts, out, "work", flag, ccMsg, payload)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Transactionprotocol *TransactionprotocolTransactor) AsyncWork(handler func(*types.Receipt, error), opts *bind.TransactOpts, flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Transactionprotocol.contract.AsyncTransact(opts, handler, "work", flag, ccMsg, payload)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Transactionprotocol *TransactionprotocolSession) Work(flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.Work(&_Transactionprotocol.TransactOpts, flag, ccMsg, payload)
}

func (_Transactionprotocol *TransactionprotocolSession) AsyncWork(handler func(*types.Receipt, error), flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Transactionprotocol.Contract.AsyncWork(handler, &_Transactionprotocol.TransactOpts, flag, ccMsg, payload)
}

// Work is a paid mutator transaction binding the contract method 0x90f248a7.
//
// Solidity: function work(uint8 flag, TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_Transactionprotocol *TransactionprotocolTransactorSession) Work(flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transactionprotocol.Contract.Work(&_Transactionprotocol.TransactOpts, flag, ccMsg, payload)
}

func (_Transactionprotocol *TransactionprotocolTransactorSession) AsyncWork(handler func(*types.Receipt, error), flag uint8, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _Transactionprotocol.Contract.AsyncWork(handler, &_Transactionprotocol.TransactOpts, flag, ccMsg, payload)
}
