package swarm

import "time"

type InstantTimer interface {
	Reset(d time.Time) bool
	Stop() bool
	Ch() <-chan time.Time
}

type Clock interface {
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

var _ Clock = RealClock{}

func (RealClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (RealClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (RealClock) InstantTimer(when time.Time) InstantTimer {
	_ = "STUB: not implemented"
	return *new(InstantTimer)
}
