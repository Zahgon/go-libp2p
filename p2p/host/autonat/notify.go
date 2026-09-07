package autonat

import (
	"github.com/libp2p/go-libp2p/core/network"

	ma "github.com/multiformats/go-multiaddr"
)

var _ network.Notifiee = (*AmbientAutoNAT)(nil)

func (as *AmbientAutoNAT) Listen(_ network.Network, _ ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (as *AmbientAutoNAT) ListenClose(_ network.Network, _ ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (as *AmbientAutoNAT) Connected(_ network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

func (as *AmbientAutoNAT) Disconnected(_ network.Network, _ network.Conn) {
	_ = "STUB: not implemented"
	return
}
