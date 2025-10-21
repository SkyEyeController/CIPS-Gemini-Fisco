// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract HomomorphicCalculator {
    // Paillier同态加密参数
    struct PublicKey {
        uint256 n;      // 模数 n = p*q
        uint256 g;      // 生成元 g
        uint256 nSquared; // n²
    }

    PublicKey public key;

    // 初始化合约，设置公钥参数
    constructor(uint256 _n, uint256 _g) {
        key.n = _n;
        key.g = _g;
        key.nSquared = _n * _n;
    }

    // 同态加法：E(m1) * E(m2) mod n² = E(m1 + m2)
    function homomorphicAdd(uint256 encryptedA, uint256 encryptedB) external view returns (uint256) {
        return (encryptedA * encryptedB) % key.nSquared;
    }

    // 同态乘法：E(m1)^k mod n² = E(m1 * k)
    function homomorphicMultiply(uint256 encryptedA, uint256 scalarK) external view returns (uint256) {
        return modExp(encryptedA, scalarK, key.nSquared);
    }

    // 模幂运算，防止溢出
    function modExp(uint256 base, uint256 exponent, uint256 modulus) internal pure returns (uint256) {
        uint256 result = 1;
        base = base % modulus;
        
        while (exponent > 0) {
            if (exponent % 2 == 1) {
                result = (result * base) % modulus;
            }
            exponent = exponent >> 1;
            base = (base * base) % modulus;
        }
        
        return result;
    }
}