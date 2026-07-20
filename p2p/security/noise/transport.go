package noise

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/sec"
	tptu "github.com/libp2p/go-libp2p/p2p/net/upgrader"
	"github.com/libp2p/go-libp2p/p2p/security/noise/pb"
)

const ID = "/noise"
const maxProtoNum = 100

type Transport struct {
	protocolID protocol.ID
	localID    peer.ID
	privateKey crypto.PrivKey
	muxers     []protocol.ID
}

var _ sec.SecureTransport = &Transport{}

func New(id protocol.ID, privkey crypto.PrivKey, muxers []tptu.StreamMuxer) (*Transport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) SecureInbound(ctx context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) SecureOutbound(ctx context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) WithSessionOptions(opts ...SessionOption) (*SessionTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) ID() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }

func matchMuxers(initiatorMuxers, responderMuxers []protocol.ID) protocol.ID {
	_ = "STUB: not implemented"
	return *new(protocol.ID)
}

type transportEarlyDataHandler struct {
	transport      *Transport
	receivedMuxers []protocol.ID
}

var _ EarlyDataHandler = &transportEarlyDataHandler{}

func newTransportEDH(t *Transport) *transportEarlyDataHandler {
	_ = "STUB: not implemented"
	return nil
}

func (i *transportEarlyDataHandler) Send(context.Context, net.Conn, peer.ID) *pb.NoiseExtensions {
	_ = "STUB: not implemented"
	return nil
}

func (i *transportEarlyDataHandler) Received(_ context.Context, _ net.Conn, extension *pb.NoiseExtensions) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *transportEarlyDataHandler) MatchMuxers(isInitiator bool) protocol.ID {
	_ = "STUB: not implemented"
	return *new(protocol.ID)
}
