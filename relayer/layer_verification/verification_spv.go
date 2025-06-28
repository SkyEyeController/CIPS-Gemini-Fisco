package layer_verification

import (
	"crossFab/utils"
	"time"

	clog "github.com/kpango/glg"
)

type VerificationSPV struct{}

func NewVerificationSPV(args [][]byte) *VerificationSPV {
	_ = args
	return &VerificationSPV{}
}

func (vs *VerificationSPV) Update() {
	for {
		clog.Info("SPV update: XXX")
		time.Sleep(100 * time.Second)
	}
}

func (vs *VerificationSPV) GenProof(cm *utils.CrossChainMessage) []byte {
	return []byte("proof of SPV")
}

func (vs *VerificationSPV) PreSend(cm *utils.CrossChainMessage) bool {
	return true
}
