package tcpreuse

import (
	"context"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func (t *ConnMgr) DialContext(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}
