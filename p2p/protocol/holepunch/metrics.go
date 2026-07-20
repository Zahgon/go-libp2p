package holepunch

import (
	"github.com/libp2p/go-libp2p/core/network"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_holepunch"

var (
	directDialsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "direct_dials_total",
			Help:      "Direct Dials Total",
		},
		[]string{"outcome"},
	)
	hpAddressOutcomesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "address_outcomes_total",
			Help:      "Hole Punch outcomes by Transport",
		},
		[]string{"side", "num_attempts", "ipv", "transport", "outcome"},
	)
	hpOutcomesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "outcomes_total",
			Help:      "Hole Punch outcomes overall",
		},
		[]string{"side", "num_attempts", "outcome"},
	)

	collectors = []prometheus.Collector{
		directDialsTotal,
		hpAddressOutcomesTotal,
		hpOutcomesTotal,
	}
)

type MetricsTracer interface {
	HolePunchFinished(side string, attemptNum int, theirAddrs []ma.Multiaddr, ourAddr []ma.Multiaddr, directConn network.ConnMultiaddrs)
	DirectDialFinished(success bool)
}

type metricsTracer struct{}

var _ MetricsTracer = &metricsTracer{}

type metricsTracerSetting struct {
	reg prometheus.Registerer
}

type MetricsTracerOption func(*metricsTracerSetting)

func WithRegisterer(reg prometheus.Registerer) MetricsTracerOption {
	_ = "STUB: not implemented"
	return *new(MetricsTracerOption)
}

func NewMetricsTracer(opts ...MetricsTracerOption) MetricsTracer {
	_ = "STUB: not implemented"
	return *new(MetricsTracer)
}

func (mt *metricsTracer) HolePunchFinished(side string, numAttempts int,
	remoteAddrs []ma.Multiaddr, localAddrs []ma.Multiaddr, directConn network.ConnMultiaddrs) {
	_ = "STUB: not implemented"
	return
}

func getNumAttemptString(numAttempt int) string { _ = "STUB: not implemented"; return "" }

func (mt *metricsTracer) DirectDialFinished(success bool) { _ = "STUB: not implemented"; return }
