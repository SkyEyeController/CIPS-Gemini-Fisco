// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./layer_app/register.sol";
import "./layer_app/interface.sol";
import "./layer_transaction/register.sol";
import "./layer_transaction/interface.sol";
import "./layer_transmission/register.sol";
import "./layer_transmission/interface.sol";
import "./layer_verification/register.sol";
import "./layer_verification/interface.sol";
import "./layer_transport/register.sol";
import "./layer_transport/interface.sol";
import "./utils/ErrorInfo.sol";
import "./utils/Types.sol";

contract Contract_Aggregator {
    // use `struct`, make `stack too deep` happy
    // 定义变量
    struct ContractState {
        uint256 chainId;
        uint256 seq;
        AppReg appReg;
        TransactionReg transactionReg;
        TransmissionReg transmissionReg;
        VerificationReg verificationReg;
        TransportReg transportReg;
    }
    ContractState private context;

    // 应用层使用
    struct AppType {
        uint256 src_AppId;
        uint256 dst_AppId;
        address appAddr;
        IApp contractApp;
    }

    // 事物层使用
    /*
    struct TransactionType {
        uint256 transactionTypeId;
        address transactionAddr;
        ITransactionProtocol contractTransaction;
    }
    */

    // 传输层使用
    struct TransmissionType {
        uint256 transmissionTypeId;
        address transmissionAddr;
        ITransmissionProtocol contractTransmission;
    }

    // 验证层使用
    struct VerificationType {
        uint256 verificationTypeId;
        address verificationAddr;
        IVericationProtocol contractVerification;
    }

    // 转发层使用
    struct TransportType {
        uint256 transportTypeId;
        address transportAddr;
        ITransportProtocol contractTransport;
    }

    // 构造时 传入参数 建立对应参数
    constructor(
        uint256 _chainId,
        address address_appReg,
        address address_transactionReg,
        address address_transmissionReg,
        address address_verificationReg,
        address address_transportReg
    ) {
        context.chainId = _chainId;
        context.seq = 0;

        context.appReg = AppReg(address_appReg);
        context.transactionReg = TransactionReg(address_transactionReg);
        context.transmissionReg = TransmissionReg(address_transmissionReg);
        context.verificationReg = VerificationReg(address_verificationReg);
        context.transportReg = TransportReg(address_transportReg);
    }

    // 用户向聚合器的 sendMsg 接口发送请求, 由该接口调用应用
    // TODO sender 权限检查
    function sendMsg(
        uint256 dstChainId,
        uint256 srcAppId,
        uint256 dstAppId,
        bytes[] memory appArgs
    ) public {
        string memory err;

        // ****************************** 应用层 ******************************
        AppType memory appVar;
        // 获取应用合约
        // RULE 对于所有层的子协议的标识符, 如果是 0, 则表示不使用该层
        (appVar.appAddr, err) = context.appReg.get(srcAppId);
        if (appVar.appAddr == address(0)) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppNotExists, err);
            return;
        }
        appVar.contractApp = IApp(appVar.appAddr);

        // 确定应用合约有效后, 创建跨链消息的空结构体
        Types.CrosschainMessage memory ccMsg;
        ccMsg.srcChainId = context.chainId;
        ccMsg.dstChainId = dstChainId;
        ccMsg.srcAppId = srcAppId;
        ccMsg.dstAppId = dstAppId;
        ccMsg.ack = false;

        // 执行应用合约
        // TODO 调整执行逻辑, 使得 app 内可以直接发送跨链消息
        // TODO 调整执行逻辑, 使得可以同时发送多笔跨链消息
        Types.CrosschainMessage memory resultApp;
        (resultApp, err) = appVar.contractApp.send_msg(ccMsg, appArgs);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppSendFailed, err);
            return;
        }
        ccMsg = resultApp;

        // ****************************** 事物层 ******************************
        // 事物层 暂时为空

        // ****************************** 传输层 ******************************
        // 获取传输协议
        TransmissionType memory transmissionVar;
        transmissionVar.transmissionTypeId = ccMsg.transmissionTypeId;
        (transmissionVar.transmissionAddr, err) = context.transmissionReg.get(
            transmissionVar.transmissionTypeId
        );
        if (transmissionVar.transmissionAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolNotExists,
                err
            );
            return;
        }
        transmissionVar.contractTransmission = ITransmissionProtocol(
            transmissionVar.transmissionAddr
        );

        // 执行传输协议
        Types.CrosschainMessage memory resultTransmission;
        (resultTransmission, err) = transmissionVar
            .contractTransmission
            .send_msg(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolSendFailed,
                err
            );
            return;
        }
        ccMsg = resultTransmission;

        // ****************************** 验证层 ******************************
        // 获取验证协议
        VerificationType memory verificationVar;
        verificationVar.verificationTypeId = ccMsg.verificationTypeId;
        (verificationVar.verificationAddr, err) = context.verificationReg.get(
            verificationVar.verificationTypeId
        );
        if (verificationVar.verificationAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolNotExists,
                err
            );
            return;
        }
        verificationVar.contractVerification = IVericationProtocol(
            verificationVar.verificationAddr
        );

        // 执行验证协议
        Types.CrosschainMessage memory resultVerification;
        (resultVerification, err) = verificationVar
            .contractVerification
            .prepare(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolPrepareFailed,
                err
            );
            return;
        }
        ccMsg = resultVerification;

        // ****************************** 转发层 ******************************
        // 获取转发协议
        TransportType memory transportVar;
        transportVar.transportTypeId = ccMsg.transportTypeId;
        (transportVar.transportAddr, err) = context.transportReg.get(
            transportVar.transportTypeId
        );
        if (transportVar.transportAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolNotExists,
                err
            );
            return;
        }
        transportVar.contractTransport = ITransportProtocol(
            transportVar.transportAddr
        );

        // RULE 跨链消息的序列号 +1
        //   因为存在不同的跨链转发协议, 所以如果由跨链转发协议来设置 seq, 那么有可能存在相同序列号的跨链消息
        //   所以, 由聚合器来给跨链消息设置一个统一的 seq
        //   转发协议自身也可以设置与协议相关的 transport_seq, 这属于协议自身控制范畴, 无所谓
        // TODO 应该为跨链消息的修改设置权限, 本层协议无权修改与本层协议无关的字段
        // TODO 描述 seq 的必要性
        context.seq += 1;
        ccMsg.seq = context.seq;

        // 执行转发协议
        Types.CrosschainMessage memory resultTransport;
        (resultTransport, err) = transportVar.contractTransport.sendOut(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolSendOutFailed,
                err
            );
            return;
        }
        ccMsg = resultTransport;
    }

    // 目的链接收跨链消息, 分为两个核心过程:
    // 1. 向上传递: 转发层到应用层
    // 2. 向下传递: 应用层到转发层
    // 最终, 转发层将应用层的响应转发出去
    // TODO sender 权限检查
    function receiveMsg(bytes[] memory data) public {
        string memory err;

        // 构造跨链消息结构体, 由转发层赋值该结构体
        Types.CrosschainMessage memory ccMsg;

        // *************************** up 转发层 up ***************************
        // 获取转发协议
        // RULE: data 的第 0 个元素必须是转发层协议的标识
        TransportType memory transportVar;
        transportVar.transportTypeId = Types.bytesToUint256(data[0]);
        (transportVar.transportAddr, err) = context.transportReg.get(
            transportVar.transportTypeId
        );
        if (transportVar.transportAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolNotExists,
                err
            );
            return;
        }
        transportVar.contractTransport = ITransportProtocol(
            transportVar.transportAddr
        );

        // 执行转发协议
        (ccMsg, err) = transportVar.contractTransport.receiveIn(data);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolReceiveInFailed,
                err
            );
            return;
        }
        // 如果 err 为空, 那么说明 receiveIn 函数执行成功, 那么 ccMsg 也已经被赋值

        // *************************** up 验证层 up ***************************
        // 获取验证协议
        VerificationType memory verificationVar;
        verificationVar.verificationTypeId = ccMsg.verificationTypeId;
        (verificationVar.verificationAddr, err) = context.verificationReg.get(
            verificationVar.verificationTypeId
        );
        if (verificationVar.verificationAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolNotExists,
                err
            );
            return;
        }
        verificationVar.contractVerification = IVericationProtocol(
            verificationVar.verificationAddr
        );

        // 执行验证协议
        bool verified;
        (verified, err) = verificationVar.contractVerification.verify(ccMsg);
        if (verified == false) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolVerifyFailed,
                err
            );
            return;
        }

        // *************************** up 传输层 up ***************************
        // 获取传输协议
        TransmissionType memory transmissionVar;
        transmissionVar.transmissionTypeId = ccMsg.transmissionTypeId;
        (transmissionVar.transmissionAddr, err) = context.transmissionReg.get(
            transmissionVar.transmissionTypeId
        );
        if (transmissionVar.transmissionAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolNotExists,
                err
            );
            return;
        }
        transmissionVar.contractTransmission = ITransmissionProtocol(
            transmissionVar.transmissionAddr
        );

        // 执行传输协议
        Types.CrosschainMessage memory resultTransmission;
        (resultTransmission, err) = transmissionVar
            .contractTransmission
            .receive_msg(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolReceiveFailed,
                err
            );
            return;
        }
        ccMsg = resultTransmission;

        // *************************** up 事物层 up ***************************
        // 获取并执行事务协议
        // TODO 先置空, 事务协议的具体设计还需要讨论
        // 事务协议执行结束

        // ****************************** 应用层 ******************************
        // 获取目的链应用
        AppType memory appVar;
        appVar.dst_AppId = ccMsg.dstAppId;
        (appVar.appAddr, err) = context.appReg.get(appVar.dst_AppId);
        if (appVar.appAddr == address(0)) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppNotExists, err);
        }
        appVar.contractApp = IApp(appVar.appAddr);

        // 执行应用逻辑
        Types.CrosschainMessage memory resultApp;
        (resultApp, err) = appVar.contractApp.receive_msg(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppReceiveFailed, err);
            return;
        }
        ccMsg = resultApp;

        // TODO 添加判断, 使得既支持"请求-响应"逻辑, 也支持"仅请求"逻辑

        // 在 receive 函数中,
        // 应用逻辑执行成功后, 应用直接对 ccMsg 赋值 ccMsg.payload_resp, ccMsg.tsctTypeId 等
        // 然后继续向下执行

        // *******************************************************************
        // *******************************************************************
        // ***************************** 向下转发 *****************************
        // *******************************************************************
        // *******************************************************************

        // ************************* down 事物层 down *************************
        // 执行事务协议
        // TODO 先置空
        // 事务协议执行结束

        // ************************* down 传输层 down *************************
        // 向下执行传输协议
        if (ccMsg.transmissionTypeId != transmissionVar.transmissionTypeId) {
            (transmissionVar.transmissionAddr, err) = context
                .transmissionReg
                .get(ccMsg.transmissionTypeId);
            if (transmissionVar.transmissionAddr == address(0)) {
                emit ErrorInfo.ExecuteError(
                    ErrorInfo.TransmissionProtocolNotExists,
                    err
                );
                return;
            }
            transmissionVar.contractTransmission = ITransmissionProtocol(
                transmissionVar.transmissionAddr
            );
        }
        (resultTransmission, err) = transmissionVar
            .contractTransmission
            .response(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolResponseFailed,
                err
            );
            return;
        }
        ccMsg = resultTransmission;

        // ************************* down 验证层 down *************************
        // 向下执行验证协议
        if (ccMsg.verificationTypeId != verificationVar.verificationTypeId) {
            (verificationVar.verificationAddr, err) = context
                .verificationReg
                .get(ccMsg.verificationTypeId);
            if (verificationVar.verificationAddr == address(0)) {
                emit ErrorInfo.ExecuteError(
                    ErrorInfo.VerificationProtocolNotExists,
                    err
                );
                return;
            }
            verificationVar.contractVerification = IVericationProtocol(
                verificationVar.verificationAddr
            );
        }
        Types.CrosschainMessage memory resultVerification;
        (resultVerification, err) = verificationVar
            .contractVerification
            .response(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolResponseFailed,
                err
            );
            return;
        }
        ccMsg = resultVerification;

        // ************************* down 转发层 down *************************
        // 向下执行转发协议
        if (ccMsg.transportTypeId != transportVar.transportTypeId) {
            (transportVar.transportAddr, err) = context.transportReg.get(
                ccMsg.transportTypeId
            );
            if (transportVar.transportAddr == address(0)) {
                emit ErrorInfo.ExecuteError(
                    ErrorInfo.TransportProtocolNotExists,
                    err
                );
                return;
            }
            transportVar.contractTransport = ITransportProtocol(
                transportVar.transportAddr
            );
        }
        Types.CrosschainMessage memory resultTransport;
        (resultTransport, err) = transportVar.contractTransport.response(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolResponseFailed,
                err
            );
            return;
        }
        ccMsg = resultTransport;
    }

    // 源链接收跨链消息, 只需要向上传递即可, 即从转发层传递到应用层
    // 如果跨链应用需要发起新的跨链消息, 则跨链应用直接调用send接口即可
    // TODO 添加send接口, 支持跨链应用继续发起跨链消息
    // TODO sender 权限检查
    // TODO 描述聚合器、注册器、子协议间的主权问题
    function acknowledgeMsg(bytes[] memory data) public {
        string memory err;

        // 构造跨链消息结构体, 由转发层赋值该结构体
        Types.CrosschainMessage memory ccMsg;

        // *************************** up 转发层 up ***************************
        // 获取转发协议
        // RULE: data 的第 0 个元素必须是转发层协议的标识
        TransportType memory transportVar;
        transportVar.transportTypeId = Types.bytesToUint256(data[0]);
        (transportVar.transportAddr, err) = context.transportReg.get(
            transportVar.transportTypeId
        );
        if (transportVar.transportAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolNotExists,
                err
            );
            return;
        }
        transportVar.contractTransport = ITransportProtocol(
            transportVar.transportAddr
        );

        // 执行转发协议
        (ccMsg, err) = transportVar.contractTransport.acknowledge(data);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransportProtocolAcknowledgeFailed,
                err
            );
            return;
        }
        // 如果 err 为空, 那么说明 acknowledge 函数执行成功, 那么 ccMsg 也已经被赋值

        // *************************** up 验证层 up ***************************
        // 获取验证协议
        VerificationType memory verificationVar;
        verificationVar.verificationTypeId = ccMsg.verificationTypeId;
        (verificationVar.verificationAddr, err) = context.verificationReg.get(
            verificationVar.verificationTypeId
        );
        if (verificationVar.verificationAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolNotExists,
                err
            );
            return;
        }
        verificationVar.contractVerification = IVericationProtocol(
            verificationVar.verificationAddr
        );

        // 执行验证协议
        bool verified;
        (verified, err) = verificationVar.contractVerification.verify(ccMsg);
        if (verified == false) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.VerificationProtocolVerifyFailed,
                err
            );
            return;
        }
        // 如果 ok, 那么说明跨链消息的有效性是成立的

        // *************************** up 传输层 up ***************************
        // 获取传输协议
        TransmissionType memory transmissionVar;
        transmissionVar.transmissionTypeId = ccMsg.transmissionTypeId;
        (transmissionVar.transmissionAddr, err) = context.transmissionReg.get(
            transmissionVar.transmissionTypeId
        );
        if (transmissionVar.transmissionAddr == address(0)) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolNotExists,
                err
            );
            return;
        }
        transmissionVar.contractTransmission = ITransmissionProtocol(
            transmissionVar.transmissionAddr
        );

        // 执行传输协议
        Types.CrosschainMessage memory resultTransmission;
        (resultTransmission, err) = transmissionVar
            .contractTransmission
            .acknowledge(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(
                ErrorInfo.TransmissionProtocolAcknowledgeFailed,
                err
            );
            return;
        }
        ccMsg = resultTransmission;

        // *************************** up 事物层 up ***************************
        // 获取并执行事务协议
        // TODO 先置空, 事务协议的具体设计还需要讨论
        // 事务协议执行结束

        // ****************************** 应用层 ******************************
        // 获取应用合约
        AppType memory appVar;
        appVar.src_AppId = ccMsg.srcAppId;
        (appVar.appAddr, err) = context.appReg.get(appVar.src_AppId);
        if (appVar.appAddr == address(0)) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppNotExists, err);
            return;
        }
        appVar.contractApp = IApp(appVar.appAddr);

        // 执行应用合约
        Types.CrosschainMessage memory resultApp;
        (resultApp, err) = appVar.contractApp.acknowledge(ccMsg);
        if (bytes(err).length != 0) {
            emit ErrorInfo.ExecuteError(ErrorInfo.AppAcknowledgeFailed, err);
            return;
        }
        ccMsg = resultApp;
    }
}
