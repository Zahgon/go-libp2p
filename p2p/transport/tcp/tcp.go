package tcp

import (
	"context"
	"net"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/net/reuseport"
	"github.com/libp2p/go-libp2p/p2p/transport/tcpreuse"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	mafmt "github.com/multiformats/go-multiaddr-fmt"
	manet "github.com/multiformats/go-multiaddr/net"
)

const defaultConnectTimeout = 5 * time.Second

var log = logging.Logger("tcp-tpt")

const keepAlivePeriod = 30 * time.Second

type canKeepAlive interface {
	SetKeepAlive(bool) error
	SetKeepAlivePeriod(time.Duration) error
}

var _ canKeepAlive = &net.TCPConn{}

var ReuseportIsAvailable = tcpreuse.ReuseportIsAvailable

func tryKeepAlive(conn net.Conn, keepAlive bool) { _ = "STUB: not implemented"; return }

func tryLinger(conn net.Conn, sec int) { _ = "STUB: not implemented"; return }

type tcpGatedMaListener struct {
	transport.GatedMaListener
	sec int
}

func (ll *tcpGatedMaListener) Accept() (manet.Conn, network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), *new(network.ConnManagementScope), nil
}

type Option func(*TcpTransport) error

func DisableReuseport() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnectionTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetrics() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialerForAddr(d DialerForAddr) Option { _ = "STUB: not implemented"; return *new(Option) }

type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

type DialerForAddr func(raddr ma.Multiaddr) (ContextDialer, error)

type TcpTransport struct {
	upgrader transport.Upgrader

	overrideDialerForAddr DialerForAddr

	disableReuseport bool
	enableMetrics    bool

	sharedTcp *tcpreuse.ConnMgr

	connectTimeout time.Duration

	rcmgr network.ResourceManager

	reuse reuseport.Transport

	metricsCollector *aggregatingCollector
}

var _ transport.Transport = &TcpTransport{}
var _ transport.DialUpdater = &TcpTransport{}

func NewTCPTransport(upgrader transport.Upgrader, rcmgr network.ResourceManager, sharedTCP *tcpreuse.ConnMgr, opts ...Option) (*TcpTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var dialMatcher = mafmt.And(mafmt.IP, mafmt.Base(ma.P_TCP))

func (t *TcpTransport) CanDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (t *TcpTransport) customDial(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (t *TcpTransport) maDial(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), nil
}

func (t *TcpTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (t *TcpTransport) DialWithUpdates(ctx context.Context, raddr ma.Multiaddr, p peer.ID, updateChan chan<- transport.DialUpdate) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (t *TcpTransport) dialWithScope(ctx context.Context, raddr ma.Multiaddr, p peer.ID, connScope network.ConnManagementScope, updateChan chan<- transport.DialUpdate) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (t *TcpTransport) UseReuseport() bool { _ = "STUB: not implemented"; return false }

func (t *TcpTransport) unsharedMAListen(laddr ma.Multiaddr) (manet.Listener, error) {
	_ = "STUB: not implemented"
	return *new(manet.Listener), nil
}

func (t *TcpTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	_ = "STUB: not implemented"
	return *new(transport.Listener), nil
}

func (t *TcpTransport) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (t *TcpTransport) Proxy() bool { _ = "STUB: not implemented"; return false }

func (t *TcpTransport) String() string { _ = "STUB: not implemented"; return "" }
