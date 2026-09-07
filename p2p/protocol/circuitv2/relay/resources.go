package relay

import (
	"time"
)

type Resources struct {
	Limit *RelayLimit

	ReservationTTL time.Duration

	MaxReservations int

	MaxCircuits int

	BufferSize int

	MaxReservationsPerPeer int

	MaxReservationsPerIP int

	MaxReservationsPerASN int
}

type RelayLimit struct {
	Duration time.Duration

	Data int64
}

func DefaultResources() Resources { _ = "STUB: not implemented"; return *new(Resources) }

func DefaultLimit() *RelayLimit { _ = "STUB: not implemented"; return nil }
