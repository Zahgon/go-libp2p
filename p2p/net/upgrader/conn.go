package upgrader

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/transport"
)

type transportConn struct {
	network.MuxedConn
	network.ConnMultiaddrs
	network.ConnSecurity
	transport transport.Transport
	scope     network.ConnManagementScope
	stat      network.ConnStats

	muxer                     protocol.ID
	security                  protocol.ID
	usedEarlyMuxerNegotiation bool
}

var _ transport.CapableConn = &transportConn{}

func (c *transportConn) As(target any) bool { _ = "STUB: not implemented"; return false }

func (t *transportConn) Transport() transport.Transport {
	_ = "STUB: not implemented"
	return *new(transport.Transport)
}

func (t *transportConn) String() string { _ = "STUB: not implemented"; return "" }

func (t *transportConn) Stat() network.ConnStats {
	_ = "STUB: not implemented"
	return *new(network.ConnStats)
}

func (t *transportConn) Scope() network.ConnScope {
	_ = "STUB: not implemented"
	return *new(network.ConnScope)
}

func (t *transportConn) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transportConn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (t *transportConn) CloseWithError(errCode network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}
