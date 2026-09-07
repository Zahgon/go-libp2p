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

type AddrBookRecord struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`

	Addrs []*AddrBookRecord_AddrEntry `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`

	CertifiedRecord *AddrBookRecord_CertifiedRecord `protobuf:"bytes,3,opt,name=certified_record,json=certifiedRecord,proto3" json:"certified_record,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AddrBookRecord) Reset() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord) String() string { _ = "STUB: not implemented"; return "" }

func (*AddrBookRecord) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*AddrBookRecord) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *AddrBookRecord) GetId() []byte { _ = "STUB: not implemented"; return nil }

func (x *AddrBookRecord) GetAddrs() []*AddrBookRecord_AddrEntry {
	_ = "STUB: not implemented"
	return nil
}

func (x *AddrBookRecord) GetCertifiedRecord() *AddrBookRecord_CertifiedRecord {
	_ = "STUB: not implemented"
	return nil
}

type AddrBookRecord_AddrEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Addr  []byte                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`

	Expiry int64 `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"`

	Ttl           int64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddrBookRecord_AddrEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord_AddrEntry) String() string { _ = "STUB: not implemented"; return "" }

func (*AddrBookRecord_AddrEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord_AddrEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*AddrBookRecord_AddrEntry) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *AddrBookRecord_AddrEntry) GetAddr() []byte { _ = "STUB: not implemented"; return nil }

func (x *AddrBookRecord_AddrEntry) GetExpiry() int64 { _ = "STUB: not implemented"; return 0 }

func (x *AddrBookRecord_AddrEntry) GetTtl() int64 { _ = "STUB: not implemented"; return 0 }

type AddrBookRecord_CertifiedRecord struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Seq uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`

	Raw           []byte `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddrBookRecord_CertifiedRecord) Reset() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord_CertifiedRecord) String() string { _ = "STUB: not implemented"; return "" }

func (*AddrBookRecord_CertifiedRecord) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *AddrBookRecord_CertifiedRecord) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*AddrBookRecord_CertifiedRecord) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *AddrBookRecord_CertifiedRecord) GetSeq() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *AddrBookRecord_CertifiedRecord) GetRaw() []byte { _ = "STUB: not implemented"; return nil }

var File_p2p_host_peerstore_pstoreds_pb_pstore_proto protoreflect.FileDescriptor

const file_p2p_host_peerstore_pstoreds_pb_pstore_proto_rawDesc = "" +
	"\n" +
	"+p2p/host/peerstore/pstoreds/pb/pstore.proto\x12\tpstore.pb\"\xb3\x02\n" +
	"\x0eAddrBookRecord\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\fR\x02id\x129\n" +
	"\x05addrs\x18\x02 \x03(\v2#.pstore.pb.AddrBookRecord.AddrEntryR\x05addrs\x12T\n" +
	"\x10certified_record\x18\x03 \x01(\v2).pstore.pb.AddrBookRecord.CertifiedRecordR\x0fcertifiedRecord\x1aI\n" +
	"\tAddrEntry\x12\x12\n" +
	"\x04addr\x18\x01 \x01(\fR\x04addr\x12\x16\n" +
	"\x06expiry\x18\x02 \x01(\x03R\x06expiry\x12\x10\n" +
	"\x03ttl\x18\x03 \x01(\x03R\x03ttl\x1a5\n" +
	"\x0fCertifiedRecord\x12\x10\n" +
	"\x03seq\x18\x01 \x01(\x04R\x03seq\x12\x10\n" +
	"\x03raw\x18\x02 \x01(\fR\x03rawB<Z:github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoreds/pbb\x06proto3"

var (
	file_p2p_host_peerstore_pstoreds_pb_pstore_proto_rawDescOnce sync.Once
	file_p2p_host_peerstore_pstoreds_pb_pstore_proto_rawDescData []byte
)

func file_p2p_host_peerstore_pstoreds_pb_pstore_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_host_peerstore_pstoreds_pb_pstore_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_p2p_host_peerstore_pstoreds_pb_pstore_proto_goTypes = []any{
	(*AddrBookRecord)(nil),
	(*AddrBookRecord_AddrEntry)(nil),
	(*AddrBookRecord_CertifiedRecord)(nil),
}
var file_p2p_host_peerstore_pstoreds_pb_pstore_proto_depIdxs = []int32{
	1,
	2,
	2,
	2,
	2,
	2,
	0,
}

func init()                                                  { file_p2p_host_peerstore_pstoreds_pb_pstore_proto_init() }
func file_p2p_host_peerstore_pstoreds_pb_pstore_proto_init() { _ = "STUB: not implemented"; return }
