package websocket

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"

	ws "github.com/gorilla/websocket"
)

var GracefulCloseTimeout = 100 * time.Millisecond

type Conn struct {
	*ws.Conn
	Scope              network.ConnManagementScope
	secure             bool
	DefaultMessageType int
	reader             io.Reader
	closeOnceVal       func() error
	laddr              ma.Multiaddr
	raddr              ma.Multiaddr

	readLock, writeLock sync.Mutex
}

var _ net.Conn = (*Conn)(nil)
var _ manet.Conn = (*Conn)(nil)

func newConn(raw *ws.Conn, secure bool, scope network.ConnManagementScope) *Conn {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) LocalMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) RemoteMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) prepNextReader() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) closeOnceFn() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
