package autorelay

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	circuitv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	circuitv2_proto "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/proto"

	ma "github.com/multiformats/go-multiaddr"
)

const protoIDv2 = circuitv2_proto.ProtoIDv2Hop

const (
	rsvpRefreshInterval = time.Minute
	rsvpExpirationSlack = 2 * time.Minute

	autorelayTag  = "autorelay"
	maxRelayAddrs = 100
)

type candidate struct {
	added           time.Time
	supportsRelayV2 bool
	ai              peer.AddrInfo
}

type relayFinder struct {
	bootTime time.Time
	host     host.Host

	conf *config

	refCount sync.WaitGroup

	ctxCancel   context.CancelFunc
	ctxCancelMx sync.Mutex

	peerSource PeerSource

	candidateFound             chan struct{}
	candidateMx                sync.Mutex
	candidates                 map[peer.ID]*candidate
	backoff                    map[peer.ID]time.Time
	maybeConnectToRelayTrigger chan struct{}

	maybeRequestNewCandidates chan struct{}

	relayReservationUpdated chan struct{}

	relayMx sync.Mutex
	relays  map[peer.ID]*circuitv2.Reservation

	circuitAddrs []ma.Multiaddr

	triggerRunScheduledWork chan struct{}
	metricsTracer           MetricsTracer

	emitter event.Emitter
}

var errAlreadyRunning = errors.New("relayFinder already running")

func newRelayFinder(host host.Host, conf *config) (*relayFinder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type scheduledWorkTimes struct {
	leastFrequentInterval       time.Duration
	nextRefresh                 time.Time
	nextBackoff                 time.Time
	nextOldCandidateCheck       time.Time
	nextAllowedCallToPeerSource time.Time
}

func (rf *relayFinder) cleanupDisconnectedPeers(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (rf *relayFinder) background(ctx context.Context) { _ = "STUB: not implemented"; return }

func (rf *relayFinder) updateAddrs() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) getCircuitAddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (rf *relayFinder) runScheduledWork(ctx context.Context, now time.Time, scheduledWork *scheduledWorkTimes, peerSourceRateLimiter chan<- struct{}) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (rf *relayFinder) clearOldCandidates(now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (rf *relayFinder) clearBackoff(now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (rf *relayFinder) findNodes(ctx context.Context, peerSourceRateLimiter <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (rf *relayFinder) notifyMaybeConnectToRelay() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) notifyMaybeNeedNewCandidates() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) notifyNewCandidate() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) notifyRelayReservationUpdated() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) handleNewNode(ctx context.Context, pi peer.AddrInfo) (added bool) {
	_ = "STUB: not implemented"
	return false
}

var errProtocolNotSupported = errors.New("doesn't speak circuit v2")

func (rf *relayFinder) tryNode(ctx context.Context, pi peer.AddrInfo) (supportsRelayV2 bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rf *relayFinder) handleNewCandidates(ctx context.Context) { _ = "STUB: not implemented"; return }

func (rf *relayFinder) maybeConnectToRelay(ctx context.Context) { _ = "STUB: not implemented"; return }

func (rf *relayFinder) connectToRelay(ctx context.Context, cand *candidate) (*circuitv2.Reservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rf *relayFinder) refreshReservations(ctx context.Context, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (rf *relayFinder) refreshRelayReservation(ctx context.Context, p peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (rf *relayFinder) usingRelay(p peer.ID) bool { _ = "STUB: not implemented"; return false }

func (rf *relayFinder) addCandidate(cand *candidate) { _ = "STUB: not implemented"; return }

func (rf *relayFinder) removeCandidate(id peer.ID) { _ = "STUB: not implemented"; return }

func (rf *relayFinder) selectCandidates() []*candidate { _ = "STUB: not implemented"; return nil }

func (rf *relayFinder) Start() error { _ = "STUB: not implemented"; return nil }

func (rf *relayFinder) Stop() error { _ = "STUB: not implemented"; return nil }

func (rf *relayFinder) initMetrics() { _ = "STUB: not implemented"; return }

func (rf *relayFinder) resetMetrics() { _ = "STUB: not implemented"; return }

func areSortedAddrsDifferent(a, b []ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }
