package layer_verification

import (
	"crossFab/utils"
	"time"

	clog "github.com/kpango/glg"
)

type VerificationNotary struct{}

func NewVerificationNotary(args [][]byte) *VerificationNotary {
	_ = args
	return &VerificationNotary{}
}

func (vn *VerificationNotary) Update() {
	for {
		clog.Info("Notary update: XXX")
		time.Sleep(100 * time.Second)
	}
}

func (vn *VerificationNotary) GenProof(cm *utils.CrossChainMessage) []byte {
	return []byte("proof of Notary")
}

func (vn *VerificationNotary) PreSend(cm *utils.CrossChainMessage) bool {
	return true
}
