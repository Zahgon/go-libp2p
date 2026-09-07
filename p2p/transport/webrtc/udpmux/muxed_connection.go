package udpmux

import (
	"context"
	"net"
	"time"
)

type packet struct {
	buf  []byte
	addr net.Addr
}

var _ net.PacketConn = &muxedConnection{}

const queueLen = 128

type muxedConnection struct {
	ctx        context.Context
	cancel     context.CancelFunc
	queue      chan packet
	mux        *UDPMux
	localUfrag string
}

var _ net.PacketConn = &muxedConnection{}

func newMuxedConnection(mux *UDPMux, localUfrag string) *muxedConnection {
	_ = "STUB: not implemented"
	return nil
}

func (c *muxedConnection) Push(buf []byte, addr net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *muxedConnection) ReadFrom(buf []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *muxedConnection) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *muxedConnection) Close() error { _ = "STUB: not implemented"; return nil }

func (c *muxedConnection) close() { _ = "STUB: not implemented"; return }

func (c *muxedConnection) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (*muxedConnection) SetDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*muxedConnection) SetReadDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*muxedConnection) SetWriteDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }
