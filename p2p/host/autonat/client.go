package autonat

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/autonat/pb"
)

func NewAutoNATClient(h host.Host, addrFunc AddrFunc, mt MetricsTracer) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

type client struct {
	h        host.Host
	addrFunc AddrFunc
	mt       MetricsTracer
}

func (c *client) DialBack(ctx context.Context, p peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

type Error struct {
	Status pb.Message_ResponseStatus
	Text   string
}

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e Error) IsDialError() bool { _ = "STUB: not implemented"; return false }

func (e Error) IsDialRefused() bool { _ = "STUB: not implemented"; return false }

func IsDialError(e error) bool { _ = "STUB: not implemented"; return false }

func IsDialRefused(e error) bool { _ = "STUB: not implemented"; return false }
