// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/ErrorInfo.sol";

contract TransmissionReg {
    event TransmissionRegister(
        address indexed transmissionAddress,
        uint256 indexed transmissionId
    );

    // 传输协议地址 -> 传输协议 ID
    mapping(address => uint256) internal addrToIndex;
    // 传输协议地址记录
    address[] internal tranmissions;

    constructor() {
        // 第一个元素为无效地址
        tranmissions.push(address(0));
    }

    function set(address transmissionAddress) public {
        if (addrToIndex[transmissionAddress] != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolExists,
                ""
            );
            return;
        }

        tranmissions.push(transmissionAddress);
        uint256 index = tranmissions.length - 1;
        addrToIndex[transmissionAddress] = index;

        emit TransmissionRegister(transmissionAddress, index);
    }
    function get(
        uint256 transmissionId
    ) public view returns (address, string memory) {
        if (transmissionId == 0) {
            return (address(0), "transmissionId must be greater than 0");
        }
        if (transmissionId > tranmissions.length - 1) {
            return (address(0), "unfound");
        }
        address addr = tranmissions[transmissionId];
        return (addr, "");
    }

    function getTransmissionId(
        address transmissionAddress
    ) public view returns (uint256, string memory) {
        if (addrToIndex[transmissionAddress] == 0) {
            return (0, "unfound");
        }
        return (addrToIndex[transmissionAddress], "");
    }
}
