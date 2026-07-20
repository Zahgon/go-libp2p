package sec

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type SecureConn interface {
	net.Conn
	network.ConnSecurity
}

type SecureTransport interface {
	SecureInbound(ctx context.Context, insecure net.Conn, p peer.ID) (SecureConn, error)

	SecureOutbound(ctx context.Context, insecure net.Conn, p peer.ID) (SecureConn, error)

	ID() protocol.ID
}

type ErrPeerIDMismatch struct {
	Expected peer.ID
	Actual   peer.ID
}

func (e ErrPeerIDMismatch) Error() string { _ = "STUB: not implemented"; return "" }

var _ error = (*ErrPeerIDMismatch)(nil)
