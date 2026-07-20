package pb

import (
	sync "sync"

	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Envelope struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	PublicKey *pb.PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`

	PayloadType []byte `protobuf:"bytes,2,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`

	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`

	Signature     []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() { _ = "STUB: not implemented"; return }

func (x *Envelope) String() string { _ = "STUB: not implemented"; return "" }

func (*Envelope) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Envelope) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Envelope) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Envelope) GetPublicKey() *pb.PublicKey { _ = "STUB: not implemented"; return nil }

func (x *Envelope) GetPayloadType() []byte { _ = "STUB: not implemented"; return nil }

func (x *Envelope) GetPayload() []byte { _ = "STUB: not implemented"; return nil }

func (x *Envelope) GetSignature() []byte { _ = "STUB: not implemented"; return nil }

var File_core_record_pb_envelope_proto protoreflect.FileDescriptor

const file_core_record_pb_envelope_proto_rawDesc = "" +
	"\n" +
	"\x1dcore/record/pb/envelope.proto\x12\trecord.pb\x1a\x1bcore/crypto/pb/crypto.proto\"\x9a\x01\n" +
	"\bEnvelope\x123\n" +
	"\n" +
	"public_key\x18\x01 \x01(\v2\x14.crypto.pb.PublicKeyR\tpublicKey\x12!\n" +
	"\fpayload_type\x18\x02 \x01(\fR\vpayloadType\x12\x18\n" +
	"\apayload\x18\x03 \x01(\fR\apayload\x12\x1c\n" +
	"\tsignature\x18\x05 \x01(\fR\tsignatureB,Z*github.com/libp2p/go-libp2p/core/record/pbb\x06proto3"

var (
	file_core_record_pb_envelope_proto_rawDescOnce sync.Once
	file_core_record_pb_envelope_proto_rawDescData []byte
)

func file_core_record_pb_envelope_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_core_record_pb_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_record_pb_envelope_proto_goTypes = []any{
	(*Envelope)(nil),
	(*pb.PublicKey)(nil),
}
var file_core_record_pb_envelope_proto_depIdxs = []int32{
	1,
	1,
	1,
	1,
	1,
	0,
}

func init()                                    { file_core_record_pb_envelope_proto_init() }
func file_core_record_pb_envelope_proto_init() { _ = "STUB: not implemented"; return }
