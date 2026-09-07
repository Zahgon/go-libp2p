package ttransport

import (
	"testing"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

var testData = []byte("this is some test data")

func SubtestProtocols(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, _ peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestBasic(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestPingPong(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestCancel(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}
