package host

import (
	"context"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"

	ma "github.com/multiformats/go-multiaddr"
)

type Host interface {
	ID() peer.ID

	Peerstore() peerstore.Peerstore

	Addrs() []ma.Multiaddr

	Network() network.Network

	Mux() protocol.Switch

	Connect(ctx context.Context, pi peer.AddrInfo) error

	SetStreamHandler(pid protocol.ID, handler network.StreamHandler)

	SetStreamHandlerMatch(protocol.ID, func(protocol.ID) bool, network.StreamHandler)

	RemoveStreamHandler(pid protocol.ID)

	NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error)

	Close() error

	ConnManager() connmgr.ConnManager

	EventBus() event.Bus
}
