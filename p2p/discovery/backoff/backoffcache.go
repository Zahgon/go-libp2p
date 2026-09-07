package backoff

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/peer"
)

type BackoffDiscovery struct {
	disc         discovery.Discovery
	stratFactory BackoffFactory
	peerCache    map[string]*backoffCache
	peerCacheMux sync.RWMutex

	parallelBufSz int
	returnedBufSz int

	clock clock
}

type BackoffDiscoveryOption func(*BackoffDiscovery) error

func NewBackoffDiscovery(disc discovery.Discovery, stratFactory BackoffFactory, opts ...BackoffDiscoveryOption) (discovery.Discovery, error) {
	_ = "STUB: not implemented"
	return *new(discovery.Discovery), nil
}

func WithBackoffDiscoverySimultaneousQueryBufferSize(size int) BackoffDiscoveryOption {
	_ = "STUB: not implemented"
	return *new(BackoffDiscoveryOption)
}

func WithBackoffDiscoveryReturnedChannelSize(size int) BackoffDiscoveryOption {
	_ = "STUB: not implemented"
	return *new(BackoffDiscoveryOption)
}

type clock interface {
	Now() time.Time
}

type realClock struct{}

func (c realClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type backoffCache struct {
	strat BackoffStrategy

	mux          sync.Mutex
	nextDiscover time.Time
	prevPeers    map[peer.ID]peer.AddrInfo
	peers        map[peer.ID]peer.AddrInfo
	sendingChs   map[chan peer.AddrInfo]int
	ongoing      bool

	clock clock
}

func (d *BackoffDiscovery) Advertise(ctx context.Context, ns string, opts ...discovery.Option) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (d *BackoffDiscovery) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findPeerDispatcher(ctx context.Context, c *backoffCache, pch <-chan peer.AddrInfo) {
	_ = "STUB: not implemented"
	return
}

func findPeerReceiver(ctx context.Context, pch, evtCh chan peer.AddrInfo, rcvPeers []peer.AddrInfo) {
	_ = "STUB: not implemented"
	return
}

func mergeAddrInfos(prevAi, newAi peer.AddrInfo) *peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}

func checkUpdates(orig, update map[peer.ID]peer.AddrInfo) bool {
	_ = "STUB: not implemented"
	return false
}
