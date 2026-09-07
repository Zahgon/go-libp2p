package util

import (
	"io"

	"github.com/libp2p/go-msgio/pbio"
	"google.golang.org/protobuf/proto"
)

type DelimitedReader struct {
	r   io.Reader
	buf []byte
}

func NewDelimitedReader(r io.Reader, maxSize int) *DelimitedReader {
	_ = "STUB: not implemented"
	return nil
}

func (d *DelimitedReader) Close() { _ = "STUB: not implemented"; return }

func (d *DelimitedReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *DelimitedReader) ReadMsg(msg proto.Message) error { _ = "STUB: not implemented"; return nil }

func NewDelimitedWriter(w io.Writer) pbio.WriteCloser {
	_ = "STUB: not implemented"
	return *new(pbio.WriteCloser)
}
