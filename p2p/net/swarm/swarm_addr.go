package swarm

import (
	"time"

	ma "github.com/multiformats/go-multiaddr"
)

func (s *Swarm) ListenAddresses() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (s *Swarm) listenAddressesNoLock() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

const ifaceAddrsCacheDuration = 1 * time.Minute

func (s *Swarm) InterfaceListenAddresses() ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
