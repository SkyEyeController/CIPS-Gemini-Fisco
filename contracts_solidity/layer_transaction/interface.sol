// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/Types.sol";

interface ITransactionProtocol {
    /*
    难以无法抽象通用的事务协议接口, 所以采用 flag 作为标记;
    事务协议内部根据 flag 的取值选择执行对应的函数
    flag:
        0, 执行初始函数
    */
    function work(
        uint8 flag,
        Types.CrosschainMessage memory ccMsg,
        bytes[] memory payload
    ) external returns (Types.CrosschainMessage memory, string memory);
}
