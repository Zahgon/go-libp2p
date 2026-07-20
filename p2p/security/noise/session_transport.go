package noise

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/sec"
	"github.com/libp2p/go-libp2p/p2p/security/noise/pb"
)

type SessionOption = func(*SessionTransport) error

func Prologue(prologue []byte) SessionOption { _ = "STUB: not implemented"; return *new(SessionOption) }

type EarlyDataHandler interface {
	Send(context.Context, net.Conn, peer.ID) *pb.NoiseExtensions

	Received(context.Context, net.Conn, *pb.NoiseExtensions) error
}

func EarlyData(initiator, responder EarlyDataHandler) SessionOption {
	_ = "STUB: not implemented"
	return *new(SessionOption)
}

func DisablePeerIDCheck() SessionOption { _ = "STUB: not implemented"; return *new(SessionOption) }

var _ sec.SecureTransport = &SessionTransport{}

type SessionTransport struct {
	t *Transport

	prologue           []byte
	disablePeerIDCheck bool

	protocolID protocol.ID

	initiatorEarlyDataHandler, responderEarlyDataHandler EarlyDataHandler
}

func (i *SessionTransport) SecureInbound(ctx context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (i *SessionTransport) SecureOutbound(ctx context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (i *SessionTransport) ID() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }
