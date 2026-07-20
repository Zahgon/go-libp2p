package main

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
)

const chatProtocol = "/libp2p/chat/1.0.0"

func chatHandler(s network.Stream) { _ = "STUB: not implemented"; return }

func chatSend(msg string, s network.Stream) error { _ = "STUB: not implemented"; return nil }

func chatInputLoop(ctx context.Context, h host.Host, donec chan struct{}) {
	_ = "STUB: not implemented"
	return
}
