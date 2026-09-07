package mocknet

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

type peernet struct {
	mocknet *mocknet

	peer    peer.ID
	ps      peerstore.Peerstore
	emitter event.Emitter

	connsByPeer map[peer.ID]map[*conn]struct{}
	connsByLink map[*link]map[*conn]struct{}

	gater connmgr.ConnectionGater

	streamHandler network.StreamHandler

	notifmu sync.Mutex
	notifs  map[network.Notifiee]struct{}

	sync.RWMutex
}

func newPeernet(m *mocknet, p peer.ID, opts PeerOptions, bus event.Bus) (*peernet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pn *peernet) Close() error { _ = "STUB: not implemented"; return nil }

func (pn *peernet) allConns() []*conn { _ = "STUB: not implemented"; return nil }

func (pn *peernet) Peerstore() peerstore.Peerstore {
	_ = "STUB: not implemented"
	return *new(peerstore.Peerstore)
}

func (pn *peernet) String() string { _ = "STUB: not implemented"; return "" }

func (pn *peernet) handleNewStream(s network.Stream) { _ = "STUB: not implemented"; return }

func (pn *peernet) DialPeer(_ context.Context, p peer.ID) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (pn *peernet) connect(p peer.ID) (*conn, error) { _ = "STUB: not implemented"; return nil, nil }

func (pn *peernet) openConn(_ peer.ID, l *link) (*conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkSecureAndUpgrade(dir network.Direction, gater connmgr.ConnectionGater, c *conn) error {
	_ = "STUB: not implemented"
	return nil
}

func addConnPair(pn1, pn2 *peernet, c1, c2 *conn) { _ = "STUB: not implemented"; return }

func (pn *peernet) remoteOpenedConn(c *conn) { _ = "STUB: not implemented"; return }

func (pn *peernet) addConn(c *conn) { _ = "STUB: not implemented"; return }

func (pn *peernet) removeConn(c *conn) { _ = "STUB: not implemented"; return }

func (pn *peernet) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (pn *peernet) Peers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (pn *peernet) Conns() []network.Conn { _ = "STUB: not implemented"; return nil }

func (pn *peernet) ConnsToPeer(p peer.ID) []network.Conn { _ = "STUB: not implemented"; return nil }

func (pn *peernet) ClosePeer(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (pn *peernet) BandwidthTotals() (in uint64, out uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (pn *peernet) Listen(addrs ...ma.Multiaddr) error { _ = "STUB: not implemented"; return nil }

func (pn *peernet) ListenAddresses() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (pn *peernet) InterfaceListenAddresses() ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pn *peernet) Connectedness(p peer.ID) network.Connectedness {
	_ = "STUB: not implemented"
	return *new(network.Connectedness)
}

func (pn *peernet) NewStream(ctx context.Context, p peer.ID) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (pn *peernet) SetStreamHandler(h network.StreamHandler) { _ = "STUB: not implemented"; return }

func (pn *peernet) Notify(f network.Notifiee) { _ = "STUB: not implemented"; return }

func (pn *peernet) StopNotify(f network.Notifiee) { _ = "STUB: not implemented"; return }

func (pn *peernet) notifyAll(notification func(f network.Notifiee)) {
	_ = "STUB: not implemented"
	return
}

func (pn *peernet) ResourceManager() network.ResourceManager {
	_ = "STUB: not implemented"
	return *new(network.ResourceManager)
}

func (pn *peernet) CanDial(_ peer.ID, _ ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }
