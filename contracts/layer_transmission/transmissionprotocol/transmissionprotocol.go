// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transmissionprotocol

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

// TransmissionprotocolABI is the input ABI used to generate the binding from.
const TransmissionprotocolABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransmissionprotocolBin is the compiled bytecode used for deploying new contracts.
var TransmissionprotocolBin = "0x608060405234801561001057600080fd5b50610b9b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633676d25b14610051578063634e5a9814610082578063a7cd3fed146100b3578063bff894a3146100e4575b600080fd5b61006b60048036038101906100669190610774565b610115565b604051610079929190610b2e565b60405180910390f35b61009c60048036038101906100979190610774565b610139565b6040516100aa929190610b2e565b60405180910390f35b6100cd60048036038101906100c89190610774565b61015d565b6040516100db929190610b2e565b60405180910390f35b6100fe60048036038101906100f99190610774565b610181565b60405161010c929190610b2e565b60405180910390f35b61011d6101a5565b6060826040518060200160405280600081525091509150915091565b6101416101a5565b6060826040518060200160405280600081525091509150915091565b6101656101a5565b6060826040518060200160405280600081525091509150915091565b6101896101a5565b6060826040518060200160405280600081525091509150915091565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61029a82610251565b810181811067ffffffffffffffff821117156102b9576102b8610262565b5b80604052505050565b60006102cc610238565b90506102d88282610291565b919050565b600080fd5b6000819050919050565b6102f5816102e2565b811461030057600080fd5b50565b600081359050610312816102ec565b92915050565b600080fd5b600067ffffffffffffffff82111561033857610337610262565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561036e5761036d610262565b5b61037782610251565b9050602081019050919050565b82818337600083830152505050565b60006103a66103a184610353565b6102c2565b9050828152602081018484840111156103c2576103c161034e565b5b6103cd848285610384565b509392505050565b600082601f8301126103ea576103e9610318565b5b81356103fa848260208601610393565b91505092915050565b60006104166104118461031d565b6102c2565b9050808382526020820190506020840283018581111561043957610438610349565b5b835b8181101561048057803567ffffffffffffffff81111561045e5761045d610318565b5b80860161046b89826103d5565b8552602085019450505060208101905061043b565b5050509392505050565b600082601f83011261049f5761049e610318565b5b81356104af848260208601610403565b91505092915050565b6000819050919050565b6104cb816104b8565b81146104d657600080fd5b50565b6000813590506104e8816104c2565b92915050565b60008115159050919050565b610503816104ee565b811461050e57600080fd5b50565b600081359050610520816104fa565b92915050565b6000610240828403121561053d5761053c61024c565b5b6105486102406102c2565b9050600061055884828501610303565b600083015250602061056c84828501610303565b602083015250604061058084828501610303565b604083015250606061059484828501610303565b60608301525060806105a884828501610303565b60808301525060a082013567ffffffffffffffff8111156105cc576105cb6102dd565b5b6105d88482850161048a565b60a08301525060c082013567ffffffffffffffff8111156105fc576105fb6102dd565b5b6106088482850161048a565b60c08301525060e061061c84828501610303565b60e08301525061010082013567ffffffffffffffff811115610641576106406102dd565b5b61064d8482850161048a565b6101008301525061012061066384828501610303565b6101208301525061014082013567ffffffffffffffff811115610689576106886102dd565b5b6106958482850161048a565b610140830152506101606106ab84828501610303565b6101608301525061018082013567ffffffffffffffff8111156106d1576106d06102dd565b5b6106dd8482850161048a565b610180830152506101a06106f384828501610303565b6101a0830152506101c082013567ffffffffffffffff811115610719576107186102dd565b5b6107258482850161048a565b6101c0830152506101e061073b848285016104d9565b6101e083015250610200610751848285016104d9565b6102008301525061022061076784828501610511565b6102208301525092915050565b60006020828403121561078a57610789610242565b5b600082013567ffffffffffffffff8111156107a8576107a7610247565b5b6107b484828501610526565b91505092915050565b6107c6816102e2565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610832578082015181840152602081019050610817565b83811115610841576000848401525b50505050565b6000610852826107f8565b61085c8185610803565b935061086c818560208601610814565b61087581610251565b840191505092915050565b600061088c8383610847565b905092915050565b6000602082019050919050565b60006108ac826107cc565b6108b681856107d7565b9350836020820285016108c8856107e8565b8060005b8581101561090457848403895281516108e58582610880565b94506108f083610894565b925060208a019950506001810190506108cc565b50829750879550505050505092915050565b61091f816104b8565b82525050565b61092e816104ee565b82525050565b60006102408301600083015161094d60008601826107bd565b50602083015161096060208601826107bd565b50604083015161097360408601826107bd565b50606083015161098660608601826107bd565b50608083015161099960808601826107bd565b5060a083015184820360a08601526109b182826108a1565b91505060c083015184820360c08601526109cb82826108a1565b91505060e08301516109e060e08601826107bd565b506101008301518482036101008601526109fa82826108a1565b915050610120830151610a116101208601826107bd565b50610140830151848203610140860152610a2b82826108a1565b915050610160830151610a426101608601826107bd565b50610180830151848203610180860152610a5c82826108a1565b9150506101a0830151610a736101a08601826107bd565b506101c08301518482036101c0860152610a8d82826108a1565b9150506101e0830151610aa46101e0860182610916565b50610200830151610ab9610200860182610916565b50610220830151610ace610220860182610925565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000610b0082610ad9565b610b0a8185610ae4565b9350610b1a818560208601610814565b610b2381610251565b840191505092915050565b60006040820190508181036000830152610b488185610934565b90508181036020830152610b5c8184610af5565b9050939250505056fea2646970667358221220f84a31e7c801f05d56ce2a56dbe74910923cbf94c789fb634d6dda2974a387f364736f6c634300080b0033"
var TransmissionprotocolSMBin = "0x"

// DeployTransmissionprotocol deploys a new contract, binding an instance of Transmissionprotocol to it.
func DeployTransmissionprotocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transmissionprotocol, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionprotocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransmissionprotocolSMBin)
	} else {
		bytecode = common.FromHex(TransmissionprotocolBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, TransmissionprotocolABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transmissionprotocol{TransmissionprotocolCaller: TransmissionprotocolCaller{contract: contract}, TransmissionprotocolTransactor: TransmissionprotocolTransactor{contract: contract}, TransmissionprotocolFilterer: TransmissionprotocolFilterer{contract: contract}}, nil
}

func AsyncDeployTransmissionprotocol(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionprotocolABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransmissionprotocolSMBin)
	} else {
		bytecode = common.FromHex(TransmissionprotocolBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, TransmissionprotocolABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transmissionprotocol is an auto generated Go binding around a Solidity contract.
type Transmissionprotocol struct {
	TransmissionprotocolCaller     // Read-only binding to the contract
	TransmissionprotocolTransactor // Write-only binding to the contract
	TransmissionprotocolFilterer   // Log filterer for contract events
}

// TransmissionprotocolCaller is an auto generated read-only Go binding around a Solidity contract.
type TransmissionprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionprotocolTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransmissionprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionprotocolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransmissionprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransmissionprotocolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransmissionprotocolSession struct {
	Contract     *Transmissionprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TransmissionprotocolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransmissionprotocolCallerSession struct {
	Contract *TransmissionprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TransmissionprotocolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransmissionprotocolTransactorSession struct {
	Contract     *TransmissionprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TransmissionprotocolRaw is an auto generated low-level Go binding around a Solidity contract.
type TransmissionprotocolRaw struct {
	Contract *Transmissionprotocol // Generic contract binding to access the raw methods on
}

// TransmissionprotocolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransmissionprotocolCallerRaw struct {
	Contract *TransmissionprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// TransmissionprotocolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransmissionprotocolTransactorRaw struct {
	Contract *TransmissionprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransmissionprotocol creates a new instance of Transmissionprotocol, bound to a specific deployed contract.
func NewTransmissionprotocol(address common.Address, backend bind.ContractBackend) (*Transmissionprotocol, error) {
	contract, err := bindTransmissionprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transmissionprotocol{TransmissionprotocolCaller: TransmissionprotocolCaller{contract: contract}, TransmissionprotocolTransactor: TransmissionprotocolTransactor{contract: contract}, TransmissionprotocolFilterer: TransmissionprotocolFilterer{contract: contract}}, nil
}

// NewTransmissionprotocolCaller creates a new read-only instance of Transmissionprotocol, bound to a specific deployed contract.
func NewTransmissionprotocolCaller(address common.Address, caller bind.ContractCaller) (*TransmissionprotocolCaller, error) {
	contract, err := bindTransmissionprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransmissionprotocolCaller{contract: contract}, nil
}

// NewTransmissionprotocolTransactor creates a new write-only instance of Transmissionprotocol, bound to a specific deployed contract.
func NewTransmissionprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*TransmissionprotocolTransactor, error) {
	contract, err := bindTransmissionprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransmissionprotocolTransactor{contract: contract}, nil
}

// NewTransmissionprotocolFilterer creates a new log filterer instance of Transmissionprotocol, bound to a specific deployed contract.
func NewTransmissionprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*TransmissionprotocolFilterer, error) {
	contract, err := bindTransmissionprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransmissionprotocolFilterer{contract: contract}, nil
}

// bindTransmissionprotocol binds a generic wrapper to an already deployed contract.
func bindTransmissionprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransmissionprotocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transmissionprotocol *TransmissionprotocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transmissionprotocol.Contract.TransmissionprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transmissionprotocol *TransmissionprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.TransmissionprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transmissionprotocol *TransmissionprotocolRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.TransmissionprotocolTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transmissionprotocol *TransmissionprotocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transmissionprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transmissionprotocol *TransmissionprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transmissionprotocol *TransmissionprotocolTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactor) Acknowledge(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Transmissionprotocol.contract.TransactWithResult(opts, out, "acknowledge", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Transmissionprotocol *TransmissionprotocolTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.contract.AsyncTransact(opts, handler, "acknowledge", ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.Acknowledge(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncAcknowledge(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x634e5a98.
//
// Solidity: function acknowledge(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactorSession) Acknowledge(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.Acknowledge(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncAcknowledge(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactor) ReceiveMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Transmissionprotocol.contract.TransactWithResult(opts, out, "receive_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Transmissionprotocol *TransmissionprotocolTransactor) AsyncReceiveMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.contract.AsyncTransact(opts, handler, "receive_msg", ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.ReceiveMsg(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncReceiveMsg(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// ReceiveMsg is a paid mutator transaction binding the contract method 0xa7cd3fed.
//
// Solidity: function receive_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactorSession) ReceiveMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.ReceiveMsg(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolTransactorSession) AsyncReceiveMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncReceiveMsg(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Transmissionprotocol.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Transmissionprotocol *TransmissionprotocolTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.Response(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncResponse(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.Response(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncResponse(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactor) SendMsg(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Transmissionprotocol.contract.TransactWithResult(opts, out, "send_msg", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Transmissionprotocol *TransmissionprotocolTransactor) AsyncSendMsg(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.contract.AsyncTransact(opts, handler, "send_msg", ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolSession) SendMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.SendMsg(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncSendMsg(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}

// SendMsg is a paid mutator transaction binding the contract method 0x3676d25b.
//
// Solidity: function send_msg(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Transmissionprotocol *TransmissionprotocolTransactorSession) SendMsg(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Transmissionprotocol.Contract.SendMsg(&_Transmissionprotocol.TransactOpts, ccMsg)
}

func (_Transmissionprotocol *TransmissionprotocolTransactorSession) AsyncSendMsg(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Transmissionprotocol.Contract.AsyncSendMsg(handler, &_Transmissionprotocol.TransactOpts, ccMsg)
}
