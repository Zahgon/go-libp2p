package rcmgr

import (
	"net"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/multiformats/go-multiaddr"
)

type Allowlist struct {
	mu sync.RWMutex

	allowedNetworks []*net.IPNet

	allowedPeerByNetwork map[peer.ID][]*net.IPNet
}

func WithAllowlistedMultiaddrs(mas []multiaddr.Multiaddr) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func newAllowlist() Allowlist { _ = "STUB: not implemented"; return *new(Allowlist) }

func toIPNet(ma multiaddr.Multiaddr) (*net.IPNet, peer.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(peer.ID), nil
}

func (al *Allowlist) Add(ma multiaddr.Multiaddr) error { _ = "STUB: not implemented"; return nil }

func (al *Allowlist) Remove(ma multiaddr.Multiaddr) error { _ = "STUB: not implemented"; return nil }

func (al *Allowlist) Allowed(ma multiaddr.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (al *Allowlist) AllowedPeerAndMultiaddr(peerID peer.ID, ma multiaddr.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}
