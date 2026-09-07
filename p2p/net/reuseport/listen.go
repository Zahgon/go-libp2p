package reuseport

import (
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

type listener struct {
	manet.Listener
	network *network
}

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (t *Transport) Listen(laddr ma.Multiaddr) (manet.Listener, error) {
	_ = "STUB: not implemented"
	return *new(manet.Listener), nil
}
