package autonat

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/host/autonat/pb"
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_autonat"

var (
	reachabilityStatus = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "reachability_status",
			Help:      "Current node reachability",
		},
	)
	reachabilityStatusConfidence = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "reachability_status_confidence",
			Help:      "Node reachability status confidence",
		},
	)
	receivedDialResponseTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "received_dial_response_total",
			Help:      "Count of dial responses for client",
		},
		[]string{"response_status"},
	)
	outgoingDialResponseTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "outgoing_dial_response_total",
			Help:      "Count of dial responses for server",
		},
		[]string{"response_status"},
	)
	outgoingDialRefusedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "outgoing_dial_refused_total",
			Help:      "Count of dial requests refused by server",
		},
		[]string{"refusal_reason"},
	)
	nextProbeTimestamp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "next_probe_timestamp",
			Help:      "Time of next probe",
		},
	)
	collectors = []prometheus.Collector{
		reachabilityStatus,
		reachabilityStatusConfidence,
		receivedDialResponseTotal,
		outgoingDialResponseTotal,
		outgoingDialRefusedTotal,
		nextProbeTimestamp,
	}
)

type MetricsTracer interface {
	ReachabilityStatus(status network.Reachability)
	ReachabilityStatusConfidence(confidence int)
	ReceivedDialResponse(status pb.Message_ResponseStatus)
	OutgoingDialResponse(status pb.Message_ResponseStatus)
	OutgoingDialRefused(reason string)
	NextProbeTime(t time.Time)
}

func getResponseStatus(status pb.Message_ResponseStatus) string {
	_ = "STUB: not implemented"
	return ""
}

const (
	rate_limited     = "rate limited"
	dial_blocked     = "dial blocked"
	no_valid_address = "no valid address"
)

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

func (mt *metricsTracer) ReachabilityStatus(status network.Reachability) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) ReachabilityStatusConfidence(confidence int) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) ReceivedDialResponse(status pb.Message_ResponseStatus) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) OutgoingDialResponse(status pb.Message_ResponseStatus) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) OutgoingDialRefused(reason string) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) NextProbeTime(t time.Time) { _ = "STUB: not implemented"; return }
