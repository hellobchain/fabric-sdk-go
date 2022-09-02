/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/fab"
)

// MockDiscoveryService is a mock discovery service used for event endpoint discovery
type MockDiscoveryService struct {
	peers []fab.Peer
}

func (s *MockDiscoveryService) SetPeers(peers []fab.CompletePeer) {
	fpeer := make([]fab.Peer, len(peers))

	for i, peer := range peers {
		fpeer[i] = peer.Peer
	}
	s.peers = fpeer
}

// NewDiscoveryService returns a new MockDiscoveryService
func NewDiscoveryService(peers ...fab.Peer) fab.DiscoveryService {
	return &MockDiscoveryService{
		peers: peers,
	}
}

// GetPeers returns a list of discovered peers
func (s *MockDiscoveryService) GetPeers() ([]fab.Peer, error) {
	return s.peers, nil
}
