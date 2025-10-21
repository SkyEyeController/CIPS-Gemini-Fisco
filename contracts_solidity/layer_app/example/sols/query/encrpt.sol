// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPrivateDataStorage {
    function getDataRecord(string memory key) external view returns (uint256 value, uint256 tag);
}

contract PrivateDataQuery {
    // 隐私数据查询（返回加密内容）
    function queryEncryptedData(
        address storageContract,
        string memory key
    ) external view returns (uint256 value, uint256 tag) {
        return IPrivateDataStorage(storageContract).getDataRecord(key);
    }
}