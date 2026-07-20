package autonatv2

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/autonatv2/pb"
	"github.com/libp2p/go-msgio/pbio"
	ma "github.com/multiformats/go-multiaddr"
)

type client struct {
	host          host.Host
	dialData      []byte
	metricsTracer MetricsTracer

	mu sync.Mutex

	dialBackQueues map[uint64]chan ma.Multiaddr
}

func newClient(s *autoNATSettings) *client { _ = "STUB: not implemented"; return nil }

func (ac *client) Start(h host.Host) { _ = "STUB: not implemented"; return }

func (ac *client) Close() { _ = "STUB: not implemented"; return }

func (ac *client) GetReachability(ctx context.Context, p peer.ID, reqs []Request) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (ac *client) getReachability(ctx context.Context, p peer.ID, reqs []Request) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func validateDialDataRequest(reqs []Request, msg *pb.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *client) newResult(resp *pb.DialResponse, reqs []Request, dialBackAddr ma.Multiaddr) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

//nolint:ineffassign

func sendDialData(dialData []byte, numBytes int, w pbio.Writer, msg *pb.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func newDialRequest(reqs []Request, nonce uint64) pb.Message {
	_ = "STUB: not implemented"
	return *new(pb.Message)
}

func (ac *client) handleDialBack(s network.Stream) { _ = "STUB: not implemented"; return }

var tlsWSAddr = ma.StringCast("/tls/ws")

func normalizeMultiaddr(addr ma.Multiaddr) ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func removeTrailing(addr ma.Multiaddr, protocolCode int) ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func (ac *client) areAddrsConsistent(connLocalAddr, dialedAddr ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}
