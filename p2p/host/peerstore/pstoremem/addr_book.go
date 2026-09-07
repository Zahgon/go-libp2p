package pstoremem

import (
	"container/heap"
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/record"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("peerstore")

type expiringAddr struct {
	Addr   ma.Multiaddr
	TTL    time.Duration
	Expiry time.Time
	Peer   peer.ID

	heapIndex int
}

func (e *expiringAddr) ExpiredBy(t time.Time) bool { _ = "STUB: not implemented"; return false }

func (e *expiringAddr) IsConnected() bool { _ = "STUB: not implemented"; return false }

func ttlIsConnected(ttl time.Duration) bool { _ = "STUB: not implemented"; return false }

type peerRecordState struct {
	Envelope *record.Envelope

	Seq uint64
}

var _ heap.Interface = &peerAddrs{}

type peerAddrs struct {
	Addrs map[peer.ID]map[string]*expiringAddr

	expiringHeap []*expiringAddr
}

func newPeerAddrs() peerAddrs { _ = "STUB: not implemented"; return *new(peerAddrs) }

func (pa *peerAddrs) Len() int           { _ = "STUB: not implemented"; return 0 }
func (pa *peerAddrs) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pa *peerAddrs) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pa *peerAddrs) Push(x any) { _ = "STUB: not implemented"; return }

func (pa *peerAddrs) Pop() any { _ = "STUB: not implemented"; return *new(any) }

func (pa *peerAddrs) Delete(a *expiringAddr) { _ = "STUB: not implemented"; return }

func (pa *peerAddrs) FindAddr(p peer.ID, addr ma.Multiaddr) (*expiringAddr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pa *peerAddrs) NextExpiry() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (pa *peerAddrs) PopIfExpired(now time.Time) (*expiringAddr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pa *peerAddrs) Update(a *expiringAddr) { _ = "STUB: not implemented"; return }

func (pa *peerAddrs) Insert(a *expiringAddr) { _ = "STUB: not implemented"; return }

func (pa *peerAddrs) NumUnconnectedAddrs() int { _ = "STUB: not implemented"; return 0 }

type clock interface {
	Now() time.Time
}

type realclock struct{}

func (rc realclock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

const (
	defaultMaxSignedPeerRecords = 100_000
	defaultMaxUnconnectedAddrs  = 1_000_000

	defaultMaxAddrsPerPeer = 64
)

type memoryAddrBook struct {
	mu                   sync.RWMutex
	addrs                peerAddrs
	signedPeerRecords    map[peer.ID]*peerRecordState
	maxUnconnectedAddrs  int
	maxSignedPeerRecords int
	maxAddrsPerPeer      int

	refCount sync.WaitGroup
	cancel   func()

	subManager *AddrSubManager
	clock      clock
}

var _ peerstore.AddrBook = (*memoryAddrBook)(nil)
var _ peerstore.CertifiedAddrBook = (*memoryAddrBook)(nil)

func NewAddrBook(opts ...AddrBookOption) *memoryAddrBook { _ = "STUB: not implemented"; return nil }

type AddrBookOption func(book *memoryAddrBook) error

func WithClock(clock clock) AddrBookOption { _ = "STUB: not implemented"; return *new(AddrBookOption) }

func WithMaxAddresses(n int) AddrBookOption { _ = "STUB: not implemented"; return *new(AddrBookOption) }

func WithMaxSignedPeerRecords(n int) AddrBookOption {
	_ = "STUB: not implemented"
	return *new(AddrBookOption)
}

func WithMaxAddressesPerPeer(n int) AddrBookOption {
	_ = "STUB: not implemented"
	return *new(AddrBookOption)
}

func (mab *memoryAddrBook) background(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mab *memoryAddrBook) Close() error { _ = "STUB: not implemented"; return nil }

func (mab *memoryAddrBook) gc() { _ = "STUB: not implemented"; return }

func (mab *memoryAddrBook) PeersWithAddrs() peer.IDSlice {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice)
}

func (mab *memoryAddrBook) AddAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) AddAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) ConsumePeerRecord(recordEnvelope *record.Envelope, ttl time.Duration) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func prevSignedAddrs(s *peerRecordState) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (mab *memoryAddrBook) maybeDeleteSignedPeerRecordUnlocked(p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) numUnconnectedAddrsForPeerUnlocked(p peer.ID) int {
	_ = "STUB: not implemented"
	return 0
}

func (mab *memoryAddrBook) evictNearestExpiryUnconnectedForPeerUnlocked(p peer.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (mab *memoryAddrBook) addAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) addAddrsUnlocked(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) SetAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) SetAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) UpdateAddrs(p peer.ID, oldTTL time.Duration, newTTL time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (mab *memoryAddrBook) Addrs(p peer.ID) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func validAddrs(now time.Time, amap map[string]*expiringAddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (mab *memoryAddrBook) GetPeerRecord(p peer.ID) *record.Envelope {
	_ = "STUB: not implemented"
	return nil
}

func (mab *memoryAddrBook) ClearAddrs(p peer.ID) { _ = "STUB: not implemented"; return }

func (mab *memoryAddrBook) AddrStream(ctx context.Context, p peer.ID) <-chan ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

type addrSub struct {
	pubch chan ma.Multiaddr
	ctx   context.Context
}

func (s *addrSub) pubAddr(a ma.Multiaddr) { _ = "STUB: not implemented"; return }

type AddrSubManager struct {
	mu   sync.RWMutex
	subs map[peer.ID][]*addrSub
}

func NewAddrSubManager() *AddrSubManager { _ = "STUB: not implemented"; return nil }

func (mgr *AddrSubManager) removeSub(p peer.ID, s *addrSub) { _ = "STUB: not implemented"; return }

func (mgr *AddrSubManager) BroadcastAddr(p peer.ID, addr ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (mgr *AddrSubManager) AddrStream(ctx context.Context, p peer.ID, initial []ma.Multiaddr) <-chan ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}
