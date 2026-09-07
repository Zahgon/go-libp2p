package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/ipfs/go-log/v2"

	p2pforge "github.com/ipshipyard/p2p-forge/client"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	ws "github.com/libp2p/go-libp2p/p2p/transport/websocket"
)

var logger = log.Logger("autotls-example")

const userAgent = "go-libp2p/example/autotls"
const identityKeyFile = "identity.key"

func main() {

	ctx := context.Background()

	log.SetLogLevel("*", "error")
	log.SetLogLevel("autotls-example", "debug")
	log.SetLogLevel("basichost", "info")
	log.SetLogLevel("autotls", "debug")
	log.SetLogLevel("p2p-forge", "debug")
	log.SetLogLevel("nat", "debug")

	certLoaded := make(chan bool, 1)

	rawLogger := logger.Desugar()

	certManager, err := p2pforge.NewP2PForgeCertMgr(

		p2pforge.WithCAEndpoint(p2pforge.DefaultCATestEndpoint),

		p2pforge.WithCertificateStorage(&certmagic.FileStorage{Path: "p2p-forge-certs"}),

		p2pforge.WithLogger(rawLogger.Sugar().Named("autotls")),

		p2pforge.WithUserAgent(userAgent),

		p2pforge.WithRegistrationDelay(10*time.Second),

		p2pforge.WithOnCertLoaded(func() {
			certLoaded <- true
		}),
	)

	if err != nil {
		panic(err)
	}

	certManager.Start()
	defer certManager.Stop()

	privKey, err := LoadIdentity(identityKeyFile)
	if err != nil {
		panic(err)
	}

	opts := []libp2p.Option{
		libp2p.Identity(privKey),
		libp2p.DisableRelay(),
		libp2p.NATPortMap(),

		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/5500",
			"/ip6/::/tcp/5500",

			fmt.Sprintf("/ip4/0.0.0.0/tcp/5500/tls/sni/*.%s/ws", p2pforge.DefaultForgeDomain),
			fmt.Sprintf("/ip6/::/tcp/5500/tls/sni/*.%s/ws", p2pforge.DefaultForgeDomain),
		),

		libp2p.Transport(tcp.NewTCPTransport),

		libp2p.ShareTCPListener(),

		libp2p.Transport(ws.New, ws.WithTLSConfig(certManager.TLSConfig())),

		libp2p.UserAgent(userAgent),

		libp2p.AddrsFactory(certManager.AddressFactory()),
	}
	h, err := libp2p.New(opts...)
	if err != nil {
		panic(err)
	}

	logger.Info("Host created with PeerID: ", h.ID())

	dhtOpts := []dht.Option{
		dht.Mode(dht.ModeClient),
		dht.BootstrapPeers(dht.GetDefaultBootstrapPeerAddrInfos()...),
	}
	dht, err := dht.New(ctx, h, dhtOpts...)
	if err != nil {
		panic(err)
	}

	go dht.Bootstrap(ctx)

	logger.Info("Addresses: ", h.Addrs())

	certManager.ProvideHost(h)

	select {
	case <-certLoaded:
		logger.Info("TLS certificate loaded ")
		logger.Info("Addresses: ", h.Addrs())
	case <-ctx.Done():
		logger.Info("Context done")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
