package websocket

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
	logging "github.com/libp2p/go-libp2p/gologshim"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/transport/tcpreuse"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var log = logging.Logger("websocket-transport")

type listener struct {
	netListener *httpNetListener
	server      http.Server
	wsUpgrader  ws.Upgrader

	isWss bool

	laddr ma.Multiaddr

	incoming chan *Conn

	closeOnce sync.Once
	closeErr  error
	closed    chan struct{}
	wsurl     *url.URL

	httpHandler http.Handler
}

var _ transport.GatedMaListener = &listener{}

func (pwma *parsedWebsocketMultiaddr) toMultiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func newListener(a ma.Multiaddr, tlsConf *tls.Config, sharedTcp *tcpreuse.ConnMgr, upgrader transport.Upgrader, handshakeTimeout time.Duration, httpHandler http.Handler, serverConfig func(*http.Server)) (*listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) serve() { _ = "STUB: not implemented"; return }

type connKey struct{}

func (l *listener) ConnContext(ctx context.Context, c net.Conn) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (l *listener) extractConnFromContext(ctx context.Context) (*negotiatingConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (l *listener) Accept() (manet.Conn, network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), *new(network.ConnManagementScope), nil
}

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *listener) Multiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

type httpNetListener struct {
	transport.GatedMaListener
	handshakeTimeout time.Duration
}

var _ net.Listener = &httpNetListener{}

func (l *httpNetListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type connWithScope struct {
	net.Conn
	Scope network.ConnManagementScope
}

func (c connWithScope) Close() error { _ = "STUB: not implemented"; return nil }

type negotiatingConn struct {
	connWithScope
	ctx        context.Context
	cancelCtx  context.CancelFunc
	stopClose  func() bool
	disarmOnce sync.Once
	timedOut   bool
}

func (c *negotiatingConn) disarm() (alive bool) { _ = "STUB: not implemented"; return false }

func (c *negotiatingConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *negotiatingConn) Unwrap() (connWithScope, error) {
	_ = "STUB: not implemented"
	return *new(connWithScope), nil
}
