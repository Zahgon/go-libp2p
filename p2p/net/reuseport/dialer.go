package reuseport

import (
	"context"
	"net"
)

type dialer struct {
	specific []*net.TCPAddr

	loopback []*net.TCPAddr

	unspecified []*net.TCPAddr
}

func (d *dialer) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func randAddr(addrs []*net.TCPAddr) *net.TCPAddr { _ = "STUB: not implemented"; return nil }

func (d *dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func newDialer(listeners map[*listener]struct{}) *dialer { _ = "STUB: not implemented"; return nil }
