// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eventlisten

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

// EventlistenABI is the input ABI used to generate the binding from.
const EventlistenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"phase\",\"type\":\"uint256\"}],\"name\":\"CmHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"CMsg\",\"type\":\"bytes[]\"}],\"name\":\"Test_sendOut\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GenCMsg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"acknowledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"assignStringToBytesArray\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emit_sendOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cmhash\",\"type\":\"bytes32\"}],\"name\":\"getAckCmByHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cmhash\",\"type\":\"bytes32\"}],\"name\":\"getReqCmByHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cmhash\",\"type\":\"bytes32\"}],\"name\":\"getRespCmByHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"receiveIn\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"response\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"ccMsg\",\"type\":\"tuple\"}],\"name\":\"sendOut\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcAppId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstAppId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadReq\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloadResp\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transmissionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transmissionPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"verificationTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"verificationPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"transportTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transportPayload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"hashReq\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashResp\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ack\",\"type\":\"bool\"}],\"internalType\":\"structTypes.CrosschainMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"test_acknowledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"test_query_ccAckMsgs\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"test_query_ccReqMsgs\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"test_query_ccRespMsgs\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"test_query_hashAckToCmIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"test_query_hashReqToCmIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"test_query_hashRespToCmIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"name\":\"test_query_seqToHashAck\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"name\":\"test_query_seqToHashReq\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"test_receiveIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"test_respones_err\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"test_response\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"test_sendOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EventlistenBin is the compiled bytecode used for deploying new contracts.
var EventlistenBin = "0x60806040523480156200001157600080fd5b506200001c620004a3565b6001819080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050190805190602001906200009892919062000536565b5060c0820151816006019080519060200190620000b792919062000536565b5060e08201518160070155610100820151816008019080519060200190620000e192919062000536565b50610120820151816009015561014082015181600a0190805190602001906200010c92919062000536565b5061016082015181600b015561018082015181600c0190805190602001906200013792919062000536565b506101a082015181600d01556101c082015181600e0190805190602001906200016292919062000536565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff02191690831515021790555050506002819080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050190805190602001906200021892919062000536565b5060c08201518160060190805190602001906200023792919062000536565b5060e082015181600701556101008201518160080190805190602001906200026192919062000536565b50610120820151816009015561014082015181600a0190805190602001906200028c92919062000536565b5061016082015181600b015561018082015181600c019080519060200190620002b792919062000536565b506101a082015181600d01556101c082015181600e019080519060200190620002e292919062000536565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff02191690831515021790555050506003819080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050190805190602001906200039892919062000536565b5060c0820151816006019080519060200190620003b792919062000536565b5060e08201518160070155610100820151816008019080519060200190620003e192919062000536565b50610120820151816009015561014082015181600a0190805190602001906200040c92919062000536565b5061016082015181600b015561018082015181600c0190805190602001906200043792919062000536565b506101a082015181600d01556101c082015181600e0190805190602001906200046292919062000536565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff02191690831515021790555050505062000720565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b8280548282559060005260206000209081019282156200058a579160200282015b8281111562000589578251829080519060200190620005789291906200059d565b509160200191906001019062000557565b5b5090506200059991906200062e565b5090565b828054620005ab90620006ea565b90600052602060002090601f016020900481019282620005cf57600085556200061b565b82601f10620005ea57805160ff19168380011785556200061b565b828001600101855582156200061b579182015b828111156200061a578251825591602001919060010190620005fd565b5b5090506200062a919062000656565b5090565b5b8082111562000652576000818162000648919062000675565b506001016200062f565b5090565b5b808211156200067157600081600090555060010162000657565b5090565b5080546200068390620006ea565b6000825580601f10620006975750620006b8565b601f016020900490600052602060002090810190620006b7919062000656565b5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200070357607f821691505b602082108114156200071a5762000719620006bb565b5b50919050565b615ea880620007306000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c8063957cffd2116100de578063ccb5ad2e11610097578063e952f03611610071578063e952f036146104f8578063f392a66e14610502578063f6b9f4b914610532578063f9120af61461054e57610173565b8063ccb5ad2e14610467578063d081dd0c14610497578063e7ef5f9b146104c757610173565b8063957cffd21461035957806396070c01146103755780639d36e127146103a6578063a6261dca146103d6578063b552adeb14610406578063bff894a31461043657610173565b80636b640869116101305780636b6408691461026157806373dae718146102915780637c822f5c146102ad5780637e0694c2146102dd578063891047471461030d57806393e00b011461032957610173565b80630470439a146101785780630884e212146101965780632369d54c146101c657806332890a6c146101e25780633ad59dbc146102135780635dde928314610231575b600080fd5b61018061056a565b60405161018d9190615174565b60405180910390f35b6101b060048036038101906101ab91906151d6565b6107a1565b6040516101bd9190615289565b60405180910390f35b6101e060048036038101906101db91906154c6565b610de2565b005b6101fc60048036038101906101f791906154c6565b610e46565b60405161020a929190615564565b60405180910390f35b61021b6110f8565b60405161022891906155dc565b60405180910390f35b61024b600480360381019061024691906151d6565b611121565b6040516102589190615289565b60405180910390f35b61027b60048036038101906102769190615623565b611762565b604051610288919061565f565b60405180910390f35b6102ab60048036038101906102a691906154c6565b61177f565b005b6102c760048036038101906102c29190615623565b6117f1565b6040516102d49190615174565b60405180910390f35b6102f760048036038101906102f291906151d6565b611df2565b6040516103049190615689565b60405180910390f35b610327600480360381019061032291906154c6565b611e0f565b005b610343600480360381019061033e91906151d6565b611ecf565b6040516103509190615689565b60405180910390f35b610373600480360381019061036e91906154c6565b611eec565b005b61038f600480360381019061038a9190615928565b611f50565b60405161039d929190615564565b60405180910390f35b6103c060048036038101906103bb9190615623565b612230565b6040516103cd9190615174565b60405180910390f35b6103f060048036038101906103eb9190615623565b612831565b6040516103fd919061565f565b60405180910390f35b610420600480360381019061041b9190615623565b61284e565b60405161042d919061565f565b60405180910390f35b610450600480360381019061044b9190615928565b61286b565b60405161045e929190615564565b60405180910390f35b610481600480360381019061047c9190615623565b612ade565b60405161048e9190615174565b60405180910390f35b6104b160048036038101906104ac9190615a12565b6130df565b6040516104be9190615289565b60405180910390f35b6104e160048036038101906104dc91906154c6565b61315d565b6040516104ef929190615564565b60405180910390f35b610500613482565b005b61051c600480360381019061051791906151d6565b6134d0565b6040516105299190615289565b60405180910390f35b61054c600480360381019061054791906154c6565b613b11565b005b61056860048036038101906105639190615a87565b613b83565b005b610572614c2d565b61057a614c2d565b6001816101a00181815250506105c46040518060400160405280601081526020017f7472616e73706f72745061796c6f6164000000000000000000000000000000008152506130df565b816101c001819052506001816101600181815250506106176040518060400160405280601381526020017f766572696669636174696f6e5061796c6f6164000000000000000000000000008152506130df565b81610180018190525060018161012001818152505061066a6040518060400160405280601381526020017f7472616e736d697373696f6e5061796c6f6164000000000000000000000000008152506130df565b81610140018190525060018160e00181815250506106bc6040518060400160405280601281526020017f7472616e73616374696f6e5061796c6f616400000000000000000000000000008152506130df565b816101000181905250600181606001818152505060018160800181815250506107196040518060400160405280600a81526020017f7061796c6f6164526571000000000000000000000000000000000000000000008152506130df565b8160a0018190525061075f6040518060400160405280600b81526020017f7061796c6f6164526573700000000000000000000000000000000000000000008152506130df565b8160c001819052506001816000018181525050600281602001818152505060018160400181815250506000816102200190151590811515815250508091505090565b6060600180549050821061080557600067ffffffffffffffff8111156107ca576107c96152b0565b5b6040519080825280602002602001820160405280156107fd57816020015b60608152602001906001900390816107e85790505b509050610ddd565b610dda6001838154811061081c5761081b615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156109395783829060005260206000200180546108ac90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546108d890615b12565b80156109255780601f106108fa57610100808354040283529160200191610925565b820191906000526020600020905b81548152906001019060200180831161090857829003601f168201915b50505050508152602001906001019061088d565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b82821015610a1257838290600052602060002001805461098590615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546109b190615b12565b80156109fe5780601f106109d3576101008083540402835291602001916109fe565b820191906000526020600020905b8154815290600101906020018083116109e157829003601f168201915b505050505081526020019060010190610966565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b82821015610af5578382906000526020600020018054610a6890615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9490615b12565b8015610ae15780601f10610ab657610100808354040283529160200191610ae1565b820191906000526020600020905b815481529060010190602001808311610ac457829003601f168201915b505050505081526020019060010190610a49565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b82821015610bd8578382906000526020600020018054610b4b90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7790615b12565b8015610bc45780601f10610b9957610100808354040283529160200191610bc4565b820191906000526020600020905b815481529060010190602001808311610ba757829003601f168201915b505050505081526020019060010190610b2c565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b82821015610cbb578382906000526020600020018054610c2e90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5a90615b12565b8015610ca75780601f10610c7c57610100808354040283529160200191610ca7565b820191906000526020600020905b815481529060010190602001808311610c8a57829003601f168201915b505050505081526020019060010190610c0f565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b82821015610d9e578382906000526020600020018054610d1190615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3d90615b12565b8015610d8a5780601f10610d5f57610100808354040283529160200191610d8a565b820191906000526020600020905b815481529060010190602001808311610d6d57829003601f168201915b505050505081526020019060010190610cf2565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff161515151581525050613bc6565b90505b919050565b6060610dec614c2d565b610df583610e46565b80935081925050506000825114610e41576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3183604051610e389190615b44565b60405180910390a25b505050565b610e4e614c2d565b6060610e58614c2d565b6000610e626110f8565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ebc5781604051806060016040528060238152602001615dfa602391399350935050506110f3565b610ec5856142ff565b91506000610ed28361479c565b90506000600460008381526020019081526020016000205414610f3257826040518060400160405280601f81526020017f696e2072656365697665496e3a2063634d736720686173206578697374656400815250945094505050506110f3565b80836101e00181815250506002839080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005019080519060200190610fb7929190614cc0565b5060c0820151816006019080519060200190610fd4929190614cc0565b5060e08201518160070155610100820151816008019080519060200190610ffc929190614cc0565b50610120820151816009015561014082015181600a019080519060200190611025929190614cc0565b5061016082015181600b015561018082015181600c01908051906020019061104e929190614cc0565b506101a082015181600d01556101c082015181600e019080519060200190611077929190614cc0565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff021916908315150217905550505060016002805490506110c39190615b95565b60046000838152602001908152602001600020819055508260405180602001604052806000815250945094505050505b915091565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600280549050821061118557600067ffffffffffffffff81111561114a576111496152b0565b5b60405190808252806020026020018201604052801561117d57816020015b60608152602001906001900390816111685790505b50905061175d565b61175a6002838154811061119c5761119b615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156112b957838290600052602060002001805461122c90615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461125890615b12565b80156112a55780601f1061127a576101008083540402835291602001916112a5565b820191906000526020600020905b81548152906001019060200180831161128857829003601f168201915b50505050508152602001906001019061120d565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b8282101561139257838290600052602060002001805461130590615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461133190615b12565b801561137e5780601f106113535761010080835404028352916020019161137e565b820191906000526020600020905b81548152906001019060200180831161136157829003601f168201915b5050505050815260200190600101906112e6565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b828210156114755783829060005260206000200180546113e890615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461141490615b12565b80156114615780601f1061143657610100808354040283529160200191611461565b820191906000526020600020905b81548152906001019060200180831161144457829003601f168201915b5050505050815260200190600101906113c9565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b828210156115585783829060005260206000200180546114cb90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546114f790615b12565b80156115445780601f1061151957610100808354040283529160200191611544565b820191906000526020600020905b81548152906001019060200180831161152757829003601f168201915b5050505050815260200190600101906114ac565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b8282101561163b5783829060005260206000200180546115ae90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546115da90615b12565b80156116275780601f106115fc57610100808354040283529160200191611627565b820191906000526020600020905b81548152906001019060200180831161160a57829003601f168201915b50505050508152602001906001019061158f565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b8282101561171e57838290600052602060002001805461169190615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546116bd90615b12565b801561170a5780601f106116df5761010080835404028352916020019161170a565b820191906000526020600020905b8154815290600101906020018083116116ed57829003601f168201915b505050505081526020019060010190611672565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff161515151581525050613bc6565b90505b919050565b600060056000838152602001908152602001600020549050919050565b6060600061178c836142ff565b9050611796614c2d565b61179f8261286b565b809450819250505060008351146117eb576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc31846040516117e29190615b44565b60405180910390a25b50505050565b6117f9614c2d565b600060066000848152602001908152602001600020549050611819614c2d565b6003828154811061182d5761182c615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b8282101561194a5783829060005260206000200180546118bd90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546118e990615b12565b80156119365780601f1061190b57610100808354040283529160200191611936565b820191906000526020600020905b81548152906001019060200180831161191957829003601f168201915b50505050508152602001906001019061189e565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b82821015611a2357838290600052602060002001805461199690615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546119c290615b12565b8015611a0f5780601f106119e457610100808354040283529160200191611a0f565b820191906000526020600020905b8154815290600101906020018083116119f257829003601f168201915b505050505081526020019060010190611977565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b82821015611b06578382906000526020600020018054611a7990615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054611aa590615b12565b8015611af25780601f10611ac757610100808354040283529160200191611af2565b820191906000526020600020905b815481529060010190602001808311611ad557829003601f168201915b505050505081526020019060010190611a5a565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b82821015611be9578382906000526020600020018054611b5c90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054611b8890615b12565b8015611bd55780601f10611baa57610100808354040283529160200191611bd5565b820191906000526020600020905b815481529060010190602001808311611bb857829003601f168201915b505050505081526020019060010190611b3d565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b82821015611ccc578382906000526020600020018054611c3f90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054611c6b90615b12565b8015611cb85780601f10611c8d57610100808354040283529160200191611cb8565b820191906000526020600020905b815481529060010190602001808311611c9b57829003601f168201915b505050505081526020019060010190611c20565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b82821015611daf578382906000526020600020018054611d2290615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054611d4e90615b12565b8015611d9b5780601f10611d7057610100808354040283529160200191611d9b565b820191906000526020600020905b815481529060010190602001808311611d7e57829003601f168201915b505050505081526020019060010190611d03565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff16151515158152505090508092505050919050565b600060076000838152602001908152602001600020549050919050565b6060611e19614c2d565b611e2283610e46565b80935081925050506000825114611e74576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3183604051611e659190615b44565b60405180910390a25050611ecc565b611e7d8161286b565b80935081925050506000825114611ec9576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3183604051611ec09190615b44565b60405180910390a25b50505b50565b600060086000838152602001908152602001600020549050919050565b6060611ef6614c2d565b611eff8361315d565b80935081925050506000825114611f4b576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3183604051611f429190615b44565b60405180910390a25b505050565b611f58614c2d565b60606000611f646110f8565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611fbd5783604051806060016040528060238152602001615dfa60239139925092505061222b565b6000801b6007600086604001518152602001908152602001600020541461201f57836040518060400160405280601881526020017f6578697374207468652073616d652063634d73672e7365710000000000000000815250925092505061222b565b600061202a8561479c565b90506001859080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050190805190602001906120a6929190614cc0565b5060c08201518160060190805190602001906120c3929190614cc0565b5060e082015181600701556101008201518160080190805190602001906120eb929190614cc0565b50610120820151816009015561014082015181600a019080519060200190612114929190614cc0565b5061016082015181600b015561018082015181600c01908051906020019061213d929190614cc0565b506101a082015181600d01556101c082015181600e019080519060200190612166929190614cc0565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff0219169083151502179055505050600180805490506121b19190615b95565b6004600083815260200190815260200160002081905550806007600087604001518152602001908152602001600020819055506001817f61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e60405160405180910390a384604051806020016040528060008152509350935050505b915091565b612238614c2d565b600060046000848152602001908152602001600020549050612258614c2d565b6001828154811061226c5761226b615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156123895783829060005260206000200180546122fc90615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461232890615b12565b80156123755780601f1061234a57610100808354040283529160200191612375565b820191906000526020600020905b81548152906001019060200180831161235857829003601f168201915b5050505050815260200190600101906122dd565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b828210156124625783829060005260206000200180546123d590615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461240190615b12565b801561244e5780601f106124235761010080835404028352916020019161244e565b820191906000526020600020905b81548152906001019060200180831161243157829003601f168201915b5050505050815260200190600101906123b6565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b828210156125455783829060005260206000200180546124b890615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546124e490615b12565b80156125315780601f1061250657610100808354040283529160200191612531565b820191906000526020600020905b81548152906001019060200180831161251457829003601f168201915b505050505081526020019060010190612499565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b8282101561262857838290600052602060002001805461259b90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546125c790615b12565b80156126145780601f106125e957610100808354040283529160200191612614565b820191906000526020600020905b8154815290600101906020018083116125f757829003601f168201915b50505050508152602001906001019061257c565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b8282101561270b57838290600052602060002001805461267e90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546126aa90615b12565b80156126f75780601f106126cc576101008083540402835291602001916126f7565b820191906000526020600020905b8154815290600101906020018083116126da57829003601f168201915b50505050508152602001906001019061265f565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b828210156127ee57838290600052602060002001805461276190615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461278d90615b12565b80156127da5780601f106127af576101008083540402835291602001916127da565b820191906000526020600020905b8154815290600101906020018083116127bd57829003601f168201915b505050505081526020019060010190612742565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff16151515158152505090508092505050919050565b600060046000838152602001908152602001600020549050919050565b600060066000838152602001908152602001600020549050919050565b612873614c2d565b6060600061287f6110f8565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146128d85783604051806060016040528060238152602001615dfa602391399250925050612ad9565b60006128e38561479c565b90508085610200018181525050600060046000876101e001518152602001908152602001600020549050806005600084815260200190815260200160002081905550856002828154811061293a57612939615ab4565b5b9060005260206000209060120201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005019080519060200190612996929190614cc0565b5060c08201518160060190805190602001906129b3929190614cc0565b5060e082015181600701556101008201518160080190805190602001906129db929190614cc0565b50610120820151816009015561014082015181600a019080519060200190612a04929190614cc0565b5061016082015181600b015561018082015181600c019080519060200190612a2d929190614cc0565b506101a082015181600d01556101c082015181600e019080519060200190612a56929190614cc0565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff0219169083151502179055509050506002827f61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e60405160405180910390a38560405180602001604052806000815250945094505050505b915091565b612ae6614c2d565b600060056000848152602001908152602001600020549050612b06614c2d565b60028281548110612b1a57612b19615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b82821015612c37578382906000526020600020018054612baa90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054612bd690615b12565b8015612c235780601f10612bf857610100808354040283529160200191612c23565b820191906000526020600020905b815481529060010190602001808311612c0657829003601f168201915b505050505081526020019060010190612b8b565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b82821015612d10578382906000526020600020018054612c8390615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054612caf90615b12565b8015612cfc5780601f10612cd157610100808354040283529160200191612cfc565b820191906000526020600020905b815481529060010190602001808311612cdf57829003601f168201915b505050505081526020019060010190612c64565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b82821015612df3578382906000526020600020018054612d6690615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054612d9290615b12565b8015612ddf5780601f10612db457610100808354040283529160200191612ddf565b820191906000526020600020905b815481529060010190602001808311612dc257829003601f168201915b505050505081526020019060010190612d47565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b82821015612ed6578382906000526020600020018054612e4990615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054612e7590615b12565b8015612ec25780601f10612e9757610100808354040283529160200191612ec2565b820191906000526020600020905b815481529060010190602001808311612ea557829003601f168201915b505050505081526020019060010190612e2a565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b82821015612fb9578382906000526020600020018054612f2c90615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054612f5890615b12565b8015612fa55780601f10612f7a57610100808354040283529160200191612fa5565b820191906000526020600020905b815481529060010190602001808311612f8857829003601f168201915b505050505081526020019060010190612f0d565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b8282101561309c57838290600052602060002001805461300f90615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461303b90615b12565b80156130885780601f1061305d57610100808354040283529160200191613088565b820191906000526020600020905b81548152906001019060200180831161306b57829003601f168201915b505050505081526020019060010190612ff0565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff16151515158152505090508092505050919050565b60606000600167ffffffffffffffff8111156130fe576130fd6152b0565b5b60405190808252806020026020018201604052801561313157816020015b606081526020019060019003908161311c5790505b509050828160008151811061314957613148615ab4565b5b602002602001018190525080915050919050565b613165614c2d565b606061316f614c2d565b60006131796110f8565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146131d35781604051806060016040528060238152602001615dfa6023913993509350505061347d565b6131dc856142ff565b91506000826040015190506000801b60086000838152602001908152602001600020541461322a5782604051806060016040528060258152602001615e4e602591399450945050505061347d565b6000801b6007600083815260200190815260200160002054141561326e5782604051806060016040528060318152602001615e1d603191399450945050505061347d565b60006132798461479c565b90506003849080600181540180825580915050600190039060005260206000209060120201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050190805190602001906132f5929190614cc0565b5060c0820151816006019080519060200190613312929190614cc0565b5060e0820151816007015561010082015181600801908051906020019061333a929190614cc0565b50610120820151816009015561014082015181600a019080519060200190613363929190614cc0565b5061016082015181600b015561018082015181600c01908051906020019061338c929190614cc0565b506101a082015181600d01556101c082015181600e0190805190602001906133b5929190614cc0565b506101e082015181600f015561020082015181601001556102208201518160110160006101000a81548160ff021916908315150217905550505060016003805490506134019190615b95565b6006600083815260200190815260200160002081905550806008600086604001518152602001908152602001600020819055506003817f61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e60405160405180910390a3836040518060200160405280600081525095509550505050505b915091565b600061348c61056a565b90507f1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e6134b882613bc6565b6040516134c59190615289565b60405180910390a150565b6060600380549050821061353457600067ffffffffffffffff8111156134f9576134f86152b0565b5b60405190808252806020026020018201604052801561352c57816020015b60608152602001906001900390816135175790505b509050613b0c565b613b096003838154811061354b5761354a615ab4565b5b906000526020600020906012020160405180610240016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156136685783829060005260206000200180546135db90615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461360790615b12565b80156136545780601f1061362957610100808354040283529160200191613654565b820191906000526020600020905b81548152906001019060200180831161363757829003601f168201915b5050505050815260200190600101906135bc565b50505050815260200160068201805480602002602001604051908101604052809291908181526020016000905b828210156137415783829060005260206000200180546136b490615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546136e090615b12565b801561372d5780601f106137025761010080835404028352916020019161372d565b820191906000526020600020905b81548152906001019060200180831161371057829003601f168201915b505050505081526020019060010190613695565b5050505081526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b8282101561382457838290600052602060002001805461379790615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546137c390615b12565b80156138105780601f106137e557610100808354040283529160200191613810565b820191906000526020600020905b8154815290600101906020018083116137f357829003601f168201915b505050505081526020019060010190613778565b50505050815260200160098201548152602001600a8201805480602002602001604051908101604052809291908181526020016000905b8282101561390757838290600052602060002001805461387a90615b12565b80601f01602080910402602001604051908101604052809291908181526020018280546138a690615b12565b80156138f35780601f106138c8576101008083540402835291602001916138f3565b820191906000526020600020905b8154815290600101906020018083116138d657829003601f168201915b50505050508152602001906001019061385b565b505050508152602001600b8201548152602001600c8201805480602002602001604051908101604052809291908181526020016000905b828210156139ea57838290600052602060002001805461395d90615b12565b80601f016020809104026020016040519081016040528092919081815260200182805461398990615b12565b80156139d65780601f106139ab576101008083540402835291602001916139d6565b820191906000526020600020905b8154815290600101906020018083116139b957829003601f168201915b50505050508152602001906001019061393e565b505050508152602001600d8201548152602001600e8201805480602002602001604051908101604052809291908181526020016000905b82821015613acd578382906000526020600020018054613a4090615b12565b80601f0160208091040260200160405190810160405280929190818152602001828054613a6c90615b12565b8015613ab95780601f10613a8e57610100808354040283529160200191613ab9565b820191906000526020600020905b815481529060010190602001808311613a9c57829003601f168201915b505050505081526020019060010190613a21565b505050508152602001600f8201548152602001601082015481526020016011820160009054906101000a900460ff161515151581525050613bc6565b90505b919050565b60606000613b1e836142ff565b9050613b28614c2d565b613b3182611f50565b80945081925050506000835114613b7d576127117fc9dc6ff1a96f4ac114b02968275a8ab4e677d894f261a74987fb58ebc283cc3184604051613b749190615b44565b60405180910390a25b50505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606060008260c00151518360a0015151846101000151518561014001515186610180015151876101c00151516010613bfe9190615bc9565b613c089190615bc9565b613c129190615bc9565b613c1c9190615bc9565b613c269190615bc9565b613c309190615bc9565b905060008167ffffffffffffffff811115613c4e57613c4d6152b0565b5b604051908082528060200260200182016040528015613c8157816020015b6060815260200190600190039081613c6c5790505b5090506000613c94856101a00151614863565b828280613ca090615c1f565b935081518110613cb357613cb2615ab4565b5b6020026020010181905250613cd0856101c001515160001b61487e565b828280613cdc90615c1f565b935081518110613cef57613cee615ab4565b5b602002602001018190525060005b856101c0015151811015613d7457856101c001518180613d1c90615c1f565b925081518110613d2f57613d2e615ab4565b5b6020026020010151838380613d4390615c1f565b945081518110613d5657613d55615ab4565b5b60200260200101819052508080613d6c90615c1f565b915050613cfd565b50613d83856101600151614863565b828280613d8f90615c1f565b935081518110613da257613da1615ab4565b5b6020026020010181905250613dbf8561018001515160001b61487e565b828280613dcb90615c1f565b935081518110613dde57613ddd615ab4565b5b602002602001018190525060005b85610180015151811015613e63578561018001518180613e0b90615c1f565b925081518110613e1e57613e1d615ab4565b5b6020026020010151838380613e3290615c1f565b945081518110613e4557613e44615ab4565b5b60200260200101819052508080613e5b90615c1f565b915050613dec565b50613e72856101200151614863565b828280613e7e90615c1f565b935081518110613e9157613e90615ab4565b5b6020026020010181905250613eae8561014001515160001b61487e565b828280613eba90615c1f565b935081518110613ecd57613ecc615ab4565b5b602002602001018190525060005b85610140015151811015613f52578561014001518180613efa90615c1f565b925081518110613f0d57613f0c615ab4565b5b6020026020010151838380613f2190615c1f565b945081518110613f3457613f33615ab4565b5b60200260200101819052508080613f4a90615c1f565b915050613edb565b50613f608560e00151614863565b828280613f6c90615c1f565b935081518110613f7f57613f7e615ab4565b5b6020026020010181905250613f9c8561010001515160001b61487e565b828280613fa890615c1f565b935081518110613fbb57613fba615ab4565b5b602002602001018190525060005b85610100015151811015614040578561010001518180613fe890615c1f565b925081518110613ffb57613ffa615ab4565b5b602002602001015183838061400f90615c1f565b94508151811061402257614021615ab4565b5b6020026020010181905250808061403890615c1f565b915050613fc9565b5061404e8560600151614863565b82828061405a90615c1f565b93508151811061406d5761406c615ab4565b5b60200260200101819052506140858560800151614863565b82828061409190615c1f565b9350815181106140a4576140a3615ab4565b5b60200260200101819052506140c08560a001515160001b61487e565b8282806140cc90615c1f565b9350815181106140df576140de615ab4565b5b602002602001018190525060005b8560a0015151811015614162578560a00151818061410a90615c1f565b92508151811061411d5761411c615ab4565b5b602002602001015183838061413190615c1f565b94508151811061414457614143615ab4565b5b6020026020010181905250808061415a90615c1f565b9150506140ed565b506141748560c001515160001b61487e565b82828061418090615c1f565b93508151811061419357614192615ab4565b5b602002602001018190525060005b8560c0015151811015614216578560c0015181806141be90615c1f565b9250815181106141d1576141d0615ab4565b5b60200260200101518383806141e590615c1f565b9450815181106141f8576141f7615ab4565b5b6020026020010181905250808061420e90615c1f565b9150506141a1565b506142248560000151614863565b82828061423090615c1f565b93508151811061424357614242615ab4565b5b602002602001018190525061425b8560200151614863565b82828061426790615c1f565b93508151811061427a57614279615ab4565b5b60200260200101819052506142928560400151614863565b82828061429e90615c1f565b9350815181106142b1576142b0615ab4565b5b60200260200101819052506142ca856102200151614954565b8282806142d690615c1f565b9350815181106142e9576142e8615ab4565b5b6020026020010181905250819350505050919050565b614307614c2d565b61430f614c2d565b60006143358460008151811061432857614327615ab4565b5b6020026020010151614a03565b826101a001818152505061434a846001614aa4565b826101c001819052506143778460018151811061436a57614369615ab4565b5b6020026020010151614a03565b6001826143849190615bc9565b61438e9190615bc9565b90506143bf846001836143a19190615bc9565b815181106143b2576143b1615ab4565b5b6020026020010151614a03565b826101600181815250506143df846002836143da9190615bc9565b614aa4565b826101800181905250614417846002836143f99190615bc9565b8151811061440a57614409615ab4565b5b6020026020010151614a03565b6002826144249190615bc9565b61442e9190615bc9565b905061445f846001836144419190615bc9565b8151811061445257614451615ab4565b5b6020026020010151614a03565b8261012001818152505061447f8460028361447a9190615bc9565b614aa4565b8261014001819052506144b7846002836144999190615bc9565b815181106144aa576144a9615ab4565b5b6020026020010151614a03565b6002826144c49190615bc9565b6144ce9190615bc9565b90506144ff846001836144e19190615bc9565b815181106144f2576144f1615ab4565b5b6020026020010151614a03565b8260e001818152505061451e846002836145199190615bc9565b614aa4565b826101000181905250614556846002836145389190615bc9565b8151811061454957614548615ab4565b5b6020026020010151614a03565b6002826145639190615bc9565b61456d9190615bc9565b905061459e846001836145809190615bc9565b8151811061459157614590615ab4565b5b6020026020010151614a03565b8260600181815250506145d6846002836145b89190615bc9565b815181106145c9576145c8615ab4565b5b6020026020010151614a03565b8260800181815250506145f5846003836145f09190615bc9565b614aa4565b8260a0018190525061462c8460038361460e9190615bc9565b8151811061461f5761461e615ab4565b5b6020026020010151614a03565b6003826146399190615bc9565b6146439190615bc9565b905061465b846001836146569190615bc9565b614aa4565b8260c00181905250614692846001836146749190615bc9565b8151811061468557614684615ab4565b5b6020026020010151614a03565b60018261469f9190615bc9565b6146a99190615bc9565b90506146da846001836146bc9190615bc9565b815181106146cd576146cc615ab4565b5b6020026020010151614a03565b826000018181525050614712846002836146f49190615bc9565b8151811061470557614704615ab4565b5b6020026020010151614a03565b82602001818152505061474a8460038361472c9190615bc9565b8151811061473d5761473c615ab4565b5b6020026020010151614a03565b826040018181525050614782846004836147649190615bc9565b8151811061477557614774615ab4565b5b6020026020010151614b99565b826102200190151590811515815250508192505050919050565b6000806147a883613bc6565b9050606060005b825181101561480957818382815181106147cc576147cb615ab4565b5b60200260200101516040516020016147e5929190615ca4565b6040516020818303038152906040529150808061480190615c1f565b9150506147af565b5060028160405161481a9190615cc8565b602060405180830381855afa158015614837573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061485a9190615cf4565b92505050919050565b606060008260001b90506148768161487e565b915050919050565b60606000602067ffffffffffffffff81111561489d5761489c6152b0565b5b6040519080825280601f01601f1916602001820160405280156148cf5781602001600182028036833780820191505090505b50905060005b602081101561494a578381602081106148f1576148f0615ab4565b5b1a60f81b82828151811061490857614907615ab4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061494290615c1f565b9150506148d5565b5080915050919050565b6060600167ffffffffffffffff811115614971576149706152b0565b5b6040519080825280601f01601f1916602001820160405280156149a35781602001600182028036833780820191505090505b509050816149b557600060f81b6149bb565b600160f81b5b816000815181106149cf576149ce615ab4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350919050565b6000602082511015614a4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a4190615d6d565b60405180910390fd5b600080600090505b6020811015614a9a57838181518110614a6e57614a6d615ab4565b5b602001015160f81c60f81b60f81c60ff16600883901b1791508080614a9290615c1f565b915050614a52565b5080915050919050565b60606000614acb848481518110614abe57614abd615ab4565b5b6020026020010151614a03565b905060008167ffffffffffffffff811115614ae957614ae86152b0565b5b604051908082528060200260200182016040528015614b1c57816020015b6060815260200190600190039081614b075790505b50905060005b82811015614b8d578581600187614b399190615bc9565b614b439190615bc9565b81518110614b5457614b53615ab4565b5b6020026020010151828281518110614b6f57614b6e615ab4565b5b60200260200101819052508080614b8590615c1f565b915050614b22565b50809250505092915050565b600080825111614bde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614bd590615dd9565b60405180910390fd5b600060f81b82600081518110614bf757614bf6615ab4565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614159050919050565b60405180610240016040528060008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016060815260200160008152602001606081526020016000815260200160608152602001600081526020016060815260200160008019168152602001600080191681526020016000151581525090565b828054828255906000526020600020908101928215614d0f579160200282015b82811115614d0e578251829080519060200190614cfe929190614d20565b5091602001919060010190614ce0565b5b509050614d1c9190614da6565b5090565b828054614d2c90615b12565b90600052602060002090601f016020900481019282614d4e5760008555614d95565b82601f10614d6757805160ff1916838001178555614d95565b82800160010185558215614d95579182015b82811115614d94578251825591602001919060010190614d79565b5b509050614da29190614dca565b5090565b5b80821115614dc65760008181614dbd9190614de7565b50600101614da7565b5090565b5b80821115614de3576000816000905550600101614dcb565b5090565b508054614df390615b12565b6000825580601f10614e055750614e24565b601f016020900490600052602060002090810190614e239190614dca565b5b50565b6000819050919050565b614e3a81614e27565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614ea6578082015181840152602081019050614e8b565b83811115614eb5576000848401525b50505050565b6000601f19601f8301169050919050565b6000614ed782614e6c565b614ee18185614e77565b9350614ef1818560208601614e88565b614efa81614ebb565b840191505092915050565b6000614f118383614ecc565b905092915050565b6000602082019050919050565b6000614f3182614e40565b614f3b8185614e4b565b935083602082028501614f4d85614e5c565b8060005b85811015614f895784840389528151614f6a8582614f05565b9450614f7583614f19565b925060208a01995050600181019050614f51565b50829750879550505050505092915050565b6000819050919050565b614fae81614f9b565b82525050565b60008115159050919050565b614fc981614fb4565b82525050565b600061024083016000830151614fe86000860182614e31565b506020830151614ffb6020860182614e31565b50604083015161500e6040860182614e31565b5060608301516150216060860182614e31565b5060808301516150346080860182614e31565b5060a083015184820360a086015261504c8282614f26565b91505060c083015184820360c08601526150668282614f26565b91505060e083015161507b60e0860182614e31565b506101008301518482036101008601526150958282614f26565b9150506101208301516150ac610120860182614e31565b506101408301518482036101408601526150c68282614f26565b9150506101608301516150dd610160860182614e31565b506101808301518482036101808601526150f78282614f26565b9150506101a083015161510e6101a0860182614e31565b506101c08301518482036101c08601526151288282614f26565b9150506101e083015161513f6101e0860182614fa5565b50610200830151615154610200860182614fa5565b50610220830151615169610220860182614fc0565b508091505092915050565b6000602082019050818103600083015261518e8184614fcf565b905092915050565b6000604051905090565b600080fd5b600080fd5b6151b381614e27565b81146151be57600080fd5b50565b6000813590506151d0816151aa565b92915050565b6000602082840312156151ec576151eb6151a0565b5b60006151fa848285016151c1565b91505092915050565b600082825260208201905092915050565b600061521f82614e40565b6152298185615203565b93508360208202850161523b85614e5c565b8060005b8581101561527757848403895281516152588582614f05565b945061526383614f19565b925060208a0199505060018101905061523f565b50829750879550505050505092915050565b600060208201905081810360008301526152a38184615214565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6152e882614ebb565b810181811067ffffffffffffffff82111715615307576153066152b0565b5b80604052505050565b600061531a615196565b905061532682826152df565b919050565b600067ffffffffffffffff821115615346576153456152b0565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561537c5761537b6152b0565b5b61538582614ebb565b9050602081019050919050565b82818337600083830152505050565b60006153b46153af84615361565b615310565b9050828152602081018484840111156153d0576153cf61535c565b5b6153db848285615392565b509392505050565b600082601f8301126153f8576153f76152ab565b5b81356154088482602086016153a1565b91505092915050565b600061542461541f8461532b565b615310565b9050808382526020820190506020840283018581111561544757615446615357565b5b835b8181101561548e57803567ffffffffffffffff81111561546c5761546b6152ab565b5b80860161547989826153e3565b85526020850194505050602081019050615449565b5050509392505050565b600082601f8301126154ad576154ac6152ab565b5b81356154bd848260208601615411565b91505092915050565b6000602082840312156154dc576154db6151a0565b5b600082013567ffffffffffffffff8111156154fa576154f96151a5565b5b61550684828501615498565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60006155368261550f565b615540818561551a565b9350615550818560208601614e88565b61555981614ebb565b840191505092915050565b6000604082019050818103600083015261557e8185614fcf565b90508181036020830152615592818461552b565b90509392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006155c68261559b565b9050919050565b6155d6816155bb565b82525050565b60006020820190506155f160008301846155cd565b92915050565b61560081614f9b565b811461560b57600080fd5b50565b60008135905061561d816155f7565b92915050565b600060208284031215615639576156386151a0565b5b60006156478482850161560e565b91505092915050565b61565981614e27565b82525050565b60006020820190506156746000830184615650565b92915050565b61568381614f9b565b82525050565b600060208201905061569e600083018461567a565b92915050565b600080fd5b600080fd5b6156b781614fb4565b81146156c257600080fd5b50565b6000813590506156d4816156ae565b92915050565b600061024082840312156156f1576156f06156a4565b5b6156fc610240615310565b9050600061570c848285016151c1565b6000830152506020615720848285016151c1565b6020830152506040615734848285016151c1565b6040830152506060615748848285016151c1565b606083015250608061575c848285016151c1565b60808301525060a082013567ffffffffffffffff8111156157805761577f6156a9565b5b61578c84828501615498565b60a08301525060c082013567ffffffffffffffff8111156157b0576157af6156a9565b5b6157bc84828501615498565b60c08301525060e06157d0848285016151c1565b60e08301525061010082013567ffffffffffffffff8111156157f5576157f46156a9565b5b61580184828501615498565b61010083015250610120615817848285016151c1565b6101208301525061014082013567ffffffffffffffff81111561583d5761583c6156a9565b5b61584984828501615498565b6101408301525061016061585f848285016151c1565b6101608301525061018082013567ffffffffffffffff811115615885576158846156a9565b5b61589184828501615498565b610180830152506101a06158a7848285016151c1565b6101a0830152506101c082013567ffffffffffffffff8111156158cd576158cc6156a9565b5b6158d984828501615498565b6101c0830152506101e06158ef8482850161560e565b6101e0830152506102006159058482850161560e565b6102008301525061022061591b848285016156c5565b6102208301525092915050565b60006020828403121561593e5761593d6151a0565b5b600082013567ffffffffffffffff81111561595c5761595b6151a5565b5b615968848285016156da565b91505092915050565b600067ffffffffffffffff82111561598c5761598b6152b0565b5b61599582614ebb565b9050602081019050919050565b60006159b56159b084615971565b615310565b9050828152602081018484840111156159d1576159d061535c565b5b6159dc848285615392565b509392505050565b600082601f8301126159f9576159f86152ab565b5b8135615a098482602086016159a2565b91505092915050565b600060208284031215615a2857615a276151a0565b5b600082013567ffffffffffffffff811115615a4657615a456151a5565b5b615a52848285016159e4565b91505092915050565b615a64816155bb565b8114615a6f57600080fd5b50565b600081359050615a8181615a5b565b92915050565b600060208284031215615a9d57615a9c6151a0565b5b6000615aab84828501615a72565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680615b2a57607f821691505b60208210811415615b3e57615b3d615ae3565b5b50919050565b60006020820190508181036000830152615b5e818461552b565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000615ba082614e27565b9150615bab83614e27565b925082821015615bbe57615bbd615b66565b5b828203905092915050565b6000615bd482614e27565b9150615bdf83614e27565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115615c1457615c13615b66565b5b828201905092915050565b6000615c2a82614e27565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415615c5d57615c5c615b66565b5b600182019050919050565b600081905092915050565b6000615c7e82614e6c565b615c888185615c68565b9350615c98818560208601614e88565b80840191505092915050565b6000615cb08285615c73565b9150615cbc8284615c73565b91508190509392505050565b6000615cd48284615c73565b915081905092915050565b600081519050615cee816155f7565b92915050565b600060208284031215615d0a57615d096151a0565b5b6000615d1884828501615cdf565b91505092915050565b7f427974657320617272617920746f6f2073686f72742e00000000000000000000600082015250565b6000615d5760168361551a565b9150615d6282615d21565b602082019050919050565b60006020820190508181036000830152615d8681615d4a565b9050919050565b7f496e70757420627974657320617272617920697320656d707479000000000000600082015250565b6000615dc3601a8361551a565b9150615dce82615d8d565b602082019050919050565b60006020820190508181036000830152615df281615db6565b905091905056fe63616c6c6572206d7573742062652061676772656761746f7220636f6e747261637421696e2061636b6e6f776c656467653a2063616e6e6f742066696e64206d61746368656420726571756573742063634d7367696e2061636b6e6f776c656467653a2063634d73672061636b206861732065786973746564a2646970667358221220fcd9ca0a549160508d8a682b727d5be1f709ec46ccaf8c28b9bff9be347e878064736f6c634300080b0033"
var EventlistenSMBin = "0x"

// DeployEventlisten deploys a new contract, binding an instance of Eventlisten to it.
func DeployEventlisten(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Eventlisten, error) {
	parsed, err := abi.JSON(strings.NewReader(EventlistenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(EventlistenSMBin)
	} else {
		bytecode = common.FromHex(EventlistenBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, EventlistenABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Eventlisten{EventlistenCaller: EventlistenCaller{contract: contract}, EventlistenTransactor: EventlistenTransactor{contract: contract}, EventlistenFilterer: EventlistenFilterer{contract: contract}}, nil
}

func AsyncDeployEventlisten(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(EventlistenABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(EventlistenSMBin)
	} else {
		bytecode = common.FromHex(EventlistenBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, EventlistenABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Eventlisten is an auto generated Go binding around a Solidity contract.
type Eventlisten struct {
	EventlistenCaller     // Read-only binding to the contract
	EventlistenTransactor // Write-only binding to the contract
	EventlistenFilterer   // Log filterer for contract events
}

// EventlistenCaller is an auto generated read-only Go binding around a Solidity contract.
type EventlistenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventlistenTransactor is an auto generated write-only Go binding around a Solidity contract.
type EventlistenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventlistenFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EventlistenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventlistenSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EventlistenSession struct {
	Contract     *Eventlisten      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventlistenCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EventlistenCallerSession struct {
	Contract *EventlistenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EventlistenTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EventlistenTransactorSession struct {
	Contract     *EventlistenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EventlistenRaw is an auto generated low-level Go binding around a Solidity contract.
type EventlistenRaw struct {
	Contract *Eventlisten // Generic contract binding to access the raw methods on
}

// EventlistenCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EventlistenCallerRaw struct {
	Contract *EventlistenCaller // Generic read-only contract binding to access the raw methods on
}

// EventlistenTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EventlistenTransactorRaw struct {
	Contract *EventlistenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventlisten creates a new instance of Eventlisten, bound to a specific deployed contract.
func NewEventlisten(address common.Address, backend bind.ContractBackend) (*Eventlisten, error) {
	contract, err := bindEventlisten(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eventlisten{EventlistenCaller: EventlistenCaller{contract: contract}, EventlistenTransactor: EventlistenTransactor{contract: contract}, EventlistenFilterer: EventlistenFilterer{contract: contract}}, nil
}

// NewEventlistenCaller creates a new read-only instance of Eventlisten, bound to a specific deployed contract.
func NewEventlistenCaller(address common.Address, caller bind.ContractCaller) (*EventlistenCaller, error) {
	contract, err := bindEventlisten(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventlistenCaller{contract: contract}, nil
}

// NewEventlistenTransactor creates a new write-only instance of Eventlisten, bound to a specific deployed contract.
func NewEventlistenTransactor(address common.Address, transactor bind.ContractTransactor) (*EventlistenTransactor, error) {
	contract, err := bindEventlisten(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventlistenTransactor{contract: contract}, nil
}

// NewEventlistenFilterer creates a new log filterer instance of Eventlisten, bound to a specific deployed contract.
func NewEventlistenFilterer(address common.Address, filterer bind.ContractFilterer) (*EventlistenFilterer, error) {
	contract, err := bindEventlisten(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventlistenFilterer{contract: contract}, nil
}

// bindEventlisten binds a generic wrapper to an already deployed contract.
func bindEventlisten(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventlistenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventlisten *EventlistenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eventlisten.Contract.EventlistenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventlisten *EventlistenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.EventlistenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventlisten *EventlistenRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.EventlistenTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventlisten *EventlistenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eventlisten.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventlisten *EventlistenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventlisten *EventlistenTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GenCMsg is a free data retrieval call binding the contract method 0x0470439a.
//
// Solidity: function GenCMsg() constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCaller) GenCMsg(opts *bind.CallOpts) (TypesCrosschainMessage, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "GenCMsg")
	return *ret0, err
}

// GenCMsg is a free data retrieval call binding the contract method 0x0470439a.
//
// Solidity: function GenCMsg() constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenSession) GenCMsg() (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GenCMsg(&_Eventlisten.CallOpts)
}

// GenCMsg is a free data retrieval call binding the contract method 0x0470439a.
//
// Solidity: function GenCMsg() constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCallerSession) GenCMsg() (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GenCMsg(&_Eventlisten.CallOpts)
}

// AssignStringToBytesArray is a free data retrieval call binding the contract method 0xd081dd0c.
//
// Solidity: function assignStringToBytesArray(string str) constant returns(bytes[])
func (_Eventlisten *EventlistenCaller) AssignStringToBytesArray(opts *bind.CallOpts, str string) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "assignStringToBytesArray", str)
	return *ret0, err
}

// AssignStringToBytesArray is a free data retrieval call binding the contract method 0xd081dd0c.
//
// Solidity: function assignStringToBytesArray(string str) constant returns(bytes[])
func (_Eventlisten *EventlistenSession) AssignStringToBytesArray(str string) ([][]byte, error) {
	return _Eventlisten.Contract.AssignStringToBytesArray(&_Eventlisten.CallOpts, str)
}

// AssignStringToBytesArray is a free data retrieval call binding the contract method 0xd081dd0c.
//
// Solidity: function assignStringToBytesArray(string str) constant returns(bytes[])
func (_Eventlisten *EventlistenCallerSession) AssignStringToBytesArray(str string) ([][]byte, error) {
	return _Eventlisten.Contract.AssignStringToBytesArray(&_Eventlisten.CallOpts, str)
}

// GetAckCmByHash is a free data retrieval call binding the contract method 0x7c822f5c.
//
// Solidity: function getAckCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCaller) GetAckCmByHash(opts *bind.CallOpts, cmhash [32]byte) (TypesCrosschainMessage, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "getAckCmByHash", cmhash)
	return *ret0, err
}

// GetAckCmByHash is a free data retrieval call binding the contract method 0x7c822f5c.
//
// Solidity: function getAckCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenSession) GetAckCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetAckCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// GetAckCmByHash is a free data retrieval call binding the contract method 0x7c822f5c.
//
// Solidity: function getAckCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCallerSession) GetAckCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetAckCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() constant returns(address output)
func (_Eventlisten *EventlistenCaller) GetAggregator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "getAggregator")
	return *ret0, err
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() constant returns(address output)
func (_Eventlisten *EventlistenSession) GetAggregator() (common.Address, error) {
	return _Eventlisten.Contract.GetAggregator(&_Eventlisten.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() constant returns(address output)
func (_Eventlisten *EventlistenCallerSession) GetAggregator() (common.Address, error) {
	return _Eventlisten.Contract.GetAggregator(&_Eventlisten.CallOpts)
}

// GetReqCmByHash is a free data retrieval call binding the contract method 0x9d36e127.
//
// Solidity: function getReqCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCaller) GetReqCmByHash(opts *bind.CallOpts, cmhash [32]byte) (TypesCrosschainMessage, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "getReqCmByHash", cmhash)
	return *ret0, err
}

// GetReqCmByHash is a free data retrieval call binding the contract method 0x9d36e127.
//
// Solidity: function getReqCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenSession) GetReqCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetReqCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// GetReqCmByHash is a free data retrieval call binding the contract method 0x9d36e127.
//
// Solidity: function getReqCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCallerSession) GetReqCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetReqCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// GetRespCmByHash is a free data retrieval call binding the contract method 0xccb5ad2e.
//
// Solidity: function getRespCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCaller) GetRespCmByHash(opts *bind.CallOpts, cmhash [32]byte) (TypesCrosschainMessage, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "getRespCmByHash", cmhash)
	return *ret0, err
}

// GetRespCmByHash is a free data retrieval call binding the contract method 0xccb5ad2e.
//
// Solidity: function getRespCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenSession) GetRespCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetRespCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// GetRespCmByHash is a free data retrieval call binding the contract method 0xccb5ad2e.
//
// Solidity: function getRespCmByHash(bytes32 cmhash) constant returns(TypesCrosschainMessage)
func (_Eventlisten *EventlistenCallerSession) GetRespCmByHash(cmhash [32]byte) (TypesCrosschainMessage, error) {
	return _Eventlisten.Contract.GetRespCmByHash(&_Eventlisten.CallOpts, cmhash)
}

// TestQueryCcAckMsgs is a free data retrieval call binding the contract method 0xf392a66e.
//
// Solidity: function test_query_ccAckMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCaller) TestQueryCcAckMsgs(opts *bind.CallOpts, index *big.Int) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_ccAckMsgs", index)
	return *ret0, err
}

// TestQueryCcAckMsgs is a free data retrieval call binding the contract method 0xf392a66e.
//
// Solidity: function test_query_ccAckMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenSession) TestQueryCcAckMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcAckMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryCcAckMsgs is a free data retrieval call binding the contract method 0xf392a66e.
//
// Solidity: function test_query_ccAckMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCallerSession) TestQueryCcAckMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcAckMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryCcReqMsgs is a free data retrieval call binding the contract method 0x0884e212.
//
// Solidity: function test_query_ccReqMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCaller) TestQueryCcReqMsgs(opts *bind.CallOpts, index *big.Int) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_ccReqMsgs", index)
	return *ret0, err
}

// TestQueryCcReqMsgs is a free data retrieval call binding the contract method 0x0884e212.
//
// Solidity: function test_query_ccReqMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenSession) TestQueryCcReqMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcReqMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryCcReqMsgs is a free data retrieval call binding the contract method 0x0884e212.
//
// Solidity: function test_query_ccReqMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCallerSession) TestQueryCcReqMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcReqMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryCcRespMsgs is a free data retrieval call binding the contract method 0x5dde9283.
//
// Solidity: function test_query_ccRespMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCaller) TestQueryCcRespMsgs(opts *bind.CallOpts, index *big.Int) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_ccRespMsgs", index)
	return *ret0, err
}

// TestQueryCcRespMsgs is a free data retrieval call binding the contract method 0x5dde9283.
//
// Solidity: function test_query_ccRespMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenSession) TestQueryCcRespMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcRespMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryCcRespMsgs is a free data retrieval call binding the contract method 0x5dde9283.
//
// Solidity: function test_query_ccRespMsgs(uint256 index) constant returns(bytes[])
func (_Eventlisten *EventlistenCallerSession) TestQueryCcRespMsgs(index *big.Int) ([][]byte, error) {
	return _Eventlisten.Contract.TestQueryCcRespMsgs(&_Eventlisten.CallOpts, index)
}

// TestQueryHashAckToCmIndex is a free data retrieval call binding the contract method 0xb552adeb.
//
// Solidity: function test_query_hashAckToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCaller) TestQueryHashAckToCmIndex(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_hashAckToCmIndex", hash)
	return *ret0, err
}

// TestQueryHashAckToCmIndex is a free data retrieval call binding the contract method 0xb552adeb.
//
// Solidity: function test_query_hashAckToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenSession) TestQueryHashAckToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashAckToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQueryHashAckToCmIndex is a free data retrieval call binding the contract method 0xb552adeb.
//
// Solidity: function test_query_hashAckToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCallerSession) TestQueryHashAckToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashAckToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQueryHashReqToCmIndex is a free data retrieval call binding the contract method 0xa6261dca.
//
// Solidity: function test_query_hashReqToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCaller) TestQueryHashReqToCmIndex(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_hashReqToCmIndex", hash)
	return *ret0, err
}

// TestQueryHashReqToCmIndex is a free data retrieval call binding the contract method 0xa6261dca.
//
// Solidity: function test_query_hashReqToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenSession) TestQueryHashReqToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashReqToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQueryHashReqToCmIndex is a free data retrieval call binding the contract method 0xa6261dca.
//
// Solidity: function test_query_hashReqToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCallerSession) TestQueryHashReqToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashReqToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQueryHashRespToCmIndex is a free data retrieval call binding the contract method 0x6b640869.
//
// Solidity: function test_query_hashRespToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCaller) TestQueryHashRespToCmIndex(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_hashRespToCmIndex", hash)
	return *ret0, err
}

// TestQueryHashRespToCmIndex is a free data retrieval call binding the contract method 0x6b640869.
//
// Solidity: function test_query_hashRespToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenSession) TestQueryHashRespToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashRespToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQueryHashRespToCmIndex is a free data retrieval call binding the contract method 0x6b640869.
//
// Solidity: function test_query_hashRespToCmIndex(bytes32 hash) constant returns(uint256)
func (_Eventlisten *EventlistenCallerSession) TestQueryHashRespToCmIndex(hash [32]byte) (*big.Int, error) {
	return _Eventlisten.Contract.TestQueryHashRespToCmIndex(&_Eventlisten.CallOpts, hash)
}

// TestQuerySeqToHashAck is a free data retrieval call binding the contract method 0x93e00b01.
//
// Solidity: function test_query_seqToHashAck(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenCaller) TestQuerySeqToHashAck(opts *bind.CallOpts, seq *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_seqToHashAck", seq)
	return *ret0, err
}

// TestQuerySeqToHashAck is a free data retrieval call binding the contract method 0x93e00b01.
//
// Solidity: function test_query_seqToHashAck(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenSession) TestQuerySeqToHashAck(seq *big.Int) ([32]byte, error) {
	return _Eventlisten.Contract.TestQuerySeqToHashAck(&_Eventlisten.CallOpts, seq)
}

// TestQuerySeqToHashAck is a free data retrieval call binding the contract method 0x93e00b01.
//
// Solidity: function test_query_seqToHashAck(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenCallerSession) TestQuerySeqToHashAck(seq *big.Int) ([32]byte, error) {
	return _Eventlisten.Contract.TestQuerySeqToHashAck(&_Eventlisten.CallOpts, seq)
}

// TestQuerySeqToHashReq is a free data retrieval call binding the contract method 0x7e0694c2.
//
// Solidity: function test_query_seqToHashReq(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenCaller) TestQuerySeqToHashReq(opts *bind.CallOpts, seq *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Eventlisten.contract.Call(opts, out, "test_query_seqToHashReq", seq)
	return *ret0, err
}

// TestQuerySeqToHashReq is a free data retrieval call binding the contract method 0x7e0694c2.
//
// Solidity: function test_query_seqToHashReq(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenSession) TestQuerySeqToHashReq(seq *big.Int) ([32]byte, error) {
	return _Eventlisten.Contract.TestQuerySeqToHashReq(&_Eventlisten.CallOpts, seq)
}

// TestQuerySeqToHashReq is a free data retrieval call binding the contract method 0x7e0694c2.
//
// Solidity: function test_query_seqToHashReq(uint256 seq) constant returns(bytes32)
func (_Eventlisten *EventlistenCallerSession) TestQuerySeqToHashReq(seq *big.Int) ([32]byte, error) {
	return _Eventlisten.Contract.TestQuerySeqToHashReq(&_Eventlisten.CallOpts, seq)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactor) Acknowledge(opts *bind.TransactOpts, data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "acknowledge", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "acknowledge", data)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenSession) Acknowledge(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.Acknowledge(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncAcknowledge(handler, &_Eventlisten.TransactOpts, data)
}

// Acknowledge is a paid mutator transaction binding the contract method 0xe7ef5f9b.
//
// Solidity: function acknowledge(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactorSession) Acknowledge(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.Acknowledge(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncAcknowledge(handler, &_Eventlisten.TransactOpts, data)
}

// EmitSendOut is a paid mutator transaction binding the contract method 0xe952f036.
//
// Solidity: function emit_sendOut() returns()
func (_Eventlisten *EventlistenTransactor) EmitSendOut(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "emit_sendOut")
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncEmitSendOut(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "emit_sendOut")
}

// EmitSendOut is a paid mutator transaction binding the contract method 0xe952f036.
//
// Solidity: function emit_sendOut() returns()
func (_Eventlisten *EventlistenSession) EmitSendOut() (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.EmitSendOut(&_Eventlisten.TransactOpts)
}

func (_Eventlisten *EventlistenSession) AsyncEmitSendOut(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncEmitSendOut(handler, &_Eventlisten.TransactOpts)
}

// EmitSendOut is a paid mutator transaction binding the contract method 0xe952f036.
//
// Solidity: function emit_sendOut() returns()
func (_Eventlisten *EventlistenTransactorSession) EmitSendOut() (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.EmitSendOut(&_Eventlisten.TransactOpts)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncEmitSendOut(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncEmitSendOut(handler, &_Eventlisten.TransactOpts)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactor) ReceiveIn(opts *bind.TransactOpts, data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "receiveIn", data)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncReceiveIn(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "receiveIn", data)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenSession) ReceiveIn(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.ReceiveIn(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncReceiveIn(handler, &_Eventlisten.TransactOpts, data)
}

// ReceiveIn is a paid mutator transaction binding the contract method 0x32890a6c.
//
// Solidity: function receiveIn(bytes[] data) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactorSession) ReceiveIn(data [][]byte) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.ReceiveIn(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncReceiveIn(handler, &_Eventlisten.TransactOpts, data)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactor) Response(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "response", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "response", ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.Response(&_Eventlisten.TransactOpts, ccMsg)
}

func (_Eventlisten *EventlistenSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncResponse(handler, &_Eventlisten.TransactOpts, ccMsg)
}

// Response is a paid mutator transaction binding the contract method 0xbff894a3.
//
// Solidity: function response(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactorSession) Response(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.Response(&_Eventlisten.TransactOpts, ccMsg)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncResponse(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncResponse(handler, &_Eventlisten.TransactOpts, ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactor) SendOut(opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(TypesCrosschainMessage)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "sendOut", ccMsg)
	return *ret0, *ret1, transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncSendOut(handler func(*types.Receipt, error), opts *bind.TransactOpts, ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "sendOut", ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenSession) SendOut(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.SendOut(&_Eventlisten.TransactOpts, ccMsg)
}

func (_Eventlisten *EventlistenSession) AsyncSendOut(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncSendOut(handler, &_Eventlisten.TransactOpts, ccMsg)
}

// SendOut is a paid mutator transaction binding the contract method 0x96070c01.
//
// Solidity: function sendOut(TypesCrosschainMessage ccMsg) returns(TypesCrosschainMessage, string)
func (_Eventlisten *EventlistenTransactorSession) SendOut(ccMsg TypesCrosschainMessage) (TypesCrosschainMessage, string, *types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.SendOut(&_Eventlisten.TransactOpts, ccMsg)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncSendOut(handler func(*types.Receipt, error), ccMsg TypesCrosschainMessage) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncSendOut(handler, &_Eventlisten.TransactOpts, ccMsg)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address input) returns()
func (_Eventlisten *EventlistenTransactor) SetAggregator(opts *bind.TransactOpts, input common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "setAggregator", input)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncSetAggregator(handler func(*types.Receipt, error), opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "setAggregator", input)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address input) returns()
func (_Eventlisten *EventlistenSession) SetAggregator(input common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.SetAggregator(&_Eventlisten.TransactOpts, input)
}

func (_Eventlisten *EventlistenSession) AsyncSetAggregator(handler func(*types.Receipt, error), input common.Address) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncSetAggregator(handler, &_Eventlisten.TransactOpts, input)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address input) returns()
func (_Eventlisten *EventlistenTransactorSession) SetAggregator(input common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.SetAggregator(&_Eventlisten.TransactOpts, input)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncSetAggregator(handler func(*types.Receipt, error), input common.Address) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncSetAggregator(handler, &_Eventlisten.TransactOpts, input)
}

// TestAcknowledge is a paid mutator transaction binding the contract method 0x957cffd2.
//
// Solidity: function test_acknowledge(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactor) TestAcknowledge(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "test_acknowledge", data)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncTestAcknowledge(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "test_acknowledge", data)
}

// TestAcknowledge is a paid mutator transaction binding the contract method 0x957cffd2.
//
// Solidity: function test_acknowledge(bytes[] data) returns()
func (_Eventlisten *EventlistenSession) TestAcknowledge(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestAcknowledge(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncTestAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestAcknowledge(handler, &_Eventlisten.TransactOpts, data)
}

// TestAcknowledge is a paid mutator transaction binding the contract method 0x957cffd2.
//
// Solidity: function test_acknowledge(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactorSession) TestAcknowledge(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestAcknowledge(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncTestAcknowledge(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestAcknowledge(handler, &_Eventlisten.TransactOpts, data)
}

// TestReceiveIn is a paid mutator transaction binding the contract method 0x2369d54c.
//
// Solidity: function test_receiveIn(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactor) TestReceiveIn(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "test_receiveIn", data)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncTestReceiveIn(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "test_receiveIn", data)
}

// TestReceiveIn is a paid mutator transaction binding the contract method 0x2369d54c.
//
// Solidity: function test_receiveIn(bytes[] data) returns()
func (_Eventlisten *EventlistenSession) TestReceiveIn(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestReceiveIn(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncTestReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestReceiveIn(handler, &_Eventlisten.TransactOpts, data)
}

// TestReceiveIn is a paid mutator transaction binding the contract method 0x2369d54c.
//
// Solidity: function test_receiveIn(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactorSession) TestReceiveIn(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestReceiveIn(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncTestReceiveIn(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestReceiveIn(handler, &_Eventlisten.TransactOpts, data)
}

// TestResponesErr is a paid mutator transaction binding the contract method 0x73dae718.
//
// Solidity: function test_respones_err(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactor) TestResponesErr(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "test_respones_err", data)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncTestResponesErr(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "test_respones_err", data)
}

// TestResponesErr is a paid mutator transaction binding the contract method 0x73dae718.
//
// Solidity: function test_respones_err(bytes[] data) returns()
func (_Eventlisten *EventlistenSession) TestResponesErr(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestResponesErr(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncTestResponesErr(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestResponesErr(handler, &_Eventlisten.TransactOpts, data)
}

// TestResponesErr is a paid mutator transaction binding the contract method 0x73dae718.
//
// Solidity: function test_respones_err(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactorSession) TestResponesErr(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestResponesErr(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncTestResponesErr(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestResponesErr(handler, &_Eventlisten.TransactOpts, data)
}

// TestResponse is a paid mutator transaction binding the contract method 0x89104747.
//
// Solidity: function test_response(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactor) TestResponse(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "test_response", data)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncTestResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "test_response", data)
}

// TestResponse is a paid mutator transaction binding the contract method 0x89104747.
//
// Solidity: function test_response(bytes[] data) returns()
func (_Eventlisten *EventlistenSession) TestResponse(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestResponse(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncTestResponse(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestResponse(handler, &_Eventlisten.TransactOpts, data)
}

// TestResponse is a paid mutator transaction binding the contract method 0x89104747.
//
// Solidity: function test_response(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactorSession) TestResponse(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestResponse(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncTestResponse(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestResponse(handler, &_Eventlisten.TransactOpts, data)
}

// TestSendOut is a paid mutator transaction binding the contract method 0xf6b9f4b9.
//
// Solidity: function test_sendOut(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactor) TestSendOut(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Eventlisten.contract.TransactWithResult(opts, out, "test_sendOut", data)
	return transaction, receipt, err
}

func (_Eventlisten *EventlistenTransactor) AsyncTestSendOut(handler func(*types.Receipt, error), opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.contract.AsyncTransact(opts, handler, "test_sendOut", data)
}

// TestSendOut is a paid mutator transaction binding the contract method 0xf6b9f4b9.
//
// Solidity: function test_sendOut(bytes[] data) returns()
func (_Eventlisten *EventlistenSession) TestSendOut(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestSendOut(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenSession) AsyncTestSendOut(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestSendOut(handler, &_Eventlisten.TransactOpts, data)
}

// TestSendOut is a paid mutator transaction binding the contract method 0xf6b9f4b9.
//
// Solidity: function test_sendOut(bytes[] data) returns()
func (_Eventlisten *EventlistenTransactorSession) TestSendOut(data [][]byte) (*types.Transaction, *types.Receipt, error) {
	return _Eventlisten.Contract.TestSendOut(&_Eventlisten.TransactOpts, data)
}

func (_Eventlisten *EventlistenTransactorSession) AsyncTestSendOut(handler func(*types.Receipt, error), data [][]byte) (*types.Transaction, error) {
	return _Eventlisten.Contract.AsyncTestSendOut(handler, &_Eventlisten.TransactOpts, data)
}

// EventlistenCmHash represents a CmHash event raised by the Eventlisten contract.
type EventlistenCmHash struct {
	Hash  [32]byte
	Phase *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchCmHash is a free log subscription operation binding the contract event 0x61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e.
//
// Solidity: event CmHash(bytes32 indexed hash, uint256 indexed phase)
func (_Eventlisten *EventlistenFilterer) WatchCmHash(fromBlock *int64, handler func(int, []types.Log), hash [32]byte, phase *big.Int) (string, error) {
	return _Eventlisten.contract.WatchLogs(fromBlock, handler, "CmHash", hash, phase)
}

func (_Eventlisten *EventlistenFilterer) WatchAllCmHash(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.contract.WatchLogs(fromBlock, handler, "CmHash")
}

// ParseCmHash is a log parse operation binding the contract event 0x61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e.
//
// Solidity: event CmHash(bytes32 indexed hash, uint256 indexed phase)
func (_Eventlisten *EventlistenFilterer) ParseCmHash(log types.Log) (*EventlistenCmHash, error) {
	event := new(EventlistenCmHash)
	if err := _Eventlisten.contract.UnpackLog(event, "CmHash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCmHash is a free log subscription operation binding the contract event 0x61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e.
//
// Solidity: event CmHash(bytes32 indexed hash, uint256 indexed phase)
func (_Eventlisten *EventlistenSession) WatchCmHash(fromBlock *int64, handler func(int, []types.Log), hash [32]byte, phase *big.Int) (string, error) {
	return _Eventlisten.Contract.WatchCmHash(fromBlock, handler, hash, phase)
}

func (_Eventlisten *EventlistenSession) WatchAllCmHash(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.Contract.WatchAllCmHash(fromBlock, handler)
}

// ParseCmHash is a log parse operation binding the contract event 0x61a7a3d37ba6dba6562dce919ee661ff6cb1718afde67d18a71df16ea7345d1e.
//
// Solidity: event CmHash(bytes32 indexed hash, uint256 indexed phase)
func (_Eventlisten *EventlistenSession) ParseCmHash(log types.Log) (*EventlistenCmHash, error) {
	return _Eventlisten.Contract.ParseCmHash(log)
}

// EventlistenTestSendOut represents a TestSendOut event raised by the Eventlisten contract.
type EventlistenTestSendOut struct {
	CMsg [][]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchTestSendOut is a free log subscription operation binding the contract event 0x1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e.
//
// Solidity: event Test_sendOut(bytes[] CMsg)
func (_Eventlisten *EventlistenFilterer) WatchTestSendOut(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.contract.WatchLogs(fromBlock, handler, "Test_sendOut")
}

func (_Eventlisten *EventlistenFilterer) WatchAllTestSendOut(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.contract.WatchLogs(fromBlock, handler, "Test_sendOut")
}

// ParseTestSendOut is a log parse operation binding the contract event 0x1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e.
//
// Solidity: event Test_sendOut(bytes[] CMsg)
func (_Eventlisten *EventlistenFilterer) ParseTestSendOut(log types.Log) (*EventlistenTestSendOut, error) {
	event := new(EventlistenTestSendOut)
	if err := _Eventlisten.contract.UnpackLog(event, "Test_sendOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTestSendOut is a free log subscription operation binding the contract event 0x1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e.
//
// Solidity: event Test_sendOut(bytes[] CMsg)
func (_Eventlisten *EventlistenSession) WatchTestSendOut(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.Contract.WatchTestSendOut(fromBlock, handler)
}

func (_Eventlisten *EventlistenSession) WatchAllTestSendOut(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Eventlisten.Contract.WatchAllTestSendOut(fromBlock, handler)
}

// ParseTestSendOut is a log parse operation binding the contract event 0x1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e.
//
// Solidity: event Test_sendOut(bytes[] CMsg)
func (_Eventlisten *EventlistenSession) ParseTestSendOut(log types.Log) (*EventlistenTestSendOut, error) {
	return _Eventlisten.Contract.ParseTestSendOut(log)
}
