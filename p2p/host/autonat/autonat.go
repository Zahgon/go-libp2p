package autonat

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("autonat")

const maxConfidence = 3

type AmbientAutoNAT struct {
	host host.Host

	*config

	ctx               context.Context
	ctxCancel         context.CancelFunc
	backgroundRunning chan struct{}

	inboundConn   chan network.Conn
	dialResponses chan error

	observations chan network.Reachability

	status atomic.Pointer[network.Reachability]

	confidence    int
	lastInbound   time.Time
	lastProbe     time.Time
	recentProbes  map[peer.ID]time.Time
	pendingProbes int
	ourAddrs      map[string]struct{}

	service *autoNATService

	emitReachabilityChanged event.Emitter
	subscriber              event.Subscription
}

type StaticAutoNAT struct {
	host         host.Host
	reachability network.Reachability
	service      *autoNATService
}

func New(h host.Host, options ...Option) (AutoNAT, error) {
	_ = "STUB: not implemented"
	return *new(AutoNAT), nil
}

func (as *AmbientAutoNAT) Status() network.Reachability {
	_ = "STUB: not implemented"
	return *new(network.Reachability)
}

func (as *AmbientAutoNAT) emitStatus() { _ = "STUB: not implemented"; return }

func ipInList(candidate ma.Multiaddr, list []ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}

func (as *AmbientAutoNAT) background() { _ = "STUB: not implemented"; return }

func (as *AmbientAutoNAT) checkAddrs() (hasNewAddr bool) { _ = "STUB: not implemented"; return false }

func (as *AmbientAutoNAT) scheduleProbe(forceProbe bool) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (as *AmbientAutoNAT) handleDialResponse(dialErr error) { _ = "STUB: not implemented"; return }

func (as *AmbientAutoNAT) recordObservation(observation network.Reachability) {
	_ = "STUB: not implemented"
	return
}

func (as *AmbientAutoNAT) tryProbe(p peer.ID) { _ = "STUB: not implemented"; return }

func (as *AmbientAutoNAT) probe(pi *peer.AddrInfo) { _ = "STUB: not implemented"; return }

func (as *AmbientAutoNAT) getPeerToProbe() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (as *AmbientAutoNAT) Close() error { _ = "STUB: not implemented"; return nil }

func (s *StaticAutoNAT) Status() network.Reachability {
	_ = "STUB: not implemented"
	return *new(network.Reachability)
}

func (s *StaticAutoNAT) Close() error { _ = "STUB: not implemented"; return nil }
