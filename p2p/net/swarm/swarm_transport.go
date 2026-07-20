package swarm

import (
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

func (s *Swarm) TransportForDialing(a ma.Multiaddr) transport.Transport {
	_ = "STUB: not implemented"
	return *new(transport.Transport)
}

func (s *Swarm) TransportForListening(a ma.Multiaddr) transport.Transport {
	_ = "STUB: not implemented"
	return *new(transport.Transport)
}

func (s *Swarm) AddTransport(t transport.Transport) error { _ = "STUB: not implemented"; return nil }
