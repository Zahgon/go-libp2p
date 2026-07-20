package discovery

import "time"

type Option func(opts *Options) error

type Options struct {
	Ttl   time.Duration
	Limit int

	Other map[any]any
}

func (opts *Options) Apply(options ...Option) error { _ = "STUB: not implemented"; return nil }

func TTL(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func Limit(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }
