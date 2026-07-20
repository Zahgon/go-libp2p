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

type Status int32

const (
	Status_UNUSED                  Status = 0
	Status_OK                      Status = 100
	Status_RESERVATION_REFUSED     Status = 200
	Status_RESOURCE_LIMIT_EXCEEDED Status = 201
	Status_PERMISSION_DENIED       Status = 202
	Status_CONNECTION_FAILED       Status = 203
	Status_NO_RESERVATION          Status = 204
	Status_MALFORMED_MESSAGE       Status = 400
	Status_UNEXPECTED_MESSAGE      Status = 401
)

var (
	Status_name = map[int32]string{
		0:   "UNUSED",
		100: "OK",
		200: "RESERVATION_REFUSED",
		201: "RESOURCE_LIMIT_EXCEEDED",
		202: "PERMISSION_DENIED",
		203: "CONNECTION_FAILED",
		204: "NO_RESERVATION",
		400: "MALFORMED_MESSAGE",
		401: "UNEXPECTED_MESSAGE",
	}
	Status_value = map[string]int32{
		"UNUSED":                  0,
		"OK":                      100,
		"RESERVATION_REFUSED":     200,
		"RESOURCE_LIMIT_EXCEEDED": 201,
		"PERMISSION_DENIED":       202,
		"CONNECTION_FAILED":       203,
		"NO_RESERVATION":          204,
		"MALFORMED_MESSAGE":       400,
		"UNEXPECTED_MESSAGE":      401,
	}
)

func (x Status) Enum() *Status { _ = "STUB: not implemented"; return nil }

func (x Status) String() string { _ = "STUB: not implemented"; return "" }

func (Status) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (Status) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x Status) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (Status) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type HopMessage_Type int32

const (
	HopMessage_RESERVE HopMessage_Type = 0
	HopMessage_CONNECT HopMessage_Type = 1
	HopMessage_STATUS  HopMessage_Type = 2
)

var (
	HopMessage_Type_name = map[int32]string{
		0: "RESERVE",
		1: "CONNECT",
		2: "STATUS",
	}
	HopMessage_Type_value = map[string]int32{
		"RESERVE": 0,
		"CONNECT": 1,
		"STATUS":  2,
	}
)

func (x HopMessage_Type) Enum() *HopMessage_Type { _ = "STUB: not implemented"; return nil }

func (x HopMessage_Type) String() string { _ = "STUB: not implemented"; return "" }

func (HopMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (HopMessage_Type) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x HopMessage_Type) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (HopMessage_Type) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type StopMessage_Type int32

const (
	StopMessage_CONNECT StopMessage_Type = 0
	StopMessage_STATUS  StopMessage_Type = 1
)

var (
	StopMessage_Type_name = map[int32]string{
		0: "CONNECT",
		1: "STATUS",
	}
	StopMessage_Type_value = map[string]int32{
		"CONNECT": 0,
		"STATUS":  1,
	}
)

func (x StopMessage_Type) Enum() *StopMessage_Type { _ = "STUB: not implemented"; return nil }

func (x StopMessage_Type) String() string { _ = "STUB: not implemented"; return "" }

func (StopMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (StopMessage_Type) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x StopMessage_Type) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (StopMessage_Type) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type HopMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Type          *HopMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=circuit.pb.HopMessage_Type,oneof" json:"type,omitempty"`
	Peer          *Peer            `protobuf:"bytes,2,opt,name=peer,proto3,oneof" json:"peer,omitempty"`
	Reservation   *Reservation     `protobuf:"bytes,3,opt,name=reservation,proto3,oneof" json:"reservation,omitempty"`
	Limit         *Limit           `protobuf:"bytes,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Status        *Status          `protobuf:"varint,5,opt,name=status,proto3,enum=circuit.pb.Status,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HopMessage) Reset() { _ = "STUB: not implemented"; return }

func (x *HopMessage) String() string { _ = "STUB: not implemented"; return "" }

func (*HopMessage) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HopMessage) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HopMessage) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HopMessage) GetType() HopMessage_Type {
	_ = "STUB: not implemented"
	return *new(HopMessage_Type)
}

func (x *HopMessage) GetPeer() *Peer { _ = "STUB: not implemented"; return nil }

func (x *HopMessage) GetReservation() *Reservation { _ = "STUB: not implemented"; return nil }

func (x *HopMessage) GetLimit() *Limit { _ = "STUB: not implemented"; return nil }

func (x *HopMessage) GetStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

type StopMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Type          *StopMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=circuit.pb.StopMessage_Type,oneof" json:"type,omitempty"`
	Peer          *Peer             `protobuf:"bytes,2,opt,name=peer,proto3,oneof" json:"peer,omitempty"`
	Limit         *Limit            `protobuf:"bytes,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Status        *Status           `protobuf:"varint,4,opt,name=status,proto3,enum=circuit.pb.Status,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopMessage) Reset() { _ = "STUB: not implemented"; return }

func (x *StopMessage) String() string { _ = "STUB: not implemented"; return "" }

func (*StopMessage) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *StopMessage) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*StopMessage) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *StopMessage) GetType() StopMessage_Type {
	_ = "STUB: not implemented"
	return *new(StopMessage_Type)
}

func (x *StopMessage) GetPeer() *Peer { _ = "STUB: not implemented"; return nil }

func (x *StopMessage) GetLimit() *Limit { _ = "STUB: not implemented"; return nil }

func (x *StopMessage) GetStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

type Peer struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Id            []byte   `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Addrs         [][]byte `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Peer) Reset() { _ = "STUB: not implemented"; return }

func (x *Peer) String() string { _ = "STUB: not implemented"; return "" }

func (*Peer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Peer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Peer) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Peer) GetId() []byte { _ = "STUB: not implemented"; return nil }

func (x *Peer) GetAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

type Reservation struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Expire        *uint64  `protobuf:"varint,1,opt,name=expire,proto3,oneof" json:"expire,omitempty"`
	Addrs         [][]byte `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	Voucher       []byte   `protobuf:"bytes,3,opt,name=voucher,proto3,oneof" json:"voucher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reservation) Reset() { _ = "STUB: not implemented"; return }

func (x *Reservation) String() string { _ = "STUB: not implemented"; return "" }

func (*Reservation) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Reservation) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Reservation) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Reservation) GetExpire() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *Reservation) GetAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

func (x *Reservation) GetVoucher() []byte { _ = "STUB: not implemented"; return nil }

type Limit struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      *uint32                `protobuf:"varint,1,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	Data          *uint64                `protobuf:"varint,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Limit) Reset() { _ = "STUB: not implemented"; return }

func (x *Limit) String() string { _ = "STUB: not implemented"; return "" }

func (*Limit) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Limit) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Limit) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Limit) GetDuration() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *Limit) GetData() uint64 { _ = "STUB: not implemented"; return 0 }

var File_p2p_protocol_circuitv2_pb_circuit_proto protoreflect.FileDescriptor

const file_p2p_protocol_circuitv2_pb_circuit_proto_rawDesc = "" +
	"\n" +
	"'p2p/protocol/circuitv2/pb/circuit.proto\x12\n" +
	"circuit.pb\"\xf1\x02\n" +
	"\n" +
	"HopMessage\x124\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1b.circuit.pb.HopMessage.TypeH\x00R\x04type\x88\x01\x01\x12)\n" +
	"\x04peer\x18\x02 \x01(\v2\x10.circuit.pb.PeerH\x01R\x04peer\x88\x01\x01\x12>\n" +
	"\vreservation\x18\x03 \x01(\v2\x17.circuit.pb.ReservationH\x02R\vreservation\x88\x01\x01\x12,\n" +
	"\x05limit\x18\x04 \x01(\v2\x11.circuit.pb.LimitH\x03R\x05limit\x88\x01\x01\x12/\n" +
	"\x06status\x18\x05 \x01(\x0e2\x12.circuit.pb.StatusH\x04R\x06status\x88\x01\x01\",\n" +
	"\x04Type\x12\v\n" +
	"\aRESERVE\x10\x00\x12\v\n" +
	"\aCONNECT\x10\x01\x12\n" +
	"\n" +
	"\x06STATUS\x10\x02B\a\n" +
	"\x05_typeB\a\n" +
	"\x05_peerB\x0e\n" +
	"\f_reservationB\b\n" +
	"\x06_limitB\t\n" +
	"\a_status\"\x96\x02\n" +
	"\vStopMessage\x125\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1c.circuit.pb.StopMessage.TypeH\x00R\x04type\x88\x01\x01\x12)\n" +
	"\x04peer\x18\x02 \x01(\v2\x10.circuit.pb.PeerH\x01R\x04peer\x88\x01\x01\x12,\n" +
	"\x05limit\x18\x03 \x01(\v2\x11.circuit.pb.LimitH\x02R\x05limit\x88\x01\x01\x12/\n" +
	"\x06status\x18\x04 \x01(\x0e2\x12.circuit.pb.StatusH\x03R\x06status\x88\x01\x01\"\x1f\n" +
	"\x04Type\x12\v\n" +
	"\aCONNECT\x10\x00\x12\n" +
	"\n" +
	"\x06STATUS\x10\x01B\a\n" +
	"\x05_typeB\a\n" +
	"\x05_peerB\b\n" +
	"\x06_limitB\t\n" +
	"\a_status\"8\n" +
	"\x04Peer\x12\x13\n" +
	"\x02id\x18\x01 \x01(\fH\x00R\x02id\x88\x01\x01\x12\x14\n" +
	"\x05addrs\x18\x02 \x03(\fR\x05addrsB\x05\n" +
	"\x03_id\"v\n" +
	"\vReservation\x12\x1b\n" +
	"\x06expire\x18\x01 \x01(\x04H\x00R\x06expire\x88\x01\x01\x12\x14\n" +
	"\x05addrs\x18\x02 \x03(\fR\x05addrs\x12\x1d\n" +
	"\avoucher\x18\x03 \x01(\fH\x01R\avoucher\x88\x01\x01B\t\n" +
	"\a_expireB\n" +
	"\n" +
	"\b_voucher\"W\n" +
	"\x05Limit\x12\x1f\n" +
	"\bduration\x18\x01 \x01(\rH\x00R\bduration\x88\x01\x01\x12\x17\n" +
	"\x04data\x18\x02 \x01(\x04H\x01R\x04data\x88\x01\x01B\v\n" +
	"\t_durationB\a\n" +
	"\x05_data*\xca\x01\n" +
	"\x06Status\x12\n" +
	"\n" +
	"\x06UNUSED\x10\x00\x12\x06\n" +
	"\x02OK\x10d\x12\x18\n" +
	"\x13RESERVATION_REFUSED\x10\xc8\x01\x12\x1c\n" +
	"\x17RESOURCE_LIMIT_EXCEEDED\x10\xc9\x01\x12\x16\n" +
	"\x11PERMISSION_DENIED\x10\xca\x01\x12\x16\n" +
	"\x11CONNECTION_FAILED\x10\xcb\x01\x12\x13\n" +
	"\x0eNO_RESERVATION\x10\xcc\x01\x12\x16\n" +
	"\x11MALFORMED_MESSAGE\x10\x90\x03\x12\x17\n" +
	"\x12UNEXPECTED_MESSAGE\x10\x91\x03B7Z5github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pbb\x06proto3"

var (
	file_p2p_protocol_circuitv2_pb_circuit_proto_rawDescOnce sync.Once
	file_p2p_protocol_circuitv2_pb_circuit_proto_rawDescData []byte
)

func file_p2p_protocol_circuitv2_pb_circuit_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_protocol_circuitv2_pb_circuit_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_p2p_protocol_circuitv2_pb_circuit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_p2p_protocol_circuitv2_pb_circuit_proto_goTypes = []any{
	(Status)(0),
	(HopMessage_Type)(0),
	(StopMessage_Type)(0),
	(*HopMessage)(nil),
	(*StopMessage)(nil),
	(*Peer)(nil),
	(*Reservation)(nil),
	(*Limit)(nil),
}
var file_p2p_protocol_circuitv2_pb_circuit_proto_depIdxs = []int32{
	1,
	5,
	6,
	7,
	0,
	2,
	5,
	7,
	0,
	9,
	9,
	9,
	9,
	0,
}

func init()                                              { file_p2p_protocol_circuitv2_pb_circuit_proto_init() }
func file_p2p_protocol_circuitv2_pb_circuit_proto_init() { _ = "STUB: not implemented"; return }
