package mocknet

import (
	"errors"
	"io"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

var streamCounter atomic.Int64

type stream struct {
	rstream *stream
	conn    *conn
	id      int64

	write     *io.PipeWriter
	read      *io.PipeReader
	toDeliver chan *transportObject

	reset  chan struct{}
	close  chan struct{}
	closed chan struct{}

	writeErr error

	protocol atomic.Pointer[protocol.ID]
	stat     network.Stats
}

var ErrClosed = errors.New("stream closed")

type transportObject struct {
	msg         []byte
	arrivalTime time.Time
}

func newStreamPair() (*stream, *stream) { _ = "STUB: not implemented"; return nil, nil }

func newStream(w *io.PipeWriter, r *io.PipeReader, dir network.Direction) *stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) ID() string { _ = "STUB: not implemented"; return "" }

func (s *stream) Protocol() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }

func (s *stream) Stat() network.Stats { _ = "STUB: not implemented"; return *new(network.Stats) }

func (s *stream) SetProtocol(proto protocol.ID) error { _ = "STUB: not implemented"; return nil }

func (s *stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (s *stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *stream) ResetWithError(_ network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) teardown() { _ = "STUB: not implemented"; return }

func (s *stream) Conn() network.Conn { _ = "STUB: not implemented"; return *new(network.Conn) }

func (s *stream) SetDeadline(_ time.Time) error      { _ = "STUB: not implemented"; return nil }
func (s *stream) SetReadDeadline(_ time.Time) error  { _ = "STUB: not implemented"; return nil }
func (s *stream) SetWriteDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) transport() { _ = "STUB: not implemented"; return }

func (s *stream) Scope() network.StreamScope {
	_ = "STUB: not implemented"
	return *new(network.StreamScope)
}

func (s *stream) cancelWrite(err error) { _ = "STUB: not implemented"; return }
