package event

import (
	"github.com/libp2p/go-libp2p/core/record"

	ma "github.com/multiformats/go-multiaddr"
)

type AddrAction int

const (
	Unknown AddrAction = iota

	Added

	Maintained

	Removed
)

type UpdatedAddress struct {
	Address ma.Multiaddr

	Action AddrAction
}

type EvtLocalAddressesUpdated struct {
	Diffs bool

	Current []UpdatedAddress

	Removed []UpdatedAddress

	SignedPeerRecord *record.Envelope
}

type EvtAutoRelayAddrsUpdated struct {
	RelayAddrs []ma.Multiaddr
}
