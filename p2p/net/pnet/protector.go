package pnet

import (
	"net"

	ipnet "github.com/libp2p/go-libp2p/core/pnet"
)

func NewProtectedConn(psk ipnet.PSK, conn net.Conn) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
