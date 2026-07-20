package holepunch

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

const (
	tracerGCInterval    = 2 * time.Minute
	tracerCacheDuration = 5 * time.Minute
)

func WithTracer(et EventTracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(mt MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsAndEventTracer(mt MetricsTracer, et EventTracer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type tracer struct {
	et   EventTracer
	mt   MetricsTracer
	self peer.ID

	refCount  sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc

	mutex sync.Mutex
	peers map[peer.ID]peerInfo
}

type peerInfo struct {
	counter int
	last    time.Time
}

type EventTracer interface {
	Trace(evt *Event)
}

type Event struct {
	Timestamp int64
	Peer      peer.ID
	Remote    peer.ID
	Type      string
	Evt       any
}

const (
	DirectDialEvtT       = "DirectDial"
	ProtocolErrorEvtT    = "ProtocolError"
	StartHolePunchEvtT   = "StartHolePunch"
	EndHolePunchEvtT     = "EndHolePunch"
	HolePunchAttemptEvtT = "HolePunchAttempt"
)

type DirectDialEvt struct {
	Success      bool
	EllapsedTime time.Duration
	Error        string `json:",omitempty"`
}

type ProtocolErrorEvt struct {
	Error string
}

type StartHolePunchEvt struct {
	RemoteAddrs []string
	RTT         time.Duration
}

type EndHolePunchEvt struct {
	Success      bool
	EllapsedTime time.Duration
	Error        string `json:",omitempty"`
}

type HolePunchAttemptEvt struct {
	Attempt int
}

func (t *tracer) DirectDialSuccessful(p peer.ID, dt time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *tracer) DirectDialFailed(p peer.ID, dt time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

func (t *tracer) ProtocolError(p peer.ID, err error) { _ = "STUB: not implemented"; return }

func (t *tracer) StartHolePunch(p peer.ID, obsAddrs []ma.Multiaddr, rtt time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *tracer) EndHolePunch(p peer.ID, dt time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

func (t *tracer) HolePunchFinished(side string, numAttempts int, theirAddrs []ma.Multiaddr, ourAddrs []ma.Multiaddr, directConn network.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *tracer) HolePunchAttempt(p peer.ID) { _ = "STUB: not implemented"; return }

func (t *tracer) gc() { _ = "STUB: not implemented"; return }

func (t *tracer) Start() { _ = "STUB: not implemented"; return }

func (t *tracer) Close() error { _ = "STUB: not implemented"; return nil }
