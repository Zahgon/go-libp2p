package upgrader

import (
	"sync"
)

func newThreshold(cutoff int) *threshold { _ = "STUB: not implemented"; return nil }

type threshold struct {
	mu   sync.Mutex
	cond sync.Cond

	count     int
	threshold int
}

func (t *threshold) Acquire() { _ = "STUB: not implemented"; return }

func (t *threshold) Release() { _ = "STUB: not implemented"; return }

func (t *threshold) Wait() { _ = "STUB: not implemented"; return }
