package autonatv2

import (
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsTracer interface {
	CompletedRequest(EventDialRequestCompleted)
	ClientCompletedRequest([]Request, Result, error)
}

const metricNamespace = "libp2p_autonatv2"

var (
	requestsCompleted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "requests_completed_total",
			Help:      "Requests Completed",
		},
		[]string{"server_error", "response_status", "dial_status", "dial_data_required", "ip_or_dns_version", "transport"},
	)
	clientRequestsCompleted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "client_requests_completed_total",
			Help:      "Client Requests Completed",
		},
		[]string{"ip_or_dns_version", "transport", "addr_count", "dial_refused", "reachability"},
	)
	clientRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "client_requests_total",
			Help:      "Client Requests Total",
		},
		[]string{"outcome"},
	)
)

type metricsTracer struct {
}

func NewMetricsTracer(reg prometheus.Registerer) MetricsTracer {
	_ = "STUB: not implemented"
	return *new(MetricsTracer)
}

func (m *metricsTracer) CompletedRequest(e EventDialRequestCompleted) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) ClientCompletedRequest(reqs []Request, result Result, err error) {
	_ = "STUB: not implemented"
	return
}

func getIPOrDNSVersion(a ma.Multiaddr) string { _ = "STUB: not implemented"; return "" }

func getErrString(e error) string { _ = "STUB: not implemented"; return "" }
