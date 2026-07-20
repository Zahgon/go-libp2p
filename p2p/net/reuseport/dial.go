package reuseport

import (
	"context"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func (t *Transport) Dial(raddr ma.Multiaddr) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (t *Transport) DialContext(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (n *network) getDialer(_ string) *dialer { _ = "STUB: not implemented"; return nil }
