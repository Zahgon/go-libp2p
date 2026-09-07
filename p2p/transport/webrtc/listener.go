package libp2pwebrtc

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/transport/webrtc/udpmux"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/pion/webrtc/v4"
)

type connMultiaddrs struct {
	local, remote ma.Multiaddr
}

var _ network.ConnMultiaddrs = &connMultiaddrs{}

func (c *connMultiaddrs) LocalMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}
func (c *connMultiaddrs) RemoteMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

const (
	candidateSetupTimeout = 10 * time.Second

	DefaultMaxInFlightConnections = 128
)

type listener struct {
	transport *WebRTCTransport

	mux *udpmux.UDPMux

	config                    webrtc.Configuration
	localFingerprint          webrtc.DTLSFingerprint
	localFingerprintMultibase string

	localAddr      net.Addr
	localMultiaddr ma.Multiaddr

	acceptQueue chan tpt.CapableConn

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

var _ tpt.Listener = &listener{}

func newListener(transport *WebRTCTransport, laddr ma.Multiaddr, socket net.PacketConn, config webrtc.Configuration) (*listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) listen() { _ = "STUB: not implemented"; return }

func (l *listener) handleCandidate(ctx context.Context, candidate udpmux.Candidate) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (l *listener) setupConnection(
	ctx context.Context, scope network.ConnManagementScope,
	remoteMultiaddr ma.Multiaddr, candidate udpmux.Candidate,
) (tConn tpt.CapableConn, err error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (l *listener) Accept() (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *listener) Multiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func addOnConnectionStateChangeCallback(pc *webrtc.PeerConnection) <-chan error {
	_ = "STUB: not implemented"
	return nil
}
