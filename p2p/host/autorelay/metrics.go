package autorelay

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_autorelay"

var (
	status = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "status",
		Help:      "relay finder active",
	})
	reservationsOpenedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservations_opened_total",
			Help:      "Reservations Opened",
		},
	)
	reservationsClosedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservations_closed_total",
			Help:      "Reservations Closed",
		},
	)
	reservationRequestsOutcomeTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "reservation_requests_outcome_total",
			Help:      "Reservation Request Outcome",
		},
		[]string{"request_type", "outcome"},
	)

	relayAddressesUpdatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "relay_addresses_updated_total",
			Help:      "Relay Addresses Updated Count",
		},
	)
	relayAddressesCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "relay_addresses_count",
			Help:      "Relay Addresses Count",
		},
	)

	candidatesCircuitV2SupportTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "candidates_circuit_v2_support_total",
			Help:      "Candidates supporting circuit v2",
		},
		[]string{"support"},
	)
	candidatesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "candidates_total",
			Help:      "Candidates Total",
		},
		[]string{"type"},
	)
	candLoopState = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "candidate_loop_state",
			Help:      "Candidate Loop State",
		},
	)

	scheduledWorkTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "scheduled_work_time",
			Help:      "Scheduled Work Times",
		},
		[]string{"work_type"},
	)

	desiredReservations = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "desired_reservations",
			Help:      "Desired Reservations",
		},
	)

	collectors = []prometheus.Collector{
		status,
		reservationsOpenedTotal,
		reservationsClosedTotal,
		reservationRequestsOutcomeTotal,
		relayAddressesUpdatedTotal,
		relayAddressesCount,
		candidatesCircuitV2SupportTotal,
		candidatesTotal,
		candLoopState,
		scheduledWorkTime,
		desiredReservations,
	}
)

type candidateLoopState int

const (
	peerSourceRateLimited candidateLoopState = iota
	waitingOnPeerChan
	waitingForTrigger
	stopped
)

type MetricsTracer interface {
	RelayFinderStatus(isActive bool)

	ReservationEnded(cnt int)
	ReservationOpened(cnt int)
	ReservationRequestFinished(isRefresh bool, err error)

	RelayAddressCount(int)
	RelayAddressUpdated()

	CandidateChecked(supportsCircuitV2 bool)
	CandidateAdded(cnt int)
	CandidateRemoved(cnt int)
	CandidateLoopState(state candidateLoopState)

	ScheduledWorkUpdated(scheduledWork *scheduledWorkTimes)

	DesiredReservations(int)
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

func (mt *metricsTracer) RelayFinderStatus(isActive bool) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ReservationEnded(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ReservationOpened(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) ReservationRequestFinished(isRefresh bool, err error) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) RelayAddressUpdated() { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) RelayAddressCount(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) CandidateChecked(supportsCircuitV2 bool) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) CandidateAdded(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) CandidateRemoved(cnt int) { _ = "STUB: not implemented"; return }

func (mt *metricsTracer) CandidateLoopState(state candidateLoopState) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) ScheduledWorkUpdated(scheduledWork *scheduledWorkTimes) {
	_ = "STUB: not implemented"
	return
}

func (mt *metricsTracer) DesiredReservations(cnt int) { _ = "STUB: not implemented"; return }

func getReservationRequestStatus(err error) string { _ = "STUB: not implemented"; return "" }

type wrappedMetricsTracer struct {
	mt MetricsTracer
}

var _ MetricsTracer = &wrappedMetricsTracer{}

func (mt *wrappedMetricsTracer) RelayFinderStatus(isActive bool) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) ReservationEnded(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) ReservationOpened(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) ReservationRequestFinished(isRefresh bool, err error) {
	_ = "STUB: not implemented"
	return
}

func (mt *wrappedMetricsTracer) RelayAddressUpdated() { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) RelayAddressCount(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) CandidateChecked(supportsCircuitV2 bool) {
	_ = "STUB: not implemented"
	return
}

func (mt *wrappedMetricsTracer) CandidateAdded(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) CandidateRemoved(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) ScheduledWorkUpdated(scheduledWork *scheduledWorkTimes) {
	_ = "STUB: not implemented"
	return
}

func (mt *wrappedMetricsTracer) DesiredReservations(cnt int) { _ = "STUB: not implemented"; return }

func (mt *wrappedMetricsTracer) CandidateLoopState(state candidateLoopState) {
	_ = "STUB: not implemented"
	return
}
