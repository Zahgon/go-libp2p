package swarm

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

type dialJob struct {
	addr    ma.Multiaddr
	peer    peer.ID
	ctx     context.Context
	resp    chan transport.DialUpdate
	timeout time.Duration
}

func (dj *dialJob) cancelled() bool { _ = "STUB: not implemented"; return false }

type dialLimiter struct {
	lk sync.Mutex

	fdConsuming int
	fdLimit     int
	waitingOnFd []*dialJob

	dialFunc dialfunc

	activePerPeer      map[peer.ID]int
	perPeerLimit       int
	waitingOnPeerLimit map[peer.ID][]*dialJob
}

type dialfunc func(context.Context, peer.ID, ma.Multiaddr, chan<- transport.DialUpdate) (transport.CapableConn, error)

func newDialLimiter(df dialfunc) *dialLimiter { _ = "STUB: not implemented"; return nil }

func newDialLimiterWithParams(df dialfunc, fdLimit, perPeerLimit int) *dialLimiter {
	_ = "STUB: not implemented"
	return nil
}

func (dl *dialLimiter) freeFDToken() { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) freePeerToken(dj *dialJob) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) finishedDial(dj *dialJob) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) shouldConsumeFd(addr ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}

func (dl *dialLimiter) addCheckFdLimit(dj *dialJob) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) addCheckPeerLimit(dj *dialJob) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) AddDialJob(dj *dialJob) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) clearAllPeerDials(p peer.ID) { _ = "STUB: not implemented"; return }

func (dl *dialLimiter) executeDial(j *dialJob) { _ = "STUB: not implemented"; return }
