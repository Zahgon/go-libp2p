package swarm

import (
	"sync"

	ma "github.com/multiformats/go-multiaddr"
)

type BlackHoleState int

const (
	blackHoleStateProbing BlackHoleState = iota
	blackHoleStateAllowed
	blackHoleStateBlocked
)

func (st BlackHoleState) String() string { _ = "STUB: not implemented"; return "" }

type BlackHoleSuccessCounter struct {
	N int

	MinSuccesses int

	Name string

	mu sync.Mutex

	requests int

	dialResults []bool

	successes int

	state BlackHoleState
}

func (b *BlackHoleSuccessCounter) RecordResult(success bool) { _ = "STUB: not implemented"; return }

func (b *BlackHoleSuccessCounter) HandleRequest() BlackHoleState {
	_ = "STUB: not implemented"
	return *new(BlackHoleState)
}

func (b *BlackHoleSuccessCounter) reset() { _ = "STUB: not implemented"; return }

func (b *BlackHoleSuccessCounter) updateState() { _ = "STUB: not implemented"; return }

func (b *BlackHoleSuccessCounter) State() BlackHoleState {
	_ = "STUB: not implemented"
	return *new(BlackHoleState)
}

type blackHoleInfo struct {
	name            string
	state           BlackHoleState
	nextProbeAfter  int
	successFraction float64
}

func (b *BlackHoleSuccessCounter) info() blackHoleInfo {
	_ = "STUB: not implemented"
	return *new(blackHoleInfo)
}

type blackHoleDetector struct {
	udp, ipv6 *BlackHoleSuccessCounter
	mt        MetricsTracer
	readOnly  bool
}

func (d *blackHoleDetector) FilterAddrs(addrs []ma.Multiaddr) (valid []ma.Multiaddr, blackHoled []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *blackHoleDetector) RecordResult(addr ma.Multiaddr, success bool) {
	_ = "STUB: not implemented"
	return
}

func (d *blackHoleDetector) getFilterState(f *BlackHoleSuccessCounter) BlackHoleState {
	_ = "STUB: not implemented"
	return *new(BlackHoleState)
}

func (d *blackHoleDetector) trackMetrics(f *BlackHoleSuccessCounter) {
	_ = "STUB: not implemented"
	return
}
