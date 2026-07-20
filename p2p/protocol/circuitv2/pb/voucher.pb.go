package pb

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReservationVoucher struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Relay         []byte  `protobuf:"bytes,1,opt,name=relay,proto3,oneof" json:"relay,omitempty"`
	Peer          []byte  `protobuf:"bytes,2,opt,name=peer,proto3,oneof" json:"peer,omitempty"`
	Expiration    *uint64 `protobuf:"varint,3,opt,name=expiration,proto3,oneof" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReservationVoucher) Reset() { _ = "STUB: not implemented"; return }

func (x *ReservationVoucher) String() string { _ = "STUB: not implemented"; return "" }

func (*ReservationVoucher) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ReservationVoucher) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ReservationVoucher) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ReservationVoucher) GetRelay() []byte { _ = "STUB: not implemented"; return nil }

func (x *ReservationVoucher) GetPeer() []byte { _ = "STUB: not implemented"; return nil }

func (x *ReservationVoucher) GetExpiration() uint64 { _ = "STUB: not implemented"; return 0 }

var File_p2p_protocol_circuitv2_pb_voucher_proto protoreflect.FileDescriptor

const file_p2p_protocol_circuitv2_pb_voucher_proto_rawDesc = "" +
	"\n" +
	"'p2p/protocol/circuitv2/pb/voucher.proto\x12\n" +
	"circuit.pb\"\x8f\x01\n" +
	"\x12ReservationVoucher\x12\x19\n" +
	"\x05relay\x18\x01 \x01(\fH\x00R\x05relay\x88\x01\x01\x12\x17\n" +
	"\x04peer\x18\x02 \x01(\fH\x01R\x04peer\x88\x01\x01\x12#\n" +
	"\n" +
	"expiration\x18\x03 \x01(\x04H\x02R\n" +
	"expiration\x88\x01\x01B\b\n" +
	"\x06_relayB\a\n" +
	"\x05_peerB\r\n" +
	"\v_expirationB7Z5github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pbb\x06proto3"

var (
	file_p2p_protocol_circuitv2_pb_voucher_proto_rawDescOnce sync.Once
	file_p2p_protocol_circuitv2_pb_voucher_proto_rawDescData []byte
)

func file_p2p_protocol_circuitv2_pb_voucher_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_protocol_circuitv2_pb_voucher_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_p2p_protocol_circuitv2_pb_voucher_proto_goTypes = []any{
	(*ReservationVoucher)(nil),
}
var file_p2p_protocol_circuitv2_pb_voucher_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}

func init()                                              { file_p2p_protocol_circuitv2_pb_voucher_proto_init() }
func file_p2p_protocol_circuitv2_pb_voucher_proto_init() { _ = "STUB: not implemented"; return }
