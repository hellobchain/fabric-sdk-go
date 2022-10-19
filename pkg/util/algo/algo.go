package algo

import (
	"encoding/pem"
	"github.com/hellobchain/newcryptosm"
	"github.com/hellobchain/newcryptosm/x509"
	"github.com/hellobchain/third_party/algo"
	"github.com/hellobchain/wswlog/wlogging"
	"sync"
)

var logger = wlogging.MustGetLoggerWithoutName()

func SetGMFlag() {
	algo.SetGMFlag()
}

func GetGMFlag() bool {
	return algo.GetGMFlag()
}

func GetDefaultHash() newcryptosm.Hash {
	return algo.GetDefaultHash()
}

func GetAlgo() string {
	return algo.GetAlgo()
}

var once sync.Once

func JudgeAlgo(pemBytes []byte) {
	once.Do(func() {
		p, _ := pem.Decode(pemBytes)
		if p == nil {
			logger.Error("pemBytes is invalid")
		}
		parseCertificate, err := x509.ParseCertificate(p.Bytes)
		if err != nil {
			logger.Error("parse cert failed", err)
			return
		}
		if parseCertificate.SignatureAlgorithm == x509.SM2WithSM3 {
			SetGMFlag()
		}
	})
}
