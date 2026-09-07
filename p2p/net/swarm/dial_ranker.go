package swarm

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	PublicTCPDelay  = 250 * time.Millisecond
	PrivateTCPDelay = 30 * time.Millisecond

	PublicQUICDelay  = 250 * time.Millisecond
	PrivateQUICDelay = 30 * time.Millisecond

	RelayDelay = 500 * time.Millisecond

	PublicOtherDelay  = 1 * time.Second
	PrivateOtherDelay = 100 * time.Millisecond
)

func NoDelayDialRanker(addrs []ma.Multiaddr) []network.AddrDelay {
	_ = "STUB: not implemented"
	return nil
}

func DefaultDialRanker(addrs []ma.Multiaddr) []network.AddrDelay {
	_ = "STUB: not implemented"
	return nil
}

func getAddrDelay(addrs []ma.Multiaddr, tcpDelay time.Duration, quicDelay time.Duration,
	otherDelay time.Duration, offset time.Duration) []network.AddrDelay {
	_ = "STUB: not implemented"
	return nil
}

func score(a ma.Multiaddr) int { _ = "STUB: not implemented"; return 0 }

func isProtocolAddr(a ma.Multiaddr, p int) bool { _ = "STUB: not implemented"; return false }

func isQUICAddr(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func filterAddrs(addrs []ma.Multiaddr, f func(a ma.Multiaddr) bool) (filtered, rest []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil, nil
}
