package backoff

import (
	"math/rand"
	"sync"
	"time"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("discovery-backoff")

type BackoffFactory func() BackoffStrategy

type BackoffStrategy interface {
	Delay() time.Duration

	Reset()
}

type Jitter func(duration, min, max time.Duration, rng *rand.Rand) time.Duration

func FullJitter(duration, min, max time.Duration, rng *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func NoJitter(duration, min, max time.Duration, _ *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type randomizedBackoff struct {
	min time.Duration
	max time.Duration
	rng *rand.Rand
}

func (b *randomizedBackoff) BoundedDelay(duration time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func boundedDuration(d, min, max time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type attemptBackoff struct {
	attempt int
	jitter  Jitter
	randomizedBackoff
}

func (b *attemptBackoff) Reset() { _ = "STUB: not implemented"; return }

func NewFixedBackoff(delay time.Duration) BackoffFactory {
	_ = "STUB: not implemented"
	return *new(BackoffFactory)
}

type fixedBackoff struct {
	delay time.Duration
}

func (b *fixedBackoff) Delay() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (b *fixedBackoff) Reset() { _ = "STUB: not implemented"; return }

func NewPolynomialBackoff(min, max time.Duration, jitter Jitter,
	timeUnits time.Duration, polyCoefs []float64, rngSrc rand.Source) BackoffFactory {
	_ = "STUB: not implemented"
	return *new(BackoffFactory)
}

type polynomialBackoff struct {
	attemptBackoff
	timeUnits time.Duration
	poly      []float64
}

func (b *polynomialBackoff) Delay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func NewExponentialBackoff(min, max time.Duration, jitter Jitter,
	timeUnits time.Duration, base float64, offset time.Duration, rngSrc rand.Source) BackoffFactory {
	_ = "STUB: not implemented"
	return *new(BackoffFactory)
}

type exponentialBackoff struct {
	attemptBackoff
	timeUnits time.Duration
	base      float64
	offset    time.Duration
}

func (b *exponentialBackoff) Delay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func NewExponentialDecorrelatedJitter(min, max time.Duration, base float64, rngSrc rand.Source) BackoffFactory {
	_ = "STUB: not implemented"
	return *new(BackoffFactory)
}

type exponentialDecorrelatedJitter struct {
	randomizedBackoff
	base      float64
	lastDelay time.Duration
}

func (b *exponentialDecorrelatedJitter) Delay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *exponentialDecorrelatedJitter) Reset() { _ = "STUB: not implemented"; return }

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) { _ = "STUB: not implemented"; return 0 }

func (r *lockedSource) Seed(seed int64) { _ = "STUB: not implemented"; return }
