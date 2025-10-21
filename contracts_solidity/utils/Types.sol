// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

library Types {
    struct CrosschainMessage {
        //公用
        uint256 srcChainId;
        uint256 dstChainId;
        uint256 seq;
        //应用层
        uint256 srcAppId;
        uint256 dstAppId;
        bytes[] payloadReq;
        bytes[] payloadResp;
        //事物层
        uint256 transactionTypeId;
        bytes[] transactionPayload;
        //传输层
        uint256 transmissionTypeId;
        bytes[] transmissionPayload;
        //验证层
        uint256 verificationTypeId;
        bytes[] verificationPayload;
        //转发层
        uint256 transportTypeId;
        bytes[] transportPayload;
        //请求向跨链消息的哈希值, 在传递的过程中不会赋值, 仅在目的链上用于消息索引, 该值不加入哈希计
        bytes32 hashReq;
        //跨链消息的hash
        bytes32 hashResp;
        bool ack;
    }

    // 将bytes类型转换为uint256，不使用内联汇编
    function bytesToUint256(bytes memory b) internal pure returns (uint256) {
        require(b.length >= 32, "Bytes array too short.");
        uint256 converted;
        for (uint i = 0; i < 32; i++) {
            converted = (converted << 8) | uint256(uint8(b[i]));
        }
        return converted;
    }

    // 辅助函数用于将bytes32转换为bytes
    function bytes32ToBytes(bytes32 data) internal pure returns (bytes memory) {
        // 初始化一个固定长度的bytes数组
        bytes memory result = new bytes(32);
        // 将bytes32数据复制到bytes数组
        for (uint i = 0; i < 32; i++) {
            result[i] = data[i];
        }
        return result;
    }

    // 辅助函数 用于 将uint256 转换为 bytes 类型 并且 长度为32
    function uint256To32Bytes(
        uint256 input
    ) internal pure returns (bytes memory) {
        bytes32 input_32btyes = bytes32(input);
        return bytes32ToBytes(input_32btyes);
    }

    // 辅助函数 用于 将bool 转换为 bytes 类型 并且长度为1
    function boolToBytes(
        bool value
    ) internal pure returns (bytes memory result) {
        result = new bytes(1); // 创建一个长度为1的bytes数组
        result[0] = value ? bytes1(0x01) : bytes1(0x00); // 根据bool值设置数组的第一个元素
        return result;
    }

    // 辅助函数 用于 将bytes 转换为 bool 类型
    function bytesToBool(bytes memory b) internal pure returns (bool result) {
        require(b.length > 0, "Input bytes array is empty");
        // 检查第一个字节不为0即可认为是true，否则是false
        result = b[0] != 0x00;
    }

    // 辅助函数 用于 将bytes[] 从 cm 中 还原回来 第一个是长度 后面的依次读回来
    // 输入初始位置 返回结束位置
    function convert_lb(
        bytes[] memory data,
        uint256 startIndex
    ) internal pure returns (bytes[] memory) {
        uint256 len;
        len = bytesToUint256(data[startIndex]);
        bytes[] memory result = new bytes[](len);
        for (uint256 i = 0; i < len; i++) {
            result[i] = data[startIndex + 1 + i];
        }
        return (result);
    }

    // 转换跨链消息 为 hash
    function cmToHash(
        CrosschainMessage memory ccMsg
    ) internal pure returns (bytes32) {
        bytes[] memory elems = cmToLB(ccMsg);

        bytes memory concatenatedBytes;
        for (uint i = 0; i < elems.length; i++) {
            concatenatedBytes = abi.encodePacked(concatenatedBytes, elems[i]);
        }

        return sha256(concatenatedBytes);
    }

    // 转换跨链消息 为 list[bytes]
    function cmToLB(
        CrosschainMessage memory ccMsg
    ) internal pure returns (bytes[] memory) {
        uint256 len;
        len =
            16 +
            ccMsg.transportPayload.length +
            ccMsg.verificationPayload.length +
            ccMsg.transmissionPayload.length +
            ccMsg.transactionPayload.length +
            ccMsg.payloadReq.length +
            ccMsg.payloadResp.length;
        //声明list
        bytes[] memory list = new bytes[](len);
        uint256 i = 0;

        // 转发层
        list[i++] = uint256To32Bytes(ccMsg.transportTypeId);
        //对于payload中，所有元素展开放在列表中, 并在首元素前放置该字段的长度
        list[i++] = bytes32ToBytes(bytes32(ccMsg.transportPayload.length));
        for (uint256 j = 0; j < ccMsg.transportPayload.length; j++) {
            list[i++] = ccMsg.transportPayload[j++];
        }

        // 验证层
        list[i++] = uint256To32Bytes(ccMsg.verificationTypeId);
        list[i++] = bytes32ToBytes(bytes32(ccMsg.verificationPayload.length));
        for (uint256 j = 0; j < ccMsg.verificationPayload.length; j++) {
            list[i++] = ccMsg.verificationPayload[j++];
        }

        //传输层
        list[i++] = uint256To32Bytes(ccMsg.transmissionTypeId);
        list[i++] = bytes32ToBytes(bytes32(ccMsg.transmissionPayload.length));
        for (uint256 j = 0; j < ccMsg.transmissionPayload.length; j++) {
            list[i++] = ccMsg.transmissionPayload[j++];
        }

        //事物层
        list[i++] = uint256To32Bytes(ccMsg.transactionTypeId);
        list[i++] = bytes32ToBytes(bytes32(ccMsg.transactionPayload.length));
        for (uint256 j = 0; j < ccMsg.transactionPayload.length; j++) {
            list[i++] = ccMsg.transactionPayload[j++];
        }

        // 应用层
        list[i++] = uint256To32Bytes(ccMsg.srcAppId);
        list[i++] = uint256To32Bytes(ccMsg.dstAppId);
        list[i++] = bytes32ToBytes(bytes32(ccMsg.payloadReq.length));
        for (uint256 j = 0; j < ccMsg.payloadReq.length; j++) {
            list[i++] = ccMsg.payloadReq[j++];
        }
        list[i++] = bytes32ToBytes(bytes32(ccMsg.payloadResp.length));
        for (uint256 j = 0; j < ccMsg.payloadResp.length; j++) {
            list[i++] = ccMsg.payloadResp[j++];
        }
        // 公共
        list[i++] = uint256To32Bytes(ccMsg.srcChainId);
        list[i++] = uint256To32Bytes(ccMsg.dstChainId);
        list[i++] = uint256To32Bytes(ccMsg.seq);
        list[i++] = boolToBytes(ccMsg.ack);
        return list;
    }

    // 转换list[bytes] 为 跨链消息
    function cmFromLB(
        bytes[] memory data
    ) internal pure returns (CrosschainMessage memory) {
        Types.CrosschainMessage memory cm;
        // 偏移量
        uint256 endIndex = 0;
        //转发层
        cm.transportTypeId = bytesToUint256(data[0]);
        cm.transportPayload = convert_lb(data, 1);
        endIndex = endIndex + 1 + bytesToUint256(data[1]);

        // 验证层
        cm.verificationTypeId = bytesToUint256(data[endIndex + 1]);
        cm.verificationPayload = convert_lb(data, endIndex + 2);
        endIndex = endIndex + 2 + bytesToUint256(data[endIndex + 2]);

        // 传输层
        cm.transmissionTypeId = bytesToUint256(data[endIndex + 1]);
        cm.transmissionPayload = convert_lb(data, endIndex + 2);
        endIndex = endIndex + 2 + bytesToUint256(data[endIndex + 2]);

        // 事务层
        cm.transactionTypeId = bytesToUint256(data[endIndex + 1]);
        cm.transactionPayload = convert_lb(data, endIndex + 2);
        endIndex = endIndex + 2 + bytesToUint256(data[endIndex + 2]);

        // 应用层
        cm.srcAppId = bytesToUint256(data[endIndex + 1]);
        cm.dstAppId = bytesToUint256(data[endIndex + 2]);
        cm.payloadReq = convert_lb(data, endIndex + 3);
        endIndex = endIndex + 3 + bytesToUint256(data[endIndex + 3]);
        cm.payloadResp = convert_lb(data, endIndex + 1);
        endIndex = endIndex + 1 + bytesToUint256(data[endIndex + 1]);

        // 公共
        cm.srcChainId = bytesToUint256(data[endIndex + 1]);
        cm.dstChainId = bytesToUint256(data[endIndex + 2]);
        cm.seq = bytesToUint256(data[endIndex + 3]);
        cm.ack = bytesToBool(data[endIndex + 4]);
        return cm;
    }
}
