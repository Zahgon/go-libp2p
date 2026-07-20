package connmgr

import (
	"time"

	"github.com/benbjohnson/clock"
)

type config struct {
	highWater     int
	lowWater      int
	gracePeriod   time.Duration
	silencePeriod time.Duration
	decayer       *DecayerCfg
	clock         clock.Clock
}

type Option func(*config) error

func DecayerConfig(opts *DecayerCfg) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClock(c clock.Clock) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGracePeriod(p time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSilencePeriod(p time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
