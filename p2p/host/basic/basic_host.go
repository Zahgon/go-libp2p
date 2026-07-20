package basichost

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/host/autonat"
	"github.com/libp2p/go-libp2p/p2p/host/pstoremanager"
	"github.com/libp2p/go-libp2p/p2p/host/relaysvc"
	"github.com/libp2p/go-libp2p/p2p/protocol/autonatv2"
	relayv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/prometheus/client_golang/prometheus"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	msmux "github.com/multiformats/go-multistream"
)

var log = logging.Logger("basichost")

var (
	DefaultNegotiationTimeout = 10 * time.Second

	DefaultAddrsFactory = func(addrs []ma.Multiaddr) []ma.Multiaddr { return addrs }
)

type AddrsFactory func([]ma.Multiaddr) []ma.Multiaddr

type BasicHost struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	closeSync sync.Once

	refCount sync.WaitGroup

	network      network.Network
	psManager    *pstoremanager.PeerstoreManager
	mux          *msmux.MultistreamMuxer[protocol.ID]
	ids          identify.IDService
	hps          *holepunch.Service
	pings        *ping.PingService
	cmgr         connmgr.ConnManager
	eventbus     event.Bus
	relayManager *relaysvc.RelayManager

	negtimeout time.Duration

	emitters struct {
		evtLocalProtocolsUpdated event.Emitter
	}

	autoNATMx sync.RWMutex
	autoNat   autonat.AutoNAT

	autonatv2      *autonatv2.AutoNAT
	addressManager *addrsManager
}

var _ host.Host = (*BasicHost)(nil)

type HostOpts struct {
	EventBus event.Bus

	MultistreamMuxer *msmux.MultistreamMuxer[protocol.ID]

	NegotiationTimeout time.Duration

	AddrsFactory AddrsFactory

	NATManager func(network.Network) NATManager

	ConnManager connmgr.ConnManager

	EnablePing bool

	EnableRelayService bool

	RelayServiceOpts []relayv2.Option

	UserAgent string

	ProtocolVersion string

	DisableSignedPeerRecord bool

	DisableNonPublicAddrPublishing bool

	EnableHolePunching bool

	HolePunchingOptions []holepunch.Option

	EnableMetrics bool

	PrometheusRegisterer prometheus.Registerer

	AutoNATv2MetricsTracker MetricsTracker

	ObservedAddrsManager ObservedAddrsManager

	AutoNATv2 *autonatv2.AutoNAT
}

func NewHost(n network.Network, opts *HostOpts) (*BasicHost, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *BasicHost) Start() { _ = "STUB: not implemented"; return }

func (h *BasicHost) newStreamHandler(s network.Stream) { _ = "STUB: not implemented"; return }

func (h *BasicHost) ID() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (h *BasicHost) Peerstore() peerstore.Peerstore {
	_ = "STUB: not implemented"
	return *new(peerstore.Peerstore)
}

func (h *BasicHost) Network() network.Network {
	_ = "STUB: not implemented"
	return *new(network.Network)
}

func (h *BasicHost) Mux() protocol.Switch { _ = "STUB: not implemented"; return *new(protocol.Switch) }

func (h *BasicHost) IDService() identify.IDService {
	_ = "STUB: not implemented"
	return *new(identify.IDService)
}

func (h *BasicHost) EventBus() event.Bus { _ = "STUB: not implemented"; return *new(event.Bus) }

func (h *BasicHost) SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (h *BasicHost) SetStreamHandlerMatch(pid protocol.ID, m func(protocol.ID) bool, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (h *BasicHost) RemoveStreamHandler(pid protocol.ID) { _ = "STUB: not implemented"; return }

func (h *BasicHost) NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (str network.Stream, strErr error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (h *BasicHost) preferredProtocol(p peer.ID, pids []protocol.ID) (protocol.ID, error) {
	_ = "STUB: not implemented"
	return *new(protocol.ID), nil
}

func (h *BasicHost) Connect(ctx context.Context, pi peer.AddrInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *BasicHost) dialPeer(ctx context.Context, p peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *BasicHost) ConnManager() connmgr.ConnManager {
	_ = "STUB: not implemented"
	return *new(connmgr.ConnManager)
}

func (h *BasicHost) Addrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (h *BasicHost) AllAddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (h *BasicHost) ConfirmedAddrs() (reachable []ma.Multiaddr, unreachable []ma.Multiaddr, unknown []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h *BasicHost) SetAutoNat(a autonat.AutoNAT) { _ = "STUB: not implemented"; return }

func (h *BasicHost) GetAutoNat() autonat.AutoNAT {
	_ = "STUB: not implemented"
	return *new(autonat.AutoNAT)
}

func (h *BasicHost) Reachability() network.Reachability {
	_ = "STUB: not implemented"
	return *new(network.Reachability)
}

func (h *BasicHost) Close() error { _ = "STUB: not implemented"; return nil }

type streamWrapper struct {
	network.Stream
	rw io.ReadWriteCloser
}

func (s *streamWrapper) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *streamWrapper) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *streamWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (s *streamWrapper) CloseWrite() error { _ = "STUB: not implemented"; return nil }
