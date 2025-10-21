module crossFab

go 1.19



require (

	// FISCO BCOS Go SDK v3
	github.com/FISCO-BCOS/go-sdk/v3 v3.0.0
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect

	// 以太坊依赖
	github.com/ethereum/go-ethereum v1.13.10
	github.com/golang/protobuf v1.5.4 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

require (
	github.com/kpango/glg v1.6.15
	github.com/spf13/cobra v1.5.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/FISCO-BCOS/bcos-c-sdk v0.0.0-20240219081048-53240138c396 // indirect
	github.com/FISCO-BCOS/crypto v0.0.0-20200202032121-bd8ab0b5d4f1 // indirect
	github.com/TarsCloud/TarsGo v1.4.5 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/ethereum/c-kzg-4844 v0.4.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kpango/fastime v1.1.9 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

// 排除冲突的 btcd 版本
exclude github.com/btcsuite/btcd v0.20.1-beta

// 强制使用特定版本解决冲突
replace github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.24.0
