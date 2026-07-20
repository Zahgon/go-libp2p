package swarm

import (
	ma "github.com/multiformats/go-multiaddr"
)

type OrderedListener interface {
	ListenOrder() int
}

func (s *Swarm) Listen(addrs ...ma.Multiaddr) error { _ = "STUB: not implemented"; return nil }

func (s *Swarm) ListenClose(addrs ...ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (s *Swarm) AddListenAddr(a ma.Multiaddr) error { _ = "STUB: not implemented"; return nil }

func containsMultiaddr(addrs []ma.Multiaddr, addr ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}
