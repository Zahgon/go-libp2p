package event

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

type EvtPeerConnectednessChanged struct {
	Peer peer.ID

	Connectedness network.Connectedness
}
