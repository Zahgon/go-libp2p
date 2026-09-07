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

type PeerRecord struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`

	Seq uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`

	Addresses     []*PeerRecord_AddressInfo `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PeerRecord) Reset() { _ = "STUB: not implemented"; return }

func (x *PeerRecord) String() string { _ = "STUB: not implemented"; return "" }

func (*PeerRecord) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PeerRecord) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PeerRecord) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *PeerRecord) GetPeerId() []byte { _ = "STUB: not implemented"; return nil }

func (x *PeerRecord) GetSeq() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *PeerRecord) GetAddresses() []*PeerRecord_AddressInfo {
	_ = "STUB: not implemented"
	return nil
}

type PeerRecord_AddressInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Multiaddr     []byte                 `protobuf:"bytes,1,opt,name=multiaddr,proto3" json:"multiaddr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PeerRecord_AddressInfo) Reset() { _ = "STUB: not implemented"; return }

func (x *PeerRecord_AddressInfo) String() string { _ = "STUB: not implemented"; return "" }

func (*PeerRecord_AddressInfo) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PeerRecord_AddressInfo) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PeerRecord_AddressInfo) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *PeerRecord_AddressInfo) GetMultiaddr() []byte { _ = "STUB: not implemented"; return nil }

var File_core_peer_pb_peer_record_proto protoreflect.FileDescriptor

const file_core_peer_pb_peer_record_proto_rawDesc = "" +
	"\n" +
	"\x1ecore/peer/pb/peer_record.proto\x12\apeer.pb\"\xa3\x01\n" +
	"\n" +
	"PeerRecord\x12\x17\n" +
	"\apeer_id\x18\x01 \x01(\fR\x06peerId\x12\x10\n" +
	"\x03seq\x18\x02 \x01(\x04R\x03seq\x12=\n" +
	"\taddresses\x18\x03 \x03(\v2\x1f.peer.pb.PeerRecord.AddressInfoR\taddresses\x1a+\n" +
	"\vAddressInfo\x12\x1c\n" +
	"\tmultiaddr\x18\x01 \x01(\fR\tmultiaddrB*Z(github.com/libp2p/go-libp2p/core/peer/pbb\x06proto3"

var (
	file_core_peer_pb_peer_record_proto_rawDescOnce sync.Once
	file_core_peer_pb_peer_record_proto_rawDescData []byte
)

func file_core_peer_pb_peer_record_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_core_peer_pb_peer_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_peer_pb_peer_record_proto_goTypes = []any{
	(*PeerRecord)(nil),
	(*PeerRecord_AddressInfo)(nil),
}
var file_core_peer_pb_peer_record_proto_depIdxs = []int32{
	1,
	1,
	1,
	1,
	1,
	0,
}

func init()                                     { file_core_peer_pb_peer_record_proto_init() }
func file_core_peer_pb_peer_record_proto_init() { _ = "STUB: not implemented"; return }
