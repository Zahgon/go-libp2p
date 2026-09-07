package rcmgr

import (
	"net/netip"

	"github.com/libp2p/go-libp2p/x/rate"
)

var defaultIPv4SubnetLimits = []rate.SubnetLimit{
	{
		PrefixLength: 32,
		Limit:        rate.Limit{RPS: 0.2, Burst: 2 * defaultMaxConcurrentConns},
	},
}

var defaultIPv6SubnetLimits = []rate.SubnetLimit{
	{
		PrefixLength: 56,
		Limit:        rate.Limit{RPS: 0.2, Burst: 2 * defaultMaxConcurrentConns},
	},
	{
		PrefixLength: 48,
		Limit:        rate.Limit{RPS: 0.5, Burst: 10 * defaultMaxConcurrentConns},
	},
}

var defaultNetworkPrefixLimits = []rate.PrefixLimit{
	{
		Prefix: netip.MustParsePrefix("127.0.0.0/8"),
		Limit:  rate.Limit{},
	},
	{
		Prefix: netip.MustParsePrefix("::1/128"),
		Limit:  rate.Limit{},
	},
}

func WithConnRateLimiters(connRateLimiter *rate.Limiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func newConnRateLimiter() *rate.Limiter { _ = "STUB: not implemented"; return nil }
