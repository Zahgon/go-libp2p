package swarm

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

var _ network.Stream = &Stream{}

type Stream struct {
	id uint64

	stream network.MuxedStream
	conn   *Conn
	scope  network.StreamManagementScope

	closeMx  sync.Mutex
	isClosed bool

	acceptStreamGoroutineCompleted bool

	protocol atomic.Pointer[protocol.ID]

	stat network.Stats
}

func (s *Stream) ID() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) String() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) Conn() network.Conn { _ = "STUB: not implemented"; return *new(network.Conn) }

func (s *Stream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) ResetWithError(errCode network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) closeAndRemoveStream() { _ = "STUB: not implemented"; return }

func (s *Stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) completeAcceptStreamGoroutine() { _ = "STUB: not implemented"; return }

func (s *Stream) Protocol() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }

func (s *Stream) SetProtocol(p protocol.ID) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) Stat() network.Stats { _ = "STUB: not implemented"; return *new(network.Stats) }

func (s *Stream) Scope() network.StreamScope {
	_ = "STUB: not implemented"
	return *new(network.StreamScope)
}
