// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/ErrorInfo.sol";

contract AppReg {
    event AppRegister(address indexed appAddr, uint256 indexed appId);

    // 应用地址 -> 应用 ID
    mapping(address => uint256) internal addrToIndex;
    // 应用地址记录
    address[] internal apps;

    constructor() {
        // 第一个元素为无效地址
        apps.push(address(0));
    }

    function set(address appAddress) public {
        if (addrToIndex[appAddress] != 0) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppExists, "");
            return;
        }

        apps.push(appAddress);
        uint256 index = apps.length - 1;
        addrToIndex[appAddress] = index;

        emit AppRegister(appAddress, index);
    }

    function get(uint256 appId) public view returns (address, string memory) {
        if (appId == 0) {
            return (address(0), "appId must be greater than 0");
        }
        if (appId >= apps.length) {
            return (address(0), "unfound");
        }
        address addr = apps[appId];
        return (addr, "");
    }

    function getAppId(
        address appAddress
    ) public view returns (uint256, string memory) {
        if (addrToIndex[appAddress] == 0) {
            return (0, "unfound");
        }
        return (addrToIndex[appAddress], "");
    }
}
