package mocknet

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

type link struct {
	mock        *mocknet
	nets        []*peernet
	opts        LinkOptions
	ratelimiter *RateLimiter

	sync.RWMutex
}

func newLink(mn *mocknet, opts LinkOptions) *link { _ = "STUB: not implemented"; return nil }

func (l *link) newConnPair(dialer *peernet) (*conn, *conn) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *link) Networks() []network.Network { _ = "STUB: not implemented"; return nil }

func (l *link) Peers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (l *link) SetOptions(o LinkOptions) { _ = "STUB: not implemented"; return }

func (l *link) Options() LinkOptions { _ = "STUB: not implemented"; return *new(LinkOptions) }

func (l *link) GetLatency() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (l *link) RateLimit(dataSize int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
