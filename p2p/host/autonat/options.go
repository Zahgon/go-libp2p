package autonat

import (
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
)

type config struct {
	host host.Host

	addressFunc       AddrFunc
	dialPolicy        dialPolicy
	dialer            network.Network
	forceReachability bool
	reachability      network.Reachability
	metricsTracer     MetricsTracer

	bootDelay          time.Duration
	retryInterval      time.Duration
	refreshInterval    time.Duration
	requestTimeout     time.Duration
	throttlePeerPeriod time.Duration

	dialTimeout         time.Duration
	maxPeerAddresses    int
	throttleGlobalMax   int
	throttlePeerMax     int
	throttleResetPeriod time.Duration
	throttleResetJitter time.Duration
}

var defaults = func(c *config) error {
	c.bootDelay = 15 * time.Second
	c.retryInterval = 90 * time.Second
	c.refreshInterval = 15 * time.Minute
	c.requestTimeout = 30 * time.Second
	c.throttlePeerPeriod = 90 * time.Second

	c.dialTimeout = 15 * time.Second
	c.maxPeerAddresses = 16
	c.throttleGlobalMax = 30
	c.throttlePeerMax = 3
	c.throttleResetPeriod = 1 * time.Minute
	c.throttleResetJitter = 15 * time.Second
	return nil
}

const maxRefreshInterval = 24 * time.Hour

func EnableService(dialer network.Network) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReachability(reachability network.Reachability) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func UsingAddresses(addrFunc AddrFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSchedule(retryInterval, refreshInterval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithoutStartupDelay() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithoutThrottling() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithThrottling(amount int, interval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPeerThrottling(amount int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(mt MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }
