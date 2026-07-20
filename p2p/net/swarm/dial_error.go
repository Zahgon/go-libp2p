package swarm

import (
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

const maxDialDialErrors = 16

type DialError struct {
	Peer       peer.ID
	DialErrors []TransportError
	Cause      error
	Skipped    int
}

func (e *DialError) Timeout() bool { _ = "STUB: not implemented"; return false }

func (e *DialError) recordErr(addr ma.Multiaddr, err error) { _ = "STUB: not implemented"; return }

func (e *DialError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *DialError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

var _ error = (*DialError)(nil)

type TransportError struct {
	Address ma.Multiaddr
	Cause   error
}

func (e *TransportError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *TransportError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var _ error = (*TransportError)(nil)
