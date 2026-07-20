package holepunch

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

func removeRelayAddrs(addrs []ma.Multiaddr) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func isRelayAddress(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func addrsToBytes(as []ma.Multiaddr) [][]byte { _ = "STUB: not implemented"; return nil }

func addrsFromBytes(bzs [][]byte) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func getDirectConnection(h host.Host, p peer.ID) network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}

func holePunchConnect(ctx context.Context, host host.Host, pi peer.AddrInfo, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}
