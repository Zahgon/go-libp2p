package config

import (
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/pnet"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	bhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	tptu "github.com/libp2p/go-libp2p/p2p/net/upgrader"
	"github.com/libp2p/go-libp2p/p2p/protocol/autonatv2"
	relayv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	"github.com/prometheus/client_golang/prometheus"

	ma "github.com/multiformats/go-multiaddr"
	"go.uber.org/fx"
)

var log = logging.Logger("p2p-config")

type AddrsFactory = bhost.AddrsFactory

type NATManagerC func(network.Network) bhost.NATManager

type RoutingC func(host.Host) (routing.PeerRouting, error)

type AutoNATConfig struct {
	ForceReachability   *network.Reachability
	EnableService       bool
	ThrottleGlobalLimit int
	ThrottlePeerLimit   int
	ThrottleInterval    time.Duration
}

type Security struct {
	ID          protocol.ID
	Constructor any
}

type Config struct {
	UserAgent string

	ProtocolVersion string

	PeerKey crypto.PrivKey

	QUICReuse          []fx.Option
	Transports         []fx.Option
	Muxers             []tptu.StreamMuxer
	SecurityTransports []Security
	Insecure           bool
	PSK                pnet.PSK

	DialTimeout time.Duration

	RelayCustom bool
	Relay       bool

	EnableRelayService bool
	RelayServiceOpts   []relayv2.Option

	ListenAddrs     []ma.Multiaddr
	AddrsFactory    bhost.AddrsFactory
	ConnectionGater connmgr.ConnectionGater

	ConnManager     connmgr.ConnManager
	ResourceManager network.ResourceManager

	NATManager NATManagerC
	Peerstore  peerstore.Peerstore
	Reporter   metrics.Reporter

	MultiaddrResolver network.MultiaddrDNSResolver

	DisablePing bool

	DisableNonPublicAddrPublishing bool

	Routing RoutingC

	EnableAutoRelay bool
	AutoRelayOpts   []autorelay.Option
	AutoNATConfig

	EnableHolePunching  bool
	HolePunchingOptions []holepunch.Option

	DisableMetrics       bool
	PrometheusRegisterer prometheus.Registerer

	DialRanker network.DialRanker

	SwarmOpts []swarm.Option

	DisableIdentifyAddressDiscovery bool

	EnableAutoNATv2 bool

	UDPBlackHoleSuccessCounter        *swarm.BlackHoleSuccessCounter
	CustomUDPBlackHoleSuccessCounter  bool
	IPv6BlackHoleSuccessCounter       *swarm.BlackHoleSuccessCounter
	CustomIPv6BlackHoleSuccessCounter bool

	UserFxOptions []fx.Option

	ShareTCPListener bool
}

func (cfg *Config) makeSwarm(eventBus event.Bus, enableMetrics bool) (*swarm.Swarm, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg *Config) makeAutoNATV2Host() (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (cfg *Config) addTransports() ([]fx.Option, error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) newBasicHost(swrm *swarm.Swarm, eventBus event.Bus, an *autonatv2.AutoNAT, o bhost.ObservedAddrsManager) (*bhost.BasicHost, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg *Config) validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) NewNode() (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func (cfg *Config) addAutoNAT(h *bhost.BasicHost) error { _ = "STUB: not implemented"; return nil }

type Option func(cfg *Config) error

func (cfg *Config) Apply(opts ...Option) error { _ = "STUB: not implemented"; return nil }
