/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chpvdr

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/client/common/discovery/dynamicdiscovery"
	"github.com/hellobchain/fabric-sdk-go/pkg/client/common/discovery/staticdiscovery"
	"github.com/hellobchain/fabric-sdk-go/pkg/client/common/selection/dynamicselection"
	"github.com/hellobchain/fabric-sdk-go/pkg/client/common/selection/fabricselection"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/options"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hellobchain/fabric-sdk-go/pkg/fab/channel/membership"
	"github.com/hellobchain/fabric-sdk-go/pkg/fab/chconfig"
	"github.com/hellobchain/fabric-sdk-go/pkg/fab/events/deliverclient"
	"github.com/hellobchain/fabric-sdk-go/pkg/util/concurrent/lazycache"
	"github.com/pkg/errors"
)

type cache interface {
	Get(lazycache.Key, ...interface{}) (interface{}, error)
	Close()
}

type contextCache struct {
	ctx                   fab.ClientContext
	eventServiceCache     cache
	discoveryServiceCache cache
	selectionServiceCache cache
	chCfgCache            cache
	membershipCache       cache
	peers                 fab.CompletePeer
}

var cfgCacheProvider = func(opts ...options.Opt) cache {
	return chconfig.NewRefCache(opts...)
}

func (c *contextCache) SetPeersOfChannel(peers fab.CompletePeer) {
	c.peers = peers
}

func newContextCache(ctx fab.ClientContext, opts []options.Opt) *contextCache {
	eventIdleTime := ctx.EndpointConfig().Timeout(fab.EventServiceIdle)
	chConfigRefresh := ctx.EndpointConfig().Timeout(fab.ChannelConfigRefresh)
	membershipRefresh := ctx.EndpointConfig().Timeout(fab.ChannelMembershipRefresh)

	c := &contextCache{
		ctx: ctx,
	}

	c.chCfgCache = cfgCacheProvider(append([]options.Opt{chconfig.WithRefreshInterval(chConfigRefresh)}, opts...)...)
	c.membershipCache = membership.NewRefCache(append([]options.Opt{membership.WithRefreshInterval(membershipRefresh)}, opts...)...)

	c.discoveryServiceCache = lazycache.New(
		"Discovery_Service_Cache",
		func(key lazycache.Key) (interface{}, error) {
			ck := key.(*cacheKey)
			return c.createDiscoveryService(ck.channelConfig, opts...)
		},
	)

	c.selectionServiceCache = lazycache.New(
		"Selection_Service_Cache",
		func(key lazycache.Key) (interface{}, error) {
			ck := key.(*cacheKey)
			return c.createSelectionService(ck.channelConfig, opts...)
		},
	)

	c.eventServiceCache = lazycache.New(
		"Event_Service_Cache",
		func(key lazycache.Key) (interface{}, error) {
			ck := key.(*eventCacheKey)
			return NewEventClientRef(
				eventIdleTime,
				func() (fab.EventClient, error) {
					return c.createEventClient(ck.channelConfig, append(ck.opts, opts...)...)
				},
			), nil
		},
	)

	return c
}

func (c *contextCache) Close() {
	logger.Debug("Closing event service cache...")
	c.eventServiceCache.Close()

	logger.Debug("Closing membership cache...")
	c.membershipCache.Close()

	logger.Debug("Closing channel configuration cache...")
	c.chCfgCache.Close()

	logger.Debug("Closing selection service cache...")
	c.selectionServiceCache.Close()

	logger.Debug("Closing discovery service cache...")
	c.discoveryServiceCache.Close()
}

func (c *contextCache) createEventClient(chConfig fab.ChannelCfg, opts ...options.Opt) (fab.EventClient, error) {
	discovery, err := c.GetDiscoveryService(chConfig.ID())
	if err != nil {
		return nil, errors.WithMessage(err, "could not get discovery service")
	}

	logger.Debugf("Using deliver events for channel [%s]", chConfig.ID())
	return deliverclient.New(c.ctx, chConfig, c.peers, discovery, opts...)
}

func (c *contextCache) createDiscoveryService(chConfig fab.ChannelCfg, opts ...options.Opt) (fab.DiscoveryService, error) {
	if chConfig.HasCapability(fab.ApplicationGroupKey, fab.V1_2Capability) {
		logger.Debugf("Using Dynamic Discovery based on V1_2 capability.")
		membership, err := c.GetMembership(chConfig.ID())
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create discovery service")
		}
		return dynamicdiscovery.NewChannelService(c.ctx, membership, chConfig.ID(), opts...)
	}
	return staticdiscovery.NewService(c.ctx.EndpointConfig(), c.peers, c.ctx.InfraProvider(), chConfig.ID())
}

func (c *contextCache) GetDiscoveryService(channelID string) (fab.DiscoveryService, error) {
	chnlCfg, err := c.GetChannelConfig(channelID)
	if err != nil {
		return nil, err
	}
	key := newCacheKey(chnlCfg)
	discoveryService, err := c.discoveryServiceCache.Get(key)
	if err != nil {
		return nil, err
	}
	ds := discoveryService.(fab.DiscoveryService)
	ds.SetPeersOfChannel(c.peers)
	return ds, nil
}

func (c *contextCache) createSelectionService(chConfig fab.ChannelCfg, opts ...options.Opt) (fab.SelectionService, error) {
	discovery, err := c.GetDiscoveryService(chConfig.ID())
	if err != nil {
		return nil, err
	}

	if chConfig.HasCapability(fab.ApplicationGroupKey, fab.V1_2Capability) {
		logger.Debugf("Using Fabric Selection based on V1_2 capability.")
		return fabricselection.New(c.ctx, chConfig.ID(), discovery, opts...)
	}
	return dynamicselection.NewService(c.ctx, chConfig.ID(), discovery)
}

func (c *contextCache) GetSelectionService(channelID string) (fab.SelectionService, error) {
	chnlCfg, err := c.GetChannelConfig(channelID)
	if err != nil {
		return nil, err
	}
	key := newCacheKey(chnlCfg)
	selectionService, err := c.selectionServiceCache.Get(key)
	if err != nil {
		return nil, err
	}
	ss := selectionService.(fab.SelectionService)
	ss.SetPeersOfChannel(c.peers)
	return ss, nil
}

// GetEventService returns the EventService.
func (c *contextCache) GetEventService(channelID string, opts ...options.Opt) (fab.EventService, error) {
	chnlCfg, err := c.GetChannelConfig(channelID)
	if err != nil {
		return nil, err
	}
	key, err := newEventCacheKey(chnlCfg, opts...)
	if err != nil {
		return nil, err
	}
	eventService, err := c.eventServiceCache.Get(key)
	if err != nil {
		return nil, err
	}
	es := eventService.(fab.EventService)
	es.SetPeersOfChannel(c.peers)
	return es, nil
}

func (c *contextCache) GetChannelConfig(channelID string) (fab.ChannelCfg, error) {
	if channelID == "" {
		// System channel
		return chconfig.NewChannelCfg(""), nil
	}
	chCfgRef, err := c.loadChannelCfgRef(channelID)
	if err != nil {
		return nil, err
	}
	chCfg, err := chCfgRef.Get()
	if err != nil {
		return nil, errors.WithMessage(err, "could not get chConfig cache reference")
	}
	return chCfg.(fab.ChannelCfg), nil
}

func (c *contextCache) loadChannelCfgRef(channelID string) (*chconfig.Ref, error) {
	key, err := chconfig.NewCacheKey(c.ctx, func(string) (fab.ChannelConfig, error) { return chconfig.New(channelID) }, channelID)
	if err != nil {
		return nil, err
	}
	cfg, err := c.chCfgCache.Get(key)
	if err != nil {
		return nil, err
	}
	chCfgRef := cfg.(*chconfig.Ref)
	chCfgRef.SetPeersOfChannel(c.peers)
	return chCfgRef, nil
}

func (c *contextCache) GetMembership(channelID string) (fab.ChannelMembership, error) {
	chCfgRef, err := c.loadChannelCfgRef(channelID)
	if err != nil {
		return nil, err
	}
	key, err := membership.NewCacheKey(membership.Context{Providers: c.ctx, EndpointConfig: c.ctx.EndpointConfig()},
		chCfgRef.Reference, channelID)
	if err != nil {
		return nil, err
	}
	ref, err := c.membershipCache.Get(key)
	if err != nil {
		return nil, err
	}

	return ref.(*membership.Ref), nil
}
