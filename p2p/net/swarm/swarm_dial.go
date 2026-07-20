package swarm

import (
	"context"
	"errors"
	"net/netip"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	mafmt "github.com/multiformats/go-multiaddr-fmt"
)

const maximumResolvedAddresses = 100

const maximumDNSADDRRecursion = 4

var (
	ErrDialBackoff = errors.New("dial backoff")

	ErrDialRefusedBlackHole = errors.New("dial refused because of black hole")

	ErrDialToSelf = errors.New("dial to self attempted")

	ErrNoTransport = errors.New("no transport for protocol")

	ErrAllDialsFailed = errors.New("all dials failed")

	ErrNoAddresses = errors.New("no addresses")

	ErrNoGoodAddresses = errors.New("no good addresses")

	ErrGaterDisallowedConnection = errors.New("gater disallows connection to peer")
)

var ErrQUICDraft29 errQUICDraft29

type errQUICDraft29 struct{}

func (errQUICDraft29) Error() string { _ = "STUB: not implemented"; return "" }

func (errQUICDraft29) Unwrap() error { _ = "STUB: not implemented"; return nil }

const DialAttempts = 1

const ConcurrentFdDials = 160

var DefaultPerPeerRateLimit = 8

type DialBackoff struct {
	entries map[peer.ID]map[string]*backoffAddr
	lock    sync.RWMutex
}

type backoffAddr struct {
	tries int
	until time.Time
}

func (db *DialBackoff) init(ctx context.Context) {
	if db.entries == nil {
		db.entries = make(map[peer.ID]map[string]*backoffAddr)
	}
	go db.background(ctx)
}

func (db *DialBackoff) background(ctx context.Context) { _ = "STUB: not implemented"; return }

func (db *DialBackoff) Backoff(p peer.ID, addr ma.Multiaddr) (backoff bool) {
	_ = "STUB: not implemented"
	return false
}

var BackoffBase = time.Second * 5

var BackoffCoef = time.Second

var BackoffMax = time.Minute * 5

func (db *DialBackoff) AddBackoff(p peer.ID, addr ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (db *DialBackoff) Clear(p peer.ID) { _ = "STUB: not implemented"; return }

func (db *DialBackoff) cleanup() { _ = "STUB: not implemented"; return }

func (s *Swarm) DialPeer(ctx context.Context, p peer.ID) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (s *Swarm) dialPeer(ctx context.Context, p peer.ID) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) dialWorkerLoop(p peer.ID, reqch <-chan dialRequest) {
	_ = "STUB: not implemented"
	return
}

func (s *Swarm) addrsForDial(ctx context.Context, p peer.ID) (goodAddrs []ma.Multiaddr, addrErrs []TransportError, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func startsWithDNSComponent(m ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func stripP2PComponent(addrs []ma.Multiaddr) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

type resolver struct {
	canResolve func(ma.Multiaddr) bool
	resolve    func(ctx context.Context, maddr ma.Multiaddr, outputLimit int) ([]ma.Multiaddr, error)
}

type resolveErr struct {
	addr ma.Multiaddr
	err  error
}

func chainResolvers(ctx context.Context, addrs []ma.Multiaddr, outputLimit int, resolvers []resolver) ([]ma.Multiaddr, []resolveErr) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) resolveAddrs(ctx context.Context, pi peer.AddrInfo) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (s *Swarm) dialNextAddr(ctx context.Context, p peer.ID, addr ma.Multiaddr, resch chan transport.DialUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Swarm) CanDial(p peer.ID, addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (s *Swarm) nonProxyAddr(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

var quicDraft29DialMatcher = mafmt.And(mafmt.IP, mafmt.Base(ma.P_UDP), mafmt.Base(ma.P_QUIC))

func (s *Swarm) filterKnownUndialables(p peer.ID, addrs []ma.Multiaddr) (goodAddrs []ma.Multiaddr, addrErrs []TransportError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Swarm) limitedDial(ctx context.Context, p peer.ID, a ma.Multiaddr, resp chan transport.DialUpdate) {
	_ = "STUB: not implemented"
	return
}

func (s *Swarm) dialAddr(ctx context.Context, p peer.ID, addr ma.Multiaddr, updCh chan<- transport.DialUpdate) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func isFdConsumingAddr(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func isRelayAddr(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func filterLowPriorityAddresses(addrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func addrPort(a ma.Multiaddr, p int) (netip.AddrPort, error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil
}
