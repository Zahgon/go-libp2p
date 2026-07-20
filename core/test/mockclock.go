package test

import (
	"sync"
	"time"
)

type MockClock struct {
	mu           sync.Mutex
	now          time.Time
	timers       []*mockInstantTimer
	advanceBySem chan struct{}
}

type mockInstantTimer struct {
	c      *MockClock
	mu     sync.Mutex
	when   time.Time
	active bool
	ch     chan time.Time
}

func (t *mockInstantTimer) Ch() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (t *mockInstantTimer) Reset(d time.Time) bool { _ = "STUB: not implemented"; return false }

func (t *mockInstantTimer) Stop() bool { _ = "STUB: not implemented"; return false }

func NewMockClock() *MockClock { _ = "STUB: not implemented"; return nil }

func (c *MockClock) InstantTimer(when time.Time) *mockInstantTimer {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *MockClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *MockClock) AdvanceBy(dur time.Duration) { _ = "STUB: not implemented"; return }
