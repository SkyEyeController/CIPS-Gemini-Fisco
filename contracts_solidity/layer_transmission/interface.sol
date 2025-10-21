// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/Types.sol";

interface ITransmissionProtocol {
    // 传输层定序
    function send_msg(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 目的链传输层接收跨链消息, 判断是否重复
    // [可选]判断顺序正确性
    function receive_msg(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 目的链传输层响应时需要执行的操作
    function response(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    // 目的链传输层响应时需要执行的操作
    function acknowledge(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);
}
