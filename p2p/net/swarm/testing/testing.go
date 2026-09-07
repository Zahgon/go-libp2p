package testing

import (
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	tptu "github.com/libp2p/go-libp2p/p2p/net/upgrader"

	ma "github.com/multiformats/go-multiaddr"
)

type config struct {
	disableReuseport    bool
	dialOnly            bool
	disableTCP          bool
	disableQUIC         bool
	disableWebTransport bool
	disableWebRTC       bool
	connectionGater     connmgr.ConnectionGater
	sk                  crypto.PrivKey
	swarmOpts           []swarm.Option
	eventBus            event.Bus
	clock
}

type clock interface {
	Now() time.Time
}

type realclock struct{}

func (rc realclock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type Option func(testing.TB, *config)

func WithClock(clock clock) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSwarmOpts(swarmOpts ...swarm.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

var OptDisableReuseport Option = func(_ testing.TB, c *config) {
	c.disableReuseport = true
}

var OptDialOnly Option = func(_ testing.TB, c *config) {
	c.dialOnly = true
}

var OptDisableTCP Option = func(_ testing.TB, c *config) {
	c.disableTCP = true
}

var OptDisableQUIC Option = func(_ testing.TB, c *config) {
	c.disableQUIC = true
}

var OptDisableWebTransport Option = func(_ testing.TB, c *config) {
	c.disableWebTransport = true
}

var OptDisableWebRTC Option = func(_ testing.TB, c *config) {
	c.disableWebRTC = true
}

func OptConnGater(cg connmgr.ConnectionGater) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func OptPeerPrivateKey(sk crypto.PrivKey) Option { _ = "STUB: not implemented"; return *new(Option) }

func EventBus(b event.Bus) Option { _ = "STUB: not implemented"; return *new(Option) }

func GenUpgrader(t testing.TB, n *swarm.Swarm, connGater connmgr.ConnectionGater, opts ...tptu.Option) transport.Upgrader {
	_ = "STUB: not implemented"
	return *new(transport.Upgrader)
}

func GenSwarm(t testing.TB, opts ...Option) *swarm.Swarm { _ = "STUB: not implemented"; return nil }

func DivulgeAddresses(a, b network.Network) { _ = "STUB: not implemented"; return }

type MockConnectionGater struct {
	Dial     func(p peer.ID, addr ma.Multiaddr) bool
	PeerDial func(p peer.ID) bool
	Accept   func(c network.ConnMultiaddrs) bool
	Secured  func(network.Direction, peer.ID, network.ConnMultiaddrs) bool
	Upgraded func(c network.Conn) (bool, control.DisconnectReason)
}

func DefaultMockConnectionGater() *MockConnectionGater { _ = "STUB: not implemented"; return nil }

func (m *MockConnectionGater) InterceptAddrDial(p peer.ID, addr ma.Multiaddr) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *MockConnectionGater) InterceptPeerDial(p peer.ID) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *MockConnectionGater) InterceptAccept(c network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *MockConnectionGater) InterceptSecured(d network.Direction, p peer.ID, c network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *MockConnectionGater) InterceptUpgraded(tc network.Conn) (allow bool, reason control.DisconnectReason) {
	_ = "STUB: not implemented"
	return false, *new(control.DisconnectReason)
}
