package algo

import (
	"github.com/wsw365904/newcryptosm"
	"github.com/wsw365904/third_party/algo"
)

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
