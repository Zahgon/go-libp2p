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

type NoiseExtensions struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	WebtransportCerthashes [][]byte               `protobuf:"bytes,1,rep,name=webtransport_certhashes,json=webtransportCerthashes" json:"webtransport_certhashes,omitempty"`
	StreamMuxers           []string               `protobuf:"bytes,2,rep,name=stream_muxers,json=streamMuxers" json:"stream_muxers,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *NoiseExtensions) Reset() { _ = "STUB: not implemented"; return }

func (x *NoiseExtensions) String() string { _ = "STUB: not implemented"; return "" }

func (*NoiseExtensions) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *NoiseExtensions) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*NoiseExtensions) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *NoiseExtensions) GetWebtransportCerthashes() [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func (x *NoiseExtensions) GetStreamMuxers() []string { _ = "STUB: not implemented"; return nil }

type NoiseHandshakePayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdentityKey   []byte                 `protobuf:"bytes,1,opt,name=identity_key,json=identityKey" json:"identity_key,omitempty"`
	IdentitySig   []byte                 `protobuf:"bytes,2,opt,name=identity_sig,json=identitySig" json:"identity_sig,omitempty"`
	Extensions    *NoiseExtensions       `protobuf:"bytes,4,opt,name=extensions" json:"extensions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoiseHandshakePayload) Reset() { _ = "STUB: not implemented"; return }

func (x *NoiseHandshakePayload) String() string { _ = "STUB: not implemented"; return "" }

func (*NoiseHandshakePayload) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *NoiseHandshakePayload) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*NoiseHandshakePayload) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *NoiseHandshakePayload) GetIdentityKey() []byte { _ = "STUB: not implemented"; return nil }

func (x *NoiseHandshakePayload) GetIdentitySig() []byte { _ = "STUB: not implemented"; return nil }

func (x *NoiseHandshakePayload) GetExtensions() *NoiseExtensions {
	_ = "STUB: not implemented"
	return nil
}

var File_p2p_security_noise_pb_payload_proto protoreflect.FileDescriptor

const file_p2p_security_noise_pb_payload_proto_rawDesc = "" +
	"\n" +
	"#p2p/security/noise/pb/payload.proto\x12\x02pb\"o\n" +
	"\x0fNoiseExtensions\x127\n" +
	"\x17webtransport_certhashes\x18\x01 \x03(\fR\x16webtransportCerthashes\x12#\n" +
	"\rstream_muxers\x18\x02 \x03(\tR\fstreamMuxers\"\x92\x01\n" +
	"\x15NoiseHandshakePayload\x12!\n" +
	"\fidentity_key\x18\x01 \x01(\fR\videntityKey\x12!\n" +
	"\fidentity_sig\x18\x02 \x01(\fR\videntitySig\x123\n" +
	"\n" +
	"extensions\x18\x04 \x01(\v2\x13.pb.NoiseExtensionsR\n" +
	"extensionsB3Z1github.com/libp2p/go-libp2p/p2p/security/noise/pb"

var (
	file_p2p_security_noise_pb_payload_proto_rawDescOnce sync.Once
	file_p2p_security_noise_pb_payload_proto_rawDescData []byte
)

func file_p2p_security_noise_pb_payload_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_security_noise_pb_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_p2p_security_noise_pb_payload_proto_goTypes = []any{
	(*NoiseExtensions)(nil),
	(*NoiseHandshakePayload)(nil),
}
var file_p2p_security_noise_pb_payload_proto_depIdxs = []int32{
	0,
	1,
	1,
	1,
	1,
	0,
}

func init()                                          { file_p2p_security_noise_pb_payload_proto_init() }
func file_p2p_security_noise_pb_payload_proto_init() { _ = "STUB: not implemented"; return }
