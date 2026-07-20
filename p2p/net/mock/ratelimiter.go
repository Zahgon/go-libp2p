package mocknet

import (
	"sync"
	"time"
)

type RateLimiter struct {
	lock         sync.Mutex
	bandwidth    float64
	allowance    float64
	maxAllowance float64
	lastUpdate   time.Time
	count        int
	duration     time.Duration
}

func NewRateLimiter(bandwidth float64) *RateLimiter { _ = "STUB: not implemented"; return nil }

func (r *RateLimiter) UpdateBandwidth(bandwidth float64) { _ = "STUB: not implemented"; return }

func (r *RateLimiter) Limit(dataSize int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
