package simlibp2p

import (
	"net"
	"sync/atomic"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peerstore"
	blankhost "github.com/libp2p/go-libp2p/p2p/host/blank"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"
	"github.com/marcopolo/simnet"
	"github.com/multiformats/go-multiaddr"
)

func MustNewHost(t *testing.T, opts ...libp2p.Option) host.Host {
	_ = "STUB: not implemented"
	return *new(host.Host)
}

type MockSourceIPSelector struct {
	ip atomic.Pointer[net.IP]
}

func (m *MockSourceIPSelector) PreferredSourceIPForDestination(_ *net.UDPAddr) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

const OneMbps = 1_000_000

func QUICSimnet(simnet *simnet.Simnet, linkSettings simnet.NodeBiDiLinkSettings, quicReuseOpts ...quicreuse.Option) libp2p.Option {
	_ = "STUB: not implemented"
	return *new(libp2p.Option)
}

type wrappedHost struct {
	blankhost.BlankHost
	ps        peerstore.Peerstore
	quicCM    *quicreuse.ConnManager
	idService identify.IDService
	connMgr   *connmgr.BasicConnMgr
}

func (h *wrappedHost) Close() error { _ = "STUB: not implemented"; return nil }

type BlankHostOpts struct {
	ConnMgr         *connmgr.BasicConnMgr
	listenMultiaddr multiaddr.Multiaddr
	simnet          *simnet.Simnet
	linkSettings    simnet.NodeBiDiLinkSettings
	quicReuseOpts   []quicreuse.Option
}

func newBlankHost(opts BlankHostOpts) (*wrappedHost, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type NodeLinkSettingsAndCount struct {
	LinkSettings simnet.NodeBiDiLinkSettings
	Count        int
}

type HostAndIdx struct {
	Host host.Host
	Idx  int
}

type SimpleLibp2pNetworkMeta struct {
	Nodes      []host.Host
	AddrToNode map[string]HostAndIdx
}

type NetworkSettings struct {
	UseBlankHost            bool
	QUICReuseOptsForHostIdx func(idx int) []quicreuse.Option
	BlankHostOptsForHostIdx func(idx int) BlankHostOpts
}

type LatencyFunc func(*simnet.Packet) time.Duration

func SimpleLibp2pNetwork(linkSettings []NodeLinkSettingsAndCount, latencyFunc LatencyFunc, networkSettings NetworkSettings) (*simnet.Simnet, *SimpleLibp2pNetworkMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GetBasicHostPair(t *testing.T) (host.Host, host.Host) {
	_ = "STUB: not implemented"
	return *new(host.Host), *new(host.Host)
}
