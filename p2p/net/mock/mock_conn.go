package mocknet

import (
	"container/list"
	"context"
	"sync"
	"sync/atomic"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

var connCounter atomic.Int64

type conn struct {
	notifLk sync.Mutex

	id int64

	local  peer.ID
	remote peer.ID

	localAddr  ma.Multiaddr
	remoteAddr ma.Multiaddr

	localPrivKey ic.PrivKey
	remotePubKey ic.PubKey

	net     *peernet
	link    *link
	rconn   *conn
	streams list.List
	stat    network.ConnStats

	closeOnce sync.Once

	isClosed atomic.Bool

	sync.RWMutex
}

func newConn(ln, rn *peernet, l *link, dir network.Direction) *conn {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *conn) ID() string { _ = "STUB: not implemented"; return "" }

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) As(_ any) bool { _ = "STUB: not implemented"; return false }

func (c *conn) teardown() { _ = "STUB: not implemented"; return }

func (c *conn) addStream(s *stream) { _ = "STUB: not implemented"; return }

func (c *conn) removeStream(s *stream) { _ = "STUB: not implemented"; return }

func (c *conn) allStreams() []network.Stream { _ = "STUB: not implemented"; return nil }

func (c *conn) remoteOpenedStream(s *stream) { _ = "STUB: not implemented"; return }

func (c *conn) openStream() *stream { _ = "STUB: not implemented"; return nil }

func (c *conn) NewStream(context.Context) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (c *conn) GetStreams() []network.Stream { _ = "STUB: not implemented"; return nil }

func (c *conn) LocalMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *conn) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemoteMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *conn) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemotePublicKey() ic.PubKey { _ = "STUB: not implemented"; return *new(ic.PubKey) }

func (c *conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (c *conn) Stat() network.ConnStats { _ = "STUB: not implemented"; return *new(network.ConnStats) }

func (c *conn) Scope() network.ConnScope { _ = "STUB: not implemented"; return *new(network.ConnScope) }

func (c *conn) CloseWithError(_ network.ConnErrorCode) error { _ = "STUB: not implemented"; return nil }
