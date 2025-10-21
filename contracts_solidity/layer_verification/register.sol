// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/ErrorInfo.sol";

contract VerificationReg {
    event VerificationRegister(
        address indexed verificationAddress,
        uint256 indexed verificationId
    );

    // 验证协议地址 -> 验证协议 ID
    mapping(address => uint256) internal addrToIndex;
    // 验证协议地址记录
    address[] internal verifications;

    constructor() {
        // 第一个元素为无效地址
        verifications.push(address(0));
    }

    function set(address verificationAddress) public {
        if (addrToIndex[verificationAddress] != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolExists,
                ""
            );
        }

        verifications.push(verificationAddress);
        uint256 index = verifications.length - 1;
        addrToIndex[verificationAddress] = index;

        emit VerificationRegister(verificationAddress, index);
    }

    function get(
        uint256 verificationId
    ) public view returns (address, string memory) {
        if (verificationId == 0) {
            return (address(0), "verificationId must be greater than 0");
        }
        if (verificationId > verifications.length - 1) {
            return (address(0), "unfound");
        }
        address addr = verifications[verificationId];
        return (addr, "");
    }

    function getVerificationId(
        address verificationAddress
    ) public view returns (uint256, string memory) {
        if (addrToIndex[verificationAddress] == 0) {
            return (0, "unfound");
        }
        return (addrToIndex[verificationAddress], "");
    }
}
