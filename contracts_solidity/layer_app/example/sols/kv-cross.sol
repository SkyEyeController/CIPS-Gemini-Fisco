// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Universal Key-Value Store
 * @dev 支持存储和查询任意字符串键值对的通用合约
 */
contract UniversalKVStore {
    // 存储键值对
    mapping(bytes32 => string) private store;
    
    // 事件：记录存储操作
    event ValueSet(string indexed key, string value);
    
    /**
     * @dev 存储键值对
     * @param key 键（字符串）
     * @param value 值（字符串）
     */
    function set(string calldata key, string calldata value) external {
        bytes32 keyHash = keccak256(abi.encodePacked(key));
        store[keyHash] = value;
        emit ValueSet(key, value);
    }
    
    /**
     * @dev 查询键对应的值
     * @param key 要查询的键
     * @return 返回键对应的值，如果不存在则返回空字符串
     */
    function get(string calldata key) external view returns (string memory) {
        bytes32 keyHash = keccak256(abi.encodePacked(key));
        return store[keyHash];
    }

}