package routing

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"

	"github.com/ipfs/go-cid"
)

type RoutingDiscovery struct {
	routing.ContentRouting
}

func NewRoutingDiscovery(router routing.ContentRouting) *RoutingDiscovery {
	_ = "STUB: not implemented"
	return nil
}

func (d *RoutingDiscovery) Advertise(ctx context.Context, ns string, opts ...discovery.Option) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (d *RoutingDiscovery) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nsToCid(ns string) (cid.Cid, error) { _ = "STUB: not implemented"; return *new(cid.Cid), nil }

func NewDiscoveryRouting(disc discovery.Discovery, opts ...discovery.Option) *DiscoveryRouting {
	_ = "STUB: not implemented"
	return nil
}

type DiscoveryRouting struct {
	discovery.Discovery
	opts []discovery.Option
}

func (r *DiscoveryRouting) Provide(ctx context.Context, c cid.Cid, bcast bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiscoveryRouting) FindProvidersAsync(ctx context.Context, c cid.Cid, limit int) <-chan peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}

func cidToNs(c cid.Cid) string { _ = "STUB: not implemented"; return "" }
