package test

import (
	"testing"

	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"

	ma "github.com/multiformats/go-multiaddr"
)

var peerstoreSuite = map[string]func(pstore.Peerstore) func(*testing.T){
	"AddrStream":               testAddrStream,
	"GetStreamBeforePeerAdded": testGetStreamBeforePeerAdded,
	"AddStreamDuplicates":      testAddrStreamDuplicates,
	"PeerstoreProtoStore":      testPeerstoreProtoStore,
	"BasicPeerstore":           testBasicPeerstore,
	"Metadata":                 testMetadata,
	"CertifiedAddrBook":        testCertifiedAddrBook,
}

type PeerstoreFactory func() (pstore.Peerstore, func())

func TestPeerstore(t *testing.T, factory PeerstoreFactory) { _ = "STUB: not implemented"; return }

func sortProtos(protos []protocol.ID) { _ = "STUB: not implemented"; return }

func testAddrStream(ps pstore.Peerstore) func(t *testing.T) { _ = "STUB: not implemented"; return nil }

func testGetStreamBeforePeerAdded(ps pstore.Peerstore) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testAddrStreamDuplicates(ps pstore.Peerstore) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testPeerstoreProtoStore(ps pstore.Peerstore) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testBasicPeerstore(ps pstore.Peerstore) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testMetadata(ps pstore.Peerstore) func(t *testing.T) { _ = "STUB: not implemented"; return nil }

func testCertifiedAddrBook(ps pstore.Peerstore) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func getAddrs(t *testing.T, n int) []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func TestPeerstoreProtoStoreLimits(t *testing.T, ps pstore.Peerstore, limit int) {
	_ = "STUB: not implemented"
	return
}
