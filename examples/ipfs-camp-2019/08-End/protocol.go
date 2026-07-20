package main

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

func sendMessage(ctx context.Context, topic *pubsub.Topic, msg string) {
	_ = "STUB: not implemented"
	return
}

func updatePeer(ctx context.Context, topic *pubsub.Topic, id peer.ID, handle string) {
	_ = "STUB: not implemented"
	return
}

func chatInputLoop(ctx context.Context, h host.Host, topic *pubsub.Topic, donec chan struct{}) {
	_ = "STUB: not implemented"
	return
}
