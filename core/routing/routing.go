package routing

import (
	"context"
	"errors"

	ci "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"

	cid "github.com/ipfs/go-cid"
)

var ErrNotFound = errors.New("routing: not found")

var ErrNotSupported = errors.New("routing: operation or key not supported")

type ContentProviding interface {
	Provide(context.Context, cid.Cid, bool) error
}

type ContentDiscovery interface {
	FindProvidersAsync(context.Context, cid.Cid, int) <-chan peer.AddrInfo
}

type ContentRouting interface {
	ContentProviding
	ContentDiscovery
}

type PeerRouting interface {
	FindPeer(context.Context, peer.ID) (peer.AddrInfo, error)
}

type ValueStore interface {
	PutValue(context.Context, string, []byte, ...Option) error

	GetValue(context.Context, string, ...Option) ([]byte, error)

	SearchValue(context.Context, string, ...Option) (<-chan []byte, error)
}

type Routing interface {
	ContentRouting
	PeerRouting
	ValueStore

	Bootstrap(context.Context) error
}

type PubKeyFetcher interface {
	GetPublicKey(context.Context, peer.ID) (ci.PubKey, error)
}

func KeyForPublicKey(id peer.ID) string { _ = "STUB: not implemented"; return "" }

func GetPublicKey(r ValueStore, ctx context.Context, p peer.ID) (ci.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ci.PubKey), nil
}
