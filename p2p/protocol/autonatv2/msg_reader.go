package autonatv2

import (
	"io"
)

type msgReader struct {
	R   io.Reader
	Buf []byte
}

func (m *msgReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *msgReader) ReadMsg() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
