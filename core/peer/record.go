package peer

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/peer/pb"
	"github.com/libp2p/go-libp2p/core/record"

	ma "github.com/multiformats/go-multiaddr"
)

var _ record.Record = (*PeerRecord)(nil)

func init() {
	record.RegisterType(&PeerRecord{})
}

const PeerRecordEnvelopeDomain = "libp2p-peer-record"

var PeerRecordEnvelopePayloadType = []byte{0x03, 0x01}

type PeerRecord struct {
	PeerID ID

	Addrs []ma.Multiaddr

	Seq uint64
}

func NewPeerRecord() *PeerRecord { _ = "STUB: not implemented"; return nil }

func PeerRecordFromAddrInfo(info AddrInfo) *PeerRecord { _ = "STUB: not implemented"; return nil }

func PeerRecordFromProtobuf(msg *pb.PeerRecord) (*PeerRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	lastTimestampMu sync.Mutex
	lastTimestamp   uint64
)

func TimestampSeq() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *PeerRecord) Domain() string { _ = "STUB: not implemented"; return "" }

func (r *PeerRecord) Codec() []byte { _ = "STUB: not implemented"; return nil }

func (r *PeerRecord) UnmarshalRecord(bytes []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeerRecord) MarshalRecord() (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PeerRecord) Equal(other *PeerRecord) bool { _ = "STUB: not implemented"; return false }

func (r *PeerRecord) ToProtobuf() (*pb.PeerRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addrsFromProtobuf(addrs []*pb.PeerRecord_AddressInfo) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func addrsToProtobuf(addrs []ma.Multiaddr) []*pb.PeerRecord_AddressInfo {
	_ = "STUB: not implemented"
	return nil
}
