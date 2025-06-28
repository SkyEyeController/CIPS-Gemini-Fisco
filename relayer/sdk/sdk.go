package sdk

import (
	"crossFab/utils"
	"os"
)

// 设置类型
// var chainType = string("pow")
// var chainType = string("poa")

/**
 * @description: 链SDK的结构体
 * @return {*}
 */
type ChainSdk struct {
	url                string
	client             *utils.FabricClient
	eventListenChan    chan map[string][]byte
	eventListenAddress utils.Address
	aggregatorAddress  utils.Address
}

/**
 * @description:初始化一个SDK
 * @param {string} wsURL
 * @param {common.Address} eventListenAddress
 * @param {common.Address} aggregatorAddress
 * @param {*ContractAggregator.ContractAggregator} aggregatorInstance
 * @return {*}
 */
func NewChainSdk(wsURL string, pk string, eventListenAddress utils.Address, aggregatorAddress utils.Address) *ChainSdk {
	//TODO：你需要在这里获取一个Fabric客户端对象，然后构造ChainSDK对象，以下代码供参考：
	//wsURL是链的地址，pk是私钥，后两者是transport和聚合合约的地址

	// chainClient := client.fabricClient{
	// 	Host:       wsURL,
	// 	Version:    "1.0",
	// 	PrivateKey: pk,
	// }

	// if !chainClient.IsConnected() {
	// 	clog.Errorf("Connect to chain error.")
	// 	os.Exit(1)
	// }

	// //构造SDK
	// return &ChainSdk{
	// 	url:                wsURL,
	// 	client:             &chainClient,
	// 	eventListenChan:    nil,
	// 	eventListenAddress: eventListenAddress,
	// 	aggregatorAddress:  aggregatorAddress,
	// }
}

// /**
//  * @description:这里不会被调用到 订阅与监听 一并 合并到 ListenEvent 中
//  * @param {common.Address} contractAddr
//  * @param {string} eventName
//  * @return {*}
//  */
// func (c *ChainSdk) SubscribeEvent(contractAddr fabric.Address, eventName string) string {
// 	clog.Error("这里不应该被调用")
// 	return "这里不应该被调用"
// }

/**
 * @description: 监听事件 从转发层合约 监听
 * @param {chanmap[string][]byte} eventListenChan
 * @return {*}
 */

func (c *ChainSdk) ListenEvent(eventListenChan chan map[string][]byte) {
	//TODO：实现事件监听机制
	client := c.client                    //获取客户端
	returnDict := make(map[string][]byte) //一个返回结果对象
	c.eventListenChan = eventListenChan

	//TODO：在这里监听链上的CmHash消息，从中获取hash和phase字段

	// var re_eventHash [32]byte
	// copy(re_eventHash[:], hs[:])
	// phase := 1
	// re_eventPhase := big.NewInt(int64(phase))

	//TODO：根据实际情况将32字节哈希和phase数值填进来
	returnDict["hash"] = re_eventHash[:]
	returnDict["phase"] = re_eventPhase.Bytes()

	c.eventListenChan <- returnDict //向通道传入你的结果
	//注意：这是一个死循环，你需要写一个循环重复以上过程
}

func file2Bytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size())
	file.Read(data)
	return data, nil
}

/**
 * @description: 接收消息，实际上是已经收到了消息，要发到自己所在的那个链上
 * @param {utils.CrosschainMessage} args_cm
 * @param {[]byte} args_proof
 * @return {*}
 */
func (c *ChainSdk) ReceiveMsg(args_cm utils.CrossChainMessage, args_proof []byte) string {
	data := args_cm
	client := c.client
	//TODO：调用聚合合约的ReceiveMsg，传入data，返回交易哈希
	//注意，你要从c.aggregatorAddress以及c.eventListenAddress中分别获取聚合合约以及transport合约地址，下同

	return txid
}

/**
 * @description:发送ACK确认
 * @param {utils.CrosschainMessage} args_cm
 * @param {[]byte} args_proof
 * @return {*}
 */
func (c *ChainSdk) AcknowledgeMsg(args_cm utils.CrossChainMessage, args_proof []byte) string {
	data := args_cm
	client := c.client
	//TODO：调用聚合合约的AcknowledgeMsg，传入data，返回交易哈希

	return txid
}

func (c *ChainSdk) QueryReqCmByHash(cmhash [32]byte) [][]byte {
	//TODO：调用transport合约的GetReqCmByHash，传入cmhash，返回查询结果

	return result
}

func (c *ChainSdk) QueryRespCmByHash(cmhash [32]byte) [][]byte {
	//TODO：调用transport合约的GetRespCmByHash，传入cmhash，返回查询结果

	return result
}

func (c *ChainSdk) QueryAckCmByHash(cmhash [32]byte) [][]byte {
	//TODO：调用transport合约的GetAckCmByHash，传入cmhash，返回查询结果

	return result
}
