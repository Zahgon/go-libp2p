package event

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/record"
	"github.com/multiformats/go-multiaddr"
)

type EvtPeerIdentificationCompleted struct {
	Peer peer.ID

	Conn network.Conn

	ListenAddrs []multiaddr.Multiaddr

	Protocols []protocol.ID

	SignedPeerRecord *record.Envelope

	AgentVersion string

	ProtocolVersion string

	ObservedAddr multiaddr.Multiaddr
}

type EvtPeerIdentificationFailed struct {
	Peer peer.ID

	Reason error
}
