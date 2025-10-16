package layer_verification

import (
	"crossFab/relayer/config"
	"crossFab/relayer/sdk"
	"crossFab/utils"
	"math/big"
	"sync"

	clog "github.com/kpango/glg"
)

func VerificationProtocolSelector(vrfy_id string, args [][]byte) VerificationInterface {
	if vrfy_id == "1" {
		return NewVerificationSPV(args)
	} else if vrfy_id == "2" {
		return NewVerificationNotary(args)
	} else {
		clog.Errorf("Unknown verification typdId: %v", vrfy_id)
		return nil
	}
}

type Verification struct {
	q_req_tspt_vrfy      chan *utils.CrossChainMessage
	q_req_vrfy_tspt      chan *utils.CmWithProof
	q_req_pre_tspt_vrfy  chan *utils.CmWithProof
	q_req_pre_vrfy_tspt  chan *utils.CmWithProof
	q_resp_tspt_vrfy     chan *utils.CrossChainMessage
	q_resp_vrfy_tspt     chan *utils.CmWithProof
	q_resp_pre_tspt_vrfy chan *utils.CmWithProof
	q_resp_pre_vrfy_tspt chan *utils.CmWithProof

	// 获取验证层地址和必要的事件名称
	verification_protocols map[string]VerificationInterface

	// 所属链的 sdk
	bcsdk *sdk.ChainSdk

	// 测试用
	Chainid *big.Int
}

func NewVerification(bcsdk *sdk.ChainSdk, config *config.Config) *Verification {
	v := &Verification{
		Chainid:                config.Chainid,
		q_req_tspt_vrfy:        config.Q_req_tspt_vrfy,
		q_req_vrfy_tspt:        config.Q_req_vrfy_tspt,
		q_req_pre_tspt_vrfy:    config.Q_req_pre_tspt_vrfy,
		q_req_pre_vrfy_tspt:    config.Q_req_pre_vrfy_tspt,
		q_resp_tspt_vrfy:       config.Q_resp_tspt_vrfy,
		q_resp_vrfy_tspt:       config.Q_resp_vrfy_tspt,
		q_resp_pre_tspt_vrfy:   config.Q_resp_pre_tspt_vrfy,
		q_resp_pre_vrfy_tspt:   config.Q_resp_pre_vrfy_tspt,
		verification_protocols: make(map[string]VerificationInterface),
		bcsdk:                  bcsdk,
	}

	for vrfy_id, vrfy_args := range config.Verifications {
		clog.Infof("Register verification protocol %s", vrfy_id)
		v.verification_protocols[vrfy_id] = VerificationProtocolSelector(vrfy_id, vrfy_args)
	}

	return v
}

func (v *Verification) watch_q_req_tspt_vrfy() {
	for {
		cm := <-v.q_req_tspt_vrfy
		clog.Debugf("%v: watch_q_req_tspt_vrfy: get cm from q_req_tspt_vrfy", v.Chainid)
		clog.Logf("%v: watch_q_req_tspt_vrfy: cm.payloadReq: %v", v.Chainid, cm)

		vrfy_id := cm.VerificationTypeId
		if verification, ok := v.verification_protocols[vrfy_id.String()]; ok {
			proof := verification.GenProof(cm)
			clog.Infof("%v: generate proof for req cm(%v): %v", v.Chainid, cm.Seq, proof)
			v.q_req_vrfy_tspt <- &utils.CmWithProof{
				Cm:    cm,
				Proof: proof,
			}
			clog.Logf("%v: cm.payloadReq after proof: %v", v.Chainid, cm.PayloadReq)
			clog.Debugf("%v: watch_q_req_tspt_vrfy: send cm with proof to q_req_vrfy_tspt", v.Chainid)
		} else {
			clog.Warnf("%v: vrfy_id %v is invalid", v.Chainid, vrfy_id)
		}
	}
}

func (v *Verification) watch_q_req_pre_tspt_vrfy() {
	for {
		item := <-v.q_req_pre_tspt_vrfy
		clog.Debugf("%v: watch_q_req_pre_tspt_vrfy: get cm with proof from q_req_pre_tspt_vrfy", v.Chainid)
		clog.Logf("%v: watch_q_req_pre_tspt_vrfy: item.cm.payloadReq: %v", v.Chainid, item.Cm.PayloadReq)

		cm := item.Cm
		proof := item.Proof
		vrfy_id := cm.VerificationTypeId
		if verification, ok := v.verification_protocols[vrfy_id.String()]; ok {
			flag := verification.PreSend(cm)
			if !flag {
				clog.Infof("%v: pre send failed for cm(%v)", v.Chainid, cm.Seq)
				continue
			}
			v.q_req_pre_vrfy_tspt <- &utils.CmWithProof{
				Cm:    cm,
				Proof: proof,
			}
			clog.Debugf("%v: watch_q_req_pre_tspt_vrfy: send cm with proof to q_req_pre_vrfy_tspt", v.Chainid)
		} else {
			clog.Warnf("%v: vrfy_id %v is invalid", v.Chainid, vrfy_id)
		}
	}
}

func (v *Verification) watch_q_resp_tspt_vrfy() {
	for {
		cm := <-v.q_resp_tspt_vrfy
		clog.Debugf("%v: watch_q_resp_tspt_vrfy: get cm from q_resp_tspt_vrfy", v.Chainid)

		vrfy_id := cm.VerificationTypeId
		if verification, ok := v.verification_protocols[vrfy_id.String()]; ok {
			proof := verification.GenProof(cm)
			clog.Infof("%v: generate proof for resp cm(%v): %v", v.Chainid, cm.Seq, proof)
			v.q_resp_vrfy_tspt <- &utils.CmWithProof{
				Cm:    cm,
				Proof: proof,
			}
			clog.Debugf("%v: watch_q_resp_tspt_vrfy: send cm with proof to q_resp_vrfy_tspt", v.Chainid)
		} else {
			clog.Warnf("%v: vrfy_id %v is invalid", v.Chainid, vrfy_id)
		}
	}
}

func (v *Verification) watch_q_resp_pre_tspt_vrfy() {
	for {
		item := <-v.q_resp_pre_tspt_vrfy
		clog.Debugf("%v: watch_q_resp_pre_tspt_vrfy: get cm with proof from q_resp_pre_tspt_vrfy", v.Chainid)

		cm := item.Cm
		proof := item.Proof
		vrfy_id := cm.VerificationTypeId
		if verification, ok := v.verification_protocols[vrfy_id.String()]; ok {
			flag := verification.PreSend(cm)
			if !flag {
				clog.Infof("%v: pre send failed for cm(%v)", v.Chainid, cm.Seq)
				continue
			}
			v.q_resp_pre_vrfy_tspt <- &utils.CmWithProof{
				Cm:    cm,
				Proof: proof,
			}
			clog.Debugf("%v: watch_q_resp_pre_tspt_vrfy: send cm with proof to q_resp_pre_vrfy_tspt", v.Chainid)
		} else {
			clog.Warnf("%v: vrfy_id %v is invalid", v.Chainid, vrfy_id)
		}
	}
}

// 启动线程，监听来自转发层消息队列
func (v *Verification) watch_q() {
	var wg sync.WaitGroup
	wg.Add(1)

	go v.watch_q_req_tspt_vrfy()
	clog.Infof("%v: start verification layer watch_q_req_tspt_vrfy thread!", v.Chainid)

	go v.watch_q_req_pre_tspt_vrfy()
	clog.Infof("%v: start verification layer watch_q_req_pre_tspt_vrfy thread!", v.Chainid)

	go v.watch_q_resp_tspt_vrfy()
	clog.Infof("%v: start verification layer watch_q_resp_tspt_vrfy thread!", v.Chainid)

	go v.watch_q_resp_pre_tspt_vrfy()
	clog.Infof("%v: start verification layer watch_q_resp_pre_tspt_vrfy thread!", v.Chainid)

	// 不考虑中断，直接阻塞协程
	wg.Wait()
}

func (v *Verification) Start() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 启动各个协议的 update 函数
	for vrfy_id, protocol := range v.verification_protocols {
		go protocol.Update()
		clog.Infof("%v: start verification protocol: %v", v.Chainid, vrfy_id)
	}

	// 监听转发层的数据
	go v.watch_q()

	wg.Wait()
}
