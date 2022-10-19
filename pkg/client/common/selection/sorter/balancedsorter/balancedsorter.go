/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package balancedsorter

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/client/common/selection/options"
	coptions "github.com/hellobchain/fabric-sdk-go/pkg/common/options"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

// New returns a peer sorter that chooses a peer according to a provided balancer.
func New(opts ...coptions.Opt) options.PeerSorter {
	params := defaultParams()
	coptions.Apply(params, opts)

	return func(peers []fab.Peer) []fab.Peer {
		return params.balancer(peers)
	}
}
