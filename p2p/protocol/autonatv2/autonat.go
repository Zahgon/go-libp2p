package autonatv2

import (
	"context"
	"errors"
	"iter"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	ServiceName      = "libp2p.autonatv2"
	DialBackProtocol = "/libp2p/autonat/2/dial-back"
	DialProtocol     = "/libp2p/autonat/2/dial-request"

	maxMsgSize            = 8192
	streamTimeout         = 15 * time.Second
	dialBackStreamTimeout = 5 * time.Second
	dialBackDialTimeout   = 10 * time.Second
	dialBackMaxMsgSize    = 1024
	minHandshakeSizeBytes = 30_000
	maxHandshakeSizeBytes = 100_000

	maxPeerAddresses = 50

	defaultThrottlePeerDuration = 2 * time.Minute
)

var (
	ErrNoPeers = errors.New("no peers for autonat v2")

	ErrPrivateAddrs = errors.New("private addresses cannot be verified with autonatv2")

	log = logging.Logger("autonatv2")
)

type Request struct {
	Addr ma.Multiaddr

	SendDialData bool
}

type Result struct {
	Addr ma.Multiaddr

	Idx int

	Reachability network.Reachability

	AllAddrsRefused bool
}

type AutoNAT struct {
	host host.Host

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	srv *server
	cli *client

	mx           sync.Mutex
	peers        *peersMap
	throttlePeer map[peer.ID]time.Time

	throttlePeerDuration time.Duration

	allowPrivateAddrs bool
}

func New(dialerHost host.Host, opts ...AutoNATOption) (*AutoNAT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (an *AutoNAT) background(sub event.Subscription) { _ = "STUB: not implemented"; return }

func (an *AutoNAT) Start(h host.Host) error { _ = "STUB: not implemented"; return nil }

func (an *AutoNAT) Close() { _ = "STUB: not implemented"; return }

func (an *AutoNAT) GetReachability(ctx context.Context, reqs []Request) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (an *AutoNAT) updatePeer(p peer.ID) { _ = "STUB: not implemented"; return }

type peersMap struct {
	peerIdx map[peer.ID]int
	peers   []peer.ID
}

func newPeersMap() *peersMap { _ = "STUB: not implemented"; return nil }

func (p *peersMap) Shuffled() iter.Seq[peer.ID] { _ = "STUB: not implemented"; return nil }

func (p *peersMap) Put(id peer.ID) { _ = "STUB: not implemented"; return }

func (p *peersMap) Delete(id peer.ID) { _ = "STUB: not implemented"; return }
