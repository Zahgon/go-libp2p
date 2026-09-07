package event

import (
	"github.com/libp2p/go-libp2p/core/network"
	ma "github.com/multiformats/go-multiaddr"
)

type EvtLocalReachabilityChanged struct {
	Reachability network.Reachability
}

type EvtHostReachableAddrsChanged struct {
	Reachable   []ma.Multiaddr
	Unreachable []ma.Multiaddr
	Unknown     []ma.Multiaddr
}
