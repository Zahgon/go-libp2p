package libp2pwebrtc

import (
	"errors"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

var errWriteAfterClose = errors.New("write after close")

const minMessageSize = 1 << 10

func (s *stream) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) sendBufferSize() int { _ = "STUB: not implemented"; return 0 }

func (s *stream) sendBufferLowThreshold() int { _ = "STUB: not implemented"; return 0 }

func (s *stream) availableSendSpace() int { _ = "STUB: not implemented"; return 0 }

func (s *stream) cancelWrite(errCode network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (s *stream) notifyWriteStateChanged() { _ = "STUB: not implemented"; return }
