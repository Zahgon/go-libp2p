package quicreuse

import (
	"context"
	"net"
	"time"
)

type nonQUICPacketConn struct {
	owningTransport RefCountedQUICTransport
	tr              QUICTransport
	ctx             context.Context
	ctxCancel       context.CancelFunc
	readCtx         context.Context
	readCancel      context.CancelFunc
}

func (n *nonQUICPacketConn) Close() error { _ = "STUB: not implemented"; return nil }

func (n *nonQUICPacketConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (n *nonQUICPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (n *nonQUICPacketConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (n *nonQUICPacketConn) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *nonQUICPacketConn) SetWriteDeadline(_ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *nonQUICPacketConn) WriteTo(p []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ net.PacketConn = &nonQUICPacketConn{}
