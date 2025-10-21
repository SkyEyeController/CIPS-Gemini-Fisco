// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/ErrorInfo.sol";

contract TransactionReg {
    event TransactionRegister(
        address indexed transactionAddr,
        uint256 indexed transactionId
    );

    mapping(address => uint256) internal addrToIndex;
    address[] internal transactions;

    constructor() {
        // 第一个元素为无效地址
        transactions.push(address(0));
    }

    function set(address transactionAddress) public {
        if (addrToIndex[transactionAddress] != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransactionProtocolExists,
                ""
            );
            return;
        }

        transactions.push(transactionAddress);
        uint256 index = transactions.length - 1;
        addrToIndex[transactionAddress] = index;

        emit TransactionRegister(transactionAddress, index);
    }

    function get(
        uint256 transactionId
    ) public view returns (address, string memory) {
        if (transactionId == 0) {
            return (address(0), "transactionId must be greater than 0");
        }
        if (transactionId > transactions.length - 1) {
            return (address(0), "unfound");
        }
        address addr = transactions[transactionId];
        return (addr, "");
    }

    function getTransactionId(
        address transactionAddress
    ) public view returns (uint256, string memory) {
        if (addrToIndex[transactionAddress] == 0) {
            return (0, "unfound");
        }
        return (addrToIndex[transactionAddress], "");
    }
}
