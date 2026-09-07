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

type HolePunch_Type int32

const (
	HolePunch_CONNECT HolePunch_Type = 100
	HolePunch_SYNC    HolePunch_Type = 300
)

var (
	HolePunch_Type_name = map[int32]string{
		100: "CONNECT",
		300: "SYNC",
	}
	HolePunch_Type_value = map[string]int32{
		"CONNECT": 100,
		"SYNC":    300,
	}
)

func (x HolePunch_Type) Enum() *HolePunch_Type { _ = "STUB: not implemented"; return nil }

func (x HolePunch_Type) String() string { _ = "STUB: not implemented"; return "" }

func (HolePunch_Type) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (HolePunch_Type) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x HolePunch_Type) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (x *HolePunch_Type) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (HolePunch_Type) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type HolePunch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *HolePunch_Type        `protobuf:"varint,1,req,name=type,enum=holepunch.pb.HolePunch_Type" json:"type,omitempty"`
	ObsAddrs      [][]byte               `protobuf:"bytes,2,rep,name=ObsAddrs" json:"ObsAddrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HolePunch) Reset() { _ = "STUB: not implemented"; return }

func (x *HolePunch) String() string { _ = "STUB: not implemented"; return "" }

func (*HolePunch) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HolePunch) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HolePunch) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HolePunch) GetType() HolePunch_Type {
	_ = "STUB: not implemented"
	return *new(HolePunch_Type)
}

func (x *HolePunch) GetObsAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

var File_p2p_protocol_holepunch_pb_holepunch_proto protoreflect.FileDescriptor

const file_p2p_protocol_holepunch_pb_holepunch_proto_rawDesc = "" +
	"\n" +
	")p2p/protocol/holepunch/pb/holepunch.proto\x12\fholepunch.pb\"y\n" +
	"\tHolePunch\x120\n" +
	"\x04type\x18\x01 \x02(\x0e2\x1c.holepunch.pb.HolePunch.TypeR\x04type\x12\x1a\n" +
	"\bObsAddrs\x18\x02 \x03(\fR\bObsAddrs\"\x1e\n" +
	"\x04Type\x12\v\n" +
	"\aCONNECT\x10d\x12\t\n" +
	"\x04SYNC\x10\xac\x02B7Z5github.com/libp2p/go-libp2p/p2p/protocol/holepunch/pb"

var (
	file_p2p_protocol_holepunch_pb_holepunch_proto_rawDescOnce sync.Once
	file_p2p_protocol_holepunch_pb_holepunch_proto_rawDescData []byte
)

func file_p2p_protocol_holepunch_pb_holepunch_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_protocol_holepunch_pb_holepunch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_p2p_protocol_holepunch_pb_holepunch_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_p2p_protocol_holepunch_pb_holepunch_proto_goTypes = []any{
	(HolePunch_Type)(0),
	(*HolePunch)(nil),
}
var file_p2p_protocol_holepunch_pb_holepunch_proto_depIdxs = []int32{
	0,
	1,
	1,
	1,
	1,
	0,
}

func init()                                                { file_p2p_protocol_holepunch_pb_holepunch_proto_init() }
func file_p2p_protocol_holepunch_pb_holepunch_proto_init() { _ = "STUB: not implemented"; return }
