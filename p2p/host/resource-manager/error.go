package rcmgr

import (
	"github.com/libp2p/go-libp2p/core/network"
)

type ErrStreamOrConnLimitExceeded struct {
	current, attempted, limit int
	err                       error
}

func (e *ErrStreamOrConnLimitExceeded) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrStreamOrConnLimitExceeded) Unwrap() error { _ = "STUB: not implemented"; return nil }

func logValuesStreamLimit(scope, edge string, dir network.Direction, stat network.ScopeStat, err error) []any {
	_ = "STUB: not implemented"
	return nil
}

func logValuesConnLimit(scope, edge string, dir network.Direction, usefd bool, stat network.ScopeStat, err error) []any {
	_ = "STUB: not implemented"
	return nil
}

type ErrMemoryLimitExceeded struct {
	current, attempted, limit int64
	priority                  uint8
	err                       error
}

func (e *ErrMemoryLimitExceeded) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrMemoryLimitExceeded) Unwrap() error { _ = "STUB: not implemented"; return nil }

func logValuesMemoryLimit(scope, edge string, stat network.ScopeStat, err error) []any {
	_ = "STUB: not implemented"
	return nil
}
