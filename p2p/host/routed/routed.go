package routedhost

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"

	logging "github.com/libp2p/go-libp2p/gologshim"

	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("routedhost")

const AddressTTL = time.Second * 10

type RoutedHost struct {
	host  host.Host
	route Routing
}

type Routing interface {
	FindPeer(context.Context, peer.ID) (peer.AddrInfo, error)
}

func Wrap(h host.Host, r Routing) *RoutedHost { _ = "STUB: not implemented"; return nil }

func (rh *RoutedHost) Connect(ctx context.Context, pi peer.AddrInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (rh *RoutedHost) findPeerAddrs(ctx context.Context, id peer.ID) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh *RoutedHost) ID() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (rh *RoutedHost) Peerstore() peerstore.Peerstore {
	_ = "STUB: not implemented"
	return *new(peerstore.Peerstore)
}

func (rh *RoutedHost) Addrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (rh *RoutedHost) Network() network.Network {
	_ = "STUB: not implemented"
	return *new(network.Network)
}

func (rh *RoutedHost) Mux() protocol.Switch {
	_ = "STUB: not implemented"
	return *new(protocol.Switch)
}

func (rh *RoutedHost) EventBus() event.Bus { _ = "STUB: not implemented"; return *new(event.Bus) }

func (rh *RoutedHost) SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (rh *RoutedHost) SetStreamHandlerMatch(pid protocol.ID, m func(protocol.ID) bool, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (rh *RoutedHost) RemoveStreamHandler(pid protocol.ID) { _ = "STUB: not implemented"; return }

func (rh *RoutedHost) NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (rh *RoutedHost) Close() error { _ = "STUB: not implemented"; return nil }

func (rh *RoutedHost) ConnManager() connmgr.ConnManager {
	_ = "STUB: not implemented"
	return *new(connmgr.ConnManager)
}

var _ (host.Host) = (*RoutedHost)(nil)
