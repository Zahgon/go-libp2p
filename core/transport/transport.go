package transport

import (
	"context"
	"errors"
	"net"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

type CapableConn interface {
	network.MuxedConn
	network.ConnSecurity
	network.ConnMultiaddrs
	network.ConnScoper

	Transport() Transport
}

type Transport interface {
	Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (CapableConn, error)

	CanDial(addr ma.Multiaddr) bool

	Listen(laddr ma.Multiaddr) (Listener, error)

	Protocols() []int

	Proxy() bool
}

type Resolver interface {
	Resolve(ctx context.Context, maddr ma.Multiaddr) ([]ma.Multiaddr, error)
}

type SkipResolver interface {
	SkipResolve(ctx context.Context, maddr ma.Multiaddr) bool
}

type Listener interface {
	Accept() (CapableConn, error)
	Close() error
	Addr() net.Addr
	Multiaddr() ma.Multiaddr
}

var ErrListenerClosed = errors.New("listener closed")

type TransportNetwork interface {
	network.Network

	AddTransport(t Transport) error
}

type GatedMaListener interface {
	Accept() (manet.Conn, network.ConnManagementScope, error)

	Close() error

	Multiaddr() ma.Multiaddr

	Addr() net.Addr
}

type Upgrader interface {
	UpgradeListener(Transport, manet.Listener) Listener

	GateMaListener(manet.Listener) GatedMaListener

	UpgradeGatedMaListener(Transport, GatedMaListener) Listener

	Upgrade(ctx context.Context, t Transport, maconn manet.Conn, dir network.Direction, p peer.ID, scope network.ConnManagementScope) (CapableConn, error)
}

type DialUpdater interface {
	DialWithUpdates(context.Context, ma.Multiaddr, peer.ID, chan<- DialUpdate) (CapableConn, error)
}

type DialUpdateKind int

const (
	UpdateKindDialFailed DialUpdateKind = iota

	UpdateKindDialSuccessful

	UpdateKindHandshakeProgressed
)

func (k DialUpdateKind) String() string { _ = "STUB: not implemented"; return "" }

type DialUpdate struct {
	Kind DialUpdateKind

	Addr ma.Multiaddr

	Conn CapableConn

	Err error
}
