package autonatv2

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/autonatv2/pb"
	"github.com/libp2p/go-msgio/pbio"

	ma "github.com/multiformats/go-multiaddr"
)

var (
	errResourceLimitExceeded = errors.New("resource limit exceeded")
	errBadRequest            = errors.New("bad request")
	errDialDataRefused       = errors.New("dial data refused")
)

type dataRequestPolicyFunc = func(observedAddr, dialAddr ma.Multiaddr) bool

type EventDialRequestCompleted struct {
	Error            error
	ResponseStatus   pb.DialResponse_ResponseStatus
	DialStatus       pb.DialStatus
	DialDataRequired bool
	DialedAddr       ma.Multiaddr
}

type server struct {
	host       host.Host
	dialerHost host.Host
	limiter    *rateLimiter

	dialDataRequestPolicy                dataRequestPolicyFunc
	amplificatonAttackPreventionDialWait time.Duration
	metricsTracer                        MetricsTracer

	now               func() time.Time
	allowPrivateAddrs bool
}

func newServer(dialer host.Host, s *autoNATSettings) *server { _ = "STUB: not implemented"; return nil }

func (as *server) Start(h host.Host) { _ = "STUB: not implemented"; return }

func (as *server) Close() { _ = "STUB: not implemented"; return }

func (as *server) handleDialRequest(s network.Stream) { _ = "STUB: not implemented"; return }

func (as *server) serveDialRequest(s network.Stream) EventDialRequestCompleted {
	_ = "STUB: not implemented"
	return *new(EventDialRequestCompleted)
}

func getDialData(w pbio.Writer, s network.Stream, msg *pb.Message, addrIdx int) error {
	_ = "STUB: not implemented"
	return nil
}

func readDialData(numBytes int, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (as *server) dialBack(ctx context.Context, p peer.ID, addr ma.Multiaddr, nonce uint64) pb.DialStatus {
	_ = "STUB: not implemented"
	return *new(pb.DialStatus)
}

type rateLimiter struct {
	PerPeerRPM int

	RPM int

	DialDataRPM int

	MaxConcurrentRequestsPerPeer int

	mu           sync.Mutex
	closed       bool
	reqs         []entry
	peerReqs     map[peer.ID][]time.Time
	dialDataReqs []time.Time

	inProgressReqs map[peer.ID]int

	now func() time.Time
}

type entry struct {
	PeerID peer.ID
	Time   time.Time
}

func (r *rateLimiter) init() {
	if r.peerReqs == nil {
		r.peerReqs = make(map[peer.ID][]time.Time)
		r.inProgressReqs = make(map[peer.ID]int)
	}
}

func (r *rateLimiter) Accept(p peer.ID) bool { _ = "STUB: not implemented"; return false }

func (r *rateLimiter) AcceptDialDataRequest() bool { _ = "STUB: not implemented"; return false }

func (r *rateLimiter) cleanup(now time.Time) { _ = "STUB: not implemented"; return }

func (r *rateLimiter) CompleteRequest(p peer.ID) { _ = "STUB: not implemented"; return }

func (r *rateLimiter) Close() { _ = "STUB: not implemented"; return }

func amplificationAttackPrevention(observedAddr, dialAddr ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}
