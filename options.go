package libp2p

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/pnet"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	relayv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"
	"github.com/prometheus/client_golang/prometheus"

	ma "github.com/multiformats/go-multiaddr"
	"go.uber.org/fx"
)

func ListenAddrStrings(s ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ListenAddrs(addrs ...ma.Multiaddr) Option { _ = "STUB: not implemented"; return *new(Option) }

func Security(name string, constructor any) Option { _ = "STUB: not implemented"; return *new(Option) }

var NoSecurity Option = func(cfg *Config) error {
	if len(cfg.SecurityTransports) > 0 {
		return fmt.Errorf("cannot use security transports with an insecure libp2p configuration")
	}
	cfg.Insecure = true
	return nil
}

func Muxer(name string, muxer network.Multiplexer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func QUICReuse(constructor any, opts ...quicreuse.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Transport(constructor any, opts ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

func Peerstore(ps peerstore.Peerstore) Option { _ = "STUB: not implemented"; return *new(Option) }

func PrivateNetwork(psk pnet.PSK) Option { _ = "STUB: not implemented"; return *new(Option) }

func BandwidthReporter(rep metrics.Reporter) Option { _ = "STUB: not implemented"; return *new(Option) }

func Identity(sk crypto.PrivKey) Option { _ = "STUB: not implemented"; return *new(Option) }

func ConnectionManager(connman connmgr.ConnManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func AddrsFactory(factory config.AddrsFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableRelay() Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableRelay() Option { _ = "STUB: not implemented"; return *new(Option) }

func EnableRelayService(opts ...relayv2.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableAutoRelay(opts ...autorelay.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableAutoRelayWithStaticRelays(static []peer.AddrInfo, opts ...autorelay.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableAutoRelayWithPeerSource(peerSource autorelay.PeerSource, opts ...autorelay.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ForceReachabilityPublic() Option { _ = "STUB: not implemented"; return *new(Option) }

func ForceReachabilityPrivate() Option { _ = "STUB: not implemented"; return *new(Option) }

func EnableNATService() Option { _ = "STUB: not implemented"; return *new(Option) }

func AutoNATServiceRateLimit(global, perPeer int, interval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ConnectionGater(cg connmgr.ConnectionGater) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ResourceManager(rcmgr network.ResourceManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NATPortMap() Option { _ = "STUB: not implemented"; return *new(Option) }

func NATManager(nm config.NATManagerC) Option { _ = "STUB: not implemented"; return *new(Option) }

func Ping(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func NonPublicAddrPublishing(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func Routing(rt config.RoutingC) Option { _ = "STUB: not implemented"; return *new(Option) }

var NoListenAddrs = func(cfg *Config) error {
	cfg.ListenAddrs = []ma.Multiaddr{}
	if !cfg.RelayCustom {
		cfg.RelayCustom = true
		cfg.Relay = false
	}
	return nil
}

var NoTransports = func(cfg *Config) error {
	cfg.Transports = []fx.Option{}
	return nil
}

func ProtocolVersion(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

func UserAgent(userAgent string) Option { _ = "STUB: not implemented"; return *new(Option) }

func MultiaddrResolver(rslv network.MultiaddrDNSResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableHolePunching(opts ...holepunch.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDialTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableMetrics() Option { _ = "STUB: not implemented"; return *new(Option) }

func PrometheusRegisterer(reg prometheus.Registerer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func DialRanker(d network.DialRanker) Option { _ = "STUB: not implemented"; return *new(Option) }

func SwarmOpts(opts ...swarm.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableIdentifyAddressDiscovery() Option { _ = "STUB: not implemented"; return *new(Option) }

func EnableAutoNATv2() Option { _ = "STUB: not implemented"; return *new(Option) }

func UDPBlackHoleSuccessCounter(f *swarm.BlackHoleSuccessCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func IPv6BlackHoleSuccessCounter(f *swarm.BlackHoleSuccessCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFxOption(opts ...fx.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func ShareTCPListener() Option { _ = "STUB: not implemented"; return *new(Option) }
