//go:build windows || riscv64 || loong64

package tcp

import (
	"github.com/libp2p/go-libp2p/core/transport"
	manet "github.com/multiformats/go-multiaddr/net"
)

type aggregatingCollector struct{}

func newTracingConn(c manet.Conn, collector *aggregatingCollector, isClient bool) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func newTracingListener(l transport.GatedMaListener, collector *aggregatingCollector) transport.GatedMaListener {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener)
}
