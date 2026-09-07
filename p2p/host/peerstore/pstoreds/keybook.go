package pstoreds

import (
	"context"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"

	ds "github.com/ipfs/go-datastore"
)

var (
	kbBase     = ds.NewKey("/peers/keys")
	pubSuffix  = ds.NewKey("/pub")
	privSuffix = ds.NewKey("/priv")
)

type dsKeyBook struct {
	ds ds.Datastore
}

var _ pstore.KeyBook = (*dsKeyBook)(nil)

func NewKeyBook(_ context.Context, store ds.Datastore, _ Options) (*dsKeyBook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kb *dsKeyBook) PubKey(p peer.ID) ic.PubKey { _ = "STUB: not implemented"; return *new(ic.PubKey) }

func (kb *dsKeyBook) AddPubKey(p peer.ID, pk ic.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (kb *dsKeyBook) PrivKey(p peer.ID) ic.PrivKey {
	_ = "STUB: not implemented"
	return *new(ic.PrivKey)
}

func (kb *dsKeyBook) AddPrivKey(p peer.ID, sk ic.PrivKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (kb *dsKeyBook) PeersWithKeys() peer.IDSlice {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice)
}

func (kb *dsKeyBook) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }

func peerToKey(p peer.ID, suffix ds.Key) ds.Key { _ = "STUB: not implemented"; return *new(ds.Key) }
