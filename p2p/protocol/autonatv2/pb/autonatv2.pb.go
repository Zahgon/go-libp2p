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

type DialStatus int32

const (
	DialStatus_UNUSED            DialStatus = 0
	DialStatus_E_DIAL_ERROR      DialStatus = 100
	DialStatus_E_DIAL_BACK_ERROR DialStatus = 101
	DialStatus_OK                DialStatus = 200
)

var (
	DialStatus_name = map[int32]string{
		0:   "UNUSED",
		100: "E_DIAL_ERROR",
		101: "E_DIAL_BACK_ERROR",
		200: "OK",
	}
	DialStatus_value = map[string]int32{
		"UNUSED":            0,
		"E_DIAL_ERROR":      100,
		"E_DIAL_BACK_ERROR": 101,
		"OK":                200,
	}
)

func (x DialStatus) Enum() *DialStatus { _ = "STUB: not implemented"; return nil }

func (x DialStatus) String() string { _ = "STUB: not implemented"; return "" }

func (DialStatus) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (DialStatus) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x DialStatus) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (DialStatus) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type DialResponse_ResponseStatus int32

const (
	DialResponse_E_INTERNAL_ERROR   DialResponse_ResponseStatus = 0
	DialResponse_E_REQUEST_REJECTED DialResponse_ResponseStatus = 100
	DialResponse_E_DIAL_REFUSED     DialResponse_ResponseStatus = 101
	DialResponse_OK                 DialResponse_ResponseStatus = 200
)

var (
	DialResponse_ResponseStatus_name = map[int32]string{
		0:   "E_INTERNAL_ERROR",
		100: "E_REQUEST_REJECTED",
		101: "E_DIAL_REFUSED",
		200: "OK",
	}
	DialResponse_ResponseStatus_value = map[string]int32{
		"E_INTERNAL_ERROR":   0,
		"E_REQUEST_REJECTED": 100,
		"E_DIAL_REFUSED":     101,
		"OK":                 200,
	}
)

func (x DialResponse_ResponseStatus) Enum() *DialResponse_ResponseStatus {
	_ = "STUB: not implemented"
	return nil
}

func (x DialResponse_ResponseStatus) String() string { _ = "STUB: not implemented"; return "" }

func (DialResponse_ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (DialResponse_ResponseStatus) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x DialResponse_ResponseStatus) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (DialResponse_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DialBackResponse_DialBackStatus int32

const (
	DialBackResponse_OK DialBackResponse_DialBackStatus = 0
)

var (
	DialBackResponse_DialBackStatus_name = map[int32]string{
		0: "OK",
	}
	DialBackResponse_DialBackStatus_value = map[string]int32{
		"OK": 0,
	}
)

func (x DialBackResponse_DialBackStatus) Enum() *DialBackResponse_DialBackStatus {
	_ = "STUB: not implemented"
	return nil
}

func (x DialBackResponse_DialBackStatus) String() string { _ = "STUB: not implemented"; return "" }

func (DialBackResponse_DialBackStatus) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (DialBackResponse_DialBackStatus) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x DialBackResponse_DialBackStatus) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (DialBackResponse_DialBackStatus) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Msg           isMessage_Msg `protobuf_oneof:"msg"`
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

func (x *Message) GetMsg() isMessage_Msg { _ = "STUB: not implemented"; return *new(isMessage_Msg) }

func (x *Message) GetDialRequest() *DialRequest { _ = "STUB: not implemented"; return nil }

func (x *Message) GetDialResponse() *DialResponse { _ = "STUB: not implemented"; return nil }

func (x *Message) GetDialDataRequest() *DialDataRequest { _ = "STUB: not implemented"; return nil }

func (x *Message) GetDialDataResponse() *DialDataResponse { _ = "STUB: not implemented"; return nil }

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_DialRequest struct {
	DialRequest *DialRequest `protobuf:"bytes,1,opt,name=dialRequest,proto3,oneof"`
}

type Message_DialResponse struct {
	DialResponse *DialResponse `protobuf:"bytes,2,opt,name=dialResponse,proto3,oneof"`
}

type Message_DialDataRequest struct {
	DialDataRequest *DialDataRequest `protobuf:"bytes,3,opt,name=dialDataRequest,proto3,oneof"`
}

type Message_DialDataResponse struct {
	DialDataResponse *DialDataResponse `protobuf:"bytes,4,opt,name=dialDataResponse,proto3,oneof"`
}

func (*Message_DialRequest) isMessage_Msg() { _ = "STUB: not implemented"; return }

func (*Message_DialResponse) isMessage_Msg() { _ = "STUB: not implemented"; return }

func (*Message_DialDataRequest) isMessage_Msg() { _ = "STUB: not implemented"; return }

func (*Message_DialDataResponse) isMessage_Msg() { _ = "STUB: not implemented"; return }

type DialRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addrs         [][]byte               `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	Nonce         uint64                 `protobuf:"fixed64,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *DialRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*DialRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialRequest) GetAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

func (x *DialRequest) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

type DialDataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AddrIdx       uint32                 `protobuf:"varint,1,opt,name=addrIdx,proto3" json:"addrIdx,omitempty"`
	NumBytes      uint64                 `protobuf:"varint,2,opt,name=numBytes,proto3" json:"numBytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialDataRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *DialDataRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*DialDataRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialDataRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialDataRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialDataRequest) GetAddrIdx() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *DialDataRequest) GetNumBytes() uint64 { _ = "STUB: not implemented"; return 0 }

type DialResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Status        DialResponse_ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=autonatv2.pb.DialResponse_ResponseStatus" json:"status,omitempty"`
	AddrIdx       uint32                      `protobuf:"varint,2,opt,name=addrIdx,proto3" json:"addrIdx,omitempty"`
	DialStatus    DialStatus                  `protobuf:"varint,3,opt,name=dialStatus,proto3,enum=autonatv2.pb.DialStatus" json:"dialStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *DialResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*DialResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialResponse) GetStatus() DialResponse_ResponseStatus {
	_ = "STUB: not implemented"
	return *new(DialResponse_ResponseStatus)
}

func (x *DialResponse) GetAddrIdx() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *DialResponse) GetDialStatus() DialStatus {
	_ = "STUB: not implemented"
	return *new(DialStatus)
}

type DialDataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialDataResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *DialDataResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*DialDataResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialDataResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialDataResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialDataResponse) GetData() []byte { _ = "STUB: not implemented"; return nil }

type DialBack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nonce         uint64                 `protobuf:"fixed64,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialBack) Reset() { _ = "STUB: not implemented"; return }

func (x *DialBack) String() string { _ = "STUB: not implemented"; return "" }

func (*DialBack) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialBack) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialBack) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialBack) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

type DialBackResponse struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Status        DialBackResponse_DialBackStatus `protobuf:"varint,1,opt,name=status,proto3,enum=autonatv2.pb.DialBackResponse_DialBackStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialBackResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *DialBackResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*DialBackResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *DialBackResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*DialBackResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *DialBackResponse) GetStatus() DialBackResponse_DialBackStatus {
	_ = "STUB: not implemented"
	return *new(DialBackResponse_DialBackStatus)
}

var File_p2p_protocol_autonatv2_pb_autonatv2_proto protoreflect.FileDescriptor

const file_p2p_protocol_autonatv2_pb_autonatv2_proto_rawDesc = "" +
	"\n" +
	")p2p/protocol/autonatv2/pb/autonatv2.proto\x12\fautonatv2.pb\"\xaa\x02\n" +
	"\aMessage\x12=\n" +
	"\vdialRequest\x18\x01 \x01(\v2\x19.autonatv2.pb.DialRequestH\x00R\vdialRequest\x12@\n" +
	"\fdialResponse\x18\x02 \x01(\v2\x1a.autonatv2.pb.DialResponseH\x00R\fdialResponse\x12I\n" +
	"\x0fdialDataRequest\x18\x03 \x01(\v2\x1d.autonatv2.pb.DialDataRequestH\x00R\x0fdialDataRequest\x12L\n" +
	"\x10dialDataResponse\x18\x04 \x01(\v2\x1e.autonatv2.pb.DialDataResponseH\x00R\x10dialDataResponseB\x05\n" +
	"\x03msg\"9\n" +
	"\vDialRequest\x12\x14\n" +
	"\x05addrs\x18\x01 \x03(\fR\x05addrs\x12\x14\n" +
	"\x05nonce\x18\x02 \x01(\x06R\x05nonce\"G\n" +
	"\x0fDialDataRequest\x12\x18\n" +
	"\aaddrIdx\x18\x01 \x01(\rR\aaddrIdx\x12\x1a\n" +
	"\bnumBytes\x18\x02 \x01(\x04R\bnumBytes\"\x82\x02\n" +
	"\fDialResponse\x12A\n" +
	"\x06status\x18\x01 \x01(\x0e2).autonatv2.pb.DialResponse.ResponseStatusR\x06status\x12\x18\n" +
	"\aaddrIdx\x18\x02 \x01(\rR\aaddrIdx\x128\n" +
	"\n" +
	"dialStatus\x18\x03 \x01(\x0e2\x18.autonatv2.pb.DialStatusR\n" +
	"dialStatus\"[\n" +
	"\x0eResponseStatus\x12\x14\n" +
	"\x10E_INTERNAL_ERROR\x10\x00\x12\x16\n" +
	"\x12E_REQUEST_REJECTED\x10d\x12\x12\n" +
	"\x0eE_DIAL_REFUSED\x10e\x12\a\n" +
	"\x02OK\x10\xc8\x01\"&\n" +
	"\x10DialDataResponse\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\" \n" +
	"\bDialBack\x12\x14\n" +
	"\x05nonce\x18\x01 \x01(\x06R\x05nonce\"s\n" +
	"\x10DialBackResponse\x12E\n" +
	"\x06status\x18\x01 \x01(\x0e2-.autonatv2.pb.DialBackResponse.DialBackStatusR\x06status\"\x18\n" +
	"\x0eDialBackStatus\x12\x06\n" +
	"\x02OK\x10\x00*J\n" +
	"\n" +
	"DialStatus\x12\n" +
	"\n" +
	"\x06UNUSED\x10\x00\x12\x10\n" +
	"\fE_DIAL_ERROR\x10d\x12\x15\n" +
	"\x11E_DIAL_BACK_ERROR\x10e\x12\a\n" +
	"\x02OK\x10\xc8\x01B7Z5github.com/libp2p/go-libp2p/p2p/protocol/autonatv2/pbb\x06proto3"

var (
	file_p2p_protocol_autonatv2_pb_autonatv2_proto_rawDescOnce sync.Once
	file_p2p_protocol_autonatv2_pb_autonatv2_proto_rawDescData []byte
)

func file_p2p_protocol_autonatv2_pb_autonatv2_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_protocol_autonatv2_pb_autonatv2_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_p2p_protocol_autonatv2_pb_autonatv2_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_p2p_protocol_autonatv2_pb_autonatv2_proto_goTypes = []any{
	(DialStatus)(0),
	(DialResponse_ResponseStatus)(0),
	(DialBackResponse_DialBackStatus)(0),
	(*Message)(nil),
	(*DialRequest)(nil),
	(*DialDataRequest)(nil),
	(*DialResponse)(nil),
	(*DialDataResponse)(nil),
	(*DialBack)(nil),
	(*DialBackResponse)(nil),
}
var file_p2p_protocol_autonatv2_pb_autonatv2_proto_depIdxs = []int32{
	4,
	6,
	5,
	7,
	1,
	0,
	2,
	7,
	7,
	7,
	7,
	0,
}

func init()                                                { file_p2p_protocol_autonatv2_pb_autonatv2_proto_init() }
func file_p2p_protocol_autonatv2_pb_autonatv2_proto_init() { _ = "STUB: not implemented"; return }
