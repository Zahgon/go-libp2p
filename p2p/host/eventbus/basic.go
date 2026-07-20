package eventbus

import (
	"log/slog"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("eventbus")

const slowConsumerWarningTimeout = time.Second

type basicBus struct {
	lk            sync.RWMutex
	nodes         map[reflect.Type]*node
	wildcard      *wildcardNode
	metricsTracer MetricsTracer
	log           *slog.Logger
}

var _ event.Bus = (*basicBus)(nil)

type emitter struct {
	n             *node
	w             *wildcardNode
	typ           reflect.Type
	closed        atomic.Bool
	dropper       func(reflect.Type)
	metricsTracer MetricsTracer
}

func (e *emitter) Emit(evt any) error { _ = "STUB: not implemented"; return nil }

func (e *emitter) Close() error { _ = "STUB: not implemented"; return nil }

func NewBus(opts ...Option) event.Bus { _ = "STUB: not implemented"; return *new(event.Bus) }

func (b *basicBus) withNode(typ reflect.Type, cb func(*node), async func(*node)) {
	_ = "STUB: not implemented"
	return
}

func (b *basicBus) tryDropNode(typ reflect.Type) { _ = "STUB: not implemented"; return }

type wildcardSub struct {
	ch            chan any
	w             *wildcardNode
	metricsTracer MetricsTracer
	name          string
	closeOnce     sync.Once
}

func (w *wildcardSub) Out() <-chan any { _ = "STUB: not implemented"; return nil }

func (w *wildcardSub) Close() error { _ = "STUB: not implemented"; return nil }

func (w *wildcardSub) Name() string { _ = "STUB: not implemented"; return "" }

type namedSink struct {
	name string
	ch   chan any
}

type sub struct {
	ch            chan any
	nodes         []*node
	dropper       func(reflect.Type)
	metricsTracer MetricsTracer
	name          string
	closeOnce     sync.Once
}

func (s *sub) Name() string { _ = "STUB: not implemented"; return "" }

func (s *sub) Out() <-chan any { _ = "STUB: not implemented"; return nil }

func (s *sub) Close() error { _ = "STUB: not implemented"; return nil }

var _ event.Subscription = (*sub)(nil)

func (b *basicBus) Subscribe(evtTypes any, opts ...event.SubscriptionOpt) (_ event.Subscription, err error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

func (b *basicBus) Emitter(evtType any, opts ...event.EmitterOpt) (e event.Emitter, err error) {
	_ = "STUB: not implemented"
	return *new(event.Emitter), nil
}

func (b *basicBus) GetAllEventTypes() []reflect.Type { _ = "STUB: not implemented"; return nil }

type wildcardNode struct {
	sync.RWMutex
	nSinks        atomic.Int32
	sinks         []*namedSink
	metricsTracer MetricsTracer
	log           *slog.Logger
}

func (n *wildcardNode) addSink(sink *namedSink) { _ = "STUB: not implemented"; return }

func (n *wildcardNode) removeSink(ch chan any) { _ = "STUB: not implemented"; return }

var wildcardType = reflect.TypeOf(event.WildcardSubscription)

func (n *wildcardNode) emit(evt any) { _ = "STUB: not implemented"; return }

type node struct {
	lk sync.Mutex

	typ reflect.Type

	nEmitters atomic.Int32

	keepLast bool
	last     any

	sinks         []*namedSink
	metricsTracer MetricsTracer
	log           *slog.Logger
}

func newNode(typ reflect.Type, metricsTracer MetricsTracer, log *slog.Logger) *node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) emit(evt any) { _ = "STUB: not implemented"; return }

func emitAndLogError(log *slog.Logger, typ reflect.Type, evt any, sink *namedSink) {
	_ = "STUB: not implemented"
	return
}

func sendSubscriberMetrics(metricsTracer MetricsTracer, sink *namedSink) {
	_ = "STUB: not implemented"
	return
}
