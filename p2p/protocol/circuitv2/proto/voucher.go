package proto

import (
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/record"
)

const RecordDomain = "libp2p-relay-rsvp"

var RecordCodec = []byte{0x03, 0x02}

func init() {
	record.RegisterType(&ReservationVoucher{})
}

type ReservationVoucher struct {
	Relay peer.ID

	Peer peer.ID

	Expiration time.Time
}

var _ record.Record = (*ReservationVoucher)(nil)

func (rv *ReservationVoucher) Domain() string { _ = "STUB: not implemented"; return "" }

func (rv *ReservationVoucher) Codec() []byte { _ = "STUB: not implemented"; return nil }

func (rv *ReservationVoucher) MarshalRecord() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rv *ReservationVoucher) UnmarshalRecord(blob []byte) error {
	_ = "STUB: not implemented"
	return nil
}
