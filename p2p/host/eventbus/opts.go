package eventbus

import (
	"log/slog"
	"sync/atomic"
)

type subSettings struct {
	buffer int
	name   string
}

var subCnt atomic.Int64

var subSettingsDefault = subSettings{
	buffer: 16,
}

func newSubSettings() subSettings { _ = "STUB: not implemented"; return *new(subSettings) }

func BufSize(n int) func(any) error { _ = "STUB: not implemented"; return nil }

func Name(name string) func(any) error { _ = "STUB: not implemented"; return nil }

type emitterSettings struct {
	makeStateful bool
}

func Stateful(s any) error { _ = "STUB: not implemented"; return nil }

type Option func(*basicBus)

func WithMetricsTracer(metricsTracer MetricsTracer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func withLogger(logger *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }
