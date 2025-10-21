// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/ErrorInfo.sol";

contract TransportReg {
    event TransportRegister(
        address indexed transportAddress,
        uint256 indexed transportId
    );

    // 转发协议地址 -> 转发协议 ID
    mapping(address => uint256) internal addrToIndex;
    // 转发协议地址记录
    address[] internal transports;

    constructor() {
        // 第一个元素为无效地址
        transports.push(address(0));
    }

    function set(address transportAddress) public {
        if (addrToIndex[transportAddress] != 0) {
            emit ErrorInfo.ExecuteError(ErrorInfo.TransportProtocolExists, "");
        }

        transports.push(transportAddress);
        uint256 index = transports.length - 1;
        addrToIndex[transportAddress] = index;

        emit TransportRegister(transportAddress, index);
    }

    function get(
        uint256 transportId
    ) public view returns (address, string memory) {
        if (transportId == 0) {
            return (address(0), "transportId must be greater than 0");
        }
        if (transportId > transports.length - 1) {
            return (address(0), "unfound");
        }
        address addr = transports[transportId];
        return (addr, "");
    }

    function getTransportId(
        address transportAddress
    ) public view returns (uint256, string memory) {
        if (addrToIndex[transportAddress] == 0) {
            return (0, "unfound");
        }
        return (addrToIndex[transportAddress], "");
    }
}
