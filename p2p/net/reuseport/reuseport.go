package reuseport

import (
	"context"
	"net"
)

var fallbackDialer net.Dialer

func reuseDial(ctx context.Context, laddr *net.TCPAddr, network, raddr string) (con net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
