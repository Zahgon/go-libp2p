package tcpreuse

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/transport/tcpreuse/internal/sampledconn"
	manet "github.com/multiformats/go-multiaddr/net"
)

type connWithScope struct {
	sampledconn.ManetTCPConnInterface
	ConnScope network.ConnManagementScope
}

func (c *connWithScope) Close() error { _ = "STUB: not implemented"; return nil }

func manetConnWithScope(c manet.Conn, scope network.ConnManagementScope) (*connWithScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
