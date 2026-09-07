package autorelay

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("autorelay")

type AutoRelay struct {
	refCount  sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc

	mx     sync.Mutex
	status network.Reachability

	relayFinder *relayFinder

	host host.Host

	metricsTracer MetricsTracer
}

func NewAutoRelay(host host.Host, opts ...Option) (*AutoRelay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *AutoRelay) IsPeerInBackoff(peerID peer.ID) bool { _ = "STUB: not implemented"; return false }

func (r *AutoRelay) Start() { _ = "STUB: not implemented"; return }

func (r *AutoRelay) background() { _ = "STUB: not implemented"; return }

func (r *AutoRelay) Close() error { _ = "STUB: not implemented"; return nil }
