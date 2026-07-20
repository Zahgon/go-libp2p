package libp2pwebrtc

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/pion/datachannel"
	"github.com/pion/webrtc/v4"
)

var _ tpt.CapableConn = &connection{}

const maxAcceptQueueLen = 256

type errConnectionTimeout struct{}

var _ net.Error = &errConnectionTimeout{}

func (errConnectionTimeout) Error() string   { _ = "STUB: not implemented"; return "" }
func (errConnectionTimeout) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (errConnectionTimeout) Temporary() bool { _ = "STUB: not implemented"; return false }

var errConnClosed = errors.New("connection closed")

type dataChannel struct {
	stream  datachannel.ReadWriteCloser
	channel *webrtc.DataChannel
}

type connection struct {
	pc        *webrtc.PeerConnection
	transport *WebRTCTransport
	scope     network.ConnManagementScope

	closeOnce sync.Once
	closeErr  error

	localPeer      peer.ID
	localMultiaddr ma.Multiaddr

	remotePeer      peer.ID
	remoteKey       ic.PubKey
	remoteMultiaddr ma.Multiaddr

	m            sync.Mutex
	streams      map[uint16]*stream
	nextStreamID atomic.Int32

	acceptQueue chan dataChannel

	ctx    context.Context
	cancel context.CancelFunc
}

func newConnection(
	direction network.Direction,
	pc *webrtc.PeerConnection,
	transport *WebRTCTransport,
	scope network.ConnManagementScope,

	localPeer peer.ID,
	localMultiaddr ma.Multiaddr,

	remotePeer peer.ID,
	remoteKey ic.PubKey,
	remoteMultiaddr ma.Multiaddr,
	incomingDataChannels chan dataChannel,
	peerConnectionClosedCh chan struct{},
) (*connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *connection) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (c *connection) Close() error { _ = "STUB: not implemented"; return nil }

func (c *connection) As(target any) bool { _ = "STUB: not implemented"; return false }

func (c *connection) CloseWithError(_ network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *connection) closeWithError(err error) { _ = "STUB: not implemented"; return }

func (c *connection) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *connection) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *connection) AcceptStream() (network.MuxedStream, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedStream), nil
}

func (c *connection) LocalPeer() peer.ID         { _ = "STUB: not implemented"; return *new(peer.ID) }
func (c *connection) RemotePeer() peer.ID        { _ = "STUB: not implemented"; return *new(peer.ID) }
func (c *connection) RemotePublicKey() ic.PubKey { _ = "STUB: not implemented"; return *new(ic.PubKey) }
func (c *connection) LocalMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}
func (c *connection) RemoteMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}
func (c *connection) Scope() network.ConnScope {
	_ = "STUB: not implemented"
	return *new(network.ConnScope)
}
func (c *connection) Transport() tpt.Transport {
	_ = "STUB: not implemented"
	return *new(tpt.Transport)
}

func (c *connection) addStream(str *stream) error { _ = "STUB: not implemented"; return nil }

func (c *connection) removeStream(id uint16) { _ = "STUB: not implemented"; return }

func (c *connection) onConnectionStateChange(state webrtc.PeerConnectionState) {
	_ = "STUB: not implemented"
	return
}

func (c *connection) detachChannel(ctx context.Context, dc *webrtc.DataChannel) (datachannel.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(datachannel.ReadWriteCloser), nil
}
