package ping

import (
	"context"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("ping")

const (
	PingSize     = 32
	pingTimeout  = 10 * time.Second
	pingDuration = 30 * time.Second

	ID = "/ipfs/ping/1.0.0"

	ServiceName = "libp2p.ping"
)

type PingService struct {
	Host host.Host
}

func NewPingService(h host.Host) *PingService { _ = "STUB: not implemented"; return nil }

func (p *PingService) PingHandler(s network.Stream) { _ = "STUB: not implemented"; return }

type Result struct {
	RTT   time.Duration
	Error error
}

func (ps *PingService) Ping(ctx context.Context, p peer.ID) <-chan Result {
	_ = "STUB: not implemented"
	return nil
}

func pingError(err error) chan Result { _ = "STUB: not implemented"; return nil }

func Ping(ctx context.Context, h host.Host, p peer.ID) <-chan Result {
	_ = "STUB: not implemented"
	return nil
}

func ping(s network.Stream, randReader io.Reader) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
