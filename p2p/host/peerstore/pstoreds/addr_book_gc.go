package pstoreds

import (
	"context"

	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

var (
	gcLookaheadBase = ds.NewKey("/peers/gc/addrs")

	purgeLookaheadQuery = query.Query{
		Prefix:   gcLookaheadBase.String(),
		Orders:   []query.Order{query.OrderByFunction(orderByTimestampInKey)},
		KeysOnly: true,
	}

	purgeStoreQuery = query.Query{
		Prefix:   addrBookBase.String(),
		Orders:   []query.Order{query.OrderByKey{}},
		KeysOnly: false,
	}

	populateLookaheadQuery = query.Query{
		Prefix:   addrBookBase.String(),
		Orders:   []query.Order{query.OrderByKey{}},
		KeysOnly: true,
	}
)

type dsAddrBookGc struct {
	ctx              context.Context
	ab               *dsAddrBook
	running          chan struct{}
	lookaheadEnabled bool
	purgeFunc        func()
	currWindowEnd    int64
}

func newAddressBookGc(ctx context.Context, ab *dsAddrBook) (*dsAddrBookGc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gc *dsAddrBookGc) background() { _ = "STUB: not implemented"; return }

func (gc *dsAddrBookGc) purgeLookahead() { _ = "STUB: not implemented"; return }

func (gc *dsAddrBookGc) purgeStore() { _ = "STUB: not implemented"; return }

func (gc *dsAddrBookGc) populateLookahead() { _ = "STUB: not implemented"; return }

func orderByTimestampInKey(a, b query.Entry) int { _ = "STUB: not implemented"; return 0 }
