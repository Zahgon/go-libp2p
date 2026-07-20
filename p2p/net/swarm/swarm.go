package swarm

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/transport"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
)

const (
	defaultDialTimeout = 15 * time.Second

	defaultDialTimeoutLocal = 5 * time.Second

	defaultNewStreamTimeout = 15 * time.Second
)

var log = logging.Logger("swarm2")

var ErrSwarmClosed = errors.New("swarm closed")

var ErrAddrFiltered = errors.New("address filtered")

var ErrDialTimeout = errors.New("dial timed out")

type Option func(*Swarm) error

func WithConnectionGater(gater connmgr.ConnectionGater) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMultiaddrResolver(resolver network.MultiaddrDNSResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMetrics(reporter metrics.Reporter) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(t MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialTimeoutLocal(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResourceManager(m network.ResourceManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDialRanker(d network.DialRanker) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithUDPBlackHoleSuccessCounter(f *BlackHoleSuccessCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithIPv6BlackHoleSuccessCounter(f *BlackHoleSuccessCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithReadOnlyBlackHoleDetector() Option { _ = "STUB: not implemented"; return *new(Option) }

type Swarm struct {
	nextConnID   atomic.Uint64
	nextStreamID atomic.Uint64

	refs sync.WaitGroup

	emitter event.Emitter

	rcmgr network.ResourceManager

	local peer.ID
	peers peerstore.Peerstore

	dialTimeout      time.Duration
	dialTimeoutLocal time.Duration

	conns struct {
		sync.RWMutex
		m map[peer.ID][]*Conn
	}

	listeners struct {
		sync.RWMutex

		ifaceListenAddres []ma.Multiaddr
		cacheEOL          time.Time

		m map[transport.Listener]struct{}
	}

	notifs struct {
		sync.RWMutex
		m map[network.Notifiee]struct{}
	}

	directConnNotifs struct {
		sync.Mutex
		m map[peer.ID][]chan struct{}
	}

	transports struct {
		sync.RWMutex
		m map[int]transport.Transport
	}

	multiaddrResolver network.MultiaddrDNSResolver

	streamh atomic.Pointer[network.StreamHandler]

	dsync   *dialSync
	backf   DialBackoff
	limiter *dialLimiter
	gater   connmgr.ConnectionGater

	closeOnce sync.Once
	ctx       context.Context
	ctxCancel context.CancelFunc

	bwc           metrics.Reporter
	metricsTracer MetricsTracer

	dialRanker network.DialRanker

	connectionEventsEmitter *connectionEventsEmitter
	udpBHF                  *BlackHoleSuccessCounter
	ipv6BHF                 *BlackHoleSuccessCounter
	bhd                     *blackHoleDetector
	readOnlyBHD             bool
}

func NewSwarm(local peer.ID, peers peerstore.Peerstore, eventBus event.Bus, opts ...Option) (*Swarm, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Swarm) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *Swarm) close() { _ = "STUB: not implemented"; return }

func (s *Swarm) addConn(tc transport.CapableConn, dir network.Direction) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) Peerstore() peerstore.Peerstore {
	_ = "STUB: not implemented"
	return *new(peerstore.Peerstore)
}

func (s *Swarm) SetStreamHandler(handler network.StreamHandler) { _ = "STUB: not implemented"; return }

func (s *Swarm) StreamHandler() network.StreamHandler {
	_ = "STUB: not implemented"
	return *new(network.StreamHandler)
}

func (s *Swarm) NewStream(ctx context.Context, p peer.ID) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (s *Swarm) waitForDirectConn(ctx context.Context, p peer.ID) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) ConnsToPeer(p peer.ID) []network.Conn { _ = "STUB: not implemented"; return nil }

func isBetterConn(a, b *Conn) bool { _ = "STUB: not implemented"; return false }

func (s *Swarm) bestConnToPeer(p peer.ID) *Conn { _ = "STUB: not implemented"; return nil }

func (s *Swarm) bestAcceptableConnToPeer(ctx context.Context, p peer.ID) *Conn {
	_ = "STUB: not implemented"
	return nil
}

func isDirectConn(c *Conn) bool { _ = "STUB: not implemented"; return false }

func (s *Swarm) Connectedness(p peer.ID) network.Connectedness {
	_ = "STUB: not implemented"
	return *new(network.Connectedness)
}

func (s *Swarm) connectednessUnlocked(p peer.ID) network.Connectedness {
	_ = "STUB: not implemented"
	return *new(network.Connectedness)
}

func (s *Swarm) Conns() []network.Conn { _ = "STUB: not implemented"; return nil }

func (s *Swarm) ClosePeer(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (s *Swarm) Peers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (s *Swarm) LocalPeer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (s *Swarm) Backoff() *DialBackoff { _ = "STUB: not implemented"; return nil }

func (s *Swarm) notifyAll(notify func(network.Notifiee)) { _ = "STUB: not implemented"; return }

func (s *Swarm) Notify(f network.Notifiee) { _ = "STUB: not implemented"; return }

func (s *Swarm) StopNotify(f network.Notifiee) { _ = "STUB: not implemented"; return }

func (s *Swarm) removeConn(c *Conn) { _ = "STUB: not implemented"; return }

func (s *Swarm) String() string { _ = "STUB: not implemented"; return "" }

func (s *Swarm) ResourceManager() network.ResourceManager {
	_ = "STUB: not implemented"
	return *new(network.ResourceManager)
}

var (
	_ network.Network            = (*Swarm)(nil)
	_ transport.TransportNetwork = (*Swarm)(nil)
)

type connWithMetrics struct {
	transport.CapableConn
	opened        time.Time
	dir           network.Direction
	metricsTracer MetricsTracer
	once          sync.Once
	closeErr      error
}

func wrapWithMetrics(capableConn transport.CapableConn, metricsTracer MetricsTracer, opened time.Time, dir network.Direction) *connWithMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (c *connWithMetrics) As(target any) bool { _ = "STUB: not implemented"; return false }

func (c *connWithMetrics) completedHandshake() { _ = "STUB: not implemented"; return }

func (c *connWithMetrics) Close() error { _ = "STUB: not implemented"; return nil }

func (c *connWithMetrics) CloseWithError(errCode network.ConnErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *connWithMetrics) Stat() network.ConnStats {
	_ = "STUB: not implemented"
	return *new(network.ConnStats)
}

var _ network.ConnStat = &connWithMetrics{}

type ResolverFromMaDNS struct {
	*madns.Resolver
}

var _ network.MultiaddrDNSResolver = ResolverFromMaDNS{}

func startsWithDNSADDR(m ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (r ResolverFromMaDNS) ResolveDNSAddr(ctx context.Context, expectedPeerID peer.ID, maddr ma.Multiaddr, recursionLimit int, outputLimit int) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResolverFromMaDNS) ResolveDNSComponent(ctx context.Context, maddr ma.Multiaddr, outputLimit int) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) AddCertHashes(listenAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}
