package tlsdiag

import (
	"net"

	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
)

func StartServer() error { _ = "STUB: not implemented"; return nil }

func handleConn(tp *libp2ptls.Transport, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}
