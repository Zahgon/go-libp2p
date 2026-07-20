package client

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var _ manet.Listener = (*Listener)(nil)

type Listener Client

func (c *Client) Listener() *Listener { _ = "STUB: not implemented"; return nil }

func (l *Listener) Accept() (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (l *Listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *Listener) Multiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (l *Listener) Close() error { _ = "STUB: not implemented"; return nil }
