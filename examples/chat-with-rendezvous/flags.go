package main

import (
	maddr "github.com/multiformats/go-multiaddr"
)

type addrList []maddr.Multiaddr

func (al *addrList) String() string { _ = "STUB: not implemented"; return "" }

func (al *addrList) Set(value string) error { _ = "STUB: not implemented"; return nil }

func StringsToAddrs(addrStrings []string) (maddrs []maddr.Multiaddr, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Config struct {
	RendezvousString string
	BootstrapPeers   addrList
	ListenAddresses  addrList
	ProtocolID       string
}

func ParseFlags() (Config, error) { _ = "STUB: not implemented"; return *new(Config), nil }
