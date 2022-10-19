/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fab

import (
	reqContext "context"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab/orderer"

	"github.com/hyperledger/fabric-protos-go/common"
)

// Orderer The Orderer class represents a peer in the target blockchain network to which
// HFC sends a block of transactions of endorsed proposals requiring ordering.
type Orderer interface {
	URL() string
	SendBroadcast(ctx reqContext.Context, envelope *SignedEnvelope) (*common.Status, error)
	SendDeliver(ctx reqContext.Context, envelope *SignedEnvelope) (chan *common.Block, chan error)
	SetSensitiveWords(ctx reqContext.Context, in *orderer.SensitiveWord) error
	QuerySensitiveWords(ctx reqContext.Context) ([]string, error)
	AddSensitiveWords(ctx reqContext.Context, in *orderer.SensitiveWord) error
	SetExcludeWords(ctx reqContext.Context, in *orderer.ExcludedSymbol) error
	QueryExcludeWords(ctx reqContext.Context) ([]string, error)
	AddExcludeWords(ctx reqContext.Context, in *orderer.ExcludedSymbol) error
}

// A SignedEnvelope can can be sent to an orderer for broadcasting
type SignedEnvelope struct {
	Payload   []byte
	Signature []byte
}
