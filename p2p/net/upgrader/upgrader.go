package upgrader

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ipnet "github.com/libp2p/go-libp2p/core/pnet"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/sec"
	"github.com/libp2p/go-libp2p/core/transport"

	manet "github.com/multiformats/go-multiaddr/net"
	mss "github.com/multiformats/go-multistream"
)

var ErrNilPeer = errors.New("nil peer")

var AcceptQueueLength = 16

const (
	defaultAcceptTimeout    = 15 * time.Second
	defaultNegotiateTimeout = 60 * time.Second
)

type Option func(*upgrader) error

func WithAcceptTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type StreamMuxer struct {
	ID    protocol.ID
	Muxer network.Multiplexer
}

type upgrader struct {
	psk       ipnet.PSK
	connGater connmgr.ConnectionGater
	rcmgr     network.ResourceManager

	muxerMuxer *mss.MultistreamMuxer[protocol.ID]
	muxers     []StreamMuxer
	muxerIDs   []protocol.ID

	security      []sec.SecureTransport
	securityMuxer *mss.MultistreamMuxer[protocol.ID]
	securityIDs   []protocol.ID

	acceptTimeout time.Duration
}

var _ transport.Upgrader = &upgrader{}

func New(security []sec.SecureTransport, muxers []StreamMuxer, psk ipnet.PSK, rcmgr network.ResourceManager, connGater connmgr.ConnectionGater, opts ...Option) (transport.Upgrader, error) {
	_ = "STUB: not implemented"
	return *new(transport.Upgrader), nil
}

func (u *upgrader) UpgradeListener(t transport.Transport, list manet.Listener) transport.Listener {
	_ = "STUB: not implemented"
	return *new(transport.Listener)
}

func (u *upgrader) GateMaListener(l manet.Listener) transport.GatedMaListener {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener)
}

func (u *upgrader) UpgradeGatedMaListener(t transport.Transport, l transport.GatedMaListener) transport.Listener {
	_ = "STUB: not implemented"
	return *new(transport.Listener)
}

func (u *upgrader) Upgrade(ctx context.Context, t transport.Transport, maconn manet.Conn, dir network.Direction, p peer.ID, connScope network.ConnManagementScope) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (u *upgrader) upgrade(ctx context.Context, t transport.Transport, maconn manet.Conn, dir network.Direction, p peer.ID, connScope network.ConnManagementScope) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (u *upgrader) setupSecurity(ctx context.Context, conn net.Conn, p peer.ID, isServer bool) (sec.SecureConn, protocol.ID, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureConn), *new(protocol.ID), nil
}

func (u *upgrader) negotiateMuxer(nc net.Conn, isServer bool) (*StreamMuxer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *upgrader) getMuxerByID(id protocol.ID) *StreamMuxer { _ = "STUB: not implemented"; return nil }

func (u *upgrader) setupMuxer(ctx context.Context, conn sec.SecureConn, server bool, scope network.PeerScope) (protocol.ID, network.MuxedConn, error) {
	_ = "STUB: not implemented"
	return *new(protocol.ID), *new(network.MuxedConn), nil
}

func (u *upgrader) getSecurityByID(id protocol.ID) sec.SecureTransport {
	_ = "STUB: not implemented"
	return *new(sec.SecureTransport)
}

func (u *upgrader) negotiateSecurity(ctx context.Context, insecure net.Conn, server bool) (sec.SecureTransport, error) {
	_ = "STUB: not implemented"
	return *new(sec.SecureTransport), nil
}
