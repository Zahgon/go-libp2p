package identify

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_identify"

var (
	pushesTriggered = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "identify_pushes_triggered_total",
			Help:      "Pushes Triggered",
		},
		[]string{"trigger"},
	)
	identify = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "identify_total",
			Help:      "Identify",
		},
		[]string{"dir"},
	)
	identifyPush = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "identify_push_total",
			Help:      "Identify Push",
		},
		[]string{"dir"},
	)
	connPushSupportTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "conn_push_support_total",
			Help:      "Identify Connection Push Support",
		},
		[]string{"support"},
	)
	protocolsCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "protocols_count",
			Help:      "Protocols Count",
		},
	)
	addrsCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "addrs_count",
			Help:      "Address Count",
		},
	)
	numProtocolsReceived = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "protocols_received",
			Help:      "Number of Protocols received",
			Buckets:   buckets,
		},
	)
	numAddrsReceived = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "addrs_received",
			Help:      "Number of addrs received",
			Buckets:   buckets,
		},
	)
	collectors = []prometheus.Collector{
		pushesTriggered,
		identify,
		identifyPush,
		connPushSupportTotal,
		protocolsCount,
		addrsCount,
		numProtocolsReceived,
		numAddrsReceived,
	}

	buckets = append(
		prometheus.LinearBuckets(1, 1, 20),
		prometheus.LinearBuckets(25, 5, 16)...,
	)
)

type MetricsTracer interface {
	TriggeredPushes(event any)

	ConnPushSupport(identifyPushSupport)

	IdentifyReceived(isPush bool, numProtocols int, numAddrs int)

	IdentifySent(isPush bool, numProtocols int, numAddrs int)
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

func (t *metricsTracer) TriggeredPushes(ev any) { _ = "STUB: not implemented"; return }

func (t *metricsTracer) IncrementPushSupport(s identifyPushSupport) {
	_ = "STUB: not implemented"
	return
}

func (t *metricsTracer) IdentifySent(isPush bool, numProtocols int, numAddrs int) {
	_ = "STUB: not implemented"
	return
}

func (t *metricsTracer) IdentifyReceived(isPush bool, numProtocols int, numAddrs int) {
	_ = "STUB: not implemented"
	return
}

func (t *metricsTracer) ConnPushSupport(support identifyPushSupport) {
	_ = "STUB: not implemented"
	return
}

func getPushSupport(s identifyPushSupport) string { _ = "STUB: not implemented"; return "" }
