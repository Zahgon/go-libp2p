package peerstore

import (
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
)

func PeerInfos(ps pstore.Peerstore, peers peer.IDSlice) []peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}

func PeerInfoIDs(pis []peer.AddrInfo) peer.IDSlice {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice)
}
