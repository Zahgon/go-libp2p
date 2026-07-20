package test

import (
	"testing"

	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

func Multiaddr(m string) ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

type peerpair struct {
	ID   peer.ID
	Addr []ma.Multiaddr
}

func RandomPeer(b *testing.B, addrCount int) *peerpair { _ = "STUB: not implemented"; return nil }

func getPeerPairs(b *testing.B, n int, addrsPerPeer int) []*peerpair {
	_ = "STUB: not implemented"
	return nil
}

func GenerateAddrs(count int) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func GeneratePeerIDs(count int) []peer.ID { _ = "STUB: not implemented"; return nil }

func AssertAddressesEqual(t *testing.T, exp, act []ma.Multiaddr) { _ = "STUB: not implemented"; return }
