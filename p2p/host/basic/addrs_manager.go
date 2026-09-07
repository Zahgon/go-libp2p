package basichost

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/record"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prometheus/client_golang/prometheus"
)

const maxObservedAddrsPerListenAddr = 3

var addrChangeTickrInterval = 5 * time.Second

const maxPeerRecordSize = 8 * 1024

type addrStore interface {
	SetAddrs(peer.ID, []ma.Multiaddr, time.Duration)
}

type ObservedAddrsManager interface {
	Addrs(minObservers int) []ma.Multiaddr
	AddrsFor(local ma.Multiaddr) []ma.Multiaddr
}

type hostAddrs struct {
	addrs            []ma.Multiaddr
	localAddrs       []ma.Multiaddr
	reachableAddrs   []ma.Multiaddr
	unreachableAddrs []ma.Multiaddr
	unknownAddrs     []ma.Multiaddr
	relayAddrs       []ma.Multiaddr
}

type addrsManager struct {
	bus                      event.Bus
	natManager               NATManager
	addrsFactory             AddrsFactory
	listenAddrs              func() []ma.Multiaddr
	addCertHashes            func([]ma.Multiaddr) []ma.Multiaddr
	observedAddrsManager     ObservedAddrsManager
	interfaceAddrs           *interfaceAddrsCache
	addrsReachabilityTracker *addrsReachabilityTracker

	triggerAddrsUpdateChan chan chan struct{}

	started atomic.Bool

	triggerReachabilityUpdate chan struct{}

	hostReachability atomic.Pointer[network.Reachability]

	addrsMx      sync.RWMutex
	currentAddrs hostAddrs

	signKey                        crypto.PrivKey
	addrStore                      addrStore
	signedRecordStore              peerstore.CertifiedAddrBook
	hostID                         peer.ID
	disableNonPublicAddrPublishing bool

	wg        sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func newAddrsManager(
	bus event.Bus,
	natmgr NATManager,
	addrsFactory AddrsFactory,
	listenAddrs func() []ma.Multiaddr,
	addCertHashes func([]ma.Multiaddr) []ma.Multiaddr,
	observedAddrsManager ObservedAddrsManager,
	client autonatv2Client,
	enableMetrics bool,
	registerer prometheus.Registerer,
	disableSignedPeerRecord bool,
	disableNonPublicAddrPublishing bool,
	signKey crypto.PrivKey,
	addrStore addrStore,
	hostID peer.ID,
) (*addrsManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *addrsManager) Start() error { _ = "STUB: not implemented"; return nil }

func (a *addrsManager) Close() { _ = "STUB: not implemented"; return }

func (a *addrsManager) NetNotifee() network.Notifiee {
	_ = "STUB: not implemented"
	return *new(network.Notifiee)
}

func (a *addrsManager) updateAddrsSync() { _ = "STUB: not implemented"; return }

func (a *addrsManager) startBackgroundWorker() (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) background(
	autoRelayAddrsSub,
	autonatReachabilitySub event.Subscription,
	emitter event.Emitter,
	localAddrsEmitter event.Emitter,
) {
	_ = "STUB: not implemented"
	return
}

func (a *addrsManager) updateAddrs(prevHostAddrs hostAddrs, relayAddrs []ma.Multiaddr) hostAddrs {
	_ = "STUB: not implemented"
	return *new(hostAddrs)
}

func (a *addrsManager) updatePeerStore(currentAddrs []ma.Multiaddr, removedAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func filterPublicAddrs(addrs []ma.Multiaddr) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func hasIPOrDNSComponent(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (a *addrsManager) notifyAddrsUpdated(emitter event.Emitter, localAddrsEmitter event.Emitter, previous, current hostAddrs) {
	_ = "STUB: not implemented"
	return
}

func (a *addrsManager) Addrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (a *addrsManager) getDialableAddrs(localAddrs, reachableAddrs, unreachableAddrs, relayAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) applyAddrsFactory(addrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) HolePunchAddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (a *addrsManager) DirectAddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (a *addrsManager) ConfirmedAddrs() (reachable []ma.Multiaddr, unreachable []ma.Multiaddr, unknown []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (a *addrsManager) getConfirmedAddrs(localAddrs []ma.Multiaddr) (reachableAddrs, unreachableAddrs, unknownAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var p2pCircuitAddr = ma.StringCast("/p2p-circuit")

func (a *addrsManager) getLocalAddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (a *addrsManager) appendInterfaceAddrs(dst []ma.Multiaddr, listenAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) appendNATAddrs(dst []ma.Multiaddr, listenAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) appendObservedAddrs(dst []ma.Multiaddr, listenAddrs, ifaceAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (a *addrsManager) makeSignedPeerRecord(addrs []ma.Multiaddr) (*record.Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *addrsManager) emitLocalAddrsUpdated(emitter event.Emitter, currentAddrs []ma.Multiaddr, lastAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func areAddrsDifferent(prev, current []ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func diffAddrs(prev, current []ma.Multiaddr) (added, maintained, removed []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func trimHostAddrList(addrs []ma.Multiaddr, maxSize int) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

const interfaceAddrsCacheTTL = time.Minute

type interfaceAddrsCache struct {
	mx          sync.RWMutex
	all         []ma.Multiaddr
	lastUpdated time.Time
}

func (i *interfaceAddrsCache) All() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (i *interfaceAddrsCache) update() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (i *interfaceAddrsCache) updateUnlocked() { _ = "STUB: not implemented"; return }

func removeNotInSource(addrs, source []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func removeInSource(addrs, source []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

type multiCloser []io.Closer

func (mc *multiCloser) Close() error { _ = "STUB: not implemented"; return nil }
