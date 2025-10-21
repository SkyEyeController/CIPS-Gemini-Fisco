// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../utils/Types.sol";

interface IApp {
    function send_msg(
        Types.CrosschainMessage memory ccMsg,
        bytes[] memory payload
    ) external returns (Types.CrosschainMessage memory, string memory);

    function receive_msg(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);

    function acknowledge(
        Types.CrosschainMessage memory ccMsg
    ) external returns (Types.CrosschainMessage memory, string memory);
}
