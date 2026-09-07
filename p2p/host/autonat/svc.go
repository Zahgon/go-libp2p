package autonat

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/autonat/pb"

	ma "github.com/multiformats/go-multiaddr"
)

var streamTimeout = 60 * time.Second

const (
	ServiceName = "libp2p.autonat"

	maxMsgSize = 4096
)

type autoNATService struct {
	instanceLock      sync.Mutex
	instance          context.CancelFunc
	backgroundRunning chan struct{}

	config *config

	mx         sync.Mutex
	reqs       map[peer.ID]int
	globalReqs int
}

func newAutoNATService(c *config) (*autoNATService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (as *autoNATService) handleStream(s network.Stream) { _ = "STUB: not implemented"; return }

func (as *autoNATService) handleDial(p peer.ID, obsaddr ma.Multiaddr, mpi *pb.Message_PeerInfo) *pb.Message_DialResponse {
	_ = "STUB: not implemented"
	return nil
}

func (as *autoNATService) doDial(pi peer.AddrInfo) *pb.Message_DialResponse {
	_ = "STUB: not implemented"
	return nil
}

func (as *autoNATService) Enable() { _ = "STUB: not implemented"; return }

func (as *autoNATService) Disable() { _ = "STUB: not implemented"; return }

func (as *autoNATService) Close() error { _ = "STUB: not implemented"; return nil }

func (as *autoNATService) background(ctx context.Context) { _ = "STUB: not implemented"; return }
