package libp2pwebtransport

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/webtransport-go"
)

const queueLen = 16
const handshakeTimeout = 10 * time.Second

type connKey struct{}

type listener struct {
	transport       *transport
	isStaticTLSConf bool
	reuseListener   quicreuse.Listener

	server webtransport.Server

	ctx       context.Context
	ctxCancel context.CancelFunc

	serverClosed chan struct{}

	addr      net.Addr
	multiaddr ma.Multiaddr

	queue chan tpt.CapableConn

	mx           sync.Mutex
	pendingConns map[*quic.Conn]*negotiatingConn
}

var _ tpt.Listener = &listener{}

func newListener(reuseListener quicreuse.Listener, t *transport, isStaticTLSConf bool) (tpt.Listener, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Listener), nil
}

func (l *listener) startHandshake(conn *quic.Conn) error { _ = "STUB: not implemented"; return nil }

type negotiatingConn struct {
	*quic.Conn
	ctx    context.Context
	cancel context.CancelFunc

	stopHandshakeTimeout func() bool
	err                  error
}

func (c *negotiatingConn) StopHandshakeTimeout() error { _ = "STUB: not implemented"; return nil }

var errTimeout = errors.New("timeout")

func (l *listener) httpHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (l *listener) httpHandlerWithConnScope(w http.ResponseWriter, r *http.Request, connScope network.ConnManagementScope) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *listener) Accept() (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (l *listener) handshake(ctx context.Context, sess *webtransport.Session) (*connSecurityMultiaddrs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *listener) Multiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }
