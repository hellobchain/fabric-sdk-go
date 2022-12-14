/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab"
)

// MockDiscoveryService is a mock discovery service used for event endpoint discovery
type MockDiscoveryService struct {
	peers []fab.Peer
}

func (s *MockDiscoveryService) SetPeersOfChannel(peers fab.CompletePeer) {
	s.peers = peers.Peers
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
