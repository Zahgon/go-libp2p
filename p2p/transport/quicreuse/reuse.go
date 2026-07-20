package quicreuse

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"time"

	"github.com/libp2p/go-netroute"
	"github.com/quic-go/quic-go"
)

type RefCountedQUICTransport interface {
	LocalAddr() net.Addr

	WriteTo([]byte, net.Addr) (int, error)

	Close() error

	DecreaseCount()
	IncreaseCount()

	Dial(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *quic.Config) (*quic.Conn, error)
	Listen(tlsConf *tls.Config, conf *quic.Config) (QUICListener, error)
}

type singleOwnerTransport struct {
	Transport QUICTransport

	packetConn net.PacketConn
}

var _ QUICTransport = &singleOwnerTransport{}
var _ RefCountedQUICTransport = (*singleOwnerTransport)(nil)

func (c *singleOwnerTransport) IncreaseCount() { _ = "STUB: not implemented"; return }
func (c *singleOwnerTransport) DecreaseCount() { _ = "STUB: not implemented"; return }
func (c *singleOwnerTransport) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func (c *singleOwnerTransport) Dial(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *quic.Config) (*quic.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *singleOwnerTransport) ReadNonQUICPacket(ctx context.Context, b []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *singleOwnerTransport) Close() error { _ = "STUB: not implemented"; return nil }

func (c *singleOwnerTransport) WriteTo(b []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *singleOwnerTransport) Listen(tlsConf *tls.Config, conf *quic.Config) (QUICListener, error) {
	_ = "STUB: not implemented"
	return *new(QUICListener), nil
}

var (
	garbageCollectInterval = 30 * time.Second
	maxUnusedDuration      = 10 * time.Second
)

type refcountedTransport struct {
	QUICTransport

	packetConn net.PacketConn

	mutex       sync.Mutex
	refCount    int
	unusedSince time.Time

	borrowDoneSignal chan struct{}

	associations map[any]map[*listener]struct{}
}

type connContextFunc = func(context.Context, *quic.ClientInfo) (context.Context, error)

func (c *refcountedTransport) associateForListener(a any, ln *listener) {
	_ = "STUB: not implemented"
	return
}

func (c *refcountedTransport) RemoveAssociationsForListener(ln *listener) {
	_ = "STUB: not implemented"
	return
}

func (c *refcountedTransport) hasAssociation(a any) bool { _ = "STUB: not implemented"; return false }

func (c *refcountedTransport) IncreaseCount() { _ = "STUB: not implemented"; return }

func (c *refcountedTransport) Close() error { _ = "STUB: not implemented"; return nil }

func (c *refcountedTransport) WriteTo(b []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *refcountedTransport) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func (c *refcountedTransport) Listen(tlsConf *tls.Config, conf *quic.Config) (QUICListener, error) {
	_ = "STUB: not implemented"
	return *new(QUICListener), nil
}

func (c *refcountedTransport) DecreaseCount() { _ = "STUB: not implemented"; return }

func (c *refcountedTransport) ShouldGarbageCollect(now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

type reuse struct {
	mutex sync.Mutex

	closeChan  chan struct{}
	gcStopChan chan struct{}

	listenUDP listenUDP

	sourceIPSelectorFn func() (SourceIPSelector, error)

	routes  SourceIPSelector
	unicast map[string]map[int]*refcountedTransport

	globalListeners map[int]*refcountedTransport

	globalDialers map[int]*refcountedTransport

	statelessResetKey   *quic.StatelessResetKey
	tokenGeneratorKey   *quic.TokenGeneratorKey
	connContext         connContextFunc
	verifySourceAddress func(addr net.Addr) bool
}

func newReuse(srk *quic.StatelessResetKey, tokenKey *quic.TokenGeneratorKey, listenUDP listenUDP, sourceIPSelectorFn func() (SourceIPSelector, error),
	connContext connContextFunc, verifySourceAddress func(addr net.Addr) bool) *reuse {
	_ = "STUB: not implemented"
	return nil
}

func (r *reuse) gc() { _ = "STUB: not implemented"; return }

func (r *reuse) TransportWithAssociationForDial(association any, network string, raddr *net.UDPAddr) (*refcountedTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *reuse) transportForDialLocked(association any, network string, source *net.IP) (*refcountedTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *reuse) AddTransport(tr *refcountedTransport, laddr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reuse) TransportForListen(network string, laddr *net.UDPAddr) (*refcountedTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *reuse) newTransport(conn net.PacketConn) *refcountedTransport {
	_ = "STUB: not implemented"
	return nil
}

func (r *reuse) Close() error { _ = "STUB: not implemented"; return nil }

type SourceIPSelector interface {
	PreferredSourceIPForDestination(dst *net.UDPAddr) (net.IP, error)
}

type netrouteSourceIPSelector struct {
	routes netroute.Router
}

func (s *netrouteSourceIPSelector) PreferredSourceIPForDestination(dst *net.UDPAddr) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}
