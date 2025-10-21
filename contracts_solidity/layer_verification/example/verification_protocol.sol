// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interface.sol";
import "../../utils/Types.sol";

contract VericationProtocol is IVericationProtocol {
    // 验证层在源链上需要执行的内容
    // 一般置空即可
    function prepare(
        Types.CrosschainMessage memory ccMsg
    ) external override returns (Types.CrosschainMessage memory, string memory) {
        return (ccMsg, "");
    }

    function update(
        bytes[] memory data
    ) external override returns (bool, string memory) {
        return (true, "");
    }

    function verify(
        Types.CrosschainMessage memory ccMsg
    ) external override returns (bool, string memory) {
        return (true, "");
    }

    // 目的链验证层响应时需要执行的操作
    // 暂时也不知道有什么用, 权当提供一个接口了
    function response(
        Types.CrosschainMessage memory ccMsg
    ) external override returns (Types.CrosschainMessage memory, string memory) {
        return (ccMsg, "");
    }
}
