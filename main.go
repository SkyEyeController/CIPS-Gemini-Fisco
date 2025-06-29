package main

import (
	"crossFab/relayer"
	"crossFab/relayer/config"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	clog "github.com/kpango/glg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
			//TODO：在这里部署你所有的跨链合约，从yamlConfig中读取配置信息，合约放置到contracts目录下
			//注意，如果你需要一些配置信息，则从yamlConfig中获取，下同
			//同时，你需要激活所有的跨链合约，参考以下海河智链的代码：

			//具体地，调用各层注册（Reg）合约的Set来注册它
			//如果你的合约不需要初始化，或者部署时即可完成初始化，可以忽略下文的Initialize调用
			//在本函数末尾输出两个合约的地址：TransportContract、AggregatorContract

			// sdk.InvokeContractSync(chainClient, AppRegContract, "Initialize", &empty.Empty{}, &empty.Empty{})
			// sdk.InvokeContractSync(chainClient, AppRegContract, "Set", utils.Base58ToAddress(AppContract), &emptypb.Empty{})
			// sdk.InvokeContractSync(chainClient, TransactionRegContract, "Initialize", &empty.Empty{}, &empty.Empty{})
			// sdk.InvokeContractSync(chainClient, TransactionRegContract, "Set", utils.Base58ToAddress(TransactionContract), &emptypb.Empty{})
			// sdk.InvokeContractSync(chainClient, TransmissionRegContract, "Initialize", &empty.Empty{}, &empty.Empty{})
			// sdk.InvokeContractSync(chainClient, TransmissionRegContract, "Set", utils.Base58ToAddress(TransmissionContract), &emptypb.Empty{})
			// sdk.InvokeContractSync(chainClient, VerificationRegContract, "Initialize", &empty.Empty{}, &empty.Empty{})
			// sdk.InvokeContractSync(chainClient, VerificationRegContract, "Set", utils.Base58ToAddress(VerificationContract), &emptypb.Empty{})
			// sdk.InvokeContractSync(chainClient, TransportRegContract, "Initialize", &empty.Empty{}, &empty.Empty{})
			// sdk.InvokeContractSync(chainClient, TransportRegContract, "Set", utils.Base58ToAddress(TransportContract), &emptypb.Empty{})
			// sdk.InvokeContractSync(chainClient, AggregatorContract, "Initialize", &h2chainpb.InternalContractState{
			// 	ChainId:         yamlConfig.Chain.ChainId,
			// 	AppReg:          utils.Base58ToAddress(AppRegContract),
			// 	TransactionReg:  utils.Base58ToAddress(TransactionRegContract),
			// 	TransmissionReg: utils.Base58ToAddress(TransmissionRegContract),
			// 	VerificationReg: utils.Base58ToAddress(VerificationRegContract),
			// 	TransportReg:    utils.Base58ToAddress(TransportRegContract),
			// }, &emptypb.Empty{})
			// clog.Infof("TransportContract Addr = %s", TransportContract)
			// clog.Infof("AggregatorContract Addr = %s", AggregatorContract)

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
			if len(args) == 0 {
				//TODO：调用聚合合约的SendMsg函数，传入DstChainId、SrcAppId、DstAppId和AppArgs
				//聚合合约地址从yamlConfig中获取
				//注意，你的后边的参数，从yamlConfig.Test中读取，以下是一个示例：

				// sdk.InvokeContractSync(chainClient, yamlConfig.Chain.AggregatorAddr, "SendMsg", &h2chainpb.AggregatorSendMsgInput{
				// 	DstChainId: yamlConfig.Test.DstChainId,
				// 	SrcAppId:   yamlConfig.Test.SrcAppId,
				// 	DstAppId:   yamlConfig.Test.DstAppId,
				// 	AppArgs:    []byte(yamlConfig.Test.AppArgs),
				// }, &empty.Empty{})
			} else if len(args) >= 3 {
				dstChainId, _ := strconv.Atoi(args[0])
				srcAppId, _ := strconv.Atoi(args[1])
				dstAppId, _ := strconv.Atoi(args[2])
				appArg := ""
				//TODO：调用聚合合约的SendMsg函数，传入DstChainId、SrcAppId、DstAppId和AppArgs
				//注意，参数从以上变量中获取
				if len(args) == 4 {
					appArg = args[3]
				}
				// sdk.InvokeContractSync(chainClient, yamlConfig.Chain.AggregatorAddr, "SendMsg", &h2chainpb.AggregatorSendMsgInput{
				// 	DstChainId: uint64(dstChainId),
				// 	SrcAppId:   uint64(srcAppId),
				// 	DstAppId:   uint64(dstAppId),
				// 	AppArgs:    []byte(appArg),
				// }, &empty.Empty{})
			}
		},
	}
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(testCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
