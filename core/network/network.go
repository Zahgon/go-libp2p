package network

import (
	"context"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

const MessageSizeMax = 1 << 22

type Direction int

const (
	DirUnknown Direction = iota

	DirInbound

	DirOutbound
)

const unrecognized = "(unrecognized)"

func (d Direction) String() string { _ = "STUB: not implemented"; return "" }

type Connectedness int

const (
	NotConnected Connectedness = iota

	Connected

	CanConnect

	CannotConnect

	Limited
)

func (c Connectedness) String() string { _ = "STUB: not implemented"; return "" }

type Reachability int

const (
	ReachabilityUnknown Reachability = iota

	ReachabilityPublic

	ReachabilityPrivate
)

func (r Reachability) String() string { _ = "STUB: not implemented"; return "" }

type ConnStats struct {
	Stats

	NumStreams int
}

type Stats struct {
	Direction Direction

	Opened time.Time

	Limited bool

	Extra map[any]any
}

type StreamHandler func(Stream)

type Network interface {
	Dialer
	io.Closer

	SetStreamHandler(StreamHandler)

	NewStream(context.Context, peer.ID) (Stream, error)

	Listen(...ma.Multiaddr) error

	ListenAddresses() []ma.Multiaddr

	InterfaceListenAddresses() ([]ma.Multiaddr, error)

	ResourceManager() ResourceManager
}

type MultiaddrDNSResolver interface {
	ResolveDNSAddr(ctx context.Context, expectedPeerID peer.ID, maddr ma.Multiaddr, recursionLimit, outputLimit int) ([]ma.Multiaddr, error)

	ResolveDNSComponent(ctx context.Context, maddr ma.Multiaddr, outputLimit int) ([]ma.Multiaddr, error)
}

type Dialer interface {
	Peerstore() peerstore.Peerstore

	LocalPeer() peer.ID

	DialPeer(context.Context, peer.ID) (Conn, error)

	ClosePeer(peer.ID) error

	Connectedness(peer.ID) Connectedness

	Peers() []peer.ID

	Conns() []Conn

	ConnsToPeer(p peer.ID) []Conn

	Notify(Notifiee)
	StopNotify(Notifiee)

	CanDial(p peer.ID, addr ma.Multiaddr) bool
}

type AddrDelay struct {
	Addr  ma.Multiaddr
	Delay time.Duration
}

type DialRanker func([]ma.Multiaddr) []AddrDelay
