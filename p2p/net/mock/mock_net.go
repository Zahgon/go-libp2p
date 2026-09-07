package mocknet

import (
	"context"
	"net"
	"sync"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

var blackholeIP6 = net.ParseIP("100::")

type mocknet struct {
	nets  map[peer.ID]*peernet
	hosts map[peer.ID]host.Host

	links map[peer.ID]map[peer.ID]map[*link]struct{}

	linkDefaults LinkOptions

	ctxCancel context.CancelFunc
	ctx       context.Context
	sync.Mutex
}

func New() Mocknet { _ = "STUB: not implemented"; return *new(Mocknet) }

func (mn *mocknet) Close() error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) GenPeer() (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (mn *mocknet) GenPeerWithOptions(opts PeerOptions) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (mn *mocknet) AddPeer(k ic.PrivKey, a ma.Multiaddr) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (mn *mocknet) AddPeerWithPeerstore(p peer.ID, ps peerstore.Peerstore) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (mn *mocknet) AddPeerWithOptions(p peer.ID, opts PeerOptions) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (mn *mocknet) addDefaults(opts *PeerOptions) error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) updatePeerstore(k ic.PrivKey, a ma.Multiaddr, ps peerstore.Peerstore) (peer.ID, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil
}

func (mn *mocknet) Peers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) Host(pid peer.ID) host.Host { _ = "STUB: not implemented"; return *new(host.Host) }

func (mn *mocknet) Net(pid peer.ID) network.Network {
	_ = "STUB: not implemented"
	return *new(network.Network)
}

func (mn *mocknet) Hosts() []host.Host { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) Nets() []network.Network { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) Links() LinkMap { _ = "STUB: not implemented"; return *new(LinkMap) }

func (mn *mocknet) LinkAll() error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) LinkPeers(p1, p2 peer.ID) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (mn *mocknet) validate(n network.Network) (*peernet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mn *mocknet) LinkNets(n1, n2 network.Network) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (mn *mocknet) Unlink(l2 Link) error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) UnlinkPeers(p1, p2 peer.ID) error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) UnlinkNets(n1, n2 network.Network) error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) linksMapGet(p1, p2 peer.ID) map[*link]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (mn *mocknet) addLink(l *link) { _ = "STUB: not implemented"; return }

func (mn *mocknet) removeLink(l *link) { _ = "STUB: not implemented"; return }

func (mn *mocknet) ConnectAllButSelf() error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) ConnectPeers(a, b peer.ID) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (mn *mocknet) ConnectNets(a, b network.Network) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (mn *mocknet) DisconnectPeers(p1, p2 peer.ID) error { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) DisconnectNets(n1, n2 network.Network) error {
	_ = "STUB: not implemented"
	return nil
}

func (mn *mocknet) LinksBetweenPeers(p1, p2 peer.ID) []Link { _ = "STUB: not implemented"; return nil }

func (mn *mocknet) LinksBetweenNets(n1, n2 network.Network) []Link {
	_ = "STUB: not implemented"
	return nil
}

func (mn *mocknet) SetLinkDefaults(o LinkOptions) { _ = "STUB: not implemented"; return }

func (mn *mocknet) LinkDefaults() LinkOptions { _ = "STUB: not implemented"; return *new(LinkOptions) }

type netSlice []network.Network

func (es netSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (es netSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (es netSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type hostSlice []host.Host

func (es hostSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (es hostSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (es hostSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
