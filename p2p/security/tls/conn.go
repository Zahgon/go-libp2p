package libp2ptls

import (
	"crypto/tls"

	ci "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/sec"
)

type conn struct {
	*tls.Conn

	localPeer       peer.ID
	remotePeer      peer.ID
	remotePubKey    ci.PubKey
	connectionState network.ConnectionState
}

var _ sec.SecureConn = &conn{}

func (c *conn) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemotePublicKey() ci.PubKey { _ = "STUB: not implemented"; return *new(ci.PubKey) }

func (c *conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}
