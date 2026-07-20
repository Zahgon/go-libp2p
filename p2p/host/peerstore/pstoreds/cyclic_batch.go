package pstoreds

import (
	"context"

	ds "github.com/ipfs/go-datastore"
)

var defaultOpsPerCyclicBatch = 20

type cyclicBatch struct {
	threshold int
	ds.Batch
	ds      ds.Batching
	pending int
}

func newCyclicBatch(ds ds.Batching, _ int) (ds.Batch, error) {
	_ = "STUB: not implemented"
	return *new(ds.Batch), nil
}

func (cb *cyclicBatch) cycle() (err error) { _ = "STUB: not implemented"; return nil }

func (cb *cyclicBatch) Put(ctx context.Context, key ds.Key, val []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *cyclicBatch) Delete(ctx context.Context, key ds.Key) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *cyclicBatch) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
