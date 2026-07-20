package network

import (
	"context"
	"errors"
	"io"
	"net"
	"time"
)

var ErrReset = errors.New("stream reset")

type StreamErrorCode uint32

type StreamError struct {
	ErrorCode      StreamErrorCode
	Remote         bool
	TransportError error
}

func (s *StreamError) Error() string { _ = "STUB: not implemented"; return "" }

func (s *StreamError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (s *StreamError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

const (
	StreamNoError                   StreamErrorCode = 0
	StreamProtocolNegotiationFailed StreamErrorCode = 0x1001
	StreamResourceLimitExceeded     StreamErrorCode = 0x1002
	StreamRateLimited               StreamErrorCode = 0x1003
	StreamProtocolViolation         StreamErrorCode = 0x1004
	StreamSupplanted                StreamErrorCode = 0x1005
	StreamGarbageCollected          StreamErrorCode = 0x1006
	StreamShutdown                  StreamErrorCode = 0x1007
	StreamGated                     StreamErrorCode = 0x1008
	StreamCodeOutOfRange            StreamErrorCode = 0x1009
)

type MuxedStream interface {
	io.Reader
	io.Writer

	io.Closer

	CloseWrite() error

	CloseRead() error

	Reset() error

	ResetWithError(errCode StreamErrorCode) error

	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
}

type MuxedConn interface {
	io.Closer

	CloseWithError(errCode ConnErrorCode) error

	IsClosed() bool

	OpenStream(context.Context) (MuxedStream, error)

	AcceptStream() (MuxedStream, error)

	As(target any) bool
}

type Multiplexer interface {
	NewConn(c net.Conn, isServer bool, scope PeerScope) (MuxedConn, error)
}
