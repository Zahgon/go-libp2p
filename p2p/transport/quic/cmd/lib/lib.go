package cmdlib

import (
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"
)

func RunClient(raddr string, p string) error { _ = "STUB: not implemented"; return nil }

func RunServer(port string, location chan peer.AddrInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func handleConn(conn tpt.CapableConn) error { _ = "STUB: not implemented"; return nil }
