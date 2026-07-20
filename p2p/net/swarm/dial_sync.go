package swarm

import (
	"context"
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
)

type dialWorkerFunc func(peer.ID, <-chan dialRequest)

var errConcurrentDialSuccessful = errors.New("concurrent dial successful")

func newDialSync(worker dialWorkerFunc) *dialSync { _ = "STUB: not implemented"; return nil }

type dialSync struct {
	mutex      sync.Mutex
	dials      map[peer.ID]*activeDial
	dialWorker dialWorkerFunc
}

type activeDial struct {
	refCnt int

	ctx         context.Context
	cancelCause func(error)

	reqch chan dialRequest
}

func (ad *activeDial) dial(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *dialSync) getActiveDial(p peer.ID) (*activeDial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *dialSync) Dial(ctx context.Context, p peer.ID) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
