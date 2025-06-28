// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vericationprotocol

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

// VericationprotocolABI is the input ABI used to generate the binding from.
const VericationprotocolABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"prepare\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VericationprotocolBin is the compiled bytecode used for deploying new contracts.
var VericationprotocolBin = "0x608060405234801561001057600080fd5b50610c19806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806322348a4714610051578063504a9bff14610082578063bff894a3146100b3578063f52a3468146100e4575b600080fd5b61006b6004803603810190610066919061076a565b610115565b604051610079929190610b24565b60405180910390f35b61009c6004803603810190610097919061076a565b610139565b6040516100aa929190610b6a565b60405180910390f35b6100cd60048036038101906100c8919061076a565b610158565b6040516100db929190610b24565b60405180910390f35b6100fe60048036038101906100f99190610b9a565b61017c565b60405161010c929190610b6a565b60405180910390f35b61011d61019b565b6060826040518060200160405280600081525091509150915091565b6000606060016040518060200160405280600081525091509150915091565b61016061019b565b6060826040518060200160405280600081525091509150915091565b6000606060016040518060200160405280600081525091509150915091565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61029082610247565b810181811067ffffffffffffffff821117156102af576102ae610258565b5b80604052505050565b60006102c261022e565b90506102ce8282610287565b919050565b600080fd5b6000819050919050565b6102eb816102d8565b81146102f657600080fd5b50565b600081359050610308816102e2565b92915050565b600080fd5b600067ffffffffffffffff82111561032e5761032d610258565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561036457610363610258565b5b61036d82610247565b9050602081019050919050565b82818337600083830152505050565b600061039c61039784610349565b6102b8565b9050828152602081018484840111156103b8576103b7610344565b5b6103c384828561037a565b509392505050565b600082601f8301126103e0576103df61030e565b5b81356103f0848260208601610389565b91505092915050565b600061040c61040784610313565b6102b8565b9050808382526020820190506020840283018581111561042f5761042e61033f565b5b835b8181101561047657803567ffffffffffffffff8111156104545761045361030e565b5b80860161046189826103cb565b85526020850194505050602081019050610431565b5050509392505050565b600082601f8301126104955761049461030e565b5b81356104a58482602086016103f9565b91505092915050565b6000819050919050565b6104c1816104ae565b81146104cc57600080fd5b50565b6000813590506104de816104b8565b92915050565b60008115159050919050565b6104f9816104e4565b811461050457600080fd5b50565b600081359050610516816104f0565b92915050565b6000610240828403121561053357610532610242565b5b61053e6102406102b8565b9050600061054e848285016102f9565b6000830152506020610562848285016102f9565b6020830152506040610576848285016102f9565b604083015250606061058a848285016102f9565b606083015250608061059e848285016102f9565b60808301525060a082013567ffffffffffffffff8111156105c2576105c16102d3565b5b6105ce84828501610480565b60a08301525060c082013567ffffffffffffffff8111156105f2576105f16102d3565b5b6105fe84828501610480565b60c08301525060e0610612848285016102f9565b60e08301525061010082013567ffffffffffffffff811115610637576106366102d3565b5b61064384828501610480565b61010083015250610120610659848285016102f9565b6101208301525061014082013567ffffffffffffffff81111561067f5761067e6102d3565b5b61068b84828501610480565b610140830152506101606106a1848285016102f9565b6101608301525061018082013567ffffffffffffffff8111156106c7576106c66102d3565b5b6106d384828501610480565b610180830152506101a06106e9848285016102f9565b6101a0830152506101c082013567ffffffffffffffff81111561070f5761070e6102d3565b5b61071b84828501610480565b6101c0830152506101e0610731848285016104cf565b6101e083015250610200610747848285016104cf565b6102008301525061022061075d84828501610507565b6102208301525092915050565b6000602082840312156107805761077f610238565b5b600082013567ffffffffffffffff81111561079e5761079d61023d565b5b6107aa8482850161051c565b91505092915050565b6107bc816102d8565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561082857808201518184015260208101905061080d565b83811115610837576000848401525b50505050565b6000610848826107ee565b61085281856107f9565b935061086281856020860161080a565b61086b81610247565b840191505092915050565b6000610882838361083d565b905092915050565b6000602082019050919050565b60006108a2826107c2565b6108ac81856107cd565b9350836020820285016108be856107de565b8060005b858110156108fa57848403895281516108db8582610876565b94506108e68361088a565b925060208a019950506001810190506108c2565b50829750879550505050505092915050565b610915816104ae565b82525050565b610924816104e4565b82525050565b60006102408301600083015161094360008601826107b3565b50602083015161095660208601826107b3565b50604083015161096960408601826107b3565b50606083015161097c60608601826107b3565b50608083015161098f60808601826107b3565b5060a083015184820360a08601526109a78282610897565b91505060c083015184820360c08601526109c18282610897565b91505060e08301516109d660e08601826107b3565b506101008301518482036101008601526109f08282610897565b915050610120830151610a076101208601826107b3565b50610140830151848203610140860152610a218282610897565b915050610160830151610a386101608601826107b3565b50610180830151848203610180860152610a528282610897565b9150506101a0830151610a696101a08601826107b3565b506101c08301518482036101c0860152610a838282610897565b9150506101e0830151610a9a6101e086018261090c565b50610200830151610aaf61020086018261090c565b50610220830151610ac461022086018261091b565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000610af682610acf565b610b008185610ada565b9350610b1081856020860161080a565b610b1981610247565b840191505092915050565b60006040820190508181036000830152610b3e818561092a565b90508181036020830152610b528184610aeb565b90509392505050565b610b64816104e4565b82525050565b6000604082019050610b7f6000830185610b5b565b8181036020830152610b918184610aeb565b90509392505050565b600060208284031215610bb057610baf610238565b5b600082013567ffffffffffffffff811115610bce57610bcd61023d565b5b610bda84828501610480565b9150509291505056fea264697066735822122095cb0551e1e417b35ab48ba2d01e6434144c4dd964bfbef20dfe51bc18cf2fef64736f6c634300080b0033"
var VericationprotocolSMBin = "0x"

// DeployVericationprotocol deploys a new contract, binding an instance of Vericationprotocol to it.
func DeployVericationprotocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Vericationprotocol, error) {
	parsed, err := abi.JSON(strings.NewReader(VericationprotocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(VericationprotocolSMBin)
	} else {
		bytecode = common.FromHex(VericationprotocolBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, VericationprotocolABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Vericationprotocol{VericationprotocolCaller: VericationprotocolCaller{contract: contract}, VericationprotocolTransactor: VericationprotocolTransactor{contract: contract}, VericationprotocolFilterer: VericationprotocolFilterer{contract: contract}}, nil
}

func AsyncDeployVericationprotocol(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(VericationprotocolABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(VericationprotocolSMBin)
	} else {
		bytecode = common.FromHex(VericationprotocolBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, VericationprotocolABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Vericationprotocol is an auto generated Go binding around a Solidity contract.
type Vericationprotocol struct {
	VericationprotocolCaller     // Read-only binding to the contract
	VericationprotocolTransactor // Write-only binding to the contract
	VericationprotocolFilterer   // Log filterer for contract events
}

// VericationprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type VericationprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VericationprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type VericationprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VericationprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type VericationprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VericationprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type VericationprotocolSession struct {
	Contract     *Vericationprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VericationprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type VericationprotocolCallerSession struct {
	Contract *VericationprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VericationprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type VericationprotocolTransactorSession struct {
	Contract     *VericationprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VericationprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type VericationprotocolRaw struct {
	Contract *Vericationprotocol // Generic contract binding to access the raw methods on
}

// VericationprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type VericationprotocolCallerRaw struct {
	Contract *VericationprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// VericationprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type VericationprotocolTransactorRaw struct {
	Contract *VericationprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVericationprotocol creates a new instance of Vericationprotocol, bound to a specific deployed contract.
func NewVericationprotocol(address common.Address, backend bind.ContractBackend) (*Vericationprotocol, error) {
	contract, err := bindVericationprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vericationprotocol{VericationprotocolCaller: VericationprotocolCaller{contract: contract}, VericationprotocolTransactor: VericationprotocolTransactor{contract: contract}, VericationprotocolFilterer: VericationprotocolFilterer{contract: contract}}, nil
}

// NewVericationprotocolCaller creates a new read-only instance of Vericationprotocol, bound to a specific deployed contract.
func NewVericationprotocolCaller(address common.Address, caller bind.ContractCaller) (*VericationprotocolCaller, error) {
	contract, err := bindVericationprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VericationprotocolCaller{contract: contract}, nil
}

// NewVericationprotocolTransactor creates a new write-only instance of Vericationprotocol, bound to a specific deployed contract.
func NewVericationprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*VericationprotocolTransactor, error) {
	contract, err := bindVericationprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VericationprotocolTransactor{contract: contract}, nil
}

// NewVericationprotocolFilterer creates a new log filterer instance of Vericationprotocol, bound to a specific deployed contract.
func NewVericationprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*VericationprotocolFilterer, error) {
	contract, err := bindVericationprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VericationprotocolFilterer{contract: contract}, nil
}

// bindVericationprotocol binds a generic wrapper to an already deployed contract.
func bindVericationprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VericationprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vericationprotocol *VericationprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vericationprotocol.Contract.VericationprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vericationprotocol *VericationprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.VericationprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vericationprotocol *VericationprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.VericationprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vericationprotocol *VericationprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vericationprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vericationprotocol *VericationprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vericationprotocol *VericationprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolTransactor) Prepare(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Vericationprotocol.contract.TransactWithResult(opts, out, "prepare", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Vericationprotocol *VericationprotocolTransactor) AsyncPrepare(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.contract.AsyncTransact(opts, handler, "prepare", ccMsg)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolSession) Prepare(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Prepare(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolSession) AsyncPrepare(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncPrepare(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}

// Prepare is a paid mutator transaction binding the contract method 0x22348a47.
//
// Solidity: function prepare(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolTransactorSession) Prepare(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Prepare(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolTransactorSession) AsyncPrepare(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncPrepare(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Vericationprotocol.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Vericationprotocol *VericationprotocolTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Response(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncResponse(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Vericationprotocol *VericationprotocolTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Response(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncResponse(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Vericationprotocol *VericationprotocolTransactor) Update(opts *bind.TransactOpts, data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Vericationprotocol.contract.TransactWithResult(opts, out, "update", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Vericationprotocol *VericationprotocolTransactor) AsyncUpdate(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Vericationprotocol.contract.AsyncTransact(opts, handler, "update", data)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Vericationprotocol *VericationprotocolSession) Update(data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Update(&_Vericationprotocol.TransactOpts, data)
}

func (_Vericationprotocol *VericationprotocolSession) AsyncUpdate(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncUpdate(handler, &_Vericationprotocol.TransactOpts, data)
}

// Update is a paid mutator transaction binding the contract method 0xf52a3468.
//
// Solidity: function update(bytes[] data) returns(bool, string)
func (_Vericationprotocol *VericationprotocolTransactorSession) Update(data [][]byte) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Update(&_Vericationprotocol.TransactOpts, data)
}

func (_Vericationprotocol *VericationprotocolTransactorSession) AsyncUpdate(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncUpdate(handler, &_Vericationprotocol.TransactOpts, data)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Vericationprotocol *VericationprotocolTransactor) Verify(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Vericationprotocol.contract.TransactWithResult(opts, out, "verify", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Vericationprotocol *VericationprotocolTransactor) AsyncVerify(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.contract.AsyncTransact(opts, handler, "verify", ccMsg)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Vericationprotocol *VericationprotocolSession) Verify(ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Verify(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolSession) AsyncVerify(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncVerify(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}

// Verify is a paid mutator transaction binding the contract method 0x504a9bff.
//
// Solidity: function verify(TypesCrosschainMessage ccMsg) returns(bool, string)
func (_Vericationprotocol *VericationprotocolTransactorSession) Verify(ccMsg TypesCrosschainMessage) (bool, string, *types.Transaction, *types.Receipt, error) {
	return _Vericationprotocol.Contract.Verify(&_Vericationprotocol.TransactOpts, ccMsg)
}

func (_Vericationprotocol *VericationprotocolTransactorSession) AsyncVerify(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Vericationprotocol.Contract.AsyncVerify(handler, &_Vericationprotocol.TransactOpts, ccMsg)
}
