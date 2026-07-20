package libp2p

import (
	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/host"
)

type Config = config.Config

type Option = config.Option

func ChainOptions(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(opts ...Option) (host.Host, error) { _ = "STUB: not implemented"; return *new(host.Host), nil }

func NewWithoutDefaults(opts ...Option) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}
