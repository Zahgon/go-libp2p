package libp2p

import (
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

type filtersConnectionGater ma.Filters

var _ connmgr.ConnectionGater = (*filtersConnectionGater)(nil)

func (f *filtersConnectionGater) InterceptAddrDial(_ peer.ID, addr ma.Multiaddr) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (f *filtersConnectionGater) InterceptPeerDial(_ peer.ID) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (f *filtersConnectionGater) InterceptAccept(connAddr network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (f *filtersConnectionGater) InterceptSecured(_ network.Direction, _ peer.ID, connAddr network.ConnMultiaddrs) (allow bool) {
	_ = "STUB: not implemented"
	return false
}

func (f *filtersConnectionGater) InterceptUpgraded(_ network.Conn) (allow bool, reason control.DisconnectReason) {
	_ = "STUB: not implemented"
	return false, *new(control.DisconnectReason)
}
