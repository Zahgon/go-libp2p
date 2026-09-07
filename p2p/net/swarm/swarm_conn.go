package swarm

import (
	"context"
	"errors"
	"sync"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

var ErrConnClosed = errors.New("connection closed")

type Conn struct {
	id    uint64
	conn  transport.CapableConn
	swarm *Swarm

	closeOnce sync.Once
	err       error

	streams struct {
		sync.Mutex
		m map[*Stream]struct{}
	}

	stat network.ConnStats
}

var _ network.Conn = &Conn{}

func (c *Conn) As(target any) bool { _ = "STUB: not implemented"; return false }

func (c *Conn) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *Conn) ID() string { _ = "STUB: not implemented"; return "" }

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) CloseWithError(errCode network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) doClose(errCode network.ConnErrorCode) { _ = "STUB: not implemented"; return }

func (c *Conn) removeStream(s *Stream) { _ = "STUB: not implemented"; return }

func (c *Conn) start() { _ = "STUB: not implemented"; return }

func (c *Conn) String() string { _ = "STUB: not implemented"; return "" }

func (c *Conn) LocalMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *Conn) RemoteMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *Conn) RemotePublicKey() ic.PubKey { _ = "STUB: not implemented"; return *new(ic.PubKey) }

func (c *Conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (c *Conn) Stat() network.ConnStats { _ = "STUB: not implemented"; return *new(network.ConnStats) }

func (c *Conn) NewStream(ctx context.Context) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (c *Conn) openAndAddStream(ctx context.Context, scope network.StreamManagementScope) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (c *Conn) addStream(ts network.MuxedStream, dir network.Direction, scope network.StreamManagementScope) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) GetStreams() []network.Stream { _ = "STUB: not implemented"; return nil }

func (c *Conn) Scope() network.ConnScope { _ = "STUB: not implemented"; return *new(network.ConnScope) }
