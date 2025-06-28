package relayer

import (
	"crossFab/relayer/config"
	"crossFab/relayer/layer_transport"
	"crossFab/relayer/layer_verification"
	sdk "crossFab/relayer/sdk"
	utils "crossFab/utils"
	"math/big"
	"strconv"
	"time"

	clog "github.com/kpango/glg"
)

func initConfig() *config.Config {
	config := config.NewConfig()

	return config
}

// 一个 Relayer 只记录与自己链相关的信息
type Relayer struct {
	protocol_addr        map[string]utils.Address
	register_addr        map[string]utils.Address
	aggregator_addr      utils.Address
	gateway_url          map[string]string
	relayer_url          string
	chain_url            string
	relayer_user_keypath string

	relayer_verification *layer_verification.Verification
	relayer_transport    *layer_transport.Transport
	// 链SDK
	bcsdk *sdk.ChainSdk
}

func NewRelayer() *Relayer {
	return &Relayer{
		protocol_addr:   make(map[string]utils.Address),
		register_addr:   make(map[string]utils.Address),
		aggregator_addr: utils.Address{},

		gateway_url:          make(map[string]string),
		relayer_url:          "",
		chain_url:            "",
		relayer_user_keypath: "",

		relayer_verification: nil,
		relayer_transport:    nil,
		bcsdk:                nil,
	}
}

func (r *Relayer) RegisterContracts(relayer_config *config.Config) {
	r.aggregator_addr = relayer_config.Aggregator_addr
	r.protocol_addr["transport"] = relayer_config.Transport_addr
}

func (r *Relayer) RegisterGateWays(relayer_config *config.Config) {
	for chain_id, serverURL := range relayer_config.Gateway_urls {
		target_chain_id, _ := new(big.Int).SetString(chain_id, 10)
		r.gateway_url[target_chain_id.String()] = serverURL
	}
	r.relayer_url = relayer_config.Gateway_server_url
}

func (r *Relayer) InitRelayer(relayer_config *config.Config) {
	r.chain_url = relayer_config.Chain_ws_url
	r.relayer_user_keypath = relayer_config.SenderPrivateKey //keypath就是私钥
	r.bcsdk = sdk.NewChainSdk(r.chain_url, r.relayer_user_keypath, r.protocol_addr["transport"], r.aggregator_addr)

	// 实例化 layer_verification 和 layer_transport
	r.relayer_verification = layer_verification.NewVerification(r.bcsdk, relayer_config)
	r.relayer_transport = layer_transport.NewTransport(r.bcsdk, relayer_config)
}

func (r *Relayer) Run(relayer_config *config.Config) {
	go r.relayer_verification.Start()
	time.Sleep(1 * time.Second)

	go r.relayer_transport.Start()
	time.Sleep(1 * time.Second)
}

func StartRelayer(yaml_config *config.YamlConfig) {
	config := initConfig()
	config.Chainid = big.NewInt((int64)(yaml_config.Chain.ChainId))
	config.Aggregator_addr = *utils.StringToAddress(yaml_config.Chain.AggregatorAddr)
	config.Transport_addr = *utils.StringToAddress(yaml_config.Chain.TransportAddr)
	config.Chain_ws_url = yaml_config.Chain.Hostname
	config.Gateway_server_url = yaml_config.Relayer.ServerUrl
	config.SenderPrivateKey = yaml_config.Chain.PrivateKey
	config.Sender = *utils.StringToAddress(yaml_config.Chain.SenderAddress)
	config.EnableRelayChain = yaml_config.Relayer.EnableRelayChain
	for i := 0; i < len(yaml_config.Relayer.Targets); i++ {
		chain_id := yaml_config.Relayer.Targets[i].ChainId
		url := yaml_config.Relayer.Targets[i].Url
		config.Gateway_urls[strconv.FormatUint(chain_id, 10)] = url
		clog.Infof("New Target: %d: %s", chain_id, url)
	}
	for i := 0; i < len(yaml_config.Relayer.Verifications); i++ {
		id := yaml_config.Relayer.Verifications[i].ID
		// _ := yaml_config.Relayer.Verifications[i].Args
		config.Verifications[strconv.FormatUint(id, 10)] = nil //暂时置空
	}
	relayer := NewRelayer()
	relayer.RegisterContracts(config)
	relayer.RegisterGateWays(config)
	relayer.InitRelayer(config)
	clog.Infof("Ready to run relayer...")
	relayer.Run(config)
}
