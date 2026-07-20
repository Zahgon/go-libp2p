package noise

import (
	"bufio"
	"context"
	"net"
	"sync"
	"time"

	"github.com/flynn/noise"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type secureSession struct {
	initiator   bool
	checkPeerID bool

	localID   peer.ID
	localKey  crypto.PrivKey
	remoteID  peer.ID
	remoteKey crypto.PubKey

	readLock  sync.Mutex
	writeLock sync.Mutex

	insecureConn   net.Conn
	insecureReader *bufio.Reader

	qseek int
	qbuf  []byte
	rlen  [2]byte

	enc *noise.CipherState
	dec *noise.CipherState

	prologue []byte

	initiatorEarlyDataHandler, responderEarlyDataHandler EarlyDataHandler

	connectionState network.ConnectionState
}

func newSecureSession(tpt *Transport, ctx context.Context, insecure net.Conn, remote peer.ID, prologue []byte, initiatorEDH, responderEDH EarlyDataHandler, initiator, checkPeerID bool) (*secureSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *secureSession) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s *secureSession) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (s *secureSession) LocalPublicKey() crypto.PubKey {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey)
}

func (s *secureSession) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s *secureSession) RemotePeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (s *secureSession) RemotePublicKey() crypto.PubKey {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey)
}

func (s *secureSession) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (s *secureSession) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *secureSession) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *secureSession) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *secureSession) Close() error { _ = "STUB: not implemented"; return nil }

func SessionWithConnState(s *secureSession, muxer protocol.ID) *secureSession {
	_ = "STUB: not implemented"
	return nil
}
