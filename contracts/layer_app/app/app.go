// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package app

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

// AppABI is the input ABI used to generate the binding from.
const AppABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AppBin is the compiled bytecode used for deploying new contracts.
var AppBin = "0x608060405234801561001057600080fd5b50610be4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063528c157814610046578063634e5a9814610077578063a7cd3fed146100a8575b600080fd5b610060600480360381019061005b9190610745565b6100d9565b60405161006e929190610b2e565b60405180910390f35b610091600480360381019061008c9190610b65565b61012e565b60405161009f929190610b2e565b60405180910390f35b6100c260048036038101906100bd9190610b65565b610152565b6040516100d0929190610b2e565b60405180910390f35b6100e1610176565b606060018460e00181815250506001846101200181815250506001846101600181815250506001846101a00181815250508360405180602001604052806000815250915091509250929050565b610136610176565b6060826040518060200160405280600081525091509150915091565b61015a610176565b6060826040518060200160405280600081525091509150915091565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61026b82610222565b810181811067ffffffffffffffff8211171561028a57610289610233565b5b80604052505050565b600061029d610209565b90506102a98282610262565b919050565b600080fd5b6000819050919050565b6102c6816102b3565b81146102d157600080fd5b50565b6000813590506102e3816102bd565b92915050565b600080fd5b600067ffffffffffffffff82111561030957610308610233565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561033f5761033e610233565b5b61034882610222565b9050602081019050919050565b82818337600083830152505050565b600061037761037284610324565b610293565b9050828152602081018484840111156103935761039261031f565b5b61039e848285610355565b509392505050565b600082601f8301126103bb576103ba6102e9565b5b81356103cb848260208601610364565b91505092915050565b60006103e76103e2846102ee565b610293565b9050808382526020820190506020840283018581111561040a5761040961031a565b5b835b8181101561045157803567ffffffffffffffff81111561042f5761042e6102e9565b5b80860161043c89826103a6565b8552602085019450505060208101905061040c565b5050509392505050565b600082601f8301126104705761046f6102e9565b5b81356104808482602086016103d4565b91505092915050565b6000819050919050565b61049c81610489565b81146104a757600080fd5b50565b6000813590506104b981610493565b92915050565b60008115159050919050565b6104d4816104bf565b81146104df57600080fd5b50565b6000813590506104f1816104cb565b92915050565b6000610240828403121561050e5761050d61021d565b5b610519610240610293565b90506000610529848285016102d4565b600083015250602061053d848285016102d4565b6020830152506040610551848285016102d4565b6040830152506060610565848285016102d4565b6060830152506080610579848285016102d4565b60808301525060a082013567ffffffffffffffff81111561059d5761059c6102ae565b5b6105a98482850161045b565b60a08301525060c082013567ffffffffffffffff8111156105cd576105cc6102ae565b5b6105d98482850161045b565b60c08301525060e06105ed848285016102d4565b60e08301525061010082013567ffffffffffffffff811115610612576106116102ae565b5b61061e8482850161045b565b61010083015250610120610634848285016102d4565b6101208301525061014082013567ffffffffffffffff81111561065a576106596102ae565b5b6106668482850161045b565b6101408301525061016061067c848285016102d4565b6101608301525061018082013567ffffffffffffffff8111156106a2576106a16102ae565b5b6106ae8482850161045b565b610180830152506101a06106c4848285016102d4565b6101a0830152506101c082013567ffffffffffffffff8111156106ea576106e96102ae565b5b6106f68482850161045b565b6101c0830152506101e061070c848285016104aa565b6101e083015250610200610722848285016104aa565b61020083015250610220610738848285016104e2565b6102208301525092915050565b6000806040838503121561075c5761075b610213565b5b600083013567ffffffffffffffff81111561077a57610779610218565b5b610786858286016104f7565b925050602083013567ffffffffffffffff8111156107a7576107a6610218565b5b6107b38582860161045b565b9150509250929050565b6107c6816102b3565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610832578082015181840152602081019050610817565b83811115610841576000848401525b50505050565b6000610852826107f8565b61085c8185610803565b935061086c818560208601610814565b61087581610222565b840191505092915050565b600061088c8383610847565b905092915050565b6000602082019050919050565b60006108ac826107cc565b6108b681856107d7565b9350836020820285016108c8856107e8565b8060005b8581101561090457848403895281516108e58582610880565b94506108f083610894565b925060208a019950506001810190506108cc565b50829750879550505050505092915050565b61091f81610489565b82525050565b61092e816104bf565b82525050565b60006102408301600083015161094d60008601826107bd565b50602083015161096060208601826107bd565b50604083015161097360408601826107bd565b50606083015161098660608601826107bd565b50608083015161099960808601826107bd565b5060a083015184820360a08601526109b182826108a1565b91505060c083015184820360c08601526109cb82826108a1565b91505060e08301516109e060e08601826107bd565b506101008301518482036101008601526109fa82826108a1565b915050610120830151610a116101208601826107bd565b50610140830151848203610140860152610a2b82826108a1565b915050610160830151610a426101608601826107bd565b50610180830151848203610180860152610a5c82826108a1565b9150506101a0830151610a736101a08601826107bd565b506101c08301518482036101c0860152610a8d82826108a1565b9150506101e0830151610aa46101e0860182610916565b50610200830151610ab9610200860182610916565b50610220830151610ace610220860182610925565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000610b0082610ad9565b610b0a8185610ae4565b9350610b1a818560208601610814565b610b2381610222565b840191505092915050565b60006040820190508181036000830152610b488185610934565b90508181036020830152610b5c8184610af5565b90509392505050565b600060208284031215610b7b57610b7a610213565b5b600082013567ffffffffffffffff811115610b9957610b98610218565b5b610ba5848285016104f7565b9150509291505056fea26469706673582212209adfb4be88cf326decc49a4d55ccce2c694a34e1063b1030e5557400f04e016964736f6c634300080b0033"
var AppSMBin = "0x"

// DeployApp deploys a new contract, binding an instance of App to it.
func DeployApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *App, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(AppSMBin)
	} else {
		bytecode = common.FromHex(AppBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, AppABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

func AsyncDeployApp(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(AppSMBin)
	} else {
		bytecode = common.FromHex(AppBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, AppABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// App is an auto generated Go binding around a Solidity contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around a Solidity contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around a Solidity contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around a Solidity contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.AppTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppTransactor) Acknowledge(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _App.contract.TransactWithResult(opts, out, "acknowledge", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_App *AppTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.contract.AsyncTransact(opts, handler, "acknowledge", ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.Acknowledge(&_App.TransactOpts, ccMsg)
}

func (_App *AppSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.Contract.AsyncAcknowledge(handler, &_App.TransactOpts, ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppTransactorSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.Acknowledge(&_App.TransactOpts, ccMsg)
}

func (_App *AppTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.Contract.AsyncAcknowledge(handler, &_App.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppTransactor) ReceiveMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _App.contract.TransactWithResult(opts, out, "receive_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_App *AppTransactor) AsyncReceiveMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.contract.AsyncTransact(opts, handler, "receive_msg", ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.ReceiveMsg(&_App.TransactOpts, ccMsg)
}

func (_App *AppSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.Contract.AsyncReceiveMsg(handler, &_App.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_App *AppTransactorSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.ReceiveMsg(&_App.TransactOpts, ccMsg)
}

func (_App *AppTransactorSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _App.Contract.AsyncReceiveMsg(handler, &_App.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_App *AppTransactor) SendMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _App.contract.TransactWithResult(opts, out, "send_msg", ccMsg, payload)
	return *ret0, *ret1, transaction, receipt, err
}

func (_App *AppTransactor) AsyncSendMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _App.contract.AsyncTransact(opts, handler, "send_msg", ccMsg, payload)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_App *AppSession) SendMsg(ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.SendMsg(&_App.TransactOpts, ccMsg, payload)
}

func (_App *AppSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _App.Contract.AsyncSendMsg(handler, &_App.TransactOpts, ccMsg, payload)
}

// SendMsg is a paid mutator transaction binding the contract method 0x528c1578.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg, bytes[] payload) returns(TypesCrosschainMessage, string)
func (_App *AppTransactorSession) SendMsg(ccMsg TypesCrosschainMessage, payload [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _App.Contract.SendMsg(&_App.TransactOpts, ccMsg, payload)
}

func (_App *AppTransactorSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage, payload [][]byte) (*types.Transaction, error) {
	return _App.Contract.AsyncSendMsg(handler, &_App.TransactOpts, ccMsg, payload)
}
