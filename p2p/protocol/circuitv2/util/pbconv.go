package util

import (
	"github.com/libp2p/go-libp2p/core/peer"
	pbv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pb"
)

func PeerToPeerInfoV2(p *pbv2.Peer) (peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return *new(peer.AddrInfo), nil
}

func PeerInfoToPeerV2(pi peer.AddrInfo) *pbv2.Peer { _ = "STUB: not implemented"; return nil }
