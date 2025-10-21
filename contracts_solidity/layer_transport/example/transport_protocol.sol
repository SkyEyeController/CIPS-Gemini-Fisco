// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interface.sol";
import "../../utils/Types.sol";

contract TransportProtocol is ITransportProtocol {
    //定义事件
    event CmHash(bytes32 indexed hash, uint256 phase);

    // 定义用到的状态变量
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

    // 转发层向外发出跨链消息
    // 1. 抛出跨链事件时, 仅抛出跨链消息的哈希
    // 2. 跨链消息存储在列表 ccmsgs 中(序号为 cmIndex), 并将 {hash: cmIndex} 存储在 hashToCmIndex 字典中
    // 3. 其中 hash 是跨链消息在"当前链"上的哈希, 之所以是当前链, 是因为:
    // 同一个跨链请求, 其在源链和目的链上的跨链消息并不完全相同(目的链上的跨链消息增加了响应数据), 导致跨链消息的哈希不同
    function sendOut(
        Types.CrosschainMessage memory ccMsg
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
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
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        // 从 data 中恢复跨链消息结构体
        Types.CrosschainMessage memory ccMsg = Types.cmFromLB(data = data);
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
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
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
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        Types.CrosschainMessage memory ccMsg = Types.cmFromLB(data = data);

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
}
