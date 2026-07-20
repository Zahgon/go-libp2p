package yamux

import (
	"io"
	"math"
	"net"

	"github.com/libp2p/go-libp2p/core/network"

	"github.com/libp2p/go-yamux/v5"
)

var DefaultTransport *Transport

const ID = "/yamux/1.0.0"

func init() {
	config := yamux.DefaultConfig()

	config.MaxStreamWindowSize = uint32(16 * 1024 * 1024)

	config.LogOutput = io.Discard

	config.ReadBufSize = 0

	config.MaxIncomingStreams = math.MaxUint32
	DefaultTransport = (*Transport)(config)
}

type Transport yamux.Config

var _ network.Multiplexer = &Transport{}

func (t *Transport) NewConn(nc net.Conn, isServer bool, scope network.PeerScope) (network.MuxedConn, error) {
	_ = "STUB: not implemented"
	return *new(network.MuxedConn), nil
}

func (t *Transport) Config() *yamux.Config { _ = "STUB: not implemented"; return nil }
