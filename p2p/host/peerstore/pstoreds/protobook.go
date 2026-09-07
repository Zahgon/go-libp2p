package pstoreds

import (
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type protoSegment struct {
	sync.RWMutex
}

type protoSegments [256]*protoSegment

func (s *protoSegments) get(p peer.ID) *protoSegment { _ = "STUB: not implemented"; return nil }

var errTooManyProtocols = errors.New("too many protocols")

type ProtoBookOption func(*dsProtoBook) error

func WithMaxProtocols(num int) ProtoBookOption {
	_ = "STUB: not implemented"
	return *new(ProtoBookOption)
}

type dsProtoBook struct {
	segments  protoSegments
	meta      pstore.PeerMetadata
	maxProtos int
}

var _ pstore.ProtoBook = (*dsProtoBook)(nil)

func NewProtoBook(meta pstore.PeerMetadata, opts ...ProtoBookOption) (*dsProtoBook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *dsProtoBook) SetProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *dsProtoBook) AddProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *dsProtoBook) GetProtocols(p peer.ID) ([]protocol.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *dsProtoBook) SupportsProtocols(p peer.ID, protos ...protocol.ID) ([]protocol.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *dsProtoBook) FirstSupportedProtocol(p peer.ID, protos ...protocol.ID) (protocol.ID, error) {
	_ = "STUB: not implemented"
	return *new(protocol.ID), nil
}

func (pb *dsProtoBook) RemoveProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *dsProtoBook) getProtocolMap(p peer.ID) (map[protocol.ID]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *dsProtoBook) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
