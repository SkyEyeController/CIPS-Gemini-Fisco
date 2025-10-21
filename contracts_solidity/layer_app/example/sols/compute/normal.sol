// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Calculator {
    // 执行四则运算（包含安全检查）
    function calculate(uint256 a, uint256 b) external pure returns (
        uint256 sum,
        uint256 difference,
        uint256 product,
        uint256 quotient
    ) {
        // 加法
        sum = a + b;
        
        // 减法（防止下溢）
        require(a >= b, "Subtraction underflow");
        difference = a - b;
        
        // 乘法
        product = a * b;
        
        // 除法（防止除以零）
        require(b != 0, "Division by zero");
        quotient = a / b;
    }
}