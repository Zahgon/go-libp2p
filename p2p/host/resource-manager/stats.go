package rcmgr

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_rcmgr"

var (
	conns = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "connections",
		Help:      "Number of Connections",
	}, []string{"dir", "scope"})

	connsInboundSystem     = conns.With(prometheus.Labels{"dir": "inbound", "scope": "system"})
	connsInboundTransient  = conns.With(prometheus.Labels{"dir": "inbound", "scope": "transient"})
	connsOutboundSystem    = conns.With(prometheus.Labels{"dir": "outbound", "scope": "system"})
	connsOutboundTransient = conns.With(prometheus.Labels{"dir": "outbound", "scope": "transient"})

	oneTenThenExpDistributionBuckets = []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 16, 32, 64, 128, 256,
	}

	peerConns = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "peer_connections",
		Buckets:   oneTenThenExpDistributionBuckets,
		Help:      "Number of connections this peer has",
	}, []string{"dir"})
	peerConnsInbound  = peerConns.With(prometheus.Labels{"dir": "inbound"})
	peerConnsOutbound = peerConns.With(prometheus.Labels{"dir": "outbound"})

	previousPeerConns = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "previous_peer_connections",
		Buckets:   oneTenThenExpDistributionBuckets,
		Help:      "Number of connections this peer previously had. This is used to get the current connection number per peer histogram by subtracting this from the peer_connections histogram",
	}, []string{"dir"})
	previousPeerConnsInbound  = previousPeerConns.With(prometheus.Labels{"dir": "inbound"})
	previousPeerConnsOutbound = previousPeerConns.With(prometheus.Labels{"dir": "outbound"})

	streams = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "streams",
		Help:      "Number of Streams",
	}, []string{"dir", "scope", "protocol"})

	peerStreams = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "peer_streams",
		Buckets:   oneTenThenExpDistributionBuckets,
		Help:      "Number of streams this peer has",
	}, []string{"dir"})
	peerStreamsInbound  = peerStreams.With(prometheus.Labels{"dir": "inbound"})
	peerStreamsOutbound = peerStreams.With(prometheus.Labels{"dir": "outbound"})

	previousPeerStreams = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "previous_peer_streams",
		Buckets:   oneTenThenExpDistributionBuckets,
		Help:      "Number of streams this peer has",
	}, []string{"dir"})
	previousPeerStreamsInbound  = previousPeerStreams.With(prometheus.Labels{"dir": "inbound"})
	previousPeerStreamsOutbound = previousPeerStreams.With(prometheus.Labels{"dir": "outbound"})

	memoryTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "memory",
		Help:      "Amount of memory reserved as reported to the Resource Manager",
	}, []string{"scope", "protocol"})

	peerMemory = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "peer_memory",
		Buckets:   memDistribution,
		Help:      "How many peers have reserved this bucket of memory, as reported to the Resource Manager",
	})
	previousPeerMemory = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "previous_peer_memory",
		Buckets:   memDistribution,
		Help:      "How many peers have previously reserved this bucket of memory, as reported to the Resource Manager",
	})

	connMemory = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "conn_memory",
		Buckets:   memDistribution,
		Help:      "How many conns have reserved this bucket of memory, as reported to the Resource Manager",
	})
	previousConnMemory = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: metricNamespace,
		Name:      "previous_conn_memory",
		Buckets:   memDistribution,
		Help:      "How many conns have previously reserved this bucket of memory, as reported to the Resource Manager",
	})

	fds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "fds",
		Help:      "Number of file descriptors reserved as reported to the Resource Manager",
	}, []string{"scope"})

	fdsSystem    = fds.With(prometheus.Labels{"scope": "system"})
	fdsTransient = fds.With(prometheus.Labels{"scope": "transient"})

	blockedResources = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "blocked_resources",
		Help:      "Number of blocked resources",
	}, []string{"dir", "scope", "resource"})

	limits = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "limit",
		Help:      "Resource manager limits",
	}, []string{"scope", "resource"})
)

var (
	memDistribution = []float64{
		1 << 10,
		4 << 10,
		32 << 10,
		1 << 20,
		32 << 20,
		256 << 20,
		512 << 20,
		1 << 30,
		2 << 30,
		4 << 30,
	}
)

func MustRegisterWith(reg prometheus.Registerer) { _ = "STUB: not implemented"; return }

func WithMetricsDisabled() Option { _ = "STUB: not implemented"; return *new(Option) }

type StatsTraceReporter struct{}

func NewStatsTraceReporter() (StatsTraceReporter, error) {
	_ = "STUB: not implemented"
	return *new(StatsTraceReporter), nil
}

func reportLimit(scope, resource string, value int64) { _ = "STUB: not implemented"; return }

func (r StatsTraceReporter) ReportSystemLimits(limiter Limiter) { _ = "STUB: not implemented"; return }

func (r StatsTraceReporter) ConsumeEvent(evt TraceEvt) { _ = "STUB: not implemented"; return }

func (r StatsTraceReporter) consumeEventWithLabelSlice(evt TraceEvt, tags *[]string) {
	_ = "STUB: not implemented"
	return
}
