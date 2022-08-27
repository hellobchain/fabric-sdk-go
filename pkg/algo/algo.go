package algo

import (
	"encoding/pem"
	"github.com/wsw365904/newcryptosm"
	"github.com/wsw365904/newcryptosm/x509"
	"github.com/wsw365904/third_party/algo"
	"github.com/wsw365904/wswlog/wlogging"
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
