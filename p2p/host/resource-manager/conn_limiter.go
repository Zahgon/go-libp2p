package rcmgr

import (
	"math"
	"net/netip"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/x/rate"
)

type ConnLimitPerSubnet struct {
	PrefixLength int

	ConnCount int
}

type NetworkPrefixLimit struct {
	Network netip.Prefix

	ConnCount int
}

var defaultMaxConcurrentConns = 8

var defaultIP4Limit = ConnLimitPerSubnet{
	ConnCount:    defaultMaxConcurrentConns,
	PrefixLength: 32,
}
var defaultIP6Limits = []ConnLimitPerSubnet{
	{
		ConnCount:    defaultMaxConcurrentConns,
		PrefixLength: 56,
	},
	{
		ConnCount:    8 * defaultMaxConcurrentConns,
		PrefixLength: 48,
	},
}

var DefaultNetworkPrefixLimitV4 = sortNetworkPrefixes([]NetworkPrefixLimit{
	{

		Network:   netip.MustParsePrefix("127.0.0.0/8"),
		ConnCount: math.MaxInt,
	},
})
var DefaultNetworkPrefixLimitV6 = sortNetworkPrefixes([]NetworkPrefixLimit{
	{

		Network:   netip.MustParsePrefix("::1/128"),
		ConnCount: math.MaxInt,
	},
})

func sortNetworkPrefixes(limits []NetworkPrefixLimit) []NetworkPrefixLimit {
	_ = "STUB: not implemented"
	return nil
}

func WithNetworkPrefixLimit(ipv4 []NetworkPrefixLimit, ipv6 []NetworkPrefixLimit) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLimitPerSubnet(ipv4 []ConnLimitPerSubnet, ipv6 []ConnLimitPerSubnet) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type connLimiter struct {
	mu sync.Mutex

	networkPrefixLimitV4    []NetworkPrefixLimit
	networkPrefixLimitV6    []NetworkPrefixLimit
	connsPerNetworkPrefixV4 []int
	connsPerNetworkPrefixV6 []int

	connLimitPerSubnetV4 []ConnLimitPerSubnet
	connLimitPerSubnetV6 []ConnLimitPerSubnet
	ip4connsPerLimit     []map[netip.Prefix]int
	ip6connsPerLimit     []map[netip.Prefix]int
}

func newConnLimiter() *connLimiter { _ = "STUB: not implemented"; return nil }

func (cl *connLimiter) addNetworkPrefixLimit(isIP6 bool, npLimit NetworkPrefixLimit) {
	_ = "STUB: not implemented"
	return
}

func (cl *connLimiter) addConn(ip netip.Addr) bool { _ = "STUB: not implemented"; return false }

func (cl *connLimiter) rmConn(ip netip.Addr) { _ = "STUB: not implemented"; return }

const handshakeDuration = 5 * time.Second

const sourceAddressRPS = float64(1.0*time.Second) / (2 * float64(handshakeDuration))

func newVerifySourceAddressRateLimiter(cl *connLimiter) *rate.Limiter {
	_ = "STUB: not implemented"
	return nil
}
