package conngater

import (
	"context"
	"net"
	"sync"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"

	"github.com/ipfs/go-datastore"
	logging "github.com/libp2p/go-libp2p/gologshim"
)

type BasicConnectionGater struct {
	sync.RWMutex

	blockedPeers   map[peer.ID]struct{}
	blockedAddrs   map[string]struct{}
	blockedSubnets map[string]*net.IPNet

	ds datastore.Datastore
}

var log = logging.Logger("net/conngater")

const (
	ns        = "/libp2p/net/conngater"
	keyPeer   = "/peer/"
	keyAddr   = "/addr/"
	keySubnet = "/subnet/"
)

func NewBasicConnectionGater(ds datastore.Datastore) (*BasicConnectionGater, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cg *BasicConnectionGater) loadRules(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (cg *BasicConnectionGater) BlockPeer(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) UnblockPeer(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) ListBlockedPeers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) BlockAddr(ip net.IP) error { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) UnblockAddr(ip net.IP) error { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) ListBlockedAddrs() []net.IP { _ = "STUB: not implemented"; return nil }

func (cg *BasicConnectionGater) BlockSubnet(ipnet *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (cg *BasicConnectionGater) UnblockSubnet(ipnet *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (cg *BasicConnectionGater) ListBlockedSubnets() []*net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

var _ connmgr.ConnectionGater = (*BasicConnectionGater)(nil)

func (cg *BasicConnectionGater) InterceptPeerDial(p peer.ID) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (cg *BasicConnectionGater) InterceptAddrDial(_ peer.ID, a ma.Multiaddr) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (cg *BasicConnectionGater) InterceptAccept(cma network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (cg *BasicConnectionGater) InterceptSecured(dir network.Direction, p peer.ID, _ network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (cg *BasicConnectionGater) InterceptUpgraded(network.Conn) (allow bool, reason control.DisconnectReason) {
	_ = "STUB: not implemented"
	return false, *new(control.DisconnectReason)
}
