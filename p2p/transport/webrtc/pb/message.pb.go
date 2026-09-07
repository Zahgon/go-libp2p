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

type Message_Flag int32

const (
	Message_FIN Message_Flag = 0

	Message_STOP_SENDING Message_Flag = 1

	Message_RESET Message_Flag = 2

	Message_FIN_ACK Message_Flag = 3
)

var (
	Message_Flag_name = map[int32]string{
		0: "FIN",
		1: "STOP_SENDING",
		2: "RESET",
		3: "FIN_ACK",
	}
	Message_Flag_value = map[string]int32{
		"FIN":          0,
		"STOP_SENDING": 1,
		"RESET":        2,
		"FIN_ACK":      3,
	}
)

func (x Message_Flag) Enum() *Message_Flag { _ = "STUB: not implemented"; return nil }

func (x Message_Flag) String() string { _ = "STUB: not implemented"; return "" }

func (Message_Flag) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (Message_Flag) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x Message_Flag) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (x *Message_Flag) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (Message_Flag) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flag          *Message_Flag          `protobuf:"varint,1,opt,name=flag,enum=Message_Flag" json:"flag,omitempty"`
	Message       []byte                 `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ErrorCode     *uint32                `protobuf:"varint,3,opt,name=errorCode" json:"errorCode,omitempty"`
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

func (x *Message) GetFlag() Message_Flag { _ = "STUB: not implemented"; return *new(Message_Flag) }

func (x *Message) GetMessage() []byte { _ = "STUB: not implemented"; return nil }

func (x *Message) GetErrorCode() uint32 { _ = "STUB: not implemented"; return 0 }

var File_p2p_transport_webrtc_pb_message_proto protoreflect.FileDescriptor

const file_p2p_transport_webrtc_pb_message_proto_rawDesc = "" +
	"\n" +
	"%p2p/transport/webrtc/pb/message.proto\"\x9f\x01\n" +
	"\aMessage\x12!\n" +
	"\x04flag\x18\x01 \x01(\x0e2\r.Message.FlagR\x04flag\x12\x18\n" +
	"\amessage\x18\x02 \x01(\fR\amessage\x12\x1c\n" +
	"\terrorCode\x18\x03 \x01(\rR\terrorCode\"9\n" +
	"\x04Flag\x12\a\n" +
	"\x03FIN\x10\x00\x12\x10\n" +
	"\fSTOP_SENDING\x10\x01\x12\t\n" +
	"\x05RESET\x10\x02\x12\v\n" +
	"\aFIN_ACK\x10\x03B5Z3github.com/libp2p/go-libp2p/p2p/transport/webrtc/pb"

var (
	file_p2p_transport_webrtc_pb_message_proto_rawDescOnce sync.Once
	file_p2p_transport_webrtc_pb_message_proto_rawDescData []byte
)

func file_p2p_transport_webrtc_pb_message_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_transport_webrtc_pb_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_p2p_transport_webrtc_pb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_p2p_transport_webrtc_pb_message_proto_goTypes = []any{
	(Message_Flag)(0),
	(*Message)(nil),
}
var file_p2p_transport_webrtc_pb_message_proto_depIdxs = []int32{
	0,
	1,
	1,
	1,
	1,
	0,
}

func init()                                            { file_p2p_transport_webrtc_pb_message_proto_init() }
func file_p2p_transport_webrtc_pb_message_proto_init() { _ = "STUB: not implemented"; return }
