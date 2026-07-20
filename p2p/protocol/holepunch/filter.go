package holepunch

import (
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func WithAddrFilter(f AddrFilter) Option { _ = "STUB: not implemented"; return *new(Option) }

type AddrFilter interface {
	FilterLocal(remoteID peer.ID, maddrs []ma.Multiaddr) []ma.Multiaddr

	FilterRemote(remoteID peer.ID, maddrs []ma.Multiaddr) []ma.Multiaddr
}
