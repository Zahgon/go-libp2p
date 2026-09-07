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

type Exchange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            []byte                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Pubkey        *pb.PublicKey          `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Exchange) Reset() { _ = "STUB: not implemented"; return }

func (x *Exchange) String() string { _ = "STUB: not implemented"; return "" }

func (*Exchange) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Exchange) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Exchange) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Exchange) GetId() []byte { _ = "STUB: not implemented"; return nil }

func (x *Exchange) GetPubkey() *pb.PublicKey { _ = "STUB: not implemented"; return nil }

var File_p2p_security_insecure_pb_plaintext_proto protoreflect.FileDescriptor

const file_p2p_security_insecure_pb_plaintext_proto_rawDesc = "" +
	"\n" +
	"(p2p/security/insecure/pb/plaintext.proto\x12\fplaintext.pb\x1a\x1bcore/crypto/pb/crypto.proto\"H\n" +
	"\bExchange\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\fR\x02id\x12,\n" +
	"\x06pubkey\x18\x02 \x01(\v2\x14.crypto.pb.PublicKeyR\x06pubkeyB6Z4github.com/libp2p/go-libp2p/p2p/security/insecure/pb"

var (
	file_p2p_security_insecure_pb_plaintext_proto_rawDescOnce sync.Once
	file_p2p_security_insecure_pb_plaintext_proto_rawDescData []byte
)

func file_p2p_security_insecure_pb_plaintext_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_security_insecure_pb_plaintext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_p2p_security_insecure_pb_plaintext_proto_goTypes = []any{
	(*Exchange)(nil),
	(*pb.PublicKey)(nil),
}
var file_p2p_security_insecure_pb_plaintext_proto_depIdxs = []int32{
	1,
	1,
	1,
	1,
	1,
	0,
}

func init()                                               { file_p2p_security_insecure_pb_plaintext_proto_init() }
func file_p2p_security_insecure_pb_plaintext_proto_init() { _ = "STUB: not implemented"; return }
