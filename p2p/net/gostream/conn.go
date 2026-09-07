package gostream

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type conn struct {
	network.Stream
	ignoreEOF bool
}

func (c *conn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func newConn(s network.Stream, ignoreEOF bool) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

func (c *conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func Dial(ctx context.Context, h host.Host, pid peer.ID, tag protocol.ID) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
