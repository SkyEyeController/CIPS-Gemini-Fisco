// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interface.sol";
import "../../utils/Types.sol";
import "../../utils/ErrorInfo.sol";

contract EventListen is ITransportProtocol {
    //定义事件
    event CmHash(bytes32 indexed hash, uint256 indexed phase);
    event Test_sendOut(bytes[] CMsg);

    // 定义用到的状态变量
    //一个地址
    address aggregatorAddress;
    // 三个列表
    Types.CrosschainMessage[] private ccReqMsgs;
    Types.CrosschainMessage[] private ccRespMsgs;
    Types.CrosschainMessage[] private ccAckMsgs;
    // 五个映射
    mapping(bytes32 => uint256) private hashReqToCmIndex;
    mapping(bytes32 => uint256) private hashRespToCmIndex;
    mapping(bytes32 => uint256) private hashAckToCmIndex;
    mapping(uint256 => bytes32) private seqToHashReq;
    mapping(uint256 => bytes32) private seqToHashAck;

    constructor() {
        // # 第 1 个元素为无效元素
        Types.CrosschainMessage memory empty;
        ccReqMsgs.push(empty);
        ccRespMsgs.push(empty);
        ccAckMsgs.push(empty);
    }

    // 检查发起者是否为合约创建者
    function setAggregator(address input) public {
        aggregatorAddress = input;
    }
    function getAggregator() public view returns (address output) {
        output = aggregatorAddress;
    }

    // 转发层向外发出跨链消息
    // 1. 抛出跨链事件时, 仅抛出跨链消息的哈希
    // 2. 跨链消息存储在列表 ccmsgs 中(序号为 cmIndex), 并将 {hash: cmIndex} 存储在 hashToCmIndex 字典中
    // 3. 其中 hash 是跨链消息在"当前链"上的哈希, 之所以是当前链, 是因为:
    // 同一个跨链请求, 其在源链和目的链上的跨链消息并不完全相同(目的链上的跨链消息增加了响应数据), 导致跨链消息的哈希不同
    function sendOut(
        Types.CrosschainMessage memory ccMsg
    ) public override returns (Types.CrosschainMessage memory, string memory) {
        // # 调用者权限检查: 只能是aggragator调用该方法
        address aggregatorAddress_holder = getAggregator();
        if (aggregatorAddress_holder != msg.sender) {
            return (ccMsg, "caller must be aggregator contract!");
        }

        //  # 检查是否存在相同 seq 的跨链请求消息
        if (seqToHashReq[ccMsg.seq] != bytes32(0)) {
            return (ccMsg, "exist the same ccMsg.seq");
        }

        bytes32 cmReqHash = Types.cmToHash(ccMsg);
        // 更新映射
        ccReqMsgs.push(ccMsg);
        // ccMsg 在源链 ccReqMsgs 中的位置
        hashReqToCmIndex[cmReqHash] = ccReqMsgs.length - 1;
        seqToHashReq[ccMsg.seq] = cmReqHash;

        // 抛出事件
        emit CmHash(cmReqHash, 1);

        return (ccMsg, "");
    }

    // 转发层收到外来的跨链消息
    // 需要解码转换为对应的跨链消息结构体
    function receiveIn(
        bytes[] memory data
    ) public override returns (Types.CrosschainMessage memory, string memory) {
        Types.CrosschainMessage memory ccMsg;
        // # 调用者权限检查: 只能是aggragator调用该方法
        address aggregatorAddress_holder = getAggregator();
        if (aggregatorAddress_holder != msg.sender) {
            return (ccMsg, "caller must be aggregator contract!");
        }

        // 从 data 中恢复跨链消息结构体
        ccMsg = Types.cmFromLB(data = data);
        // 判断跨链消息是否重复
        bytes32 cmhash = Types.cmToHash(ccMsg = ccMsg);
        if (hashReqToCmIndex[cmhash] != 0) {
            return (ccMsg, "in receiveIn: ccMsg has existed");
        }
        // 存储到本地
        ccMsg.hashReq = cmhash;
        // 为了能够在 response 函数中, 还能找到该消息
        ccRespMsgs.push(ccMsg);
        hashReqToCmIndex[cmhash] = ccRespMsgs.length - 1;
        // ccMsg 在目的链 ccmsgs 中的位置
        return (ccMsg, "");
    }

    // 目的链向源链响应跨链消息
    // 和 sendOut 可能重复, 但是作为一个"请求-响应"式协议, 最好还是分开
    function response(
        Types.CrosschainMessage memory ccMsg
    ) public override returns (Types.CrosschainMessage memory, string memory) {
        // # 调用者权限检查: 只能是aggragator调用该方法
        address aggregatorAddress_holder = getAggregator();
        if (aggregatorAddress_holder != msg.sender) {
            return (ccMsg, "caller must be aggregator contract!");
        }

        // 计算跨链消息响应向哈希值
        bytes32 cmRespHash = Types.cmToHash(ccMsg);
        ccMsg.hashResp = cmRespHash;
        // 更新已保存的请求向跨链消息(响应数据发生了变化)
        uint256 cmIndex = hashReqToCmIndex[ccMsg.hashReq];
        hashRespToCmIndex[cmRespHash] = cmIndex;
        ccRespMsgs[cmIndex] = ccMsg;

        // 抛出响应事件
        // phase = 2
        emit CmHash(cmRespHash, 2);

        return (ccMsg, "");
    }

    // 目的链向源链响应跨链消息
    // 和 receiveIn 可能重复, 原因同上
    function acknowledge(
        bytes[] memory data
    ) public override returns (Types.CrosschainMessage memory, string memory) {
        Types.CrosschainMessage memory ccMsg;
        // # 调用者权限检查: 只能是aggragator调用该方法
        address aggregatorAddress_holder = getAggregator();
        if (aggregatorAddress_holder != msg.sender) {
            return (ccMsg, "caller must be aggregator contract!");
        }

        ccMsg = Types.cmFromLB(data = data);
        // 判断跨链消息是否已经存储在 ack 队列中
        uint256 seq = ccMsg.seq;
        if (seqToHashAck[seq] != bytes32(0)) {
            return (ccMsg, "in acknowledge: ccMsg ack has existed");
        }
        // 根据 ccMsg.seq 判断 ack 跨链消息是否有与之匹配的 req 跨链消息
        if (seqToHashReq[seq] == bytes32(0)) {
            return (ccMsg, "in acknowledge: cannot find matched request ccMsg");
        }

        // 已确定跨链消息 ccMsg 是需要处理的 ack 消息
        // 计算响应向跨链消息哈希值 cmAckHash, 存储到 ack 队列中(未知为 cmAckIndex), 并做好映射 {cmAckHash: cmAckIndex}
        bytes32 cmAckHash = Types.cmToHash(ccMsg);
        ccAckMsgs.push(ccMsg);
        hashAckToCmIndex[cmAckHash] = ccAckMsgs.length - 1;
        seqToHashAck[ccMsg.seq] = cmAckHash;
        // 为重复检测做判断

        // 抛出响应事件
        // phase = 3
        emit CmHash(cmAckHash, 3);

        return (ccMsg, "");
    }

    // ************************************************************************************************
    // ************************************************************************************************
    // ***************************************  后续内容用于获取值  ***************************************
    // ************************************************************************************************
    // ************************************************************************************************

    function getReqCmByHash(
        bytes32 cmhash
    ) public view returns (Types.CrosschainMessage memory) {
        uint256 cmIndex = hashReqToCmIndex[cmhash];
        Types.CrosschainMessage memory ccMsg;
        ccMsg = ccReqMsgs[cmIndex];
        return ccMsg;
    }

    function getRespCmByHash(
        bytes32 cmhash
    ) public view returns (Types.CrosschainMessage memory) {
        uint256 cmIndex = hashRespToCmIndex[cmhash];
        Types.CrosschainMessage memory ccMsg;
        ccMsg = ccRespMsgs[cmIndex];
        return ccMsg;
    }

    function getAckCmByHash(
        bytes32 cmhash
    ) public view returns (Types.CrosschainMessage memory) {
        uint256 cmIndex = hashAckToCmIndex[cmhash];
        Types.CrosschainMessage memory ccMsg;
        ccMsg = ccAckMsgs[cmIndex];
        return ccMsg;
    }

    // ************************************************************************************************
    // ************************************************************************************************
    // ***************************************  后续内容用于测试  ***************************************
    // ************************************************************************************************
    // ************************************************************************************************

    //经过验证 链上抛出的数据报文与链下go生成的一致
    function emit_sendOut() public {
        Types.CrosschainMessage memory ccMsg = GenCMsg();
        emit Test_sendOut(Types.cmToLB(ccMsg));
    }

    function GenCMsg() public pure returns (Types.CrosschainMessage memory) {
        Types.CrosschainMessage memory cm;
        cm.transportTypeId = 1;
        cm.transportPayload = assignStringToBytesArray("transportPayload");
        cm.verificationTypeId = 1;
        cm.verificationPayload = assignStringToBytesArray(
            "verificationPayload"
        );
        cm.transmissionTypeId = 1;
        cm.transmissionPayload = assignStringToBytesArray(
            "transmissionPayload"
        );
        cm.transactionTypeId = 1;
        cm.transactionPayload = assignStringToBytesArray("transactionPayload");
        cm.srcAppId = 1;
        cm.dstAppId = 1;
        cm.payloadReq = assignStringToBytesArray("payloadReq");
        cm.payloadResp = assignStringToBytesArray("payloadResp");
        cm.srcChainId = 1;
        cm.dstChainId = 2;
        cm.seq = 1;
        cm.ack = false;
        return cm;
    }

    function assignStringToBytesArray(
        string memory str
    ) public pure returns (bytes[] memory) {
        bytes[] memory byteArray = new bytes[](1); // 创建一个包含一个元素的bytes数组
        byteArray[0] = bytes(str); // 将string转换为bytes并赋值
        return byteArray;
    }

    // ************************************************************************************************
    // ************************************************************************************************
    // ***************************************  后续内容用于测试  ***************************************
    // ************************************************************************************************
    // ************************************************************************************************

    function test_sendOut(bytes[] memory data) public {
        string memory err;
        Types.CrosschainMessage memory ccMsg = Types.cmFromLB(data = data);
        Types.CrosschainMessage memory returnCcMsg;
        (returnCcMsg, err) = sendOut(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
        }
    }

    function test_query_seqToHashReq(
        uint256 seq
    ) public view returns (bytes32) {
        return seqToHashReq[seq];
    }

    function test_query_hashReqToCmIndex(
        bytes32 hash
    ) public view returns (uint256) {
        return hashReqToCmIndex[hash];
    }

    function test_query_ccReqMsgs(
        uint256 index
    ) public view returns (bytes[] memory) {
        if (index >= ccReqMsgs.length) {
            return new bytes[](0);
        } else {
            return Types.cmToLB(ccReqMsgs[index]);
        }
    }

    function test_receiveIn(bytes[] memory data) public {
        string memory err;
        Types.CrosschainMessage memory returnCcMsg;
        (returnCcMsg, err) = receiveIn(data);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
        }
    }

    function test_query_hashRespToCmIndex(
        bytes32 hash
    ) public view returns (uint256) {
        return hashRespToCmIndex[hash];
    }

    function test_query_ccRespMsgs(
        uint256 index
    ) public view returns (bytes[] memory) {
        if (index >= ccRespMsgs.length) {
            return new bytes[](0);
        } else {
            return Types.cmToLB(ccRespMsgs[index]);
        }
    }

    // 注意, 要先执行 receiveIn 函数
    // 因为 receiveIn 和 response 是在一笔交易中执行的
    // 且 receiveIn 函数的存储操作会影响 response 函数中的读取操作
    function test_response(bytes[] memory data) public {
        string memory err;
        Types.CrosschainMessage memory returnCcMsg;
        (returnCcMsg, err) = receiveIn(data);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
            return;
        }
        (returnCcMsg, err) = response(returnCcMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
        }
    }

    // 仅用于测试是否验证调用者
    function test_respones_err(bytes[] memory data) public {
        string memory err;
        Types.CrosschainMessage memory ccMsg = Types.cmFromLB(data = data);
        Types.CrosschainMessage memory respCcMsg;
        (respCcMsg, err) = response(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
        }
    }

    function test_acknowledge(bytes[] memory data) public {
        string memory err;
        Types.CrosschainMessage memory respCcMsg;
        (respCcMsg, err) = acknowledge(data);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(10001, err);
        }
    }

    function test_query_seqToHashAck(
        uint256 seq
    ) public view returns (bytes32) {
        return seqToHashAck[seq];
    }

    function test_query_hashAckToCmIndex(
        bytes32 hash
    ) public view returns (uint256) {
        return hashAckToCmIndex[hash];
    }

    function test_query_ccAckMsgs(
        uint256 index
    ) public view returns (bytes[] memory) {
        if (index >= ccAckMsgs.length) {
            return new bytes[](0);
        } else {
            return Types.cmToLB(ccAckMsgs[index]);
        }
    }
}
