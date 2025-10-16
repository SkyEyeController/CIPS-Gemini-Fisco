package config

import (
	"crossFab/utils"
	"math/big"
)

type CmWithProof struct {
	Cm    *utils.CrossChainMessage
	Proof []byte
}

type YamlConfig struct {
	Chain struct {
		ChainId        uint64 `yaml:"chain_id"`
		Hostname       string `yaml:"hostname"`
		SenderAddress  string `yaml:"sender_address"`
		PrivateKey     string `yaml:"private_key"`
		TransportAddr  string `yaml:"transport_addr"`
		AggregatorAddr string `yaml:"aggregator_addr"`
		AppAddr        string `yaml:"app_addr"`
	} `yaml:"chain"`
	Relayer struct {
		ServerUrl        string               `yaml:"server"`
		EnableRelayChain bool                 `yaml:"enable_relaychain"`
		Verifications    []VerifiactionConfig `yaml:"verifications"`
		Targets          []Relayer            `yaml:"targets"`
	} `yaml:"relayer"`
	Test struct {
		DstChainId uint64 `yaml:"dst_chain_id"`
		SrcAppId   uint64 `yaml:"src_app_id"`
		DstAppId   uint64 `yaml:"dst_app_id"`
		AppArgs    string `yaml:"app_args"`
	} `yaml:"test"`
}

type VerifiactionConfig struct {
	ID   uint64 `yaml:"id"`
	Args string `yaml:"args"`
}

type Relayer struct {
	ChainId uint64 `yaml:"chain_id"`
	Url     string `yaml:"url"`
}

type Config struct {
	Chainid              *big.Int
	Sender               utils.Address
	SenderPrivateKey     string
	Chain_ws_url         string
	Q_req_tspt_vrfy      chan *utils.CrossChainMessage
	Q_req_vrfy_tspt      chan *utils.CmWithProof
	Q_req_pre_tspt_vrfy  chan *utils.CmWithProof
	Q_req_pre_vrfy_tspt  chan *utils.CmWithProof
	Q_resp_tspt_vrfy     chan *utils.CrossChainMessage
	Q_resp_vrfy_tspt     chan *utils.CmWithProof
	Q_resp_pre_tspt_vrfy chan *utils.CmWithProof
	Q_resp_pre_vrfy_tspt chan *utils.CmWithProof
	// 验证协议的标识 -> 验证协议的相关参数(如验证协议的地址、事件名称等)
	Verifications       map[string][][]byte
	Transport_addr      utils.Address
	Transport_eventName string
	Aggregator_addr     utils.Address
	// chainid -> url, 其他网关服务的 url
	Gateway_urls map[string]string
	// 当前网关服务的 url
	Gateway_server_url string
	EnableRelayChain   bool

	//unknown usage
	Chain_rpc_urls []string
	Chain_ws_urls  []string
}

// NewConfig initializes and returns a new instance of Config
func NewConfig() *Config {
	return &Config{
		Q_req_tspt_vrfy:      make(chan *utils.CrossChainMessage),
		Q_req_vrfy_tspt:      make(chan *utils.CmWithProof),
		Q_req_pre_tspt_vrfy:  make(chan *utils.CmWithProof),
		Q_req_pre_vrfy_tspt:  make(chan *utils.CmWithProof),
		Q_resp_tspt_vrfy:     make(chan *utils.CrossChainMessage),
		Q_resp_vrfy_tspt:     make(chan *utils.CmWithProof),
		Q_resp_pre_tspt_vrfy: make(chan *utils.CmWithProof),
		Q_resp_pre_vrfy_tspt: make(chan *utils.CmWithProof),

		Verifications:  make(map[string][][]byte),
		Chain_rpc_urls: []string{},
		Chain_ws_urls:  []string{},
		Gateway_urls:   make(map[string]string),
	}
}

// SetQReqTsptVrfy sets the queue for request transport verify
func (c *Config) Set_q_tspt_vrfy(q chan *utils.CrossChainMessage) *Config {
	c.Q_req_tspt_vrfy = q
	return c
}

// SetQVrfyTspt sets the queue for verify transport
func (c *Config) Set_q_vrfy_tspt(q chan *utils.CmWithProof) *Config {
	c.Q_req_vrfy_tspt = q
	return c
}

// SetRPCUrls sets the RPC URLs
func (c *Config) Set_rpc_urls(urls []string) *Config {
	c.Chain_rpc_urls = urls
	return c
}

// SetWSUrls sets the WebSocket URLs
func (c *Config) Set_ws_urls(urls []string) *Config {
	c.Chain_ws_urls = urls
	return c
}

// SetTransportContractAddress sets the transport contract address
func (c *Config) Set_transport_contract_address(addr utils.Address) *Config {
	c.Transport_addr = addr
	return c
}

// SetTransportEventName sets the name of the transport event
func (c *Config) Set_transport_event_name(name string) *Config {
	c.Transport_eventName = name
	return c
}
