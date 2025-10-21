import random
from Crypto.Util import number

# 生成1024位的Paillier密钥对
def generate_paillier_keys(key_length=64):
    p = number.getPrime(key_length // 2)
    q = number.getPrime(key_length // 2)
    n = p * q
    λ = (p - 1) * (q - 1)
    g = n + 1  # 简化的g值选择
    μ = pow(λ, -1, n)  # 计算λ的模逆元
    return {"public_key": (n, g), "private_key": (λ, μ)}

# keys = generate_paillier_keys()
n = 6290536192560178763
g = 6290536192560178764
λ = 6290536187518758120
μ = 2660698920544575751

# n, g = keys["public_key"]
# λ, μ = keys["private_key"]

print(f"公钥: n = {n}, g = {g}")
print(f"私钥: λ = {λ}, μ = {μ}")

# # 加密函数
# def encrypt(plaintext, n, g):
#     n_squared = n * n
#     r = random.randint(1, n - 1)  # 随机数r
#     c = (pow(g, plaintext, n_squared) * pow(r, n, n_squared)) % n_squared
#     return c

# # 加密数据
# a, b, k = 20, 10, 2  # 示例数据：a=5, b=3, 标量k=2
# encrypted_a = encrypt(a, n, g)
# encrypted_b = encrypt(b, n, g)

# print(f"加密结果: E(a) = {encrypted_a}, E(b) = {encrypted_b}")


# 解密函数
def decrypt(ciphertext, λ, μ, n):
    n_squared = n * n
    plaintext = ((pow(ciphertext, λ, n_squared) - 1) // n * μ) % n
    return plaintext

# 解密链上计算结果
encrypted_sum = 1598420809420265871383619227169790083
encrypted_product = 22956890701136816078461401019314214532
decrypted_sum = decrypt(encrypted_sum, λ, μ, n)      # 结果应为 a + b = 8
decrypted_product = decrypt(encrypted_product, λ, μ, n)  # 结果应为 a * k = 10
print(f"解密结果: a + b = {decrypted_sum}")
# print(f"解密结果: a + b = {decrypted_sum}, a * k = {decrypted_product}")