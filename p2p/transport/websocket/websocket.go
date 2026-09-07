package websocket

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/transport/tcpreuse"

	ma "github.com/multiformats/go-multiaddr"
	mafmt "github.com/multiformats/go-multiaddr-fmt"
	manet "github.com/multiformats/go-multiaddr/net"
)

var WsFmt = mafmt.And(mafmt.TCP, mafmt.Base(ma.P_WS))

var dialMatcher = mafmt.And(
	mafmt.Or(mafmt.IP, mafmt.DNS),
	mafmt.Base(ma.P_TCP),
	mafmt.Or(
		mafmt.Base(ma.P_WS),
		mafmt.And(
			mafmt.Or(
				mafmt.And(
					mafmt.Base(ma.P_TLS),
					mafmt.Base(ma.P_SNI)),
				mafmt.Base(ma.P_TLS),
			),
			mafmt.Base(ma.P_WS)),
		mafmt.Base(ma.P_WSS)))

var (
	wssComponent, _ = ma.NewComponent("wss", "")
	tlsComponent, _ = ma.NewComponent("tls", "")
	wsComponent, _  = ma.NewComponent("ws", "")
	tlsWsAddr       = ma.Multiaddr{*tlsComponent, *wsComponent}
)

func init() {
	manet.RegisterFromNetAddr(ParseWebsocketNetAddr, "websocket")
	manet.RegisterToNetAddr(ConvertWebsocketMultiaddrToNetAddr, "ws")
	manet.RegisterToNetAddr(ConvertWebsocketMultiaddrToNetAddr, "wss")
}

type Option func(*WebsocketTransport) error

func WithTLSClientConfig(c *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLSConfig(conf *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

var defaultHandshakeTimeout = 15 * time.Second

var defaultHTTPIdleTimeout = 30 * time.Second

func WithHandshakeTimeout(timeout time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHTTPHandler(h http.Handler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPServerConfig(fn func(*http.Server)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type WebsocketTransport struct {
	upgrader         transport.Upgrader
	rcmgr            network.ResourceManager
	tlsClientConf    *tls.Config
	tlsConf          *tls.Config
	sharedTcp        *tcpreuse.ConnMgr
	handshakeTimeout time.Duration
	httpHandler      http.Handler
	httpServerConfig func(*http.Server)
}

var _ transport.Transport = (*WebsocketTransport)(nil)

func New(u transport.Upgrader, rcmgr network.ResourceManager, sharedTCP *tcpreuse.ConnMgr, opts ...Option) (*WebsocketTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *WebsocketTransport) CanDial(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (t *WebsocketTransport) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (t *WebsocketTransport) Proxy() bool { _ = "STUB: not implemented"; return false }

func (t *WebsocketTransport) Resolve(_ context.Context, maddr ma.Multiaddr) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *WebsocketTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (t *WebsocketTransport) dialWithScope(ctx context.Context, raddr ma.Multiaddr, p peer.ID, connScope network.ConnManagementScope) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (t *WebsocketTransport) maDial(ctx context.Context, raddr ma.Multiaddr, scope network.ConnManagementScope) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (t *WebsocketTransport) gatedMaListen(a ma.Multiaddr) (transport.GatedMaListener, error) {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener), nil
}

func (t *WebsocketTransport) Listen(a ma.Multiaddr) (transport.Listener, error) {
	_ = "STUB: not implemented"
	return *new(transport.Listener), nil
}

type transportListener struct {
	transport.Listener
}

type capableConn struct {
	transport.CapableConn
}

func (c *capableConn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}

func (l *transportListener) Accept() (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}
