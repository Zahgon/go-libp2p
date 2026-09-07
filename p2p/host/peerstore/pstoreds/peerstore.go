package pstoreds

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

type Options struct {
	CacheSize uint

	MaxProtocols int

	MaxAddrsPerPeer int

	GCPurgeInterval time.Duration

	GCLookaheadInterval time.Duration

	GCInitialDelay time.Duration

	Clock clock
}

func DefaultOpts() Options { _ = "STUB: not implemented"; return *new(Options) }

type pstoreds struct {
	peerstore.Metrics

	*dsKeyBook
	*dsAddrBook
	*dsProtoBook
	*dsPeerMetadata
}

var _ peerstore.Peerstore = &pstoreds{}

func NewPeerstore(ctx context.Context, store ds.Batching, opts Options) (*pstoreds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uniquePeerIds(ds ds.Datastore, prefix ds.Key, extractor func(result query.Result) string) (peer.IDSlice, error) {
	_ = "STUB: not implemented"
	return *new(peer.IDSlice), nil
}

func (ps *pstoreds) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (ps *pstoreds) Peers() peer.IDSlice { _ = "STUB: not implemented"; return *new(peer.IDSlice) }

func (ps *pstoreds) PeerInfo(p peer.ID) peer.AddrInfo {
	_ = "STUB: not implemented"
	return *new(peer.AddrInfo)
}

func (ps *pstoreds) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
