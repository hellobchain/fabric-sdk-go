package hash

import (
	"github.com/hellobchain/fabric-sdk-go/internal/github.com/hyperledger/fabric/protoutil"
	cb "github.com/hyperledger/fabric-protos-go/common"
)

func ComputeCurrentHash(b *cb.BlockHeader) []byte {
	return protoutil.BlockHeaderHash(b)
}
