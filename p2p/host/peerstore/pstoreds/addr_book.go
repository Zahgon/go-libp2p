package pstoreds

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/record"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoreds/pb"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoremem"

	ds "github.com/ipfs/go-datastore"
	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

type ttlWriteMode int

const (
	ttlOverride ttlWriteMode = iota
	ttlExtend
)

var (
	log = logging.Logger("peerstore/ds")

	addrBookBase = ds.NewKey("/peers/addrs")
)

type addrsRecord struct {
	sync.RWMutex
	*pb.AddrBookRecord
	dirty bool
}

func (r *addrsRecord) flush(write ds.Write) (err error) { _ = "STUB: not implemented"; return nil }

func (r *addrsRecord) clean(now time.Time) (chgd bool) { _ = "STUB: not implemented"; return false }

func (r *addrsRecord) hasExpiredAddrs(now int64) bool { _ = "STUB: not implemented"; return false }

func removeExpired(entries []*pb.AddrBookRecord_AddrEntry, now int64) []*pb.AddrBookRecord_AddrEntry {
	_ = "STUB: not implemented"
	return nil
}

type dsAddrBook struct {
	ctx  context.Context
	opts Options

	cache       cache[peer.ID, *addrsRecord]
	ds          ds.Batching
	gc          *dsAddrBookGc
	subsManager *pstoremem.AddrSubManager

	childrenDone sync.WaitGroup
	cancelFn     func()

	clock clock
}

type clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}

type realclock struct{}

func (rc realclock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (rc realclock) After(d time.Duration) <-chan time.Time { _ = "STUB: not implemented"; return nil }

var _ pstore.AddrBook = (*dsAddrBook)(nil)
var _ pstore.CertifiedAddrBook = (*dsAddrBook)(nil)

func NewAddrBook(ctx context.Context, store ds.Batching, opts Options) (ab *dsAddrBook, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ab *dsAddrBook) Close() error { _ = "STUB: not implemented"; return nil }

func (ab *dsAddrBook) loadRecord(id peer.ID, cache bool, update bool) (pr *addrsRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ab *dsAddrBook) AddAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ab *dsAddrBook) AddAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ab *dsAddrBook) ConsumePeerRecord(recordEnvelope *record.Envelope, ttl time.Duration) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ab *dsAddrBook) supersededSignedAddrs(p peer.ID, newAddrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func ttlIsConnected(ttl time.Duration) bool { _ = "STUB: not implemented"; return false }

func (ab *dsAddrBook) latestPeerRecordSeq(p peer.ID) uint64 { _ = "STUB: not implemented"; return 0 }

func (ab *dsAddrBook) storeSignedPeerRecord(p peer.ID, envelope *record.Envelope, rec *peer.PeerRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (ab *dsAddrBook) GetPeerRecord(p peer.ID) *record.Envelope {
	_ = "STUB: not implemented"
	return nil
}

func (ab *dsAddrBook) SetAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ab *dsAddrBook) SetAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ab *dsAddrBook) UpdateAddrs(p peer.ID, oldTTL time.Duration, newTTL time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ab *dsAddrBook) Addrs(p peer.ID) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (ab *dsAddrBook) PeersWithAddrs() peer.IDSlice {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice)
}

func (ab *dsAddrBook) AddrStream(ctx context.Context, p peer.ID) <-chan ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (ab *dsAddrBook) ClearAddrs(p peer.ID) { _ = "STUB: not implemented"; return }

func (ab *dsAddrBook) setAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration, mode ttlWriteMode, _ bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func deleteInPlace(s []*pb.AddrBookRecord_AddrEntry, addrs []ma.Multiaddr) []*pb.AddrBookRecord_AddrEntry {
	_ = "STUB: not implemented"
	return nil
}

func (ab *dsAddrBook) deleteAddrs(p peer.ID, addrs []ma.Multiaddr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func cleanAddrs(addrs []ma.Multiaddr, pid peer.ID) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}
