package autonat

import (
	"github.com/libp2p/go-libp2p/core/host"

	ma "github.com/multiformats/go-multiaddr"
)

type dialPolicy struct {
	allowSelfDials bool
	host           host.Host
}

func (d *dialPolicy) skipDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (d *dialPolicy) skipPeer(addrs []ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }
