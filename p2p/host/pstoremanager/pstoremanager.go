package pstoremanager

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peerstore"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("pstoremanager")

type Option func(*PeerstoreManager) error

func WithGracePeriod(p time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCleanupInterval(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type PeerstoreManager struct {
	pstore   peerstore.Peerstore
	eventBus event.Bus
	network  network.Network

	cancel   context.CancelFunc
	refCount sync.WaitGroup

	gracePeriod     time.Duration
	cleanupInterval time.Duration
}

func NewPeerstoreManager(pstore peerstore.Peerstore, eventBus event.Bus, network network.Network, opts ...Option) (*PeerstoreManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PeerstoreManager) Start() { _ = "STUB: not implemented"; return }

func (m *PeerstoreManager) background(ctx context.Context, sub event.Subscription) {
	_ = "STUB: not implemented"
	return
}

func (m *PeerstoreManager) Close() error { _ = "STUB: not implemented"; return nil }
