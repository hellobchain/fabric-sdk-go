package hash

import (
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/wsw365904/fabric-sdk-go/internal/github.com/hyperledger/fabric/protoutil"
)

func ComputeCurrentHash(b *cb.BlockHeader) []byte {
	return protoutil.BlockHeaderHash(b)
}
