package yamux

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"

	"github.com/libp2p/go-yamux/v5"
)

type stream yamux.Stream

var _ network.MuxedStream = &stream{}

func parseError(err error) error { _ = "STUB: not implemented"; return nil }

func (s *stream) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *stream) ResetWithError(errCode network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s *stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (s *stream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) yamux() *yamux.Stream { _ = "STUB: not implemented"; return nil }
