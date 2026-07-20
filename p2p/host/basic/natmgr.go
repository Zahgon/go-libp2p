package basichost

import (
	"context"
	"io"
	"net/netip"
	"sync"

	"github.com/libp2p/go-libp2p/core/network"
	inat "github.com/libp2p/go-libp2p/p2p/net/nat"

	ma "github.com/multiformats/go-multiaddr"
)

type NATManager interface {
	GetMapping(ma.Multiaddr) ma.Multiaddr
	HasDiscoveredNAT() bool
	io.Closer
}

func NewNATManager(net network.Network) NATManager {
	_ = "STUB: not implemented"
	return *new(NATManager)
}

type entry struct {
	protocol string
	port     int
}

type nat interface {
	AddMapping(ctx context.Context, protocol string, port int) error
	RemoveMapping(ctx context.Context, protocol string, port int) error
	GetMapping(protocol string, port int) (netip.AddrPort, bool)
	io.Closer
}

var discoverNAT = func(ctx context.Context) (nat, error) { return inat.DiscoverNAT(ctx) }

type natManager struct {
	net   network.Network
	natMx sync.RWMutex
	nat   nat

	syncFlag chan struct{}

	tracked map[entry]bool

	refCount  sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func newNATManager(net network.Network) *natManager { _ = "STUB: not implemented"; return nil }

func (nmgr *natManager) Close() error { _ = "STUB: not implemented"; return nil }

func (nmgr *natManager) HasDiscoveredNAT() bool { _ = "STUB: not implemented"; return false }

func (nmgr *natManager) background(ctx context.Context) { _ = "STUB: not implemented"; return }

func (nmgr *natManager) sync() { _ = "STUB: not implemented"; return }

func (nmgr *natManager) doSync() { _ = "STUB: not implemented"; return }

func (nmgr *natManager) GetMapping(addr ma.Multiaddr) ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

type nmgrNetNotifiee natManager

func (nn *nmgrNetNotifiee) natManager() *natManager              { _ = "STUB: not implemented"; return nil }
func (nn *nmgrNetNotifiee) Listen(network.Network, ma.Multiaddr) { _ = "STUB: not implemented"; return }
func (nn *nmgrNetNotifiee) ListenClose(_ network.Network, _ ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}
func (nn *nmgrNetNotifiee) Connected(network.Network, network.Conn) {
	_ = "STUB: not implemented"
	return
}
func (nn *nmgrNetNotifiee) Disconnected(network.Network, network.Conn) {
	_ = "STUB: not implemented"
	return
}
