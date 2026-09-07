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

type Message_MessageType int32

const (
	Message_DIAL          Message_MessageType = 0
	Message_DIAL_RESPONSE Message_MessageType = 1
)

var (
	Message_MessageType_name = map[int32]string{
		0: "DIAL",
		1: "DIAL_RESPONSE",
	}
	Message_MessageType_value = map[string]int32{
		"DIAL":          0,
		"DIAL_RESPONSE": 1,
	}
)

func (x Message_MessageType) Enum() *Message_MessageType { _ = "STUB: not implemented"; return nil }

func (x Message_MessageType) String() string { _ = "STUB: not implemented"; return "" }

func (Message_MessageType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (Message_MessageType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x Message_MessageType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (x *Message_MessageType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Message_ResponseStatus int32

const (
	Message_OK               Message_ResponseStatus = 0
	Message_E_DIAL_ERROR     Message_ResponseStatus = 100
	Message_E_DIAL_REFUSED   Message_ResponseStatus = 101
	Message_E_BAD_REQUEST    Message_ResponseStatus = 200
	Message_E_INTERNAL_ERROR Message_ResponseStatus = 300
)

var (
	Message_ResponseStatus_name = map[int32]string{
		0:   "OK",
		100: "E_DIAL_ERROR",
		101: "E_DIAL_REFUSED",
		200: "E_BAD_REQUEST",
		300: "E_INTERNAL_ERROR",
	}
	Message_ResponseStatus_value = map[string]int32{
		"OK":               0,
		"E_DIAL_ERROR":     100,
		"E_DIAL_REFUSED":   101,
		"E_BAD_REQUEST":    200,
		"E_INTERNAL_ERROR": 300,
	}
)

func (x Message_ResponseStatus) Enum() *Message_ResponseStatus {
	_ = "STUB: not implemented"
	return nil
}

func (x Message_ResponseStatus) String() string { _ = "STUB: not implemented"; return "" }

func (Message_ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (Message_ResponseStatus) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x Message_ResponseStatus) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (x *Message_ResponseStatus) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (Message_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *Message_MessageType   `protobuf:"varint,1,opt,name=type,enum=autonat.pb.Message_MessageType" json:"type,omitempty"`
	Dial          *Message_Dial          `protobuf:"bytes,2,opt,name=dial" json:"dial,omitempty"`
	DialResponse  *Message_DialResponse  `protobuf:"bytes,3,opt,name=dialResponse" json:"dialResponse,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() { _ = "STUB: not implemented"; return }

func (x *Message) String() string { _ = "STUB: not implemented"; return "" }

func (*Message) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Message) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Message) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Message) GetType() Message_MessageType {
	_ = "STUB: not implemented"
	return *new(Message_MessageType)
}

func (x *Message) GetDial() *Message_Dial { _ = "STUB: not implemented"; return nil }

func (x *Message) GetDialResponse() *Message_DialResponse { _ = "STUB: not implemented"; return nil }

type Message_PeerInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            []byte                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Addrs         [][]byte               `protobuf:"bytes,2,rep,name=addrs" json:"addrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message_PeerInfo) Reset() { _ = "STUB: not implemented"; return }

func (x *Message_PeerInfo) String() string { _ = "STUB: not implemented"; return "" }

func (*Message_PeerInfo) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Message_PeerInfo) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Message_PeerInfo) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Message_PeerInfo) GetId() []byte { _ = "STUB: not implemented"; return nil }

func (x *Message_PeerInfo) GetAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

type Message_Dial struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Peer          *Message_PeerInfo      `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message_Dial) Reset() { _ = "STUB: not implemented"; return }

func (x *Message_Dial) String() string { _ = "STUB: not implemented"; return "" }

func (*Message_Dial) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Message_Dial) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Message_Dial) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Message_Dial) GetPeer() *Message_PeerInfo { _ = "STUB: not implemented"; return nil }

type Message_DialResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Status        *Message_ResponseStatus `protobuf:"varint,1,opt,name=status,enum=autonat.pb.Message_ResponseStatus" json:"status,omitempty"`
	StatusText    *string                 `protobuf:"bytes,2,opt,name=statusText" json:"statusText,omitempty"`
	Addr          []byte                  `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message_DialResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *Message_DialResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*Message_DialResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Message_DialResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Message_DialResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *Message_DialResponse) GetStatus() Message_ResponseStatus {
	_ = "STUB: not implemented"
	return *new(Message_ResponseStatus)
}

func (x *Message_DialResponse) GetStatusText() string { _ = "STUB: not implemented"; return "" }

func (x *Message_DialResponse) GetAddr() []byte { _ = "STUB: not implemented"; return nil }

var File_p2p_host_autonat_pb_autonat_proto protoreflect.FileDescriptor

const file_p2p_host_autonat_pb_autonat_proto_rawDesc = "" +
	"\n" +
	"!p2p/host/autonat/pb/autonat.proto\x12\n" +
	"autonat.pb\"\xb5\x04\n" +
	"\aMessage\x123\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1f.autonat.pb.Message.MessageTypeR\x04type\x12,\n" +
	"\x04dial\x18\x02 \x01(\v2\x18.autonat.pb.Message.DialR\x04dial\x12D\n" +
	"\fdialResponse\x18\x03 \x01(\v2 .autonat.pb.Message.DialResponseR\fdialResponse\x1a0\n" +
	"\bPeerInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\fR\x02id\x12\x14\n" +
	"\x05addrs\x18\x02 \x03(\fR\x05addrs\x1a8\n" +
	"\x04Dial\x120\n" +
	"\x04peer\x18\x01 \x01(\v2\x1c.autonat.pb.Message.PeerInfoR\x04peer\x1a~\n" +
	"\fDialResponse\x12:\n" +
	"\x06status\x18\x01 \x01(\x0e2\".autonat.pb.Message.ResponseStatusR\x06status\x12\x1e\n" +
	"\n" +
	"statusText\x18\x02 \x01(\tR\n" +
	"statusText\x12\x12\n" +
	"\x04addr\x18\x03 \x01(\fR\x04addr\"*\n" +
	"\vMessageType\x12\b\n" +
	"\x04DIAL\x10\x00\x12\x11\n" +
	"\rDIAL_RESPONSE\x10\x01\"i\n" +
	"\x0eResponseStatus\x12\x06\n" +
	"\x02OK\x10\x00\x12\x10\n" +
	"\fE_DIAL_ERROR\x10d\x12\x12\n" +
	"\x0eE_DIAL_REFUSED\x10e\x12\x12\n" +
	"\rE_BAD_REQUEST\x10\xc8\x01\x12\x15\n" +
	"\x10E_INTERNAL_ERROR\x10\xac\x02B1Z/github.com/libp2p/go-libp2p/p2p/host/autonat/pb"

var (
	file_p2p_host_autonat_pb_autonat_proto_rawDescOnce sync.Once
	file_p2p_host_autonat_pb_autonat_proto_rawDescData []byte
)

func file_p2p_host_autonat_pb_autonat_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_host_autonat_pb_autonat_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_p2p_host_autonat_pb_autonat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_p2p_host_autonat_pb_autonat_proto_goTypes = []any{
	(Message_MessageType)(0),
	(Message_ResponseStatus)(0),
	(*Message)(nil),
	(*Message_PeerInfo)(nil),
	(*Message_Dial)(nil),
	(*Message_DialResponse)(nil),
}
var file_p2p_host_autonat_pb_autonat_proto_depIdxs = []int32{
	0,
	4,
	5,
	3,
	1,
	5,
	5,
	5,
	5,
	0,
}

func init()                                        { file_p2p_host_autonat_pb_autonat_proto_init() }
func file_p2p_host_autonat_pb_autonat_proto_init() { _ = "STUB: not implemented"; return }
