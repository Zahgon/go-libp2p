package relay

import (
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

var (
	errTooManyReservations       = errors.New("too many reservations")
	errTooManyReservationsForIP  = errors.New("too many peers for IP address")
	errTooManyReservationsForASN = errors.New("too many peers for ASN")
)

type peerWithExpiry struct {
	Expiry time.Time
	Peer   peer.ID
}

type constraints struct {
	rc *Resources

	mutex sync.Mutex
	total []peerWithExpiry
	ips   map[string][]peerWithExpiry
	asns  map[uint32][]peerWithExpiry
}

func newConstraints(rc *Resources) *constraints { _ = "STUB: not implemented"; return nil }

func (c *constraints) Reserve(p peer.ID, a ma.Multiaddr, expiry time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *constraints) cleanup(now time.Time) { _ = "STUB: not implemented"; return }

func (c *constraints) cleanupPeer(p peer.ID) { _ = "STUB: not implemented"; return }
