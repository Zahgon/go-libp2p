package pstoremem

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
)

type memoryPeerMetadata struct {
	ds     map[peer.ID]map[string]any
	dslock sync.RWMutex
}

var _ pstore.PeerMetadata = (*memoryPeerMetadata)(nil)

func NewPeerMetadata() *memoryPeerMetadata { _ = "STUB: not implemented"; return nil }

func (ps *memoryPeerMetadata) Put(p peer.ID, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *memoryPeerMetadata) Get(p peer.ID, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (ps *memoryPeerMetadata) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
