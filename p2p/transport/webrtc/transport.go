package libp2pwebrtc

import (
	"context"
	"crypto"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/pnet"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/security/noise"

	ma "github.com/multiformats/go-multiaddr"

	"github.com/pion/datachannel"
	"github.com/pion/webrtc/v4"
)

var webrtcComponent *ma.Component

func init() {
	var err error
	webrtcComponent, err = ma.NewComponent(ma.ProtocolWithCode(ma.P_WEBRTC_DIRECT).Name, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Executable()
	}
}

const (
	handshakeChannelNegotiated = true

	handshakeChannelID = uint16(0)
)

const (
	DefaultDisconnectedTimeout = 20 * time.Second
	DefaultFailedTimeout       = 30 * time.Second
	DefaultKeepaliveTimeout    = 15 * time.Second

	sctpReceiveBufferSize = 10 * maxReceiveMessageSize
)

type WebRTCTransport struct {
	webrtcConfig webrtc.Configuration
	rcmgr        network.ResourceManager
	gater        connmgr.ConnectionGater
	privKey      ic.PrivKey
	noiseTpt     *noise.Transport
	localPeerId  peer.ID

	listenUDP func(network string, laddr *net.UDPAddr) (net.PacketConn, error)

	dialerVersion int

	peerConnectionTimeouts iceTimeouts

	maxInFlightConnections uint32
}

var _ tpt.Transport = &WebRTCTransport{}

type Option func(*WebRTCTransport) error

func WithDialerVersion(version int) Option { _ = "STUB: not implemented"; return *new(Option) }

type iceTimeouts struct {
	Disconnect time.Duration
	Failed     time.Duration
	Keepalive  time.Duration
}

type ListenUDPFn func(network string, laddr *net.UDPAddr) (net.PacketConn, error)

func New(privKey ic.PrivKey, psk pnet.PSK, gater connmgr.ConnectionGater, rcmgr network.ResourceManager, listenUDP ListenUDPFn, opts ...Option) (*WebRTCTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *WebRTCTransport) ListenOrder() int { _ = "STUB: not implemented"; return 0 }

func (t *WebRTCTransport) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (t *WebRTCTransport) Proxy() bool { _ = "STUB: not implemented"; return false }

func (t *WebRTCTransport) CanDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (t *WebRTCTransport) Listen(addr ma.Multiaddr) (tpt.Listener, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Listener), nil
}

func (t *WebRTCTransport) listenSocket(socket net.PacketConn) (tpt.Listener, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Listener), nil
}

func (t *WebRTCTransport) Dial(ctx context.Context, remoteMultiaddr ma.Multiaddr, p peer.ID) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (t *WebRTCTransport) dial(ctx context.Context, scope network.ConnManagementScope, remoteMultiaddr ma.Multiaddr, p peer.ID) (tConn tpt.CapableConn, err error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func genUfrag() string { _ = "STUB: not implemented"; return "" }

func genV2ClientCredentials() (ufrag, pwd string) { _ = "STUB: not implemented"; return "", "" }

func randIceString(n int) string { _ = "STUB: not implemented"; return "" }

func (t *WebRTCTransport) getCertificateFingerprint() (webrtc.DTLSFingerprint, error) {
	_ = "STUB: not implemented"
	return *new(webrtc.DTLSFingerprint), nil
}

func (t *WebRTCTransport) generateNoisePrologue(pc *webrtc.PeerConnection, hash crypto.Hash, inbound bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *WebRTCTransport) noiseHandshake(ctx context.Context, pc *webrtc.PeerConnection, s *stream, peer peer.ID, hash crypto.Hash, inbound bool) (ic.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ic.PubKey), nil
}

func (t *WebRTCTransport) AddCertHashes(addr ma.Multiaddr) (ma.Multiaddr, bool) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), false
}

type netConnWrapper struct {
	*stream
}

func (netConnWrapper) LocalAddr() net.Addr  { _ = "STUB: not implemented"; return *new(net.Addr) }
func (netConnWrapper) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
func (w netConnWrapper) Close() error       { _ = "STUB: not implemented"; return nil }

func detachHandshakeDataChannel(ctx context.Context, dc *webrtc.DataChannel) (datachannel.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(datachannel.ReadWriteCloser), nil
}

type webRTCConnection struct {
	PeerConnection         *webrtc.PeerConnection
	HandshakeDataChannel   *webrtc.DataChannel
	IncomingDataChannels   chan dataChannel
	PeerConnectionClosedCh chan struct{}
}

func newWebRTCConnection(settings webrtc.SettingEngine, config webrtc.Configuration) (webRTCConnection, error) {
	_ = "STUB: not implemented"
	return *new(webRTCConnection), nil
}

func IsWebRTCDirectMultiaddr(addr ma.Multiaddr) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}
