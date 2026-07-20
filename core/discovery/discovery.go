package discovery

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

type Advertiser interface {
	Advertise(ctx context.Context, ns string, opts ...Option) (time.Duration, error)
}

type Discoverer interface {
	FindPeers(ctx context.Context, ns string, opts ...Option) (<-chan peer.AddrInfo, error)
}

type Discovery interface {
	Advertiser
	Discoverer
}
