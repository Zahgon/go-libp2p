package noise

import (
	"golang.org/x/crypto/chacha20poly1305"
)

const MaxTransportMsgLength = 0xffff

const MaxPlaintextLength = MaxTransportMsgLength - chacha20poly1305.Overhead

const LengthPrefixLength = 2

func (s *secureSession) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *secureSession) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *secureSession) readNextInsecureMsgLen() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *secureSession) readNextMsgInsecure(buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *secureSession) writeMsgInsecure(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
