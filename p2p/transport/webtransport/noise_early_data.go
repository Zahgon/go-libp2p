package libp2pwebtransport

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-libp2p/p2p/security/noise/pb"
)

type earlyDataHandler struct {
	earlyData *pb.NoiseExtensions
	receive   func(extensions *pb.NoiseExtensions) error
}

var _ noise.EarlyDataHandler = &earlyDataHandler{}

func newEarlyDataSender(earlyData *pb.NoiseExtensions) noise.EarlyDataHandler {
	_ = "STUB: not implemented"
	return *new(noise.EarlyDataHandler)
}

func newEarlyDataReceiver(receive func(*pb.NoiseExtensions) error) noise.EarlyDataHandler {
	_ = "STUB: not implemented"
	return *new(noise.EarlyDataHandler)
}

func (e *earlyDataHandler) Send(context.Context, net.Conn, peer.ID) *pb.NoiseExtensions {
	_ = "STUB: not implemented"
	return nil
}

func (e *earlyDataHandler) Received(_ context.Context, _ net.Conn, ext *pb.NoiseExtensions) error {
	_ = "STUB: not implemented"
	return nil
}
