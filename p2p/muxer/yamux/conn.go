package yamux

import (
	"context"

	"github.com/libp2p/go-libp2p/core/network"

	"github.com/libp2p/go-yamux/v5"
)

type conn yamux.Session

var _ network.MuxedConn = &conn{}

func (c *conn) As(target any) bool { _ = "STUB: not implemented"; return false }

func NewMuxedConn(m *yamux.Session) network.MuxedConn {
	_ = "STUB: not implemented"
	return *new(network.MuxedConn)
}

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) CloseWithError(errCode network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *conn) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) AcceptStream() (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) yamux() *yamux.Session { _ = "STUB: not implemented"; return nil }
