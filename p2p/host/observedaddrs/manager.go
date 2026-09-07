package observedaddrs

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	logging "github.com/libp2p/go-libp2p/gologshim"
	basichost "github.com/libp2p/go-libp2p/p2p/host/basic"

	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("observedaddrs")

var ActivationThresh = 4

var (
	observedAddrManagerWorkerChannelSize = 16

	natTypeChangeTickrInterval = 1 * time.Minute
)

const maxExternalThinWaistAddrsPerLocalAddr = 3

type thinWaist struct {
	Addr, TW, Rest ma.Multiaddr
}

var errTW = errors.New("not a thinwaist address")

func thinWaistForm(a ma.Multiaddr) (thinWaist, error) {
	_ = "STUB: not implemented"
	return *new(thinWaist), nil
}

func getObserver(a ma.Multiaddr) (string, error) { _ = "STUB: not implemented"; return "", nil }

type connMultiaddrs interface {
	network.ConnMultiaddrs
	IsClosed() bool
}

const observerSetCacheSize = 10

type observerSet struct {
	ObservedTWAddr ma.Multiaddr
	ObservedBy     map[string]int

	mu               sync.RWMutex
	cachedMultiaddrs map[string]ma.Multiaddr
}

func (s *observerSet) cacheMultiaddr(addr ma.Multiaddr) ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

type observation struct {
	conn     connMultiaddrs
	observed ma.Multiaddr
}

type Manager struct {
	listenAddrs func() []ma.Multiaddr

	wch chan observation

	eventbus event.Bus

	wg         sync.WaitGroup
	ctx        context.Context
	ctxCancel  context.CancelFunc
	stopNotify func()

	mu sync.RWMutex

	externalAddrs map[string]map[string]*observerSet

	connObservedTWAddrs map[connMultiaddrs]ma.Multiaddr
}

var _ basichost.ObservedAddrsManager = (*Manager)(nil)

func NewManager(eventbus event.Bus, net network.Network) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newManagerWithListenAddrs(bus event.Bus, listenAddrs func() []ma.Multiaddr) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Manager) Start(n network.Network) { _ = "STUB: not implemented"; return }

func (o *Manager) AddrsFor(addr ma.Multiaddr) (addrs []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil
}

func (o *Manager) appendInferredAddrs(twToObserverSets map[string][]*observerSet, addrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (o *Manager) Addrs(minObservers int) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (o *Manager) getTopExternalAddrs(localTWStr string, minObservers int) []*observerSet {
	_ = "STUB: not implemented"
	return nil
}

func (o *Manager) eventHandler(identifySub event.Subscription, natEmitter event.Emitter) {
	_ = "STUB: not implemented"
	return
}

func (o *Manager) worker() { _ = "STUB: not implemented"; return }

func (o *Manager) shouldRecordObservation(conn connMultiaddrs, observed ma.Multiaddr) (shouldRecord bool, localTW thinWaist, observedTW thinWaist) {
	_ = "STUB: not implemented"
	return false, *new(thinWaist), *new(thinWaist)
}

func (o *Manager) maybeRecordObservation(conn connMultiaddrs, observed ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (o *Manager) recordObservationUnlocked(conn connMultiaddrs, localTW, observedTW thinWaist) {
	_ = "STUB: not implemented"
	return
}

func (o *Manager) removeExternalAddrsUnlocked(observer, localTWStr, observedTWStr string) {
	_ = "STUB: not implemented"
	return
}

func (o *Manager) addExternalAddrsUnlocked(observedTWAddr ma.Multiaddr, observer, localTWStr, observedTWStr string) {
	_ = "STUB: not implemented"
	return
}

func (o *Manager) removeConn(conn connMultiaddrs) { _ = "STUB: not implemented"; return }

func (o *Manager) getNATType() (tcpNATType, udpNATType network.NATDeviceType) {
	_ = "STUB: not implemented"
	return *new(network.NATDeviceType), *new(network.NATDeviceType)
}

func (o *Manager) Close() error { _ = "STUB: not implemented"; return nil }

func hasConsistentTransport(aTW, bTW ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func isRelayedAddress(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }
