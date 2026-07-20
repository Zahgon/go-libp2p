package pstoremem

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
)

type pstoremem struct {
	peerstore.Metrics

	*memoryKeyBook
	*memoryAddrBook
	*memoryProtoBook
	*memoryPeerMetadata
}

var _ peerstore.Peerstore = &pstoremem{}

type Option any

func NewPeerstore(opts ...Option) (ps *pstoremem, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *pstoremem) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (ps *pstoremem) Peers() peer.IDSlice { _ = "STUB: not implemented"; return *new(peer.IDSlice) }

func (ps *pstoremem) PeerInfo(p peer.ID) peer.AddrInfo {
	_ = "STUB: not implemented"
	return *new(peer.AddrInfo)
}

func (ps *pstoremem) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
