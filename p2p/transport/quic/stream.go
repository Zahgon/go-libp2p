package libp2pquic

import (
	"github.com/libp2p/go-libp2p/core/network"

	"github.com/quic-go/quic-go"
)

const (
	reset quic.StreamErrorCode = 0
)

type stream struct {
	*quic.Stream
}

var _ network.MuxedStream = stream{}

func parseStreamError(err error) error { _ = "STUB: not implemented"; return nil }

func (s stream) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s stream) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s stream) ResetWithError(errCode network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }
