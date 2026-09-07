package pstoremem

import (
	ma "github.com/multiformats/go-multiaddr"
)

func isFDCostlyTransport(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

type addrList []ma.Multiaddr

func (al addrList) Len() int      { _ = "STUB: not implemented"; return 0 }
func (al addrList) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (al addrList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
