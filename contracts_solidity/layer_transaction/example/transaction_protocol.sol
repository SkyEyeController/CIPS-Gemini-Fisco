// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interface.sol";
import "../../utils/Types.sol";

contract TransactionProtocol is ITransactionProtocol {
    function work(
        uint8 flag,
        Types.CrosschainMessage memory ccMsg,
        bytes[] memory payload
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        return (ccMsg, "");
    }
}
