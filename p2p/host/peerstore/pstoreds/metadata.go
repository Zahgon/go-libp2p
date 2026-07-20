package pstoreds

import (
	"context"
	"encoding/gob"

	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"

	ds "github.com/ipfs/go-datastore"
)

var pmBase = ds.NewKey("/peers/metadata")

type dsPeerMetadata struct {
	ds ds.Datastore
}

var _ pstore.PeerMetadata = (*dsPeerMetadata)(nil)

func init() {

	gob.Register(make(map[protocol.ID]struct{}))
}

func NewPeerMetadata(_ context.Context, store ds.Datastore, _ Options) (*dsPeerMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pm *dsPeerMetadata) Get(p peer.ID, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (pm *dsPeerMetadata) Put(p peer.ID, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (pm *dsPeerMetadata) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
