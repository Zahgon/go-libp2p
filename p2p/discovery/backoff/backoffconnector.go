package backoff

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	lru "github.com/hashicorp/golang-lru/v2"
)

type BackoffConnector struct {
	cache      *lru.TwoQueueCache[peer.ID, *connCacheData]
	host       host.Host
	connTryDur time.Duration
	backoff    BackoffFactory
	mux        sync.Mutex
}

func NewBackoffConnector(h host.Host, cacheSize int, connectionTryDuration time.Duration, backoff BackoffFactory) (*BackoffConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type connCacheData struct {
	nextTry time.Time
	strat   BackoffStrategy
}

func (c *BackoffConnector) Connect(ctx context.Context, peerCh <-chan peer.AddrInfo) {
	_ = "STUB: not implemented"
	return
}
