package connmgr

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("connmgr")

type BasicConnMgr struct {
	*decayer

	clock clock.Clock

	cfg      *config
	segments segments

	plk       sync.RWMutex
	protected map[peer.ID]map[string]struct{}

	trimMutex sync.Mutex
	connCount atomic.Int32

	trimCount uint64

	lastTrimMu sync.RWMutex
	lastTrim   time.Time

	refCount                sync.WaitGroup
	ctx                     context.Context
	cancel                  func()
	unregisterMemoryWatcher func()
}

var (
	_ connmgr.ConnManager = (*BasicConnMgr)(nil)
	_ connmgr.Decayer     = (*BasicConnMgr)(nil)
)

type segment struct {
	sync.Mutex
	peers map[peer.ID]*peerInfo
}

type segments struct {
	bucketsMu sync.Mutex
	buckets   [256]*segment
}

func (ss *segments) get(p peer.ID) *segment { _ = "STUB: not implemented"; return nil }

func (ss *segments) countPeers() (count int) { _ = "STUB: not implemented"; return 0 }

func (s *segment) tagInfoFor(p peer.ID, now time.Time) *peerInfo {
	_ = "STUB: not implemented"
	return nil
}

func NewConnManager(low, hi int, opts ...Option) (*BasicConnMgr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cm *BasicConnMgr) ForceTrim() { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) Close() error { _ = "STUB: not implemented"; return nil }

func (cm *BasicConnMgr) Protect(id peer.ID, tag string) { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) Unprotect(id peer.ID, tag string) (protected bool) {
	_ = "STUB: not implemented"
	return false
}

func (cm *BasicConnMgr) IsProtected(id peer.ID, tag string) (protected bool) {
	_ = "STUB: not implemented"
	return false
}

func (cm *BasicConnMgr) CheckLimit(systemLimit connmgr.GetConnLimiter) error {
	_ = "STUB: not implemented"
	return nil
}

type peerInfo struct {
	id       peer.ID
	tags     map[string]int
	decaying map[*decayingTag]*connmgr.DecayingValue

	value int
	temp  bool

	conns map[network.Conn]time.Time

	firstSeen time.Time
}

type peerInfos []*peerInfo

func (p peerInfos) SortByValueAndStreams(segments *segments, sortByMoreStreams bool) {
	_ = "STUB: not implemented"
	return
}

func (cm *BasicConnMgr) TrimOpenConns(_ context.Context) { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) background() { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) doTrim() { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) trim() { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) getConnsToCloseEmergency(target int) []network.Conn {
	_ = "STUB: not implemented"
	return nil
}

func (cm *BasicConnMgr) getConnsToClose() []network.Conn { _ = "STUB: not implemented"; return nil }

func (cm *BasicConnMgr) GetTagInfo(p peer.ID) *connmgr.TagInfo {
	_ = "STUB: not implemented"
	return nil
}

func (cm *BasicConnMgr) TagPeer(p peer.ID, tag string, val int) { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) UntagPeer(p peer.ID, tag string) { _ = "STUB: not implemented"; return }

func (cm *BasicConnMgr) UpsertTag(p peer.ID, tag string, upsert func(int) int) {
	_ = "STUB: not implemented"
	return
}

type CMInfo struct {
	LowWater int

	HighWater int

	LastTrim time.Time

	GracePeriod time.Duration

	ConnCount int
}

func (cm *BasicConnMgr) GetInfo() CMInfo { _ = "STUB: not implemented"; return *new(CMInfo) }

func (cm *BasicConnMgr) Notifee() network.Notifiee {
	_ = "STUB: not implemented"
	return *new(network.Notifiee)
}

type cmNotifee BasicConnMgr

func (nn *cmNotifee) cm() *BasicConnMgr { _ = "STUB: not implemented"; return nil }

func (nn *cmNotifee) Connected(_ network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

func (nn *cmNotifee) Disconnected(_ network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

func (nn *cmNotifee) Listen(_ network.Network, _ ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (nn *cmNotifee) ListenClose(_ network.Network, _ ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}
