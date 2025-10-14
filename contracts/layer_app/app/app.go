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
const AppABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ValueSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"receive_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"}],\"name\":\"send_msg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AppBin is the compiled bytecode used for deploying new contracts.
var AppBin = "0x608060405234801561001057600080fd5b50611a7f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063528c15781461005c578063634e5a981461008d578063693ec85e146100be578063a7cd3fed146100ee578063e942b5161461011f575b600080fd5b61007660048036038101906100719190611221565b61013b565b60405161008492919061160a565b60405180910390f35b6100a760048036038101906100a29190611641565b610199565b6040516100b592919061160a565b60405180910390f35b6100d860048036038101906100d391906116e5565b6101bd565b6040516100e59190611732565b60405180910390f35b61010860048036038101906101039190611641565b610290565b60405161011692919061160a565b60405180910390f35b61013960048036038101906101349190611754565b610a82565b005b610143610b29565b606060018460e00181815250506001846101200181815250506001846101600181815250506001846101a0018181525050828460a001819052508360405180602001604052806000815250915091509250929050565b6101a1610b29565b6060826040518060200160405280600081525091509150915091565b6060600083836040516020016101d4929190611805565b60405160208183030381529060405280519060200120905060008082815260200190815260200160002080546102099061184d565b80601f01602080910402602001604051908101604052809291908181526020018280546102359061184d565b80156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b505050505091505092915050565b610298610b29565b606060008360a001516000815181106102b4576102b361187f565b5b602002602001015190506000816000815181106102d4576102d361187f565b5b602001015160f81c60f81b90507f2800000000000000000000000000000000000000000000000000000000000000817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141561042a5760006002835161033b91906118dd565b67ffffffffffffffff81111561035457610353610d0f565b5b6040519080825280601f01601f1916602001820160405280156103865781602001600182028036833780820191505090505b5090506000600190505b6001845161039e91906118dd565b811015610424578381815181106103b8576103b761187f565b5b602001015160f81c60f81b826001836103d191906118dd565b815181106103e2576103e161187f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061041c90611911565b915050610390565b50809250505b600080600060405180602001604052806000815250905060405180602001604052806000815250915060405180602001604052806000815250925060008060005b8751811015610701577f2c000000000000000000000000000000000000000000000000000000000000008882815181106104a8576104a761187f565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614156106ee5760008314156105d05781816104f091906118dd565b67ffffffffffffffff81111561050957610508610d0f565b5b6040519080825280601f01601f19166020018201604052801561053b5781602001600182028036833780820191505090505b50955060008290505b818110156105ca5788818151811061055f5761055e61187f565b5b602001015160f81c60f81b87848361057791906118dd565b815181106105885761058761187f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806105c290611911565b915050610544565b506106c2565b60018314156106c15781816105e591906118dd565b67ffffffffffffffff8111156105fe576105fd610d0f565b5b6040519080825280601f01601f1916602001820160405280156106305781602001600182028036833780820191505090505b50945060008290505b818110156106bf578881815181106106545761065361187f565b5b602001015160f81c60f81b86848361066c91906118dd565b8151811061067d5761067c61187f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806106b790611911565b915050610639565b505b5b82806106cd90611911565b9350506001816106dd919061195a565b915060028311156106ed57610701565b5b80806106f990611911565b91505061046b565b506002821480156107125750865181105b156108015780875161072491906118dd565b67ffffffffffffffff81111561073d5761073c610d0f565b5b6040519080825280601f01601f19166020018201604052801561076f5781602001600182028036833780820191505090505b50925060008190505b87518110156107ff578781815181106107945761079361187f565b5b602001015160f81c60f81b8483836107ac91906118dd565b815181106107bd576107bc61187f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806107f790611911565b915050610778565b505b7fd2f67e6aeaad1ab7487a680eb9d3363a597afa7a3de33fa9bf3ae6edcb88435d858051906020012014156108e1576000849050600084905060008260405160200161084d91906119e1565b60405160208183030381529060405280519060200120905081600080838152602001908152602001600020908051906020019061088b929190610bbc565b508260405161089a91906119e1565b60405180910390207fc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f5836040516108d19190611732565b60405180910390a2505050610a61565b7f6817c00f03de8b5bd58d2016b59d251c13056b989171c5852949903bc043bc2785805190602001201415610a6057600084905060008160405160200161092891906119e1565b6040516020818303038152906040528051906020012090506000806000838152602001908152602001600020805461095f9061184d565b80601f016020809104026020016040519081016040528092919081815260200182805461098b9061184d565b80156109d85780601f106109ad576101008083540402835291602001916109d8565b820191906000526020600020905b8154815290600101906020018083116109bb57829003601f168201915b505050505090506000600167ffffffffffffffff8111156109fc576109fb610d0f565b5b604051908082528060200260200182016040528015610a2f57816020015b6060815260200190600190039081610a1a5790505b5090508181600081518110610a4757610a4661187f565b5b6020026020010181905250808e60c00181905250505050505b5b89604051806020016040528060008152509850985050505050505050915091565b60008484604051602001610a97929190611805565b60405160208183030381529060405280519060200120905082826000808481526020019081526020016000209190610ad0929190610c42565b508484604051610ae1929190611805565b60405180910390207fc4df459909ff9875f4edd846ae16975efbc51b1b2930ffe2144f25bd303948f58484604051610b1a929190611a25565b60405180910390a25050505050565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b828054610bc89061184d565b90600052602060002090601f016020900481019282610bea5760008555610c31565b82601f10610c0357805160ff1916838001178555610c31565b82800160010185558215610c31579182015b82811115610c30578251825591602001919060010190610c15565b5b509050610c3e9190610cc8565b5090565b828054610c4e9061184d565b90600052602060002090601f016020900481019282610c705760008555610cb7565b82601f10610c8957803560ff1916838001178555610cb7565b82800160010185558215610cb7579182015b82811115610cb6578235825591602001919060010190610c9b565b5b509050610cc49190610cc8565b5090565b5b80821115610ce1576000816000905550600101610cc9565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d4782610cfe565b810181811067ffffffffffffffff82111715610d6657610d65610d0f565b5b80604052505050565b6000610d79610ce5565b9050610d858282610d3e565b919050565b600080fd5b6000819050919050565b610da281610d8f565b8114610dad57600080fd5b50565b600081359050610dbf81610d99565b92915050565b600080fd5b600067ffffffffffffffff821115610de557610de4610d0f565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff821115610e1b57610e1a610d0f565b5b610e2482610cfe565b9050602081019050919050565b82818337600083830152505050565b6000610e53610e4e84610e00565b610d6f565b905082815260208101848484011115610e6f57610e6e610dfb565b5b610e7a848285610e31565b509392505050565b600082601f830112610e9757610e96610dc5565b5b8135610ea7848260208601610e40565b91505092915050565b6000610ec3610ebe84610dca565b610d6f565b90508083825260208201905060208402830185811115610ee657610ee5610df6565b5b835b81811015610f2d57803567ffffffffffffffff811115610f0b57610f0a610dc5565b5b808601610f188982610e82565b85526020850194505050602081019050610ee8565b5050509392505050565b600082601f830112610f4c57610f4b610dc5565b5b8135610f5c848260208601610eb0565b91505092915050565b6000819050919050565b610f7881610f65565b8114610f8357600080fd5b50565b600081359050610f9581610f6f565b92915050565b60008115159050919050565b610fb081610f9b565b8114610fbb57600080fd5b50565b600081359050610fcd81610fa7565b92915050565b60006102408284031215610fea57610fe9610cf9565b5b610ff5610240610d6f565b9050600061100584828501610db0565b600083015250602061101984828501610db0565b602083015250604061102d84828501610db0565b604083015250606061104184828501610db0565b606083015250608061105584828501610db0565b60808301525060a082013567ffffffffffffffff81111561107957611078610d8a565b5b61108584828501610f37565b60a08301525060c082013567ffffffffffffffff8111156110a9576110a8610d8a565b5b6110b584828501610f37565b60c08301525060e06110c984828501610db0565b60e08301525061010082013567ffffffffffffffff8111156110ee576110ed610d8a565b5b6110fa84828501610f37565b6101008301525061012061111084828501610db0565b6101208301525061014082013567ffffffffffffffff81111561113657611135610d8a565b5b61114284828501610f37565b6101408301525061016061115884828501610db0565b6101608301525061018082013567ffffffffffffffff81111561117e5761117d610d8a565b5b61118a84828501610f37565b610180830152506101a06111a084828501610db0565b6101a0830152506101c082013567ffffffffffffffff8111156111c6576111c5610d8a565b5b6111d284828501610f37565b6101c0830152506101e06111e884828501610f86565b6101e0830152506102006111fe84828501610f86565b6102008301525061022061121484828501610fbe565b6102208301525092915050565b6000806040838503121561123857611237610cef565b5b600083013567ffffffffffffffff81111561125657611255610cf4565b5b61126285828601610fd3565b925050602083013567ffffffffffffffff81111561128357611282610cf4565b5b61128f85828601610f37565b9150509250929050565b6112a281610d8f565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561130e5780820151818401526020810190506112f3565b8381111561131d576000848401525b50505050565b600061132e826112d4565b61133881856112df565b93506113488185602086016112f0565b61135181610cfe565b840191505092915050565b60006113688383611323565b905092915050565b6000602082019050919050565b6000611388826112a8565b61139281856112b3565b9350836020820285016113a4856112c4565b8060005b858110156113e057848403895281516113c1858261135c565b94506113cc83611370565b925060208a019950506001810190506113a8565b50829750879550505050505092915050565b6113fb81610f65565b82525050565b61140a81610f9b565b82525050565b6000610240830160008301516114296000860182611299565b50602083015161143c6020860182611299565b50604083015161144f6040860182611299565b5060608301516114626060860182611299565b5060808301516114756080860182611299565b5060a083015184820360a086015261148d828261137d565b91505060c083015184820360c08601526114a7828261137d565b91505060e08301516114bc60e0860182611299565b506101008301518482036101008601526114d6828261137d565b9150506101208301516114ed610120860182611299565b50610140830151848203610140860152611507828261137d565b91505061016083015161151e610160860182611299565b50610180830151848203610180860152611538828261137d565b9150506101a083015161154f6101a0860182611299565b506101c08301518482036101c0860152611569828261137d565b9150506101e08301516115806101e08601826113f2565b506102008301516115956102008601826113f2565b506102208301516115aa610220860182611401565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b60006115dc826115b5565b6115e681856115c0565b93506115f68185602086016112f0565b6115ff81610cfe565b840191505092915050565b600060408201905081810360008301526116248185611410565b9050818103602083015261163881846115d1565b90509392505050565b60006020828403121561165757611656610cef565b5b600082013567ffffffffffffffff81111561167557611674610cf4565b5b61168184828501610fd3565b91505092915050565b600080fd5b60008083601f8401126116a5576116a4610dc5565b5b8235905067ffffffffffffffff8111156116c2576116c161168a565b5b6020830191508360018202830111156116de576116dd610df6565b5b9250929050565b600080602083850312156116fc576116fb610cef565b5b600083013567ffffffffffffffff81111561171a57611719610cf4565b5b6117268582860161168f565b92509250509250929050565b6000602082019050818103600083015261174c81846115d1565b905092915050565b6000806000806040858703121561176e5761176d610cef565b5b600085013567ffffffffffffffff81111561178c5761178b610cf4565b5b6117988782880161168f565b9450945050602085013567ffffffffffffffff8111156117bb576117ba610cf4565b5b6117c78782880161168f565b925092505092959194509250565b600081905092915050565b60006117ec83856117d5565b93506117f9838584610e31565b82840190509392505050565b60006118128284866117e0565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061186557607f821691505b602082108114156118795761187861181e565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118e882610d8f565b91506118f383610d8f565b925082821015611906576119056118ae565b5b828203905092915050565b600061191c82610d8f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561194f5761194e6118ae565b5b600182019050919050565b600061196582610d8f565b915061197083610d8f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156119a5576119a46118ae565b5b828201905092915050565b60006119bb826115b5565b6119c581856117d5565b93506119d58185602086016112f0565b80840191505092915050565b60006119ed82846119b0565b915081905092915050565b6000611a0483856115c0565b9350611a11838584610e31565b611a1a83610cfe565b840190509392505050565b60006020820190508181036000830152611a408184866119f8565b9050939250505056fea26469706673582212205169d0e4853273bac2b04cd2e6335debe4175608a82f7d548f6810caf3f4515c64736f6c634300080b0033"
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
