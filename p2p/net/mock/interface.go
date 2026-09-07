package mocknet

import (
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

type PeerOptions struct {
	ps peerstore.Peerstore

	gater connmgr.ConnectionGater
}

type Mocknet interface {
	GenPeer() (host.Host, error)
	GenPeerWithOptions(PeerOptions) (host.Host, error)

	AddPeer(ic.PrivKey, ma.Multiaddr) (host.Host, error)
	AddPeerWithPeerstore(peer.ID, peerstore.Peerstore) (host.Host, error)
	AddPeerWithOptions(peer.ID, PeerOptions) (host.Host, error)

	Peers() []peer.ID
	Net(peer.ID) network.Network
	Nets() []network.Network
	Host(peer.ID) host.Host
	Hosts() []host.Host
	Links() LinkMap
	LinksBetweenPeers(a, b peer.ID) []Link
	LinksBetweenNets(a, b network.Network) []Link

	LinkPeers(peer.ID, peer.ID) (Link, error)
	LinkNets(network.Network, network.Network) (Link, error)
	Unlink(Link) error
	UnlinkPeers(peer.ID, peer.ID) error
	UnlinkNets(network.Network, network.Network) error

	SetLinkDefaults(LinkOptions)
	LinkDefaults() LinkOptions

	ConnectPeers(peer.ID, peer.ID) (network.Conn, error)
	ConnectNets(network.Network, network.Network) (network.Conn, error)
	DisconnectPeers(peer.ID, peer.ID) error
	DisconnectNets(network.Network, network.Network) error
	LinkAll() error
	ConnectAllButSelf() error

	io.Closer
}

type LinkOptions struct {
	Latency   time.Duration
	Bandwidth float64
}

type Link interface {
	Networks() []network.Network
	Peers() []peer.ID

	SetOptions(LinkOptions)
	Options() LinkOptions
}

type LinkMap map[string]map[string]map[Link]struct{}

type Printer interface {
	MocknetLinks(mn Mocknet)
	NetworkConns(ni network.Network)
}

func PrinterTo(w io.Writer) Printer { _ = "STUB: not implemented"; return *new(Printer) }
