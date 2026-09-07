package quicreuse

import (
	"context"
	"net"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/quic-go/quic-go"
)

type Option func(*ConnManager) error

type listenUDP func(network string, laddr *net.UDPAddr) (net.PacketConn, error)

func OverrideListenUDP(f listenUDP) Option { _ = "STUB: not implemented"; return *new(Option) }

func OverrideSourceIPSelector(f func() (SourceIPSelector, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithQlogTracerDir(dir string) Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableReuseport() Option { _ = "STUB: not implemented"; return *new(Option) }

func ConnContext(f func(ctx context.Context, clientInfo *quic.ClientInfo) (context.Context, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func VerifySourceAddress(f func(addr net.Addr) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func EnableMetrics(reg prometheus.Registerer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
