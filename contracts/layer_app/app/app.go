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
const AppABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"Msgsendsuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ValueSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AppBin is the compiled bytecode used for deploying new contracts.
var AppBin = "0x608060405234801561001057600080fd5b50611590806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063528c15781461005c578063634e5a981461008d578063693ec85e146100be578063a7cd3fed146100ee578063e942b5161461011f575b600080fd5b61007660048036038101906100719190610cb3565b61013b565b60405161008492919061109c565b60405180910390f35b6100a760048036038101906100a291906110d3565b6101e3565b6040516100b592919061109c565b60405180910390f35b6100d860048036038101906100d39190611177565b610207565b6040516100e591906111c4565b60405180910390f35b610108600480360381019061010391906110d3565b6102da565b60405161011692919061109c565b60405180910390f35b610139600480360381019061013491906111e6565b61059a565b005b610143610641565b606060018460e00181815250506001846101200181815250506001846101600181815250506001846101a0018181525050828460a0018190525060008460a001515111156101c7577f7b2abf4495a2a10686d70c8e6f97d93e16251f33ee4e869ddd7f5f234d31882f8460a001516040516101be91906112ed565b60405180910390a15b8360405180602001604052806000815250915091509250929050565b6101eb610641565b6060826040518060200160405280600081525091509150915091565b60606000838360405160200161021e92919061133f565b604051602081830303815290604052805190602001209050600080828152602001908152602001600020805461025390611387565b80601f016020809104026020016040519081016040528092919081815260200182805461027f90611387565b80156102cc5780601f106102a1576101008083540402835291602001916102cc565b820191906000526020600020905b8154815290600101906020018083116102af57829003601f168201915b505050505091505092915050565b6102e2610641565b606060028360a001515114156104155760003073ffffffffffffffffffffffffffffffffffffffff1663693ec85e8560a00151600181518110610328576103276113b9565b5b60200260200101516040518263ffffffff1660e01b815260040161034c91906111c4565b600060405180830381865afa158015610369573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103929190611489565b9050600167ffffffffffffffff8111156103af576103ae6107a1565b5b6040519080825280602002602001820160405280156103e257816020015b60608152602001906001900390816103cd5790505b508460c00181905250808460c00151600081518110610404576104036113b9565b5b602002602001018190525050610580565b60038360a0015151141561057f573073ffffffffffffffffffffffffffffffffffffffff1663e942b5168460a00151600181518110610457576104566113b9565b5b60200260200101518560a00151600281518110610477576104766113b9565b5b60200260200101516040518363ffffffff1660e01b815260040161049c9291906114d2565b600060405180830381600087803b1580156104b657600080fd5b505af11580156104ca573d6000803e3d6000fd5b50505050600167ffffffffffffffff8111156104e9576104e86107a1565b5b60405190808252806020026020018201604052801561051c57816020015b60608152602001906001900390816105075790505b508360c001819052506040518060400160405280600b81526020017f73657420737563636573730000000000000000000000000000000000000000008152508360c00151600081518110610573576105726113b9565b5b60200260200101819052505b5b826040518060200160405280600081525091509150915091565b600084846040516020016105af92919061133f565b604051602081830303815290604052805190602001209050828260008084815260200190815260200160002091906105e89291906106d4565b5084846040516105f992919061133f565b60405180910390207fc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f58484604051610632929190611536565b60405180910390a25050505050565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b8280546106e090611387565b90600052602060002090601f0160209004810192826107025760008555610749565b82601f1061071b57803560ff1916838001178555610749565b82800160010185558215610749579182015b8281111561074857823582559160200191906001019061072d565b5b509050610756919061075a565b5090565b5b8082111561077357600081600090555060010161075b565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107d982610790565b810181811067ffffffffffffffff821117156107f8576107f76107a1565b5b80604052505050565b600061080b610777565b905061081782826107d0565b919050565b600080fd5b6000819050919050565b61083481610821565b811461083f57600080fd5b50565b6000813590506108518161082b565b92915050565b600080fd5b600067ffffffffffffffff821115610877576108766107a1565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff8211156108ad576108ac6107a1565b5b6108b682610790565b9050602081019050919050565b82818337600083830152505050565b60006108e56108e084610892565b610801565b9050828152602081018484840111156109015761090061088d565b5b61090c8482856108c3565b509392505050565b600082601f83011261092957610928610857565b5b81356109398482602086016108d2565b91505092915050565b60006109556109508461085c565b610801565b9050808382526020820190506020840283018581111561097857610977610888565b5b835b818110156109bf57803567ffffffffffffffff81111561099d5761099c610857565b5b8086016109aa8982610914565b8552602085019450505060208101905061097a565b5050509392505050565b600082601f8301126109de576109dd610857565b5b81356109ee848260208601610942565b91505092915050565b6000819050919050565b610a0a816109f7565b8114610a1557600080fd5b50565b600081359050610a2781610a01565b92915050565b60008115159050919050565b610a4281610a2d565b8114610a4d57600080fd5b50565b600081359050610a5f81610a39565b92915050565b60006102408284031215610a7c57610a7b61078b565b5b610a87610240610801565b90506000610a9784828501610842565b6000830152506020610aab84828501610842565b6020830152506040610abf84828501610842565b6040830152506060610ad384828501610842565b6060830152506080610ae784828501610842565b60808301525060a082013567ffffffffffffffff811115610b0b57610b0a61081c565b5b610b17848285016109c9565b60a08301525060c082013567ffffffffffffffff811115610b3b57610b3a61081c565b5b610b47848285016109c9565b60c08301525060e0610b5b84828501610842565b60e08301525061010082013567ffffffffffffffff811115610b8057610b7f61081c565b5b610b8c848285016109c9565b61010083015250610120610ba284828501610842565b6101208301525061014082013567ffffffffffffffff811115610bc857610bc761081c565b5b610bd4848285016109c9565b61014083015250610160610bea84828501610842565b6101608301525061018082013567ffffffffffffffff811115610c1057610c0f61081c565b5b610c1c848285016109c9565b610180830152506101a0610c3284828501610842565b6101a0830152506101c082013567ffffffffffffffff811115610c5857610c5761081c565b5b610c64848285016109c9565b6101c0830152506101e0610c7a84828501610a18565b6101e083015250610200610c9084828501610a18565b61020083015250610220610ca684828501610a50565b6102208301525092915050565b60008060408385031215610cca57610cc9610781565b5b600083013567ffffffffffffffff811115610ce857610ce7610786565b5b610cf485828601610a65565b925050602083013567ffffffffffffffff811115610d1557610d14610786565b5b610d21858286016109c9565b9150509250929050565b610d3481610821565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610da0578082015181840152602081019050610d85565b83811115610daf576000848401525b50505050565b6000610dc082610d66565b610dca8185610d71565b9350610dda818560208601610d82565b610de381610790565b840191505092915050565b6000610dfa8383610db5565b905092915050565b6000602082019050919050565b6000610e1a82610d3a565b610e248185610d45565b935083602082028501610e3685610d56565b8060005b85811015610e725784840389528151610e538582610dee565b9450610e5e83610e02565b925060208a01995050600181019050610e3a565b50829750879550505050505092915050565b610e8d816109f7565b82525050565b610e9c81610a2d565b82525050565b600061024083016000830151610ebb6000860182610d2b565b506020830151610ece6020860182610d2b565b506040830151610ee16040860182610d2b565b506060830151610ef46060860182610d2b565b506080830151610f076080860182610d2b565b5060a083015184820360a0860152610f1f8282610e0f565b91505060c083015184820360c0860152610f398282610e0f565b91505060e0830151610f4e60e0860182610d2b565b50610100830151848203610100860152610f688282610e0f565b915050610120830151610f7f610120860182610d2b565b50610140830151848203610140860152610f998282610e0f565b915050610160830151610fb0610160860182610d2b565b50610180830151848203610180860152610fca8282610e0f565b9150506101a0830151610fe16101a0860182610d2b565b506101c08301518482036101c0860152610ffb8282610e0f565b9150506101e08301516110126101e0860182610e84565b50610200830151611027610200860182610e84565b5061022083015161103c610220860182610e93565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b600061106e82611047565b6110788185611052565b9350611088818560208601610d82565b61109181610790565b840191505092915050565b600060408201905081810360008301526110b68185610ea2565b905081810360208301526110ca8184611063565b90509392505050565b6000602082840312156110e9576110e8610781565b5b600082013567ffffffffffffffff81111561110757611106610786565b5b61111384828501610a65565b91505092915050565b600080fd5b60008083601f84011261113757611136610857565b5b8235905067ffffffffffffffff8111156111545761115361111c565b5b6020830191508360018202830111156111705761116f610888565b5b9250929050565b6000806020838503121561118e5761118d610781565b5b600083013567ffffffffffffffff8111156111ac576111ab610786565b5b6111b885828601611121565b92509250509250929050565b600060208201905081810360008301526111de8184611063565b905092915050565b60008060008060408587031215611200576111ff610781565b5b600085013567ffffffffffffffff81111561121e5761121d610786565b5b61122a87828801611121565b9450945050602085013567ffffffffffffffff81111561124d5761124c610786565b5b61125987828801611121565b925092505092959194509250565b600082825260208201905092915050565b600061128382610d3a565b61128d8185611267565b93508360208202850161129f85610d56565b8060005b858110156112db57848403895281516112bc8582610dee565b94506112c783610e02565b925060208a019950506001810190506112a3565b50829750879550505050505092915050565b600060208201905081810360008301526113078184611278565b905092915050565b600081905092915050565b6000611326838561130f565b93506113338385846108c3565b82840190509392505050565b600061134c82848661131a565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061139f57607f821691505b602082108114156113b3576113b2611358565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600067ffffffffffffffff821115611403576114026107a1565b5b61140c82610790565b9050602081019050919050565b600061142c611427846113e8565b610801565b9050828152602081018484840111156114485761144761088d565b5b611453848285610d82565b509392505050565b600082601f8301126114705761146f610857565b5b8151611480848260208601611419565b91505092915050565b60006020828403121561149f5761149e610781565b5b600082015167ffffffffffffffff8111156114bd576114bc610786565b5b6114c98482850161145b565b91505092915050565b600060408201905081810360008301526114ec8185611063565b905081810360208301526115008184611063565b90509392505050565b60006115158385611052565b93506115228385846108c3565b61152b83610790565b840190509392505050565b60006020820190508181036000830152611551818486611509565b9050939250505056fea2646970667358221220edc6eeee159d743a8dc7d9896b3da709eddce0f3e4dd0956e22c3d1a83bd8c3464736f6c634300080b0033"
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

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_App *AppCaller) Get(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _App.contract.Call(opts, out, "get", key)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_App *AppSession) Get(key string) (string, error) {
	return _App.Contract.Get(&_App.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(string)
func (_App *AppCallerSession) Get(key string) (string, error) {
	return _App.Contract.Get(&_App.CallOpts, key)
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

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_App *AppTransactor) Set(opts *bind.TransactOpts, key string, value string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _App.contract.TransactWithResult(opts, out, "set", key, value)
	return transaction, receipt, err
}

func (_App *AppTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _App.contract.AsyncTransact(opts, handler, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_App *AppSession) Set(key string, value string) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.Set(&_App.TransactOpts, key, value)
}

func (_App *AppSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _App.Contract.AsyncSet(handler, &_App.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_App *AppTransactorSession) Set(key string, value string) (*types.Transaction, *types.Receipt, error) {
	return _App.Contract.Set(&_App.TransactOpts, key, value)
}

func (_App *AppTransactorSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _App.Contract.AsyncSet(handler, &_App.TransactOpts, key, value)
}

// AppMsgsendsuccess represents a Msgsendsuccess event raised by the App contract.
type AppMsgsendsuccess struct {
	Payload [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchMsgsendsuccess is a free log subscription operation binding the contract event 0x7b2abf4495a2a10686d70c8e6f97d93e16251f33ee4e869ddd7f5f234d31882f.
//
// Solidity: event Msgsendsuccess(bytes[] payload)
func (_App *AppFilterer) WatchMsgsendsuccess(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.contract.WatchLogs(fromBlock, handler, "Msgsendsuccess")
}

func (_App *AppFilterer) WatchAllMsgsendsuccess(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.contract.WatchLogs(fromBlock, handler, "Msgsendsuccess")
}

// ParseMsgsendsuccess is a log parse operation binding the contract event 0x7b2abf4495a2a10686d70c8e6f97d93e16251f33ee4e869ddd7f5f234d31882f.
//
// Solidity: event Msgsendsuccess(bytes[] payload)
func (_App *AppFilterer) ParseMsgsendsuccess(log types.Log) (*AppMsgsendsuccess, error) {
	event := new(AppMsgsendsuccess)
	if err := _App.contract.UnpackLog(event, "Msgsendsuccess", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchMsgsendsuccess is a free log subscription operation binding the contract event 0x7b2abf4495a2a10686d70c8e6f97d93e16251f33ee4e869ddd7f5f234d31882f.
//
// Solidity: event Msgsendsuccess(bytes[] payload)
func (_App *AppSession) WatchMsgsendsuccess(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.Contract.WatchMsgsendsuccess(fromBlock, handler)
}

func (_App *AppSession) WatchAllMsgsendsuccess(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.Contract.WatchAllMsgsendsuccess(fromBlock, handler)
}

// ParseMsgsendsuccess is a log parse operation binding the contract event 0x7b2abf4495a2a10686d70c8e6f97d93e16251f33ee4e869ddd7f5f234d31882f.
//
// Solidity: event Msgsendsuccess(bytes[] payload)
func (_App *AppSession) ParseMsgsendsuccess(log types.Log) (*AppMsgsendsuccess, error) {
	return _App.Contract.ParseMsgsendsuccess(log)
}

// AppValueSet represents a ValueSet event raised by the App contract.
type AppValueSet struct {
	Key   common.Hash
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchValueSet is a free log subscription operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_App *AppFilterer) WatchValueSet(fromBlock *int64, handler func(int, []types.Log), key string) (string, error) {
	return _App.contract.WatchLogs(fromBlock, handler, "ValueSet", key)
}

func (_App *AppFilterer) WatchAllValueSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.contract.WatchLogs(fromBlock, handler, "ValueSet")
}

// ParseValueSet is a log parse operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_App *AppFilterer) ParseValueSet(log types.Log) (*AppValueSet, error) {
	event := new(AppValueSet)
	if err := _App.contract.UnpackLog(event, "ValueSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchValueSet is a free log subscription operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_App *AppSession) WatchValueSet(fromBlock *int64, handler func(int, []types.Log), key string) (string, error) {
	return _App.Contract.WatchValueSet(fromBlock, handler, key)
}

func (_App *AppSession) WatchAllValueSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _App.Contract.WatchAllValueSet(fromBlock, handler)
}

// ParseValueSet is a log parse operation binding the contract event 0xc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5.
//
// Solidity: event ValueSet(string indexed key, string value)
func (_App *AppSession) ParseValueSet(log types.Log) (*AppValueSet, error) {
	return _App.Contract.ParseValueSet(log)
}
