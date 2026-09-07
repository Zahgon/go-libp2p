package libp2p

import (
	"crypto/rand"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoremem"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	tls "github.com/libp2p/go-libp2p/p2p/security/tls"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	libp2pwebrtc "github.com/libp2p/go-libp2p/p2p/transport/webrtc"
	ws "github.com/libp2p/go-libp2p/p2p/transport/websocket"
	webtransport "github.com/libp2p/go-libp2p/p2p/transport/webtransport"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/multiformats/go-multiaddr"
)

var DefaultSecurity = ChainOptions(
	Security(tls.ID, tls.New),
	Security(noise.ID, noise.New),
)

var DefaultMuxers = Muxer(yamux.ID, yamux.DefaultTransport)

var DefaultTransports = ChainOptions(
	Transport(tcp.NewTCPTransport),
	Transport(quic.NewTransport),
	Transport(ws.New),
	Transport(webtransport.New),
	Transport(libp2pwebrtc.New),
)

var DefaultPrivateTransports = ChainOptions(
	Transport(tcp.NewTCPTransport),
	Transport(ws.New),
)

var DefaultPeerstore Option = func(cfg *Config) error {
	ps, err := pstoremem.NewPeerstore()
	if err != nil {
		return err
	}
	return cfg.Apply(Peerstore(ps))
}

var RandomIdentity = func(cfg *Config) error {
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return err
	}
	return cfg.Apply(Identity(priv))
}

var DefaultListenAddrs = func(cfg *Config) error {
	addrs := []string{
		"/ip4/0.0.0.0/tcp/0",
		"/ip4/0.0.0.0/udp/0/quic-v1",
		"/ip4/0.0.0.0/udp/0/quic-v1/webtransport",
		"/ip4/0.0.0.0/udp/0/webrtc-direct",
		"/ip6/::/tcp/0",
		"/ip6/::/udp/0/quic-v1",
		"/ip6/::/udp/0/quic-v1/webtransport",
		"/ip6/::/udp/0/webrtc-direct",
	}
	listenAddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, s := range addrs {
		addr, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			return err
		}
		listenAddrs = append(listenAddrs, addr)
	}
	return cfg.Apply(ListenAddrs(listenAddrs...))
}

var DefaultEnableRelay = func(cfg *Config) error {
	return cfg.Apply(EnableRelay())
}

var DefaultResourceManager = func(cfg *Config) error {

	limits := rcmgr.DefaultLimits
	SetDefaultServiceLimits(&limits)
	mgr, err := rcmgr.NewResourceManager(rcmgr.NewFixedLimiter(limits.AutoScale()))
	if err != nil {
		return err
	}

	return cfg.Apply(ResourceManager(mgr))
}

var DefaultConnectionManager = func(cfg *Config) error {
	mgr, err := connmgr.NewConnManager(160, 192)
	if err != nil {
		return err
	}

	return cfg.Apply(ConnectionManager(mgr))
}

var DefaultPrometheusRegisterer = func(cfg *Config) error {
	return cfg.Apply(PrometheusRegisterer(prometheus.DefaultRegisterer))
}

var defaultUDPBlackHoleDetector = func(cfg *Config) error {

	return cfg.Apply(UDPBlackHoleSuccessCounter(&swarm.BlackHoleSuccessCounter{N: 100, MinSuccesses: 5, Name: "UDP"}))
}

var defaultIPv6BlackHoleDetector = func(cfg *Config) error {

	return cfg.Apply(IPv6BlackHoleSuccessCounter(&swarm.BlackHoleSuccessCounter{N: 100, MinSuccesses: 5, Name: "IPv6"}))
}

var defaults = []struct {
	fallback func(cfg *Config) bool
	opt      Option
}{
	{
		fallback: func(cfg *Config) bool { return cfg.Transports == nil && cfg.ListenAddrs == nil },
		opt:      DefaultListenAddrs,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Transports == nil && cfg.PSK == nil },
		opt:      DefaultTransports,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Transports == nil && cfg.PSK != nil },
		opt:      DefaultPrivateTransports,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Muxers == nil },
		opt:      DefaultMuxers,
	},
	{
		fallback: func(cfg *Config) bool { return !cfg.Insecure && cfg.SecurityTransports == nil },
		opt:      DefaultSecurity,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.PeerKey == nil },
		opt:      RandomIdentity,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Peerstore == nil },
		opt:      DefaultPeerstore,
	},
	{
		fallback: func(cfg *Config) bool { return !cfg.RelayCustom },
		opt:      DefaultEnableRelay,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.ResourceManager == nil },
		opt:      DefaultResourceManager,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.ConnManager == nil },
		opt:      DefaultConnectionManager,
	},
	{
		fallback: func(cfg *Config) bool { return !cfg.DisableMetrics && cfg.PrometheusRegisterer == nil },
		opt:      DefaultPrometheusRegisterer,
	},
	{
		fallback: func(cfg *Config) bool {
			return !cfg.CustomUDPBlackHoleSuccessCounter && cfg.UDPBlackHoleSuccessCounter == nil
		},
		opt: defaultUDPBlackHoleDetector,
	},
	{
		fallback: func(cfg *Config) bool {
			return !cfg.CustomIPv6BlackHoleSuccessCounter && cfg.IPv6BlackHoleSuccessCounter == nil
		},
		opt: defaultIPv6BlackHoleDetector,
	},
}

var Defaults Option = func(cfg *Config) error {
	for _, def := range defaults {
		if err := cfg.Apply(def.opt); err != nil {
			return err
		}
	}
	return nil
}

var FallbackDefaults Option = func(cfg *Config) error {
	for _, def := range defaults {
		if !def.fallback(cfg) {
			continue
		}
		if err := cfg.Apply(def.opt); err != nil {
			return err
		}
	}
	return nil
}
