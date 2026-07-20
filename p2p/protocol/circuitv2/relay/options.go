package relay

import (
	"github.com/multiformats/go-multiaddr"
)

type Option func(*Relay) error

func WithResources(rc Resources) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLimit(limit *RelayLimit) Option { _ = "STUB: not implemented"; return *new(Option) }

type ReservationAddressFilterFunc func(addr multiaddr.Multiaddr) (include bool)

func WithReservationAddressFilter(filter ReservationAddressFilterFunc) (option Option) {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithInfiniteLimits() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithACL(acl ACLFilter) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetricsTracer(mt MetricsTracer) Option { _ = "STUB: not implemented"; return *new(Option) }
