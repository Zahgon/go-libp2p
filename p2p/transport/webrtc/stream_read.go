package libp2pwebrtc

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

func (s *stream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) setDataChannelReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s *stream) closeRead(errCode network.StreamErrorCode, remote bool) error {
	_ = "STUB: not implemented"
	return nil
}
