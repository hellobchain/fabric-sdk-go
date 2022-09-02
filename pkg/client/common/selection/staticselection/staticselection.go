/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package staticselection

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/client/common/selection/options"
	copts "github.com/wsw365904/fabric-sdk-go/pkg/common/options"
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/wsw365904/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

// SelectionService implements static selection service
type SelectionService struct {
	discoveryService fab.DiscoveryService
	peers            []fab.CompletePeer
}

func (s *SelectionService) SetPeers(peers []fab.CompletePeer) {
	s.peers = peers
}

// NewService creates a static selection service
func NewService(discovery fab.DiscoveryService) (fab.SelectionService, error) {
	return &SelectionService{discoveryService: discovery}, nil
}

// GetEndorsersForChaincode returns a set of endorsing peers
func (s *SelectionService) GetEndorsersForChaincode(chaincodes []*fab.ChaincodeCall, opts ...copts.Opt) ([]fab.Peer, error) {
	params := options.NewParams(opts)

	channelPeers, err := s.discoveryService.GetPeers()
	if err != nil {
		logger.Errorf("Error retrieving peers from discovery service: %s", err)
		return nil, nil
	}

	// Apply peer filter if provided
	if params.PeerFilter != nil {
		var peers []fab.Peer
		for _, peer := range channelPeers {
			if params.PeerFilter(peer) {
				peers = append(peers, peer)
			}
		}
		channelPeers = peers
	}

	if params.PeerSorter != nil {
		peers := make([]fab.Peer, len(channelPeers))
		copy(peers, channelPeers)
		channelPeers = params.PeerSorter(peers)
	}

	if logger.IsEnabledFor(wlogging.PayloadLevel + 1) {
		str := ""
		for i, peer := range channelPeers {
			str += peer.URL()
			if i+1 < len(channelPeers) {
				str += ","
			}
		}
		logger.Debugf("Available peers: [%s]", str)
	}

	return channelPeers, nil
}
