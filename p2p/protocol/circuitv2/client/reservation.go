package client

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	pbv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pb"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/proto"

	ma "github.com/multiformats/go-multiaddr"
)

var ReserveTimeout = time.Minute

type Reservation struct {
	Expiration time.Time

	Addrs []ma.Multiaddr

	LimitDuration time.Duration

	LimitData uint64

	Voucher *proto.ReservationVoucher
}

type ReservationError struct {
	Status pbv2.Status

	Reason string

	err error
}

func (re ReservationError) Error() string { _ = "STUB: not implemented"; return "" }

func (re ReservationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func Reserve(ctx context.Context, h host.Host, ai peer.AddrInfo) (*Reservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
