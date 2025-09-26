package main

import (
	"crossFab/relayer"
	"crossFab/relayer/config"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	clog "github.com/kpango/glg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"context"
	"encoding/hex"
	"math/big"

	contract_aggregator "crossFab/contracts"
	"crossFab/contracts/layer_app/app"
	"crossFab/contracts/layer_app/appreg"
	"crossFab/contracts/layer_transaction/transactionprotocol"
	"crossFab/contracts/layer_transaction/transactionreg"
	"crossFab/contracts/layer_transmission/transmissionprotocol"
	"crossFab/contracts/layer_transmission/transmissionreg"
	"crossFab/contracts/layer_transport/eventlisten"
	"crossFab/contracts/layer_transport/transportprotocol"
	"crossFab/contracts/layer_transport/transportreg"
	"crossFab/contracts/layer_verification/vericationprotocol"
	"crossFab/contracts/layer_verification/verificationreg"
	"crossFab/contracts/universalkvstore"
	"crossFab/contracts/utils/errorinfo"
	"crossFab/contracts/utils/types"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	fisco_types "github.com/FISCO-BCOS/go-sdk/v3/types"
)

func readYAML() config.YamlConfig {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		clog.Fatalf("Error reading file: %v", err)
		os.Exit(1)
	}
	var yamlConfig config.YamlConfig
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		clog.Fatalf("Error parsing YAML %v", err)
		os.Exit(1)
	}
	return yamlConfig
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "crossFab",
		Short: "Cross-Chain Relayer for Fisco",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Helper: crossFab -h")
		},
	}
	var deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy smart contracts",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			yamlConfig := readYAML()
			//连接配置 - 从yamlConfig中读取配置信息
			privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")

			// 创建客户端配置对象
			config := &client.Config{
				IsSMCrypto:  false,
				GroupID:     "group0",
				PrivateKey:  privateKey,
				Host:        "127.0.0.1",
				Port:        20200,
				TLSCaFile:   "./ca.crt",
				TLSKeyFile:  "./sdk.key",
				TLSCertFile: "./sdk.crt",
			}

			client, err := client.DialContext(context.Background(), config)
			if err != nil {
				clog.Fatalf("Failed to connect to FISCO BCOS: %v", err)
			}
			clog.Info("Successfully connected to FISCO BCOS")

			// =================== 第一步：部署基础合约 ===================
			clog.Info("=================Step 1: Deploy Base Contracts===============")

			// 部署 Types 合约
			typesAddress, typesReceipt, _, err := types.DeployTypes(client.GetTransactOpts(), client)
			if err != nil || typesReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy Types contract: %v", err)
			}
			clog.Infof("✅ Types deployed: %s", typesAddress.Hex())

			// 部署 ErrorInfo 合约
			errorInfoAddress, errorInfoReceipt, _, err := errorinfo.DeployErrorinfo(client.GetTransactOpts(), client)
			if err != nil || errorInfoReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy ErrorInfo contract: %v", err)
			}
			clog.Infof("✅ ErrorInfo deployed: %s", errorInfoAddress.Hex())

			// =================== 第二步：部署注册合约 ===================
			clog.Info("=================Step 2: Deploy Registry Contracts===============")

			//部署 VerificationReg 合约
			verificationRegAddress, verificationRegReceipt, verificationRegInstance, err := verificationreg.DeployVerificationreg(client.GetTransactOpts(), client)
			if err != nil || verificationRegReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy VerificationReg: %v", err)
			}
			clog.Infof("✅ VerificationReg deployed: %s", verificationRegAddress.Hex())

			//部署 TransportReg 合约
			transportRegAddress, transportRegReceipt, transportRegInstance, err := transportreg.DeployTransportreg(client.GetTransactOpts(), client)
			if err != nil || transportRegReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransportReg: %v", err)
			}
			clog.Infof("✅ TransportReg deployed: %s", transportRegAddress.Hex())

			//部署 TransmissionReg 合约
			transmissionRegAddress, transmissionRegReceipt, transmissionRegInstance, err := transmissionreg.DeployTransmissionreg(client.GetTransactOpts(), client)
			if err != nil || transmissionRegReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransmissionReg: %v", err)
			}
			clog.Infof("✅ TransmissionReg deployed: %s", transmissionRegAddress.Hex())

			//部署 TransactionReg 合约
			transactionRegAddress, transactionRegReceipt, transactionRegInstance, err := transactionreg.DeployTransactionreg(client.GetTransactOpts(), client)
			if err != nil || transactionRegReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransactionReg: %v", err)
			}
			clog.Infof("✅ TransactionReg deployed: %s", transactionRegAddress.Hex())

			// 部署 AppReg 合约
			appRegAddress, appRegReceipt, appRegInstance, err := appreg.DeployAppreg(client.GetTransactOpts(), client)
			if err != nil || appRegReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy AppReg: %v", err)
			}
			clog.Infof("✅ AppReg deployed: %s", appRegAddress.Hex())

			// =================== 第三步：部署协议实现合约 ===================
			clog.Info("=================Step 3: Deploy Protocol Implementations===============")

			// 部署验证协议实现
			verificationImplAddress, verificationImplReceipt, _, err := vericationprotocol.DeployVericationprotocol(client.GetTransactOpts(), client)
			if err != nil || verificationImplReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy VerificationProtocol: %v", err)
			}
			clog.Infof("✅ VerificationProtocol deployed: %s", verificationImplAddress.Hex())

			// 部署转发协议实现
			transportImplAddress, transportImplReceipt, _, err := transportprotocol.DeployTransportprotocol(client.GetTransactOpts(), client)
			if err != nil || transportImplReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransportProtocol: %v", err)
			}
			clog.Infof("✅ TransportProtocol deployed: %s", transportImplAddress.Hex())

			// 部署传输协议实现
			transmissionImplAddress, transmissionImplReceipt, _, err := transmissionprotocol.DeployTransmissionprotocol(client.GetTransactOpts(), client)
			if err != nil || transmissionImplReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransmissionProtocol: %v", err)
			}
			clog.Infof("✅ TransmissionProtocol deployed: %s", transmissionImplAddress.Hex())

			//部署事务协议实现
			transactionImplAddress, transactionImplReceipt, _, err := transactionprotocol.DeployTransactionprotocol(client.GetTransactOpts(), client)
			if err != nil || transactionImplReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy TransactionProtocol: %v", err)
			}
			clog.Infof("✅ TransactionProtocol deployed: %s", transactionImplAddress.Hex())

			// 部署应用协议实现
			appImplAddress, appImplReceipt, _, err := app.DeployApp(client.GetTransactOpts(), client)
			if err != nil || appImplReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy App: %v", err)
			}
			clog.Infof("✅ App deployed: %s", appImplAddress.Hex())

			// 部署事件监听协议（这是我们的TransportContract）
			eventListenAddress, eventListenReceipt, _, err := eventlisten.DeployEventlisten(client.GetTransactOpts(), client)
			if err != nil || eventListenReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy EventListen: %v", err)
			}
			clog.Infof("✅ EventListen deployed: %s", eventListenAddress.Hex())

			// =================== 第四步：注册所有协议到系统 ===================
			clog.Info("=================Step 4: Register All Protocols===============")

			// 具体地，调用各层注册（Reg）合约的Set来注册它

			// 注册验证协议
			verificationRegSession := &verificationreg.VerificationregSession{
				Contract:     verificationRegInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, verRegReceipt, err := verificationRegSession.Set(verificationImplAddress)
			if err != nil || verRegReceipt.Status != 0 {
				clog.Fatalf("Failed to register verification protocol: %v", err)
			}
			clog.Infof("✅ Verification protocol registered")

			//注册转发协议 - 使用EventListen作为转发层协议
			transportRegSession := &transportreg.TransportregSession{
				Contract:     transportRegInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, transpRegReceipt, err := transportRegSession.Set(eventListenAddress) // 注册EventListen而不是TransportProtocol
			if err != nil || transpRegReceipt.Status != 0 {
				clog.Fatalf("Failed to register transport protocol: %v", err)
			}
			clog.Infof("✅ Transport protocol (EventListen) registered")

			//注册传输协议
			transmissionRegSession := &transmissionreg.TransmissionregSession{
				Contract:     transmissionRegInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, transmissionRegReceipt, err = transmissionRegSession.Set(transmissionImplAddress)
			if err != nil || transmissionRegReceipt.Status != 0 {
				clog.Fatalf("Failed to register transmission protocol: %v", err)
			}
			clog.Infof("✅ Transmission protocol registered")

			//注册事务协议
			transactionRegSession := &transactionreg.TransactionregSession{
				Contract:     transactionRegInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, transactionRegReceipt, err = transactionRegSession.Set(transactionImplAddress)
			if err != nil || transactionRegReceipt.Status != 0 {
				clog.Fatalf("Failed to register transaction protocol: %v", err)
			}
			clog.Infof("✅ Transaction protocol registered")

			//注册应用协议
			appRegSession := &appreg.AppregSession{
				Contract:     appRegInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, appRegReceipt, err = appRegSession.Set(appImplAddress)
			if err != nil || appRegReceipt.Status != 0 {
				clog.Fatalf("Failed to register app protocol: %v", err)
			}
			clog.Infof("✅ App protocol registered")

			// =================== 第五步：部署聚合器合约 ===================
			clog.Info("=================Step 5: Deploy Cross-Chain Aggregator===============")

			//部署聚合器合约 - 跨链系统的核心协调器
			chainId := big.NewInt(int64(yamlConfig.Chain.ChainId))
			aggregatorAddress, aggregatorReceipt, _, err := contract_aggregator.DeployContractAggregator(
				client.GetTransactOpts(),
				client,
				chainId,
				appRegAddress,
				transactionRegAddress,
				transmissionRegAddress,
				verificationRegAddress,
				transportRegAddress,
			)
			if err != nil || aggregatorReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy Contract_Aggregator: %v", err)
			}
			clog.Infof("🎯 Cross-Chain Aggregator deployed: %s", aggregatorAddress.Hex())

			// 配置EventListen合约，设置聚合器地址
			eventListenInstance, err := eventlisten.NewEventlisten(eventListenAddress, client)
			if err != nil {
				clog.Fatalf("Failed to create EventListen instance: %v", err)
			}
			eventListenSession := &eventlisten.EventlistenSession{
				Contract:     eventListenInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}
			_, setAggregatorReceipt, err := eventListenSession.SetAggregator(aggregatorAddress)
			if err != nil || setAggregatorReceipt.Status != 0 {
				clog.Fatalf("Failed to set aggregator address in EventListen: %v", err)
			}
			clog.Infof("✅ EventListen aggregator address configured")

			// 验证配置 - 读取EventListen中的聚合器地址
			currentAggregator, err := eventListenSession.GetAggregator()
			if err != nil {
				clog.Warnf("Failed to read aggregator address from EventListen: %v", err)
			} else {
				clog.Infof("🔍 EventListen aggregator address: %s", currentAggregator.Hex())
				if currentAggregator.Hex() == aggregatorAddress.Hex() {
					clog.Infof("✅ Aggregator address verification passed")
				} else {
					clog.Errorf("❌ Aggregator address mismatch!")
				}
			}
			// =================== 第六步：部署通用键值存储合约 ===================
			clog.Info("=================Step 6: Deploy UniversalKVStore===============")
			//部署通用键值存储合约
			kvStoreAddress, kvStoreReceipt, _, err := universalkvstore.DeployUniversalkvstore(client.GetTransactOpts(), client)
			if err != nil || kvStoreReceipt.Status != 0 {
				clog.Fatalf("Failed to deploy UniversalKVStore: %v", err)
			}
			clog.Infof("✅ UniversalKVStore deployed: %s", kvStoreAddress.Hex())
			// =================== 第七步：部署完成，输出合约地址 ===================

			// 在本函数末尾输出两个合约的地址：TransportContract、AggregatorContract
			clog.Infof("TransportContract Addr = %s", eventListenAddress.Hex())
			clog.Infof("AggregatorContract Addr = %s", aggregatorAddress.Hex())
			clog.Infof("UniversalKVStore Addr = %s", kvStoreAddress.Hex())
			// 将合约地址写入配置文件
			yamlConfig.Chain.TransportAddr = eventListenAddress.Hex()
			yamlConfig.Chain.AggregatorAddr = aggregatorAddress.Hex()

			// 部署完成总结
			clog.Info("All cross-chain contracts deployed and configured successfully!")
		},
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the relayer",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			yamlConfig := readYAML()
			go relayer.StartRelayer(&yamlConfig)
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
			<-stop
		},
	}
	testCmd := &cobra.Command{
		Use:   "test <dstChainId> <srcAppId> <dstAppId> <appArg>",
		Short: "Send a cross-chain message to chain",
		Args:  cobra.MaximumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			yamlConfig := readYAML()
			if yamlConfig.Chain.AggregatorAddr == "" {
				clog.Logf("You should give the address of aggregator contract.")
				os.Exit(-1)
			}

			// 从配置初始化FISCO BCOS客户端
			privateKey, err := hex.DecodeString( /*yamlConfig.Chain.PrivateKey*/ "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
			if err != nil {
				clog.Fatalf("Failed to decode private key: %v", err)
			}

			// 创建客户端配置对象 - 从yamlConfig读取配置
			config := &client.Config{
				IsSMCrypto:  false,       // 不使用国密算法（可从yamlConfig读取）
				GroupID:     "group0",    // 群组ID（可从yamlConfig读取）
				PrivateKey:  privateKey,  // 解码后的私钥
				Host:        "127.0.0.1", // FISCO BCOS节点IP地址（可从yamlConfig读取）
				Port:        20200,       // FISCO BCOS节点RPC端口（可从yamlConfig读取）
				TLSCaFile:   "./ca.crt",  // TLS根证书文件路径（可从yamlConfig读取）
				TLSKeyFile:  "./sdk.key", // TLS客户端私钥文件路径（可从yamlConfig读取）
				TLSCertFile: "./sdk.crt", // TLS客户端证书文件路径（可从yamlConfig读取）
			}

			// 建立与FISCO BCOS节点的连接
			client, err := client.DialContext(context.Background(), config)
			if err != nil {
				clog.Fatalf("Failed to connect to FISCO BCOS: %v", err)
			}

			// 创建聚合器合约实例
			aggregatorInstance, err := contract_aggregator.NewContractAggregator(
				common.HexToAddress(yamlConfig.Chain.AggregatorAddr), client)
			if err != nil {
				clog.Fatalf("Failed to create aggregator instance: %v", err)
			}

			// 创建聚合器会话
			aggregatorSession := &contract_aggregator.ContractAggregatorSession{
				Contract:     aggregatorInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}

			if len(args) == 0 {
				// TODO：调用聚合合约的SendMsg函数，传入DstChainId、SrcAppId、DstAppId和AppArgs
				// 聚合合约地址从yamlConfig中获取
				// 注意，你的后边的参数，从yamlConfig.Test中读取
				clog.Info("Sending cross-chain message with test configuration...")

				// 从yamlConfig.Test中读取测试参数
				dstChainId := big.NewInt(int64(yamlConfig.Test.DstChainId))
				srcAppId := big.NewInt(int64(yamlConfig.Test.SrcAppId))
				dstAppId := big.NewInt(int64(yamlConfig.Test.DstAppId))
				appArgs := [][]byte{[]byte(yamlConfig.Test.AppArgs)}

				clog.Infof("DEBUG: dstChainId=%v, srcAppId=%v, dstAppId=%v, appArgs=%v\n", dstChainId, srcAppId, dstAppId, appArgs)

				// 调用聚合合约的SendMsg函数
				_, receipt, err := aggregatorSession.SendMsg(dstChainId, srcAppId, dstAppId, appArgs)

				if err != nil {
					clog.Errorf("❌ SendMsg failed: %v", err)
					os.Exit(-1)
				} else if receipt.Status == 0 {
					// 交易成功
					clog.Infof("✅ Cross-chain message sent successfully!")
					clog.Infof("   Destination Chain: %d", yamlConfig.Test.DstChainId)
					clog.Infof("   Source App: %d", yamlConfig.Test.SrcAppId)
					clog.Infof("   Destination App: %d", yamlConfig.Test.DstAppId)
					clog.Infof("   App Args: %s", yamlConfig.Test.AppArgs)
					clog.Infof("   Transaction Hash: %s", receipt.TransactionHash)
					clog.Infof("   Gas Used: %s", receipt.GasUsed)

					// 详细输出事件日志信息
					if len(receipt.Logs) > 0 {
						clog.Infof("📋 Event Logs Found: %d", len(receipt.Logs))
						for i, log := range receipt.Logs {
							clog.Infof("   Log %d:", i)
							clog.Infof("     Contract: %s", log.Address)
							clog.Infof("     Topics: %d", len(log.Topics))
							for j, topic := range log.Topics {
								clog.Infof("       Topic[%d]: %s", j, topic)
							}
							clog.Infof("     Data: %s", log.Data)
						}
					} else {
						clog.Warnf("⚠️  No event logs found in transaction!")
					}
				} else {
					// 交易失败
					clog.Errorf("❌ SendMsg transaction failed with status: %d", receipt.Status)
					os.Exit(-1)
				}
			} else if len(args) >= 3 {
				dstChainId, _ := strconv.Atoi(args[0])
				srcAppId, _ := strconv.Atoi(args[1])
				dstAppId, _ := strconv.Atoi(args[2])
				appArg := ""

				// TODO：调用聚合合约的SendMsg函数，传入DstChainId、SrcAppId、DstAppId和AppArgs
				// 注意，参数从以上变量中获取
				if len(args) == 4 {
					appArg = args[3]
				}

				clog.Info("Sending cross-chain message with command line arguments...")
				clog.Infof("   Destination Chain: %d", dstChainId)
				clog.Infof("   Source App: %d", srcAppId)
				clog.Infof("   Destination App: %d", dstAppId)
				clog.Infof("   App Args: %s", appArg)

				// 转换参数为合约需要的类型
				bigDstChainId := big.NewInt(int64(dstChainId))
				bigSrcAppId := big.NewInt(int64(srcAppId))
				bigDstAppId := big.NewInt(int64(dstAppId))
				appArgs := [][]byte{[]byte(appArg)}

				// 调用聚合合约的SendMsg函数
				_, receipt, err := aggregatorSession.SendMsg(bigDstChainId, bigSrcAppId, bigDstAppId, appArgs)

				if err != nil {
					clog.Errorf("❌ SendMsg failed: %v", err)
					os.Exit(-1)
				} else if receipt.Status == 0 {
					// 交易成功
					clog.Infof("✅ Cross-chain message sent successfully!")
					clog.Infof("   Transaction Hash: %s", receipt.TransactionHash)
					clog.Infof("   Gas Used: %s", receipt.GasUsed)

					// 输出事件日志信息
					if len(receipt.Logs) > 0 {
						clog.Infof("   Event Logs: %d", len(receipt.Logs))
						for i, log := range receipt.Logs {
							clog.Infof("   Log %d: Contract=%s, Topics=%d", i, log.Address, len(log.Topics))
						}
					}
				} else {
					// 交易失败
					clog.Errorf("❌ SendMsg transaction failed with status: %d", receipt.Status)
					os.Exit(-1)
				}
			} else {
				// 参数不足
				clog.Errorf("❌ Invalid arguments. Usage:")
				clog.Errorf("   crossFab test                              # Use config.yml test parameters")
				clog.Errorf("   crossFab test <dstChainId> <srcAppId> <dstAppId> [appArg]  # Use command line parameters")
				os.Exit(-1)
			}
		},
	}
	debugCmd := &cobra.Command{
		Use:   "debug",
		Short: "Debug EventListen contract and event subscription",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			yamlConfig := readYAML()
			if yamlConfig.Chain.TransportAddr == "" {
				clog.Logf("You should give the address of transport contract in config.yml.")
				clog.Logf("Please run 'deploy' command first to get the transport contract address.")
				os.Exit(-1)
			}

			// 初始化FISCO BCOS客户端
			privateKey, err := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
			if err != nil {
				clog.Fatalf("Failed to decode private key: %v", err)
			}

			config := &client.Config{
				IsSMCrypto:  false,
				GroupID:     "group0",
				PrivateKey:  privateKey,
				Host:        "127.0.0.1",
				Port:        20200,
				TLSCaFile:   "./ca.crt",
				TLSKeyFile:  "./sdk.key",
				TLSCertFile: "./sdk.crt",
			}

			client, err := client.DialContext(context.Background(), config)
			if err != nil {
				clog.Fatalf("Failed to connect to FISCO BCOS: %v", err)
			}

			// 创建EventListen合约实例
			eventListenInstance, err := eventlisten.NewEventlisten(
				common.HexToAddress(yamlConfig.Chain.TransportAddr), client)
			if err != nil {
				clog.Fatalf("Failed to create EventListen instance: %v", err)
			}

			eventListenSession := &eventlisten.EventlistenSession{
				Contract:     eventListenInstance,
				CallOpts:     *client.GetCallOpts(),
				TransactOpts: *client.GetTransactOpts(),
			}

			// 1. 检查聚合器地址
			clog.Info("🔍 Step 1: Checking EventListen configuration...")
			aggregatorAddr, err := eventListenSession.GetAggregator()
			if err != nil {
				clog.Errorf("Failed to get aggregator address: %v", err)
			} else {
				clog.Infof("   Aggregator Address: %s", aggregatorAddr.Hex())
				clog.Infof("   Expected Address: %s", yamlConfig.Chain.AggregatorAddr)
				if aggregatorAddr.Hex() == yamlConfig.Chain.AggregatorAddr {
					clog.Infof("✅ Aggregator address matches")
				} else {
					clog.Errorf("❌ Aggregator address mismatch!")
				}
			}

			// 2. 测试事件订阅
			clog.Info("🔍 Step 2: Testing event subscription...")

			// 获取当前区块号
			latestBlockNumber, err := client.GetBlockNumber(context.Background())
			if err != nil {
				clog.Errorf("Failed to get latest block number: %v", err)
				latestBlockNumber = 0
			}
			clog.Infof("   Current block number: %d", latestBlockNumber)

			// 计算CmHash事件签名
			eventSignature := "CmHash(bytes32,uint256)"
			expectedTopic := common.BytesToHash(crypto.Keccak256([]byte(eventSignature))).Hex()
			clog.Infof("   Expected CmHash event topic: %s", expectedTopic)

			// 设置事件监听
			eventListenChan := make(chan map[string][]byte, 10)
			subscriptionDone := make(chan bool)

			go func() {
				clog.Info("🔍 Step 3: Starting event subscription test...")
				var eventLogParams fisco_types.EventLogParams
				eventLogParams.FromBlock = 1 // 从区块1开始监听
				eventLogParams.ToBlock = -1  // 监听到最新区块
				eventLogParams.Addresses = []string{strings.ToLower(yamlConfig.Chain.TransportAddr)}
				// 监听Transport合约地址
				eventLogParams.Topics = []string{} // 监听所有事件

				clog.Infof("   Subscription params: FromBlock=%d, ToBlock=-1", 1)
				clog.Infof("   Monitoring contract: %s", yamlConfig.Chain.TransportAddr)

				taskId, err := client.SubscribeEventLogs(context.Background(), eventLogParams,
					func(status int, logs []fisco_types.Log) {
						clog.Infof("📥 Event subscription callback: status=%d, logs=%d", status, len(logs))

						if status != 0 {
							clog.Errorf("Event subscription error, status: %d", status)
							return
						}

						for i, eventLog := range logs {
							clog.Infof("   Event %d:", i)
							clog.Infof("     Contract: %s", eventLog.Address)
							clog.Infof("     Topics: %d", len(eventLog.Topics))

							for j, topic := range eventLog.Topics {
								clog.Infof("       Topic[%d]: %s", j, topic.Hex())
							}
							clog.Infof("     Data: %s", eventLog.Data)

							// 检查是否是CmHash事件或其他事件
							if len(eventLog.Topics) >= 1 {
								topicHex := eventLog.Topics[0].Hex()
								if topicHex == expectedTopic {
									clog.Infof("✅ Found CmHash event!")
									eventListenChan <- map[string][]byte{"type": []byte("CmHash"), "found": []byte("true")}
								} else if topicHex == "0x1ace2b42299d2f9f1ffdeefaf822c85d2b263f6105d7cf4a3e482f14032fb52e" {
									clog.Infof("✅ Found Test_sendOut event!")
									eventListenChan <- map[string][]byte{"type": []byte("Test_sendOut"), "found": []byte("true")}
								} else {
									clog.Infof("ℹ️ Found other event: %s", topicHex)
									eventListenChan <- map[string][]byte{"type": []byte("other"), "topic": []byte(topicHex)}
								}
							}
						}
					})

				if err != nil {
					clog.Errorf("Failed to subscribe to events: %v", err)
					subscriptionDone <- false
					return
				}

				clog.Infof("✅ Event subscription started with taskId: %s", taskId)
				subscriptionDone <- true

				// 保持订阅活跃
				select {}
			}()

			// 等待订阅启动
			if success := <-subscriptionDone; !success {
				clog.Errorf("❌ Event subscription failed")
				return
			}

			// 3. 手动触发事件 - 测试 emit_sendOut (Test_sendOut 事件)
			clog.Info("🔍 Step 4a: Testing emit_sendOut (Test_sendOut event)...")
			_, emitReceipt, err := eventListenSession.EmitSendOut()
			if err != nil {
				clog.Errorf("Failed to call emit_sendOut: %v", err)
			} else {
				clog.Infof("✅ emit_sendOut called successfully")
				clog.Infof("   Transaction Hash: %s", emitReceipt.TransactionHash)
				clog.Infof("   Gas Used: %s", emitReceipt.GasUsed)

				if len(emitReceipt.Logs) > 0 {
					clog.Infof("📋 Direct Event Logs: %d", len(emitReceipt.Logs))
					for i, log := range emitReceipt.Logs {
						clog.Infof("   Log %d: Contract=%s, Topics=%d", i, log.Address, len(log.Topics))
						for j, topic := range log.Topics {
							clog.Infof("     Topic[%d]: %s", j, topic)
						}
					}
				} else {
					clog.Warnf("⚠️  No direct event logs found!")
				}
			}

			// 4. 通过聚合器触发正常的业务流程 - 测试真正的 CmHash 事件
			clog.Info("🔍 Step 4b: Testing sendMsg via Aggregator (CmHash event)...")
			if yamlConfig.Chain.AggregatorAddr == "" {
				clog.Errorf("❌ Aggregator address not configured in config.yml")
			} else {
				// 创建聚合器合约实例
				aggregatorInstance, err := contract_aggregator.NewContractAggregator(
					common.HexToAddress(yamlConfig.Chain.AggregatorAddr), client)
				if err != nil {
					clog.Errorf("Failed to create aggregator instance: %v", err)
				} else {
					// 创建聚合器会话
					aggregatorSession := &contract_aggregator.ContractAggregatorSession{
						Contract:     aggregatorInstance,
						CallOpts:     *client.GetCallOpts(),
						TransactOpts: *client.GetTransactOpts(),
					}

					// 使用测试参数发送跨链消息
					dstChainId := big.NewInt(int64(yamlConfig.Test.DstChainId))
					srcAppId := big.NewInt(int64(yamlConfig.Test.SrcAppId))
					dstAppId := big.NewInt(int64(yamlConfig.Test.DstAppId))
					appArgs := [][]byte{[]byte(yamlConfig.Test.AppArgs)}

					clog.Infof("   Calling aggregator.sendMsg with:")
					clog.Infof("     dstChainId: %d", yamlConfig.Test.DstChainId)
					clog.Infof("     srcAppId: %d", yamlConfig.Test.SrcAppId)
					clog.Infof("     dstAppId: %d", yamlConfig.Test.DstAppId)
					clog.Infof("     appArgs: %s", yamlConfig.Test.AppArgs)

					// 调用聚合器的SendMsg函数
					_, sendReceipt, err := aggregatorSession.SendMsg(dstChainId, srcAppId, dstAppId, appArgs)

					if err != nil {
						clog.Errorf("❌ SendMsg failed: %v", err)
					} else if sendReceipt.Status == 0 {
						clog.Infof("✅ Aggregator SendMsg called successfully")
						clog.Infof("   Transaction Hash: %s", sendReceipt.TransactionHash)
						clog.Infof("   Gas Used: %s", sendReceipt.GasUsed)

						if len(sendReceipt.Logs) > 0 {
							clog.Infof("📋 Aggregator Event Logs: %d", len(sendReceipt.Logs))
							for i, log := range sendReceipt.Logs {
								clog.Infof("   Log %d: Contract=%s, Topics=%d", i, log.Address, len(log.Topics))
								for j, topic := range log.Topics {
									clog.Infof("     Topic[%d]: %s", j, topic)
									// 检查是否是CmHash事件
									if topic == expectedTopic {
										clog.Infof("🎉 Found CmHash event in aggregator transaction!")
									}
								}
							}
						} else {
							clog.Warnf("⚠️  No event logs found in aggregator transaction!")
						}
					} else {
						clog.Errorf("❌ SendMsg transaction failed with status: %d", sendReceipt.Status)
					}
				}
			}

			// 5. 等待事件回调
			clog.Info("🔍 Step 5: Waiting for event callback (15 seconds)...")
			eventCount := 0
			timeout := time.After(15 * time.Second)

			for {
				select {
				case event := <-eventListenChan:
					eventCount++
					clog.Infof("🎉 Event %d detected via subscription! %v", eventCount, event)
					if eventCount >= 2 { // 期望收到两个事件：Test_sendOut 和 CmHash
						clog.Infof("✅ All expected events received!")
						return
					}
				case <-timeout:
					if eventCount > 0 {
						clog.Infof("⏰ TIMEOUT: Received %d events via subscription after 15 seconds", eventCount)
					} else {
						clog.Warnf("⏰ TIMEOUT: No events detected via subscription after 15 seconds")
					}
					return
				}
			}
		},
	}

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(debugCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
