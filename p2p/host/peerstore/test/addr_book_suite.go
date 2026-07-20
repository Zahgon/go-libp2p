package test

import (
	"testing"

	pstore "github.com/libp2p/go-libp2p/core/peerstore"

	mockClock "github.com/benbjohnson/clock"
)

var addressBookSuite = map[string]func(book pstore.AddrBook, clk *mockClock.Mock) func(*testing.T){
	"AddAddress":           testAddAddress,
	"Clear":                testClearWorks,
	"SetNegativeTTLClears": testSetNegativeTTLClears,
	"UpdateTTLs":           testUpdateTTLs,
	"NilAddrsDontBreak":    testNilAddrsDontBreak,
	"AddressesExpire":      testAddressesExpire,
	"ClearWithIter":        testClearWithIterator,
	"PeersWithAddresses":   testPeersWithAddrs,
	"CertifiedAddresses":   testCertifiedAddresses,
}

type AddrBookFactory func() (pstore.AddrBook, func())

func TestAddrBook(t *testing.T, factory AddrBookFactory, clk *mockClock.Mock) {
	_ = "STUB: not implemented"
	return
}

func testAddAddress(ab pstore.AddrBook, clk *mockClock.Mock) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testClearWorks(ab pstore.AddrBook, _ *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testSetNegativeTTLClears(m pstore.AddrBook, _ *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testUpdateTTLs(m pstore.AddrBook, clk *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testNilAddrsDontBreak(m pstore.AddrBook, _ *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testAddressesExpire(m pstore.AddrBook, clk *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testClearWithIterator(m pstore.AddrBook, _ *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testPeersWithAddrs(m pstore.AddrBook, _ *mockClock.Mock) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func testCertifiedAddresses(m pstore.AddrBook, clk *mockClock.Mock) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}
