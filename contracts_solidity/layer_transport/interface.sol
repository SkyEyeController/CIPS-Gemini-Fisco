// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/Types.sol";

interface ITransportProtocol {
    // 转发层向外发出跨链消息
    function sendOut(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 转发层收到外来的跨链消息
    // 需要解码转换为对应的跨链消息结构体
    function receiveIn(
        bytes[] memory data
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 目的链向源链响应跨链消息
    // 和 sendOut 可能重复, 但是作为一个"请求-响应"式协议, 最好还是分开
    function response(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 目的链向源链响应跨链消息
    // 和 receiveIn 可能重复, 原因同上
    function acknowledge(
        bytes[] memory data
    ) external returns (Types.CrosschainMessage memory, string memory);
}
