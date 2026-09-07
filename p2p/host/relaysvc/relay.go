package relaysvc

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	relayv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
)

type RelayManager struct {
	host host.Host

	mutex sync.Mutex
	relay *relayv2.Relay
	opts  []relayv2.Option

	refCount  sync.WaitGroup
	ctxCancel context.CancelFunc
}

func NewRelayManager(host host.Host, opts ...relayv2.Option) *RelayManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *RelayManager) background(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *RelayManager) reachabilityChanged(r network.Reachability) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *RelayManager) Close() error { _ = "STUB: not implemented"; return nil }
