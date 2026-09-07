package libp2ptls

import (
	"context"
	"crypto/tls"
	"net"

	ci "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/sec"
	tptu "github.com/libp2p/go-libp2p/p2p/net/upgrader"
)

const ID = "/tls/1.0.0"

type Transport struct {
	identity *Identity

	localPeer  peer.ID
	privKey    ci.PrivKey
	muxers     []protocol.ID
	protocolID protocol.ID
}

var _ sec.SecureTransport = &Transport{}

func New(id protocol.ID, key ci.PrivKey, muxers []tptu.StreamMuxer) (*Transport, error) {
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

func (t *Transport) handshake(ctx context.Context, tlsConn *tls.Conn, keyCh <-chan ci.PubKey) (_sconn sec.SecureConn, err error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) setupConn(tlsConn *tls.Conn, remotePubKey ci.PubKey) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) ID() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }
