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

type KeyType int32

const (
	KeyType_RSA       KeyType = 0
	KeyType_Ed25519   KeyType = 1
	KeyType_Secp256k1 KeyType = 2
	KeyType_ECDSA     KeyType = 3
)

var (
	KeyType_name = map[int32]string{
		0: "RSA",
		1: "Ed25519",
		2: "Secp256k1",
		3: "ECDSA",
	}
	KeyType_value = map[string]int32{
		"RSA":       0,
		"Ed25519":   1,
		"Secp256k1": 2,
		"ECDSA":     3,
	}
)

func (x KeyType) Enum() *KeyType { _ = "STUB: not implemented"; return nil }

func (x KeyType) String() string { _ = "STUB: not implemented"; return "" }

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (KeyType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x KeyType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (x *KeyType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (KeyType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type PublicKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *KeyType               `protobuf:"varint,1,req,name=Type,enum=crypto.pb.KeyType" json:"Type,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,req,name=Data" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicKey) Reset() { _ = "STUB: not implemented"; return }

func (x *PublicKey) String() string { _ = "STUB: not implemented"; return "" }

func (*PublicKey) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PublicKey) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *PublicKey) GetType() KeyType { _ = "STUB: not implemented"; return *new(KeyType) }

func (x *PublicKey) GetData() []byte { _ = "STUB: not implemented"; return nil }

type PrivateKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *KeyType               `protobuf:"varint,1,req,name=Type,enum=crypto.pb.KeyType" json:"Type,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,req,name=Data" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateKey) Reset() { _ = "STUB: not implemented"; return }

func (x *PrivateKey) String() string { _ = "STUB: not implemented"; return "" }

func (*PrivateKey) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PrivateKey) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *PrivateKey) GetType() KeyType { _ = "STUB: not implemented"; return *new(KeyType) }

func (x *PrivateKey) GetData() []byte { _ = "STUB: not implemented"; return nil }

var File_core_crypto_pb_crypto_proto protoreflect.FileDescriptor

const file_core_crypto_pb_crypto_proto_rawDesc = "" +
	"\n" +
	"\x1bcore/crypto/pb/crypto.proto\x12\tcrypto.pb\"G\n" +
	"\tPublicKey\x12&\n" +
	"\x04Type\x18\x01 \x02(\x0e2\x12.crypto.pb.KeyTypeR\x04Type\x12\x12\n" +
	"\x04Data\x18\x02 \x02(\fR\x04Data\"H\n" +
	"\n" +
	"PrivateKey\x12&\n" +
	"\x04Type\x18\x01 \x02(\x0e2\x12.crypto.pb.KeyTypeR\x04Type\x12\x12\n" +
	"\x04Data\x18\x02 \x02(\fR\x04Data*9\n" +
	"\aKeyType\x12\a\n" +
	"\x03RSA\x10\x00\x12\v\n" +
	"\aEd25519\x10\x01\x12\r\n" +
	"\tSecp256k1\x10\x02\x12\t\n" +
	"\x05ECDSA\x10\x03B,Z*github.com/libp2p/go-libp2p/core/crypto/pb"

var (
	file_core_crypto_pb_crypto_proto_rawDescOnce sync.Once
	file_core_crypto_pb_crypto_proto_rawDescData []byte
)

func file_core_crypto_pb_crypto_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_core_crypto_pb_crypto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_crypto_pb_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_crypto_pb_crypto_proto_goTypes = []any{
	(KeyType)(0),
	(*PublicKey)(nil),
	(*PrivateKey)(nil),
}
var file_core_crypto_pb_crypto_proto_depIdxs = []int32{
	0,
	0,
	2,
	2,
	2,
	2,
	0,
}

func init()                                  { file_core_crypto_pb_crypto_proto_init() }
func file_core_crypto_pb_crypto_proto_init() { _ = "STUB: not implemented"; return }
