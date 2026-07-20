package libp2pquic

import (
	"net"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
)

type listener struct {
	reuseListener   quicreuse.Listener
	transport       *transport
	rcmgr           network.ResourceManager
	privKey         ic.PrivKey
	localPeer       peer.ID
	localMultiaddrs map[quic.Version]ma.Multiaddr
}

func newListener(ln quicreuse.Listener, t *transport, localPeer peer.ID, key ic.PrivKey, rcmgr network.ResourceManager) (listener, error) {
	_ = "STUB: not implemented"
	return *new(listener), nil
}

func (l *listener) Accept() (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (l *listener) wrapConn(qconn *quic.Conn) (*conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) wrapConnWithScope(qconn *quic.Conn, connScope network.ConnManagementScope, remoteMultiaddr ma.Multiaddr) (*conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
