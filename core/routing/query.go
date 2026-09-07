package routing

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
)

type QueryEventType int

var QueryEventBufferSize = 16

const (
	SendingQuery QueryEventType = iota

	PeerResponse

	FinalPeer

	QueryError

	Provider

	Value

	AddingPeer

	DialingPeer
)

type QueryEvent struct {
	ID        peer.ID
	Type      QueryEventType
	Responses []*peer.AddrInfo
	Extra     string
}

type routingQueryKey struct{}
type eventChannel struct {
	mu  sync.Mutex
	ctx context.Context
	ch  chan<- *QueryEvent
}

func (e *eventChannel) waitThenClose() { _ = "STUB: not implemented"; return }

func (e *eventChannel) send(ctx context.Context, ev *QueryEvent) { _ = "STUB: not implemented"; return }

func RegisterForQueryEvents(ctx context.Context) (context.Context, <-chan *QueryEvent) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func PublishQueryEvent(ctx context.Context, ev *QueryEvent) { _ = "STUB: not implemented"; return }

func cloneForPublish(ev *QueryEvent) *QueryEvent { _ = "STUB: not implemented"; return nil }

func SubscribesToQueryEvents(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
