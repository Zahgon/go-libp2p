package rate

import (
	"container/heap"
	"net/netip"
	"slices"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"golang.org/x/time/rate"
)

type Limit struct {
	RPS float64

	Burst int
}

type PrefixLimit struct {
	Prefix netip.Prefix
	Limit
}

type SubnetLimit struct {
	PrefixLength int
	Limit
}

type Limiter struct {
	NetworkPrefixLimits []PrefixLimit

	GlobalLimit Limit

	SubnetRateLimiter SubnetLimiter

	initOnce             sync.Once
	globalBucket         *rate.Limiter
	networkPrefixBuckets []*rate.Limiter
}

func (r *Limiter) init() {
	r.initOnce.Do(func() {
		if r.GlobalLimit.RPS == 0 {
			r.globalBucket = rate.NewLimiter(rate.Inf, 0)
		} else {
			r.globalBucket = rate.NewLimiter(rate.Limit(r.GlobalLimit.RPS), r.GlobalLimit.Burst)
		}

		r.NetworkPrefixLimits = slices.Clone(r.NetworkPrefixLimits)

		slices.SortFunc(r.NetworkPrefixLimits, func(a, b PrefixLimit) int { return b.Prefix.Bits() - a.Prefix.Bits() })
		r.networkPrefixBuckets = make([]*rate.Limiter, 0, len(r.NetworkPrefixLimits))
		for _, limit := range r.NetworkPrefixLimits {
			if limit.RPS == 0 {
				r.networkPrefixBuckets = append(r.networkPrefixBuckets, rate.NewLimiter(rate.Inf, 0))
			} else {
				r.networkPrefixBuckets = append(r.networkPrefixBuckets, rate.NewLimiter(rate.Limit(limit.RPS), limit.Burst))
			}
		}
	})
}

func (r *Limiter) Limit(f func(s network.Stream)) func(s network.Stream) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Limiter) Allow(ipAddr netip.Addr) bool { _ = "STUB: not implemented"; return false }

type SubnetLimiter struct {
	IPv4SubnetLimits []SubnetLimit

	IPv6SubnetLimits []SubnetLimit

	GracePeriod time.Duration

	initOnce  sync.Once
	mx        sync.Mutex
	ipv4Heaps []*bucketHeap
	ipv6Heaps []*bucketHeap
}

func (s *SubnetLimiter) init() {
	s.initOnce.Do(func() {

		slices.SortFunc(s.IPv4SubnetLimits, func(a, b SubnetLimit) int { return b.PrefixLength - a.PrefixLength })
		slices.SortFunc(s.IPv6SubnetLimits, func(a, b SubnetLimit) int { return b.PrefixLength - a.PrefixLength })

		s.ipv4Heaps = make([]*bucketHeap, len(s.IPv4SubnetLimits))
		for i := range s.IPv4SubnetLimits {
			s.ipv4Heaps[i] = &bucketHeap{
				prefixBucket:  make([]prefixBucketWithExpiry, 0),
				prefixToIndex: make(map[netip.Prefix]int),
			}
			heap.Init(s.ipv4Heaps[i])
		}

		s.ipv6Heaps = make([]*bucketHeap, len(s.IPv6SubnetLimits))
		for i := range s.IPv6SubnetLimits {
			s.ipv6Heaps[i] = &bucketHeap{
				prefixBucket:  make([]prefixBucketWithExpiry, 0),
				prefixToIndex: make(map[netip.Prefix]int),
			}
			heap.Init(s.ipv6Heaps[i])
		}
	})
}

func (s *SubnetLimiter) Allow(ipAddr netip.Addr, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SubnetLimiter) cleanUp(now time.Time) { _ = "STUB: not implemented"; return }

type tokenBucket struct {
	*rate.Limiter
}

func (b *tokenBucket) FullAt(now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

type prefixBucketWithExpiry struct {
	tokenBucket
	Prefix netip.Prefix
	Expiry time.Time
}

type bucketHeap struct {
	prefixBucket  []prefixBucketWithExpiry
	prefixToIndex map[netip.Prefix]int
}

var _ heap.Interface = (*bucketHeap)(nil)

func (h *bucketHeap) Upsert(b prefixBucketWithExpiry) { _ = "STUB: not implemented"; return }

func (h *bucketHeap) Get(prefix netip.Prefix) prefixBucketWithExpiry {
	_ = "STUB: not implemented"
	return *new(prefixBucketWithExpiry)
}

func (h *bucketHeap) Expire(expiry time.Time) { _ = "STUB: not implemented"; return }

func (h *bucketHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *bucketHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (h *bucketHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h *bucketHeap) Push(x any) { _ = "STUB: not implemented"; return }

func (h *bucketHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }
