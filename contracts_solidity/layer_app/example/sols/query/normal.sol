// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IDataStorage {
    function getData(string memory key) external view returns (uint256);
}

contract DataQuery {
    // 通过合约地址和 key 查询外部合约数据
    function getValue(address contractAddress, string memory key) external view returns (uint256) {
        return IDataStorage(contractAddress).getData(key);
    }
}