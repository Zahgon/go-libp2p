package connmgr

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/benbjohnson/clock"
)

var DefaultResolution = 1 * time.Minute

type bumpCmd struct {
	peer  peer.ID
	tag   *decayingTag
	delta int
}

type removeCmd struct {
	peer peer.ID
	tag  *decayingTag
}

type decayer struct {
	cfg   *DecayerCfg
	mgr   *BasicConnMgr
	clock clock.Clock

	tagsMu    sync.Mutex
	knownTags map[string]*decayingTag

	lastTick atomic.Pointer[time.Time]

	bumpTagCh   chan bumpCmd
	removeTagCh chan removeCmd
	closeTagCh  chan *decayingTag

	closeCh chan struct{}
	doneCh  chan struct{}
	err     error
}

var _ connmgr.Decayer = (*decayer)(nil)

type DecayerCfg struct {
	Resolution time.Duration
	Clock      clock.Clock
}

func (cfg *DecayerCfg) WithDefaults() *DecayerCfg { _ = "STUB: not implemented"; return nil }

func NewDecayer(cfg *DecayerCfg, mgr *BasicConnMgr) (*decayer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *decayer) RegisterDecayingTag(name string, interval time.Duration, decayFn connmgr.DecayFn, bumpFn connmgr.BumpFn) (connmgr.DecayingTag, error) {
	_ = "STUB: not implemented"
	return *new(connmgr.DecayingTag), nil
}

func (d *decayer) Close() error { _ = "STUB: not implemented"; return nil }

func (d *decayer) process() { _ = "STUB: not implemented"; return }

type decayingTag struct {
	trkr     *decayer
	name     string
	interval time.Duration
	nextTick time.Time
	decayFn  connmgr.DecayFn
	bumpFn   connmgr.BumpFn

	closed atomic.Bool
}

var _ connmgr.DecayingTag = (*decayingTag)(nil)

func (t *decayingTag) Name() string { _ = "STUB: not implemented"; return "" }

func (t *decayingTag) Interval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (t *decayingTag) Bump(p peer.ID, delta int) error { _ = "STUB: not implemented"; return nil }

func (t *decayingTag) Remove(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (t *decayingTag) Close() error { _ = "STUB: not implemented"; return nil }
