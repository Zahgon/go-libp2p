package autorelay

import (
	"context"
	"errors"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

type PeerSource func(ctx context.Context, num int) <-chan peer.AddrInfo

type config struct {
	clock      ClockWithInstantTimer
	peerSource PeerSource

	minInterval time.Duration

	minCandidates int

	maxCandidates int

	bootDelay time.Duration

	backoff time.Duration

	desiredRelays int

	maxCandidateAge  time.Duration
	setMinCandidates bool

	metricsTracer MetricsTracer
}

var defaultConfig = config{
	clock:           RealClock{},
	minCandidates:   4,
	maxCandidates:   20,
	bootDelay:       3 * time.Minute,
	backoff:         time.Hour,
	desiredRelays:   2,
	maxCandidateAge: 30 * time.Minute,
	minInterval:     30 * time.Second,
}

var (
	errAlreadyHavePeerSource = errors.New("can only use a single WithPeerSource or WithStaticRelays")
)

type Option func(*config) error

func WithStaticRelays(static []peer.AddrInfo) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPeerSource(f PeerSource) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNumRelays(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxCandidates(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMinCandidates(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBootDelay(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBackoff(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxCandidateAge(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type InstantTimer interface {
	Reset(d time.Time) bool
	Stop() bool
	Ch() <-chan time.Time
}

type ClockWithInstantTimer interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	InstantTimer(when time.Time) InstantTimer
}

type RealTimer struct{ t *time.Timer }

var _ InstantTimer = (*RealTimer)(nil)

func (t RealTimer) Ch() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (t RealTimer) Reset(d time.Time) bool { _ = "STUB: not implemented"; return false }

func (t RealTimer) Stop() bool { _ = "STUB: not implemented"; return false }

type RealClock struct{}

var _ ClockWithInstantTimer = RealClock{}

func (RealClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (RealClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (RealClock) InstantTimer(when time.Time) InstantTimer {
	_ = "STUB: not implemented"
	return *new(InstantTimer)
}

func WithClock(cl ClockWithInstantTimer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMinInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(mt MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }
