// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DataStorage {
    // 数据存储映射
    mapping(bytes32 => uint256) private _data;
    // 独立记录key存在状态的映射（解决零值存储问题）
    mapping(bytes32 => bool) private _keyExists;

    // 存储数据时检查重复
    function setData(string memory key, uint256 value) external {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        
        // 检查key是否已存在
        require(!_keyExists[hashedKey], "Key already exists");
        
        _data[hashedKey] = value;
        _keyExists[hashedKey] = true; // 标记为已存在
    }

    // 查询数据（返回默认值 0 当 key 不存在时）
    function getData(string memory key) external view returns (uint256) {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        return _data[hashedKey];
    }

    // 新增：检查key是否存在的独立方法
    function keyExists(string memory key) external view returns (bool) {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        return _keyExists[hashedKey];
    }
}