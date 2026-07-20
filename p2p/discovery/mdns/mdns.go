package mdns

import (
	"context"
	"io"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/zeroconf/v2"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	ServiceName   = "_p2p._udp"
	mdnsDomain    = "local"
	dnsaddrPrefix = "dnsaddr="
)

var log = logging.Logger("mdns")

type Service interface {
	Start() error
	io.Closer
}

type Notifee interface {
	HandlePeerFound(peer.AddrInfo)
}

type mdnsService struct {
	host        host.Host
	serviceName string
	peerName    string

	ctx       context.Context
	ctxCancel context.CancelFunc

	resolverWG sync.WaitGroup
	server     *zeroconf.Server

	notifee Notifee
}

func NewMdnsService(host host.Host, serviceName string, notifee Notifee) *mdnsService {
	_ = "STUB: not implemented"
	return nil
}

func (s *mdnsService) Start() error { _ = "STUB: not implemented"; return nil }

func (s *mdnsService) Close() error { _ = "STUB: not implemented"; return nil }

func (s *mdnsService) getIPs(addrs []ma.Multiaddr) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func containsUnsuitableProtocol(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func isSuitableForMDNS(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (s *mdnsService) startServer() error { _ = "STUB: not implemented"; return nil }

func (s *mdnsService) startResolver(ctx context.Context) { _ = "STUB: not implemented"; return }

func randomString(l int) string { _ = "STUB: not implemented"; return "" }
