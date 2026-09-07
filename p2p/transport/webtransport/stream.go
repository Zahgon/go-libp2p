package libp2pwebtransport

import (
	"net"

	"github.com/libp2p/go-libp2p/core/network"

	"github.com/quic-go/webtransport-go"
)

const (
	reset webtransport.StreamErrorCode = 0
)

type webtransportStream struct {
	*webtransport.Stream
	wsess *webtransport.Session
}

var _ net.Conn = webtransportStream{}

func (s webtransportStream) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s webtransportStream) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type stream struct {
	*webtransport.Stream
}

var _ network.MuxedStream = stream{}

func (s stream) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s stream) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s stream) ResetWithError(_ network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s stream) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (s stream) CloseWrite() error { _ = "STUB: not implemented"; return nil }
