package utils

import (
	"crypto/sha256"
	"math/big"
	
)

type FabricClient struct {
	//TODO：这里实现你的Fabric的客户端对象结构，用于连接链、部署以及调用合约
}

type CrossChainMessage struct {
	//TODO：在这里实现你的跨链消息包结构，以下是一个参考

	// 公用
	SrcChainId *big.Int
	DstChainId *big.Int
	Seq        *big.Int
	// 应用层
	SrcAppId    *big.Int
	DstAppId    *big.Int
	PayloadReq  [][]byte
	PayloadResp [][]byte
	// 事务层
	TransactionTypeId  *big.Int
	TransactionPayload [][]byte
	// 传输层
	TransmissionTypeId  *big.Int
	TransmissionPayload [][]byte
	// 验证层
	VerificationTypeId  *big.Int
	VerificationPayload [][]byte
	// 转发层
	TransportTypeId  *big.Int
	TransportPayload [][]byte
	// 请求向跨链消息的哈希值, 在传递的过程中不会赋值, 仅在目的链上用于消息索引, 该值不加入哈希计算
	HashReq [32]byte
	// 跨链消息的hash
	HashResp [32]byte
	Ack      bool
}

type Address struct {
	//TODO：在这里实现你需要的合约以及用户地址结构，如果两者不统一，则分成两个类：UserAddress以及ContractAddress
}

func StringToAddress(addr string) *Address {
	//TODO：字符串转地址
	return nil
}

func AddressToBase58(address *Address) string {
	//TODO：地址转字符串'
	return ""
}

// 封装验证层接收和返回的内容
// 为 cm 和一个 []byte 类型的 proof
type CmWithProof struct {
	Cm    *CrossChainMessage
	Proof []byte
}

// toFixedBytes 将*big.Int转换为指定大小的字节数组，采用大端序
func toFixedBytes(n *big.Int, size int) []byte {
	buf := make([]byte, size)
	nBytes := n.Bytes()
	copy(buf[size-len(nBytes):], nBytes)
	return buf
}

// toBoolBytes 将布尔值转换为1字节的大端表示
func toBoolBytes(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

// addListField 向结果中添加列表字段，包括字段的长度和所有元素
func addListField(result *[][]byte, list [][]byte) {
	lengthBytes := toFixedBytes(big.NewInt(int64(len(list))), 32)
	*result = append(*result, lengthBytes)
	*result = append(*result, list...)
}

// 将 list[bytes] 转换为字节数组
func convertLb(data [][]byte, startIndex *big.Int) ([][]byte, *big.Int) {
	var result [][]byte
	payloadLength := bytesToBigInt(data[startIndex.Int64()])
	for i := big.NewInt(1); i.Cmp(payloadLength) <= 0; i = new(big.Int).Add(i, big.NewInt(1)) {
		result = append(result, data[new(big.Int).Add(startIndex, i).Int64()])
	}
	// nextIndex := new(big.Int).Add(startIndex, new(big.Int).Add(payloadLength, big.NewInt(1)))
	nextIndex := new(big.Int).Add(startIndex, new(big.Int).Add(payloadLength, big.NewInt(0)))

	return result, nextIndex
}

// bytesToBigInt 将字节数组转换为 big.Int
func bytesToBigInt(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

// bytesToBool 将字节数组转换为布尔值
func bytesToBool(bytes []byte) bool {
	// 如果字节数组的长度为 0 或第一个字节为 0，则返回 false，否则返回 true
	return len(bytes) > 0 && bytes[0] != 0
}

// cmToLb 将CrosschainMessage转换为字节列表
func CmToLb(ccMsg CrossChainMessage) [][]byte {
	var result [][]byte

	// 按照给定的顺序，依次添加字段到结果中
	result = append(result, toFixedBytes(ccMsg.TransportTypeId, 32))
	addListField(&result, ccMsg.TransportPayload)

	result = append(result, toFixedBytes(ccMsg.VerificationTypeId, 32))
	addListField(&result, ccMsg.VerificationPayload)

	result = append(result, toFixedBytes(ccMsg.TransmissionTypeId, 32))
	addListField(&result, ccMsg.TransmissionPayload)

	result = append(result, toFixedBytes(ccMsg.TransactionTypeId, 32))
	addListField(&result, ccMsg.TransactionPayload)

	result = append(result, toFixedBytes(ccMsg.SrcAppId, 32))
	result = append(result, toFixedBytes(ccMsg.DstAppId, 32))
	addListField(&result, ccMsg.PayloadReq)
	addListField(&result, ccMsg.PayloadResp)

	result = append(result, toFixedBytes(ccMsg.SrcChainId, 32))
	result = append(result, toFixedBytes(ccMsg.DstChainId, 32))
	result = append(result, toFixedBytes(ccMsg.Seq, 32))
	result = append(result, toBoolBytes(ccMsg.Ack))

	return result
}

func CmToHash(ccMsg CrossChainMessage) [32]byte {
	elems := CmToLb(ccMsg)
	hash := sha256.New()
	for _, elem := range elems {
		hash.Write(elem)
	}
	var hashArr [32]byte
	copy(hashArr[:], hash.Sum(nil)) // 将[]byte切片的内容拷贝到[32]byte数组
	return hashArr
}

func CmFromLb(data [][]byte) CrossChainMessage {
	var cm CrossChainMessage
	var endIndex *big.Int

	// 转发层
	cm.TransportTypeId = bytesToBigInt(data[0])
	transportPayload, nextIndex := convertLb(data, big.NewInt(1))
	cm.TransportPayload = transportPayload
	endIndex = nextIndex

	// 验证层
	cm.VerificationTypeId = bytesToBigInt(data[endIndex.Int64()+1])
	verificationPayload, _ := convertLb(data, new(big.Int).Add(endIndex, big.NewInt(2)))
	cm.VerificationPayload = verificationPayload
	endIndex.Add(endIndex, new(big.Int).Add(big.NewInt(2), bytesToBigInt(data[endIndex.Int64()+2])))

	// 传输层
	cm.TransmissionTypeId = bytesToBigInt(data[endIndex.Int64()+1])
	transmissionPayload, _ := convertLb(data, new(big.Int).Add(endIndex, big.NewInt(2)))
	cm.TransmissionPayload = transmissionPayload
	endIndex.Add(endIndex, new(big.Int).Add(big.NewInt(2), bytesToBigInt(data[endIndex.Int64()+2])))

	// 事务层
	cm.TransactionTypeId = bytesToBigInt(data[endIndex.Int64()+1])
	transactionPayload, _ := convertLb(data, new(big.Int).Add(endIndex, big.NewInt(2)))
	cm.TransactionPayload = transactionPayload
	endIndex.Add(endIndex, new(big.Int).Add(big.NewInt(2), bytesToBigInt(data[endIndex.Int64()+2])))

	// 应用层
	cm.SrcAppId = bytesToBigInt(data[endIndex.Int64()+1])
	cm.DstAppId = bytesToBigInt(data[endIndex.Int64()+2])
	payloadReq, _ := convertLb(data, new(big.Int).Add(endIndex, big.NewInt(3)))
	cm.PayloadReq = payloadReq
	endIndex.Add(endIndex, new(big.Int).Add(big.NewInt(3), bytesToBigInt(data[endIndex.Int64()+3])))
	payloadResp, _ := convertLb(data, new(big.Int).Add(endIndex, big.NewInt(1)))
	cm.PayloadResp = payloadResp
	endIndex.Add(endIndex, new(big.Int).Add(big.NewInt(1), bytesToBigInt(data[endIndex.Int64()+1])))

	// 公共
	cm.SrcChainId = bytesToBigInt(data[endIndex.Int64()+1])
	cm.DstChainId = bytesToBigInt(data[endIndex.Int64()+2])
	cm.Seq = bytesToBigInt(data[endIndex.Int64()+3])
	cm.Ack = bytesToBool(data[endIndex.Int64()+4])

	return cm
}

// newAddress 生成一个新的地址，通过对字符串"any address"使用SHA256散列算法
// func NewAddress() *Address {
// 	// 对字符串"any address"进行SHA256散列
// 	hash := sha256.Sum256([]byte("any address"))
// 	// 将散列值转换为common.Address类型并返回
// 	// 注意：这里假设common.Address可以直接从字节切片创建，具体实现可能有所不同
// 	return StringToAddress(utils.GetAddressByBytes(hash[:]))
// }
