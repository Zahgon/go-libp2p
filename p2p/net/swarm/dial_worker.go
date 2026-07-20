package swarm

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

type dialRequest struct {
	ctx context.Context

	resch chan dialResponse
}

type dialResponse struct {
	conn *Conn

	err error
}

type pendRequest struct {
	req dialRequest

	err *DialError

	addrs map[string]struct{}
}

type addrDial struct {
	addr ma.Multiaddr

	ctx context.Context

	conn *Conn

	err error

	dialed bool

	createdAt time.Time

	dialRankingDelay time.Duration

	expectedTCPUpgradeTime time.Time
}

type dialWorker struct {
	s    *Swarm
	peer peer.ID

	reqch <-chan dialRequest

	pendingRequests map[*pendRequest]struct{}

	trackedDials map[string]*addrDial

	resch chan tpt.DialUpdate

	connected bool

	wg sync.WaitGroup
	cl Clock
}

func newDialWorker(s *Swarm, p peer.ID, reqch <-chan dialRequest, cl Clock) *dialWorker {
	_ = "STUB: not implemented"
	return nil
}

func (w *dialWorker) loop() { _ = "STUB: not implemented"; return }

func (w *dialWorker) dispatchError(ad *addrDial, err error) { _ = "STUB: not implemented"; return }

func (w *dialWorker) rankAddrs(addrs []ma.Multiaddr, isSimConnect bool) []network.AddrDelay {
	_ = "STUB: not implemented"
	return nil
}

type dialQueue struct {
	q []network.AddrDelay
}

func newDialQueue() *dialQueue { _ = "STUB: not implemented"; return nil }

func (dq *dialQueue) Add(adelay network.AddrDelay) { _ = "STUB: not implemented"; return }

func (dq *dialQueue) UpdateOrAdd(adelay network.AddrDelay) { _ = "STUB: not implemented"; return }

func (dq *dialQueue) NextBatch() []network.AddrDelay { _ = "STUB: not implemented"; return nil }

func (dq *dialQueue) top() network.AddrDelay {
	_ = "STUB: not implemented"
	return *new(network.AddrDelay)
}

func (dq *dialQueue) Len() int { _ = "STUB: not implemented"; return 0 }
