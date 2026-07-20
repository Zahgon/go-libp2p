package peer

import (
	"fmt"

	ma "github.com/multiformats/go-multiaddr"
)

type AddrInfo struct {
	ID    ID
	Addrs []ma.Multiaddr
}

var _ fmt.Stringer = AddrInfo{}

func (pi AddrInfo) String() string { _ = "STUB: not implemented"; return "" }

var ErrInvalidAddr = fmt.Errorf("invalid p2p multiaddr")

func AddrInfosFromP2pAddrs(maddrs ...ma.Multiaddr) ([]AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SplitAddr(m ma.Multiaddr) (transport ma.Multiaddr, id ID) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), *new(ID)
}

func IDFromP2PAddr(m ma.Multiaddr) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func AddrInfoFromString(s string) (*AddrInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func AddrInfoFromP2pAddr(m ma.Multiaddr) (*AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddrInfoToP2pAddrs(pi *AddrInfo) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pi *AddrInfo) Loggable() map[string]any { _ = "STUB: not implemented"; return nil }

func AddrInfosToIDs(pis []AddrInfo) []ID { _ = "STUB: not implemented"; return nil }
