package layer_verification

import "crossFab/utils"

type VerificationInterface interface {
	// 用于持续更新该验证协议的必要内容，比如 spv 的区块头数据
	Update()

	GenProof(cm *utils.CrossChainMessage) []byte

	// 主要起到过滤的作用
	PreSend(cm *utils.CrossChainMessage) bool
}
