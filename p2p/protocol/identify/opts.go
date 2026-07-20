package identify

import "time"

type config struct {
	protocolVersion         string
	userAgent               string
	disableSignedPeerRecord bool
	metricsTracer           MetricsTracer
	timeout                 time.Duration
}

type Option func(*config)

func ProtocolVersion(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

func UserAgent(ua string) Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableSignedPeerRecord() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(tr MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
