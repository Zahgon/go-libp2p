package pstoremem

import (
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type protoSegment struct {
	sync.RWMutex
	protocols map[peer.ID]map[protocol.ID]struct{}
}

type protoSegments [256]*protoSegment

func (s *protoSegments) get(p peer.ID) *protoSegment { _ = "STUB: not implemented"; return nil }

var errTooManyProtocols = errors.New("too many protocols")

type memoryProtoBook struct {
	segments protoSegments

	maxProtos int
}

var _ pstore.ProtoBook = (*memoryProtoBook)(nil)

type ProtoBookOption func(book *memoryProtoBook) error

func WithMaxProtocols(num int) ProtoBookOption {
	_ = "STUB: not implemented"
	return *new(ProtoBookOption)
}

func NewProtoBook(opts ...ProtoBookOption) (*memoryProtoBook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *memoryProtoBook) SetProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *memoryProtoBook) AddProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *memoryProtoBook) GetProtocols(p peer.ID) ([]protocol.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *memoryProtoBook) RemoveProtocols(p peer.ID, protos ...protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *memoryProtoBook) SupportsProtocols(p peer.ID, protos ...protocol.ID) ([]protocol.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pb *memoryProtoBook) FirstSupportedProtocol(p peer.ID, protos ...protocol.ID) (protocol.ID, error) {
	_ = "STUB: not implemented"
	return *new(protocol.ID), nil
}

func (pb *memoryProtoBook) RemovePeer(p peer.ID) { _ = "STUB: not implemented"; return }
