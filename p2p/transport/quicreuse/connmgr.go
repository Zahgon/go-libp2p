package quicreuse

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"sync"

	"github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/quic-go/quic-go"
)

var log = gologshim.Logger("quicreuse")

type QUICListener interface {
	Accept(ctx context.Context) (*quic.Conn, error)
	Close() error
	Addr() net.Addr
}

var _ QUICListener = &quic.Listener{}

type QUICTransport interface {
	Listen(tlsConf *tls.Config, conf *quic.Config) (QUICListener, error)
	Dial(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *quic.Config) (*quic.Conn, error)
	WriteTo(b []byte, addr net.Addr) (int, error)
	ReadNonQUICPacket(ctx context.Context, b []byte) (int, net.Addr, error)
	io.Closer
}

type ConnManager struct {
	reuseUDP4       *reuse
	reuseUDP6       *reuse
	enableReuseport bool

	listenUDP          listenUDP
	sourceIPSelectorFn func() (SourceIPSelector, error)

	enableMetrics bool
	registerer    prometheus.Registerer

	serverConfig *quic.Config
	clientConfig *quic.Config

	quicListenersMu sync.Mutex
	quicListeners   map[string]quicListenerEntry

	srk         quic.StatelessResetKey
	tokenKey    quic.TokenGeneratorKey
	connContext connContextFunc

	verifySourceAddress func(addr net.Addr) bool

	qlogTracerDir string
}

type quicListenerEntry struct {
	refCount int
	ln       *quicListener
}

func defaultListenUDP(network string, laddr *net.UDPAddr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

func defaultSourceIPSelectorFn() (SourceIPSelector, error) {
	_ = "STUB: not implemented"
	return *new(SourceIPSelector), nil
}

const (
	unverifiedAddressNewConnectionRPS   = 1000
	unverifiedAddressNewConnectionBurst = 1000
)

func NewConnManager(statelessResetKey quic.StatelessResetKey, tokenKey quic.TokenGeneratorKey, opts ...Option) (*ConnManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnManager) getReuse(network string) (*reuse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnManager) LendTransport(network string, tr QUICTransport, conn net.PacketConn) (<-chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnManager) ListenQUIC(addr ma.Multiaddr, tlsConf *tls.Config, allowWindowIncrease func(conn *quic.Conn, delta uint64) bool) (Listener, error) {
	_ = "STUB: not implemented"
	return *new(Listener), nil
}

func (c *ConnManager) ListenQUICAndAssociate(association any, addr ma.Multiaddr, tlsConf *tls.Config, allowWindowIncrease func(conn *quic.Conn, delta uint64) bool) (Listener, error) {
	_ = "STUB: not implemented"
	return *new(Listener), nil
}

func (c *ConnManager) onListenerClosed(key string) { _ = "STUB: not implemented"; return }

func (c *ConnManager) SharedNonQUICPacketConn(_ string, laddr *net.UDPAddr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

func (c *ConnManager) transportForListen(network string, laddr *net.UDPAddr) (RefCountedQUICTransport, error) {
	_ = "STUB: not implemented"
	return *new(RefCountedQUICTransport), nil
}

type associationKey struct{}

func WithAssociation(ctx context.Context, association any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *ConnManager) DialQUIC(ctx context.Context, raddr ma.Multiaddr, tlsConf *tls.Config, allowWindowIncrease func(conn *quic.Conn, delta uint64) bool) (*quic.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnManager) TransportForDial(network string, raddr *net.UDPAddr) (RefCountedQUICTransport, error) {
	_ = "STUB: not implemented"
	return *new(RefCountedQUICTransport), nil
}

func (c *ConnManager) TransportWithAssociationForDial(association any, network string, raddr *net.UDPAddr) (RefCountedQUICTransport, error) {
	_ = "STUB: not implemented"
	return *new(RefCountedQUICTransport), nil
}

func (c *ConnManager) newSingleOwnerTransport(conn net.PacketConn) *singleOwnerTransport {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConnManager) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (c *ConnManager) Close() error { _ = "STUB: not implemented"; return nil }

func (c *ConnManager) ClientConfig() *quic.Config { _ = "STUB: not implemented"; return nil }

type wrappedQUICTransport struct {
	*quic.Transport
}

var _ QUICTransport = (*wrappedQUICTransport)(nil)

func (t *wrappedQUICTransport) Listen(tlsConf *tls.Config, conf *quic.Config) (QUICListener, error) {
	_ = "STUB: not implemented"
	return *new(QUICListener), nil
}

func newQUICTransport(
	conn net.PacketConn,
	tokenGeneratorKey *quic.TokenGeneratorKey,
	statelessResetKey *quic.StatelessResetKey,
	connContext connContextFunc,
	verifySourceAddress func(addr net.Addr) bool,
) *quic.Transport {
	_ = "STUB: not implemented"
	return nil
}
