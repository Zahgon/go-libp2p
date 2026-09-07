package eventbus

import (
	"reflect"

	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "libp2p_eventbus"

var (
	eventsEmitted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "events_emitted_total",
			Help:      "Events Emitted",
		},
		[]string{"event"},
	)
	totalSubscribers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "subscribers_total",
			Help:      "Number of subscribers for an event type",
		},
		[]string{"event"},
	)
	subscriberQueueLength = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "subscriber_queue_length",
			Help:      "Subscriber queue length",
		},
		[]string{"subscriber_name"},
	)
	subscriberQueueFull = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Name:      "subscriber_queue_full",
			Help:      "Subscriber Queue completely full",
		},
		[]string{"subscriber_name"},
	)
	subscriberEventQueued = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "subscriber_event_queued",
			Help:      "Event Queued for subscriber",
		},
		[]string{"subscriber_name"},
	)
	collectors = []prometheus.Collector{
		eventsEmitted,
		totalSubscribers,
		subscriberQueueLength,
		subscriberQueueFull,
		subscriberEventQueued,
	}
)

type MetricsTracer interface {
	EventEmitted(typ reflect.Type)

	AddSubscriber(typ reflect.Type)

	RemoveSubscriber(typ reflect.Type)

	SubscriberQueueLength(name string, n int)

	SubscriberQueueFull(name string, isFull bool)

	SubscriberEventQueued(name string)
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

func (m *metricsTracer) EventEmitted(typ reflect.Type) { _ = "STUB: not implemented"; return }

func (m *metricsTracer) AddSubscriber(typ reflect.Type) { _ = "STUB: not implemented"; return }

func (m *metricsTracer) RemoveSubscriber(typ reflect.Type) { _ = "STUB: not implemented"; return }

func (m *metricsTracer) SubscriberQueueLength(name string, n int) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) SubscriberQueueFull(name string, isFull bool) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsTracer) SubscriberEventQueued(name string) { _ = "STUB: not implemented"; return }
