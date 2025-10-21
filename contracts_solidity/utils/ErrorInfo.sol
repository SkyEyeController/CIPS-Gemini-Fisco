// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

library ErrorInfo {
    event ExecuteError(uint256 indexed code, string message);

    uint8 constant AppExists = 101;
    uint8 constant AppNotExists = 102;
    uint8 constant AppSendFailed = 103;
    uint8 constant AppReceiveFailed = 104;
    uint8 constant AppAcknowledgeFailed = 105;

    uint8 constant TransactionProtocolExists = 111;
    uint8 constant TransactionProtocolNotExists = 112;
    uint8 constant TransactionProtocolExecuteFailed = 113;

    uint8 constant TransmissionProtocolExists = 121;
    uint8 constant TransmissionProtocolNotExists = 122;
    uint8 constant TransmissionProtocolSendFailed = 123;
    uint8 constant TransmissionProtocolReceiveFailed = 124;
    uint8 constant TransmissionProtocolResponseFailed = 125;
    uint8 constant TransmissionProtocolAcknowledgeFailed = 126;

    uint8 constant VerificationProtocolExists = 131;
    uint8 constant VerificationProtocolNotExists = 132;
    uint8 constant VerificationProtocolPrepareFailed = 133;
    uint8 constant VerificationProtocolUpdateFailed = 134;
    uint8 constant VerificationProtocolVerifyFailed = 135;
    uint8 constant VerificationProtocolResponseFailed = 135;

    uint8 constant TransportProtocolExists = 141;
    uint8 constant TransportProtocolNotExists = 142;
    uint8 constant TransportProtocolSendOutFailed = 143;
    uint8 constant TransportProtocolReceiveInFailed = 144;
    uint8 constant TransportProtocolResponseFailed = 145;
    uint8 constant TransportProtocolAcknowledgeFailed = 146;
}
