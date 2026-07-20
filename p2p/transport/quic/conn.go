package libp2pquic

import (
	"context"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
)

type conn struct {
	quicConn  *quic.Conn
	transport *transport
	scope     network.ConnManagementScope

	localPeer      peer.ID
	localMultiaddr ma.Multiaddr

	remotePeerID    peer.ID
	remotePubKey    ic.PubKey
	remoteMultiaddr ma.Multiaddr
}

func (c *conn) As(target any) bool { _ = "STUB: not implemented"; return false }

var _ tpt.CapableConn = &conn{}

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) CloseWithError(errCode network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) closeWithError(errCode quic.ApplicationErrorCode, errString string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *conn) allowWindowIncrease(size uint64) bool { _ = "STUB: not implemented"; return false }

func (c *conn) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) AcceptStream() (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *conn) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (c *conn) RemotePublicKey() ic.PubKey { _ = "STUB: not implemented"; return *new(ic.PubKey) }

func (c *conn) LocalMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *conn) RemoteMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *conn) Transport() tpt.Transport { _ = "STUB: not implemented"; return *new(tpt.Transport) }

func (c *conn) Scope() network.ConnScope { _ = "STUB: not implemented"; return *new(network.ConnScope) }

func (c *conn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}
