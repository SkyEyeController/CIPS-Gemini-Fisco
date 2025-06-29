package sdk

import (
	"context"
	contract_aggregator "crossFab/contracts"
	"crossFab/contracts/layer_transport/eventlisten"
	"crossFab/utils"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	client             *utils.FiscoClient
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

	privateKey, err := hex.DecodeString(pk) // 传入私钥
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
		os.Exit(1)
	}

	// 创建客户端配置对象
	config := &client.Config{
		IsSMCrypto:  false,       // 不使用国密算法
		GroupID:     "group0",    // 群组ID，FISCO BCOS的逻辑分区
		PrivateKey:  privateKey,  // 账户私钥，用于签名交易
		Host:        "127.0.0.1", // 应该是IP地址，不是wsURL
		Port:        20200,       // FISCO BCOS节点RPC端口
		TLSCaFile:   "./ca.crt",  // TLS根证书文件路径
		TLSKeyFile:  "./sdk.key", // TLS客户端私钥文件路径
		TLSCertFile: "./sdk.crt", // TLS客户端证书文件路径

		//这里证书路径问题需要记得提问
	}

	// 🌐 建立与FISCO BCOS节点的连接
	fiscoClient, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to connect:", err) // 连接失败则退出程序
	}
	fmt.Println("Successfully connected to FISCO BCOS") // 连接成功提示

	// 创建包装的FiscoClient
	chainClient := &utils.FiscoClient{
		Client:        fiscoClient,
		Config:        config,
		ChainId:       big.NewInt(1),
		PrivateKeyHex: pk,
		Connected:     true,
	}

	return &ChainSdk{
		url:                wsURL,
		client:             chainClient, // 使用包装的FiscoClient
		eventListenChan:    nil,
		eventListenAddress: eventListenAddress,
		aggregatorAddress:  aggregatorAddress,
	}
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
	client := c.client.Client // 获取底层FISCO客户端
	c.eventListenChan = eventListenChan

	// 构造事件订阅的过滤器
	// 获取当前最新区块号
	latestBlockNumber, err := client.GetBlockNumber(context.Background())
	if err != nil {
		log.Printf("Failed to get latest block number: %v", err)
		latestBlockNumber = 0 // 如果获取失败，从0开始
	}

	eventLogParams := types.EventLogParams{
		FromBlock: latestBlockNumber,                            // 从最新区块开始监听
		ToBlock:   -1,                                           // 到最新区块 (-1 表示 latest)
		Addresses: []string{c.eventListenAddress.Address.Hex()}, // 监听的合约地址
		Topics: []string{
			"0x" + hex.EncodeToString(crypto.Keccak256([]byte("CmHash(bytes32,uint256)"))), // CmHash事件签名
		},
	}

	log.Println("Started listening for CmHash events...")

	// 使用SubscribeEventLogs函数订阅事件
	taskId, err := client.SubscribeEventLogs(context.Background(), eventLogParams, func(status int, logs []types.Log) {
		if status != 0 {
			log.Printf("Event subscription error, status: %d", status)
			return
		}

		// 处理每个事件日志
		for _, eventLog := range logs {
			// 解析CmHash事件
			if len(eventLog.Topics) >= 3 {
				// Topics[0] 是事件签名
				// Topics[1] 是 hash (bytes32, indexed)
				// Topics[2] 是 phase (uint256, indexed)

				var re_eventHash [32]byte
				copy(re_eventHash[:], eventLog.Topics[1].Bytes())

				// 解析phase
				phase := new(big.Int)
				phase.SetBytes(eventLog.Topics[2].Bytes())
				re_eventPhase := phase

				// 构造返回字典
				returnDict := make(map[string][]byte)
				returnDict["hash"] = re_eventHash[:]
				returnDict["phase"] = re_eventPhase.Bytes()

				// 发送到通道
				select {
				case c.eventListenChan <- returnDict:
					log.Printf("CmHash event detected: hash=%x, phase=%d", re_eventHash, phase.Int64())
				default:
					log.Printf("Event channel full, skipping event")
				}
			}
		}
	})

	if err != nil {
		log.Printf("Failed to subscribe to events: %v", err)
		return
	}

	log.Printf("Event subscription started with taskId: %s", taskId)

	// 保持goroutine运行，直到需要停止
	select {}
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

	// 将CrossChainMessage转换为合约需要的字节列表格式
	cmBytes := utils.CmToLb(data)

	// 创建聚合器合约实例
	aggregatorInstance, err := contract_aggregator.NewContractAggregator(
		c.aggregatorAddress.Address, client.Client)
	if err != nil {
		log.Printf("Failed to create aggregator instance: %v", err)
		return ""
	}

	// 创建聚合器会话
	aggregatorSession := &contract_aggregator.ContractAggregatorSession{
		Contract:     aggregatorInstance,
		CallOpts:     *client.Client.GetCallOpts(),
		TransactOpts: *client.Client.GetTransactOpts(),
	}

	// 调用聚合合约的ReceiveMsg，传入cmBytes作为data参数
	_, receipt, err := aggregatorSession.ReceiveMsg(cmBytes)
	if err != nil {
		log.Printf("Failed to call ReceiveMsg: %v", err)
		return ""
	}

	if receipt.Status != 0 {
		log.Printf("ReceiveMsg transaction failed with status: %d", receipt.Status)
		return ""
	}

	txid := receipt.TransactionHash
	log.Printf("ReceiveMsg successful, txid: %s", txid)
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

	// 将CrossChainMessage转换为合约需要的字节列表格式
	cmBytes := utils.CmToLb(data)

	// 创建聚合器合约实例
	aggregatorInstance, err := contract_aggregator.NewContractAggregator(
		c.aggregatorAddress.Address, client.Client)
	if err != nil {
		log.Printf("Failed to create aggregator instance: %v", err)
		return ""
	}

	// 创建聚合器会话
	aggregatorSession := &contract_aggregator.ContractAggregatorSession{
		Contract:     aggregatorInstance,
		CallOpts:     *client.Client.GetCallOpts(),
		TransactOpts: *client.Client.GetTransactOpts(),
	}

	// 调用聚合合约的AcknowledgeMsg，传入data，返回交易哈希
	_, receipt, err := aggregatorSession.AcknowledgeMsg(cmBytes)
	if err != nil {
		log.Printf("Failed to call AcknowledgeMsg: %v", err)
		return ""
	}

	if receipt.Status != 0 {
		log.Printf("AcknowledgeMsg transaction failed with status: %d", receipt.Status)
		return ""
	}

	txid := receipt.TransactionHash
	log.Printf("AcknowledgeMsg successful, txid: %s", txid)
	return txid
}

func (c *ChainSdk) QueryReqCmByHash(cmhash [32]byte) [][]byte {
	// 创建transport合约实例
	transportInstance, err := eventlisten.NewEventlisten(
		c.eventListenAddress.Address, c.client.Client)
	if err != nil {
		log.Printf("Failed to create transport instance: %v", err)
		return nil
	}

	// 创建transport合约会话
	transportSession := &eventlisten.EventlistenSession{
		Contract:     transportInstance,
		CallOpts:     *c.client.Client.GetCallOpts(),
		TransactOpts: *c.client.Client.GetTransactOpts(),
	}

	// 调用transport合约的GetReqCmByHash，传入cmhash，返回查询结果
	result, err := transportSession.GetReqCmByHash(cmhash)
	if err != nil {
		log.Printf("Failed to query req cm by hash: %v", err)
		return nil
	}

	// 将 TypesCrosschainMessage 转换为 [][]byte
	return utils.CmToLb(utils.CrossChainMessage{
		SrcChainId:          result.SrcChainId,
		DstChainId:          result.DstChainId,
		Seq:                 result.Seq,
		SrcAppId:            result.SrcAppId,
		DstAppId:            result.DstAppId,
		PayloadReq:          result.PayloadReq,
		PayloadResp:         result.PayloadResp,
		TransactionTypeId:   result.TransactionTypeId,
		TransactionPayload:  result.TransactionPayload,
		TransmissionTypeId:  result.TransmissionTypeId,
		TransmissionPayload: result.TransmissionPayload,
		VerificationTypeId:  result.VerificationTypeId,
		VerificationPayload: result.VerificationPayload,
		TransportTypeId:     result.TransportTypeId,
		TransportPayload:    result.TransportPayload,
		HashReq:             result.HashReq,
		HashResp:            result.HashResp,
		Ack:                 result.Ack,
	})
}

func (c *ChainSdk) QueryRespCmByHash(cmHash [32]byte) [][]byte {
	// 创建event合约实例
	eventInstance, err := eventlisten.NewEventlisten(
		c.eventListenAddress.Address, c.client.Client)
	if err != nil {
		log.Printf("Failed to create event instance: %v", err)
		return nil
	}

	// 创建event合约会话
	eventSession := &eventlisten.EventlistenSession{
		Contract:     eventInstance,
		CallOpts:     *c.client.Client.GetCallOpts(),
		TransactOpts: *c.client.Client.GetTransactOpts(),
	}

	// 调用event合约的GetRespCmByHash，传入cmHash，返回查询结果
	eventResult, err := eventSession.GetRespCmByHash(cmHash)
	if err != nil {
		log.Printf("Failed to query resp cm by hash: %v", err)
		return nil
	}

	// 将 TypesCrosschainMessage 转换为 [][]byte
	return utils.CmToLb(utils.CrossChainMessage{
		SrcChainId:          eventResult.SrcChainId,
		DstChainId:          eventResult.DstChainId,
		Seq:                 eventResult.Seq,
		SrcAppId:            eventResult.SrcAppId,
		DstAppId:            eventResult.DstAppId,
		PayloadReq:          eventResult.PayloadReq,
		PayloadResp:         eventResult.PayloadResp,
		TransactionTypeId:   eventResult.TransactionTypeId,
		TransactionPayload:  eventResult.TransactionPayload,
		TransmissionTypeId:  eventResult.TransmissionTypeId,
		TransmissionPayload: eventResult.TransmissionPayload,
		VerificationTypeId:  eventResult.VerificationTypeId,
		VerificationPayload: eventResult.VerificationPayload,
		TransportTypeId:     eventResult.TransportTypeId,
		TransportPayload:    eventResult.TransportPayload,
		HashReq:             eventResult.HashReq,
		HashResp:            eventResult.HashResp,
		Ack:                 eventResult.Ack,
	})
}

func (c *ChainSdk) QueryAckCmByHash(cmHash [32]byte) [][]byte {
	// 创建event合约实例
	eventInstance, err := eventlisten.NewEventlisten(
		c.eventListenAddress.Address, c.client.Client)
	if err != nil {
		log.Printf("Failed to create event instance: %v", err)
		return nil
	}

	// 创建event合约会话
	eventSession := &eventlisten.EventlistenSession{
		Contract:     eventInstance,
		CallOpts:     *c.client.Client.GetCallOpts(),
		TransactOpts: *c.client.Client.GetTransactOpts(),
	}

	// 调用event合约的GetAckCmByHash，传入cmHash，返回查询结果
	eventResult, err := eventSession.GetAckCmByHash(cmHash)
	if err != nil {
		log.Printf("Failed to query ack cm by hash: %v", err)
		return nil
	}

	// 将 TypesCrosschainMessage 转换为 [][]byte
	return utils.CmToLb(utils.CrossChainMessage{
		SrcChainId:          eventResult.SrcChainId,
		DstChainId:          eventResult.DstChainId,
		Seq:                 eventResult.Seq,
		SrcAppId:            eventResult.SrcAppId,
		DstAppId:            eventResult.DstAppId,
		PayloadReq:          eventResult.PayloadReq,
		PayloadResp:         eventResult.PayloadResp,
		TransactionTypeId:   eventResult.TransactionTypeId,
		TransactionPayload:  eventResult.TransactionPayload,
		TransmissionTypeId:  eventResult.TransmissionTypeId,
		TransmissionPayload: eventResult.TransmissionPayload,
		VerificationTypeId:  eventResult.VerificationTypeId,
		VerificationPayload: eventResult.VerificationPayload,
		TransportTypeId:     eventResult.TransportTypeId,
		TransportPayload:    eventResult.TransportPayload,
		HashReq:             eventResult.HashReq,
		HashResp:            eventResult.HashResp,
		Ack:                 eventResult.Ack,
	})
}
