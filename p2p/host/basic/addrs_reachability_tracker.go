package basichost

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/protocol/autonatv2"
	ma "github.com/multiformats/go-multiaddr"
)

type autonatv2Client interface {
	GetReachability(ctx context.Context, reqs []autonatv2.Request) (autonatv2.Result, error)
}

const (
	maxAddrsPerRequest = 10

	maxTrackedAddrs = 50

	defaultMaxConcurrency = 5

	newAddrsProbeDelay = 1 * time.Second
)

type addrsReachabilityTracker struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	client autonatv2Client

	reachabilityUpdateCh chan struct{}
	maxConcurrency       int
	newAddrsProbeDelay   time.Duration
	probeManager         *probeManager
	newAddrs             chan []ma.Multiaddr
	clock                clock.Clock
	metricsTracker       MetricsTracker

	mx               sync.Mutex
	reachableAddrs   []ma.Multiaddr
	unreachableAddrs []ma.Multiaddr
	unknownAddrs     []ma.Multiaddr
}

func newAddrsReachabilityTracker(client autonatv2Client, reachabilityUpdateCh chan struct{}, cl clock.Clock, metricsTracker MetricsTracker) *addrsReachabilityTracker {
	_ = "STUB: not implemented"
	return nil
}

func (r *addrsReachabilityTracker) UpdateAddrs(addrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (r *addrsReachabilityTracker) ConfirmedAddrs() (reachableAddrs, unreachableAddrs, unknownAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *addrsReachabilityTracker) Start() error { _ = "STUB: not implemented"; return nil }

func (r *addrsReachabilityTracker) Close() error { _ = "STUB: not implemented"; return nil }

const (
	defaultReachabilityRefreshInterval = 5 * time.Minute

	maxBackoffInterval = 5 * time.Minute

	backoffStartInterval = 5 * time.Second
)

func (r *addrsReachabilityTracker) background() { _ = "STUB: not implemented"; return }

func newBackoffInterval(current time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *addrsReachabilityTracker) appendConfirmedAddrs(reachable, unreachable, unknown []ma.Multiaddr) (reachableAddrs, unreachableAddrs, unknownAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *addrsReachabilityTracker) notify() { _ = "STUB: not implemented"; return }

func (r *addrsReachabilityTracker) updateTrackedAddrs(addrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

type probe = []autonatv2.Request

const probeTimeout = 30 * time.Second

type reachabilityTask struct {
	Cancel context.CancelFunc

	BackoffCh chan bool
}

func (r *addrsReachabilityTracker) refreshReachability() reachabilityTask {
	_ = "STUB: not implemented"
	return *new(reachabilityTask)
}

var errTooManyConsecutiveFailures = errors.New("too many consecutive failures")

type errCountingClient struct {
	autonatv2Client
	MaxConsecutiveErrors int
	mx                   sync.Mutex
	consecutiveErrors    int
}

func (c *errCountingClient) GetReachability(ctx context.Context, reqs probe) (autonatv2.Result, error) {
	_ = "STUB: not implemented"
	return *new(autonatv2.Result), nil
}

const maxConsecutiveErrors = 20

func isErrorPersistent(err error) bool { _ = "STUB: not implemented"; return false }

const (
	recentProbeInterval = 10 * time.Minute

	maxConsecutiveRefusals = 5

	maxRecentDialsPerAddr = 10

	targetConfidence = 3

	minConfidence = 2

	maxRecentDialsWindow = targetConfidence + 2

	highConfidenceAddrProbeInterval = 1 * time.Hour

	maxProbeResultTTL = maxRecentDialsWindow * highConfidenceAddrProbeInterval
)

type probeManager struct {
	now func() time.Time

	mx                    sync.Mutex
	inProgressProbes      map[string]int
	inProgressProbesTotal int
	statuses              map[string]*addrStatus
	primaryAddrs          []ma.Multiaddr
	secondaryAddrs        []ma.Multiaddr
}

func newProbeManager(now func() time.Time) *probeManager { _ = "STUB: not implemented"; return nil }

func (m *probeManager) AppendConfirmedAddrs(reachable, unreachable, unknown []ma.Multiaddr) (reachableAddrs, unreachableAddrs, unknownAddrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *probeManager) UpdateAddrs(addrs []ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (m *probeManager) GetProbe() probe { _ = "STUB: not implemented"; return *new(probe) }

func (m *probeManager) getFirstProbeAddrIdx(addrs []ma.Multiaddr, now time.Time) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *probeManager) appendRequestsToProbe(reqs probe, addrs []ma.Multiaddr, st int, skipStart bool, now time.Time) probe {
	_ = "STUB: not implemented"
	return *new(probe)
}

func (m *probeManager) MarkProbeInProgress(reqs probe) { _ = "STUB: not implemented"; return }

func (m *probeManager) InProgressProbes() int { _ = "STUB: not implemented"; return 0 }

func (m *probeManager) CompleteProbe(reqs probe, res autonatv2.Result, err error) {
	_ = "STUB: not implemented"
	return
}

type dialOutcome struct {
	Success bool
	At      time.Time
}

type addrStatus struct {
	Addr                ma.Multiaddr
	primary             *addrStatus
	lastRefusalTime     time.Time
	consecutiveRefusals int
	dialTimes           []time.Time
	outcomes            []dialOutcome
}

func (s *addrStatus) Reachability() network.Reachability {
	_ = "STUB: not implemented"
	return *new(network.Reachability)
}

func (s *addrStatus) RequiredProbeCount(now time.Time) int { _ = "STUB: not implemented"; return 0 }

func (s *addrStatus) requiredProbeCountForConfirmation(now time.Time) int {
	_ = "STUB: not implemented"
	return 0
}

func (s *addrStatus) AddRefusal(now time.Time) { _ = "STUB: not implemented"; return }

func (s *addrStatus) AddOutcome(at time.Time, rch network.Reachability, windowSize int) {
	_ = "STUB: not implemented"
	return
}

func (s *addrStatus) RemoveBefore(t time.Time) { _ = "STUB: not implemented"; return }

func (s *addrStatus) recentDialCount(now time.Time) int { _ = "STUB: not implemented"; return 0 }

func (s *addrStatus) reachabilityAndCounts() (rch network.Reachability, successes int, failures int) {
	_ = "STUB: not implemented"
	return *new(network.Reachability), 0, 0
}

var errNotTW = errors.New("not a thinwaist address")

func thinWaistPart(a ma.Multiaddr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func assignPrimaryAddrs(statuses map[string]*addrStatus) { _ = "STUB: not implemented"; return }
