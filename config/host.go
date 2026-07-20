package config

import (
	basichost "github.com/libp2p/go-libp2p/p2p/host/basic"
	routed "github.com/libp2p/go-libp2p/p2p/host/routed"

	"go.uber.org/fx"
)

type closableBasicHost struct {
	*fx.App
	*basichost.BasicHost
}

func (h *closableBasicHost) Close() error { _ = "STUB: not implemented"; return nil }

type closableRoutedHost struct {
	closableBasicHost
	*routed.RoutedHost
}

func (h *closableRoutedHost) Close() error { _ = "STUB: not implemented"; return nil }
