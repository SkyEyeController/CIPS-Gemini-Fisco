package layer_transport

import (
	"bytes"
	relay_http "crossFab/internal"
	"crossFab/relayer/config"
	"crossFab/relayer/sdk"
	"crossFab/utils"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"encoding/binary"
	"encoding/json"

	clog "github.com/kpango/glg"
)

/******************************** Gateway Data ********************************/

// 用于网关间数据传输
// 设计平台、语言无关的编解码方法
// 暂时采用 json 来实现
type GatewayData struct {
	Message [][]byte `json:"message"`
	Proof   []byte   `json:"proof"`
}

func NewGatewayData(message [][]byte, proof []byte) *GatewayData {
	return &GatewayData{
		Message: message,
		Proof:   proof,
	}
}

// 编码 GatewayData 为 json 格式
func (gd *GatewayData) Encode() []byte {
	// message_bytes := make([][]byte, len(gd.Message))
	// for i, m := range gd.Message {
	// 	message_bytes[i] = []byte(hex.EncodeToString(m))
	// }
	// data := NewGatewayData(message_bytes, []byte(hex.EncodeToString(gd.Proof)))
	// res, err := json.Marshal(data)

	res, err := json.Marshal(gd)
	if err != nil {
		clog.Fatalf("Encode GatewayData fail: %v", err)
	}

	return res
}

// 解码 json 格式的数据为 GatewayData
func (gd *GatewayData) Decode(data []byte) {
	// receive := NewGatewayData(nil, nil)
	// err := json.Unmarshal(data, &receive)
	// if err != nil {
	// 	clog.Fatalf("Decode GatewayData fail: %v", err)
	// }

	// gd.Message = make([][]byte, len(receive.Message))
	// for i, m := range receive.Message {
	// 	gd.Message[i], err = hex.DecodeString(string(m))
	// 	if err != nil {
	// 		clog.Fatalf("Decode GatewayData Message fail: %v", err)
	// 	}
	// }

	// gd.Proof, err = hex.DecodeString(string(receive.Proof))
	// if err != nil {
	// 	clog.Fatalf("Decode GatewayData Proof fail: %v", err)
	// }

	err := json.Unmarshal(data, &gd)
	if err != nil {
		clog.Fatalf("Decode GatewayData fail: %v", err)
	}
}

func (gd *GatewayData) GetMessage() [][]byte {
	return gd.Message
}
func (gd *GatewayData) GetProof() []byte {
	return gd.Proof
}

/******************************** Gateway Api Client ********************************/

// 用于向另一个网关发送消息
type GatewayApiClient struct {
	relay_client *relay_http.HttpClient
}

func NewGatewayApiClient(server_url string) *GatewayApiClient {
	return &GatewayApiClient{
		relay_client: relay_http.NewHttpClient(server_url),
	}
}

// 向目的链网关发送原链的请求消息
// 包含跨链消息和证明
func (gc *GatewayApiClient) SendReqCm(cm *utils.CrossChainMessage, proof []byte) ([]byte, string) {
	clog.Logf("cm.PayloadReq: %v", strings.Trim(string(cm.PayloadReq[0]), "\x00"))
	gwData := NewGatewayData(utils.CmToLb(*cm), proof)
	gwDataEncoded := gwData.Encode()
	clog.Logf("gwDataEncoded: %v", strings.Trim(string(gwDataEncoded), "\x00"))

	path := "/cm/req"
	response, err := http.Post(gc.relay_client.GetURL()+path, "application/json", bytes.NewBuffer(gwDataEncoded))
	if err != nil {
		clog.Fatalf("use POST to post data error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Sprintln("requests path:", path, "; requests error.")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		clog.Fatalf("read POST response error: %v", err)
	}

	// clog.Debugf("SendReqCm: send reqCm to gateway, path: %s/cm/req", gc.relay_client.GetURL())
	return body, ""
}

// 向原链网关发送目的链的响应消息
// 包含跨链消息和证明
func (gc *GatewayApiClient) SendRespCm(cm *utils.CrossChainMessage, proof []byte) ([]byte, string) {
	gwData := NewGatewayData(utils.CmToLb(*cm), proof)
	gwDataEncoded := gwData.Encode()

	path := "/cm/resp"
	response, err := http.Post(gc.relay_client.GetURL()+path, "application/json", bytes.NewBuffer(gwDataEncoded))
	if err != nil {
		clog.Fatalf("use POST to post data error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Sprintln("requests path:", path, "; requests error.")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		clog.Fatalf("read POST response error: %v", err)
	}

	// clog.Debugf("SendRespCm: send respCm to gateway, path: %s/cm/resp", gc.relay_client.GetURL())
	return body, ""
}

func (gc *GatewayApiClient) Register_other_gateway(chain_id *big.Int, gateway_url string) ([]byte, string) {
	data := struct {
		Chain_id    *big.Int `json:"chain_id"`
		Gateway_url string   `json:"gateway_url"`
	}{
		Chain_id:    chain_id,
		Gateway_url: gateway_url,
	}
	dataEncoded, err := json.Marshal(data)
	if err != nil {
		clog.Fatalf("Encode data error: %v", err)
	}

	path := "/register/gateway"
	response := gc.relay_client.POST(path, dataEncoded)
	if response.StatusCode != 200 {
		return nil, fmt.Sprintln("requests path:", path, "; requests error.")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		clog.Fatalf("read POST response error: %v", err)
	}
	return body, ""
}

/******************************** Gateway Api Server ********************************/

// gateway api server 基于 api server 实现
// 和 transport 的内核分开，不然容易出现循环导入的问题
type GatewayApiServer struct {
	url          string
	relay_server *relay_http.HttpEngine
	transport    *Transport
}

// 下面为按照 CIPS-Standards 直接实现的内容
// 需要修改以满足 http 的接口需求
/*
func (gs *GatewayApiServer) Cmreq(data []byte) bool {
	gwData := NewGatewayData(nil, nil)
	gwData.Decode(data)
	cm := utils.CmFromLb(gwData.GetMessage())
	proof := gwData.GetProof()
	return gs.transport.ReceivceReqCmFromOthers(cm, proof)
}

func (gs *GatewayApiServer) Cmresp(data []byte) bool {
	gwData := NewGatewayData(nil, nil)
	gwData.Decode(data)
	cm := utils.CmFromLb(gwData.GetMessage())
	proof := gwData.GetProof()
	return gs.transport.ReceivceRespCmFromOthers(cm, proof)
}

func (gs *GatewayApiServer) RegisterGateway(data []byte) bool {
	var receive struct {
		Chain_id    *big.Int `json:"chain_id"`
		Gateway_url string   `json:"gateway_url"`
	}
	err := json.Unmarshal(data, &receive)
	if err != nil {
		clog.Fatalf("Decode data error: %v", err)
	}
	chain_id := receive.Chain_id
	gateway_url := receive.Gateway_url
	return gs.transport.RegisterGateway(chain_id, gateway_url)
}
*/

/* 以下内容为修改后满足 http 接口需求的内容 */

func NewGatewayApiServer(url string, transport *Transport) *GatewayApiServer {
	relay_server := relay_http.NewHttpEngine()

	cmreq := func(writer http.ResponseWriter, request *http.Request) {
		// 接收的数据格式为 GatewayData
		gwData := NewGatewayData(nil, nil)
		err := json.NewDecoder(request.Body).Decode(gwData)
		if err != nil {
			clog.Fatalf("Decode GatewayData error: %v", err)
		}

		// 从 GatewayData 中获取跨链消息和证明
		cm := utils.CmFromLb(gwData.Message)
		proof := gwData.Proof
		clog.Debugf("%v: cmreq handler get request from /cm/req", transport.chainid)
		clog.Logf("%v: cmreq handler cm.payloadReq: %v", transport.chainid, cm.PayloadReq)
		// 将返回值接入 response 中
		if transport.ReceivceReqCmFromOthers(&cm, proof) {
			// 写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "200",
				Msg:  "success",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
			clog.Debugf("%v: cmreq handler response successfully, from /cm/req", transport.chainid)
		} else {
			// 错误，写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "500",
				Msg:  "error",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
			clog.Warnf("%v: cmreq handler response error, from /cm/req", transport.chainid)
		}
	}

	cmresp := func(writer http.ResponseWriter, request *http.Request) {
		// 接收的数据格式为 GatewayData
		gwData := NewGatewayData(nil, nil)
		err := json.NewDecoder(request.Body).Decode(gwData)
		if err != nil {
			clog.Fatalf("Decode GatewayData error: %v", err)
		}
		defer request.Body.Close()

		// 从 GatewayData 中获取跨链消息和证明
		cm := utils.CmFromLb(gwData.GetMessage())
		proof := gwData.GetProof()
		clog.Debugf("%v: cmresp handler get request from /cm/resp", transport.chainid)

		// 将返回值接入 response 中
		if transport.ReceivceRespCmFromOthers(&cm, proof) {
			// 写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "200",
				Msg:  "success",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
			clog.Debugf("%v: cmresp handler response successfully, from /cm/resp", transport.chainid)
		} else {
			// 错误，写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "500",
				Msg:  "error",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
			clog.Warnf("%v: cmresp handler response error, from /cm/resp", transport.chainid)
		}
	}

	registerGateway := func(writer http.ResponseWriter, request *http.Request) {
		// 接受的数据格式如下
		var receive struct {
			Chain_id    *big.Int `json:"chain_id"`
			Gateway_url string   `json:"gateway_url"`
		}
		// 接收并解析数据
		err := json.NewDecoder(request.Body).Decode(&receive)
		if err != nil {
			clog.Fatalf("Decode data error: %v", err)
		}
		defer request.Body.Close()

		// 从 receive 中获取 chain_id 和 gateway_url
		chain_id := receive.Chain_id
		gateway_url := receive.Gateway_url

		// 将返回值接入 response 中
		if transport.RegisterGateway(chain_id.String(), gateway_url) {
			// 写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "200",
				Msg:  "success",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
		} else {
			// 错误，写入 http.ResponseWriter
			resp := struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: "500",
				Msg:  "error",
			}
			if err := json.NewEncoder(writer).Encode(resp); err != nil {
				clog.Error(err)
			}
		}
	}

	relay_server.POST("/cm/req", cmreq)
	relay_server.POST("/cm/resp", cmresp)
	relay_server.POST("/register/gateway", registerGateway)

	return &GatewayApiServer{
		url:          url,
		relay_server: relay_server,
		transport:    transport,
	}
}

/******************************** Transport ********************************/

// 传输层网关
type Transport struct {
	chainid *big.Int
	sender  utils.Address
	// FIXME: 完整 url or 端口号？
	gateway_server_url string

	// 层间交互队列
	// tspt 代表 transport, 即转发层
	// vrfy 代表 verificaiton, 即验证层
	// q_xxx_tspt_vrfy 代表转发层和验证层间的消息队列, 该队列的生产者为转发层, 消费者为验证层, 即数据从转发层流向验证层
	// q_xxx_vrfy_tspt 代表转发层和验证层间的消息队列, 该队列的生产者为验证层, 消费者为转发层, 即数据从验证层流向转发层
	q_req_tspt_vrfy      chan *utils.CrossChainMessage
	q_req_vrfy_tspt      chan *utils.CmWithProof
	q_req_pre_tspt_vrfy  chan *utils.CmWithProof
	q_req_pre_vrfy_tspt  chan *utils.CmWithProof
	q_resp_tspt_vrfy     chan *utils.CrossChainMessage
	q_resp_vrfy_tspt     chan *utils.CmWithProof
	q_resp_pre_tspt_vrfy chan *utils.CmWithProof
	q_resp_pre_vrfy_tspt chan *utils.CmWithProof

	// 用于检查发送到链上的交易是否正确执行
	// 将哈希保存在队列中, 然后轮询其是否执行成功
	// 如果执行失败, 则执行相应的策略
	q_req_txhashes  chan string
	q_resp_txhashes chan string

	// 所属链的 sdk
	bcsdk *sdk.ChainSdk

	// 转发层地址和必要的事件名称
	// 转发层暂时仅使用 eventListen 协议, 不考虑多种协议
	transport_addr      utils.Address
	transport_eventName string

	// 聚合器地址和必要的事件名称
	aggregator_addr utils.Address

	// 创建网关链接客户端, 便于和其他链的网关交互
	// 暂时设置每个 gateway 仅一个 client
	// client 改成 list, 支持多个 client
	gateway_clients    map[string]*GatewayApiClient
	enable_relay_chain bool
}

// FIXME: 待 sdk 完善后需进一步修改
func NewTransport(bcsdk *sdk.ChainSdk, config *config.Config) *Transport {
	gateway_clients := make(map[string]*GatewayApiClient)

	for chainid, url := range config.Gateway_urls {
		gateway_clients[chainid] = NewGatewayApiClient(url)
	}

	return &Transport{
		chainid:            config.Chainid,
		sender:             config.Sender,
		gateway_server_url: config.Gateway_server_url,

		q_req_tspt_vrfy:      config.Q_req_tspt_vrfy,
		q_req_vrfy_tspt:      config.Q_req_vrfy_tspt,
		q_req_pre_tspt_vrfy:  config.Q_req_pre_tspt_vrfy,
		q_req_pre_vrfy_tspt:  config.Q_req_pre_vrfy_tspt,
		q_resp_tspt_vrfy:     config.Q_resp_tspt_vrfy,
		q_resp_vrfy_tspt:     config.Q_resp_vrfy_tspt,
		q_resp_pre_tspt_vrfy: config.Q_resp_pre_tspt_vrfy,
		q_resp_pre_vrfy_tspt: config.Q_resp_pre_vrfy_tspt,

		q_req_txhashes:  make(chan string),
		q_resp_txhashes: make(chan string),

		bcsdk: bcsdk,

		transport_addr:      config.Transport_addr,
		transport_eventName: config.Transport_eventName,

		aggregator_addr: config.Aggregator_addr,

		gateway_clients:    gateway_clients,
		enable_relay_chain: config.EnableRelayChain,
	}
}

// 注册网关
func (t *Transport) RegisterGateway(chain_id string, gateway_url string) bool {
	client := NewGatewayApiClient(gateway_url)
	t.gateway_clients[chain_id] = client
	return true
}

// 监听链上的跨链消息，并转发至验证层的队列 q_tspt_vrfy
// FIXME: 这里需要根据 bcsdk 返回的数据进行进一步修改
func (t *Transport) Watch() {
	// event_name := "CmHash"
	// event_key := t.bcsdk.SubscribeEvent(t.transport_addr, event_name)

	cm_event_chan := make(chan map[string][]byte)
	go t.bcsdk.ListenEvent(cm_event_chan)

	for {
		cm_event := <-cm_event_chan
		// go binary 的 []byte 转 int 对 []byte 长度有要求
		phase_bytes := append(cm_event["phase"], []byte{0, 0, 0, 0}...)
		phase := binary.LittleEndian.Uint32(phase_bytes)
		cmhash := cm_event["hash"]
		var cmhash32 [32]byte
		copy(cmhash32[:], cmhash)
		if phase == 1 {
			// 源链上的请求消息
			clog.Debugf("Obtained request cmhash on the source chain(chainid: %v, cmhash: %x)", t.chainid, cmhash)
			cmLb := t.bcsdk.QueryReqCmByHash(cmhash32)
			cm := utils.CmFromLb(cmLb)

			clog.Debugf("%v: Watch: send cm to q_req_tspt_vrfy", t.chainid)
			t.q_req_tspt_vrfy <- &cm
		} else if phase == 2 {
			// 目的链上发出的响应消息
			clog.Infof("Obtained response cmhash on the target chain(chainid: %v, cmhash: %x)", t.chainid, cmhash)
			cmLb := t.bcsdk.QueryRespCmByHash(cmhash32)
			cm := utils.CmFromLb(cmLb)

			clog.Debugf("%v: Watch: send cm to q_resp_tspt_vrfy", t.chainid)
			t.q_resp_tspt_vrfy <- &cm
		} else if phase == 3 {
			// 源链上发出的确认消息
			clog.Infof("Obtained ack cmhash on the source chain(chainid: %v, cmhash: %x)", t.chainid, cmhash)
			// 对于 ack 消息，网关不需要做任何事情
		} else {
			clog.Errorf("Error. Invalid cm phase: %v", phase)
		}
	}
}

// 作为交互逻辑中的源链网关，从验证层队列 q_req_vrfy_tspt 获取证明，并发送给附属于目的链的网关
func (t *Transport) SendReqToTargetRelay() {
	for {
		// 1. 接收验证层返回的内容，验证层传回的内容包括跨链消息和证明
		item := <-t.q_req_vrfy_tspt
		clog.Debugf("%v: SendReqToTargetRelay: get cm with proof from q_req_vrfy_tspt", t.chainid)

		cm := item.Cm
		proof := item.Proof

		DstChainId := ""
		if t.enable_relay_chain {
			DstChainId = "30000"
			clog.Debugf("Enable Relay Chain, transport to 30000")
		} else {
			DstChainId = cm.DstChainId.String()
		}

		// 2. 寻找目的链网关
		if t.gateway_clients[DstChainId] == nil {
			clog.Errorf("cm.dstChainId(%v) not in gateway_clients", DstChainId)
			clog.Errorf("cm.seq: %v", cm.Seq)
			continue
		}
		gw_api_client := t.gateway_clients[DstChainId]

		// 3. 向目的链网关发送 reqCm，返回值 content 的内容就是 reviceReqCmFromOthers() 函数的返回值
		// 如果涉及到编解码，那么编解码操作应该放在 GatewayApiClient 中做统一的编解码操作
		// 这样，使得 sendReqCm 操作就仿佛是在调用 receiveReqCmFromOthers() 一样
		clog.Debugf("%v: SendReqToTargetRelay: send cm with proof to gateway", t.chainid)
		content, err := gw_api_client.SendReqCm(cm, proof)
		if err != "" {
			clog.Errorf("%v: %v", t.chainid, err)
		}
		clog.Debugf("%v: SendReqToTargetRelay: send cm with proof to gateway successfully get response", t.chainid)

		// content 在这里有何作用？
		_ = content
	}
}

// 作为交互逻辑中的目的链网关，从验证层队列 q_resp_vrfy_tspt 获取证明，并发送给附属于源链的网关
func (t *Transport) SendRespToSourceRelay() {
	for {
		// 1. 接收验证层返回的内容，验证层传回的内容包括跨链消息和证明
		item := <-t.q_resp_vrfy_tspt
		clog.Debugf("%v: SendRespToSourceRelay: get cm with proof from q_resp_vrfy_tspt", t.chainid)

		cm := item.Cm
		proof := item.Proof
		// fmt.Printf("cm.srcChainId: %v\n", cm.SrcChainId)
		// fmt.Printf("cm.seq: %v\n", cm.Seq)

		SrcChainId := ""
		if t.enable_relay_chain {
			SrcChainId = "30000"
			clog.Debugf("Enable Relay Chain, transport to 30000")
		} else {
			SrcChainId = cm.SrcChainId.String()
		}

		// 2. 寻找源链网关
		if t.gateway_clients[SrcChainId] == nil {
			fmt.Printf("cm.srcChainId(%v) not in gateway_clients\n", cm.SrcChainId)
			fmt.Printf("cm.seq: %v\n", cm.Seq)
			continue
		}
		gw_api_client := t.gateway_clients[SrcChainId]

		// 3. 向源链网关发送 respCm，返回值 content 的内容就是 reviceRespCmFromOthers() 函数的返回值
		clog.Debugf("%v: SendRespToSourceRelay: send cm with proof to gateway", t.chainid)
		content, err := gw_api_client.SendRespCm(cm, proof)
		if err != "" {
			clog.Errorf("%v: %v", t.chainid, err)
		}
		clog.Debugf("%v: SendRespToSourceRelay: send cm with proof to gateway successfully get response", t.chainid)

		// content 在这里有何作用？
		_ = content
	}
}

// 作为交互逻辑中的目的链网关, 通过 GatewayApiServer 接收来自其他源链网关转发层的数据
func (t *Transport) ReceivceReqCmFromOthers(cm *utils.CrossChainMessage, proof []byte) bool {
	// 判断目的链标识是否相同
	if cm.DstChainId.Cmp(t.chainid) != 0 {
		clog.Errorf("%v: cm.dstChainId(%v) != t.chainid(%v)\n", t.chainid, cm.DstChainId, t.chainid)
		return false
	}

	clog.Debugf("%v: ReceivceReqCmFromOthers: send item to q_req_pre_tspt_vrfy", t.chainid)
	t.q_req_pre_tspt_vrfy <- &utils.CmWithProof{
		Cm:    cm,
		Proof: proof,
	}
	return true
}

// 作为交互逻辑中的源链网关, 通过 GatewayApiServer 接收来自其他目的链网关转发层的数据
func (t *Transport) ReceivceRespCmFromOthers(cm *utils.CrossChainMessage, proof []byte) bool {
	// 判断源链标识是否相同
	if cm.SrcChainId.Cmp(t.chainid) != 0 {
		clog.Errorf("%v: cm.srcChainId(%v) != t.chainid(%v)\n", t.chainid, cm.SrcChainId, t.chainid)
		return false
	}

	clog.Debugf("%v: ReceivceRespCmFromOthers: send item to q_resp_pre_tspt_vrfy", t.chainid)
	t.q_resp_pre_tspt_vrfy <- &utils.CmWithProof{
		Cm:    cm,
		Proof: proof,
	}
	return true
}

// 作为交互逻辑中的目的链网关, 从验证层的 q_req_pre_vrfy_tspt 队列中获取数据并编码发送到本链(也就是目的链)
func (t *Transport) SendReqToTargetChain() {
	for {
		// 1. 接收验证层返回的内容, 验证层传回的内容包括跨链消息和证明
		// 错误处理, 假设上级协议执行失败, 那么如何响应错误
		item := <-t.q_req_pre_vrfy_tspt
		clog.Debugf("%v: SendReqToTargetChain: get cm with proof from q_req_pre_vrfy_tspt", t.chainid)
		clog.Logf("%v: SendReqToTargetChain: cm.PayloadReq: %v", t.chainid, string(item.Cm.PayloadReq[0]))
		cm := item.Cm
		proof := item.Proof

		// 2. 向自己所在的链(也就是交互逻辑的目的链)发送请求交易
		txhash := t.bcsdk.ReceiveMsg(*cm, proof)
		// 在 dealReqTxHashes 中处理
		t.q_req_txhashes <- txhash
		clog.Debugf("%v: SendReqToTargetChain: send txhash to q_req_txhashes", t.chainid)
	}
}

// 作为交互逻辑中的源链网关, 从验证层的 q_resp_pre_vrfy_tspt 队列中获取数据并编码发送到本链(也就是源链)
func (t *Transport) SendRespToSourceChain() {
	for {
		// 1. 接收验证层返回的内容, 验证层传回的内容包括跨链消息和证明
		// 错误处理, 假设上级协议执行失败, 那么如何响应错误
		item := <-t.q_resp_pre_vrfy_tspt
		clog.Debugf("%v: SendRespToSourceChain: get cm with proof from q_resp_pre_vrfy_tspt", t.chainid)

		cm := item.Cm
		proof := item.Proof

		// 2. 向自己所在的链(也就是交互逻辑的源链)发送响应交易
		txhash := t.bcsdk.AcknowledgeMsg(*cm, proof)
		t.q_resp_txhashes <- txhash
		clog.Debugf("%v: SendRespToSourceChain: send txhash to q_resp_txhashes", t.chainid)
	}
}

// 作为交互逻辑中的目的链, 检查发送到本链的跨链交易是否成功, 从而判断是否应该重传或者执行其他操作
func (t *Transport) DealReqTxHashes() {
	for {
		txhash := <-t.q_req_txhashes
		clog.Debugf("%v: DealReqTxHashes: get txhash from q_req_txhashes", t.chainid)

		// 处理跨链请求交易哈希
		// NOW，在主逻辑中添加线程
		clog.Debugf("get req txhash: %v", txhash)
	}
}

// 作为交互逻辑中的源链, 检查发送到本链的跨链交易是否成功, 从而判断是否应该重传或者执行其他操作
func (t *Transport) DealRespTxHashes() {
	for {
		txhash := <-t.q_resp_txhashes
		clog.Debugf("%v: DealRespTxHashes: get txhash from q_resp_txhashes", t.chainid)

		// 处理跨链响应交易哈希
		// NOW，在主逻辑中添加线程
		clog.Debugf("get resp txhash: %v", txhash)
	}
}

// 主要分为四个线程
// (1) watch 线程, 用于向链上获取最新的跨链消息
// (2) gatewayServer 线程, 用于接收来自其他网关的消息
// (3) reqHashes 处理线程
// (4) respHashes 处理线程
func (t *Transport) Start() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 1. 启动 watch 线程
	go t.Watch()
	clog.Infof("%v: start transport layer watch thread!", t.chainid)

	// 2. 启动 server 线程
	server := NewGatewayApiServer(t.gateway_server_url, t)
	parts := strings.Split(t.gateway_server_url, ":")
	port := parts[len(parts)-1]
	// 使用端口号启动 server
	// 因此要求在 gateway_server_url 中必须显示的给出端口号
	go server.relay_server.Run(":" + port)
	clog.Infof("%v: start transport layer gateway server thread! url: %s", t.chainid, t.gateway_server_url)

	// 3.1 启动线程 sendReqToTargetRelay
	go t.SendReqToTargetRelay()
	clog.Infof("%v: start transport layer sendReqToTargetRelay thread!", t.chainid)

	// 3.2 启动线程 sendRespToSourceRelay
	go t.SendRespToSourceRelay()
	clog.Infof("%v: start transport layer sendRespToSourceRelay thread!", t.chainid)

	// 3.3 启动线程 sendReqToTargetChain
	go t.SendReqToTargetChain()
	clog.Infof("%v: start transport layer sendReqToTargetChain thread!", t.chainid)

	// 3.4 启动线程 sendRespToSourceChain
	go t.SendRespToSourceChain()
	clog.Infof("%v: start transport layer sendRespToSourceChain thread!", t.chainid)

	// 4. 处理 req 哈希
	go t.DealReqTxHashes()
	clog.Infof("%v: start transport layer dealReqTxHashes thread!", t.chainid)

	// 4. 处理 resp 哈希
	go t.DealRespTxHashes()
	clog.Infof("%v: start transport layer dealRespTxHashes thread!", t.chainid)

	wg.Wait()
}
