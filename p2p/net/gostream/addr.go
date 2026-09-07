package gostream

import "github.com/libp2p/go-libp2p/core/peer"

type addr struct{ id peer.ID }

func (a *addr) Network() string { _ = "STUB: not implemented"; return "" }

func (a *addr) String() string { _ = "STUB: not implemented"; return "" }
