// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EnhancedDataStorage {
    // 数据存储结构体
    struct DataRecord {
        uint256 value;     // 存储值
        uint256 tag;       // 隐私标签（0=明文，1=密文）
    }
    
    // 存储主数据
    mapping(bytes32 => DataRecord) private _data;
    // 独立记录key存在状态
    mapping(bytes32 => bool) private _keyExists;

    // 事件日志
    event DataStored(bytes32 indexed keyHash, uint256 tag);
    event TagUpdated(bytes32 indexed keyHash, uint256 newTag);

    // 存储明文数据（自动标记tag=0）
    function setPlainData(string memory key, uint256 value) external {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        require(!_keyExists[hashedKey], "Key already exists");
        
        _data[hashedKey] = DataRecord(value, 0);
        _keyExists[hashedKey] = true;
        emit DataStored(hashedKey, 0);
    }

    // 存储加密数据（强制标记tag=1）
    function setEncryptedData(string memory key, uint256 cipherValue) external {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        require(!_keyExists[hashedKey], "Key already exists");
        
        _data[hashedKey] = DataRecord(cipherValue, 1);
        _keyExists[hashedKey] = true;
        emit DataStored(hashedKey, 1);
    }

    // 更新数据标签（仅允许0/1）
    function updateTag(string memory key, uint256 newTag) external {
        require(newTag <= 1, "Invalid tag value");
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        require(_keyExists[hashedKey], "Key not exist");
        
        _data[hashedKey].tag = newTag;
        emit TagUpdated(hashedKey, newTag);
    }

    // 获取数据值（兼容旧接口）
    function getData(string memory key) external view returns (uint256) {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        return _data[hashedKey].value;
    }

    // 获取完整数据记录
    function getDataRecord(string memory key) external view returns (uint256 value, uint256 tag) {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        require(_keyExists[hashedKey], "Key not exist");
        DataRecord storage record = _data[hashedKey];
        return (record.value, record.tag);
    }

    // 检查密钥存在状态
    function keyExists(string memory key) external view returns (bool) {
        bytes32 hashedKey = keccak256(abi.encodePacked(key));
        return _keyExists[hashedKey];
    }
}