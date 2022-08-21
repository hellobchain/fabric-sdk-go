package algo

import (
	flogging "github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/sdkpatch/logbridge"
	"github.com/hyperledger/fabric-sdk-go/third_party/smalgo/x509"
	"github.com/spf13/viper"
)

var logger = flogging.MustGetLogger("pkg.algo")

func SetGMFlag() {
	logger.Info("SetGMFlag")
	viper.Set("GMFlag", true)
}

func GetGMFlag() bool {
	algoFlag := viper.GetBool("GMFlag")
	logger.Info("GetGMFlag:", algoFlag)
	return algoFlag
}

func GetDefaultHash() x509.Hash {
	if GetGMFlag() {
		return x509.SM3
	} else {
		return x509.SHA256
	}
}

func GetAlgo() string {
	if GetGMFlag() {
		return "sm2"
	} else {
		return "ecdsa"
	}
}
