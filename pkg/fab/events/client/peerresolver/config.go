/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peerresolver

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hellobchain/fabric-sdk-go/pkg/fab/events/client/lbp"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

// GetBalancer returns the configured load balancer
func GetBalancer(policy fab.EventServicePolicy) lbp.LoadBalancePolicy {
	switch policy.Balancer {
	case fab.RoundRobin:
		logger.Debugf("Using round-robin load balancer.")
		return lbp.NewRoundRobin()
	case fab.Random:
		logger.Debugf("Using random load balancer.")
		return lbp.NewRandom()
	default:
		logger.Debugf("Balancer not specified. Using random load balancer.")
		return lbp.NewRandom()
	}
}
