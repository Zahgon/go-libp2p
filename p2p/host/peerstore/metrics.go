package peerstore

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

var LatencyEWMASmoothing = 0.1

type metrics struct {
	mutex  sync.RWMutex
	latmap map[peer.ID]time.Duration
}

func NewMetrics() *metrics { _ = "STUB: not implemented"; return nil }

func (m *metrics) RecordLatency(p peer.ID, next time.Duration) { _ = "STUB: not implemented"; return }

func (m *metrics) LatencyEWMA(p peer.ID) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (m *metrics) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
