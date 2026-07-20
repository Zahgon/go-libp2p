package holepunch

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"
	ma "github.com/multiformats/go-multiaddr"
)

var ErrHolePunchActive = errors.New("another hole punching attempt to this peer is active")

const maxRetries = 3

type holePuncher struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	host     host.Host
	refCount sync.WaitGroup

	ids         identify.IDService
	listenAddrs func() []ma.Multiaddr

	directDialTimeout time.Duration

	activeMx sync.Mutex
	active   map[peer.ID]struct{}

	closeMx sync.RWMutex
	closed  bool

	tracer *tracer
	filter AddrFilter
}

func newHolePuncher(h host.Host, ids identify.IDService, listenAddrs func() []ma.Multiaddr, tracer *tracer, filter AddrFilter) *holePuncher {
	_ = "STUB: not implemented"
	return nil
}

func (hp *holePuncher) beginDirectConnect(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (hp *holePuncher) DirectConnect(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (hp *holePuncher) directConnect(rp peer.ID) error { _ = "STUB: not implemented"; return nil }

func (hp *holePuncher) initiateHolePunch(rp peer.ID) ([]ma.Multiaddr, []ma.Multiaddr, time.Duration, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(time.Duration), nil
}

func (hp *holePuncher) initiateHolePunchImpl(str network.Stream) ([]ma.Multiaddr, []ma.Multiaddr, time.Duration, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(time.Duration), nil
}

func (hp *holePuncher) Close() error { _ = "STUB: not implemented"; return nil }

type netNotifiee holePuncher

func (nn *netNotifiee) Connected(_ network.Network, conn network.Conn) {
	_ = "STUB: not implemented"
	return
}

func (nn *netNotifiee) Disconnected(_ network.Network, _ network.Conn) {
	_ = "STUB: not implemented"
	return
}
func (nn *netNotifiee) Listen(_ network.Network, _ ma.Multiaddr) { _ = "STUB: not implemented"; return }
func (nn *netNotifiee) ListenClose(_ network.Network, _ ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}
