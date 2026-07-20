package relay

import (
	"time"

	pbv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pb"
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_relaysvc"

var (
	status = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "status",
			Help:      "Relay Status",
		},
	)

	reservationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservations_total",
			Help:      "Relay Reservation Request",
		},
		[]string{"type"},
	)
	reservationRequestResponseStatusTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservation_request_response_status_total",
			Help:      "Relay Reservation Request Response Status",
		},
		[]string{"status"},
	)
	reservationRejectionsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservation_rejections_total",
			Help:      "Relay Reservation Rejected Reason",
		},
		[]string{"reason"},
	)

	connectionsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "connections_total",
			Help:      "Relay Connection Total",
		},
		[]string{"type"},
	)
	connectionRequestResponseStatusTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "connection_request_response_status_total",
			Help:      "Relay Connection Request Status",
		},
		[]string{"status"},
	)
	connectionRejectionsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "connection_rejections_total",
			Help:      "Relay Connection Rejected Reason",
		},
		[]string{"reason"},
	)
	connectionDurationSeconds = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: metricNamespace,
			Name:      "connection_duration_seconds",
			Help:      "Relay Connection Duration",
		},
	)

	dataTransferredBytesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "data_transferred_bytes_total",
			Help:      "Bytes Transferred Total",
		},
	)

	collectors = []prometheus.Collector{
		status,
		reservationsTotal,
		reservationRequestResponseStatusTotal,
		reservationRejectionsTotal,
		connectionsTotal,
		connectionRequestResponseStatusTotal,
		connectionRejectionsTotal,
		connectionDurationSeconds,
		dataTransferredBytesTotal,
	}
)

const (
	requestStatusOK       = "ok"
	requestStatusRejected = "rejected"
	requestStatusError    = "error"
)

type MetricsTracer interface {
	RelayStatus(enabled bool)

	ConnectionOpened()

	ConnectionClosed(d time.Duration)

	ConnectionRequestHandled(status pbv2.Status)

	ReservationAllowed(isRenewal bool)

	ReservationClosed(cnt int)

	ReservationRequestHandled(status pbv2.Status)

	BytesTransferred(cnt int)
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

func (mt *metricsTracer) RelayStatus(enabled bool) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ConnectionOpened() { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ConnectionClosed(d time.Duration) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ConnectionRequestHandled(status pbv2.Status) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) ReservationAllowed(isRenewal bool) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ReservationClosed(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ReservationRequestHandled(status pbv2.Status) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) BytesTransferred(cnt int) { _ = "STUB: not implemented"; return }

func getResponseStatus(status pbv2.Status) string { _ = "STUB: not implemented"; return "" }

func getRejectionReason(status pbv2.Status) string { _ = "STUB: not implemented"; return "" }
