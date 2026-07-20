package insecure

import (
	"context"
	"io"
	"net"

	ci "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/sec"
	"github.com/libp2p/go-libp2p/p2p/security/insecure/pb"
)

const ID = "/plaintext/2.0.0"

type Transport struct {
	id         peer.ID
	key        ci.PrivKey
	protocolID protocol.ID
}

var _ sec.SecureTransport = &Transport{}

func NewWithIdentity(protocolID protocol.ID, id peer.ID, key ci.PrivKey) *Transport {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (t *Transport) SecureInbound(_ context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) SecureOutbound(_ context.Context, insecure net.Conn, p peer.ID) (sec.SecureConn, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), nil
}

func (t *Transport) ID() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }

type Conn struct {
	net.Conn

	local, remote             peer.ID
	localPubKey, remotePubKey ci.PubKey
}

func makeExchangeMessage(pubkey ci.PubKey) (*pb.Exchange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ic *Conn) runHandshakeSync() error { _ = "STUB: not implemented"; return nil }

func readWriteMsg(rw io.ReadWriter, out *pb.Exchange) (*pb.Exchange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ic *Conn) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (ic *Conn) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (ic *Conn) RemotePublicKey() ci.PubKey { _ = "STUB: not implemented"; return *new(ci.PubKey) }

func (ic *Conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

var _ sec.SecureTransport = (*Transport)(nil)
var _ sec.SecureConn = (*Conn)(nil)
