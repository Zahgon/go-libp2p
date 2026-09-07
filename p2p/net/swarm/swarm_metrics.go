package swarm

import (
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"

	ma "github.com/multiformats/go-multiaddr"

	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_swarm"

var (
	connsOpened = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "connections_opened_total",
			Help:      "Connections Opened",
		},
		[]string{"dir", "transport", "security", "muxer", "early_muxer", "ip_version"},
	)
	keyTypes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "key_types_total",
			Help:      "key type",
		},
		[]string{"dir", "key_type"},
	)
	connsClosed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "connections_closed_total",
			Help:      "Connections Closed",
		},
		[]string{"dir", "transport", "security", "muxer", "early_muxer", "ip_version"},
	)
	dialError = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "dial_errors_total",
			Help:      "Dial Error",
		},
		[]string{"transport", "error", "ip_version"},
	)
	connDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "connection_duration_seconds",
			Help:      "Duration of a Connection",
			Buckets:   prometheus.ExponentialBuckets(1.0/16, 2, 25),
		},
		[]string{"dir", "transport", "security", "muxer", "early_muxer", "ip_version"},
	)
	connHandshakeLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "handshake_latency_seconds",
			Help:      "Duration of the libp2p Handshake",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.3, 35),
		},
		[]string{"transport", "security", "muxer", "early_muxer", "ip_version"},
	)
	dialsPerPeer = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "dials_per_peer_total",
			Help:      "Number of addresses dialed per peer",
		},
		[]string{"outcome", "num_dials"},
	)
	dialLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "dial_latency_seconds",
			Help:      "time taken to establish connection with the peer",
			Buckets:   []float64{0.001, 0.01, 0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.75, 1, 2},
		},
		[]string{"outcome", "num_dials"},
	)
	dialRankingDelay = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "dial_ranking_delay_seconds",
			Help:      "delay introduced by the dial ranking logic",
			Buckets:   []float64{0.001, 0.01, 0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.75, 1, 2},
		},
	)
	blackHoleSuccessCounterState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "black_hole_filter_state",
			Help:      "State of the black hole filter",
		},
		[]string{"name"},
	)
	blackHoleSuccessCounterSuccessFraction = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "black_hole_filter_success_fraction",
			Help:      "Fraction of successful dials among the last n requests",
		},
		[]string{"name"},
	)
	blackHoleSuccessCounterNextRequestAllowedAfter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "black_hole_filter_next_request_allowed_after",
			Help:      "Number of requests after which the next request will be allowed",
		},
		[]string{"name"},
	)
	collectors = []prometheus.Collector{
		connsOpened,
		keyTypes,
		connsClosed,
		dialError,
		connDuration,
		connHandshakeLatency,
		dialsPerPeer,
		dialRankingDelay,
		dialLatency,
		blackHoleSuccessCounterSuccessFraction,
		blackHoleSuccessCounterState,
		blackHoleSuccessCounterNextRequestAllowedAfter,
	}
)

type MetricsTracer interface {
	OpenedConnection(network.Direction, crypto.PubKey, network.ConnectionState, ma.Multiaddr)
	ClosedConnection(network.Direction, time.Duration, network.ConnectionState, ma.Multiaddr)
	CompletedHandshake(time.Duration, network.ConnectionState, ma.Multiaddr)
	FailedDialing(ma.Multiaddr, error, error)
	DialCompleted(success bool, totalDials int, latency time.Duration)
	DialRankingDelay(d time.Duration)
	UpdatedBlackHoleSuccessCounter(name string, state BlackHoleState, nextProbeAfter int, successFraction float64)
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

func appendConnectionState(tags []string, cs network.ConnectionState) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m *metricsTracer) OpenedConnection(dir network.Direction, p crypto.PubKey, cs network.ConnectionState, laddr ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) ClosedConnection(dir network.Direction, duration time.Duration, cs network.ConnectionState, laddr ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) CompletedHandshake(t time.Duration, cs network.ConnectionState, laddr ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) FailedDialing(addr ma.Multiaddr, dialErr error, cause error) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) DialCompleted(success bool, totalDials int, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) DialRankingDelay(d time.Duration) { _ = "STUB: not implemented"; return }

func (m *metricsTracer) UpdatedBlackHoleSuccessCounter(name string, state BlackHoleState,
	nextProbeAfter int, successFraction float64) {
	_ = "STUB: not implemented"
	return
}
