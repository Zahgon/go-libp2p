package autorelay

import (
	ma "github.com/multiformats/go-multiaddr"
)

func cleanupAddressSet(addrs []ma.Multiaddr) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func isRelayAddr(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func hasAddrsplosion(addrs []ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func addrKeyAndPort(a ma.Multiaddr) (string, int) { _ = "STUB: not implemented"; return "", 0 }

func sanitizeAddrsplodedSet(public, private []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}
