package gostream

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type listener struct {
	host     host.Host
	ctx      context.Context
	tag      protocol.ID
	cancel   func()
	streamCh chan network.Stream

	ignoreEOF bool
}

func (l *listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func Listen(h host.Host, tag protocol.ID, opts ...ListenerOption) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

type ListenerOption func(*listener) error

func IgnoreEOF() ListenerOption { _ = "STUB: not implemented"; return *new(ListenerOption) }
