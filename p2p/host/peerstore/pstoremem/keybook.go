package pstoremem

import (
	"sync"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
)

type memoryKeyBook struct {
	sync.RWMutex
	pks map[peer.ID]ic.PubKey
	sks map[peer.ID]ic.PrivKey
}

var _ pstore.KeyBook = (*memoryKeyBook)(nil)

func NewKeyBook() *memoryKeyBook { _ = "STUB: not implemented"; return nil }

func (mkb *memoryKeyBook) PeersWithKeys() peer.IDSlice {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice)
}

func (mkb *memoryKeyBook) PubKey(p peer.ID) ic.PubKey {
	_ = "STUB: not implemented"
	return *new(ic.PubKey)
}

func (mkb *memoryKeyBook) AddPubKey(p peer.ID, pk ic.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (mkb *memoryKeyBook) PrivKey(p peer.ID) ic.PrivKey {
	_ = "STUB: not implemented"
	return *new(ic.PrivKey)
}

func (mkb *memoryKeyBook) AddPrivKey(p peer.ID, sk ic.PrivKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (mkb *memoryKeyBook) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
