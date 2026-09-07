package basichost

import (
	ma "github.com/multiformats/go-multiaddr"

	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_host_addrs"

var (
	reachableAddrs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "reachable",
			Help:      "Number of reachable addresses by transport",
		},
		[]string{"ipv", "transport"},
	)
	unreachableAddrs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "unreachable",
			Help:      "Number of unreachable addresses by transport",
		},
		[]string{"ipv", "transport"},
	)
	unknownAddrs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "unknown",
			Help:      "Number of addresses with unknown reachability by transport",
		},
		[]string{"ipv", "transport"},
	)
	collectors = []prometheus.Collector{
		reachableAddrs,
		unreachableAddrs,
		unknownAddrs,
	}
)

type MetricsTracker interface {
	ConfirmedAddrsChanged(reachable, unreachable, unknown []ma.Multiaddr)

	ReachabilityTrackerClosed()
}

type metricsTracker struct {
	prevReachableCounts   map[metricKey]int
	prevUnreachableCounts map[metricKey]int
	prevUnknownCounts     map[metricKey]int
	currentReachable      map[metricKey]int
	currentUnreachable    map[metricKey]int
	currentUnknown        map[metricKey]int
}

var _ MetricsTracker = &metricsTracker{}

type metricsTrackerSetting struct {
	reg prometheus.Registerer
}

type metricsTrackerOption func(*metricsTrackerSetting)

func withRegisterer(reg prometheus.Registerer) metricsTrackerOption {
	_ = "STUB: not implemented"
	return *new(metricsTrackerOption)
}

type metricKey struct {
	ipv       string
	transport string
}

func newMetricsTracker(opts ...metricsTrackerOption) MetricsTracker {
	_ = "STUB: not implemented"
	return *new(MetricsTracker)
}

func (t *metricsTracker) ReachabilityTrackerClosed() { _ = "STUB: not implemented"; return }

func (t *metricsTracker) ConfirmedAddrsChanged(reachable, unreachable, unknown []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func updateMetric(metric *prometheus.GaugeVec, addrs []ma.Multiaddr, current map[metricKey]int, prev map[metricKey]int) {
	_ = "STUB: not implemented"
	return
}

func resetMetric(metric *prometheus.GaugeVec, current map[metricKey]int, prev map[metricKey]int) {
	_ = "STUB: not implemented"
	return
}
