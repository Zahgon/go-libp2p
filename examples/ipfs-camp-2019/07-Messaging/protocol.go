package main

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func chatInputLoop(ctx context.Context, topic *pubsub.Topic, donec chan struct{}) {
	_ = "STUB: not implemented"
	return
}
