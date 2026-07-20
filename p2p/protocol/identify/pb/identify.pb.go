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

type Identify struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ProtocolVersion *string `protobuf:"bytes,5,opt,name=protocolVersion" json:"protocolVersion,omitempty"`

	AgentVersion *string `protobuf:"bytes,6,opt,name=agentVersion" json:"agentVersion,omitempty"`

	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`

	ListenAddrs [][]byte `protobuf:"bytes,2,rep,name=listenAddrs" json:"listenAddrs,omitempty"`

	ObservedAddr []byte `protobuf:"bytes,4,opt,name=observedAddr" json:"observedAddr,omitempty"`

	Protocols []string `protobuf:"bytes,3,rep,name=protocols" json:"protocols,omitempty"`

	SignedPeerRecord []byte `protobuf:"bytes,8,opt,name=signedPeerRecord" json:"signedPeerRecord,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Identify) Reset() { _ = "STUB: not implemented"; return }

func (x *Identify) String() string { _ = "STUB: not implemented"; return "" }

func (*Identify) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Identify) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Identify) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Identify) GetProtocolVersion() string { _ = "STUB: not implemented"; return "" }

func (x *Identify) GetAgentVersion() string { _ = "STUB: not implemented"; return "" }

func (x *Identify) GetPublicKey() []byte { _ = "STUB: not implemented"; return nil }

func (x *Identify) GetListenAddrs() [][]byte { _ = "STUB: not implemented"; return nil }

func (x *Identify) GetObservedAddr() []byte { _ = "STUB: not implemented"; return nil }

func (x *Identify) GetProtocols() []string { _ = "STUB: not implemented"; return nil }

func (x *Identify) GetSignedPeerRecord() []byte { _ = "STUB: not implemented"; return nil }

var File_p2p_protocol_identify_pb_identify_proto protoreflect.FileDescriptor

const file_p2p_protocol_identify_pb_identify_proto_rawDesc = "" +
	"\n" +
	"'p2p/protocol/identify/pb/identify.proto\x12\videntify.pb\"\x86\x02\n" +
	"\bIdentify\x12(\n" +
	"\x0fprotocolVersion\x18\x05 \x01(\tR\x0fprotocolVersion\x12\"\n" +
	"\fagentVersion\x18\x06 \x01(\tR\fagentVersion\x12\x1c\n" +
	"\tpublicKey\x18\x01 \x01(\fR\tpublicKey\x12 \n" +
	"\vlistenAddrs\x18\x02 \x03(\fR\vlistenAddrs\x12\"\n" +
	"\fobservedAddr\x18\x04 \x01(\fR\fobservedAddr\x12\x1c\n" +
	"\tprotocols\x18\x03 \x03(\tR\tprotocols\x12*\n" +
	"\x10signedPeerRecord\x18\b \x01(\fR\x10signedPeerRecordB6Z4github.com/libp2p/go-libp2p/p2p/protocol/identify/pb"

var (
	file_p2p_protocol_identify_pb_identify_proto_rawDescOnce sync.Once
	file_p2p_protocol_identify_pb_identify_proto_rawDescData []byte
)

func file_p2p_protocol_identify_pb_identify_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_p2p_protocol_identify_pb_identify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_p2p_protocol_identify_pb_identify_proto_goTypes = []any{
	(*Identify)(nil),
}
var file_p2p_protocol_identify_pb_identify_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}

func init()                                              { file_p2p_protocol_identify_pb_identify_proto_init() }
func file_p2p_protocol_identify_pb_identify_proto_init() { _ = "STUB: not implemented"; return }
