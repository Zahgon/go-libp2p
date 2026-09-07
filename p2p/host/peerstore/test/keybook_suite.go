package test

import (
	"testing"

	pstore "github.com/libp2p/go-libp2p/core/peerstore"
)

var keyBookSuite = map[string]func(kb pstore.KeyBook) func(*testing.T){
	"AddGetPrivKey":         testKeybookPrivKey,
	"AddGetPubKey":          testKeyBookPubKey,
	"PeersWithKeys":         testKeyBookPeers,
	"PubKeyAddedOnRetrieve": testInlinedPubKeyAddedOnRetrieve,
	"Delete":                testKeyBookDelete,
}

type KeyBookFactory func() (pstore.KeyBook, func())

func TestKeyBook(t *testing.T, factory KeyBookFactory) { _ = "STUB: not implemented"; return }

func testKeybookPrivKey(kb pstore.KeyBook) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testKeyBookPubKey(kb pstore.KeyBook) func(t *testing.T) { _ = "STUB: not implemented"; return nil }

func testKeyBookPeers(kb pstore.KeyBook) func(t *testing.T) { _ = "STUB: not implemented"; return nil }

func testInlinedPubKeyAddedOnRetrieve(kb pstore.KeyBook) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testKeyBookDelete(kb pstore.KeyBook) func(t *testing.T) { _ = "STUB: not implemented"; return nil }

var keybookBenchmarkSuite = map[string]func(kb pstore.KeyBook) func(*testing.B){
	"PubKey":        benchmarkPubKey,
	"AddPubKey":     benchmarkAddPubKey,
	"PrivKey":       benchmarkPrivKey,
	"AddPrivKey":    benchmarkAddPrivKey,
	"PeersWithKeys": benchmarkPeersWithKeys,
}

func BenchmarkKeyBook(b *testing.B, factory KeyBookFactory) { _ = "STUB: not implemented"; return }

func benchmarkPubKey(kb pstore.KeyBook) func(*testing.B) { _ = "STUB: not implemented"; return nil }

func benchmarkAddPubKey(kb pstore.KeyBook) func(*testing.B) { _ = "STUB: not implemented"; return nil }

func benchmarkPrivKey(kb pstore.KeyBook) func(*testing.B) { _ = "STUB: not implemented"; return nil }

func benchmarkAddPrivKey(kb pstore.KeyBook) func(*testing.B) { _ = "STUB: not implemented"; return nil }

func benchmarkPeersWithKeys(kb pstore.KeyBook) func(*testing.B) {
	_ = "STUB: not implemented"
	return nil
}
