// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interface.sol";
import "../../utils/Types.sol";

// 仅导入当前存在的合约文件并按实际 contract 名重命名
import {UniversalKVStore as KVCrossContract} from "./sols/kv-cross.sol";

import {DataStorage as StorageNormalContract} from "./sols/storage/normal.sol";
import {EnhancedDataStorage as StorageEncrptContract} from "./sols/storage/encrpt.sol";

import {DataQuery as QueryNormalContract} from "./sols/query/normal.sol";
import {PrivateDataQuery as QueryEncrptContract} from "./sols/query/encrpt.sol";

import {Calculator as ComputeNormalContract} from "./sols/compute/normal.sol";
import {HomomorphicCalculator as ComputeEncrptContract} from "./sols/compute/encrpt.sol";

contract App is IApp {
    // 合约实例注册表（使用明文名称作为键）
    mapping(string => address) private contractRegistry;

    // 实例变量（只保留存在的合约类型）
    KVCrossContract private kvCross;

    StorageNormalContract private storageNormal;
    StorageEncrptContract private storageEncrpt;

    QueryNormalContract private queryNormal;
    QueryEncrptContract private queryEncrpt;

    ComputeNormalContract private computeNormal;
    ComputeEncrptContract private computeEncrpt;

    event ContractCalled(
        string contractName,
        string functionName,
        bytes[] params
    );
    event Msgsendsuccess(bytes[] payload);
    // 新增：路由错误事件
    event RouteError(string contractName, string functionName, string reason);

    constructor() {
        // 在构造中 new 实例（只对存在的合约）
        kvCross = new KVCrossContract();

        storageNormal = new StorageNormalContract();
        storageEncrpt = new StorageEncrptContract();

        queryNormal = new QueryNormalContract();
        queryEncrpt = new QueryEncrptContract();

        computeNormal = new ComputeNormalContract();
        computeEncrpt = new ComputeEncrptContract(
            6290536192560178763,
            6290536192560178764
        ); // 示例公钥参数

        // 注册地址（与现有合约对应的名称）
        contractRegistry["kv-cross"] = address(kvCross);

        contractRegistry["DataStorage"] = address(storageNormal);
        contractRegistry["EnhancedDataStorage"] = address(storageEncrpt);

        contractRegistry["DataQuery"] = address(queryNormal);
        contractRegistry["PrivateDataQuery"] = address(queryEncrpt);

        contractRegistry["Calculator"] = address(computeNormal);
        contractRegistry["HomomorphicCalculator"] = address(computeEncrpt);
    }

    // 与 IApp 接口匹配的 send_msg
    function send_msg(
        Types.CrosschainMessage memory ccMsg,
        bytes[] memory payload
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        ccMsg.transactionTypeId = 1;
        ccMsg.transmissionTypeId = 1;
        ccMsg.verificationTypeId = 1;
        ccMsg.transportTypeId = 1;
        ccMsg.payloadReq = payload;

        if (ccMsg.payloadReq.length > 0) {
            emit Msgsendsuccess(ccMsg.payloadReq);
        }
        return (ccMsg, "");
    }

    /**
     * payloadReq 结构：
     * [0] = 合约名称 (例如 "DataStorage" / "kv-cross")
     * [1] = 函数名 (例如 "set" / "get" / "query" / "compute")
     * [2...] = 函数参数（每项为 bytes）
     */
    function receive_msg(
        Types.CrosschainMessage memory ccMsg
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        require(
            ccMsg.payloadReq.length >= 2,
            "Invalid payload: need at least 2 elements"
        );

        string memory contractName = string(ccMsg.payloadReq[0]);
        string memory functionName = string(ccMsg.payloadReq[1]);

        address target = contractRegistry[contractName];
        //require(target != address(0), "Contract not registered");

        emit ContractCalled(contractName, functionName, ccMsg.payloadReq);

        bytes memory result = routeCall(
            contractName,
            target,
            functionName,
            ccMsg.payloadReq
        );

        ccMsg.payloadReq = new bytes[](0); // 清空请求负载

        ccMsg.payloadResp = new bytes[](1);
        ccMsg.payloadResp[0] = result;

        return (ccMsg, "");
    }

    // 路由总入口（按合约名分流）
    function routeCall(
        string memory contractName,
        address target,
        string memory functionName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 nameHash = keccak256(bytes(contractName));

        if (nameHash == keccak256(bytes("kv-cross"))) {
            return routeKVCrossCall(functionName, params);
        }

        // storage 系列 - 区分 normal 和 encrpt
        if (nameHash == keccak256(bytes("DataStorage"))) {
            return routeStorageNormalCall(target, functionName, params);
        }
        if (nameHash == keccak256(bytes("EnhancedDataStorage"))) {
            return routeStorageEncrptCall(target, functionName, params);
        }

        // query 系列 - 区分 normal 和 encrpt
        if (nameHash == keccak256(bytes("DataQuery"))) {
            return routeQueryNormalCall(target, functionName, params);
        }
        if (nameHash == keccak256(bytes("PrivateDataQuery"))) {
            return routeQueryEncrptCall(target, functionName, params);
        }

        // compute 系列 - 区分 normal 和 encrpt
        if (nameHash == keccak256(bytes("Calculator"))) {
            return routeComputeNormalCall(target, functionName, params);
        }
        if (nameHash == keccak256(bytes("HomomorphicCalculator"))) {
            return routeComputeEncrptCall(target, functionName, params);
        }

        revert("Unknown contract");
    }

    // kv-cross 的路由（直接调用 kvCross 实例）
    function routeKVCrossCall(
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        if (fHash == keccak256(bytes("set"))) {
            require(params.length >= 4, "set requires key and value");
            kvCross.set(string(params[2]), string(params[3]));
            return bytes("Accepted");
        }

        if (fHash == keccak256(bytes("get"))) {
            require(params.length >= 3, "get requires key");
            string memory v = kvCross.get(string(params[2]));
            return bytes(v);
        }

        revert("Unknown kv-cross function");
    }
    // 将 bytes 里的十进制字符串转为 uint256；若是 32 字节则按 ABI 解码
    function parseUint256(bytes memory data) private pure returns (uint256) {
        if (data.length == 32) {
            return abi.decode(data, (uint256));
        }
        uint256 result = 0;
        for (uint256 i = 0; i < data.length; i++) {
            uint8 c = uint8(data[i]);
            require(c >= 48 && c <= 57, "Invalid uint string");
            result = result * 10 + (c - 48);
        }
        return result;
    }
    event LogBytes(bytes data);
    function decodeHextoAddress(
        bytes memory hexEncodedAddress
    ) public returns (address) {
        // 1. 将 bytes 转为 ASCII 字符串
        emit LogBytes(hexEncodedAddress);
        string memory asciiEncodedAddress = string(hexEncodedAddress);
        bytes memory addressBytes = bytes(asciiEncodedAddress);
        require(addressBytes.length >= 42, "Invalid encoded address length");

        // 2. 手动提取 40 位地址部分（跳过 "0x"）
        bytes memory extractedAddress = new bytes(40);
        for (uint i = 2; i < 42; i++) {
            extractedAddress[i - 2] = addressBytes[i];
        }

        // 3. 转换为 address 类型
        return parseAddress(string(extractedAddress));
    }

    // 辅助函数：将 40 位十六进制字符串转为 address
    function parseAddress(string memory _a) internal pure returns (address) {
        bytes memory tmp = bytes(_a);
        uint160 iaddr = 0;
        for (uint i = 0; i < 40; i++) {
            iaddr *= 16;
            if (tmp[i] >= 0x30 && tmp[i] <= 0x39) {
                iaddr += uint8(tmp[i]) - 0x30;
            } else if (tmp[i] >= 0x41 && tmp[i] <= 0x46) {
                iaddr += uint8(tmp[i]) - 0x37;
            } else if (tmp[i] >= 0x61 && tmp[i] <= 0x66) {
                iaddr += uint8(tmp[i]) - 0x57;
            } else {
                revert("Invalid address character");
            }
        }
        return address(iaddr);
    }

    // Storage Normal 路由 - DataStorage 合约
    function routeStorageNormalCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));
        StorageNormalContract s = StorageNormalContract(target); // 使用注册表中的地址

        if (fHash == keccak256(bytes("setData"))) {
            require(params.length >= 4, "setData requires key and value");
            string memory key = string(params[2]);
            uint256 value = parseUint256(params[3]); // <-- 改为容错解析
            // 用 try/catch 捕获下游合约的 revert
            try s.setData(key, value) {
                return bytes("setData success");
            } catch Error(string memory reason) {
                emit RouteError("DataStorage", "setData", reason);
                revert(reason);
            } catch (bytes memory /*lowLevelData*/) {
                emit RouteError("DataStorage", "setData", "low-level error");
                revert("DataStorage.setData failed");
            }
        }

        if (fHash == keccak256(bytes("getData"))) {
            require(params.length >= 3, "getData requires key");
            string memory key = string(params[2]);
            try s.getData(key) returns (uint256 data) {
                return abi.encode(data);
            } catch Error(string memory reason) {
                emit RouteError("DataStorage", "getData", reason);
                revert(reason);
            } catch (bytes memory /*lowLevelData*/) {
                emit RouteError("DataStorage", "getData", "low-level error");
                revert("DataStorage.getData failed");
            }
        }

        if (fHash == keccak256(bytes("keyExists"))) {
            require(params.length >= 3, "keyExists requires key");
            string memory key = string(params[2]);
            try s.keyExists(key) returns (bool exists) {
                return abi.encode(exists);
            } catch Error(string memory reason) {
                emit RouteError("DataStorage", "keyExists", reason);
                revert(reason);
            } catch (bytes memory /*lowLevelData*/) {
                emit RouteError("DataStorage", "keyExists", "low-level error");
                revert("DataStorage.keyExists failed");
            }
        }

        revert("Unknown DataStorage function");
    }

    // Storage Encrpt 路由 - EnhancedDataStorage 合约
    function routeStorageEncrptCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        if (fHash == keccak256(bytes("setPlainData"))) {
            require(params.length >= 4, "setPlainData requires key and value");
            string memory key = string(params[2]);
            uint256 value = parseUint256(params[3]); // <-- 改为容错解析
            storageEncrpt.setPlainData(key, value);
            return bytes("setPlainData success");
        }

        if (fHash == keccak256(bytes("setEncryptedData"))) {
            require(
                params.length >= 4,
                "setEncryptedData requires key and cipherValue"
            );
            string memory key = string(params[2]);
            uint256 cipherValue = parseUint256(params[3]); // <-- 改为容错解析
            storageEncrpt.setEncryptedData(key, cipherValue);
            return bytes("setEncryptedData success");
        }

        if (fHash == keccak256(bytes("updateTag"))) {
            require(params.length >= 4, "updateTag requires key and newTag");
            string memory key = string(params[2]);
            uint256 newTag = parseUint256(params[3]); // <-- 改为容错解析
            storageEncrpt.updateTag(key, newTag);
            return bytes("updateTag success");
        }

        if (fHash == keccak256(bytes("getData"))) {
            require(params.length >= 3, "getData requires key");
            string memory key = string(params[2]);
            uint256 data = storageEncrpt.getData(key);
            return abi.encode(data);
        }

        if (fHash == keccak256(bytes("getDataRecord"))) {
            require(params.length >= 3, "getDataRecord requires key");
            string memory key = string(params[2]);
            (uint256 value, uint256 tag) = storageEncrpt.getDataRecord(key);
            return abi.encode(value, tag);
        }

        if (fHash == keccak256(bytes("keyExists"))) {
            require(params.length >= 3, "keyExists requires key");
            string memory key = string(params[2]);
            bool exists = storageEncrpt.keyExists(key); // <-- 修正：原来误用 storageNormal
            return abi.encode(exists);
        }

        revert("Unknown EnhancedDataStorage function");
    }

    // Query Normal 路由 - DataQuery 合约
    function routeQueryNormalCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        if (fHash == keccak256(bytes("getValue"))) {
            //return bytes("1008611");
            require(params.length >= 4, "getValue requires 4 params");
            
            // 解析合约地址
            address contractAddress;
            try this.decodeHextoAddress(params[2]) returns (
                address decodedAddress
            ) {
                contractAddress = decodedAddress;
            } catch {
                revert("Invalid contractAddress format");
            }

            // 解析键
            string memory key = string(params[3]);

            // 调用外部合约
            uint256 value;
            try queryNormal.getValue(contractAddress, key) returns (uint256 v) {
                value = v;
            } catch (bytes memory) {
                revert("getValue call failed");
            }
            return abi.encode(value);
        }

        revert("Unknown DataQuery function");
    }

    // Query Encrpt 路由 - PrivateDataQuery 合约
    function routeQueryEncrptCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        if (fHash == keccak256(bytes("queryEncryptedData"))) {
            require(
                params.length >= 4,
                "queryEncryptedData requires storageContract and dataHash"
            );
            address storageContract = decodeHextoAddress(params[2]); // 修改：只接受 0x 字符串
            string memory key = string(params[3]);
            (uint256 value, uint256 tag) = queryEncrpt.queryEncryptedData(
                storageContract,
                key
            );
            return abi.encode(value, tag);
        }

        revert("Unknown PrivateDataQuery function");
    }

    // Compute Normal 路由 - Calculator 合约
    function routeComputeNormalCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        // calculate(uint256 a, uint256 b) returns (uint256 sum, uint256 difference, uint256 product, uint256 quotient)
        if (fHash == keccak256(bytes("calculate"))) {
            require(params.length >= 4, "calculate requires two numbers");
            uint256 a = parseUint256(params[2]); // <-- 改为容错解析
            uint256 b = parseUint256(params[3]); // <-- 改为容错解析
            (
                uint256 sum,
                uint256 difference,
                uint256 product,
                uint256 quotient
            ) = computeNormal.calculate(a, b);
            return abi.encode(sum, difference, product, quotient);
        }

        revert("Unknown Calculator function");
    }

    // Compute Encrpt 路由
    function routeComputeEncrptCall(
        address target,
        string memory funcName,
        bytes[] memory params
    ) private returns (bytes memory) {
        bytes32 fHash = keccak256(bytes(funcName));

        // homomorphicAdd(uint256 encryptedA, uint256 encryptedB) returns (uint256)
        if (fHash == keccak256(bytes("homomorphicAdd"))) {
            require(
                params.length >= 4,
                "homomorphicAdd requires two encrypted values"
            );
            uint256 encryptedA = parseUint256(params[2]); // <-- 改为容错解析
            uint256 encryptedB = parseUint256(params[3]); // <-- 改为容错解析
            uint256 result = computeEncrpt.homomorphicAdd(
                encryptedA,
                encryptedB
            );
            return abi.encode(result);
        }

        // homomorphicMultiply(uint256 encryptedA, uint256 scalarK) returns (uint256)
        if (fHash == keccak256(bytes("homomorphicMultiply"))) {
            require(
                params.length >= 4,
                "homomorphicMultiply requires encrypted value and scalar"
            );
            uint256 encryptedA = parseUint256(params[2]); // <-- 改为容错解析
            uint256 scalarK = parseUint256(params[3]); // <-- 改为容错解析
            uint256 result = computeEncrpt.homomorphicMultiply(
                encryptedA,
                scalarK
            );
            return abi.encode(result);
        }

        revert("Unknown HomomorphicCalculator function");
    }

    function acknowledge(
        Types.CrosschainMessage memory ccMsg
    )
        external
        override
        returns (Types.CrosschainMessage memory, string memory)
    {
        return (ccMsg, "");
    }

    // 注册/查询合约地址
    function registerContract(string memory name, address addr) external {
        require(addr != address(0), "zero address");
        contractRegistry[name] = addr;
    }

    function getContractAddress(
        string memory name
    ) external view returns (address) {
        return contractRegistry[name];
    }
}
