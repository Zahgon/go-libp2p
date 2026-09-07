package libp2pwebtransport

import (
	"context"

	"github.com/libp2p/go-libp2p/core/network"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/webtransport-go"
)

type connSecurityMultiaddrs struct {
	network.ConnSecurity
	network.ConnMultiaddrs
}

type connMultiaddrs struct {
	local, remote ma.Multiaddr
}

var _ network.ConnMultiaddrs = &connMultiaddrs{}

func (c *connMultiaddrs) LocalMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}
func (c *connMultiaddrs) RemoteMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

type conn struct {
	*connSecurityMultiaddrs

	transport *transport
	session   *webtransport.Session

	scope network.ConnManagementScope
	qconn *quic.Conn
}

var _ tpt.CapableConn = &conn{}

func newConn(tr *transport, sess *webtransport.Session, sconn *connSecurityMultiaddrs, scope network.ConnManagementScope, qconn *quic.Conn) *conn {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) AcceptStream() (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) allowWindowIncrease(size uint64) bool { _ = "STUB: not implemented"; return false }

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) CloseWithError(_ network.ConnErrorCode) error { _ = "STUB: not implemented"; return nil }

func (c *conn) IsClosed() bool           { _ = "STUB: not implemented"; return false }
func (c *conn) Scope() network.ConnScope { _ = "STUB: not implemented"; return *new(network.ConnScope) }
func (c *conn) Transport() tpt.Transport { _ = "STUB: not implemented"; return *new(tpt.Transport) }

func (c *conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (c *conn) As(target any) bool { _ = "STUB: not implemented"; return false }
